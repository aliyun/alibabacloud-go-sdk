// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBudgetUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertTriggered(v string) *DescribeBudgetUsageResponseBody
	GetAlertTriggered() *string
	SetDimensionRefId(v string) *DescribeBudgetUsageResponseBody
	GetDimensionRefId() *string
	SetDimensionType(v string) *DescribeBudgetUsageResponseBody
	GetDimensionType() *string
	SetExceeded(v string) *DescribeBudgetUsageResponseBody
	GetExceeded() *string
	SetGmtCreated(v string) *DescribeBudgetUsageResponseBody
	GetGmtCreated() *string
	SetGmtModified(v string) *DescribeBudgetUsageResponseBody
	GetGmtModified() *string
	SetGwClusterId(v string) *DescribeBudgetUsageResponseBody
	GetGwClusterId() *string
	SetPolicies(v []*DescribeBudgetUsageResponseBodyPolicies) *DescribeBudgetUsageResponseBody
	GetPolicies() []*DescribeBudgetUsageResponseBodyPolicies
	SetRequestId(v string) *DescribeBudgetUsageResponseBody
	GetRequestId() *string
	SetUsedPoints(v string) *DescribeBudgetUsageResponseBody
	GetUsedPoints() *string
	SetWindowEndTime(v string) *DescribeBudgetUsageResponseBody
	GetWindowEndTime() *string
	SetWindowStartTime(v string) *DescribeBudgetUsageResponseBody
	GetWindowStartTime() *string
}

type DescribeBudgetUsageResponseBody struct {
	// Indicates whether a budget alert is triggered.
	//
	// example:
	//
	// false
	AlertTriggered *string `json:"AlertTriggered,omitempty" xml:"AlertTriggered,omitempty"`
	// The dimension reference ID of the current budget.
	//
	// example:
	//
	// ref-001
	DimensionRefId *string `json:"DimensionRefId,omitempty" xml:"DimensionRefId,omitempty"`
	// The dimension type of the current budget.
	//
	// example:
	//
	// cluster
	DimensionType *string `json:"DimensionType,omitempty" xml:"DimensionType,omitempty"`
	// Indicates whether the budget is exceeded.
	//
	// example:
	//
	// false
	Exceeded *string `json:"Exceeded,omitempty" xml:"Exceeded,omitempty"`
	// The creation time of the budget usage record in RFC 3339 UTC format.
	//
	// example:
	//
	// 2025-01-01T00:00:00Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The last modification time of the budget usage record in RFC 3339 UTC format.
	//
	// example:
	//
	// 2025-01-15T12:00:00Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the gateway instance.
	//
	// example:
	//
	// pc-123
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The list of associated budget policies.
	Policies []*DescribeBudgetUsageResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of budget points used in the current statistical window.
	//
	// example:
	//
	// 5000
	UsedPoints *string `json:"UsedPoints,omitempty" xml:"UsedPoints,omitempty"`
	// The end time of the budget statistical window in RFC 3339 UTC format.
	//
	// example:
	//
	// 2025-01-31T23:59:59Z
	WindowEndTime *string `json:"WindowEndTime,omitempty" xml:"WindowEndTime,omitempty"`
	// The start time of the budget statistical window in RFC 3339 UTC format.
	//
	// example:
	//
	// 2025-01-01T00:00:00Z
	WindowStartTime *string `json:"WindowStartTime,omitempty" xml:"WindowStartTime,omitempty"`
}

func (s DescribeBudgetUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBudgetUsageResponseBody) GetAlertTriggered() *string {
	return s.AlertTriggered
}

func (s *DescribeBudgetUsageResponseBody) GetDimensionRefId() *string {
	return s.DimensionRefId
}

func (s *DescribeBudgetUsageResponseBody) GetDimensionType() *string {
	return s.DimensionType
}

func (s *DescribeBudgetUsageResponseBody) GetExceeded() *string {
	return s.Exceeded
}

func (s *DescribeBudgetUsageResponseBody) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeBudgetUsageResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeBudgetUsageResponseBody) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeBudgetUsageResponseBody) GetPolicies() []*DescribeBudgetUsageResponseBodyPolicies {
	return s.Policies
}

func (s *DescribeBudgetUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBudgetUsageResponseBody) GetUsedPoints() *string {
	return s.UsedPoints
}

func (s *DescribeBudgetUsageResponseBody) GetWindowEndTime() *string {
	return s.WindowEndTime
}

func (s *DescribeBudgetUsageResponseBody) GetWindowStartTime() *string {
	return s.WindowStartTime
}

func (s *DescribeBudgetUsageResponseBody) SetAlertTriggered(v string) *DescribeBudgetUsageResponseBody {
	s.AlertTriggered = &v
	return s
}

func (s *DescribeBudgetUsageResponseBody) SetDimensionRefId(v string) *DescribeBudgetUsageResponseBody {
	s.DimensionRefId = &v
	return s
}

func (s *DescribeBudgetUsageResponseBody) SetDimensionType(v string) *DescribeBudgetUsageResponseBody {
	s.DimensionType = &v
	return s
}

func (s *DescribeBudgetUsageResponseBody) SetExceeded(v string) *DescribeBudgetUsageResponseBody {
	s.Exceeded = &v
	return s
}

func (s *DescribeBudgetUsageResponseBody) SetGmtCreated(v string) *DescribeBudgetUsageResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *DescribeBudgetUsageResponseBody) SetGmtModified(v string) *DescribeBudgetUsageResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeBudgetUsageResponseBody) SetGwClusterId(v string) *DescribeBudgetUsageResponseBody {
	s.GwClusterId = &v
	return s
}

func (s *DescribeBudgetUsageResponseBody) SetPolicies(v []*DescribeBudgetUsageResponseBodyPolicies) *DescribeBudgetUsageResponseBody {
	s.Policies = v
	return s
}

func (s *DescribeBudgetUsageResponseBody) SetRequestId(v string) *DescribeBudgetUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBudgetUsageResponseBody) SetUsedPoints(v string) *DescribeBudgetUsageResponseBody {
	s.UsedPoints = &v
	return s
}

func (s *DescribeBudgetUsageResponseBody) SetWindowEndTime(v string) *DescribeBudgetUsageResponseBody {
	s.WindowEndTime = &v
	return s
}

func (s *DescribeBudgetUsageResponseBody) SetWindowStartTime(v string) *DescribeBudgetUsageResponseBody {
	s.WindowStartTime = &v
	return s
}

func (s *DescribeBudgetUsageResponseBody) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBudgetUsageResponseBodyPolicies struct {
	// The usage percentage that triggers a budget alert.
	//
	// example:
	//
	// 0.8
	AlertThresholdPct *string `json:"AlertThresholdPct,omitempty" xml:"AlertThresholdPct,omitempty"`
	// The budget dimension reference ID to which the policy applies.
	//
	// example:
	//
	// ref-001
	BudgetDimensionRefId *string `json:"BudgetDimensionRefId,omitempty" xml:"BudgetDimensionRefId,omitempty"`
	// The budget dimension type to which the policy applies.
	//
	// example:
	//
	// cluster
	BudgetDimensionType *string `json:"BudgetDimensionType,omitempty" xml:"BudgetDimensionType,omitempty"`
	// The number of budget points set by the policy.
	//
	// example:
	//
	// 10000
	BudgetPoints *string `json:"BudgetPoints,omitempty" xml:"BudgetPoints,omitempty"`
	// The ID of the budget policy.
	//
	// example:
	//
	// bp-001
	BudgetPolicyId *string `json:"BudgetPolicyId,omitempty" xml:"BudgetPolicyId,omitempty"`
	// The budget type.
	//
	// example:
	//
	// token
	BudgetType *string `json:"BudgetType,omitempty" xml:"BudgetType,omitempty"`
	// The day of the month when the budget is reset. For example, 1 indicates the first day of each month.
	//
	// example:
	//
	// 1
	ResetDayOfMonth *string `json:"ResetDayOfMonth,omitempty" xml:"ResetDayOfMonth,omitempty"`
	// The status of the budget policy.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBudgetUsageResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetUsageResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *DescribeBudgetUsageResponseBodyPolicies) GetAlertThresholdPct() *string {
	return s.AlertThresholdPct
}

func (s *DescribeBudgetUsageResponseBodyPolicies) GetBudgetDimensionRefId() *string {
	return s.BudgetDimensionRefId
}

func (s *DescribeBudgetUsageResponseBodyPolicies) GetBudgetDimensionType() *string {
	return s.BudgetDimensionType
}

func (s *DescribeBudgetUsageResponseBodyPolicies) GetBudgetPoints() *string {
	return s.BudgetPoints
}

func (s *DescribeBudgetUsageResponseBodyPolicies) GetBudgetPolicyId() *string {
	return s.BudgetPolicyId
}

func (s *DescribeBudgetUsageResponseBodyPolicies) GetBudgetType() *string {
	return s.BudgetType
}

func (s *DescribeBudgetUsageResponseBodyPolicies) GetResetDayOfMonth() *string {
	return s.ResetDayOfMonth
}

func (s *DescribeBudgetUsageResponseBodyPolicies) GetStatus() *string {
	return s.Status
}

func (s *DescribeBudgetUsageResponseBodyPolicies) SetAlertThresholdPct(v string) *DescribeBudgetUsageResponseBodyPolicies {
	s.AlertThresholdPct = &v
	return s
}

func (s *DescribeBudgetUsageResponseBodyPolicies) SetBudgetDimensionRefId(v string) *DescribeBudgetUsageResponseBodyPolicies {
	s.BudgetDimensionRefId = &v
	return s
}

func (s *DescribeBudgetUsageResponseBodyPolicies) SetBudgetDimensionType(v string) *DescribeBudgetUsageResponseBodyPolicies {
	s.BudgetDimensionType = &v
	return s
}

func (s *DescribeBudgetUsageResponseBodyPolicies) SetBudgetPoints(v string) *DescribeBudgetUsageResponseBodyPolicies {
	s.BudgetPoints = &v
	return s
}

func (s *DescribeBudgetUsageResponseBodyPolicies) SetBudgetPolicyId(v string) *DescribeBudgetUsageResponseBodyPolicies {
	s.BudgetPolicyId = &v
	return s
}

func (s *DescribeBudgetUsageResponseBodyPolicies) SetBudgetType(v string) *DescribeBudgetUsageResponseBodyPolicies {
	s.BudgetType = &v
	return s
}

func (s *DescribeBudgetUsageResponseBodyPolicies) SetResetDayOfMonth(v string) *DescribeBudgetUsageResponseBodyPolicies {
	s.ResetDayOfMonth = &v
	return s
}

func (s *DescribeBudgetUsageResponseBodyPolicies) SetStatus(v string) *DescribeBudgetUsageResponseBodyPolicies {
	s.Status = &v
	return s
}

func (s *DescribeBudgetUsageResponseBodyPolicies) Validate() error {
	return dara.Validate(s)
}
