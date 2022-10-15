package main

import "fmt"

func main() {
	const eventName = "Test name"
	const tickets = 50
	var remainingTickets = 50
	fmt.Println("Welcome to the", eventName, "booking application.")
	fmt.Println("We have a total of", tickets, "of wich", remainingTickets, "are remaining.")
}
