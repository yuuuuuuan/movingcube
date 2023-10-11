package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	"log"
)

const (
	screenWidth  = 600
	screenHeight = 500
	cubeWidth    = 50
	cubeHeight   = 50
	stoneWidth   = 100
	stoneHeight  = 100
	movingSpeed  = 2
)

type Input struct {
	msg string
}

type Cube struct {
	width  int
	height int
	x      float64
	y      float64
	img    *ebiten.Image
}

func (i *Input) Update(cube *Cube, stone *Cube) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && CubeCollision(cube, stone, "left") {
		cube.x -= movingSpeed
		i.msg = "left pressed\n"
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) && CubeCollision(cube, stone, "right") {
		cube.x += movingSpeed
		i.msg = "right pressed\n"
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) && CubeCollision(cube, stone, "up") {
		cube.y -= movingSpeed
		i.msg = "up pressed\n"
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) && CubeCollision(cube, stone, "down") {
		cube.y += movingSpeed
		i.msg = "down pressed\n"
	} else {
		i.msg = "none\n"

	}
}

func NewCube() *Cube {
	rect := image.Rect(0, 0, cubeWidth, cubeHeight)
	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(color.White)
	cube := &Cube{
		cubeWidth,
		cubeHeight,
		20,
		20,
		img,
	}
	return cube
}

func (cube *Cube) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(cube.x, cube.y)
	screen.DrawImage(cube.img, op)
}

func NewStone() *Cube {
	rect := image.Rect(0, 0, stoneWidth, stoneHeight)
	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(color.White)
	cube := &Cube{
		stoneWidth,
		stoneHeight,
		screenWidth / 2,
		stoneHeight / 2,
		img,
	}
	return cube
}

func CubeCollision(cube1, cube2 *Cube, stonedir string) bool {
	cube1x := cube1.x
	cube1y := cube1.y
	cube2x := cube2.x
	cube2y := cube2.y
	switch stonedir {
	case "up":
		if (cube1y-cubeHeight) > cube2y || (cube1x+cubeWidth) < cube2x || cube1x > (cube2x+stoneWidth) {
			return true
		}
	case "down":
		if cube1y < (cube2y-stoneHeight) || (cube1x+cubeWidth) < cube2x || cube1x > (cube2x+stoneWidth) {
			return true
		}
	case "left":
		if (cube1x+cubeWidth) < cube2x || cube1y < (cube2y-stoneHeight) || (cube1y-cubeHeight) > cube2y {
			return true
		}
	case "right":
		if cube1x > (cube1x-stoneWidth) || cube1y < (cube2y-stoneHeight) || (cube1y-cubeHeight) > cube2y {
			return true
		}
	}
	return false

}

type Game struct {
	input *Input
	cube  *Cube
	stone *Cube
}

func (g *Game) Update() error {
	g.input.Update(g.cube, g.stone)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(screen, g.input.msg)
	g.cube.Draw(screen)
	g.stone.Draw(screen)
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() *Game {
	// 相同的代码省略...
	return &Game{
		input: &Input{""},
		cube:  NewCube(),
		stone: NewStone(),
	}
}

func main() {

	//游戏显示背景（全部的可视范围）
	ebiten.SetWindowSize(960, 640)
	//title标题设置
	ebiten.SetWindowTitle("这是一个游戏")
	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
