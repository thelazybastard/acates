package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	acatesApp := app.New()
	acatesWindow := acatesApp.NewWindow("Acates")

	homePage := container.NewVBox(
		widget.NewLabel("HEllo world"),
	)

	settingsPage := container.NewVBox(
		widget.NewLabel("HEllo World"),
	)

	root := container.NewAppTabs(
		container.NewTabItem("Home", homePage),
		container.NewTabItem("Settings", settingsPage),
	)
	root.SetTabLocation(container.TabLocationLeading)

	acatesWindow.SetContent(root)
	acatesWindow.ShowAndRun()
}
