package main

import (
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"log"
	"os"
	"strings"
	"syscall"
	"time"
	"github.com/0xAX/notificator"
)

func main() {
	forbidden := os.Args[1:]

	if len(forbidden) == 0{
		fmt.Println("Usage: `killer process1 process2`")
		os.Exit(1)
	}

	notify := notificator.New(notificator.Options{
		DefaultIcon: "",
		AppName:     "Killer",
	})

	notificationsSend := map[string]time.Time{}


	for {

		processList, err := ps.Processes()
		if err != nil {
			log.Fatal("ps.Processes() Failed, are you using windows?")
		}

		// map ages
		for _, p := range processList {
			pName := strings.ToLower(p.Executable())

			for _, f := range forbidden {
				containsName :=  strings.Contains(pName, f)
				if !containsName {
					continue
				}
				percentage := float64(len(f))/float64(len(pName))
				if percentage < 0.5{
					continue
				}

				log.Printf("Kill %d\t%s\n", p.Pid(), p.Executable())

				err := syscall.Kill(p.Pid(), 9)
				if err != nil {
					log.Println("Could not kill:", err)
					continue
				}

				noti :=  fmt.Sprintf("%q is forbidden and killed",p.Executable())

				if t,ok := notificationsSend[noti]; ok && time.Now().Sub(t) < time.Second{
					continue
				}

				notify.Push("Forbidden",noti, "", notificator.UR_NORMAL)
				notificationsSend[noti] = time.Now()
			}
		}
		time.Sleep(time.Second)
	}

}