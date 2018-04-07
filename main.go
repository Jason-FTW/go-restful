package main

import (
	"go-restful/api"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Router endpoints and API
	router.HandleFunc("", api.GetAccounts)
}