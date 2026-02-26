// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductSpecValue interface {
	dara.Model
	String() string
	GoString() string
	SetValue(v string) *ProductSpecValue
	GetValue() *string
	SetValueAlias(v string) *ProductSpecValue
	GetValueAlias() *string
	SetValueId(v int64) *ProductSpecValue
	GetValueId() *int64
}

type ProductSpecValue struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// example:
	//
	// 秘色
	ValueAlias *string `json:"valueAlias,omitempty" xml:"valueAlias,omitempty"`
	// example:
	//
	// 1000
	ValueId *int64 `json:"valueId,omitempty" xml:"valueId,omitempty"`
}

func (s ProductSpecValue) String() string {
	return dara.Prettify(s)
}

func (s ProductSpecValue) GoString() string {
	return s.String()
}

func (s *ProductSpecValue) GetValue() *string {
	return s.Value
}

func (s *ProductSpecValue) GetValueAlias() *string {
	return s.ValueAlias
}

func (s *ProductSpecValue) GetValueId() *int64 {
	return s.ValueId
}

func (s *ProductSpecValue) SetValue(v string) *ProductSpecValue {
	s.Value = &v
	return s
}

func (s *ProductSpecValue) SetValueAlias(v string) *ProductSpecValue {
	s.ValueAlias = &v
	return s
}

func (s *ProductSpecValue) SetValueId(v int64) *ProductSpecValue {
	s.ValueId = &v
	return s
}

func (s *ProductSpecValue) Validate() error {
	return dara.Validate(s)
}
