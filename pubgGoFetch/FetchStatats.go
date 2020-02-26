package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiKey = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIyNDUyYjJkMC0zNDZjLTAxMzgtMmU4Zi02MzE5NDU3ZGU0YzUiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNTgyMDIzNzc2LCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Im5pY2stcm9zZW5kYWwtIn0.n-q-FKt_KBvhxryV7gg7eNNgrQ79AbfFw3YJA8_Qjw8"

//const apiAddress = "https://api.pubg.com/shards/steam/players?filter[playerNames]=GoldenGipsy"
const apiAddress = "https://api.pubg.com/shards/steam/players?filter[playerNames]=LackOfHonor,GoldenGipsy"

func main() {
	res, err := funcName()
	if err == nil {
		fmt.Println("no err")
		//fmt.Println(res)
		body, err := ioutil.ReadAll(res.Body)
		if err == nil {
			fmt.Println(string(body))
		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println("err")
		fmt.Println(err)
	}

}

func funcName() (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiAddress, nil)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Accept", "application/vnd.api+json")
	return client.Do(req)
}
