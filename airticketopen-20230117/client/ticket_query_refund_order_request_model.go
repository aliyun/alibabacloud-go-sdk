// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryRefundOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *TicketQueryRefundOrderRequest
	GetAccountNo() *int64
	SetDistributorOrderId(v string) *TicketQueryRefundOrderRequest
	GetDistributorOrderId() *string
}

type TicketQueryRefundOrderRequest struct {
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
}

func (s TicketQueryRefundOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryRefundOrderRequest) GoString() string {
	return s.String()
}

func (s *TicketQueryRefundOrderRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *TicketQueryRefundOrderRequest) GetDistributorOrderId() *string {
	return s.DistributorOrderId
}

func (s *TicketQueryRefundOrderRequest) SetAccountNo(v int64) *TicketQueryRefundOrderRequest {
	s.AccountNo = &v
	return s
}

func (s *TicketQueryRefundOrderRequest) SetDistributorOrderId(v string) *TicketQueryRefundOrderRequest {
	s.DistributorOrderId = &v
	return s
}

func (s *TicketQueryRefundOrderRequest) Validate() error {
	return dara.Validate(s)
}
