package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
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

		var body bytes.Buffer

		if err := json.Indent(&body, buf, " >", "  "); err != nil {
			panic(err)
		}
		log.Println(body.String())
	})))
}
