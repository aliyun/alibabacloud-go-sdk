// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *RenewDBInstanceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RenewDBInstanceResponseBody
	GetRequestId() *string
}

type RenewDBInstanceResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 203317xxxxxxxx
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B118EF45-9633-4EE3-8405-42ED4373721B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewDBInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewDBInstanceResponseBody) SetOrderId(v string) *RenewDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewDBInstanceResponseBody) SetRequestId(v string) *RenewDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
