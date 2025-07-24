// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttribute interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *Attribute
	GetKey() *string
	SetValue(v string) *Attribute
	GetValue() *string
}

type Attribute struct {
	// 键。
	//
	// example:
	//
	// currentYarnRangerPluginState
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 值。
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Attribute) String() string {
	return dara.Prettify(s)
}

func (s Attribute) GoString() string {
	return s.String()
}

func (s *Attribute) GetKey() *string {
	return s.Key
}

func (s *Attribute) GetValue() *string {
	return s.Value
}

func (s *Attribute) SetKey(v string) *Attribute {
	s.Key = &v
	return s
}

func (s *Attribute) SetValue(v string) *Attribute {
	s.Value = &v
	return s
}

func (s *Attribute) Validate() error {
	return dara.Validate(s)
}
