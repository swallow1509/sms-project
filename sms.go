package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	accountSID = "TwilioSID"
	authToken  = "b57f7e3e1678236762b7bdca5fe22d90"
	urlString  = "https://api.twilio.com/2010-04-01/Accounts/accountSID/Messages.json"
)

func main() {

	// Basic data for our message
	v := url.Values{}
	v.Set("To", "+Receiver")
	v.Set("From", "+VirtualTwilioSenderNumber")
	v.Set("Body", "You are pretty Welcome to put your custom SMS here!")
	reqBody := *strings.NewReader(v.Encode())

	// New client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlString, &reqBody)
	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
}
