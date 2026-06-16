// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFederatedCredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFederatedCredentialProvider(v *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) *GetFederatedCredentialProviderResponseBody
	GetFederatedCredentialProvider() *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider
	SetRequestId(v string) *GetFederatedCredentialProviderResponseBody
	GetRequestId() *string
}

type GetFederatedCredentialProviderResponseBody struct {
	// The federated credential provider.
	FederatedCredentialProvider *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider `json:"FederatedCredentialProvider,omitempty" xml:"FederatedCredentialProvider,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFederatedCredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFederatedCredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetFederatedCredentialProviderResponseBody) GetFederatedCredentialProvider() *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider {
	return s.FederatedCredentialProvider
}

func (s *GetFederatedCredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFederatedCredentialProviderResponseBody) SetFederatedCredentialProvider(v *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) *GetFederatedCredentialProviderResponseBody {
	s.FederatedCredentialProvider = v
	return s
}

func (s *GetFederatedCredentialProviderResponseBody) SetRequestId(v string) *GetFederatedCredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBody) Validate() error {
	if s.FederatedCredentialProvider != nil {
		if err := s.FederatedCredentialProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider struct {
	CloudIdPProviderConfig *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderCloudIdPProviderConfig `json:"CloudIdPProviderConfig,omitempty" xml:"CloudIdPProviderConfig,omitempty" type:"Struct"`
	// The time when the provider was created.
	//
	// example:
	//
	// 1729061324000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the federated credential provider.
	//
	// example:
	//
	// fcp_asd123XXX
	FederatedCredentialProviderId *string `json:"FederatedCredentialProviderId,omitempty" xml:"FederatedCredentialProviderId,omitempty"`
	// The name of the federated credential provider.
	//
	// example:
	//
	// test
	FederatedCredentialProviderName *string `json:"FederatedCredentialProviderName,omitempty" xml:"FederatedCredentialProviderName,omitempty"`
	// The type of the federated credential provider.
	//
	// example:
	//
	// pkcs7
	FederatedCredentialProviderType *string `json:"FederatedCredentialProviderType,omitempty" xml:"FederatedCredentialProviderType,omitempty"`
	// The ID of the EIAM instance.
	//
	// example:
	//
	// idaas_qlbbighyxbu42xl7eninojXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the network access endpoint.
	//
	// example:
	//
	// inae_public
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// The OpenID Connect (OIDC) configuration.
	OidcProviderConfig *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig `json:"OidcProviderConfig,omitempty" xml:"OidcProviderConfig,omitempty" type:"Struct"`
	// The PKCS#7 configuration.
	Pkcs7ProviderConfig *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig `json:"Pkcs7ProviderConfig,omitempty" xml:"Pkcs7ProviderConfig,omitempty" type:"Struct"`
	// The private certificate authority (CA) configuration.
	PrivateCaProviderConfig *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig `json:"PrivateCaProviderConfig,omitempty" xml:"PrivateCaProviderConfig,omitempty" type:"Struct"`
	// The status of the federated credential provider.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the provider was last updated.
	//
	// example:
	//
	// 1729061324000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) String() string {
	return dara.Prettify(s)
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) GoString() string {
	return s.String()
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) GetCloudIdPProviderConfig() *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderCloudIdPProviderConfig {
	return s.CloudIdPProviderConfig
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) GetDescription() *string {
	return s.Description
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) GetFederatedCredentialProviderName() *string {
	return s.FederatedCredentialProviderName
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) GetFederatedCredentialProviderType() *string {
	return s.FederatedCredentialProviderType
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) GetOidcProviderConfig() *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig {
	return s.OidcProviderConfig
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) GetPkcs7ProviderConfig() *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig {
	return s.Pkcs7ProviderConfig
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) GetPrivateCaProviderConfig() *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig {
	return s.PrivateCaProviderConfig
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) GetStatus() *string {
	return s.Status
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) SetCloudIdPProviderConfig(v *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderCloudIdPProviderConfig) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider {
	s.CloudIdPProviderConfig = v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) SetCreateTime(v int64) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider {
	s.CreateTime = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) SetDescription(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider {
	s.Description = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) SetFederatedCredentialProviderId(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) SetFederatedCredentialProviderName(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider {
	s.FederatedCredentialProviderName = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) SetFederatedCredentialProviderType(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider {
	s.FederatedCredentialProviderType = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) SetInstanceId(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider {
	s.InstanceId = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) SetNetworkAccessEndpointId(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) SetOidcProviderConfig(v *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider {
	s.OidcProviderConfig = v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) SetPkcs7ProviderConfig(v *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider {
	s.Pkcs7ProviderConfig = v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) SetPrivateCaProviderConfig(v *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider {
	s.PrivateCaProviderConfig = v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) SetStatus(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider {
	s.Status = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) SetUpdateTime(v int64) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider {
	s.UpdateTime = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProvider) Validate() error {
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

type GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderCloudIdPProviderConfig struct {
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderCloudIdPProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderCloudIdPProviderConfig) GoString() string {
	return s.String()
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderCloudIdPProviderConfig) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderCloudIdPProviderConfig) SetIdentityProviderId(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderCloudIdPProviderConfig {
	s.IdentityProviderId = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderCloudIdPProviderConfig) Validate() error {
	return dara.Validate(s)
}

type GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig struct {
	// A list of audiences for the OIDC credential.
	Audiences []*string `json:"Audiences,omitempty" xml:"Audiences,omitempty" type:"Repeated"`
	// The dynamically obtained JWKS.
	//
	// example:
	//
	// https://example.com
	DynamicJwks *string `json:"DynamicJwks,omitempty" xml:"DynamicJwks,omitempty"`
	// The issuer of the OIDC credential.
	//
	// example:
	//
	// http://test.idaas.com/v2/oauth
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The time of the last JWKS retrieval.
	//
	// example:
	//
	// 1729061324000
	JwksLastObtainedTime *int64 `json:"JwksLastObtainedTime,omitempty" xml:"JwksLastObtainedTime,omitempty"`
	// The source of the JSON Web Key Set (JWKS).
	//
	// example:
	//
	// static
	JwksSource *string `json:"JwksSource,omitempty" xml:"JwksSource,omitempty"`
	// The JWKS endpoint.
	//
	// example:
	//
	// https://example.com
	JwksUri *string `json:"JwksUri,omitempty" xml:"JwksUri,omitempty"`
	// The static JWKS content.
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
	// The trust condition for the OIDC provider.
	//
	// example:
	//
	// IsNullOrEmpty("jwt.issuer")
	TrustCondition *string `json:"TrustCondition,omitempty" xml:"TrustCondition,omitempty"`
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) GoString() string {
	return s.String()
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) GetAudiences() []*string {
	return s.Audiences
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) GetDynamicJwks() *string {
	return s.DynamicJwks
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) GetIssuer() *string {
	return s.Issuer
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) GetJwksLastObtainedTime() *int64 {
	return s.JwksLastObtainedTime
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) GetJwksSource() *string {
	return s.JwksSource
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) GetJwksUri() *string {
	return s.JwksUri
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) GetStaticJwks() *string {
	return s.StaticJwks
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) GetTrustCondition() *string {
	return s.TrustCondition
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) SetAudiences(v []*string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig {
	s.Audiences = v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) SetDynamicJwks(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig {
	s.DynamicJwks = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) SetIssuer(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig {
	s.Issuer = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) SetJwksLastObtainedTime(v int64) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig {
	s.JwksLastObtainedTime = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) SetJwksSource(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig {
	s.JwksSource = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) SetJwksUri(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig {
	s.JwksUri = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) SetStaticJwks(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig {
	s.StaticJwks = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) SetTrustCondition(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig {
	s.TrustCondition = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderOidcProviderConfig) Validate() error {
	return dara.Validate(s)
}

type GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig struct {
	// A list of PKCS#7 certificates.
	Certificates []*GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The Cryptographic Message Syntax (CMS) verification mode.
	//
	// example:
	//
	// cert
	CmsVerificationMode *string `json:"CmsVerificationMode,omitempty" xml:"CmsVerificationMode,omitempty"`
	// The validity period of the signature, in seconds.
	//
	// example:
	//
	// 3600
	SignatureEffectiveTime *int64 `json:"SignatureEffectiveTime,omitempty" xml:"SignatureEffectiveTime,omitempty"`
	// An expression that specifies the signing time.
	//
	// example:
	//
	// pkcs7.payload.jsonData.audience.signingTime
	SigningTimeValueExpression *string `json:"SigningTimeValueExpression,omitempty" xml:"SigningTimeValueExpression,omitempty"`
	// The source of the certificate trust anchor.
	//
	// example:
	//
	// alibaba_cloud
	TrustAnchorSource *string `json:"TrustAnchorSource,omitempty" xml:"TrustAnchorSource,omitempty"`
	// The trust condition for the PKCS#7 provider.
	//
	// example:
	//
	// IsNullOrEmpty("certNo")
	TrustCondition *string `json:"TrustCondition,omitempty" xml:"TrustCondition,omitempty"`
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) GoString() string {
	return s.String()
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) GetCertificates() []*GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates {
	return s.Certificates
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) GetCmsVerificationMode() *string {
	return s.CmsVerificationMode
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) GetSignatureEffectiveTime() *int64 {
	return s.SignatureEffectiveTime
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) GetSigningTimeValueExpression() *string {
	return s.SigningTimeValueExpression
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) GetTrustAnchorSource() *string {
	return s.TrustAnchorSource
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) GetTrustCondition() *string {
	return s.TrustCondition
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) SetCertificates(v []*GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig {
	s.Certificates = v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) SetCmsVerificationMode(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig {
	s.CmsVerificationMode = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) SetSignatureEffectiveTime(v int64) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig {
	s.SignatureEffectiveTime = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) SetSigningTimeValueExpression(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig {
	s.SigningTimeValueExpression = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) SetTrustAnchorSource(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig {
	s.TrustAnchorSource = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) SetTrustCondition(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig {
	s.TrustCondition = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfig) Validate() error {
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

type GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates struct {
	// The certificate metadata.
	CertificateMetadata *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificatesCertificateMetadata `json:"CertificateMetadata,omitempty" xml:"CertificateMetadata,omitempty" type:"Struct"`
	// The content of the certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIE+zCCA0egAwIBAgIJAJZY0ZY0ZY0Z
	//
	// -----END CERTIFICATE-----
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The fingerprint of the certificate.
	//
	// example:
	//
	// 2b18947a6a9fc7764fd8b5fb18a863b0c6daxxx
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates) String() string {
	return dara.Prettify(s)
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates) GoString() string {
	return s.String()
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates) GetCertificateMetadata() *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificatesCertificateMetadata {
	return s.CertificateMetadata
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates) GetContent() *string {
	return s.Content
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates) SetCertificateMetadata(v *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificatesCertificateMetadata) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates {
	s.CertificateMetadata = v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates) SetContent(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates {
	s.Content = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates) SetFingerprint(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates {
	s.Fingerprint = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificates) Validate() error {
	if s.CertificateMetadata != nil {
		if err := s.CertificateMetadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificatesCertificateMetadata struct {
	// The time when the certificate expires.
	//
	// example:
	//
	// 1729061324000
	NotAfter *int64 `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The time when the certificate becomes valid.
	//
	// example:
	//
	// 1729061324000
	NotBefore *int64 `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificatesCertificateMetadata) String() string {
	return dara.Prettify(s)
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificatesCertificateMetadata) GoString() string {
	return s.String()
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificatesCertificateMetadata) GetNotAfter() *int64 {
	return s.NotAfter
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificatesCertificateMetadata) GetNotBefore() *int64 {
	return s.NotBefore
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificatesCertificateMetadata) SetNotAfter(v int64) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificatesCertificateMetadata {
	s.NotAfter = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificatesCertificateMetadata) SetNotBefore(v int64) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificatesCertificateMetadata {
	s.NotBefore = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPkcs7ProviderConfigCertificatesCertificateMetadata) Validate() error {
	return dara.Validate(s)
}

type GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig struct {
	// A list of root certificates.
	Certificates []*GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The method for obtaining the root certificate.
	//
	// example:
	//
	// custom
	TrustAnchorSource *string `json:"TrustAnchorSource,omitempty" xml:"TrustAnchorSource,omitempty"`
	// The trust condition for the root certificate.
	//
	// example:
	//
	// IsNullOrEmpty("certNo")
	TrustCondition *string `json:"TrustCondition,omitempty" xml:"TrustCondition,omitempty"`
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig) GoString() string {
	return s.String()
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig) GetCertificates() []*GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates {
	return s.Certificates
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig) GetTrustAnchorSource() *string {
	return s.TrustAnchorSource
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig) GetTrustCondition() *string {
	return s.TrustCondition
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig) SetCertificates(v []*GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig {
	s.Certificates = v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig) SetTrustAnchorSource(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig {
	s.TrustAnchorSource = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig) SetTrustCondition(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig {
	s.TrustCondition = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfig) Validate() error {
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

type GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates struct {
	// The certificate metadata.
	CertificateMetadata *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificatesCertificateMetadata `json:"CertificateMetadata,omitempty" xml:"CertificateMetadata,omitempty" type:"Struct"`
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
	// The fingerprint of the root certificate.
	//
	// example:
	//
	// 2b18947a6a9fc7764fd8b5fb18a863b0c6daxxx
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates) String() string {
	return dara.Prettify(s)
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates) GoString() string {
	return s.String()
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates) GetCertificateMetadata() *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificatesCertificateMetadata {
	return s.CertificateMetadata
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates) GetContent() *string {
	return s.Content
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates) SetCertificateMetadata(v *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificatesCertificateMetadata) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates {
	s.CertificateMetadata = v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates) SetContent(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates {
	s.Content = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates) SetFingerprint(v string) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates {
	s.Fingerprint = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificates) Validate() error {
	if s.CertificateMetadata != nil {
		if err := s.CertificateMetadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificatesCertificateMetadata struct {
	// The time when the certificate expires.
	//
	// example:
	//
	// 1729061324000
	NotAfter *int64 `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The time when the certificate becomes valid.
	//
	// example:
	//
	// 1729061324000
	NotBefore *int64 `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificatesCertificateMetadata) String() string {
	return dara.Prettify(s)
}

func (s GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificatesCertificateMetadata) GoString() string {
	return s.String()
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificatesCertificateMetadata) GetNotAfter() *int64 {
	return s.NotAfter
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificatesCertificateMetadata) GetNotBefore() *int64 {
	return s.NotBefore
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificatesCertificateMetadata) SetNotAfter(v int64) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificatesCertificateMetadata {
	s.NotAfter = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificatesCertificateMetadata) SetNotBefore(v int64) *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificatesCertificateMetadata {
	s.NotBefore = &v
	return s
}

func (s *GetFederatedCredentialProviderResponseBodyFederatedCredentialProviderPrivateCaProviderConfigCertificatesCertificateMetadata) Validate() error {
	return dara.Validate(s)
}
