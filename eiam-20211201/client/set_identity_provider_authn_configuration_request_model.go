// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdentityProviderAuthnConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCreateUserConfig(v *SetIdentityProviderAuthnConfigurationRequestAutoCreateUserConfig) *SetIdentityProviderAuthnConfigurationRequest
	GetAutoCreateUserConfig() *SetIdentityProviderAuthnConfigurationRequestAutoCreateUserConfig
	SetAutoUpdateUserConfig(v *SetIdentityProviderAuthnConfigurationRequestAutoUpdateUserConfig) *SetIdentityProviderAuthnConfigurationRequest
	GetAutoUpdateUserConfig() *SetIdentityProviderAuthnConfigurationRequestAutoUpdateUserConfig
	SetBindingConfig(v *SetIdentityProviderAuthnConfigurationRequestBindingConfig) *SetIdentityProviderAuthnConfigurationRequest
	GetBindingConfig() *SetIdentityProviderAuthnConfigurationRequestBindingConfig
	SetIdentityProviderId(v string) *SetIdentityProviderAuthnConfigurationRequest
	GetIdentityProviderId() *string
	SetInstanceId(v string) *SetIdentityProviderAuthnConfigurationRequest
	GetInstanceId() *string
	SetLdapAuthnConfig(v *SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig) *SetIdentityProviderAuthnConfigurationRequest
	GetLdapAuthnConfig() *SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig
}

type SetIdentityProviderAuthnConfigurationRequest struct {
	// 自动创建账户账户规则配置。
	AutoCreateUserConfig *SetIdentityProviderAuthnConfigurationRequestAutoCreateUserConfig `json:"AutoCreateUserConfig,omitempty" xml:"AutoCreateUserConfig,omitempty" type:"Struct"`
	AutoUpdateUserConfig *SetIdentityProviderAuthnConfigurationRequestAutoUpdateUserConfig `json:"AutoUpdateUserConfig,omitempty" xml:"AutoUpdateUserConfig,omitempty" type:"Struct"`
	// 账户绑定规则配置
	BindingConfig *SetIdentityProviderAuthnConfigurationRequestBindingConfig `json:"BindingConfig,omitempty" xml:"BindingConfig,omitempty" type:"Struct"`
	// IDaaS的身份提供方主键id
	//
	// This parameter is required.
	//
	// example:
	//
	// idp_11111
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// AD/LDAP配置
	LdapAuthnConfig *SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig `json:"LdapAuthnConfig,omitempty" xml:"LdapAuthnConfig,omitempty" type:"Struct"`
}

func (s SetIdentityProviderAuthnConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderAuthnConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderAuthnConfigurationRequest) GetAutoCreateUserConfig() *SetIdentityProviderAuthnConfigurationRequestAutoCreateUserConfig {
	return s.AutoCreateUserConfig
}

func (s *SetIdentityProviderAuthnConfigurationRequest) GetAutoUpdateUserConfig() *SetIdentityProviderAuthnConfigurationRequestAutoUpdateUserConfig {
	return s.AutoUpdateUserConfig
}

func (s *SetIdentityProviderAuthnConfigurationRequest) GetBindingConfig() *SetIdentityProviderAuthnConfigurationRequestBindingConfig {
	return s.BindingConfig
}

func (s *SetIdentityProviderAuthnConfigurationRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *SetIdentityProviderAuthnConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetIdentityProviderAuthnConfigurationRequest) GetLdapAuthnConfig() *SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig {
	return s.LdapAuthnConfig
}

func (s *SetIdentityProviderAuthnConfigurationRequest) SetAutoCreateUserConfig(v *SetIdentityProviderAuthnConfigurationRequestAutoCreateUserConfig) *SetIdentityProviderAuthnConfigurationRequest {
	s.AutoCreateUserConfig = v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequest) SetAutoUpdateUserConfig(v *SetIdentityProviderAuthnConfigurationRequestAutoUpdateUserConfig) *SetIdentityProviderAuthnConfigurationRequest {
	s.AutoUpdateUserConfig = v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequest) SetBindingConfig(v *SetIdentityProviderAuthnConfigurationRequestBindingConfig) *SetIdentityProviderAuthnConfigurationRequest {
	s.BindingConfig = v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequest) SetIdentityProviderId(v string) *SetIdentityProviderAuthnConfigurationRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequest) SetInstanceId(v string) *SetIdentityProviderAuthnConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequest) SetLdapAuthnConfig(v *SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig) *SetIdentityProviderAuthnConfigurationRequest {
	s.LdapAuthnConfig = v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequest) Validate() error {
	if s.AutoCreateUserConfig != nil {
		if err := s.AutoCreateUserConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AutoUpdateUserConfig != nil {
		if err := s.AutoUpdateUserConfig.Validate(); err != nil {
			return err
		}
	}
	if s.BindingConfig != nil {
		if err := s.BindingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LdapAuthnConfig != nil {
		if err := s.LdapAuthnConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetIdentityProviderAuthnConfigurationRequestAutoCreateUserConfig struct {
	// 自动创建账户是否开启
	//
	// example:
	//
	// disabled
	AutoCreateUserStatus        *string   `json:"AutoCreateUserStatus,omitempty" xml:"AutoCreateUserStatus,omitempty"`
	TargetOrganizationalUnitIds []*string `json:"TargetOrganizationalUnitIds,omitempty" xml:"TargetOrganizationalUnitIds,omitempty" type:"Repeated"`
}

func (s SetIdentityProviderAuthnConfigurationRequestAutoCreateUserConfig) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderAuthnConfigurationRequestAutoCreateUserConfig) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderAuthnConfigurationRequestAutoCreateUserConfig) GetAutoCreateUserStatus() *string {
	return s.AutoCreateUserStatus
}

func (s *SetIdentityProviderAuthnConfigurationRequestAutoCreateUserConfig) GetTargetOrganizationalUnitIds() []*string {
	return s.TargetOrganizationalUnitIds
}

func (s *SetIdentityProviderAuthnConfigurationRequestAutoCreateUserConfig) SetAutoCreateUserStatus(v string) *SetIdentityProviderAuthnConfigurationRequestAutoCreateUserConfig {
	s.AutoCreateUserStatus = &v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequestAutoCreateUserConfig) SetTargetOrganizationalUnitIds(v []*string) *SetIdentityProviderAuthnConfigurationRequestAutoCreateUserConfig {
	s.TargetOrganizationalUnitIds = v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequestAutoCreateUserConfig) Validate() error {
	return dara.Validate(s)
}

type SetIdentityProviderAuthnConfigurationRequestAutoUpdateUserConfig struct {
	// example:
	//
	// disabled
	AutoUpdateUserStatus *string `json:"AutoUpdateUserStatus,omitempty" xml:"AutoUpdateUserStatus,omitempty"`
}

func (s SetIdentityProviderAuthnConfigurationRequestAutoUpdateUserConfig) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderAuthnConfigurationRequestAutoUpdateUserConfig) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderAuthnConfigurationRequestAutoUpdateUserConfig) GetAutoUpdateUserStatus() *string {
	return s.AutoUpdateUserStatus
}

func (s *SetIdentityProviderAuthnConfigurationRequestAutoUpdateUserConfig) SetAutoUpdateUserStatus(v string) *SetIdentityProviderAuthnConfigurationRequestAutoUpdateUserConfig {
	s.AutoUpdateUserStatus = &v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequestAutoUpdateUserConfig) Validate() error {
	return dara.Validate(s)
}

type SetIdentityProviderAuthnConfigurationRequestBindingConfig struct {
	// 自动匹配账户的规则
	AutoMatchUserProfileExpressions []*SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions `json:"AutoMatchUserProfileExpressions,omitempty" xml:"AutoMatchUserProfileExpressions,omitempty" type:"Repeated"`
	// 自动匹配账户是否开启
	//
	// example:
	//
	// disabled
	AutoMatchUserStatus *string `json:"AutoMatchUserStatus,omitempty" xml:"AutoMatchUserStatus,omitempty"`
	// 用户手动绑定账户功能是否开启
	//
	// example:
	//
	// enabled
	MappingBindingStatus *string `json:"MappingBindingStatus,omitempty" xml:"MappingBindingStatus,omitempty"`
}

func (s SetIdentityProviderAuthnConfigurationRequestBindingConfig) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderAuthnConfigurationRequestBindingConfig) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfig) GetAutoMatchUserProfileExpressions() []*SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions {
	return s.AutoMatchUserProfileExpressions
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfig) GetAutoMatchUserStatus() *string {
	return s.AutoMatchUserStatus
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfig) GetMappingBindingStatus() *string {
	return s.MappingBindingStatus
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfig) SetAutoMatchUserProfileExpressions(v []*SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions) *SetIdentityProviderAuthnConfigurationRequestBindingConfig {
	s.AutoMatchUserProfileExpressions = v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfig) SetAutoMatchUserStatus(v string) *SetIdentityProviderAuthnConfigurationRequestBindingConfig {
	s.AutoMatchUserStatus = &v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfig) SetMappingBindingStatus(v string) *SetIdentityProviderAuthnConfigurationRequestBindingConfig {
	s.MappingBindingStatus = &v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfig) Validate() error {
	if s.AutoMatchUserProfileExpressions != nil {
		for _, item := range s.AutoMatchUserProfileExpressions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions struct {
	// 表达式的类型
	//
	// This parameter is required.
	//
	// example:
	//
	// filed
	ExpressionMappingType *string `json:"ExpressionMappingType,omitempty" xml:"ExpressionMappingType,omitempty"`
	// 映射属性取值表达式
	//
	// This parameter is required.
	//
	// example:
	//
	// idpUser.userId
	SourceValueExpression *string `json:"SourceValueExpression,omitempty" xml:"SourceValueExpression,omitempty"`
	// 映射目标属性名称
	//
	// This parameter is required.
	//
	// example:
	//
	// user.username
	TargetField *string `json:"TargetField,omitempty" xml:"TargetField,omitempty"`
	// 映射目标属性名称
	//
	// example:
	//
	// username
	TargetFieldDescription *string `json:"TargetFieldDescription,omitempty" xml:"TargetFieldDescription,omitempty"`
}

func (s SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions) GetExpressionMappingType() *string {
	return s.ExpressionMappingType
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions) GetSourceValueExpression() *string {
	return s.SourceValueExpression
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions) GetTargetField() *string {
	return s.TargetField
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions) GetTargetFieldDescription() *string {
	return s.TargetFieldDescription
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions) SetExpressionMappingType(v string) *SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions {
	s.ExpressionMappingType = &v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions) SetSourceValueExpression(v string) *SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions {
	s.SourceValueExpression = &v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions) SetTargetField(v string) *SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions {
	s.TargetField = &v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions) SetTargetFieldDescription(v string) *SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions {
	s.TargetFieldDescription = &v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequestBindingConfigAutoMatchUserProfileExpressions) Validate() error {
	return dara.Validate(s)
}

type SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig struct {
	// 是否支持自动更新密码
	//
	// example:
	//
	// enabled
	AutoUpdatePasswordStatus *string `json:"AutoUpdatePasswordStatus,omitempty" xml:"AutoUpdatePasswordStatus,omitempty"`
	// 用户登录标识
	//
	// example:
	//
	// email
	UserLoginIdentifier *string `json:"UserLoginIdentifier,omitempty" xml:"UserLoginIdentifier,omitempty"`
	// 用户ObjectClass
	//
	// example:
	//
	// posixAccount
	UserObjectClass *string `json:"UserObjectClass,omitempty" xml:"UserObjectClass,omitempty"`
}

func (s SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig) GetAutoUpdatePasswordStatus() *string {
	return s.AutoUpdatePasswordStatus
}

func (s *SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig) GetUserLoginIdentifier() *string {
	return s.UserLoginIdentifier
}

func (s *SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig) GetUserObjectClass() *string {
	return s.UserObjectClass
}

func (s *SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig) SetAutoUpdatePasswordStatus(v string) *SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig {
	s.AutoUpdatePasswordStatus = &v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig) SetUserLoginIdentifier(v string) *SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig {
	s.UserLoginIdentifier = &v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig) SetUserObjectClass(v string) *SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig {
	s.UserObjectClass = &v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationRequestLdapAuthnConfig) Validate() error {
	return dara.Validate(s)
}
