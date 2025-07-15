// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDesktopsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *RenewDesktopsResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RenewDesktopsResponseBody
	GetRequestId() *string
}

type RenewDesktopsResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 20844399755****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewDesktopsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *RenewDesktopsResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewDesktopsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewDesktopsResponseBody) SetOrderId(v string) *RenewDesktopsResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewDesktopsResponseBody) SetRequestId(v string) *RenewDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewDesktopsResponseBody) Validate() error {
	return dara.Validate(s)
}
