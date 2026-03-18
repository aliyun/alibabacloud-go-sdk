// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationTokenId(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody
	GetAuthenticationTokenId() *string
	SetAuthenticationTokenType(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody
	GetAuthenticationTokenType() *string
	SetConsumerId(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody
	GetConsumerId() *string
	SetConsumerType(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody
	GetConsumerType() *string
	SetCreateTime(v int64) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody
	GetCreateTime() *int64
	SetCreatorId(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody
	GetCreatorId() *string
	SetCreatorType(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody
	GetCreatorType() *string
	SetCredentialProviderId(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody
	GetCredentialProviderId() *string
	SetExpirationTime(v int64) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody
	GetExpirationTime() *int64
	SetInstanceId(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody
	GetInstanceId() *string
	SetJwtContent(v *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBodyJwtContent) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody
	GetJwtContent() *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBodyJwtContent
	SetRevoked(v bool) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody
	GetRevoked() *bool
	SetUpdateTime(v int64) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody
	GetUpdateTime() *int64
}

type ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody struct {
	// example:
	//
	// atntkn_01kqflm0sxxx8nmdc1cb5dskxxxxx
	AuthenticationTokenId *string `json:"authenticationTokenId,omitempty" xml:"authenticationTokenId,omitempty"`
	// example:
	//
	// jwt
	AuthenticationTokenType *string `json:"authenticationTokenType,omitempty" xml:"authenticationTokenType,omitempty"`
	// example:
	//
	// test_jwt_subject
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// example:
	//
	// custom
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
	InstanceId *string                                                                `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	JwtContent *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBodyJwtContent `json:"jwtContent,omitempty" xml:"jwtContent,omitempty" type:"Struct"`
	// example:
	//
	// false
	Revoked *bool `json:"revoked,omitempty" xml:"revoked,omitempty"`
	// example:
	//
	// 1649830225000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) GetAuthenticationTokenId() *string {
	return s.AuthenticationTokenId
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) GetAuthenticationTokenType() *string {
	return s.AuthenticationTokenType
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) GetConsumerType() *string {
	return s.ConsumerType
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) GetCreatorType() *string {
	return s.CreatorType
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) GetCredentialProviderId() *string {
	return s.CredentialProviderId
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) GetJwtContent() *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBodyJwtContent {
	return s.JwtContent
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) GetRevoked() *bool {
	return s.Revoked
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) SetAuthenticationTokenId(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody {
	s.AuthenticationTokenId = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) SetAuthenticationTokenType(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody {
	s.AuthenticationTokenType = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) SetConsumerId(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody {
	s.ConsumerId = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) SetConsumerType(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody {
	s.ConsumerType = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) SetCreateTime(v int64) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody {
	s.CreateTime = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) SetCreatorId(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody {
	s.CreatorId = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) SetCreatorType(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody {
	s.CreatorType = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) SetCredentialProviderId(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody {
	s.CredentialProviderId = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) SetExpirationTime(v int64) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody {
	s.ExpirationTime = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) SetInstanceId(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) SetJwtContent(v *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBodyJwtContent) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody {
	s.JwtContent = v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) SetRevoked(v bool) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody {
	s.Revoked = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) SetUpdateTime(v int64) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBody) Validate() error {
	if s.JwtContent != nil {
		if err := s.JwtContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBodyJwtContent struct {
	// example:
	//
	// sk-Nx2vzxxxxxxxxxxxxxxxxx
	DerivedShortToken *string `json:"derivedShortToken,omitempty" xml:"derivedShortToken,omitempty"`
	// example:
	//
	// eyJhbGciOixxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	JwtValue *string `json:"jwtValue,omitempty" xml:"jwtValue,omitempty"`
}

func (s ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBodyJwtContent) String() string {
	return dara.Prettify(s)
}

func (s ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBodyJwtContent) GoString() string {
	return s.String()
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBodyJwtContent) GetDerivedShortToken() *string {
	return s.DerivedShortToken
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBodyJwtContent) GetJwtValue() *string {
	return s.JwtValue
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBodyJwtContent) SetDerivedShortToken(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBodyJwtContent {
	s.DerivedShortToken = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBodyJwtContent) SetJwtValue(v string) *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBodyJwtContent {
	s.JwtValue = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenByDerivedShortTokenResponseBodyJwtContent) Validate() error {
	return dara.Validate(s)
}
