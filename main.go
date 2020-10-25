package main

import (
	"fmt"
	"github.com/0xAX/notificator"
	ps "github.com/mitchellh/go-ps"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	forbidden := os.Args[1:]

	if len(forbidden) == 0 {
		fmt.Println("Usage: `killer process1 process2`")
		os.Exit(1)
	}

	notify := notificator.New(notificator.Options{
		DefaultIcon: "",
		AppName:     "Killer",
	})

	notificationsSend := map[string]time.Time{}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	go func() {
		for range c {
			fmt.Println(" ...nice try. Now go back to work ;)")
		}
	}()

	for {

		processList, err := ps.Processes()
		if err != nil {
			log.Fatal("ps.Processes() Failed, are you using windows?")
		}

		// map ages
		for _, p := range processList {
			pName := strings.ToLower(p.Executable())

			for _, f := range forbidden {
				containsName := strings.Contains(pName, f)
				if !containsName {
					continue
				}
				percentage := float64(len(f)) / float64(len(pName))
				if percentage < 0.5 {
					continue
				}

				log.Printf("Kill %d\t%s\n", p.Pid(), p.Executable())

				err := syscall.Kill(p.Pid(), 9)
				if err != nil {
					log.Println("Could not kill:", err)
					continue
				}

				noti := fmt.Sprintf("%q is forbidden and killed", p.Executable())

				if t, ok := notificationsSend[noti]; ok && time.Now().Sub(t) < time.Second {
					continue
				}

				notify.Push("Forbidden", noti, "", notificator.UR_CRITICAL)
				notificationsSend[noti] = time.Now()
			}
		}
		time.Sleep(time.Second)
	}

}
