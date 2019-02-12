package main

import (
	"flag"
	"fmt"
	"math"
)

const (
	width, height = 1200, 800
	cells         = 200
	xyrange       = 30.0
	zrange        = 70.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6 // 30 degrees
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var eq = flag.String("eq", "orig", "orig or monkey or egg")

func main() {
	flag.Parse()
	var f func(x, y float64) float64
	switch *eq {
	case "well":
		f = well
	case "egg":
		f = egg
	case "monkey":
		f = monkeySaddle
	case "orig":
		f = orig
	default:
		f = orig
	}
	fmt.Printf("<svg xmlns='http://wwww.w3.org/2000/svg' "+
		"style='stroke: grey; fill:white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, colour, ok1 := corner(i+1, j, f)
			bx, by, colour, ok2 := corner(i, j, f)
			cx, cy, colour, ok3 := corner(i, j+1, f)
			dx, dy, colour, ok4 := corner(i+1, j+1, f)
			if ok1 && ok2 && ok3 && ok4 {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s' />\n",
					ax, ay, bx, by, cx, cy, dx, dy, colour)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i int, j int, f func(x, y float64) float64) (float64, float64, string, bool) {
	colour := "#00ff00"
	x := xyrange * (float64(i)/(cells-1) - 0.5)
	y := xyrange * (float64(j)/(cells-1) - 0.5)

	z := math.Mod(f(x, y), height)
	z = zrange * (float64(z) / (cells - 1))

	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, colour, false
	}

	if z > 0.1 {
		colour = "#ff0000"
	}

	if z < 0.005 {
		colour = "#0000ff"
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, colour, true
}

func orig(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func well(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r * math.Sin(x)
}

func saddle(x, y float64) float64 {
	z := math.Pow(x, 2) - math.Pow(y, 2)
	return z
}

func egg(x, y float64) float64 {
	z := math.Sin(x) * math.Sin(y)
	return z
}

func monkeySaddle(x, y float64) float64 {
	//z := x*x - y*y
	z := math.Pow(x, 3) - 3*x*math.Pow(y, 2)
	return z
}