// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConditionalAccessPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateConditionalAccessPolicyRequest
	GetClientToken() *string
	SetConditionalAccessPolicyId(v string) *UpdateConditionalAccessPolicyRequest
	GetConditionalAccessPolicyId() *string
	SetConditionalAccessPolicyName(v string) *UpdateConditionalAccessPolicyRequest
	GetConditionalAccessPolicyName() *string
	SetConditionsConfig(v *UpdateConditionalAccessPolicyRequestConditionsConfig) *UpdateConditionalAccessPolicyRequest
	GetConditionsConfig() *UpdateConditionalAccessPolicyRequestConditionsConfig
	SetDecisionConfig(v *UpdateConditionalAccessPolicyRequestDecisionConfig) *UpdateConditionalAccessPolicyRequest
	GetDecisionConfig() *UpdateConditionalAccessPolicyRequestDecisionConfig
	SetDecisionType(v string) *UpdateConditionalAccessPolicyRequest
	GetDecisionType() *string
	SetInstanceId(v string) *UpdateConditionalAccessPolicyRequest
	GetInstanceId() *string
	SetPriority(v int32) *UpdateConditionalAccessPolicyRequest
	GetPriority() *int32
}

type UpdateConditionalAccessPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// client-examplexxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Conditional Access Policy ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cap_11111
	ConditionalAccessPolicyId *string `json:"ConditionalAccessPolicyId,omitempty" xml:"ConditionalAccessPolicyId,omitempty"`
	// Conditional Access Policy Name
	//
	// This parameter is required.
	//
	// example:
	//
	// My Policy
	ConditionalAccessPolicyName *string `json:"ConditionalAccessPolicyName,omitempty" xml:"ConditionalAccessPolicyName,omitempty"`
	// Conditional Access Policy Condition Content Configuration
	ConditionsConfig *UpdateConditionalAccessPolicyRequestConditionsConfig `json:"ConditionsConfig,omitempty" xml:"ConditionsConfig,omitempty" type:"Struct"`
	// Conditional Access Policy Action Configuration
	DecisionConfig *UpdateConditionalAccessPolicyRequestDecisionConfig `json:"DecisionConfig,omitempty" xml:"DecisionConfig,omitempty" type:"Struct"`
	// Conditional Access Policy Execution Type
	//
	// This parameter is required.
	//
	// example:
	//
	// enforcement
	DecisionType *string `json:"DecisionType,omitempty" xml:"DecisionType,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Priority of the conditional access policy
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s UpdateConditionalAccessPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConditionalAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateConditionalAccessPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateConditionalAccessPolicyRequest) GetConditionalAccessPolicyId() *string {
	return s.ConditionalAccessPolicyId
}

func (s *UpdateConditionalAccessPolicyRequest) GetConditionalAccessPolicyName() *string {
	return s.ConditionalAccessPolicyName
}

func (s *UpdateConditionalAccessPolicyRequest) GetConditionsConfig() *UpdateConditionalAccessPolicyRequestConditionsConfig {
	return s.ConditionsConfig
}

func (s *UpdateConditionalAccessPolicyRequest) GetDecisionConfig() *UpdateConditionalAccessPolicyRequestDecisionConfig {
	return s.DecisionConfig
}

func (s *UpdateConditionalAccessPolicyRequest) GetDecisionType() *string {
	return s.DecisionType
}

func (s *UpdateConditionalAccessPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateConditionalAccessPolicyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateConditionalAccessPolicyRequest) SetClientToken(v string) *UpdateConditionalAccessPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateConditionalAccessPolicyRequest) SetConditionalAccessPolicyId(v string) *UpdateConditionalAccessPolicyRequest {
	s.ConditionalAccessPolicyId = &v
	return s
}

func (s *UpdateConditionalAccessPolicyRequest) SetConditionalAccessPolicyName(v string) *UpdateConditionalAccessPolicyRequest {
	s.ConditionalAccessPolicyName = &v
	return s
}

func (s *UpdateConditionalAccessPolicyRequest) SetConditionsConfig(v *UpdateConditionalAccessPolicyRequestConditionsConfig) *UpdateConditionalAccessPolicyRequest {
	s.ConditionsConfig = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequest) SetDecisionConfig(v *UpdateConditionalAccessPolicyRequestDecisionConfig) *UpdateConditionalAccessPolicyRequest {
	s.DecisionConfig = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequest) SetDecisionType(v string) *UpdateConditionalAccessPolicyRequest {
	s.DecisionType = &v
	return s
}

func (s *UpdateConditionalAccessPolicyRequest) SetInstanceId(v string) *UpdateConditionalAccessPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateConditionalAccessPolicyRequest) SetPriority(v int32) *UpdateConditionalAccessPolicyRequest {
	s.Priority = &v
	return s
}

func (s *UpdateConditionalAccessPolicyRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateConditionalAccessPolicyRequestConditionsConfig struct {
	// Target Applications for the Conditional Access Policy
	Applications *UpdateConditionalAccessPolicyRequestConditionsConfigApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	// Network zones for conditional access policy
	NetworkZones *UpdateConditionalAccessPolicyRequestConditionsConfigNetworkZones `json:"NetworkZones,omitempty" xml:"NetworkZones,omitempty" type:"Struct"`
	// Target Users for the Conditional Access Policy
	Users *UpdateConditionalAccessPolicyRequestConditionsConfigUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s UpdateConditionalAccessPolicyRequestConditionsConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateConditionalAccessPolicyRequestConditionsConfig) GoString() string {
	return s.String()
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfig) GetApplications() *UpdateConditionalAccessPolicyRequestConditionsConfigApplications {
	return s.Applications
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfig) GetNetworkZones() *UpdateConditionalAccessPolicyRequestConditionsConfigNetworkZones {
	return s.NetworkZones
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfig) GetUsers() *UpdateConditionalAccessPolicyRequestConditionsConfigUsers {
	return s.Users
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfig) SetApplications(v *UpdateConditionalAccessPolicyRequestConditionsConfigApplications) *UpdateConditionalAccessPolicyRequestConditionsConfig {
	s.Applications = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfig) SetNetworkZones(v *UpdateConditionalAccessPolicyRequestConditionsConfigNetworkZones) *UpdateConditionalAccessPolicyRequestConditionsConfig {
	s.NetworkZones = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfig) SetUsers(v *UpdateConditionalAccessPolicyRequestConditionsConfigUsers) *UpdateConditionalAccessPolicyRequestConditionsConfig {
	s.Users = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateConditionalAccessPolicyRequestConditionsConfigApplications struct {
	// Excluded Applications
	ExcludeApplications []*string `json:"ExcludeApplications,omitempty" xml:"ExcludeApplications,omitempty" type:"Repeated"`
	// Included Applications
	IncludeApplications []*string `json:"IncludeApplications,omitempty" xml:"IncludeApplications,omitempty" type:"Repeated"`
}

func (s UpdateConditionalAccessPolicyRequestConditionsConfigApplications) String() string {
	return dara.Prettify(s)
}

func (s UpdateConditionalAccessPolicyRequestConditionsConfigApplications) GoString() string {
	return s.String()
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigApplications) GetExcludeApplications() []*string {
	return s.ExcludeApplications
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigApplications) GetIncludeApplications() []*string {
	return s.IncludeApplications
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigApplications) SetExcludeApplications(v []*string) *UpdateConditionalAccessPolicyRequestConditionsConfigApplications {
	s.ExcludeApplications = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigApplications) SetIncludeApplications(v []*string) *UpdateConditionalAccessPolicyRequestConditionsConfigApplications {
	s.IncludeApplications = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigApplications) Validate() error {
	return dara.Validate(s)
}

type UpdateConditionalAccessPolicyRequestConditionsConfigNetworkZones struct {
	// Excluded network zones
	ExcludeNetworkZones []*string `json:"ExcludeNetworkZones,omitempty" xml:"ExcludeNetworkZones,omitempty" type:"Repeated"`
	// Included network zones
	IncludeNetworkZones []*string `json:"IncludeNetworkZones,omitempty" xml:"IncludeNetworkZones,omitempty" type:"Repeated"`
}

func (s UpdateConditionalAccessPolicyRequestConditionsConfigNetworkZones) String() string {
	return dara.Prettify(s)
}

func (s UpdateConditionalAccessPolicyRequestConditionsConfigNetworkZones) GoString() string {
	return s.String()
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigNetworkZones) GetExcludeNetworkZones() []*string {
	return s.ExcludeNetworkZones
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigNetworkZones) GetIncludeNetworkZones() []*string {
	return s.IncludeNetworkZones
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigNetworkZones) SetExcludeNetworkZones(v []*string) *UpdateConditionalAccessPolicyRequestConditionsConfigNetworkZones {
	s.ExcludeNetworkZones = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigNetworkZones) SetIncludeNetworkZones(v []*string) *UpdateConditionalAccessPolicyRequestConditionsConfigNetworkZones {
	s.IncludeNetworkZones = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigNetworkZones) Validate() error {
	return dara.Validate(s)
}

type UpdateConditionalAccessPolicyRequestConditionsConfigUsers struct {
	// Excluded user groups
	ExcludeGroups []*string `json:"ExcludeGroups,omitempty" xml:"ExcludeGroups,omitempty" type:"Repeated"`
	// Excluded organizations
	ExcludeOrganizationalUnits []*string `json:"ExcludeOrganizationalUnits,omitempty" xml:"ExcludeOrganizationalUnits,omitempty" type:"Repeated"`
	// Excluded Users
	ExcludeUsers []*string `json:"ExcludeUsers,omitempty" xml:"ExcludeUsers,omitempty" type:"Repeated"`
	// Included User Groups
	IncludeGroups []*string `json:"IncludeGroups,omitempty" xml:"IncludeGroups,omitempty" type:"Repeated"`
	// Included organizations
	IncludeOrganizationalUnits []*string `json:"IncludeOrganizationalUnits,omitempty" xml:"IncludeOrganizationalUnits,omitempty" type:"Repeated"`
	// Included Users
	IncludeUsers []*string `json:"IncludeUsers,omitempty" xml:"IncludeUsers,omitempty" type:"Repeated"`
}

func (s UpdateConditionalAccessPolicyRequestConditionsConfigUsers) String() string {
	return dara.Prettify(s)
}

func (s UpdateConditionalAccessPolicyRequestConditionsConfigUsers) GoString() string {
	return s.String()
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigUsers) GetExcludeGroups() []*string {
	return s.ExcludeGroups
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigUsers) GetExcludeOrganizationalUnits() []*string {
	return s.ExcludeOrganizationalUnits
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigUsers) GetExcludeUsers() []*string {
	return s.ExcludeUsers
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigUsers) GetIncludeGroups() []*string {
	return s.IncludeGroups
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigUsers) GetIncludeOrganizationalUnits() []*string {
	return s.IncludeOrganizationalUnits
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigUsers) GetIncludeUsers() []*string {
	return s.IncludeUsers
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigUsers) SetExcludeGroups(v []*string) *UpdateConditionalAccessPolicyRequestConditionsConfigUsers {
	s.ExcludeGroups = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigUsers) SetExcludeOrganizationalUnits(v []*string) *UpdateConditionalAccessPolicyRequestConditionsConfigUsers {
	s.ExcludeOrganizationalUnits = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigUsers) SetExcludeUsers(v []*string) *UpdateConditionalAccessPolicyRequestConditionsConfigUsers {
	s.ExcludeUsers = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigUsers) SetIncludeGroups(v []*string) *UpdateConditionalAccessPolicyRequestConditionsConfigUsers {
	s.IncludeGroups = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigUsers) SetIncludeOrganizationalUnits(v []*string) *UpdateConditionalAccessPolicyRequestConditionsConfigUsers {
	s.IncludeOrganizationalUnits = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigUsers) SetIncludeUsers(v []*string) *UpdateConditionalAccessPolicyRequestConditionsConfigUsers {
	s.IncludeUsers = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestConditionsConfigUsers) Validate() error {
	return dara.Validate(s)
}

type UpdateConditionalAccessPolicyRequestDecisionConfig struct {
	// Whether to enable session reuse
	//
	// example:
	//
	// enabled
	ActiveSessionReuseStatus *string `json:"ActiveSessionReuseStatus,omitempty" xml:"ActiveSessionReuseStatus,omitempty"`
	// Conditional Access Policy Decision Action
	//
	// example:
	//
	// allow
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// Conditional Access Policy Re-authentication Interval (seconds)
	//
	// example:
	//
	// 300
	MfaAuthenticationIntervalSeconds *int64 `json:"MfaAuthenticationIntervalSeconds,omitempty" xml:"MfaAuthenticationIntervalSeconds,omitempty"`
	// Allowed MFA types for the Conditional Access Policy
	MfaAuthenticationMethods []*string `json:"MfaAuthenticationMethods,omitempty" xml:"MfaAuthenticationMethods,omitempty" type:"Repeated"`
	// Conditional Access Policy MFA Type
	//
	// example:
	//
	// directly_access
	MfaType *string `json:"MfaType,omitempty" xml:"MfaType,omitempty"`
}

func (s UpdateConditionalAccessPolicyRequestDecisionConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateConditionalAccessPolicyRequestDecisionConfig) GoString() string {
	return s.String()
}

func (s *UpdateConditionalAccessPolicyRequestDecisionConfig) GetActiveSessionReuseStatus() *string {
	return s.ActiveSessionReuseStatus
}

func (s *UpdateConditionalAccessPolicyRequestDecisionConfig) GetEffect() *string {
	return s.Effect
}

func (s *UpdateConditionalAccessPolicyRequestDecisionConfig) GetMfaAuthenticationIntervalSeconds() *int64 {
	return s.MfaAuthenticationIntervalSeconds
}

func (s *UpdateConditionalAccessPolicyRequestDecisionConfig) GetMfaAuthenticationMethods() []*string {
	return s.MfaAuthenticationMethods
}

func (s *UpdateConditionalAccessPolicyRequestDecisionConfig) GetMfaType() *string {
	return s.MfaType
}

func (s *UpdateConditionalAccessPolicyRequestDecisionConfig) SetActiveSessionReuseStatus(v string) *UpdateConditionalAccessPolicyRequestDecisionConfig {
	s.ActiveSessionReuseStatus = &v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestDecisionConfig) SetEffect(v string) *UpdateConditionalAccessPolicyRequestDecisionConfig {
	s.Effect = &v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestDecisionConfig) SetMfaAuthenticationIntervalSeconds(v int64) *UpdateConditionalAccessPolicyRequestDecisionConfig {
	s.MfaAuthenticationIntervalSeconds = &v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestDecisionConfig) SetMfaAuthenticationMethods(v []*string) *UpdateConditionalAccessPolicyRequestDecisionConfig {
	s.MfaAuthenticationMethods = v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestDecisionConfig) SetMfaType(v string) *UpdateConditionalAccessPolicyRequestDecisionConfig {
	s.MfaType = &v
	return s
}

func (s *UpdateConditionalAccessPolicyRequestDecisionConfig) Validate() error {
	return dara.Validate(s)
}
