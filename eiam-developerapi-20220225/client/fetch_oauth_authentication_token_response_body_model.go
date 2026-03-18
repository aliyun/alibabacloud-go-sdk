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
	SetRevoked(v bool) *FetchOAuthAuthenticationTokenResponseBody
	GetRevoked() *bool
	SetUpdateTime(v int64) *FetchOAuthAuthenticationTokenResponseBody
	GetUpdateTime() *int64
}

type FetchOAuthAuthenticationTokenResponseBody struct {
	// example:
	//
	// atntkn_01kqflm0sxxx8nmdc1cb5dskxxxxx
	AuthenticationTokenId *string `json:"authenticationTokenId,omitempty" xml:"authenticationTokenId,omitempty"`
	// example:
	//
	// oauth_access_token
	AuthenticationTokenType *string `json:"authenticationTokenType,omitempty" xml:"authenticationTokenType,omitempty"`
	// example:
	//
	// app_ngtkgrrxxxxktg5eao6z4xxxxx
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// example:
	//
	// application
	ConsumerType *string `json:"consumerType,omitempty" xml:"consumerType,omitempty"`
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// app_ngtkgrrxxxxktg5eao6z4xxxxx
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// application
	CreatorType *string `json:"creatorType,omitempty" xml:"creatorType,omitempty"`
	// example:
	//
	// atp_01kr2cmj5gxxx4fvmls2e93dxxxxx
	CredentialProviderId *string `json:"credentialProviderId,omitempty" xml:"credentialProviderId,omitempty"`
	// example:
	//
	// 1772693568000
	ExpirationTime *int64 `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	// EIAM实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId              *string                                                           `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	OauthAccessTokenContent *FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent `json:"oauthAccessTokenContent,omitempty" xml:"oauthAccessTokenContent,omitempty" type:"Struct"`
	// example:
	//
	// false
	Revoked *bool `json:"revoked,omitempty" xml:"revoked,omitempty"`
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
	return nil
}

type FetchOAuthAuthenticationTokenResponseBodyOauthAccessTokenContent struct {
	// example:
	//
	// DgEBAGP2xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AccessTokenValue *string `json:"accessTokenValue,omitempty" xml:"accessTokenValue,omitempty"`
	// example:
	//
	// example:test_01 example:test_02
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
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
