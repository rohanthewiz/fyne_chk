package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Entry Widget")
	w.Resize(fyne.Size{Width: 500, Height: 500})

	// Entry
	lbl := widget.NewLabel("Text Entry")
	in := widget.NewEntry() // Text input
	in.SetPlaceHolder("Enter two decimal numbers...")

	result := widget.NewEntry()

	saveBtn := widget.NewButton("Save",
		func() {
			// Parse entry
			var a, b float64
			_, err := fmt.Sscanf(in.Text, "%f %f", &a, &b)
			if err != nil {
				log.Println("Error reading the input", err.Error())
				return
			}
			// Return the result
			sum := a + b
			result.SetText("> " + fmt.Sprintf("%0.1f", sum))
		})

	con := container.NewVBox(
		lbl, in, widget.NewSeparator(),
		layout.NewSpacer(),
		saveBtn, widget.NewSeparator(), result)

	w.SetContent(con)
	w.ShowAndRun()
}
