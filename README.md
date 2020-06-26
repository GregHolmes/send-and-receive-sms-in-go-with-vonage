![Vonage][logo]

# Send and Receive SMS In Go With Vonage

The code found in `main.go` provides two urls `/send-sms` and `/webhook/inbound-sms`. 
Triggering `/send-sms` will send an SMS to the `To` number defined in `.env`.

**Table of Contents**

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Clone the Repository](#clone-the-repository)
  - [Database Credentials](#database-credentials)
  - [Run Docker](#run-docker)
  - [Install Third Party Libraries](#install-third-party-libraries)
  - [Test Run the Application](#test-run-the-application)
- [Code of Conduct](#code-of-conduct)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

- [Go](https://golang.org/)
- [A Vonage (formally Nexmo) account](https://dashboard.nexmo.com/sign-up?utm_source=DEV_REL&utm_medium=github&utm_campaign=send-and-receive-sms-in-go-with-vonage)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Getting Started

### Clone the Repository

Run the following two commands to clone this repository and change directory into the repository directory.

```
git clone https://github.com/GregHolmes/send-and-receive-sms-in-go-with-vonage
cd send-and-receive-sms-in-go-with-vonage
```

### Credentials

Rename the `.env.example` file to `.env` and populate the values. The VONAGE_API_KEY and VONAGE_API_SECRET can be found on your [Dashboard](https://dashboard.nexmo.com/sign-up?utm_source=DEV_REL&utm_medium=github&utm_campaign=send-and-receive-sms-in-go-with-vonage).

```env
VONAGE_API_KEY=
VONAGE_API_SECRET=
FROM=<Your Vonage number>
TO=<The recipients number>
```

### Install Third Party Libraries

To install the libraries used, run the following commands:

```go
	go get github.com/gorilla/mux
	go get github.com/gorilla/schema
	go get github.com/joho/godotenv
	go get github.com/nexmo-community/nexmo-go
```

### Test Run the Application

Go to: [http://localhost:8080/send-sms](http://localhost:8080/send-sms) in your browser, if configured correctly, you should have sent yourself an SMS. Check your phone!

## Code of Conduct

In the interest of fostering an open and welcoming environment, we strive to make participation in our project and our community a harassment-free experience for everyone. Please check out our [Code of Conduct][coc] in full.

## Contributing
We :heart: contributions from everyone! Check out the [Contributing Guidelines][contributing] for more information.

[![contributions welcome][contribadge]][issues]

## License

This project is subject to the [MIT License][license]

[logo]: vonage_logo.png "Vonage"
[contribadge]: https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat "Contributions Welcome"

[coc]: CODE_OF_CONDUCT.md "Code of Conduct"
[contributing]: CONTRIBUTING.md "Contributing"
[license]: LICENSE "MIT License"

[issues]: ./../../issues "Issues"
