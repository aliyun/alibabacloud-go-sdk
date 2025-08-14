// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdInsertionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListAdInsertionsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListAdInsertionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAdInsertionsRequest
	GetNextToken() *string
	SetPageNo(v int64) *ListAdInsertionsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListAdInsertionsRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListAdInsertionsRequest
	GetSortBy() *string
}

type ListAdInsertionsRequest struct {
	// The configuration name. Fuzzy match is supported.
	//
	// example:
	//
	// ad
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The maximum number of entries to retrieve in a subsequent request. If this parameter is used, the pagination parameters become invalid. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used in the next request to retrieve a new page of results. If this parameter is used, the pagination parameters become invalid.
	//
	// example:
	//
	// ******8EqYpQbZ6Eh7+Zz8DxVYoQ*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting order of the configurations by creation time. asc: ascending. desc: descending.
	//
	// example:
	//
	// asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListAdInsertionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAdInsertionsRequest) GoString() string {
	return s.String()
}

func (s *ListAdInsertionsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAdInsertionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAdInsertionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAdInsertionsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListAdInsertionsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAdInsertionsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListAdInsertionsRequest) SetKeyword(v string) *ListAdInsertionsRequest {
	s.Keyword = &v
	return s
}

func (s *ListAdInsertionsRequest) SetMaxResults(v int32) *ListAdInsertionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAdInsertionsRequest) SetNextToken(v string) *ListAdInsertionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAdInsertionsRequest) SetPageNo(v int64) *ListAdInsertionsRequest {
	s.PageNo = &v
	return s
}

func (s *ListAdInsertionsRequest) SetPageSize(v int64) *ListAdInsertionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAdInsertionsRequest) SetSortBy(v string) *ListAdInsertionsRequest {
	s.SortBy = &v
	return s
}

func (s *ListAdInsertionsRequest) Validate() error {
	return dara.Validate(s)
}
