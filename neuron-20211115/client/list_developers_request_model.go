// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDevelopersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *ListDevelopersRequest
	GetAccountIds() []*string
	SetEnterpriseId(v int64) *ListDevelopersRequest
	GetEnterpriseId() *int64
	SetName(v string) *ListDevelopersRequest
	GetName() *string
	SetOrderBy(v string) *ListDevelopersRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListDevelopersRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListDevelopersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDevelopersRequest
	GetPageSize() *int32
	SetRoleId(v int64) *ListDevelopersRequest
	GetRoleId() *int64
}

type ListDevelopersRequest struct {
	AccountIds     []*string `json:"accountIds,omitempty" xml:"accountIds,omitempty" type:"Repeated"`
	EnterpriseId   *int64    `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	Name           *string   `json:"name,omitempty" xml:"name,omitempty"`
	OrderBy        *string   `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string   `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32    `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RoleId         *int64    `json:"roleId,omitempty" xml:"roleId,omitempty"`
}

func (s ListDevelopersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDevelopersRequest) GoString() string {
	return s.String()
}

func (s *ListDevelopersRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *ListDevelopersRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListDevelopersRequest) GetName() *string {
	return s.Name
}

func (s *ListDevelopersRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDevelopersRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListDevelopersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDevelopersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDevelopersRequest) GetRoleId() *int64 {
	return s.RoleId
}

func (s *ListDevelopersRequest) SetAccountIds(v []*string) *ListDevelopersRequest {
	s.AccountIds = v
	return s
}

func (s *ListDevelopersRequest) SetEnterpriseId(v int64) *ListDevelopersRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListDevelopersRequest) SetName(v string) *ListDevelopersRequest {
	s.Name = &v
	return s
}

func (s *ListDevelopersRequest) SetOrderBy(v string) *ListDevelopersRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDevelopersRequest) SetOrderDirection(v string) *ListDevelopersRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListDevelopersRequest) SetPageNumber(v int32) *ListDevelopersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDevelopersRequest) SetPageSize(v int32) *ListDevelopersRequest {
	s.PageSize = &v
	return s
}

func (s *ListDevelopersRequest) SetRoleId(v int64) *ListDevelopersRequest {
	s.RoleId = &v
	return s
}

func (s *ListDevelopersRequest) Validate() error {
	return dara.Validate(s)
}
