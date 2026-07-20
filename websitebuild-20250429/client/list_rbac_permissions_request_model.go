// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRbacPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListRbacPermissionsRequest
	GetBizId() *string
	SetMaxResults(v int32) *ListRbacPermissionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRbacPermissionsRequest
	GetNextToken() *string
	SetOrderColumn(v string) *ListRbacPermissionsRequest
	GetOrderColumn() *string
	SetOrderType(v string) *ListRbacPermissionsRequest
	GetOrderType() *string
	SetPageNum(v int32) *ListRbacPermissionsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListRbacPermissionsRequest
	GetPageSize() *int32
}

type ListRbacPermissionsRequest struct {
	// The business ID.
	//
	// example:
	//
	// WS20250801154628000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The number of entries per query.
	//
	// Valid values: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. This parameter is empty if no more results exist.
	//
	// example:
	//
	// AAAAARbaCuN6hiD08qrLdwJ9Fh3BFw8paIJ7ylB6A7Qn9JjM
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The field used for sorting.
	//
	// example:
	//
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// The sort type. Valid values: ASC and DESC.
	//
	// example:
	//
	// DESC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRbacPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRbacPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListRbacPermissionsRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListRbacPermissionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRbacPermissionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRbacPermissionsRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListRbacPermissionsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListRbacPermissionsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListRbacPermissionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRbacPermissionsRequest) SetBizId(v string) *ListRbacPermissionsRequest {
	s.BizId = &v
	return s
}

func (s *ListRbacPermissionsRequest) SetMaxResults(v int32) *ListRbacPermissionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRbacPermissionsRequest) SetNextToken(v string) *ListRbacPermissionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRbacPermissionsRequest) SetOrderColumn(v string) *ListRbacPermissionsRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListRbacPermissionsRequest) SetOrderType(v string) *ListRbacPermissionsRequest {
	s.OrderType = &v
	return s
}

func (s *ListRbacPermissionsRequest) SetPageNum(v int32) *ListRbacPermissionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListRbacPermissionsRequest) SetPageSize(v int32) *ListRbacPermissionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRbacPermissionsRequest) Validate() error {
	return dara.Validate(s)
}
