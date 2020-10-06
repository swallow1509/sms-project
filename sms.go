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
	accountSID = "AC8a33c89d0a3b3f53d455c55e720f4805"
	authToken  = "b57f7e3e1678236762b7bdca5fe22d90"
	urlString  = "https://api.twilio.com/2010-04-01/Accounts/AC8a33c89d0a3b3f53d455c55e720f4805/Messages.json"
)

func main() {

	// Basic data for our message
	v := url.Values{}
	v.Set("To", "+998998280098")
	v.Set("From", "+16266289463")
	v.Set("Body", "hehehe Umidjon qaleysiz?")
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
