package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Proxy proxy

swagger:model Proxy
*/
type Proxy struct {

	/* format
	 */
	Format *string `json:"format,omitempty"`

	/* query

	Required: true
	*/
	Query *string `json:"query"`
}

// Validate validates this proxy
func (m *Proxy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormat(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var proxyTypeFormatPropEnum []interface{}

// prop value enum
func (m *Proxy) validateFormatEnum(path, location string, value string) error {
	if proxyTypeFormatPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["raw"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			proxyTypeFormatPropEnum = append(proxyTypeFormatPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, proxyTypeFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Proxy) validateFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", *m.Format); err != nil {
		return err
	}

	return nil
}

func (m *Proxy) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}
