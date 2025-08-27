// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIeHotelBillSettlementQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *IeHotelBillSettlementQueryRequest
	GetBillBatch() *string
	SetCategory(v int32) *IeHotelBillSettlementQueryRequest
	GetCategory() *int32
	SetOrderId(v int64) *IeHotelBillSettlementQueryRequest
	GetOrderId() *int64
	SetPageNo(v int32) *IeHotelBillSettlementQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *IeHotelBillSettlementQueryRequest
	GetPageSize() *int32
	SetPeriodEnd(v string) *IeHotelBillSettlementQueryRequest
	GetPeriodEnd() *string
	SetPeriodStart(v string) *IeHotelBillSettlementQueryRequest
	GetPeriodStart() *string
	SetScrollId(v string) *IeHotelBillSettlementQueryRequest
	GetScrollId() *string
	SetScrollMod(v bool) *IeHotelBillSettlementQueryRequest
	GetScrollMod() *bool
}

type IeHotelBillSettlementQueryRequest struct {
	BillBatch *string `json:"bill_batch,omitempty" xml:"bill_batch,omitempty"`
	// example:
	//
	// 12
	Category *int32 `json:"category,omitempty" xml:"category,omitempty"`
	OrderId  *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"page_no,omitempty" xml:"page_no,omitempty"`
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

func (s IeHotelBillSettlementQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s IeHotelBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *IeHotelBillSettlementQueryRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *IeHotelBillSettlementQueryRequest) GetCategory() *int32 {
	return s.Category
}

func (s *IeHotelBillSettlementQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *IeHotelBillSettlementQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *IeHotelBillSettlementQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *IeHotelBillSettlementQueryRequest) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *IeHotelBillSettlementQueryRequest) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *IeHotelBillSettlementQueryRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *IeHotelBillSettlementQueryRequest) GetScrollMod() *bool {
	return s.ScrollMod
}

func (s *IeHotelBillSettlementQueryRequest) SetBillBatch(v string) *IeHotelBillSettlementQueryRequest {
	s.BillBatch = &v
	return s
}

func (s *IeHotelBillSettlementQueryRequest) SetCategory(v int32) *IeHotelBillSettlementQueryRequest {
	s.Category = &v
	return s
}

func (s *IeHotelBillSettlementQueryRequest) SetOrderId(v int64) *IeHotelBillSettlementQueryRequest {
	s.OrderId = &v
	return s
}

func (s *IeHotelBillSettlementQueryRequest) SetPageNo(v int32) *IeHotelBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *IeHotelBillSettlementQueryRequest) SetPageSize(v int32) *IeHotelBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *IeHotelBillSettlementQueryRequest) SetPeriodEnd(v string) *IeHotelBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *IeHotelBillSettlementQueryRequest) SetPeriodStart(v string) *IeHotelBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

func (s *IeHotelBillSettlementQueryRequest) SetScrollId(v string) *IeHotelBillSettlementQueryRequest {
	s.ScrollId = &v
	return s
}

func (s *IeHotelBillSettlementQueryRequest) SetScrollMod(v bool) *IeHotelBillSettlementQueryRequest {
	s.ScrollMod = &v
	return s
}

func (s *IeHotelBillSettlementQueryRequest) Validate() error {
	return dara.Validate(s)
}
