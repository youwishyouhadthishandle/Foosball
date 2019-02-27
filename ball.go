package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	radius = 8
)

type ball struct {
	circle
	tex    *sdl.Texture
	xv, yv float64
}

func (ball *ball) draw(renderer *sdl.Renderer) *sdl.Renderer {
	renderer.Copy(ball.tex,
		&sdl.Rect{0, 0, 2 * radius, 2 * radius},
		&sdl.Rect{int32(ball.x), int32(ball.y), 2 * radius, 2 * radius})
	return renderer
}

func newBall(renderer *sdl.Renderer, x, y int32) (bal ball, err error) {

	img.Init(img.INIT_JPG | img.INIT_PNG)
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
	BallImg, err := img.Load("Ball.png")
	if err != nil {
		fmt.Println(err)
		return ball{}, fmt.Errorf("%v", err)
	}
	defer BallImg.Free()
	bal.tex, err = renderer.CreateTextureFromSurface(BallImg)
	if err != nil {
		fmt.Println(err)
		return ball{}, fmt.Errorf("%v", err)
	}

	bal.x = float64(x)
	bal.y = float64(y)
	bal.radius = radius
	bal.xv = 3
	bal.yv = 1
	return bal, nil
}

func (ball *ball) update() {
	ball.x += ball.xv
	ball.y += ball.yv

	goal, index := ball.collidesWall()
	if index != -1 {
		onCollisionWithWall(ball, index)
	} else {
		fmt.Println(goal)
	}
}
