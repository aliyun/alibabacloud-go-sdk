// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseRatePlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PurchaseRatePlanResponseBody
	GetInstanceId() *string
	SetOrderId(v string) *PurchaseRatePlanResponseBody
	GetOrderId() *string
	SetRequestId(v string) *PurchaseRatePlanResponseBody
	GetRequestId() *string
}

type PurchaseRatePlanResponseBody struct {
	// Instance ID.
	//
	// example:
	//
	// esa-site-ads11w
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Order ID.
	//
	// example:
	//
	// 123123
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 30423A7F-A83D-1E24-B80E-86DD25790758
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PurchaseRatePlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PurchaseRatePlanResponseBody) GoString() string {
	return s.String()
}

func (s *PurchaseRatePlanResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PurchaseRatePlanResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *PurchaseRatePlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PurchaseRatePlanResponseBody) SetInstanceId(v string) *PurchaseRatePlanResponseBody {
	s.InstanceId = &v
	return s
}

func (s *PurchaseRatePlanResponseBody) SetOrderId(v string) *PurchaseRatePlanResponseBody {
	s.OrderId = &v
	return s
}

func (s *PurchaseRatePlanResponseBody) SetRequestId(v string) *PurchaseRatePlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *PurchaseRatePlanResponseBody) Validate() error {
	return dara.Validate(s)
}
