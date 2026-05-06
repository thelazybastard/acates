package main

import (
	"net/url"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	acatesApp := app.New()
	acatesWindow := acatesApp.NewWindow("Acates")

	recipesPage := container.NewVBox(
		widget.NewLabel("HEllo world"),
	)

	trackerPage := container.NewVBox()
	trackerPage.Add(
		widget.NewButton("Add", func() {
			entry := widget.NewEntry()
			entry.SetPlaceHolder("Enter Ingredient here....")

			addedIngredient := []*widget.FormItem{
				widget.NewFormItem("Item", entry),
			}

			dialog.ShowForm("Add ingredient", "Add", "Cancel", addedIngredient, func(ok bool) {
				if ok && entry.Text != "" {

					contents := container.NewHBox()

					label := widget.NewLabel(entry.Text)
					button := widget.NewButton("Delete", func() {
						contents.Hide()
					})

					contents.Add(label)
					contents.Add(button)
					trackerPage.Add(contents)

					trackerPage.Refresh()
				}
			}, acatesWindow)
		}),
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
