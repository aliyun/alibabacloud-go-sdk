// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTokenByAuthorizationServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredentialName(v string) *GenerateTokenByAuthorizationServerRequest
	GetApplicationFederatedCredentialName() *string
	SetClientAssertion(v string) *GenerateTokenByAuthorizationServerRequest
	GetClientAssertion() *string
	SetClientAssertionType(v string) *GenerateTokenByAuthorizationServerRequest
	GetClientAssertionType() *string
	SetClientId(v string) *GenerateTokenByAuthorizationServerRequest
	GetClientId() *string
	SetClientSecret(v string) *GenerateTokenByAuthorizationServerRequest
	GetClientSecret() *string
	SetClientX509(v string) *GenerateTokenByAuthorizationServerRequest
	GetClientX509() *string
	SetClientX509Chain(v string) *GenerateTokenByAuthorizationServerRequest
	GetClientX509Chain() *string
	SetCode(v string) *GenerateTokenByAuthorizationServerRequest
	GetCode() *string
	SetCodeVerifier(v string) *GenerateTokenByAuthorizationServerRequest
	GetCodeVerifier() *string
	SetDeviceCode(v string) *GenerateTokenByAuthorizationServerRequest
	GetDeviceCode() *string
	SetGrantType(v string) *GenerateTokenByAuthorizationServerRequest
	GetGrantType() *string
	SetPassword(v string) *GenerateTokenByAuthorizationServerRequest
	GetPassword() *string
	SetRedirectUri(v string) *GenerateTokenByAuthorizationServerRequest
	GetRedirectUri() *string
	SetRefreshToken(v string) *GenerateTokenByAuthorizationServerRequest
	GetRefreshToken() *string
	SetScope(v string) *GenerateTokenByAuthorizationServerRequest
	GetScope() *string
	SetUsername(v string) *GenerateTokenByAuthorizationServerRequest
	GetUsername() *string
}

type GenerateTokenByAuthorizationServerRequest struct {
	// Federated application credential name.
	//
	// example:
	//
	// testxxxxx
	ApplicationFederatedCredentialName *string `json:"application_federated_credential_name,omitempty" xml:"application_federated_credential_name,omitempty"`
	// Client assertion.
	//
	// example:
	//
	// eyJraWQiOiJLRVlLZ0Iyxxxxx
	ClientAssertion *string `json:"client_assertion,omitempty" xml:"client_assertion,omitempty"`
	// Client assertion type.
	//
	// example:
	//
	// urn:ietf:params:oauth:client-assertion-type:jwt-bearer
	ClientAssertionType *string `json:"client_assertion_type,omitempty" xml:"client_assertion_type,omitempty"`
	// Client ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// Client key.
	//
	// example:
	//
	// CSEHDcHcrUKHw1CuxkJEHPveWRxxxxx
	ClientSecret *string `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
	// Client certificate.
	//
	// example:
	//
	// testxxxxx
	ClientX509 *string `json:"client_x509,omitempty" xml:"client_x509,omitempty"`
	// Intermediate certificate list.
	//
	// example:
	//
	// testxxxxx
	ClientX509Chain *string `json:"client_x509_chain,omitempty" xml:"client_x509_chain,omitempty"`
	// Authorization code. Required when grant_type is authorization_code.
	//
	// example:
	//
	// CO541xY59EsKniV2wvWDXZ4jiKxxxxx
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// You can validate the code.
	//
	// example:
	//
	// xxxxx
	CodeVerifier *string `json:"code_verifier,omitempty" xml:"code_verifier,omitempty"`
	// Device code. Required when grant_type is urn:ietf:params:oauth:grant-type:device_code.
	//
	// example:
	//
	// DCxxxxxx
	DeviceCode *string `json:"device_code,omitempty" xml:"device_code,omitempty"`
	// Grant type.
	//
	// This parameter is required.
	//
	// example:
	//
	// authorization_code
	GrantType *string `json:"grant_type,omitempty" xml:"grant_type,omitempty"`
	// Password.
	//
	// example:
	//
	// testxxxxx
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// Redirection URI. Required when grant_type is authorization_code. Must match the redirect_uri used in the authorization code request.
	//
	// example:
	//
	// https://example.com/xxxxx
	RedirectUri *string `json:"redirect_uri,omitempty" xml:"redirect_uri,omitempty"`
	// Refresh token.
	//
	// example:
	//
	// RTxxxxx
	RefreshToken *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	// Scope.
	//
	// example:
	//
	// openid
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// Username.
	//
	// example:
	//
	// userxxxxx
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GenerateTokenByAuthorizationServerRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateTokenByAuthorizationServerRequest) GoString() string {
	return s.String()
}

func (s *GenerateTokenByAuthorizationServerRequest) GetApplicationFederatedCredentialName() *string {
	return s.ApplicationFederatedCredentialName
}

func (s *GenerateTokenByAuthorizationServerRequest) GetClientAssertion() *string {
	return s.ClientAssertion
}

func (s *GenerateTokenByAuthorizationServerRequest) GetClientAssertionType() *string {
	return s.ClientAssertionType
}

func (s *GenerateTokenByAuthorizationServerRequest) GetClientId() *string {
	return s.ClientId
}

func (s *GenerateTokenByAuthorizationServerRequest) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *GenerateTokenByAuthorizationServerRequest) GetClientX509() *string {
	return s.ClientX509
}

func (s *GenerateTokenByAuthorizationServerRequest) GetClientX509Chain() *string {
	return s.ClientX509Chain
}

func (s *GenerateTokenByAuthorizationServerRequest) GetCode() *string {
	return s.Code
}

func (s *GenerateTokenByAuthorizationServerRequest) GetCodeVerifier() *string {
	return s.CodeVerifier
}

func (s *GenerateTokenByAuthorizationServerRequest) GetDeviceCode() *string {
	return s.DeviceCode
}

func (s *GenerateTokenByAuthorizationServerRequest) GetGrantType() *string {
	return s.GrantType
}

func (s *GenerateTokenByAuthorizationServerRequest) GetPassword() *string {
	return s.Password
}

func (s *GenerateTokenByAuthorizationServerRequest) GetRedirectUri() *string {
	return s.RedirectUri
}

func (s *GenerateTokenByAuthorizationServerRequest) GetRefreshToken() *string {
	return s.RefreshToken
}

func (s *GenerateTokenByAuthorizationServerRequest) GetScope() *string {
	return s.Scope
}

func (s *GenerateTokenByAuthorizationServerRequest) GetUsername() *string {
	return s.Username
}

func (s *GenerateTokenByAuthorizationServerRequest) SetApplicationFederatedCredentialName(v string) *GenerateTokenByAuthorizationServerRequest {
	s.ApplicationFederatedCredentialName = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) SetClientAssertion(v string) *GenerateTokenByAuthorizationServerRequest {
	s.ClientAssertion = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) SetClientAssertionType(v string) *GenerateTokenByAuthorizationServerRequest {
	s.ClientAssertionType = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) SetClientId(v string) *GenerateTokenByAuthorizationServerRequest {
	s.ClientId = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) SetClientSecret(v string) *GenerateTokenByAuthorizationServerRequest {
	s.ClientSecret = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) SetClientX509(v string) *GenerateTokenByAuthorizationServerRequest {
	s.ClientX509 = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) SetClientX509Chain(v string) *GenerateTokenByAuthorizationServerRequest {
	s.ClientX509Chain = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) SetCode(v string) *GenerateTokenByAuthorizationServerRequest {
	s.Code = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) SetCodeVerifier(v string) *GenerateTokenByAuthorizationServerRequest {
	s.CodeVerifier = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) SetDeviceCode(v string) *GenerateTokenByAuthorizationServerRequest {
	s.DeviceCode = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) SetGrantType(v string) *GenerateTokenByAuthorizationServerRequest {
	s.GrantType = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) SetPassword(v string) *GenerateTokenByAuthorizationServerRequest {
	s.Password = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) SetRedirectUri(v string) *GenerateTokenByAuthorizationServerRequest {
	s.RedirectUri = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) SetRefreshToken(v string) *GenerateTokenByAuthorizationServerRequest {
	s.RefreshToken = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) SetScope(v string) *GenerateTokenByAuthorizationServerRequest {
	s.Scope = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) SetUsername(v string) *GenerateTokenByAuthorizationServerRequest {
	s.Username = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerRequest) Validate() error {
	return dara.Validate(s)
}
