// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UploadHostLogsHandlerFunc turns a function with the right signature into a upload host logs handler
type UploadHostLogsHandlerFunc func(UploadHostLogsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UploadHostLogsHandlerFunc) Handle(params UploadHostLogsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UploadHostLogsHandler interface for that can handle valid upload host logs params
type UploadHostLogsHandler interface {
	Handle(UploadHostLogsParams, interface{}) middleware.Responder
}

// NewUploadHostLogs creates a new http.Handler for the upload host logs operation
func NewUploadHostLogs(ctx *middleware.Context, handler UploadHostLogsHandler) *UploadHostLogs {
	return &UploadHostLogs{Context: ctx, Handler: handler}
}

/*UploadHostLogs swagger:route POST /clusters/{cluster_id}/hosts/{host_id}/logs installer uploadHostLogs

Agent API to upload logs.

*/
type UploadHostLogs struct {
	Context *middleware.Context
	Handler UploadHostLogsHandler
}

func (o *UploadHostLogs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUploadHostLogsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
