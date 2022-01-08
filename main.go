package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Sruthi's App")
	w.Resize(fyne.Size{
		Width:  500,
		Height: 500,
	})

	top := canvas.NewText("top bar", color.RGBA{
		R: 250,
		G: 20,
		B: 250,
		A: 0,
	})
	left := canvas.NewText("left", color.White)
	middle := canvas.NewText("content", color.White)

	content := container.New(layout.NewBorderLayout(top, nil, left, nil),
		top, left, middle)
	w.SetContent(content)
	w.ShowAndRun()
}
