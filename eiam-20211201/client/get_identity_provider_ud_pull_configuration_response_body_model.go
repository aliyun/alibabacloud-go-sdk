// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderUdPullConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetIdentityProviderUdPullConfigurationResponseBody
	GetRequestId() *string
	SetUdPullConfiguration(v *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) *GetIdentityProviderUdPullConfigurationResponseBody
	GetUdPullConfiguration() *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration
}

type GetIdentityProviderUdPullConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Inbound Synchronization Configuration Information
	UdPullConfiguration *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration `json:"UdPullConfiguration,omitempty" xml:"UdPullConfiguration,omitempty" type:"Struct"`
}

func (s GetIdentityProviderUdPullConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderUdPullConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderUdPullConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentityProviderUdPullConfigurationResponseBody) GetUdPullConfiguration() *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration {
	return s.UdPullConfiguration
}

func (s *GetIdentityProviderUdPullConfigurationResponseBody) SetRequestId(v string) *GetIdentityProviderUdPullConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBody) SetUdPullConfiguration(v *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) *GetIdentityProviderUdPullConfigurationResponseBody {
	s.UdPullConfiguration = v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBody) Validate() error {
	if s.UdPullConfiguration != nil {
		if err := s.UdPullConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration struct {
	// Group Synchronization Status
	//
	// Possible values:
	//
	// Disabled: disabled
	//
	// Enabled: enabled
	//
	// example:
	//
	// enabled
	GroupSyncStatus *string `json:"GroupSyncStatus,omitempty" xml:"GroupSyncStatus,omitempty"`
	// Identity provider ID
	//
	// example:
	//
	// idp_my664lwkhpicbyzirog3xxxxx
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// Incremental Callback Status: Whether to process incremental callback data from the IdP
	//
	// example:
	//
	// enabled
	IncrementalCallbackStatus *string `json:"IncrementalCallbackStatus,omitempty" xml:"IncrementalCallbackStatus,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// LDAP Synchronization Side Related Configuration Information
	LdapUdPullConfig *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig `json:"LdapUdPullConfig,omitempty" xml:"LdapUdPullConfig,omitempty" type:"Struct"`
	// Scheduled sync configuration
	PeriodicSyncConfig *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig `json:"PeriodicSyncConfig,omitempty" xml:"PeriodicSyncConfig,omitempty" type:"Struct"`
	// Scheduled Validation Status: Whether to periodically validate data discrepancies between IDaaS and the Identity Provider. Possible values:
	//
	// Disabled: disabled
	//
	// Enabled: enabled
	//
	// example:
	//
	// enabled
	PeriodicSyncStatus *string `json:"PeriodicSyncStatus,omitempty" xml:"PeriodicSyncStatus,omitempty"`
	// Inbound Synchronization Protection Rule Configuration
	PullProtectedRule *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule `json:"PullProtectedRule,omitempty" xml:"PullProtectedRule,omitempty" type:"Struct"`
	// Synchronization Scope Configuration Information
	UdSyncScopeConfig *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationUdSyncScopeConfig `json:"UdSyncScopeConfig,omitempty" xml:"UdSyncScopeConfig,omitempty" type:"Struct"`
}

func (s GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) GetGroupSyncStatus() *string {
	return s.GroupSyncStatus
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) GetIncrementalCallbackStatus() *string {
	return s.IncrementalCallbackStatus
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) GetLdapUdPullConfig() *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig {
	return s.LdapUdPullConfig
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) GetPeriodicSyncConfig() *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig {
	return s.PeriodicSyncConfig
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) GetPeriodicSyncStatus() *string {
	return s.PeriodicSyncStatus
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) GetPullProtectedRule() *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule {
	return s.PullProtectedRule
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) GetUdSyncScopeConfig() *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationUdSyncScopeConfig {
	return s.UdSyncScopeConfig
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) SetGroupSyncStatus(v string) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration {
	s.GroupSyncStatus = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) SetIdentityProviderId(v string) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration {
	s.IdentityProviderId = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) SetIncrementalCallbackStatus(v string) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration {
	s.IncrementalCallbackStatus = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) SetInstanceId(v string) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration {
	s.InstanceId = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) SetLdapUdPullConfig(v *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration {
	s.LdapUdPullConfig = v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) SetPeriodicSyncConfig(v *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration {
	s.PeriodicSyncConfig = v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) SetPeriodicSyncStatus(v string) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration {
	s.PeriodicSyncStatus = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) SetPullProtectedRule(v *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration {
	s.PullProtectedRule = v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) SetUdSyncScopeConfig(v *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationUdSyncScopeConfig) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration {
	s.UdSyncScopeConfig = v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfiguration) Validate() error {
	if s.LdapUdPullConfig != nil {
		if err := s.LdapUdPullConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PeriodicSyncConfig != nil {
		if err := s.PeriodicSyncConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PullProtectedRule != nil {
		if err := s.PullProtectedRule.Validate(); err != nil {
			return err
		}
	}
	if s.UdSyncScopeConfig != nil {
		if err := s.UdSyncScopeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig struct {
	// Group Member Identifier
	//
	// example:
	//
	// group
	GroupMemberAttributeName *string `json:"GroupMemberAttributeName,omitempty" xml:"GroupMemberAttributeName,omitempty"`
	// Group ObjectClass
	//
	// example:
	//
	// member
	GroupObjectClass *string `json:"GroupObjectClass,omitempty" xml:"GroupObjectClass,omitempty"`
	// Group Custom Filter
	//
	// example:
	//
	// (|(cn=test)(group=test@test.com))
	GroupObjectClassCustomFilter *string `json:"GroupObjectClassCustomFilter,omitempty" xml:"GroupObjectClassCustomFilter,omitempty"`
	// Organization ObjectClass
	//
	// example:
	//
	// ou,top
	OrganizationUnitObjectClass *string `json:"OrganizationUnitObjectClass,omitempty" xml:"OrganizationUnitObjectClass,omitempty"`
	// User ObjectClass
	//
	// example:
	//
	// ou,top
	UserObjectClass *string `json:"UserObjectClass,omitempty" xml:"UserObjectClass,omitempty"`
	// User ObjectClass Custom Filter
	//
	// example:
	//
	// (|(cn=test)(mail=test@test.com))
	UserObjectClassCustomFilter *string `json:"UserObjectClassCustomFilter,omitempty" xml:"UserObjectClassCustomFilter,omitempty"`
}

func (s GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) GetGroupMemberAttributeName() *string {
	return s.GroupMemberAttributeName
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) GetGroupObjectClass() *string {
	return s.GroupObjectClass
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) GetGroupObjectClassCustomFilter() *string {
	return s.GroupObjectClassCustomFilter
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) GetOrganizationUnitObjectClass() *string {
	return s.OrganizationUnitObjectClass
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) GetUserObjectClass() *string {
	return s.UserObjectClass
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) GetUserObjectClassCustomFilter() *string {
	return s.UserObjectClassCustomFilter
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) SetGroupMemberAttributeName(v string) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig {
	s.GroupMemberAttributeName = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) SetGroupObjectClass(v string) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig {
	s.GroupObjectClass = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) SetGroupObjectClassCustomFilter(v string) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig {
	s.GroupObjectClassCustomFilter = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) SetOrganizationUnitObjectClass(v string) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig {
	s.OrganizationUnitObjectClass = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) SetUserObjectClass(v string) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig {
	s.UserObjectClass = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) SetUserObjectClassCustomFilter(v string) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig {
	s.UserObjectClassCustomFilter = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationLdapUdPullConfig) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig struct {
	// Cron expression
	//
	// example:
	//
	// 0 45 1 	- 	- ?
	PeriodicSyncCron *string `json:"PeriodicSyncCron,omitempty" xml:"PeriodicSyncCron,omitempty"`
	// Execution time slots, for example 3,5, meaning the task runs once between 03:00–04:00 and once between 05:00–06:00.
	//
	// example:
	//
	// [3,5]
	PeriodicSyncTimes *int32 `json:"PeriodicSyncTimes,omitempty" xml:"PeriodicSyncTimes,omitempty"`
	// type
	//
	// example:
	//
	// cron
	PeriodicSyncType *string `json:"PeriodicSyncType,omitempty" xml:"PeriodicSyncType,omitempty"`
}

func (s GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig) GetPeriodicSyncCron() *string {
	return s.PeriodicSyncCron
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig) GetPeriodicSyncTimes() *int32 {
	return s.PeriodicSyncTimes
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig) GetPeriodicSyncType() *string {
	return s.PeriodicSyncType
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig) SetPeriodicSyncCron(v string) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig {
	s.PeriodicSyncCron = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig) SetPeriodicSyncTimes(v int32) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig {
	s.PeriodicSyncTimes = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig) SetPeriodicSyncType(v string) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig {
	s.PeriodicSyncType = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPeriodicSyncConfig) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule struct {
	// Group Deletion Threshold: If the number of deleted groups exceeds this value, the synchronization task will be terminated.
	//
	// example:
	//
	// 10
	GroupDeletedThreshold *int32 `json:"GroupDeletedThreshold,omitempty" xml:"GroupDeletedThreshold,omitempty"`
	// Organization Deletion Threshold: If the number of deleted organizations exceeds this value, the synchronization task will be terminated.
	//
	// example:
	//
	// 10
	OrganizationalUnitDeletedThreshold *int32 `json:"OrganizationalUnitDeletedThreshold,omitempty" xml:"OrganizationalUnitDeletedThreshold,omitempty"`
	// Account Deletion Threshold: If the number of deleted users exceeds this value, the synchronization task will be terminated.
	//
	// example:
	//
	// 30
	UserDeletedThreshold *int32 `json:"UserDeletedThreshold,omitempty" xml:"UserDeletedThreshold,omitempty"`
}

func (s GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule) GetGroupDeletedThreshold() *int32 {
	return s.GroupDeletedThreshold
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule) GetOrganizationalUnitDeletedThreshold() *int32 {
	return s.OrganizationalUnitDeletedThreshold
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule) GetUserDeletedThreshold() *int32 {
	return s.UserDeletedThreshold
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule) SetGroupDeletedThreshold(v int32) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule {
	s.GroupDeletedThreshold = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule) SetOrganizationalUnitDeletedThreshold(v int32) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule {
	s.OrganizationalUnitDeletedThreshold = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule) SetUserDeletedThreshold(v int32) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule {
	s.UserDeletedThreshold = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationPullProtectedRule) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationUdSyncScopeConfig struct {
	// Synchronization Source Node
	SourceScopes []*string `json:"SourceScopes,omitempty" xml:"SourceScopes,omitempty" type:"Repeated"`
	// Synchronization Target Node
	//
	// example:
	//
	// ou_asjdfhaskfhw213mnsj33sXXX
	TargetScope *string `json:"TargetScope,omitempty" xml:"TargetScope,omitempty"`
}

func (s GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationUdSyncScopeConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationUdSyncScopeConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationUdSyncScopeConfig) GetSourceScopes() []*string {
	return s.SourceScopes
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationUdSyncScopeConfig) GetTargetScope() *string {
	return s.TargetScope
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationUdSyncScopeConfig) SetSourceScopes(v []*string) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationUdSyncScopeConfig {
	s.SourceScopes = v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationUdSyncScopeConfig) SetTargetScope(v string) *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationUdSyncScopeConfig {
	s.TargetScope = &v
	return s
}

func (s *GetIdentityProviderUdPullConfigurationResponseBodyUdPullConfigurationUdSyncScopeConfig) Validate() error {
	return dara.Validate(s)
}
