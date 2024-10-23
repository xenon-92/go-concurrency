package usingSleep

import (
	"fmt"
	"time"
)

func my1() {
	words := []string{
		"alpha",
		"beta",
		"delta",
		"theta",
		"kappa",
		"pi",
		"silon",
	}

	for i, str := range words {
		go displaySomething(fmt.Sprintf("%d --> %s", i, str))
	}
	time.Sleep(1 * time.Second)
	displaySomething("this is two")

}
