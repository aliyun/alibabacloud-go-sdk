// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListGroupsRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListGroupsRequest
	GetLang() *string
	SetNextToken(v string) *ListGroupsRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListGroupsRequest
	GetPageSize() *int32
	SetUseNextToken(v bool) *ListGroupsRequest
	GetUseNextToken() *bool
}

type ListGroupsRequest struct {
	// The page number of the current page to return. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 89
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The token for the next query. If NextToken is empty, no additional results exist. If NextToken has a value, the value indicates the token to use for the next query.
	//
	// example:
	//
	// 1426C575705AE8545E8360A6EFA3B***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The maximum number of entries to return on each page in a paging query. Default value: 20. Maximum value: 2000.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to use the NextToken method to retrieve the vulnerability list data. If this parameter is used, TotalCount is no longer returned. Valid values:
	//
	// - **true**: Uses the NextToken method.
	//
	// - **false**: Does not use the NextToken method.
	//
	// example:
	//
	// true
	UseNextToken *bool `json:"UseNextToken,omitempty" xml:"UseNextToken,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListGroupsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGroupsRequest) GetUseNextToken() *bool {
	return s.UseNextToken
}

func (s *ListGroupsRequest) SetCurrentPage(v int32) *ListGroupsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListGroupsRequest) SetLang(v string) *ListGroupsRequest {
	s.Lang = &v
	return s
}

func (s *ListGroupsRequest) SetNextToken(v string) *ListGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListGroupsRequest) SetPageSize(v int32) *ListGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupsRequest) SetUseNextToken(v bool) *ListGroupsRequest {
	s.UseNextToken = &v
	return s
}

func (s *ListGroupsRequest) Validate() error {
	return dara.Validate(s)
}
