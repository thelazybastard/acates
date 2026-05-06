package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	acatesApp := app.New()
	acatesWindow := acatesApp.NewWindow("Acates")

	recipesPage := container.NewVBox(
		widget.NewLabel("HEllo world"),
	)

	trackerPage := container.NewVBox(
		widget.NewLabel("Hello world"),
	)

	suggestionsPage := container.NewVBox(
		widget.NewLabel("Hello world"),
	)

	aboutPage := container.NewVBox(
		widget.NewLabel("HEllo World"),
	)

	root := container.NewAppTabs(
		container.NewTabItem("Home", recipesPage),
		container.NewTabItem("Tracker", trackerPage),
		container.NewTabItem("Suggestions", suggestionsPage),
		container.NewTabItem("Settings", aboutPage),
	)
	root.SetTabLocation(container.TabLocationLeading)

	acatesWindow.SetContent(root)
	acatesWindow.ShowAndRun()
}
