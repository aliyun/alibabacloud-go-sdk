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
	// The alert threshold, as a percentage. Valid values: 0 to 100.
	//
	// example:
	//
	// 10
	AlertThresholdPct *string `json:"AlertThresholdPct,omitempty" xml:"AlertThresholdPct,omitempty"`
	// The ID of the dimension object. This parameter is required if `BudgetType` is set to `ConsumerTotal` or `ConsumerGroupTotal`.
	//
	// example:
	//
	// mi-xxxxx
	BudgetDimensionRefId *string `json:"BudgetDimensionRefId,omitempty" xml:"BudgetDimensionRefId,omitempty"`
	// The number of budget points.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	BudgetPoints *string `json:"BudgetPoints,omitempty" xml:"BudgetPoints,omitempty"`
	// The budget type. Valid values:
	//
	// - `GlobalTotal`: global total budget
	//
	// - `ConsumerTotal`: consumer total budget
	//
	// - `ConsumerGroupTotal`: consumer group total budget
	//
	// This parameter is required.
	//
	// example:
	//
	// ConsumerTotal
	BudgetType *string `json:"BudgetType,omitempty" xml:"BudgetType,omitempty"`
	// The gateway cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The day of the month on which the budget resets. Valid values: 1 to 28.
	//
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
