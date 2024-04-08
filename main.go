package main

import (
	"os"

	"github.com/charmbracelet/log"
)

type WideReceiver struct {
	Name string
	Team string
}

var wideReceivers = []WideReceiver{
	{"Justin Jefferson", "Minnesota Vikings"},
	{"Davante Adams", "Las Vegas Raiders"},
	{"Stefon Diggs", "Buffalo Bills"},
	{"A.J. Brown", "Philadelphia Eagles"},
	{"Jaylen Waddle", "Miami Dolphins"},
	{"Tyreek Hill", "Miami Dolphins"},
	{"CeeDee Lamb", "Dallas Cowboys"},
	{"Keenan Allen", "Los Angeles Chargers"},
	{"Terry McLaurin", "Washington Commanders"},
	{"Deebo Samuel", "San Francisco 49ers"},
	{"Mike Evans", "Tampa Bay Buccaneers"},
	{"Chris Godwin", "Tampa Bay Buccaneers"},
	{"Amon-Ra St. Brown", "Detroit Lions"},
	{"DK Metcalf", "Seattle Seahawks"},
	{"Michael Pittman Jr.", "Indianapolis Colts"},
}

func main() {

	pyramid := [15]string{}

	for i := 0; i < 15; i++ {
		pyramid[i] = wideReceivers[i].Name
	}

	logger := log.New(os.Stdout)
	logger.SetFormatter(log.JSONFormatter)
	logger.SetLevel(log.InfoLevel)

	logger.Info(pyramid)
}
