// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// CreateSource defines model for CreateSource.
type CreateSource struct {
	Name    string `binding:"required" json:"name"`
	Summary string `binding:"required" json:"summary"`
	Tags    string `binding:"required" json:"tags"`
	Uri     string `binding:"required" json:"uri"`
}

// Pong defines model for Pong.
type Pong struct {
	Pong string `json:"pong"`
}

// Source defines model for Source.
type Source struct {
	Name      string  `binding:"required" json:"name"`
	Score     int     `binding:"required" json:"score"`
	Summary   string  `binding:"required" json:"summary"`
	Tags      string  `binding:"required" json:"tags"`
	Uri       string  `binding:"required" json:"uri"`
	UriDigest *string `binding:"required" gorm:"primary_key" json:"uriDigest,omitempty"`
}

// CreateSourceJSONRequestBody defines body for CreateSource for application/json ContentType.
type CreateSourceJSONRequestBody = CreateSource

// UpdateSourceJSONRequestBody defines body for UpdateSource for application/json ContentType.
type UpdateSourceJSONRequestBody = CreateSource

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /api/v1/sources)
	CreateSource(c *gin.Context)

	// (DELETE /api/v1/sources/{uriDigest})
	DeleteSource(c *gin.Context, uriDigest string)

	// (GET /api/v1/sources/{uriDigest})
	GetSource(c *gin.Context, uriDigest string)

	// (PUT /api/v1/sources/{uriDigest})
	UpdateSource(c *gin.Context, uriDigest string)

	// (GET /ping)
	GetPing(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// CreateSource operation middleware
func (siw *ServerInterfaceWrapper) CreateSource(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateSource(c)
}

// DeleteSource operation middleware
func (siw *ServerInterfaceWrapper) DeleteSource(c *gin.Context) {

	var err error

	// ------------- Path parameter "uriDigest" -------------
	var uriDigest string

	err = runtime.BindStyledParameterWithOptions("simple", "uriDigest", c.Param("uriDigest"), &uriDigest, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uriDigest: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteSource(c, uriDigest)
}

// GetSource operation middleware
func (siw *ServerInterfaceWrapper) GetSource(c *gin.Context) {

	var err error

	// ------------- Path parameter "uriDigest" -------------
	var uriDigest string

	err = runtime.BindStyledParameterWithOptions("simple", "uriDigest", c.Param("uriDigest"), &uriDigest, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uriDigest: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetSource(c, uriDigest)
}

// UpdateSource operation middleware
func (siw *ServerInterfaceWrapper) UpdateSource(c *gin.Context) {

	var err error

	// ------------- Path parameter "uriDigest" -------------
	var uriDigest string

	err = runtime.BindStyledParameterWithOptions("simple", "uriDigest", c.Param("uriDigest"), &uriDigest, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uriDigest: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateSource(c, uriDigest)
}

// GetPing operation middleware
func (siw *ServerInterfaceWrapper) GetPing(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetPing(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/api/v1/sources", wrapper.CreateSource)
	router.DELETE(options.BaseURL+"/api/v1/sources/:uriDigest", wrapper.DeleteSource)
	router.GET(options.BaseURL+"/api/v1/sources/:uriDigest", wrapper.GetSource)
	router.PUT(options.BaseURL+"/api/v1/sources/:uriDigest", wrapper.UpdateSource)
	router.GET(options.BaseURL+"/ping", wrapper.GetPing)
}
