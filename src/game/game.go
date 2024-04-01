package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/alexlopezt/asteroids/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600

	meteorSpawnTime = 1 * time.Second
)

var rnd *rand.Rand = rand.New(rand.NewSource(time.Now().UnixMilli()))

type Game struct {
	player           *Player
	meteorSpawnTimer *Timer
	meteors          []*Meteor

	// bullets          []*Bullet

	score int

	baseVelocity  float64
	velocityTimer *Timer
}

func init() {

}
func NewGame() *Game {
	return &Game{
		player:           NewPlayer(),
		meteorSpawnTimer: NewTimer(meteorSpawnTime),
		velocityTimer:    NewTimer(meteorSpeedUpTime),
		baseVelocity:     baseMeteorVelocity,
	}
}
func (g *Game) Update() error {
	g.velocityTimer.Update()
	if g.velocityTimer.IsReady() {
		g.velocityTimer.Reset()
		g.baseVelocity += meteorSpeedUpAmount
	}

	g.player.Update()

	g.meteorSpawnTimer.Update()

	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		g.meteors = append(g.meteors, NewMeteor(g.baseVelocity))
		for _, m := range g.meteors {
			m.Update()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	drawSoil(screen)

	g.player.Draw(screen)

	// for _, m := range g.meteors {
	// 	m.Draw(screen)
	// }

	text.Draw(screen, fmt.Sprintf("%06d", g.score), assets.ScoreFont, ScreenWidth/2-100, 50, color.White)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func drawSoil(screen *ebiten.Image) {
	sprite := assets.Soil
	bounds := sprite.Bounds()
	i := 0
	for {
		y := float64(bounds.Dy() * i)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(20, y)
		screen.DrawImage(assets.Soil, op)
		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(ScreenWidth-20.0-float64(sprite.Bounds().Dx()), y)
		screen.DrawImage(sprite, op2)

		if y > ScreenHeight {
			break
		}
		i++
	}
}
