// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *CreateOrderResponseBody
	GetInstanceIds() []*string
	SetOrderId(v string) *CreateOrderResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateOrderResponseBody
	GetRequestId() *string
}

type CreateOrderResponseBody struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 12809858
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 054AF571-C86F-533F-8A7B-F62206FA4367
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateOrderResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrderResponseBody) SetInstanceIds(v []*string) *CreateOrderResponseBody {
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
