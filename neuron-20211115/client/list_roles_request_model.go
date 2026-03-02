// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizerId(v string) *ListRolesRequest
	GetAuthorizerId() *string
	SetAuthorizerType(v string) *ListRolesRequest
	GetAuthorizerType() *string
	SetEnterpriseId(v int64) *ListRolesRequest
	GetEnterpriseId() *int64
	SetName(v string) *ListRolesRequest
	GetName() *string
	SetOrderBy(v string) *ListRolesRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListRolesRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListRolesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRolesRequest
	GetPageSize() *int32
	SetPlatform(v string) *ListRolesRequest
	GetPlatform() *string
	SetRoleIds(v []*int64) *ListRolesRequest
	GetRoleIds() []*int64
	SetRoleType(v string) *ListRolesRequest
	GetRoleType() *string
}

type ListRolesRequest struct {
	AuthorizerId   *string  `json:"authorizerId,omitempty" xml:"authorizerId,omitempty"`
	AuthorizerType *string  `json:"authorizerType,omitempty" xml:"authorizerType,omitempty"`
	EnterpriseId   *int64   `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	Name           *string  `json:"name,omitempty" xml:"name,omitempty"`
	OrderBy        *string  `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string  `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Platform       *string  `json:"platform,omitempty" xml:"platform,omitempty"`
	RoleIds        []*int64 `json:"roleIds,omitempty" xml:"roleIds,omitempty" type:"Repeated"`
	RoleType       *string  `json:"roleType,omitempty" xml:"roleType,omitempty"`
}

func (s ListRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) GetAuthorizerId() *string {
	return s.AuthorizerId
}

func (s *ListRolesRequest) GetAuthorizerType() *string {
	return s.AuthorizerType
}

func (s *ListRolesRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListRolesRequest) GetName() *string {
	return s.Name
}

func (s *ListRolesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListRolesRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListRolesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRolesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRolesRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ListRolesRequest) GetRoleIds() []*int64 {
	return s.RoleIds
}

func (s *ListRolesRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *ListRolesRequest) SetAuthorizerId(v string) *ListRolesRequest {
	s.AuthorizerId = &v
	return s
}

func (s *ListRolesRequest) SetAuthorizerType(v string) *ListRolesRequest {
	s.AuthorizerType = &v
	return s
}

func (s *ListRolesRequest) SetEnterpriseId(v int64) *ListRolesRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListRolesRequest) SetName(v string) *ListRolesRequest {
	s.Name = &v
	return s
}

func (s *ListRolesRequest) SetOrderBy(v string) *ListRolesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListRolesRequest) SetOrderDirection(v string) *ListRolesRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListRolesRequest) SetPageNumber(v int32) *ListRolesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRolesRequest) SetPageSize(v int32) *ListRolesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRolesRequest) SetPlatform(v string) *ListRolesRequest {
	s.Platform = &v
	return s
}

func (s *ListRolesRequest) SetRoleIds(v []*int64) *ListRolesRequest {
	s.RoleIds = v
	return s
}

func (s *ListRolesRequest) SetRoleType(v string) *ListRolesRequest {
	s.RoleType = &v
	return s
}

func (s *ListRolesRequest) Validate() error {
	return dara.Validate(s)
}
