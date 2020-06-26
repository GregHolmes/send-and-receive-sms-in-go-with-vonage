package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/joho/godotenv"
	"github.com/nexmo-community/nexmo-go"
)

var decoder = schema.NewDecoder()

type SmsObject struct {
	Msisdn           string `schema:"msisdn"`
	To               string `schema:"to"`
	MessageId        string `schema:"messageId"`
	Text             string `schema:"text"`
	Type             string `schema:"type"`
	Keyword          string `schema:"keyword"`
	ApiKey           string `schema:"api-key"`
	MessageTimestamp string `schema:"message-timestamp"`
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func receiveSms(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/webhook/inbound-sms" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		var smsObject SmsObject

		err := decoder.Decode(&smsObject, r.URL.Query())
		if err != nil {
			log.Println("Error in GET parameters : ", err)
		} else {
			fmt.Println("Received an SMS from: " + smsObject.Msisdn + " with the body: " + smsObject.Text)
		}
	case "POST":
		var smsObject SmsObject

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(reqBody, &smsObject)
		fmt.Println("Received an SMS from: " + smsObject.Msisdn + " with the body: " + smsObject.Text)
	default:
		fmt.Println(w, "Sorry, only GET and POST methods are supported.")
	}
}

func sendSms(w http.ResponseWriter, r *http.Request) {

	// The commented out code below is the previous functionality from our public branch
	// auth := nexmo.NewAuthSet()
	// auth.SetAPISecret(goDotEnvVariable("VONAGE_API_KEY"), goDotEnvVariable("VONAGE_API_SECRET"))

	// client := nexmo.NewClient(http.DefaultClient, auth)

	// smsContent := nexmo.SendSMSRequest{
	// 	From: goDotEnvVariable("FROM"),
	// 	To:   goDotEnvVariable("TO"),
	// 	Text: "This is a message sent from Go!",
	// }

	// smsResponse, _, err := client.SMS.SendSMS(smsContent)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	auth := nexmo.CreateAuthFromKeySecret(goDotEnvVariable("VONAGE_API_KEY"), goDotEnvVariable("VONAGE_API_SECRET"))
	smsClient := nexmo.NewSMSClient(auth)
	response, err := smsClient.Send(goDotEnvVariable("FROM"), goDotEnvVariable("TO"), "This is a message from golang", nexmo.SMSOpts{})

	if err != nil {
		panic(err)
	}

	if response.Messages[0].Status == "0" {
		fmt.Println("Account Balance: " + response.Messages[0].RemainingBalance)
	}

	fmt.Println("Status:", response.Messages[0].Status)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/webhook/inbound-sms", receiveSms)
	router.HandleFunc("/send-sms", sendSms)
	log.Fatal(http.ListenAndServe(":8080", router))
}
