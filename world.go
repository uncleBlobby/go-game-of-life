package main

import (
	"fmt"
	"math/rand"
	"time"
)

type World struct {
	Cells       [][]Cell
	Generations int
	Living      int
	FrameTime   time.Duration
}

func InitWorld() World {
	var world World

	for j := 0; j < WORLD_HEIGHT; j++ {
		var line Line
		for i := 0; i < 80; i++ {
			cell := Cell{false, Position{X: i, Y: j}, ' '}
			line.Cells = append(line.Cells, cell)
		}
		world.Cells = append(world.Cells, line.Cells)
	}

	world.RandomizeSeed()
	world.Generations = 0

	return world
}

func (w *World) DrawLoop() {
	for _, line := range w.Cells {
		for i, cell := range line {
			if i < len(line)-1 {
				cell.Draw()
			}
			if i == len(line)-1 {
				fmt.Println()
			}
		}
	}
	w.CountLiving()
	fmt.Printf("Generation #%d\n", w.Generations)
	fmt.Printf("Living: %d\n", w.Living)
	fmt.Printf("FrameTime: %v\n", w.FrameTime)
	SleepClear(TICK_RATE)
}

func (w *World) RandomizeSeed() {
	for i, line := range w.Cells {
		for j := range line {
			num := rand.Int()
			if num%10 == 0 {
				w.Cells[i][j].Alive = true
				w.Cells[i][j].Char = rune('\u25A0')
			}
			//fmt.Println()
		}
	}
}

func (w *World) ProcessGeneration() {
	start := time.Now()
	var newWorld = w
	for j := 0; j < WORLD_HEIGHT; j++ {
		for i := 0; i < WORLD_WIDTH; i++ {
			w.Cells[j][i].Life(newWorld)
			//newWorld.Cells[j][i] = w.Cells[j][i]
		}
	}
	w.Generations += 1
	w.FrameTime = time.Since(start)

	//w = &newWorld
}

func (w *World) CountLiving() {
	w.Living = 0
	for _, line := range w.Cells {
		for _, cell := range line {
			if cell.Alive {
				w.Living += 1
			}
		}
	}
}
