package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/groomer/gTalk/db"
	"github.com/groomer/gTalk/router"
)

func init() {
	db.Connect()
}

func main() {
	mux := router.SetupRouter()

	log.Fatal(http.ListenAndServe(":3000", mux))
	fmt.Println("gTalk Server")
}
