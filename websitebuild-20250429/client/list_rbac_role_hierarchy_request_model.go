// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRbacRoleHierarchyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListRbacRoleHierarchyRequest
	GetBizId() *string
	SetMaxResults(v int32) *ListRbacRoleHierarchyRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRbacRoleHierarchyRequest
	GetNextToken() *string
	SetOrderColumn(v string) *ListRbacRoleHierarchyRequest
	GetOrderColumn() *string
	SetOrderType(v string) *ListRbacRoleHierarchyRequest
	GetOrderType() *string
	SetPageNum(v int32) *ListRbacRoleHierarchyRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListRbacRoleHierarchyRequest
	GetPageSize() *int32
}

type ListRbacRoleHierarchyRequest struct {
	// example:
	//
	// WS20250731233102000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAARbaCuN6hiD08qrLdwJ9Fh3BFw8paIJ7ylB6A7Qn9JjM
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// CreationTime
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// example:
	//
	// DOWNGRADE
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 0
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRbacRoleHierarchyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRbacRoleHierarchyRequest) GoString() string {
	return s.String()
}

func (s *ListRbacRoleHierarchyRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListRbacRoleHierarchyRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRbacRoleHierarchyRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRbacRoleHierarchyRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListRbacRoleHierarchyRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListRbacRoleHierarchyRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListRbacRoleHierarchyRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRbacRoleHierarchyRequest) SetBizId(v string) *ListRbacRoleHierarchyRequest {
	s.BizId = &v
	return s
}

func (s *ListRbacRoleHierarchyRequest) SetMaxResults(v int32) *ListRbacRoleHierarchyRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRbacRoleHierarchyRequest) SetNextToken(v string) *ListRbacRoleHierarchyRequest {
	s.NextToken = &v
	return s
}

func (s *ListRbacRoleHierarchyRequest) SetOrderColumn(v string) *ListRbacRoleHierarchyRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListRbacRoleHierarchyRequest) SetOrderType(v string) *ListRbacRoleHierarchyRequest {
	s.OrderType = &v
	return s
}

func (s *ListRbacRoleHierarchyRequest) SetPageNum(v int32) *ListRbacRoleHierarchyRequest {
	s.PageNum = &v
	return s
}

func (s *ListRbacRoleHierarchyRequest) SetPageSize(v int32) *ListRbacRoleHierarchyRequest {
	s.PageSize = &v
	return s
}

func (s *ListRbacRoleHierarchyRequest) Validate() error {
	return dara.Validate(s)
}
