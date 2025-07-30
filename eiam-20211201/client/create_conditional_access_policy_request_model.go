// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConditionalAccessPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateConditionalAccessPolicyRequest
	GetClientToken() *string
	SetConditionalAccessPolicyName(v string) *CreateConditionalAccessPolicyRequest
	GetConditionalAccessPolicyName() *string
	SetConditionalAccessPolicyType(v string) *CreateConditionalAccessPolicyRequest
	GetConditionalAccessPolicyType() *string
	SetConditionsConfig(v *CreateConditionalAccessPolicyRequestConditionsConfig) *CreateConditionalAccessPolicyRequest
	GetConditionsConfig() *CreateConditionalAccessPolicyRequestConditionsConfig
	SetDecisionConfig(v *CreateConditionalAccessPolicyRequestDecisionConfig) *CreateConditionalAccessPolicyRequest
	GetDecisionConfig() *CreateConditionalAccessPolicyRequestDecisionConfig
	SetDecisionType(v string) *CreateConditionalAccessPolicyRequest
	GetDecisionType() *string
	SetDescription(v string) *CreateConditionalAccessPolicyRequest
	GetDescription() *string
	SetEvaluateAt(v string) *CreateConditionalAccessPolicyRequest
	GetEvaluateAt() *string
	SetInstanceId(v string) *CreateConditionalAccessPolicyRequest
	GetInstanceId() *string
	SetPriority(v int32) *CreateConditionalAccessPolicyRequest
	GetPriority() *int32
}

type CreateConditionalAccessPolicyRequest struct {
	// Idempotent token.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Conditional access policy name
	//
	// This parameter is required.
	//
	// example:
	//
	// My Conditional Access Policy
	ConditionalAccessPolicyName *string `json:"ConditionalAccessPolicyName,omitempty" xml:"ConditionalAccessPolicyName,omitempty"`
	// Type of the conditional access policy, with the following options:
	//
	// arn:alibaba:idaas:authn:access:policy:system: System policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// arn:alibaba:idaas:authn:access:policy:system
	ConditionalAccessPolicyType *string `json:"ConditionalAccessPolicyType,omitempty" xml:"ConditionalAccessPolicyType,omitempty"`
	// Condition content configuration for the conditional access policy
	ConditionsConfig *CreateConditionalAccessPolicyRequestConditionsConfig `json:"ConditionsConfig,omitempty" xml:"ConditionsConfig,omitempty" type:"Struct"`
	// Action configuration for the conditional access policy
	DecisionConfig *CreateConditionalAccessPolicyRequestDecisionConfig `json:"DecisionConfig,omitempty" xml:"DecisionConfig,omitempty" type:"Struct"`
	// Execution type of the conditional access policy, with the following options:
	//
	// enforcement: Enforce the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// enforcement
	DecisionType *string `json:"DecisionType,omitempty" xml:"DecisionType,omitempty"`
	// Description of the conditional access policy
	//
	// example:
	//
	// Test Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Execution point of the conditional access policy, with the following options:
	//
	// - arn:alibaba:idaas:authn:access:rule:eval_at:after_step1: Allow.
	//
	// This parameter is required.
	//
	// example:
	//
	// arn:alibaba:idaas:authn:access:rule:eval_at:after_step1
	EvaluateAt *string `json:"EvaluateAt,omitempty" xml:"EvaluateAt,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Priority of the conditional access policy, lower values indicate higher priority
	//
	// Minimum value: 1
	//
	// Maximum value: 100
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s CreateConditionalAccessPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConditionalAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateConditionalAccessPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateConditionalAccessPolicyRequest) GetConditionalAccessPolicyName() *string {
	return s.ConditionalAccessPolicyName
}

func (s *CreateConditionalAccessPolicyRequest) GetConditionalAccessPolicyType() *string {
	return s.ConditionalAccessPolicyType
}

func (s *CreateConditionalAccessPolicyRequest) GetConditionsConfig() *CreateConditionalAccessPolicyRequestConditionsConfig {
	return s.ConditionsConfig
}

func (s *CreateConditionalAccessPolicyRequest) GetDecisionConfig() *CreateConditionalAccessPolicyRequestDecisionConfig {
	return s.DecisionConfig
}

func (s *CreateConditionalAccessPolicyRequest) GetDecisionType() *string {
	return s.DecisionType
}

func (s *CreateConditionalAccessPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateConditionalAccessPolicyRequest) GetEvaluateAt() *string {
	return s.EvaluateAt
}

func (s *CreateConditionalAccessPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateConditionalAccessPolicyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateConditionalAccessPolicyRequest) SetClientToken(v string) *CreateConditionalAccessPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateConditionalAccessPolicyRequest) SetConditionalAccessPolicyName(v string) *CreateConditionalAccessPolicyRequest {
	s.ConditionalAccessPolicyName = &v
	return s
}

func (s *CreateConditionalAccessPolicyRequest) SetConditionalAccessPolicyType(v string) *CreateConditionalAccessPolicyRequest {
	s.ConditionalAccessPolicyType = &v
	return s
}

func (s *CreateConditionalAccessPolicyRequest) SetConditionsConfig(v *CreateConditionalAccessPolicyRequestConditionsConfig) *CreateConditionalAccessPolicyRequest {
	s.ConditionsConfig = v
	return s
}

func (s *CreateConditionalAccessPolicyRequest) SetDecisionConfig(v *CreateConditionalAccessPolicyRequestDecisionConfig) *CreateConditionalAccessPolicyRequest {
	s.DecisionConfig = v
	return s
}

func (s *CreateConditionalAccessPolicyRequest) SetDecisionType(v string) *CreateConditionalAccessPolicyRequest {
	s.DecisionType = &v
	return s
}

func (s *CreateConditionalAccessPolicyRequest) SetDescription(v string) *CreateConditionalAccessPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateConditionalAccessPolicyRequest) SetEvaluateAt(v string) *CreateConditionalAccessPolicyRequest {
	s.EvaluateAt = &v
	return s
}

func (s *CreateConditionalAccessPolicyRequest) SetInstanceId(v string) *CreateConditionalAccessPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateConditionalAccessPolicyRequest) SetPriority(v int32) *CreateConditionalAccessPolicyRequest {
	s.Priority = &v
	return s
}

func (s *CreateConditionalAccessPolicyRequest) Validate() error {
	return dara.Validate(s)
}

type CreateConditionalAccessPolicyRequestConditionsConfig struct {
	// Target applications for the conditional access policy
	Applications *CreateConditionalAccessPolicyRequestConditionsConfigApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	// Network zones for conditional access policy
	NetworkZones *CreateConditionalAccessPolicyRequestConditionsConfigNetworkZones `json:"NetworkZones,omitempty" xml:"NetworkZones,omitempty" type:"Struct"`
	// Target users of the conditional access policy
	Users *CreateConditionalAccessPolicyRequestConditionsConfigUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s CreateConditionalAccessPolicyRequestConditionsConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateConditionalAccessPolicyRequestConditionsConfig) GoString() string {
	return s.String()
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfig) GetApplications() *CreateConditionalAccessPolicyRequestConditionsConfigApplications {
	return s.Applications
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfig) GetNetworkZones() *CreateConditionalAccessPolicyRequestConditionsConfigNetworkZones {
	return s.NetworkZones
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfig) GetUsers() *CreateConditionalAccessPolicyRequestConditionsConfigUsers {
	return s.Users
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfig) SetApplications(v *CreateConditionalAccessPolicyRequestConditionsConfigApplications) *CreateConditionalAccessPolicyRequestConditionsConfig {
	s.Applications = v
	return s
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfig) SetNetworkZones(v *CreateConditionalAccessPolicyRequestConditionsConfigNetworkZones) *CreateConditionalAccessPolicyRequestConditionsConfig {
	s.NetworkZones = v
	return s
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfig) SetUsers(v *CreateConditionalAccessPolicyRequestConditionsConfigUsers) *CreateConditionalAccessPolicyRequestConditionsConfig {
	s.Users = v
	return s
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfig) Validate() error {
	return dara.Validate(s)
}

type CreateConditionalAccessPolicyRequestConditionsConfigApplications struct {
	// Excluded applications
	ExcludeApplications []*string `json:"ExcludeApplications,omitempty" xml:"ExcludeApplications,omitempty" type:"Repeated"`
	// Included applications
	IncludeApplications []*string `json:"IncludeApplications,omitempty" xml:"IncludeApplications,omitempty" type:"Repeated"`
}

func (s CreateConditionalAccessPolicyRequestConditionsConfigApplications) String() string {
	return dara.Prettify(s)
}

func (s CreateConditionalAccessPolicyRequestConditionsConfigApplications) GoString() string {
	return s.String()
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigApplications) GetExcludeApplications() []*string {
	return s.ExcludeApplications
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigApplications) GetIncludeApplications() []*string {
	return s.IncludeApplications
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigApplications) SetExcludeApplications(v []*string) *CreateConditionalAccessPolicyRequestConditionsConfigApplications {
	s.ExcludeApplications = v
	return s
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigApplications) SetIncludeApplications(v []*string) *CreateConditionalAccessPolicyRequestConditionsConfigApplications {
	s.IncludeApplications = v
	return s
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigApplications) Validate() error {
	return dara.Validate(s)
}

type CreateConditionalAccessPolicyRequestConditionsConfigNetworkZones struct {
	// Excluded network zones
	ExcludeNetworkZones []*string `json:"ExcludeNetworkZones,omitempty" xml:"ExcludeNetworkZones,omitempty" type:"Repeated"`
	// Included network zones
	IncludeNetworkZones []*string `json:"IncludeNetworkZones,omitempty" xml:"IncludeNetworkZones,omitempty" type:"Repeated"`
}

func (s CreateConditionalAccessPolicyRequestConditionsConfigNetworkZones) String() string {
	return dara.Prettify(s)
}

func (s CreateConditionalAccessPolicyRequestConditionsConfigNetworkZones) GoString() string {
	return s.String()
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigNetworkZones) GetExcludeNetworkZones() []*string {
	return s.ExcludeNetworkZones
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigNetworkZones) GetIncludeNetworkZones() []*string {
	return s.IncludeNetworkZones
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigNetworkZones) SetExcludeNetworkZones(v []*string) *CreateConditionalAccessPolicyRequestConditionsConfigNetworkZones {
	s.ExcludeNetworkZones = v
	return s
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigNetworkZones) SetIncludeNetworkZones(v []*string) *CreateConditionalAccessPolicyRequestConditionsConfigNetworkZones {
	s.IncludeNetworkZones = v
	return s
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigNetworkZones) Validate() error {
	return dara.Validate(s)
}

type CreateConditionalAccessPolicyRequestConditionsConfigUsers struct {
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
	// Selected user
	IncludeUsers []*string `json:"IncludeUsers,omitempty" xml:"IncludeUsers,omitempty" type:"Repeated"`
}

func (s CreateConditionalAccessPolicyRequestConditionsConfigUsers) String() string {
	return dara.Prettify(s)
}

func (s CreateConditionalAccessPolicyRequestConditionsConfigUsers) GoString() string {
	return s.String()
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigUsers) GetExcludeGroups() []*string {
	return s.ExcludeGroups
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigUsers) GetExcludeOrganizationalUnits() []*string {
	return s.ExcludeOrganizationalUnits
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigUsers) GetExcludeUsers() []*string {
	return s.ExcludeUsers
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigUsers) GetIncludeGroups() []*string {
	return s.IncludeGroups
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigUsers) GetIncludeOrganizationalUnits() []*string {
	return s.IncludeOrganizationalUnits
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigUsers) GetIncludeUsers() []*string {
	return s.IncludeUsers
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigUsers) SetExcludeGroups(v []*string) *CreateConditionalAccessPolicyRequestConditionsConfigUsers {
	s.ExcludeGroups = v
	return s
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigUsers) SetExcludeOrganizationalUnits(v []*string) *CreateConditionalAccessPolicyRequestConditionsConfigUsers {
	s.ExcludeOrganizationalUnits = v
	return s
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigUsers) SetExcludeUsers(v []*string) *CreateConditionalAccessPolicyRequestConditionsConfigUsers {
	s.ExcludeUsers = v
	return s
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigUsers) SetIncludeGroups(v []*string) *CreateConditionalAccessPolicyRequestConditionsConfigUsers {
	s.IncludeGroups = v
	return s
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigUsers) SetIncludeOrganizationalUnits(v []*string) *CreateConditionalAccessPolicyRequestConditionsConfigUsers {
	s.IncludeOrganizationalUnits = v
	return s
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigUsers) SetIncludeUsers(v []*string) *CreateConditionalAccessPolicyRequestConditionsConfigUsers {
	s.IncludeUsers = v
	return s
}

func (s *CreateConditionalAccessPolicyRequestConditionsConfigUsers) Validate() error {
	return dara.Validate(s)
}

type CreateConditionalAccessPolicyRequestDecisionConfig struct {
	// Whether to enable session reuse
	//
	// example:
	//
	// enabled
	ActiveSessionReuseStatus *string `json:"ActiveSessionReuseStatus,omitempty" xml:"ActiveSessionReuseStatus,omitempty"`
	// Decision action for the conditional access policy, with the following options:
	//
	// - allow: Allow.
	//
	// - deny: Deny.
	//
	// example:
	//
	// allow or deny
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// Re-authentication interval (in seconds) for the conditional access policy
	//
	// - Maximum MFA re-authentication interval: 86400
	//
	// - Minimum MFA re-authentication interval: 300
	//
	// example:
	//
	// 500
	MfaAuthenticationIntervalSeconds *int64 `json:"MfaAuthenticationIntervalSeconds,omitempty" xml:"MfaAuthenticationIntervalSeconds,omitempty"`
	// Allowed MFA types for the conditional access policy, with the following options:
	//
	// - ia_otp_sms: SMS verification code
	//
	// - ia_otp_email: Email verification code
	//
	// - ia_totp: OTP dynamic password
	//
	// - ia_webauthn: WebAuthn
	MfaAuthenticationMethods []*string `json:"MfaAuthenticationMethods,omitempty" xml:"MfaAuthenticationMethods,omitempty" type:"Repeated"`
	// MFA type for the conditional access policy, with the following options:
	//
	// - directly_access: Direct access
	//
	// - mfa_required: MFA required
	//
	// example:
	//
	// directly_access
	MfaType *string `json:"MfaType,omitempty" xml:"MfaType,omitempty"`
}

func (s CreateConditionalAccessPolicyRequestDecisionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateConditionalAccessPolicyRequestDecisionConfig) GoString() string {
	return s.String()
}

func (s *CreateConditionalAccessPolicyRequestDecisionConfig) GetActiveSessionReuseStatus() *string {
	return s.ActiveSessionReuseStatus
}

func (s *CreateConditionalAccessPolicyRequestDecisionConfig) GetEffect() *string {
	return s.Effect
}

func (s *CreateConditionalAccessPolicyRequestDecisionConfig) GetMfaAuthenticationIntervalSeconds() *int64 {
	return s.MfaAuthenticationIntervalSeconds
}

func (s *CreateConditionalAccessPolicyRequestDecisionConfig) GetMfaAuthenticationMethods() []*string {
	return s.MfaAuthenticationMethods
}

func (s *CreateConditionalAccessPolicyRequestDecisionConfig) GetMfaType() *string {
	return s.MfaType
}

func (s *CreateConditionalAccessPolicyRequestDecisionConfig) SetActiveSessionReuseStatus(v string) *CreateConditionalAccessPolicyRequestDecisionConfig {
	s.ActiveSessionReuseStatus = &v
	return s
}

func (s *CreateConditionalAccessPolicyRequestDecisionConfig) SetEffect(v string) *CreateConditionalAccessPolicyRequestDecisionConfig {
	s.Effect = &v
	return s
}

func (s *CreateConditionalAccessPolicyRequestDecisionConfig) SetMfaAuthenticationIntervalSeconds(v int64) *CreateConditionalAccessPolicyRequestDecisionConfig {
	s.MfaAuthenticationIntervalSeconds = &v
	return s
}

func (s *CreateConditionalAccessPolicyRequestDecisionConfig) SetMfaAuthenticationMethods(v []*string) *CreateConditionalAccessPolicyRequestDecisionConfig {
	s.MfaAuthenticationMethods = v
	return s
}

func (s *CreateConditionalAccessPolicyRequestDecisionConfig) SetMfaType(v string) *CreateConditionalAccessPolicyRequestDecisionConfig {
	s.MfaType = &v
	return s
}

func (s *CreateConditionalAccessPolicyRequestDecisionConfig) Validate() error {
	return dara.Validate(s)
}
