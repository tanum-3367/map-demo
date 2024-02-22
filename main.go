package main

import (
	"github.com/fogleman/gg"
)

func main() {
	const Width = 600
	const Height = 400

	// set up objects
	room := Room{X: 50, Y: 50, Width: 500, Height: 300}
	objects := []Object{
		{X: 100, Y: 100, Width: 50, Height: 50},
		{X: 200, Y: 200, Width: 50, Height: 50},
	}
	robot := Robot{X: 300, Y: 200, Radius: 30}

	// set up a new canvas
	dc := gg.NewContext(Width, Height)

	// set background color
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// draw room
	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(room.X, room.Y, room.Width, room.Height)
	dc.Stroke()

	// draw objects
	dc.SetRGB(1, 0, 0)
	for _, object := range objects {
		dc.DrawRectangle(object.X, object.Y, object.Width, object.Height)
		dc.Stroke()
	}

	// draw robot
	dc.SetRGB(0, 0, 1)
	dc.DrawCircle(robot.X, robot.Y, robot.Radius)
	dc.Stroke()

	dc.SavePNG("room_layout.png")
}
