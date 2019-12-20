// ============================================= //
//             Trash Guy Script                  //
//                 (> ^_^)>                      //
//          Made by Painor (t.me/painor)         //
//                                               //
// ============================================= // ========================== //
// Copyright (C) 2019 Painor (https://t.me/painor)                             //
// Permission is hereby granted, free of charge, to any person obtaining a     //
// copy of this software and associated documentation files (the "Software"),  //
// to deal in the Software without restriction, including without limitation   //
// the rights to use, copy, modify, merge, publish, distribute, sublicense,    //
// and/or sell copies of the Software, and to permit persons to whom the       //
// Software is furnished to do so, subject to the following conditions: The    //
// above copyright notice and this permission notice shall be included in all  //
// copies or substantial portions of the Software.                             //
// =========================================================================== //
package src

import (
	"sort"
	"strings"
	"unicode/utf8"
)

func makeRange(min, max, step int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i + step
	}
	return a
}
func sum(input []int) int {
	sum := 0
	for i := range input {
		sum += input[i]
	}
	return sum
}
func reverse(ss []int) []int {
	var ok []int
	copy(ss, ok)

	last := len(ok) - 1
	for i := 0; i < len(ok)/2; i++ {
		ok[i], ok[last-i] = ok[last-i], ok[i]
	}
	return ok
}

func calcMedian(n []int) int {
	sort.Ints(n) // sort the numbers

	mNumber := len(n) / 2

	if mNumber%2 == 0 {
		return (n[mNumber-1] + n[mNumber]) / 2
	}

	return n[mNumber]
}

func wrap(sprite Sprite, canvas []string) string {
	if sprite.wrapper != "" {
		return strings.Join(canvas[:], " ")
	} else {
		return strings.Join(canvas[:], " ")
	}
}

func frameGroupValues(sprites Sprite) GroupValues {
	minFgSize := 6
	maxFgSizeIndex := minFgSize + len(sprites.trashItems)*2
	fgSizes := makeRange(minFgSize+1, maxFgSizeIndex+1, 2)

	totalFrameCount := sum(fgSizes)
	return GroupValues{fgSizes, totalFrameCount}
}
func converter(slc Slice, spr Sprite, index int) Position {

	groupValues := frameGroupValues(spr)
	tempFrameIndex := 0
	if index-slc.frameStart < slc.frameMax {
		for i := 0; i < len(groupValues.fgSizes); i++ {
			fg := makeRange(2, groupValues.fgSizes[i]+2, 1)
			midFg := calcMedian(fg)
			upperFg := fg[0 : midFg-2]

			lowerFg := reverse(upperFg)
			lowerFg = append(lowerFg, 1, 0)
			for pos := range upperFg {
				if tempFrameIndex == index {
					return Position{pos, true, i}
				}
				tempFrameIndex++
			}
			for pos := range lowerFg {
				if tempFrameIndex == index {
					return Position{pos, false, i}
				}
				tempFrameIndex++
			}

		}
	}
	panic("HELP PLEASE I AM OUT OF RANGE AAAAAAAAAAAAAAA")
}

func getFrame(slc Slice, spr Sprite, index int) string {
	framePosition := converter(slc, spr, index)
	truncItems := spr.trashItems[framePosition.itemIndex:len(spr.trashItems)]

	missingItemsLen := 0
	for _, x := range spr.trashItems[0:framePosition.itemIndex] {
		missingItemsLen += utf8.RuneCountInString(x)
	}

	var padding []string
	for i := 0; i < missingItemsLen; i++ {
		padding = append(padding, spr.spacer)
	}
	//Create a dynamic canvas while each item disappears
	var canvas []string
	canvas = append(canvas, spr.spriteTrash)
	canvas = append(canvas, spr.spacer)
	canvas = append(canvas, spr.spacer)
	canvas = append(canvas, spr.spacer)
	for _, x := range padding {
		canvas = append(canvas, x)
	}
	for _, x := range truncItems {
		canvas = append(canvas, x)
	}
	itemTruncateLength := -len(truncItems)
	lastIndex := len(canvas) - (-itemTruncateLength)

	if framePosition.itemIndex < len(spr.trashItems) && framePosition.position < len(canvas) {

		// Start sequence, forward motion, going right
		if framePosition.forward {
			if framePosition.position < lastIndex {
				// Start from second space after the trash can

				canvas[framePosition.position] = spr.spriteRight
				// Snapshot the frames of the animation going right
				return wrap(spr, canvas)
			} else {
				// End of forward motion, look left with item
				// Set item position in front of trash guy
				canvas[framePosition.position-1] = canvas[lastIndex]
				// Set position of trash guy where item was
				canvas[framePosition.position] = spr.spriteLeft
				// Snapshot frame looking across at trash can
				return wrap(spr, canvas)
			}
		} else { // Reverse motion, going left
			// Going left with item towards trash can
			if framePosition.position > 0 {
				canvas[framePosition.position] = spr.spriteLeft
				// Place item in front while not yet at the trash can
				if canvas[framePosition.position-1] != spr.spriteTrash {
					canvas[framePosition.position-1] = canvas[lastIndex]
					// Temporarily remove item from pile while holding it
					canvas[lastIndex] = spr.spacer

				} else {
					// If trash can reached, replace spacing of missing item
					if len(spr.spacer) == 1 {
						lastItemLen := len(canvas[lastIndex])
						var tempCanvas []string
						tempCanvas = canvas[0:lastIndex]
						for i := 0; i < lastItemLen; i++ {
							tempCanvas = append(tempCanvas, spr.spacer)
						}
						for _, x := range canvas[lastIndex+1:] {
							tempCanvas = append(tempCanvas, x)
						}
						canvas = tempCanvas
					} else {
						// Unknown spacer size, use as directed
						canvas[lastIndex] = spr.spacer
					}
					// Snapshot the frame looking right
					return wrap(spr, canvas)
				}
			}
		}
	}
	return ""
}
