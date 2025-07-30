// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderDetail(v *GetIdentityProviderResponseBodyIdentityProviderDetail) *GetIdentityProviderResponseBody
	GetIdentityProviderDetail() *GetIdentityProviderResponseBodyIdentityProviderDetail
	SetRequestId(v string) *GetIdentityProviderResponseBody
	GetRequestId() *string
}

type GetIdentityProviderResponseBody struct {
	// Identity provider Information.
	IdentityProviderDetail *GetIdentityProviderResponseBodyIdentityProviderDetail `json:"IdentityProviderDetail,omitempty" xml:"IdentityProviderDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBody) GetIdentityProviderDetail() *GetIdentityProviderResponseBodyIdentityProviderDetail {
	return s.IdentityProviderDetail
}

func (s *GetIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentityProviderResponseBody) SetIdentityProviderDetail(v *GetIdentityProviderResponseBodyIdentityProviderDetail) *GetIdentityProviderResponseBody {
	s.IdentityProviderDetail = v
	return s
}

func (s *GetIdentityProviderResponseBody) SetRequestId(v string) *GetIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentityProviderResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyIdentityProviderDetail struct {
	// Advanced configuration capability.
	//
	// Value range:
	//
	// Disabled: disabled
	//
	// Enable: enabled
	//
	// example:
	//
	// disabled
	AdvancedStatus *string `json:"AdvancedStatus,omitempty" xml:"AdvancedStatus,omitempty"`
	// The corresponding identity provider product, e.g., Okta, Google, or Azure AD. Possible values:
	//
	// DingTalk: urn:alibaba:idaas:idp:alibaba:dingtalk
	//
	// LDAP: urn:alibaba:idaas:idp:unknown:ldap
	//
	// Alibaba Cloud IDaaS: urn:alibaba:idaas:idp:alibaba:idaas
	//
	// WeCom (Enterprise WeChat): urn:alibaba:idaas:idp:tencent:wecom
	//
	// Lark (Feishu): urn:alibaba:idaas:idp:bytedance:lark
	//
	// Active Directory: urn:alibaba:idaas:idp:microsoft:ad
	//
	// Azure Active Directory: urn:alibaba:idaas:idp:microsoft:aad
	//
	// Alibaba Cloud SASE: urn:alibaba:idaas:idp:alibaba:sase
	//
	// example:
	//
	// urn:alibaba:idaas:idp:bytedance:lark
	AuthnSourceSupplier *string `json:"AuthnSourceSupplier,omitempty" xml:"AuthnSourceSupplier,omitempty"`
	// Authentication type — OIDC or SAML. Possible values:
	//
	// OIDC: urn:alibaba:idaas:authntype:oidc
	//
	// SAML: urn:alibaba:idaas:authntype:saml2
	//
	// example:
	//
	// urn:alibaba:idaas:authntype:oidc
	AuthnSourceType *string `json:"AuthnSourceType,omitempty" xml:"AuthnSourceType,omitempty"`
	// Whether the corresponding IdP supports authentication. Value range:
	//
	// Disabled: disabled
	//
	// Enabled: enabled
	//
	// example:
	//
	// disabled
	AuthnStatus *string `json:"AuthnStatus,omitempty" xml:"AuthnStatus,omitempty"`
	// The time when the version was created.
	//
	// example:
	//
	// 1726021079000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Identity provider description.
	//
	// example:
	//
	// for poc test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// DingTalk Basic Configuration
	DingtalkAppConfig *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig `json:"DingtalkAppConfig,omitempty" xml:"DingtalkAppConfig,omitempty" type:"Struct"`
	// DingTalk synchronous configuration.
	DingtalkProvisioningConfig *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig `json:"DingtalkProvisioningConfig,omitempty" xml:"DingtalkProvisioningConfig,omitempty" type:"Struct"`
	// Identity provider external ID.
	//
	// example:
	//
	// idp_xxxx
	IdentityProviderExternalId *string `json:"IdentityProviderExternalId,omitempty" xml:"IdentityProviderExternalId,omitempty"`
	// Identity provider ID.
	//
	// example:
	//
	// idp_mwpcwnhrimlr2horx7xgg7xxxx
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// Identity provider name.
	//
	// example:
	//
	// test
	IdentityProviderName *string `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
	// Identity provider type.
	//
	// example:
	//
	// urn:alibaba:idaas:idp:alibaba:dingtalk:push
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// idaas_x2df3bak3uwnapqm6xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Lark configuration.
	LarkConfig *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig `json:"LarkConfig,omitempty" xml:"LarkConfig,omitempty" type:"Struct"`
	// Last status check result.
	//
	// example:
	//
	// success
	LastStatusCheckJobResult *string `json:"LastStatusCheckJobResult,omitempty" xml:"LastStatusCheckJobResult,omitempty"`
	// AD/LDAP Identity provider information.
	LdapConfig *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig `json:"LdapConfig,omitempty" xml:"LdapConfig,omitempty" type:"Struct"`
	// The reason why write operations on the instance are locked.
	//
	// example:
	//
	// financial
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The URL of the application logo.
	//
	// example:
	//
	// https://img.alicdn.com/imgextra/i4/O1CN01OB8fJj22fpoZm4sd0_!!6000000007148-2-tps-149-xxx.png
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// The unique identifier of the network access endpoint.
	//
	// example:
	//
	// nae_mx4vsadfe6govkqkwckxxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// OIDC IdP configuration.
	OidcConfig *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig `json:"OidcConfig,omitempty" xml:"OidcConfig,omitempty" type:"Struct"`
	// Sync in configuration.
	UdPullConfig *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig `json:"UdPullConfig,omitempty" xml:"UdPullConfig,omitempty" type:"Struct"`
	// Indicates whether the IDaaS EIAM system supports UD (User Directory) synchronization.
	//
	// example:
	//
	// disabled
	UdPullStatus *string `json:"UdPullStatus,omitempty" xml:"UdPullStatus,omitempty"`
	// Outbound synchronization configuration.
	UdPushConfig *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfig `json:"UdPushConfig,omitempty" xml:"UdPushConfig,omitempty" type:"Struct"`
	// Outbound synchronization capability.
	//
	// example:
	//
	// disabled
	UdPushStatus *string `json:"UdPushStatus,omitempty" xml:"UdPushStatus,omitempty"`
	// The time when the serviceInstance  was last updated.
	//
	// example:
	//
	// 1726021079000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// WeCom configuration.
	WeComConfig *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig `json:"WeComConfig,omitempty" xml:"WeComConfig,omitempty" type:"Struct"`
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetail) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetail) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetAdvancedStatus() *string {
	return s.AdvancedStatus
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetAuthnSourceSupplier() *string {
	return s.AuthnSourceSupplier
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetAuthnSourceType() *string {
	return s.AuthnSourceType
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetAuthnStatus() *string {
	return s.AuthnStatus
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetDescription() *string {
	return s.Description
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetDingtalkAppConfig() *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig {
	return s.DingtalkAppConfig
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetDingtalkProvisioningConfig() *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig {
	return s.DingtalkProvisioningConfig
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetIdentityProviderExternalId() *string {
	return s.IdentityProviderExternalId
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetLarkConfig() *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig {
	return s.LarkConfig
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetLastStatusCheckJobResult() *string {
	return s.LastStatusCheckJobResult
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetLdapConfig() *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig {
	return s.LdapConfig
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetLockReason() *string {
	return s.LockReason
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetLogoUrl() *string {
	return s.LogoUrl
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetOidcConfig() *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig {
	return s.OidcConfig
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetUdPullConfig() *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig {
	return s.UdPullConfig
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetUdPullStatus() *string {
	return s.UdPullStatus
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetUdPushConfig() *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfig {
	return s.UdPushConfig
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetUdPushStatus() *string {
	return s.UdPushStatus
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) GetWeComConfig() *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig {
	return s.WeComConfig
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetAdvancedStatus(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.AdvancedStatus = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetAuthnSourceSupplier(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.AuthnSourceSupplier = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetAuthnSourceType(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.AuthnSourceType = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetAuthnStatus(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.AuthnStatus = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetCreateTime(v int64) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.CreateTime = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetDescription(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.Description = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetDingtalkAppConfig(v *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.DingtalkAppConfig = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetDingtalkProvisioningConfig(v *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.DingtalkProvisioningConfig = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetIdentityProviderExternalId(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.IdentityProviderExternalId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetIdentityProviderId(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.IdentityProviderId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetIdentityProviderName(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.IdentityProviderName = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetIdentityProviderType(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.IdentityProviderType = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetInstanceId(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.InstanceId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetLarkConfig(v *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.LarkConfig = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetLastStatusCheckJobResult(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.LastStatusCheckJobResult = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetLdapConfig(v *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.LdapConfig = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetLockReason(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.LockReason = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetLogoUrl(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.LogoUrl = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetNetworkAccessEndpointId(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetOidcConfig(v *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.OidcConfig = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetUdPullConfig(v *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.UdPullConfig = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetUdPullStatus(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.UdPullStatus = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetUdPushConfig(v *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfig) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.UdPushConfig = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetUdPushStatus(v string) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.UdPushStatus = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetUpdateTime(v int64) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.UpdateTime = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) SetWeComConfig(v *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig) *GetIdentityProviderResponseBodyIdentityProviderDetail {
	s.WeComConfig = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetail) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig struct {
	// The AppKey for the application.
	//
	// example:
	//
	// 41reopmwoy9s
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The details of the application secret.
	//
	// example:
	//
	// REOQ6Cl55kriOd8NOBeqWYLKpHR4p6fdZxxxx
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// DingTalk corpId.
	//
	// example:
	//
	// 3756043633237690761
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// DingTalk Version.
	//
	// example:
	//
	// public_dingtalk
	DingtalkVersion *string `json:"DingtalkVersion,omitempty" xml:"DingtalkVersion,omitempty"`
	// DingTalk  encrypt key.
	//
	// example:
	//
	// 29003eb11d0a28b4802a6f02fb8aa25dff730e2ac26ffdxxx
	EncryptKey *string `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
	// DingTalk  verification token.
	//
	// example:
	//
	// 5ba9c127a7abe029003eb11d0a28b4802a6f02fb8aa25dff730e2ac26ffxxxxx
	VerificationToken *string `json:"VerificationToken,omitempty" xml:"VerificationToken,omitempty"`
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) GetAppKey() *string {
	return s.AppKey
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) GetAppSecret() *string {
	return s.AppSecret
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) GetCorpId() *string {
	return s.CorpId
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) GetDingtalkVersion() *string {
	return s.DingtalkVersion
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) GetEncryptKey() *string {
	return s.EncryptKey
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) GetVerificationToken() *string {
	return s.VerificationToken
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) SetAppKey(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig {
	s.AppKey = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) SetAppSecret(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig {
	s.AppSecret = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) SetCorpId(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig {
	s.CorpId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) SetDingtalkVersion(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig {
	s.DingtalkVersion = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) SetEncryptKey(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig {
	s.EncryptKey = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) SetVerificationToken(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig {
	s.VerificationToken = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkAppConfig) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig struct {
	// List of authorized DingTalk departments.
	AuthedDepartmentIds []*GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedDepartmentIds `json:"AuthedDepartmentIds,omitempty" xml:"AuthedDepartmentIds,omitempty" type:"Repeated"`
	// Authorized DingTalk account list.
	AuthedUsers []*GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedUsers `json:"AuthedUsers,omitempty" xml:"AuthedUsers,omitempty" type:"Repeated"`
	// DingTalk enterprise corpId.
	//
	// example:
	//
	// ding_xxxxx
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// The name of the company.
	//
	// example:
	//
	// test_xxx
	CorpName *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig) GetAuthedDepartmentIds() []*GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedDepartmentIds {
	return s.AuthedDepartmentIds
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig) GetAuthedUsers() []*GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedUsers {
	return s.AuthedUsers
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig) GetCorpId() *string {
	return s.CorpId
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig) GetCorpName() *string {
	return s.CorpName
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig) SetAuthedDepartmentIds(v []*GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedDepartmentIds) *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig {
	s.AuthedDepartmentIds = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig) SetAuthedUsers(v []*GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedUsers) *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig {
	s.AuthedUsers = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig) SetCorpId(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig {
	s.CorpId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig) SetCorpName(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig {
	s.CorpName = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfig) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedDepartmentIds struct {
	// Department ID.
	//
	// example:
	//
	// 123xxx444
	DeptId *string `json:"DeptId,omitempty" xml:"DeptId,omitempty"`
	// Department name.
	//
	// example:
	//
	// test_xxx
	DeptName *string `json:"DeptName,omitempty" xml:"DeptName,omitempty"`
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedDepartmentIds) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedDepartmentIds) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedDepartmentIds) GetDeptId() *string {
	return s.DeptId
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedDepartmentIds) GetDeptName() *string {
	return s.DeptName
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedDepartmentIds) SetDeptId(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedDepartmentIds {
	s.DeptId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedDepartmentIds) SetDeptName(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedDepartmentIds {
	s.DeptName = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedDepartmentIds) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedUsers struct {
	// DingTalk user name.
	//
	// example:
	//
	// zhangsan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// DingTalk user id.
	//
	// example:
	//
	// 13030833392920xxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedUsers) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedUsers) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedUsers) GetName() *string {
	return s.Name
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedUsers) GetUserId() *string {
	return s.UserId
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedUsers) SetName(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedUsers {
	s.Name = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedUsers) SetUserId(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedUsers {
	s.UserId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailDingtalkProvisioningConfigAuthedUsers) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig struct {
	// The application ID.
	//
	// example:
	//
	// cli_a7a99f53a317xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// ***
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// Feishu encryptKey.
	//
	// example:
	//
	// c5db46da8d4b751a7878e5d670402fb60e4d2391de3fa01f7c9e6353f6d1xxxxx
	EncryptKey *string `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
	// Feishu enterprise code.
	//
	// example:
	//
	// FX1231xxxx
	EnterpriseNumber *string `json:"EnterpriseNumber,omitempty" xml:"EnterpriseNumber,omitempty"`
	// Feishu verificationToken.
	//
	// example:
	//
	// c5db46da8d4b751a7878e5d670402fb60e4d2391de3fa01f7c9e6353f6d1xxxxx
	VerificationToken *string `json:"VerificationToken,omitempty" xml:"VerificationToken,omitempty"`
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig) GetAppId() *string {
	return s.AppId
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig) GetAppSecret() *string {
	return s.AppSecret
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig) GetEncryptKey() *string {
	return s.EncryptKey
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig) GetEnterpriseNumber() *string {
	return s.EnterpriseNumber
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig) GetVerificationToken() *string {
	return s.VerificationToken
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig) SetAppId(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig {
	s.AppId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig) SetAppSecret(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig {
	s.AppSecret = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig) SetEncryptKey(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig {
	s.EncryptKey = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig) SetEnterpriseNumber(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig {
	s.EnterpriseNumber = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig) SetVerificationToken(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig {
	s.VerificationToken = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLarkConfig) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig struct {
	// Administrator password.
	//
	// example:
	//
	// XXXX
	AdministratorPassword *string `json:"AdministratorPassword,omitempty" xml:"AdministratorPassword,omitempty"`
	// Administrator username.
	//
	// example:
	//
	// example.com
	AdministratorUsername *string `json:"AdministratorUsername,omitempty" xml:"AdministratorUsername,omitempty"`
	// Whether to verify the fingerprint certificate.
	//
	// example:
	//
	// enabled
	CertificateFingerprintStatus *string `json:"CertificateFingerprintStatus,omitempty" xml:"CertificateFingerprintStatus,omitempty"`
	// Certificate fingerprint list.
	CertificateFingerprints []*string `json:"CertificateFingerprints,omitempty" xml:"CertificateFingerprints,omitempty" type:"Repeated"`
	// Ldap protocol.
	//
	// example:
	//
	// ldap
	LdapProtocol *string `json:"LdapProtocol,omitempty" xml:"LdapProtocol,omitempty"`
	// ldap server host.
	//
	// example:
	//
	// 127.xx.xx.100
	LdapServerHost *string `json:"LdapServerHost,omitempty" xml:"LdapServerHost,omitempty"`
	// ldap server port.
	//
	// example:
	//
	// 389
	LdapServerPort *int32 `json:"LdapServerPort,omitempty" xml:"LdapServerPort,omitempty"`
	// StartTls status.
	//
	// example:
	//
	// enabled
	StartTlsStatus *string `json:"StartTlsStatus,omitempty" xml:"StartTlsStatus,omitempty"`
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) GetAdministratorPassword() *string {
	return s.AdministratorPassword
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) GetAdministratorUsername() *string {
	return s.AdministratorUsername
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) GetCertificateFingerprintStatus() *string {
	return s.CertificateFingerprintStatus
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) GetCertificateFingerprints() []*string {
	return s.CertificateFingerprints
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) GetLdapProtocol() *string {
	return s.LdapProtocol
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) GetLdapServerHost() *string {
	return s.LdapServerHost
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) GetLdapServerPort() *int32 {
	return s.LdapServerPort
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) GetStartTlsStatus() *string {
	return s.StartTlsStatus
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) SetAdministratorPassword(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig {
	s.AdministratorPassword = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) SetAdministratorUsername(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig {
	s.AdministratorUsername = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) SetCertificateFingerprintStatus(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig {
	s.CertificateFingerprintStatus = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) SetCertificateFingerprints(v []*string) *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig {
	s.CertificateFingerprints = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) SetLdapProtocol(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig {
	s.LdapProtocol = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) SetLdapServerHost(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig {
	s.LdapServerHost = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) SetLdapServerPort(v int32) *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig {
	s.LdapServerPort = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) SetStartTlsStatus(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig {
	s.StartTlsStatus = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailLdapConfig) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig struct {
	// OIDC client authentication configuration.
	AuthnParam *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam `json:"AuthnParam,omitempty" xml:"AuthnParam,omitempty" type:"Struct"`
	// OIDC endpoint configuration.
	EndpointConfig *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty" type:"Struct"`
	// OIDC authorization scope list.
	//
	// example:
	//
	// openid
	GrantScopes []*string `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
	// OIDC authorization grant type.
	//
	// example:
	//
	// authorization_code
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
	// Supported PKCE code challenge methods.
	//
	// example:
	//
	// S256
	PkceChallengeMethod *string `json:"PkceChallengeMethod,omitempty" xml:"PkceChallengeMethod,omitempty"`
	// Whether to use PKCE in authorization code grant flow.
	//
	// example:
	//
	// true
	PkceRequired *bool `json:"PkceRequired,omitempty" xml:"PkceRequired,omitempty"`
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) GetAuthnParam() *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam {
	return s.AuthnParam
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) GetEndpointConfig() *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig {
	return s.EndpointConfig
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) GetGrantScopes() []*string {
	return s.GrantScopes
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) GetGrantType() *string {
	return s.GrantType
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) GetPkceChallengeMethod() *string {
	return s.PkceChallengeMethod
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) GetPkceRequired() *bool {
	return s.PkceRequired
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) SetAuthnParam(v *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam) *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig {
	s.AuthnParam = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) SetEndpointConfig(v *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig) *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig {
	s.EndpointConfig = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) SetGrantScopes(v []*string) *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig {
	s.GrantScopes = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) SetGrantType(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig {
	s.GrantType = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) SetPkceChallengeMethod(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig {
	s.PkceChallengeMethod = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) SetPkceRequired(v bool) *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig {
	s.PkceRequired = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfig) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam struct {
	// OIDC/OAuth2 authentication method.
	//
	// example:
	//
	// client_secret_post
	AuthnMethod *string `json:"AuthnMethod,omitempty" xml:"AuthnMethod,omitempty"`
	// The client ID of the device whose access credential you want to query.
	//
	// example:
	//
	// mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The application secret registered with the OIDC authentication service.
	//
	// example:
	//
	// CSEHDddddddxxxxuxkJEHPveWRXBGqVqRsxxxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam) GetAuthnMethod() *string {
	return s.AuthnMethod
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam) GetClientId() *string {
	return s.ClientId
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam) SetAuthnMethod(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam {
	s.AuthnMethod = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam) SetClientId(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam {
	s.ClientId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam) SetClientSecret(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam {
	s.ClientSecret = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigAuthnParam) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig struct {
	// OAuth2 authorization endpoint.
	//
	// example:
	//
	// https://example.com/oauth/authorize
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	// The CA that issued the certificate.
	//
	// example:
	//
	// https://example.com/oauth
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// Jwks uri.
	//
	// example:
	//
	// https://example.com/oauth/jwks
	JwksUri *string `json:"JwksUri,omitempty" xml:"JwksUri,omitempty"`
	// Token endpoint.
	//
	// example:
	//
	// https://example.com/oauth/token
	TokenEndpoint *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
	// OIDC user info endpoint.
	//
	// example:
	//
	// https://example.com/oauth/userinfo
	UserinfoEndpoint *string `json:"UserinfoEndpoint,omitempty" xml:"UserinfoEndpoint,omitempty"`
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig) GetIssuer() *string {
	return s.Issuer
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig) GetJwksUri() *string {
	return s.JwksUri
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig) GetUserinfoEndpoint() *string {
	return s.UserinfoEndpoint
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig) SetAuthorizationEndpoint(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig) SetIssuer(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig {
	s.Issuer = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig) SetJwksUri(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig {
	s.JwksUri = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig) SetTokenEndpoint(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig {
	s.TokenEndpoint = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig) SetUserinfoEndpoint(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig {
	s.UserinfoEndpoint = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailOidcConfigEndpointConfig) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig struct {
	// Whether to enable group synchronization. Possible values:
	//
	// Disabled: disabled
	//
	// Enabled: enabled
	//
	// example:
	//
	// disabled
	GroupSyncStatus *string `json:"GroupSyncStatus,omitempty" xml:"GroupSyncStatus,omitempty"`
	// Incremental callback status: Whether to process incremental callback data from the IdP.
	//
	// example:
	//
	// disabled
	IncrementalCallbackStatus *string `json:"IncrementalCallbackStatus,omitempty" xml:"IncrementalCallbackStatus,omitempty"`
	// Inbound synchronization configuration Information.
	UdSyncScopeConfig *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfigUdSyncScopeConfig `json:"UdSyncScopeConfig,omitempty" xml:"UdSyncScopeConfig,omitempty" type:"Struct"`
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig) GetGroupSyncStatus() *string {
	return s.GroupSyncStatus
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig) GetIncrementalCallbackStatus() *string {
	return s.IncrementalCallbackStatus
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig) GetUdSyncScopeConfig() *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfigUdSyncScopeConfig {
	return s.UdSyncScopeConfig
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig) SetGroupSyncStatus(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig {
	s.GroupSyncStatus = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig) SetIncrementalCallbackStatus(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig {
	s.IncrementalCallbackStatus = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig) SetUdSyncScopeConfig(v *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfigUdSyncScopeConfig) *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig {
	s.UdSyncScopeConfig = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfig) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfigUdSyncScopeConfig struct {
	// Synchronization source node.
	SourceScopes []*string `json:"SourceScopes,omitempty" xml:"SourceScopes,omitempty" type:"Repeated"`
	// Synchronization target node.
	//
	// example:
	//
	// ou_123xxxx
	TargetScope *string `json:"TargetScope,omitempty" xml:"TargetScope,omitempty"`
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfigUdSyncScopeConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfigUdSyncScopeConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfigUdSyncScopeConfig) GetSourceScopes() []*string {
	return s.SourceScopes
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfigUdSyncScopeConfig) GetTargetScope() *string {
	return s.TargetScope
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfigUdSyncScopeConfig) SetSourceScopes(v []*string) *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfigUdSyncScopeConfig {
	s.SourceScopes = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfigUdSyncScopeConfig) SetTargetScope(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfigUdSyncScopeConfig {
	s.TargetScope = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPullConfigUdSyncScopeConfig) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfig struct {
	// Incremental callback status: Whether to process incremental callback data from the IdP.
	//
	// example:
	//
	// disabled
	IncrementalCallbackStatus *string `json:"IncrementalCallbackStatus,omitempty" xml:"IncrementalCallbackStatus,omitempty"`
	// Outbound synchronization configuration Information.
	UdSyncScopeConfigs []*GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfigUdSyncScopeConfigs `json:"UdSyncScopeConfigs,omitempty" xml:"UdSyncScopeConfigs,omitempty" type:"Repeated"`
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfig) GetIncrementalCallbackStatus() *string {
	return s.IncrementalCallbackStatus
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfig) GetUdSyncScopeConfigs() []*GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfigUdSyncScopeConfigs {
	return s.UdSyncScopeConfigs
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfig) SetIncrementalCallbackStatus(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfig {
	s.IncrementalCallbackStatus = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfig) SetUdSyncScopeConfigs(v []*GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfigUdSyncScopeConfigs) *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfig {
	s.UdSyncScopeConfigs = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfig) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfigUdSyncScopeConfigs struct {
	// Synchronization source node.
	SourceScopes []*string `json:"SourceScopes,omitempty" xml:"SourceScopes,omitempty" type:"Repeated"`
	// Synchronization target node.
	//
	// example:
	//
	// ou_123xxxx
	TargetScope *string `json:"TargetScope,omitempty" xml:"TargetScope,omitempty"`
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfigUdSyncScopeConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfigUdSyncScopeConfigs) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfigUdSyncScopeConfigs) GetSourceScopes() []*string {
	return s.SourceScopes
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfigUdSyncScopeConfigs) GetTargetScope() *string {
	return s.TargetScope
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfigUdSyncScopeConfigs) SetSourceScopes(v []*string) *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfigUdSyncScopeConfigs {
	s.SourceScopes = v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfigUdSyncScopeConfigs) SetTargetScope(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfigUdSyncScopeConfigs {
	s.TargetScope = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailUdPushConfigUdSyncScopeConfigs) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig struct {
	// The ID of the load generator. This parameter is disabled.
	//
	// example:
	//
	// 1242350
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// Authorization callback domain.
	//
	// example:
	//
	// https://example.com/xxxx
	AuthorizeCallbackDomain *string `json:"AuthorizeCallbackDomain,omitempty" xml:"AuthorizeCallbackDomain,omitempty"`
	// CorpId.
	//
	// example:
	//
	// 356201295345457xxxxx
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// Corp secret.
	//
	// example:
	//
	// weaseiszjskejskaj12sjeszojxxxx
	CorpSecret *string `json:"CorpSecret,omitempty" xml:"CorpSecret,omitempty"`
	// Trusted domain.
	//
	// example:
	//
	// https://example.com
	TrustableDomain *string `json:"TrustableDomain,omitempty" xml:"TrustableDomain,omitempty"`
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig) GetAgentId() *string {
	return s.AgentId
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig) GetAuthorizeCallbackDomain() *string {
	return s.AuthorizeCallbackDomain
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig) GetCorpId() *string {
	return s.CorpId
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig) GetCorpSecret() *string {
	return s.CorpSecret
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig) GetTrustableDomain() *string {
	return s.TrustableDomain
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig) SetAgentId(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig {
	s.AgentId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig) SetAuthorizeCallbackDomain(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig {
	s.AuthorizeCallbackDomain = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig) SetCorpId(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig {
	s.CorpId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig) SetCorpSecret(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig {
	s.CorpSecret = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig) SetTrustableDomain(v string) *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig {
	s.TrustableDomain = &v
	return s
}

func (s *GetIdentityProviderResponseBodyIdentityProviderDetailWeComConfig) Validate() error {
	return dara.Validate(s)
}
