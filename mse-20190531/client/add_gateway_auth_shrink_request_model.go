// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayAuthShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddGatewayAuthShrinkRequest
	GetAcceptLanguage() *string
	SetAuthResourceConfig(v string) *AddGatewayAuthShrinkRequest
	GetAuthResourceConfig() *string
	SetAuthResourceListShrink(v string) *AddGatewayAuthShrinkRequest
	GetAuthResourceListShrink() *string
	SetAuthResourceMode(v int32) *AddGatewayAuthShrinkRequest
	GetAuthResourceMode() *int32
	SetClientId(v string) *AddGatewayAuthShrinkRequest
	GetClientId() *string
	SetClientSecret(v string) *AddGatewayAuthShrinkRequest
	GetClientSecret() *string
	SetCookieDomain(v string) *AddGatewayAuthShrinkRequest
	GetCookieDomain() *string
	SetExternalAuthZJSONShrink(v string) *AddGatewayAuthShrinkRequest
	GetExternalAuthZJSONShrink() *string
	SetGatewayUniqueId(v string) *AddGatewayAuthShrinkRequest
	GetGatewayUniqueId() *string
	SetIsWhite(v bool) *AddGatewayAuthShrinkRequest
	GetIsWhite() *bool
	SetIssuer(v string) *AddGatewayAuthShrinkRequest
	GetIssuer() *string
	SetJwks(v string) *AddGatewayAuthShrinkRequest
	GetJwks() *string
	SetLoginUrl(v string) *AddGatewayAuthShrinkRequest
	GetLoginUrl() *string
	SetName(v string) *AddGatewayAuthShrinkRequest
	GetName() *string
	SetRedirectUrl(v string) *AddGatewayAuthShrinkRequest
	GetRedirectUrl() *string
	SetScopesListShrink(v string) *AddGatewayAuthShrinkRequest
	GetScopesListShrink() *string
	SetStatus(v bool) *AddGatewayAuthShrinkRequest
	GetStatus() *bool
	SetSub(v string) *AddGatewayAuthShrinkRequest
	GetSub() *string
	SetTokenName(v string) *AddGatewayAuthShrinkRequest
	GetTokenName() *string
	SetTokenNamePrefix(v string) *AddGatewayAuthShrinkRequest
	GetTokenNamePrefix() *string
	SetTokenPass(v bool) *AddGatewayAuthShrinkRequest
	GetTokenPass() *bool
	SetTokenPosition(v string) *AddGatewayAuthShrinkRequest
	GetTokenPosition() *string
	SetType(v string) *AddGatewayAuthShrinkRequest
	GetType() *string
}

type AddGatewayAuthShrinkRequest struct {
	// example:
	//
	// zh
	AcceptLanguage     *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AuthResourceConfig *string `json:"AuthResourceConfig,omitempty" xml:"AuthResourceConfig,omitempty"`
	// The information about the resource to be authorized.
	AuthResourceListShrink *string `json:"AuthResourceList,omitempty" xml:"AuthResourceList,omitempty"`
	// example:
	//
	// 1
	AuthResourceMode *int32 `json:"AuthResourceMode,omitempty" xml:"AuthResourceMode,omitempty"`
	// The application ID registered with the OIDC authentication service.
	//
	// example:
	//
	// 23460e2fdd9bf9ad106****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The application secret registered with the OIDC authentication service.
	//
	// example:
	//
	// 123****
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The domain name of the cookie. After the authentication is passed, the cookie is sent to the specified domain name to maintain the logon status. For example, if you set `Cookie-domain` to a.example.com, the cookie is sent to the domain name `a.example.com`. If you set `Cookie-domain` to .example.com, the cookie is sent to all subdomains of `example.com`.
	//
	// example:
	//
	// test.com
	CookieDomain *string `json:"CookieDomain,omitempty" xml:"CookieDomain,omitempty"`
	// The information about the custom authentication service.
	ExternalAuthZJSONShrink *string `json:"ExternalAuthZJSON,omitempty" xml:"ExternalAuthZJSON,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-492af9b04bb4474cae9d645be850e3d7
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// Specifies whether to enable the whitelist feature.
	//
	// example:
	//
	// true
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// The iss value of JWT claims, which indicates the issuer. You must make sure that the value of this parameter is the same as the iss value in the payload of JWT claims.
	//
	// example:
	//
	// testing@secure.istio.io
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The JWT public key. The JSON format is supported.
	//
	// example:
	//
	// {"keys":[{"e":"AQAB","kid":"DHFbpoIUqrY8t2zpA2qXfCmr5VO5ZEr4RzHU_-envvQ","kty":"RSA","n":"xAE7eB6qugXyCAG3yhh7pkDkT65pHymX-P7KfIupjf59vsdo91bSP9C8H07pSAGQO1MV_xFj9VswgsCg4R6otmg5PV2He95lZdHtOcU5DXIg_pbhLdKXbi66GlVeK6ABZOUW3WYtnNHD-91gVuoeJT_DwtGGcp4ignkgXfkiEm4sw-4sfb4qdt5oLbyVpmW6x9cfa7vs2WTfURiCrBoUqgBo_-4WTiULmmHSGZHOjzwa8WtrtOQGsAFjIbno85jp6MnGGGZPYZbDAa_b3y5u-YpW7ypZrvD8BgtKVjgtQgZhLAGezMt0ua3DRrWnKqTZ0BJ_EyxOGuHJrLsn00fnMQ"}]}
	Jwks *string `json:"Jwks,omitempty" xml:"Jwks,omitempty"`
	// The URL that is used to log on to the IDaaS instance.
	//
	// example:
	//
	// ***
	LoginUrl *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	// The name.
	//
	// example:
	//
	// jwt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The redirect URL.
	//
	// example:
	//
	// https://test-.com/oauth2/callback
	RedirectUrl *string `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	// The OIDC scope.
	ScopesListShrink *string `json:"ScopesList,omitempty" xml:"ScopesList,omitempty"`
	// The status.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sub value of JWT claims, which indicates the subject. You must make sure that the value of this parameter is the same as the sub value in the payload of JWT claims. If you do not set this parameter or leave it empty, the default value, which is the value of the Issuer parameter, is used.
	//
	// example:
	//
	// testing@secure.istio.io
	Sub *string `json:"Sub,omitempty" xml:"Sub,omitempty"`
	// The name of the parameter that is required to verify a token. By default, a token is prefixed with Bearer and stored in the authorization header. Example: `Authorization: Bearer token`.
	//
	// example:
	//
	// Authorization
	TokenName *string `json:"TokenName,omitempty" xml:"TokenName,omitempty"`
	// The name prefix of the parameter that is required to verify a token. By default, a token is prefixed with Bearer and stored in the authorization header. Example: `Authorization: Bearer token`
	//
	// example:
	//
	// Bearer
	TokenNamePrefix *string `json:"TokenNamePrefix,omitempty" xml:"TokenNamePrefix,omitempty"`
	// Specifies whether to enable pass-through.
	//
	// example:
	//
	// true
	TokenPass *bool `json:"TokenPass,omitempty" xml:"TokenPass,omitempty"`
	// The position of the parameter that is required to verify a token. By default, a token is prefixed with Bearer and stored in the authorization header. Example: `Authorization: Bearer token`.
	//
	// example:
	//
	// HEADER
	TokenPosition *string `json:"TokenPosition,omitempty" xml:"TokenPosition,omitempty"`
	// The authentication type. JSON Web Token (JWT) authentication, OpenID Connect (OIDC) authentication, Identity as a Service (IDaaS) authentication, or custom authentication are supported.
	//
	// example:
	//
	// JWT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddGatewayAuthShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayAuthShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddGatewayAuthShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddGatewayAuthShrinkRequest) GetAuthResourceConfig() *string {
	return s.AuthResourceConfig
}

func (s *AddGatewayAuthShrinkRequest) GetAuthResourceListShrink() *string {
	return s.AuthResourceListShrink
}

func (s *AddGatewayAuthShrinkRequest) GetAuthResourceMode() *int32 {
	return s.AuthResourceMode
}

func (s *AddGatewayAuthShrinkRequest) GetClientId() *string {
	return s.ClientId
}

func (s *AddGatewayAuthShrinkRequest) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *AddGatewayAuthShrinkRequest) GetCookieDomain() *string {
	return s.CookieDomain
}

func (s *AddGatewayAuthShrinkRequest) GetExternalAuthZJSONShrink() *string {
	return s.ExternalAuthZJSONShrink
}

func (s *AddGatewayAuthShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddGatewayAuthShrinkRequest) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *AddGatewayAuthShrinkRequest) GetIssuer() *string {
	return s.Issuer
}

func (s *AddGatewayAuthShrinkRequest) GetJwks() *string {
	return s.Jwks
}

func (s *AddGatewayAuthShrinkRequest) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *AddGatewayAuthShrinkRequest) GetName() *string {
	return s.Name
}

func (s *AddGatewayAuthShrinkRequest) GetRedirectUrl() *string {
	return s.RedirectUrl
}

func (s *AddGatewayAuthShrinkRequest) GetScopesListShrink() *string {
	return s.ScopesListShrink
}

func (s *AddGatewayAuthShrinkRequest) GetStatus() *bool {
	return s.Status
}

func (s *AddGatewayAuthShrinkRequest) GetSub() *string {
	return s.Sub
}

func (s *AddGatewayAuthShrinkRequest) GetTokenName() *string {
	return s.TokenName
}

func (s *AddGatewayAuthShrinkRequest) GetTokenNamePrefix() *string {
	return s.TokenNamePrefix
}

func (s *AddGatewayAuthShrinkRequest) GetTokenPass() *bool {
	return s.TokenPass
}

func (s *AddGatewayAuthShrinkRequest) GetTokenPosition() *string {
	return s.TokenPosition
}

func (s *AddGatewayAuthShrinkRequest) GetType() *string {
	return s.Type
}

func (s *AddGatewayAuthShrinkRequest) SetAcceptLanguage(v string) *AddGatewayAuthShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetAuthResourceConfig(v string) *AddGatewayAuthShrinkRequest {
	s.AuthResourceConfig = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetAuthResourceListShrink(v string) *AddGatewayAuthShrinkRequest {
	s.AuthResourceListShrink = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetAuthResourceMode(v int32) *AddGatewayAuthShrinkRequest {
	s.AuthResourceMode = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetClientId(v string) *AddGatewayAuthShrinkRequest {
	s.ClientId = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetClientSecret(v string) *AddGatewayAuthShrinkRequest {
	s.ClientSecret = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetCookieDomain(v string) *AddGatewayAuthShrinkRequest {
	s.CookieDomain = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetExternalAuthZJSONShrink(v string) *AddGatewayAuthShrinkRequest {
	s.ExternalAuthZJSONShrink = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetGatewayUniqueId(v string) *AddGatewayAuthShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetIsWhite(v bool) *AddGatewayAuthShrinkRequest {
	s.IsWhite = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetIssuer(v string) *AddGatewayAuthShrinkRequest {
	s.Issuer = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetJwks(v string) *AddGatewayAuthShrinkRequest {
	s.Jwks = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetLoginUrl(v string) *AddGatewayAuthShrinkRequest {
	s.LoginUrl = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetName(v string) *AddGatewayAuthShrinkRequest {
	s.Name = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetRedirectUrl(v string) *AddGatewayAuthShrinkRequest {
	s.RedirectUrl = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetScopesListShrink(v string) *AddGatewayAuthShrinkRequest {
	s.ScopesListShrink = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetStatus(v bool) *AddGatewayAuthShrinkRequest {
	s.Status = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetSub(v string) *AddGatewayAuthShrinkRequest {
	s.Sub = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetTokenName(v string) *AddGatewayAuthShrinkRequest {
	s.TokenName = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetTokenNamePrefix(v string) *AddGatewayAuthShrinkRequest {
	s.TokenNamePrefix = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetTokenPass(v bool) *AddGatewayAuthShrinkRequest {
	s.TokenPass = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetTokenPosition(v string) *AddGatewayAuthShrinkRequest {
	s.TokenPosition = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) SetType(v string) *AddGatewayAuthShrinkRequest {
	s.Type = &v
	return s
}

func (s *AddGatewayAuthShrinkRequest) Validate() error {
	return dara.Validate(s)
}
