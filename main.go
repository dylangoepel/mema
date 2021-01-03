package main

import (
    "github.com/gotk3/gotk3/gtk"
    "github.com/gotk3/gotk3/gdk"
    "github.com/gotk3/gotk3/glib"
    "os"
    "strconv"
    "log"
    "math/rand"
    "time"
)

var mainLabel *gtk.Label
var speedLabel *gtk.Label
var countLabel *gtk.Label
var win *gtk.Window

var t int = 0

const PrepareDelay uint = 500
var Interval uint = 1000
var Count uint = 3

var CurrentChallenge []int = nil
var Handles []glib.SourceHandle = nil

func min(i1, i2 int) int {
    if i1 < i2 {
        return i1
    } else {
        return i2
    }
}

func presentMainText(text string) {
    width, height := win.GetSize()
    height = height * 4 / 5
    fontSize := min(width, height) * 500
    mainLabel.SetMarkup("<span size=\"" + strconv.Itoa(fontSize) + "\" font_family=\"monospace\">" + text + "</span>")
}
func presentCurrentConfiguration() {
    width, height := win.GetSize()
    height = height / 5
    width = width / 2
    fontSize := min(width, height) * 250
    speedLabel.SetMarkup("<span size=\"" + strconv.Itoa(fontSize) + "\" font_family=\"monospace\">" + strconv.Itoa(int(Interval)) + "</span>")
    countLabel.SetMarkup("<span size=\"" + strconv.Itoa(fontSize) + "\" font_family=\"monospace\">" + strconv.Itoa(int(Count)) + "</span>")
}

func showNumber(num int) {
    presentMainText("")
    glib.TimeoutAdd(100, presentMainText, strconv.Itoa(num))
}

func generateChallenge() []int {
    rand.Seed(time.Now().UnixNano())
    numbers := make([]int, Count)
    for i := uint(0); i < Count; i++ {
        numbers[i] = int(uint(rand.Uint32()) % 50 + 1)
    }
    return numbers
}

func cleanupChallengePresentation() {
    presentMainText("")
    Handles = nil
}

func presentChallenge(numbers []int) {
    Handles = make([]glib.SourceHandle, len(numbers) + 1)
    presentMainText("")
    for i, n := range numbers {
        Handles[i], _ = glib.TimeoutAdd(PrepareDelay + Interval * uint(i), showNumber, n)
    }
    Handles[len(numbers)], _ = glib.TimeoutAdd(PrepareDelay + Interval * uint(len(numbers)), cleanupChallengePresentation)
}

func presentSolution(numbers []int) {
    sum := 0
    for _, n := range numbers {
        sum += n
    }

    presentMainText("<u>" + strconv.Itoa(sum) + "</u>")
}

func keyPress(w *gtk.Window, k *gdk.Event) {
    eventKey := gdk.EventKeyNewFromEvent(k)
    if eventKey.KeyVal() == gdk.KEY_space {
        if  CurrentChallenge == nil {
            CurrentChallenge = generateChallenge()
            presentChallenge(CurrentChallenge)
        } else if Handles == nil {
            presentSolution(CurrentChallenge)
            CurrentChallenge = nil
        }
    } else if eventKey.KeyVal() == gdk.KEY_plus {
        Interval += 100
    } else if eventKey.KeyVal() == gdk.KEY_minus && Interval > 200 {
        Interval -= 100
    } else if eventKey.KeyVal() == gdk.KEY_Up {
        Count++
    } else if eventKey.KeyVal() == gdk.KEY_Down && Count > 2 {
        Count--
    }
    presentCurrentConfiguration()
}


func main() {
    gtk.Init(&os.Args)

    var err error
    win, err = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
    if err != nil {
        log.Fatal("Unabel to create window:", err)
    }

    win.SetTitle("mema")
    win.Connect("destroy", func() {
        gtk.MainQuit()
    })
    win.Connect("key_press_event", keyPress)
    win.Connect("check-resize", func() {
        presentMainText(mainLabel.GetLabel())
        presentCurrentConfiguration()
    })

    mainLabel, err = gtk.LabelNew("")
    if err != nil {
        log.Fatal("Unable to create main label:", err)
    }
    mainLabel.SetJustify(gtk.JUSTIFY_CENTER)

    speedLabel, err = gtk.LabelNew("")
    if err != nil {
        log.Fatal("Unable to create speed label:", err)
    }
    countLabel, err = gtk.LabelNew("")
    if err != nil {
        log.Fatal("Unable to create count label:", err)
    }

    grid, err := gtk.GridNew()
    grid.SetRowHomogeneous(true)
    grid.SetColumnHomogeneous(true)
    grid.Attach(gtk.IWidget(mainLabel), 0, 0, 2, 4)
    grid.Attach(gtk.IWidget(speedLabel), 0, 4, 1, 1)
    grid.Attach(gtk.IWidget(countLabel), 1, 4, 1, 1)

    win.Add(grid)

    win.ShowAll()

    presentCurrentConfiguration()

    gtk.Main()
}
