package apitools

import (
	"github.com/labstack/echo/v4"
)

type Handler interface {
	Middleware() []echo.MiddlewareFunc
	Handle() echo.HandlerFunc
	Method() Method
	Path() string
	SetHandle(h echo.HandlerFunc)
	SetMethod(m Method)
	SetPath(p string)
	SetRequest(r interface{})
	Request() interface{}
	Response() []responseOperation
	AddResponse(rop responseOperation)
}

func NewHandler() Handler {
	return &RouteHandler{}
}

type responseOperation struct {
	Schema      interface{}
	ContentType string
	Status      int
}

func ResponseOperation(status int, schema interface{}) responseOperation {
	return responseOperation{
		Status:      status,
		Schema:      schema,
		ContentType: "text/json",
	}
}

func ResponseOperationWithType(
	status int,
	schema interface{},
	contentType string,
) responseOperation {
	return responseOperation{
		Status:      status,
		Schema:      schema,
		ContentType: contentType,
	}
}

type RouteHandler struct {
	request     interface{}
	response    []responseOperation
	handler     echo.HandlerFunc
	path        string
	middlewares []echo.MiddlewareFunc
	method      Method
}

func (r *RouteHandler) Response() []responseOperation {
	return r.response
}

func (r *RouteHandler) Request() interface{} {
	return r.request
}

func (r *RouteHandler) SetRequest(req interface{}) {
	r.request = req
}

func (r *RouteHandler) AddResponse(rop responseOperation) {
	r.response = append(r.response, rop)
}

func (r *RouteHandler) Middleware() (_ []echo.MiddlewareFunc) {
	return r.middlewares
}

func (r *RouteHandler) AddMiddleware(m echo.MiddlewareFunc) {
	r.middlewares = append(r.middlewares, m)
}

func (r *RouteHandler) Handle() (_ echo.HandlerFunc) {
	return r.handler
}

func (r *RouteHandler) Method() (_ Method) {
	return r.method
}

func (r *RouteHandler) Path() (_ string) {
	return r.path
}

func (r *RouteHandler) SetHandle(h echo.HandlerFunc) {
	r.handler = h
}

func (r *RouteHandler) SetMethod(m Method) {
	r.method = m
}

func (r *RouteHandler) SetPath(p string) {
	r.path = p
}
