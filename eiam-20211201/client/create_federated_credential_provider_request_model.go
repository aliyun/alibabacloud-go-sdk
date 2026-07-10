// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFederatedCredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudIdPProviderConfig(v *CreateFederatedCredentialProviderRequestCloudIdPProviderConfig) *CreateFederatedCredentialProviderRequest
	GetCloudIdPProviderConfig() *CreateFederatedCredentialProviderRequestCloudIdPProviderConfig
	SetDescription(v string) *CreateFederatedCredentialProviderRequest
	GetDescription() *string
	SetFederatedCredentialProviderName(v string) *CreateFederatedCredentialProviderRequest
	GetFederatedCredentialProviderName() *string
	SetFederatedCredentialProviderType(v string) *CreateFederatedCredentialProviderRequest
	GetFederatedCredentialProviderType() *string
	SetInstanceId(v string) *CreateFederatedCredentialProviderRequest
	GetInstanceId() *string
	SetNetworkAccessEndpointId(v string) *CreateFederatedCredentialProviderRequest
	GetNetworkAccessEndpointId() *string
	SetOidcProviderConfig(v *CreateFederatedCredentialProviderRequestOidcProviderConfig) *CreateFederatedCredentialProviderRequest
	GetOidcProviderConfig() *CreateFederatedCredentialProviderRequestOidcProviderConfig
	SetPkcs7ProviderConfig(v *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) *CreateFederatedCredentialProviderRequest
	GetPkcs7ProviderConfig() *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig
	SetPrivateCaProviderConfig(v *CreateFederatedCredentialProviderRequestPrivateCaProviderConfig) *CreateFederatedCredentialProviderRequest
	GetPrivateCaProviderConfig() *CreateFederatedCredentialProviderRequestPrivateCaProviderConfig
}

type CreateFederatedCredentialProviderRequest struct {
	CloudIdPProviderConfig *CreateFederatedCredentialProviderRequestCloudIdPProviderConfig `json:"CloudIdPProviderConfig,omitempty" xml:"CloudIdPProviderConfig,omitempty" type:"Struct"`
	// The description of the federated trust source.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the federated trust source.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	FederatedCredentialProviderName *string `json:"FederatedCredentialProviderName,omitempty" xml:"FederatedCredentialProviderName,omitempty"`
	// The type of the federated trust source.
	//
	// This parameter is required.
	//
	// example:
	//
	// pkcs7
	FederatedCredentialProviderType *string `json:"FederatedCredentialProviderType,omitempty" xml:"FederatedCredentialProviderType,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the network access endpoint.
	//
	// example:
	//
	// nae_example_id
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// The OIDC configuration.
	OidcProviderConfig *CreateFederatedCredentialProviderRequestOidcProviderConfig `json:"OidcProviderConfig,omitempty" xml:"OidcProviderConfig,omitempty" type:"Struct"`
	// The PKCS7 configuration.
	Pkcs7ProviderConfig *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig `json:"Pkcs7ProviderConfig,omitempty" xml:"Pkcs7ProviderConfig,omitempty" type:"Struct"`
	// The private CA configuration.
	PrivateCaProviderConfig *CreateFederatedCredentialProviderRequestPrivateCaProviderConfig `json:"PrivateCaProviderConfig,omitempty" xml:"PrivateCaProviderConfig,omitempty" type:"Struct"`
}

func (s CreateFederatedCredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFederatedCredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *CreateFederatedCredentialProviderRequest) GetCloudIdPProviderConfig() *CreateFederatedCredentialProviderRequestCloudIdPProviderConfig {
	return s.CloudIdPProviderConfig
}

func (s *CreateFederatedCredentialProviderRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFederatedCredentialProviderRequest) GetFederatedCredentialProviderName() *string {
	return s.FederatedCredentialProviderName
}

func (s *CreateFederatedCredentialProviderRequest) GetFederatedCredentialProviderType() *string {
	return s.FederatedCredentialProviderType
}

func (s *CreateFederatedCredentialProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateFederatedCredentialProviderRequest) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *CreateFederatedCredentialProviderRequest) GetOidcProviderConfig() *CreateFederatedCredentialProviderRequestOidcProviderConfig {
	return s.OidcProviderConfig
}

func (s *CreateFederatedCredentialProviderRequest) GetPkcs7ProviderConfig() *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig {
	return s.Pkcs7ProviderConfig
}

func (s *CreateFederatedCredentialProviderRequest) GetPrivateCaProviderConfig() *CreateFederatedCredentialProviderRequestPrivateCaProviderConfig {
	return s.PrivateCaProviderConfig
}

func (s *CreateFederatedCredentialProviderRequest) SetCloudIdPProviderConfig(v *CreateFederatedCredentialProviderRequestCloudIdPProviderConfig) *CreateFederatedCredentialProviderRequest {
	s.CloudIdPProviderConfig = v
	return s
}

func (s *CreateFederatedCredentialProviderRequest) SetDescription(v string) *CreateFederatedCredentialProviderRequest {
	s.Description = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequest) SetFederatedCredentialProviderName(v string) *CreateFederatedCredentialProviderRequest {
	s.FederatedCredentialProviderName = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequest) SetFederatedCredentialProviderType(v string) *CreateFederatedCredentialProviderRequest {
	s.FederatedCredentialProviderType = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequest) SetInstanceId(v string) *CreateFederatedCredentialProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequest) SetNetworkAccessEndpointId(v string) *CreateFederatedCredentialProviderRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequest) SetOidcProviderConfig(v *CreateFederatedCredentialProviderRequestOidcProviderConfig) *CreateFederatedCredentialProviderRequest {
	s.OidcProviderConfig = v
	return s
}

func (s *CreateFederatedCredentialProviderRequest) SetPkcs7ProviderConfig(v *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) *CreateFederatedCredentialProviderRequest {
	s.Pkcs7ProviderConfig = v
	return s
}

func (s *CreateFederatedCredentialProviderRequest) SetPrivateCaProviderConfig(v *CreateFederatedCredentialProviderRequestPrivateCaProviderConfig) *CreateFederatedCredentialProviderRequest {
	s.PrivateCaProviderConfig = v
	return s
}

func (s *CreateFederatedCredentialProviderRequest) Validate() error {
	if s.CloudIdPProviderConfig != nil {
		if err := s.CloudIdPProviderConfig.Validate(); err != nil {
			return err
		}
	}
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

type CreateFederatedCredentialProviderRequestCloudIdPProviderConfig struct {
	// example:
	//
	// idp_m7hk2hbnew5ir3ycfabug6xxxx
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
}

func (s CreateFederatedCredentialProviderRequestCloudIdPProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateFederatedCredentialProviderRequestCloudIdPProviderConfig) GoString() string {
	return s.String()
}

func (s *CreateFederatedCredentialProviderRequestCloudIdPProviderConfig) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *CreateFederatedCredentialProviderRequestCloudIdPProviderConfig) SetIdentityProviderId(v string) *CreateFederatedCredentialProviderRequestCloudIdPProviderConfig {
	s.IdentityProviderId = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequestCloudIdPProviderConfig) Validate() error {
	return dara.Validate(s)
}

type CreateFederatedCredentialProviderRequestOidcProviderConfig struct {
	// The list of audiences.
	Audiences []*string `json:"Audiences,omitempty" xml:"Audiences,omitempty" type:"Repeated"`
	// Issuer
	//
	// example:
	//
	// https://example.com
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The source of the JWKS.
	//
	// example:
	//
	// static
	JwksSource *string `json:"JwksSource,omitempty" xml:"JwksSource,omitempty"`
	// The JWKS endpoint.
	//
	// example:
	//
	// https://example.com/jwks
	JwksUri *string `json:"JwksUri,omitempty" xml:"JwksUri,omitempty"`
	// The statically retrieved JWKS.
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
	//       "n": "qrsfFfSZngqKOxVE29ZIR4SXkwKq029B3HLDAZui_Pwaxwn8FssR9QdwsljZS06BTDp10vhPgqMB7s7TmHulL3I4WuSB-l4uXXXXX"
	//
	//     }
	//
	//   ]
	//
	// }
	StaticJwks *string `json:"StaticJwks,omitempty" xml:"StaticJwks,omitempty"`
	// The trust condition.
	//
	// example:
	//
	// IsNullOrEmpty("jwt.issuer")
	TrustCondition *string `json:"TrustCondition,omitempty" xml:"TrustCondition,omitempty"`
}

func (s CreateFederatedCredentialProviderRequestOidcProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateFederatedCredentialProviderRequestOidcProviderConfig) GoString() string {
	return s.String()
}

func (s *CreateFederatedCredentialProviderRequestOidcProviderConfig) GetAudiences() []*string {
	return s.Audiences
}

func (s *CreateFederatedCredentialProviderRequestOidcProviderConfig) GetIssuer() *string {
	return s.Issuer
}

func (s *CreateFederatedCredentialProviderRequestOidcProviderConfig) GetJwksSource() *string {
	return s.JwksSource
}

func (s *CreateFederatedCredentialProviderRequestOidcProviderConfig) GetJwksUri() *string {
	return s.JwksUri
}

func (s *CreateFederatedCredentialProviderRequestOidcProviderConfig) GetStaticJwks() *string {
	return s.StaticJwks
}

func (s *CreateFederatedCredentialProviderRequestOidcProviderConfig) GetTrustCondition() *string {
	return s.TrustCondition
}

func (s *CreateFederatedCredentialProviderRequestOidcProviderConfig) SetAudiences(v []*string) *CreateFederatedCredentialProviderRequestOidcProviderConfig {
	s.Audiences = v
	return s
}

func (s *CreateFederatedCredentialProviderRequestOidcProviderConfig) SetIssuer(v string) *CreateFederatedCredentialProviderRequestOidcProviderConfig {
	s.Issuer = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequestOidcProviderConfig) SetJwksSource(v string) *CreateFederatedCredentialProviderRequestOidcProviderConfig {
	s.JwksSource = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequestOidcProviderConfig) SetJwksUri(v string) *CreateFederatedCredentialProviderRequestOidcProviderConfig {
	s.JwksUri = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequestOidcProviderConfig) SetStaticJwks(v string) *CreateFederatedCredentialProviderRequestOidcProviderConfig {
	s.StaticJwks = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequestOidcProviderConfig) SetTrustCondition(v string) *CreateFederatedCredentialProviderRequestOidcProviderConfig {
	s.TrustCondition = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequestOidcProviderConfig) Validate() error {
	return dara.Validate(s)
}

type CreateFederatedCredentialProviderRequestPkcs7ProviderConfig struct {
	// The list of PKCS7 certificates.
	Certificates []*CreateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The CMS verification mode.
	//
	// example:
	//
	// cert_chain
	CmsVerificationMode *string `json:"CmsVerificationMode,omitempty" xml:"CmsVerificationMode,omitempty"`
	// The validity period of the signature.
	//
	// example:
	//
	// 1200
	SignatureEffectiveTime *int64 `json:"SignatureEffectiveTime,omitempty" xml:"SignatureEffectiveTime,omitempty"`
	// The expression used to retrieve the signing time.
	//
	// example:
	//
	// pkcs7.signingTime
	SigningTimeValueExpression *string `json:"SigningTimeValueExpression,omitempty" xml:"SigningTimeValueExpression,omitempty"`
	// The source of the certificate trust anchor.
	//
	// example:
	//
	// custom
	TrustAnchorSource *string `json:"TrustAnchorSource,omitempty" xml:"TrustAnchorSource,omitempty"`
	// The trust condition.
	//
	// example:
	//
	// IsNullOrEmpty("jwt.issuer")
	TrustCondition *string `json:"TrustCondition,omitempty" xml:"TrustCondition,omitempty"`
}

func (s CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) GoString() string {
	return s.String()
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) GetCertificates() []*CreateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates {
	return s.Certificates
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) GetCmsVerificationMode() *string {
	return s.CmsVerificationMode
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) GetSignatureEffectiveTime() *int64 {
	return s.SignatureEffectiveTime
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) GetSigningTimeValueExpression() *string {
	return s.SigningTimeValueExpression
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) GetTrustAnchorSource() *string {
	return s.TrustAnchorSource
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) GetTrustCondition() *string {
	return s.TrustCondition
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) SetCertificates(v []*CreateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates) *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig {
	s.Certificates = v
	return s
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) SetCmsVerificationMode(v string) *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig {
	s.CmsVerificationMode = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) SetSignatureEffectiveTime(v int64) *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig {
	s.SignatureEffectiveTime = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) SetSigningTimeValueExpression(v string) *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig {
	s.SigningTimeValueExpression = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) SetTrustAnchorSource(v string) *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig {
	s.TrustAnchorSource = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) SetTrustCondition(v string) *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig {
	s.TrustCondition = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfig) Validate() error {
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

type CreateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates struct {
	// The content of the root certificate.
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

func (s CreateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates) String() string {
	return dara.Prettify(s)
}

func (s CreateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates) GoString() string {
	return s.String()
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates) GetContent() *string {
	return s.Content
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates) SetContent(v string) *CreateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates {
	s.Content = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequestPkcs7ProviderConfigCertificates) Validate() error {
	return dara.Validate(s)
}

type CreateFederatedCredentialProviderRequestPrivateCaProviderConfig struct {
	// The list of root certificates.
	Certificates []*CreateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The method used to retrieve the root certificate.
	//
	// example:
	//
	// custom
	TrustAnchorSource *string `json:"TrustAnchorSource,omitempty" xml:"TrustAnchorSource,omitempty"`
	// The trust condition for the root certificate.
	//
	// example:
	//
	// IsNullOrEmpty("jwt.issuer")
	TrustCondition *string `json:"TrustCondition,omitempty" xml:"TrustCondition,omitempty"`
}

func (s CreateFederatedCredentialProviderRequestPrivateCaProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateFederatedCredentialProviderRequestPrivateCaProviderConfig) GoString() string {
	return s.String()
}

func (s *CreateFederatedCredentialProviderRequestPrivateCaProviderConfig) GetCertificates() []*CreateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates {
	return s.Certificates
}

func (s *CreateFederatedCredentialProviderRequestPrivateCaProviderConfig) GetTrustAnchorSource() *string {
	return s.TrustAnchorSource
}

func (s *CreateFederatedCredentialProviderRequestPrivateCaProviderConfig) GetTrustCondition() *string {
	return s.TrustCondition
}

func (s *CreateFederatedCredentialProviderRequestPrivateCaProviderConfig) SetCertificates(v []*CreateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates) *CreateFederatedCredentialProviderRequestPrivateCaProviderConfig {
	s.Certificates = v
	return s
}

func (s *CreateFederatedCredentialProviderRequestPrivateCaProviderConfig) SetTrustAnchorSource(v string) *CreateFederatedCredentialProviderRequestPrivateCaProviderConfig {
	s.TrustAnchorSource = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequestPrivateCaProviderConfig) SetTrustCondition(v string) *CreateFederatedCredentialProviderRequestPrivateCaProviderConfig {
	s.TrustCondition = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequestPrivateCaProviderConfig) Validate() error {
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

type CreateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates struct {
	// The content of the root certificate.
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

func (s CreateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates) String() string {
	return dara.Prettify(s)
}

func (s CreateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates) GoString() string {
	return s.String()
}

func (s *CreateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates) GetContent() *string {
	return s.Content
}

func (s *CreateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates) SetContent(v string) *CreateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates {
	s.Content = &v
	return s
}

func (s *CreateFederatedCredentialProviderRequestPrivateCaProviderConfigCertificates) Validate() error {
	return dara.Validate(s)
}
