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
	grid      [int64(width / sandGrainSize)][int64(height / sandGrainSize)]bool
	colorGrid [int64(width / sandGrainSize)][int64(height / sandGrainSize)]color.RGBA
	sandColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
)

type Game struct {
	iterationCount int
}

func rainbowColorForIteration(iteration int) color.RGBA {
	// Calculate RGB values for the given iteration
	r := uint8(255*(1+iteration)/510) % 255
	g := uint8(255*(1+2*iteration)/510) % 255
	b := uint8(255*(1+3*iteration)/510) % 255

	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func (g *Game) Update() error {

	nextGrid := grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] {
				if j+1 < len(grid[i]) && !grid[i][j+1] { //if there is not a sand grain or the floor below, move the grain there
					nextGrid[i][j] = false
					nextGrid[i][j+1] = true
					colorGrid[i][j+1] = sandColor
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
						colorGrid[i-1][j+1] = sandColor
					} else if rightEmpty {
						nextGrid[i][j] = false
						nextGrid[i+1][j+1] = true
						colorGrid[i+1][j+1] = sandColor
					}
				}
			}
		}
	}
	grid = nextGrid

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) { //we place a grain of sand if mouse is clicked
		x, y := ebiten.CursorPosition()
		grid[int64(x/sandGrainSize)][int64(y/sandGrainSize)] = true
	}

	// sandColor = color.RGBA{R: 255, G: 165, B: 0, A: 255}

	sandColor = rainbowColorForIteration(g.iterationCount)
	g.iterationCount++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] {
				vector.DrawFilledRect(screen, float32(i*sandGrainSize), float32(j*sandGrainSize), sandGrainSize, sandGrainSize, colorGrid[i][j], false)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {

	grid[int64((width/sandGrainSize)/2)][int64((height/sandGrainSize)/2)] = true
	// colorGrid [int64(width / sandGrainSize)][int64(height / sandGrainSize)] = color.RGBA{R: 255, G: 165, B: 0, A: 255}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Falling Sand")
	if err := ebiten.RunGame(&Game{}); err != nil {
		fmt.Println(err)
	}
}
