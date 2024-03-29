package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/groomer/gTalk/controller"
	"github.com/groomer/gTalk/db"
	"github.com/groomer/gTalk/router"
)

func Ticker24Hour() {
	ticker := time.NewTicker(24 * time.Hour)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("24 hour ticker")
				controller.GetReminders()
			case <-quit:
				fmt.Println("24 hour ticker stop")
				ticker.Stop()
				return
			}
		}
	}()
}

func init() {
	db.Connect()
	controller.MessageStopper = make(map[uint64](chan bool))
	controller.GetReminders()

	now := time.Now()
	hour := (23 - now.Hour()) % 24
	minute := (60 - now.Minute())
	if minute == 60 {
		hour += 1
		minute = 0
	} else {
		minute %= 60
	}

	everyMidnight := time.NewTicker(time.Duration(hour)*time.Hour + time.Duration(minute)*time.Minute)
	go func() {
		for {
			select {
			case <-everyMidnight.C:
				fmt.Println("first Reminder call!")
				controller.GetReminders()
				go Ticker24Hour()
				return
			}
		}
	}()
}

func main() {
	mux := router.SetupRouter()

	log.Fatal(http.ListenAndServe(":3000", mux))
	fmt.Println("gTalk Server")
}
