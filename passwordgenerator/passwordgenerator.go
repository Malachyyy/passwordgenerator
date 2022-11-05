package main

import (
	"fmt"
	"log"

	"github.com/sethvargo/go-password/password"
)

func generatepassword(lengte int) {
	password, err := password.Generate(lengte, 3, 3, false, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(password)
}

func main() {
	fmt.Println("Hoe lang is het wachtwoord? ")

	var lengte int

	fmt.Scanln(&lengte)

	generatepassword(lengte)
}
