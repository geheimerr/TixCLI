package main

import (
	"fmt"
	"strings"
	"time"
)

var usertix uint
var firstName string
var lastName string
var userEmail string

func main() {
	greetUsers()
	confName := "GoCon"
	const conftix int = 50
	var remtix uint = 50
	var bookings = make([]map[string]string, 0)

	fmt.Printf("Get your tickets to attend %v\n", confName)
	fmt.Printf("Total GoCon tickets: %v    Remaining GoCon tickets: %v\n", conftix, remtix)
	fmt.Println("--------------------------------------------------")

	for remtix > 0 {

		fmt.Print("Please enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Please enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email ID: ")
		fmt.Scan(&userEmail)

		fmt.Printf("Hi %v, how many tickets do you want? ", firstName)
		fmt.Scan(&usertix)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(userEmail, "@")
		isValidTicketNum := usertix > 0 && usertix <= remtix

		if isValidName && isValidEmail && isValidTicketNum {
			remtix -= usertix

			var userData = make(map[string]string)
			userData["firstName"] = firstName
			userData["lastName"] = lastName
			userData["email"] = userEmail
			userData["tickets"] = fmt.Sprintf("%d", usertix)

			bookings = append(bookings, userData)

			fmt.Printf("\nThank you, %v! %v tickets purchased successfully.\n", firstName, usertix)
			fmt.Printf("A confirmation has been sent to %v\n", userEmail)
			fmt.Printf("%v tickets remaining for %v.\n", remtix, confName)

			firstNames := []string{}
			for _, booking := range bookings {
				firstNames = append(firstNames, booking["firstName"])
			}
			fmt.Printf("Current bookings: %v\n", firstNames)

		} else {
			fmt.Println("\n--- Invalid Input ---")
			if !isValidName {
				fmt.Println("First name or last name is too short.")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered is missing an '@'.")
			}
			if !isValidTicketNum {
				fmt.Printf("You requested %v tickets, but we only have %v available or your request was invalid.\n", usertix, remtix)
			}
			fmt.Println("Please try again.")
		}
		fmt.Println("--------------------------------------------------")
	}

	fmt.Println("Sorry, all tickets for GoCon have been sold out. Thank you!")
}

func greetUsers() {
	fmt.Println("Welcome to TixCLI.")
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
}
