package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
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
		}

		time.Sleep(5 * time.Second)
	}
}
