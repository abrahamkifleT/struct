package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userfristName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY)")

	var appUser user

	appUser = user{
		firstName: userfristName,
		lastName:  userlastName,
		birthDate: userbirthdate,
		createdAt: time.Now(),
	}

	outputUserDetails(appUser)
}

func outputUserDetails(u user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
