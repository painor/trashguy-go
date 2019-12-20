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

type FrameGroup struct {
	FgSizes         []int
	TotalFrameCount int
}
type Conversion struct {
	Position  int
	Forward   bool
	ItemIndex int
}
type Sprite struct {
	trashItems  []string
	spriteTrash string
	spriteLeft  string
	spriteRight string
	spacer      string
	wrapper     string
}
type GroupValues struct {
	fgSizes         []int
	totalFrameCount int
}
type Slice struct {
	frameStart int
	frameMax   int
}

type Position struct {
	position  int
	forward   bool
	itemIndex int
}

type Symbols struct {
	DEFAULT_INPUT []string
	SPACER_DEFAULT string
	SPACER_WIDE string
	SPACER_EMOJI string
	WRAPPER_MONOSPACE string
	WRAPPER_BLOCK_MONO  string
	SPRITE_CAN string
	SPRITE_LEFT string
	SPRITE_RIGHT string
}

type Options struct {
	spriteCan string
	spriteLeft string
	spriteRight string
	spacer string
	wrapper string
	frameStart int
	framesMax int
}