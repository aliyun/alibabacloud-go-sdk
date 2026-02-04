// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderUdPushConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetIdentityProviderUdPushConfigurationResponseBody
	GetRequestId() *string
	SetUdPushConfiguration(v *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) *GetIdentityProviderUdPushConfigurationResponseBody
	GetUdPushConfiguration() *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration
}

type GetIdentityProviderUdPushConfigurationResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId           *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UdPushConfiguration *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration `json:"UdPushConfiguration,omitempty" xml:"UdPushConfiguration,omitempty" type:"Struct"`
}

func (s GetIdentityProviderUdPushConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderUdPushConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderUdPushConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentityProviderUdPushConfigurationResponseBody) GetUdPushConfiguration() *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration {
	return s.UdPushConfiguration
}

func (s *GetIdentityProviderUdPushConfigurationResponseBody) SetRequestId(v string) *GetIdentityProviderUdPushConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBody) SetUdPushConfiguration(v *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) *GetIdentityProviderUdPushConfigurationResponseBody {
	s.UdPushConfiguration = v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBody) Validate() error {
	if s.UdPushConfiguration != nil {
		if err := s.UdPushConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration struct {
	// IDaaS EIAM 身份提供方ID
	//
	// example:
	//
	// idp_na2rzpyc67zr7ixdfy35zgrxxx
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// 增量回调状态，是否处理来自IdP的增量回调数据
	//
	// example:
	//
	// enabled
	IncrementalCallbackStatus *string `json:"IncrementalCallbackStatus,omitempty" xml:"IncrementalCallbackStatus,omitempty"`
	// IDaaS EIAM 实例Id
	//
	// example:
	//
	// idaas_7vdv3olzk36gymwtlaq6fixxxx
	InstanceId         *string                                                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LdapUdPushConfig   *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig   `json:"LdapUdPushConfig,omitempty" xml:"LdapUdPushConfig,omitempty" type:"Struct"`
	PeriodicSyncConfig *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig `json:"PeriodicSyncConfig,omitempty" xml:"PeriodicSyncConfig,omitempty" type:"Struct"`
	// example:
	//
	// disabled
	PeriodicSyncStatus *string `json:"PeriodicSyncStatus,omitempty" xml:"PeriodicSyncStatus,omitempty"`
	// 同步出配置信息
	UdSyncScopeConfigs []*GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationUdSyncScopeConfigs `json:"UdSyncScopeConfigs,omitempty" xml:"UdSyncScopeConfigs,omitempty" type:"Repeated"`
}

func (s GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) GetIncrementalCallbackStatus() *string {
	return s.IncrementalCallbackStatus
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) GetLdapUdPushConfig() *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig {
	return s.LdapUdPushConfig
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) GetPeriodicSyncConfig() *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig {
	return s.PeriodicSyncConfig
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) GetPeriodicSyncStatus() *string {
	return s.PeriodicSyncStatus
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) GetUdSyncScopeConfigs() []*GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationUdSyncScopeConfigs {
	return s.UdSyncScopeConfigs
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) SetIdentityProviderId(v string) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration {
	s.IdentityProviderId = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) SetIncrementalCallbackStatus(v string) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration {
	s.IncrementalCallbackStatus = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) SetInstanceId(v string) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration {
	s.InstanceId = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) SetLdapUdPushConfig(v *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration {
	s.LdapUdPushConfig = v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) SetPeriodicSyncConfig(v *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration {
	s.PeriodicSyncConfig = v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) SetPeriodicSyncStatus(v string) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration {
	s.PeriodicSyncStatus = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) SetUdSyncScopeConfigs(v []*GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationUdSyncScopeConfigs) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration {
	s.UdSyncScopeConfigs = v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfiguration) Validate() error {
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

type GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig struct {
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

func (s GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig) GetOrganizationUnitObjectClass() *string {
	return s.OrganizationUnitObjectClass
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig) GetOrganizationalUnitRdn() *string {
	return s.OrganizationalUnitRdn
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig) GetPasswordSyncStatus() *string {
	return s.PasswordSyncStatus
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig) GetUserObjectClass() *string {
	return s.UserObjectClass
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig) GetUserRdn() *string {
	return s.UserRdn
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig) SetOrganizationUnitObjectClass(v string) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig {
	s.OrganizationUnitObjectClass = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig) SetOrganizationalUnitRdn(v string) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig {
	s.OrganizationalUnitRdn = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig) SetPasswordSyncStatus(v string) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig {
	s.PasswordSyncStatus = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig) SetUserObjectClass(v string) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig {
	s.UserObjectClass = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig) SetUserRdn(v string) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig {
	s.UserRdn = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationLdapUdPushConfig) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig struct {
	// example:
	//
	// 0 45 1 	- 	- ?
	PeriodicSyncCron  *string   `json:"PeriodicSyncCron,omitempty" xml:"PeriodicSyncCron,omitempty"`
	PeriodicSyncTimes []*string `json:"PeriodicSyncTimes,omitempty" xml:"PeriodicSyncTimes,omitempty" type:"Repeated"`
	// example:
	//
	// cron
	PeriodicSyncType *string `json:"PeriodicSyncType,omitempty" xml:"PeriodicSyncType,omitempty"`
}

func (s GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig) GetPeriodicSyncCron() *string {
	return s.PeriodicSyncCron
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig) GetPeriodicSyncTimes() []*string {
	return s.PeriodicSyncTimes
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig) GetPeriodicSyncType() *string {
	return s.PeriodicSyncType
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig) SetPeriodicSyncCron(v string) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig {
	s.PeriodicSyncCron = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig) SetPeriodicSyncTimes(v []*string) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig {
	s.PeriodicSyncTimes = v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig) SetPeriodicSyncType(v string) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig {
	s.PeriodicSyncType = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationPeriodicSyncConfig) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationUdSyncScopeConfigs struct {
	// 同步来源节点
	SourceScopes []*string `json:"SourceScopes,omitempty" xml:"SourceScopes,omitempty" type:"Repeated"`
	// 同步目标节点
	//
	// example:
	//
	// 604352338
	TargetScope *string `json:"TargetScope,omitempty" xml:"TargetScope,omitempty"`
}

func (s GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationUdSyncScopeConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationUdSyncScopeConfigs) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationUdSyncScopeConfigs) GetSourceScopes() []*string {
	return s.SourceScopes
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationUdSyncScopeConfigs) GetTargetScope() *string {
	return s.TargetScope
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationUdSyncScopeConfigs) SetSourceScopes(v []*string) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationUdSyncScopeConfigs {
	s.SourceScopes = v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationUdSyncScopeConfigs) SetTargetScope(v string) *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationUdSyncScopeConfigs {
	s.TargetScope = &v
	return s
}

func (s *GetIdentityProviderUdPushConfigurationResponseBodyUdPushConfigurationUdSyncScopeConfigs) Validate() error {
	return dara.Validate(s)
}
