// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderNum(v int64) *OrderDetailRequest
	GetOrderNum() *int64
	SetOutOrderNum(v string) *OrderDetailRequest
	GetOutOrderNum() *string
}

type OrderDetailRequest struct {
	// order number created by book
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// external order number(customized by buyer when book)
	//
	// example:
	//
	// x091-2023-0220-j-0001
	OutOrderNum *string `json:"out_order_num,omitempty" xml:"out_order_num,omitempty"`
}

func (s OrderDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailRequest) GoString() string {
	return s.String()
}

func (s *OrderDetailRequest) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *OrderDetailRequest) GetOutOrderNum() *string {
	return s.OutOrderNum
}

func (s *OrderDetailRequest) SetOrderNum(v int64) *OrderDetailRequest {
	s.OrderNum = &v
	return s
}

func (s *OrderDetailRequest) SetOutOrderNum(v string) *OrderDetailRequest {
	s.OutOrderNum = &v
	return s
}

func (s *OrderDetailRequest) Validate() error {
	return dara.Validate(s)
}
