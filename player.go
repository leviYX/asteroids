package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-asteroids/assets"
)

type Player struct {
	sprite *ebiten.Image
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite
	p := &Player{
		sprite: sprite,
	}
	return p
}

func (p *Player) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	// 1. 取屏幕尺寸
	// 方式 A：如果 Game 结构体里存了 ScreenWidth/ScreenHeight，直接用
	// sw, sh := game.ScreenWidth, game.ScreenHeight
	// 方式 B：直接查窗口大小
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()

	// 2. 取图片尺寸
	bw, bh := p.sprite.Bounds().Dx(), p.sprite.Bounds().Dy()

	// 3. 计算左上角坐标
	x := float64(sw-bw) / 2 // 水平居中
	y := float64(sh - bh)   // 贴底
	options.GeoM.Translate(x, y)
	screen.DrawImage(p.sprite, options)
}

func (p *Player) Update() {

}
