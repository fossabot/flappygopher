package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type bird struct {
	time     int
	textures []*sdl.Texture
}

func newBird(r *sdl.Renderer) (*bird, error) {

	var textures []*sdl.Texture
	for i := 1; i <= 4; i++ {
		path := fmt.Sprintf("res/imgs/bird_frame_%d.png", i)
		texture, err := img.LoadTexture(r, path)
		if err != nil {
			return nil, fmt.Errorf("could not load bird frame: %v", err)
		}
		textures = append(textures, texture)
	}
	return &bird{textures: textures}, nil
}