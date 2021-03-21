package controllers

import (
	"github.com/revel/revel"
)

type Hotels struct {
	*revel.Controller
}

func (c Hotels) Index() revel.Result {
	var query = c.Params.Query
	return c.Render(query)
}

func (c Hotels) Show() revel.Result {
	var id string = c.Params.Get("id")
	return c.Render(id)
}
