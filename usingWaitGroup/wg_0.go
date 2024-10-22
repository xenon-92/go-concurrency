package usingWaitGroup

import (
	"fmt"
	"sync"
)

func wg0() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"theta",
		"kappa",
		"pi",
		"silon",
	}
	wg.Add(len(words))

	for i, val := range words {
		go displaySomething(fmt.Sprintf("%d --> %s", i, val), &wg)
	}
	wg.Wait()

	wg.Add(1)
	displaySomething("this is the second thing to be printed", &wg)
	wg.Wait()
}

func displaySomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
