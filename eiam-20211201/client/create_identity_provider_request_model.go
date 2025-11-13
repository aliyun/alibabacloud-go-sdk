// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthnConfig(v *CreateIdentityProviderRequestAuthnConfig) *CreateIdentityProviderRequest
	GetAuthnConfig() *CreateIdentityProviderRequestAuthnConfig
	SetAutoCreateUserConfig(v *CreateIdentityProviderRequestAutoCreateUserConfig) *CreateIdentityProviderRequest
	GetAutoCreateUserConfig() *CreateIdentityProviderRequestAutoCreateUserConfig
	SetAutoUpdateUserConfig(v *CreateIdentityProviderRequestAutoUpdateUserConfig) *CreateIdentityProviderRequest
	GetAutoUpdateUserConfig() *CreateIdentityProviderRequestAutoUpdateUserConfig
	SetBindingConfig(v *CreateIdentityProviderRequestBindingConfig) *CreateIdentityProviderRequest
	GetBindingConfig() *CreateIdentityProviderRequestBindingConfig
	SetClientToken(v string) *CreateIdentityProviderRequest
	GetClientToken() *string
	SetDingtalkAppConfig(v *CreateIdentityProviderRequestDingtalkAppConfig) *CreateIdentityProviderRequest
	GetDingtalkAppConfig() *CreateIdentityProviderRequestDingtalkAppConfig
	SetIdentityProviderName(v string) *CreateIdentityProviderRequest
	GetIdentityProviderName() *string
	SetIdentityProviderType(v string) *CreateIdentityProviderRequest
	GetIdentityProviderType() *string
	SetInstanceId(v string) *CreateIdentityProviderRequest
	GetInstanceId() *string
	SetLarkConfig(v *CreateIdentityProviderRequestLarkConfig) *CreateIdentityProviderRequest
	GetLarkConfig() *CreateIdentityProviderRequestLarkConfig
	SetLdapConfig(v *CreateIdentityProviderRequestLdapConfig) *CreateIdentityProviderRequest
	GetLdapConfig() *CreateIdentityProviderRequestLdapConfig
	SetLogoUrl(v string) *CreateIdentityProviderRequest
	GetLogoUrl() *string
	SetNetworkAccessEndpointId(v string) *CreateIdentityProviderRequest
	GetNetworkAccessEndpointId() *string
	SetOidcConfig(v *CreateIdentityProviderRequestOidcConfig) *CreateIdentityProviderRequest
	GetOidcConfig() *CreateIdentityProviderRequestOidcConfig
	SetUdPullConfig(v *CreateIdentityProviderRequestUdPullConfig) *CreateIdentityProviderRequest
	GetUdPullConfig() *CreateIdentityProviderRequestUdPullConfig
	SetUdPushConfig(v *CreateIdentityProviderRequestUdPushConfig) *CreateIdentityProviderRequest
	GetUdPushConfig() *CreateIdentityProviderRequestUdPushConfig
	SetWeComConfig(v *CreateIdentityProviderRequestWeComConfig) *CreateIdentityProviderRequest
	GetWeComConfig() *CreateIdentityProviderRequestWeComConfig
}

type CreateIdentityProviderRequest struct {
	// Authentication configuration information.
	AuthnConfig *CreateIdentityProviderRequestAuthnConfig `json:"AuthnConfig,omitempty" xml:"AuthnConfig,omitempty" type:"Struct"`
	// Auto-create account rule configuration.
	AutoCreateUserConfig *CreateIdentityProviderRequestAutoCreateUserConfig `json:"AutoCreateUserConfig,omitempty" xml:"AutoCreateUserConfig,omitempty" type:"Struct"`
	// Auto-update account rule configuration.
	AutoUpdateUserConfig *CreateIdentityProviderRequestAutoUpdateUserConfig `json:"AutoUpdateUserConfig,omitempty" xml:"AutoUpdateUserConfig,omitempty" type:"Struct"`
	// OIDC identity provider account binding rule configuration.
	BindingConfig *CreateIdentityProviderRequestBindingConfig `json:"BindingConfig,omitempty" xml:"BindingConfig,omitempty" type:"Struct"`
	// Idp client token.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// DingTalk configuration information.
	DingtalkAppConfig *CreateIdentityProviderRequestDingtalkAppConfig `json:"DingtalkAppConfig,omitempty" xml:"DingtalkAppConfig,omitempty" type:"Struct"`
	// Identity provider name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	IdentityProviderName *string `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
	// Identity provider synchronization type.
	//
	// - Inbound to DingTalk: urn:alibaba:idaas:idp:alibaba:dingtalk:pull
	//
	// - Outbound to DingTalk: urn:alibaba:idaas:idp:alibaba:dingtalk:push
	//
	// - Inbound to WeCom: urn:alibaba:idaas:idp:tencent:wecom:pull
	//
	// - Inbound to Lark: urn:alibaba:idaas:idp:bytedance:lark:pull
	//
	// - Inbound to AD: urn:alibaba:idaas:idp:microsoft:ad:pull
	//
	// - Inbound to LDAP: urn:alibaba:idaas:idp:unknown:ldap:pull
	//
	// - Standard OIDC: urn:alibaba:idaas:idp:standard:oidc
	//
	// - SASE Custom OIDC: urn:alibaba:idaas:idp:alibaba:sase
	//
	// This parameter is required.
	//
	// example:
	//
	// urn:alibaba:idaas:idp:alibaba:dingtalk:push
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Lark (Feishu) configuration information.
	LarkConfig *CreateIdentityProviderRequestLarkConfig `json:"LarkConfig,omitempty" xml:"LarkConfig,omitempty" type:"Struct"`
	// AD/LDAP configuration information.
	LdapConfig *CreateIdentityProviderRequestLdapConfig `json:"LdapConfig,omitempty" xml:"LdapConfig,omitempty" type:"Struct"`
	// IdP logo url.
	//
	// example:
	//
	// xxxx-image://idaas_23aqr2ye554csg33dqpch5eu3q/tmp/d17d9adc-a943-45e7-ba0c-2838dddexxxxx
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// The unique identifier of the network access endpoint.
	//
	// example:
	//
	// nae_examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// OIDC IdP configuration.
	OidcConfig *CreateIdentityProviderRequestOidcConfig `json:"OidcConfig,omitempty" xml:"OidcConfig,omitempty" type:"Struct"`
	// Inbound synchronization configuration information.
	UdPullConfig *CreateIdentityProviderRequestUdPullConfig `json:"UdPullConfig,omitempty" xml:"UdPullConfig,omitempty" type:"Struct"`
	// Outbound synchronization configuration information.
	UdPushConfig *CreateIdentityProviderRequestUdPushConfig `json:"UdPushConfig,omitempty" xml:"UdPushConfig,omitempty" type:"Struct"`
	// WeCom configuration information.
	WeComConfig *CreateIdentityProviderRequestWeComConfig `json:"WeComConfig,omitempty" xml:"WeComConfig,omitempty" type:"Struct"`
}

func (s CreateIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequest) GetAuthnConfig() *CreateIdentityProviderRequestAuthnConfig {
	return s.AuthnConfig
}

func (s *CreateIdentityProviderRequest) GetAutoCreateUserConfig() *CreateIdentityProviderRequestAutoCreateUserConfig {
	return s.AutoCreateUserConfig
}

func (s *CreateIdentityProviderRequest) GetAutoUpdateUserConfig() *CreateIdentityProviderRequestAutoUpdateUserConfig {
	return s.AutoUpdateUserConfig
}

func (s *CreateIdentityProviderRequest) GetBindingConfig() *CreateIdentityProviderRequestBindingConfig {
	return s.BindingConfig
}

func (s *CreateIdentityProviderRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateIdentityProviderRequest) GetDingtalkAppConfig() *CreateIdentityProviderRequestDingtalkAppConfig {
	return s.DingtalkAppConfig
}

func (s *CreateIdentityProviderRequest) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *CreateIdentityProviderRequest) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *CreateIdentityProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateIdentityProviderRequest) GetLarkConfig() *CreateIdentityProviderRequestLarkConfig {
	return s.LarkConfig
}

func (s *CreateIdentityProviderRequest) GetLdapConfig() *CreateIdentityProviderRequestLdapConfig {
	return s.LdapConfig
}

func (s *CreateIdentityProviderRequest) GetLogoUrl() *string {
	return s.LogoUrl
}

func (s *CreateIdentityProviderRequest) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *CreateIdentityProviderRequest) GetOidcConfig() *CreateIdentityProviderRequestOidcConfig {
	return s.OidcConfig
}

func (s *CreateIdentityProviderRequest) GetUdPullConfig() *CreateIdentityProviderRequestUdPullConfig {
	return s.UdPullConfig
}

func (s *CreateIdentityProviderRequest) GetUdPushConfig() *CreateIdentityProviderRequestUdPushConfig {
	return s.UdPushConfig
}

func (s *CreateIdentityProviderRequest) GetWeComConfig() *CreateIdentityProviderRequestWeComConfig {
	return s.WeComConfig
}

func (s *CreateIdentityProviderRequest) SetAuthnConfig(v *CreateIdentityProviderRequestAuthnConfig) *CreateIdentityProviderRequest {
	s.AuthnConfig = v
	return s
}

func (s *CreateIdentityProviderRequest) SetAutoCreateUserConfig(v *CreateIdentityProviderRequestAutoCreateUserConfig) *CreateIdentityProviderRequest {
	s.AutoCreateUserConfig = v
	return s
}

func (s *CreateIdentityProviderRequest) SetAutoUpdateUserConfig(v *CreateIdentityProviderRequestAutoUpdateUserConfig) *CreateIdentityProviderRequest {
	s.AutoUpdateUserConfig = v
	return s
}

func (s *CreateIdentityProviderRequest) SetBindingConfig(v *CreateIdentityProviderRequestBindingConfig) *CreateIdentityProviderRequest {
	s.BindingConfig = v
	return s
}

func (s *CreateIdentityProviderRequest) SetClientToken(v string) *CreateIdentityProviderRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIdentityProviderRequest) SetDingtalkAppConfig(v *CreateIdentityProviderRequestDingtalkAppConfig) *CreateIdentityProviderRequest {
	s.DingtalkAppConfig = v
	return s
}

func (s *CreateIdentityProviderRequest) SetIdentityProviderName(v string) *CreateIdentityProviderRequest {
	s.IdentityProviderName = &v
	return s
}

func (s *CreateIdentityProviderRequest) SetIdentityProviderType(v string) *CreateIdentityProviderRequest {
	s.IdentityProviderType = &v
	return s
}

func (s *CreateIdentityProviderRequest) SetInstanceId(v string) *CreateIdentityProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateIdentityProviderRequest) SetLarkConfig(v *CreateIdentityProviderRequestLarkConfig) *CreateIdentityProviderRequest {
	s.LarkConfig = v
	return s
}

func (s *CreateIdentityProviderRequest) SetLdapConfig(v *CreateIdentityProviderRequestLdapConfig) *CreateIdentityProviderRequest {
	s.LdapConfig = v
	return s
}

func (s *CreateIdentityProviderRequest) SetLogoUrl(v string) *CreateIdentityProviderRequest {
	s.LogoUrl = &v
	return s
}

func (s *CreateIdentityProviderRequest) SetNetworkAccessEndpointId(v string) *CreateIdentityProviderRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *CreateIdentityProviderRequest) SetOidcConfig(v *CreateIdentityProviderRequestOidcConfig) *CreateIdentityProviderRequest {
	s.OidcConfig = v
	return s
}

func (s *CreateIdentityProviderRequest) SetUdPullConfig(v *CreateIdentityProviderRequestUdPullConfig) *CreateIdentityProviderRequest {
	s.UdPullConfig = v
	return s
}

func (s *CreateIdentityProviderRequest) SetUdPushConfig(v *CreateIdentityProviderRequestUdPushConfig) *CreateIdentityProviderRequest {
	s.UdPushConfig = v
	return s
}

func (s *CreateIdentityProviderRequest) SetWeComConfig(v *CreateIdentityProviderRequestWeComConfig) *CreateIdentityProviderRequest {
	s.WeComConfig = v
	return s
}

func (s *CreateIdentityProviderRequest) Validate() error {
	if s.AuthnConfig != nil {
		if err := s.AuthnConfig.Validate(); err != nil {
			return err
		}
	}
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
	if s.DingtalkAppConfig != nil {
		if err := s.DingtalkAppConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LarkConfig != nil {
		if err := s.LarkConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LdapConfig != nil {
		if err := s.LdapConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OidcConfig != nil {
		if err := s.OidcConfig.Validate(); err != nil {
			return err
		}
	}
	if s.UdPullConfig != nil {
		if err := s.UdPullConfig.Validate(); err != nil {
			return err
		}
	}
	if s.UdPushConfig != nil {
		if err := s.UdPushConfig.Validate(); err != nil {
			return err
		}
	}
	if s.WeComConfig != nil {
		if err := s.WeComConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIdentityProviderRequestAuthnConfig struct {
	// Whether the corresponding IdP supports authentication. Value range:
	//
	// - Disabled: disabled
	//
	// - Enabled: enabled
	//
	// example:
	//
	// enabled
	AuthnStatus *string `json:"AuthnStatus,omitempty" xml:"AuthnStatus,omitempty"`
	// Whether automatic password update is supported. Value range:
	//
	// - Disabled: disabled
	//
	// - Enabled: enabled
	//
	// example:
	//
	// enabled
	AutoUpdatePasswordStatus *string `json:"AutoUpdatePasswordStatus,omitempty" xml:"AutoUpdatePasswordStatus,omitempty"`
}

func (s CreateIdentityProviderRequestAuthnConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestAuthnConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestAuthnConfig) GetAuthnStatus() *string {
	return s.AuthnStatus
}

func (s *CreateIdentityProviderRequestAuthnConfig) GetAutoUpdatePasswordStatus() *string {
	return s.AutoUpdatePasswordStatus
}

func (s *CreateIdentityProviderRequestAuthnConfig) SetAuthnStatus(v string) *CreateIdentityProviderRequestAuthnConfig {
	s.AuthnStatus = &v
	return s
}

func (s *CreateIdentityProviderRequestAuthnConfig) SetAutoUpdatePasswordStatus(v string) *CreateIdentityProviderRequestAuthnConfig {
	s.AutoUpdatePasswordStatus = &v
	return s
}

func (s *CreateIdentityProviderRequestAuthnConfig) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderRequestAutoCreateUserConfig struct {
	// Whether auto-creation of accounts is enabled. Possible values:
	//
	// - Disabled: disabled
	//
	// - Enabled: enabled
	//
	// example:
	//
	// disabled
	AutoCreateUserStatus *string `json:"AutoCreateUserStatus,omitempty" xml:"AutoCreateUserStatus,omitempty"`
	// Target organizational unit IDs collection.
	TargetOrganizationalUnitIds []*string `json:"TargetOrganizationalUnitIds,omitempty" xml:"TargetOrganizationalUnitIds,omitempty" type:"Repeated"`
}

func (s CreateIdentityProviderRequestAutoCreateUserConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestAutoCreateUserConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestAutoCreateUserConfig) GetAutoCreateUserStatus() *string {
	return s.AutoCreateUserStatus
}

func (s *CreateIdentityProviderRequestAutoCreateUserConfig) GetTargetOrganizationalUnitIds() []*string {
	return s.TargetOrganizationalUnitIds
}

func (s *CreateIdentityProviderRequestAutoCreateUserConfig) SetAutoCreateUserStatus(v string) *CreateIdentityProviderRequestAutoCreateUserConfig {
	s.AutoCreateUserStatus = &v
	return s
}

func (s *CreateIdentityProviderRequestAutoCreateUserConfig) SetTargetOrganizationalUnitIds(v []*string) *CreateIdentityProviderRequestAutoCreateUserConfig {
	s.TargetOrganizationalUnitIds = v
	return s
}

func (s *CreateIdentityProviderRequestAutoCreateUserConfig) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderRequestAutoUpdateUserConfig struct {
	// Whether auto-updating of accounts is enabled. Possible values:
	//
	// - Disabled: disabled
	//
	// - Enabled: enabled
	//
	// example:
	//
	// disabled
	AutoUpdateUserStatus *string `json:"AutoUpdateUserStatus,omitempty" xml:"AutoUpdateUserStatus,omitempty"`
}

func (s CreateIdentityProviderRequestAutoUpdateUserConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestAutoUpdateUserConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestAutoUpdateUserConfig) GetAutoUpdateUserStatus() *string {
	return s.AutoUpdateUserStatus
}

func (s *CreateIdentityProviderRequestAutoUpdateUserConfig) SetAutoUpdateUserStatus(v string) *CreateIdentityProviderRequestAutoUpdateUserConfig {
	s.AutoUpdateUserStatus = &v
	return s
}

func (s *CreateIdentityProviderRequestAutoUpdateUserConfig) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderRequestBindingConfig struct {
	// List of rules for automatically matching accounts.
	AutoMatchUserProfileExpressions []*CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions `json:"AutoMatchUserProfileExpressions,omitempty" xml:"AutoMatchUserProfileExpressions,omitempty" type:"Repeated"`
	// Whether automatic account matching is enabled. Value range:
	//
	// - Disabled: disabled
	//
	// - Enabled: enabled
	//
	// example:
	//
	// disabled
	AutoMatchUserStatus *string `json:"AutoMatchUserStatus,omitempty" xml:"AutoMatchUserStatus,omitempty"`
	// Whether the user manual account binding function is enabled. Value range:
	//
	// - Disabled: disabled
	//
	// - Enabled: enabled
	//
	// example:
	//
	// enabled
	MappingBindingStatus *string `json:"MappingBindingStatus,omitempty" xml:"MappingBindingStatus,omitempty"`
}

func (s CreateIdentityProviderRequestBindingConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestBindingConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestBindingConfig) GetAutoMatchUserProfileExpressions() []*CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions {
	return s.AutoMatchUserProfileExpressions
}

func (s *CreateIdentityProviderRequestBindingConfig) GetAutoMatchUserStatus() *string {
	return s.AutoMatchUserStatus
}

func (s *CreateIdentityProviderRequestBindingConfig) GetMappingBindingStatus() *string {
	return s.MappingBindingStatus
}

func (s *CreateIdentityProviderRequestBindingConfig) SetAutoMatchUserProfileExpressions(v []*CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions) *CreateIdentityProviderRequestBindingConfig {
	s.AutoMatchUserProfileExpressions = v
	return s
}

func (s *CreateIdentityProviderRequestBindingConfig) SetAutoMatchUserStatus(v string) *CreateIdentityProviderRequestBindingConfig {
	s.AutoMatchUserStatus = &v
	return s
}

func (s *CreateIdentityProviderRequestBindingConfig) SetMappingBindingStatus(v string) *CreateIdentityProviderRequestBindingConfig {
	s.MappingBindingStatus = &v
	return s
}

func (s *CreateIdentityProviderRequestBindingConfig) Validate() error {
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

type CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions struct {
	// Type of the expression. Value range:
	//
	// - Field: filed
	//
	// - Expression: expression
	//
	// example:
	//
	// filed
	ExpressionMappingType *string `json:"ExpressionMappingType,omitempty" xml:"ExpressionMappingType,omitempty"`
	// Expression for the mapped attribute value.
	//
	// example:
	//
	// idpUser.phoneNumber
	SourceValueExpression *string `json:"SourceValueExpression,omitempty" xml:"SourceValueExpression,omitempty"`
	// Name of the target attribute.
	//
	// example:
	//
	// user.username
	TargetField *string `json:"TargetField,omitempty" xml:"TargetField,omitempty"`
	// Description of the target attribute.
	//
	// example:
	//
	// user.username
	TargetFieldDescription *string `json:"TargetFieldDescription,omitempty" xml:"TargetFieldDescription,omitempty"`
}

func (s CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions) GetExpressionMappingType() *string {
	return s.ExpressionMappingType
}

func (s *CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions) GetSourceValueExpression() *string {
	return s.SourceValueExpression
}

func (s *CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions) GetTargetField() *string {
	return s.TargetField
}

func (s *CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions) GetTargetFieldDescription() *string {
	return s.TargetFieldDescription
}

func (s *CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions) SetExpressionMappingType(v string) *CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions {
	s.ExpressionMappingType = &v
	return s
}

func (s *CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions) SetSourceValueExpression(v string) *CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions {
	s.SourceValueExpression = &v
	return s
}

func (s *CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions) SetTargetField(v string) *CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions {
	s.TargetField = &v
	return s
}

func (s *CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions) SetTargetFieldDescription(v string) *CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions {
	s.TargetFieldDescription = &v
	return s
}

func (s *CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderRequestDingtalkAppConfig struct {
	// AppKey of the DingTalk application.
	//
	// example:
	//
	// Xczngvfemo4e
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// AppSecret of the DingTalk application.
	//
	// example:
	//
	// 5d405a12a6f84ad4ab05ee09axxxx
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// CorpId of the DingTalk application.
	//
	// example:
	//
	// 3075680424786133505
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// DingTalk edition. Valid values:
	//
	// public_dingtalk – Standard DingTalk.
	//
	// private_dingtalk – Dedicated DingTalk.
	//
	// example:
	//
	// public_dingtalk
	DingtalkVersion *string `json:"DingtalkVersion,omitempty" xml:"DingtalkVersion,omitempty"`
	// DingTalk encrypt key.
	//
	// example:
	//
	// 29003eb11d0a28b4802a6f02fb8aa25dff730e2ac26ffd200dxxxx
	EncryptKey *string `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
	// DingTalk verification token.
	//
	// example:
	//
	// 5ba9c127a7abe029003eb11d0a28b4802a6f02fb8aa25dff730e2ac26ffd200dxxxx
	VerificationToken *string `json:"VerificationToken,omitempty" xml:"VerificationToken,omitempty"`
}

func (s CreateIdentityProviderRequestDingtalkAppConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestDingtalkAppConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestDingtalkAppConfig) GetAppKey() *string {
	return s.AppKey
}

func (s *CreateIdentityProviderRequestDingtalkAppConfig) GetAppSecret() *string {
	return s.AppSecret
}

func (s *CreateIdentityProviderRequestDingtalkAppConfig) GetCorpId() *string {
	return s.CorpId
}

func (s *CreateIdentityProviderRequestDingtalkAppConfig) GetDingtalkVersion() *string {
	return s.DingtalkVersion
}

func (s *CreateIdentityProviderRequestDingtalkAppConfig) GetEncryptKey() *string {
	return s.EncryptKey
}

func (s *CreateIdentityProviderRequestDingtalkAppConfig) GetVerificationToken() *string {
	return s.VerificationToken
}

func (s *CreateIdentityProviderRequestDingtalkAppConfig) SetAppKey(v string) *CreateIdentityProviderRequestDingtalkAppConfig {
	s.AppKey = &v
	return s
}

func (s *CreateIdentityProviderRequestDingtalkAppConfig) SetAppSecret(v string) *CreateIdentityProviderRequestDingtalkAppConfig {
	s.AppSecret = &v
	return s
}

func (s *CreateIdentityProviderRequestDingtalkAppConfig) SetCorpId(v string) *CreateIdentityProviderRequestDingtalkAppConfig {
	s.CorpId = &v
	return s
}

func (s *CreateIdentityProviderRequestDingtalkAppConfig) SetDingtalkVersion(v string) *CreateIdentityProviderRequestDingtalkAppConfig {
	s.DingtalkVersion = &v
	return s
}

func (s *CreateIdentityProviderRequestDingtalkAppConfig) SetEncryptKey(v string) *CreateIdentityProviderRequestDingtalkAppConfig {
	s.EncryptKey = &v
	return s
}

func (s *CreateIdentityProviderRequestDingtalkAppConfig) SetVerificationToken(v string) *CreateIdentityProviderRequestDingtalkAppConfig {
	s.VerificationToken = &v
	return s
}

func (s *CreateIdentityProviderRequestDingtalkAppConfig) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderRequestLarkConfig struct {
	// Lark (Feishu) app appId.
	//
	// example:
	//
	// cli_xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Lark (Feishu) app secret.
	//
	// example:
	//
	// KiiLzh5Dueh4wbLxxxx
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// Lark (Feishu) encrypt key.
	//
	// example:
	//
	// 29003eb11d0a28b4802a6f02fb8aa25dff730e2ac26ffd200dxxxx
	EncryptKey *string `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
	// Lark (Feishu) enterprise number.
	//
	// example:
	//
	// FSX123111xxx
	EnterpriseNumber *string `json:"EnterpriseNumber,omitempty" xml:"EnterpriseNumber,omitempty"`
	// Lark (Feishu)  verification token.
	//
	// example:
	//
	// 5ba9c127a7abe029003eb11d0a28b4802a6f02fb8aa25dff730e2ac26ffd200dxxxx
	VerificationToken *string `json:"VerificationToken,omitempty" xml:"VerificationToken,omitempty"`
}

func (s CreateIdentityProviderRequestLarkConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestLarkConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestLarkConfig) GetAppId() *string {
	return s.AppId
}

func (s *CreateIdentityProviderRequestLarkConfig) GetAppSecret() *string {
	return s.AppSecret
}

func (s *CreateIdentityProviderRequestLarkConfig) GetEncryptKey() *string {
	return s.EncryptKey
}

func (s *CreateIdentityProviderRequestLarkConfig) GetEnterpriseNumber() *string {
	return s.EnterpriseNumber
}

func (s *CreateIdentityProviderRequestLarkConfig) GetVerificationToken() *string {
	return s.VerificationToken
}

func (s *CreateIdentityProviderRequestLarkConfig) SetAppId(v string) *CreateIdentityProviderRequestLarkConfig {
	s.AppId = &v
	return s
}

func (s *CreateIdentityProviderRequestLarkConfig) SetAppSecret(v string) *CreateIdentityProviderRequestLarkConfig {
	s.AppSecret = &v
	return s
}

func (s *CreateIdentityProviderRequestLarkConfig) SetEncryptKey(v string) *CreateIdentityProviderRequestLarkConfig {
	s.EncryptKey = &v
	return s
}

func (s *CreateIdentityProviderRequestLarkConfig) SetEnterpriseNumber(v string) *CreateIdentityProviderRequestLarkConfig {
	s.EnterpriseNumber = &v
	return s
}

func (s *CreateIdentityProviderRequestLarkConfig) SetVerificationToken(v string) *CreateIdentityProviderRequestLarkConfig {
	s.VerificationToken = &v
	return s
}

func (s *CreateIdentityProviderRequestLarkConfig) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderRequestLdapConfig struct {
	// Administrator password.
	//
	// example:
	//
	// xxxx
	AdministratorPassword *string `json:"AdministratorPassword,omitempty" xml:"AdministratorPassword,omitempty"`
	// Administrator username.
	//
	// example:
	//
	// DC=example,DC=com
	AdministratorUsername *string `json:"AdministratorUsername,omitempty" xml:"AdministratorUsername,omitempty"`
	// Whether to verify the certificate fingerprint. Value range:
	//
	// - Disabled: disabled
	//
	// - Enabled: enabled
	//
	// example:
	//
	// enabled
	CertificateFingerprintStatus *string `json:"CertificateFingerprintStatus,omitempty" xml:"CertificateFingerprintStatus,omitempty"`
	// List of certificate fingerprints.
	CertificateFingerprints []*string `json:"CertificateFingerprints,omitempty" xml:"CertificateFingerprints,omitempty" type:"Repeated"`
	// Group member attribute name.
	//
	// example:
	//
	// member
	GroupMemberAttributeName *string `json:"GroupMemberAttributeName,omitempty" xml:"GroupMemberAttributeName,omitempty"`
	// Group ObjectClass.
	//
	// example:
	//
	// group
	GroupObjectClass *string `json:"GroupObjectClass,omitempty" xml:"GroupObjectClass,omitempty"`
	// Custom filter for Group ObjectClass.
	//
	// example:
	//
	// (|(cn=test)(group=test@test.com))
	GroupObjectClassCustomFilter *string `json:"GroupObjectClassCustomFilter,omitempty" xml:"GroupObjectClassCustomFilter,omitempty"`
	// Communication protocol.
	//
	// example:
	//
	// ldap
	LdapProtocol *string `json:"LdapProtocol,omitempty" xml:"LdapProtocol,omitempty"`
	// AD/LDAP server address.
	//
	// example:
	//
	// 123.xx.xx.89
	LdapServerHost *string `json:"LdapServerHost,omitempty" xml:"LdapServerHost,omitempty"`
	// AD/LDAP port number.
	//
	// example:
	//
	// 636
	LdapServerPort *int32 `json:"LdapServerPort,omitempty" xml:"LdapServerPort,omitempty"`
	// Organization Unit ObjectClass.
	//
	// example:
	//
	// organizationUnit,top
	OrganizationUnitObjectClass *string `json:"OrganizationUnitObjectClass,omitempty" xml:"OrganizationUnitObjectClass,omitempty"`
	// example:
	//
	// ou
	OrganizationalUnitRdn *string `json:"OrganizationalUnitRdn,omitempty" xml:"OrganizationalUnitRdn,omitempty"`
	// example:
	//
	// enabled
	PasswordSyncStatus *string `json:"PasswordSyncStatus,omitempty" xml:"PasswordSyncStatus,omitempty"`
	// Whether startTLS is enabled. Value range:
	//
	// - Disabled: disabled
	//
	// - Enabled: enabled
	//
	// example:
	//
	// enabled
	StartTlsStatus *string `json:"StartTlsStatus,omitempty" xml:"StartTlsStatus,omitempty"`
	// User login identifier.
	//
	// example:
	//
	// userPrincipalName, mail
	UserLoginIdentifier *string `json:"UserLoginIdentifier,omitempty" xml:"UserLoginIdentifier,omitempty"`
	// User ObjectClass.
	//
	// example:
	//
	// person,user
	UserObjectClass *string `json:"UserObjectClass,omitempty" xml:"UserObjectClass,omitempty"`
	// Custom filter for User ObjectClass.
	//
	// example:
	//
	// (|(cn=test)(mail=test@test.com))
	UserObjectClassCustomFilter *string `json:"UserObjectClassCustomFilter,omitempty" xml:"UserObjectClassCustomFilter,omitempty"`
	// example:
	//
	// cn
	UserRdn *string `json:"UserRdn,omitempty" xml:"UserRdn,omitempty"`
}

func (s CreateIdentityProviderRequestLdapConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestLdapConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestLdapConfig) GetAdministratorPassword() *string {
	return s.AdministratorPassword
}

func (s *CreateIdentityProviderRequestLdapConfig) GetAdministratorUsername() *string {
	return s.AdministratorUsername
}

func (s *CreateIdentityProviderRequestLdapConfig) GetCertificateFingerprintStatus() *string {
	return s.CertificateFingerprintStatus
}

func (s *CreateIdentityProviderRequestLdapConfig) GetCertificateFingerprints() []*string {
	return s.CertificateFingerprints
}

func (s *CreateIdentityProviderRequestLdapConfig) GetGroupMemberAttributeName() *string {
	return s.GroupMemberAttributeName
}

func (s *CreateIdentityProviderRequestLdapConfig) GetGroupObjectClass() *string {
	return s.GroupObjectClass
}

func (s *CreateIdentityProviderRequestLdapConfig) GetGroupObjectClassCustomFilter() *string {
	return s.GroupObjectClassCustomFilter
}

func (s *CreateIdentityProviderRequestLdapConfig) GetLdapProtocol() *string {
	return s.LdapProtocol
}

func (s *CreateIdentityProviderRequestLdapConfig) GetLdapServerHost() *string {
	return s.LdapServerHost
}

func (s *CreateIdentityProviderRequestLdapConfig) GetLdapServerPort() *int32 {
	return s.LdapServerPort
}

func (s *CreateIdentityProviderRequestLdapConfig) GetOrganizationUnitObjectClass() *string {
	return s.OrganizationUnitObjectClass
}

func (s *CreateIdentityProviderRequestLdapConfig) GetOrganizationalUnitRdn() *string {
	return s.OrganizationalUnitRdn
}

func (s *CreateIdentityProviderRequestLdapConfig) GetPasswordSyncStatus() *string {
	return s.PasswordSyncStatus
}

func (s *CreateIdentityProviderRequestLdapConfig) GetStartTlsStatus() *string {
	return s.StartTlsStatus
}

func (s *CreateIdentityProviderRequestLdapConfig) GetUserLoginIdentifier() *string {
	return s.UserLoginIdentifier
}

func (s *CreateIdentityProviderRequestLdapConfig) GetUserObjectClass() *string {
	return s.UserObjectClass
}

func (s *CreateIdentityProviderRequestLdapConfig) GetUserObjectClassCustomFilter() *string {
	return s.UserObjectClassCustomFilter
}

func (s *CreateIdentityProviderRequestLdapConfig) GetUserRdn() *string {
	return s.UserRdn
}

func (s *CreateIdentityProviderRequestLdapConfig) SetAdministratorPassword(v string) *CreateIdentityProviderRequestLdapConfig {
	s.AdministratorPassword = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetAdministratorUsername(v string) *CreateIdentityProviderRequestLdapConfig {
	s.AdministratorUsername = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetCertificateFingerprintStatus(v string) *CreateIdentityProviderRequestLdapConfig {
	s.CertificateFingerprintStatus = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetCertificateFingerprints(v []*string) *CreateIdentityProviderRequestLdapConfig {
	s.CertificateFingerprints = v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetGroupMemberAttributeName(v string) *CreateIdentityProviderRequestLdapConfig {
	s.GroupMemberAttributeName = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetGroupObjectClass(v string) *CreateIdentityProviderRequestLdapConfig {
	s.GroupObjectClass = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetGroupObjectClassCustomFilter(v string) *CreateIdentityProviderRequestLdapConfig {
	s.GroupObjectClassCustomFilter = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetLdapProtocol(v string) *CreateIdentityProviderRequestLdapConfig {
	s.LdapProtocol = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetLdapServerHost(v string) *CreateIdentityProviderRequestLdapConfig {
	s.LdapServerHost = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetLdapServerPort(v int32) *CreateIdentityProviderRequestLdapConfig {
	s.LdapServerPort = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetOrganizationUnitObjectClass(v string) *CreateIdentityProviderRequestLdapConfig {
	s.OrganizationUnitObjectClass = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetOrganizationalUnitRdn(v string) *CreateIdentityProviderRequestLdapConfig {
	s.OrganizationalUnitRdn = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetPasswordSyncStatus(v string) *CreateIdentityProviderRequestLdapConfig {
	s.PasswordSyncStatus = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetStartTlsStatus(v string) *CreateIdentityProviderRequestLdapConfig {
	s.StartTlsStatus = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetUserLoginIdentifier(v string) *CreateIdentityProviderRequestLdapConfig {
	s.UserLoginIdentifier = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetUserObjectClass(v string) *CreateIdentityProviderRequestLdapConfig {
	s.UserObjectClass = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetUserObjectClassCustomFilter(v string) *CreateIdentityProviderRequestLdapConfig {
	s.UserObjectClassCustomFilter = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) SetUserRdn(v string) *CreateIdentityProviderRequestLdapConfig {
	s.UserRdn = &v
	return s
}

func (s *CreateIdentityProviderRequestLdapConfig) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderRequestOidcConfig struct {
	// OIDC client authentication configuration.
	AuthnParam *CreateIdentityProviderRequestOidcConfigAuthnParam `json:"AuthnParam,omitempty" xml:"AuthnParam,omitempty" type:"Struct"`
	// OIDC endpoint configuration.
	EndpointConfig *CreateIdentityProviderRequestOidcConfigEndpointConfig `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty" type:"Struct"`
	// OIDC grant scopes collection.
	//
	// example:
	//
	// openid
	GrantScopes []*string `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
	// OIDC grant type.
	//
	// example:
	//
	// authorization_code
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
	// PKCE algorithm. Possible values:
	//
	// - SHA256: S256
	//
	// - Plain text: plain
	//
	// example:
	//
	// S256
	PkceChallengeMethod *string `json:"PkceChallengeMethod,omitempty" xml:"PkceChallengeMethod,omitempty"`
	// Whether to use PKCE in the AuthorizationCode grant mode.
	//
	// example:
	//
	// true
	PkceRequired *bool `json:"PkceRequired,omitempty" xml:"PkceRequired,omitempty"`
}

func (s CreateIdentityProviderRequestOidcConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestOidcConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestOidcConfig) GetAuthnParam() *CreateIdentityProviderRequestOidcConfigAuthnParam {
	return s.AuthnParam
}

func (s *CreateIdentityProviderRequestOidcConfig) GetEndpointConfig() *CreateIdentityProviderRequestOidcConfigEndpointConfig {
	return s.EndpointConfig
}

func (s *CreateIdentityProviderRequestOidcConfig) GetGrantScopes() []*string {
	return s.GrantScopes
}

func (s *CreateIdentityProviderRequestOidcConfig) GetGrantType() *string {
	return s.GrantType
}

func (s *CreateIdentityProviderRequestOidcConfig) GetPkceChallengeMethod() *string {
	return s.PkceChallengeMethod
}

func (s *CreateIdentityProviderRequestOidcConfig) GetPkceRequired() *bool {
	return s.PkceRequired
}

func (s *CreateIdentityProviderRequestOidcConfig) SetAuthnParam(v *CreateIdentityProviderRequestOidcConfigAuthnParam) *CreateIdentityProviderRequestOidcConfig {
	s.AuthnParam = v
	return s
}

func (s *CreateIdentityProviderRequestOidcConfig) SetEndpointConfig(v *CreateIdentityProviderRequestOidcConfigEndpointConfig) *CreateIdentityProviderRequestOidcConfig {
	s.EndpointConfig = v
	return s
}

func (s *CreateIdentityProviderRequestOidcConfig) SetGrantScopes(v []*string) *CreateIdentityProviderRequestOidcConfig {
	s.GrantScopes = v
	return s
}

func (s *CreateIdentityProviderRequestOidcConfig) SetGrantType(v string) *CreateIdentityProviderRequestOidcConfig {
	s.GrantType = &v
	return s
}

func (s *CreateIdentityProviderRequestOidcConfig) SetPkceChallengeMethod(v string) *CreateIdentityProviderRequestOidcConfig {
	s.PkceChallengeMethod = &v
	return s
}

func (s *CreateIdentityProviderRequestOidcConfig) SetPkceRequired(v bool) *CreateIdentityProviderRequestOidcConfig {
	s.PkceRequired = &v
	return s
}

func (s *CreateIdentityProviderRequestOidcConfig) Validate() error {
	if s.AuthnParam != nil {
		if err := s.AuthnParam.Validate(); err != nil {
			return err
		}
	}
	if s.EndpointConfig != nil {
		if err := s.EndpointConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIdentityProviderRequestOidcConfigAuthnParam struct {
	// OIDC authentication method. Value range:
	//
	// - client_secret_basic
	//
	// - client_secret_post
	//
	// example:
	//
	// client_secret_post
	AuthnMethod *string `json:"AuthnMethod,omitempty" xml:"AuthnMethod,omitempty"`
	// The ID of the client.
	//
	// example:
	//
	// mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The  secret of the client.
	//
	// example:
	//
	// CSEHDddddddxxxxuxkJEHPveWRXBGqVqRsxxxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s CreateIdentityProviderRequestOidcConfigAuthnParam) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestOidcConfigAuthnParam) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestOidcConfigAuthnParam) GetAuthnMethod() *string {
	return s.AuthnMethod
}

func (s *CreateIdentityProviderRequestOidcConfigAuthnParam) GetClientId() *string {
	return s.ClientId
}

func (s *CreateIdentityProviderRequestOidcConfigAuthnParam) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *CreateIdentityProviderRequestOidcConfigAuthnParam) SetAuthnMethod(v string) *CreateIdentityProviderRequestOidcConfigAuthnParam {
	s.AuthnMethod = &v
	return s
}

func (s *CreateIdentityProviderRequestOidcConfigAuthnParam) SetClientId(v string) *CreateIdentityProviderRequestOidcConfigAuthnParam {
	s.ClientId = &v
	return s
}

func (s *CreateIdentityProviderRequestOidcConfigAuthnParam) SetClientSecret(v string) *CreateIdentityProviderRequestOidcConfigAuthnParam {
	s.ClientSecret = &v
	return s
}

func (s *CreateIdentityProviderRequestOidcConfigAuthnParam) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderRequestOidcConfigEndpointConfig struct {
	// OIDC authorization endpoint.
	//
	// example:
	//
	// https://example.com/auth/authorize
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// OIDC issuer information.
	//
	// example:
	//
	// https://example.com/auth
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// OIDC jwks uri.
	//
	// example:
	//
	// https://example.com/auth/jwks
	JwksUri *string `json:"JwksUri,omitempty" xml:"JwksUri,omitempty"`
	// OIDC token endpoint.
	//
	// example:
	//
	// https://example.com/auth/token
	TokenEndpoint *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
	// OIDC user info endpoint.
	//
	// example:
	//
	// https://example.com/auth/userinfo
	UserinfoEndpoint *string `json:"UserinfoEndpoint,omitempty" xml:"UserinfoEndpoint,omitempty"`
}

func (s CreateIdentityProviderRequestOidcConfigEndpointConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestOidcConfigEndpointConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestOidcConfigEndpointConfig) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *CreateIdentityProviderRequestOidcConfigEndpointConfig) GetIssuer() *string {
	return s.Issuer
}

func (s *CreateIdentityProviderRequestOidcConfigEndpointConfig) GetJwksUri() *string {
	return s.JwksUri
}

func (s *CreateIdentityProviderRequestOidcConfigEndpointConfig) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *CreateIdentityProviderRequestOidcConfigEndpointConfig) GetUserinfoEndpoint() *string {
	return s.UserinfoEndpoint
}

func (s *CreateIdentityProviderRequestOidcConfigEndpointConfig) SetAuthorizationEndpoint(v string) *CreateIdentityProviderRequestOidcConfigEndpointConfig {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *CreateIdentityProviderRequestOidcConfigEndpointConfig) SetIssuer(v string) *CreateIdentityProviderRequestOidcConfigEndpointConfig {
	s.Issuer = &v
	return s
}

func (s *CreateIdentityProviderRequestOidcConfigEndpointConfig) SetJwksUri(v string) *CreateIdentityProviderRequestOidcConfigEndpointConfig {
	s.JwksUri = &v
	return s
}

func (s *CreateIdentityProviderRequestOidcConfigEndpointConfig) SetTokenEndpoint(v string) *CreateIdentityProviderRequestOidcConfigEndpointConfig {
	s.TokenEndpoint = &v
	return s
}

func (s *CreateIdentityProviderRequestOidcConfigEndpointConfig) SetUserinfoEndpoint(v string) *CreateIdentityProviderRequestOidcConfigEndpointConfig {
	s.UserinfoEndpoint = &v
	return s
}

func (s *CreateIdentityProviderRequestOidcConfigEndpointConfig) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderRequestUdPullConfig struct {
	// Whether group synchronization is supported. The default value is disabled. Possible values:
	//
	// - Disabled: disabled
	//
	// - Enabled: enabled
	//
	// example:
	//
	// disabled
	GroupSyncStatus *string `json:"GroupSyncStatus,omitempty" xml:"GroupSyncStatus,omitempty"`
	// Incremental callback status, indicating whether to process incremental callback data from the IdP. Possible values:
	//
	// - Disabled: disabled
	//
	// - Enabled: enabled
	//
	// example:
	//
	// disabled
	IncrementalCallbackStatus *string `json:"IncrementalCallbackStatus,omitempty" xml:"IncrementalCallbackStatus,omitempty"`
	// Scheduled configuration verification.
	PeriodicSyncConfig *CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig `json:"PeriodicSyncConfig,omitempty" xml:"PeriodicSyncConfig,omitempty" type:"Struct"`
	// Periodic check status, indicating whether to periodically check the data differences between EIAM and the identity provider. Possible values:
	//
	// - Disabled: disabled
	//
	// - Enabled: enabled
	//
	// example:
	//
	// disabled
	PeriodicSyncStatus *string `json:"PeriodicSyncStatus,omitempty" xml:"PeriodicSyncStatus,omitempty"`
	// Synchronization scope configuration information.
	UdSyncScopeConfig *CreateIdentityProviderRequestUdPullConfigUdSyncScopeConfig `json:"UdSyncScopeConfig,omitempty" xml:"UdSyncScopeConfig,omitempty" type:"Struct"`
}

func (s CreateIdentityProviderRequestUdPullConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestUdPullConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestUdPullConfig) GetGroupSyncStatus() *string {
	return s.GroupSyncStatus
}

func (s *CreateIdentityProviderRequestUdPullConfig) GetIncrementalCallbackStatus() *string {
	return s.IncrementalCallbackStatus
}

func (s *CreateIdentityProviderRequestUdPullConfig) GetPeriodicSyncConfig() *CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig {
	return s.PeriodicSyncConfig
}

func (s *CreateIdentityProviderRequestUdPullConfig) GetPeriodicSyncStatus() *string {
	return s.PeriodicSyncStatus
}

func (s *CreateIdentityProviderRequestUdPullConfig) GetUdSyncScopeConfig() *CreateIdentityProviderRequestUdPullConfigUdSyncScopeConfig {
	return s.UdSyncScopeConfig
}

func (s *CreateIdentityProviderRequestUdPullConfig) SetGroupSyncStatus(v string) *CreateIdentityProviderRequestUdPullConfig {
	s.GroupSyncStatus = &v
	return s
}

func (s *CreateIdentityProviderRequestUdPullConfig) SetIncrementalCallbackStatus(v string) *CreateIdentityProviderRequestUdPullConfig {
	s.IncrementalCallbackStatus = &v
	return s
}

func (s *CreateIdentityProviderRequestUdPullConfig) SetPeriodicSyncConfig(v *CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig) *CreateIdentityProviderRequestUdPullConfig {
	s.PeriodicSyncConfig = v
	return s
}

func (s *CreateIdentityProviderRequestUdPullConfig) SetPeriodicSyncStatus(v string) *CreateIdentityProviderRequestUdPullConfig {
	s.PeriodicSyncStatus = &v
	return s
}

func (s *CreateIdentityProviderRequestUdPullConfig) SetUdSyncScopeConfig(v *CreateIdentityProviderRequestUdPullConfigUdSyncScopeConfig) *CreateIdentityProviderRequestUdPullConfig {
	s.UdSyncScopeConfig = v
	return s
}

func (s *CreateIdentityProviderRequestUdPullConfig) Validate() error {
	if s.PeriodicSyncConfig != nil {
		if err := s.PeriodicSyncConfig.Validate(); err != nil {
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

type CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig struct {
	// cron expression.
	//
	// example:
	//
	// 0 45 1 	- 	- ?
	PeriodicSyncCron *string `json:"PeriodicSyncCron,omitempty" xml:"PeriodicSyncCron,omitempty"`
	// Collection of time points.
	PeriodicSyncTimes []*int32 `json:"PeriodicSyncTimes,omitempty" xml:"PeriodicSyncTimes,omitempty" type:"Repeated"`
	// type.
	//
	// example:
	//
	// cron
	PeriodicSyncType *string `json:"PeriodicSyncType,omitempty" xml:"PeriodicSyncType,omitempty"`
}

func (s CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig) GetPeriodicSyncCron() *string {
	return s.PeriodicSyncCron
}

func (s *CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig) GetPeriodicSyncTimes() []*int32 {
	return s.PeriodicSyncTimes
}

func (s *CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig) GetPeriodicSyncType() *string {
	return s.PeriodicSyncType
}

func (s *CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig) SetPeriodicSyncCron(v string) *CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig {
	s.PeriodicSyncCron = &v
	return s
}

func (s *CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig) SetPeriodicSyncTimes(v []*int32) *CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig {
	s.PeriodicSyncTimes = v
	return s
}

func (s *CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig) SetPeriodicSyncType(v string) *CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig {
	s.PeriodicSyncType = &v
	return s
}

func (s *CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderRequestUdPullConfigUdSyncScopeConfig struct {
	// List of source nodes for synchronization.
	SourceScopes []*string `json:"SourceScopes,omitempty" xml:"SourceScopes,omitempty" type:"Repeated"`
	// Synchronize target node, and fill in the IDaaS organization ID.
	//
	// example:
	//
	// ou_lyhyy6p7yf7mdrdiq5xxxx
	TargetScope *string `json:"TargetScope,omitempty" xml:"TargetScope,omitempty"`
}

func (s CreateIdentityProviderRequestUdPullConfigUdSyncScopeConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestUdPullConfigUdSyncScopeConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestUdPullConfigUdSyncScopeConfig) GetSourceScopes() []*string {
	return s.SourceScopes
}

func (s *CreateIdentityProviderRequestUdPullConfigUdSyncScopeConfig) GetTargetScope() *string {
	return s.TargetScope
}

func (s *CreateIdentityProviderRequestUdPullConfigUdSyncScopeConfig) SetSourceScopes(v []*string) *CreateIdentityProviderRequestUdPullConfigUdSyncScopeConfig {
	s.SourceScopes = v
	return s
}

func (s *CreateIdentityProviderRequestUdPullConfigUdSyncScopeConfig) SetTargetScope(v string) *CreateIdentityProviderRequestUdPullConfigUdSyncScopeConfig {
	s.TargetScope = &v
	return s
}

func (s *CreateIdentityProviderRequestUdPullConfigUdSyncScopeConfig) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderRequestUdPushConfig struct {
	// Incremental callback status. This field is reserved and currently not in use; please ignore it.
	//
	// example:
	//
	// disabled
	IncrementalCallbackStatus *string                                                      `json:"IncrementalCallbackStatus,omitempty" xml:"IncrementalCallbackStatus,omitempty"`
	PeriodicSyncConfig        *CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig `json:"PeriodicSyncConfig,omitempty" xml:"PeriodicSyncConfig,omitempty" type:"Struct"`
	// Periodic check status. This field is currently not in use, please ignore it.
	//
	// example:
	//
	// disabled
	PeriodicSyncStatus *string `json:"PeriodicSyncStatus,omitempty" xml:"PeriodicSyncStatus,omitempty"`
	// Outbound synchronization configuration information.
	UdSyncScopeConfigs []*CreateIdentityProviderRequestUdPushConfigUdSyncScopeConfigs `json:"UdSyncScopeConfigs,omitempty" xml:"UdSyncScopeConfigs,omitempty" type:"Repeated"`
}

func (s CreateIdentityProviderRequestUdPushConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestUdPushConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestUdPushConfig) GetIncrementalCallbackStatus() *string {
	return s.IncrementalCallbackStatus
}

func (s *CreateIdentityProviderRequestUdPushConfig) GetPeriodicSyncConfig() *CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig {
	return s.PeriodicSyncConfig
}

func (s *CreateIdentityProviderRequestUdPushConfig) GetPeriodicSyncStatus() *string {
	return s.PeriodicSyncStatus
}

func (s *CreateIdentityProviderRequestUdPushConfig) GetUdSyncScopeConfigs() []*CreateIdentityProviderRequestUdPushConfigUdSyncScopeConfigs {
	return s.UdSyncScopeConfigs
}

func (s *CreateIdentityProviderRequestUdPushConfig) SetIncrementalCallbackStatus(v string) *CreateIdentityProviderRequestUdPushConfig {
	s.IncrementalCallbackStatus = &v
	return s
}

func (s *CreateIdentityProviderRequestUdPushConfig) SetPeriodicSyncConfig(v *CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig) *CreateIdentityProviderRequestUdPushConfig {
	s.PeriodicSyncConfig = v
	return s
}

func (s *CreateIdentityProviderRequestUdPushConfig) SetPeriodicSyncStatus(v string) *CreateIdentityProviderRequestUdPushConfig {
	s.PeriodicSyncStatus = &v
	return s
}

func (s *CreateIdentityProviderRequestUdPushConfig) SetUdSyncScopeConfigs(v []*CreateIdentityProviderRequestUdPushConfigUdSyncScopeConfigs) *CreateIdentityProviderRequestUdPushConfig {
	s.UdSyncScopeConfigs = v
	return s
}

func (s *CreateIdentityProviderRequestUdPushConfig) Validate() error {
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

type CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig struct {
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

func (s CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig) GetPeriodicSyncCron() *string {
	return s.PeriodicSyncCron
}

func (s *CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig) GetPeriodicSyncTimes() []*int32 {
	return s.PeriodicSyncTimes
}

func (s *CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig) GetPeriodicSyncType() *string {
	return s.PeriodicSyncType
}

func (s *CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig) SetPeriodicSyncCron(v string) *CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig {
	s.PeriodicSyncCron = &v
	return s
}

func (s *CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig) SetPeriodicSyncTimes(v []*int32) *CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig {
	s.PeriodicSyncTimes = v
	return s
}

func (s *CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig) SetPeriodicSyncType(v string) *CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig {
	s.PeriodicSyncType = &v
	return s
}

func (s *CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderRequestUdPushConfigUdSyncScopeConfigs struct {
	// List of source nodes for synchronization.
	SourceScopes []*string `json:"SourceScopes,omitempty" xml:"SourceScopes,omitempty" type:"Repeated"`
	// Target node for synchronization.
	//
	// example:
	//
	// ou_lyhyy6p7yf7mdrdiq5xxxx
	TargetScope *string `json:"TargetScope,omitempty" xml:"TargetScope,omitempty"`
}

func (s CreateIdentityProviderRequestUdPushConfigUdSyncScopeConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestUdPushConfigUdSyncScopeConfigs) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestUdPushConfigUdSyncScopeConfigs) GetSourceScopes() []*string {
	return s.SourceScopes
}

func (s *CreateIdentityProviderRequestUdPushConfigUdSyncScopeConfigs) GetTargetScope() *string {
	return s.TargetScope
}

func (s *CreateIdentityProviderRequestUdPushConfigUdSyncScopeConfigs) SetSourceScopes(v []*string) *CreateIdentityProviderRequestUdPushConfigUdSyncScopeConfigs {
	s.SourceScopes = v
	return s
}

func (s *CreateIdentityProviderRequestUdPushConfigUdSyncScopeConfigs) SetTargetScope(v string) *CreateIdentityProviderRequestUdPushConfigUdSyncScopeConfigs {
	s.TargetScope = &v
	return s
}

func (s *CreateIdentityProviderRequestUdPushConfigUdSyncScopeConfigs) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderRequestWeComConfig struct {
	// Agent ID of the self-built WeCom application.
	//
	// example:
	//
	// 278231941749863339
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// Authorization callback domain.
	//
	// example:
	//
	// https://xxx.aliyunidaas.com/xxxx
	AuthorizeCallbackDomain *string `json:"AuthorizeCallbackDomain,omitempty" xml:"AuthorizeCallbackDomain,omitempty"`
	// Corp ID of the self-built WeCom application.
	//
	// example:
	//
	// 3756043633237690761
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// Corp Secret of the self-built WeCom application.
	//
	// example:
	//
	// CSEHDddddddxxxxuxkJEHPveWRXBGqVqRsxxxx
	CorpSecret *string `json:"CorpSecret,omitempty" xml:"CorpSecret,omitempty"`
	// Trusted domain.
	//
	// example:
	//
	// https://xxx.aliyunidaas.com/
	TrustableDomain *string `json:"TrustableDomain,omitempty" xml:"TrustableDomain,omitempty"`
}

func (s CreateIdentityProviderRequestWeComConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestWeComConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestWeComConfig) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateIdentityProviderRequestWeComConfig) GetAuthorizeCallbackDomain() *string {
	return s.AuthorizeCallbackDomain
}

func (s *CreateIdentityProviderRequestWeComConfig) GetCorpId() *string {
	return s.CorpId
}

func (s *CreateIdentityProviderRequestWeComConfig) GetCorpSecret() *string {
	return s.CorpSecret
}

func (s *CreateIdentityProviderRequestWeComConfig) GetTrustableDomain() *string {
	return s.TrustableDomain
}

func (s *CreateIdentityProviderRequestWeComConfig) SetAgentId(v string) *CreateIdentityProviderRequestWeComConfig {
	s.AgentId = &v
	return s
}

func (s *CreateIdentityProviderRequestWeComConfig) SetAuthorizeCallbackDomain(v string) *CreateIdentityProviderRequestWeComConfig {
	s.AuthorizeCallbackDomain = &v
	return s
}

func (s *CreateIdentityProviderRequestWeComConfig) SetCorpId(v string) *CreateIdentityProviderRequestWeComConfig {
	s.CorpId = &v
	return s
}

func (s *CreateIdentityProviderRequestWeComConfig) SetCorpSecret(v string) *CreateIdentityProviderRequestWeComConfig {
	s.CorpSecret = &v
	return s
}

func (s *CreateIdentityProviderRequestWeComConfig) SetTrustableDomain(v string) *CreateIdentityProviderRequestWeComConfig {
	s.TrustableDomain = &v
	return s
}

func (s *CreateIdentityProviderRequestWeComConfig) Validate() error {
	return dara.Validate(s)
}
