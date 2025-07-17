// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiBackendMatchCondition interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *HttpApiBackendMatchCondition
	GetKey() *string
	SetOperator(v string) *HttpApiBackendMatchCondition
	GetOperator() *string
	SetType(v string) *HttpApiBackendMatchCondition
	GetType() *string
	SetValue(v string) *HttpApiBackendMatchCondition
	GetValue() *string
}

type HttpApiBackendMatchCondition struct {
	// example:
	//
	// color
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// equal
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// Query
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// gray
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HttpApiBackendMatchCondition) String() string {
	return dara.Prettify(s)
}

func (s HttpApiBackendMatchCondition) GoString() string {
	return s.String()
}

func (s *HttpApiBackendMatchCondition) GetKey() *string {
	return s.Key
}

func (s *HttpApiBackendMatchCondition) GetOperator() *string {
	return s.Operator
}

func (s *HttpApiBackendMatchCondition) GetType() *string {
	return s.Type
}

func (s *HttpApiBackendMatchCondition) GetValue() *string {
	return s.Value
}

func (s *HttpApiBackendMatchCondition) SetKey(v string) *HttpApiBackendMatchCondition {
	s.Key = &v
	return s
}

func (s *HttpApiBackendMatchCondition) SetOperator(v string) *HttpApiBackendMatchCondition {
	s.Operator = &v
	return s
}

func (s *HttpApiBackendMatchCondition) SetType(v string) *HttpApiBackendMatchCondition {
	s.Type = &v
	return s
}

func (s *HttpApiBackendMatchCondition) SetValue(v string) *HttpApiBackendMatchCondition {
	s.Value = &v
	return s
}

func (s *HttpApiBackendMatchCondition) Validate() error {
	return dara.Validate(s)
}
