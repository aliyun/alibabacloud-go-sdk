// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrder interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *Order
	GetCreateTime() *string
	SetOrderId(v string) *Order
	GetOrderId() *string
}

type Order struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s Order) String() string {
	return dara.Prettify(s)
}

func (s Order) GoString() string {
	return s.String()
}

func (s *Order) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Order) GetOrderId() *string {
	return s.OrderId
}

func (s *Order) SetCreateTime(v string) *Order {
	s.CreateTime = &v
	return s
}

func (s *Order) SetOrderId(v string) *Order {
	s.OrderId = &v
	return s
}

func (s *Order) Validate() error {
	return dara.Validate(s)
}
