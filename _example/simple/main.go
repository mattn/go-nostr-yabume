package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mattn/go-nostr-yabume"
)

func main() {
	client, err := yabume.NewClient("https://api.yabu.me")
	if err != nil {
		log.Fatal(err)
	}
	id := "note1a0vtesecwwwfpeycecjne2pdduvcpa77vf9cjfclg84msmytacss4z9784"
	resp, err := client.GetV0EventsId(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	eventId, err := yabume.ParseGetV0EventsIdResponse(resp)
	if err != nil {
		log.Fatal(err)
	}
	ev := eventId.JSON200
	fmt.Println(ev.Content)
}
