package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

// draw main function that sets up box sizes and spacing for numbers.
func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	w, h := termbox.Size()
	bWidth := w * 15 / 100
	bHeight := h / 3
	offY := (h - bHeight) / 2
	spacing := w * 1 / 10
	hourTen := spacing
	hourOne := bWidth + spacing
	colonSpace := bWidth*2 + spacing
	minTen := bWidth*3 + spacing
	minOne := bWidth*4 + spacing
	barray := [5]int{hourTen, hourOne, colonSpace, minTen, minOne}
	cTime(bWidth, bHeight, barray, offY)
	termbox.Flush()
}

// cTime gets the current time and parses each digit to the printing functions.
func cTime(bWidth int, bHeight int, place [5]int, offY int) {
	now := time.Now()
	hour, minute, _ := now.Clock()
	one := (hour / 10) % 10
	two := hour % 10
	three := (minute / 10) % 10
	four := minute % 10
	order := [5]int{one, two, -1, three, four}
	for x := 0; x < 5; x++ {
		switch order[x] {
		case -1:
			colon(bWidth, bHeight, place[x], offY)
		case 0:
			zero(bWidth, bHeight, place[x], offY)
		case 1:
			ones(bWidth, bHeight, place[x], offY)
		case 2:
			twos(bWidth, bHeight, place[x], offY)
		case 3:
			threes(bWidth, bHeight, place[x], offY)
		case 4:
			fours(bWidth, bHeight, place[x], offY)
		case 5:
			five(bWidth, bHeight, place[x], offY)
		case 6:
			six(bWidth, bHeight, place[x], offY)
		case 7:
			seven(bWidth, bHeight, place[x], offY)
		case 8:
			eight(bWidth, bHeight, place[x], offY)
		case 9:
			nine(bWidth, bHeight, place[x], offY)
		}
	}
}

// colon draws a colon in a WxH size box.
func colon(w int, h int, offX int, offY int) {
	for x := 0; x <= w/8; x++ {
		for y := 0; y <= h*1/8; y++ {
			termbox.SetCell((x + offX + w/2 - w/8), y+offY+h*1/8, ' ', termbox.ColorDefault,
				termbox.ColorWhite)
		}
	}
	for x := 0; x <= w/8; x++ {
		for y := 0; y <= h*1/8; y++ {
			termbox.SetCell((x + offX + w/2 - w/8), y+offY+h*6/8, ' ', termbox.ColorDefault,
				termbox.ColorWhite)
		}
	}
}

// zero draws a zero in a WxH size box.
func zero(w int, h int, offX int, offY int) {
	topcross(w, h, offX, offY)
	leftupperdown(w, h, offX, offY)
	rightupperdown(w, h, offX, offY)
	rightlowerdown(w, h, offX, offY)
	leftlowerdown(w, h, offX, offY)
	botcross(w, h, offX, offY)
}

// one draws a one in a WxH size box.
func ones(w int, h int, offX int, offY int) {
	rightupperdown(w, h, offX, offY)
	rightlowerdown(w, h, offX, offY)

}

// twos draws a two in a WxH size box.
func twos(w int, h int, offX int, offY int) {
	topcross(w, h, offX, offY)
	rightupperdown(w, h, offX, offY)
	midcross(w, h, offX, offY)
	leftlowerdown(w, h, offX, offY)
	botcross(w, h, offX, offY)
}

// threes draws a three in a WxH size box.
func threes(w int, h int, offX int, offY int) {
	topcross(w, h, offX, offY)
	rightupperdown(w, h, offX, offY)
	midcross(w, h, offX, offY)
	rightlowerdown(w, h, offX, offY)
	botcross(w, h, offX, offY)
}

// fours draws a four in a WxH size box.
func fours(w int, h int, offX int, offY int) {
	leftupperdown(w, h, offX, offY)
	rightupperdown(w, h, offX, offY)
	midcross(w, h, offX, offY)
	rightlowerdown(w, h, offX, offY)
}

// five draws a five in a WxH size box.
func five(w int, h int, offX int, offY int) {
	topcross(w, h, offX, offY)
	leftupperdown(w, h, offX, offY)
	midcross(w, h, offX, offY)
	rightlowerdown(w, h, offX, offY)
	botcross(w, h, offX, offY)
}

// six draws a six in a WxH size box.
func six(w int, h int, offX int, offY int) {
	topcross(w, h, offX, offY)
	leftupperdown(w, h, offX, offY)
	midcross(w, h, offX, offY)
	rightlowerdown(w, h, offX, offY)
	leftlowerdown(w, h, offX, offY)
	botcross(w, h, offX, offY)
}

// seven draws a seven in a WxH size box.
func seven(w int, h int, offX int, offY int) {
	topcross(w, h, offX, offY)
	rightupperdown(w, h, offX, offY)
	rightlowerdown(w, h, offX, offY)
}

// eight draws a eight in a WxH size box.
func eight(w int, h int, offX int, offY int) {
	topcross(w, h, offX, offY)
	leftupperdown(w, h, offX, offY)
	rightupperdown(w, h, offX, offY)
	midcross(w, h, offX, offY)
	rightlowerdown(w, h, offX, offY)
	leftlowerdown(w, h, offX, offY)
	botcross(w, h, offX, offY)
}

// nine draws a nine in a WxH size box.
func nine(w int, h int, offX int, offY int) {
	topcross(w, h, offX, offY)
	leftupperdown(w, h, offX, offY)
	rightupperdown(w, h, offX, offY)
	midcross(w, h, offX, offY)
	rightlowerdown(w, h, offX, offY)
}

// topcross helper function that draws a top horizontal bar.
func topcross(w int, h int, offX int, offY int) {
	for x := 0; x <= w*7/8; x++ {
		for y := 0; y <= h/10; y++ {
			termbox.SetCell((x + offX), y+(offY), ' ', termbox.ColorDefault,
				termbox.ColorWhite)
		}
	}
}

// midcross helper function that draws a mid horizontal bar.
func midcross(w int, h int, offX int, offY int) {
	for x := 0; x <= w*7/8; x++ {
		for y := 0; y <= h/10; y++ {
			termbox.SetCell((x + offX), y+(offY+h/2), ' ', termbox.ColorDefault,
				termbox.ColorWhite)
		}
	}
}

// botcross helper function that draws a bottom horizontal bar.
func botcross(w int, h int, offX int, offY int) {
	for x := 0; x <= w*7/8; x++ {
		for y := 0; y <= h/10; y++ {
			termbox.SetCell((x + offX), y+(offY+h-h/10+1), ' ', termbox.ColorDefault,
				termbox.ColorWhite)
		}
	}
}

// leftupperdown helper function that draws a left upper vertical bar.
func leftupperdown(w int, h int, offX int, offY int) {
	for x := 0; x <= w/6; x++ {
		for y := 0; y <= h/2; y++ {
			termbox.SetCell(x+offX, y+(offY), ' ', termbox.ColorDefault,
				termbox.ColorWhite)
		}
	}
}

// leftlowerdown helper function that draws a left lower vertical bar.
func leftlowerdown(w int, h int, offX int, offY int) {
	for x := 0; x <= w/6; x++ {
		for y := 0; y <= h/2; y++ {
			termbox.SetCell(x+offX, y+(offY+(h/2+1)), ' ', termbox.ColorDefault,
				termbox.ColorWhite)
		}
	}
}

// rightupperdown helper function that draws a right upper vertical bar.
func rightupperdown(w int, h int, offX int, offY int) {
	for x := 0; x <= w/6; x++ {
		for y := 0; y <= h/2; y++ {
			termbox.SetCell(x+offX+w*7/8-w/6, y+(offY), ' ', termbox.ColorDefault,
				termbox.ColorWhite)
		}
	}
}

// rightlowerdown helper function that draws a right lower vertical bar.
func rightlowerdown(w int, h int, offX int, offY int) {
	for x := 0; x <= w/6; x++ {
		for y := 0; y <= h/2; y++ {
			termbox.SetCell(x+offX+w*7/8-w/6, y+(offY+h/2+1), ' ', termbox.ColorDefault,
				termbox.ColorWhite)
		}
	}

}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	draw()
loop:
	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			}
		default:
			draw()
			time.Sleep(10 * time.Millisecond)
		}
	}
}
