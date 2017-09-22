package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		Anleitung()
		return
	}

	command := strings.ToLower(os.Args[1])
	if command == "remainingtill" {
		RemainingTill()
		return
	}

	fmt.Println("Unknown argument " + command)
	os.Exit(1)
}

func Anleitung() {
	fmt.Println("PomodoroGo - Das Pomodoro-Tool (shoff, 2017-09-22)")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("   - RemainingTill xx (e.g. RemainingTill 18)")
	fmt.Println("     Remaining Pomodoros that are possible until xx o'clock.")
	fmt.Println("")
}

func RemainingTill() {
	if len(os.Args) != 3 {
		fmt.Println("Parameters not ok")
		os.Exit(1)
	}

	untilHour, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println("Could not parse your hour parameter.")
		os.Exit(1)
	}

	var currentTime int
	currentTime = time.Now().Hour()*60 + time.Now().Minute()
	untilInMinutes := untilHour * 60
	remaining := (int(untilInMinutes) - currentTime) / 30
	fmt.Printf("There are %v pomodoros remaining till %v o'clock.\n", remaining, untilHour)
}
