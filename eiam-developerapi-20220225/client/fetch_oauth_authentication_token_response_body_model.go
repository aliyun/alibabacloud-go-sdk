// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchOAuthAuthenticationTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationTokenId(v string) *FetchOAuthAuthenticationTokenResponseBody
	GetAuthenticationTokenId() *string
	SetAuthenticationTokenType(v string) *FetchOAuthAuthenticationTokenResponseBody
	GetAuthenticationTokenType() *string
	SetConsumerId(v string) *FetchOAuthAuthenticationTokenResponseBody
	GetConsumerId() *string
	SetConsumerType(v string) *FetchOAuthAuthenticationTokenResponseBody
	GetConsumerType() *string
	SetCreateTime(v int64) *FetchOAuthAuthenticationTokenResponseBody
	GetCreateTime() *int64
	SetCreatorId(v string) *FetchOAuthAuthenticationTokenResponseBody
	GetCreatorId() *string
	SetCreatorType(v string) *FetchOAuthAuthenticationTokenResponseBody
	GetCreatorType() *string
	SetCredentialProviderId(v string) *FetchOAuthAuthenticationTokenResponseBody
	GetCredentialProviderId() *string
	SetExpirationTime(v int64) *FetchOAuthAuthenticationTokenResponseBody
	GetExpirationTime() *int64
	SetInstanceId(v string) *FetchOAuthAuthenticationTokenResponseBody
	GetInstanceId() *string
	SetOauthAccessTokenContent(v *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent) *FetchOAuthAuthenticationTokenResponseBody
	GetOauthAccessTokenContent() *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent
	SetOauthAuthorizationSession(v *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession) *FetchOAuthAuthenticationTokenResponseBody
	GetOauthAuthorizationSession() *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession
	SetRevoked(v bool) *FetchOAuthAuthenticationTokenResponseBody
	GetRevoked() *bool
	SetUpdateTime(v int64) *FetchOAuthAuthenticationTokenResponseBody
	GetUpdateTime() *int64
}

type FetchOAuthAuthenticationTokenResponseBody struct {
	// The authentication token ID.
	//
	// example:
	//
	// atntkn_01kqflm0sxxx8nmdc1cb5dskxxxxx
	AuthenticationTokenId *string `json:"authenticationTokenId,omitempty" xml:"authenticationTokenId,omitempty"`
	// The authentication token type.
	//
	// > The value is fixed as `oauth_access_token`, indicating an OAuth Access Token type authentication token.
	//
	// example:
	//
	// oauth_access_token
	AuthenticationTokenType *string `json:"authenticationTokenType,omitempty" xml:"authenticationTokenType,omitempty"`
	// The consumer ID of the authentication token.
	//
	// example:
	//
	// app_ngtkgrrxxxxktg5eao6z4xxxxx
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The consumer type of the authentication token.
	//
	// example:
	//
	// application
	ConsumerType *string `json:"consumerType,omitempty" xml:"consumerType,omitempty"`
	// The creation time of the authentication token. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The creator ID of the authentication token.
	//
	// example:
	//
	// app_ngtkgrrxxxxktg5eao6z4xxxxx
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The creator type of the authentication token.
	//
	// example:
	//
	// application
	CreatorType *string `json:"creatorType,omitempty" xml:"creatorType,omitempty"`
	// The credential provider ID.
	//
	// example:
	//
	// atp_01kr2cmj5gxxx4fvmls2e93dxxxxx
	CredentialProviderId *string `json:"credentialProviderId,omitempty" xml:"credentialProviderId,omitempty"`
	// The expiration time of the authentication token. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1772693568000
	ExpirationTime *int64 `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The content of the OAuth Access Token type authentication token.
	OauthAccessTokenContent *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent `json:"oauthAccessTokenContent,omitempty" xml:"oauthAccessTokenContent,omitempty" type:"Struct"`
	// The authorization session of the OAuth user_federation flow. Returned during first-time authorization or when user interaction is required.
	OauthAuthorizationSession *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession `json:"oauthAuthorizationSession,omitempty" xml:"oauthAuthorizationSession,omitempty" type:"Struct"`
	// Indicates whether the authentication token is revoked.
	//
	// example:
	//
	// false
	Revoked *bool `json:"revoked,omitempty" xml:"revoked,omitempty"`
	// The update time of the authentication token. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830225000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s FetchOAuthAuthenticationTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FetchOAuthAuthenticationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *FetchOAuthAuthenticationTokenResponseBody) GetAuthenticationTokenId() *string {
	return s.AuthenticationTokenId
}

func (s *FetchOAuthAuthenticationTokenResponseBody) GetAuthenticationTokenType() *string {
	return s.AuthenticationTokenType
}

func (s *FetchOAuthAuthenticationTokenResponseBody) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *FetchOAuthAuthenticationTokenResponseBody) GetConsumerType() *string {
	return s.ConsumerType
}

func (s *FetchOAuthAuthenticationTokenResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *FetchOAuthAuthenticationTokenResponseBody) GetCreatorId() *string {
	return s.CreatorId
}

func (s *FetchOAuthAuthenticationTokenResponseBody) GetCreatorType() *string {
	return s.CreatorType
}

func (s *FetchOAuthAuthenticationTokenResponseBody) GetCredentialProviderId() *string {
	return s.CredentialProviderId
}

func (s *FetchOAuthAuthenticationTokenResponseBody) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *FetchOAuthAuthenticationTokenResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *FetchOAuthAuthenticationTokenResponseBody) GetOauthAccessTokenContent() *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent {
	return s.OauthAccessTokenContent
}

func (s *FetchOAuthAuthenticationTokenResponseBody) GetOauthAuthorizationSession() *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession {
	return s.OauthAuthorizationSession
}

func (s *FetchOAuthAuthenticationTokenResponseBody) GetRevoked() *bool {
	return s.Revoked
}

func (s *FetchOAuthAuthenticationTokenResponseBody) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *FetchOAuthAuthenticationTokenResponseBody) SetAuthenticationTokenId(v string) *FetchOAuthAuthenticationTokenResponseBody {
	s.AuthenticationTokenId = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBody) SetAuthenticationTokenType(v string) *FetchOAuthAuthenticationTokenResponseBody {
	s.AuthenticationTokenType = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBody) SetConsumerId(v string) *FetchOAuthAuthenticationTokenResponseBody {
	s.ConsumerId = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBody) SetConsumerType(v string) *FetchOAuthAuthenticationTokenResponseBody {
	s.ConsumerType = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBody) SetCreateTime(v int64) *FetchOAuthAuthenticationTokenResponseBody {
	s.CreateTime = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBody) SetCreatorId(v string) *FetchOAuthAuthenticationTokenResponseBody {
	s.CreatorId = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBody) SetCreatorType(v string) *FetchOAuthAuthenticationTokenResponseBody {
	s.CreatorType = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBody) SetCredentialProviderId(v string) *FetchOAuthAuthenticationTokenResponseBody {
	s.CredentialProviderId = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBody) SetExpirationTime(v int64) *FetchOAuthAuthenticationTokenResponseBody {
	s.ExpirationTime = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBody) SetInstanceId(v string) *FetchOAuthAuthenticationTokenResponseBody {
	s.InstanceId = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBody) SetOauthAccessTokenContent(v *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent) *FetchOAuthAuthenticationTokenResponseBody {
	s.OauthAccessTokenContent = v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBody) SetOauthAuthorizationSession(v *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession) *FetchOAuthAuthenticationTokenResponseBody {
	s.OauthAuthorizationSession = v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBody) SetRevoked(v bool) *FetchOAuthAuthenticationTokenResponseBody {
	s.Revoked = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBody) SetUpdateTime(v int64) *FetchOAuthAuthenticationTokenResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBody) Validate() error {
	if s.OauthAccessTokenContent != nil {
		if err := s.OauthAccessTokenContent.Validate(); err != nil {
			return err
		}
	}
	if s.OauthAuthorizationSession != nil {
		if err := s.OauthAuthorizationSession.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent struct {
	// The access_token field in the OAuth protocol token endpoint response.
	//
	// example:
	//
	// DgEBAGP2xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AccessTokenValue *string `json:"accessTokenValue,omitempty" xml:"accessTokenValue,omitempty"`
	// The scope field in the OAuth protocol token endpoint response.
	//
	// example:
	//
	// example:test_01 example:test_02
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The token_type field in the OAuth protocol token endpoint response.
	//
	// example:
	//
	// Bearer
	TokenType *string `json:"tokenType,omitempty" xml:"tokenType,omitempty"`
}

func (s FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent) String() string {
	return dara.Prettify(s)
}

func (s FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent) GoString() string {
	return s.String()
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent) GetAccessTokenValue() *string {
	return s.AccessTokenValue
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent) GetScope() *string {
	return s.Scope
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent) GetTokenType() *string {
	return s.TokenType
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent) SetAccessTokenValue(v string) *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent {
	s.AccessTokenValue = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent) SetScope(v string) *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent {
	s.Scope = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent) SetTokenType(v string) *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent {
	s.TokenType = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent) Validate() error {
	return dara.Validate(s)
}

type FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession struct {
	// The user authorization URL.
	//
	// example:
	//
	// https://login.dingtalk.com/oauth2/auth?client_id=...
	AuthorizationUrl *string `json:"authorizationUrl,omitempty" xml:"authorizationUrl,omitempty"`
	// The authorization session ID.
	//
	// example:
	//
	// atpoas_01l6losojlojbbv01adsq56xxxxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The authorization session status.
	//
	// example:
	//
	// pending
	SessionStatus *string `json:"sessionStatus,omitempty" xml:"sessionStatus,omitempty"`
	// The authorization session URI.
	//
	// example:
	//
	// urn:ietf:params:oauth:request_uri:atpoas_01l6ljnvrpc5niakl3gj3amxxxxxx
	SessionUri *string `json:"sessionUri,omitempty" xml:"sessionUri,omitempty"`
}

func (s FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession) String() string {
	return dara.Prettify(s)
}

func (s FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession) GoString() string {
	return s.String()
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession) GetAuthorizationUrl() *string {
	return s.AuthorizationUrl
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession) GetSessionId() *string {
	return s.SessionId
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession) GetSessionUri() *string {
	return s.SessionUri
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession) SetAuthorizationUrl(v string) *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession {
	s.AuthorizationUrl = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession) SetSessionId(v string) *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession {
	s.SessionId = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession) SetSessionStatus(v string) *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession {
	s.SessionStatus = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession) SetSessionUri(v string) *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession {
	s.SessionUri = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenResponseBodyOauthAuthorizationSession) Validate() error {
	return dara.Validate(s)
}
