package main

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
	//icon_emoji string
}

func post(text string) {
	m := Message{text, "#support", "fogbugz", "http://www.fogcreek.com/images/fogbugz/pricing/kiwi.png"}
	b, err := json.Marshal(m)
	fmt.Printf("%s => %s (err=%i) \n", m, b, err)
	http.Post(webhookurl, "text/json", bytes.NewReader(b))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received: %s!\n", r.URL)
	s, _ := url.QueryUnescape(r.URL.String()[1:])
	s = strings.Replace(s, "http:/", "http://", 1)
	fmt.Printf("Decoded as: %s!\n", s)
	post(s)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <slack-webhook-url>\n", os.Args[0])
		fmt.Printf("e.g: %s https://company.slack.com/services/hooks/incoming-webhook?token=loremipsum\n", os.Args[0])
		fmt.Printf("\nConfigure a fogbugz URL Trigger like to send on Case events to:\n")
		fmt.Printf("http://your-proxy-host:10333//{CaseNumber}: {EventType} - {AssignedToName} - <http:/your-fogbugz-host/default.asp?{CaseNumber}|{Title}>")
		return
	}
	webhookurl = os.Args[1]
	fmt.Printf("Using: %s\n", webhookurl)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":10333", nil)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
