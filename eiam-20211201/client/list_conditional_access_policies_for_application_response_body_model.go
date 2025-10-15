// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConditionalAccessPoliciesForApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConditionalAccessPolicies(v []*ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) *ListConditionalAccessPoliciesForApplicationResponseBody
	GetConditionalAccessPolicies() []*ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies
	SetRequestId(v string) *ListConditionalAccessPoliciesForApplicationResponseBody
	GetRequestId() *string
}

type ListConditionalAccessPoliciesForApplicationResponseBody struct {
	ConditionalAccessPolicies []*ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies `json:"ConditionalAccessPolicies,omitempty" xml:"ConditionalAccessPolicies,omitempty" type:"Repeated"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConditionalAccessPoliciesForApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBody) GetConditionalAccessPolicies() []*ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies {
	return s.ConditionalAccessPolicies
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBody) SetConditionalAccessPolicies(v []*ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) *ListConditionalAccessPoliciesForApplicationResponseBody {
	s.ConditionalAccessPolicies = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBody) SetRequestId(v string) *ListConditionalAccessPoliciesForApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBody) Validate() error {
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

type ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies struct {
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
	ConditionsConfig *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig `json:"ConditionsConfig,omitempty" xml:"ConditionsConfig,omitempty" type:"Struct"`
	// 创建时间
	//
	// example:
	//
	// 1741857554000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// IDaaS EIAM 条件访问策略执行Action
	DecisionConfig *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig `json:"DecisionConfig,omitempty" xml:"DecisionConfig,omitempty" type:"Struct"`
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
	// idaas_oynbcyaaejuik6b37eldzxxxxx
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

func (s ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) GetConditionalAccessPolicyId() *string {
	return s.ConditionalAccessPolicyId
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) GetConditionalAccessPolicyName() *string {
	return s.ConditionalAccessPolicyName
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) GetConditionalAccessPolicyType() *string {
	return s.ConditionalAccessPolicyType
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) GetConditionsConfig() *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig {
	return s.ConditionsConfig
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) GetDecisionConfig() *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig {
	return s.DecisionConfig
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) GetDecisionType() *string {
	return s.DecisionType
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) GetDescription() *string {
	return s.Description
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) GetEvaluateAt() *string {
	return s.EvaluateAt
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) GetLastUpdatedTime() *int64 {
	return s.LastUpdatedTime
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) GetStatus() *string {
	return s.Status
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) SetConditionalAccessPolicyId(v string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies {
	s.ConditionalAccessPolicyId = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) SetConditionalAccessPolicyName(v string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies {
	s.ConditionalAccessPolicyName = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) SetConditionalAccessPolicyType(v string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies {
	s.ConditionalAccessPolicyType = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) SetConditionsConfig(v *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies {
	s.ConditionsConfig = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) SetCreateTime(v int64) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) SetDecisionConfig(v *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies {
	s.DecisionConfig = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) SetDecisionType(v string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies {
	s.DecisionType = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) SetDescription(v string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies {
	s.Description = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) SetEvaluateAt(v string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies {
	s.EvaluateAt = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) SetInstanceId(v string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies {
	s.InstanceId = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) SetLastUpdatedTime(v int64) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies {
	s.LastUpdatedTime = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) SetPriority(v int32) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies {
	s.Priority = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) SetStatus(v string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies {
	s.Status = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPolicies) Validate() error {
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

type ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig struct {
	// 条件访问策略目标应用
	Applications *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	// 条件访问策略网络区域
	NetworkZones *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones `json:"NetworkZones,omitempty" xml:"NetworkZones,omitempty" type:"Struct"`
	// 条件访问策略目标用户
	Users *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig) GetApplications() *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigApplications {
	return s.Applications
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig) GetNetworkZones() *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones {
	return s.NetworkZones
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig) GetUsers() *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	return s.Users
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig) SetApplications(v *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigApplications) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig {
	s.Applications = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig) SetNetworkZones(v *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig {
	s.NetworkZones = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig) SetUsers(v *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig {
	s.Users = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfig) Validate() error {
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

type ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigApplications struct {
	// 排除的应用
	ExcludeApplications []*string `json:"ExcludeApplications,omitempty" xml:"ExcludeApplications,omitempty" type:"Repeated"`
	// 选择的应用
	IncludeApplications []*string `json:"IncludeApplications,omitempty" xml:"IncludeApplications,omitempty" type:"Repeated"`
}

func (s ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigApplications) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigApplications) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigApplications) GetExcludeApplications() []*string {
	return s.ExcludeApplications
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigApplications) GetIncludeApplications() []*string {
	return s.IncludeApplications
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigApplications) SetExcludeApplications(v []*string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigApplications {
	s.ExcludeApplications = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigApplications) SetIncludeApplications(v []*string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigApplications {
	s.IncludeApplications = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigApplications) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones struct {
	// 排除的网络区域
	ExcludeNetworkZones []*string `json:"ExcludeNetworkZones,omitempty" xml:"ExcludeNetworkZones,omitempty" type:"Repeated"`
	// 选择的网络区域
	IncludeNetworkZones []*string `json:"IncludeNetworkZones,omitempty" xml:"IncludeNetworkZones,omitempty" type:"Repeated"`
}

func (s ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) GetExcludeNetworkZones() []*string {
	return s.ExcludeNetworkZones
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) GetIncludeNetworkZones() []*string {
	return s.IncludeNetworkZones
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) SetExcludeNetworkZones(v []*string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones {
	s.ExcludeNetworkZones = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) SetIncludeNetworkZones(v []*string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones {
	s.IncludeNetworkZones = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers struct {
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

func (s ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetExcludeGroups() []*string {
	return s.ExcludeGroups
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetExcludeOrganizationalUnits() []*string {
	return s.ExcludeOrganizationalUnits
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetExcludeUsers() []*string {
	return s.ExcludeUsers
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetIncludeGroups() []*string {
	return s.IncludeGroups
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetIncludeOrganizationalUnits() []*string {
	return s.IncludeOrganizationalUnits
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetIncludeUsers() []*string {
	return s.IncludeUsers
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetExcludeGroups(v []*string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.ExcludeGroups = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetExcludeOrganizationalUnits(v []*string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.ExcludeOrganizationalUnits = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetExcludeUsers(v []*string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.ExcludeUsers = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetIncludeGroups(v []*string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.IncludeGroups = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetIncludeOrganizationalUnits(v []*string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.IncludeOrganizationalUnits = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetIncludeUsers(v []*string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.IncludeUsers = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesConditionsConfigUsers) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig struct {
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

func (s ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig) GetActiveSessionReuseStatus() *string {
	return s.ActiveSessionReuseStatus
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig) GetEffect() *string {
	return s.Effect
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig) GetMfaAuthenticationIntervalSeconds() *int64 {
	return s.MfaAuthenticationIntervalSeconds
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig) GetMfaAuthenticationMethods() []*string {
	return s.MfaAuthenticationMethods
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig) GetMfaType() *string {
	return s.MfaType
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig) SetActiveSessionReuseStatus(v string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.ActiveSessionReuseStatus = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig) SetEffect(v string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.Effect = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig) SetMfaAuthenticationIntervalSeconds(v int64) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.MfaAuthenticationIntervalSeconds = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig) SetMfaAuthenticationMethods(v []*string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.MfaAuthenticationMethods = v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig) SetMfaType(v string) *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.MfaType = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationResponseBodyConditionalAccessPoliciesDecisionConfig) Validate() error {
	return dara.Validate(s)
}
