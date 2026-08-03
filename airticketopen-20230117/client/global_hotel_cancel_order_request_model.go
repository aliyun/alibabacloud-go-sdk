// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelCancelOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelCancelOrderRequest
	GetAccountNo() *int64
	SetOrderNo(v string) *GlobalHotelCancelOrderRequest
	GetOrderNo() *string
	SetTracerId(v string) *GlobalHotelCancelOrderRequest
	GetTracerId() *string
}

type GlobalHotelCancelOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SO202606290001
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelCancelOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCancelOrderRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelCancelOrderRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelCancelOrderRequest) GetOrderNo() *string {
	return s.OrderNo
}

func (s *GlobalHotelCancelOrderRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCancelOrderRequest) SetAccountNo(v int64) *GlobalHotelCancelOrderRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelCancelOrderRequest) SetOrderNo(v string) *GlobalHotelCancelOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *GlobalHotelCancelOrderRequest) SetTracerId(v string) *GlobalHotelCancelOrderRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCancelOrderRequest) Validate() error {
	return dara.Validate(s)
}
