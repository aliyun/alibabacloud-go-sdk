// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainBillSettlementQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *TrainBillSettlementQueryRequest
	GetBillBatch() *string
	SetOrderId(v int64) *TrainBillSettlementQueryRequest
	GetOrderId() *int64
	SetPageNo(v int32) *TrainBillSettlementQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *TrainBillSettlementQueryRequest
	GetPageSize() *int32
	SetPeriodEnd(v string) *TrainBillSettlementQueryRequest
	GetPeriodEnd() *string
	SetPeriodStart(v string) *TrainBillSettlementQueryRequest
	GetPeriodStart() *string
	SetScrollId(v string) *TrainBillSettlementQueryRequest
	GetScrollId() *string
	SetScrollMod(v bool) *TrainBillSettlementQueryRequest
	GetScrollMod() *bool
}

type TrainBillSettlementQueryRequest struct {
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
	// 20
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

func (s TrainBillSettlementQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *TrainBillSettlementQueryRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *TrainBillSettlementQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *TrainBillSettlementQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *TrainBillSettlementQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *TrainBillSettlementQueryRequest) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *TrainBillSettlementQueryRequest) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *TrainBillSettlementQueryRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *TrainBillSettlementQueryRequest) GetScrollMod() *bool {
	return s.ScrollMod
}

func (s *TrainBillSettlementQueryRequest) SetBillBatch(v string) *TrainBillSettlementQueryRequest {
	s.BillBatch = &v
	return s
}

func (s *TrainBillSettlementQueryRequest) SetOrderId(v int64) *TrainBillSettlementQueryRequest {
	s.OrderId = &v
	return s
}

func (s *TrainBillSettlementQueryRequest) SetPageNo(v int32) *TrainBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *TrainBillSettlementQueryRequest) SetPageSize(v int32) *TrainBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *TrainBillSettlementQueryRequest) SetPeriodEnd(v string) *TrainBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *TrainBillSettlementQueryRequest) SetPeriodStart(v string) *TrainBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

func (s *TrainBillSettlementQueryRequest) SetScrollId(v string) *TrainBillSettlementQueryRequest {
	s.ScrollId = &v
	return s
}

func (s *TrainBillSettlementQueryRequest) SetScrollMod(v bool) *TrainBillSettlementQueryRequest {
	s.ScrollMod = &v
	return s
}

func (s *TrainBillSettlementQueryRequest) Validate() error {
	return dara.Validate(s)
}
