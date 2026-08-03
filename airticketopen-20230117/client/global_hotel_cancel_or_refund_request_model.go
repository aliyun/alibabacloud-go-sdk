// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelCancelOrRefundRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelCancelOrRefundRequest
	GetAccountNo() *int64
	SetOrderNo(v string) *GlobalHotelCancelOrRefundRequest
	GetOrderNo() *string
	SetTracerId(v string) *GlobalHotelCancelOrRefundRequest
	GetTracerId() *string
}

type GlobalHotelCancelOrRefundRequest struct {
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

func (s GlobalHotelCancelOrRefundRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCancelOrRefundRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelCancelOrRefundRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelCancelOrRefundRequest) GetOrderNo() *string {
	return s.OrderNo
}

func (s *GlobalHotelCancelOrRefundRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCancelOrRefundRequest) SetAccountNo(v int64) *GlobalHotelCancelOrRefundRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelCancelOrRefundRequest) SetOrderNo(v string) *GlobalHotelCancelOrRefundRequest {
	s.OrderNo = &v
	return s
}

func (s *GlobalHotelCancelOrRefundRequest) SetTracerId(v string) *GlobalHotelCancelOrRefundRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCancelOrRefundRequest) Validate() error {
	return dara.Validate(s)
}
