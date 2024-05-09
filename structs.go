package main

import (
	"fmt"
)

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthday: ")

	appUser, err := newUser(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()
	appUser.clearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) (inputValue string) {
	fmt.Print(promptText)
	fmt.Scanln(&inputValue)
	return inputValue
}
