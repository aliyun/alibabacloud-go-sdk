// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddGatewayAuthRequest
	GetAcceptLanguage() *string
	SetAuthResourceConfig(v string) *AddGatewayAuthRequest
	GetAuthResourceConfig() *string
	SetAuthResourceList(v []*AddGatewayAuthRequestAuthResourceList) *AddGatewayAuthRequest
	GetAuthResourceList() []*AddGatewayAuthRequestAuthResourceList
	SetAuthResourceMode(v int32) *AddGatewayAuthRequest
	GetAuthResourceMode() *int32
	SetClientId(v string) *AddGatewayAuthRequest
	GetClientId() *string
	SetClientSecret(v string) *AddGatewayAuthRequest
	GetClientSecret() *string
	SetCookieDomain(v string) *AddGatewayAuthRequest
	GetCookieDomain() *string
	SetExternalAuthZJSON(v *AddGatewayAuthRequestExternalAuthZJSON) *AddGatewayAuthRequest
	GetExternalAuthZJSON() *AddGatewayAuthRequestExternalAuthZJSON
	SetGatewayUniqueId(v string) *AddGatewayAuthRequest
	GetGatewayUniqueId() *string
	SetIsWhite(v bool) *AddGatewayAuthRequest
	GetIsWhite() *bool
	SetIssuer(v string) *AddGatewayAuthRequest
	GetIssuer() *string
	SetJwks(v string) *AddGatewayAuthRequest
	GetJwks() *string
	SetLoginUrl(v string) *AddGatewayAuthRequest
	GetLoginUrl() *string
	SetName(v string) *AddGatewayAuthRequest
	GetName() *string
	SetRedirectUrl(v string) *AddGatewayAuthRequest
	GetRedirectUrl() *string
	SetScopesList(v []*string) *AddGatewayAuthRequest
	GetScopesList() []*string
	SetStatus(v bool) *AddGatewayAuthRequest
	GetStatus() *bool
	SetSub(v string) *AddGatewayAuthRequest
	GetSub() *string
	SetTokenName(v string) *AddGatewayAuthRequest
	GetTokenName() *string
	SetTokenNamePrefix(v string) *AddGatewayAuthRequest
	GetTokenNamePrefix() *string
	SetTokenPass(v bool) *AddGatewayAuthRequest
	GetTokenPass() *bool
	SetTokenPosition(v string) *AddGatewayAuthRequest
	GetTokenPosition() *string
	SetType(v string) *AddGatewayAuthRequest
	GetType() *string
}

type AddGatewayAuthRequest struct {
	// example:
	//
	// zh
	AcceptLanguage     *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AuthResourceConfig *string `json:"AuthResourceConfig,omitempty" xml:"AuthResourceConfig,omitempty"`
	// The information about the resource to be authorized.
	AuthResourceList []*AddGatewayAuthRequestAuthResourceList `json:"AuthResourceList,omitempty" xml:"AuthResourceList,omitempty" type:"Repeated"`
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
	ExternalAuthZJSON *AddGatewayAuthRequestExternalAuthZJSON `json:"ExternalAuthZJSON,omitempty" xml:"ExternalAuthZJSON,omitempty" type:"Struct"`
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
	ScopesList []*string `json:"ScopesList,omitempty" xml:"ScopesList,omitempty" type:"Repeated"`
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

func (s AddGatewayAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayAuthRequest) GoString() string {
	return s.String()
}

func (s *AddGatewayAuthRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddGatewayAuthRequest) GetAuthResourceConfig() *string {
	return s.AuthResourceConfig
}

func (s *AddGatewayAuthRequest) GetAuthResourceList() []*AddGatewayAuthRequestAuthResourceList {
	return s.AuthResourceList
}

func (s *AddGatewayAuthRequest) GetAuthResourceMode() *int32 {
	return s.AuthResourceMode
}

func (s *AddGatewayAuthRequest) GetClientId() *string {
	return s.ClientId
}

func (s *AddGatewayAuthRequest) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *AddGatewayAuthRequest) GetCookieDomain() *string {
	return s.CookieDomain
}

func (s *AddGatewayAuthRequest) GetExternalAuthZJSON() *AddGatewayAuthRequestExternalAuthZJSON {
	return s.ExternalAuthZJSON
}

func (s *AddGatewayAuthRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddGatewayAuthRequest) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *AddGatewayAuthRequest) GetIssuer() *string {
	return s.Issuer
}

func (s *AddGatewayAuthRequest) GetJwks() *string {
	return s.Jwks
}

func (s *AddGatewayAuthRequest) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *AddGatewayAuthRequest) GetName() *string {
	return s.Name
}

func (s *AddGatewayAuthRequest) GetRedirectUrl() *string {
	return s.RedirectUrl
}

func (s *AddGatewayAuthRequest) GetScopesList() []*string {
	return s.ScopesList
}

func (s *AddGatewayAuthRequest) GetStatus() *bool {
	return s.Status
}

func (s *AddGatewayAuthRequest) GetSub() *string {
	return s.Sub
}

func (s *AddGatewayAuthRequest) GetTokenName() *string {
	return s.TokenName
}

func (s *AddGatewayAuthRequest) GetTokenNamePrefix() *string {
	return s.TokenNamePrefix
}

func (s *AddGatewayAuthRequest) GetTokenPass() *bool {
	return s.TokenPass
}

func (s *AddGatewayAuthRequest) GetTokenPosition() *string {
	return s.TokenPosition
}

func (s *AddGatewayAuthRequest) GetType() *string {
	return s.Type
}

func (s *AddGatewayAuthRequest) SetAcceptLanguage(v string) *AddGatewayAuthRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddGatewayAuthRequest) SetAuthResourceConfig(v string) *AddGatewayAuthRequest {
	s.AuthResourceConfig = &v
	return s
}

func (s *AddGatewayAuthRequest) SetAuthResourceList(v []*AddGatewayAuthRequestAuthResourceList) *AddGatewayAuthRequest {
	s.AuthResourceList = v
	return s
}

func (s *AddGatewayAuthRequest) SetAuthResourceMode(v int32) *AddGatewayAuthRequest {
	s.AuthResourceMode = &v
	return s
}

func (s *AddGatewayAuthRequest) SetClientId(v string) *AddGatewayAuthRequest {
	s.ClientId = &v
	return s
}

func (s *AddGatewayAuthRequest) SetClientSecret(v string) *AddGatewayAuthRequest {
	s.ClientSecret = &v
	return s
}

func (s *AddGatewayAuthRequest) SetCookieDomain(v string) *AddGatewayAuthRequest {
	s.CookieDomain = &v
	return s
}

func (s *AddGatewayAuthRequest) SetExternalAuthZJSON(v *AddGatewayAuthRequestExternalAuthZJSON) *AddGatewayAuthRequest {
	s.ExternalAuthZJSON = v
	return s
}

func (s *AddGatewayAuthRequest) SetGatewayUniqueId(v string) *AddGatewayAuthRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddGatewayAuthRequest) SetIsWhite(v bool) *AddGatewayAuthRequest {
	s.IsWhite = &v
	return s
}

func (s *AddGatewayAuthRequest) SetIssuer(v string) *AddGatewayAuthRequest {
	s.Issuer = &v
	return s
}

func (s *AddGatewayAuthRequest) SetJwks(v string) *AddGatewayAuthRequest {
	s.Jwks = &v
	return s
}

func (s *AddGatewayAuthRequest) SetLoginUrl(v string) *AddGatewayAuthRequest {
	s.LoginUrl = &v
	return s
}

func (s *AddGatewayAuthRequest) SetName(v string) *AddGatewayAuthRequest {
	s.Name = &v
	return s
}

func (s *AddGatewayAuthRequest) SetRedirectUrl(v string) *AddGatewayAuthRequest {
	s.RedirectUrl = &v
	return s
}

func (s *AddGatewayAuthRequest) SetScopesList(v []*string) *AddGatewayAuthRequest {
	s.ScopesList = v
	return s
}

func (s *AddGatewayAuthRequest) SetStatus(v bool) *AddGatewayAuthRequest {
	s.Status = &v
	return s
}

func (s *AddGatewayAuthRequest) SetSub(v string) *AddGatewayAuthRequest {
	s.Sub = &v
	return s
}

func (s *AddGatewayAuthRequest) SetTokenName(v string) *AddGatewayAuthRequest {
	s.TokenName = &v
	return s
}

func (s *AddGatewayAuthRequest) SetTokenNamePrefix(v string) *AddGatewayAuthRequest {
	s.TokenNamePrefix = &v
	return s
}

func (s *AddGatewayAuthRequest) SetTokenPass(v bool) *AddGatewayAuthRequest {
	s.TokenPass = &v
	return s
}

func (s *AddGatewayAuthRequest) SetTokenPosition(v string) *AddGatewayAuthRequest {
	s.TokenPosition = &v
	return s
}

func (s *AddGatewayAuthRequest) SetType(v string) *AddGatewayAuthRequest {
	s.Type = &v
	return s
}

func (s *AddGatewayAuthRequest) Validate() error {
	if s.AuthResourceList != nil {
		for _, item := range s.AuthResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ExternalAuthZJSON != nil {
		if err := s.ExternalAuthZJSON.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddGatewayAuthRequestAuthResourceList struct {
	AuthResourceHeaderList []*AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList `json:"AuthResourceHeaderList,omitempty" xml:"AuthResourceHeaderList,omitempty" type:"Repeated"`
	// The domain ID.
	//
	// example:
	//
	// 1
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// example:
	//
	// true
	IgnoreCase *bool `json:"IgnoreCase,omitempty" xml:"IgnoreCase,omitempty"`
	// example:
	//
	// EQUAL
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The request path.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s AddGatewayAuthRequestAuthResourceList) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayAuthRequestAuthResourceList) GoString() string {
	return s.String()
}

func (s *AddGatewayAuthRequestAuthResourceList) GetAuthResourceHeaderList() []*AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList {
	return s.AuthResourceHeaderList
}

func (s *AddGatewayAuthRequestAuthResourceList) GetDomainId() *int64 {
	return s.DomainId
}

func (s *AddGatewayAuthRequestAuthResourceList) GetIgnoreCase() *bool {
	return s.IgnoreCase
}

func (s *AddGatewayAuthRequestAuthResourceList) GetMatchType() *string {
	return s.MatchType
}

func (s *AddGatewayAuthRequestAuthResourceList) GetPath() *string {
	return s.Path
}

func (s *AddGatewayAuthRequestAuthResourceList) SetAuthResourceHeaderList(v []*AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList) *AddGatewayAuthRequestAuthResourceList {
	s.AuthResourceHeaderList = v
	return s
}

func (s *AddGatewayAuthRequestAuthResourceList) SetDomainId(v int64) *AddGatewayAuthRequestAuthResourceList {
	s.DomainId = &v
	return s
}

func (s *AddGatewayAuthRequestAuthResourceList) SetIgnoreCase(v bool) *AddGatewayAuthRequestAuthResourceList {
	s.IgnoreCase = &v
	return s
}

func (s *AddGatewayAuthRequestAuthResourceList) SetMatchType(v string) *AddGatewayAuthRequestAuthResourceList {
	s.MatchType = &v
	return s
}

func (s *AddGatewayAuthRequestAuthResourceList) SetPath(v string) *AddGatewayAuthRequestAuthResourceList {
	s.Path = &v
	return s
}

func (s *AddGatewayAuthRequestAuthResourceList) Validate() error {
	if s.AuthResourceHeaderList != nil {
		for _, item := range s.AuthResourceHeaderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList struct {
	// example:
	//
	// x-req
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	// example:
	//
	// EQUAL
	HeaderMethod *string `json:"HeaderMethod,omitempty" xml:"HeaderMethod,omitempty"`
	// example:
	//
	// 123
	HeaderValue *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
}

func (s AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList) GoString() string {
	return s.String()
}

func (s *AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList) GetHeaderMethod() *string {
	return s.HeaderMethod
}

func (s *AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList) GetHeaderValue() *string {
	return s.HeaderValue
}

func (s *AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList) SetHeaderKey(v string) *AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList {
	s.HeaderKey = &v
	return s
}

func (s *AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList) SetHeaderMethod(v string) *AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList {
	s.HeaderMethod = &v
	return s
}

func (s *AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList) SetHeaderValue(v string) *AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList {
	s.HeaderValue = &v
	return s
}

func (s *AddGatewayAuthRequestAuthResourceListAuthResourceHeaderList) Validate() error {
	return dara.Validate(s)
}

type AddGatewayAuthRequestExternalAuthZJSON struct {
	// The header that can be carried in an authentication request.
	AllowRequestHeaders []*string `json:"AllowRequestHeaders,omitempty" xml:"AllowRequestHeaders,omitempty" type:"Repeated"`
	// The header that can be retained in an authentication response.
	AllowUpstreamHeaders []*string `json:"AllowUpstreamHeaders,omitempty" xml:"AllowUpstreamHeaders,omitempty" type:"Repeated"`
	// example:
	//
	// 4000000
	BodyMaxBytes *int32 `json:"BodyMaxBytes,omitempty" xml:"BodyMaxBytes,omitempty"`
	// Specifies whether the gateway allows a client request when the authentication server is unavailable. If a connection to the authentication server fails to be established or a 5xx error code is returned, the authentication server is unavailable.
	//
	// example:
	//
	// true
	IsRestrict *bool `json:"IsRestrict,omitempty" xml:"IsRestrict,omitempty"`
	// The path of the authentication API provided by the authentication service. The path supports the prefix match method.
	//
	// example:
	//
	// /auth
	PrefixPath *string `json:"PrefixPath,omitempty" xml:"PrefixPath,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 1
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The timeout period. Unit: seconds.
	//
	// example:
	//
	// 100
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The header that stores a token in an authentication request. In most cases, a token is stored in the Authorization or Cookie header.
	//
	// example:
	//
	// Authorization
	TokenKey         *string `json:"TokenKey,omitempty" xml:"TokenKey,omitempty"`
	WithRematchRoute *bool   `json:"WithRematchRoute,omitempty" xml:"WithRematchRoute,omitempty"`
	// example:
	//
	// true
	WithRequestBody *bool `json:"WithRequestBody,omitempty" xml:"WithRequestBody,omitempty"`
}

func (s AddGatewayAuthRequestExternalAuthZJSON) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayAuthRequestExternalAuthZJSON) GoString() string {
	return s.String()
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) GetAllowRequestHeaders() []*string {
	return s.AllowRequestHeaders
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) GetAllowUpstreamHeaders() []*string {
	return s.AllowUpstreamHeaders
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) GetBodyMaxBytes() *int32 {
	return s.BodyMaxBytes
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) GetIsRestrict() *bool {
	return s.IsRestrict
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) GetPrefixPath() *string {
	return s.PrefixPath
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) GetTimeout() *int32 {
	return s.Timeout
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) GetTokenKey() *string {
	return s.TokenKey
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) GetWithRematchRoute() *bool {
	return s.WithRematchRoute
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) GetWithRequestBody() *bool {
	return s.WithRequestBody
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) SetAllowRequestHeaders(v []*string) *AddGatewayAuthRequestExternalAuthZJSON {
	s.AllowRequestHeaders = v
	return s
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) SetAllowUpstreamHeaders(v []*string) *AddGatewayAuthRequestExternalAuthZJSON {
	s.AllowUpstreamHeaders = v
	return s
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) SetBodyMaxBytes(v int32) *AddGatewayAuthRequestExternalAuthZJSON {
	s.BodyMaxBytes = &v
	return s
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) SetIsRestrict(v bool) *AddGatewayAuthRequestExternalAuthZJSON {
	s.IsRestrict = &v
	return s
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) SetPrefixPath(v string) *AddGatewayAuthRequestExternalAuthZJSON {
	s.PrefixPath = &v
	return s
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) SetServiceId(v int64) *AddGatewayAuthRequestExternalAuthZJSON {
	s.ServiceId = &v
	return s
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) SetTimeout(v int32) *AddGatewayAuthRequestExternalAuthZJSON {
	s.Timeout = &v
	return s
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) SetTokenKey(v string) *AddGatewayAuthRequestExternalAuthZJSON {
	s.TokenKey = &v
	return s
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) SetWithRematchRoute(v bool) *AddGatewayAuthRequestExternalAuthZJSON {
	s.WithRematchRoute = &v
	return s
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) SetWithRequestBody(v bool) *AddGatewayAuthRequestExternalAuthZJSON {
	s.WithRequestBody = &v
	return s
}

func (s *AddGatewayAuthRequestExternalAuthZJSON) Validate() error {
	return dara.Validate(s)
}
