// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAdditionalBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *RenewAdditionalBandwidthResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RenewAdditionalBandwidthResponseBody
	GetRequestId() *string
}

type RenewAdditionalBandwidthResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 2084452111111
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewAdditionalBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewAdditionalBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *RenewAdditionalBandwidthResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewAdditionalBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewAdditionalBandwidthResponseBody) SetOrderId(v string) *RenewAdditionalBandwidthResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewAdditionalBandwidthResponseBody) SetRequestId(v string) *RenewAdditionalBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewAdditionalBandwidthResponseBody) Validate() error {
	return dara.Validate(s)
}
