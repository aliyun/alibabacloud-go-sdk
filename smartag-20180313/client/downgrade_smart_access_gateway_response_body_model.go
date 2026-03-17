// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradeSmartAccessGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *DowngradeSmartAccessGatewayResponseBody
	GetOrderId() *string
	SetRequestId(v string) *DowngradeSmartAccessGatewayResponseBody
	GetRequestId() *string
}

type DowngradeSmartAccessGatewayResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 20337777855****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A6B9EB0F-57DB-4843-A372-04678ABF490E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DowngradeSmartAccessGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DowngradeSmartAccessGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DowngradeSmartAccessGatewayResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *DowngradeSmartAccessGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DowngradeSmartAccessGatewayResponseBody) SetOrderId(v string) *DowngradeSmartAccessGatewayResponseBody {
	s.OrderId = &v
	return s
}

func (s *DowngradeSmartAccessGatewayResponseBody) SetRequestId(v string) *DowngradeSmartAccessGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DowngradeSmartAccessGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
