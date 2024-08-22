package main

import "log"

func assertGenerationError(err error) {
	if err != nil {
		log.Fatalf("generation error: %s", err.Error())
	}
}
