// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeSmartAccessGatewaySoftwareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *UpgradeSmartAccessGatewaySoftwareResponseBody
	GetOrderId() *string
	SetRequestId(v string) *UpgradeSmartAccessGatewaySoftwareResponseBody
	GetRequestId() *string
}

type UpgradeSmartAccessGatewaySoftwareResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 20697688135****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 97A4F8A5-603E-4C3B-A91E-17CD87090EA9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeSmartAccessGatewaySoftwareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeSmartAccessGatewaySoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeSmartAccessGatewaySoftwareResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *UpgradeSmartAccessGatewaySoftwareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeSmartAccessGatewaySoftwareResponseBody) SetOrderId(v string) *UpgradeSmartAccessGatewaySoftwareResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeSmartAccessGatewaySoftwareResponseBody) SetRequestId(v string) *UpgradeSmartAccessGatewaySoftwareResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeSmartAccessGatewaySoftwareResponseBody) Validate() error {
	return dara.Validate(s)
}
