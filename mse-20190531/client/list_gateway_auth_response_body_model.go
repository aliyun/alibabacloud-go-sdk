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
	// example:
	//
	// 200
	Code *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListGatewayAuthResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EC1EED4A-B147-597B-B949-FD3131AB4298
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   []*ListGatewayAuthResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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
	// example:
	//
	// 0
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
	CookieDomain  *string                                             `json:"CookieDomain,omitempty" xml:"CookieDomain,omitempty"`
	ExternalAuthZ *ListGatewayAuthResponseBodyDataResultExternalAuthZ `json:"ExternalAuthZ,omitempty" xml:"ExternalAuthZ,omitempty" type:"Struct"`
	// example:
	//
	// 1
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// gw-e2d226bba4b2445c9e29fa7f8216****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 2021-09-13 19:24:23
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2021-09-13 19:24:23
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 549
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
	// {\\"keys\\":[{\\"kty\\":\\"oct\\",\\"k\\":\\"9V9lpiuAQsME1efQChI0kEQz6fdVlJbDRFFa1lvEB_A\\",\\"alg\\":\\"HS256\\"}]}
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
	// http://test.com/oauth2/callback
	RedirectUrl *string `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	// example:
	//
	// ["openid"]
	ScopesList *string `json:"ScopesList,omitempty" xml:"ScopesList,omitempty"`
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
	AllowRequestHeaders  []*string `json:"AllowRequestHeaders,omitempty" xml:"AllowRequestHeaders,omitempty" type:"Repeated"`
	AllowUpstreamHeaders []*string `json:"AllowUpstreamHeaders,omitempty" xml:"AllowUpstreamHeaders,omitempty" type:"Repeated"`
	// example:
	//
	// 1024
	BodyMaxBytes *int32 `json:"BodyMaxBytes,omitempty" xml:"BodyMaxBytes,omitempty"`
	// example:
	//
	// true
	IsRestrict *bool `json:"IsRestrict,omitempty" xml:"IsRestrict,omitempty"`
	// example:
	//
	// /auth
	PrefixPath *string                                                    `json:"PrefixPath,omitempty" xml:"PrefixPath,omitempty"`
	Service    *ListGatewayAuthResponseBodyDataResultExternalAuthZService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	// example:
	//
	// 45186
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 30
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
	// example:
	//
	// com.mse.console.test.3YXO
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// updatetime
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
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
