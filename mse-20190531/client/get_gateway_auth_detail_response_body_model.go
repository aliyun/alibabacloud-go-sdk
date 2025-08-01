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
	// example:
	//
	// 200
	Code *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetGatewayAuthDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9C96CDF8-9E6C-XXXX-XXXX-8F87A10117E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
}

type GetGatewayAuthDetailResponseBodyData struct {
	AuthResourceConfig *string `json:"AuthResourceConfig,omitempty" xml:"AuthResourceConfig,omitempty"`
	// example:
	//
	// 0
	AuthResourceMode *int32 `json:"AuthResourceMode,omitempty" xml:"AuthResourceMode,omitempty"`
	// example:
	//
	// example-app
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// xxxxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// example:
	//
	// hello.com
	CookieDomain  *string                                            `json:"CookieDomain,omitempty" xml:"CookieDomain,omitempty"`
	ExternalAuthZ *GetGatewayAuthDetailResponseBodyDataExternalAuthZ `json:"ExternalAuthZ,omitempty" xml:"ExternalAuthZ,omitempty" type:"Struct"`
	// example:
	//
	// 2274
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// gw-6f0dbd108a0249d2b675b3ef50b*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 2024-02-19T02:41:03.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-02-19T02:41:03.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1100
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// example:
	//
	// https://example.com/auth
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// example:
	//
	// {\\n  \\"keys\\":[\\n    {\\n      \\"kty\\": \\"RSA\\",\\n      \\"e\\": \\"AQAB\\",\\n      \\"use\\": \\"sig\\",\\n      \\"kid\\": \\"1rGufmH1YN8rqM9ZOLgo7eEST3AnL89Y-m-XGFioLoA\\",\\n      \\"alg\\": \\"RS256\\",\\n      \\"n\\": \\"rM2GIc0YTMqwNCwXnjKbW5QndkCEZgyLu3uQUnyZF7HvMTekiTvQg_39mg3dV1eaYYkYfZBogyroJBqAQXhk6VVCxlBjFVp2xstJPVWngMOOlcafwN_BKdN-EQ06O_Uu__e7gNKI3DunkNk0cNaFETE7d4meRYyTlgEzYgsrW05_ufR0BKoddL3E5JsCpUxRjH9ICbodBx0U74W6Dcci-R2EA1DBrEcboE6n90uoJs6UJNriAK_71nAsYonihU5aQFFnyPTkJHfRwHK6JlME6rn-b-rpLSpdyc6U1nOFZP2DEpz8U5FrYoLYSZIU-MQGxDhCnGc_rxl2IyP9B2qcCQ\\"\\n    }\\n  ]\\n}\\n
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
	// https://yourdomain/path
	RedirectUrl  *string                                             `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	ResourceList []*GetGatewayAuthDetailResponseBodyDataResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// example:
	//
	// ["openid","email"]
	ScopesList *string `json:"ScopesList,omitempty" xml:"ScopesList,omitempty"`
	// example:
	//
	// false
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// https://example.com/auth
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
	return dara.Validate(s)
}

type GetGatewayAuthDetailResponseBodyDataExternalAuthZ struct {
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
	PrefixPath *string                                                   `json:"PrefixPath,omitempty" xml:"PrefixPath,omitempty"`
	Service    *GetGatewayAuthDetailResponseBodyDataExternalAuthZService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	// example:
	//
	// 15300
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 10
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
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
	return dara.Validate(s)
}

type GetGatewayAuthDetailResponseBodyDataExternalAuthZService struct {
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// httpbin-auth-service
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// K8S
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
	// example:
	//
	// 2274
	AuthId                 *int64                                                                    `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	AuthResourceHeaderList []*GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList `json:"AuthResourceHeaderList,omitempty" xml:"AuthResourceHeaderList,omitempty" type:"Repeated"`
	// example:
	//
	// 1765
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2274
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// gw-6f0dbd108a0249d2b675b3ef50b*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 2024-02-19T03:32:38.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-02-19T03:32:38.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1303
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IgnoreCase *bool `json:"IgnoreCase,omitempty" xml:"IgnoreCase,omitempty"`
	// example:
	//
	// true
	IsWhite *bool `json:"IsWhite,omitempty" xml:"IsWhite,omitempty"`
	// example:
	//
	// EQUAL
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
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
	return dara.Validate(s)
}

type GetGatewayAuthDetailResponseBodyDataResourceListAuthResourceHeaderList struct {
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
