// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCacheReserveSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateCacheReserveSpecResponseBody
	GetInstanceId() *string
	SetOrderId(v string) *UpdateCacheReserveSpecResponseBody
	GetOrderId() *string
	SetRequestId(v string) *UpdateCacheReserveSpecResponseBody
	GetRequestId() *string
}

type UpdateCacheReserveSpecResponseBody struct {
	// Instance ID.
	//
	// example:
	//
	// esa-cr-9tuv*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Order ID.
	//
	// example:
	//
	// 2223332122***
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 40423A7F-A83D-1E24-B80E-86DD25790759
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCacheReserveSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCacheReserveSpecResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCacheReserveSpecResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCacheReserveSpecResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *UpdateCacheReserveSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCacheReserveSpecResponseBody) SetInstanceId(v string) *UpdateCacheReserveSpecResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateCacheReserveSpecResponseBody) SetOrderId(v string) *UpdateCacheReserveSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdateCacheReserveSpecResponseBody) SetRequestId(v string) *UpdateCacheReserveSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCacheReserveSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
