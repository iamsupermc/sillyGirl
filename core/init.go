package core

import (
	"bufio"
	"os"
	"regexp"
	"time"
)

var Duration time.Duration

func init() {
	killp()
	for _, arg := range os.Args {
		if arg == "-d" {
			initStore()
			Daemon()
		}
	}
	initStore()
	initToHandleMessage()
	InitReplies()
	file, err := os.Open(ExecPath + "/conf/sets.conf")
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if regexp.MustCompile(`^set`).MatchString(line) {
				Senders <- &Faker{
					Message: line,
				}
			}
		}
		file.Close()
	}
	initSys()
	Duration = time.Duration(sillyGirl.GetInt("duration", 5)) * time.Second
}
