// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConditionalAccessPoliciesForNetworkZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConditionalAccessPolicies(v []*ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) *ListConditionalAccessPoliciesForNetworkZoneResponseBody
	GetConditionalAccessPolicies() []*ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies
	SetRequestId(v string) *ListConditionalAccessPoliciesForNetworkZoneResponseBody
	GetRequestId() *string
}

type ListConditionalAccessPoliciesForNetworkZoneResponseBody struct {
	// Collection of conditional access policies
	ConditionalAccessPolicies []*ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies `json:"ConditionalAccessPolicies,omitempty" xml:"ConditionalAccessPolicies,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponseBody) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBody) GetConditionalAccessPolicies() []*ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies {
	return s.ConditionalAccessPolicies
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBody) SetConditionalAccessPolicies(v []*ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) *ListConditionalAccessPoliciesForNetworkZoneResponseBody {
	s.ConditionalAccessPolicies = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBody) SetRequestId(v string) *ListConditionalAccessPoliciesForNetworkZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies struct {
	// Conditional access policy ID
	//
	// example:
	//
	// cp_xxxxx
	ConditionalAccessPolicyId *string `json:"ConditionalAccessPolicyId,omitempty" xml:"ConditionalAccessPolicyId,omitempty"`
	// Conditional access policy name
	//
	// example:
	//
	// My Policy
	ConditionalAccessPolicyName *string `json:"ConditionalAccessPolicyName,omitempty" xml:"ConditionalAccessPolicyName,omitempty"`
	// Type of the conditional access policy
	//
	// example:
	//
	// arn:alibaba:idaas:authn:access:policy:system
	ConditionalAccessPolicyType *string `json:"ConditionalAccessPolicyType,omitempty" xml:"ConditionalAccessPolicyType,omitempty"`
	// Content of the conditional access policy
	ConditionsConfig *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig `json:"ConditionsConfig,omitempty" xml:"ConditionsConfig,omitempty" type:"Struct"`
	// Creation Time
	//
	// example:
	//
	// 1741857554000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Action of the conditional access policy
	DecisionConfig *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig `json:"DecisionConfig,omitempty" xml:"DecisionConfig,omitempty" type:"Struct"`
	// Execution type of the conditional access policy
	//
	// example:
	//
	// enforcement
	DecisionType *string `json:"DecisionType,omitempty" xml:"DecisionType,omitempty"`
	// Description of the conditional access policy
	//
	// example:
	//
	// terraform-example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Execution point of the conditional access policy
	//
	// example:
	//
	// arn:alibaba:idaas:authn:access:rule:eval_at:after_step1
	EvaluateAt *string `json:"EvaluateAt,omitempty" xml:"EvaluateAt,omitempty"`
	// Instance ID
	//
	// example:
	//
	// idaas_oynbcyaaejuik6b37eldz4pinu
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Last Updated Time
	//
	// example:
	//
	// 1741857554000
	LastUpdatedTime *int64 `json:"LastUpdatedTime,omitempty" xml:"LastUpdatedTime,omitempty"`
	// Priority
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Enable or disable status of the conditional access policy
	//
	// example:
	//
	// disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) GetConditionalAccessPolicyId() *string {
	return s.ConditionalAccessPolicyId
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) GetConditionalAccessPolicyName() *string {
	return s.ConditionalAccessPolicyName
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) GetConditionalAccessPolicyType() *string {
	return s.ConditionalAccessPolicyType
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) GetConditionsConfig() *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig {
	return s.ConditionsConfig
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) GetDecisionConfig() *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig {
	return s.DecisionConfig
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) GetDecisionType() *string {
	return s.DecisionType
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) GetDescription() *string {
	return s.Description
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) GetEvaluateAt() *string {
	return s.EvaluateAt
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) GetLastUpdatedTime() *int64 {
	return s.LastUpdatedTime
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) GetStatus() *string {
	return s.Status
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) SetConditionalAccessPolicyId(v string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies {
	s.ConditionalAccessPolicyId = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) SetConditionalAccessPolicyName(v string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies {
	s.ConditionalAccessPolicyName = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) SetConditionalAccessPolicyType(v string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies {
	s.ConditionalAccessPolicyType = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) SetConditionsConfig(v *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies {
	s.ConditionsConfig = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) SetCreateTime(v int64) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) SetDecisionConfig(v *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies {
	s.DecisionConfig = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) SetDecisionType(v string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies {
	s.DecisionType = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) SetDescription(v string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies {
	s.Description = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) SetEvaluateAt(v string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies {
	s.EvaluateAt = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) SetInstanceId(v string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies {
	s.InstanceId = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) SetLastUpdatedTime(v int64) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies {
	s.LastUpdatedTime = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) SetPriority(v int32) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies {
	s.Priority = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) SetStatus(v string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies {
	s.Status = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPolicies) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig struct {
	// Target applications of the conditional access policy
	Applications *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	// Network zones for conditional access policies
	NetworkZones *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones `json:"NetworkZones,omitempty" xml:"NetworkZones,omitempty" type:"Struct"`
	// Target users of the conditional access policy
	Users *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig) GetApplications() *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigApplications {
	return s.Applications
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig) GetNetworkZones() *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones {
	return s.NetworkZones
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig) GetUsers() *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	return s.Users
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig) SetApplications(v *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigApplications) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig {
	s.Applications = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig) SetNetworkZones(v *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig {
	s.NetworkZones = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig) SetUsers(v *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig {
	s.Users = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfig) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigApplications struct {
	// Excluded applications
	ExcludeApplications []*string `json:"ExcludeApplications,omitempty" xml:"ExcludeApplications,omitempty" type:"Repeated"`
	// Selected applications
	IncludeApplications []*string `json:"IncludeApplications,omitempty" xml:"IncludeApplications,omitempty" type:"Repeated"`
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigApplications) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigApplications) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigApplications) GetExcludeApplications() []*string {
	return s.ExcludeApplications
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigApplications) GetIncludeApplications() []*string {
	return s.IncludeApplications
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigApplications) SetExcludeApplications(v []*string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigApplications {
	s.ExcludeApplications = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigApplications) SetIncludeApplications(v []*string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigApplications {
	s.IncludeApplications = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigApplications) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones struct {
	// Excluded network zones
	ExcludeNetworkZones []*string `json:"ExcludeNetworkZones,omitempty" xml:"ExcludeNetworkZones,omitempty" type:"Repeated"`
	// Included network zones
	IncludeNetworkZones []*string `json:"IncludeNetworkZones,omitempty" xml:"IncludeNetworkZones,omitempty" type:"Repeated"`
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) GetExcludeNetworkZones() []*string {
	return s.ExcludeNetworkZones
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) GetIncludeNetworkZones() []*string {
	return s.IncludeNetworkZones
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) SetExcludeNetworkZones(v []*string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones {
	s.ExcludeNetworkZones = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) SetIncludeNetworkZones(v []*string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones {
	s.IncludeNetworkZones = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers struct {
	// Excluded user groups
	ExcludeGroups []*string `json:"ExcludeGroups,omitempty" xml:"ExcludeGroups,omitempty" type:"Repeated"`
	// Excluded organizations
	ExcludeOrganizationalUnits []*string `json:"ExcludeOrganizationalUnits,omitempty" xml:"ExcludeOrganizationalUnits,omitempty" type:"Repeated"`
	// Excluded Users
	ExcludeUsers []*string `json:"ExcludeUsers,omitempty" xml:"ExcludeUsers,omitempty" type:"Repeated"`
	// Selected user groups
	IncludeGroups []*string `json:"IncludeGroups,omitempty" xml:"IncludeGroups,omitempty" type:"Repeated"`
	// Selected organizations
	IncludeOrganizationalUnits []*string `json:"IncludeOrganizationalUnits,omitempty" xml:"IncludeOrganizationalUnits,omitempty" type:"Repeated"`
	// Selected users
	IncludeUsers []*string `json:"IncludeUsers,omitempty" xml:"IncludeUsers,omitempty" type:"Repeated"`
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetExcludeGroups() []*string {
	return s.ExcludeGroups
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetExcludeOrganizationalUnits() []*string {
	return s.ExcludeOrganizationalUnits
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetExcludeUsers() []*string {
	return s.ExcludeUsers
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetIncludeGroups() []*string {
	return s.IncludeGroups
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetIncludeOrganizationalUnits() []*string {
	return s.IncludeOrganizationalUnits
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetIncludeUsers() []*string {
	return s.IncludeUsers
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetExcludeGroups(v []*string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.ExcludeGroups = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetExcludeOrganizationalUnits(v []*string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.ExcludeOrganizationalUnits = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetExcludeUsers(v []*string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.ExcludeUsers = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetIncludeGroups(v []*string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.IncludeGroups = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetIncludeOrganizationalUnits(v []*string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.IncludeOrganizationalUnits = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetIncludeUsers(v []*string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.IncludeUsers = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesConditionsConfigUsers) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig struct {
	// Whether session reuse is enabled
	//
	// example:
	//
	// enabled
	ActiveSessionReuseStatus *string `json:"ActiveSessionReuseStatus,omitempty" xml:"ActiveSessionReuseStatus,omitempty"`
	// Decision action of the conditional access policy
	//
	// example:
	//
	// allow
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// Re-authentication interval for conditional access policy (seconds)
	//
	// example:
	//
	// 300
	MfaAuthenticationIntervalSeconds *int64 `json:"MfaAuthenticationIntervalSeconds,omitempty" xml:"MfaAuthenticationIntervalSeconds,omitempty"`
	// MFA types allowed by the conditional access policy
	MfaAuthenticationMethods []*string `json:"MfaAuthenticationMethods,omitempty" xml:"MfaAuthenticationMethods,omitempty" type:"Repeated"`
	// MFA type of the conditional access policy
	//
	// example:
	//
	// directly_access
	MfaType *string `json:"MfaType,omitempty" xml:"MfaType,omitempty"`
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig) GetActiveSessionReuseStatus() *string {
	return s.ActiveSessionReuseStatus
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig) GetEffect() *string {
	return s.Effect
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig) GetMfaAuthenticationIntervalSeconds() *int64 {
	return s.MfaAuthenticationIntervalSeconds
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig) GetMfaAuthenticationMethods() []*string {
	return s.MfaAuthenticationMethods
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig) GetMfaType() *string {
	return s.MfaType
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig) SetActiveSessionReuseStatus(v string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.ActiveSessionReuseStatus = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig) SetEffect(v string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.Effect = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig) SetMfaAuthenticationIntervalSeconds(v int64) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.MfaAuthenticationIntervalSeconds = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig) SetMfaAuthenticationMethods(v []*string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.MfaAuthenticationMethods = v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig) SetMfaType(v string) *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.MfaType = &v
	return s
}

func (s *ListConditionalAccessPoliciesForNetworkZoneResponseBodyConditionalAccessPoliciesDecisionConfig) Validate() error {
	return dara.Validate(s)
}
