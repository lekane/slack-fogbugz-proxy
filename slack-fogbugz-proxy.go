package main

// Written by Kim Blomqvist <kim.blomqvist@lekane.com>

/*
Copyright (c) 2014 Lekane Oy. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Lekane Oy nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

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
