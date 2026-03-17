// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradeSmartAccessGatewaySoftwareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *DowngradeSmartAccessGatewaySoftwareResponseBody
	GetOrderId() *string
	SetRequestId(v string) *DowngradeSmartAccessGatewaySoftwareResponseBody
	GetRequestId() *string
}

type DowngradeSmartAccessGatewaySoftwareResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 204595234160786
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3145AF24-1A5E-4AB7-90DA-7201FDD90B8D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DowngradeSmartAccessGatewaySoftwareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DowngradeSmartAccessGatewaySoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *DowngradeSmartAccessGatewaySoftwareResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *DowngradeSmartAccessGatewaySoftwareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DowngradeSmartAccessGatewaySoftwareResponseBody) SetOrderId(v string) *DowngradeSmartAccessGatewaySoftwareResponseBody {
	s.OrderId = &v
	return s
}

func (s *DowngradeSmartAccessGatewaySoftwareResponseBody) SetRequestId(v string) *DowngradeSmartAccessGatewaySoftwareResponseBody {
	s.RequestId = &v
	return s
}

func (s *DowngradeSmartAccessGatewaySoftwareResponseBody) Validate() error {
	return dara.Validate(s)
}
