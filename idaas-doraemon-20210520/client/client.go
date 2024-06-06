// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateAuthenticatorRegistrationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11122223333
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// webauthn
	AuthenticatorType *string `json:"AuthenticatorType,omitempty" xml:"AuthenticatorType,omitempty"`
	// example:
	//
	// 由具体认证方式决定
	ClientExtendParamsJson *string `json:"ClientExtendParamsJson,omitempty" xml:"ClientExtendParamsJson,omitempty"`
	// example:
	//
	// 由具体认证方式决定
	ClientExtendParamsJsonSign *string `json:"ClientExtendParamsJsonSign,omitempty" xml:"ClientExtendParamsJsonSign,omitempty"`
	// example:
	//
	// 注册上下文
	RegistrationContext *string `json:"RegistrationContext,omitempty" xml:"RegistrationContext,omitempty"`
	// example:
	//
	// 由具体认证方式决定
	ServerExtendParamsJson *string `json:"ServerExtendParamsJson,omitempty" xml:"ServerExtendParamsJson,omitempty"`
	// example:
	//
	// 张三
	UserDisplayName *string `json:"UserDisplayName,omitempty" xml:"UserDisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user-11122223333
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// ZhangSan
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateAuthenticatorRegistrationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthenticatorRegistrationRequest) GoString() string {
	return s.String()
}

func (s *CreateAuthenticatorRegistrationRequest) SetApplicationExternalId(v string) *CreateAuthenticatorRegistrationRequest {
	s.ApplicationExternalId = &v
	return s
}

func (s *CreateAuthenticatorRegistrationRequest) SetAuthenticatorType(v string) *CreateAuthenticatorRegistrationRequest {
	s.AuthenticatorType = &v
	return s
}

func (s *CreateAuthenticatorRegistrationRequest) SetClientExtendParamsJson(v string) *CreateAuthenticatorRegistrationRequest {
	s.ClientExtendParamsJson = &v
	return s
}

func (s *CreateAuthenticatorRegistrationRequest) SetClientExtendParamsJsonSign(v string) *CreateAuthenticatorRegistrationRequest {
	s.ClientExtendParamsJsonSign = &v
	return s
}

func (s *CreateAuthenticatorRegistrationRequest) SetRegistrationContext(v string) *CreateAuthenticatorRegistrationRequest {
	s.RegistrationContext = &v
	return s
}

func (s *CreateAuthenticatorRegistrationRequest) SetServerExtendParamsJson(v string) *CreateAuthenticatorRegistrationRequest {
	s.ServerExtendParamsJson = &v
	return s
}

func (s *CreateAuthenticatorRegistrationRequest) SetUserDisplayName(v string) *CreateAuthenticatorRegistrationRequest {
	s.UserDisplayName = &v
	return s
}

func (s *CreateAuthenticatorRegistrationRequest) SetUserId(v string) *CreateAuthenticatorRegistrationRequest {
	s.UserId = &v
	return s
}

func (s *CreateAuthenticatorRegistrationRequest) SetUserName(v string) *CreateAuthenticatorRegistrationRequest {
	s.UserName = &v
	return s
}

type CreateAuthenticatorRegistrationResponseBody struct {
	ChallengeBase64 *string `json:"ChallengeBase64,omitempty" xml:"ChallengeBase64,omitempty"`
	Options         *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// DA0A9BF3-029D-5968-B876-0DF49777B22D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAuthenticatorRegistrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthenticatorRegistrationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAuthenticatorRegistrationResponseBody) SetChallengeBase64(v string) *CreateAuthenticatorRegistrationResponseBody {
	s.ChallengeBase64 = &v
	return s
}

func (s *CreateAuthenticatorRegistrationResponseBody) SetOptions(v string) *CreateAuthenticatorRegistrationResponseBody {
	s.Options = &v
	return s
}

func (s *CreateAuthenticatorRegistrationResponseBody) SetRequestId(v string) *CreateAuthenticatorRegistrationResponseBody {
	s.RequestId = &v
	return s
}

type CreateAuthenticatorRegistrationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAuthenticatorRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAuthenticatorRegistrationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthenticatorRegistrationResponse) GoString() string {
	return s.String()
}

func (s *CreateAuthenticatorRegistrationResponse) SetHeaders(v map[string]*string) *CreateAuthenticatorRegistrationResponse {
	s.Headers = v
	return s
}

func (s *CreateAuthenticatorRegistrationResponse) SetStatusCode(v int32) *CreateAuthenticatorRegistrationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAuthenticatorRegistrationResponse) SetBody(v *CreateAuthenticatorRegistrationResponseBody) *CreateAuthenticatorRegistrationResponse {
	s.Body = v
	return s
}

type CreateUserAuthenticateOptionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11122223333
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// webauthn
	AuthenticatorType *string `json:"AuthenticatorType,omitempty" xml:"AuthenticatorType,omitempty"`
	// example:
	//
	// 1111222
	BindHashBase64 *string `json:"BindHashBase64,omitempty" xml:"BindHashBase64,omitempty"`
	// example:
	//
	// 由具体认证方式决定
	ClientExtendParamsJson *string `json:"ClientExtendParamsJson,omitempty" xml:"ClientExtendParamsJson,omitempty"`
	// example:
	//
	// 由具体认证方式决定
	ClientExtendParamsJsonSign *string `json:"ClientExtendParamsJsonSign,omitempty" xml:"ClientExtendParamsJsonSign,omitempty"`
	// example:
	//
	// 由具体认证方式决定
	ServerExtendParamsJson *string `json:"ServerExtendParamsJson,omitempty" xml:"ServerExtendParamsJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user-11122223333
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateUserAuthenticateOptionsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserAuthenticateOptionsRequest) GoString() string {
	return s.String()
}

func (s *CreateUserAuthenticateOptionsRequest) SetApplicationExternalId(v string) *CreateUserAuthenticateOptionsRequest {
	s.ApplicationExternalId = &v
	return s
}

func (s *CreateUserAuthenticateOptionsRequest) SetAuthenticatorType(v string) *CreateUserAuthenticateOptionsRequest {
	s.AuthenticatorType = &v
	return s
}

func (s *CreateUserAuthenticateOptionsRequest) SetBindHashBase64(v string) *CreateUserAuthenticateOptionsRequest {
	s.BindHashBase64 = &v
	return s
}

func (s *CreateUserAuthenticateOptionsRequest) SetClientExtendParamsJson(v string) *CreateUserAuthenticateOptionsRequest {
	s.ClientExtendParamsJson = &v
	return s
}

func (s *CreateUserAuthenticateOptionsRequest) SetClientExtendParamsJsonSign(v string) *CreateUserAuthenticateOptionsRequest {
	s.ClientExtendParamsJsonSign = &v
	return s
}

func (s *CreateUserAuthenticateOptionsRequest) SetServerExtendParamsJson(v string) *CreateUserAuthenticateOptionsRequest {
	s.ServerExtendParamsJson = &v
	return s
}

func (s *CreateUserAuthenticateOptionsRequest) SetUserId(v string) *CreateUserAuthenticateOptionsRequest {
	s.UserId = &v
	return s
}

type CreateUserAuthenticateOptionsResponseBody struct {
	// example:
	//
	// 2343246776
	ChallengeBase64 *string `json:"ChallengeBase64,omitempty" xml:"ChallengeBase64,omitempty"`
	// example:
	//
	// {\\"subscribeEvent\\":false,\\"installAgent\\":false}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// D0F570B0-89CE-07F2-B41E-F97E089E13C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUserAuthenticateOptionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserAuthenticateOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserAuthenticateOptionsResponseBody) SetChallengeBase64(v string) *CreateUserAuthenticateOptionsResponseBody {
	s.ChallengeBase64 = &v
	return s
}

func (s *CreateUserAuthenticateOptionsResponseBody) SetOptions(v string) *CreateUserAuthenticateOptionsResponseBody {
	s.Options = &v
	return s
}

func (s *CreateUserAuthenticateOptionsResponseBody) SetRequestId(v string) *CreateUserAuthenticateOptionsResponseBody {
	s.RequestId = &v
	return s
}

type CreateUserAuthenticateOptionsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserAuthenticateOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserAuthenticateOptionsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserAuthenticateOptionsResponse) GoString() string {
	return s.String()
}

func (s *CreateUserAuthenticateOptionsResponse) SetHeaders(v map[string]*string) *CreateUserAuthenticateOptionsResponse {
	s.Headers = v
	return s
}

func (s *CreateUserAuthenticateOptionsResponse) SetStatusCode(v int32) *CreateUserAuthenticateOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserAuthenticateOptionsResponse) SetBody(v *CreateUserAuthenticateOptionsResponseBody) *CreateUserAuthenticateOptionsResponse {
	s.Body = v
	return s
}

type DeregisterAuthenticatorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11122223333
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a00712a4e56196755059356c290793031xHdUrCPXUS
	AuthenticatorUuid *string `json:"AuthenticatorUuid,omitempty" xml:"AuthenticatorUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user-11122223333
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeregisterAuthenticatorRequest) String() string {
	return tea.Prettify(s)
}

func (s DeregisterAuthenticatorRequest) GoString() string {
	return s.String()
}

func (s *DeregisterAuthenticatorRequest) SetApplicationExternalId(v string) *DeregisterAuthenticatorRequest {
	s.ApplicationExternalId = &v
	return s
}

func (s *DeregisterAuthenticatorRequest) SetAuthenticatorUuid(v string) *DeregisterAuthenticatorRequest {
	s.AuthenticatorUuid = &v
	return s
}

func (s *DeregisterAuthenticatorRequest) SetUserId(v string) *DeregisterAuthenticatorRequest {
	s.UserId = &v
	return s
}

type DeregisterAuthenticatorResponseBody struct {
	// example:
	//
	// D0F570B0-89CE-07F2-B41E-F97E089E13C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeregisterAuthenticatorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeregisterAuthenticatorResponseBody) GoString() string {
	return s.String()
}

func (s *DeregisterAuthenticatorResponseBody) SetRequestId(v string) *DeregisterAuthenticatorResponseBody {
	s.RequestId = &v
	return s
}

type DeregisterAuthenticatorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeregisterAuthenticatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeregisterAuthenticatorResponse) String() string {
	return tea.Prettify(s)
}

func (s DeregisterAuthenticatorResponse) GoString() string {
	return s.String()
}

func (s *DeregisterAuthenticatorResponse) SetHeaders(v map[string]*string) *DeregisterAuthenticatorResponse {
	s.Headers = v
	return s
}

func (s *DeregisterAuthenticatorResponse) SetStatusCode(v int32) *DeregisterAuthenticatorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeregisterAuthenticatorResponse) SetBody(v *DeregisterAuthenticatorResponseBody) *DeregisterAuthenticatorResponse {
	s.Body = v
	return s
}

type FetchAccessTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// secAuthApp
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eyJhdHRhY2htZW50IjoicGxhdGZvcm0iLCJ1c2VyVmVyaWZpY2F0aW9uIjoicmVxdWlyZWQifQ==
	MobileExtendParamsJson *string `json:"MobileExtendParamsJson,omitempty" xml:"MobileExtendParamsJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a3984bc27776bb6f65f82276f75d6a12d56e1b41e98ed6975d852b14309b8a20
	MobileExtendParamsJsonSign *string `json:"MobileExtendParamsJsonSign,omitempty" xml:"MobileExtendParamsJsonSign,omitempty"`
	// example:
	//
	// eyJhdHRhY2htZW50IjoicGxhdGZvcm0iLCJ1c2VyVmVyaWZpY2F0aW9uIjoicmVxdWlyZWQifQ==
	ServerExtendParamsJson *string `json:"ServerExtendParamsJson,omitempty" xml:"ServerExtendParamsJson,omitempty"`
	// example:
	//
	// 6ba923abf6749536b7fd31c6a31360fboOFZQ8WbzmG
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 124.127.214.94
	XClientIp *string `json:"XClientIp,omitempty" xml:"XClientIp,omitempty"`
}

func (s FetchAccessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *FetchAccessTokenRequest) SetApplicationExternalId(v string) *FetchAccessTokenRequest {
	s.ApplicationExternalId = &v
	return s
}

func (s *FetchAccessTokenRequest) SetMobileExtendParamsJson(v string) *FetchAccessTokenRequest {
	s.MobileExtendParamsJson = &v
	return s
}

func (s *FetchAccessTokenRequest) SetMobileExtendParamsJsonSign(v string) *FetchAccessTokenRequest {
	s.MobileExtendParamsJsonSign = &v
	return s
}

func (s *FetchAccessTokenRequest) SetServerExtendParamsJson(v string) *FetchAccessTokenRequest {
	s.ServerExtendParamsJson = &v
	return s
}

func (s *FetchAccessTokenRequest) SetUserId(v string) *FetchAccessTokenRequest {
	s.UserId = &v
	return s
}

func (s *FetchAccessTokenRequest) SetXClientIp(v string) *FetchAccessTokenRequest {
	s.XClientIp = &v
	return s
}

type FetchAccessTokenResponseBody struct {
	// example:
	//
	// Operation.Success
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *FetchAccessTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BD908DED-849D-5FDB-9C9B-52B1D3F9DC41
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FetchAccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FetchAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *FetchAccessTokenResponseBody) SetCode(v string) *FetchAccessTokenResponseBody {
	s.Code = &v
	return s
}

func (s *FetchAccessTokenResponseBody) SetData(v *FetchAccessTokenResponseBodyData) *FetchAccessTokenResponseBody {
	s.Data = v
	return s
}

func (s *FetchAccessTokenResponseBody) SetMessage(v string) *FetchAccessTokenResponseBody {
	s.Message = &v
	return s
}

func (s *FetchAccessTokenResponseBody) SetRequestId(v string) *FetchAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchAccessTokenResponseBody) SetSuccess(v bool) *FetchAccessTokenResponseBody {
	s.Success = &v
	return s
}

type FetchAccessTokenResponseBodyData struct {
	// example:
	//
	// eyciOi...jV_PTsxgaNA
	AccessToken *string `json:"Access_token,omitempty" xml:"Access_token,omitempty"`
	// example:
	//
	// 7200
	ExpiresIn *string `json:"Expires_in,omitempty" xml:"Expires_in,omitempty"`
	// example:
	//
	// NiwiZW...wcmlzZUlk
	IdToken *string `json:"Id_token,omitempty" xml:"Id_token,omitempty"`
	// example:
	//
	// 7200
	RefreshToken *string `json:"Refresh_token,omitempty" xml:"Refresh_token,omitempty"`
	// example:
	//
	// openid
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// Bearer
	TokenType *string `json:"Token_type,omitempty" xml:"Token_type,omitempty"`
}

func (s FetchAccessTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FetchAccessTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *FetchAccessTokenResponseBodyData) SetAccessToken(v string) *FetchAccessTokenResponseBodyData {
	s.AccessToken = &v
	return s
}

func (s *FetchAccessTokenResponseBodyData) SetExpiresIn(v string) *FetchAccessTokenResponseBodyData {
	s.ExpiresIn = &v
	return s
}

func (s *FetchAccessTokenResponseBodyData) SetIdToken(v string) *FetchAccessTokenResponseBodyData {
	s.IdToken = &v
	return s
}

func (s *FetchAccessTokenResponseBodyData) SetRefreshToken(v string) *FetchAccessTokenResponseBodyData {
	s.RefreshToken = &v
	return s
}

func (s *FetchAccessTokenResponseBodyData) SetScope(v string) *FetchAccessTokenResponseBodyData {
	s.Scope = &v
	return s
}

func (s *FetchAccessTokenResponseBodyData) SetTokenType(v string) *FetchAccessTokenResponseBodyData {
	s.TokenType = &v
	return s
}

type FetchAccessTokenResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s FetchAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *FetchAccessTokenResponse) SetHeaders(v map[string]*string) *FetchAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *FetchAccessTokenResponse) SetStatusCode(v int32) *FetchAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchAccessTokenResponse) SetBody(v *FetchAccessTokenResponseBody) *FetchAccessTokenResponse {
	s.Body = v
	return s
}

type GetAuthenticatorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11122223333
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1331311
	AuthenticatorUuid *string `json:"AuthenticatorUuid,omitempty" xml:"AuthenticatorUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAuthenticatorRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthenticatorRequest) GoString() string {
	return s.String()
}

func (s *GetAuthenticatorRequest) SetApplicationExternalId(v string) *GetAuthenticatorRequest {
	s.ApplicationExternalId = &v
	return s
}

func (s *GetAuthenticatorRequest) SetAuthenticatorUuid(v string) *GetAuthenticatorRequest {
	s.AuthenticatorUuid = &v
	return s
}

func (s *GetAuthenticatorRequest) SetUserId(v string) *GetAuthenticatorRequest {
	s.UserId = &v
	return s
}

type GetAuthenticatorResponseBody struct {
	Authenticator *GetAuthenticatorResponseBodyAuthenticator `json:"Authenticator,omitempty" xml:"Authenticator,omitempty" type:"Struct"`
	// example:
	//
	// 123131312313
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuthenticatorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthenticatorResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthenticatorResponseBody) SetAuthenticator(v *GetAuthenticatorResponseBodyAuthenticator) *GetAuthenticatorResponseBody {
	s.Authenticator = v
	return s
}

func (s *GetAuthenticatorResponseBody) SetRequestId(v string) *GetAuthenticatorResponseBody {
	s.RequestId = &v
	return s
}

type GetAuthenticatorResponseBodyAuthenticator struct {
	// example:
	//
	// ceshi
	AuthenticatorName *string `json:"AuthenticatorName,omitempty" xml:"AuthenticatorName,omitempty"`
	AuthenticatorUuid *string `json:"AuthenticatorUuid,omitempty" xml:"AuthenticatorUuid,omitempty"`
	// example:
	//
	// 1313
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// example:
	//
	// 123123
	CustomAuthenticator *string `json:"CustomAuthenticator,omitempty" xml:"CustomAuthenticator,omitempty"`
	// example:
	//
	// 127.0.0.1
	LastVerifySourceIp *string `json:"LastVerifySourceIp,omitempty" xml:"LastVerifySourceIp,omitempty"`
	// example:
	//
	// 123123
	LastVerifyTime *int64 `json:"LastVerifyTime,omitempty" xml:"LastVerifyTime,omitempty"`
	// example:
	//
	// 123123
	LastVerifyUserAgent *string `json:"LastVerifyUserAgent,omitempty" xml:"LastVerifyUserAgent,omitempty"`
	// example:
	//
	// 127.0.0.1
	RegisterSourceIp *string `json:"RegisterSourceIp,omitempty" xml:"RegisterSourceIp,omitempty"`
	// example:
	//
	// 1231313123
	RegisterTime *int64 `json:"RegisterTime,omitempty" xml:"RegisterTime,omitempty"`
	// example:
	//
	// TOTP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAuthenticatorResponseBodyAuthenticator) String() string {
	return tea.Prettify(s)
}

func (s GetAuthenticatorResponseBodyAuthenticator) GoString() string {
	return s.String()
}

func (s *GetAuthenticatorResponseBodyAuthenticator) SetAuthenticatorName(v string) *GetAuthenticatorResponseBodyAuthenticator {
	s.AuthenticatorName = &v
	return s
}

func (s *GetAuthenticatorResponseBodyAuthenticator) SetAuthenticatorUuid(v string) *GetAuthenticatorResponseBodyAuthenticator {
	s.AuthenticatorUuid = &v
	return s
}

func (s *GetAuthenticatorResponseBodyAuthenticator) SetCredentialId(v string) *GetAuthenticatorResponseBodyAuthenticator {
	s.CredentialId = &v
	return s
}

func (s *GetAuthenticatorResponseBodyAuthenticator) SetCustomAuthenticator(v string) *GetAuthenticatorResponseBodyAuthenticator {
	s.CustomAuthenticator = &v
	return s
}

func (s *GetAuthenticatorResponseBodyAuthenticator) SetLastVerifySourceIp(v string) *GetAuthenticatorResponseBodyAuthenticator {
	s.LastVerifySourceIp = &v
	return s
}

func (s *GetAuthenticatorResponseBodyAuthenticator) SetLastVerifyTime(v int64) *GetAuthenticatorResponseBodyAuthenticator {
	s.LastVerifyTime = &v
	return s
}

func (s *GetAuthenticatorResponseBodyAuthenticator) SetLastVerifyUserAgent(v string) *GetAuthenticatorResponseBodyAuthenticator {
	s.LastVerifyUserAgent = &v
	return s
}

func (s *GetAuthenticatorResponseBodyAuthenticator) SetRegisterSourceIp(v string) *GetAuthenticatorResponseBodyAuthenticator {
	s.RegisterSourceIp = &v
	return s
}

func (s *GetAuthenticatorResponseBodyAuthenticator) SetRegisterTime(v int64) *GetAuthenticatorResponseBodyAuthenticator {
	s.RegisterTime = &v
	return s
}

func (s *GetAuthenticatorResponseBodyAuthenticator) SetType(v string) *GetAuthenticatorResponseBodyAuthenticator {
	s.Type = &v
	return s
}

type GetAuthenticatorResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthenticatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthenticatorResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthenticatorResponse) GoString() string {
	return s.String()
}

func (s *GetAuthenticatorResponse) SetHeaders(v map[string]*string) *GetAuthenticatorResponse {
	s.Headers = v
	return s
}

func (s *GetAuthenticatorResponse) SetStatusCode(v int32) *GetAuthenticatorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthenticatorResponse) SetBody(v *GetAuthenticatorResponseBody) *GetAuthenticatorResponse {
	s.Body = v
	return s
}

type ListAuthenticationLogsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// webauth_demo
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// example:
	//
	// TOTP
	AuthenticatorType *string `json:"AuthenticatorType,omitempty" xml:"AuthenticatorType,omitempty"`
	// example:
	//
	// 0fcd5fce3fdc42052633a65eb242f10aZ8nzoQMpV3m
	AuthenticatorUuid *string `json:"AuthenticatorUuid,omitempty" xml:"AuthenticatorUuid,omitempty"`
	// example:
	//
	// scimcred-003o4e2wqxlaxjmjyz96
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// example:
	//
	// 1640144294
	FromTime *int64 `json:"FromTime,omitempty" xml:"FromTime,omitempty"`
	// example:
	//
	// login
	LogTag *string `json:"LogTag,omitempty" xml:"LogTag,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1640144694
	ToTime *int64 `json:"ToTime,omitempty" xml:"ToTime,omitempty"`
	// example:
	//
	// 10e8c224
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAuthenticationLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAuthenticationLogsRequest) GoString() string {
	return s.String()
}

func (s *ListAuthenticationLogsRequest) SetApplicationExternalId(v string) *ListAuthenticationLogsRequest {
	s.ApplicationExternalId = &v
	return s
}

func (s *ListAuthenticationLogsRequest) SetAuthenticatorType(v string) *ListAuthenticationLogsRequest {
	s.AuthenticatorType = &v
	return s
}

func (s *ListAuthenticationLogsRequest) SetAuthenticatorUuid(v string) *ListAuthenticationLogsRequest {
	s.AuthenticatorUuid = &v
	return s
}

func (s *ListAuthenticationLogsRequest) SetCredentialId(v string) *ListAuthenticationLogsRequest {
	s.CredentialId = &v
	return s
}

func (s *ListAuthenticationLogsRequest) SetFromTime(v int64) *ListAuthenticationLogsRequest {
	s.FromTime = &v
	return s
}

func (s *ListAuthenticationLogsRequest) SetLogTag(v string) *ListAuthenticationLogsRequest {
	s.LogTag = &v
	return s
}

func (s *ListAuthenticationLogsRequest) SetPageNumber(v int64) *ListAuthenticationLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthenticationLogsRequest) SetPageSize(v int64) *ListAuthenticationLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthenticationLogsRequest) SetToTime(v int64) *ListAuthenticationLogsRequest {
	s.ToTime = &v
	return s
}

func (s *ListAuthenticationLogsRequest) SetUserId(v string) *ListAuthenticationLogsRequest {
	s.UserId = &v
	return s
}

type ListAuthenticationLogsResponseBody struct {
	AuthenticationLogContent []*ListAuthenticationLogsResponseBodyAuthenticationLogContent `json:"AuthenticationLogContent,omitempty" xml:"AuthenticationLogContent,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// F620E712-F325-520D-BE1A-E2421B062F57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAuthenticationLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAuthenticationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthenticationLogsResponseBody) SetAuthenticationLogContent(v []*ListAuthenticationLogsResponseBodyAuthenticationLogContent) *ListAuthenticationLogsResponseBody {
	s.AuthenticationLogContent = v
	return s
}

func (s *ListAuthenticationLogsResponseBody) SetPageNumber(v int64) *ListAuthenticationLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAuthenticationLogsResponseBody) SetPageSize(v int64) *ListAuthenticationLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAuthenticationLogsResponseBody) SetRequestId(v string) *ListAuthenticationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthenticationLogsResponseBody) SetTotalCount(v int64) *ListAuthenticationLogsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAuthenticationLogsResponseBodyAuthenticationLogContent struct {
	// example:
	//
	// 1588301417442737
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// sms_auth
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// example:
	//
	// 0b6ba5da7bd4cef233ae7b48b5e5dcc1yx9aipQp0hD
	ApplicationUuid *string `json:"ApplicationUuid,omitempty" xml:"ApplicationUuid,omitempty"`
	// example:
	//
	// send
	AuthenticationAction *string `json:"AuthenticationAction,omitempty" xml:"AuthenticationAction,omitempty"`
	// example:
	//
	// 1647233491138
	AuthenticationTime *int64 `json:"AuthenticationTime,omitempty" xml:"AuthenticationTime,omitempty"`
	// example:
	//
	// TOTP-DefaultName
	AuthenticatorName *string `json:"AuthenticatorName,omitempty" xml:"AuthenticatorName,omitempty"`
	// example:
	//
	// TOTP
	AuthenticatorType *string `json:"AuthenticatorType,omitempty" xml:"AuthenticatorType,omitempty"`
	// example:
	//
	// c810731ad3af06e45171395abf19cf9cb00E7i97PhE
	AuthenticatorUuid *string `json:"AuthenticatorUuid,omitempty" xml:"AuthenticatorUuid,omitempty"`
	// example:
	//
	// 3519
	ChallengeBase64 *string `json:"ChallengeBase64,omitempty" xml:"ChallengeBase64,omitempty"`
	// example:
	//
	// scimcred-00w9gcvpwueg7geg1a87
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Request was denied due to business check permission
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// khj
	LogParams *string `json:"LogParams,omitempty" xml:"LogParams,omitempty"`
	// example:
	//
	// login
	LogTag *string `json:"LogTag,omitempty" xml:"LogTag,omitempty"`
	// example:
	//
	// {"name":"lili"}
	RawAuthenticationContext *string `json:"RawAuthenticationContext,omitempty" xml:"RawAuthenticationContext,omitempty"`
	// example:
	//
	// 8cefc5fe5d10dffffdcff011a6b34a863hHmgm0SIrh
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// aoxis
	UserAgent *string `json:"UserAgent,omitempty" xml:"UserAgent,omitempty"`
	// example:
	//
	// xhtseqml
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 10.23.11.56
	UserSourceIp *string `json:"UserSourceIp,omitempty" xml:"UserSourceIp,omitempty"`
	// example:
	//
	// true
	VerifyResult *bool `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s ListAuthenticationLogsResponseBodyAuthenticationLogContent) String() string {
	return tea.Prettify(s)
}

func (s ListAuthenticationLogsResponseBodyAuthenticationLogContent) GoString() string {
	return s.String()
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetAliUid(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.AliUid = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetApplicationExternalId(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.ApplicationExternalId = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetApplicationUuid(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.ApplicationUuid = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetAuthenticationAction(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.AuthenticationAction = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetAuthenticationTime(v int64) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.AuthenticationTime = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetAuthenticatorName(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.AuthenticatorName = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetAuthenticatorType(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.AuthenticatorType = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetAuthenticatorUuid(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.AuthenticatorUuid = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetChallengeBase64(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.ChallengeBase64 = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetCredentialId(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.CredentialId = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetErrorCode(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.ErrorCode = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetErrorMessage(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.ErrorMessage = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetLogParams(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.LogParams = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetLogTag(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.LogTag = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetRawAuthenticationContext(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.RawAuthenticationContext = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetTenantId(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.TenantId = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetUserAgent(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.UserAgent = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetUserId(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.UserId = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetUserSourceIp(v string) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.UserSourceIp = &v
	return s
}

func (s *ListAuthenticationLogsResponseBodyAuthenticationLogContent) SetVerifyResult(v bool) *ListAuthenticationLogsResponseBodyAuthenticationLogContent {
	s.VerifyResult = &v
	return s
}

type ListAuthenticationLogsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthenticationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthenticationLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAuthenticationLogsResponse) GoString() string {
	return s.String()
}

func (s *ListAuthenticationLogsResponse) SetHeaders(v map[string]*string) *ListAuthenticationLogsResponse {
	s.Headers = v
	return s
}

func (s *ListAuthenticationLogsResponse) SetStatusCode(v int32) *ListAuthenticationLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthenticationLogsResponse) SetBody(v *ListAuthenticationLogsResponseBody) *ListAuthenticationLogsResponse {
	s.Body = v
	return s
}

type ListAuthenticatorOpsLogsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// webauth_demo
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// example:
	//
	// webauthn
	AuthenticatorType *string `json:"AuthenticatorType,omitempty" xml:"AuthenticatorType,omitempty"`
	// example:
	//
	// c810731ad3af06e45171395abf19cf9cb00E7i97PhE
	AuthenticatorUuid *string `json:"AuthenticatorUuid,omitempty" xml:"AuthenticatorUuid,omitempty"`
	// example:
	//
	// 1647232996712
	FromTime *int64 `json:"FromTime,omitempty" xml:"FromTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1647232996812
	ToTime *int64 `json:"ToTime,omitempty" xml:"ToTime,omitempty"`
	// example:
	//
	// wsnidfr
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAuthenticatorOpsLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAuthenticatorOpsLogsRequest) GoString() string {
	return s.String()
}

func (s *ListAuthenticatorOpsLogsRequest) SetApplicationExternalId(v string) *ListAuthenticatorOpsLogsRequest {
	s.ApplicationExternalId = &v
	return s
}

func (s *ListAuthenticatorOpsLogsRequest) SetAuthenticatorType(v string) *ListAuthenticatorOpsLogsRequest {
	s.AuthenticatorType = &v
	return s
}

func (s *ListAuthenticatorOpsLogsRequest) SetAuthenticatorUuid(v string) *ListAuthenticatorOpsLogsRequest {
	s.AuthenticatorUuid = &v
	return s
}

func (s *ListAuthenticatorOpsLogsRequest) SetFromTime(v int64) *ListAuthenticatorOpsLogsRequest {
	s.FromTime = &v
	return s
}

func (s *ListAuthenticatorOpsLogsRequest) SetPageNumber(v int64) *ListAuthenticatorOpsLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthenticatorOpsLogsRequest) SetPageSize(v int64) *ListAuthenticatorOpsLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthenticatorOpsLogsRequest) SetToTime(v int64) *ListAuthenticatorOpsLogsRequest {
	s.ToTime = &v
	return s
}

func (s *ListAuthenticatorOpsLogsRequest) SetUserId(v string) *ListAuthenticatorOpsLogsRequest {
	s.UserId = &v
	return s
}

type ListAuthenticatorOpsLogsResponseBody struct {
	AuthenticationLogContent []*ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent `json:"AuthenticationLogContent,omitempty" xml:"AuthenticationLogContent,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 39FF2C8B-73E7-5BA5-AC25-20EEAC37BAB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAuthenticatorOpsLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAuthenticatorOpsLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthenticatorOpsLogsResponseBody) SetAuthenticationLogContent(v []*ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) *ListAuthenticatorOpsLogsResponseBody {
	s.AuthenticationLogContent = v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBody) SetPageNumber(v int64) *ListAuthenticatorOpsLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBody) SetPageSize(v int64) *ListAuthenticatorOpsLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBody) SetRequestId(v string) *ListAuthenticatorOpsLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBody) SetTotalCount(v int64) *ListAuthenticatorOpsLogsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent struct {
	// example:
	//
	// 1218758215717204
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// webauth_demo
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// example:
	//
	// 2bba0564edce2ff90c544dc623d1c1f5iPpCGAqTioA
	ApplicationUuid *string `json:"ApplicationUuid,omitempty" xml:"ApplicationUuid,omitempty"`
	// example:
	//
	// TOTP-DefaultName
	AuthenticatorName *string `json:"AuthenticatorName,omitempty" xml:"AuthenticatorName,omitempty"`
	// example:
	//
	// TOTP
	AuthenticatorType *string `json:"AuthenticatorType,omitempty" xml:"AuthenticatorType,omitempty"`
	// example:
	//
	// 0fcd5fce3fdc42052633a65eb242f10aZ8nzoQMpV3m
	AuthenticatorUuid *string `json:"AuthenticatorUuid,omitempty" xml:"AuthenticatorUuid,omitempty"`
	// example:
	//
	// exfa
	ChallengeBase64 *string `json:"ChallengeBase64,omitempty" xml:"ChallengeBase64,omitempty"`
	// example:
	//
	// scimcred-00t3zfhiy2ylzvm2i287
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// example:
	//
	// Failed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// user not exist, alipayId=2088022534411015
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 1
	LogParams *string `json:"LogParams,omitempty" xml:"LogParams,omitempty"`
	// example:
	//
	// ListTagResources
	OperationAction *string `json:"OperationAction,omitempty" xml:"OperationAction,omitempty"`
	// example:
	//
	// True
	OperationResult *bool `json:"OperationResult,omitempty" xml:"OperationResult,omitempty"`
	// example:
	//
	// 1647309893086
	OperationTime *int64  `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
	RawContext    *string `json:"RawContext,omitempty" xml:"RawContext,omitempty"`
	// example:
	//
	// ed0222867189ad2bacfa746843d4b9692mlWC9sU3zl
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 10.13.1.56
	UserAgent *string `json:"UserAgent,omitempty" xml:"UserAgent,omitempty"`
	// example:
	//
	// 3b323
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 10.13.1.56
	UserSourceIp *string `json:"UserSourceIp,omitempty" xml:"UserSourceIp,omitempty"`
}

func (s ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) String() string {
	return tea.Prettify(s)
}

func (s ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) GoString() string {
	return s.String()
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetAliUid(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.AliUid = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetApplicationExternalId(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.ApplicationExternalId = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetApplicationUuid(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.ApplicationUuid = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetAuthenticatorName(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.AuthenticatorName = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetAuthenticatorType(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.AuthenticatorType = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetAuthenticatorUuid(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.AuthenticatorUuid = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetChallengeBase64(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.ChallengeBase64 = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetCredentialId(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.CredentialId = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetErrorCode(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.ErrorCode = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetErrorMessage(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.ErrorMessage = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetLogParams(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.LogParams = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetOperationAction(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.OperationAction = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetOperationResult(v bool) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.OperationResult = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetOperationTime(v int64) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.OperationTime = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetRawContext(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.RawContext = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetTenantId(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.TenantId = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetUserAgent(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.UserAgent = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetUserId(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.UserId = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent) SetUserSourceIp(v string) *ListAuthenticatorOpsLogsResponseBodyAuthenticationLogContent {
	s.UserSourceIp = &v
	return s
}

type ListAuthenticatorOpsLogsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthenticatorOpsLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthenticatorOpsLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAuthenticatorOpsLogsResponse) GoString() string {
	return s.String()
}

func (s *ListAuthenticatorOpsLogsResponse) SetHeaders(v map[string]*string) *ListAuthenticatorOpsLogsResponse {
	s.Headers = v
	return s
}

func (s *ListAuthenticatorOpsLogsResponse) SetStatusCode(v int32) *ListAuthenticatorOpsLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthenticatorOpsLogsResponse) SetBody(v *ListAuthenticatorOpsLogsResponseBody) *ListAuthenticatorOpsLogsResponse {
	s.Body = v
	return s
}

type ListAuthenticatorsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11122223333
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// example:
	//
	// TOTP
	AuthenticatorType *string `json:"AuthenticatorType,omitempty" xml:"AuthenticatorType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 5
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user-11122223333
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAuthenticatorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAuthenticatorsRequest) GoString() string {
	return s.String()
}

func (s *ListAuthenticatorsRequest) SetApplicationExternalId(v string) *ListAuthenticatorsRequest {
	s.ApplicationExternalId = &v
	return s
}

func (s *ListAuthenticatorsRequest) SetAuthenticatorType(v string) *ListAuthenticatorsRequest {
	s.AuthenticatorType = &v
	return s
}

func (s *ListAuthenticatorsRequest) SetPageNumber(v int64) *ListAuthenticatorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthenticatorsRequest) SetPageSize(v int64) *ListAuthenticatorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthenticatorsRequest) SetUserId(v string) *ListAuthenticatorsRequest {
	s.UserId = &v
	return s
}

type ListAuthenticatorsResponseBody struct {
	Authenticator []*ListAuthenticatorsResponseBodyAuthenticator `json:"Authenticator,omitempty" xml:"Authenticator,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 5
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 531FA210-6F3D-5368-9AE2-AD7B9D1D5CAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 13
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAuthenticatorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAuthenticatorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthenticatorsResponseBody) SetAuthenticator(v []*ListAuthenticatorsResponseBodyAuthenticator) *ListAuthenticatorsResponseBody {
	s.Authenticator = v
	return s
}

func (s *ListAuthenticatorsResponseBody) SetPageNumber(v int64) *ListAuthenticatorsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAuthenticatorsResponseBody) SetPageSize(v int64) *ListAuthenticatorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAuthenticatorsResponseBody) SetRequestId(v string) *ListAuthenticatorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthenticatorsResponseBody) SetTotalCount(v int64) *ListAuthenticatorsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAuthenticatorsResponseBodyAuthenticator struct {
	// example:
	//
	// webauth_demo
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// example:
	//
	// TOTP-DefaultName
	AuthenticatorName *string `json:"AuthenticatorName,omitempty" xml:"AuthenticatorName,omitempty"`
	// example:
	//
	// 0fcd5fce3fdc42052633a65eb242f10aZ8nzoQMpV3m
	AuthenticatorUuid *string `json:"AuthenticatorUuid,omitempty" xml:"AuthenticatorUuid,omitempty"`
	// example:
	//
	// scimcred-00skzv5noh34r5aoci76
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// example:
	//
	// 123123
	LastVerifyTime *int64 `json:"LastVerifyTime,omitempty" xml:"LastVerifyTime,omitempty"`
	// example:
	//
	// 1342800000000
	RegisterTime *int64 `json:"RegisterTime,omitempty" xml:"RegisterTime,omitempty"`
	// example:
	//
	// TOTP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAuthenticatorsResponseBodyAuthenticator) String() string {
	return tea.Prettify(s)
}

func (s ListAuthenticatorsResponseBodyAuthenticator) GoString() string {
	return s.String()
}

func (s *ListAuthenticatorsResponseBodyAuthenticator) SetApplicationExternalId(v string) *ListAuthenticatorsResponseBodyAuthenticator {
	s.ApplicationExternalId = &v
	return s
}

func (s *ListAuthenticatorsResponseBodyAuthenticator) SetAuthenticatorName(v string) *ListAuthenticatorsResponseBodyAuthenticator {
	s.AuthenticatorName = &v
	return s
}

func (s *ListAuthenticatorsResponseBodyAuthenticator) SetAuthenticatorUuid(v string) *ListAuthenticatorsResponseBodyAuthenticator {
	s.AuthenticatorUuid = &v
	return s
}

func (s *ListAuthenticatorsResponseBodyAuthenticator) SetCredentialId(v string) *ListAuthenticatorsResponseBodyAuthenticator {
	s.CredentialId = &v
	return s
}

func (s *ListAuthenticatorsResponseBodyAuthenticator) SetLastVerifyTime(v int64) *ListAuthenticatorsResponseBodyAuthenticator {
	s.LastVerifyTime = &v
	return s
}

func (s *ListAuthenticatorsResponseBodyAuthenticator) SetRegisterTime(v int64) *ListAuthenticatorsResponseBodyAuthenticator {
	s.RegisterTime = &v
	return s
}

func (s *ListAuthenticatorsResponseBodyAuthenticator) SetType(v string) *ListAuthenticatorsResponseBodyAuthenticator {
	s.Type = &v
	return s
}

type ListAuthenticatorsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthenticatorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthenticatorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAuthenticatorsResponse) GoString() string {
	return s.String()
}

func (s *ListAuthenticatorsResponse) SetHeaders(v map[string]*string) *ListAuthenticatorsResponse {
	s.Headers = v
	return s
}

func (s *ListAuthenticatorsResponse) SetStatusCode(v int32) *ListAuthenticatorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthenticatorsResponse) SetBody(v *ListAuthenticatorsResponseBody) *ListAuthenticatorsResponse {
	s.Body = v
	return s
}

type ListCostUnitOrdersRequest struct {
	// example:
	//
	// 2021-11-16
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// example:
	//
	// 2021-12-04 00
	FinalDate *string `json:"FinalDate,omitempty" xml:"FinalDate,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCostUnitOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCostUnitOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListCostUnitOrdersRequest) SetBeginDate(v string) *ListCostUnitOrdersRequest {
	s.BeginDate = &v
	return s
}

func (s *ListCostUnitOrdersRequest) SetFinalDate(v string) *ListCostUnitOrdersRequest {
	s.FinalDate = &v
	return s
}

func (s *ListCostUnitOrdersRequest) SetPageNumber(v int32) *ListCostUnitOrdersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCostUnitOrdersRequest) SetPageSize(v int32) *ListCostUnitOrdersRequest {
	s.PageSize = &v
	return s
}

type ListCostUnitOrdersResponseBody struct {
	Items []*ListCostUnitOrdersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// A82621A5-715C-5486-80E9-9262EE43BA55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// example:
	//
	// 2
	TotalPages *int64 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListCostUnitOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCostUnitOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListCostUnitOrdersResponseBody) SetItems(v []*ListCostUnitOrdersResponseBodyItems) *ListCostUnitOrdersResponseBody {
	s.Items = v
	return s
}

func (s *ListCostUnitOrdersResponseBody) SetPageSize(v int64) *ListCostUnitOrdersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCostUnitOrdersResponseBody) SetRequestId(v string) *ListCostUnitOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCostUnitOrdersResponseBody) SetTotalElements(v int64) *ListCostUnitOrdersResponseBody {
	s.TotalElements = &v
	return s
}

func (s *ListCostUnitOrdersResponseBody) SetTotalPages(v int64) *ListCostUnitOrdersResponseBody {
	s.TotalPages = &v
	return s
}

type ListCostUnitOrdersResponseBodyItems struct {
	// example:
	//
	// 215222637610732
	AliOrderCode *string `json:"AliOrderCode,omitempty" xml:"AliOrderCode,omitempty"`
	// example:
	//
	// idaas_auth_public_cn-pl32pciy101
	AliOrderInstanceId *string `json:"AliOrderInstanceId,omitempty" xml:"AliOrderInstanceId,omitempty"`
	// example:
	//
	// 1646028448000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Exhausted  *bool  `json:"Exhausted,omitempty" xml:"Exhausted,omitempty"`
	// example:
	//
	// 1923321600000
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// VALID
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// example:
	//
	// 1923321600000
	RefundTime *int64 `json:"RefundTime,omitempty" xml:"RefundTime,omitempty"`
	// example:
	//
	// 100000
	TotalCostUnit *int64 `json:"TotalCostUnit,omitempty" xml:"TotalCostUnit,omitempty"`
	// example:
	//
	// 100
	UsedCostUnit *int64 `json:"UsedCostUnit,omitempty" xml:"UsedCostUnit,omitempty"`
}

func (s ListCostUnitOrdersResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListCostUnitOrdersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListCostUnitOrdersResponseBodyItems) SetAliOrderCode(v string) *ListCostUnitOrdersResponseBodyItems {
	s.AliOrderCode = &v
	return s
}

func (s *ListCostUnitOrdersResponseBodyItems) SetAliOrderInstanceId(v string) *ListCostUnitOrdersResponseBodyItems {
	s.AliOrderInstanceId = &v
	return s
}

func (s *ListCostUnitOrdersResponseBodyItems) SetCreateTime(v int64) *ListCostUnitOrdersResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *ListCostUnitOrdersResponseBodyItems) SetExhausted(v bool) *ListCostUnitOrdersResponseBodyItems {
	s.Exhausted = &v
	return s
}

func (s *ListCostUnitOrdersResponseBodyItems) SetExpiredTime(v int64) *ListCostUnitOrdersResponseBodyItems {
	s.ExpiredTime = &v
	return s
}

func (s *ListCostUnitOrdersResponseBodyItems) SetOrderStatus(v string) *ListCostUnitOrdersResponseBodyItems {
	s.OrderStatus = &v
	return s
}

func (s *ListCostUnitOrdersResponseBodyItems) SetRefundTime(v int64) *ListCostUnitOrdersResponseBodyItems {
	s.RefundTime = &v
	return s
}

func (s *ListCostUnitOrdersResponseBodyItems) SetTotalCostUnit(v int64) *ListCostUnitOrdersResponseBodyItems {
	s.TotalCostUnit = &v
	return s
}

func (s *ListCostUnitOrdersResponseBodyItems) SetUsedCostUnit(v int64) *ListCostUnitOrdersResponseBodyItems {
	s.UsedCostUnit = &v
	return s
}

type ListCostUnitOrdersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCostUnitOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCostUnitOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCostUnitOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListCostUnitOrdersResponse) SetHeaders(v map[string]*string) *ListCostUnitOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListCostUnitOrdersResponse) SetStatusCode(v int32) *ListCostUnitOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCostUnitOrdersResponse) SetBody(v *ListCostUnitOrdersResponseBody) *ListCostUnitOrdersResponse {
	s.Body = v
	return s
}

type ListOrderConsumeStatisticRecordsRequest struct {
	// example:
	//
	// 211283925690732
	AliOrderCode          *string `json:"AliOrderCode,omitempty" xml:"AliOrderCode,omitempty"`
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// 2022-05-31
	StatisticTimeMax *string `json:"StatisticTimeMax,omitempty" xml:"StatisticTimeMax,omitempty"`
	// example:
	//
	// 2022-05-31
	StatisticTimeMin *string `json:"StatisticTimeMin,omitempty" xml:"StatisticTimeMin,omitempty"`
}

func (s ListOrderConsumeStatisticRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrderConsumeStatisticRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListOrderConsumeStatisticRecordsRequest) SetAliOrderCode(v string) *ListOrderConsumeStatisticRecordsRequest {
	s.AliOrderCode = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsRequest) SetApplicationExternalId(v string) *ListOrderConsumeStatisticRecordsRequest {
	s.ApplicationExternalId = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsRequest) SetPageNumber(v int32) *ListOrderConsumeStatisticRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsRequest) SetPageSize(v int32) *ListOrderConsumeStatisticRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsRequest) SetServiceCode(v string) *ListOrderConsumeStatisticRecordsRequest {
	s.ServiceCode = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsRequest) SetStatisticTimeMax(v string) *ListOrderConsumeStatisticRecordsRequest {
	s.StatisticTimeMax = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsRequest) SetStatisticTimeMin(v string) *ListOrderConsumeStatisticRecordsRequest {
	s.StatisticTimeMin = &v
	return s
}

type ListOrderConsumeStatisticRecordsResponseBody struct {
	Items []*ListOrderConsumeStatisticRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1655384783875
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// example:
	//
	// 10
	TotalPages *int64 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListOrderConsumeStatisticRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrderConsumeStatisticRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrderConsumeStatisticRecordsResponseBody) SetItems(v []*ListOrderConsumeStatisticRecordsResponseBodyItems) *ListOrderConsumeStatisticRecordsResponseBody {
	s.Items = v
	return s
}

func (s *ListOrderConsumeStatisticRecordsResponseBody) SetPageSize(v int64) *ListOrderConsumeStatisticRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsResponseBody) SetRequestId(v string) *ListOrderConsumeStatisticRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsResponseBody) SetTotalElements(v int64) *ListOrderConsumeStatisticRecordsResponseBody {
	s.TotalElements = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsResponseBody) SetTotalPages(v int64) *ListOrderConsumeStatisticRecordsResponseBody {
	s.TotalPages = &v
	return s
}

type ListOrderConsumeStatisticRecordsResponseBodyItems struct {
	// example:
	//
	// 23567876823
	AliOrderCode *string `json:"AliOrderCode,omitempty" xml:"AliOrderCode,omitempty"`
	// example:
	//
	// sms_auth
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// example:
	//
	// 1
	ChargedCount *int64 `json:"ChargedCount,omitempty" xml:"ChargedCount,omitempty"`
	// example:
	//
	// SMS
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// 1655388050
	StatisticTime *int64 `json:"StatisticTime,omitempty" xml:"StatisticTime,omitempty"`
	// example:
	//
	// 10
	TotalPrice *int64 `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
	// example:
	//
	// 10
	UnitPrice *int64 `json:"UnitPrice,omitempty" xml:"UnitPrice,omitempty"`
}

func (s ListOrderConsumeStatisticRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListOrderConsumeStatisticRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListOrderConsumeStatisticRecordsResponseBodyItems) SetAliOrderCode(v string) *ListOrderConsumeStatisticRecordsResponseBodyItems {
	s.AliOrderCode = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsResponseBodyItems) SetApplicationExternalId(v string) *ListOrderConsumeStatisticRecordsResponseBodyItems {
	s.ApplicationExternalId = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsResponseBodyItems) SetChargedCount(v int64) *ListOrderConsumeStatisticRecordsResponseBodyItems {
	s.ChargedCount = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsResponseBodyItems) SetServiceCode(v string) *ListOrderConsumeStatisticRecordsResponseBodyItems {
	s.ServiceCode = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsResponseBodyItems) SetStatisticTime(v int64) *ListOrderConsumeStatisticRecordsResponseBodyItems {
	s.StatisticTime = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsResponseBodyItems) SetTotalPrice(v int64) *ListOrderConsumeStatisticRecordsResponseBodyItems {
	s.TotalPrice = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsResponseBodyItems) SetUnitPrice(v int64) *ListOrderConsumeStatisticRecordsResponseBodyItems {
	s.UnitPrice = &v
	return s
}

type ListOrderConsumeStatisticRecordsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrderConsumeStatisticRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrderConsumeStatisticRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrderConsumeStatisticRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListOrderConsumeStatisticRecordsResponse) SetHeaders(v map[string]*string) *ListOrderConsumeStatisticRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListOrderConsumeStatisticRecordsResponse) SetStatusCode(v int32) *ListOrderConsumeStatisticRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrderConsumeStatisticRecordsResponse) SetBody(v *ListOrderConsumeStatisticRecordsResponseBody) *ListOrderConsumeStatisticRecordsResponse {
	s.Body = v
	return s
}

type ListPwnedPasswordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 000000
	PrefixHexPasswordSha1Hash *string `json:"PrefixHexPasswordSha1Hash,omitempty" xml:"PrefixHexPasswordSha1Hash,omitempty"`
}

func (s ListPwnedPasswordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPwnedPasswordsRequest) GoString() string {
	return s.String()
}

func (s *ListPwnedPasswordsRequest) SetPrefixHexPasswordSha1Hash(v string) *ListPwnedPasswordsRequest {
	s.PrefixHexPasswordSha1Hash = &v
	return s
}

type ListPwnedPasswordsResponseBody struct {
	// example:
	//
	// 2
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize           *int64                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PwnedPasswordInfos []*ListPwnedPasswordsResponseBodyPwnedPasswordInfos `json:"PwnedPasswordInfos,omitempty" xml:"PwnedPasswordInfos,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// D0F570B0-89CE-07F2-B41E-F97E089E13C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 80
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPwnedPasswordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPwnedPasswordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPwnedPasswordsResponseBody) SetPageNumber(v int64) *ListPwnedPasswordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPwnedPasswordsResponseBody) SetPageSize(v int64) *ListPwnedPasswordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPwnedPasswordsResponseBody) SetPwnedPasswordInfos(v []*ListPwnedPasswordsResponseBodyPwnedPasswordInfos) *ListPwnedPasswordsResponseBody {
	s.PwnedPasswordInfos = v
	return s
}

func (s *ListPwnedPasswordsResponseBody) SetRequestId(v string) *ListPwnedPasswordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPwnedPasswordsResponseBody) SetTotalCount(v int64) *ListPwnedPasswordsResponseBody {
	s.TotalCount = &v
	return s
}

type ListPwnedPasswordsResponseBodyPwnedPasswordInfos struct {
	// example:
	//
	// 0000
	HexPasswordSha1Hash *string `json:"HexPasswordSha1Hash,omitempty" xml:"HexPasswordSha1Hash,omitempty"`
	// example:
	//
	// 84
	PwnedCount *int64 `json:"PwnedCount,omitempty" xml:"PwnedCount,omitempty"`
}

func (s ListPwnedPasswordsResponseBodyPwnedPasswordInfos) String() string {
	return tea.Prettify(s)
}

func (s ListPwnedPasswordsResponseBodyPwnedPasswordInfos) GoString() string {
	return s.String()
}

func (s *ListPwnedPasswordsResponseBodyPwnedPasswordInfos) SetHexPasswordSha1Hash(v string) *ListPwnedPasswordsResponseBodyPwnedPasswordInfos {
	s.HexPasswordSha1Hash = &v
	return s
}

func (s *ListPwnedPasswordsResponseBodyPwnedPasswordInfos) SetPwnedCount(v int64) *ListPwnedPasswordsResponseBodyPwnedPasswordInfos {
	s.PwnedCount = &v
	return s
}

type ListPwnedPasswordsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPwnedPasswordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPwnedPasswordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPwnedPasswordsResponse) GoString() string {
	return s.String()
}

func (s *ListPwnedPasswordsResponse) SetHeaders(v map[string]*string) *ListPwnedPasswordsResponse {
	s.Headers = v
	return s
}

func (s *ListPwnedPasswordsResponse) SetStatusCode(v int32) *ListPwnedPasswordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPwnedPasswordsResponse) SetBody(v *ListPwnedPasswordsResponseBody) *ListPwnedPasswordsResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11122223333
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// example:
	//
	// user-11122223333
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetApplicationExternalId(v string) *ListUsersRequest {
	s.ApplicationExternalId = &v
	return s
}

func (s *ListUsersRequest) SetUserId(v string) *ListUsersRequest {
	s.UserId = &v
	return s
}

type ListUsersResponseBody struct {
	// example:
	//
	// D0F570B0-89CE-07F2-B41E-F97E089E13C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 13
	TotalCount *int64                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Users      []*ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int64) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v []*ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

type ListUsersResponseBodyUsers struct {
	UserDisplayName *string `json:"UserDisplayName,omitempty" xml:"UserDisplayName,omitempty"`
	// example:
	//
	// 2323232676856745
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// lili
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) SetUserDisplayName(v string) *ListUsersResponseBodyUsers {
	s.UserDisplayName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserId(v string) *ListUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserName(v string) *ListUsersResponseBodyUsers {
	s.UserName = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type QuerySmsReportsRequest struct {
	// example:
	//
	// doraemon_xxxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// eventId_xxxxxx
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
}

func (s QuerySmsReportsRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsReportsRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsReportsRequest) SetAppId(v string) *QuerySmsReportsRequest {
	s.AppId = &v
	return s
}

func (s *QuerySmsReportsRequest) SetEventId(v string) *QuerySmsReportsRequest {
	s.EventId = &v
	return s
}

type QuerySmsReportsResponseBody struct {
	// example:
	//
	// 4581B16-xxxxx
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SmsReports []*QuerySmsReportsResponseBodySmsReports `json:"SmsReports,omitempty" xml:"SmsReports,omitempty" type:"Repeated"`
	// example:
	//
	// 15
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
}

func (s QuerySmsReportsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsReportsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsReportsResponseBody) SetRequestId(v string) *QuerySmsReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsReportsResponseBody) SetSmsReports(v []*QuerySmsReportsResponseBodySmsReports) *QuerySmsReportsResponseBody {
	s.SmsReports = v
	return s
}

func (s *QuerySmsReportsResponseBody) SetTotalElements(v int64) *QuerySmsReportsResponseBody {
	s.TotalElements = &v
	return s
}

type QuerySmsReportsResponseBodySmsReports struct {
	// example:
	//
	// appid
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChargedCount *int32  `json:"ChargedCount,omitempty" xml:"ChargedCount,omitempty"`
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// event_xxxx
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// 156xxxxxxxx
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// "d36xxxxx3573772f159f616ebJ8KGPBg7Hr::appId::123456::Bird"
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// example:
	//
	// DELIVRD
	Stat *string `json:"Stat,omitempty" xml:"Stat,omitempty"`
	// example:
	//
	// d36a370f33f037233573772f15xxxxxx
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 3514
	Tid *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// 20220823171446
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QuerySmsReportsResponseBodySmsReports) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsReportsResponseBodySmsReports) GoString() string {
	return s.String()
}

func (s *QuerySmsReportsResponseBodySmsReports) SetAppId(v string) *QuerySmsReportsResponseBodySmsReports {
	s.AppId = &v
	return s
}

func (s *QuerySmsReportsResponseBodySmsReports) SetChargedCount(v int32) *QuerySmsReportsResponseBodySmsReports {
	s.ChargedCount = &v
	return s
}

func (s *QuerySmsReportsResponseBodySmsReports) SetCode(v string) *QuerySmsReportsResponseBodySmsReports {
	s.Code = &v
	return s
}

func (s *QuerySmsReportsResponseBodySmsReports) SetEventId(v string) *QuerySmsReportsResponseBodySmsReports {
	s.EventId = &v
	return s
}

func (s *QuerySmsReportsResponseBodySmsReports) SetMobile(v string) *QuerySmsReportsResponseBodySmsReports {
	s.Mobile = &v
	return s
}

func (s *QuerySmsReportsResponseBodySmsReports) SetSn(v string) *QuerySmsReportsResponseBodySmsReports {
	s.Sn = &v
	return s
}

func (s *QuerySmsReportsResponseBodySmsReports) SetStat(v string) *QuerySmsReportsResponseBodySmsReports {
	s.Stat = &v
	return s
}

func (s *QuerySmsReportsResponseBodySmsReports) SetTenantId(v string) *QuerySmsReportsResponseBodySmsReports {
	s.TenantId = &v
	return s
}

func (s *QuerySmsReportsResponseBodySmsReports) SetTid(v string) *QuerySmsReportsResponseBodySmsReports {
	s.Tid = &v
	return s
}

func (s *QuerySmsReportsResponseBodySmsReports) SetTime(v string) *QuerySmsReportsResponseBodySmsReports {
	s.Time = &v
	return s
}

type QuerySmsReportsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmsReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySmsReportsResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsReportsResponse) GoString() string {
	return s.String()
}

func (s *QuerySmsReportsResponse) SetHeaders(v map[string]*string) *QuerySmsReportsResponse {
	s.Headers = v
	return s
}

func (s *QuerySmsReportsResponse) SetStatusCode(v int32) *QuerySmsReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmsReportsResponse) SetBody(v *QuerySmsReportsResponseBody) *QuerySmsReportsResponse {
	s.Body = v
	return s
}

type QuerySmsUpsResponseBody struct {
	RequestId     *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SmsUps        []*QuerySmsUpsResponseBodySmsUps `json:"SmsUps,omitempty" xml:"SmsUps,omitempty" type:"Repeated"`
	TotalElements *int64                           `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
}

func (s QuerySmsUpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsUpsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsUpsResponseBody) SetRequestId(v string) *QuerySmsUpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsUpsResponseBody) SetSmsUps(v []*QuerySmsUpsResponseBodySmsUps) *QuerySmsUpsResponseBody {
	s.SmsUps = v
	return s
}

func (s *QuerySmsUpsResponseBody) SetTotalElements(v int64) *QuerySmsUpsResponseBody {
	s.TotalElements = &v
	return s
}

type QuerySmsUpsResponseBodySmsUps struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DestCode    *string `json:"DestCode,omitempty" xml:"DestCode,omitempty"`
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	SendTime    *string `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	SequenceId  *string `json:"SequenceId,omitempty" xml:"SequenceId,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QuerySmsUpsResponseBodySmsUps) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsUpsResponseBodySmsUps) GoString() string {
	return s.String()
}

func (s *QuerySmsUpsResponseBodySmsUps) SetContent(v string) *QuerySmsUpsResponseBodySmsUps {
	s.Content = &v
	return s
}

func (s *QuerySmsUpsResponseBodySmsUps) SetDestCode(v string) *QuerySmsUpsResponseBodySmsUps {
	s.DestCode = &v
	return s
}

func (s *QuerySmsUpsResponseBodySmsUps) SetPhoneNumber(v string) *QuerySmsUpsResponseBodySmsUps {
	s.PhoneNumber = &v
	return s
}

func (s *QuerySmsUpsResponseBodySmsUps) SetSendTime(v string) *QuerySmsUpsResponseBodySmsUps {
	s.SendTime = &v
	return s
}

func (s *QuerySmsUpsResponseBodySmsUps) SetSequenceId(v string) *QuerySmsUpsResponseBodySmsUps {
	s.SequenceId = &v
	return s
}

func (s *QuerySmsUpsResponseBodySmsUps) SetTenantId(v string) *QuerySmsUpsResponseBodySmsUps {
	s.TenantId = &v
	return s
}

type QuerySmsUpsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmsUpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySmsUpsResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySmsUpsResponse) GoString() string {
	return s.String()
}

func (s *QuerySmsUpsResponse) SetHeaders(v map[string]*string) *QuerySmsUpsResponse {
	s.Headers = v
	return s
}

func (s *QuerySmsUpsResponse) SetStatusCode(v int32) *QuerySmsUpsResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmsUpsResponse) SetBody(v *QuerySmsUpsResponseBody) *QuerySmsUpsResponse {
	s.Body = v
	return s
}

type RegisterAuthenticatorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11122223333
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// example:
	//
	// ZhangSan\\"s authenticator
	AuthenticatorName *string `json:"AuthenticatorName,omitempty" xml:"AuthenticatorName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// webauthn
	AuthenticatorType *string `json:"AuthenticatorType,omitempty" xml:"AuthenticatorType,omitempty"`
	// example:
	//
	// 由具体认证方式决定
	ClientExtendParamsJson *string `json:"ClientExtendParamsJson,omitempty" xml:"ClientExtendParamsJson,omitempty"`
	// example:
	//
	// 由具体认证方式决定
	ClientExtendParamsJsonSign *string `json:"ClientExtendParamsJsonSign,omitempty" xml:"ClientExtendParamsJsonSign,omitempty"`
	// example:
	//
	// 123123
	LogParams *string `json:"LogParams,omitempty" xml:"LogParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {}
	RegistrationContext *string `json:"RegistrationContext,omitempty" xml:"RegistrationContext,omitempty"`
	// example:
	//
	// AAAAdda
	RequireChallengeBase64 *string `json:"RequireChallengeBase64,omitempty" xml:"RequireChallengeBase64,omitempty"`
	// example:
	//
	// 由具体认证方式决定
	ServerExtendParamsJson *string `json:"ServerExtendParamsJson,omitempty" xml:"ServerExtendParamsJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user-11122223333
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserSourceIp *string `json:"UserSourceIp,omitempty" xml:"UserSourceIp,omitempty"`
}

func (s RegisterAuthenticatorRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterAuthenticatorRequest) GoString() string {
	return s.String()
}

func (s *RegisterAuthenticatorRequest) SetApplicationExternalId(v string) *RegisterAuthenticatorRequest {
	s.ApplicationExternalId = &v
	return s
}

func (s *RegisterAuthenticatorRequest) SetAuthenticatorName(v string) *RegisterAuthenticatorRequest {
	s.AuthenticatorName = &v
	return s
}

func (s *RegisterAuthenticatorRequest) SetAuthenticatorType(v string) *RegisterAuthenticatorRequest {
	s.AuthenticatorType = &v
	return s
}

func (s *RegisterAuthenticatorRequest) SetClientExtendParamsJson(v string) *RegisterAuthenticatorRequest {
	s.ClientExtendParamsJson = &v
	return s
}

func (s *RegisterAuthenticatorRequest) SetClientExtendParamsJsonSign(v string) *RegisterAuthenticatorRequest {
	s.ClientExtendParamsJsonSign = &v
	return s
}

func (s *RegisterAuthenticatorRequest) SetLogParams(v string) *RegisterAuthenticatorRequest {
	s.LogParams = &v
	return s
}

func (s *RegisterAuthenticatorRequest) SetRegistrationContext(v string) *RegisterAuthenticatorRequest {
	s.RegistrationContext = &v
	return s
}

func (s *RegisterAuthenticatorRequest) SetRequireChallengeBase64(v string) *RegisterAuthenticatorRequest {
	s.RequireChallengeBase64 = &v
	return s
}

func (s *RegisterAuthenticatorRequest) SetServerExtendParamsJson(v string) *RegisterAuthenticatorRequest {
	s.ServerExtendParamsJson = &v
	return s
}

func (s *RegisterAuthenticatorRequest) SetUserId(v string) *RegisterAuthenticatorRequest {
	s.UserId = &v
	return s
}

func (s *RegisterAuthenticatorRequest) SetUserSourceIp(v string) *RegisterAuthenticatorRequest {
	s.UserSourceIp = &v
	return s
}

type RegisterAuthenticatorResponseBody struct {
	// example:
	//
	// 123123123
	AuthenticatorUuid *string `json:"AuthenticatorUuid,omitempty" xml:"AuthenticatorUuid,omitempty"`
	EtasResponseSting *string `json:"EtasResponseSting,omitempty" xml:"EtasResponseSting,omitempty"`
	// example:
	//
	// 23123
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterAuthenticatorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterAuthenticatorResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterAuthenticatorResponseBody) SetAuthenticatorUuid(v string) *RegisterAuthenticatorResponseBody {
	s.AuthenticatorUuid = &v
	return s
}

func (s *RegisterAuthenticatorResponseBody) SetEtasResponseSting(v string) *RegisterAuthenticatorResponseBody {
	s.EtasResponseSting = &v
	return s
}

func (s *RegisterAuthenticatorResponseBody) SetRequestId(v string) *RegisterAuthenticatorResponseBody {
	s.RequestId = &v
	return s
}

type RegisterAuthenticatorResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterAuthenticatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterAuthenticatorResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterAuthenticatorResponse) GoString() string {
	return s.String()
}

func (s *RegisterAuthenticatorResponse) SetHeaders(v map[string]*string) *RegisterAuthenticatorResponse {
	s.Headers = v
	return s
}

func (s *RegisterAuthenticatorResponse) SetStatusCode(v int32) *RegisterAuthenticatorResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterAuthenticatorResponse) SetBody(v *RegisterAuthenticatorResponseBody) *RegisterAuthenticatorResponse {
	s.Body = v
	return s
}

type ServiceInvokeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// A0000003
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SimpleAuthInit
	DoraemonAction *string `json:"DoraemonAction,omitempty" xml:"DoraemonAction,omitempty"`
	// example:
	//
	// eyJhdHRhY2htZW50IjoicGxhdGZvcm0iLCJ1c2VyVmVyaWZpY2F0aW9uIjoicmVxdWlyZWQifQ==
	MobileExtendParamsJson *string `json:"MobileExtendParamsJson,omitempty" xml:"MobileExtendParamsJson,omitempty"`
	// example:
	//
	// ba6f5596f00102cb1a7d971f8390b7ccb28e6bc4e8694d051f9299ccc69e1186
	MobileExtendParamsJsonSign *string `json:"MobileExtendParamsJsonSign,omitempty" xml:"MobileExtendParamsJsonSign,omitempty"`
	// example:
	//
	// eyJhdHRhY2htZW50IjoicGxhdGZvcm0iLCJ1c2VyVmVyaWZpY2F0aW9uIjoicmVxdWlyZWQifQ==
	ServerExtendParamsJson *string `json:"ServerExtendParamsJson,omitempty" xml:"ServerExtendParamsJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// IFAA
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// false
	TestModel *bool `json:"TestModel,omitempty" xml:"TestModel,omitempty"`
	// example:
	//
	// 127.0.0.1
	XClientIp *string `json:"XClientIp,omitempty" xml:"XClientIp,omitempty"`
}

func (s ServiceInvokeRequest) String() string {
	return tea.Prettify(s)
}

func (s ServiceInvokeRequest) GoString() string {
	return s.String()
}

func (s *ServiceInvokeRequest) SetApplicationExternalId(v string) *ServiceInvokeRequest {
	s.ApplicationExternalId = &v
	return s
}

func (s *ServiceInvokeRequest) SetDoraemonAction(v string) *ServiceInvokeRequest {
	s.DoraemonAction = &v
	return s
}

func (s *ServiceInvokeRequest) SetMobileExtendParamsJson(v string) *ServiceInvokeRequest {
	s.MobileExtendParamsJson = &v
	return s
}

func (s *ServiceInvokeRequest) SetMobileExtendParamsJsonSign(v string) *ServiceInvokeRequest {
	s.MobileExtendParamsJsonSign = &v
	return s
}

func (s *ServiceInvokeRequest) SetServerExtendParamsJson(v string) *ServiceInvokeRequest {
	s.ServerExtendParamsJson = &v
	return s
}

func (s *ServiceInvokeRequest) SetServiceCode(v string) *ServiceInvokeRequest {
	s.ServiceCode = &v
	return s
}

func (s *ServiceInvokeRequest) SetTestModel(v bool) *ServiceInvokeRequest {
	s.TestModel = &v
	return s
}

func (s *ServiceInvokeRequest) SetXClientIp(v string) *ServiceInvokeRequest {
	s.XClientIp = &v
	return s
}

type ServiceInvokeResponseBody struct {
	// example:
	//
	// Operation.Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// "{\\"ifaa\\"{\\"code\\":0,\\"message\\":\\"eyJo....BWUFBQUFBIn19\\"},\\"version\\":\\"1.0.0\\",\\"transaction\\":{\\"id\\":\\"sIFAA-ios\\",\\"type\\":\\"sIFAA-ios\\",\\"payload\\":\\"\\"},\\"action\\":\\"request/auth\\"}" }
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// qacdfhigvbcvb
	IdToken *string `json:"IdToken,omitempty" xml:"IdToken,omitempty"`
	// example:
	//
	// Operation.Failure.ApiInvoke.ServiceInArrears
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F8F63CAB-A96C-5017-AFB0-324593223817
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ServiceInvokeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ServiceInvokeResponseBody) GoString() string {
	return s.String()
}

func (s *ServiceInvokeResponseBody) SetCode(v string) *ServiceInvokeResponseBody {
	s.Code = &v
	return s
}

func (s *ServiceInvokeResponseBody) SetData(v string) *ServiceInvokeResponseBody {
	s.Data = &v
	return s
}

func (s *ServiceInvokeResponseBody) SetEventId(v string) *ServiceInvokeResponseBody {
	s.EventId = &v
	return s
}

func (s *ServiceInvokeResponseBody) SetIdToken(v string) *ServiceInvokeResponseBody {
	s.IdToken = &v
	return s
}

func (s *ServiceInvokeResponseBody) SetMessage(v string) *ServiceInvokeResponseBody {
	s.Message = &v
	return s
}

func (s *ServiceInvokeResponseBody) SetRequestId(v string) *ServiceInvokeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ServiceInvokeResponseBody) SetSuccess(v bool) *ServiceInvokeResponseBody {
	s.Success = &v
	return s
}

type ServiceInvokeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ServiceInvokeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ServiceInvokeResponse) String() string {
	return tea.Prettify(s)
}

func (s ServiceInvokeResponse) GoString() string {
	return s.String()
}

func (s *ServiceInvokeResponse) SetHeaders(v map[string]*string) *ServiceInvokeResponse {
	s.Headers = v
	return s
}

func (s *ServiceInvokeResponse) SetStatusCode(v int32) *ServiceInvokeResponse {
	s.StatusCode = &v
	return s
}

func (s *ServiceInvokeResponse) SetBody(v *ServiceInvokeResponseBody) *ServiceInvokeResponse {
	s.Body = v
	return s
}

type UpdateAuthenticatorAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11122223333
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TOTP-DefaultName
	AuthenticatorName *string `json:"AuthenticatorName,omitempty" xml:"AuthenticatorName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c810731ad3af06e45171395abf19cf9cb00E7i97PhE
	AuthenticatorUuid *string `json:"AuthenticatorUuid,omitempty" xml:"AuthenticatorUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user-11122223333
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateAuthenticatorAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuthenticatorAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthenticatorAttributeRequest) SetApplicationExternalId(v string) *UpdateAuthenticatorAttributeRequest {
	s.ApplicationExternalId = &v
	return s
}

func (s *UpdateAuthenticatorAttributeRequest) SetAuthenticatorName(v string) *UpdateAuthenticatorAttributeRequest {
	s.AuthenticatorName = &v
	return s
}

func (s *UpdateAuthenticatorAttributeRequest) SetAuthenticatorUuid(v string) *UpdateAuthenticatorAttributeRequest {
	s.AuthenticatorUuid = &v
	return s
}

func (s *UpdateAuthenticatorAttributeRequest) SetUserId(v string) *UpdateAuthenticatorAttributeRequest {
	s.UserId = &v
	return s
}

type UpdateAuthenticatorAttributeResponseBody struct {
	// example:
	//
	// D0F570B0-89CE-07F2-B41E-F97E089E13C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAuthenticatorAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuthenticatorAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuthenticatorAttributeResponseBody) SetRequestId(v string) *UpdateAuthenticatorAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAuthenticatorAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuthenticatorAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuthenticatorAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuthenticatorAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthenticatorAttributeResponse) SetHeaders(v map[string]*string) *UpdateAuthenticatorAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthenticatorAttributeResponse) SetStatusCode(v int32) *UpdateAuthenticatorAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuthenticatorAttributeResponse) SetBody(v *UpdateAuthenticatorAttributeResponseBody) *UpdateAuthenticatorAttributeResponse {
	s.Body = v
	return s
}

type VerifyIdTokenRequest struct {
	// example:
	//
	// webauth_demo
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// jwtIdToken
	//
	// example:
	//
	// eyJraWQiOiIzNj
	JwtIdToken *string `json:"JwtIdToken,omitempty" xml:"JwtIdToken,omitempty"`
}

func (s VerifyIdTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyIdTokenRequest) GoString() string {
	return s.String()
}

func (s *VerifyIdTokenRequest) SetApplicationExternalId(v string) *VerifyIdTokenRequest {
	s.ApplicationExternalId = &v
	return s
}

func (s *VerifyIdTokenRequest) SetJwtIdToken(v string) *VerifyIdTokenRequest {
	s.JwtIdToken = &v
	return s
}

type VerifyIdTokenResponseBody struct {
	// example:
	//
	// F620E712-F325-520D-BE1A-E2421B062F57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s VerifyIdTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyIdTokenResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyIdTokenResponseBody) SetRequestId(v string) *VerifyIdTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyIdTokenResponseBody) SetUserId(v string) *VerifyIdTokenResponseBody {
	s.UserId = &v
	return s
}

type VerifyIdTokenResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyIdTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyIdTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyIdTokenResponse) GoString() string {
	return s.String()
}

func (s *VerifyIdTokenResponse) SetHeaders(v map[string]*string) *VerifyIdTokenResponse {
	s.Headers = v
	return s
}

func (s *VerifyIdTokenResponse) SetStatusCode(v int32) *VerifyIdTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyIdTokenResponse) SetBody(v *VerifyIdTokenResponseBody) *VerifyIdTokenResponse {
	s.Body = v
	return s
}

type VerifyUserAuthenticationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11122223333
	ApplicationExternalId *string `json:"ApplicationExternalId,omitempty" xml:"ApplicationExternalId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {}
	AuthenticationContext *string `json:"AuthenticationContext,omitempty" xml:"AuthenticationContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// webauthn
	AuthenticatorType *string `json:"AuthenticatorType,omitempty" xml:"AuthenticatorType,omitempty"`
	// example:
	//
	// 由具体认证方式决定
	ClientExtendParamsJson *string `json:"ClientExtendParamsJson,omitempty" xml:"ClientExtendParamsJson,omitempty"`
	// example:
	//
	// 由具体认证方式决定
	ClientExtendParamsJsonSign *string `json:"ClientExtendParamsJsonSign,omitempty" xml:"ClientExtendParamsJsonSign,omitempty"`
	// example:
	//
	// 123123
	LogParams *string `json:"LogParams,omitempty" xml:"LogParams,omitempty"`
	// example:
	//
	// login
	LogTag *string `json:"LogTag,omitempty" xml:"LogTag,omitempty"`
	// example:
	//
	// AAAAdda
	RequireBindHashBase64 *string `json:"RequireBindHashBase64,omitempty" xml:"RequireBindHashBase64,omitempty"`
	// example:
	//
	// AAAAdda
	RequireChallengeBase64 *string `json:"RequireChallengeBase64,omitempty" xml:"RequireChallengeBase64,omitempty"`
	// example:
	//
	// 由具体认证方式决定
	ServerExtendParamsJson *string `json:"ServerExtendParamsJson,omitempty" xml:"ServerExtendParamsJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user-11122223333
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserSourceIp *string `json:"UserSourceIp,omitempty" xml:"UserSourceIp,omitempty"`
}

func (s VerifyUserAuthenticationRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyUserAuthenticationRequest) GoString() string {
	return s.String()
}

func (s *VerifyUserAuthenticationRequest) SetApplicationExternalId(v string) *VerifyUserAuthenticationRequest {
	s.ApplicationExternalId = &v
	return s
}

func (s *VerifyUserAuthenticationRequest) SetAuthenticationContext(v string) *VerifyUserAuthenticationRequest {
	s.AuthenticationContext = &v
	return s
}

func (s *VerifyUserAuthenticationRequest) SetAuthenticatorType(v string) *VerifyUserAuthenticationRequest {
	s.AuthenticatorType = &v
	return s
}

func (s *VerifyUserAuthenticationRequest) SetClientExtendParamsJson(v string) *VerifyUserAuthenticationRequest {
	s.ClientExtendParamsJson = &v
	return s
}

func (s *VerifyUserAuthenticationRequest) SetClientExtendParamsJsonSign(v string) *VerifyUserAuthenticationRequest {
	s.ClientExtendParamsJsonSign = &v
	return s
}

func (s *VerifyUserAuthenticationRequest) SetLogParams(v string) *VerifyUserAuthenticationRequest {
	s.LogParams = &v
	return s
}

func (s *VerifyUserAuthenticationRequest) SetLogTag(v string) *VerifyUserAuthenticationRequest {
	s.LogTag = &v
	return s
}

func (s *VerifyUserAuthenticationRequest) SetRequireBindHashBase64(v string) *VerifyUserAuthenticationRequest {
	s.RequireBindHashBase64 = &v
	return s
}

func (s *VerifyUserAuthenticationRequest) SetRequireChallengeBase64(v string) *VerifyUserAuthenticationRequest {
	s.RequireChallengeBase64 = &v
	return s
}

func (s *VerifyUserAuthenticationRequest) SetServerExtendParamsJson(v string) *VerifyUserAuthenticationRequest {
	s.ServerExtendParamsJson = &v
	return s
}

func (s *VerifyUserAuthenticationRequest) SetUserId(v string) *VerifyUserAuthenticationRequest {
	s.UserId = &v
	return s
}

func (s *VerifyUserAuthenticationRequest) SetUserSourceIp(v string) *VerifyUserAuthenticationRequest {
	s.UserSourceIp = &v
	return s
}

type VerifyUserAuthenticationResponseBody struct {
	AuthenticateResultInfo *VerifyUserAuthenticationResponseBodyAuthenticateResultInfo `json:"AuthenticateResultInfo,omitempty" xml:"AuthenticateResultInfo,omitempty" type:"Struct"`
	// example:
	//
	// zmtdjb
	EtasSDKString *string `json:"EtasSDKString,omitempty" xml:"EtasSDKString,omitempty"`
	// example:
	//
	// qacdfhigvbcvb
	IdToken *string `json:"IdToken,omitempty" xml:"IdToken,omitempty"`
	// example:
	//
	// 83ED3BC7-B36E-53B5-83D4-152CF2088DB9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	VerifyResult *bool `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s VerifyUserAuthenticationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyUserAuthenticationResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyUserAuthenticationResponseBody) SetAuthenticateResultInfo(v *VerifyUserAuthenticationResponseBodyAuthenticateResultInfo) *VerifyUserAuthenticationResponseBody {
	s.AuthenticateResultInfo = v
	return s
}

func (s *VerifyUserAuthenticationResponseBody) SetEtasSDKString(v string) *VerifyUserAuthenticationResponseBody {
	s.EtasSDKString = &v
	return s
}

func (s *VerifyUserAuthenticationResponseBody) SetIdToken(v string) *VerifyUserAuthenticationResponseBody {
	s.IdToken = &v
	return s
}

func (s *VerifyUserAuthenticationResponseBody) SetRequestId(v string) *VerifyUserAuthenticationResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyUserAuthenticationResponseBody) SetVerifyResult(v bool) *VerifyUserAuthenticationResponseBody {
	s.VerifyResult = &v
	return s
}

type VerifyUserAuthenticationResponseBodyAuthenticateResultInfo struct {
	// example:
	//
	// 10e8c224555bba8bfb780badebc4a08cc1Z6LmC6Hp9
	BindHashBase64 *string `json:"BindHashBase64,omitempty" xml:"BindHashBase64,omitempty"`
	// example:
	//
	// scimcred-000vfaknfwlp8kesl520
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// example:
	//
	// 3b32
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s VerifyUserAuthenticationResponseBodyAuthenticateResultInfo) String() string {
	return tea.Prettify(s)
}

func (s VerifyUserAuthenticationResponseBodyAuthenticateResultInfo) GoString() string {
	return s.String()
}

func (s *VerifyUserAuthenticationResponseBodyAuthenticateResultInfo) SetBindHashBase64(v string) *VerifyUserAuthenticationResponseBodyAuthenticateResultInfo {
	s.BindHashBase64 = &v
	return s
}

func (s *VerifyUserAuthenticationResponseBodyAuthenticateResultInfo) SetCredentialId(v string) *VerifyUserAuthenticationResponseBodyAuthenticateResultInfo {
	s.CredentialId = &v
	return s
}

func (s *VerifyUserAuthenticationResponseBodyAuthenticateResultInfo) SetUserId(v string) *VerifyUserAuthenticationResponseBodyAuthenticateResultInfo {
	s.UserId = &v
	return s
}

type VerifyUserAuthenticationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyUserAuthenticationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyUserAuthenticationResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyUserAuthenticationResponse) GoString() string {
	return s.String()
}

func (s *VerifyUserAuthenticationResponse) SetHeaders(v map[string]*string) *VerifyUserAuthenticationResponse {
	s.Headers = v
	return s
}

func (s *VerifyUserAuthenticationResponse) SetStatusCode(v int32) *VerifyUserAuthenticationResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyUserAuthenticationResponse) SetBody(v *VerifyUserAuthenticationResponseBody) *VerifyUserAuthenticationResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou": tea.String("idaas-doraemon.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("idaas-doraemon"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建认证器请求
//
// @param request - CreateAuthenticatorRegistrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAuthenticatorRegistrationResponse
func (client *Client) CreateAuthenticatorRegistrationWithOptions(request *CreateAuthenticatorRegistrationRequest, runtime *util.RuntimeOptions) (_result *CreateAuthenticatorRegistrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationExternalId)) {
		query["ApplicationExternalId"] = request.ApplicationExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticatorType)) {
		query["AuthenticatorType"] = request.AuthenticatorType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientExtendParamsJson)) {
		query["ClientExtendParamsJson"] = request.ClientExtendParamsJson
	}

	if !tea.BoolValue(util.IsUnset(request.ClientExtendParamsJsonSign)) {
		query["ClientExtendParamsJsonSign"] = request.ClientExtendParamsJsonSign
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrationContext)) {
		query["RegistrationContext"] = request.RegistrationContext
	}

	if !tea.BoolValue(util.IsUnset(request.ServerExtendParamsJson)) {
		query["ServerExtendParamsJson"] = request.ServerExtendParamsJson
	}

	if !tea.BoolValue(util.IsUnset(request.UserDisplayName)) {
		query["UserDisplayName"] = request.UserDisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAuthenticatorRegistration"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAuthenticatorRegistrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建认证器请求
//
// @param request - CreateAuthenticatorRegistrationRequest
//
// @return CreateAuthenticatorRegistrationResponse
func (client *Client) CreateAuthenticatorRegistration(request *CreateAuthenticatorRegistrationRequest) (_result *CreateAuthenticatorRegistrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAuthenticatorRegistrationResponse{}
	_body, _err := client.CreateAuthenticatorRegistrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建用户认证请求
//
// @param request - CreateUserAuthenticateOptionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserAuthenticateOptionsResponse
func (client *Client) CreateUserAuthenticateOptionsWithOptions(request *CreateUserAuthenticateOptionsRequest, runtime *util.RuntimeOptions) (_result *CreateUserAuthenticateOptionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationExternalId)) {
		query["ApplicationExternalId"] = request.ApplicationExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticatorType)) {
		query["AuthenticatorType"] = request.AuthenticatorType
	}

	if !tea.BoolValue(util.IsUnset(request.BindHashBase64)) {
		query["BindHashBase64"] = request.BindHashBase64
	}

	if !tea.BoolValue(util.IsUnset(request.ClientExtendParamsJson)) {
		query["ClientExtendParamsJson"] = request.ClientExtendParamsJson
	}

	if !tea.BoolValue(util.IsUnset(request.ClientExtendParamsJsonSign)) {
		query["ClientExtendParamsJsonSign"] = request.ClientExtendParamsJsonSign
	}

	if !tea.BoolValue(util.IsUnset(request.ServerExtendParamsJson)) {
		query["ServerExtendParamsJson"] = request.ServerExtendParamsJson
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserAuthenticateOptions"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserAuthenticateOptionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建用户认证请求
//
// @param request - CreateUserAuthenticateOptionsRequest
//
// @return CreateUserAuthenticateOptionsResponse
func (client *Client) CreateUserAuthenticateOptions(request *CreateUserAuthenticateOptionsRequest) (_result *CreateUserAuthenticateOptionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserAuthenticateOptionsResponse{}
	_body, _err := client.CreateUserAuthenticateOptionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除认证器
//
// @param request - DeregisterAuthenticatorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeregisterAuthenticatorResponse
func (client *Client) DeregisterAuthenticatorWithOptions(request *DeregisterAuthenticatorRequest, runtime *util.RuntimeOptions) (_result *DeregisterAuthenticatorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationExternalId)) {
		query["ApplicationExternalId"] = request.ApplicationExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticatorUuid)) {
		query["AuthenticatorUuid"] = request.AuthenticatorUuid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeregisterAuthenticator"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeregisterAuthenticatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除认证器
//
// @param request - DeregisterAuthenticatorRequest
//
// @return DeregisterAuthenticatorResponse
func (client *Client) DeregisterAuthenticator(request *DeregisterAuthenticatorRequest) (_result *DeregisterAuthenticatorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeregisterAuthenticatorResponse{}
	_body, _err := client.DeregisterAuthenticatorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取access_token
//
// @param request - FetchAccessTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchAccessTokenResponse
func (client *Client) FetchAccessTokenWithOptions(request *FetchAccessTokenRequest, runtime *util.RuntimeOptions) (_result *FetchAccessTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationExternalId)) {
		query["ApplicationExternalId"] = request.ApplicationExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.MobileExtendParamsJson)) {
		query["MobileExtendParamsJson"] = request.MobileExtendParamsJson
	}

	if !tea.BoolValue(util.IsUnset(request.MobileExtendParamsJsonSign)) {
		query["MobileExtendParamsJsonSign"] = request.MobileExtendParamsJsonSign
	}

	if !tea.BoolValue(util.IsUnset(request.ServerExtendParamsJson)) {
		query["ServerExtendParamsJson"] = request.ServerExtendParamsJson
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.XClientIp)) {
		query["XClientIp"] = request.XClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FetchAccessToken"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FetchAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取access_token
//
// @param request - FetchAccessTokenRequest
//
// @return FetchAccessTokenResponse
func (client *Client) FetchAccessToken(request *FetchAccessTokenRequest) (_result *FetchAccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FetchAccessTokenResponse{}
	_body, _err := client.FetchAccessTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个认证器
//
// @param request - GetAuthenticatorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthenticatorResponse
func (client *Client) GetAuthenticatorWithOptions(request *GetAuthenticatorRequest, runtime *util.RuntimeOptions) (_result *GetAuthenticatorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationExternalId)) {
		query["ApplicationExternalId"] = request.ApplicationExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticatorUuid)) {
		query["AuthenticatorUuid"] = request.AuthenticatorUuid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAuthenticator"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuthenticatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个认证器
//
// @param request - GetAuthenticatorRequest
//
// @return GetAuthenticatorResponse
func (client *Client) GetAuthenticator(request *GetAuthenticatorRequest) (_result *GetAuthenticatorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuthenticatorResponse{}
	_body, _err := client.GetAuthenticatorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列表查询认证事件日志
//
// @param request - ListAuthenticationLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthenticationLogsResponse
func (client *Client) ListAuthenticationLogsWithOptions(request *ListAuthenticationLogsRequest, runtime *util.RuntimeOptions) (_result *ListAuthenticationLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationExternalId)) {
		query["ApplicationExternalId"] = request.ApplicationExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticatorType)) {
		query["AuthenticatorType"] = request.AuthenticatorType
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticatorUuid)) {
		query["AuthenticatorUuid"] = request.AuthenticatorUuid
	}

	if !tea.BoolValue(util.IsUnset(request.CredentialId)) {
		query["CredentialId"] = request.CredentialId
	}

	if !tea.BoolValue(util.IsUnset(request.FromTime)) {
		query["FromTime"] = request.FromTime
	}

	if !tea.BoolValue(util.IsUnset(request.LogTag)) {
		query["LogTag"] = request.LogTag
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ToTime)) {
		query["ToTime"] = request.ToTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAuthenticationLogs"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAuthenticationLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表查询认证事件日志
//
// @param request - ListAuthenticationLogsRequest
//
// @return ListAuthenticationLogsResponse
func (client *Client) ListAuthenticationLogs(request *ListAuthenticationLogsRequest) (_result *ListAuthenticationLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAuthenticationLogsResponse{}
	_body, _err := client.ListAuthenticationLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列表查询认证器操作日志
//
// @param request - ListAuthenticatorOpsLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthenticatorOpsLogsResponse
func (client *Client) ListAuthenticatorOpsLogsWithOptions(request *ListAuthenticatorOpsLogsRequest, runtime *util.RuntimeOptions) (_result *ListAuthenticatorOpsLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationExternalId)) {
		query["ApplicationExternalId"] = request.ApplicationExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticatorType)) {
		query["AuthenticatorType"] = request.AuthenticatorType
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticatorUuid)) {
		query["AuthenticatorUuid"] = request.AuthenticatorUuid
	}

	if !tea.BoolValue(util.IsUnset(request.FromTime)) {
		query["FromTime"] = request.FromTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ToTime)) {
		query["ToTime"] = request.ToTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAuthenticatorOpsLogs"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAuthenticatorOpsLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表查询认证器操作日志
//
// @param request - ListAuthenticatorOpsLogsRequest
//
// @return ListAuthenticatorOpsLogsResponse
func (client *Client) ListAuthenticatorOpsLogs(request *ListAuthenticatorOpsLogsRequest) (_result *ListAuthenticatorOpsLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAuthenticatorOpsLogsResponse{}
	_body, _err := client.ListAuthenticatorOpsLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列表查询认证器
//
// @param request - ListAuthenticatorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthenticatorsResponse
func (client *Client) ListAuthenticatorsWithOptions(request *ListAuthenticatorsRequest, runtime *util.RuntimeOptions) (_result *ListAuthenticatorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationExternalId)) {
		query["ApplicationExternalId"] = request.ApplicationExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticatorType)) {
		query["AuthenticatorType"] = request.AuthenticatorType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAuthenticators"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAuthenticatorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表查询认证器
//
// @param request - ListAuthenticatorsRequest
//
// @return ListAuthenticatorsResponse
func (client *Client) ListAuthenticators(request *ListAuthenticatorsRequest) (_result *ListAuthenticatorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAuthenticatorsResponse{}
	_body, _err := client.ListAuthenticatorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询按量计费订单列表
//
// @param request - ListCostUnitOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCostUnitOrdersResponse
func (client *Client) ListCostUnitOrdersWithOptions(request *ListCostUnitOrdersRequest, runtime *util.RuntimeOptions) (_result *ListCostUnitOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginDate)) {
		query["BeginDate"] = request.BeginDate
	}

	if !tea.BoolValue(util.IsUnset(request.FinalDate)) {
		query["FinalDate"] = request.FinalDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCostUnitOrders"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCostUnitOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询按量计费订单列表
//
// @param request - ListCostUnitOrdersRequest
//
// @return ListCostUnitOrdersResponse
func (client *Client) ListCostUnitOrders(request *ListCostUnitOrdersRequest) (_result *ListCostUnitOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCostUnitOrdersResponse{}
	_body, _err := client.ListCostUnitOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用量消费统计记录列表
//
// @param request - ListOrderConsumeStatisticRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrderConsumeStatisticRecordsResponse
func (client *Client) ListOrderConsumeStatisticRecordsWithOptions(request *ListOrderConsumeStatisticRecordsRequest, runtime *util.RuntimeOptions) (_result *ListOrderConsumeStatisticRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliOrderCode)) {
		query["AliOrderCode"] = request.AliOrderCode
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationExternalId)) {
		query["ApplicationExternalId"] = request.ApplicationExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.StatisticTimeMax)) {
		query["StatisticTimeMax"] = request.StatisticTimeMax
	}

	if !tea.BoolValue(util.IsUnset(request.StatisticTimeMin)) {
		query["StatisticTimeMin"] = request.StatisticTimeMin
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOrderConsumeStatisticRecords"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOrderConsumeStatisticRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用量消费统计记录列表
//
// @param request - ListOrderConsumeStatisticRecordsRequest
//
// @return ListOrderConsumeStatisticRecordsResponse
func (client *Client) ListOrderConsumeStatisticRecords(request *ListOrderConsumeStatisticRecordsRequest) (_result *ListOrderConsumeStatisticRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrderConsumeStatisticRecordsResponse{}
	_body, _err := client.ListOrderConsumeStatisticRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 弱密码检测
//
// @param request - ListPwnedPasswordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPwnedPasswordsResponse
func (client *Client) ListPwnedPasswordsWithOptions(request *ListPwnedPasswordsRequest, runtime *util.RuntimeOptions) (_result *ListPwnedPasswordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PrefixHexPasswordSha1Hash)) {
		query["PrefixHexPasswordSha1Hash"] = request.PrefixHexPasswordSha1Hash
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPwnedPasswords"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPwnedPasswordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 弱密码检测
//
// @param request - ListPwnedPasswordsRequest
//
// @return ListPwnedPasswordsResponse
func (client *Client) ListPwnedPasswords(request *ListPwnedPasswordsRequest) (_result *ListPwnedPasswordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPwnedPasswordsResponse{}
	_body, _err := client.ListPwnedPasswordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询应用用户
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationExternalId)) {
		query["ApplicationExternalId"] = request.ApplicationExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询应用用户
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 短信回执查询
//
// @param request - QuerySmsReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsReportsResponse
func (client *Client) QuerySmsReportsWithOptions(request *QuerySmsReportsRequest, runtime *util.RuntimeOptions) (_result *QuerySmsReportsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySmsReports"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySmsReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 短信回执查询
//
// @param request - QuerySmsReportsRequest
//
// @return QuerySmsReportsResponse
func (client *Client) QuerySmsReports(request *QuerySmsReportsRequest) (_result *QuerySmsReportsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySmsReportsResponse{}
	_body, _err := client.QuerySmsReportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 短信上行查询
//
// @param request - QuerySmsUpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsUpsResponse
func (client *Client) QuerySmsUpsWithOptions(runtime *util.RuntimeOptions) (_result *QuerySmsUpsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QuerySmsUps"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySmsUpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 短信上行查询
//
// @return QuerySmsUpsResponse
func (client *Client) QuerySmsUps() (_result *QuerySmsUpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySmsUpsResponse{}
	_body, _err := client.QuerySmsUpsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 注册认证器
//
// @param request - RegisterAuthenticatorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterAuthenticatorResponse
func (client *Client) RegisterAuthenticatorWithOptions(request *RegisterAuthenticatorRequest, runtime *util.RuntimeOptions) (_result *RegisterAuthenticatorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationExternalId)) {
		query["ApplicationExternalId"] = request.ApplicationExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticatorName)) {
		query["AuthenticatorName"] = request.AuthenticatorName
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticatorType)) {
		query["AuthenticatorType"] = request.AuthenticatorType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientExtendParamsJson)) {
		query["ClientExtendParamsJson"] = request.ClientExtendParamsJson
	}

	if !tea.BoolValue(util.IsUnset(request.ClientExtendParamsJsonSign)) {
		query["ClientExtendParamsJsonSign"] = request.ClientExtendParamsJsonSign
	}

	if !tea.BoolValue(util.IsUnset(request.LogParams)) {
		query["LogParams"] = request.LogParams
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrationContext)) {
		query["RegistrationContext"] = request.RegistrationContext
	}

	if !tea.BoolValue(util.IsUnset(request.RequireChallengeBase64)) {
		query["RequireChallengeBase64"] = request.RequireChallengeBase64
	}

	if !tea.BoolValue(util.IsUnset(request.ServerExtendParamsJson)) {
		query["ServerExtendParamsJson"] = request.ServerExtendParamsJson
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserSourceIp)) {
		query["UserSourceIp"] = request.UserSourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterAuthenticator"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterAuthenticatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 注册认证器
//
// @param request - RegisterAuthenticatorRequest
//
// @return RegisterAuthenticatorResponse
func (client *Client) RegisterAuthenticator(request *RegisterAuthenticatorRequest) (_result *RegisterAuthenticatorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterAuthenticatorResponse{}
	_body, _err := client.RegisterAuthenticatorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 认证接口
//
// @param request - ServiceInvokeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ServiceInvokeResponse
func (client *Client) ServiceInvokeWithOptions(request *ServiceInvokeRequest, runtime *util.RuntimeOptions) (_result *ServiceInvokeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationExternalId)) {
		query["ApplicationExternalId"] = request.ApplicationExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.DoraemonAction)) {
		query["DoraemonAction"] = request.DoraemonAction
	}

	if !tea.BoolValue(util.IsUnset(request.MobileExtendParamsJson)) {
		query["MobileExtendParamsJson"] = request.MobileExtendParamsJson
	}

	if !tea.BoolValue(util.IsUnset(request.MobileExtendParamsJsonSign)) {
		query["MobileExtendParamsJsonSign"] = request.MobileExtendParamsJsonSign
	}

	if !tea.BoolValue(util.IsUnset(request.ServerExtendParamsJson)) {
		query["ServerExtendParamsJson"] = request.ServerExtendParamsJson
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		query["ServiceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.TestModel)) {
		query["TestModel"] = request.TestModel
	}

	if !tea.BoolValue(util.IsUnset(request.XClientIp)) {
		query["XClientIp"] = request.XClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ServiceInvoke"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ServiceInvokeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 认证接口
//
// @param request - ServiceInvokeRequest
//
// @return ServiceInvokeResponse
func (client *Client) ServiceInvoke(request *ServiceInvokeRequest) (_result *ServiceInvokeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ServiceInvokeResponse{}
	_body, _err := client.ServiceInvokeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新认证器名字
//
// @param request - UpdateAuthenticatorAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuthenticatorAttributeResponse
func (client *Client) UpdateAuthenticatorAttributeWithOptions(request *UpdateAuthenticatorAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateAuthenticatorAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationExternalId)) {
		query["ApplicationExternalId"] = request.ApplicationExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticatorName)) {
		query["AuthenticatorName"] = request.AuthenticatorName
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticatorUuid)) {
		query["AuthenticatorUuid"] = request.AuthenticatorUuid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAuthenticatorAttribute"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAuthenticatorAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新认证器名字
//
// @param request - UpdateAuthenticatorAttributeRequest
//
// @return UpdateAuthenticatorAttributeResponse
func (client *Client) UpdateAuthenticatorAttribute(request *UpdateAuthenticatorAttributeRequest) (_result *UpdateAuthenticatorAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAuthenticatorAttributeResponse{}
	_body, _err := client.UpdateAuthenticatorAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证id_token
//
// @param request - VerifyIdTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyIdTokenResponse
func (client *Client) VerifyIdTokenWithOptions(request *VerifyIdTokenRequest, runtime *util.RuntimeOptions) (_result *VerifyIdTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationExternalId)) {
		query["ApplicationExternalId"] = request.ApplicationExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.JwtIdToken)) {
		query["JwtIdToken"] = request.JwtIdToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyIdToken"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyIdTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证id_token
//
// @param request - VerifyIdTokenRequest
//
// @return VerifyIdTokenResponse
func (client *Client) VerifyIdToken(request *VerifyIdTokenRequest) (_result *VerifyIdTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyIdTokenResponse{}
	_body, _err := client.VerifyIdTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 认证用户
//
// @param request - VerifyUserAuthenticationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyUserAuthenticationResponse
func (client *Client) VerifyUserAuthenticationWithOptions(request *VerifyUserAuthenticationRequest, runtime *util.RuntimeOptions) (_result *VerifyUserAuthenticationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationExternalId)) {
		query["ApplicationExternalId"] = request.ApplicationExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticationContext)) {
		query["AuthenticationContext"] = request.AuthenticationContext
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticatorType)) {
		query["AuthenticatorType"] = request.AuthenticatorType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientExtendParamsJson)) {
		query["ClientExtendParamsJson"] = request.ClientExtendParamsJson
	}

	if !tea.BoolValue(util.IsUnset(request.ClientExtendParamsJsonSign)) {
		query["ClientExtendParamsJsonSign"] = request.ClientExtendParamsJsonSign
	}

	if !tea.BoolValue(util.IsUnset(request.LogParams)) {
		query["LogParams"] = request.LogParams
	}

	if !tea.BoolValue(util.IsUnset(request.LogTag)) {
		query["LogTag"] = request.LogTag
	}

	if !tea.BoolValue(util.IsUnset(request.RequireBindHashBase64)) {
		query["RequireBindHashBase64"] = request.RequireBindHashBase64
	}

	if !tea.BoolValue(util.IsUnset(request.RequireChallengeBase64)) {
		query["RequireChallengeBase64"] = request.RequireChallengeBase64
	}

	if !tea.BoolValue(util.IsUnset(request.ServerExtendParamsJson)) {
		query["ServerExtendParamsJson"] = request.ServerExtendParamsJson
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserSourceIp)) {
		query["UserSourceIp"] = request.UserSourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyUserAuthentication"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyUserAuthenticationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 认证用户
//
// @param request - VerifyUserAuthenticationRequest
//
// @return VerifyUserAuthenticationResponse
func (client *Client) VerifyUserAuthentication(request *VerifyUserAuthenticationRequest) (_result *VerifyUserAuthenticationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyUserAuthenticationResponse{}
	_body, _err := client.VerifyUserAuthenticationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
