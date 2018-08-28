package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Account structure
type Account struct {
	Name string
}

func main() {

	accounts := &[]Account{
		Account{"Salary"},
		Account{"Savings"},
		Account{"Stocks"},
		Account{"Commodities"},
	}

	http.HandleFunc("/accounts", func(w http.ResponseWriter, r *http.Request) {
		switch m := (*r).Method; m {
		case http.MethodGet:
			bytes, _ := json.Marshal(accounts)
			fmt.Fprint(w, string(bytes))
		default:
			serveEmptyResponse(w, r)
		}
	})

	http.HandleFunc("/", serveEmptyResponse)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveEmptyResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}
