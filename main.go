package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/nicklaw5/helix"
)

func isUserOnline(username string) (bool, error) {
	log.Println("Creating client")
	client, err := helix.NewClient(&helix.Options{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	})

	if err != nil {
		return false, err
	}

	log.Println("Requesting App Access Token")
	resp, err := client.RequestAppAccessToken([]string{})
	if err != nil {
		return false, err
	}
	log.Println("Setting App Access Token")
	client.SetAppAccessToken(resp.Data.AccessToken)

	log.Printf("Fetching streams with user_login=%s", username)
	resp2, err := client.GetStreams(&helix.StreamsParams{
		UserLogins: []string{username},
	})
	log.Println(resp2)
	if err != nil {
		return false, err
	}

	if len(resp2.Data.Streams) > 0 {
		return true, nil
	}

	return false, nil
}

func getLastNonEmptyPart(parts []string) string {
	for i := len(parts) - 1; i >= 0; i-- {
		if len(parts[i]) > 0 {
			return parts[i]
		}
	}
	return ""
}

func isSafe(name string) bool {
	matched, err := regexp.Match("^[a-zA-Z0-9-_.]+$", []byte(name))
	if err != nil {
		log.Println(err)
		return false
	}

	return matched
}

func onlineHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	user := getLastNonEmptyPart(parts)
	log.Printf("Got username %v from %v", user, r.URL.Path)

	if !isSafe(user) {
		log.Fatalf("Username %v seemed unsafe, so bailing out", user)
	}

	if len(user) > 0 {
		online, err := isUserOnline(user)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Add("Content-Type", "image/svg+xml")
		if online {
			http.ServeFile(w, r, "online.svg")
		} else {
			http.ServeFile(w, r, "offline.svg")
		}
	}
}

func nilHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/favicon.ico", nilHandler)
	http.HandleFunc("/", onlineHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
