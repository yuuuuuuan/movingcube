package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	"log"
)

type dir string

const (
	screenWidth  = 600
	screenHeight = 500
	cubeWidth    = 50
	cubeHeight   = 50
	stoneWidth   = 100
	stoneHeight  = 100
	movingSpeed  = 1
)

const (
	keyboardnone = iota
	keyboardup
	keyboarddown
	keyboardleft
	keyboardright
)

type Input struct {
	keyboardstate int
	msg           string
}

type Cube struct {
	width  int
	height int
	x      int
	y      int
	img    *ebiten.Image
}

func (i *Input) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		i.keyboardstate = keyboardleft
		i.msg = "left pressed\n"
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		i.keyboardstate = keyboardright
		i.msg = "right pressed\n"
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		i.keyboardstate = keyboardup
		i.msg = "up pressed\n"
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		i.keyboardstate = keyboarddown
		i.msg = "down pressed\n"
	} else {
		i.keyboardstate = keyboardnone
		i.msg = "none\n"

	}
}

func (c *Cube) Cubemotivation(i *Input, s *Cube) {
	dir, _ := CubeCollision(c, s)
	switch i.keyboardstate {
	case keyboardup:
		if dir == "down" {
			return
		}
		c.y -= movingSpeed
	case keyboarddown:
		if dir == "up" {
			return
		}
		c.y += movingSpeed
	case keyboardleft:
		if dir == "right" {
			return
		}
		c.x -= movingSpeed
	case keyboardright:
		if dir == "left" {
			return
		}
		c.x += movingSpeed
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

func (c *Cube) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.x), float64(c.y))
	screen.DrawImage(c.img, op)
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

func CubeCollision(cube1, cube2 *Cube) (dir, bool) {
	ax := cube1.x
	ay := cube1.y
	aw := cubeWidth
	ah := cubeHeight
	bx := cube2.x
	by := cube2.y
	bw := stoneWidth
	bh := stoneHeight
	//left
	if (ax+aw) == bx && ((by < ay && ay < (by+bh)) || (by < (ay+ah) && (ay+ah) < (by+bh))) {
		return "left", true
	}
	//right
	if ax == (bx+bw) && ((by < ay && ay < (by+bh)) || (by < (ay+ah) && (ay+ah) < (by+bh))) {
		return "right", true
	}
	//up
	if (ay+ah) == by && ((bx < ax && ax < (bx+bw)) || (bx < (ax+aw) && (ax+aw) < (bx+bw))) {
		return "up", true
	}
	//down
	if ay == (by+bh) && ((bx < ax && ax < (bx+bw)) || (bx < (ax+aw) && (ax+aw) < (bx+bw))) {
		return "down", true
	}
	return "none", false
}

type Game struct {
	input *Input
	cube  *Cube
	stone *Cube
}

func (g *Game) Update() error {
	g.input.Update()
	g.cube.Cubemotivation(g.input, g.stone)
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
	return &Game{
		input: &Input{
			0,
			""},
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
