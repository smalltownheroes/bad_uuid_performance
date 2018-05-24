package main

import (
	"flag"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

func main() {
	var max = flag.Int("m", 1000, "Number of iterations")
	flag.Parse()

	start := time.Now()
	for i := 0; i != *max; i++ {
		someUuids()
	}
	t := time.Now()
	fmt.Printf("Ran %d loops in %s\n", *max, t.Sub(start))
}

func someUuids() {
	_ = uuid.Must(uuid.NewV4()).String()
	_ = uuid.Must(uuid.NewV4()).String()
	_ = uuid.Must(uuid.NewV4()).String()
	_ = uuid.Must(uuid.NewV4()).String()
}
