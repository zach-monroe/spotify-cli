/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/zach-monroe/spotify-cli/spot"

	"time"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found!")
	}
}

var lastTimeIGotAuth time.Time

type SpotToken struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

// function for recalling locally cached tokens. Invoked later in the "cachedToken" var.
func loadToken() (SpotToken, error) {

	var tokenToLoad SpotToken

	file, err := os.ReadFile(".token.json")
	if err != nil {
		return SpotToken{}, nil
	}
	err = json.Unmarshal(file, &tokenToLoad)
	if err != nil {
		return tokenToLoad, err
	}
	return tokenToLoad, nil

}

// function for saving token to a local json file
func saveToken(token SpotToken) error {

	data, err := json.Marshal(token)
	if err != nil {
		return err
	}
	err = os.WriteFile(".token.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	cachedToken, err := loadToken()
	finalToken := cachedToken
	if err != nil {
		log.Panic(err)
	}

	if cachedToken.Token != "" && time.Since(cachedToken.CreatedAt) > 1*time.Hour {
		token, err := spot.GetAuth()
		if err != nil {
			log.Panic(err)
		}
		now := time.Now()
		yourToken := SpotToken{
			Token:     token,
			CreatedAt: now,
		}
		err = saveToken(yourToken)
		if err != nil {
			log.Panic(err)
		}
		finalToken = yourToken
	}

	fmt.Printf(finalToken.Token)
	//cmd.Execute()

}
