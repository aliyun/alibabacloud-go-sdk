// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayAuthShrinkRequest
	GetAcceptLanguage() *string
	SetAuthResourceConfig(v string) *UpdateGatewayAuthShrinkRequest
	GetAuthResourceConfig() *string
	SetAuthResourceListShrink(v string) *UpdateGatewayAuthShrinkRequest
	GetAuthResourceListShrink() *string
	SetAuthResourceMode(v int32) *UpdateGatewayAuthShrinkRequest
	GetAuthResourceMode() *int32
	SetClientId(v string) *UpdateGatewayAuthShrinkRequest
	GetClientId() *string
	SetClientSecret(v string) *UpdateGatewayAuthShrinkRequest
	GetClientSecret() *string
	SetCookieDomain(v string) *UpdateGatewayAuthShrinkRequest
	GetCookieDomain() *string
	SetDeleteResourceIdListShrink(v string) *UpdateGatewayAuthShrinkRequest
	GetDeleteResourceIdListShrink() *string
	SetExternalAuthZJSONShrink(v string) *UpdateGatewayAuthShrinkRequest
	GetExternalAuthZJSONShrink() *string
	SetGatewayUniqueId(v string) *UpdateGatewayAuthShrinkRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayAuthShrinkRequest
	GetId() *int64
	SetIsWhite(v bool) *UpdateGatewayAuthShrinkRequest
	GetIsWhite() *bool
	SetIssuer(v string) *UpdateGatewayAuthShrinkRequest
	GetIssuer() *string
	SetJwks(v string) *UpdateGatewayAuthShrinkRequest
	GetJwks() *string
	SetLoginUrl(v string) *UpdateGatewayAuthShrinkRequest
	GetLoginUrl() *string
	SetName(v string) *UpdateGatewayAuthShrinkRequest
	GetName() *string
	SetRedirectUrl(v string) *UpdateGatewayAuthShrinkRequest
	GetRedirectUrl() *string
	SetScopesListShrink(v string) *UpdateGatewayAuthShrinkRequest
	GetScopesListShrink() *string
	SetStatus(v bool) *UpdateGatewayAuthShrinkRequest
	GetStatus() *bool
	SetSub(v string) *UpdateGatewayAuthShrinkRequest
	GetSub() *string
	SetTokenName(v string) *UpdateGatewayAuthShrinkRequest
	GetTokenName() *string
	SetTokenNamePrefix(v string) *UpdateGatewayAuthShrinkRequest
	GetTokenNamePrefix() *string
	SetTokenPass(v bool) *UpdateGatewayAuthShrinkRequest
	GetTokenPass() *bool
	SetTokenPosition(v string) *UpdateGatewayAuthShrinkRequest
	GetTokenPosition() *string
	SetType(v string) *UpdateGatewayAuthShrinkRequest
	GetType() *string
}

type UpdateGatewayAuthShrinkRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// \\"\\"
	AuthResourceConfig     *string `json:"AuthResourceConfig,omitempty" xml:"AuthResourceConfig,omitempty"`
	AuthResourceListShrink *string `json:"AuthResourceList,omitempty" xml:"AuthResourceList,omitempty"`
	// example:
	//
	// 1
	AuthResourceMode *int32 `json:"AuthResourceMode,omitempty" xml:"AuthResourceMode,omitempty"`
	// example:
	//
	// app_mnvxaavggw7hcdcnr6usi6***
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// CS6EYfx3k9yTRR9EtQ2MXWP97P6UAUwFg4teoWJ19Z****
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// example:
	//
	// test.com
	CookieDomain               *string `json:"CookieDomain,omitempty" xml:"CookieDomain,omitempty"`
	DeleteResourceIdListShrink *string `json:"DeleteResourceIdList,omitempty" xml:"DeleteResourceIdList,omitempty"`
	ExternalAuthZJSONShrink    *string `json:"ExternalAuthZJSON,omitempty" xml:"ExternalAuthZJSON,omitempty"`
	// example:
	//
	// gw-9cdcf8e4f58144059e73ff4c5ef9****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 719
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// example:
	//
	// test
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// example:
	//
	// {"keys":[{"e":"AQAB","kid":"DHFbpoIUqrY8t2zpA2qXfCmr5VO5ZEr4RzHU_-envvQ","kty":"RSA","n":"xAE7eB6qugXyCAG3yhh7pkDkT65pHymX-P7KfIupjf59vsdo91bSP9C8H07pSAGQO1MV_xFj9VswgsCg4R6otmg5PV2He95lZdHtOcU5DXIg_pbhLdKXbi66GlVeK6ABZOUW3WYtnNHD-91gVuoeJT_DwtGGcp4ignkgXfkiEm4sw-4sfb4qdt5oLbyVpmW6x9cfa7vs2WTfURiCrBoUqgBo_-4WTiULmmHSGZHOjzwa8WtrtOQGsAFjIbno85jp6MnGGGZPYZbDAa_b3y5u-YpW7ypZrvD8BgtKVjgtQgZhLAGezMt0ua3DRrWnKqTZ0BJ_EyxOGuHJrLsn00fnMQ"}]}
	Jwks *string `json:"Jwks,omitempty" xml:"Jwks,omitempty"`
	// example:
	//
	// https://daxxxxcn.aliyunidaas.com/
	LoginUrl *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// https://test-.com/oauth2/callback
	RedirectUrl      *string `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	ScopesListShrink *string `json:"ScopesList,omitempty" xml:"ScopesList,omitempty"`
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// example-app
	Sub *string `json:"Sub,omitempty" xml:"Sub,omitempty"`
	// example:
	//
	// Authorization
	TokenName *string `json:"TokenName,omitempty" xml:"TokenName,omitempty"`
	// example:
	//
	// Bearer
	TokenNamePrefix *string `json:"TokenNamePrefix,omitempty" xml:"TokenNamePrefix,omitempty"`
	// example:
	//
	// true
	TokenPass *bool `json:"TokenPass,omitempty" xml:"TokenPass,omitempty"`
	// example:
	//
	// HEADER
	TokenPosition *string `json:"TokenPosition,omitempty" xml:"TokenPosition,omitempty"`
	// example:
	//
	// ExternalAuthZ
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateGatewayAuthShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayAuthShrinkRequest) GetAuthResourceConfig() *string {
	return s.AuthResourceConfig
}

func (s *UpdateGatewayAuthShrinkRequest) GetAuthResourceListShrink() *string {
	return s.AuthResourceListShrink
}

func (s *UpdateGatewayAuthShrinkRequest) GetAuthResourceMode() *int32 {
	return s.AuthResourceMode
}

func (s *UpdateGatewayAuthShrinkRequest) GetClientId() *string {
	return s.ClientId
}

func (s *UpdateGatewayAuthShrinkRequest) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *UpdateGatewayAuthShrinkRequest) GetCookieDomain() *string {
	return s.CookieDomain
}

func (s *UpdateGatewayAuthShrinkRequest) GetDeleteResourceIdListShrink() *string {
	return s.DeleteResourceIdListShrink
}

func (s *UpdateGatewayAuthShrinkRequest) GetExternalAuthZJSONShrink() *string {
	return s.ExternalAuthZJSONShrink
}

func (s *UpdateGatewayAuthShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayAuthShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayAuthShrinkRequest) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *UpdateGatewayAuthShrinkRequest) GetIssuer() *string {
	return s.Issuer
}

func (s *UpdateGatewayAuthShrinkRequest) GetJwks() *string {
	return s.Jwks
}

func (s *UpdateGatewayAuthShrinkRequest) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *UpdateGatewayAuthShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateGatewayAuthShrinkRequest) GetRedirectUrl() *string {
	return s.RedirectUrl
}

func (s *UpdateGatewayAuthShrinkRequest) GetScopesListShrink() *string {
	return s.ScopesListShrink
}

func (s *UpdateGatewayAuthShrinkRequest) GetStatus() *bool {
	return s.Status
}

func (s *UpdateGatewayAuthShrinkRequest) GetSub() *string {
	return s.Sub
}

func (s *UpdateGatewayAuthShrinkRequest) GetTokenName() *string {
	return s.TokenName
}

func (s *UpdateGatewayAuthShrinkRequest) GetTokenNamePrefix() *string {
	return s.TokenNamePrefix
}

func (s *UpdateGatewayAuthShrinkRequest) GetTokenPass() *bool {
	return s.TokenPass
}

func (s *UpdateGatewayAuthShrinkRequest) GetTokenPosition() *string {
	return s.TokenPosition
}

func (s *UpdateGatewayAuthShrinkRequest) GetType() *string {
	return s.Type
}

func (s *UpdateGatewayAuthShrinkRequest) SetAcceptLanguage(v string) *UpdateGatewayAuthShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetAuthResourceConfig(v string) *UpdateGatewayAuthShrinkRequest {
	s.AuthResourceConfig = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetAuthResourceListShrink(v string) *UpdateGatewayAuthShrinkRequest {
	s.AuthResourceListShrink = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetAuthResourceMode(v int32) *UpdateGatewayAuthShrinkRequest {
	s.AuthResourceMode = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetClientId(v string) *UpdateGatewayAuthShrinkRequest {
	s.ClientId = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetClientSecret(v string) *UpdateGatewayAuthShrinkRequest {
	s.ClientSecret = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetCookieDomain(v string) *UpdateGatewayAuthShrinkRequest {
	s.CookieDomain = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetDeleteResourceIdListShrink(v string) *UpdateGatewayAuthShrinkRequest {
	s.DeleteResourceIdListShrink = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetExternalAuthZJSONShrink(v string) *UpdateGatewayAuthShrinkRequest {
	s.ExternalAuthZJSONShrink = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetGatewayUniqueId(v string) *UpdateGatewayAuthShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetId(v int64) *UpdateGatewayAuthShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetIsWhite(v bool) *UpdateGatewayAuthShrinkRequest {
	s.IsWhite = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetIssuer(v string) *UpdateGatewayAuthShrinkRequest {
	s.Issuer = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetJwks(v string) *UpdateGatewayAuthShrinkRequest {
	s.Jwks = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetLoginUrl(v string) *UpdateGatewayAuthShrinkRequest {
	s.LoginUrl = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetName(v string) *UpdateGatewayAuthShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetRedirectUrl(v string) *UpdateGatewayAuthShrinkRequest {
	s.RedirectUrl = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetScopesListShrink(v string) *UpdateGatewayAuthShrinkRequest {
	s.ScopesListShrink = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetStatus(v bool) *UpdateGatewayAuthShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetSub(v string) *UpdateGatewayAuthShrinkRequest {
	s.Sub = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetTokenName(v string) *UpdateGatewayAuthShrinkRequest {
	s.TokenName = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetTokenNamePrefix(v string) *UpdateGatewayAuthShrinkRequest {
	s.TokenNamePrefix = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetTokenPass(v bool) *UpdateGatewayAuthShrinkRequest {
	s.TokenPass = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetTokenPosition(v string) *UpdateGatewayAuthShrinkRequest {
	s.TokenPosition = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) SetType(v string) *UpdateGatewayAuthShrinkRequest {
	s.Type = &v
	return s
}

func (s *UpdateGatewayAuthShrinkRequest) Validate() error {
	return dara.Validate(s)
}
