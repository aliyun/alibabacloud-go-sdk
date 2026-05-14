// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBudgetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetName(v string) *CreateBudgetShrinkRequest
	GetBudgetName() *string
	SetBudgetType(v string) *CreateBudgetShrinkRequest
	GetBudgetType() *string
	SetComment(v string) *CreateBudgetShrinkRequest
	GetComment() *string
	SetCycleEndPeriod(v string) *CreateBudgetShrinkRequest
	GetCycleEndPeriod() *string
	SetCycleQuotaShrink(v string) *CreateBudgetShrinkRequest
	GetCycleQuotaShrink() *string
	SetCycleStartPeriod(v string) *CreateBudgetShrinkRequest
	GetCycleStartPeriod() *string
	SetCycleType(v string) *CreateBudgetShrinkRequest
	GetCycleType() *string
	SetEcIdAccountIdsShrink(v string) *CreateBudgetShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetMetric(v string) *CreateBudgetShrinkRequest
	GetMetric() *string
	SetNbid(v string) *CreateBudgetShrinkRequest
	GetNbid() *string
	SetQueryFilterShrink(v string) *CreateBudgetShrinkRequest
	GetQueryFilterShrink() *string
	SetQuota(v string) *CreateBudgetShrinkRequest
	GetQuota() *string
	SetQuotaType(v string) *CreateBudgetShrinkRequest
	GetQuotaType() *string
	SetWarnConfsShrink(v string) *CreateBudgetShrinkRequest
	GetWarnConfsShrink() *string
}

type CreateBudgetShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Department_dev_budget
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
	Metric            *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Nbid              *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	QueryFilterShrink *string `json:"QueryFilter,omitempty" xml:"QueryFilter,omitempty"`
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

func (s CreateBudgetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBudgetShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBudgetShrinkRequest) GetBudgetName() *string {
	return s.BudgetName
}

func (s *CreateBudgetShrinkRequest) GetBudgetType() *string {
	return s.BudgetType
}

func (s *CreateBudgetShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateBudgetShrinkRequest) GetCycleEndPeriod() *string {
	return s.CycleEndPeriod
}

func (s *CreateBudgetShrinkRequest) GetCycleQuotaShrink() *string {
	return s.CycleQuotaShrink
}

func (s *CreateBudgetShrinkRequest) GetCycleStartPeriod() *string {
	return s.CycleStartPeriod
}

func (s *CreateBudgetShrinkRequest) GetCycleType() *string {
	return s.CycleType
}

func (s *CreateBudgetShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *CreateBudgetShrinkRequest) GetMetric() *string {
	return s.Metric
}

func (s *CreateBudgetShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *CreateBudgetShrinkRequest) GetQueryFilterShrink() *string {
	return s.QueryFilterShrink
}

func (s *CreateBudgetShrinkRequest) GetQuota() *string {
	return s.Quota
}

func (s *CreateBudgetShrinkRequest) GetQuotaType() *string {
	return s.QuotaType
}

func (s *CreateBudgetShrinkRequest) GetWarnConfsShrink() *string {
	return s.WarnConfsShrink
}

func (s *CreateBudgetShrinkRequest) SetBudgetName(v string) *CreateBudgetShrinkRequest {
	s.BudgetName = &v
	return s
}

func (s *CreateBudgetShrinkRequest) SetBudgetType(v string) *CreateBudgetShrinkRequest {
	s.BudgetType = &v
	return s
}

func (s *CreateBudgetShrinkRequest) SetComment(v string) *CreateBudgetShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateBudgetShrinkRequest) SetCycleEndPeriod(v string) *CreateBudgetShrinkRequest {
	s.CycleEndPeriod = &v
	return s
}

func (s *CreateBudgetShrinkRequest) SetCycleQuotaShrink(v string) *CreateBudgetShrinkRequest {
	s.CycleQuotaShrink = &v
	return s
}

func (s *CreateBudgetShrinkRequest) SetCycleStartPeriod(v string) *CreateBudgetShrinkRequest {
	s.CycleStartPeriod = &v
	return s
}

func (s *CreateBudgetShrinkRequest) SetCycleType(v string) *CreateBudgetShrinkRequest {
	s.CycleType = &v
	return s
}

func (s *CreateBudgetShrinkRequest) SetEcIdAccountIdsShrink(v string) *CreateBudgetShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *CreateBudgetShrinkRequest) SetMetric(v string) *CreateBudgetShrinkRequest {
	s.Metric = &v
	return s
}

func (s *CreateBudgetShrinkRequest) SetNbid(v string) *CreateBudgetShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *CreateBudgetShrinkRequest) SetQueryFilterShrink(v string) *CreateBudgetShrinkRequest {
	s.QueryFilterShrink = &v
	return s
}

func (s *CreateBudgetShrinkRequest) SetQuota(v string) *CreateBudgetShrinkRequest {
	s.Quota = &v
	return s
}

func (s *CreateBudgetShrinkRequest) SetQuotaType(v string) *CreateBudgetShrinkRequest {
	s.QuotaType = &v
	return s
}

func (s *CreateBudgetShrinkRequest) SetWarnConfsShrink(v string) *CreateBudgetShrinkRequest {
	s.WarnConfsShrink = &v
	return s
}

func (s *CreateBudgetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
