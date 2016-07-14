package main

import (
	tl "github.com/ilackarms/termloop"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	g := tl.NewGame()
	dat, err := ioutil.ReadFile("lorry.txt")
	check(err)
	e := tl.NewEntityFromCanvas(1, 1, tl.CanvasFromString(string(dat)))
	g.Screen().AddEntity(e)
	g.Start()
}
