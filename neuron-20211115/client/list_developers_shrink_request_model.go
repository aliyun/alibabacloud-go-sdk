// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDevelopersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIdsShrink(v string) *ListDevelopersShrinkRequest
	GetAccountIdsShrink() *string
	SetEnterpriseId(v int64) *ListDevelopersShrinkRequest
	GetEnterpriseId() *int64
	SetName(v string) *ListDevelopersShrinkRequest
	GetName() *string
	SetOrderBy(v string) *ListDevelopersShrinkRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListDevelopersShrinkRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListDevelopersShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDevelopersShrinkRequest
	GetPageSize() *int32
	SetRoleId(v int64) *ListDevelopersShrinkRequest
	GetRoleId() *int64
}

type ListDevelopersShrinkRequest struct {
	AccountIdsShrink *string `json:"accountIds,omitempty" xml:"accountIds,omitempty"`
	EnterpriseId     *int64  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	OrderBy          *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection   *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber       *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize         *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RoleId           *int64  `json:"roleId,omitempty" xml:"roleId,omitempty"`
}

func (s ListDevelopersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDevelopersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDevelopersShrinkRequest) GetAccountIdsShrink() *string {
	return s.AccountIdsShrink
}

func (s *ListDevelopersShrinkRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListDevelopersShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListDevelopersShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDevelopersShrinkRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListDevelopersShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDevelopersShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDevelopersShrinkRequest) GetRoleId() *int64 {
	return s.RoleId
}

func (s *ListDevelopersShrinkRequest) SetAccountIdsShrink(v string) *ListDevelopersShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *ListDevelopersShrinkRequest) SetEnterpriseId(v int64) *ListDevelopersShrinkRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListDevelopersShrinkRequest) SetName(v string) *ListDevelopersShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListDevelopersShrinkRequest) SetOrderBy(v string) *ListDevelopersShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDevelopersShrinkRequest) SetOrderDirection(v string) *ListDevelopersShrinkRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListDevelopersShrinkRequest) SetPageNumber(v int32) *ListDevelopersShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDevelopersShrinkRequest) SetPageSize(v int32) *ListDevelopersShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDevelopersShrinkRequest) SetRoleId(v int64) *ListDevelopersShrinkRequest {
	s.RoleId = &v
	return s
}

func (s *ListDevelopersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
