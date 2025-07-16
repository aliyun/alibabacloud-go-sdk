// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDentriesAppPropertiesValue interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DentriesAppPropertiesValue
	GetName() *string
	SetValue(v string) *DentriesAppPropertiesValue
	GetValue() *string
	SetVisibility(v string) *DentriesAppPropertiesValue
	GetVisibility() *string
}

type DentriesAppPropertiesValue struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// PRIVATE
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DentriesAppPropertiesValue) String() string {
	return dara.Prettify(s)
}

func (s DentriesAppPropertiesValue) GoString() string {
	return s.String()
}

func (s *DentriesAppPropertiesValue) GetName() *string {
	return s.Name
}

func (s *DentriesAppPropertiesValue) GetValue() *string {
	return s.Value
}

func (s *DentriesAppPropertiesValue) GetVisibility() *string {
	return s.Visibility
}

func (s *DentriesAppPropertiesValue) SetName(v string) *DentriesAppPropertiesValue {
	s.Name = &v
	return s
}

func (s *DentriesAppPropertiesValue) SetValue(v string) *DentriesAppPropertiesValue {
	s.Value = &v
	return s
}

func (s *DentriesAppPropertiesValue) SetVisibility(v string) *DentriesAppPropertiesValue {
	s.Visibility = &v
	return s
}

func (s *DentriesAppPropertiesValue) Validate() error {
	return dara.Validate(s)
}
