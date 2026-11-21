package main

import (
	"encoding/json"
	"context"
	"fmt"
	"net/http"
	"time"
	"os"

	"github.com/S1lverTongue/go-github-tracker/internal/types"
)

func makeUserRequest(username string, ctx context.Context) (*http.Request, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

	return req, nil
}

func makeUserEventsRequest(username string, ctx context.Context) (*http.Request, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

	return req, nil
}

func printPushEvent(user string, response types.GitEvent, event types.PushEvent) {
	fmt.Printf("%s pushed to %s on branch %s\n", user, response.Repo.Name, event.Ref)
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	if len(os.Args) < 2 {
		fmt.Println("No username specified")
		os.Exit(1)
	}

	user := os.Args[1]

	req, err := makeUserEventsRequest(user, ctx)

	if err != nil {
		panic(err)
	}

	var client = &http.Client{}
	resp, err := client.Do(req)

	if resp.StatusCode != 200 {
		fmt.Println("User not Found")
		os.Exit(1)
	}

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var gitResponse []types.GitEvent
	if err := json.NewDecoder(resp.Body).Decode(&gitResponse); err != nil {
		panic(err)
	}

	for _, response := range gitResponse {
		switch response.Type {
		case "PushEvent":
			var event types.PushEvent
			json.Unmarshal(response.Payload, &event)
			printPushEvent(user, response, event)
		default:
			fmt.Println(response.Type)
		}
	}
}