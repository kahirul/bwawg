package main

import (
	"gopkg.in/unrolled/render.v1"
	"log"
	"net/http"
)

type Action func(rw http.ResponseWriter, r *http.Request) error

type AppController struct{}

func (c *AppController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Print(r.Header)
		if err := a(rw, r); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}

type MyController struct {
	AppController
	*render.Render
}

func (c *MyController) Index(rw http.ResponseWriter, r *http.Request) error {
	c.JSON(rw, 200, map[string]string{"msg": "Hello from Index"})
	return nil
}

func (c *MyController) Home(rw http.ResponseWriter, r *http.Request) error {
	c.JSON(rw, 200, map[string]string{"msg": "Hello from Home"})
	return nil
}

func main() {
	c := &MyController{Render: render.New(render.Options{})}

	mux := http.NewServeMux()
	mux.Handle("/", c.Action(c.Index))
	mux.Handle("/home", c.Action(c.Home))

	log.Print("Listening at :8080")
	http.ListenAndServe(":8080", mux)
}
