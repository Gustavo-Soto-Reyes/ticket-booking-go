package main

import "fmt"

func main() {
	conferenceName := "Generic Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to our booking application for %v\n", conferenceName)
	fmt.Printf("We have %v tickets, and %v remaining tickets\n", conferenceTickets, remainingTickets)
	fmt.Println("Buy your tickets here")

	var bookings []string

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// getting user input
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for this %v", remainingTickets, conferenceName)
}
