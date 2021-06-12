package main

import (
	"github.com/chipgo/observer/club"
)

func main() {
	MUFanClub := club.NewFanClub("Manchester United")

	tom := club.NewFan("Tom")
	bob := club.NewFan("Bob")
	alice := club.NewFan("Alice")
	mike := club.NewFan("Mike")

	MUFanClub.RegisterObservers(tom, bob, alice, mike)

	MUFanClub.SetPoint(89)
	MUFanClub.RemoveObserver(mike)
	MUFanClub.SetPoint(92)
}
