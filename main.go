package main

import (
	"fmt"
	"strconv"
	"strings"
)

var bookings = make([]map[string]string, 0)

const conferenceTickets = 50

var remainingTickets uint = 50

func main() {
	conferenceName := "Generic Conference"

	fmt.Printf("Welcome to our booking application for %v\n", conferenceName)
	fmt.Printf("We have %v tickets, and %v remaining tickets\n", conferenceTickets, remainingTickets)
	fmt.Println("Buy your tickets here")

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for true { // you can define a condition after for

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email, conferenceName)
			firstNames := getFirstName()
			fmt.Printf("Here are the first names %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Out conference is booked out. Come back next time ")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}

	}
}

func greetUsers(confName string, tickets int, remaining uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
}

func getFirstName() []string {
	firstnames := []string{}
	for _, booking := range bookings {
		firstnames = append(firstnames, booking["firstName"])

	}
	return firstnames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for this %v\n", remainingTickets, conferenceName)
}
