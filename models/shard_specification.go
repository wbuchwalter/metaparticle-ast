// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ShardSpecification shard specification
// swagger:model shardSpecification

type ShardSpecification struct {

	// field path
	FieldPath string `json:"fieldPath,omitempty"`

	// shards
	Shards int32 `json:"shards,omitempty"`

	// url pattern
	URLPattern string `json:"urlPattern,omitempty"`
}

/* polymorph shardSpecification fieldPath false */

/* polymorph shardSpecification shards false */

/* polymorph shardSpecification urlPattern false */

// Validate validates this shard specification
func (m *ShardSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ShardSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShardSpecification) UnmarshalBinary(b []byte) error {
	var res ShardSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}