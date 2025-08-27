// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightBillSettlementQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *FlightBillSettlementQueryRequest
	GetBillBatch() *string
	SetOrderId(v int64) *FlightBillSettlementQueryRequest
	GetOrderId() *int64
	SetPageNo(v int32) *FlightBillSettlementQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *FlightBillSettlementQueryRequest
	GetPageSize() *int32
	SetPeriodEnd(v string) *FlightBillSettlementQueryRequest
	GetPeriodEnd() *string
	SetPeriodStart(v string) *FlightBillSettlementQueryRequest
	GetPeriodStart() *string
	SetScrollId(v string) *FlightBillSettlementQueryRequest
	GetScrollId() *string
	SetScrollMod(v bool) *FlightBillSettlementQueryRequest
	GetScrollMod() *bool
}

type FlightBillSettlementQueryRequest struct {
	BillBatch *string `json:"bill_batch,omitempty" xml:"bill_batch,omitempty"`
	OrderId   *int64  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 2021-10-01
	PeriodEnd *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	// example:
	//
	// 2021-10-01
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
	ScrollId    *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	ScrollMod   *bool   `json:"scroll_mod,omitempty" xml:"scroll_mod,omitempty"`
}

func (s FlightBillSettlementQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *FlightBillSettlementQueryRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *FlightBillSettlementQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightBillSettlementQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *FlightBillSettlementQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *FlightBillSettlementQueryRequest) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *FlightBillSettlementQueryRequest) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *FlightBillSettlementQueryRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *FlightBillSettlementQueryRequest) GetScrollMod() *bool {
	return s.ScrollMod
}

func (s *FlightBillSettlementQueryRequest) SetBillBatch(v string) *FlightBillSettlementQueryRequest {
	s.BillBatch = &v
	return s
}

func (s *FlightBillSettlementQueryRequest) SetOrderId(v int64) *FlightBillSettlementQueryRequest {
	s.OrderId = &v
	return s
}

func (s *FlightBillSettlementQueryRequest) SetPageNo(v int32) *FlightBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *FlightBillSettlementQueryRequest) SetPageSize(v int32) *FlightBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *FlightBillSettlementQueryRequest) SetPeriodEnd(v string) *FlightBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *FlightBillSettlementQueryRequest) SetPeriodStart(v string) *FlightBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

func (s *FlightBillSettlementQueryRequest) SetScrollId(v string) *FlightBillSettlementQueryRequest {
	s.ScrollId = &v
	return s
}

func (s *FlightBillSettlementQueryRequest) SetScrollMod(v bool) *FlightBillSettlementQueryRequest {
	s.ScrollMod = &v
	return s
}

func (s *FlightBillSettlementQueryRequest) Validate() error {
	return dara.Validate(s)
}
