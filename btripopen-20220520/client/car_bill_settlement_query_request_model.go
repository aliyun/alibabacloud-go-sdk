// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarBillSettlementQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *CarBillSettlementQueryRequest
	GetBillBatch() *string
	SetOrderId(v int64) *CarBillSettlementQueryRequest
	GetOrderId() *int64
	SetPageNo(v int32) *CarBillSettlementQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *CarBillSettlementQueryRequest
	GetPageSize() *int32
	SetPeriodEnd(v string) *CarBillSettlementQueryRequest
	GetPeriodEnd() *string
	SetPeriodStart(v string) *CarBillSettlementQueryRequest
	GetPeriodStart() *string
	SetScrollId(v string) *CarBillSettlementQueryRequest
	GetScrollId() *string
	SetScrollMod(v bool) *CarBillSettlementQueryRequest
	GetScrollMod() *bool
}

type CarBillSettlementQueryRequest struct {
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
	// 30
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 2022-07-01
	PeriodEnd *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	// example:
	//
	// 2022-07-01
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
	ScrollId    *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	ScrollMod   *bool   `json:"scroll_mod,omitempty" xml:"scroll_mod,omitempty"`
}

func (s CarBillSettlementQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s CarBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *CarBillSettlementQueryRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *CarBillSettlementQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CarBillSettlementQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *CarBillSettlementQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CarBillSettlementQueryRequest) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *CarBillSettlementQueryRequest) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *CarBillSettlementQueryRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *CarBillSettlementQueryRequest) GetScrollMod() *bool {
	return s.ScrollMod
}

func (s *CarBillSettlementQueryRequest) SetBillBatch(v string) *CarBillSettlementQueryRequest {
	s.BillBatch = &v
	return s
}

func (s *CarBillSettlementQueryRequest) SetOrderId(v int64) *CarBillSettlementQueryRequest {
	s.OrderId = &v
	return s
}

func (s *CarBillSettlementQueryRequest) SetPageNo(v int32) *CarBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *CarBillSettlementQueryRequest) SetPageSize(v int32) *CarBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *CarBillSettlementQueryRequest) SetPeriodEnd(v string) *CarBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *CarBillSettlementQueryRequest) SetPeriodStart(v string) *CarBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

func (s *CarBillSettlementQueryRequest) SetScrollId(v string) *CarBillSettlementQueryRequest {
	s.ScrollId = &v
	return s
}

func (s *CarBillSettlementQueryRequest) SetScrollMod(v bool) *CarBillSettlementQueryRequest {
	s.ScrollMod = &v
	return s
}

func (s *CarBillSettlementQueryRequest) Validate() error {
	return dara.Validate(s)
}
