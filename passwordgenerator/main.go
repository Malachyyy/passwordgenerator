package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	Digits = "0123456789"

	Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
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

func Confirmatie() string {
	var (
		Hoofdletters string
		Nummers      string
		Symbolen     string
	)

	Combinatie := strings.Builder{}

	fmt.Println("Wilt u hoofdletters? [y|n] ")
	fmt.Scanln(&Hoofdletters)

	fmt.Println("Wilt u nummers? [y|n] ")
	fmt.Scanln(&Nummers)

	fmt.Println("Wilt u symbolen? [y|n] ")
	fmt.Scanln(&Symbolen)

	strings.ToLower(Hoofdletters)
	strings.ToLower(Nummers)
	strings.ToLower(Symbolen)

	if Hoofdletters == "y" {
		Combinatie.WriteString(string(UpperLetters))
	}

	if Nummers == "y" {
		Combinatie.WriteString(string(Digits))
	}

	if Symbolen == "y" {
		Combinatie.WriteString(string(Symbols))
	}
	return Combinatie.String()
}

func main() {
	fmt.Println("Kies de lengte van het wachtwoord: ")

	var PasswordLength int
	fmt.Scanln(&PasswordLength)

	Combinations := LowerLetters + Confirmatie()

	Password := GenerateRandomString(PasswordLength, Combinations)

	fmt.Println(Password)
}
