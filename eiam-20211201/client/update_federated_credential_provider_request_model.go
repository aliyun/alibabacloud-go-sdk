// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFederatedCredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFederatedCredentialProviderId(v string) *UpdateFederatedCredentialProviderRequest
	GetFederatedCredentialProviderId() *string
	SetFederatedCredentialProviderName(v string) *UpdateFederatedCredentialProviderRequest
	GetFederatedCredentialProviderName() *string
	SetInstanceId(v string) *UpdateFederatedCredentialProviderRequest
	GetInstanceId() *string
	SetNetworkAccessEndpointId(v string) *UpdateFederatedCredentialProviderRequest
	GetNetworkAccessEndpointId() *string
	SetOidcProviderConfig(v *UpdateFederatedCredentialProviderRequestOidcProviderConfig) *UpdateFederatedCredentialProviderRequest
	GetOidcProviderConfig() *UpdateFederatedCredentialProviderRequestOidcProviderConfig
	SetPkcs7ProviderConfig(v *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) *UpdateFederatedCredentialProviderRequest
	GetPkcs7ProviderConfig() *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig
	SetPrivateCaProviderConfig(v *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig) *UpdateFederatedCredentialProviderRequest
	GetPrivateCaProviderConfig() *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig
}

type UpdateFederatedCredentialProviderRequest struct {
	// 联邦凭证提供方ID
	//
	// This parameter is required.
	//
	// example:
	//
	// fcp_mkv7rgt4d7i4u7zqtzev2mxxxx
	FederatedCredentialProviderId *string `json:"FederatedCredentialProviderId,omitempty" xml:"FederatedCredentialProviderId,omitempty"`
	// 联邦凭证提供方名称
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	FederatedCredentialProviderName *string `json:"FederatedCredentialProviderName,omitempty" xml:"FederatedCredentialProviderName,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 网络端点ID
	//
	// example:
	//
	// nae_public
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// OIDC配置
	OidcProviderConfig *UpdateFederatedCredentialProviderRequestOidcProviderConfig `json:"OidcProviderConfig,omitempty" xml:"OidcProviderConfig,omitempty" type:"Struct"`
	// PKCS7配置
	Pkcs7ProviderConfig *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig `json:"Pkcs7ProviderConfig,omitempty" xml:"Pkcs7ProviderConfig,omitempty" type:"Struct"`
	// 私有CA配置
	PrivateCaProviderConfig *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig `json:"PrivateCaProviderConfig,omitempty" xml:"PrivateCaProviderConfig,omitempty" type:"Struct"`
}

func (s UpdateFederatedCredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFederatedCredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateFederatedCredentialProviderRequest) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *UpdateFederatedCredentialProviderRequest) GetFederatedCredentialProviderName() *string {
	return s.FederatedCredentialProviderName
}

func (s *UpdateFederatedCredentialProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateFederatedCredentialProviderRequest) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *UpdateFederatedCredentialProviderRequest) GetOidcProviderConfig() *UpdateFederatedCredentialProviderRequestOidcProviderConfig {
	return s.OidcProviderConfig
}

func (s *UpdateFederatedCredentialProviderRequest) GetPkcs7ProviderConfig() *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig {
	return s.Pkcs7ProviderConfig
}

func (s *UpdateFederatedCredentialProviderRequest) GetPrivateCaProviderConfig() *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig {
	return s.PrivateCaProviderConfig
}

func (s *UpdateFederatedCredentialProviderRequest) SetFederatedCredentialProviderId(v string) *UpdateFederatedCredentialProviderRequest {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequest) SetFederatedCredentialProviderName(v string) *UpdateFederatedCredentialProviderRequest {
	s.FederatedCredentialProviderName = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequest) SetInstanceId(v string) *UpdateFederatedCredentialProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequest) SetNetworkAccessEndpointId(v string) *UpdateFederatedCredentialProviderRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequest) SetOidcProviderConfig(v *UpdateFederatedCredentialProviderRequestOidcProviderConfig) *UpdateFederatedCredentialProviderRequest {
	s.OidcProviderConfig = v
	return s
}

func (s *UpdateFederatedCredentialProviderRequest) SetPkcs7ProviderConfig(v *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) *UpdateFederatedCredentialProviderRequest {
	s.Pkcs7ProviderConfig = v
	return s
}

func (s *UpdateFederatedCredentialProviderRequest) SetPrivateCaProviderConfig(v *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig) *UpdateFederatedCredentialProviderRequest {
	s.PrivateCaProviderConfig = v
	return s
}

func (s *UpdateFederatedCredentialProviderRequest) Validate() error {
	if s.OidcProviderConfig != nil {
		if err := s.OidcProviderConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Pkcs7ProviderConfig != nil {
		if err := s.Pkcs7ProviderConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateCaProviderConfig != nil {
		if err := s.PrivateCaProviderConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateFederatedCredentialProviderRequestOidcProviderConfig struct {
	Audiences []*string `json:"Audiences,omitempty" xml:"Audiences,omitempty" type:"Repeated"`
	// Jwks来源
	//
	// This parameter is required.
	//
	// example:
	//
	// static
	JwksSource *string `json:"JwksSource,omitempty" xml:"JwksSource,omitempty"`
	// JWKS 端点
	//
	// example:
	//
	// https://example.com/jwks
	JwksUri *string `json:"JwksUri,omitempty" xml:"JwksUri,omitempty"`
	// 静态获取的jwks
	//
	// example:
	//
	// {
	//
	//   "keys": [
	//
	//     {
	//
	//       "kty": "RSA",
	//
	//       "e": "AQAB",
	//
	//       "use": "sig",
	//
	//       "kid": "KEY2RzsjRrimRASiAhCjBo18YwDoxpYHnHtv",
	//
	//       "n": "qrsfFfSZngqKOxVE29ZIR4SXkwKq029B3HLDAZui_Pwaxwn8FssR9QdwsljZS06BTDp10vhPgqMB7s7TmHulL3I4WuSB-l4uXTXXXX"
	//
	//     }
	//
	//   ]
	//
	// }
	StaticJwks *string `json:"StaticJwks,omitempty" xml:"StaticJwks,omitempty"`
	// 信任条件
	//
	// example:
	//
	// IsNullOrEmpty("jwt.issuer")
	TrustCondition *string `json:"TrustCondition,omitempty" xml:"TrustCondition,omitempty"`
}

func (s UpdateFederatedCredentialProviderRequestOidcProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateFederatedCredentialProviderRequestOidcProviderConfig) GoString() string {
	return s.String()
}

func (s *UpdateFederatedCredentialProviderRequestOidcProviderConfig) GetAudiences() []*string {
	return s.Audiences
}

func (s *UpdateFederatedCredentialProviderRequestOidcProviderConfig) GetJwksSource() *string {
	return s.JwksSource
}

func (s *UpdateFederatedCredentialProviderRequestOidcProviderConfig) GetJwksUri() *string {
	return s.JwksUri
}

func (s *UpdateFederatedCredentialProviderRequestOidcProviderConfig) GetStaticJwks() *string {
	return s.StaticJwks
}

func (s *UpdateFederatedCredentialProviderRequestOidcProviderConfig) GetTrustCondition() *string {
	return s.TrustCondition
}

func (s *UpdateFederatedCredentialProviderRequestOidcProviderConfig) SetAudiences(v []*string) *UpdateFederatedCredentialProviderRequestOidcProviderConfig {
	s.Audiences = v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestOidcProviderConfig) SetJwksSource(v string) *UpdateFederatedCredentialProviderRequestOidcProviderConfig {
	s.JwksSource = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestOidcProviderConfig) SetJwksUri(v string) *UpdateFederatedCredentialProviderRequestOidcProviderConfig {
	s.JwksUri = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestOidcProviderConfig) SetStaticJwks(v string) *UpdateFederatedCredentialProviderRequestOidcProviderConfig {
	s.StaticJwks = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestOidcProviderConfig) SetTrustCondition(v string) *UpdateFederatedCredentialProviderRequestOidcProviderConfig {
	s.TrustCondition = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestOidcProviderConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig struct {
	// pkcs7证书列表
	Certificates []*UpdateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// CMS验证模式
	//
	// example:
	//
	// cert
	CmsVerificationMode *string `json:"CmsVerificationMode,omitempty" xml:"CmsVerificationMode,omitempty"`
	// 签名有效期, 单位秒，1200
	//
	// example:
	//
	// 1200
	SignatureEffectiveTime *int64 `json:"SignatureEffectiveTime,omitempty" xml:"SignatureEffectiveTime,omitempty"`
	// example:
	//
	// pkcs7.signingTime
	SigningTimeValueExpression *string `json:"SigningTimeValueExpression,omitempty" xml:"SigningTimeValueExpression,omitempty"`
	// 证书信任锚点来源
	//
	// This parameter is required.
	//
	// example:
	//
	// custom
	TrustAnchorSource *string `json:"TrustAnchorSource,omitempty" xml:"TrustAnchorSource,omitempty"`
	// 信任条件
	//
	// example:
	//
	// IsNullOrEmpty("jwt.issuer")
	TrustCondition *string `json:"TrustCondition,omitempty" xml:"TrustCondition,omitempty"`
}

func (s UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) GoString() string {
	return s.String()
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) GetCertificates() []*UpdateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates {
	return s.Certificates
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) GetCmsVerificationMode() *string {
	return s.CmsVerificationMode
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) GetSignatureEffectiveTime() *int64 {
	return s.SignatureEffectiveTime
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) GetSigningTimeValueExpression() *string {
	return s.SigningTimeValueExpression
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) GetTrustAnchorSource() *string {
	return s.TrustAnchorSource
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) GetTrustCondition() *string {
	return s.TrustCondition
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) SetCertificates(v []*UpdateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates) *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig {
	s.Certificates = v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) SetCmsVerificationMode(v string) *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig {
	s.CmsVerificationMode = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) SetSignatureEffectiveTime(v int64) *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig {
	s.SignatureEffectiveTime = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) SetSigningTimeValueExpression(v string) *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig {
	s.SigningTimeValueExpression = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) SetTrustAnchorSource(v string) *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig {
	s.TrustAnchorSource = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) SetTrustCondition(v string) *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig {
	s.TrustCondition = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfig) Validate() error {
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

type UpdateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates struct {
	// Root证书内容
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIE+zCCA0egAwIBAgIJAJZY0ZY0ZY0Z
	//
	// -----END CERTIFICATE-----
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s UpdateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates) String() string {
	return dara.Prettify(s)
}

func (s UpdateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates) GoString() string {
	return s.String()
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates) GetContent() *string {
	return s.Content
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates) SetContent(v string) *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates {
	s.Content = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates) Validate() error {
	return dara.Validate(s)
}

type UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig struct {
	// Root证书列表
	Certificates []*UpdateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// Root证书获取方式
	//
	// This parameter is required.
	//
	// example:
	//
	// custom
	TrustAnchorSource *string `json:"TrustAnchorSource,omitempty" xml:"TrustAnchorSource,omitempty"`
	// Root证书的信任条件
	//
	// example:
	//
	// IsNullOrEmpty("jwt.issuer")
	TrustCondition *string `json:"TrustCondition,omitempty" xml:"TrustCondition,omitempty"`
}

func (s UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig) GoString() string {
	return s.String()
}

func (s *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig) GetCertificates() []*UpdateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates {
	return s.Certificates
}

func (s *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig) GetTrustAnchorSource() *string {
	return s.TrustAnchorSource
}

func (s *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig) GetTrustCondition() *string {
	return s.TrustCondition
}

func (s *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig) SetCertificates(v []*UpdateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates) *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig {
	s.Certificates = v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig) SetTrustAnchorSource(v string) *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig {
	s.TrustAnchorSource = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig) SetTrustCondition(v string) *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig {
	s.TrustCondition = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfig) Validate() error {
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

type UpdateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates struct {
	// Root证书内容
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIE+zCCA0egAwIBAgIJAJZY0ZY0ZY0Z
	//
	// -----END CERTIFICATE-----
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s UpdateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates) String() string {
	return dara.Prettify(s)
}

func (s UpdateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates) GoString() string {
	return s.String()
}

func (s *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates) GetContent() *string {
	return s.Content
}

func (s *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates) SetContent(v string) *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates {
	s.Content = &v
	return s
}

func (s *UpdateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates) Validate() error {
	return dara.Validate(s)
}
