package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// EndpointConfig endpoint config
// swagger:model EndpointConfig
type EndpointConfig struct {

	// address
	// Required: true
	Address string `json:"address"`

	// aliases
	// Required: true
	Aliases []string `json:"aliases"`

	// container
	// Required: true
	Container string `json:"container"`

	// direct
	Direct bool `json:"direct,omitempty"`

	// gateway
	// Required: true
	Gateway string `json:"gateway"`

	// id
	// Required: true
	ID string `json:"id"`

	// name
	// Required: true
	Name string `json:"name"`

	// nameservers
	// Required: true
	Nameservers []string `json:"nameservers"`

	// ports
	// Required: true
	Ports []string `json:"ports"`

	// scope
	// Required: true
	Scope string `json:"scope"`

	// trust
	Trust string `json:"trust,omitempty"`
}

// Validate validates this endpoint config
func (m *EndpointConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAliases(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContainer(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNameservers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointConfig) validateAddress(formats strfmt.Registry) error {

	if err := validate.RequiredString("address", "body", string(m.Address)); err != nil {
		return err
	}

	return nil
}

func (m *EndpointConfig) validateAliases(formats strfmt.Registry) error {

	if err := validate.Required("aliases", "body", m.Aliases); err != nil {
		return err
	}

	return nil
}

func (m *EndpointConfig) validateContainer(formats strfmt.Registry) error {

	if err := validate.RequiredString("container", "body", string(m.Container)); err != nil {
		return err
	}

	return nil
}

func (m *EndpointConfig) validateGateway(formats strfmt.Registry) error {

	if err := validate.RequiredString("gateway", "body", string(m.Gateway)); err != nil {
		return err
	}

	return nil
}

func (m *EndpointConfig) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *EndpointConfig) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *EndpointConfig) validateNameservers(formats strfmt.Registry) error {

	if err := validate.Required("nameservers", "body", m.Nameservers); err != nil {
		return err
	}

	return nil
}

func (m *EndpointConfig) validatePorts(formats strfmt.Registry) error {

	if err := validate.Required("ports", "body", m.Ports); err != nil {
		return err
	}

	return nil
}

func (m *EndpointConfig) validateScope(formats strfmt.Registry) error {

	if err := validate.RequiredString("scope", "body", string(m.Scope)); err != nil {
		return err
	}

	return nil
}
