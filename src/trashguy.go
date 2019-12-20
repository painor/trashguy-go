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
	"strings"
)

var symbols = Symbols{
	[]string{"\U0001F353", "\U0001F34A", "\U0001F345"},
	"\u0020",
	"\u2800\u0020",
	"\u2796",
	"`",
	"```",
	"\U0001F5D1",
	"<(^_^ <)",
	"(> ^_^)>",
}
var DEFAULT_OPTIONS = Options{
	spriteCan:   symbols.SPRITE_CAN,
	spriteLeft:  symbols.SPRITE_LEFT,
	spriteRight: symbols.SPRITE_RIGHT,
	spacer:      symbols.SPACER_DEFAULT,
	wrapper:     "",
	frameStart:  0,
	framesMax:   -1,
}

type TrashGuy struct {
	TrashItems   string
	Options      Options
	Sprites      Sprite
	Slices       Slice
	FrameStart   int
	Index        int
	DefaultIndex int
}

func (this *TrashGuy) Init() {
	if this.TrashItems == "" {
		panic("HELP I NEED ITEMSSSSSSS")
	}
	sanTrashItems := strings.Split(this.TrashItems, " ")
	this.Sprites = Sprite{
		trashItems:  sanTrashItems,
		spriteTrash: this.Options.spriteCan,
		spriteLeft:  this.Options.spriteLeft,
		spriteRight: this.Options.spriteRight,
		spacer:      this.Options.spacer,
		wrapper:     this.Options.wrapper,
	}
	this.FrameStart = this.Options.frameStart
	this.Index = this.Options.frameStart - 1
	this.DefaultIndex = this.Index
	maxTotalFrames := this.length()
	maxAvailableFrames := maxTotalFrames - this.FrameStart
	if this.Options.framesMax == -1 {
		this.Options.framesMax = maxAvailableFrames
	}

	this.Slices = Slice{
		frameStart: this.Options.framesMax,
		frameMax:   this.Options.framesMax,
	}
}
func (this *TrashGuy) length() int {
	return frameGroupValues(this.Sprites).totalFrameCount
}
func (this *TrashGuy) getNext() string {
	this.Index++
	return getFrame(this.Slices, this.Sprites, this.Index)
}
func (this *TrashGuy) Animate() string {
	var ar []string

	for {
		item := this.getNext()
		if item == "" {
			break
		}
		ar = append(ar, item)
	}
	return strings.Join(ar,"\n")
}
