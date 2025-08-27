// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelBillSettlementQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *HotelBillSettlementQueryRequest
	GetBillBatch() *string
	SetOrderId(v int64) *HotelBillSettlementQueryRequest
	GetOrderId() *int64
	SetPageNo(v int32) *HotelBillSettlementQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *HotelBillSettlementQueryRequest
	GetPageSize() *int32
	SetPeriodEnd(v string) *HotelBillSettlementQueryRequest
	GetPeriodEnd() *string
	SetPeriodStart(v string) *HotelBillSettlementQueryRequest
	GetPeriodStart() *string
	SetScrollId(v string) *HotelBillSettlementQueryRequest
	GetScrollId() *string
	SetScrollMod(v bool) *HotelBillSettlementQueryRequest
	GetScrollMod() *bool
}

type HotelBillSettlementQueryRequest struct {
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
	// 2021-10-02
	PeriodEnd *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	// example:
	//
	// 2021-10-01
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
	ScrollId    *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	ScrollMod   *bool   `json:"scroll_mod,omitempty" xml:"scroll_mod,omitempty"`
}

func (s HotelBillSettlementQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *HotelBillSettlementQueryRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *HotelBillSettlementQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *HotelBillSettlementQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *HotelBillSettlementQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *HotelBillSettlementQueryRequest) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *HotelBillSettlementQueryRequest) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *HotelBillSettlementQueryRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *HotelBillSettlementQueryRequest) GetScrollMod() *bool {
	return s.ScrollMod
}

func (s *HotelBillSettlementQueryRequest) SetBillBatch(v string) *HotelBillSettlementQueryRequest {
	s.BillBatch = &v
	return s
}

func (s *HotelBillSettlementQueryRequest) SetOrderId(v int64) *HotelBillSettlementQueryRequest {
	s.OrderId = &v
	return s
}

func (s *HotelBillSettlementQueryRequest) SetPageNo(v int32) *HotelBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *HotelBillSettlementQueryRequest) SetPageSize(v int32) *HotelBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *HotelBillSettlementQueryRequest) SetPeriodEnd(v string) *HotelBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *HotelBillSettlementQueryRequest) SetPeriodStart(v string) *HotelBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

func (s *HotelBillSettlementQueryRequest) SetScrollId(v string) *HotelBillSettlementQueryRequest {
	s.ScrollId = &v
	return s
}

func (s *HotelBillSettlementQueryRequest) SetScrollMod(v bool) *HotelBillSettlementQueryRequest {
	s.ScrollMod = &v
	return s
}

func (s *HotelBillSettlementQueryRequest) Validate() error {
	return dara.Validate(s)
}
