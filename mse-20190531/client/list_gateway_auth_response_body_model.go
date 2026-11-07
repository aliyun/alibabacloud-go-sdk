// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListGatewayAuthResponseBody
	GetCode() *int32
	SetData(v *ListGatewayAuthResponseBodyData) *ListGatewayAuthResponseBody
	GetData() *ListGatewayAuthResponseBodyData
	SetHttpStatusCode(v int32) *ListGatewayAuthResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListGatewayAuthResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewayAuthResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGatewayAuthResponseBody
	GetSuccess() *bool
}

type ListGatewayAuthResponseBody struct {
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListGatewayAuthResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Request processed successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EC1EED4A-B147-597B-B949-FD3131AB4298
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGatewayAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListGatewayAuthResponseBody) GetData() *ListGatewayAuthResponseBodyData {
	return s.Data
}

func (s *ListGatewayAuthResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGatewayAuthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayAuthResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGatewayAuthResponseBody) SetCode(v int32) *ListGatewayAuthResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayAuthResponseBody) SetData(v *ListGatewayAuthResponseBodyData) *ListGatewayAuthResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayAuthResponseBody) SetHttpStatusCode(v int32) *ListGatewayAuthResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGatewayAuthResponseBody) SetMessage(v string) *ListGatewayAuthResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayAuthResponseBody) SetRequestId(v string) *ListGatewayAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayAuthResponseBody) SetSuccess(v bool) *ListGatewayAuthResponseBody {
	s.Success = &v
	return s
}

func (s *ListGatewayAuthResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayAuthResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned results.
	Result []*ListGatewayAuthResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 9
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListGatewayAuthResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayAuthResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayAuthResponseBodyData) GetResult() []*ListGatewayAuthResponseBodyDataResult {
	return s.Result
}

func (s *ListGatewayAuthResponseBodyData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *ListGatewayAuthResponseBodyData) SetPageNumber(v int32) *ListGatewayAuthResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayAuthResponseBodyData) SetPageSize(v int32) *ListGatewayAuthResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGatewayAuthResponseBodyData) SetResult(v []*ListGatewayAuthResponseBodyDataResult) *ListGatewayAuthResponseBodyData {
	s.Result = v
	return s
}

func (s *ListGatewayAuthResponseBodyData) SetTotalSize(v int64) *ListGatewayAuthResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListGatewayAuthResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayAuthResponseBodyDataResult struct {
	// The authorization resource mode. Valid values:
	//
	// - 0: simple mode.
	//
	// - 1: complex mode.
	//
	// example:
	//
	// 0
	AuthResourceMode *int32 `json:"AuthResourceMode,omitempty" xml:"AuthResourceMode,omitempty"`
	// The application ID registered in the service. This parameter is used for OIDC or IDaaS.
	//
	// example:
	//
	// app_mnvxaavggw7hcdcnr6usi6***
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The application secret registered in the service. This parameter is used for OIDC or IDaaS.
	//
	// example:
	//
	// CS6EYfx3k9yTRR9EtQ2MXWP97P6UAUwFg4teoWJ19Z****
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The domain name of the cookie. After authentication, the cookie is sent to the specified domain name to maintain the logon status. This parameter is used for OIDC or IDaaS.
	//
	// example:
	//
	// test.com
	CookieDomain *string `json:"CookieDomain,omitempty" xml:"CookieDomain,omitempty"`
	// The custom authentication information.
	ExternalAuthZ *ListGatewayAuthResponseBodyDataResultExternalAuthZ `json:"ExternalAuthZ,omitempty" xml:"ExternalAuthZ,omitempty" type:"Struct"`
	// The gateway ID.
	//
	// example:
	//
	// 1
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-e2d226bba4b2445c9e29fa7f8216****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-09-13 19:24:23
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2021-09-13 19:24:23
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The authentication ID.
	//
	// example:
	//
	// 549
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The authorization blacklists and whitelists mode. Valid values:
	//
	// - true: whitelist mode. Hosts and paths in the whitelist can be accessed without verification, while all other hosts and paths require verification.
	//
	// - false: blacklist mode. Hosts and paths in the blacklist require verification, while all other hosts and paths can be accessed directly.
	//
	// example:
	//
	// true
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// Used when the authentication type is JWT or OIDC:
	//
	// - JWT: the iss (issuer) claim in the JWT, which represents the issuer.
	//
	// - OIDC: the iss (issuer) claim in the OIDC token, which represents the issuer.
	//
	// example:
	//
	// https://example.com/auth
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The JWT public key in JSON format.
	//
	// example:
	//
	// {\\"keys\\":[{\\"kty\\":\\"oct\\",\\"k\\":\\"9V9lpiuAQsME1efQChI0kEQz6fdVlJbDRFFa1lvEB_A\\",\\"alg\\":\\"HS256\\"}]}
	Jwks *string `json:"Jwks,omitempty" xml:"Jwks,omitempty"`
	// The URL of the user logon page for the IDaaS instance. This parameter is used for IDaaS.
	//
	// example:
	//
	// https://daxxxxcn.aliyunidaas.com/
	LoginUrl *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	// The name of the authentication rule.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The redirect URL after successful authorization. This parameter is used for OIDC or IDaaS and must be consistent with the redirect URL configured in OIDC or IDaaS.
	//
	// example:
	//
	// http://test.com/oauth2/callback
	RedirectUrl *string `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	// The OIDC authorization scopes. This parameter is used for OIDC.
	//
	// example:
	//
	// ["openid"]
	ScopesList *string `json:"ScopesList,omitempty" xml:"ScopesList,omitempty"`
	// The enabling status. Valid values:
	//
	// - false: disabled.
	//
	// - true: enabled.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// Used when the authentication type is JWT. The sub (subject) claim in the JWT, which represents the subject.
	//
	// example:
	//
	// example-app
	Sub *string `json:"Sub,omitempty" xml:"Sub,omitempty"`
	// The location where the JWT token is stored.
	//
	// By default, the token to be verified is prefixed with Bearer and placed in the Authorization header, for example, `Authorization: Bearer token`.
	//
	// example:
	//
	// Authorization
	TokenName *string `json:"TokenName,omitempty" xml:"TokenName,omitempty"`
	// The prefix of the JWT token.
	//
	// By default, the token to be verified is prefixed with Bearer and placed in the Authorization header, for example, `Authorization: Bearer token`.
	//
	// example:
	//
	// Bearer
	TokenNamePrefix *string `json:"TokenNamePrefix,omitempty" xml:"TokenNamePrefix,omitempty"`
	// Specifies whether to pass through the token.
	//
	// example:
	//
	// true
	TokenPass *bool `json:"TokenPass,omitempty" xml:"TokenPass,omitempty"`
	// The JWT token position. Valid values:
	//
	// - HEADER: transmitted through the header.
	//
	// By default, the token to be verified is prefixed with Bearer and placed in the Authorization header, for example, Authorization: Bearer token.
	//
	// example:
	//
	// HEADER
	TokenPosition *string `json:"TokenPosition,omitempty" xml:"TokenPosition,omitempty"`
	// The authentication type. Valid values:
	//
	// - JWT
	//
	// - OIDC
	//
	// - IDaaS
	//
	// - ExternalAuthZ: self-managed authentication service
	//
	// example:
	//
	// JWT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayAuthResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthResponseBodyDataResult) GetAuthResourceMode() *int32 {
	return s.AuthResourceMode
}

func (s *ListGatewayAuthResponseBodyDataResult) GetClientId() *string {
	return s.ClientId
}

func (s *ListGatewayAuthResponseBodyDataResult) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *ListGatewayAuthResponseBodyDataResult) GetCookieDomain() *string {
	return s.CookieDomain
}

func (s *ListGatewayAuthResponseBodyDataResult) GetExternalAuthZ() *ListGatewayAuthResponseBodyDataResultExternalAuthZ {
	return s.ExternalAuthZ
}

func (s *ListGatewayAuthResponseBodyDataResult) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *ListGatewayAuthResponseBodyDataResult) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayAuthResponseBodyDataResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListGatewayAuthResponseBodyDataResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListGatewayAuthResponseBodyDataResult) GetId() *int64 {
	return s.Id
}

func (s *ListGatewayAuthResponseBodyDataResult) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *ListGatewayAuthResponseBodyDataResult) GetIssuer() *string {
	return s.Issuer
}

func (s *ListGatewayAuthResponseBodyDataResult) GetJwks() *string {
	return s.Jwks
}

func (s *ListGatewayAuthResponseBodyDataResult) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *ListGatewayAuthResponseBodyDataResult) GetName() *string {
	return s.Name
}

func (s *ListGatewayAuthResponseBodyDataResult) GetRedirectUrl() *string {
	return s.RedirectUrl
}

func (s *ListGatewayAuthResponseBodyDataResult) GetScopesList() *string {
	return s.ScopesList
}

func (s *ListGatewayAuthResponseBodyDataResult) GetStatus() *bool {
	return s.Status
}

func (s *ListGatewayAuthResponseBodyDataResult) GetSub() *string {
	return s.Sub
}

func (s *ListGatewayAuthResponseBodyDataResult) GetTokenName() *string {
	return s.TokenName
}

func (s *ListGatewayAuthResponseBodyDataResult) GetTokenNamePrefix() *string {
	return s.TokenNamePrefix
}

func (s *ListGatewayAuthResponseBodyDataResult) GetTokenPass() *bool {
	return s.TokenPass
}

func (s *ListGatewayAuthResponseBodyDataResult) GetTokenPosition() *string {
	return s.TokenPosition
}

func (s *ListGatewayAuthResponseBodyDataResult) GetType() *string {
	return s.Type
}

func (s *ListGatewayAuthResponseBodyDataResult) SetAuthResourceMode(v int32) *ListGatewayAuthResponseBodyDataResult {
	s.AuthResourceMode = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetClientId(v string) *ListGatewayAuthResponseBodyDataResult {
	s.ClientId = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetClientSecret(v string) *ListGatewayAuthResponseBodyDataResult {
	s.ClientSecret = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetCookieDomain(v string) *ListGatewayAuthResponseBodyDataResult {
	s.CookieDomain = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetExternalAuthZ(v *ListGatewayAuthResponseBodyDataResultExternalAuthZ) *ListGatewayAuthResponseBodyDataResult {
	s.ExternalAuthZ = v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetGatewayId(v int64) *ListGatewayAuthResponseBodyDataResult {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetGatewayUniqueId(v string) *ListGatewayAuthResponseBodyDataResult {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetGmtCreate(v string) *ListGatewayAuthResponseBodyDataResult {
	s.GmtCreate = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetGmtModified(v string) *ListGatewayAuthResponseBodyDataResult {
	s.GmtModified = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetId(v int64) *ListGatewayAuthResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetIsWhite(v bool) *ListGatewayAuthResponseBodyDataResult {
	s.IsWhite = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetIssuer(v string) *ListGatewayAuthResponseBodyDataResult {
	s.Issuer = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetJwks(v string) *ListGatewayAuthResponseBodyDataResult {
	s.Jwks = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetLoginUrl(v string) *ListGatewayAuthResponseBodyDataResult {
	s.LoginUrl = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetName(v string) *ListGatewayAuthResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetRedirectUrl(v string) *ListGatewayAuthResponseBodyDataResult {
	s.RedirectUrl = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetScopesList(v string) *ListGatewayAuthResponseBodyDataResult {
	s.ScopesList = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetStatus(v bool) *ListGatewayAuthResponseBodyDataResult {
	s.Status = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetSub(v string) *ListGatewayAuthResponseBodyDataResult {
	s.Sub = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetTokenName(v string) *ListGatewayAuthResponseBodyDataResult {
	s.TokenName = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetTokenNamePrefix(v string) *ListGatewayAuthResponseBodyDataResult {
	s.TokenNamePrefix = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetTokenPass(v bool) *ListGatewayAuthResponseBodyDataResult {
	s.TokenPass = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetTokenPosition(v string) *ListGatewayAuthResponseBodyDataResult {
	s.TokenPosition = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) SetType(v string) *ListGatewayAuthResponseBodyDataResult {
	s.Type = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResult) Validate() error {
	if s.ExternalAuthZ != nil {
		if err := s.ExternalAuthZ.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayAuthResponseBodyDataResultExternalAuthZ struct {
	// The names of the allowed request headers.
	AllowRequestHeaders []*string `json:"AllowRequestHeaders,omitempty" xml:"AllowRequestHeaders,omitempty" type:"Repeated"`
	// The names of the allowed response headers.
	AllowUpstreamHeaders []*string `json:"AllowUpstreamHeaders,omitempty" xml:"AllowUpstreamHeaders,omitempty" type:"Repeated"`
	// The maximum number of bytes for the request body.
	//
	// example:
	//
	// 1024
	BodyMaxBytes *int32 `json:"BodyMaxBytes,omitempty" xml:"BodyMaxBytes,omitempty"`
	// The custom authentication mode. Valid values:
	//
	// - true: strict mode. When the authentication service is unavailable (connection establishment fails or a 5xx status code is returned), the gateway rejects the client request.
	//
	// - false: loose mode. When the authentication service is unavailable (connection establishment fails or a 5xx status code is returned), the gateway allows the client request.
	//
	// example:
	//
	// true
	IsRestrict *bool `json:"IsRestrict,omitempty" xml:"IsRestrict,omitempty"`
	// The path of the authentication API provided by the authentication service. The API path must be prefix-matched.
	//
	// example:
	//
	// /auth
	PrefixPath *string `json:"PrefixPath,omitempty" xml:"PrefixPath,omitempty"`
	// The authentication service information.
	Service *ListGatewayAuthResponseBodyDataResultExternalAuthZService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	// The ID of the authentication service.
	//
	// example:
	//
	// 45186
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The timeout period of the authentication service. Unit: seconds.
	//
	// example:
	//
	// 30
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The token location, which indicates the header in the request message where the token is located. Common values include Authorization and Cookie.
	//
	// example:
	//
	// Authorization
	TokenKey *string `json:"TokenKey,omitempty" xml:"TokenKey,omitempty"`
	// Specifies whether to rematch the route after authentication is completed.
	//
	// example:
	//
	// true
	WithRematchRoute *bool `json:"WithRematchRoute,omitempty" xml:"WithRematchRoute,omitempty"`
	// Specifies whether to include the request body.
	//
	// example:
	//
	// true
	WithRequestBody *bool `json:"WithRequestBody,omitempty" xml:"WithRequestBody,omitempty"`
}

func (s ListGatewayAuthResponseBodyDataResultExternalAuthZ) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthResponseBodyDataResultExternalAuthZ) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) GetAllowRequestHeaders() []*string {
	return s.AllowRequestHeaders
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) GetAllowUpstreamHeaders() []*string {
	return s.AllowUpstreamHeaders
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) GetBodyMaxBytes() *int32 {
	return s.BodyMaxBytes
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) GetIsRestrict() *bool {
	return s.IsRestrict
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) GetPrefixPath() *string {
	return s.PrefixPath
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) GetService() *ListGatewayAuthResponseBodyDataResultExternalAuthZService {
	return s.Service
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) GetTokenKey() *string {
	return s.TokenKey
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) GetWithRematchRoute() *bool {
	return s.WithRematchRoute
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) GetWithRequestBody() *bool {
	return s.WithRequestBody
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) SetAllowRequestHeaders(v []*string) *ListGatewayAuthResponseBodyDataResultExternalAuthZ {
	s.AllowRequestHeaders = v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) SetAllowUpstreamHeaders(v []*string) *ListGatewayAuthResponseBodyDataResultExternalAuthZ {
	s.AllowUpstreamHeaders = v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) SetBodyMaxBytes(v int32) *ListGatewayAuthResponseBodyDataResultExternalAuthZ {
	s.BodyMaxBytes = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) SetIsRestrict(v bool) *ListGatewayAuthResponseBodyDataResultExternalAuthZ {
	s.IsRestrict = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) SetPrefixPath(v string) *ListGatewayAuthResponseBodyDataResultExternalAuthZ {
	s.PrefixPath = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) SetService(v *ListGatewayAuthResponseBodyDataResultExternalAuthZService) *ListGatewayAuthResponseBodyDataResultExternalAuthZ {
	s.Service = v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) SetServiceId(v int64) *ListGatewayAuthResponseBodyDataResultExternalAuthZ {
	s.ServiceId = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) SetTimeout(v int32) *ListGatewayAuthResponseBodyDataResultExternalAuthZ {
	s.Timeout = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) SetTokenKey(v string) *ListGatewayAuthResponseBodyDataResultExternalAuthZ {
	s.TokenKey = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) SetWithRematchRoute(v bool) *ListGatewayAuthResponseBodyDataResultExternalAuthZ {
	s.WithRematchRoute = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) SetWithRequestBody(v bool) *ListGatewayAuthResponseBodyDataResultExternalAuthZ {
	s.WithRequestBody = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZ) Validate() error {
	if s.Service != nil {
		if err := s.Service.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayAuthResponseBodyDataResultExternalAuthZService struct {
	// The name of the service group.
	//
	// example:
	//
	// com.mse.console.test.3YXO
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The authentication service name.
	//
	// example:
	//
	// updatetime
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source type of the authentication service.
	//
	// example:
	//
	// MSE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListGatewayAuthResponseBodyDataResultExternalAuthZService) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthResponseBodyDataResultExternalAuthZService) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZService) GetGroupName() *string {
	return s.GroupName
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZService) GetName() *string {
	return s.Name
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZService) GetNamespace() *string {
	return s.Namespace
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZService) GetSourceType() *string {
	return s.SourceType
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZService) SetGroupName(v string) *ListGatewayAuthResponseBodyDataResultExternalAuthZService {
	s.GroupName = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZService) SetName(v string) *ListGatewayAuthResponseBodyDataResultExternalAuthZService {
	s.Name = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZService) SetNamespace(v string) *ListGatewayAuthResponseBodyDataResultExternalAuthZService {
	s.Namespace = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZService) SetSourceType(v string) *ListGatewayAuthResponseBodyDataResultExternalAuthZService {
	s.SourceType = &v
	return s
}

func (s *ListGatewayAuthResponseBodyDataResultExternalAuthZService) Validate() error {
	return dara.Validate(s)
}
