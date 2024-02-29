package apitools

import (
	"github.com/labstack/echo/v4"
)

// should be a controller
type HandlerGroup interface {
	Path() string
	Routes() []Handler
	Middlewares() []echo.MiddlewareFunc
	AddHandler(Handler)
	AddMiddleware(echo.MiddlewareFunc)
	SetPath(string)
}

func NewHandlerGroup() HandlerGroup {
	return &ModuleHandler{}
}

type ModuleHandler struct {
	path        string
	routes      []Handler
	middlewares []echo.MiddlewareFunc
}

func (m *ModuleHandler) Path() (_ string) {
	return m.path
}

func (m *ModuleHandler) Routes() (_ []Handler) {
	return m.routes
}

func (m *ModuleHandler) Middlewares() (_ []echo.MiddlewareFunc) {
	return m.middlewares
}

func (m *ModuleHandler) AddHandler(r Handler) {
	m.routes = append(m.routes, r)
}

func (m *ModuleHandler) AddMiddleware(mid echo.MiddlewareFunc) {
	m.middlewares = append(m.middlewares, mid)
}

func (m *ModuleHandler) SetPath(path string) {
	m.path = path
}
