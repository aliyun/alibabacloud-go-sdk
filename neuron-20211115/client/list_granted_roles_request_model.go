// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGrantedRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizerId(v string) *ListGrantedRolesRequest
	GetAuthorizerId() *string
	SetAuthorizerType(v string) *ListGrantedRolesRequest
	GetAuthorizerType() *string
	SetEnterpriseId(v int64) *ListGrantedRolesRequest
	GetEnterpriseId() *int64
	SetName(v string) *ListGrantedRolesRequest
	GetName() *string
}

type ListGrantedRolesRequest struct {
	AuthorizerId   *string `json:"authorizerId,omitempty" xml:"authorizerId,omitempty"`
	AuthorizerType *string `json:"authorizerType,omitempty" xml:"authorizerType,omitempty"`
	EnterpriseId   *int64  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListGrantedRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGrantedRolesRequest) GoString() string {
	return s.String()
}

func (s *ListGrantedRolesRequest) GetAuthorizerId() *string {
	return s.AuthorizerId
}

func (s *ListGrantedRolesRequest) GetAuthorizerType() *string {
	return s.AuthorizerType
}

func (s *ListGrantedRolesRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListGrantedRolesRequest) GetName() *string {
	return s.Name
}

func (s *ListGrantedRolesRequest) SetAuthorizerId(v string) *ListGrantedRolesRequest {
	s.AuthorizerId = &v
	return s
}

func (s *ListGrantedRolesRequest) SetAuthorizerType(v string) *ListGrantedRolesRequest {
	s.AuthorizerType = &v
	return s
}

func (s *ListGrantedRolesRequest) SetEnterpriseId(v int64) *ListGrantedRolesRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListGrantedRolesRequest) SetName(v string) *ListGrantedRolesRequest {
	s.Name = &v
	return s
}

func (s *ListGrantedRolesRequest) Validate() error {
	return dara.Validate(s)
}
