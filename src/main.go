package main

import "fmt"

func main() {
	const conferenceTickets uint = 50
	conferenceName := "Go Conference"
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var email string
	var numTickets uint
	pointRemainingTickets := &remainingTickets

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email:")
	fmt.Scan(&email)
	fmt.Println("Please enter the amount of tickets you would like:")
	fmt.Scan(&numTickets)
	fmt.Printf("%v %v you have purchased %v tickets.  We will send conformation to: %v\n", firstName, lastName, numTickets, email)

	*pointRemainingTickets = remainingTickets - numTickets

	fmt.Printf("There are %v tickets still available.\n", remainingTickets)
}
