// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseDDoSInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PurchaseDDoSInstanceResponseBody
	GetInstanceId() *string
	SetOrderId(v string) *PurchaseDDoSInstanceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *PurchaseDDoSInstanceResponseBody
	GetRequestId() *string
}

type PurchaseDDoSInstanceResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// esa-ddos-9tuv*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2223332122***
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PurchaseDDoSInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PurchaseDDoSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PurchaseDDoSInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PurchaseDDoSInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *PurchaseDDoSInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PurchaseDDoSInstanceResponseBody) SetInstanceId(v string) *PurchaseDDoSInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *PurchaseDDoSInstanceResponseBody) SetOrderId(v string) *PurchaseDDoSInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *PurchaseDDoSInstanceResponseBody) SetRequestId(v string) *PurchaseDDoSInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PurchaseDDoSInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
