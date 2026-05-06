package main

import (
	"net/url"

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
		widget.NewCard(
			"Acates",
			"Version 1.0.0",
			container.NewVBox(
				widget.NewLabel("Author: Monish H. Giani"),
				container.NewHBox(
					widget.NewLabel("License:"),
					widget.NewHyperlink("GNU General Public Licence v3.0", parseUrl("https://github.com/thelazybastard/acates/blob/main/LICENSE")),
				),
				widget.NewHyperlink("Source Code", parseUrl("https://github.com/thelazybastard/acates")),
			),
		),
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

func parseUrl(website string) *url.URL {
	u, _ := url.Parse(website)
	return u
}
