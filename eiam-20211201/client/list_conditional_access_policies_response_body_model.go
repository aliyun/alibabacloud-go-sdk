// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConditionalAccessPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConditionalAccessPolicies(v []*ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) *ListConditionalAccessPoliciesResponseBody
	GetConditionalAccessPolicies() []*ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies
	SetNextToken(v string) *ListConditionalAccessPoliciesResponseBody
	GetNextToken() *string
	SetPreviousToken(v string) *ListConditionalAccessPoliciesResponseBody
	GetPreviousToken() *string
	SetRequestId(v string) *ListConditionalAccessPoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListConditionalAccessPoliciesResponseBody
	GetTotalCount() *int64
}

type ListConditionalAccessPoliciesResponseBody struct {
	// Collection of conditional access policies
	ConditionalAccessPolicies []*ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies `json:"ConditionalAccessPolicies,omitempty" xml:"ConditionalAccessPolicies,omitempty" type:"Repeated"`
	// The token value returned by this call for the next page query.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Previous page query token (Token)
	//
	// example:
	//
	// PTxxxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of items in the list.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListConditionalAccessPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesResponseBody) GetConditionalAccessPolicies() []*ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies {
	return s.ConditionalAccessPolicies
}

func (s *ListConditionalAccessPoliciesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConditionalAccessPoliciesResponseBody) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListConditionalAccessPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConditionalAccessPoliciesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListConditionalAccessPoliciesResponseBody) SetConditionalAccessPolicies(v []*ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) *ListConditionalAccessPoliciesResponseBody {
	s.ConditionalAccessPolicies = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBody) SetNextToken(v string) *ListConditionalAccessPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBody) SetPreviousToken(v string) *ListConditionalAccessPoliciesResponseBody {
	s.PreviousToken = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBody) SetRequestId(v string) *ListConditionalAccessPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBody) SetTotalCount(v int64) *ListConditionalAccessPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies struct {
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
	ConditionsConfig *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig `json:"ConditionsConfig,omitempty" xml:"ConditionsConfig,omitempty" type:"Struct"`
	// Creation time
	//
	// example:
	//
	// 1741857554000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Action of the conditional access policy
	DecisionConfig *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig `json:"DecisionConfig,omitempty" xml:"DecisionConfig,omitempty" type:"Struct"`
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
	// My Policy Description
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
	// idaas_ksvv5c7f2l6uzh6oqspeks23ni
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Last updated time
	//
	// example:
	//
	// 1741857554000
	LastUpdatedTime *int64 `json:"LastUpdatedTime,omitempty" xml:"LastUpdatedTime,omitempty"`
	// Priority, 1-100
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Enable or disable status of the conditional access policy
	//
	// example:
	//
	// disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) GetConditionalAccessPolicyId() *string {
	return s.ConditionalAccessPolicyId
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) GetConditionalAccessPolicyName() *string {
	return s.ConditionalAccessPolicyName
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) GetConditionalAccessPolicyType() *string {
	return s.ConditionalAccessPolicyType
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) GetConditionsConfig() *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig {
	return s.ConditionsConfig
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) GetDecisionConfig() *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig {
	return s.DecisionConfig
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) GetDecisionType() *string {
	return s.DecisionType
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) GetDescription() *string {
	return s.Description
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) GetEvaluateAt() *string {
	return s.EvaluateAt
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) GetLastUpdatedTime() *int64 {
	return s.LastUpdatedTime
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) GetStatus() *string {
	return s.Status
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) SetConditionalAccessPolicyId(v string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies {
	s.ConditionalAccessPolicyId = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) SetConditionalAccessPolicyName(v string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies {
	s.ConditionalAccessPolicyName = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) SetConditionalAccessPolicyType(v string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies {
	s.ConditionalAccessPolicyType = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) SetConditionsConfig(v *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies {
	s.ConditionsConfig = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) SetCreateTime(v int64) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) SetDecisionConfig(v *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies {
	s.DecisionConfig = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) SetDecisionType(v string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies {
	s.DecisionType = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) SetDescription(v string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies {
	s.Description = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) SetEvaluateAt(v string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies {
	s.EvaluateAt = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) SetInstanceId(v string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies {
	s.InstanceId = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) SetLastUpdatedTime(v int64) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies {
	s.LastUpdatedTime = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) SetPriority(v int32) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies {
	s.Priority = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) SetStatus(v string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies {
	s.Status = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPolicies) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig struct {
	// Target applications of the conditional access policy
	Applications *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	// Network zones for conditional access policies
	NetworkZones *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones `json:"NetworkZones,omitempty" xml:"NetworkZones,omitempty" type:"Struct"`
	// Target users of the conditional access policy
	Users *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig) GetApplications() *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigApplications {
	return s.Applications
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig) GetNetworkZones() *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones {
	return s.NetworkZones
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig) GetUsers() *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	return s.Users
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig) SetApplications(v *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigApplications) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig {
	s.Applications = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig) SetNetworkZones(v *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig {
	s.NetworkZones = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig) SetUsers(v *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig {
	s.Users = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfig) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigApplications struct {
	// Excluded applications
	ExcludeApplications []*string `json:"ExcludeApplications,omitempty" xml:"ExcludeApplications,omitempty" type:"Repeated"`
	// Selected applications
	IncludeApplications []*string `json:"IncludeApplications,omitempty" xml:"IncludeApplications,omitempty" type:"Repeated"`
}

func (s ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigApplications) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigApplications) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigApplications) GetExcludeApplications() []*string {
	return s.ExcludeApplications
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigApplications) GetIncludeApplications() []*string {
	return s.IncludeApplications
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigApplications) SetExcludeApplications(v []*string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigApplications {
	s.ExcludeApplications = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigApplications) SetIncludeApplications(v []*string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigApplications {
	s.IncludeApplications = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigApplications) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones struct {
	// Excluded network zones
	ExcludeNetworkZones []*string `json:"ExcludeNetworkZones,omitempty" xml:"ExcludeNetworkZones,omitempty" type:"Repeated"`
	// Included network ranges
	IncludeNetworkZones []*string `json:"IncludeNetworkZones,omitempty" xml:"IncludeNetworkZones,omitempty" type:"Repeated"`
}

func (s ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) GetExcludeNetworkZones() []*string {
	return s.ExcludeNetworkZones
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) GetIncludeNetworkZones() []*string {
	return s.IncludeNetworkZones
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) SetExcludeNetworkZones(v []*string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones {
	s.ExcludeNetworkZones = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) SetIncludeNetworkZones(v []*string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones {
	s.IncludeNetworkZones = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigNetworkZones) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers struct {
	// Excluded user groups
	ExcludeGroups []*string `json:"ExcludeGroups,omitempty" xml:"ExcludeGroups,omitempty" type:"Repeated"`
	// Excluded organizations
	ExcludeOrganizationalUnits []*string `json:"ExcludeOrganizationalUnits,omitempty" xml:"ExcludeOrganizationalUnits,omitempty" type:"Repeated"`
	// Excluded users
	ExcludeUsers []*string `json:"ExcludeUsers,omitempty" xml:"ExcludeUsers,omitempty" type:"Repeated"`
	// Included user groups
	IncludeGroups []*string `json:"IncludeGroups,omitempty" xml:"IncludeGroups,omitempty" type:"Repeated"`
	// Included organizations
	IncludeOrganizationalUnits []*string `json:"IncludeOrganizationalUnits,omitempty" xml:"IncludeOrganizationalUnits,omitempty" type:"Repeated"`
	// Selected users
	IncludeUsers []*string `json:"IncludeUsers,omitempty" xml:"IncludeUsers,omitempty" type:"Repeated"`
}

func (s ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetExcludeGroups() []*string {
	return s.ExcludeGroups
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetExcludeOrganizationalUnits() []*string {
	return s.ExcludeOrganizationalUnits
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetExcludeUsers() []*string {
	return s.ExcludeUsers
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetIncludeGroups() []*string {
	return s.IncludeGroups
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetIncludeOrganizationalUnits() []*string {
	return s.IncludeOrganizationalUnits
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) GetIncludeUsers() []*string {
	return s.IncludeUsers
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetExcludeGroups(v []*string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.ExcludeGroups = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetExcludeOrganizationalUnits(v []*string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.ExcludeOrganizationalUnits = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetExcludeUsers(v []*string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.ExcludeUsers = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetIncludeGroups(v []*string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.IncludeGroups = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetIncludeOrganizationalUnits(v []*string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.IncludeOrganizationalUnits = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) SetIncludeUsers(v []*string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers {
	s.IncludeUsers = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesConditionsConfigUsers) Validate() error {
	return dara.Validate(s)
}

type ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig struct {
	// Whether to enable session reuse for secondary authentication
	//
	// example:
	//
	// disabled
	ActiveSessionReuseStatus *string `json:"ActiveSessionReuseStatus,omitempty" xml:"ActiveSessionReuseStatus,omitempty"`
	// Decision action of the conditional access policy:
	//
	// deny  Deny
	//
	// allow Allow
	//
	// example:
	//
	// deny
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// Re-authentication interval for the conditional access policy (in seconds) 300-86400
	//
	// example:
	//
	// 300
	MfaAuthenticationIntervalSeconds *int64 `json:"MfaAuthenticationIntervalSeconds,omitempty" xml:"MfaAuthenticationIntervalSeconds,omitempty"`
	// MFA types allowed by the conditional access policy
	MfaAuthenticationMethods []*string `json:"MfaAuthenticationMethods,omitempty" xml:"MfaAuthenticationMethods,omitempty" type:"Repeated"`
	// Conditional Access Policy Mfa Type
	//
	// example:
	//
	// directly_access
	MfaType *string `json:"MfaType,omitempty" xml:"MfaType,omitempty"`
}

func (s ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig) GetActiveSessionReuseStatus() *string {
	return s.ActiveSessionReuseStatus
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig) GetEffect() *string {
	return s.Effect
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig) GetMfaAuthenticationIntervalSeconds() *int64 {
	return s.MfaAuthenticationIntervalSeconds
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig) GetMfaAuthenticationMethods() []*string {
	return s.MfaAuthenticationMethods
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig) GetMfaType() *string {
	return s.MfaType
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig) SetActiveSessionReuseStatus(v string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.ActiveSessionReuseStatus = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig) SetEffect(v string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.Effect = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig) SetMfaAuthenticationIntervalSeconds(v int64) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.MfaAuthenticationIntervalSeconds = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig) SetMfaAuthenticationMethods(v []*string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.MfaAuthenticationMethods = v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig) SetMfaType(v string) *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig {
	s.MfaType = &v
	return s
}

func (s *ListConditionalAccessPoliciesResponseBodyConditionalAccessPoliciesDecisionConfig) Validate() error {
	return dara.Validate(s)
}
