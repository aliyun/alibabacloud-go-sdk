// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketPayOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *TicketPayOrderRequest
	GetAccountNo() *int64
	SetDistributorOrderId(v string) *TicketPayOrderRequest
	GetDistributorOrderId() *string
}

type TicketPayOrderRequest struct {
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

func (s TicketPayOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketPayOrderRequest) GoString() string {
	return s.String()
}

func (s *TicketPayOrderRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *TicketPayOrderRequest) GetDistributorOrderId() *string {
	return s.DistributorOrderId
}

func (s *TicketPayOrderRequest) SetAccountNo(v int64) *TicketPayOrderRequest {
	s.AccountNo = &v
	return s
}

func (s *TicketPayOrderRequest) SetDistributorOrderId(v string) *TicketPayOrderRequest {
	s.DistributorOrderId = &v
	return s
}

func (s *TicketPayOrderRequest) Validate() error {
	return dara.Validate(s)
}
