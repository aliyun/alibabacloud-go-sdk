// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketCheckRefundRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *TicketCheckRefundRequest
	GetAccountNo() *int64
	SetDistributorOrderId(v string) *TicketCheckRefundRequest
	GetDistributorOrderId() *string
	SetRefundReason(v string) *TicketCheckRefundRequest
	GetRefundReason() *string
	SetRefundRemark(v string) *TicketCheckRefundRequest
	GetRefundRemark() *string
}

type TicketCheckRefundRequest struct {
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

func (s TicketCheckRefundRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketCheckRefundRequest) GoString() string {
	return s.String()
}

func (s *TicketCheckRefundRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *TicketCheckRefundRequest) GetDistributorOrderId() *string {
	return s.DistributorOrderId
}

func (s *TicketCheckRefundRequest) GetRefundReason() *string {
	return s.RefundReason
}

func (s *TicketCheckRefundRequest) GetRefundRemark() *string {
	return s.RefundRemark
}

func (s *TicketCheckRefundRequest) SetAccountNo(v int64) *TicketCheckRefundRequest {
	s.AccountNo = &v
	return s
}

func (s *TicketCheckRefundRequest) SetDistributorOrderId(v string) *TicketCheckRefundRequest {
	s.DistributorOrderId = &v
	return s
}

func (s *TicketCheckRefundRequest) SetRefundReason(v string) *TicketCheckRefundRequest {
	s.RefundReason = &v
	return s
}

func (s *TicketCheckRefundRequest) SetRefundRemark(v string) *TicketCheckRefundRequest {
	s.RefundRemark = &v
	return s
}

func (s *TicketCheckRefundRequest) Validate() error {
	return dara.Validate(s)
}
