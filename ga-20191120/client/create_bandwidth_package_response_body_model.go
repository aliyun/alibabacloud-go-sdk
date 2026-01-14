// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBandwidthPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *CreateBandwidthPackageResponseBody
	GetBandwidthPackageId() *string
	SetOrderId(v string) *CreateBandwidthPackageResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateBandwidthPackageResponseBody
	GetRequestId() *string
}

type CreateBandwidthPackageResponseBody struct {
	// The ID of the bandwidth plan.
	//
	// example:
	//
	// gbwp-bp1sgzldyj6b4q7cx****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The order ID.
	//
	// If bills are not automatically paid, you must go to the Order Center to complete the payments.
	//
	// This parameter is returned only if ChargeType is set to PREPAY. If AutoPay is set to false, you must go to the [Order Center](https://usercenter2-intl.aliyun.com/order/list) to complete the payment.
	//
	// example:
	//
	// 208257****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4B6DBBB0-2D01-4C6A-A384-4129266E6B78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBandwidthPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBandwidthPackageResponseBody) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *CreateBandwidthPackageResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateBandwidthPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBandwidthPackageResponseBody) SetBandwidthPackageId(v string) *CreateBandwidthPackageResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *CreateBandwidthPackageResponseBody) SetOrderId(v string) *CreateBandwidthPackageResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateBandwidthPackageResponseBody) SetRequestId(v string) *CreateBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBandwidthPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
