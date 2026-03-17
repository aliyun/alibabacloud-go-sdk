// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeSmartAccessGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *UpgradeSmartAccessGatewayResponseBody
	GetOrderId() *string
	SetRequestId(v string) *UpgradeSmartAccessGatewayResponseBody
	GetRequestId() *string
}

type UpgradeSmartAccessGatewayResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 203384676330296
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 45F07029-1783-4B2D-B4CE-45B9EAA58440
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeSmartAccessGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeSmartAccessGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeSmartAccessGatewayResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *UpgradeSmartAccessGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeSmartAccessGatewayResponseBody) SetOrderId(v string) *UpgradeSmartAccessGatewayResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeSmartAccessGatewayResponseBody) SetRequestId(v string) *UpgradeSmartAccessGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeSmartAccessGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
