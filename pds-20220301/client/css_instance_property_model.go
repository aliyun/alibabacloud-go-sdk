// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCssInstanceProperty interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CssInstanceProperty
	GetCode() *string
	SetGlobalKey(v string) *CssInstanceProperty
	GetGlobalKey() *string
	SetName(v string) *CssInstanceProperty
	GetName() *string
	SetUnit(v string) *CssInstanceProperty
	GetUnit() *string
	SetValue(v string) *CssInstanceProperty
	GetValue() *string
}

type CssInstanceProperty struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	GlobalKey *string `json:"globalKey,omitempty" xml:"globalKey,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Unit      *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CssInstanceProperty) String() string {
	return dara.Prettify(s)
}

func (s CssInstanceProperty) GoString() string {
	return s.String()
}

func (s *CssInstanceProperty) GetCode() *string {
	return s.Code
}

func (s *CssInstanceProperty) GetGlobalKey() *string {
	return s.GlobalKey
}

func (s *CssInstanceProperty) GetName() *string {
	return s.Name
}

func (s *CssInstanceProperty) GetUnit() *string {
	return s.Unit
}

func (s *CssInstanceProperty) GetValue() *string {
	return s.Value
}

func (s *CssInstanceProperty) SetCode(v string) *CssInstanceProperty {
	s.Code = &v
	return s
}

func (s *CssInstanceProperty) SetGlobalKey(v string) *CssInstanceProperty {
	s.GlobalKey = &v
	return s
}

func (s *CssInstanceProperty) SetName(v string) *CssInstanceProperty {
	s.Name = &v
	return s
}

func (s *CssInstanceProperty) SetUnit(v string) *CssInstanceProperty {
	s.Unit = &v
	return s
}

func (s *CssInstanceProperty) SetValue(v string) *CssInstanceProperty {
	s.Value = &v
	return s
}

func (s *CssInstanceProperty) Validate() error {
	return dara.Validate(s)
}
