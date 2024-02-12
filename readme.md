# go-game-of-life

This project is meant to be an experiment in reproducing Conway's Game of Life using the Go programming language and a simple TUI.

The description, taken from Wikipedia, is below:

The universe of the Game of Life is an infinite, two-dimensional orthogonal grid of square cells, each of which is in one of two
possible states: either alive or dead.  Every cell interacts with its eight neighbors which are the cells that are adjacent in any
direction (including the diagonal).  

The initial pattern constitutes the seed of the system.  The first generation is created by applying the following rules to every cell 
in the system.  Each generation should be a pure function of the preceding one.

The rules of Conway's Game of Life are as follows:
  1. Any live cell with fewer than two live neighbors dies, as if by underpopulation.
  2. Any live cell with two or three live neighbors lives on to the next generation.
  3. Any live cell with more than three live neighbors dies, as if by overpopulation.
  4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

## data structures

World (System)
  The entire world in which the game takes place.
Cell
  The singular atom of life and death.
Neighbor
  A cell that lives and dies up against another in any direction.
NeighborSet
  A set of all neighbors surrounding a given cell.
Generation
  The entire set of cells for each iteration of the game.

## functions

clear the screen

draw the screen

initialize (seed)

process the entire system per generation

tick length
