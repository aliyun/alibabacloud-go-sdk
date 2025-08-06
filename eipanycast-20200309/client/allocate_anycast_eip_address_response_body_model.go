// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateAnycastEipAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnycastId(v string) *AllocateAnycastEipAddressResponseBody
	GetAnycastId() *string
	SetOrderId(v string) *AllocateAnycastEipAddressResponseBody
	GetOrderId() *string
	SetRequestId(v string) *AllocateAnycastEipAddressResponseBody
	GetRequestId() *string
}

type AllocateAnycastEipAddressResponseBody struct {
	// The ID of the Anycast EIP.
	//
	// example:
	//
	// aeip-bp1ix34fralt4ykf3****
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 1422000****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FBDB18D8-E91E-4978-8D6C-6E2E3EE10133
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateAnycastEipAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateAnycastEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateAnycastEipAddressResponseBody) GetAnycastId() *string {
	return s.AnycastId
}

func (s *AllocateAnycastEipAddressResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *AllocateAnycastEipAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateAnycastEipAddressResponseBody) SetAnycastId(v string) *AllocateAnycastEipAddressResponseBody {
	s.AnycastId = &v
	return s
}

func (s *AllocateAnycastEipAddressResponseBody) SetOrderId(v string) *AllocateAnycastEipAddressResponseBody {
	s.OrderId = &v
	return s
}

func (s *AllocateAnycastEipAddressResponseBody) SetRequestId(v string) *AllocateAnycastEipAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateAnycastEipAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
