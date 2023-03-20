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
	var Tickets tickets.Tickets

	/* ------------------------------- RecoverData ------------------------------ */
	err := Tickets.RecoverData()
	if err != nil {
		log.Fatal(err)
	}
	/* ------------------------- Flights to destinations ------------------------ */
	destination := "Brazil"
	total := Tickets.GetTotalTickets(destination)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Flights to %s: %d\n", destination, total)

	/* ---------------------------- Flights for times --------------------------- */
	chEarly := make(chan int)
	chMorning := make(chan int)
	chAfternoon := make(chan int)
	chNight := make(chan int)
	chError := make(chan error)
	go func() {
		Tickets.GetCountByPeriod(Early, chEarly, chError)
		Tickets.GetCountByPeriod(Morning, chMorning, chError)
		Tickets.GetCountByPeriod(Afternoon, chAfternoon, chError)
		Tickets.GetCountByPeriod(Night, chNight, chError)
	}()
	select {
	case err := <-chError:
		log.Fatal(err)
	case total1 := <-chEarly:
		fmt.Printf("Flights at EarlyMornings: %d\n", total1)
	}
	select {
	case err := <-chError:
		log.Fatal(err)
	case total2 := <-chMorning:
		fmt.Printf("Flights at Mornings: %d\n", total2)
	}
	select {
	case err := <-chError:
		log.Fatal(err)
	case total3 := <-chAfternoon:
		fmt.Printf("Flights at Afternoons: %d\n", total3)
	}
	select {
	case err := <-chError:
		log.Fatal(err)
	case total4 := <-chNight:
		fmt.Printf("Flights at Nights: %d\n", total4)
	}

	/* ------------------------- Average per destination ------------------------ */
	average, err := Tickets.AverageDestination(destination)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Average flights to %s: %v\n", destination, average)
}
