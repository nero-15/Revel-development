package controllers

import (
	"net/http"

	"github.com/revel/revel"
)

type Hotels struct {
	*revel.Controller
}

func (c Hotels) Index() revel.Result {
	new_cookie := &http.Cookie{Name: "foo", Value: "Bar"}
	c.SetCookie(new_cookie)
	var query = c.Params.Query
	return c.Render(query)
}

func (c Hotels) Show() revel.Result {
	var id string = c.Params.Get("id")
	return c.Render(id)
}
