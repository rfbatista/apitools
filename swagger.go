package apitools

import (
	"log"

	"github.com/swaggest/openapi-go"
	"github.com/swaggest/openapi-go/openapi3"
	"go.uber.org/zap"
)

func NewSwaggerRegister(p *zap.Logger) *SwaggerRegister {
	return &SwaggerRegister{log: p}
}

type SwaggerRegister struct {
	log         *zap.Logger
	title       string
	version     string
	description string
	handlers    []Handler
}

func (s *SwaggerRegister) AddHandler(h Handler) {
	s.log.Debug("add handler")
	s.handlers = append(s.handlers, h)
}

func (s *SwaggerRegister) SetTitle(t string) {
	s.title = t
}

func (s *SwaggerRegister) SetVersion(v string) {
	s.version = v
}

func (s *SwaggerRegister) SetDescription(d string) {
	s.description = d
}

func (s *SwaggerRegister) CreateSpecs() ([]byte, error) {
	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.0.3"}
	reflector.Spec.Info.
		WithTitle(s.title).
		WithVersion(s.version).
		WithDescription(s.description)
	s.log.Debug("creating doc", zap.Int("handlers", len(s.handlers)))
	for _, handle := range s.handlers {
		op, err := reflector.NewOperationContext(handle.Method().String(), handle.Path())
		if err != nil {
			return nil, err
		}
		op.AddReqStructure(handle.Request())
		for _, v := range handle.Response() {
			op.AddRespStructure(v.Schema, func(cu *openapi.ContentUnit) {
				cu.HTTPStatus = v.Status
				cu.ContentType = v.ContentType
			})
		}
		err = reflector.AddOperation(op)
		if err != nil {
			return nil, err
		}
	}
	schema, err := reflector.Spec.MarshalYAML()
	if err != nil {
		log.Fatal(err)
	}
	return schema, nil
}
