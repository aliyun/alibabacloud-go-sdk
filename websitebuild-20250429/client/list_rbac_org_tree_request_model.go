// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRbacOrgTreeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListRbacOrgTreeRequest
	GetBizId() *string
	SetMaxResults(v int32) *ListRbacOrgTreeRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRbacOrgTreeRequest
	GetNextToken() *string
	SetOrderColumn(v string) *ListRbacOrgTreeRequest
	GetOrderColumn() *string
	SetOrderType(v string) *ListRbacOrgTreeRequest
	GetOrderType() *string
	SetPageNum(v int32) *ListRbacOrgTreeRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListRbacOrgTreeRequest
	GetPageSize() *int32
}

type ListRbacOrgTreeRequest struct {
	// example:
	//
	// WD20250703155602000001
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
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// example:
	//
	// BUY
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

func (s ListRbacOrgTreeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRbacOrgTreeRequest) GoString() string {
	return s.String()
}

func (s *ListRbacOrgTreeRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListRbacOrgTreeRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRbacOrgTreeRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRbacOrgTreeRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListRbacOrgTreeRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListRbacOrgTreeRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListRbacOrgTreeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRbacOrgTreeRequest) SetBizId(v string) *ListRbacOrgTreeRequest {
	s.BizId = &v
	return s
}

func (s *ListRbacOrgTreeRequest) SetMaxResults(v int32) *ListRbacOrgTreeRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRbacOrgTreeRequest) SetNextToken(v string) *ListRbacOrgTreeRequest {
	s.NextToken = &v
	return s
}

func (s *ListRbacOrgTreeRequest) SetOrderColumn(v string) *ListRbacOrgTreeRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListRbacOrgTreeRequest) SetOrderType(v string) *ListRbacOrgTreeRequest {
	s.OrderType = &v
	return s
}

func (s *ListRbacOrgTreeRequest) SetPageNum(v int32) *ListRbacOrgTreeRequest {
	s.PageNum = &v
	return s
}

func (s *ListRbacOrgTreeRequest) SetPageSize(v int32) *ListRbacOrgTreeRequest {
	s.PageSize = &v
	return s
}

func (s *ListRbacOrgTreeRequest) Validate() error {
	return dara.Validate(s)
}
