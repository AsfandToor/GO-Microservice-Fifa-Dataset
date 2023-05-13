package controller

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net/http"
	"strings"
)

func GetAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	query := `	
		{
			"query": {
				"match_all": {} 
			}
		}
	`

	var b strings.Builder
	b.WriteString(query)
	read := strings.NewReader(b.String())

	fmt.Println(read)

	res, err := es.Search(
		es.Search.WithBody(read),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	// Decode the response body into a map
	var data map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}

	// Encode the map as JSON and write it to the response
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func GetQueryData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	params := r.URL.Query()

	var b strings.Builder

	for key, values := range params {
		for _, value := range values {
			b.WriteString(fmt.Sprintf(`{ "query": { "match": { "%v": "%v" } } }`, key, value))
		}
	}

	read := strings.NewReader(b.String())

	fmt.Println(read)

	res, err := es.Search(
		es.Search.WithBody(read),
		es.Search.WithPretty(),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	// Decode the response body into a map
	var data map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}

	// Encode the map as JSON and write it to the response
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
