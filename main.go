package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello!")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.Show()

	w2 := a.NewWindow("Clock!")

	clock := widget.NewLabel("")
	updateTime(clock)
	w2.SetContent(clock)
	w2.Resize(fyne.NewSize(300, 100))
	w2.Show()

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	a.Run()
	//w.ShowAndRun()

}

/*w2.SetContent(widget.NewButton("Open new", func() {
	w3 := a.NewWindow("Third")
	w3.SetContent(widget.NewLabel("Third"))
	w3.Show()
}))
button to open another window

*/
//app.Quit closes the app
//can put coe here to run after the app is closed(?)
//app.Run() runs the windo, window.Show() opens another window
//window.SetMaster(), defines the main window, if it closes all windows are closed too
//create button = window.SeteContent(widget.NewButton("label on the button", func(){}))

/* get login and password for deleteFromAnlix(?)
"fyne.io/fyne/v2/container"
func makeUI() (*widget.Label, *widget.Entry) {
	return widget.NewLabel("Hello world!"),
		widget.NewEntry()
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello Person")

	w.SetContent(container.NewVBox(makeUI()))
	w.ShowAndRun()
}
*/
