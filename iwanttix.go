package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"strings"
	"time"
)

func sending_email(gmail string, password string) {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		gmail,
		password,
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		gmail,
		[]string{gmail},
		[]byte("Go Go !! To get the ticket !!"),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// smtp settings
	gmail := flag.String("g", "foo@bar.com", "your gmail account")
	password := flag.String("p", "foo123", "your password")

	flag.Parse()

	if flag.NFlag() != 2 {
		flag.Usage()
		return
	}

	for true {
		response, err := http.Get("http://golang.kktix.cc/events/gtg13")
		if err != nil {
			fmt.Printf("%s", err)
			continue
		}

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			continue
		}

		no_ticket := strings.Contains(string(contents), "額滿")
		if !no_ticket {
			fmt.Printf("\a幹有票了！ (%s) \n", time.Now())
			sending_email(*gmail, *password)
		}

		time.Sleep(5 * time.Second)
	}
}
