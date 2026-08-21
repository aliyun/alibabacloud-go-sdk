// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketCancelOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *TicketCancelOrderRequest
	GetAccountNo() *int64
	SetDistributorOrderId(v string) *TicketCancelOrderRequest
	GetDistributorOrderId() *string
}

type TicketCancelOrderRequest struct {
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

func (s TicketCancelOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketCancelOrderRequest) GoString() string {
	return s.String()
}

func (s *TicketCancelOrderRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *TicketCancelOrderRequest) GetDistributorOrderId() *string {
	return s.DistributorOrderId
}

func (s *TicketCancelOrderRequest) SetAccountNo(v int64) *TicketCancelOrderRequest {
	s.AccountNo = &v
	return s
}

func (s *TicketCancelOrderRequest) SetDistributorOrderId(v string) *TicketCancelOrderRequest {
	s.DistributorOrderId = &v
	return s
}

func (s *TicketCancelOrderRequest) Validate() error {
	return dara.Validate(s)
}
