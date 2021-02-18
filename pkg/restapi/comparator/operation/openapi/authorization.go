// Code generated by go-swagger; DO NOT EDIT.

package openapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Authorization An authorization.
//
// swagger:model Authorization
type Authorization struct {

	// auth token
	// Required: true
	AuthToken *string `json:"authToken"`

	// id
	ID string `json:"id,omitempty"`

	// requesting party DID
	// Required: true
	RequestingParty *string `json:"requestingParty"`

	// scope
	// Required: true
	Scope []*Scope `json:"scope"`
}

// Validate validates this authorization
func (m *Authorization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestingParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Authorization) validateAuthToken(formats strfmt.Registry) error {

	if err := validate.Required("authToken", "body", m.AuthToken); err != nil {
		return err
	}

	return nil
}

func (m *Authorization) validateRequestingParty(formats strfmt.Registry) error {

	if err := validate.Required("requestingParty", "body", m.RequestingParty); err != nil {
		return err
	}

	return nil
}

func (m *Authorization) validateScope(formats strfmt.Registry) error {

	if err := validate.Required("scope", "body", m.Scope); err != nil {
		return err
	}

	for i := 0; i < len(m.Scope); i++ {
		if swag.IsZero(m.Scope[i]) { // not required
			continue
		}

		if m.Scope[i] != nil {
			if err := m.Scope[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scope" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this authorization based on the context it is used
func (m *Authorization) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Authorization) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Scope); i++ {

		if m.Scope[i] != nil {
			if err := m.Scope[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scope" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Authorization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Authorization) UnmarshalBinary(b []byte) error {
	var res Authorization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}