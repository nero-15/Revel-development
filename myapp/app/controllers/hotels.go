package controllers

import (
	"github.com/revel/revel"
)

type Hotels struct {
	*revel.Controller
}

func (c Hotels) Index() revel.Result {
	return c.Render()
}

func (c Hotels) Show() revel.Result {
	var id string = c.Params.Get("id")
	return c.Render(id)
}
