// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviders(v []*ListIdentityProvidersResponseBodyIdentityProviders) *ListIdentityProvidersResponseBody
	GetIdentityProviders() []*ListIdentityProvidersResponseBodyIdentityProviders
	SetRequestId(v string) *ListIdentityProvidersResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListIdentityProvidersResponseBody
	GetTotalCount() *int64
}

type ListIdentityProvidersResponseBody struct {
	// The list of identity providers.
	IdentityProviders []*ListIdentityProvidersResponseBodyIdentityProviders `json:"IdentityProviders,omitempty" xml:"IdentityProviders,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIdentityProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersResponseBody) GetIdentityProviders() []*ListIdentityProvidersResponseBodyIdentityProviders {
	return s.IdentityProviders
}

func (s *ListIdentityProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIdentityProvidersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIdentityProvidersResponseBody) SetIdentityProviders(v []*ListIdentityProvidersResponseBodyIdentityProviders) *ListIdentityProvidersResponseBody {
	s.IdentityProviders = v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetRequestId(v string) *ListIdentityProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetTotalCount(v int64) *ListIdentityProvidersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) Validate() error {
	if s.IdentityProviders != nil {
		for _, item := range s.IdentityProviders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIdentityProvidersResponseBodyIdentityProviders struct {
	// Specifies whether the advanced configuration capability is enabled. Valid values:
	//
	// - disabled: Disabled.
	//
	// - enabled: Enabled.
	//
	// example:
	//
	// disabled
	AdvancedStatus *string `json:"AdvancedStatus,omitempty" xml:"AdvancedStatus,omitempty"`
	// The authentication source product, such as Okta, Google, or Azure AD. Valid values:
	//
	// - urn:alibaba:idaas:idp:alibaba:dingtalk: DingTalk.
	//
	// - urn:alibaba:idaas:idp:unknown:ldap: LDAP.
	//
	// - urn:alibaba:idaas:idp:alibaba:idaas: Alibaba Cloud IDaaS.
	//
	// - urn:alibaba:idaas:idp:tencent:wecom: WeCom.
	//
	// - urn:alibaba:idaas:idp:bytedance:lark: Lark.
	//
	// - urn:alibaba:idaas:idp:microsoft:ad: Active Directory.
	//
	// - urn:alibaba:idaas:idp:microsoft:aad: Azure Active Directory.
	//
	// - urn:alibaba:idaas:idp:alibaba:sase: Alibaba Cloud SASE.
	//
	// example:
	//
	// urn:alibaba:idaas:idp:bytedance:lark
	AuthnSourceSupplier *string `json:"AuthnSourceSupplier,omitempty" xml:"AuthnSourceSupplier,omitempty"`
	// The authentication method type, which can be OIDC or SAML. Valid values:
	//
	// - urn:alibaba:idaas:authntype:oidc: OIDC.
	//
	// - urn:alibaba:idaas:authntype:saml2: SAML.
	//
	// example:
	//
	// urn:alibaba:idaas:authntype:oidc
	AuthnSourceType *string `json:"AuthnSourceType,omitempty" xml:"AuthnSourceType,omitempty"`
	// Specifies whether the IdP supports authentication. Valid values:
	//
	// - disabled: Disabled.
	//
	// - enabled: Enabled.
	//
	// example:
	//
	// disabled
	AuthnStatus *string `json:"AuthnStatus,omitempty" xml:"AuthnStatus,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1712561597000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the identity provider.
	//
	// example:
	//
	// None
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The external ID of the identity provider.
	//
	// example:
	//
	// test_123
	IdentityProviderExternalId *string `json:"IdentityProviderExternalId,omitempty" xml:"IdentityProviderExternalId,omitempty"`
	// The identity provider ID.
	//
	// example:
	//
	// idp_m5b5wd5s2hpq4t6iaehhXXX
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// The name of the identity provider.
	//
	// example:
	//
	// test
	IdentityProviderName *string `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
	// The identity provider synchronization type. Valid values:
	//
	// - urn:alibaba:idaas:idp:alibaba:dingtalk:pull: Inbound DingTalk.
	//
	// - urn:alibaba:idaas:idp:alibaba:dingtalk:push: Outbound DingTalk.
	//
	// - urn:alibaba:idaas:idp:tencent:wecom:pull: Inbound WeCom.
	//
	// - urn:alibaba:idaas:idp:bytedance:lark:pull: Inbound Lark.
	//
	// - urn:alibaba:idaas:idp:microsoft:ad:pull: Inbound Active Directory.
	//
	// - urn:alibaba:idaas:idp:unknown:ldap:pull: Inbound LDAP.
	//
	// - urn:alibaba:idaas:idp:standard:oidc: Standard OIDC.
	//
	// - urn:alibaba:idaas:idp:alibaba:sase: SASE custom OIDC.
	//
	// example:
	//
	// urn:alibaba:idaas:idp:bytedance:lark:pull
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	// The incremental callback status, which specifies whether to process incremental callback data from the IdP. Valid values:
	//
	// - disabled: Disabled.
	//
	// - enabled: Enabled.
	//
	// example:
	//
	// enabled
	IncrementalCallbackStatus *string `json:"IncrementalCallbackStatus,omitempty" xml:"IncrementalCallbackStatus,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_pbf4dth34l2qb7mydpntXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The result of the last status check. A synchronization task can be triggered only when the identity provider status check result is success before synchronization.
	//
	// example:
	//
	// success
	LastStatusCheckJobResult *string `json:"LastStatusCheckJobResult,omitempty" xml:"LastStatusCheckJobResult,omitempty"`
	// The lock reason.
	//
	// example:
	//
	// financial
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The custom logo URL of the identity provider.
	//
	// example:
	//
	// https://cdn-cn-hangzhou.aliyunidaas.com/xx/logos/xx
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// The periodic verification status, which specifies whether to periodically verify data differences between EIAM and the identity provider.
	//
	// example:
	//
	// disabled
	PeriodicSyncStatus *string `json:"PeriodicSyncStatus,omitempty" xml:"PeriodicSyncStatus,omitempty"`
	// Specifies whether the inbound synchronization feature is supported. Valid values:
	//
	// - disabled: Disabled.
	//
	// - enabled: Enabled.
	//
	// example:
	//
	// disabled
	UdPullStatus *string `json:"UdPullStatus,omitempty" xml:"UdPullStatus,omitempty"`
	// The synchronization target node.
	//
	// example:
	//
	// ou_2buqmxsa3ltyqkjgpwfijurXXX
	UdPullTargetScope *string `json:"UdPullTargetScope,omitempty" xml:"UdPullTargetScope,omitempty"`
	// Specifies whether the outbound synchronization capability is enabled. Valid values:
	//
	// - disabled: Disabled.
	//
	// - enabled: Enabled.
	//
	// example:
	//
	// disabled
	UdPushStatus *string `json:"UdPushStatus,omitempty" xml:"UdPushStatus,omitempty"`
	// The update time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1712561597000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListIdentityProvidersResponseBodyIdentityProviders) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersResponseBodyIdentityProviders) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetAdvancedStatus() *string {
	return s.AdvancedStatus
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetAuthnSourceSupplier() *string {
	return s.AuthnSourceSupplier
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetAuthnSourceType() *string {
	return s.AuthnSourceType
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetAuthnStatus() *string {
	return s.AuthnStatus
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetDescription() *string {
	return s.Description
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetIdentityProviderExternalId() *string {
	return s.IdentityProviderExternalId
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetIncrementalCallbackStatus() *string {
	return s.IncrementalCallbackStatus
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetLastStatusCheckJobResult() *string {
	return s.LastStatusCheckJobResult
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetLockReason() *string {
	return s.LockReason
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetLogoUrl() *string {
	return s.LogoUrl
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetPeriodicSyncStatus() *string {
	return s.PeriodicSyncStatus
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetUdPullStatus() *string {
	return s.UdPullStatus
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetUdPullTargetScope() *string {
	return s.UdPullTargetScope
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetUdPushStatus() *string {
	return s.UdPushStatus
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetAdvancedStatus(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.AdvancedStatus = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetAuthnSourceSupplier(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.AuthnSourceSupplier = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetAuthnSourceType(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.AuthnSourceType = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetAuthnStatus(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.AuthnStatus = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetCreateTime(v int64) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.CreateTime = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetDescription(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.Description = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetIdentityProviderExternalId(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.IdentityProviderExternalId = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetIdentityProviderId(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.IdentityProviderId = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetIdentityProviderName(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.IdentityProviderName = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetIdentityProviderType(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.IdentityProviderType = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetIncrementalCallbackStatus(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.IncrementalCallbackStatus = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetInstanceId(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.InstanceId = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetLastStatusCheckJobResult(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.LastStatusCheckJobResult = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetLockReason(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.LockReason = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetLogoUrl(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.LogoUrl = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetPeriodicSyncStatus(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.PeriodicSyncStatus = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetUdPullStatus(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.UdPullStatus = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetUdPullTargetScope(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.UdPullTargetScope = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetUdPushStatus(v string) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.UdPushStatus = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) SetUpdateTime(v int64) *ListIdentityProvidersResponseBodyIdentityProviders {
	s.UpdateTime = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyIdentityProviders) Validate() error {
	return dara.Validate(s)
}
