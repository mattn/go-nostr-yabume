package main

import (
	"context"
	"io"
	"log"
	"os"

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

	io.Copy(os.Stdout, resp.Body)
}
