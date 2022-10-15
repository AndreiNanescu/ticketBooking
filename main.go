package main

import "fmt"

func main() {
	const eventName = "Test name"
	const tickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to the %v booking application.\n", eventName)
	fmt.Printf("We have a total of %v tickets of which %v are available.\n", tickets, remainingTickets)
	fmt.Println("Get your tickets here.")
}
