package main

import (
	"fmt"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Turtle Appreciation Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These are all of our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("All tickets have now been booked")
				break
			}
		} else {
			fmt.Println("------\nInvalid input:")
			if !isValidName {
				fmt.Println("First or last name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email must contain @ symbol")
			}
			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets, you tried to buy %v tickets\n", remainingTickets, userTickets)
			}
			fmt.Println("-------------")
		}
	}
}

func greetUser() {
	fmt.Printf("\nWelcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets remaining out of %v\n", remainingTickets, conferenceTickets)
	fmt.Printf("You can buy all the %v tickets you need here!\n-----\n", conferenceName)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create struct for user
	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. We will send your tickets to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// simulate time delay to send ticket, to practice concurrency
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("----------")
	fmt.Printf("Sending ticket:\n%v \nto %v\n", ticket, email)
	fmt.Println("----------")
}
