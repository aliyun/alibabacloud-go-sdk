// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpDeliveryHeaderParamValue interface {
	dara.Model
	String() string
	GoString() string
	SetStaticValue(v string) *HttpDeliveryHeaderParamValue
	GetStaticValue() *string
}

type HttpDeliveryHeaderParamValue struct {
	// The static variable.
	//
	// example:
	//
	// alicdn
	StaticValue *string `json:"StaticValue,omitempty" xml:"StaticValue,omitempty"`
}

func (s HttpDeliveryHeaderParamValue) String() string {
	return dara.Prettify(s)
}

func (s HttpDeliveryHeaderParamValue) GoString() string {
	return s.String()
}

func (s *HttpDeliveryHeaderParamValue) GetStaticValue() *string {
	return s.StaticValue
}

func (s *HttpDeliveryHeaderParamValue) SetStaticValue(v string) *HttpDeliveryHeaderParamValue {
	s.StaticValue = &v
	return s
}

func (s *HttpDeliveryHeaderParamValue) Validate() error {
	return dara.Validate(s)
}
