// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomParams interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CustomParams
	GetName() *string
	SetProperties(v []*Property) *CustomParams
	GetProperties() []*Property
}

type CustomParams struct {
	// example:
	//
	// Normalize
	Name       *string     `json:"Name,omitempty" xml:"Name,omitempty"`
	Properties []*Property `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
}

func (s CustomParams) String() string {
	return dara.Prettify(s)
}

func (s CustomParams) GoString() string {
	return s.String()
}

func (s *CustomParams) GetName() *string {
	return s.Name
}

func (s *CustomParams) GetProperties() []*Property {
	return s.Properties
}

func (s *CustomParams) SetName(v string) *CustomParams {
	s.Name = &v
	return s
}

func (s *CustomParams) SetProperties(v []*Property) *CustomParams {
	s.Properties = v
	return s
}

func (s *CustomParams) Validate() error {
	return dara.Validate(s)
}
