package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	Digits = "0123456789"

	Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"

	AllCharSet = LowerLetters + UpperLetters + Digits + Symbols
)

func GenerateRandomString(PasswordLength int, Combinations string) string {
	ret := make([]byte, PasswordLength)
	for i := 0; i < PasswordLength; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(len(Combinations))))
		if err != nil {
			fmt.Println(err)
		}
		ret[i] = Combinations[random.Int64()]
	}

	return string(ret)
}

func main() {
	fmt.Println("Kies de lengte van het wachtwoord: ")

	var lengte int
	fmt.Scanln(&lengte)

	PasswordLength := lengte
	Combinations := AllCharSet
	Password := GenerateRandomString(PasswordLength, Combinations)

	fmt.Println(Password)
}
