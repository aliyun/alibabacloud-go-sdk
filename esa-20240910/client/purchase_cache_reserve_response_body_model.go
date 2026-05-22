// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseCacheReserveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PurchaseCacheReserveResponseBody
	GetInstanceId() *string
	SetOrderId(v string) *PurchaseCacheReserveResponseBody
	GetOrderId() *string
	SetRequestId(v string) *PurchaseCacheReserveResponseBody
	GetRequestId() *string
}

type PurchaseCacheReserveResponseBody struct {
	// Instance ID.
	//
	// example:
	//
	// xcdn-ad*****s11w
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Order ID.
	//
	// example:
	//
	// 31223****11
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 30423A7F-A83D-1E24-B80E-86DD25790758
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PurchaseCacheReserveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PurchaseCacheReserveResponseBody) GoString() string {
	return s.String()
}

func (s *PurchaseCacheReserveResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PurchaseCacheReserveResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *PurchaseCacheReserveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PurchaseCacheReserveResponseBody) SetInstanceId(v string) *PurchaseCacheReserveResponseBody {
	s.InstanceId = &v
	return s
}

func (s *PurchaseCacheReserveResponseBody) SetOrderId(v string) *PurchaseCacheReserveResponseBody {
	s.OrderId = &v
	return s
}

func (s *PurchaseCacheReserveResponseBody) SetRequestId(v string) *PurchaseCacheReserveResponseBody {
	s.RequestId = &v
	return s
}

func (s *PurchaseCacheReserveResponseBody) Validate() error {
	return dara.Validate(s)
}
