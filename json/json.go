package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Company struct {
	Name              string    `json:"name"`
	FoundedOn         time.Time `json:"founded_on"`
	NumberOfEmployees int       `json:"number_of_employees"`
	Stock             Stock     `json:"stock"`
}

type Stock struct {
	Ticker      string    `json:"ticker"`
	LastUpdated time.Time `json:"last_updated"`
}

func main() {
	marshal()
	fmt.Println()
	unmarshal()
}

func marshal() {
	company := getCompany()
	json, err := json.Marshal(company)

	if err != nil {
		log.Fatalf("Unable to marshal struct\n%s", err)
	}

	//'string(json)' converts a array of bytes to a string.
	fmt.Println("marshal() function\n", string(json))
}

func unmarshal() {
	//It creates a empty company
	company := Company{}

	jsonCompany := `{
		"name":"Boston Properties",
		"founded_on":"1970-12-25T00:00:00-03:00",
		"number_of_employees":3000,
		"stock":{
			"ticker":"NYSE: BXP",
			"last_updated":"2022-03-22T22:47:51.8167377-03:00"
			}
		}`

	//json.Unmarshal() has a local scope. So, we use "&" to change the same
	//memory address of 'company' variable present in the main thread.
	json.Unmarshal([]byte(jsonCompany), &company)
	fmt.Println("unmarshal() function\n", company)
}

func getCompany() Company {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	foundedOn := time.Date(1970, 12, 25, 0, 0, 0, 0, loc)
	lastUpdated := time.Now()

	company := Company{
		Name:              "Boston Properties",
		FoundedOn:         foundedOn,
		NumberOfEmployees: 3000,
		Stock: Stock{
			Ticker:      "NYSE: BXP",
			LastUpdated: lastUpdated,
		},
	}
	return company
}
