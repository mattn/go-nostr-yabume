# go-nostr-yabume

yabu.me client for Go

## Usage

```go
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
```

## Installation

```
go get github.com/mattn/go-nostr-yabume
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
