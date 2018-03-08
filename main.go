package main

import (
	"bufio"
	"flag"
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

	// 2. We get the logs to send to CronCron

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		data := url.Values{}
		data.Set("logs", scanner.Text())
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
			log.Fatalln("Wrong token")

		default:
			log.Printf("Failed sending logs to CronCron: %v", resp.Status)
		}
	}

	// Check for errors during `Scan`. End of file is
	// expected and not reported by `Scan` as an error.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
