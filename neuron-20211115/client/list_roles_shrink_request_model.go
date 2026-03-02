// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizerId(v string) *ListRolesShrinkRequest
	GetAuthorizerId() *string
	SetAuthorizerType(v string) *ListRolesShrinkRequest
	GetAuthorizerType() *string
	SetEnterpriseId(v int64) *ListRolesShrinkRequest
	GetEnterpriseId() *int64
	SetName(v string) *ListRolesShrinkRequest
	GetName() *string
	SetOrderBy(v string) *ListRolesShrinkRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListRolesShrinkRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListRolesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRolesShrinkRequest
	GetPageSize() *int32
	SetPlatform(v string) *ListRolesShrinkRequest
	GetPlatform() *string
	SetRoleIdsShrink(v string) *ListRolesShrinkRequest
	GetRoleIdsShrink() *string
	SetRoleType(v string) *ListRolesShrinkRequest
	GetRoleType() *string
}

type ListRolesShrinkRequest struct {
	AuthorizerId   *string `json:"authorizerId,omitempty" xml:"authorizerId,omitempty"`
	AuthorizerType *string `json:"authorizerType,omitempty" xml:"authorizerType,omitempty"`
	EnterpriseId   *int64  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Platform       *string `json:"platform,omitempty" xml:"platform,omitempty"`
	RoleIdsShrink  *string `json:"roleIds,omitempty" xml:"roleIds,omitempty"`
	RoleType       *string `json:"roleType,omitempty" xml:"roleType,omitempty"`
}

func (s ListRolesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRolesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListRolesShrinkRequest) GetAuthorizerId() *string {
	return s.AuthorizerId
}

func (s *ListRolesShrinkRequest) GetAuthorizerType() *string {
	return s.AuthorizerType
}

func (s *ListRolesShrinkRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListRolesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListRolesShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListRolesShrinkRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListRolesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRolesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRolesShrinkRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ListRolesShrinkRequest) GetRoleIdsShrink() *string {
	return s.RoleIdsShrink
}

func (s *ListRolesShrinkRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *ListRolesShrinkRequest) SetAuthorizerId(v string) *ListRolesShrinkRequest {
	s.AuthorizerId = &v
	return s
}

func (s *ListRolesShrinkRequest) SetAuthorizerType(v string) *ListRolesShrinkRequest {
	s.AuthorizerType = &v
	return s
}

func (s *ListRolesShrinkRequest) SetEnterpriseId(v int64) *ListRolesShrinkRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListRolesShrinkRequest) SetName(v string) *ListRolesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListRolesShrinkRequest) SetOrderBy(v string) *ListRolesShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListRolesShrinkRequest) SetOrderDirection(v string) *ListRolesShrinkRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListRolesShrinkRequest) SetPageNumber(v int32) *ListRolesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRolesShrinkRequest) SetPageSize(v int32) *ListRolesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListRolesShrinkRequest) SetPlatform(v string) *ListRolesShrinkRequest {
	s.Platform = &v
	return s
}

func (s *ListRolesShrinkRequest) SetRoleIdsShrink(v string) *ListRolesShrinkRequest {
	s.RoleIdsShrink = &v
	return s
}

func (s *ListRolesShrinkRequest) SetRoleType(v string) *ListRolesShrinkRequest {
	s.RoleType = &v
	return s
}

func (s *ListRolesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
