package pnp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// GetToken function is used to get the Auth Token
func GetToken(endPoint string, clientID string, clientSecret string) map[string]interface{} {
	requestBody, err := json.Marshal(map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     clientID,
		"client_secret": clientSecret})

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(endPoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	return result
}
