// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainJwtAuthenticationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationTokenId(v string) *ObtainJwtAuthenticationTokenRequest
	GetAuthenticationTokenId() *string
	SetConsumerId(v string) *ObtainJwtAuthenticationTokenRequest
	GetConsumerId() *string
}

type ObtainJwtAuthenticationTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// atntkn_01kqflm0sxxx8nmdc1cb5dskxxxxx
	AuthenticationTokenId *string `json:"authenticationTokenId,omitempty" xml:"authenticationTokenId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_jwt_subject
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
}

func (s ObtainJwtAuthenticationTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ObtainJwtAuthenticationTokenRequest) GoString() string {
	return s.String()
}

func (s *ObtainJwtAuthenticationTokenRequest) GetAuthenticationTokenId() *string {
	return s.AuthenticationTokenId
}

func (s *ObtainJwtAuthenticationTokenRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ObtainJwtAuthenticationTokenRequest) SetAuthenticationTokenId(v string) *ObtainJwtAuthenticationTokenRequest {
	s.AuthenticationTokenId = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenRequest) SetConsumerId(v string) *ObtainJwtAuthenticationTokenRequest {
	s.ConsumerId = &v
	return s
}

func (s *ObtainJwtAuthenticationTokenRequest) Validate() error {
	return dara.Validate(s)
}
