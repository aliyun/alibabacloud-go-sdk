// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewNetworkPackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *RenewNetworkPackagesResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RenewNetworkPackagesResponseBody
	GetRequestId() *string
}

type RenewNetworkPackagesResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 214726268900640
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewNetworkPackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewNetworkPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *RenewNetworkPackagesResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewNetworkPackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewNetworkPackagesResponseBody) SetOrderId(v string) *RenewNetworkPackagesResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewNetworkPackagesResponseBody) SetRequestId(v string) *RenewNetworkPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewNetworkPackagesResponseBody) Validate() error {
	return dara.Validate(s)
}
