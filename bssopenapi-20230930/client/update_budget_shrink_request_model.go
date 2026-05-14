// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBudgetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetName(v string) *UpdateBudgetShrinkRequest
	GetBudgetName() *string
	SetBudgetType(v string) *UpdateBudgetShrinkRequest
	GetBudgetType() *string
	SetComment(v string) *UpdateBudgetShrinkRequest
	GetComment() *string
	SetCycleEndPeriod(v string) *UpdateBudgetShrinkRequest
	GetCycleEndPeriod() *string
	SetCycleQuotaShrink(v string) *UpdateBudgetShrinkRequest
	GetCycleQuotaShrink() *string
	SetCycleStartPeriod(v string) *UpdateBudgetShrinkRequest
	GetCycleStartPeriod() *string
	SetCycleType(v string) *UpdateBudgetShrinkRequest
	GetCycleType() *string
	SetEcIdAccountIdsShrink(v string) *UpdateBudgetShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetMetric(v string) *UpdateBudgetShrinkRequest
	GetMetric() *string
	SetNbid(v string) *UpdateBudgetShrinkRequest
	GetNbid() *string
	SetOriginalBudgetName(v string) *UpdateBudgetShrinkRequest
	GetOriginalBudgetName() *string
	SetQueryFilterShrink(v string) *UpdateBudgetShrinkRequest
	GetQueryFilterShrink() *string
	SetQuota(v string) *UpdateBudgetShrinkRequest
	GetQuota() *string
	SetQuotaType(v string) *UpdateBudgetShrinkRequest
	GetQuotaType() *string
	SetWarnConfsShrink(v string) *UpdateBudgetShrinkRequest
	GetWarnConfsShrink() *string
}

type UpdateBudgetShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// NewBudgetName
	BudgetName *string `json:"BudgetName,omitempty" xml:"BudgetName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CONSUME
	BudgetType *string `json:"BudgetType,omitempty" xml:"BudgetType,omitempty"`
	Comment    *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-12
	CycleEndPeriod   *string `json:"CycleEndPeriod,omitempty" xml:"CycleEndPeriod,omitempty"`
	CycleQuotaShrink *string `json:"CycleQuota,omitempty" xml:"CycleQuota,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-01
	CycleStartPeriod *string `json:"CycleStartPeriod,omitempty" xml:"CycleStartPeriod,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MONTHLY
	CycleType            *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// REQUIRE_AMOUNT
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Nbid   *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OriginalBudgetName
	OriginalBudgetName *string `json:"OriginalBudgetName,omitempty" xml:"OriginalBudgetName,omitempty"`
	QueryFilterShrink  *string `json:"QueryFilter,omitempty" xml:"QueryFilter,omitempty"`
	// example:
	//
	// 1000
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FIXED
	QuotaType       *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	WarnConfsShrink *string `json:"WarnConfs,omitempty" xml:"WarnConfs,omitempty"`
}

func (s UpdateBudgetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBudgetShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateBudgetShrinkRequest) GetBudgetName() *string {
	return s.BudgetName
}

func (s *UpdateBudgetShrinkRequest) GetBudgetType() *string {
	return s.BudgetType
}

func (s *UpdateBudgetShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateBudgetShrinkRequest) GetCycleEndPeriod() *string {
	return s.CycleEndPeriod
}

func (s *UpdateBudgetShrinkRequest) GetCycleQuotaShrink() *string {
	return s.CycleQuotaShrink
}

func (s *UpdateBudgetShrinkRequest) GetCycleStartPeriod() *string {
	return s.CycleStartPeriod
}

func (s *UpdateBudgetShrinkRequest) GetCycleType() *string {
	return s.CycleType
}

func (s *UpdateBudgetShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *UpdateBudgetShrinkRequest) GetMetric() *string {
	return s.Metric
}

func (s *UpdateBudgetShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *UpdateBudgetShrinkRequest) GetOriginalBudgetName() *string {
	return s.OriginalBudgetName
}

func (s *UpdateBudgetShrinkRequest) GetQueryFilterShrink() *string {
	return s.QueryFilterShrink
}

func (s *UpdateBudgetShrinkRequest) GetQuota() *string {
	return s.Quota
}

func (s *UpdateBudgetShrinkRequest) GetQuotaType() *string {
	return s.QuotaType
}

func (s *UpdateBudgetShrinkRequest) GetWarnConfsShrink() *string {
	return s.WarnConfsShrink
}

func (s *UpdateBudgetShrinkRequest) SetBudgetName(v string) *UpdateBudgetShrinkRequest {
	s.BudgetName = &v
	return s
}

func (s *UpdateBudgetShrinkRequest) SetBudgetType(v string) *UpdateBudgetShrinkRequest {
	s.BudgetType = &v
	return s
}

func (s *UpdateBudgetShrinkRequest) SetComment(v string) *UpdateBudgetShrinkRequest {
	s.Comment = &v
	return s
}

func (s *UpdateBudgetShrinkRequest) SetCycleEndPeriod(v string) *UpdateBudgetShrinkRequest {
	s.CycleEndPeriod = &v
	return s
}

func (s *UpdateBudgetShrinkRequest) SetCycleQuotaShrink(v string) *UpdateBudgetShrinkRequest {
	s.CycleQuotaShrink = &v
	return s
}

func (s *UpdateBudgetShrinkRequest) SetCycleStartPeriod(v string) *UpdateBudgetShrinkRequest {
	s.CycleStartPeriod = &v
	return s
}

func (s *UpdateBudgetShrinkRequest) SetCycleType(v string) *UpdateBudgetShrinkRequest {
	s.CycleType = &v
	return s
}

func (s *UpdateBudgetShrinkRequest) SetEcIdAccountIdsShrink(v string) *UpdateBudgetShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *UpdateBudgetShrinkRequest) SetMetric(v string) *UpdateBudgetShrinkRequest {
	s.Metric = &v
	return s
}

func (s *UpdateBudgetShrinkRequest) SetNbid(v string) *UpdateBudgetShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *UpdateBudgetShrinkRequest) SetOriginalBudgetName(v string) *UpdateBudgetShrinkRequest {
	s.OriginalBudgetName = &v
	return s
}

func (s *UpdateBudgetShrinkRequest) SetQueryFilterShrink(v string) *UpdateBudgetShrinkRequest {
	s.QueryFilterShrink = &v
	return s
}

func (s *UpdateBudgetShrinkRequest) SetQuota(v string) *UpdateBudgetShrinkRequest {
	s.Quota = &v
	return s
}

func (s *UpdateBudgetShrinkRequest) SetQuotaType(v string) *UpdateBudgetShrinkRequest {
	s.QuotaType = &v
	return s
}

func (s *UpdateBudgetShrinkRequest) SetWarnConfsShrink(v string) *UpdateBudgetShrinkRequest {
	s.WarnConfsShrink = &v
	return s
}

func (s *UpdateBudgetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
