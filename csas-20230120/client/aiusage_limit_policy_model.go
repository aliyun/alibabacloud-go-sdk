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
	Description   *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate     *string   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *string   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	LimitPolicyId *string   `json:"LimitPolicyId,omitempty" xml:"LimitPolicyId,omitempty"`
	LimitValue    *int64    `json:"LimitValue,omitempty" xml:"LimitValue,omitempty"`
	MetricType    *string   `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Name          *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Priority      *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResetPeriod   *string   `json:"ResetPeriod,omitempty" xml:"ResetPeriod,omitempty"`
	ServiceIds    []*string `json:"ServiceIds,omitempty" xml:"ServiceIds,omitempty" type:"Repeated"`
	Status        *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds  []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
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
