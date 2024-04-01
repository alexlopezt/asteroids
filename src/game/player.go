package game

import (
	"math"

	"github.com/alexlopezt/asteroids/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	sprite         *ebiten.Image
	playerPosition Point
	rotation       float64
}

func NewPlayer() *Player {

	sprite := assets.PlayerSprite
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Point{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
	}
	return &Player{
		playerPosition: pos,
		sprite:         sprite,
	}
}

func (p *Player) Update() {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}
}
func (p *Player) UpdateNotUsed() {
	speed := 5.0

	var delta Point

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		delta.Y = speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		delta.Y = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		delta.X = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		delta.X = speed
	}

	// Check for diagonal movement
	if delta.X != 0 && delta.Y != 0 {
		factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	p.playerPosition.X += delta.X
	p.playerPosition.Y += delta.Y

	// p.updateRotation()
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()

	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.playerPosition.X, p.playerPosition.Y)

	screen.DrawImage(assets.PlayerSprite, op)
}
