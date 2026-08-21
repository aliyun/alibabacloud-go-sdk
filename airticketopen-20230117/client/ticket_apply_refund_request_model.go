// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketApplyRefundRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *TicketApplyRefundRequest
	GetAccountNo() *int64
	SetDistributorOrderId(v string) *TicketApplyRefundRequest
	GetDistributorOrderId() *string
	SetRefundReason(v string) *TicketApplyRefundRequest
	GetRefundReason() *string
	SetRefundRemark(v string) *TicketApplyRefundRequest
	GetRefundRemark() *string
}

type TicketApplyRefundRequest struct {
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
	// 123456
	DistributorOrderId *string `json:"DistributorOrderId,omitempty" xml:"DistributorOrderId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 计划有变
	RefundReason *string `json:"RefundReason,omitempty" xml:"RefundReason,omitempty"`
	// example:
	//
	// 档期冲突
	RefundRemark *string `json:"RefundRemark,omitempty" xml:"RefundRemark,omitempty"`
}

func (s TicketApplyRefundRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketApplyRefundRequest) GoString() string {
	return s.String()
}

func (s *TicketApplyRefundRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *TicketApplyRefundRequest) GetDistributorOrderId() *string {
	return s.DistributorOrderId
}

func (s *TicketApplyRefundRequest) GetRefundReason() *string {
	return s.RefundReason
}

func (s *TicketApplyRefundRequest) GetRefundRemark() *string {
	return s.RefundRemark
}

func (s *TicketApplyRefundRequest) SetAccountNo(v int64) *TicketApplyRefundRequest {
	s.AccountNo = &v
	return s
}

func (s *TicketApplyRefundRequest) SetDistributorOrderId(v string) *TicketApplyRefundRequest {
	s.DistributorOrderId = &v
	return s
}

func (s *TicketApplyRefundRequest) SetRefundReason(v string) *TicketApplyRefundRequest {
	s.RefundReason = &v
	return s
}

func (s *TicketApplyRefundRequest) SetRefundRemark(v string) *TicketApplyRefundRequest {
	s.RefundRemark = &v
	return s
}

func (s *TicketApplyRefundRequest) Validate() error {
	return dara.Validate(s)
}
