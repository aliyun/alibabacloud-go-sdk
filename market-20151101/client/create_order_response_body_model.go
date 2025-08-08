// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v *CreateOrderResponseBodyInstanceIds) *CreateOrderResponseBody
	GetInstanceIds() *CreateOrderResponseBodyInstanceIds
	SetOrderId(v string) *CreateOrderResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateOrderResponseBody
	GetRequestId() *string
}

type CreateOrderResponseBody struct {
	InstanceIds *CreateOrderResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	// example:
	//
	// 202********0117
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 4ca591b5-bc30-4fa7-aeaa-c4d0ec5d24ed
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) GetInstanceIds() *CreateOrderResponseBodyInstanceIds {
	return s.InstanceIds
}

func (s *CreateOrderResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrderResponseBody) SetInstanceIds(v *CreateOrderResponseBodyInstanceIds) *CreateOrderResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateOrderResponseBody) SetOrderId(v string) *CreateOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateOrderResponseBody) SetRequestId(v string) *CreateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateOrderResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s CreateOrderResponseBodyInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBodyInstanceIds) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *CreateOrderResponseBodyInstanceIds) SetInstanceId(v []*string) *CreateOrderResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

func (s *CreateOrderResponseBodyInstanceIds) Validate() error {
	return dara.Validate(s)
}
