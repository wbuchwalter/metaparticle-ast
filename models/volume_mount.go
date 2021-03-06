// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VolumeMount volume mount
// swagger:model volumeMount
type VolumeMount struct {

	// mount path
	// Required: true
	MountPath *string `json:"mountPath"`

	// name
	// Required: true
	Name *string `json:"name"`

	// sub path
	SubPath string `json:"subPath,omitempty"`
}

// Validate validates this volume mount
func (m *VolumeMount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMountPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeMount) validateMountPath(formats strfmt.Registry) error {

	if err := validate.Required("mountPath", "body", m.MountPath); err != nil {
		return err
	}

	return nil
}

func (m *VolumeMount) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolumeMount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeMount) UnmarshalBinary(b []byte) error {
	var res VolumeMount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
