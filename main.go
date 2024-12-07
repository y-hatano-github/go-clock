package main

import (
	"math"
	"time"

	termbox "github.com/nsf/termbox-go"
	c "github.com/y-hatano-github/coordin"
)

const centerX = 30
const centerY = 12
const axisX = 20
const axisY = 10

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetOutputMode(termbox.Output256)

	key := make(chan string)
	go keyEvent(key)

loop:
	for {
		select {
		case k := <-key:
			if k == "esc" {
				break loop
			}
		default:
		}
		drawClock()
	}
}

func keyEvent(key chan string) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				key <- "esc"
			case termbox.KeyCtrlC:
				key <- "esc"
			default:
				key <- string(ev.Ch)
			}
		}
	}
}

func drawClock() {
	t := time.Now()
	termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)

	h := func(time, cx, cy, h, v, dg int) c.Points {
		d := float64(dg*time - 90)
		x := float64(h) * math.Cos(float64(d)*3.14/180)
		y := float64(v) * math.Sin(float64(d)*3.14/180)

		return c.Line(c.Point{X: cx, Y: cy}, c.Point{X: cx + int(x), Y: cy + int(y)})
	}

	hm := c.Circled(centerX, centerY, axisX-2, axisY-1, 30) // HourMark

	hh := h(t.Hour()*5+int(t.Minute()/12), centerX, centerY, axisX-6, axisY-3, 360/60) // hours hand
	mh := h(t.Minute(), centerX, centerY, axisX, axisY, 360/60)                        // minutes hand
	sh := h(t.Second(), centerX, centerY, axisX, axisY, 360/60)                        // seconds hand
	f, _ := c.Circle(centerX, centerY, axisX, axisY)                                   // frame

	setCell(hm, termbox.ColorWhite, termbox.ColorRed)
	setCell(hh, termbox.ColorWhite, termbox.ColorBlue)
	setCell(mh, termbox.ColorWhite, termbox.ColorCyan)
	setCell(sh, termbox.ColorWhite, termbox.ColorGreen)
	setCell(f, termbox.ColorWhite, termbox.ColorWhite)
	setCell(c.Points{c.Point{X: centerX, Y: centerY}}, termbox.ColorWhite, termbox.ColorWhite)

	termbox.Flush()
}

func setCell(ps c.Points, fg termbox.Attribute, bg termbox.Attribute) {
	for _, p := range ps {
		termbox.SetCell(p.X, p.Y, ' ', fg, bg)
	}
}
