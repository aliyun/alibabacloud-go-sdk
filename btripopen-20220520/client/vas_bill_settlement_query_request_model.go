// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVasBillSettlementQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *VasBillSettlementQueryRequest
	GetBillBatch() *string
	SetCooperatorId(v string) *VasBillSettlementQueryRequest
	GetCooperatorId() *string
	SetOrderId(v int64) *VasBillSettlementQueryRequest
	GetOrderId() *int64
	SetPageNo(v int32) *VasBillSettlementQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *VasBillSettlementQueryRequest
	GetPageSize() *int32
	SetPeriodEnd(v string) *VasBillSettlementQueryRequest
	GetPeriodEnd() *string
	SetPeriodStart(v string) *VasBillSettlementQueryRequest
	GetPeriodStart() *string
	SetScrollId(v string) *VasBillSettlementQueryRequest
	GetScrollId() *string
	SetScrollMod(v bool) *VasBillSettlementQueryRequest
	GetScrollMod() *bool
}

type VasBillSettlementQueryRequest struct {
	// example:
	//
	// 20250501
	BillBatch *string `json:"bill_batch,omitempty" xml:"bill_batch,omitempty"`
	// example:
	//
	// cooperator_alibtrip
	CooperatorId *string `json:"cooperator_id,omitempty" xml:"cooperator_id,omitempty"`
	// example:
	//
	// 12345
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
	// 30
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 2022-07-01
	PeriodEnd *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	// example:
	//
	// 2021-10-01
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
	// example:
	//
	// CAESBgoEIgIIABgAIhkKFwMSAAAAMUw4MDAwMDAwMDA3MjdkMzgw
	ScrollId *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	// example:
	//
	// true
	ScrollMod *bool `json:"scroll_mod,omitempty" xml:"scroll_mod,omitempty"`
}

func (s VasBillSettlementQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s VasBillSettlementQueryRequest) GoString() string {
	return s.String()
}

func (s *VasBillSettlementQueryRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *VasBillSettlementQueryRequest) GetCooperatorId() *string {
	return s.CooperatorId
}

func (s *VasBillSettlementQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *VasBillSettlementQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *VasBillSettlementQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *VasBillSettlementQueryRequest) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *VasBillSettlementQueryRequest) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *VasBillSettlementQueryRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *VasBillSettlementQueryRequest) GetScrollMod() *bool {
	return s.ScrollMod
}

func (s *VasBillSettlementQueryRequest) SetBillBatch(v string) *VasBillSettlementQueryRequest {
	s.BillBatch = &v
	return s
}

func (s *VasBillSettlementQueryRequest) SetCooperatorId(v string) *VasBillSettlementQueryRequest {
	s.CooperatorId = &v
	return s
}

func (s *VasBillSettlementQueryRequest) SetOrderId(v int64) *VasBillSettlementQueryRequest {
	s.OrderId = &v
	return s
}

func (s *VasBillSettlementQueryRequest) SetPageNo(v int32) *VasBillSettlementQueryRequest {
	s.PageNo = &v
	return s
}

func (s *VasBillSettlementQueryRequest) SetPageSize(v int32) *VasBillSettlementQueryRequest {
	s.PageSize = &v
	return s
}

func (s *VasBillSettlementQueryRequest) SetPeriodEnd(v string) *VasBillSettlementQueryRequest {
	s.PeriodEnd = &v
	return s
}

func (s *VasBillSettlementQueryRequest) SetPeriodStart(v string) *VasBillSettlementQueryRequest {
	s.PeriodStart = &v
	return s
}

func (s *VasBillSettlementQueryRequest) SetScrollId(v string) *VasBillSettlementQueryRequest {
	s.ScrollId = &v
	return s
}

func (s *VasBillSettlementQueryRequest) SetScrollMod(v bool) *VasBillSettlementQueryRequest {
	s.ScrollMod = &v
	return s
}

func (s *VasBillSettlementQueryRequest) Validate() error {
	return dara.Validate(s)
}
