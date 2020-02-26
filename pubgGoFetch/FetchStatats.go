package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"
)

const apiKey = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIyNDUyYjJkMC0zNDZjLTAxMzgtMmU4Zi02MzE5NDU3ZGU0YzUiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNTgyMDIzNzc2LCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6Im5pY2stcm9zZW5kYWwtIn0.n-q-FKt_KBvhxryV7gg7eNNgrQ79AbfFw3YJA8_Qjw8"

const apiAddress = "https://api.pubg.com/shards/steam/players?filter[playerNames]=LackOfHonor,GoldenGipsy"

func main() {
	res, err := funcName()
	if err == nil {
		body, err := ioutil.ReadAll(res.Body)
		if err == nil {
			var record TeamRecord
			json.Unmarshal(body, &record)
			for _, p := range record.Data {
				name, matches := getPlayerMatchsIds(p)
				fmt.Println(name, "played in: ", len(matches))
			}

		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println("err")
		fmt.Println(err)
	}

}

func getPlayerMatchsIds(playerOne PlayerRecord) (string, []string) {
	//should remove all non matches from the collection..
	sort.Slice(playerOne.Relationships.Matches.Data, func(i, j int) bool {
		return playerOne.Relationships.Matches.Data[i].Type == "match"
	})
	var matchIds []string
	for _, m := range playerOne.Relationships.Matches.Data {
		matchIds = append(matchIds, m.ID)
	}
	return playerOne.Attributes.Name, matchIds
}

func funcName() (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiAddress, nil)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Accept", "application/vnd.api+json")
	return client.Do(req)
}

type PlayerRecord struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Name         string      `json:"name"`
		Stats        interface{} `json:"stats"`
		TitleID      string      `json:"titleId"`
		ShardID      string      `json:"shardId"`
		CreatedAt    time.Time   `json:"createdAt"`
		UpdatedAt    time.Time   `json:"updatedAt"`
		PatchVersion string      `json:"patchVersion"`
	} `json:"attributes"`
	Relationships struct {
		Assets struct {
			Data []interface{} `json:"data"`
		} `json:"assets"`
		Matches struct {
			Data []struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"matches"`
	} `json:"relationships"`
	Links struct {
		Self   string `json:"self"`
		Schema string `json:"schema"`
	} `json:"links"`
}

type TeamRecord struct {
	Data []PlayerRecord
}
