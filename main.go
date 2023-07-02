package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func main() {
	opt := option.WithCredentialsFile("testfb-ae3ac-firebase-adminsdk-w90ma-a547943d06.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		errStr := fmt.Sprintf("error initializing app: %v", err)
		log.Fatal(errStr)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		errStr := fmt.Sprintf("error initializing app: %v", err)
		log.Fatal(errStr)
	}

	data := struct {
		EventId    int `json:"event_id"`
		GameId     int `json:"game_id"`
		GameStatus int `json:"game_status"`
		WalaBet    int `json:"wala_bet"`
		WalaOdd    int `json:"wala_odd"`
		MeronBet   int `json:"meron_bet"`
		MeronOdd   int `json:"meron_odd"`
	}{
		EventId:    1,
		GameId:     1,
		GameStatus: 0,
		WalaBet:    11000,
		WalaOdd:    1,
		MeronBet:   4500,
		MeronOdd:   2,
	}

	_, err = client.Collection("current-game-status").Doc("game-status").Set(context.Background(), &data)
	if err != nil {
		log.Print(err)
	}
	defer client.Close()

	fmt.Println("inside firestore")
}
