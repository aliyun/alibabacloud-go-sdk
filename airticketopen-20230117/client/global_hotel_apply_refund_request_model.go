// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelApplyRefundRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelApplyRefundRequest
	GetAccountNo() *int64
	SetOrderNo(v string) *GlobalHotelApplyRefundRequest
	GetOrderNo() *string
	SetRefundReason(v string) *GlobalHotelApplyRefundRequest
	GetRefundReason() *string
	SetTracerId(v string) *GlobalHotelApplyRefundRequest
	GetTracerId() *string
}

type GlobalHotelApplyRefundRequest struct {
	// The distributor account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// The sales order number.
	//
	// This parameter is required.
	//
	// example:
	//
	// SO202606290001
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	// The refund reason.
	//
	// example:
	//
	// GUEST_REQUEST
	RefundReason *string `json:"RefundReason,omitempty" xml:"RefundReason,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelApplyRefundRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelApplyRefundRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelApplyRefundRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelApplyRefundRequest) GetOrderNo() *string {
	return s.OrderNo
}

func (s *GlobalHotelApplyRefundRequest) GetRefundReason() *string {
	return s.RefundReason
}

func (s *GlobalHotelApplyRefundRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelApplyRefundRequest) SetAccountNo(v int64) *GlobalHotelApplyRefundRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelApplyRefundRequest) SetOrderNo(v string) *GlobalHotelApplyRefundRequest {
	s.OrderNo = &v
	return s
}

func (s *GlobalHotelApplyRefundRequest) SetRefundReason(v string) *GlobalHotelApplyRefundRequest {
	s.RefundReason = &v
	return s
}

func (s *GlobalHotelApplyRefundRequest) SetTracerId(v string) *GlobalHotelApplyRefundRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelApplyRefundRequest) Validate() error {
	return dara.Validate(s)
}
