// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealBillSettlementQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *MealBillSettlementQueryRequest
	GetBillBatch() *string
	SetOrderId(v int64) *MealBillSettlementQueryRequest
	GetOrderId() *int64
	SetPageNo(v int32) *MealBillSettlementQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *MealBillSettlementQueryRequest
	GetPageSize() *int32
	SetPeriodEnd(v string) *MealBillSettlementQueryRequest
	GetPeriodEnd() *string
	SetPeriodStart(v string) *MealBillSettlementQueryRequest
	GetPeriodStart() *string
	SetScrollId(v string) *MealBillSettlementQueryRequest
	GetScrollId() *string
	SetScrollMod(v bool) *MealBillSettlementQueryRequest
	GetScrollMod() *bool
}

type MealBillSettlementQueryRequest struct {
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
	// 2022-07-02
	PeriodEnd *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	// example:
	//
	// 2022-07-01
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
	ScrollId    *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	ScrollMod   *bool   `json:"scroll_mod,omitempty" xml:"scroll_mod,omitempty"`
}

func (s MealBillSettlementQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s MealBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *MealBillSettlementQueryRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *MealBillSettlementQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *MealBillSettlementQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *MealBillSettlementQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *MealBillSettlementQueryRequest) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *MealBillSettlementQueryRequest) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *MealBillSettlementQueryRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *MealBillSettlementQueryRequest) GetScrollMod() *bool {
	return s.ScrollMod
}

func (s *MealBillSettlementQueryRequest) SetBillBatch(v string) *MealBillSettlementQueryRequest {
	s.BillBatch = &v
	return s
}

func (s *MealBillSettlementQueryRequest) SetOrderId(v int64) *MealBillSettlementQueryRequest {
	s.OrderId = &v
	return s
}

func (s *MealBillSettlementQueryRequest) SetPageNo(v int32) *MealBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *MealBillSettlementQueryRequest) SetPageSize(v int32) *MealBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *MealBillSettlementQueryRequest) SetPeriodEnd(v string) *MealBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *MealBillSettlementQueryRequest) SetPeriodStart(v string) *MealBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

func (s *MealBillSettlementQueryRequest) SetScrollId(v string) *MealBillSettlementQueryRequest {
	s.ScrollId = &v
	return s
}

func (s *MealBillSettlementQueryRequest) SetScrollMod(v bool) *MealBillSettlementQueryRequest {
	s.ScrollMod = &v
	return s
}

func (s *MealBillSettlementQueryRequest) Validate() error {
	return dara.Validate(s)
}
