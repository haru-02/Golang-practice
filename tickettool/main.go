package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) 
var firstName string
var lastName string
var email string
var userTickets uint

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
} 

var wg = sync.WaitGroup{}

func main() {

	greetUsers()
		firstName, lastName, email, userTickets = getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput()

		if  isValidEmail && isValidName && isValidTicketNumber {
			
			bookTicket()
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The slice is : %v \n",firstNames)


			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
			//	break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("your email is invalid, please try again.")
			}
			if !isValidTicketNumber {
				fmt.Println("requested tickets exceed availability, please try again.")
			}
		}
		wg.Wait()
}

func greetUsers() {
	fmt.Printf("\nWelcome to %v booking application.\n", conferenceName)	
	fmt.Printf("We have a total capacity of %v seats and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("please get your tickets here to attend")
} 

func getFirstNames() []string {
	firstnames := []string{}
	for _, booking := range bookings {
		firstnames = append(firstnames, booking.firstName)
	}
	return firstnames;
}

func getUserInput () (string, string, string, uint){

	fmt.Print("\nplease enter your first name : ")
	fmt.Scan(&firstName)
	fmt.Print("please enter your last name : ")
	fmt.Scan(&lastName)
	fmt.Print("please enter your email : ")
	fmt.Scan(&email)
	fmt.Print("please enter the number of tickets booking : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket()  {
	remainingTickets -= userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
		
	fmt.Printf("thank you %v %v for booking %v tickets. you will recieve a conformation email at %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("there are %v tickets remaining for the %v.\n", remainingTickets, conferenceName)
} 

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v \n", ticket, email)
	fmt.Println("########################")
	wg.Done()
}