package main

import (
	"fmt"
	"log"

	"github.com/pavles6/projector-go/pkg/projector"
)

func main() {
	opts, err := projector.GetOpts()

	if err != nil {
		log.Fatalf("Unable to get options %v\n", err)
	}

	fmt.Printf("%+v", opts)
}
