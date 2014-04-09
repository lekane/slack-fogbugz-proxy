package main

// Copyright Lekane Oy. All rights reserved.
// Written by Kim Blomqvist <kim.blomqvist@lekane.com>

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var webhookurl string

type Message struct {
	Text     string `json:"text"`
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Icon_url string `json:"icon_url"`
}

func post(text string) {
	m := Message{text, "#support", "fogbugz", "http://www.fogcreek.com/images/fogbugz/pricing/kiwi.png"}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("ERROR '%s' marshalling message: %s\n", err, m)
		return
	}
	fmt.Printf("Posting: %s\n", b)
	http.Post(webhookurl, "text/json", bytes.NewReader(b))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("-------------\n")
	fmt.Printf("Received: %s\n", r.URL)
	s, _ := url.QueryUnescape(r.URL.String()[1:])
	s = strings.Replace(s, "http:/", "http://", 1)
	fmt.Printf("Decoded as: %s\n", s)
	post(s)
	fmt.Printf("-------------\n")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <slack-webhook-url>\n", os.Args[0])
		fmt.Printf("e.g: %s https://company.slack.com/services/hooks/incoming-webhook?token=loremipsum\n", os.Args[0])
		fmt.Printf("\nConfigure a fogbugz URL Trigger like to send on Case events to:\n")
		fmt.Printf("http://your-proxy-host:10333//{CaseNumber}: {EventType} - {AssignedToName} - <http:/your-fogbugz-host/default.asp?{CaseNumber}|{Title}>\n")
		return
	}
	webhookurl = os.Args[1]
	port := ":10333"
	fmt.Printf("Listening to port: %s\n", port)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
