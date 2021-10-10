package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/flexzuu/website/query/client"
	"github.com/flexzuu/website/query/markdown"

	"github.com/joho/godotenv"
)

func main() {
	// Load variables from env file
	_ = godotenv.Load()

	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	endpoint := os.Getenv("ENDPOINT")
	if endpoint == "" {
		log.Fatal("must set ENDPOINT")
	}

	key := os.Getenv("TOKEN")
	if key == "" {
		log.Fatal("must set TOKEN")
	}

	basePath := os.Getenv("BASE_PATH")
	if basePath == "" {
		log.Fatal("must set BASE_PATH")
	}

	c := client.New(endpoint, key)

	posts, err := c.Posts(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts.Posts {
		fileName := fmt.Sprintf("%s.md", post.Slug)

		file, err := os.Create(path.Join(basePath, "post", fileName))
		if err != nil {
			log.Fatalf("could not create File %s: %s", fileName, err.Error())
		}
		err = markdown.PostTemplate.Execute(file, post)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
