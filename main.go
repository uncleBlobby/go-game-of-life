package main

import (
	"os"
	"os/exec"
	"time"
)

const (
	WORLD_HEIGHT = 24
	WORLD_WIDTH  = 80
	ORIGIN_X     = 0
	ORIGIN_Y     = 0
	TICK_RATE    = "100ms"
	LIVE_CHAR    = rune('\u25A0')
	DEAD_CHAR    = ' '
)

func main() {
	Clear()

	world := InitWorld()

	for {
		world.DrawLoop()
		world.ProcessGeneration()
	}

}

func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func SleepClear(t string) {
	dur, _ := time.ParseDuration(t)
	time.Sleep(dur)
	Clear()
}
