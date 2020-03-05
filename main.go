package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var data_feed interface{}

func collector() {

	url := "https://api.twitter.com/1.1/tweets/search/30day/dev.json?result_type=mixed"
	method := "POST"
	payload := strings.NewReader("{\n	\"query\":\"#openbanking OR #apifirst OR #devops OR #cloudfirst OR #microservices OR #apigateway OR #oauth OR #swagger OR #raml OR #openapis\",\n	\"maxResults\":\"10\"\n}")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer <TWITTER-TOKEN>")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)


	fmt.Println(string(body))

}


func main() {

	collector()

}
