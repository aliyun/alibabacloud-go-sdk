// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdentityProviderUdPullConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupSyncStatus(v string) *SetIdentityProviderUdPullConfigurationRequest
	GetGroupSyncStatus() *string
	SetIdentityProviderId(v string) *SetIdentityProviderUdPullConfigurationRequest
	GetIdentityProviderId() *string
	SetIncrementalCallbackStatus(v string) *SetIdentityProviderUdPullConfigurationRequest
	GetIncrementalCallbackStatus() *string
	SetInstanceId(v string) *SetIdentityProviderUdPullConfigurationRequest
	GetInstanceId() *string
	SetLdapUdPullConfig(v *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) *SetIdentityProviderUdPullConfigurationRequest
	GetLdapUdPullConfig() *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig
	SetPeriodicSyncConfig(v *SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig) *SetIdentityProviderUdPullConfigurationRequest
	GetPeriodicSyncConfig() *SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig
	SetPeriodicSyncStatus(v string) *SetIdentityProviderUdPullConfigurationRequest
	GetPeriodicSyncStatus() *string
	SetPullProtectedRule(v *SetIdentityProviderUdPullConfigurationRequestPullProtectedRule) *SetIdentityProviderUdPullConfigurationRequest
	GetPullProtectedRule() *SetIdentityProviderUdPullConfigurationRequestPullProtectedRule
	SetUdSyncScopeConfig(v *SetIdentityProviderUdPullConfigurationRequestUdSyncScopeConfig) *SetIdentityProviderUdPullConfigurationRequest
	GetUdSyncScopeConfig() *SetIdentityProviderUdPullConfigurationRequestUdSyncScopeConfig
}

type SetIdentityProviderUdPullConfigurationRequest struct {
	// The group synchronization status. Valid values:
	//
	// - disabled: Disabled.
	//
	// - enabled: Enabled.
	//
	// example:
	//
	// disabled
	GroupSyncStatus *string `json:"GroupSyncStatus,omitempty" xml:"GroupSyncStatus,omitempty"`
	// The identity provider ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idp_my664lwkhpicbyzirog3xxxxx
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// Specifies whether to process incremental callback data from the IdP. Valid values:
	//
	// - disabled: Disabled.
	//
	// - enabled: Enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// disabled
	IncrementalCallbackStatus *string `json:"IncrementalCallbackStatus,omitempty" xml:"IncrementalCallbackStatus,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The LDAP synchronization configuration.
	LdapUdPullConfig *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig `json:"LdapUdPullConfig,omitempty" xml:"LdapUdPullConfig,omitempty" type:"Struct"`
	// The periodic synchronization configuration.
	PeriodicSyncConfig *SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig `json:"PeriodicSyncConfig,omitempty" xml:"PeriodicSyncConfig,omitempty" type:"Struct"`
	// Specifies whether to periodically verify data differences between IDaaS and the identity provider. Valid values:
	//
	// - disabled: Disabled.
	//
	// - enabled: Enabled.
	//
	// example:
	//
	// disabled
	PeriodicSyncStatus *string `json:"PeriodicSyncStatus,omitempty" xml:"PeriodicSyncStatus,omitempty"`
	// The inbound synchronization protection rule.
	PullProtectedRule *SetIdentityProviderUdPullConfigurationRequestPullProtectedRule `json:"PullProtectedRule,omitempty" xml:"PullProtectedRule,omitempty" type:"Struct"`
	// The inbound synchronization scope configuration.
	UdSyncScopeConfig *SetIdentityProviderUdPullConfigurationRequestUdSyncScopeConfig `json:"UdSyncScopeConfig,omitempty" xml:"UdSyncScopeConfig,omitempty" type:"Struct"`
}

func (s SetIdentityProviderUdPullConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderUdPullConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderUdPullConfigurationRequest) GetGroupSyncStatus() *string {
	return s.GroupSyncStatus
}

func (s *SetIdentityProviderUdPullConfigurationRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *SetIdentityProviderUdPullConfigurationRequest) GetIncrementalCallbackStatus() *string {
	return s.IncrementalCallbackStatus
}

func (s *SetIdentityProviderUdPullConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetIdentityProviderUdPullConfigurationRequest) GetLdapUdPullConfig() *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig {
	return s.LdapUdPullConfig
}

func (s *SetIdentityProviderUdPullConfigurationRequest) GetPeriodicSyncConfig() *SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig {
	return s.PeriodicSyncConfig
}

func (s *SetIdentityProviderUdPullConfigurationRequest) GetPeriodicSyncStatus() *string {
	return s.PeriodicSyncStatus
}

func (s *SetIdentityProviderUdPullConfigurationRequest) GetPullProtectedRule() *SetIdentityProviderUdPullConfigurationRequestPullProtectedRule {
	return s.PullProtectedRule
}

func (s *SetIdentityProviderUdPullConfigurationRequest) GetUdSyncScopeConfig() *SetIdentityProviderUdPullConfigurationRequestUdSyncScopeConfig {
	return s.UdSyncScopeConfig
}

func (s *SetIdentityProviderUdPullConfigurationRequest) SetGroupSyncStatus(v string) *SetIdentityProviderUdPullConfigurationRequest {
	s.GroupSyncStatus = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequest) SetIdentityProviderId(v string) *SetIdentityProviderUdPullConfigurationRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequest) SetIncrementalCallbackStatus(v string) *SetIdentityProviderUdPullConfigurationRequest {
	s.IncrementalCallbackStatus = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequest) SetInstanceId(v string) *SetIdentityProviderUdPullConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequest) SetLdapUdPullConfig(v *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) *SetIdentityProviderUdPullConfigurationRequest {
	s.LdapUdPullConfig = v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequest) SetPeriodicSyncConfig(v *SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig) *SetIdentityProviderUdPullConfigurationRequest {
	s.PeriodicSyncConfig = v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequest) SetPeriodicSyncStatus(v string) *SetIdentityProviderUdPullConfigurationRequest {
	s.PeriodicSyncStatus = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequest) SetPullProtectedRule(v *SetIdentityProviderUdPullConfigurationRequestPullProtectedRule) *SetIdentityProviderUdPullConfigurationRequest {
	s.PullProtectedRule = v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequest) SetUdSyncScopeConfig(v *SetIdentityProviderUdPullConfigurationRequestUdSyncScopeConfig) *SetIdentityProviderUdPullConfigurationRequest {
	s.UdSyncScopeConfig = v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequest) Validate() error {
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

type SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig struct {
	// The group member identifier.
	//
	// example:
	//
	// member
	GroupMemberAttributeName *string `json:"GroupMemberAttributeName,omitempty" xml:"GroupMemberAttributeName,omitempty"`
	// The group ObjectClass.
	//
	// example:
	//
	// group
	GroupObjectClass *string `json:"GroupObjectClass,omitempty" xml:"GroupObjectClass,omitempty"`
	// The custom filter for groups.
	//
	// example:
	//
	// (|(cn=test)(group=test@test.com))
	GroupObjectClassCustomFilter *string `json:"GroupObjectClassCustomFilter,omitempty" xml:"GroupObjectClassCustomFilter,omitempty"`
	// The organizational unit ObjectClass.
	//
	// example:
	//
	// organizationUnit,top
	OrganizationUnitObjectClass *string `json:"OrganizationUnitObjectClass,omitempty" xml:"OrganizationUnitObjectClass,omitempty"`
	// The user ObjectClass.
	//
	// example:
	//
	// userPrincipalName, mail
	UserObjectClass *string `json:"UserObjectClass,omitempty" xml:"UserObjectClass,omitempty"`
	// The custom filter for users.
	//
	// example:
	//
	// (|(cn=test)(mail=test@test.com))
	UserObjectClassCustomFilter *string `json:"UserObjectClassCustomFilter,omitempty" xml:"UserObjectClassCustomFilter,omitempty"`
}

func (s SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) GetGroupMemberAttributeName() *string {
	return s.GroupMemberAttributeName
}

func (s *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) GetGroupObjectClass() *string {
	return s.GroupObjectClass
}

func (s *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) GetGroupObjectClassCustomFilter() *string {
	return s.GroupObjectClassCustomFilter
}

func (s *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) GetOrganizationUnitObjectClass() *string {
	return s.OrganizationUnitObjectClass
}

func (s *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) GetUserObjectClass() *string {
	return s.UserObjectClass
}

func (s *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) GetUserObjectClassCustomFilter() *string {
	return s.UserObjectClassCustomFilter
}

func (s *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) SetGroupMemberAttributeName(v string) *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig {
	s.GroupMemberAttributeName = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) SetGroupObjectClass(v string) *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig {
	s.GroupObjectClass = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) SetGroupObjectClassCustomFilter(v string) *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig {
	s.GroupObjectClassCustomFilter = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) SetOrganizationUnitObjectClass(v string) *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig {
	s.OrganizationUnitObjectClass = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) SetUserObjectClass(v string) *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig {
	s.UserObjectClass = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) SetUserObjectClassCustomFilter(v string) *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig {
	s.UserObjectClassCustomFilter = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequestLdapUdPullConfig) Validate() error {
	return dara.Validate(s)
}

type SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig struct {
	// The cron expression. This parameter is required when periodicSyncType is set to cron.
	//
	// example:
	//
	// 0 45 1 	- 	- ?
	PeriodicSyncCron *string `json:"PeriodicSyncCron,omitempty" xml:"PeriodicSyncCron,omitempty"`
	// The execution time points. This parameter is required when periodicSyncType is set to time. For example, [3, 5] indicates that the task runs once between 3:00 and 4:00 and once between 5:00 and 6:00.
	PeriodicSyncTimes []*int32 `json:"PeriodicSyncTimes,omitempty" xml:"PeriodicSyncTimes,omitempty" type:"Repeated"`
	// The type of periodic synchronization.
	//
	// example:
	//
	// cron
	PeriodicSyncType *string `json:"PeriodicSyncType,omitempty" xml:"PeriodicSyncType,omitempty"`
}

func (s SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig) GetPeriodicSyncCron() *string {
	return s.PeriodicSyncCron
}

func (s *SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig) GetPeriodicSyncTimes() []*int32 {
	return s.PeriodicSyncTimes
}

func (s *SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig) GetPeriodicSyncType() *string {
	return s.PeriodicSyncType
}

func (s *SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig) SetPeriodicSyncCron(v string) *SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig {
	s.PeriodicSyncCron = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig) SetPeriodicSyncTimes(v []*int32) *SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig {
	s.PeriodicSyncTimes = v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig) SetPeriodicSyncType(v string) *SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig {
	s.PeriodicSyncType = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequestPeriodicSyncConfig) Validate() error {
	return dara.Validate(s)
}

type SetIdentityProviderUdPullConfigurationRequestPullProtectedRule struct {
	// The threshold for the number of deleted groups. If the number of deleted groups exceeds this value, the synchronization task is terminated.
	//
	// example:
	//
	// 10
	GroupDeletedThreshold *int32 `json:"GroupDeletedThreshold,omitempty" xml:"GroupDeletedThreshold,omitempty"`
	// The threshold for the number of deleted organizational units. If the number of deleted organizational units exceeds this value, the synchronization task is terminated.
	//
	// example:
	//
	// 10
	OrganizationalUnitDeletedThreshold *int32 `json:"OrganizationalUnitDeletedThreshold,omitempty" xml:"OrganizationalUnitDeletedThreshold,omitempty"`
	// The threshold for the number of deleted accounts. If the number of deleted users exceeds this value, the synchronization task is terminated.
	//
	// example:
	//
	// 30
	UserDeletedThreshold *int32 `json:"UserDeletedThreshold,omitempty" xml:"UserDeletedThreshold,omitempty"`
}

func (s SetIdentityProviderUdPullConfigurationRequestPullProtectedRule) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderUdPullConfigurationRequestPullProtectedRule) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderUdPullConfigurationRequestPullProtectedRule) GetGroupDeletedThreshold() *int32 {
	return s.GroupDeletedThreshold
}

func (s *SetIdentityProviderUdPullConfigurationRequestPullProtectedRule) GetOrganizationalUnitDeletedThreshold() *int32 {
	return s.OrganizationalUnitDeletedThreshold
}

func (s *SetIdentityProviderUdPullConfigurationRequestPullProtectedRule) GetUserDeletedThreshold() *int32 {
	return s.UserDeletedThreshold
}

func (s *SetIdentityProviderUdPullConfigurationRequestPullProtectedRule) SetGroupDeletedThreshold(v int32) *SetIdentityProviderUdPullConfigurationRequestPullProtectedRule {
	s.GroupDeletedThreshold = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequestPullProtectedRule) SetOrganizationalUnitDeletedThreshold(v int32) *SetIdentityProviderUdPullConfigurationRequestPullProtectedRule {
	s.OrganizationalUnitDeletedThreshold = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequestPullProtectedRule) SetUserDeletedThreshold(v int32) *SetIdentityProviderUdPullConfigurationRequestPullProtectedRule {
	s.UserDeletedThreshold = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequestPullProtectedRule) Validate() error {
	return dara.Validate(s)
}

type SetIdentityProviderUdPullConfigurationRequestUdSyncScopeConfig struct {
	// The list of source nodes for synchronization.
	SourceScopes []*string `json:"SourceScopes,omitempty" xml:"SourceScopes,omitempty" type:"Repeated"`
	// The target node for synchronization.
	//
	// example:
	//
	// ou_asdaq1addsxzdq1XXX
	TargetScope *string `json:"TargetScope,omitempty" xml:"TargetScope,omitempty"`
}

func (s SetIdentityProviderUdPullConfigurationRequestUdSyncScopeConfig) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderUdPullConfigurationRequestUdSyncScopeConfig) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderUdPullConfigurationRequestUdSyncScopeConfig) GetSourceScopes() []*string {
	return s.SourceScopes
}

func (s *SetIdentityProviderUdPullConfigurationRequestUdSyncScopeConfig) GetTargetScope() *string {
	return s.TargetScope
}

func (s *SetIdentityProviderUdPullConfigurationRequestUdSyncScopeConfig) SetSourceScopes(v []*string) *SetIdentityProviderUdPullConfigurationRequestUdSyncScopeConfig {
	s.SourceScopes = v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequestUdSyncScopeConfig) SetTargetScope(v string) *SetIdentityProviderUdPullConfigurationRequestUdSyncScopeConfig {
	s.TargetScope = &v
	return s
}

func (s *SetIdentityProviderUdPullConfigurationRequestUdSyncScopeConfig) Validate() error {
	return dara.Validate(s)
}
