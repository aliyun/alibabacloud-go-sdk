// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrderDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *GetOrderDetailRequest
	GetOrderId() *string
	SetOwnerId(v int64) *GetOrderDetailRequest
	GetOwnerId() *int64
}

type GetOrderDetailRequest struct {
	// The order ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 32453453
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s GetOrderDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetOrderDetailRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *GetOrderDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetOrderDetailRequest) SetOrderId(v string) *GetOrderDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetOrderDetailRequest) SetOwnerId(v int64) *GetOrderDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *GetOrderDetailRequest) Validate() error {
	return dara.Validate(s)
}
