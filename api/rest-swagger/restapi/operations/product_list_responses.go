// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Smart-Purveyance-Tracker/backend/api/rest-swagger/models"
)

// ProductListOKCode is the HTTP code returned for type ProductListOK
const ProductListOKCode int = 200

/*ProductListOK returns list of product

swagger:response productListOK
*/
type ProductListOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Product `json:"body,omitempty"`
}

// NewProductListOK creates ProductListOK with default headers values
func NewProductListOK() *ProductListOK {

	return &ProductListOK{}
}

// WithPayload adds the payload to the product list o k response
func (o *ProductListOK) WithPayload(payload []*models.Product) *ProductListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product list o k response
func (o *ProductListOK) SetPayload(payload []*models.Product) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Product, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*ProductListDefault error

swagger:response productListDefault
*/
type ProductListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewProductListDefault creates ProductListDefault with default headers values
func NewProductListDefault(code int) *ProductListDefault {
	if code <= 0 {
		code = 500
	}

	return &ProductListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the product list default response
func (o *ProductListDefault) WithStatusCode(code int) *ProductListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the product list default response
func (o *ProductListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the product list default response
func (o *ProductListDefault) WithPayload(payload *models.Error) *ProductListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the product list default response
func (o *ProductListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProductListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
