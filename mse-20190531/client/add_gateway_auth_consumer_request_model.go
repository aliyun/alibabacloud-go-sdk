// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayAuthConsumerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddGatewayAuthConsumerRequest
	GetAcceptLanguage() *string
	SetDescription(v string) *AddGatewayAuthConsumerRequest
	GetDescription() *string
	SetEncodeType(v string) *AddGatewayAuthConsumerRequest
	GetEncodeType() *string
	SetGatewayUniqueId(v string) *AddGatewayAuthConsumerRequest
	GetGatewayUniqueId() *string
	SetJwks(v string) *AddGatewayAuthConsumerRequest
	GetJwks() *string
	SetKeyName(v string) *AddGatewayAuthConsumerRequest
	GetKeyName() *string
	SetKeyValue(v string) *AddGatewayAuthConsumerRequest
	GetKeyValue() *string
	SetName(v string) *AddGatewayAuthConsumerRequest
	GetName() *string
	SetTokenName(v string) *AddGatewayAuthConsumerRequest
	GetTokenName() *string
	SetTokenPass(v bool) *AddGatewayAuthConsumerRequest
	GetTokenPass() *bool
	SetTokenPosition(v string) *AddGatewayAuthConsumerRequest
	GetTokenPosition() *string
	SetTokenPrefix(v string) *AddGatewayAuthConsumerRequest
	GetTokenPrefix() *string
	SetType(v string) *AddGatewayAuthConsumerRequest
	GetType() *string
}

type AddGatewayAuthConsumerRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The description of the consumer.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The encryption type. Valid values:
	//
	// 	- RSA
	//
	// 	- OCT
	//
	// example:
	//
	// RSA
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// The unique ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-c70622ff52fe49beb29bea9a6f52****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The JWT public key. The JSON format is supported.
	//
	// example:
	//
	// {"keys":[{"e":"AQAB","kid":"DHFbpoIUqrY8t2zpA2qXfCmr5VO5ZEr4RzHU_-envvQ","kty":"RSA","n":"xAE7eB6qugXyCAG3yhh7pkDkT65pHymX-P7KfIupjf59vsdo91bSP9C8H07pSAGQO1MV_xFj9VswgsCg4R6otmg5PV2He95lZdHtOcU5DXIg_pbhLdKXbi66GlVeK6ABZOUW3WYtnNHD-91gVuoeJT_DwtGGcp4ignkgXfkiEm4sw-4sfb4qdt5oLbyVpmW6x9cfa7vs2WTfURiCrBoUqgBo_-4WTiULmmHSGZHOjzwa8WtrtOQGsAFjIbno85jp6MnGGGZPYZbDAa_b3y5u-YpW7ypZrvD8BgtKVjgtQgZhLAGezMt0ua3DRrWnKqTZ0BJ_EyxOGuHJrLsn00fnMQ"}]}
	Jwks *string `json:"Jwks,omitempty" xml:"Jwks,omitempty"`
	// The name of the key used for JWT-based identity authentication.
	//
	// example:
	//
	// iss
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The value of the key used for JWT-based identity authentication.
	//
	// example:
	//
	// abcd
	KeyValue *string `json:"KeyValue,omitempty" xml:"KeyValue,omitempty"`
	// The name of the consumer.
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The names of the parameters that are required to verify each token. By default, each token is prefixed with Bearer and stored in the Authorization header, such as `Authorization: Bearer <Content of a token>`.
	//
	// example:
	//
	// Authorization
	TokenName *string `json:"TokenName,omitempty" xml:"TokenName,omitempty"`
	// Specifies whether to enable pass-through.
	//
	// example:
	//
	// true
	TokenPass *bool `json:"TokenPass,omitempty" xml:"TokenPass,omitempty"`
	// The positions of the parameters that are required to verify each token. By default, each token is prefixed with Bearer and stored in the Authorization header, such as `Authorization: Bearer <Content of a token>`.
	//
	// example:
	//
	// HEADER
	TokenPosition *string `json:"TokenPosition,omitempty" xml:"TokenPosition,omitempty"`
	// The prefixes of the parameters that are required to verify each token. By default, each token is prefixed with Bearer and stored in the Authorization header, such as `Authorization: Bearer <Content of a token>`.
	//
	// example:
	//
	// Bearer
	TokenPrefix *string `json:"TokenPrefix,omitempty" xml:"TokenPrefix,omitempty"`
	// The authentication type. Valid values:
	//
	// 	- JWT
	//
	// This parameter is required.
	//
	// example:
	//
	// JWT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddGatewayAuthConsumerRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayAuthConsumerRequest) GoString() string {
	return s.String()
}

func (s *AddGatewayAuthConsumerRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddGatewayAuthConsumerRequest) GetDescription() *string {
	return s.Description
}

func (s *AddGatewayAuthConsumerRequest) GetEncodeType() *string {
	return s.EncodeType
}

func (s *AddGatewayAuthConsumerRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddGatewayAuthConsumerRequest) GetJwks() *string {
	return s.Jwks
}

func (s *AddGatewayAuthConsumerRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *AddGatewayAuthConsumerRequest) GetKeyValue() *string {
	return s.KeyValue
}

func (s *AddGatewayAuthConsumerRequest) GetName() *string {
	return s.Name
}

func (s *AddGatewayAuthConsumerRequest) GetTokenName() *string {
	return s.TokenName
}

func (s *AddGatewayAuthConsumerRequest) GetTokenPass() *bool {
	return s.TokenPass
}

func (s *AddGatewayAuthConsumerRequest) GetTokenPosition() *string {
	return s.TokenPosition
}

func (s *AddGatewayAuthConsumerRequest) GetTokenPrefix() *string {
	return s.TokenPrefix
}

func (s *AddGatewayAuthConsumerRequest) GetType() *string {
	return s.Type
}

func (s *AddGatewayAuthConsumerRequest) SetAcceptLanguage(v string) *AddGatewayAuthConsumerRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddGatewayAuthConsumerRequest) SetDescription(v string) *AddGatewayAuthConsumerRequest {
	s.Description = &v
	return s
}

func (s *AddGatewayAuthConsumerRequest) SetEncodeType(v string) *AddGatewayAuthConsumerRequest {
	s.EncodeType = &v
	return s
}

func (s *AddGatewayAuthConsumerRequest) SetGatewayUniqueId(v string) *AddGatewayAuthConsumerRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddGatewayAuthConsumerRequest) SetJwks(v string) *AddGatewayAuthConsumerRequest {
	s.Jwks = &v
	return s
}

func (s *AddGatewayAuthConsumerRequest) SetKeyName(v string) *AddGatewayAuthConsumerRequest {
	s.KeyName = &v
	return s
}

func (s *AddGatewayAuthConsumerRequest) SetKeyValue(v string) *AddGatewayAuthConsumerRequest {
	s.KeyValue = &v
	return s
}

func (s *AddGatewayAuthConsumerRequest) SetName(v string) *AddGatewayAuthConsumerRequest {
	s.Name = &v
	return s
}

func (s *AddGatewayAuthConsumerRequest) SetTokenName(v string) *AddGatewayAuthConsumerRequest {
	s.TokenName = &v
	return s
}

func (s *AddGatewayAuthConsumerRequest) SetTokenPass(v bool) *AddGatewayAuthConsumerRequest {
	s.TokenPass = &v
	return s
}

func (s *AddGatewayAuthConsumerRequest) SetTokenPosition(v string) *AddGatewayAuthConsumerRequest {
	s.TokenPosition = &v
	return s
}

func (s *AddGatewayAuthConsumerRequest) SetTokenPrefix(v string) *AddGatewayAuthConsumerRequest {
	s.TokenPrefix = &v
	return s
}

func (s *AddGatewayAuthConsumerRequest) SetType(v string) *AddGatewayAuthConsumerRequest {
	s.Type = &v
	return s
}

func (s *AddGatewayAuthConsumerRequest) Validate() error {
	return dara.Validate(s)
}
