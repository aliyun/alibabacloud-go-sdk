// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAIUsageLimitPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *AIUsageLimitPolicy
	GetDescription() *string
	SetGmtCreate(v string) *AIUsageLimitPolicy
	GetGmtCreate() *string
	SetGmtModified(v string) *AIUsageLimitPolicy
	GetGmtModified() *string
	SetLimitPolicyId(v string) *AIUsageLimitPolicy
	GetLimitPolicyId() *string
	SetLimitValue(v int64) *AIUsageLimitPolicy
	GetLimitValue() *int64
	SetMetricType(v string) *AIUsageLimitPolicy
	GetMetricType() *string
	SetName(v string) *AIUsageLimitPolicy
	GetName() *string
	SetPriority(v int32) *AIUsageLimitPolicy
	GetPriority() *int32
	SetResetPeriod(v string) *AIUsageLimitPolicy
	GetResetPeriod() *string
	SetServiceIds(v []*string) *AIUsageLimitPolicy
	GetServiceIds() []*string
	SetStatus(v string) *AIUsageLimitPolicy
	GetStatus() *string
	SetUserGroupIds(v []*string) *AIUsageLimitPolicy
	GetUserGroupIds() []*string
}

type AIUsageLimitPolicy struct {
	// A brief description of the policy\\"s purpose or scope.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp (in UTC) when the policy was created, formatted as `YYYY-MM-DDThh:mm:ssZ`. This is a system-generated, read-only property.
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp (in UTC) when the policy was last modified, formatted as `YYYY-MM-DDThh:mm:ssZ`. This is a system-generated, read-only property.
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The unique identifier for the usage limit policy. This is a system-generated, read-only value.
	LimitPolicyId *string `json:"LimitPolicyId,omitempty" xml:"LimitPolicyId,omitempty"`
	// The maximum value for the specified `MetricType` allowed within the `ResetPeriod`. Once this limit is reached, further requests are throttled or rejected.
	LimitValue *int64 `json:"LimitValue,omitempty" xml:"LimitValue,omitempty"`
	// The type of metric the limit applies to, such as the number of API requests, tokens processed, or compute units consumed.
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// A user-friendly name for the policy. This helps you identify the policy in a list.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority of the policy, used to determine the evaluation order when multiple policies apply to the same request. A lower number indicates a higher priority.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The time window during which the usage count is accumulated before it resets. For example: `Hour`, `Day`, or `Month`.
	ResetPeriod *string `json:"ResetPeriod,omitempty" xml:"ResetPeriod,omitempty"`
	// A list of service IDs that this policy applies to. The policy is enforced only for requests made to these services.
	ServiceIds []*string `json:"ServiceIds,omitempty" xml:"ServiceIds,omitempty" type:"Repeated"`
	// The status of the policy. Valid values are `Enabled` and `Disabled`. A disabled policy is not enforced.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// A list of user group IDs that this policy applies to. The policy is enforced only for users who belong to these groups.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
}

func (s AIUsageLimitPolicy) String() string {
	return dara.Prettify(s)
}

func (s AIUsageLimitPolicy) GoString() string {
	return s.String()
}

func (s *AIUsageLimitPolicy) GetDescription() *string {
	return s.Description
}

func (s *AIUsageLimitPolicy) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *AIUsageLimitPolicy) GetGmtModified() *string {
	return s.GmtModified
}

func (s *AIUsageLimitPolicy) GetLimitPolicyId() *string {
	return s.LimitPolicyId
}

func (s *AIUsageLimitPolicy) GetLimitValue() *int64 {
	return s.LimitValue
}

func (s *AIUsageLimitPolicy) GetMetricType() *string {
	return s.MetricType
}

func (s *AIUsageLimitPolicy) GetName() *string {
	return s.Name
}

func (s *AIUsageLimitPolicy) GetPriority() *int32 {
	return s.Priority
}

func (s *AIUsageLimitPolicy) GetResetPeriod() *string {
	return s.ResetPeriod
}

func (s *AIUsageLimitPolicy) GetServiceIds() []*string {
	return s.ServiceIds
}

func (s *AIUsageLimitPolicy) GetStatus() *string {
	return s.Status
}

func (s *AIUsageLimitPolicy) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *AIUsageLimitPolicy) SetDescription(v string) *AIUsageLimitPolicy {
	s.Description = &v
	return s
}

func (s *AIUsageLimitPolicy) SetGmtCreate(v string) *AIUsageLimitPolicy {
	s.GmtCreate = &v
	return s
}

func (s *AIUsageLimitPolicy) SetGmtModified(v string) *AIUsageLimitPolicy {
	s.GmtModified = &v
	return s
}

func (s *AIUsageLimitPolicy) SetLimitPolicyId(v string) *AIUsageLimitPolicy {
	s.LimitPolicyId = &v
	return s
}

func (s *AIUsageLimitPolicy) SetLimitValue(v int64) *AIUsageLimitPolicy {
	s.LimitValue = &v
	return s
}

func (s *AIUsageLimitPolicy) SetMetricType(v string) *AIUsageLimitPolicy {
	s.MetricType = &v
	return s
}

func (s *AIUsageLimitPolicy) SetName(v string) *AIUsageLimitPolicy {
	s.Name = &v
	return s
}

func (s *AIUsageLimitPolicy) SetPriority(v int32) *AIUsageLimitPolicy {
	s.Priority = &v
	return s
}

func (s *AIUsageLimitPolicy) SetResetPeriod(v string) *AIUsageLimitPolicy {
	s.ResetPeriod = &v
	return s
}

func (s *AIUsageLimitPolicy) SetServiceIds(v []*string) *AIUsageLimitPolicy {
	s.ServiceIds = v
	return s
}

func (s *AIUsageLimitPolicy) SetStatus(v string) *AIUsageLimitPolicy {
	s.Status = &v
	return s
}

func (s *AIUsageLimitPolicy) SetUserGroupIds(v []*string) *AIUsageLimitPolicy {
	s.UserGroupIds = v
	return s
}

func (s *AIUsageLimitPolicy) Validate() error {
	return dara.Validate(s)
}
