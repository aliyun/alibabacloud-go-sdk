// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductExtendProperty interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *ProductExtendProperty
	GetKey() *string
	SetValue(v string) *ProductExtendProperty
	GetValue() *string
}

type ProductExtendProperty struct {
	// example:
	//
	// ss_picture_scene
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ProductExtendProperty) String() string {
	return dara.Prettify(s)
}

func (s ProductExtendProperty) GoString() string {
	return s.String()
}

func (s *ProductExtendProperty) GetKey() *string {
	return s.Key
}

func (s *ProductExtendProperty) GetValue() *string {
	return s.Value
}

func (s *ProductExtendProperty) SetKey(v string) *ProductExtendProperty {
	s.Key = &v
	return s
}

func (s *ProductExtendProperty) SetValue(v string) *ProductExtendProperty {
	s.Value = &v
	return s
}

func (s *ProductExtendProperty) Validate() error {
	return dara.Validate(s)
}
