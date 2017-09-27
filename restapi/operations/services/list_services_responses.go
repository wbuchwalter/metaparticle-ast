// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/metaparticle-io/metaparticle-ast/models"
)

// ListServicesOKCode is the HTTP code returned for type ListServicesOK
const ListServicesOKCode int = 200

/*ListServicesOK list the services

swagger:response listServicesOK
*/
type ListServicesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Service `json:"body,omitempty"`
}

// NewListServicesOK creates ListServicesOK with default headers values
func NewListServicesOK() *ListServicesOK {
	return &ListServicesOK{}
}

// WithPayload adds the payload to the list services o k response
func (o *ListServicesOK) WithPayload(payload []*models.Service) *ListServicesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list services o k response
func (o *ListServicesOK) SetPayload(payload []*models.Service) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListServicesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Service, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}