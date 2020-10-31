package main

import (
	"fmt"

	"github.com/kyroy/kdtree"
	"github.com/kyroy/kdtree/points"
)

func main() {
	fmt.Println("go-reverse-geocoder")

	tree := kdtree.New([]kdtree.Point{
		points.NewPoint([]float64{0, 2}, "0, 2"),
		points.NewPoint([]float64{1, 2}, "1, 2"),
		points.NewPoint([]float64{0, 3}, "0, 3"),
		points.NewPoint([]float64{2, 3}, "2, 3"),
		points.NewPoint([]float64{2, 3}, "2, 3"),
		points.NewPoint([]float64{-1, -1}, "-1 -1"),
	})

	fmt.Println(tree.KNN(&points.Point{Coordinates: []float64{0, 0}}, 3))
}
