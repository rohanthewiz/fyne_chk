package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Entry Widget")

	w.Resize(fyne.Size{
		Width:  500,
		Height: 500,
	})

	lbl := widget.NewLabel("Text Entry")
	in := widget.NewEntry()
	in.SetPlaceHolder("Enter text...")
	saveBtn := widget.NewButton("Save",
		func() {println("You entered:", in.Text)})

	content := container.NewVBox(lbl, in, saveBtn)

	w.SetContent(content)
	w.ShowAndRun()
}
