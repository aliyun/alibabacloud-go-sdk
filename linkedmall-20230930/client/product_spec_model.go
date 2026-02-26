// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductSpec interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *ProductSpec
	GetKey() *string
	SetKeyId(v int64) *ProductSpec
	GetKeyId() *int64
	SetValues(v []*ProductSpecValue) *ProductSpec
	GetValues() []*ProductSpecValue
}

type ProductSpec struct {
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 1000
	KeyId  *int64              `json:"keyId,omitempty" xml:"keyId,omitempty"`
	Values []*ProductSpecValue `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ProductSpec) String() string {
	return dara.Prettify(s)
}

func (s ProductSpec) GoString() string {
	return s.String()
}

func (s *ProductSpec) GetKey() *string {
	return s.Key
}

func (s *ProductSpec) GetKeyId() *int64 {
	return s.KeyId
}

func (s *ProductSpec) GetValues() []*ProductSpecValue {
	return s.Values
}

func (s *ProductSpec) SetKey(v string) *ProductSpec {
	s.Key = &v
	return s
}

func (s *ProductSpec) SetKeyId(v int64) *ProductSpec {
	s.KeyId = &v
	return s
}

func (s *ProductSpec) SetValues(v []*ProductSpecValue) *ProductSpec {
	s.Values = v
	return s
}

func (s *ProductSpec) Validate() error {
	if s.Values != nil {
		for _, item := range s.Values {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
