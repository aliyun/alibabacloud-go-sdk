// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpDeliveryQueryParamValue interface {
	dara.Model
	String() string
	GoString() string
	SetStaticValue(v string) *HttpDeliveryQueryParamValue
	GetStaticValue() *string
}

type HttpDeliveryQueryParamValue struct {
	// The value of the custom HTTP delivery query string parameter.
	//
	// > Key-map. The value can be a static value, a dynamic function, or a dynamic value.
	//
	// example:
	//
	// auth_token: sk-***
	StaticValue *string `json:"StaticValue,omitempty" xml:"StaticValue,omitempty"`
}

func (s HttpDeliveryQueryParamValue) String() string {
	return dara.Prettify(s)
}

func (s HttpDeliveryQueryParamValue) GoString() string {
	return s.String()
}

func (s *HttpDeliveryQueryParamValue) GetStaticValue() *string {
	return s.StaticValue
}

func (s *HttpDeliveryQueryParamValue) SetStaticValue(v string) *HttpDeliveryQueryParamValue {
	s.StaticValue = &v
	return s
}

func (s *HttpDeliveryQueryParamValue) Validate() error {
	return dara.Validate(s)
}
