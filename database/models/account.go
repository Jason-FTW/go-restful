package models

import (
	"time"
	"github.com/satori/go.uuid"
)

// Account is the structure for the account model
// Note that we are using bson instead of json, as that is what mgo handles
type Account struct {
	ID				uuid.UUID	`bson:"id"`
	FirstName		string		`bson:"firstname"`
	LastName		string		`bson:"firstname"`

	Username		string		`bson:"username"`
	Password		string		`bson:"password"`
	CreatedOn		time.Time	`bson:"createdOn"`
	LastModified	time.Time	`bson:"lastModified"`
}

// NewAccount is a constructor that returns a new struct instance of an Account
// This function initializes the default values for UUID as a new UUIDv4; CreatedOn and LastModified as time.Now()
func NewAccount(firstName, lastName, username, password string) Account {
	now := time.Now()
	return Account{uuid.NewV4(), firstName, lastName, username, password, now, now}
}