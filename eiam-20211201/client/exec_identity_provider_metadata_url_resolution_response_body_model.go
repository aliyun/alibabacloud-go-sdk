// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecIdentityProviderMetadataUrlResolutionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetIdentityProviderMetadata(v *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadata) *ExecIdentityProviderMetadataUrlResolutionResponseBody
  GetIdentityProviderMetadata() *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadata 
  SetRequestId(v string) *ExecIdentityProviderMetadataUrlResolutionResponseBody
  GetRequestId() *string 
}

type ExecIdentityProviderMetadataUrlResolutionResponseBody struct {
  IdentityProviderMetadata *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadata `json:"IdentityProviderMetadata,omitempty" xml:"IdentityProviderMetadata,omitempty" type:"Struct"`
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecIdentityProviderMetadataUrlResolutionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecIdentityProviderMetadataUrlResolutionResponseBody) GoString() string {
  return s.String()
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBody) GetIdentityProviderMetadata() *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadata  {
  return s.IdentityProviderMetadata
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBody) SetIdentityProviderMetadata(v *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadata) *ExecIdentityProviderMetadataUrlResolutionResponseBody {
  s.IdentityProviderMetadata = v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBody) SetRequestId(v string) *ExecIdentityProviderMetadataUrlResolutionResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBody) Validate() error {
  if s.IdentityProviderMetadata != nil {
    if err := s.IdentityProviderMetadata.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadata struct {
  // OIDC IdP的Meta信息。
  OidcOpenIdConfiguration *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration `json:"OidcOpenIdConfiguration,omitempty" xml:"OidcOpenIdConfiguration,omitempty" type:"Struct"`
  SamlMetadataConfiguration *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration `json:"SamlMetadataConfiguration,omitempty" xml:"SamlMetadataConfiguration,omitempty" type:"Struct"`
}

func (s ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadata) String() string {
  return dara.Prettify(s)
}

func (s ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadata) GoString() string {
  return s.String()
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadata) GetOidcOpenIdConfiguration() *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration  {
  return s.OidcOpenIdConfiguration
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadata) GetSamlMetadataConfiguration() *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration  {
  return s.SamlMetadataConfiguration
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadata) SetOidcOpenIdConfiguration(v *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration) *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadata {
  s.OidcOpenIdConfiguration = v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadata) SetSamlMetadataConfiguration(v *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration) *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadata {
  s.SamlMetadataConfiguration = v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadata) Validate() error {
  if s.OidcOpenIdConfiguration != nil {
    if err := s.OidcOpenIdConfiguration.Validate(); err != nil {
      return err
    }
  }
  if s.SamlMetadataConfiguration != nil {
    if err := s.SamlMetadataConfiguration.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration struct {
  // oAuth2 授权端点。
  // 
  // example:
  // 
  // https://demo.com/oauth2/default/v1/authorize
  AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
  // OIDC issuer信息。
  // 
  // example:
  // 
  // https://demo.com/fe974231-3454-4b70-9326-70fb71e41bce/v2.0/
  Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
  // OIDC jwks地址。
  // 
  // example:
  // 
  // https://demo.com/oauth2/v1/keys
  JwksUri *string `json:"JwksUri,omitempty" xml:"JwksUri,omitempty"`
  // oAuth2 Token端点。
  // 
  // example:
  // 
  // https://demo.com/api/bff/v1.2/developer/oidc/token
  TokenEndpoint *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
  // OIDC 用户信息端点。
  // 
  // example:
  // 
  // https://demo.com/api/bff/v1.2/developer/oidc/userinfo
  UserinfoEndpoint *string `json:"UserinfoEndpoint,omitempty" xml:"UserinfoEndpoint,omitempty"`
}

func (s ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration) String() string {
  return dara.Prettify(s)
}

func (s ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration) GoString() string {
  return s.String()
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration) GetAuthorizationEndpoint() *string  {
  return s.AuthorizationEndpoint
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration) GetIssuer() *string  {
  return s.Issuer
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration) GetJwksUri() *string  {
  return s.JwksUri
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration) GetTokenEndpoint() *string  {
  return s.TokenEndpoint
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration) GetUserinfoEndpoint() *string  {
  return s.UserinfoEndpoint
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration) SetAuthorizationEndpoint(v string) *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration {
  s.AuthorizationEndpoint = &v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration) SetIssuer(v string) *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration {
  s.Issuer = &v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration) SetJwksUri(v string) *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration {
  s.JwksUri = &v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration) SetTokenEndpoint(v string) *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration {
  s.TokenEndpoint = &v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration) SetUserinfoEndpoint(v string) *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration {
  s.UserinfoEndpoint = &v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataOidcOpenIdConfiguration) Validate() error {
  return dara.Validate(s)
}

type ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration struct {
  Certificates []*ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfigurationCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
  // example:
  // 
  // http://dc.test.com/adfs/services/trust
  IdPEntityId *string `json:"IdPEntityId,omitempty" xml:"IdPEntityId,omitempty"`
  // example:
  // 
  // https://dc.test.com/adfs/ls/
  IdPSsoUrl *string `json:"IdPSsoUrl,omitempty" xml:"IdPSsoUrl,omitempty"`
}

func (s ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration) String() string {
  return dara.Prettify(s)
}

func (s ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration) GoString() string {
  return s.String()
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration) GetCertificates() []*ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfigurationCertificates  {
  return s.Certificates
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration) GetIdPEntityId() *string  {
  return s.IdPEntityId
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration) GetIdPSsoUrl() *string  {
  return s.IdPSsoUrl
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration) SetCertificates(v []*ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfigurationCertificates) *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration {
  s.Certificates = v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration) SetIdPEntityId(v string) *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration {
  s.IdPEntityId = &v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration) SetIdPSsoUrl(v string) *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration {
  s.IdPSsoUrl = &v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfiguration) Validate() error {
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

type ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfigurationCertificates struct {
  // example:
  // 
  // -----BEGIN CERTIFICATE----- MIIE+zCCA0egAwIBAgIJAJZY0ZY0ZY0Z -----END CERTIFICATE-----
  Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfigurationCertificates) String() string {
  return dara.Prettify(s)
}

func (s ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfigurationCertificates) GoString() string {
  return s.String()
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfigurationCertificates) GetContent() *string  {
  return s.Content
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfigurationCertificates) SetContent(v string) *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfigurationCertificates {
  s.Content = &v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponseBodyIdentityProviderMetadataSamlMetadataConfigurationCertificates) Validate() error {
  return dara.Validate(s)
}

