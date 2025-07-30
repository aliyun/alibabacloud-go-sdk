// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPropertyKey(v string) *CreatePropertyRequest
	GetPropertyKey() *string
	SetPropertyValues(v []*string) *CreatePropertyRequest
	GetPropertyValues() []*string
}

type CreatePropertyRequest struct {
	// The property name.
	//
	// This parameter is required.
	//
	// example:
	//
	// department
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The values of the property. You can specify up to 50 values for a property.
	PropertyValues []*string `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Repeated"`
}

func (s CreatePropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePropertyRequest) GoString() string {
	return s.String()
}

func (s *CreatePropertyRequest) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *CreatePropertyRequest) GetPropertyValues() []*string {
	return s.PropertyValues
}

func (s *CreatePropertyRequest) SetPropertyKey(v string) *CreatePropertyRequest {
	s.PropertyKey = &v
	return s
}

func (s *CreatePropertyRequest) SetPropertyValues(v []*string) *CreatePropertyRequest {
	s.PropertyValues = v
	return s
}

func (s *CreatePropertyRequest) Validate() error {
	return dara.Validate(s)
}
