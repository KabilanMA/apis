package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v37/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func getLanguages(username *string, accessToken string) (map[string]int, error) {

	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, tokenSource)

	client := github.NewClient(tc)

	// Get user repos
	repos, _, err := client.Repositories.List(ctx, *username, nil)
	if err != nil {
		return nil, err
	}

	// Count languages in each repo
	languages := make(map[string]int)
	for _, repo := range repos {
		langs, _, err := client.Repositories.ListLanguages(ctx, *username, repo.GetName())

		if err != nil {
			return nil, err
		}

		for lang := range langs {
			languages[lang]++
		}
	}

	return languages, nil
}

func main() {
	username := "kabilanma"

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	github_acT := os.Getenv("GITHUB_ACCESS_TOKEN")
	if github_acT == "" {
		fmt.Println("GITHUB_ACCESS_TOKEN IS NOT SET")
		return
	}

	languages, err := getLanguages(&username, github_acT)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Languages used by %s:\n", username)
	for lang, count := range languages {
		fmt.Printf("%s: %d\n", lang, count)
	}
}
