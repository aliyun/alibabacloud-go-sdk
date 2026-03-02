// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizationResource interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizerId(v string) *AuthorizationResource
	GetAuthorizerId() *string
	SetAuthorizerType(v string) *AuthorizationResource
	GetAuthorizerType() *string
}

type AuthorizationResource struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthorizerId *string `json:"authorizerId,omitempty" xml:"authorizerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// developer
	AuthorizerType *string `json:"authorizerType,omitempty" xml:"authorizerType,omitempty"`
}

func (s AuthorizationResource) String() string {
	return dara.Prettify(s)
}

func (s AuthorizationResource) GoString() string {
	return s.String()
}

func (s *AuthorizationResource) GetAuthorizerId() *string {
	return s.AuthorizerId
}

func (s *AuthorizationResource) GetAuthorizerType() *string {
	return s.AuthorizerType
}

func (s *AuthorizationResource) SetAuthorizerId(v string) *AuthorizationResource {
	s.AuthorizerId = &v
	return s
}

func (s *AuthorizationResource) SetAuthorizerType(v string) *AuthorizationResource {
	s.AuthorizerType = &v
	return s
}

func (s *AuthorizationResource) Validate() error {
	return dara.Validate(s)
}
