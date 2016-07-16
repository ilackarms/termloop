package termloop

import "github.com/pborman/uuid"

// Provides a general Drawable to be rendered.
type Entity struct {
	Canvas Canvas `json:"Canvas"`
	X      int `json:"X"`
	Y      int `json:"Y"`
	Width  int `json:"Width"`
	Height int `json:"Height"`
	UUID   string `json:"UUID"`
}

// NewEntity creates a new Entity, with position (x, y) and size
// (width, height).
// Returns a pointer to the new Entity.
func NewEntity(x, y, width, height int) *Entity {
	canvas := NewCanvas(width, height)
	e := Entity{X: x, Y: y, Width: width, Height: height,
		Canvas: canvas, UUID: uuid.New()}
	return &e
}

// NewEntityFromCanvas returns a pointer to a new Entity, with
// position (x, y) and Canvas c. Width and height are calculated
// using the Canvas.
func NewEntityFromCanvas(x, y int, c Canvas) *Entity {
	e := Entity{
		X:      x,
		Y:      y,
		Canvas: c,
		Width:  len(c),
		Height: len(c[0]),
		UUID: uuid.New(),
	}
	return &e
}

// Returns the UUID for a drawable
func (e *Entity) GetUUID() string {
	return e.UUID
}

// Draw draws the entity to its current position on the screen.
// This is usually called every frame.
func (e *Entity) Draw(s *Screen) {
	for i := 0; i < e.Width; i++ {
		for j := 0; j < e.Height; j++ {
			s.RenderCell(e.X +i, e.Y +j, &e.Canvas[i][j])
		}
	}
}

func (e *Entity) Tick(ev Event) {}

// Position returns the (x, y) coordinates of the Entity.
func (e *Entity) Position() (int, int) {
	return e.X, e.Y
}

// Size returns the width and height of the entity, in characters.
func (e *Entity) Size() (int, int) {
	return e.Width, e.Height
}

// SetPosition sets the x and y coordinates of the Entity.
func (e *Entity) SetPosition(x, y int) {
	e.X = x
	e.Y = y
}

// SetCell updates the attribute of the Cell at x, y to match those of c.
// The coordinates are relative to the entity itself, not the Screen.
func (e *Entity) SetCell(x, y int, c *Cell) {
	renderCell(&e.Canvas[x][y], c)
}

// Fill fills the canvas of the Entity with
// a Cell c.
func (e *Entity) Fill(c *Cell) {
	for i := range e.Canvas {
		for j := range e.Canvas[i] {
			renderCell(&e.Canvas[i][j], c)
		}
	}
}

// ApplyCanvas takes a pointer to a Canvas, c, and applies this canvas
// over the top of the Entity's canvas. Any new values in c will overwrite
// those in the entity.
func (e *Entity) ApplyCanvas(c *Canvas) {
	for i := 0; i < min(len(e.Canvas), len(*c)); i++ {
		for j := 0; j < min(len(e.Canvas[0]), len((*c)[0])); j++ {
			renderCell(&e.Canvas[i][j], &(*c)[i][j])
		}
	}
}
