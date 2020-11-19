// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Smart-Purveyance-Tracker/backend/api/rest-swagger/models"
)

// GetProductOKCode is the HTTP code returned for type GetProductOK
const GetProductOKCode int = 200

/*GetProductOK get product

swagger:response getProductOK
*/
type GetProductOK struct {

	/*
	  In: Body
	*/
	Payload *models.Product `json:"body,omitempty"`
}

// NewGetProductOK creates GetProductOK with default headers values
func NewGetProductOK() *GetProductOK {

	return &GetProductOK{}
}

// WithPayload adds the payload to the get product o k response
func (o *GetProductOK) WithPayload(payload *models.Product) *GetProductOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get product o k response
func (o *GetProductOK) SetPayload(payload *models.Product) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetProductDefault error

swagger:response getProductDefault
*/
type GetProductDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetProductDefault creates GetProductDefault with default headers values
func NewGetProductDefault(code int) *GetProductDefault {
	if code <= 0 {
		code = 500
	}

	return &GetProductDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get product default response
func (o *GetProductDefault) WithStatusCode(code int) *GetProductDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get product default response
func (o *GetProductDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get product default response
func (o *GetProductDefault) WithPayload(payload *models.Error) *GetProductDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get product default response
func (o *GetProductDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
