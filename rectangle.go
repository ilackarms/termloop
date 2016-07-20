package termloop

import "github.com/pborman/uuid"

// A type representing a 2D rectangle, with position, size and color.
type Rectangle struct {
	X      int `json:"X"`
	Y      int `json:"Y"`
	Width  int `json:"Width"`
	Height int `json:"Height"`
	Color  Attr `json:"Color"`
	UUID   string `json:"UUID"`
	Ch     rune    `json:"Ch"`
}

// NewRectangle creates a new Rectangle at position (x, y), with size
// (width, height) and color color.
// Returns a pointer to the new Rectangle.
func NewRectangle(x, y, w, h int, color Attr) *Rectangle {
	r := Rectangle{X: x, Y: y, Width: w, Height: h, Color: color, UUID: uuid.New(), Ch: ' '}
	return &r
}

// Returns the UUID for a drawable
func (r *Rectangle) GetUUID() string {
	return r.UUID
}

// Draws the Rectangle r onto Screen s.
func (r *Rectangle) Draw(s *Screen) {
	for i := 0; i < r.Width; i++ {
		for j := 0; j < r.Height; j++ {
			s.RenderCell(r.X +i, r.Y +j, &Cell{Bg: r.Color, Ch: r.Ch})
		}
	}
}

func (r *Rectangle) Tick(ev Event) {}

// Size returns the width and height in characters of the Rectangle.
func (r *Rectangle) Size() (int, int) {
	return r.Width, r.Height
}

// Position returns the x and y coordinates of the Rectangle.
func (r *Rectangle) Position() (int, int) {
	return r.X, r.Y
}

// SetPosition sets the coordinates of the Rectangle to be x and y.
func (r *Rectangle) SetPosition(x, y int) {
	r.X = x
	r.Y = y
}

// SetSize sets the width and height of the Rectangle to be w and h.
func (r *Rectangle) SetSize(w, h int) {
	r.Width = w
	r.Height = h
}

// Color returns the color of the Rectangle.
func (r *Rectangle) GetColor() Attr {
	return r.Color
}

// SetColor sets the color of the Rectangle.
func (r *Rectangle) SetColor(color Attr) {
	r.Color = color
}
