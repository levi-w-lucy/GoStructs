package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (appUser user) OutputUserDetails() {
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthDate)
}

func (appUser *user) clearUserName() {
	appUser.firstName = ""
	appUser.lastName = ""
}

func newUser(firstName, lastName, birthDate string) (*user, error) {
	errMessage := ""
	if firstName == "" {
		errMessage += "First name was missing. "
	}

	if lastName == "" {
		errMessage += "Last name was missing. "
	}

	if birthDate == "" {
		errMessage += "Birth Date was missing."
	}

	if errMessage != "" {
		return nil, errors.New("First name, last name, and birth date are required. " + errMessage)
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
