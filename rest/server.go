package main

import (
	"go-restful/api"

	"github.com/gorilla/mux"

	"time"
	"github.com/globalsign/mgo"
	"go-restful/constants"
	"go-restful/database"
)

// Environment: one of (dev/development, prod/production)
// Do NOT MODIFY HERE. @see ./env.go
const env = constants.Env

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Objectify the mgo dialing information. We will use this to connect to the database
	mgoInfo := &mgo.DialInfo {
		Addrs:		[] string {constants.DBHosts},
		Timeout:	60 * time.Second,
		Database: 	constants.DBName,
		Username:	constants.DBUserName,
		Password:	constants.DBPassword,
	}

	// Initialize the mgo driver with the above information
	session, err := mgo.DialWithInfo(mgoInfo)
	session.SetMode(mgo.Monotonic, true)

	// Panic and log err if there is an error
	if err != nil {
		panic(err)
	}

	// Initialize the database
	err = database.Init(session)

	// Close the mgo session once the program has concluded
	defer session.Close()

	// Router endpoints and API

	// Account API
	router.HandleFunc("/api/accounts", api.GetAccounts).Methods("GET")
	router.HandleFunc("/api/account/{id}", api.GetAccount).Methods("GET")
	router.HandleFunc("/api/account", api.CreateAccount).Methods("POST")
	router.HandleFunc("/api/account", api.DeleteAccount).Methods("DELETE")
}