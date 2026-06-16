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
	SetSamlConfig(v *CreateIdentityProviderRequestSamlConfig) *CreateIdentityProviderRequest
	GetSamlConfig() *CreateIdentityProviderRequestSamlConfig
	SetUdPullConfig(v *CreateIdentityProviderRequestUdPullConfig) *CreateIdentityProviderRequest
	GetUdPullConfig() *CreateIdentityProviderRequestUdPullConfig
	SetUdPushConfig(v *CreateIdentityProviderRequestUdPushConfig) *CreateIdentityProviderRequest
	GetUdPushConfig() *CreateIdentityProviderRequestUdPushConfig
	SetWeComConfig(v *CreateIdentityProviderRequestWeComConfig) *CreateIdentityProviderRequest
	GetWeComConfig() *CreateIdentityProviderRequestWeComConfig
}

type CreateIdentityProviderRequest struct {
	// The authentication configurations.
	AuthnConfig *CreateIdentityProviderRequestAuthnConfig `json:"AuthnConfig,omitempty" xml:"AuthnConfig,omitempty" type:"Struct"`
	// The rule configurations for automatic account creation.
	AutoCreateUserConfig *CreateIdentityProviderRequestAutoCreateUserConfig `json:"AutoCreateUserConfig,omitempty" xml:"AutoCreateUserConfig,omitempty" type:"Struct"`
	// The rule configurations for automatic account updates.
	AutoUpdateUserConfig *CreateIdentityProviderRequestAutoUpdateUserConfig `json:"AutoUpdateUserConfig,omitempty" xml:"AutoUpdateUserConfig,omitempty" type:"Struct"`
	// The account binding rule configurations for the OIDC identity provider.
	BindingConfig *CreateIdentityProviderRequestBindingConfig `json:"BindingConfig,omitempty" xml:"BindingConfig,omitempty" type:"Struct"`
	// A client token used to ensure the idempotence of the request. Generate a unique value from your client for each request. The ClientToken can only contain ASCII characters. Note: If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may be different for each API request.
	//
	// example:
	//
	// clientToken_20250704_Axxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The DingTalk configurations.
	DingtalkAppConfig *CreateIdentityProviderRequestDingtalkAppConfig `json:"DingtalkAppConfig,omitempty" xml:"DingtalkAppConfig,omitempty" type:"Struct"`
	// The name of the identity provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	IdentityProviderName *string `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
	// The synchronization type of the identity provider.
	//
	// - Inbound DingTalk: urn:alibaba:idaas:idp:alibaba:dingtalk:pull
	//
	// - Outbound DingTalk: urn:alibaba:idaas:idp:alibaba:dingtalk:push
	//
	// - Inbound WeCom: urn:alibaba:idaas:idp:tencent:wecom:pull
	//
	// - Inbound Lark: urn:alibaba:idaas:idp:bytedance:lark:pull
	//
	// - Inbound AD: urn:alibaba:idaas:idp:microsoft:ad:pull
	//
	// - Inbound LDAP: urn:alibaba:idaas:idp:unknown:ldap:pull
	//
	// - Standard OIDC: urn:alibaba:idaas:idp:standard:oidc
	//
	// - SASE-specific OIDC: urn:alibaba:idaas:idp:alibaba:sase
	//
	// This parameter is required.
	//
	// example:
	//
	// urn:alibaba:idaas:idp:alibaba:dingtalk:push
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Lark configurations.
	LarkConfig *CreateIdentityProviderRequestLarkConfig `json:"LarkConfig,omitempty" xml:"LarkConfig,omitempty" type:"Struct"`
	// The AD/LDAP configurations.
	LdapConfig *CreateIdentityProviderRequestLdapConfig `json:"LdapConfig,omitempty" xml:"LdapConfig,omitempty" type:"Struct"`
	// The URL of the application logo.
	//
	// example:
	//
	// xxxx-image://xxxx_23aqr2ye554csg33dqpch5eu3q/tmp/d17d9adc-a943-45e7-ba0c-2838dddea678
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// The network endpoint ID.
	//
	// example:
	//
	// nae_examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// The OIDC IdP configurations.
	OidcConfig *CreateIdentityProviderRequestOidcConfig `json:"OidcConfig,omitempty" xml:"OidcConfig,omitempty" type:"Struct"`
	SamlConfig *CreateIdentityProviderRequestSamlConfig `json:"SamlConfig,omitempty" xml:"SamlConfig,omitempty" type:"Struct"`
	// The inbound synchronization configurations.
	UdPullConfig *CreateIdentityProviderRequestUdPullConfig `json:"UdPullConfig,omitempty" xml:"UdPullConfig,omitempty" type:"Struct"`
	// The outbound synchronization configurations.
	UdPushConfig *CreateIdentityProviderRequestUdPushConfig `json:"UdPushConfig,omitempty" xml:"UdPushConfig,omitempty" type:"Struct"`
	// The WeCom configurations.
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

func (s *CreateIdentityProviderRequest) GetSamlConfig() *CreateIdentityProviderRequestSamlConfig {
	return s.SamlConfig
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

func (s *CreateIdentityProviderRequest) SetSamlConfig(v *CreateIdentityProviderRequestSamlConfig) *CreateIdentityProviderRequest {
	s.SamlConfig = v
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
	if s.SamlConfig != nil {
		if err := s.SamlConfig.Validate(); err != nil {
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
	// Specifies whether the IdP supports authentication. Valid values:
	//
	// - disabled
	//
	// - enabled
	//
	// example:
	//
	// enabled
	AuthnStatus *string `json:"AuthnStatus,omitempty" xml:"AuthnStatus,omitempty"`
	// Specifies whether to automatically update passwords. Valid values:
	//
	// - disabled
	//
	// - enabled
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
	// Specifies whether to automatically create accounts. Valid values:
	//
	// - disabled
	//
	// - enabled
	//
	// example:
	//
	// disabled
	AutoCreateUserStatus *string `json:"AutoCreateUserStatus,omitempty" xml:"AutoCreateUserStatus,omitempty"`
	// The collection of target organizational unit IDs.
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
	// Specifies whether to automatically update accounts. Valid values:
	//
	// - disabled
	//
	// - enabled
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
	// The list of rules for automatically matching accounts.
	AutoMatchUserProfileExpressions []*CreateIdentityProviderRequestBindingConfigAutoMatchUserProfileExpressions `json:"AutoMatchUserProfileExpressions,omitempty" xml:"AutoMatchUserProfileExpressions,omitempty" type:"Repeated"`
	// Specifies whether to automatically match accounts. Valid values:
	//
	// - disabled
	//
	// - enabled
	//
	// example:
	//
	// disabled
	AutoMatchUserStatus *string `json:"AutoMatchUserStatus,omitempty" xml:"AutoMatchUserStatus,omitempty"`
	// Specifies whether to allow users to manually bind accounts. Valid values:
	//
	// - disabled
	//
	// - enabled
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
	// The type of the expression. Valid values:
	//
	// - field
	//
	// - expression
	//
	// example:
	//
	// field
	ExpressionMappingType *string `json:"ExpressionMappingType,omitempty" xml:"ExpressionMappingType,omitempty"`
	// The expression for the value of the mapped attribute.
	//
	// example:
	//
	// idpUser.phoneNumber
	SourceValueExpression *string `json:"SourceValueExpression,omitempty" xml:"SourceValueExpression,omitempty"`
	// The name of the target mapped attribute.
	//
	// example:
	//
	// user.username
	TargetField *string `json:"TargetField,omitempty" xml:"TargetField,omitempty"`
	// The name of the mapping\\"s target property.
	//
	// example:
	//
	// username
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
	// The AppKey of the DingTalk application.
	//
	// example:
	//
	// Xczngvfemo4e
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The AppSecret of the DingTalk application.
	//
	// example:
	//
	// 5d405a12a6f84ad4ab05ee09axxxx
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// The CorpId of the DingTalk application.
	//
	// example:
	//
	// 307568042478613xxxx
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// The DingTalk version. Valid values:
	//
	// - public_dingtalk: Standard DingTalk
	//
	// - private_dingtalk: Enterprise DingTalk
	//
	// example:
	//
	// public_dingtalk
	DingtalkVersion *string `json:"DingtalkVersion,omitempty" xml:"DingtalkVersion,omitempty"`
	// The EncryptKey of the DingTalk application.
	//
	// example:
	//
	// VkdWw91mdkrjVFr3ObNwefap21dfxxxx
	EncryptKey *string `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
	// The VerificationToken of the DingTalk application.
	//
	// example:
	//
	// myDingApp_VerifyTokenxxxxx
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
	// The AppId of the Lark application.
	//
	// example:
	//
	// cli_xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The App Secret of the Lark application.
	//
	// example:
	//
	// KiiLzh5Dueh4wbLxxxx
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// The EncryptKey of the custom Lark application.
	//
	// example:
	//
	// VkdWw91mdkrjVFr3ObNwefap21dfxxxx
	EncryptKey *string `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
	// The enterprise code of Lark.
	//
	// example:
	//
	// FSX123111xxx
	EnterpriseNumber *string `json:"EnterpriseNumber,omitempty" xml:"EnterpriseNumber,omitempty"`
	// The VerificationToken of the custom Lark application.
	//
	// example:
	//
	// feishuVerifyTokenxxxxx
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
	// The administrator password.
	//
	// example:
	//
	// xxxx
	AdministratorPassword *string `json:"AdministratorPassword,omitempty" xml:"AdministratorPassword,omitempty"`
	// The administrator account.
	//
	// example:
	//
	// DC=example,DC=com
	AdministratorUsername *string `json:"AdministratorUsername,omitempty" xml:"AdministratorUsername,omitempty"`
	// Specifies whether to verify the certificate fingerprint. Valid values:
	//
	// - disabled
	//
	// - enabled
	//
	// example:
	//
	// enabled
	CertificateFingerprintStatus *string `json:"CertificateFingerprintStatus,omitempty" xml:"CertificateFingerprintStatus,omitempty"`
	// The list of certificate fingerprints.
	CertificateFingerprints []*string `json:"CertificateFingerprints,omitempty" xml:"CertificateFingerprints,omitempty" type:"Repeated"`
	// The group member identifier.
	//
	// example:
	//
	// member
	GroupMemberAttributeName *string `json:"GroupMemberAttributeName,omitempty" xml:"GroupMemberAttributeName,omitempty"`
	// The objectClass for groups.
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
	// The communication protocol.
	//
	// example:
	//
	// ldap
	LdapProtocol *string `json:"LdapProtocol,omitempty" xml:"LdapProtocol,omitempty"`
	// The address of the AD/LDAP server.
	//
	// example:
	//
	// 123.xx.xx.89
	LdapServerHost *string `json:"LdapServerHost,omitempty" xml:"LdapServerHost,omitempty"`
	// The port number of the AD/LDAP server.
	//
	// example:
	//
	// 636
	LdapServerPort *int32 `json:"LdapServerPort,omitempty" xml:"LdapServerPort,omitempty"`
	// The objectClass for organizational units.
	//
	// example:
	//
	// organizationUnit,top
	OrganizationUnitObjectClass *string `json:"OrganizationUnitObjectClass,omitempty" xml:"OrganizationUnitObjectClass,omitempty"`
	// The RDN for organizational units.
	//
	// example:
	//
	// ou
	OrganizationalUnitRdn *string `json:"OrganizationalUnitRdn,omitempty" xml:"OrganizationalUnitRdn,omitempty"`
	// The switch for password synchronization.
	//
	// example:
	//
	// enabled
	PasswordSyncStatus *string `json:"PasswordSyncStatus,omitempty" xml:"PasswordSyncStatus,omitempty"`
	// Specifies whether to enable StartTLS. Valid values:
	//
	// - disabled
	//
	// - enabled
	//
	// example:
	//
	// enabled
	StartTlsStatus *string `json:"StartTlsStatus,omitempty" xml:"StartTlsStatus,omitempty"`
	// The user logon identifier.
	//
	// example:
	//
	// userPrincipalName, mail
	UserLoginIdentifier *string `json:"UserLoginIdentifier,omitempty" xml:"UserLoginIdentifier,omitempty"`
	// The objectClass for users.
	//
	// example:
	//
	// person,user
	UserObjectClass *string `json:"UserObjectClass,omitempty" xml:"UserObjectClass,omitempty"`
	// The custom filter for users.
	//
	// example:
	//
	// (|(cn=test)(mail=test@test.com))
	UserObjectClassCustomFilter *string `json:"UserObjectClassCustomFilter,omitempty" xml:"UserObjectClassCustomFilter,omitempty"`
	// The RDN for users.
	//
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
	// The OIDC client authentication configurations.
	AuthnParam *CreateIdentityProviderRequestOidcConfigAuthnParam `json:"AuthnParam,omitempty" xml:"AuthnParam,omitempty" type:"Struct"`
	// The OIDC endpoint configurations.
	EndpointConfig *CreateIdentityProviderRequestOidcConfigEndpointConfig `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty" type:"Struct"`
	// The collection of OIDC authorization scopes.
	//
	// example:
	//
	// openid
	GrantScopes []*string `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
	// The OIDC grant type.
	//
	// example:
	//
	// authorization_code
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
	// The PKCE algorithm. Valid values:
	//
	// - S256: SHA256
	//
	// - plain: Plaintext
	//
	// example:
	//
	// S256
	PkceChallengeMethod *string `json:"PkceChallengeMethod,omitempty" xml:"PkceChallengeMethod,omitempty"`
	// Specifies whether to use PKCE in the Authorization Code grant type.
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
	// The OIDC authentication method. Valid values:
	//
	// - client_secret_basic
	//
	// - client_secret_post
	//
	// example:
	//
	// client_secret_post
	AuthnMethod *string `json:"AuthnMethod,omitempty" xml:"AuthnMethod,omitempty"`
	// The OIDC client ID.
	//
	// example:
	//
	// mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The OpenID Connect (OIDC) client secret.
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
	// The OIDC authorization endpoint.
	//
	// example:
	//
	// https://example.com/auth/authorize
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The OIDC issuer.
	//
	// example:
	//
	// https://example.com/auth
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The OIDC JWKS URI.
	//
	// example:
	//
	// https://example.com/auth/jwks
	JwksUri *string `json:"JwksUri,omitempty" xml:"JwksUri,omitempty"`
	// The OIDC token endpoint.
	//
	// example:
	//
	// https://example.com/auth/token
	TokenEndpoint *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
	// The OIDC user information endpoint.
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

type CreateIdentityProviderRequestSamlConfig struct {
	// example:
	//
	// HTTP-REDIRECT
	BindingMethod *string                                                `json:"BindingMethod,omitempty" xml:"BindingMethod,omitempty"`
	Certificates  []*CreateIdentityProviderRequestSamlConfigCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// example:
	//
	// http://dc.test.com/adfs/services/trust
	IdPEntityId *string `json:"IdPEntityId,omitempty" xml:"IdPEntityId,omitempty"`
	// example:
	//
	// https://dc.test.com/adfs/ls/
	IdPSsoUrl *string `json:"IdPSsoUrl,omitempty" xml:"IdPSsoUrl,omitempty"`
	// example:
	//
	// 180
	MaxClockSkew *int64 `json:"MaxClockSkew,omitempty" xml:"MaxClockSkew,omitempty"`
	// example:
	//
	// true
	RequireRequestSigned *bool `json:"RequireRequestSigned,omitempty" xml:"RequireRequestSigned,omitempty"`
	WantAssertionsSigned *bool `json:"WantAssertionsSigned,omitempty" xml:"WantAssertionsSigned,omitempty"`
	WantResponseSigned   *bool `json:"WantResponseSigned,omitempty" xml:"WantResponseSigned,omitempty"`
}

func (s CreateIdentityProviderRequestSamlConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestSamlConfig) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestSamlConfig) GetBindingMethod() *string {
	return s.BindingMethod
}

func (s *CreateIdentityProviderRequestSamlConfig) GetCertificates() []*CreateIdentityProviderRequestSamlConfigCertificates {
	return s.Certificates
}

func (s *CreateIdentityProviderRequestSamlConfig) GetIdPEntityId() *string {
	return s.IdPEntityId
}

func (s *CreateIdentityProviderRequestSamlConfig) GetIdPSsoUrl() *string {
	return s.IdPSsoUrl
}

func (s *CreateIdentityProviderRequestSamlConfig) GetMaxClockSkew() *int64 {
	return s.MaxClockSkew
}

func (s *CreateIdentityProviderRequestSamlConfig) GetRequireRequestSigned() *bool {
	return s.RequireRequestSigned
}

func (s *CreateIdentityProviderRequestSamlConfig) GetWantAssertionsSigned() *bool {
	return s.WantAssertionsSigned
}

func (s *CreateIdentityProviderRequestSamlConfig) GetWantResponseSigned() *bool {
	return s.WantResponseSigned
}

func (s *CreateIdentityProviderRequestSamlConfig) SetBindingMethod(v string) *CreateIdentityProviderRequestSamlConfig {
	s.BindingMethod = &v
	return s
}

func (s *CreateIdentityProviderRequestSamlConfig) SetCertificates(v []*CreateIdentityProviderRequestSamlConfigCertificates) *CreateIdentityProviderRequestSamlConfig {
	s.Certificates = v
	return s
}

func (s *CreateIdentityProviderRequestSamlConfig) SetIdPEntityId(v string) *CreateIdentityProviderRequestSamlConfig {
	s.IdPEntityId = &v
	return s
}

func (s *CreateIdentityProviderRequestSamlConfig) SetIdPSsoUrl(v string) *CreateIdentityProviderRequestSamlConfig {
	s.IdPSsoUrl = &v
	return s
}

func (s *CreateIdentityProviderRequestSamlConfig) SetMaxClockSkew(v int64) *CreateIdentityProviderRequestSamlConfig {
	s.MaxClockSkew = &v
	return s
}

func (s *CreateIdentityProviderRequestSamlConfig) SetRequireRequestSigned(v bool) *CreateIdentityProviderRequestSamlConfig {
	s.RequireRequestSigned = &v
	return s
}

func (s *CreateIdentityProviderRequestSamlConfig) SetWantAssertionsSigned(v bool) *CreateIdentityProviderRequestSamlConfig {
	s.WantAssertionsSigned = &v
	return s
}

func (s *CreateIdentityProviderRequestSamlConfig) SetWantResponseSigned(v bool) *CreateIdentityProviderRequestSamlConfig {
	s.WantResponseSigned = &v
	return s
}

func (s *CreateIdentityProviderRequestSamlConfig) Validate() error {
	if s.Certificates != nil {
		for _, item := range s.Certificates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateIdentityProviderRequestSamlConfigCertificates struct {
	// example:
	//
	// -----BEGIN CERTIFICATE----- MIIC0jCCAbqgAwIBAgIQXXXXX-----END CERTIFICATE-----
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s CreateIdentityProviderRequestSamlConfigCertificates) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderRequestSamlConfigCertificates) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderRequestSamlConfigCertificates) GetContent() *string {
	return s.Content
}

func (s *CreateIdentityProviderRequestSamlConfigCertificates) SetContent(v string) *CreateIdentityProviderRequestSamlConfigCertificates {
	s.Content = &v
	return s
}

func (s *CreateIdentityProviderRequestSamlConfigCertificates) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderRequestUdPullConfig struct {
	// Specifies whether to synchronize groups. The default value is disabled. Valid values:
	//
	// - disabled
	//
	// - enabled
	//
	// example:
	//
	// disabled
	GroupSyncStatus *string `json:"GroupSyncStatus,omitempty" xml:"GroupSyncStatus,omitempty"`
	// The status of incremental callbacks. Specifies whether to process incremental callback data from the identity provider (IdP). Valid values:
	//
	// - disabled
	//
	// - enabled
	//
	// example:
	//
	// disabled
	IncrementalCallbackStatus *string `json:"IncrementalCallbackStatus,omitempty" xml:"IncrementalCallbackStatus,omitempty"`
	// The configuration for periodic synchronization.
	PeriodicSyncConfig *CreateIdentityProviderRequestUdPullConfigPeriodicSyncConfig `json:"PeriodicSyncConfig,omitempty" xml:"PeriodicSyncConfig,omitempty" type:"Struct"`
	// The status of periodic synchronization. Specifies whether to periodically check for data differences between IDaaS and the IdP. Valid values:
	//
	// - disabled
	//
	// - enabled
	//
	// example:
	//
	// disabled
	PeriodicSyncStatus *string `json:"PeriodicSyncStatus,omitempty" xml:"PeriodicSyncStatus,omitempty"`
	// The synchronization scope configurations.
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
	// The cron expression.
	//
	// example:
	//
	// 0 45 1 	- 	- ?
	PeriodicSyncCron *string `json:"PeriodicSyncCron,omitempty" xml:"PeriodicSyncCron,omitempty"`
	// The collection of running time points.
	PeriodicSyncTimes []*int32 `json:"PeriodicSyncTimes,omitempty" xml:"PeriodicSyncTimes,omitempty" type:"Repeated"`
	// The type.
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
	// The list of source nodes for synchronization.
	SourceScopes []*string `json:"SourceScopes,omitempty" xml:"SourceScopes,omitempty" type:"Repeated"`
	// The target node for synchronization. Enter the IDaaS organization ID.
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
	// The status of incremental callbacks. This parameter is not in use. Ignore this parameter.
	//
	// example:
	//
	// disabled
	IncrementalCallbackStatus *string `json:"IncrementalCallbackStatus,omitempty" xml:"IncrementalCallbackStatus,omitempty"`
	// The periodic synchronization configurations.
	PeriodicSyncConfig *CreateIdentityProviderRequestUdPushConfigPeriodicSyncConfig `json:"PeriodicSyncConfig,omitempty" xml:"PeriodicSyncConfig,omitempty" type:"Struct"`
	// The status of periodic synchronization. This parameter is not in use. Ignore this parameter.
	//
	// example:
	//
	// disabled
	PeriodicSyncStatus *string `json:"PeriodicSyncStatus,omitempty" xml:"PeriodicSyncStatus,omitempty"`
	// The configurations of the outbound synchronization scope.
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
	// The cron expression.
	//
	// example:
	//
	// 0 45 1 	- 	- ?
	PeriodicSyncCron *string `json:"PeriodicSyncCron,omitempty" xml:"PeriodicSyncCron,omitempty"`
	// The collection of running time points.
	PeriodicSyncTimes []*int32 `json:"PeriodicSyncTimes,omitempty" xml:"PeriodicSyncTimes,omitempty" type:"Repeated"`
	// The type.
	//
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
	// The list of source nodes for synchronization.
	SourceScopes []*string `json:"SourceScopes,omitempty" xml:"SourceScopes,omitempty" type:"Repeated"`
	// The target node for synchronization.
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
	// The agent ID of the custom WeCom application.
	//
	// example:
	//
	// 278231941749863339
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The authorized callback domain.
	//
	// example:
	//
	// https://xxx.aliyunidaas.com/xxxx
	AuthorizeCallbackDomain *string `json:"AuthorizeCallbackDomain,omitempty" xml:"AuthorizeCallbackDomain,omitempty"`
	// The CorpId of the custom WeCom application.
	//
	// example:
	//
	// 3756043633237690761
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// The CorpSecret of the custom WeCom application.
	//
	// example:
	//
	// CSEHDddddddxxxxuxkJEHPveWRXBGqVqRsxxxx
	CorpSecret *string `json:"CorpSecret,omitempty" xml:"CorpSecret,omitempty"`
	// The trusted domain name.
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
