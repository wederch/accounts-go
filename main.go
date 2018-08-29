package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Account structure
type Account struct {
	Name string
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	accounts := &[]Account{
		Account{"Salary"},
		Account{"Savings"},
		Account{"Stocks"},
		Account{"Commodities"},
		Account{"FX"},
		Account{"Funds"},
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

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func serveEmptyResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}
