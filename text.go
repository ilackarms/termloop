package termloop

// Text represents a string that can be drawn to the screen.
type Text struct {
	X      int `json:"X"`
	Y      int `json:"Y"`
	Fg     Attr `json:"Fg"`
	Bg     Attr `json:"Bg"`
	Text   []rune `json:"Text"`
	Canvas []Cell `json:"Canvas"`
}

// NewText creates a new Text, at position (x, y). It sets the Text's
// background and foreground colors to fg and bg respectively, and sets the
// Text's text to be text.
// Returns a pointer to the new Text.
func NewText(x, y int, text string, fg, bg Attr) *Text {
	str := []rune(text)
	c := make([]Cell, len(str))
	for i := range c {
		c[i] = Cell{Ch: str[i], Fg: fg, Bg: bg}
	}
	return &Text{
		X:      x,
		Y:      y,
		Fg:     fg,
		Bg:     bg,
		Text:   str,
		Canvas: c,
	}
}

func (t *Text) Tick(ev Event) {}

// Draw draws the Text to the Screen s.
func (t *Text) Draw(s *Screen) {
	w, _ := t.Size()
	for i := 0; i < w; i++ {
		s.RenderCell(t.X +i, t.Y, &t.Canvas[i])
	}
}

// Position returns the (x, y) coordinates of the Text.
func (t *Text) Position() (int, int) {
	return t.X, t.Y
}

// Size returns the width and height of the Text.
func (t *Text) Size() (int, int) {
	return len(t.Text), 1
}

// SetPosition sets the coordinates of the Text to be (x, y).
func (t *Text) SetPosition(x, y int) {
	t.X = x
	t.Y = y
}

// Text returns the text of the Text.
func (t *Text) GetText() string {
	return string(t.Text)
}

// SetText sets the text of the Text to be text.
func (t *Text) SetText(text string) {
	t.Text = []rune(text)
	c := make([]Cell, len(t.Text))
	for i := range c {
		c[i] = Cell{Ch: t.Text[i], Fg: t.Fg, Bg: t.Bg}
	}
	t.Canvas = c
}

// Color returns the (foreground, background) colors of the Text.
func (t *Text) Color() (Attr, Attr) {
	return t.Fg, t.Bg
}

// SetColor sets the (foreground, background) colors of the Text
// to fg, bg respectively.
func (t *Text) SetColor(fg, bg Attr) {
	t.Fg = fg
	t.Bg = bg
	for i := range t.Canvas {
		t.Canvas[i].Fg = fg
		t.Canvas[i].Bg = bg
	}
}
