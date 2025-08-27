// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorFlightBillSettlementQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *CooperatorFlightBillSettlementQueryRequest
	GetBillBatch() *string
	SetCooperatorId(v string) *CooperatorFlightBillSettlementQueryRequest
	GetCooperatorId() *string
	SetOrderId(v int64) *CooperatorFlightBillSettlementQueryRequest
	GetOrderId() *int64
	SetPageNo(v int32) *CooperatorFlightBillSettlementQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *CooperatorFlightBillSettlementQueryRequest
	GetPageSize() *int32
	SetPeriodEnd(v string) *CooperatorFlightBillSettlementQueryRequest
	GetPeriodEnd() *string
	SetPeriodStart(v string) *CooperatorFlightBillSettlementQueryRequest
	GetPeriodStart() *string
	SetScrollId(v string) *CooperatorFlightBillSettlementQueryRequest
	GetScrollId() *string
	SetScrollMod(v bool) *CooperatorFlightBillSettlementQueryRequest
	GetScrollMod() *bool
}

type CooperatorFlightBillSettlementQueryRequest struct {
	BillBatch *string `json:"bill_batch,omitempty" xml:"bill_batch,omitempty"`
	// example:
	//
	// cooperator_alibtrip
	CooperatorId *string `json:"cooperator_id,omitempty" xml:"cooperator_id,omitempty"`
	OrderId      *int64  `json:"order_id,omitempty" xml:"order_id,omitempty"`
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

func (s CooperatorFlightBillSettlementQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s CooperatorFlightBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *CooperatorFlightBillSettlementQueryRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *CooperatorFlightBillSettlementQueryRequest) GetCooperatorId() *string {
	return s.CooperatorId
}

func (s *CooperatorFlightBillSettlementQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CooperatorFlightBillSettlementQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *CooperatorFlightBillSettlementQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CooperatorFlightBillSettlementQueryRequest) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *CooperatorFlightBillSettlementQueryRequest) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *CooperatorFlightBillSettlementQueryRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *CooperatorFlightBillSettlementQueryRequest) GetScrollMod() *bool {
	return s.ScrollMod
}

func (s *CooperatorFlightBillSettlementQueryRequest) SetBillBatch(v string) *CooperatorFlightBillSettlementQueryRequest {
	s.BillBatch = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryRequest) SetCooperatorId(v string) *CooperatorFlightBillSettlementQueryRequest {
	s.CooperatorId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryRequest) SetOrderId(v int64) *CooperatorFlightBillSettlementQueryRequest {
	s.OrderId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryRequest) SetPageNo(v int32) *CooperatorFlightBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryRequest) SetPageSize(v int32) *CooperatorFlightBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryRequest) SetPeriodEnd(v string) *CooperatorFlightBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryRequest) SetPeriodStart(v string) *CooperatorFlightBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryRequest) SetScrollId(v string) *CooperatorFlightBillSettlementQueryRequest {
	s.ScrollId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryRequest) SetScrollMod(v bool) *CooperatorFlightBillSettlementQueryRequest {
	s.ScrollMod = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryRequest) Validate() error {
	return dara.Validate(s)
}
