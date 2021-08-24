package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

// Credentials stores all of our access/api tokens
// and secret keys needed for authentication against
// the twitter REST API

type Credentials struct {
	API_KEY       string
	API_SECRET    string
	ACCESS_TOKEN  string
	ACCESS_SECRET string
}

// get_client is a function that will return a twitter client
// that can be used to send tweets, or to stream new tweets
// this takes a Credential structure which contains the keys to authenticate
// and return a pointer to a twitter Client (or throws an error)

func get_client(creds *Credentials) (*twitter.Client, error) {
	fmt.Println("Let's give this twitter thing a try...")

	config := oauth1.NewConfig(creds.API_KEY, creds.API_SECRET)
	token := oauth1.NewToken(creds.ACCESS_TOKEN, creds.ACCESS_SECRET)

	http_client := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(http_client)

	// verify credentials

	verify := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// get user and check credentials worked
	user, _, e := client.Accounts.VerifyCredentials(verify)
	if e != nil {
		return nil, e
	}

	fmt.Printf("User:%+v\n", user.ScreenName)

	return client, nil

}

func main() {
	fmt.Println("Ini's Go twitter bot")

	cmd_args := os.Args[1]

	ld := godotenv.Load("config.env")
	if ld != nil {
		log.Fatal("Error loading env file")
	}

	creds := Credentials{
		API_KEY:       os.Getenv("API_KEY"),
		API_SECRET:    os.Getenv("API_SECRET"),
		ACCESS_TOKEN:  os.Getenv("ACCESS_TOKEN"),
		ACCESS_SECRET: os.Getenv("ACCESS_SECRET"),
	}
	fmt.Printf("%+v\n", creds)

	client, e := get_client(&creds)
	if e != nil {
		log.Println("Error getting Twitter Client")
		log.Println(e)
	}
	tweet, resp, err := client.Statuses.Update(cmd_args, nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("response code: %+v\n", resp.Status)
	log.Printf("tweet that you tweeted: %+v\n", tweet.Text)
}
