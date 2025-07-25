package main

import (
	"log"
	"os"

	"github.com/fflewddur/ltbsky"
)

func main() {
	BasicExample()
	ImageExample()
	ImageAndFacetsExample()
}

func BasicExample() {
	log.Println("bsky basic example")
	server := os.Getenv("BSKY_SERVER")
	handle := os.Getenv("BSKY_HANDLE")
	password := os.Getenv("BSKY_PASSWORD")
	client, err := ltbsky.NewClient(server, handle, password)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	log.Printf("Client created: %+v", client)
	postBuilder := ltbsky.NewPostBuilder("Hello, world!")
	uri, err := client.Post(postBuilder)
	if err != nil {
		log.Fatalf("Error posting: %v", err)
	}
	log.Printf("Post created with URI: %s", uri)
	log.Println("bsky test completed")
}

func ImageExample() {
	log.Println("bsky image example")
	server := os.Getenv("BSKY_SERVER")
	handle := os.Getenv("BSKY_HANDLE")
	password := os.Getenv("BSKY_PASSWORD")
	client, err := ltbsky.NewClient(server, handle, password)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	postBuilder := ltbsky.NewPostBuilder("Hello with image!")
	postBuilder.AddLang("en")
	postBuilder.AddImageFromPath("./test-data/bsky-go-1.png", "Alt text for image")

	uri, err := client.Post(postBuilder)
	if err != nil {
		log.Fatalf("Error posting: %v", err)
	}
	log.Printf("Post created with URI: %s", uri)
	log.Println("bsky image example completed")
}

func ImageAndFacetsExample() {
	log.Println("bsky image and facets example")
	server := os.Getenv("BSKY_SERVER")
	handle := os.Getenv("BSKY_HANDLE")
	password := os.Getenv("BSKY_PASSWORD")
	client, err := ltbsky.NewClient(server, handle, password)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	postBuilder := ltbsky.NewPostBuilder("Hello with image and facets! https://go.dev @golang.org")
	postBuilder.AddLang("en")
	postBuilder.AddImageFromPath("./test-data/bsky-go-1.png", "Alt text for image")

	uri, err := client.Post(postBuilder)
	if err != nil {
		log.Fatalf("Error posting: %v", err)
	}
	log.Printf("Post created with URI: %s", uri)
	log.Println("bsky image and facets example completed")
}
