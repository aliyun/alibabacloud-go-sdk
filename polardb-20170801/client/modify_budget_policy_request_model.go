// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBudgetPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertThresholdPct(v string) *ModifyBudgetPolicyRequest
	GetAlertThresholdPct() *string
	SetBudgetPoints(v string) *ModifyBudgetPolicyRequest
	GetBudgetPoints() *string
	SetBudgetPolicyId(v string) *ModifyBudgetPolicyRequest
	GetBudgetPolicyId() *string
	SetGwClusterId(v string) *ModifyBudgetPolicyRequest
	GetGwClusterId() *string
	SetRegionId(v string) *ModifyBudgetPolicyRequest
	GetRegionId() *string
	SetResetDayOfMonth(v string) *ModifyBudgetPolicyRequest
	GetResetDayOfMonth() *string
}

type ModifyBudgetPolicyRequest struct {
	// example:
	//
	// 80
	AlertThresholdPct *string `json:"AlertThresholdPct,omitempty" xml:"AlertThresholdPct,omitempty"`
	// example:
	//
	// 10000
	BudgetPoints *string `json:"BudgetPoints,omitempty" xml:"BudgetPoints,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 023aacc1effc4b56bb154bfbec6ba9**
	BudgetPolicyId *string `json:"BudgetPolicyId,omitempty" xml:"BudgetPolicyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2
	ResetDayOfMonth *string `json:"ResetDayOfMonth,omitempty" xml:"ResetDayOfMonth,omitempty"`
}

func (s ModifyBudgetPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBudgetPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBudgetPolicyRequest) GetAlertThresholdPct() *string {
	return s.AlertThresholdPct
}

func (s *ModifyBudgetPolicyRequest) GetBudgetPoints() *string {
	return s.BudgetPoints
}

func (s *ModifyBudgetPolicyRequest) GetBudgetPolicyId() *string {
	return s.BudgetPolicyId
}

func (s *ModifyBudgetPolicyRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *ModifyBudgetPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyBudgetPolicyRequest) GetResetDayOfMonth() *string {
	return s.ResetDayOfMonth
}

func (s *ModifyBudgetPolicyRequest) SetAlertThresholdPct(v string) *ModifyBudgetPolicyRequest {
	s.AlertThresholdPct = &v
	return s
}

func (s *ModifyBudgetPolicyRequest) SetBudgetPoints(v string) *ModifyBudgetPolicyRequest {
	s.BudgetPoints = &v
	return s
}

func (s *ModifyBudgetPolicyRequest) SetBudgetPolicyId(v string) *ModifyBudgetPolicyRequest {
	s.BudgetPolicyId = &v
	return s
}

func (s *ModifyBudgetPolicyRequest) SetGwClusterId(v string) *ModifyBudgetPolicyRequest {
	s.GwClusterId = &v
	return s
}

func (s *ModifyBudgetPolicyRequest) SetRegionId(v string) *ModifyBudgetPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBudgetPolicyRequest) SetResetDayOfMonth(v string) *ModifyBudgetPolicyRequest {
	s.ResetDayOfMonth = &v
	return s
}

func (s *ModifyBudgetPolicyRequest) Validate() error {
	return dara.Validate(s)
}
