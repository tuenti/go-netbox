// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VirtualChassis virtual chassis
// swagger:model VirtualChassis
type VirtualChassis struct {

	// Domain
	// Max Length: 30
	Domain string `json:"domain,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// master
	// Required: true
	Master *NestedDevice `json:"master"`
}

// Validate validates this virtual chassis
func (m *VirtualChassis) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualChassis) validateDomain(formats strfmt.Registry) error {

	if swag.IsZero(m.Domain) { // not required
		return nil
	}

	if err := validate.MaxLength("domain", "body", string(m.Domain), 30); err != nil {
		return err
	}

	return nil
}

func (m *VirtualChassis) validateMaster(formats strfmt.Registry) error {

	if err := validate.Required("master", "body", m.Master); err != nil {
		return err
	}

	if m.Master != nil {
		if err := m.Master.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("master")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualChassis) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualChassis) UnmarshalBinary(b []byte) error {
	var res VirtualChassis
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
