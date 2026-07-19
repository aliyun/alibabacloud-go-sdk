// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRbacRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListRbacRolesRequest
	GetBizId() *string
	SetMaxResults(v int32) *ListRbacRolesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRbacRolesRequest
	GetNextToken() *string
	SetOrderColumn(v string) *ListRbacRolesRequest
	GetOrderColumn() *string
	SetOrderType(v string) *ListRbacRolesRequest
	GetOrderType() *string
	SetPageNum(v int32) *ListRbacRolesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListRbacRolesRequest
	GetPageSize() *int32
}

type ListRbacRolesRequest struct {
	// example:
	//
	// WS20250801154628000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFh3Xqm+JgZ/U9Jyb7wdVr9LWk80Tghn5UZjbcWEVEderBcbVF+Y6PS0i8PpCL4PQZ3e0C9oEH0Asd4tJEuGtkl2WuKdiWZpEwadNydQdJPFM=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRbacRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRbacRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRbacRolesRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListRbacRolesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRbacRolesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRbacRolesRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListRbacRolesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListRbacRolesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListRbacRolesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRbacRolesRequest) SetBizId(v string) *ListRbacRolesRequest {
	s.BizId = &v
	return s
}

func (s *ListRbacRolesRequest) SetMaxResults(v int32) *ListRbacRolesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRbacRolesRequest) SetNextToken(v string) *ListRbacRolesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRbacRolesRequest) SetOrderColumn(v string) *ListRbacRolesRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListRbacRolesRequest) SetOrderType(v string) *ListRbacRolesRequest {
	s.OrderType = &v
	return s
}

func (s *ListRbacRolesRequest) SetPageNum(v int32) *ListRbacRolesRequest {
	s.PageNum = &v
	return s
}

func (s *ListRbacRolesRequest) SetPageSize(v int32) *ListRbacRolesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRbacRolesRequest) Validate() error {
	return dara.Validate(s)
}
