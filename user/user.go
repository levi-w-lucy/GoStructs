package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (appUser User) OutputUserDetails() {
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthDate)
}

func (appUser *User) ClearUserName() {
	appUser.firstName = ""
	appUser.lastName = ""
}

func NewUser(firstName, lastName, birthDate string) (*User, error) {
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

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
