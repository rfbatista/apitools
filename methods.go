package apitools

import (
	"fmt"
)

type Method int

const (
	POST Method = iota
	GET
	PATCH
	PUT
	DELETE
)

func (m Method) String() string {
	switch m {
	case POST:
		return "POST"
	case GET:
		return "GET"
	case PATCH:
		return "PATCH"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	default:
		return fmt.Sprintf("Unknown Method: %d", m)
	}
}

type ResponseType string

const (
	HTML ResponseType = "text/html"
	JSON ResponseType = "text/json"
)
