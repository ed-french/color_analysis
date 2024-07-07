package main

import (
	"fmt"
	"time"
)

func main() {
	filenames := []string{"flower.jpg",
		"cake.JPG",
		"road.JPG",
		"fire.JPG",
		"ipad.jpg",
		"smash.JPG",
		"124.JPG",
		"20120807_114539.jpg",
		"43C2612F-F3C9-4549-95E4-E1ED21DD6E16.jpg",
		"6-15-2014_100.JPG",
		"__Epson_26032016142042 (2).jpg",
		"20140806_121541.jpg"}

	for _, filename := range filenames {
		fmt.Print(filename)
		matt, err := get_matt_color(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(matt)
		time.Sleep(4 * time.Second)

	}

}
