package main

import (
	"math/rand"

	"github.com/google/uuid"
)

type Account struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Number    int64   `json:"number"`
	Balance   float64 `json:"balance"`
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		ID:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Int()),
	}
}
