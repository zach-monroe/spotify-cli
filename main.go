/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"log"

	"github.com/zach-monroe/spotify-cli/spot"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found!")
	}
}

func main() {
	body, err := spot.GetAuth()
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf(body)
	//cmd.Execute()

}
