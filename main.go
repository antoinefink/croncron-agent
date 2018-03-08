package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const pushLogsAPIEndpoint = "https://api.croncron.io/push_logs"

func main() {

	// 1. We make sure the run token is defined.

	token := flag.String("token", "", "Token of the run (required)")
	flag.Parse()

	if *token == "" {
		*token = os.Getenv("CRONCRON_TOKEN")
	}

	if *token == "" {
		log.Fatal("Run token shouldn't be empty")
	}

	// 2. We get the log to send to CronCron

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() == 0 {
		log.Fatal("Nothing send to standard input. Check our documentation.")
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	// 3. We send the log to CronCron using the API

	data := url.Values{}
	data.Set("logs", string(input))
	data.Set("token", *token)

	req, err := http.NewRequest("POST", pushLogsAPIEndpoint, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	switch resp.StatusCode {
	case 200:
		// All good! Let's stay quiet.

	case 403:
		log.Println("Wrong token")

	default:
		log.Printf("Failed sending logs to CronCron: %v", resp.Status)
	}
}
