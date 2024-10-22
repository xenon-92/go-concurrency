package main

import (
	_ "github.com/xenon-92/go-concurrency/usingSleep"
	"github.com/xenon-92/go-concurrency/usingWaitGroup"
)

func main() {
	// usingSleep.RunMain()
	usingWaitGroup.RunMain()
}
