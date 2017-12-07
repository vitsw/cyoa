package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
	"os"
	"github.com/vitsw/cyoa"
)

func main() {
	storyFile := flag.String("story", "gopher.json", "file with json story")
	tplFile := flag.String("tpl", "default", "custom template file")
	port := flag.Int("port", 8080, "webserver port")
	flag.Parse()

	jsonF, err := os.Open(*storyFile)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonF.Close()

	story, err := cyoa.JSONStory(jsonF)
	if err != nil {
		log.Fatal(err)
	}

	if *tplFile == "default" {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), cyoa.NewHandler(*story)))
	} else {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), cyoa.NewHandlerWithTemplate(*story, *tplFile)))
	}
}
