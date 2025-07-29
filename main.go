package main

import "fmt"

func main() {
	confName := "GoCon"
	const conftix int = 50
	var remtix uint = 50
	var bookings []string

	fmt.Println("Welcome to TixCLI.")
	fmt.Printf("Get your tickets to attend %v\n", confName)
	fmt.Printf("Total GoCon tickets: %v    Remaining GoCon tickets: %v\n", conftix, remtix)
	fmt.Println("--------------------------------------------------")

	for {
		if remtix == 0 {
			fmt.Println("Sorry, all tickets for GoCon have been sold out. Thank you!")
			break
		}

		var userName string
		var usertix uint
		var userEmail string

		fmt.Print("Please enter your username: ")
		fmt.Scan(&userName)

		fmt.Printf("Hi %v, how many tickets do you want? ", userName)
		fmt.Scan(&usertix)

		if usertix > remtix {
			fmt.Printf("We're sorry, you can't purchase %v tickets. We only have %v remaining.\n", usertix, remtix)
			fmt.Println("--------------------------------------------------")
			continue
		}

		fmt.Print("Enter your email ID. The purchased tickets will be sent to this address: ")
		fmt.Scan(&userEmail)

		remtix -= usertix
		bookings = append(bookings, userName)

		fmt.Printf("\nThank you, %v! %v tickets purchased successfully.\n", userName, usertix)
		fmt.Printf("A confirmation has been sent to %v\n", userEmail)
		fmt.Printf("Current bookings: %v\n", bookings)
		fmt.Printf("%v tickets remaining for %v.\n", remtix, confName)
		fmt.Println("--------------------------------------------------")
	}
}
