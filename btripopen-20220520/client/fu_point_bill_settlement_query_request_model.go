// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFuPointBillSettlementQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *FuPointBillSettlementQueryRequest
	GetBillBatch() *string
	SetCooperatorId(v string) *FuPointBillSettlementQueryRequest
	GetCooperatorId() *string
	SetOrderId(v int64) *FuPointBillSettlementQueryRequest
	GetOrderId() *int64
	SetPageNo(v int32) *FuPointBillSettlementQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *FuPointBillSettlementQueryRequest
	GetPageSize() *int32
	SetPeriodEnd(v string) *FuPointBillSettlementQueryRequest
	GetPeriodEnd() *string
	SetPeriodStart(v string) *FuPointBillSettlementQueryRequest
	GetPeriodStart() *string
	SetScrollId(v string) *FuPointBillSettlementQueryRequest
	GetScrollId() *string
	SetScrollMod(v bool) *FuPointBillSettlementQueryRequest
	GetScrollMod() *bool
}

type FuPointBillSettlementQueryRequest struct {
	// example:
	//
	// 20240101
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
	// example:
	//
	// 1qwe
	ScrollId *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	// example:
	//
	// true
	ScrollMod *bool `json:"scroll_mod,omitempty" xml:"scroll_mod,omitempty"`
}

func (s FuPointBillSettlementQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s FuPointBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *FuPointBillSettlementQueryRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *FuPointBillSettlementQueryRequest) GetCooperatorId() *string {
	return s.CooperatorId
}

func (s *FuPointBillSettlementQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FuPointBillSettlementQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *FuPointBillSettlementQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *FuPointBillSettlementQueryRequest) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *FuPointBillSettlementQueryRequest) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *FuPointBillSettlementQueryRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *FuPointBillSettlementQueryRequest) GetScrollMod() *bool {
	return s.ScrollMod
}

func (s *FuPointBillSettlementQueryRequest) SetBillBatch(v string) *FuPointBillSettlementQueryRequest {
	s.BillBatch = &v
	return s
}

func (s *FuPointBillSettlementQueryRequest) SetCooperatorId(v string) *FuPointBillSettlementQueryRequest {
	s.CooperatorId = &v
	return s
}

func (s *FuPointBillSettlementQueryRequest) SetOrderId(v int64) *FuPointBillSettlementQueryRequest {
	s.OrderId = &v
	return s
}

func (s *FuPointBillSettlementQueryRequest) SetPageNo(v int32) *FuPointBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *FuPointBillSettlementQueryRequest) SetPageSize(v int32) *FuPointBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *FuPointBillSettlementQueryRequest) SetPeriodEnd(v string) *FuPointBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *FuPointBillSettlementQueryRequest) SetPeriodStart(v string) *FuPointBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

func (s *FuPointBillSettlementQueryRequest) SetScrollId(v string) *FuPointBillSettlementQueryRequest {
	s.ScrollId = &v
	return s
}

func (s *FuPointBillSettlementQueryRequest) SetScrollMod(v bool) *FuPointBillSettlementQueryRequest {
	s.ScrollMod = &v
	return s
}

func (s *FuPointBillSettlementQueryRequest) Validate() error {
	return dara.Validate(s)
}
