package main

import (
	"math"
	"time"

	termbox "github.com/nsf/termbox-go"
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
	drawCircle(centerX, centerY, axisX+1, axisY+1, ' ', termbox.ColorWhite, termbox.ColorWhite)
	drawHand(t.Second(), centerX, centerY, axisX, axisY, 360/60, ' ', termbox.ColorWhite, termbox.ColorGreen)
	drawHand(t.Minute(), centerX, centerY, axisX, axisY, 360/60, ' ', termbox.ColorWhite, termbox.ColorCyan)
	drawHand(t.Hour(), centerX, centerY, axisX-2, axisY-2, 360/12, ' ', termbox.ColorWhite, termbox.ColorBlue)
	termbox.SetCell(centerX, centerY, ' ', termbox.ColorWhite, termbox.ColorWhite)
	termbox.Flush()
}

func drawCircle(cx, cy, h, v int, ch rune, fg termbox.Attribute, bg termbox.Attribute) {
	for i := 0; i <= 90; i++ {
		x := float64(h) * math.Cos(float64(i)*3.14/180)
		y := float64(v) * math.Sin(float64(i)*3.14/180)
		termbox.SetCell(cx+int(x), cy-int(y), ch, fg, bg) // 0 to 90 degrees
		termbox.SetCell(cx+int(x), cy+int(y), ch, fg, bg) // 90 to 180 degrees
		termbox.SetCell(cx-int(x), cy+int(y), ch, fg, bg) // 180 to 270 degrees
		termbox.SetCell(cx-int(x), cy-int(y), ch, fg, bg) // 270 to 360 degrees
	}
}

func drawLine(x1, y1, x2, y2 int, ch rune, fg termbox.Attribute, bg termbox.Attribute) {

	var dx, dy, sx, sy int

	if x2 > x1 {
		sx = 1
	} else {
		sx = -1
	}
	if x2 > x1 {
		dx = x2 - x1
	} else {
		dx = x1 - x2
	}
	if y2 > y1 {
		sy = 1
	} else {
		sy = -1
	}
	if y2 > y1 {
		dy = y2 - y1
	} else {
		dy = y1 - y2
	}

	x := x1
	y := y1

	if dx >= dy {
		e := -dx
		for i := 0; i <= dx; i++ {
			termbox.SetCell(x, y, ch, fg, bg)
			x += sx
			e += 2 * dy
			if e >= 0 {
				y += sy
				e -= 2 * dx
			}
		}

	} else {
		e := -dy
		for i := 0; i <= dy; i++ {
			termbox.SetCell(x, y, ch, fg, bg)
			y += sy
			e += 2 * dx
			if e >= 0 {
				x += sx
				e -= 2 * dy
			}
		}
	}
}

func drawHand(time, cx, cy, h, v, dg int, ch rune, fg termbox.Attribute, bg termbox.Attribute) {

	x := float64(h) * math.Cos(float64(dg*time-90)*3.14/180)
	y := float64(v) * math.Sin(float64(dg*time-90)*3.14/180)

	drawLine(cx, cy, cx+int(x), cy+int(y), ch, fg, bg)
}
