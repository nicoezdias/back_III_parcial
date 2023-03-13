package tickets

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Id                 int
	Name               string
	Email              string
	DestinationCountry string
	FlightTime         string
	Price              int
}

/* -------------------------------- Exercise 1 ------------------------------- */
func GetTotalTickets(destination string) (int, error) {
	total := 0
	Tickets, err := RecoverData()
	if err != nil {
		return total, err
	}
	for _, v := range Tickets {
		if v.DestinationCountry == destination {
			total++
		}
	}
	return total, nil
}

/* -------------------------------- Exercise 2 ------------------------------- */
func GetCountByPeriod(time string, chTotal chan int, chErr chan error) {
	var condition func(int) bool
	total := 0
	Tickets, err := RecoverData()
	if err != nil {
		chErr <- err
	}
	switch time {
	case "EarlyMorning":
		condition = GetEarlyMornings
	case "Morning":
		condition = GetMornings
	case "Afternoon":
		condition = GetAfternoons
	case "Night":
		condition = GetNights
	default:
		chErr <- errors.New("Error: The time entered is incorrect")
	}
	for _, v := range Tickets {
		hour, err := strconv.Atoi(strings.Split(string(v.FlightTime), ":")[0])
		if err != nil {
			chErr <- errors.New("Error: Can not convert strings to type int")
		}
		if condition(hour) {
			total++
		}
	}
	chTotal <- total
}
func GetEarlyMornings(hour int) bool {
	return hour >= 0 && hour <= 6
}
func GetMornings(hour int) bool {
	return hour >= 7 && hour <= 12
}
func GetAfternoons(hour int) bool {
	return hour >= 13 && hour <= 19
}
func GetNights(hour int) bool {
	return hour >= 20 && hour <= 23
}

/* -------------------------------- Exercise 3 ------------------------------- */
func AverageDestination(destination string) (float64, error) {
	Tickets, err1 := RecoverData()
	if err1 != nil {
		return 0, err1
	}
	total, err2 := GetTotalTickets(destination)
	if err2 != nil {
		return 0, err2
	}
	cantFlights := len(Tickets)
	return float64(total) / float64(cantFlights), nil
}

/* ------------------------------- RecoverData ------------------------------ */
func RecoverData() ([]Ticket, error) {
	var tickets = []Ticket{}
	res, err := os.ReadFile("./tickets.csv")
	if err != nil {
		return []Ticket{}, errors.New("Error: Can not read file")
	}
	data := strings.Split(string(res), "\n")
	for _, d := range data {
		if len(d) > 0 {
			var tick Ticket
			cat := strings.Split(d, ",")
			intVar, err := strconv.Atoi(cat[0])
			if err != nil {
				return []Ticket{}, errors.New("Error: Can not convert strings to type int")
			}
			tick.Id = intVar
			tick.Name = cat[1]
			tick.Email = cat[2]
			tick.DestinationCountry = cat[3]
			tick.FlightTime = cat[4]
			intVar, err = strconv.Atoi(cat[5])
			if err != nil {
				return []Ticket{}, errors.New("Error: Can not convert strings to type int")
			}
			tick.Price = intVar
			tickets = append(tickets, tick)
		}
	}
	return tickets, nil
}
