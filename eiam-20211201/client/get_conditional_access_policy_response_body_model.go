// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConditionalAccessPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConditionalAccessPolicy(v *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) *GetConditionalAccessPolicyResponseBody
	GetConditionalAccessPolicy() *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy
	SetRequestId(v string) *GetConditionalAccessPolicyResponseBody
	GetRequestId() *string
}

type GetConditionalAccessPolicyResponseBody struct {
	// The details of the conditional access policy.
	ConditionalAccessPolicy *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy `json:"ConditionalAccessPolicy,omitempty" xml:"ConditionalAccessPolicy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConditionalAccessPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConditionalAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetConditionalAccessPolicyResponseBody) GetConditionalAccessPolicy() *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy {
	return s.ConditionalAccessPolicy
}

func (s *GetConditionalAccessPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConditionalAccessPolicyResponseBody) SetConditionalAccessPolicy(v *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) *GetConditionalAccessPolicyResponseBody {
	s.ConditionalAccessPolicy = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBody) SetRequestId(v string) *GetConditionalAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBody) Validate() error {
	if s.ConditionalAccessPolicy != nil {
		if err := s.ConditionalAccessPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy struct {
	// The ID of the conditional access policy.
	//
	// example:
	//
	// cp_xxxxx
	ConditionalAccessPolicyId *string `json:"ConditionalAccessPolicyId,omitempty" xml:"ConditionalAccessPolicyId,omitempty"`
	// The name of the conditional access policy.
	//
	// example:
	//
	// My policy
	ConditionalAccessPolicyName *string `json:"ConditionalAccessPolicyName,omitempty" xml:"ConditionalAccessPolicyName,omitempty"`
	// The type of the conditional access policy.
	//
	// example:
	//
	// arn:alibaba:idaas:authn:access:policy:system
	ConditionalAccessPolicyType *string `json:"ConditionalAccessPolicyType,omitempty" xml:"ConditionalAccessPolicyType,omitempty"`
	// The content of the conditional access policy.
	ConditionsConfig *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig `json:"ConditionsConfig,omitempty" xml:"ConditionsConfig,omitempty" type:"Struct"`
	// The time when the policy was created.
	//
	// example:
	//
	// 1741857554000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The action of the conditional access policy.
	DecisionConfig *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig `json:"DecisionConfig,omitempty" xml:"DecisionConfig,omitempty" type:"Struct"`
	// The execution type of the conditional access policy.
	//
	// example:
	//
	// enforcement
	DecisionType *string `json:"DecisionType,omitempty" xml:"DecisionType,omitempty"`
	// The description of the conditional access policy.
	//
	// example:
	//
	// ga access port for ecs: internal-cn-hangzhou-docker-builder-2(i-bp19g1pheaailkk1xvr6)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The policy enforcement point of the conditional access policy.
	//
	// example:
	//
	// arn:alibaba:idaas:authn:access:rule:eval_at:after_step1
	EvaluateAt *string `json:"EvaluateAt,omitempty" xml:"EvaluateAt,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_qnx6fbrinlecptl5hld23lfkvy
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the policy was last updated.
	//
	// example:
	//
	// 1741857554000
	LastUpdatedTime *int64 `json:"LastUpdatedTime,omitempty" xml:"LastUpdatedTime,omitempty"`
	// The priority.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The status of the conditional access policy.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) GoString() string {
	return s.String()
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) GetConditionalAccessPolicyId() *string {
	return s.ConditionalAccessPolicyId
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) GetConditionalAccessPolicyName() *string {
	return s.ConditionalAccessPolicyName
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) GetConditionalAccessPolicyType() *string {
	return s.ConditionalAccessPolicyType
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) GetConditionsConfig() *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig {
	return s.ConditionsConfig
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) GetDecisionConfig() *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig {
	return s.DecisionConfig
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) GetDecisionType() *string {
	return s.DecisionType
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) GetDescription() *string {
	return s.Description
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) GetEvaluateAt() *string {
	return s.EvaluateAt
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) GetLastUpdatedTime() *int64 {
	return s.LastUpdatedTime
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) GetPriority() *int32 {
	return s.Priority
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) GetStatus() *string {
	return s.Status
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) SetConditionalAccessPolicyId(v string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy {
	s.ConditionalAccessPolicyId = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) SetConditionalAccessPolicyName(v string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy {
	s.ConditionalAccessPolicyName = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) SetConditionalAccessPolicyType(v string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy {
	s.ConditionalAccessPolicyType = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) SetConditionsConfig(v *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy {
	s.ConditionsConfig = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) SetCreateTime(v int64) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy {
	s.CreateTime = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) SetDecisionConfig(v *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy {
	s.DecisionConfig = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) SetDecisionType(v string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy {
	s.DecisionType = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) SetDescription(v string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy {
	s.Description = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) SetEvaluateAt(v string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy {
	s.EvaluateAt = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) SetInstanceId(v string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy {
	s.InstanceId = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) SetLastUpdatedTime(v int64) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy {
	s.LastUpdatedTime = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) SetPriority(v int32) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy {
	s.Priority = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) SetStatus(v string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy {
	s.Status = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicy) Validate() error {
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

type GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig struct {
	// The applications to which the conditional access policy applies.
	Applications *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	// The network zones of the conditional access policy.
	NetworkZones *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigNetworkZones `json:"NetworkZones,omitempty" xml:"NetworkZones,omitempty" type:"Struct"`
	// The users to whom the conditional access policy applies.
	Users *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig) String() string {
	return dara.Prettify(s)
}

func (s GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig) GoString() string {
	return s.String()
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig) GetApplications() *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigApplications {
	return s.Applications
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig) GetNetworkZones() *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigNetworkZones {
	return s.NetworkZones
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig) GetUsers() *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers {
	return s.Users
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig) SetApplications(v *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigApplications) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig {
	s.Applications = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig) SetNetworkZones(v *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigNetworkZones) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig {
	s.NetworkZones = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig) SetUsers(v *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig {
	s.Users = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfig) Validate() error {
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

type GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigApplications struct {
	// The excluded applications.
	ExcludeApplications []*string `json:"ExcludeApplications,omitempty" xml:"ExcludeApplications,omitempty" type:"Repeated"`
	// The selected applications.
	IncludeApplications []*string `json:"IncludeApplications,omitempty" xml:"IncludeApplications,omitempty" type:"Repeated"`
}

func (s GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigApplications) String() string {
	return dara.Prettify(s)
}

func (s GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigApplications) GoString() string {
	return s.String()
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigApplications) GetExcludeApplications() []*string {
	return s.ExcludeApplications
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigApplications) GetIncludeApplications() []*string {
	return s.IncludeApplications
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigApplications) SetExcludeApplications(v []*string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigApplications {
	s.ExcludeApplications = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigApplications) SetIncludeApplications(v []*string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigApplications {
	s.IncludeApplications = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigApplications) Validate() error {
	return dara.Validate(s)
}

type GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigNetworkZones struct {
	// The excluded network zones.
	ExcludeNetworkZones []*string `json:"ExcludeNetworkZones,omitempty" xml:"ExcludeNetworkZones,omitempty" type:"Repeated"`
	// The selected network zones.
	IncludeNetworkZones []*string `json:"IncludeNetworkZones,omitempty" xml:"IncludeNetworkZones,omitempty" type:"Repeated"`
}

func (s GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigNetworkZones) String() string {
	return dara.Prettify(s)
}

func (s GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigNetworkZones) GoString() string {
	return s.String()
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigNetworkZones) GetExcludeNetworkZones() []*string {
	return s.ExcludeNetworkZones
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigNetworkZones) GetIncludeNetworkZones() []*string {
	return s.IncludeNetworkZones
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigNetworkZones) SetExcludeNetworkZones(v []*string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigNetworkZones {
	s.ExcludeNetworkZones = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigNetworkZones) SetIncludeNetworkZones(v []*string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigNetworkZones {
	s.IncludeNetworkZones = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigNetworkZones) Validate() error {
	return dara.Validate(s)
}

type GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers struct {
	// The excluded user groups.
	ExcludeGroups []*string `json:"ExcludeGroups,omitempty" xml:"ExcludeGroups,omitempty" type:"Repeated"`
	// The excluded organizations.
	ExcludeOrganizationalUnits []*string `json:"ExcludeOrganizationalUnits,omitempty" xml:"ExcludeOrganizationalUnits,omitempty" type:"Repeated"`
	// The excluded users.
	ExcludeUsers []*string `json:"ExcludeUsers,omitempty" xml:"ExcludeUsers,omitempty" type:"Repeated"`
	// The selected user groups.
	IncludeGroups []*string `json:"IncludeGroups,omitempty" xml:"IncludeGroups,omitempty" type:"Repeated"`
	// The selected organizations.
	IncludeOrganizationalUnits []*string `json:"IncludeOrganizationalUnits,omitempty" xml:"IncludeOrganizationalUnits,omitempty" type:"Repeated"`
	// The selected users.
	IncludeUsers []*string `json:"IncludeUsers,omitempty" xml:"IncludeUsers,omitempty" type:"Repeated"`
}

func (s GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) String() string {
	return dara.Prettify(s)
}

func (s GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) GoString() string {
	return s.String()
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) GetExcludeGroups() []*string {
	return s.ExcludeGroups
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) GetExcludeOrganizationalUnits() []*string {
	return s.ExcludeOrganizationalUnits
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) GetExcludeUsers() []*string {
	return s.ExcludeUsers
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) GetIncludeGroups() []*string {
	return s.IncludeGroups
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) GetIncludeOrganizationalUnits() []*string {
	return s.IncludeOrganizationalUnits
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) GetIncludeUsers() []*string {
	return s.IncludeUsers
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) SetExcludeGroups(v []*string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers {
	s.ExcludeGroups = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) SetExcludeOrganizationalUnits(v []*string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers {
	s.ExcludeOrganizationalUnits = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) SetExcludeUsers(v []*string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers {
	s.ExcludeUsers = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) SetIncludeGroups(v []*string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers {
	s.IncludeGroups = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) SetIncludeOrganizationalUnits(v []*string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers {
	s.IncludeOrganizationalUnits = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) SetIncludeUsers(v []*string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers {
	s.IncludeUsers = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyConditionsConfigUsers) Validate() error {
	return dara.Validate(s)
}

type GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig struct {
	// Indicates whether session reuse is enabled.
	//
	// example:
	//
	// enabled
	ActiveSessionReuseStatus *string `json:"ActiveSessionReuseStatus,omitempty" xml:"ActiveSessionReuseStatus,omitempty"`
	// The decision action of the conditional access policy.
	//
	// example:
	//
	// allow
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// The interval in seconds at which re-authentication is required for the conditional access policy.
	//
	// example:
	//
	// 300
	MfaAuthenticationIntervalSeconds *int64 `json:"MfaAuthenticationIntervalSeconds,omitempty" xml:"MfaAuthenticationIntervalSeconds,omitempty"`
	// The MFA methods that are allowed by the conditional access policy.
	MfaAuthenticationMethods []*string `json:"MfaAuthenticationMethods,omitempty" xml:"MfaAuthenticationMethods,omitempty" type:"Repeated"`
	// The multi-factor authentication (MFA) type of the conditional access policy.
	//
	// example:
	//
	// directly_access
	MfaType *string `json:"MfaType,omitempty" xml:"MfaType,omitempty"`
}

func (s GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig) GoString() string {
	return s.String()
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig) GetActiveSessionReuseStatus() *string {
	return s.ActiveSessionReuseStatus
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig) GetEffect() *string {
	return s.Effect
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig) GetMfaAuthenticationIntervalSeconds() *int64 {
	return s.MfaAuthenticationIntervalSeconds
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig) GetMfaAuthenticationMethods() []*string {
	return s.MfaAuthenticationMethods
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig) GetMfaType() *string {
	return s.MfaType
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig) SetActiveSessionReuseStatus(v string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig {
	s.ActiveSessionReuseStatus = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig) SetEffect(v string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig {
	s.Effect = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig) SetMfaAuthenticationIntervalSeconds(v int64) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig {
	s.MfaAuthenticationIntervalSeconds = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig) SetMfaAuthenticationMethods(v []*string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig {
	s.MfaAuthenticationMethods = v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig) SetMfaType(v string) *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig {
	s.MfaType = &v
	return s
}

func (s *GetConditionalAccessPolicyResponseBodyConditionalAccessPolicyDecisionConfig) Validate() error {
	return dara.Validate(s)
}
