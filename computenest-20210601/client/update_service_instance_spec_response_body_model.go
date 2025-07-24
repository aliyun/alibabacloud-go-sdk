// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *UpdateServiceInstanceSpecResponseBody
	GetOrderId() *string
	SetRequestId(v string) *UpdateServiceInstanceSpecResponseBody
	GetRequestId() *string
}

type UpdateServiceInstanceSpecResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 23396265896****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceInstanceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateServiceInstanceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceInstanceSpecResponseBody) SetOrderId(v string) *UpdateServiceInstanceSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdateServiceInstanceSpecResponseBody) SetRequestId(v string) *UpdateServiceInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceInstanceSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
