// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBudgetPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertThresholdPct(v string) *CreateBudgetPolicyRequest
	GetAlertThresholdPct() *string
	SetBudgetDimensionRefId(v string) *CreateBudgetPolicyRequest
	GetBudgetDimensionRefId() *string
	SetBudgetPoints(v string) *CreateBudgetPolicyRequest
	GetBudgetPoints() *string
	SetBudgetType(v string) *CreateBudgetPolicyRequest
	GetBudgetType() *string
	SetGwClusterId(v string) *CreateBudgetPolicyRequest
	GetGwClusterId() *string
	SetRegionId(v string) *CreateBudgetPolicyRequest
	GetRegionId() *string
	SetResetDayOfMonth(v string) *CreateBudgetPolicyRequest
	GetResetDayOfMonth() *string
}

type CreateBudgetPolicyRequest struct {
	// example:
	//
	// 10
	AlertThresholdPct *string `json:"AlertThresholdPct,omitempty" xml:"AlertThresholdPct,omitempty"`
	// example:
	//
	// mi-xxxxx
	BudgetDimensionRefId *string `json:"BudgetDimensionRefId,omitempty" xml:"BudgetDimensionRefId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	BudgetPoints *string `json:"BudgetPoints,omitempty" xml:"BudgetPoints,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ConsumerTotal
	BudgetType *string `json:"BudgetType,omitempty" xml:"BudgetType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ResetDayOfMonth *string `json:"ResetDayOfMonth,omitempty" xml:"ResetDayOfMonth,omitempty"`
}

func (s CreateBudgetPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBudgetPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateBudgetPolicyRequest) GetAlertThresholdPct() *string {
	return s.AlertThresholdPct
}

func (s *CreateBudgetPolicyRequest) GetBudgetDimensionRefId() *string {
	return s.BudgetDimensionRefId
}

func (s *CreateBudgetPolicyRequest) GetBudgetPoints() *string {
	return s.BudgetPoints
}

func (s *CreateBudgetPolicyRequest) GetBudgetType() *string {
	return s.BudgetType
}

func (s *CreateBudgetPolicyRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *CreateBudgetPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBudgetPolicyRequest) GetResetDayOfMonth() *string {
	return s.ResetDayOfMonth
}

func (s *CreateBudgetPolicyRequest) SetAlertThresholdPct(v string) *CreateBudgetPolicyRequest {
	s.AlertThresholdPct = &v
	return s
}

func (s *CreateBudgetPolicyRequest) SetBudgetDimensionRefId(v string) *CreateBudgetPolicyRequest {
	s.BudgetDimensionRefId = &v
	return s
}

func (s *CreateBudgetPolicyRequest) SetBudgetPoints(v string) *CreateBudgetPolicyRequest {
	s.BudgetPoints = &v
	return s
}

func (s *CreateBudgetPolicyRequest) SetBudgetType(v string) *CreateBudgetPolicyRequest {
	s.BudgetType = &v
	return s
}

func (s *CreateBudgetPolicyRequest) SetGwClusterId(v string) *CreateBudgetPolicyRequest {
	s.GwClusterId = &v
	return s
}

func (s *CreateBudgetPolicyRequest) SetRegionId(v string) *CreateBudgetPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBudgetPolicyRequest) SetResetDayOfMonth(v string) *CreateBudgetPolicyRequest {
	s.ResetDayOfMonth = &v
	return s
}

func (s *CreateBudgetPolicyRequest) Validate() error {
	return dara.Validate(s)
}
