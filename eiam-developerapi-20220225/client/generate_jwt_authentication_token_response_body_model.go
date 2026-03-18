// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateJwtAuthenticationTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationTokenId(v string) *GenerateJwtAuthenticationTokenResponseBody
	GetAuthenticationTokenId() *string
	SetAuthenticationTokenType(v string) *GenerateJwtAuthenticationTokenResponseBody
	GetAuthenticationTokenType() *string
	SetConsumerId(v string) *GenerateJwtAuthenticationTokenResponseBody
	GetConsumerId() *string
	SetConsumerType(v string) *GenerateJwtAuthenticationTokenResponseBody
	GetConsumerType() *string
	SetCreateTime(v int64) *GenerateJwtAuthenticationTokenResponseBody
	GetCreateTime() *int64
	SetCreatorId(v string) *GenerateJwtAuthenticationTokenResponseBody
	GetCreatorId() *string
	SetCreatorType(v string) *GenerateJwtAuthenticationTokenResponseBody
	GetCreatorType() *string
	SetCredentialProviderId(v string) *GenerateJwtAuthenticationTokenResponseBody
	GetCredentialProviderId() *string
	SetExpirationTime(v int64) *GenerateJwtAuthenticationTokenResponseBody
	GetExpirationTime() *int64
	SetInstanceId(v string) *GenerateJwtAuthenticationTokenResponseBody
	GetInstanceId() *string
	SetJwtContent(v *GenerateJwtAuthenticationTokenResponseBodyJwtContent) *GenerateJwtAuthenticationTokenResponseBody
	GetJwtContent() *GenerateJwtAuthenticationTokenResponseBodyJwtContent
	SetRevoked(v bool) *GenerateJwtAuthenticationTokenResponseBody
	GetRevoked() *bool
	SetUpdateTime(v int64) *GenerateJwtAuthenticationTokenResponseBody
	GetUpdateTime() *int64
}

type GenerateJwtAuthenticationTokenResponseBody struct {
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
	InstanceId *string                                               `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	JwtContent *GenerateJwtAuthenticationTokenResponseBodyJwtContent `json:"jwtContent,omitempty" xml:"jwtContent,omitempty" type:"Struct"`
	// example:
	//
	// false
	Revoked *bool `json:"revoked,omitempty" xml:"revoked,omitempty"`
	// example:
	//
	// 1649830225000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GenerateJwtAuthenticationTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateJwtAuthenticationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateJwtAuthenticationTokenResponseBody) GetAuthenticationTokenId() *string {
	return s.AuthenticationTokenId
}

func (s *GenerateJwtAuthenticationTokenResponseBody) GetAuthenticationTokenType() *string {
	return s.AuthenticationTokenType
}

func (s *GenerateJwtAuthenticationTokenResponseBody) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *GenerateJwtAuthenticationTokenResponseBody) GetConsumerType() *string {
	return s.ConsumerType
}

func (s *GenerateJwtAuthenticationTokenResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GenerateJwtAuthenticationTokenResponseBody) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GenerateJwtAuthenticationTokenResponseBody) GetCreatorType() *string {
	return s.CreatorType
}

func (s *GenerateJwtAuthenticationTokenResponseBody) GetCredentialProviderId() *string {
	return s.CredentialProviderId
}

func (s *GenerateJwtAuthenticationTokenResponseBody) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *GenerateJwtAuthenticationTokenResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateJwtAuthenticationTokenResponseBody) GetJwtContent() *GenerateJwtAuthenticationTokenResponseBodyJwtContent {
	return s.JwtContent
}

func (s *GenerateJwtAuthenticationTokenResponseBody) GetRevoked() *bool {
	return s.Revoked
}

func (s *GenerateJwtAuthenticationTokenResponseBody) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GenerateJwtAuthenticationTokenResponseBody) SetAuthenticationTokenId(v string) *GenerateJwtAuthenticationTokenResponseBody {
	s.AuthenticationTokenId = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponseBody) SetAuthenticationTokenType(v string) *GenerateJwtAuthenticationTokenResponseBody {
	s.AuthenticationTokenType = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponseBody) SetConsumerId(v string) *GenerateJwtAuthenticationTokenResponseBody {
	s.ConsumerId = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponseBody) SetConsumerType(v string) *GenerateJwtAuthenticationTokenResponseBody {
	s.ConsumerType = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponseBody) SetCreateTime(v int64) *GenerateJwtAuthenticationTokenResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponseBody) SetCreatorId(v string) *GenerateJwtAuthenticationTokenResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponseBody) SetCreatorType(v string) *GenerateJwtAuthenticationTokenResponseBody {
	s.CreatorType = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponseBody) SetCredentialProviderId(v string) *GenerateJwtAuthenticationTokenResponseBody {
	s.CredentialProviderId = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponseBody) SetExpirationTime(v int64) *GenerateJwtAuthenticationTokenResponseBody {
	s.ExpirationTime = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponseBody) SetInstanceId(v string) *GenerateJwtAuthenticationTokenResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponseBody) SetJwtContent(v *GenerateJwtAuthenticationTokenResponseBodyJwtContent) *GenerateJwtAuthenticationTokenResponseBody {
	s.JwtContent = v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponseBody) SetRevoked(v bool) *GenerateJwtAuthenticationTokenResponseBody {
	s.Revoked = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponseBody) SetUpdateTime(v int64) *GenerateJwtAuthenticationTokenResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponseBody) Validate() error {
	if s.JwtContent != nil {
		if err := s.JwtContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateJwtAuthenticationTokenResponseBodyJwtContent struct {
	// example:
	//
	// sk-Nx2vzxxxxxxxxxxxxxxxxx
	DerivedShortToken *string `json:"derivedShortToken,omitempty" xml:"derivedShortToken,omitempty"`
	// example:
	//
	// eyJhbGciOixxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	JwtValue *string `json:"jwtValue,omitempty" xml:"jwtValue,omitempty"`
}

func (s GenerateJwtAuthenticationTokenResponseBodyJwtContent) String() string {
	return dara.Prettify(s)
}

func (s GenerateJwtAuthenticationTokenResponseBodyJwtContent) GoString() string {
	return s.String()
}

func (s *GenerateJwtAuthenticationTokenResponseBodyJwtContent) GetDerivedShortToken() *string {
	return s.DerivedShortToken
}

func (s *GenerateJwtAuthenticationTokenResponseBodyJwtContent) GetJwtValue() *string {
	return s.JwtValue
}

func (s *GenerateJwtAuthenticationTokenResponseBodyJwtContent) SetDerivedShortToken(v string) *GenerateJwtAuthenticationTokenResponseBodyJwtContent {
	s.DerivedShortToken = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponseBodyJwtContent) SetJwtValue(v string) *GenerateJwtAuthenticationTokenResponseBodyJwtContent {
	s.JwtValue = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenResponseBodyJwtContent) Validate() error {
	return dara.Validate(s)
}
