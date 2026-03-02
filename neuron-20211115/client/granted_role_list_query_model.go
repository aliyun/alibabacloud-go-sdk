// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantedRoleListQuery interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizerId(v string) *GrantedRoleListQuery
	GetAuthorizerId() *string
	SetAuthorizerType(v string) *GrantedRoleListQuery
	GetAuthorizerType() *string
	SetName(v string) *GrantedRoleListQuery
	GetName() *string
}

type GrantedRoleListQuery struct {
	AuthorizerId   *string `json:"authorizerId,omitempty" xml:"authorizerId,omitempty"`
	AuthorizerType *string `json:"authorizerType,omitempty" xml:"authorizerType,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GrantedRoleListQuery) String() string {
	return dara.Prettify(s)
}

func (s GrantedRoleListQuery) GoString() string {
	return s.String()
}

func (s *GrantedRoleListQuery) GetAuthorizerId() *string {
	return s.AuthorizerId
}

func (s *GrantedRoleListQuery) GetAuthorizerType() *string {
	return s.AuthorizerType
}

func (s *GrantedRoleListQuery) GetName() *string {
	return s.Name
}

func (s *GrantedRoleListQuery) SetAuthorizerId(v string) *GrantedRoleListQuery {
	s.AuthorizerId = &v
	return s
}

func (s *GrantedRoleListQuery) SetAuthorizerType(v string) *GrantedRoleListQuery {
	s.AuthorizerType = &v
	return s
}

func (s *GrantedRoleListQuery) SetName(v string) *GrantedRoleListQuery {
	s.Name = &v
	return s
}

func (s *GrantedRoleListQuery) Validate() error {
	return dara.Validate(s)
}
