package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	width         = 640
	height        = 480
	sandGrainSize = 10
)

var (
	grid [int64(width / sandGrainSize)][int64(height / sandGrainSize)]bool
)

type Game struct {
}

func (g *Game) Update() error {

	nextGrid := grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] {
				if j+1 < len(grid[i]) && !grid[i][j+1] { //if there is not a sand grain or the floor below, move the grain there
					nextGrid[i][j] = false
					nextGrid[i][j+1] = true
				} else if j+1 < len(grid[i]) && grid[i][j+1] { //if there is a sand grain below and is not the floor move to a side
					leftEmpty := i-1 >= 0 && !grid[i-1][j+1]
					rightEmpty := i+1 < len(grid) && !grid[i+1][j+1]

					if leftEmpty && rightEmpty {
						if rand.Intn(2) == 1 { //random left or right
							leftEmpty = false
						} else {
							rightEmpty = false
						}
					}

					if leftEmpty {
						nextGrid[i][j] = false
						nextGrid[i-1][j+1] = true
					} else if rightEmpty {
						nextGrid[i][j] = false
						nextGrid[i+1][j+1] = true
					}
				}
			}
		}
	}
	grid = nextGrid

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		grid[int64(x/sandGrainSize)][int64(y/sandGrainSize)] = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] {
				vector.DrawFilledRect(screen, float32(i*sandGrainSize), float32(j*sandGrainSize), sandGrainSize, sandGrainSize, color.White, false)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {

	grid[int64((width/sandGrainSize)/2)][int64((height/sandGrainSize)/2)] = true

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Falling Sand")
	if err := ebiten.RunGame(&Game{}); err != nil {
		fmt.Println(err)
	}
}
