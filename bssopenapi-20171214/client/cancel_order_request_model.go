// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *CancelOrderRequest
	GetOrderId() *string
	SetOwnerId(v int64) *CancelOrderRequest
	GetOwnerId() *int64
}

type CancelOrderRequest struct {
	// The ID of the order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 73465432785
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CancelOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelOrderRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *CancelOrderRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelOrderRequest) SetOrderId(v string) *CancelOrderRequest {
	s.OrderId = &v
	return s
}

func (s *CancelOrderRequest) SetOwnerId(v int64) *CancelOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelOrderRequest) Validate() error {
	return dara.Validate(s)
}
