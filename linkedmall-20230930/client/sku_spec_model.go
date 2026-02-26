// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkuSpec interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *SkuSpec
	GetKey() *string
	SetKeyId(v int64) *SkuSpec
	GetKeyId() *int64
	SetValue(v string) *SkuSpec
	GetValue() *string
	SetValueAlias(v string) *SkuSpec
	GetValueAlias() *string
	SetValueId(v int64) *SkuSpec
	GetValueId() *int64
}

type SkuSpec struct {
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 1000
	KeyId *int64  `json:"keyId,omitempty" xml:"keyId,omitempty"`
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

func (s SkuSpec) String() string {
	return dara.Prettify(s)
}

func (s SkuSpec) GoString() string {
	return s.String()
}

func (s *SkuSpec) GetKey() *string {
	return s.Key
}

func (s *SkuSpec) GetKeyId() *int64 {
	return s.KeyId
}

func (s *SkuSpec) GetValue() *string {
	return s.Value
}

func (s *SkuSpec) GetValueAlias() *string {
	return s.ValueAlias
}

func (s *SkuSpec) GetValueId() *int64 {
	return s.ValueId
}

func (s *SkuSpec) SetKey(v string) *SkuSpec {
	s.Key = &v
	return s
}

func (s *SkuSpec) SetKeyId(v int64) *SkuSpec {
	s.KeyId = &v
	return s
}

func (s *SkuSpec) SetValue(v string) *SkuSpec {
	s.Value = &v
	return s
}

func (s *SkuSpec) SetValueAlias(v string) *SkuSpec {
	s.ValueAlias = &v
	return s
}

func (s *SkuSpec) SetValueId(v int64) *SkuSpec {
	s.ValueId = &v
	return s
}

func (s *SkuSpec) Validate() error {
	return dara.Validate(s)
}
