// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyRefundRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *ApplyRefundRequest
	GetAccountNo() *int64
	SetOrderNo(v string) *ApplyRefundRequest
	GetOrderNo() *string
	SetRefundReason(v string) *ApplyRefundRequest
	GetRefundReason() *string
	SetTracerId(v string) *ApplyRefundRequest
	GetTracerId() *string
}

type ApplyRefundRequest struct {
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
	// GUEST_REQUEST
	RefundReason *string `json:"RefundReason,omitempty" xml:"RefundReason,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ApplyRefundRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyRefundRequest) GoString() string {
	return s.String()
}

func (s *ApplyRefundRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *ApplyRefundRequest) GetOrderNo() *string {
	return s.OrderNo
}

func (s *ApplyRefundRequest) GetRefundReason() *string {
	return s.RefundReason
}

func (s *ApplyRefundRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *ApplyRefundRequest) SetAccountNo(v int64) *ApplyRefundRequest {
	s.AccountNo = &v
	return s
}

func (s *ApplyRefundRequest) SetOrderNo(v string) *ApplyRefundRequest {
	s.OrderNo = &v
	return s
}

func (s *ApplyRefundRequest) SetRefundReason(v string) *ApplyRefundRequest {
	s.RefundReason = &v
	return s
}

func (s *ApplyRefundRequest) SetTracerId(v string) *ApplyRefundRequest {
	s.TracerId = &v
	return s
}

func (s *ApplyRefundRequest) Validate() error {
	return dara.Validate(s)
}
