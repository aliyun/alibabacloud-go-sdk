// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeTags(v bool) *ListAccountsRequest
	GetIncludeTags() *bool
	SetMaxResults(v int32) *ListAccountsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAccountsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListAccountsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAccountsRequest
	GetPageSize() *int32
	SetQueryKeyword(v string) *ListAccountsRequest
	GetQueryKeyword() *string
	SetTag(v []*ListAccountsRequestTag) *ListAccountsRequest
	GetTag() []*ListAccountsRequestTag
}

type ListAccountsRequest struct {
	// Specifies whether to return information about tags. Valid values:
	//
	// 	- false (default value)
	//
	// 	- true
	//
	// example:
	//
	// true
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The number of entries per page. After you configure this parameter, token-based paging is preferentially used.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword used for the query, such as the display name of a member.
	//
	// Fuzzy match is supported.
	//
	// example:
	//
	// Advance
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	// The tags. This parameter specifies a filter condition.
	Tag []*ListAccountsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListAccountsRequest) GetIncludeTags() *bool {
	return s.IncludeTags
}

func (s *ListAccountsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAccountsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccountsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAccountsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAccountsRequest) GetQueryKeyword() *string {
	return s.QueryKeyword
}

func (s *ListAccountsRequest) GetTag() []*ListAccountsRequestTag {
	return s.Tag
}

func (s *ListAccountsRequest) SetIncludeTags(v bool) *ListAccountsRequest {
	s.IncludeTags = &v
	return s
}

func (s *ListAccountsRequest) SetMaxResults(v int32) *ListAccountsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAccountsRequest) SetNextToken(v string) *ListAccountsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAccountsRequest) SetPageNumber(v int32) *ListAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAccountsRequest) SetPageSize(v int32) *ListAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAccountsRequest) SetQueryKeyword(v string) *ListAccountsRequest {
	s.QueryKeyword = &v
	return s
}

func (s *ListAccountsRequest) SetTag(v []*ListAccountsRequestTag) *ListAccountsRequest {
	s.Tag = v
	return s
}

func (s *ListAccountsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccountsRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tag_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAccountsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsRequestTag) GoString() string {
	return s.String()
}

func (s *ListAccountsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListAccountsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListAccountsRequestTag) SetKey(v string) *ListAccountsRequestTag {
	s.Key = &v
	return s
}

func (s *ListAccountsRequestTag) SetValue(v string) *ListAccountsRequestTag {
	s.Value = &v
	return s
}

func (s *ListAccountsRequestTag) Validate() error {
	return dara.Validate(s)
}
