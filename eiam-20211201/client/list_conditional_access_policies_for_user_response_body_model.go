// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConditionalAccessPoliciesForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConditionalAccessPolicies(v []*ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) *ListConditionalAccessPoliciesForUserResponseBody
	GetConditionalAccessPolicies() []*ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies
	SetRequestId(v string) *ListConditionalAccessPoliciesForUserResponseBody
	GetRequestId() *string
}

type ListConditionalAccessPoliciesForUserResponseBody struct {
	ConditionalAccessPolicies []*ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies `json:"ConditionalAccessPolicies,omitempty" xml:"ConditionalAccessPolicies,omitempty" type:"Repeated"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConditionalAccessPoliciesForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForUserResponseBody) GetConditionalAccessPolicies() []*ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies {
	return s.ConditionalAccessPolicies
}

func (s *ListConditionalAccessPoliciesForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConditionalAccessPoliciesForUserResponseBody) SetConditionalAccessPolicies(v []*ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) *ListConditionalAccessPoliciesForUserResponseBody {
	s.ConditionalAccessPolicies = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBody) SetRequestId(v string) *ListConditionalAccessPoliciesForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBody) Validate() error {
	if s.ConditionalAccessPolicies != nil {
		for _, item := range s.ConditionalAccessPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies struct {
	// IDaaS EIAM 条件访问策略Id
	//
	// example:
	//
	// cap_m5etun43kejxphsbke6mjxxxxx
	ConditionalAccessPolicyId *string `json:"ConditionalAccessPolicyId,omitempty" xml:"ConditionalAccessPolicyId,omitempty"`
	// IDaaS EIAM 条件访问策略名称
	ConditionalAccessPolicyName *string `json:"ConditionalAccessPolicyName,omitempty" xml:"ConditionalAccessPolicyName,omitempty"`
	// IDaaS EIAM 条件访问策略类型
	//
	// example:
	//
	// arn:alibaba:idaas:authn:access:policy:system
	ConditionalAccessPolicyType *string `json:"ConditionalAccessPolicyType,omitempty" xml:"ConditionalAccessPolicyType,omitempty"`
	// IDaaS EIAM 条件访问策略内容
	ConditionsConfig *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig `json:"ConditionsConfig,omitempty" xml:"ConditionsConfig,omitempty" type:"Struct"`
	// 创建时间
	//
	// example:
	//
	// 1741857554000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// IDaaS EIAM 条件访问策略执行Action
	DecisionConfig *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig `json:"DecisionConfig,omitempty" xml:"DecisionConfig,omitempty" type:"Struct"`
	// IDaaS EIAM 条件访问策略执行类型
	//
	// example:
	//
	// enforcement
	DecisionType *string `json:"DecisionType,omitempty" xml:"DecisionType,omitempty"`
	// IDaaS EIAM 条件访问策略描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// IDaaS EIAM 条件访问策略执行点
	//
	// example:
	//
	// arn:alibaba:idaas:authn:access:rule:eval_at:after_step1
	EvaluateAt *string `json:"EvaluateAt,omitempty" xml:"EvaluateAt,omitempty"`
	// IDaaS EIAM 实例Id
	//
	// example:
	//
	// idaas_oynbcyaaejuik6b37eldz4xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 更新时间
	//
	// example:
	//
	// 1741857554000
	LastUpdatedTime *int64 `json:"LastUpdatedTime,omitempty" xml:"LastUpdatedTime,omitempty"`
	// 优先级
	//
	// example:
	//
	// 100
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// IDaaS EIAM 条件访问策略启用禁用状态
	//
	// example:
	//
	// disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) GetConditionalAccessPolicyId() *string {
	return s.ConditionalAccessPolicyId
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) GetConditionalAccessPolicyName() *string {
	return s.ConditionalAccessPolicyName
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) GetConditionalAccessPolicyType() *string {
	return s.ConditionalAccessPolicyType
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) GetConditionsConfig() *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig {
	return s.ConditionsConfig
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) GetDecisionConfig() *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig {
	return s.DecisionConfig
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) GetDecisionType() *string {
	return s.DecisionType
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) GetDescription() *string {
	return s.Description
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) GetEvaluateAt() *string {
	return s.EvaluateAt
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) GetLastUpdatedTime() *int64 {
	return s.LastUpdatedTime
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) GetStatus() *string {
	return s.Status
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) SetConditionalAccessPolicyId(v string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies {
	s.ConditionalAccessPolicyId = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) SetConditionalAccessPolicyName(v string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies {
	s.ConditionalAccessPolicyName = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) SetConditionalAccessPolicyType(v string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies {
	s.ConditionalAccessPolicyType = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) SetConditionsConfig(v *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies {
	s.ConditionsConfig = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) SetCreateTime(v int64) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) SetDecisionConfig(v *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies {
	s.DecisionConfig = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) SetDecisionType(v string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies {
	s.DecisionType = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) SetDescription(v string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies {
	s.Description = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) SetEvaluateAt(v string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies {
	s.EvaluateAt = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) SetInstanceId(v string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies {
	s.InstanceId = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) SetLastUpdatedTime(v int64) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies {
	s.LastUpdatedTime = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) SetPriority(v int32) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies {
	s.Priority = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) SetStatus(v string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies {
	s.Status = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPolicies) Validate() error {
	if s.ConditionsConfig != nil {
		if err := s.ConditionsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DecisionConfig != nil {
		if err := s.DecisionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig struct {
	// 条件访问策略目标应用
	Applications *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	// 条件访问策略网络区域
	NetworkZones *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones `json:"NetworkZones,omitempty" xml:"NetworkZones,omitempty" type:"Struct"`
	// 条件访问策略目标用户
	Users *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig) GetApplications() *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigApplications {
	return s.Applications
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig) GetNetworkZones() *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones {
	return s.NetworkZones
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig) GetUsers() *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	return s.Users
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig) SetApplications(v *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigApplications) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig {
	s.Applications = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig) SetNetworkZones(v *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig {
	s.NetworkZones = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig) SetUsers(v *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig {
	s.Users = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfig) Validate() error {
	if s.Applications != nil {
		if err := s.Applications.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkZones != nil {
		if err := s.NetworkZones.Validate(); err != nil {
			return err
		}
	}
	if s.Users != nil {
		if err := s.Users.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigApplications struct {
	// 排除的应用
	ExcludeApplications []*string `json:"ExcludeApplications,omitempty" xml:"ExcludeApplications,omitempty" type:"Repeated"`
	// 选择的应用
	IncludeApplications []*string `json:"IncludeApplications,omitempty" xml:"IncludeApplications,omitempty" type:"Repeated"`
}

func (s ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigApplications) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigApplications) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigApplications) GetExcludeApplications() []*string {
	return s.ExcludeApplications
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigApplications) GetIncludeApplications() []*string {
	return s.IncludeApplications
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigApplications) SetExcludeApplications(v []*string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigApplications {
	s.ExcludeApplications = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigApplications) SetIncludeApplications(v []*string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigApplications {
	s.IncludeApplications = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigApplications) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones struct {
	// 排除的网络区域
	ExcludeNetworkZones []*string `json:"ExcludeNetworkZones,omitempty" xml:"ExcludeNetworkZones,omitempty" type:"Repeated"`
	// 选择的网络区域
	IncludeNetworkZones []*string `json:"IncludeNetworkZones,omitempty" xml:"IncludeNetworkZones,omitempty" type:"Repeated"`
}

func (s ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) GetExcludeNetworkZones() []*string {
	return s.ExcludeNetworkZones
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) GetIncludeNetworkZones() []*string {
	return s.IncludeNetworkZones
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) SetExcludeNetworkZones(v []*string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones {
	s.ExcludeNetworkZones = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) SetIncludeNetworkZones(v []*string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones {
	s.IncludeNetworkZones = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers struct {
	// 排除的用户组
	ExcludeGroups []*string `json:"ExcludeGroups,omitempty" xml:"ExcludeGroups,omitempty" type:"Repeated"`
	// 排除的组织
	ExcludeOrganizationalUnits []*string `json:"ExcludeOrganizationalUnits,omitempty" xml:"ExcludeOrganizationalUnits,omitempty" type:"Repeated"`
	// 排除的用户
	ExcludeUsers []*string `json:"ExcludeUsers,omitempty" xml:"ExcludeUsers,omitempty" type:"Repeated"`
	// 选择的用户组
	IncludeGroups []*string `json:"IncludeGroups,omitempty" xml:"IncludeGroups,omitempty" type:"Repeated"`
	// 选择的组织
	IncludeOrganizationalUnits []*string `json:"IncludeOrganizationalUnits,omitempty" xml:"IncludeOrganizationalUnits,omitempty" type:"Repeated"`
	// 选择的用户
	IncludeUsers []*string `json:"IncludeUsers,omitempty" xml:"IncludeUsers,omitempty" type:"Repeated"`
}

func (s ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetExcludeGroups() []*string {
	return s.ExcludeGroups
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetExcludeOrganizationalUnits() []*string {
	return s.ExcludeOrganizationalUnits
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetExcludeUsers() []*string {
	return s.ExcludeUsers
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetIncludeGroups() []*string {
	return s.IncludeGroups
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetIncludeOrganizationalUnits() []*string {
	return s.IncludeOrganizationalUnits
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetIncludeUsers() []*string {
	return s.IncludeUsers
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetExcludeGroups(v []*string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.ExcludeGroups = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetExcludeOrganizationalUnits(v []*string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.ExcludeOrganizationalUnits = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetExcludeUsers(v []*string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.ExcludeUsers = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetIncludeGroups(v []*string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.IncludeGroups = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetIncludeOrganizationalUnits(v []*string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.IncludeOrganizationalUnits = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetIncludeUsers(v []*string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.IncludeUsers = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesConditionsConfigUsers) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig struct {
	// example:
	//
	// disabled
	ActiveSessionReuseStatus *string `json:"ActiveSessionReuseStatus,omitempty" xml:"ActiveSessionReuseStatus,omitempty"`
	// IDaaS EIAM 条件访问策略决策Action
	//
	// example:
	//
	// allow
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// IDaaS EIAM 条件访问策略重复认证时间间隔(秒)
	//
	// example:
	//
	// 300
	MfaAuthenticationIntervalSeconds *int64 `json:"MfaAuthenticationIntervalSeconds,omitempty" xml:"MfaAuthenticationIntervalSeconds,omitempty"`
	// IDaaS EIAM 条件访问策略允许使用的mfa类型
	MfaAuthenticationMethods []*string `json:"MfaAuthenticationMethods,omitempty" xml:"MfaAuthenticationMethods,omitempty" type:"Repeated"`
	// IDaaS EIAM 条件访问策略Mfa类型
	//
	// example:
	//
	// directly_access
	MfaType *string `json:"MfaType,omitempty" xml:"MfaType,omitempty"`
}

func (s ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig) GetActiveSessionReuseStatus() *string {
	return s.ActiveSessionReuseStatus
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig) GetEffect() *string {
	return s.Effect
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig) GetMfaAuthenticationIntervalSeconds() *int64 {
	return s.MfaAuthenticationIntervalSeconds
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig) GetMfaAuthenticationMethods() []*string {
	return s.MfaAuthenticationMethods
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig) GetMfaType() *string {
	return s.MfaType
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig) SetActiveSessionReuseStatus(v string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.ActiveSessionReuseStatus = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig) SetEffect(v string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.Effect = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig) SetMfaAuthenticationIntervalSeconds(v int64) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.MfaAuthenticationIntervalSeconds = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig) SetMfaAuthenticationMethods(v []*string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.MfaAuthenticationMethods = v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig) SetMfaType(v string) *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.MfaType = &v
	return s
}

func (s *ListConditionalAccessPoliciesForUserResponseBodyConditionalAccessPoliciesDecisionConfig) Validate() error {
	return dara.Validate(s)
}
