package main

import (
	"github.com/sago35/koebiten"
	"github.com/sago35/koebiten/hardware"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *koebiten.Image) {
	koebiten.Println("Hello world")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 128, 64
}

func main() {
	koebiten.SetHardware(hardware.Device)

	game := &Game{}
	koebiten.RunGame(game)
}
