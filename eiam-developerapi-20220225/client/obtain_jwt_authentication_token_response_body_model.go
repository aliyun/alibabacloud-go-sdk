// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainJwtAuthenticationTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationTokenId(v string) *ObtainJwtAuthenticationTokenResponseBody
	GetAuthenticationTokenId() *string
	SetAuthenticationTokenType(v string) *ObtainJwtAuthenticationTokenResponseBody
	GetAuthenticationTokenType() *string
	SetConsumerId(v string) *ObtainJwtAuthenticationTokenResponseBody
	GetConsumerId() *string
	SetConsumerType(v string) *ObtainJwtAuthenticationTokenResponseBody
	GetConsumerType() *string
	SetCreateTime(v int64) *ObtainJwtAuthenticationTokenResponseBody
	GetCreateTime() *int64
	SetCreatorId(v string) *ObtainJwtAuthenticationTokenResponseBody
	GetCreatorId() *string
	SetCreatorType(v string) *ObtainJwtAuthenticationTokenResponseBody
	GetCreatorType() *string
	SetCredentialProviderId(v string) *ObtainJwtAuthenticationTokenResponseBody
	GetCredentialProviderId() *string
	SetExpirationTime(v int64) *ObtainJwtAuthenticationTokenResponseBody
	GetExpirationTime() *int64
	SetInstanceId(v string) *ObtainJwtAuthenticationTokenResponseBody
	GetInstanceId() *string
	SetJwtContent(v *ObtainJwtAuthenticationTokenResponseBodyJwtContent) *ObtainJwtAuthenticationTokenResponseBody
	GetJwtContent() *ObtainJwtAuthenticationTokenResponseBodyJwtContent
	SetRevoked(v bool) *ObtainJwtAuthenticationTokenResponseBody
	GetRevoked() *bool
	SetUpdateTime(v int64) *ObtainJwtAuthenticationTokenResponseBody
	GetUpdateTime() *int64
}

type ObtainJwtAuthenticationTokenResponseBody struct {
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
	InstanceId *string                                             `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	JwtContent *ObtainJwtAuthenticationTokenResponseBodyJwtContent `json:"jwtContent,omitempty" xml:"jwtContent,omitempty" type:"Struct"`
	// example:
	//
	// false
	Revoked *bool `json:"revoked,omitempty" xml:"revoked,omitempty"`
	// example:
	//
	// 1649830225000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ObtainJwtAuthenticationTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ObtainJwtAuthenticationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ObtainJwtAuthenticationTokenResponseBody) GetAuthenticationTokenId() *string {
	return s.AuthenticationTokenId
}

func (s *ObtainJwtAuthenticationTokenResponseBody) GetAuthenticationTokenType() *string {
	return s.AuthenticationTokenType
}

func (s *ObtainJwtAuthenticationTokenResponseBody) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ObtainJwtAuthenticationTokenResponseBody) GetConsumerType() *string {
	return s.ConsumerType
}

func (s *ObtainJwtAuthenticationTokenResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ObtainJwtAuthenticationTokenResponseBody) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ObtainJwtAuthenticationTokenResponseBody) GetCreatorType() *string {
	return s.CreatorType
}

func (s *ObtainJwtAuthenticationTokenResponseBody) GetCredentialProviderId() *string {
	return s.CredentialProviderId
}

func (s *ObtainJwtAuthenticationTokenResponseBody) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *ObtainJwtAuthenticationTokenResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ObtainJwtAuthenticationTokenResponseBody) GetJwtContent() *ObtainJwtAuthenticationTokenResponseBodyJwtContent {
	return s.JwtContent
}

func (s *ObtainJwtAuthenticationTokenResponseBody) GetRevoked() *bool {
	return s.Revoked
}

func (s *ObtainJwtAuthenticationTokenResponseBody) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ObtainJwtAuthenticationTokenResponseBody) SetAuthenticationTokenId(v string) *ObtainJwtAuthenticationTokenResponseBody {
	s.AuthenticationTokenId = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponseBody) SetAuthenticationTokenType(v string) *ObtainJwtAuthenticationTokenResponseBody {
	s.AuthenticationTokenType = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponseBody) SetConsumerId(v string) *ObtainJwtAuthenticationTokenResponseBody {
	s.ConsumerId = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponseBody) SetConsumerType(v string) *ObtainJwtAuthenticationTokenResponseBody {
	s.ConsumerType = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponseBody) SetCreateTime(v int64) *ObtainJwtAuthenticationTokenResponseBody {
	s.CreateTime = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponseBody) SetCreatorId(v string) *ObtainJwtAuthenticationTokenResponseBody {
	s.CreatorId = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponseBody) SetCreatorType(v string) *ObtainJwtAuthenticationTokenResponseBody {
	s.CreatorType = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponseBody) SetCredentialProviderId(v string) *ObtainJwtAuthenticationTokenResponseBody {
	s.CredentialProviderId = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponseBody) SetExpirationTime(v int64) *ObtainJwtAuthenticationTokenResponseBody {
	s.ExpirationTime = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponseBody) SetInstanceId(v string) *ObtainJwtAuthenticationTokenResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponseBody) SetJwtContent(v *ObtainJwtAuthenticationTokenResponseBodyJwtContent) *ObtainJwtAuthenticationTokenResponseBody {
	s.JwtContent = v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponseBody) SetRevoked(v bool) *ObtainJwtAuthenticationTokenResponseBody {
	s.Revoked = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponseBody) SetUpdateTime(v int64) *ObtainJwtAuthenticationTokenResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponseBody) Validate() error {
	if s.JwtContent != nil {
		if err := s.JwtContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObtainJwtAuthenticationTokenResponseBodyJwtContent struct {
	// example:
	//
	// sk-Nx2vzxxxxxxxxxxxxxxxxx
	DerivedShortToken *string `json:"derivedShortToken,omitempty" xml:"derivedShortToken,omitempty"`
	// example:
	//
	// eyJhbGciOixxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	JwtValue *string `json:"jwtValue,omitempty" xml:"jwtValue,omitempty"`
}

func (s ObtainJwtAuthenticationTokenResponseBodyJwtContent) String() string {
	return dara.Prettify(s)
}

func (s ObtainJwtAuthenticationTokenResponseBodyJwtContent) GoString() string {
	return s.String()
}

func (s *ObtainJwtAuthenticationTokenResponseBodyJwtContent) GetDerivedShortToken() *string {
	return s.DerivedShortToken
}

func (s *ObtainJwtAuthenticationTokenResponseBodyJwtContent) GetJwtValue() *string {
	return s.JwtValue
}

func (s *ObtainJwtAuthenticationTokenResponseBodyJwtContent) SetDerivedShortToken(v string) *ObtainJwtAuthenticationTokenResponseBodyJwtContent {
	s.DerivedShortToken = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponseBodyJwtContent) SetJwtValue(v string) *ObtainJwtAuthenticationTokenResponseBodyJwtContent {
	s.JwtValue = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenResponseBodyJwtContent) Validate() error {
	return dara.Validate(s)
}
