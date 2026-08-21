// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryPriceStockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *TicketQueryPriceStockRequest
	GetAccountNo() *int64
	SetEndDate(v string) *TicketQueryPriceStockRequest
	GetEndDate() *string
	SetProductId(v string) *TicketQueryPriceStockRequest
	GetProductId() *string
	SetStartDate(v string) *TicketQueryPriceStockRequest
	GetStartDate() *string
}

type TicketQueryPriceStockRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// example:
	//
	// 2026-10-30
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 2026-10-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s TicketQueryPriceStockRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryPriceStockRequest) GoString() string {
	return s.String()
}

func (s *TicketQueryPriceStockRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *TicketQueryPriceStockRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *TicketQueryPriceStockRequest) GetProductId() *string {
	return s.ProductId
}

func (s *TicketQueryPriceStockRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *TicketQueryPriceStockRequest) SetAccountNo(v int64) *TicketQueryPriceStockRequest {
	s.AccountNo = &v
	return s
}

func (s *TicketQueryPriceStockRequest) SetEndDate(v string) *TicketQueryPriceStockRequest {
	s.EndDate = &v
	return s
}

func (s *TicketQueryPriceStockRequest) SetProductId(v string) *TicketQueryPriceStockRequest {
	s.ProductId = &v
	return s
}

func (s *TicketQueryPriceStockRequest) SetStartDate(v string) *TicketQueryPriceStockRequest {
	s.StartDate = &v
	return s
}

func (s *TicketQueryPriceStockRequest) Validate() error {
	return dara.Validate(s)
}
