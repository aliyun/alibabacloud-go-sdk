// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseBotInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PurchaseBotInstanceResponseBody
	GetInstanceId() *string
	SetOrderId(v string) *PurchaseBotInstanceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *PurchaseBotInstanceResponseBody
	GetRequestId() *string
}

type PurchaseBotInstanceResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// esa-bot-9v9x3o2***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 31223****11
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PurchaseBotInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PurchaseBotInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PurchaseBotInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PurchaseBotInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *PurchaseBotInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PurchaseBotInstanceResponseBody) SetInstanceId(v string) *PurchaseBotInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *PurchaseBotInstanceResponseBody) SetOrderId(v string) *PurchaseBotInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *PurchaseBotInstanceResponseBody) SetRequestId(v string) *PurchaseBotInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PurchaseBotInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
