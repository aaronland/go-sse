package main

import (
	"flag"
	"fmt"
	"github.com/sfomuseum/go-sse"
	"log"	
)

func cb(raw []byte) error {
	fmt.Println(string(raw))
	return nil
}

func main() {

	endpoint := flag.String("endpoint", "", "")

	flag.Parse()

	client, err := sse.NewClient(*endpoint)

	if err != nil {
		log.Fatal(err)
	}

	for {
		err := client.Listen(cb)

		if err != nil {
			log.Fatal(err)
		}
	}

}
