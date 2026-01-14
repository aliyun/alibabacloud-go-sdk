// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayAuthRequest
	GetAcceptLanguage() *string
	SetAuthResourceConfig(v string) *UpdateGatewayAuthRequest
	GetAuthResourceConfig() *string
	SetAuthResourceList(v []*UpdateGatewayAuthRequestAuthResourceList) *UpdateGatewayAuthRequest
	GetAuthResourceList() []*UpdateGatewayAuthRequestAuthResourceList
	SetAuthResourceMode(v int32) *UpdateGatewayAuthRequest
	GetAuthResourceMode() *int32
	SetClientId(v string) *UpdateGatewayAuthRequest
	GetClientId() *string
	SetClientSecret(v string) *UpdateGatewayAuthRequest
	GetClientSecret() *string
	SetCookieDomain(v string) *UpdateGatewayAuthRequest
	GetCookieDomain() *string
	SetDeleteResourceIdList(v []*int64) *UpdateGatewayAuthRequest
	GetDeleteResourceIdList() []*int64
	SetExternalAuthZJSON(v *UpdateGatewayAuthRequestExternalAuthZJSON) *UpdateGatewayAuthRequest
	GetExternalAuthZJSON() *UpdateGatewayAuthRequestExternalAuthZJSON
	SetGatewayUniqueId(v string) *UpdateGatewayAuthRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayAuthRequest
	GetId() *int64
	SetIsWhite(v bool) *UpdateGatewayAuthRequest
	GetIsWhite() *bool
	SetIssuer(v string) *UpdateGatewayAuthRequest
	GetIssuer() *string
	SetJwks(v string) *UpdateGatewayAuthRequest
	GetJwks() *string
	SetLoginUrl(v string) *UpdateGatewayAuthRequest
	GetLoginUrl() *string
	SetName(v string) *UpdateGatewayAuthRequest
	GetName() *string
	SetRedirectUrl(v string) *UpdateGatewayAuthRequest
	GetRedirectUrl() *string
	SetScopesList(v []*string) *UpdateGatewayAuthRequest
	GetScopesList() []*string
	SetStatus(v bool) *UpdateGatewayAuthRequest
	GetStatus() *bool
	SetSub(v string) *UpdateGatewayAuthRequest
	GetSub() *string
	SetTokenName(v string) *UpdateGatewayAuthRequest
	GetTokenName() *string
	SetTokenNamePrefix(v string) *UpdateGatewayAuthRequest
	GetTokenNamePrefix() *string
	SetTokenPass(v bool) *UpdateGatewayAuthRequest
	GetTokenPass() *bool
	SetTokenPosition(v string) *UpdateGatewayAuthRequest
	GetTokenPosition() *string
	SetType(v string) *UpdateGatewayAuthRequest
	GetType() *string
}

type UpdateGatewayAuthRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// \\"\\"
	AuthResourceConfig *string                                     `json:"AuthResourceConfig,omitempty" xml:"AuthResourceConfig,omitempty"`
	AuthResourceList   []*UpdateGatewayAuthRequestAuthResourceList `json:"AuthResourceList,omitempty" xml:"AuthResourceList,omitempty" type:"Repeated"`
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
	CookieDomain         *string                                    `json:"CookieDomain,omitempty" xml:"CookieDomain,omitempty"`
	DeleteResourceIdList []*int64                                   `json:"DeleteResourceIdList,omitempty" xml:"DeleteResourceIdList,omitempty" type:"Repeated"`
	ExternalAuthZJSON    *UpdateGatewayAuthRequestExternalAuthZJSON `json:"ExternalAuthZJSON,omitempty" xml:"ExternalAuthZJSON,omitempty" type:"Struct"`
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
	RedirectUrl *string   `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	ScopesList  []*string `json:"ScopesList,omitempty" xml:"ScopesList,omitempty" type:"Repeated"`
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

func (s UpdateGatewayAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayAuthRequest) GetAuthResourceConfig() *string {
	return s.AuthResourceConfig
}

func (s *UpdateGatewayAuthRequest) GetAuthResourceList() []*UpdateGatewayAuthRequestAuthResourceList {
	return s.AuthResourceList
}

func (s *UpdateGatewayAuthRequest) GetAuthResourceMode() *int32 {
	return s.AuthResourceMode
}

func (s *UpdateGatewayAuthRequest) GetClientId() *string {
	return s.ClientId
}

func (s *UpdateGatewayAuthRequest) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *UpdateGatewayAuthRequest) GetCookieDomain() *string {
	return s.CookieDomain
}

func (s *UpdateGatewayAuthRequest) GetDeleteResourceIdList() []*int64 {
	return s.DeleteResourceIdList
}

func (s *UpdateGatewayAuthRequest) GetExternalAuthZJSON() *UpdateGatewayAuthRequestExternalAuthZJSON {
	return s.ExternalAuthZJSON
}

func (s *UpdateGatewayAuthRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayAuthRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayAuthRequest) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *UpdateGatewayAuthRequest) GetIssuer() *string {
	return s.Issuer
}

func (s *UpdateGatewayAuthRequest) GetJwks() *string {
	return s.Jwks
}

func (s *UpdateGatewayAuthRequest) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *UpdateGatewayAuthRequest) GetName() *string {
	return s.Name
}

func (s *UpdateGatewayAuthRequest) GetRedirectUrl() *string {
	return s.RedirectUrl
}

func (s *UpdateGatewayAuthRequest) GetScopesList() []*string {
	return s.ScopesList
}

func (s *UpdateGatewayAuthRequest) GetStatus() *bool {
	return s.Status
}

func (s *UpdateGatewayAuthRequest) GetSub() *string {
	return s.Sub
}

func (s *UpdateGatewayAuthRequest) GetTokenName() *string {
	return s.TokenName
}

func (s *UpdateGatewayAuthRequest) GetTokenNamePrefix() *string {
	return s.TokenNamePrefix
}

func (s *UpdateGatewayAuthRequest) GetTokenPass() *bool {
	return s.TokenPass
}

func (s *UpdateGatewayAuthRequest) GetTokenPosition() *string {
	return s.TokenPosition
}

func (s *UpdateGatewayAuthRequest) GetType() *string {
	return s.Type
}

func (s *UpdateGatewayAuthRequest) SetAcceptLanguage(v string) *UpdateGatewayAuthRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetAuthResourceConfig(v string) *UpdateGatewayAuthRequest {
	s.AuthResourceConfig = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetAuthResourceList(v []*UpdateGatewayAuthRequestAuthResourceList) *UpdateGatewayAuthRequest {
	s.AuthResourceList = v
	return s
}

func (s *UpdateGatewayAuthRequest) SetAuthResourceMode(v int32) *UpdateGatewayAuthRequest {
	s.AuthResourceMode = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetClientId(v string) *UpdateGatewayAuthRequest {
	s.ClientId = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetClientSecret(v string) *UpdateGatewayAuthRequest {
	s.ClientSecret = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetCookieDomain(v string) *UpdateGatewayAuthRequest {
	s.CookieDomain = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetDeleteResourceIdList(v []*int64) *UpdateGatewayAuthRequest {
	s.DeleteResourceIdList = v
	return s
}

func (s *UpdateGatewayAuthRequest) SetExternalAuthZJSON(v *UpdateGatewayAuthRequestExternalAuthZJSON) *UpdateGatewayAuthRequest {
	s.ExternalAuthZJSON = v
	return s
}

func (s *UpdateGatewayAuthRequest) SetGatewayUniqueId(v string) *UpdateGatewayAuthRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetId(v int64) *UpdateGatewayAuthRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetIsWhite(v bool) *UpdateGatewayAuthRequest {
	s.IsWhite = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetIssuer(v string) *UpdateGatewayAuthRequest {
	s.Issuer = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetJwks(v string) *UpdateGatewayAuthRequest {
	s.Jwks = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetLoginUrl(v string) *UpdateGatewayAuthRequest {
	s.LoginUrl = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetName(v string) *UpdateGatewayAuthRequest {
	s.Name = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetRedirectUrl(v string) *UpdateGatewayAuthRequest {
	s.RedirectUrl = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetScopesList(v []*string) *UpdateGatewayAuthRequest {
	s.ScopesList = v
	return s
}

func (s *UpdateGatewayAuthRequest) SetStatus(v bool) *UpdateGatewayAuthRequest {
	s.Status = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetSub(v string) *UpdateGatewayAuthRequest {
	s.Sub = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetTokenName(v string) *UpdateGatewayAuthRequest {
	s.TokenName = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetTokenNamePrefix(v string) *UpdateGatewayAuthRequest {
	s.TokenNamePrefix = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetTokenPass(v bool) *UpdateGatewayAuthRequest {
	s.TokenPass = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetTokenPosition(v string) *UpdateGatewayAuthRequest {
	s.TokenPosition = &v
	return s
}

func (s *UpdateGatewayAuthRequest) SetType(v string) *UpdateGatewayAuthRequest {
	s.Type = &v
	return s
}

func (s *UpdateGatewayAuthRequest) Validate() error {
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

type UpdateGatewayAuthRequestAuthResourceList struct {
	AuthResourceHeaderList []*UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList `json:"AuthResourceHeaderList,omitempty" xml:"AuthResourceHeaderList,omitempty" type:"Repeated"`
	// example:
	//
	// 1765
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// example:
	//
	// 1
	Id         *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	IgnoreCase *bool  `json:"IgnoreCase,omitempty" xml:"IgnoreCase,omitempty"`
	// example:
	//
	// EQUAL
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s UpdateGatewayAuthRequestAuthResourceList) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthRequestAuthResourceList) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthRequestAuthResourceList) GetAuthResourceHeaderList() []*UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList {
	return s.AuthResourceHeaderList
}

func (s *UpdateGatewayAuthRequestAuthResourceList) GetDomainId() *int64 {
	return s.DomainId
}

func (s *UpdateGatewayAuthRequestAuthResourceList) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayAuthRequestAuthResourceList) GetIgnoreCase() *bool {
	return s.IgnoreCase
}

func (s *UpdateGatewayAuthRequestAuthResourceList) GetMatchType() *string {
	return s.MatchType
}

func (s *UpdateGatewayAuthRequestAuthResourceList) GetPath() *string {
	return s.Path
}

func (s *UpdateGatewayAuthRequestAuthResourceList) SetAuthResourceHeaderList(v []*UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList) *UpdateGatewayAuthRequestAuthResourceList {
	s.AuthResourceHeaderList = v
	return s
}

func (s *UpdateGatewayAuthRequestAuthResourceList) SetDomainId(v int64) *UpdateGatewayAuthRequestAuthResourceList {
	s.DomainId = &v
	return s
}

func (s *UpdateGatewayAuthRequestAuthResourceList) SetId(v int64) *UpdateGatewayAuthRequestAuthResourceList {
	s.Id = &v
	return s
}

func (s *UpdateGatewayAuthRequestAuthResourceList) SetIgnoreCase(v bool) *UpdateGatewayAuthRequestAuthResourceList {
	s.IgnoreCase = &v
	return s
}

func (s *UpdateGatewayAuthRequestAuthResourceList) SetMatchType(v string) *UpdateGatewayAuthRequestAuthResourceList {
	s.MatchType = &v
	return s
}

func (s *UpdateGatewayAuthRequestAuthResourceList) SetPath(v string) *UpdateGatewayAuthRequestAuthResourceList {
	s.Path = &v
	return s
}

func (s *UpdateGatewayAuthRequestAuthResourceList) Validate() error {
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

type UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList struct {
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

func (s UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList) GetHeaderMethod() *string {
	return s.HeaderMethod
}

func (s *UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList) GetHeaderValue() *string {
	return s.HeaderValue
}

func (s *UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList) SetHeaderKey(v string) *UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList {
	s.HeaderKey = &v
	return s
}

func (s *UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList) SetHeaderMethod(v string) *UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList {
	s.HeaderMethod = &v
	return s
}

func (s *UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList) SetHeaderValue(v string) *UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList {
	s.HeaderValue = &v
	return s
}

func (s *UpdateGatewayAuthRequestAuthResourceListAuthResourceHeaderList) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayAuthRequestExternalAuthZJSON struct {
	AllowRequestHeaders  []*string `json:"AllowRequestHeaders,omitempty" xml:"AllowRequestHeaders,omitempty" type:"Repeated"`
	AllowUpstreamHeaders []*string `json:"AllowUpstreamHeaders,omitempty" xml:"AllowUpstreamHeaders,omitempty" type:"Repeated"`
	// example:
	//
	// 4000000
	BodyMaxBytes *int32 `json:"BodyMaxBytes,omitempty" xml:"BodyMaxBytes,omitempty"`
	// example:
	//
	// true
	IsRestrict *bool `json:"IsRestrict,omitempty" xml:"IsRestrict,omitempty"`
	// example:
	//
	// /auth
	PrefixPath *string `json:"PrefixPath,omitempty" xml:"PrefixPath,omitempty"`
	// example:
	//
	// 37396
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 10
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// Authorization
	TokenKey *string `json:"TokenKey,omitempty" xml:"TokenKey,omitempty"`
	// example:
	//
	// true
	WithRematchRoute *bool `json:"WithRematchRoute,omitempty" xml:"WithRematchRoute,omitempty"`
	// example:
	//
	// true
	WithRequestBody *bool `json:"WithRequestBody,omitempty" xml:"WithRequestBody,omitempty"`
}

func (s UpdateGatewayAuthRequestExternalAuthZJSON) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthRequestExternalAuthZJSON) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) GetAllowRequestHeaders() []*string {
	return s.AllowRequestHeaders
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) GetAllowUpstreamHeaders() []*string {
	return s.AllowUpstreamHeaders
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) GetBodyMaxBytes() *int32 {
	return s.BodyMaxBytes
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) GetIsRestrict() *bool {
	return s.IsRestrict
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) GetPrefixPath() *string {
	return s.PrefixPath
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) GetTokenKey() *string {
	return s.TokenKey
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) GetWithRematchRoute() *bool {
	return s.WithRematchRoute
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) GetWithRequestBody() *bool {
	return s.WithRequestBody
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) SetAllowRequestHeaders(v []*string) *UpdateGatewayAuthRequestExternalAuthZJSON {
	s.AllowRequestHeaders = v
	return s
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) SetAllowUpstreamHeaders(v []*string) *UpdateGatewayAuthRequestExternalAuthZJSON {
	s.AllowUpstreamHeaders = v
	return s
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) SetBodyMaxBytes(v int32) *UpdateGatewayAuthRequestExternalAuthZJSON {
	s.BodyMaxBytes = &v
	return s
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) SetIsRestrict(v bool) *UpdateGatewayAuthRequestExternalAuthZJSON {
	s.IsRestrict = &v
	return s
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) SetPrefixPath(v string) *UpdateGatewayAuthRequestExternalAuthZJSON {
	s.PrefixPath = &v
	return s
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) SetServiceId(v int64) *UpdateGatewayAuthRequestExternalAuthZJSON {
	s.ServiceId = &v
	return s
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) SetTimeout(v int32) *UpdateGatewayAuthRequestExternalAuthZJSON {
	s.Timeout = &v
	return s
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) SetTokenKey(v string) *UpdateGatewayAuthRequestExternalAuthZJSON {
	s.TokenKey = &v
	return s
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) SetWithRematchRoute(v bool) *UpdateGatewayAuthRequestExternalAuthZJSON {
	s.WithRematchRoute = &v
	return s
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) SetWithRequestBody(v bool) *UpdateGatewayAuthRequestExternalAuthZJSON {
	s.WithRequestBody = &v
	return s
}

func (s *UpdateGatewayAuthRequestExternalAuthZJSON) Validate() error {
	return dara.Validate(s)
}
