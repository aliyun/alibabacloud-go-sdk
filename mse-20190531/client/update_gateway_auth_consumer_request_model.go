// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayAuthConsumerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayAuthConsumerRequest
	GetAcceptLanguage() *string
	SetDescription(v string) *UpdateGatewayAuthConsumerRequest
	GetDescription() *string
	SetEncodeType(v string) *UpdateGatewayAuthConsumerRequest
	GetEncodeType() *string
	SetGatewayUniqueId(v string) *UpdateGatewayAuthConsumerRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayAuthConsumerRequest
	GetId() *int64
	SetJwks(v string) *UpdateGatewayAuthConsumerRequest
	GetJwks() *string
	SetKeyName(v string) *UpdateGatewayAuthConsumerRequest
	GetKeyName() *string
	SetKeyValue(v string) *UpdateGatewayAuthConsumerRequest
	GetKeyValue() *string
	SetTokenName(v string) *UpdateGatewayAuthConsumerRequest
	GetTokenName() *string
	SetTokenPass(v bool) *UpdateGatewayAuthConsumerRequest
	GetTokenPass() *bool
	SetTokenPosition(v string) *UpdateGatewayAuthConsumerRequest
	GetTokenPosition() *string
	SetTokenPrefix(v string) *UpdateGatewayAuthConsumerRequest
	GetTokenPrefix() *string
}

type UpdateGatewayAuthConsumerRequest struct {
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
	// Description
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
	// gw-90392d768a3847a7b804c505254d****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the consumer.
	//
	// This parameter is required.
	//
	// example:
	//
	// 63
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
}

func (s UpdateGatewayAuthConsumerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayAuthConsumerRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayAuthConsumerRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayAuthConsumerRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGatewayAuthConsumerRequest) GetEncodeType() *string {
	return s.EncodeType
}

func (s *UpdateGatewayAuthConsumerRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayAuthConsumerRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayAuthConsumerRequest) GetJwks() *string {
	return s.Jwks
}

func (s *UpdateGatewayAuthConsumerRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *UpdateGatewayAuthConsumerRequest) GetKeyValue() *string {
	return s.KeyValue
}

func (s *UpdateGatewayAuthConsumerRequest) GetTokenName() *string {
	return s.TokenName
}

func (s *UpdateGatewayAuthConsumerRequest) GetTokenPass() *bool {
	return s.TokenPass
}

func (s *UpdateGatewayAuthConsumerRequest) GetTokenPosition() *string {
	return s.TokenPosition
}

func (s *UpdateGatewayAuthConsumerRequest) GetTokenPrefix() *string {
	return s.TokenPrefix
}

func (s *UpdateGatewayAuthConsumerRequest) SetAcceptLanguage(v string) *UpdateGatewayAuthConsumerRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayAuthConsumerRequest) SetDescription(v string) *UpdateGatewayAuthConsumerRequest {
	s.Description = &v
	return s
}

func (s *UpdateGatewayAuthConsumerRequest) SetEncodeType(v string) *UpdateGatewayAuthConsumerRequest {
	s.EncodeType = &v
	return s
}

func (s *UpdateGatewayAuthConsumerRequest) SetGatewayUniqueId(v string) *UpdateGatewayAuthConsumerRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayAuthConsumerRequest) SetId(v int64) *UpdateGatewayAuthConsumerRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayAuthConsumerRequest) SetJwks(v string) *UpdateGatewayAuthConsumerRequest {
	s.Jwks = &v
	return s
}

func (s *UpdateGatewayAuthConsumerRequest) SetKeyName(v string) *UpdateGatewayAuthConsumerRequest {
	s.KeyName = &v
	return s
}

func (s *UpdateGatewayAuthConsumerRequest) SetKeyValue(v string) *UpdateGatewayAuthConsumerRequest {
	s.KeyValue = &v
	return s
}

func (s *UpdateGatewayAuthConsumerRequest) SetTokenName(v string) *UpdateGatewayAuthConsumerRequest {
	s.TokenName = &v
	return s
}

func (s *UpdateGatewayAuthConsumerRequest) SetTokenPass(v bool) *UpdateGatewayAuthConsumerRequest {
	s.TokenPass = &v
	return s
}

func (s *UpdateGatewayAuthConsumerRequest) SetTokenPosition(v string) *UpdateGatewayAuthConsumerRequest {
	s.TokenPosition = &v
	return s
}

func (s *UpdateGatewayAuthConsumerRequest) SetTokenPrefix(v string) *UpdateGatewayAuthConsumerRequest {
	s.TokenPrefix = &v
	return s
}

func (s *UpdateGatewayAuthConsumerRequest) Validate() error {
	return dara.Validate(s)
}
