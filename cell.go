package main

import "fmt"

type Line struct {
	Cells []Cell
}

type Position struct {
	X int
	Y int
}

type Cell struct {
	Alive    bool
	Position Position
	Char     rune
}

func (c Cell) Draw() {
	fmt.Print(string(c.Char))
}

func (c Cell) GetNeighbors(w *World) []Cell {
	neighbors := []Cell{}

	if c.Position.X > ORIGIN_X {
		if c.Position.Y > ORIGIN_Y {
			// left up
			neighbors = append(neighbors, w.Cells[c.Position.Y-1][c.Position.X-1])
		}

		// left mid
		neighbors = append(neighbors, w.Cells[c.Position.Y][c.Position.X-1])
		if c.Position.Y < WORLD_HEIGHT-1 {
			// left down
			neighbors = append(neighbors, w.Cells[c.Position.Y+1][c.Position.X-1])
		}

	}

	if c.Position.X < WORLD_WIDTH-1 {
		if c.Position.Y > ORIGIN_Y {
			// right up
			neighbors = append(neighbors, w.Cells[c.Position.Y-1][c.Position.X+1])
		}

		// right mid
		neighbors = append(neighbors, w.Cells[c.Position.Y][c.Position.X+1])

		if c.Position.Y < WORLD_HEIGHT-1 {
			// right down
			neighbors = append(neighbors, w.Cells[c.Position.Y+1][c.Position.X+1])
		}

	}

	if c.Position.Y < WORLD_HEIGHT-1 {
		// down mid
		neighbors = append(neighbors, w.Cells[c.Position.Y+1][c.Position.X])
	}

	if c.Position.Y > ORIGIN_Y {
		// up mid
		neighbors = append(neighbors, w.Cells[c.Position.Y-1][c.Position.X])
	}

	return neighbors
}

func (c *Cell) CountLiveNeighbors(w *World) int {
	neighbors := c.GetNeighbors(w)

	sum := 0

	for _, n := range neighbors {
		if n.Alive {
			sum += 1
		}
	}

	return sum
}

func (c *Cell) Life(w *World) {
	liveNeighbors := c.CountLiveNeighbors(w)

	if c.Alive {
		if liveNeighbors < 2 {
			c.Alive = false
			c.Char = DEAD_CHAR
		}
		if liveNeighbors == 2 || liveNeighbors == 3 {
			c.Alive = true
			c.Char = LIVE_CHAR
		}
		if liveNeighbors > 3 {
			c.Alive = false
			c.Char = DEAD_CHAR
		}
	}

	if !c.Alive && liveNeighbors == 3 {
		c.Alive = true
		c.Char = LIVE_CHAR
	}
}
