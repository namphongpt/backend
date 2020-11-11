// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Smart-Purveyance-Tracker/backend/api/rest-swagger/models"
)

// SignupOKCode is the HTTP code returned for type SignupOK
const SignupOKCode int = 200

/*SignupOK returns user

swagger:response signupOK
*/
type SignupOK struct {
	/*

	 */
	Authenthication string `json:"Authenthication"`

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewSignupOK creates SignupOK with default headers values
func NewSignupOK() *SignupOK {

	return &SignupOK{}
}

// WithAuthenthication adds the authenthication to the signup o k response
func (o *SignupOK) WithAuthenthication(authenthication string) *SignupOK {
	o.Authenthication = authenthication
	return o
}

// SetAuthenthication sets the authenthication to the signup o k response
func (o *SignupOK) SetAuthenthication(authenthication string) {
	o.Authenthication = authenthication
}

// WithPayload adds the payload to the signup o k response
func (o *SignupOK) WithPayload(payload *models.User) *SignupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the signup o k response
func (o *SignupOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SignupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Authenthication

	authenthication := o.Authenthication
	if authenthication != "" {
		rw.Header().Set("Authenthication", authenthication)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*SignupDefault error

swagger:response signupDefault
*/
type SignupDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSignupDefault creates SignupDefault with default headers values
func NewSignupDefault(code int) *SignupDefault {
	if code <= 0 {
		code = 500
	}

	return &SignupDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the signup default response
func (o *SignupDefault) WithStatusCode(code int) *SignupDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the signup default response
func (o *SignupDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the signup default response
func (o *SignupDefault) WithPayload(payload *models.Error) *SignupDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the signup default response
func (o *SignupDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SignupDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}