// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorHotelBillSettlementQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *CooperatorHotelBillSettlementQueryRequest
	GetBillBatch() *string
	SetCooperatorId(v string) *CooperatorHotelBillSettlementQueryRequest
	GetCooperatorId() *string
	SetOrderId(v int64) *CooperatorHotelBillSettlementQueryRequest
	GetOrderId() *int64
	SetPageNo(v int32) *CooperatorHotelBillSettlementQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *CooperatorHotelBillSettlementQueryRequest
	GetPageSize() *int32
	SetPeriodEnd(v string) *CooperatorHotelBillSettlementQueryRequest
	GetPeriodEnd() *string
	SetPeriodStart(v string) *CooperatorHotelBillSettlementQueryRequest
	GetPeriodStart() *string
	SetScrollId(v string) *CooperatorHotelBillSettlementQueryRequest
	GetScrollId() *string
	SetScrollMod(v bool) *CooperatorHotelBillSettlementQueryRequest
	GetScrollMod() *bool
}

type CooperatorHotelBillSettlementQueryRequest struct {
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

func (s CooperatorHotelBillSettlementQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s CooperatorHotelBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *CooperatorHotelBillSettlementQueryRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *CooperatorHotelBillSettlementQueryRequest) GetCooperatorId() *string {
	return s.CooperatorId
}

func (s *CooperatorHotelBillSettlementQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CooperatorHotelBillSettlementQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *CooperatorHotelBillSettlementQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CooperatorHotelBillSettlementQueryRequest) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *CooperatorHotelBillSettlementQueryRequest) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *CooperatorHotelBillSettlementQueryRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *CooperatorHotelBillSettlementQueryRequest) GetScrollMod() *bool {
	return s.ScrollMod
}

func (s *CooperatorHotelBillSettlementQueryRequest) SetBillBatch(v string) *CooperatorHotelBillSettlementQueryRequest {
	s.BillBatch = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryRequest) SetCooperatorId(v string) *CooperatorHotelBillSettlementQueryRequest {
	s.CooperatorId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryRequest) SetOrderId(v int64) *CooperatorHotelBillSettlementQueryRequest {
	s.OrderId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryRequest) SetPageNo(v int32) *CooperatorHotelBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryRequest) SetPageSize(v int32) *CooperatorHotelBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryRequest) SetPeriodEnd(v string) *CooperatorHotelBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryRequest) SetPeriodStart(v string) *CooperatorHotelBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryRequest) SetScrollId(v string) *CooperatorHotelBillSettlementQueryRequest {
	s.ScrollId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryRequest) SetScrollMod(v bool) *CooperatorHotelBillSettlementQueryRequest {
	s.ScrollMod = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryRequest) Validate() error {
	return dara.Validate(s)
}
