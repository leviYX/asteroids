package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player *Player
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}
	game.player = NewPlayer(game)

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("飞机大战七夕特供版")

	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
