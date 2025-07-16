// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDentryAppPropertiesValue interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DentryAppPropertiesValue
	GetName() *string
	SetValue(v string) *DentryAppPropertiesValue
	GetValue() *string
	SetVisibility(v string) *DentryAppPropertiesValue
	GetVisibility() *string
}

type DentryAppPropertiesValue struct {
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// PUBLIC
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DentryAppPropertiesValue) String() string {
	return dara.Prettify(s)
}

func (s DentryAppPropertiesValue) GoString() string {
	return s.String()
}

func (s *DentryAppPropertiesValue) GetName() *string {
	return s.Name
}

func (s *DentryAppPropertiesValue) GetValue() *string {
	return s.Value
}

func (s *DentryAppPropertiesValue) GetVisibility() *string {
	return s.Visibility
}

func (s *DentryAppPropertiesValue) SetName(v string) *DentryAppPropertiesValue {
	s.Name = &v
	return s
}

func (s *DentryAppPropertiesValue) SetValue(v string) *DentryAppPropertiesValue {
	s.Value = &v
	return s
}

func (s *DentryAppPropertiesValue) SetVisibility(v string) *DentryAppPropertiesValue {
	s.Visibility = &v
	return s
}

func (s *DentryAppPropertiesValue) Validate() error {
	return dara.Validate(s)
}
