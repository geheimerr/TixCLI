package main

import "fmt"

func main() {
	var confName string = "GoCon"
	const conftix int = 50
	var remtix uint = 50

	fmt.Println("Welcome to TixCLI.")
	fmt.Println("Get your tickets to attend", confName)
	fmt.Println("Total GoCon tickets:", conftix, "		Remaining GoCon tickets:", remtix)

	var userName string
	var usertix uint
	var userEmail string
	fmt.Printf("Please enter your username: %v", userName)
	fmt.Scan(&userName)
	fmt.Printf("Hi %v, how many tickets do you want? ", userName)
	fmt.Scan(&usertix)
	fmt.Print("Enter your email ID. The purchased tickets shall be sent the email ID entered: ")
	fmt.Scan(&userEmail)
	fmt.Printf("%v tickets purchased succesfully. Your tickets have been sent to %v \n", usertix, userEmail)
	var updatedtixCount = remtix - usertix
	var bookings []string
	bookings = append(bookings, userName)

	fmt.Printf("Current bookings: %v\n", bookings)
	fmt.Printf("%v tickets remaining\n", updatedtixCount)

}
