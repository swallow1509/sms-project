package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

//Const values as initial varibles
const (
	accountSID = "AC8a33c89d0a3b3f53d455c55e720f4805"
	authToken  = "5edae12bc1a68a01d5050fdd50dc0c37"
	uslString  = "https://api.twilio.com/2010-04-01/Accounts/" + accountSID + "/Messages.json"
)

func main() {

	//Basic data for message
	val := url.Values{}
	val.Set("To", "+998990777558")
	val.Set("From", "+16266289463")
	val.Set("Body", "Sardorbek qaleysiz? Men Umidjonman. Amerikadan SMS jonatyapman")
	reqBody := *strings.NewReader(val.Encode())

	//New client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", uslString, &reqBody)
	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")

	//Create request
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)

}
