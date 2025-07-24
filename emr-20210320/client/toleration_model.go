// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToleration interface {
	dara.Model
	String() string
	GoString() string
	SetEffect(v string) *Toleration
	GetEffect() *string
	SetKey(v string) *Toleration
	GetKey() *string
	SetOperator(v string) *Toleration
	GetOperator() *string
	SetValue(v string) *Toleration
	GetValue() *string
}

type Toleration struct {
	Effect   *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Toleration) String() string {
	return dara.Prettify(s)
}

func (s Toleration) GoString() string {
	return s.String()
}

func (s *Toleration) GetEffect() *string {
	return s.Effect
}

func (s *Toleration) GetKey() *string {
	return s.Key
}

func (s *Toleration) GetOperator() *string {
	return s.Operator
}

func (s *Toleration) GetValue() *string {
	return s.Value
}

func (s *Toleration) SetEffect(v string) *Toleration {
	s.Effect = &v
	return s
}

func (s *Toleration) SetKey(v string) *Toleration {
	s.Key = &v
	return s
}

func (s *Toleration) SetOperator(v string) *Toleration {
	s.Operator = &v
	return s
}

func (s *Toleration) SetValue(v string) *Toleration {
	s.Value = &v
	return s
}

func (s *Toleration) Validate() error {
	return dara.Validate(s)
}
