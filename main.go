package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://api.twitter.com/1.1/tweets/search/30day/dev.json?result_type=mixed"
	method := "POST"
	payload := strings.NewReader("{\n	\"query\":\"#openbanking OR #apifirst OR #devops OR #cloudfirst OR #microservices OR #apigateway OR #oauth OR #swagger OR #raml OR #openapis\",\n	\"maxResults\":\"10\"\n}")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAAEjDCgEAAAAAVtCArfmG%2FHYi752sKmUHFaaqZUc%3DyTr3Mn58g28egkjS7jwJHY8h5fOXYK5K4KcMDt1BN7R6M5tlbs")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
