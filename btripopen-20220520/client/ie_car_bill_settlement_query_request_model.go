// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIeCarBillSettlementQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *IeCarBillSettlementQueryRequest
	GetBillBatch() *string
	SetBillRecordTimeEnd(v string) *IeCarBillSettlementQueryRequest
	GetBillRecordTimeEnd() *string
	SetBillRecordTimeStart(v string) *IeCarBillSettlementQueryRequest
	GetBillRecordTimeStart() *string
	SetOrderId(v int64) *IeCarBillSettlementQueryRequest
	GetOrderId() *int64
	SetPageNo(v int32) *IeCarBillSettlementQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *IeCarBillSettlementQueryRequest
	GetPageSize() *int32
	SetPeriodEnd(v string) *IeCarBillSettlementQueryRequest
	GetPeriodEnd() *string
	SetPeriodStart(v string) *IeCarBillSettlementQueryRequest
	GetPeriodStart() *string
	SetScrollId(v string) *IeCarBillSettlementQueryRequest
	GetScrollId() *string
	SetScrollMod(v bool) *IeCarBillSettlementQueryRequest
	GetScrollMod() *bool
}

type IeCarBillSettlementQueryRequest struct {
	// example:
	//
	// 20250725
	BillBatch           *string `json:"bill_batch,omitempty" xml:"bill_batch,omitempty"`
	BillRecordTimeEnd   *string `json:"bill_record_time_end,omitempty" xml:"bill_record_time_end,omitempty"`
	BillRecordTimeStart *string `json:"bill_record_time_start,omitempty" xml:"bill_record_time_start,omitempty"`
	// example:
	//
	// 1002002203361199686
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
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
	// 10
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
	// CAESBgoEIgIIABgAIhkKFwMSAAAAMUw4MDAwMDAwMDA3MzA1MGJj
	ScrollId *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	// example:
	//
	// true
	ScrollMod *bool `json:"scroll_mod,omitempty" xml:"scroll_mod,omitempty"`
}

func (s IeCarBillSettlementQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s IeCarBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *IeCarBillSettlementQueryRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *IeCarBillSettlementQueryRequest) GetBillRecordTimeEnd() *string {
	return s.BillRecordTimeEnd
}

func (s *IeCarBillSettlementQueryRequest) GetBillRecordTimeStart() *string {
	return s.BillRecordTimeStart
}

func (s *IeCarBillSettlementQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *IeCarBillSettlementQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *IeCarBillSettlementQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *IeCarBillSettlementQueryRequest) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *IeCarBillSettlementQueryRequest) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *IeCarBillSettlementQueryRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *IeCarBillSettlementQueryRequest) GetScrollMod() *bool {
	return s.ScrollMod
}

func (s *IeCarBillSettlementQueryRequest) SetBillBatch(v string) *IeCarBillSettlementQueryRequest {
	s.BillBatch = &v
	return s
}

func (s *IeCarBillSettlementQueryRequest) SetBillRecordTimeEnd(v string) *IeCarBillSettlementQueryRequest {
	s.BillRecordTimeEnd = &v
	return s
}

func (s *IeCarBillSettlementQueryRequest) SetBillRecordTimeStart(v string) *IeCarBillSettlementQueryRequest {
	s.BillRecordTimeStart = &v
	return s
}

func (s *IeCarBillSettlementQueryRequest) SetOrderId(v int64) *IeCarBillSettlementQueryRequest {
	s.OrderId = &v
	return s
}

func (s *IeCarBillSettlementQueryRequest) SetPageNo(v int32) *IeCarBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *IeCarBillSettlementQueryRequest) SetPageSize(v int32) *IeCarBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *IeCarBillSettlementQueryRequest) SetPeriodEnd(v string) *IeCarBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *IeCarBillSettlementQueryRequest) SetPeriodStart(v string) *IeCarBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

func (s *IeCarBillSettlementQueryRequest) SetScrollId(v string) *IeCarBillSettlementQueryRequest {
	s.ScrollId = &v
	return s
}

func (s *IeCarBillSettlementQueryRequest) SetScrollMod(v bool) *IeCarBillSettlementQueryRequest {
	s.ScrollMod = &v
	return s
}

func (s *IeCarBillSettlementQueryRequest) Validate() error {
	return dara.Validate(s)
}
