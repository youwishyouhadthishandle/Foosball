package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

const (
	playerSpeed   = 5
	playerWidth   = 26
	playerHeight  = 30
	boundaryWidth = 29
)

type team struct {
	goalKeeper [1]player
	defence    [2]player
	mid        [5]player
	attack     [3]player
}

type player struct {
	circle
	tex *sdl.Texture
}

func newteam(renderer *sdl.Renderer, teamid int32) (t team, err error) {
	offset := int32(0)
	if teamid == 2 {
		offset = boxWidth - 1 - 2*playerWidth
	}
	t.goalKeeper[0], err = newplayer(renderer, int32(math.Abs(float64(61-playerWidth-offset))), boxHeight/2-playerHeight, teamid)
	if err != nil {
		fmt.Println(err)
		return team{}, fmt.Errorf("%v", err)
	}
	for i := range t.defence {
		t.defence[i], err = newplayer(renderer, int32(math.Abs(float64(136-playerWidth-offset))), int32(boxHeight*(i+1))/3-playerHeight, teamid)
		if err != nil {
			fmt.Println(err)
			return team{}, fmt.Errorf("%v", err)
		}
	}
	for i := range t.mid {
		t.mid[i], err = newplayer(renderer, int32(math.Abs(float64(286-playerWidth-offset))), int32(boxHeight*(i+1))/6-playerHeight, teamid)
		if err != nil {
			fmt.Println(err)
			return team{}, fmt.Errorf("%v", err)
		}
	}
	for i := range t.attack {
		t.attack[i], err = newplayer(renderer, int32(math.Abs(float64(436-playerWidth-offset))), int32(boxHeight*(i+1))/4-playerHeight, teamid)
		if err != nil {
			fmt.Println(err)
			return team{}, fmt.Errorf("%v", err)
		}
	}
	return t, nil
}

func newplayer(renderer *sdl.Renderer, x, y, teamid int32) (p player, err error) {

	img.Init(img.INIT_JPG | img.INIT_PNG)
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
	playerImg, err := img.Load("Player_Red.png")
	if teamid == 2 {
		playerImg, err = img.Load("Player_Blue.png")
	}
	if err != nil {
		fmt.Println(err)
		return player{}, fmt.Errorf("%v", err)
	}
	defer playerImg.Free()
	p.tex, err = renderer.CreateTextureFromSurface(playerImg)
	if err != nil {
		fmt.Println(err)
		return player{}, fmt.Errorf("%v", err)
	}

	p.x = float64(x + playerWidth/2)
	p.y = float64(y + playerHeight/2)
	p.radius = 13
	return p, nil
}

func playerDraw(p *player, renderer *sdl.Renderer) *sdl.Renderer {
	renderer.Copy(p.tex,
		&sdl.Rect{0, 0, playerWidth, playerHeight},
		&sdl.Rect{int32(p.x), int32(p.y), playerWidth, playerHeight})
	return renderer
}

func (t *team) draw(renderer *sdl.Renderer) {
	playerDraw(&t.goalKeeper[0], renderer)
	for i := range t.defence {
		playerDraw(&t.defence[i], renderer)
	}
	for i := range t.mid {
		playerDraw(&t.mid[i], renderer)
	}
	for i := range t.attack {
		playerDraw(&t.attack[i], renderer)
	}
}

func (t *team) update(lastStick []player, lastMotion int32) ([]player, int32) {
	keys := sdl.GetKeyboardState()
	var stick1 = t.goalKeeper[0:1]
	var stick2 = t.defence[0:2]
	var stick3 = t.mid[0:5]
	var stick4 = t.attack[0:3]
	if keys[sdl.SCANCODE_A] == 1 {
		lastStick = stick1
	} else if keys[sdl.SCANCODE_S] == 1 {
		lastStick = stick2
	} else if keys[sdl.SCANCODE_D] == 1 {
		lastStick = stick3
	} else if keys[sdl.SCANCODE_F] == 1 {
		lastStick = stick4
	}
	if keys[sdl.SCANCODE_UP] == 1 {
		lastMotion = 1
		if lastStick[0].y > boundaryWidth {
			for i := range lastStick {
				if lastStick[i].y > boundaryWidth {
					lastStick[i].y -= playerSpeed
				}
			}
		}
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		lastMotion = -1
		if lastStick[len(lastStick)-1].y < boxHeight-boundaryWidth-playerHeight-1 {
			for i := range lastStick {
				if lastStick[i].y < boxHeight-playerHeight-boundaryWidth-1 {
					lastStick[i].y += playerSpeed
				}
			}
		}
	}
	return lastStick, lastMotion
}
