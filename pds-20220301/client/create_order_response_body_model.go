// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateOrderResponseBody
	GetInstanceId() *string
	SetOrderId(v string) *CreateOrderResponseBody
	GetOrderId() *string
}

type CreateOrderResponseBody struct {
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	OrderId    *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateOrderResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateOrderResponseBody) SetInstanceId(v string) *CreateOrderResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateOrderResponseBody) SetOrderId(v string) *CreateOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
