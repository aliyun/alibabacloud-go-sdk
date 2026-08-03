// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelPayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelPayRequest
	GetAccountNo() *int64
	SetOrderNo(v string) *GlobalHotelPayRequest
	GetOrderNo() *string
	SetTracerId(v string) *GlobalHotelPayRequest
	GetTracerId() *string
}

type GlobalHotelPayRequest struct {
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

func (s GlobalHotelPayRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelPayRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelPayRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelPayRequest) GetOrderNo() *string {
	return s.OrderNo
}

func (s *GlobalHotelPayRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelPayRequest) SetAccountNo(v int64) *GlobalHotelPayRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelPayRequest) SetOrderNo(v string) *GlobalHotelPayRequest {
	s.OrderNo = &v
	return s
}

func (s *GlobalHotelPayRequest) SetTracerId(v string) *GlobalHotelPayRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelPayRequest) Validate() error {
	return dara.Validate(s)
}
