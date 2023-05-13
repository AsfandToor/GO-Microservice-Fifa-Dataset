package main

import (
	"QueryService/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.GetAllData).Methods("GET")
	r.HandleFunc("/fifa", controller.GetQueryData).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", r))
}
