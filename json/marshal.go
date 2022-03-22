package main

import (
	"encoding/json"
	"log"
	"time"
)

type Company struct {
	Name              string    `json:"name"`
	Ticker            string    `json:"ticker"`
	FoundedOn         time.Time `json:"founded_on"`
	LastUpdated       time.Time `json:"last_updated"`
	NumberOfEmployees int       `json:"number_of_employees"`
}

func main() {
	company := getCompany()
	json, err := json.Marshal(company)

	if err != nil {
		log.Fatalf("Unable to marshal struct\n%s", err)
	}

	//'string(json)' converts a array of bytes to a string.
	println(string(json))
}

func getCompany() Company {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	foundedOn := time.Date(1970, 12, 25, 0, 0, 0, 0, loc)
	lastUpdated := time.Now()

	company := Company{
		Name:              "Boston Properties",
		Ticker:            "NYSE: BXP",
		FoundedOn:         foundedOn,
		LastUpdated:       lastUpdated,
		NumberOfEmployees: 3000,
	}
	return company
}
