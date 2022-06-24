package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/prometheus/alertmanager/notify/webhook"
)

func main() {
	listenAddr := flag.String("listen-address", ":12345", "The address to listen on for incoming webhook requests.")
	flag.Parse()

	log.Printf("Starting up on %s to listen for incoming webhook messages...", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		buf, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()

		message := &webhook.Message{}
		if err := json.Unmarshal(buf, message); err != nil {
			panic(err)
		}

		fmt.Printf("Received %d alert(s) for group key %v:\n", len(message.Alerts), message.GroupKey)
		for i, alert := range message.Alerts {
			fmt.Printf("\t%d. %s: %v\n", i+1, alert.Status, alert.Labels)
		}
	})))
}
