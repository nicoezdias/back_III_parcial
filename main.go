package main

import (
	"entregaback/internal/tickets"
	"fmt"
	"log"
)

const (
	Early     string = "EarlyMorning"
	Morning   string = "Morning"
	Afternoon string = "Afternoon"
	Night     string = "Night"
)

func main() {
	/* ------------------------- Flights to destinations ------------------------ */
	destination := "Brazil"
	total, err1 := tickets.GetTotalTickets(destination)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("Flights to %s: %d\n", destination, total)

	/* ---------------------------- Flights for times --------------------------- */
	chEarly := make(chan int)
	chMorning := make(chan int)
	chAfternoon := make(chan int)
	chNight := make(chan int)
	chError := make(chan error)

	go tickets.GetCountByPeriod(Early, chEarly, chError)
	select {
	case err := <-chError:
		log.Fatal(err)
	case total1 := <-chEarly:
		fmt.Printf("Flights at EarlyMornings: %d\n", total1)
	}

	go tickets.GetCountByPeriod(Morning, chMorning, chError)
	select {
	case err := <-chError:
		log.Fatal(err)
	case total2 := <-chMorning:
		fmt.Printf("Flights at Mornings: %d\n", total2)
	}

	go tickets.GetCountByPeriod(Afternoon, chAfternoon, chError)
	select {
	case err := <-chError:
		log.Fatal(err)
	case total3 := <-chAfternoon:
		fmt.Printf("Flights at Afternoons: %d\n", total3)
	}

	go tickets.GetCountByPeriod(Night, chNight, chError)
	select {
	case err := <-chError:
		log.Fatal(err)
	case total4 := <-chNight:
		fmt.Printf("Flights at Nights: %d\n", total4)
	}

	/* ------------------------- Average per destination ------------------------ */
	average, err2 := tickets.AverageDestination(destination)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("Average flights to %s: %v\n", destination, average)
}
