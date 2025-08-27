// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIeFlightBillSettlementQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *IeFlightBillSettlementQueryRequest
	GetBillBatch() *string
	SetOrderId(v int64) *IeFlightBillSettlementQueryRequest
	GetOrderId() *int64
	SetPageNo(v int32) *IeFlightBillSettlementQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *IeFlightBillSettlementQueryRequest
	GetPageSize() *int32
	SetPeriodEnd(v string) *IeFlightBillSettlementQueryRequest
	GetPeriodEnd() *string
	SetPeriodStart(v string) *IeFlightBillSettlementQueryRequest
	GetPeriodStart() *string
	SetScrollId(v string) *IeFlightBillSettlementQueryRequest
	GetScrollId() *string
	SetScrollMod(v bool) *IeFlightBillSettlementQueryRequest
	GetScrollMod() *bool
}

type IeFlightBillSettlementQueryRequest struct {
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
	// 100
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 2021-10-02
	PeriodEnd *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	// example:
	//
	// 2021-10-01
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
	ScrollId    *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	ScrollMod   *bool   `json:"scroll_mod,omitempty" xml:"scroll_mod,omitempty"`
}

func (s IeFlightBillSettlementQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s IeFlightBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *IeFlightBillSettlementQueryRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *IeFlightBillSettlementQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *IeFlightBillSettlementQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *IeFlightBillSettlementQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *IeFlightBillSettlementQueryRequest) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *IeFlightBillSettlementQueryRequest) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *IeFlightBillSettlementQueryRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *IeFlightBillSettlementQueryRequest) GetScrollMod() *bool {
	return s.ScrollMod
}

func (s *IeFlightBillSettlementQueryRequest) SetBillBatch(v string) *IeFlightBillSettlementQueryRequest {
	s.BillBatch = &v
	return s
}

func (s *IeFlightBillSettlementQueryRequest) SetOrderId(v int64) *IeFlightBillSettlementQueryRequest {
	s.OrderId = &v
	return s
}

func (s *IeFlightBillSettlementQueryRequest) SetPageNo(v int32) *IeFlightBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *IeFlightBillSettlementQueryRequest) SetPageSize(v int32) *IeFlightBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *IeFlightBillSettlementQueryRequest) SetPeriodEnd(v string) *IeFlightBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *IeFlightBillSettlementQueryRequest) SetPeriodStart(v string) *IeFlightBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

func (s *IeFlightBillSettlementQueryRequest) SetScrollId(v string) *IeFlightBillSettlementQueryRequest {
	s.ScrollId = &v
	return s
}

func (s *IeFlightBillSettlementQueryRequest) SetScrollMod(v bool) *IeFlightBillSettlementQueryRequest {
	s.ScrollMod = &v
	return s
}

func (s *IeFlightBillSettlementQueryRequest) Validate() error {
	return dara.Validate(s)
}
