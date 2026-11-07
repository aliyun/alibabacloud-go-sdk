// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayAuthDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetGatewayAuthDetailResponseBody
	GetCode() *int32
	SetData(v *GetGatewayAuthDetailResponseBodyData) *GetGatewayAuthDetailResponseBody
	GetData() *GetGatewayAuthDetailResponseBodyData
	SetHttpStatusCode(v int32) *GetGatewayAuthDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetGatewayAuthDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayAuthDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGatewayAuthDetailResponseBody
	GetSuccess() *bool
}

type GetGatewayAuthDetailResponseBody struct {
	// The response code. It synchronizes with the Success parameter. If Success is true, the code is 200.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The authentication details.
	Data *GetGatewayAuthDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9C96CDF8-9E6C-XXXX-XXXX-8F87A10117E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayAuthDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayAuthDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayAuthDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetGatewayAuthDetailResponseBody) GetData() *GetGatewayAuthDetailResponseBodyData {
	return s.Data
}

func (s *GetGatewayAuthDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetGatewayAuthDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayAuthDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayAuthDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGatewayAuthDetailResponseBody) SetCode(v int32) *GetGatewayAuthDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBody) SetData(v *GetGatewayAuthDetailResponseBodyData) *GetGatewayAuthDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayAuthDetailResponseBody) SetHttpStatusCode(v int32) *GetGatewayAuthDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBody) SetMessage(v string) *GetGatewayAuthDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBody) SetRequestId(v string) *GetGatewayAuthDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBody) SetSuccess(v bool) *GetGatewayAuthDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayAuthDetailResponseBodyData struct {
	// The YAML configuration for the complex mode.
	//
	// example:
	//
	// permissions:
	//
	//   - not_rule:
	//
	//       or_rules:
	//
	//         rules:
	//
	//         - and_rules:
	//
	//             rules:
	//
	//             - url_path:
	//
	//                 path:
	//
	//                   ignore_case: false
	//
	//                   prefix: /version
	//
	//         - and_rules:
	//
	//             rules:
	//
	//             - url_path:
	//
	//                 path:
	//
	//                   ignore_case: false
	//
	//                   prefix: /header
	AuthResourceConfig *string `json:"AuthResourceConfig,omitempty" xml:"AuthResourceConfig,omitempty"`
	// The authorization resource mode. Valid values:
	//
	// - 0: Simple mode.
	//
	// - 1: Complex mode.
	//
	// example:
	//
	// 0
	AuthResourceMode *int32 `json:"AuthResourceMode,omitempty" xml:"AuthResourceMode,omitempty"`
	// The application ID registered for the service. This parameter is used for OIDC and IDaaS.
	//
	// example:
	//
	// example-app
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The application secret registered for the service. This parameter is used for OIDC and IDaaS.
	//
	// example:
	//
	// xxxxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The domain name of the cookie. After authentication, the cookie is sent to the specified domain name to maintain the logon state. This parameter is used for OIDC and IDaaS.
	//
	// example:
	//
	// hello.com
	CookieDomain *string `json:"CookieDomain,omitempty" xml:"CookieDomain,omitempty"`
	// The custom authentication information.
	ExternalAuthZ *GetGatewayAuthDetailResponseBodyDataExternalAuthZ `json:"ExternalAuthZ,omitempty" xml:"ExternalAuthZ,omitempty" type:"Struct"`
	// The ID of the gateway to which the authentication belongs.
	//
	// example:
	//
	// 2274
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway to which the authentication belongs.
	//
	// example:
	//
	// gw-6f0dbd108a0249d2b675b3ef50b*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The GMT time when the authentication was created.
	//
	// example:
	//
	// 2024-02-19T02:41:03.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The GMT time when the authentication was modified.
	//
	// example:
	//
	// 2024-02-19T02:41:03.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The authentication ID.
	//
	// example:
	//
	// 1100
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Authorization blacklist and whitelist mode:
	//
	// - true: whitelist mode (hosts and paths in the list can be accessed without authentication, while all others require authentication).
	//
	// - false: blacklist mode (hosts and paths in the list require authentication, while all others can be accessed directly).
	//
	// example:
	//
	// true
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// Used when the authentication type is JWT or OIDC. For JWT, it represents the iss (issuer) claim in the JWT. For OIDC, it represents the iss (issuer) claim in the OIDC token.
	//
	// example:
	//
	// https://example.com/auth
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The JWT public key in JSON format.
	//
	// example:
	//
	// {\\n  \\"keys\\":[\\n    {\\n      \\"kty\\": \\"RSA\\",\\n      \\"e\\": \\"AQAB\\",\\n      \\"use\\": \\"sig\\",\\n      \\"kid\\": \\"1rGufmH1YN8rqM9ZOLgo7eEST3AnL89Y-m-XGFioLoA\\",\\n      \\"alg\\": \\"RS256\\",\\n      \\"n\\": \\"rM2GIc0YTMqwNCwXnjKbW5QndkCEZgyLu3uQUnyZF7HvMTekiTvQg_39mg3dV1eaYYkYfZBogyroJBqAQXhk6VVCxlBjFVp2xstJPVWngMOOlcafwN_BKdN-EQ06O_Uu__e7gNKI3DunkNk0cNaFETE7d4meRYyTlgEzYgsrW05_ufR0BKoddL3E5JsCpUxRjH9ICbodBx0U74W6Dcci-R2EA1DBrEcboE6n90uoJs6UJNriAK_71nAsYonihU5aQFFnyPTkJHfRwHK6JlME6rn-b-rpLSpdyc6U1nOFZP2DEpz8U5FrYoLYSZIU-MQGxDhCnGc_rxl2IyP9B2qcCQ\\"\\n    }\\n  ]\\n}\\n
	Jwks *string `json:"Jwks,omitempty" xml:"Jwks,omitempty"`
	// The URL of the user logon page for the IDaaS instance. This parameter is used for IDaaS.
	//
	// example:
	//
	// https://daxxxxcn.aliyunidaas.com/
	LoginUrl *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	// The authentication name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The redirect URL after successful authorization. This parameter is used for OIDC and IDaaS, and must be consistent with the redirect URL configured in OIDC or IDaaS.
	//
	// example:
	//
	// https://yourdomain/path
	RedirectUrl *string `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	// The list of authorization rules for authentication. Multiple rule conditions have an OR relationship, and multiple matching items within a rule condition have an AND relationship.
	ResourceList []*GetGatewayAuthDetailResponseBodyDataResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// The OIDC scopes. This parameter is used for OIDC.
	//
	// example:
	//
	// ["openid","email"]
	ScopesList *string `json:"ScopesList,omitempty" xml:"ScopesList,omitempty"`
	// Indicates whether authentication is enabled.
	//
	// example:
	//
	// false
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// Used when the authentication type is JWT. It represents the sub (subject) claim in the JWT.
	//
	// example:
	//
	// https://example.com/auth
	Sub *string `json:"Sub,omitempty" xml:"Sub,omitempty"`
	// The location where the JWT token is stored.
	//
	// example:
	//
	// Authorization
	TokenName *string `json:"TokenName,omitempty" xml:"TokenName,omitempty"`
	// The prefix of the JWT token.
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
	// The position of the JWT token. Valid values:
	//
	// - HEADER: Transmitted through the HTTP header.
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
	// - ExternalAuthZ (self-managed authentication service)
	//
	// example:
	//
	// JWT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetGatewayAuthDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayAuthDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayAuthDetailResponseBodyData) GetAuthResourceConfig() *string {
	return s.AuthResourceConfig
}

func (s *GetGatewayAuthDetailResponseBodyData) GetAuthResourceMode() *int32 {
	return s.AuthResourceMode
}

func (s *GetGatewayAuthDetailResponseBodyData) GetClientId() *string {
	return s.ClientId
}

func (s *GetGatewayAuthDetailResponseBodyData) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *GetGatewayAuthDetailResponseBodyData) GetCookieDomain() *string {
	return s.CookieDomain
}

func (s *GetGatewayAuthDetailResponseBodyData) GetExternalAuthZ() *GetGatewayAuthDetailResponseBodyDataExternalAuthZ {
	return s.ExternalAuthZ
}

func (s *GetGatewayAuthDetailResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *GetGatewayAuthDetailResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayAuthDetailResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGatewayAuthDetailResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGatewayAuthDetailResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetGatewayAuthDetailResponseBodyData) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *GetGatewayAuthDetailResponseBodyData) GetIssuer() *string {
	return s.Issuer
}

func (s *GetGatewayAuthDetailResponseBodyData) GetJwks() *string {
	return s.Jwks
}

func (s *GetGatewayAuthDetailResponseBodyData) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *GetGatewayAuthDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetGatewayAuthDetailResponseBodyData) GetRedirectUrl() *string {
	return s.RedirectUrl
}

func (s *GetGatewayAuthDetailResponseBodyData) GetResourceList() []*GetGatewayAuthDetailResponseBodyDataResourceList {
	return s.ResourceList
}

func (s *GetGatewayAuthDetailResponseBodyData) GetScopesList() *string {
	return s.ScopesList
}

func (s *GetGatewayAuthDetailResponseBodyData) GetStatus() *bool {
	return s.Status
}

func (s *GetGatewayAuthDetailResponseBodyData) GetSub() *string {
	return s.Sub
}

func (s *GetGatewayAuthDetailResponseBodyData) GetTokenName() *string {
	return s.TokenName
}

func (s *GetGatewayAuthDetailResponseBodyData) GetTokenNamePrefix() *string {
	return s.TokenNamePrefix
}

func (s *GetGatewayAuthDetailResponseBodyData) GetTokenPass() *bool {
	return s.TokenPass
}

func (s *GetGatewayAuthDetailResponseBodyData) GetTokenPosition() *string {
	return s.TokenPosition
}

func (s *GetGatewayAuthDetailResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetGatewayAuthDetailResponseBodyData) SetAuthResourceConfig(v string) *GetGatewayAuthDetailResponseBodyData {
	s.AuthResourceConfig = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetAuthResourceMode(v int32) *GetGatewayAuthDetailResponseBodyData {
	s.AuthResourceMode = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetClientId(v string) *GetGatewayAuthDetailResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetClientSecret(v string) *GetGatewayAuthDetailResponseBodyData {
	s.ClientSecret = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetCookieDomain(v string) *GetGatewayAuthDetailResponseBodyData {
	s.CookieDomain = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetExternalAuthZ(v *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) *GetGatewayAuthDetailResponseBodyData {
	s.ExternalAuthZ = v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetGatewayId(v int64) *GetGatewayAuthDetailResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetGatewayUniqueId(v string) *GetGatewayAuthDetailResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetGmtCreate(v string) *GetGatewayAuthDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetGmtModified(v string) *GetGatewayAuthDetailResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetId(v int64) *GetGatewayAuthDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetIsWhite(v bool) *GetGatewayAuthDetailResponseBodyData {
	s.IsWhite = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetIssuer(v string) *GetGatewayAuthDetailResponseBodyData {
	s.Issuer = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetJwks(v string) *GetGatewayAuthDetailResponseBodyData {
	s.Jwks = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetLoginUrl(v string) *GetGatewayAuthDetailResponseBodyData {
	s.LoginUrl = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetName(v string) *GetGatewayAuthDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetRedirectUrl(v string) *GetGatewayAuthDetailResponseBodyData {
	s.RedirectUrl = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetResourceList(v []*GetGatewayAuthDetailResponseBodyDataResourceList) *GetGatewayAuthDetailResponseBodyData {
	s.ResourceList = v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetScopesList(v string) *GetGatewayAuthDetailResponseBodyData {
	s.ScopesList = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetStatus(v bool) *GetGatewayAuthDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetSub(v string) *GetGatewayAuthDetailResponseBodyData {
	s.Sub = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetTokenName(v string) *GetGatewayAuthDetailResponseBodyData {
	s.TokenName = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetTokenNamePrefix(v string) *GetGatewayAuthDetailResponseBodyData {
	s.TokenNamePrefix = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetTokenPass(v bool) *GetGatewayAuthDetailResponseBodyData {
	s.TokenPass = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetTokenPosition(v string) *GetGatewayAuthDetailResponseBodyData {
	s.TokenPosition = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) SetType(v string) *GetGatewayAuthDetailResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyData) Validate() error {
	if s.ExternalAuthZ != nil {
		if err := s.ExternalAuthZ.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceList != nil {
		for _, item := range s.ResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGatewayAuthDetailResponseBodyDataExternalAuthZ struct {
	// The allowed request headers.
	AllowRequestHeaders []*string `json:"AllowRequestHeaders,omitempty" xml:"AllowRequestHeaders,omitempty" type:"Repeated"`
	// The allowed response headers.
	AllowUpstreamHeaders []*string `json:"AllowUpstreamHeaders,omitempty" xml:"AllowUpstreamHeaders,omitempty" type:"Repeated"`
	// The maximum size of the request body in bytes.
	//
	// example:
	//
	// 4000000
	BodyMaxBytes *int32 `json:"BodyMaxBytes,omitempty" xml:"BodyMaxBytes,omitempty"`
	// Specifies the custom authentication mode. Valid values:
	//
	// - true: Strict mode. If the authentication service is unavailable (for example, the connection fails or a 5xx status code is returned), the gateway rejects the client request.
	//
	// - false: Loose mode. If the authentication service is unavailable (for example, the connection fails or a 5xx status code is returned), the gateway allows the client request.
	//
	// example:
	//
	// true
	IsRestrict *bool `json:"IsRestrict,omitempty" xml:"IsRestrict,omitempty"`
	// The path of the authentication API.
	//
	// example:
	//
	// /auth
	PrefixPath *string `json:"PrefixPath,omitempty" xml:"PrefixPath,omitempty"`
	// The authentication service information.
	Service *GetGatewayAuthDetailResponseBodyDataExternalAuthZService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	// The ID of the authentication service.
	//
	// example:
	//
	// 15300
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The timeout period of the authentication service, in seconds.
	//
	// example:
	//
	// 10
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The position of the token, which indicates the HTTP header that contains the token in the request. Common values include Authorization and Cookie.
	//
	// example:
	//
	// Authorization
	TokenKey *string `json:"TokenKey,omitempty" xml:"TokenKey,omitempty"`
	// Specifies whether to rematch the route after authentication is completed.
	WithRematchRoute *bool `json:"WithRematchRoute,omitempty" xml:"WithRematchRoute,omitempty"`
	// Specifies whether the request body is allowed.
	//
	// example:
	//
	// true
	WithRequestBody *bool `json:"WithRequestBody,omitempty" xml:"WithRequestBody,omitempty"`
}

func (s GetGatewayAuthDetailResponseBodyDataExternalAuthZ) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayAuthDetailResponseBodyDataExternalAuthZ) GoString() string {
	return s.String()
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) GetAllowRequestHeaders() []*string {
	return s.AllowRequestHeaders
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) GetAllowUpstreamHeaders() []*string {
	return s.AllowUpstreamHeaders
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) GetBodyMaxBytes() *int32 {
	return s.BodyMaxBytes
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) GetIsRestrict() *bool {
	return s.IsRestrict
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) GetPrefixPath() *string {
	return s.PrefixPath
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) GetService() *GetGatewayAuthDetailResponseBodyDataExternalAuthZService {
	return s.Service
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) GetTokenKey() *string {
	return s.TokenKey
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) GetWithRematchRoute() *bool {
	return s.WithRematchRoute
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) GetWithRequestBody() *bool {
	return s.WithRequestBody
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) SetAllowRequestHeaders(v []*string) *GetGatewayAuthDetailResponseBodyDataExternalAuthZ {
	s.AllowRequestHeaders = v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) SetAllowUpstreamHeaders(v []*string) *GetGatewayAuthDetailResponseBodyDataExternalAuthZ {
	s.AllowUpstreamHeaders = v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) SetBodyMaxBytes(v int32) *GetGatewayAuthDetailResponseBodyDataExternalAuthZ {
	s.BodyMaxBytes = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) SetIsRestrict(v bool) *GetGatewayAuthDetailResponseBodyDataExternalAuthZ {
	s.IsRestrict = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) SetPrefixPath(v string) *GetGatewayAuthDetailResponseBodyDataExternalAuthZ {
	s.PrefixPath = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) SetService(v *GetGatewayAuthDetailResponseBodyDataExternalAuthZService) *GetGatewayAuthDetailResponseBodyDataExternalAuthZ {
	s.Service = v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) SetServiceId(v int64) *GetGatewayAuthDetailResponseBodyDataExternalAuthZ {
	s.ServiceId = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) SetTimeout(v int32) *GetGatewayAuthDetailResponseBodyDataExternalAuthZ {
	s.Timeout = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) SetTokenKey(v string) *GetGatewayAuthDetailResponseBodyDataExternalAuthZ {
	s.TokenKey = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) SetWithRematchRoute(v bool) *GetGatewayAuthDetailResponseBodyDataExternalAuthZ {
	s.WithRematchRoute = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) SetWithRequestBody(v bool) *GetGatewayAuthDetailResponseBodyDataExternalAuthZ {
	s.WithRequestBody = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZ) Validate() error {
	if s.Service != nil {
		if err := s.Service.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayAuthDetailResponseBodyDataExternalAuthZService struct {
	// The group to which the authentication service belongs.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The authentication service name.
	//
	// example:
	//
	// httpbin-auth-service
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace of the authentication service.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source type of the authentication service.
	//
	// example:
	//
	// K8s
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s GetGatewayAuthDetailResponseBodyDataExternalAuthZService) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayAuthDetailResponseBodyDataExternalAuthZService) GoString() string {
	return s.String()
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZService) GetGroupName() *string {
	return s.GroupName
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZService) GetName() *string {
	return s.Name
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZService) GetNamespace() *string {
	return s.Namespace
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZService) GetSourceType() *string {
	return s.SourceType
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZService) SetGroupName(v string) *GetGatewayAuthDetailResponseBodyDataExternalAuthZService {
	s.GroupName = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZService) SetName(v string) *GetGatewayAuthDetailResponseBodyDataExternalAuthZService {
	s.Name = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZService) SetNamespace(v string) *GetGatewayAuthDetailResponseBodyDataExternalAuthZService {
	s.Namespace = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZService) SetSourceType(v string) *GetGatewayAuthDetailResponseBodyDataExternalAuthZService {
	s.SourceType = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataExternalAuthZService) Validate() error {
	return dara.Validate(s)
}

type GetGatewayAuthDetailResponseBodyDataResourceList struct {
	// The ID of the authentication to which the rule belongs.
	//
	// example:
	//
	// 2274
	AuthId *int64 `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// The request headers of the rule. This parameter is currently only used for custom authentication.
	AuthResourceHeaderList []*GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList `json:"AuthResourceHeaderList,omitempty" xml:"AuthResourceHeaderList,omitempty" type:"Repeated"`
	// The ID of the domain name in the rule.
	//
	// example:
	//
	// 1765
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The domain name in the rule.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the gateway to which the rule belongs.
	//
	// example:
	//
	// 2274
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway to which the rule belongs.
	//
	// example:
	//
	// gw-6f0dbd108a0249d2b675b3ef50b*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The GMT time when the authorization rule was created.
	//
	// example:
	//
	// 2024-02-19T03:32:38.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The GMT time when the authorization rule was last modified.
	//
	// example:
	//
	// 2024-02-19T03:32:38.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the authorization rule.
	//
	// example:
	//
	// 1303
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether case sensitivity is enabled for the path. Valid values:
	//
	// - true: Case sensitivity is enabled.
	//
	// - false: Case sensitivity is disabled.
	//
	// example:
	//
	// true
	IgnoreCase *bool `json:"IgnoreCase,omitempty" xml:"IgnoreCase,omitempty"`
	// The authorization blacklist and whitelist mode. Valid values:
	//
	// - true: Whitelist mode. Hosts and paths in the list can be accessed without verification, while all others require verification.
	//
	// - false: Blacklist mode. Hosts and paths in the list require verification, while all others can be accessed directly.
	//
	// example:
	//
	// true
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// The matching type for the request path. Valid values:
	//
	// - EQUAL: Exact match.
	//
	// - PRE: Prefix match.
	//
	// - ERGULAR: Regular expression match.
	//
	// example:
	//
	// EQUAL
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The matching path of the rule.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s GetGatewayAuthDetailResponseBodyDataResourceList) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayAuthDetailResponseBodyDataResourceList) GoString() string {
	return s.String()
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) GetAuthId() *int64 {
	return s.AuthId
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) GetAuthResourceHeaderList() []*GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList {
	return s.AuthResourceHeaderList
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) GetDomainId() *int64 {
	return s.DomainId
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) GetDomainName() *string {
	return s.DomainName
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) GetId() *int64 {
	return s.Id
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) GetIgnoreCase() *bool {
	return s.IgnoreCase
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) GetIsWhite() *bool {
	return s.IsWhite
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) GetMatchType() *string {
	return s.MatchType
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) GetPath() *string {
	return s.Path
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) SetAuthId(v int64) *GetGatewayAuthDetailResponseBodyDataResourceList {
	s.AuthId = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) SetAuthResourceHeaderList(v []*GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList) *GetGatewayAuthDetailResponseBodyDataResourceList {
	s.AuthResourceHeaderList = v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) SetDomainId(v int64) *GetGatewayAuthDetailResponseBodyDataResourceList {
	s.DomainId = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) SetDomainName(v string) *GetGatewayAuthDetailResponseBodyDataResourceList {
	s.DomainName = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) SetGatewayId(v int64) *GetGatewayAuthDetailResponseBodyDataResourceList {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) SetGatewayUniqueId(v string) *GetGatewayAuthDetailResponseBodyDataResourceList {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) SetGmtCreate(v string) *GetGatewayAuthDetailResponseBodyDataResourceList {
	s.GmtCreate = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) SetGmtModified(v string) *GetGatewayAuthDetailResponseBodyDataResourceList {
	s.GmtModified = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) SetId(v int64) *GetGatewayAuthDetailResponseBodyDataResourceList {
	s.Id = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) SetIgnoreCase(v bool) *GetGatewayAuthDetailResponseBodyDataResourceList {
	s.IgnoreCase = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) SetIsWhite(v bool) *GetGatewayAuthDetailResponseBodyDataResourceList {
	s.IsWhite = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) SetMatchType(v string) *GetGatewayAuthDetailResponseBodyDataResourceList {
	s.MatchType = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) SetPath(v string) *GetGatewayAuthDetailResponseBodyDataResourceList {
	s.Path = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceList) Validate() error {
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

type GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList struct {
	// The name of the request header.
	//
	// example:
	//
	// x-req
	HeaderKey *string `json:"HeaderKey,omitempty" xml:"HeaderKey,omitempty"`
	// The matching condition. Valid values:
	//
	// - EQUAL: Equal to.
	//
	// - NOT_EQUAL: Not equal to.
	//
	// - EXIST: Exists.
	//
	// - NOT_EXIST: Does not exist.
	//
	// - INCLUDE: Includes.
	//
	// - EXCLUDE: Excludes.
	//
	// - PREFIX: Prefix.
	//
	// - SUFFIX: Suffix.
	//
	// - REGREX: Regular expression.
	//
	// example:
	//
	// EQUAL
	HeaderMethod *string `json:"HeaderMethod,omitempty" xml:"HeaderMethod,omitempty"`
	// The value of the request header.
	//
	// example:
	//
	// 123
	HeaderValue *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
}

func (s GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList) GoString() string {
	return s.String()
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList) GetHeaderKey() *string {
	return s.HeaderKey
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList) GetHeaderMethod() *string {
	return s.HeaderMethod
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList) GetHeaderValue() *string {
	return s.HeaderValue
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList) SetHeaderKey(v string) *GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList {
	s.HeaderKey = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList) SetHeaderMethod(v string) *GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList {
	s.HeaderMethod = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList) SetHeaderValue(v string) *GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList {
	s.HeaderValue = &v
	return s
}

func (s *GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList) Validate() error {
	return dara.Validate(s)
}
