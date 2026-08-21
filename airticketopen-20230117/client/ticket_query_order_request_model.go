// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *TicketQueryOrderRequest
	GetAccountNo() *int64
	SetDistributorOrderId(v string) *TicketQueryOrderRequest
	GetDistributorOrderId() *string
}

type TicketQueryOrderRequest struct {
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

func (s TicketQueryOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryOrderRequest) GoString() string {
	return s.String()
}

func (s *TicketQueryOrderRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *TicketQueryOrderRequest) GetDistributorOrderId() *string {
	return s.DistributorOrderId
}

func (s *TicketQueryOrderRequest) SetAccountNo(v int64) *TicketQueryOrderRequest {
	s.AccountNo = &v
	return s
}

func (s *TicketQueryOrderRequest) SetDistributorOrderId(v string) *TicketQueryOrderRequest {
	s.DistributorOrderId = &v
	return s
}

func (s *TicketQueryOrderRequest) Validate() error {
	return dara.Validate(s)
}
