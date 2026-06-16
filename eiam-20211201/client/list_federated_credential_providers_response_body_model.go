// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFederatedCredentialProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFederatedCredentialProviders(v []*ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) *ListFederatedCredentialProvidersResponseBody
	GetFederatedCredentialProviders() []*ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders
	SetMaxResults(v int32) *ListFederatedCredentialProvidersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListFederatedCredentialProvidersResponseBody
	GetNextToken() *string
	SetPreviousToken(v string) *ListFederatedCredentialProvidersResponseBody
	GetPreviousToken() *string
	SetRequestId(v string) *ListFederatedCredentialProvidersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListFederatedCredentialProvidersResponseBody
	GetTotalCount() *int32
}

type ListFederatedCredentialProvidersResponseBody struct {
	// The list of federated credential providers.
	FederatedCredentialProviders []*ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders `json:"FederatedCredentialProviders,omitempty" xml:"FederatedCredentialProviders,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to retrieve the next page of results. This parameter is empty if all results have been returned.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The token for the previous page of results.
	//
	// example:
	//
	// PTxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
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
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFederatedCredentialProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFederatedCredentialProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListFederatedCredentialProvidersResponseBody) GetFederatedCredentialProviders() []*ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders {
	return s.FederatedCredentialProviders
}

func (s *ListFederatedCredentialProvidersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListFederatedCredentialProvidersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFederatedCredentialProvidersResponseBody) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListFederatedCredentialProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFederatedCredentialProvidersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFederatedCredentialProvidersResponseBody) SetFederatedCredentialProviders(v []*ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) *ListFederatedCredentialProvidersResponseBody {
	s.FederatedCredentialProviders = v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBody) SetMaxResults(v int32) *ListFederatedCredentialProvidersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBody) SetNextToken(v string) *ListFederatedCredentialProvidersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBody) SetPreviousToken(v string) *ListFederatedCredentialProvidersResponseBody {
	s.PreviousToken = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBody) SetRequestId(v string) *ListFederatedCredentialProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBody) SetTotalCount(v int32) *ListFederatedCredentialProvidersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBody) Validate() error {
	if s.FederatedCredentialProviders != nil {
		for _, item := range s.FederatedCredentialProviders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders struct {
	CloudIdPProviderConfig *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersCloudIdPProviderConfig `json:"CloudIdPProviderConfig,omitempty" xml:"CloudIdPProviderConfig,omitempty" type:"Struct"`
	// The provider\\"s creation time.
	//
	// example:
	//
	// 1729061324000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The provider\\"s description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the federated credential provider.
	//
	// example:
	//
	// fcp_asda123XXX
	FederatedCredentialProviderId *string `json:"FederatedCredentialProviderId,omitempty" xml:"FederatedCredentialProviderId,omitempty"`
	// The name of the federated credential provider.
	//
	// example:
	//
	// pkcs7test
	FederatedCredentialProviderName *string `json:"FederatedCredentialProviderName,omitempty" xml:"FederatedCredentialProviderName,omitempty"`
	// The type of the federated credential provider.
	//
	// example:
	//
	// pkcs7
	FederatedCredentialProviderType *string `json:"FederatedCredentialProviderType,omitempty" xml:"FederatedCredentialProviderType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_dd4n3rnknybjjxuu5gq6ovqxXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the network access endpoint.
	//
	// example:
	//
	// inae_public
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// The OpenID Connect (OIDC) configuration.
	OidcProviderConfig *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig `json:"OidcProviderConfig,omitempty" xml:"OidcProviderConfig,omitempty" type:"Struct"`
	// The PKCS7 configuration.
	Pkcs7ProviderConfig *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig `json:"Pkcs7ProviderConfig,omitempty" xml:"Pkcs7ProviderConfig,omitempty" type:"Struct"`
	// The private CA configuration.
	PrivateCaProviderConfig *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig `json:"PrivateCaProviderConfig,omitempty" xml:"PrivateCaProviderConfig,omitempty" type:"Struct"`
	// The provider\\"s status.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The provider\\"s last update time.
	//
	// example:
	//
	// 1729061324000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) String() string {
	return dara.Prettify(s)
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) GoString() string {
	return s.String()
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) GetCloudIdPProviderConfig() *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersCloudIdPProviderConfig {
	return s.CloudIdPProviderConfig
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) GetDescription() *string {
	return s.Description
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) GetFederatedCredentialProviderName() *string {
	return s.FederatedCredentialProviderName
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) GetFederatedCredentialProviderType() *string {
	return s.FederatedCredentialProviderType
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) GetOidcProviderConfig() *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig {
	return s.OidcProviderConfig
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) GetPkcs7ProviderConfig() *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig {
	return s.Pkcs7ProviderConfig
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) GetPrivateCaProviderConfig() *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig {
	return s.PrivateCaProviderConfig
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) GetStatus() *string {
	return s.Status
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) SetCloudIdPProviderConfig(v *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersCloudIdPProviderConfig) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders {
	s.CloudIdPProviderConfig = v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) SetCreateTime(v int64) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders {
	s.CreateTime = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) SetDescription(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders {
	s.Description = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) SetFederatedCredentialProviderId(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) SetFederatedCredentialProviderName(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders {
	s.FederatedCredentialProviderName = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) SetFederatedCredentialProviderType(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders {
	s.FederatedCredentialProviderType = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) SetInstanceId(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders {
	s.InstanceId = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) SetNetworkAccessEndpointId(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) SetOidcProviderConfig(v *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders {
	s.OidcProviderConfig = v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) SetPkcs7ProviderConfig(v *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders {
	s.Pkcs7ProviderConfig = v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) SetPrivateCaProviderConfig(v *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders {
	s.PrivateCaProviderConfig = v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) SetStatus(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders {
	s.Status = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) SetUpdateTime(v int64) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders {
	s.UpdateTime = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProviders) Validate() error {
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

type ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersCloudIdPProviderConfig struct {
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersCloudIdPProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersCloudIdPProviderConfig) GoString() string {
	return s.String()
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersCloudIdPProviderConfig) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersCloudIdPProviderConfig) SetIdentityProviderId(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersCloudIdPProviderConfig {
	s.IdentityProviderId = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersCloudIdPProviderConfig) Validate() error {
	return dara.Validate(s)
}

type ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig struct {
	// The list of audiences for the OIDC credential.
	Audiences []*string `json:"Audiences,omitempty" xml:"Audiences,omitempty" type:"Repeated"`
	// The dynamically obtained JWKS.
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
	DynamicJwks *string `json:"DynamicJwks,omitempty" xml:"DynamicJwks,omitempty"`
	// The issuer.
	//
	// example:
	//
	// https://example.com
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The timestamp of the last JWKS retrieval.
	//
	// example:
	//
	// 1729061324000
	JwksLastObtainedTime *int64 `json:"JwksLastObtainedTime,omitempty" xml:"JwksLastObtainedTime,omitempty"`
	// The JWKS source.
	//
	// example:
	//
	// static
	JwksSource *string `json:"JwksSource,omitempty" xml:"JwksSource,omitempty"`
	// The JSON Web Key Set (JWKS) endpoint.
	//
	// example:
	//
	// https://example.com
	JwksUri *string `json:"JwksUri,omitempty" xml:"JwksUri,omitempty"`
	// The statically obtained JWKS.
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
	// The trust condition.
	//
	// example:
	//
	// IsNullOrEmpty("jwt.issuer")
	TrustCondition *string `json:"TrustCondition,omitempty" xml:"TrustCondition,omitempty"`
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) GoString() string {
	return s.String()
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) GetAudiences() []*string {
	return s.Audiences
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) GetDynamicJwks() *string {
	return s.DynamicJwks
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) GetIssuer() *string {
	return s.Issuer
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) GetJwksLastObtainedTime() *int64 {
	return s.JwksLastObtainedTime
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) GetJwksSource() *string {
	return s.JwksSource
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) GetJwksUri() *string {
	return s.JwksUri
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) GetStaticJwks() *string {
	return s.StaticJwks
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) GetTrustCondition() *string {
	return s.TrustCondition
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) SetAudiences(v []*string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig {
	s.Audiences = v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) SetDynamicJwks(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig {
	s.DynamicJwks = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) SetIssuer(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig {
	s.Issuer = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) SetJwksLastObtainedTime(v int64) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig {
	s.JwksLastObtainedTime = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) SetJwksSource(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig {
	s.JwksSource = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) SetJwksUri(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig {
	s.JwksUri = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) SetStaticJwks(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig {
	s.StaticJwks = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) SetTrustCondition(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig {
	s.TrustCondition = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersOidcProviderConfig) Validate() error {
	return dara.Validate(s)
}

type ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig struct {
	// A list of PKCS7 certificates.
	Certificates []*ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The Cryptographic Message Syntax (CMS) verification mode.
	//
	// example:
	//
	// cert
	CmsVerificationMode *string `json:"CmsVerificationMode,omitempty" xml:"CmsVerificationMode,omitempty"`
	// The validity period of the signature.
	//
	// example:
	//
	// 3600
	SignatureEffectiveTime *int64 `json:"SignatureEffectiveTime,omitempty" xml:"SignatureEffectiveTime,omitempty"`
	// The expression used to obtain the signing time.
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
	// The trust condition.
	//
	// example:
	//
	// IsNullOrEmpty("certNo")
	TrustCondition *string `json:"TrustCondition,omitempty" xml:"TrustCondition,omitempty"`
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) GoString() string {
	return s.String()
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) GetCertificates() []*ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates {
	return s.Certificates
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) GetCmsVerificationMode() *string {
	return s.CmsVerificationMode
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) GetSignatureEffectiveTime() *int64 {
	return s.SignatureEffectiveTime
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) GetSigningTimeValueExpression() *string {
	return s.SigningTimeValueExpression
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) GetTrustAnchorSource() *string {
	return s.TrustAnchorSource
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) GetTrustCondition() *string {
	return s.TrustCondition
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) SetCertificates(v []*ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig {
	s.Certificates = v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) SetCmsVerificationMode(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig {
	s.CmsVerificationMode = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) SetSignatureEffectiveTime(v int64) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig {
	s.SignatureEffectiveTime = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) SetSigningTimeValueExpression(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig {
	s.SigningTimeValueExpression = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) SetTrustAnchorSource(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig {
	s.TrustAnchorSource = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) SetTrustCondition(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig {
	s.TrustCondition = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfig) Validate() error {
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

type ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates struct {
	// The metadata of the certificate.
	CertificateMetadata *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificatesCertificateMetadata `json:"CertificateMetadata,omitempty" xml:"CertificateMetadata,omitempty" type:"Struct"`
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
	// The certificate fingerprint.
	//
	// example:
	//
	// 2b18947a6a9fc7764fd8b5fb18a863b0c6daxxx
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates) String() string {
	return dara.Prettify(s)
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates) GoString() string {
	return s.String()
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates) GetCertificateMetadata() *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificatesCertificateMetadata {
	return s.CertificateMetadata
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates) GetContent() *string {
	return s.Content
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates) SetCertificateMetadata(v *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificatesCertificateMetadata) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates {
	s.CertificateMetadata = v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates) SetContent(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates {
	s.Content = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates) SetFingerprint(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates {
	s.Fingerprint = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificates) Validate() error {
	if s.CertificateMetadata != nil {
		if err := s.CertificateMetadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificatesCertificateMetadata struct {
	// The expiration time.
	//
	// example:
	//
	// 1729061324000
	NotAfter *int64 `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The validity start time.
	//
	// example:
	//
	// 1729061324000
	NotBefore *int64 `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificatesCertificateMetadata) String() string {
	return dara.Prettify(s)
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificatesCertificateMetadata) GoString() string {
	return s.String()
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificatesCertificateMetadata) GetNotAfter() *int64 {
	return s.NotAfter
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificatesCertificateMetadata) GetNotBefore() *int64 {
	return s.NotBefore
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificatesCertificateMetadata) SetNotAfter(v int64) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificatesCertificateMetadata {
	s.NotAfter = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificatesCertificateMetadata) SetNotBefore(v int64) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificatesCertificateMetadata {
	s.NotBefore = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPkcs7ProviderConfigCertificatesCertificateMetadata) Validate() error {
	return dara.Validate(s)
}

type ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig struct {
	// A list of root certificates.
	Certificates []*ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The method for obtaining the root certificate.
	//
	// example:
	//
	// custom
	TrustAnchorSource *string `json:"TrustAnchorSource,omitempty" xml:"TrustAnchorSource,omitempty"`
	// The trust condition.
	//
	// example:
	//
	// IsNullOrEmpty("certNo")
	TrustCondition *string `json:"TrustCondition,omitempty" xml:"TrustCondition,omitempty"`
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig) GoString() string {
	return s.String()
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig) GetCertificates() []*ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates {
	return s.Certificates
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig) GetTrustAnchorSource() *string {
	return s.TrustAnchorSource
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig) GetTrustCondition() *string {
	return s.TrustCondition
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig) SetCertificates(v []*ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig {
	s.Certificates = v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig) SetTrustAnchorSource(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig {
	s.TrustAnchorSource = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig) SetTrustCondition(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig {
	s.TrustCondition = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfig) Validate() error {
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

type ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates struct {
	// The metadata of the certificate.
	CertificateMetadata *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificatesCertificateMetadata `json:"CertificateMetadata,omitempty" xml:"CertificateMetadata,omitempty" type:"Struct"`
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

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates) String() string {
	return dara.Prettify(s)
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates) GoString() string {
	return s.String()
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates) GetCertificateMetadata() *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificatesCertificateMetadata {
	return s.CertificateMetadata
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates) GetContent() *string {
	return s.Content
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates) SetCertificateMetadata(v *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificatesCertificateMetadata) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates {
	s.CertificateMetadata = v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates) SetContent(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates {
	s.Content = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates) SetFingerprint(v string) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates {
	s.Fingerprint = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificates) Validate() error {
	if s.CertificateMetadata != nil {
		if err := s.CertificateMetadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificatesCertificateMetadata struct {
	// The expiration time.
	//
	// example:
	//
	// 1729061324000
	NotAfter *int64 `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The validity start time.
	//
	// example:
	//
	// 1729061324000
	NotBefore *int64 `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificatesCertificateMetadata) String() string {
	return dara.Prettify(s)
}

func (s ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificatesCertificateMetadata) GoString() string {
	return s.String()
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificatesCertificateMetadata) GetNotAfter() *int64 {
	return s.NotAfter
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificatesCertificateMetadata) GetNotBefore() *int64 {
	return s.NotBefore
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificatesCertificateMetadata) SetNotAfter(v int64) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificatesCertificateMetadata {
	s.NotAfter = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificatesCertificateMetadata) SetNotBefore(v int64) *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificatesCertificateMetadata {
	s.NotBefore = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponseBodyFederatedCredentialProvidersPrivateCaProviderConfigCertificatesCertificateMetadata) Validate() error {
	return dara.Validate(s)
}
