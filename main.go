package main

import (
	"fmt"
	"strings"
)

func main() {
	gatheringName := "IT собрание"
	const gatheringTickets int = 50
	var remainingTickets uint = 50

	bookings := []string{}

	fmt.Printf("Data Types of conferenceName is: %T , conferenceTickets is: %T and remainingTickets is: %T\n", gatheringName, gatheringTickets, remainingTickets)
	fmt.Printf("Welcome to our %v booking Application\n", gatheringName)
	fmt.Printf("We have total numbers of %v tickets and %v are still available\n", gatheringTickets, remainingTickets)
	fmt.Println("Get Your Ticket Now")

	for remainingTickets > 0 && len(bookings) < 50 { //возвращаем количество записей
		//задаем переменные
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidaeTickets := userTickets > 0 && userTickets <= remainingTickets

	}
}
