// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ProductListHandlerFunc turns a function with the right signature into a product list handler
type ProductListHandlerFunc func(ProductListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ProductListHandlerFunc) Handle(params ProductListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ProductListHandler interface for that can handle valid product list params
type ProductListHandler interface {
	Handle(ProductListParams, interface{}) middleware.Responder
}

// NewProductList creates a new http.Handler for the product list operation
func NewProductList(ctx *middleware.Context, handler ProductListHandler) *ProductList {
	return &ProductList{Context: ctx, Handler: handler}
}

/*ProductList swagger:route GET /product/list productList

ProductList product list API

*/
type ProductList struct {
	Context *middleware.Context
	Handler ProductListHandler
}

func (o *ProductList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewProductListParams()

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
