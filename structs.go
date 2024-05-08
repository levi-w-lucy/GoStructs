package main

import "fmt"

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthday: ")

	fmt.Println(firstName, lastName, birthdate)
}

func getUserData(promptText string) (inputValue string) {
	fmt.Print(promptText)
	fmt.Scan(&inputValue)
	return inputValue
}
