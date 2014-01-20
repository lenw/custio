package custio

import (
	"net/http"
	"bytes"
	"log"
	"net/url"
	"strings"
)

// PUT https://track.customer.io/api/v1/customers/{CUSTOMER_ID}
// {"email":"customer@example.com","name":"Bob","plan":"premium","array":["1","2","3"]}

const (
	cioUrlCustomers = "https://track.customer.io/api/v1/customers/"
)

func Customer(siteId, apiKey, id, email, name string) {
	values := make(url.Values)
	values.Set("name", name)
	values.Set("email", email)
	request, _ := http.NewRequest("PUT", cioUrlCustomers+id, strings.NewReader(values.Encode()))
	request.SetBasicAuth(siteId, apiKey)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(request)
	log.Printf("CIO -> R %v : E %v", resp, err)

}


func Track(siteId, apiKey, id, eventName string, data map[string]string) {
	values := make(url.Values)
	values.Set("name", eventName)
	for k, v := range data {
		values.Set("data["+k+"]", v)
	}

	request, _ := http.NewRequest("POST", cioUrlCustomers+id+"/events", strings.NewReader(values.Encode()))
	client := http.Client{}
	request.SetBasicAuth(siteId, apiKey)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(request)
	log.Printf("CIO -> R %v : E %v", resp, err)

}
