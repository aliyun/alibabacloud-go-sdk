// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewVirtualBridgeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *RenewVirtualBridgeResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *RenewVirtualBridgeResponseBody
	GetRequestId() *string
}

type RenewVirtualBridgeResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 214552063030752
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewVirtualBridgeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewVirtualBridgeResponseBody) GoString() string {
	return s.String()
}

func (s *RenewVirtualBridgeResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *RenewVirtualBridgeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewVirtualBridgeResponseBody) SetOrderId(v int64) *RenewVirtualBridgeResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewVirtualBridgeResponseBody) SetRequestId(v string) *RenewVirtualBridgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewVirtualBridgeResponseBody) Validate() error {
	return dara.Validate(s)
}
