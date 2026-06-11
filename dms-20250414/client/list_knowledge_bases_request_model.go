// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v string) *ListKnowledgeBasesRequest
	GetFilters() *string
	SetMaxResults(v int32) *ListKnowledgeBasesRequest
	GetMaxResults() *int32
	SetNamePattern(v string) *ListKnowledgeBasesRequest
	GetNamePattern() *string
	SetNextToken(v string) *ListKnowledgeBasesRequest
	GetNextToken() *string
	SetSortFieldName(v string) *ListKnowledgeBasesRequest
	GetSortFieldName() *string
	SetSortOrder(v string) *ListKnowledgeBasesRequest
	GetSortOrder() *string
	SetTag(v string) *ListKnowledgeBasesRequest
	GetTag() *string
}

type ListKnowledgeBasesRequest struct {
	// The filter conditions for the knowledge bases, specified as a JSON string. The only supported key is `state`. Valid values are `0` and `1`.
	//
	// example:
	//
	// {"state":1}
	Filters *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The maximum number of entries to return on each page. Use this parameter with the `NextToken` parameter to implement pagination.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A keyword to search for in the names of knowledge bases.
	//
	// example:
	//
	// order
	NamePattern *string `json:"NamePattern,omitempty" xml:"NamePattern,omitempty"`
	// The token used to retrieve the next page of results. Valid values:
	//
	// - Omit this parameter for the first request.
	//
	// - If the previous response returned a **NextToken*	- value, use it to retrieve the next page of results.
	//
	// example:
	//
	// zCXSmY0CJbybp6FZV7vo0Wjw64X-*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort field. Valid values:
	//
	// - `id`: Sorts by knowledge base ID. This is the default.
	//
	// - `name`: Sorts by knowledge base name.
	//
	// example:
	//
	// name
	SortFieldName *string `json:"SortFieldName,omitempty" xml:"SortFieldName,omitempty"`
	// The sort order. Valid values:
	//
	// - **ASC**: Ascending order. This is the default.
	//
	// - **DESC**: Descending order.
	//
	// example:
	//
	// Desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The tag of the knowledge base. In DataAgent, this is the space ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1dq7qod8hxtt1***
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListKnowledgeBasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBasesRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBasesRequest) GetFilters() *string {
	return s.Filters
}

func (s *ListKnowledgeBasesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBasesRequest) GetNamePattern() *string {
	return s.NamePattern
}

func (s *ListKnowledgeBasesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBasesRequest) GetSortFieldName() *string {
	return s.SortFieldName
}

func (s *ListKnowledgeBasesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListKnowledgeBasesRequest) GetTag() *string {
	return s.Tag
}

func (s *ListKnowledgeBasesRequest) SetFilters(v string) *ListKnowledgeBasesRequest {
	s.Filters = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetMaxResults(v int32) *ListKnowledgeBasesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetNamePattern(v string) *ListKnowledgeBasesRequest {
	s.NamePattern = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetNextToken(v string) *ListKnowledgeBasesRequest {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetSortFieldName(v string) *ListKnowledgeBasesRequest {
	s.SortFieldName = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetSortOrder(v string) *ListKnowledgeBasesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetTag(v string) *ListKnowledgeBasesRequest {
	s.Tag = &v
	return s
}

func (s *ListKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
