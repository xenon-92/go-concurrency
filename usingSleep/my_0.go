package usingSleep

import (
	"time"
)

func my0() {
	go displaySomething("this is one")
	time.Sleep(1 * time.Second) // using time.sleep to wait on the results
	displaySomething("this is two")
}
