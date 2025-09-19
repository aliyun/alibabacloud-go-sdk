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
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 89
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the return value of NextToken is empty, no next query is to be sent. If a value of NextToken is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// 1426C575705AE8545E8360A6EFA3B***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 2000.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to use NextToken to query vulnerabilities. If you set this parameter to true, TotalCount is not returned. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
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
