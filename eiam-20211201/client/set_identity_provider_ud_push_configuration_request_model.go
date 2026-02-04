// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdentityProviderUdPushConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderId(v string) *SetIdentityProviderUdPushConfigurationRequest
	GetIdentityProviderId() *string
	SetIncrementalCallbackStatus(v string) *SetIdentityProviderUdPushConfigurationRequest
	GetIncrementalCallbackStatus() *string
	SetInstanceId(v string) *SetIdentityProviderUdPushConfigurationRequest
	GetInstanceId() *string
	SetLdapUdPushConfig(v *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig) *SetIdentityProviderUdPushConfigurationRequest
	GetLdapUdPushConfig() *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig
	SetPeriodicSyncConfig(v *SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig) *SetIdentityProviderUdPushConfigurationRequest
	GetPeriodicSyncConfig() *SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig
	SetPeriodicSyncStatus(v string) *SetIdentityProviderUdPushConfigurationRequest
	GetPeriodicSyncStatus() *string
	SetUdSyncScopeConfigs(v []*SetIdentityProviderUdPushConfigurationRequestUdSyncScopeConfigs) *SetIdentityProviderUdPushConfigurationRequest
	GetUdSyncScopeConfigs() []*SetIdentityProviderUdPushConfigurationRequestUdSyncScopeConfigs
}

type SetIdentityProviderUdPushConfigurationRequest struct {
	// IDaaS的身份提供方主键id
	//
	// This parameter is required.
	//
	// example:
	//
	// idp_11111
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// 增量回调状态，是否处理来自IdP的增量回调数据
	//
	// This parameter is required.
	//
	// example:
	//
	// disabled
	IncrementalCallbackStatus *string `json:"IncrementalCallbackStatus,omitempty" xml:"IncrementalCallbackStatus,omitempty"`
	// IDaaS EIAM的实例id
	//
	// This parameter is required.
	//
	// example:
	//
	// eiam-111ccc1111
	InstanceId         *string                                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LdapUdPushConfig   *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig   `json:"LdapUdPushConfig,omitempty" xml:"LdapUdPushConfig,omitempty" type:"Struct"`
	PeriodicSyncConfig *SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig `json:"PeriodicSyncConfig,omitempty" xml:"PeriodicSyncConfig,omitempty" type:"Struct"`
	// example:
	//
	// disabled
	PeriodicSyncStatus *string `json:"PeriodicSyncStatus,omitempty" xml:"PeriodicSyncStatus,omitempty"`
	// 同步出配置信息
	UdSyncScopeConfigs []*SetIdentityProviderUdPushConfigurationRequestUdSyncScopeConfigs `json:"UdSyncScopeConfigs,omitempty" xml:"UdSyncScopeConfigs,omitempty" type:"Repeated"`
}

func (s SetIdentityProviderUdPushConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderUdPushConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderUdPushConfigurationRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *SetIdentityProviderUdPushConfigurationRequest) GetIncrementalCallbackStatus() *string {
	return s.IncrementalCallbackStatus
}

func (s *SetIdentityProviderUdPushConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetIdentityProviderUdPushConfigurationRequest) GetLdapUdPushConfig() *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig {
	return s.LdapUdPushConfig
}

func (s *SetIdentityProviderUdPushConfigurationRequest) GetPeriodicSyncConfig() *SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig {
	return s.PeriodicSyncConfig
}

func (s *SetIdentityProviderUdPushConfigurationRequest) GetPeriodicSyncStatus() *string {
	return s.PeriodicSyncStatus
}

func (s *SetIdentityProviderUdPushConfigurationRequest) GetUdSyncScopeConfigs() []*SetIdentityProviderUdPushConfigurationRequestUdSyncScopeConfigs {
	return s.UdSyncScopeConfigs
}

func (s *SetIdentityProviderUdPushConfigurationRequest) SetIdentityProviderId(v string) *SetIdentityProviderUdPushConfigurationRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequest) SetIncrementalCallbackStatus(v string) *SetIdentityProviderUdPushConfigurationRequest {
	s.IncrementalCallbackStatus = &v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequest) SetInstanceId(v string) *SetIdentityProviderUdPushConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequest) SetLdapUdPushConfig(v *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig) *SetIdentityProviderUdPushConfigurationRequest {
	s.LdapUdPushConfig = v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequest) SetPeriodicSyncConfig(v *SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig) *SetIdentityProviderUdPushConfigurationRequest {
	s.PeriodicSyncConfig = v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequest) SetPeriodicSyncStatus(v string) *SetIdentityProviderUdPushConfigurationRequest {
	s.PeriodicSyncStatus = &v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequest) SetUdSyncScopeConfigs(v []*SetIdentityProviderUdPushConfigurationRequestUdSyncScopeConfigs) *SetIdentityProviderUdPushConfigurationRequest {
	s.UdSyncScopeConfigs = v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequest) Validate() error {
	if s.LdapUdPushConfig != nil {
		if err := s.LdapUdPushConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PeriodicSyncConfig != nil {
		if err := s.PeriodicSyncConfig.Validate(); err != nil {
			return err
		}
	}
	if s.UdSyncScopeConfigs != nil {
		for _, item := range s.UdSyncScopeConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig struct {
	// example:
	//
	// ou,top
	OrganizationUnitObjectClass *string `json:"OrganizationUnitObjectClass,omitempty" xml:"OrganizationUnitObjectClass,omitempty"`
	// example:
	//
	// ou
	OrganizationalUnitRdn *string `json:"OrganizationalUnitRdn,omitempty" xml:"OrganizationalUnitRdn,omitempty"`
	// example:
	//
	// enabled
	PasswordSyncStatus *string `json:"PasswordSyncStatus,omitempty" xml:"PasswordSyncStatus,omitempty"`
	// example:
	//
	// user,top
	UserObjectClass *string `json:"UserObjectClass,omitempty" xml:"UserObjectClass,omitempty"`
	// example:
	//
	// cn
	UserRdn *string `json:"UserRdn,omitempty" xml:"UserRdn,omitempty"`
}

func (s SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig) GetOrganizationUnitObjectClass() *string {
	return s.OrganizationUnitObjectClass
}

func (s *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig) GetOrganizationalUnitRdn() *string {
	return s.OrganizationalUnitRdn
}

func (s *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig) GetPasswordSyncStatus() *string {
	return s.PasswordSyncStatus
}

func (s *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig) GetUserObjectClass() *string {
	return s.UserObjectClass
}

func (s *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig) GetUserRdn() *string {
	return s.UserRdn
}

func (s *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig) SetOrganizationUnitObjectClass(v string) *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig {
	s.OrganizationUnitObjectClass = &v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig) SetOrganizationalUnitRdn(v string) *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig {
	s.OrganizationalUnitRdn = &v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig) SetPasswordSyncStatus(v string) *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig {
	s.PasswordSyncStatus = &v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig) SetUserObjectClass(v string) *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig {
	s.UserObjectClass = &v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig) SetUserRdn(v string) *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig {
	s.UserRdn = &v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequestLdapUdPushConfig) Validate() error {
	return dara.Validate(s)
}

type SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig struct {
	// example:
	//
	// 0 45 1 	- 	- ?
	PeriodicSyncCron  *string  `json:"PeriodicSyncCron,omitempty" xml:"PeriodicSyncCron,omitempty"`
	PeriodicSyncTimes []*int32 `json:"PeriodicSyncTimes,omitempty" xml:"PeriodicSyncTimes,omitempty" type:"Repeated"`
	// example:
	//
	// cron
	PeriodicSyncType *string `json:"PeriodicSyncType,omitempty" xml:"PeriodicSyncType,omitempty"`
}

func (s SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig) GetPeriodicSyncCron() *string {
	return s.PeriodicSyncCron
}

func (s *SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig) GetPeriodicSyncTimes() []*int32 {
	return s.PeriodicSyncTimes
}

func (s *SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig) GetPeriodicSyncType() *string {
	return s.PeriodicSyncType
}

func (s *SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig) SetPeriodicSyncCron(v string) *SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig {
	s.PeriodicSyncCron = &v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig) SetPeriodicSyncTimes(v []*int32) *SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig {
	s.PeriodicSyncTimes = v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig) SetPeriodicSyncType(v string) *SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig {
	s.PeriodicSyncType = &v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequestPeriodicSyncConfig) Validate() error {
	return dara.Validate(s)
}

type SetIdentityProviderUdPushConfigurationRequestUdSyncScopeConfigs struct {
	// 同步来源节点
	SourceScopes []*string `json:"SourceScopes,omitempty" xml:"SourceScopes,omitempty" type:"Repeated"`
	// 同步目标节点
	//
	// example:
	//
	// 6537211
	TargetScope *string `json:"TargetScope,omitempty" xml:"TargetScope,omitempty"`
}

func (s SetIdentityProviderUdPushConfigurationRequestUdSyncScopeConfigs) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderUdPushConfigurationRequestUdSyncScopeConfigs) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderUdPushConfigurationRequestUdSyncScopeConfigs) GetSourceScopes() []*string {
	return s.SourceScopes
}

func (s *SetIdentityProviderUdPushConfigurationRequestUdSyncScopeConfigs) GetTargetScope() *string {
	return s.TargetScope
}

func (s *SetIdentityProviderUdPushConfigurationRequestUdSyncScopeConfigs) SetSourceScopes(v []*string) *SetIdentityProviderUdPushConfigurationRequestUdSyncScopeConfigs {
	s.SourceScopes = v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequestUdSyncScopeConfigs) SetTargetScope(v string) *SetIdentityProviderUdPushConfigurationRequestUdSyncScopeConfigs {
	s.TargetScope = &v
	return s
}

func (s *SetIdentityProviderUdPushConfigurationRequestUdSyncScopeConfigs) Validate() error {
	return dara.Validate(s)
}
