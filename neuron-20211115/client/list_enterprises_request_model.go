// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterprisesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListEnterprisesRequest
	GetName() *string
	SetOrderBy(v string) *ListEnterprisesRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListEnterprisesRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListEnterprisesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEnterprisesRequest
	GetPageSize() *int32
}

type ListEnterprisesRequest struct {
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListEnterprisesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnterprisesRequest) GoString() string {
	return s.String()
}

func (s *ListEnterprisesRequest) GetName() *string {
	return s.Name
}

func (s *ListEnterprisesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListEnterprisesRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListEnterprisesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEnterprisesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEnterprisesRequest) SetName(v string) *ListEnterprisesRequest {
	s.Name = &v
	return s
}

func (s *ListEnterprisesRequest) SetOrderBy(v string) *ListEnterprisesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListEnterprisesRequest) SetOrderDirection(v string) *ListEnterprisesRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListEnterprisesRequest) SetPageNumber(v int32) *ListEnterprisesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEnterprisesRequest) SetPageSize(v int32) *ListEnterprisesRequest {
	s.PageSize = &v
	return s
}

func (s *ListEnterprisesRequest) Validate() error {
	return dara.Validate(s)
}
