package main

import (
	"fmt"
	"log"
	"os"
	"readwise-app/core"

	"github.com/atotto/clipboard"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	tokenReadwise := os.Getenv("READWISE_TOKEN")
	tokenOpenAI := os.Getenv("OPENAI_TOKEN")
	if tokenReadwise == "" || tokenOpenAI == "" {
		log.Fatal("READWISE_TOKEN or OPENAI_TOKEN not set")
	}

	quote, err := core.GetFavoriteQuote(tokenReadwise)
	if err != nil {
		log.Fatal("error getting favorite quote")
	}

	prompt, err := core.GetReflectionPrompt(tokenOpenAI, quote)
	if err != nil {
		log.Fatal("error getting reflection prompt")
	}

	fmt.Println(prompt)

	fmt.Println("Copying text to clipboard...")
	if err := clipboard.WriteAll(prompt); err != nil {
		log.Fatal("error copying to clipboard")
	}
}
