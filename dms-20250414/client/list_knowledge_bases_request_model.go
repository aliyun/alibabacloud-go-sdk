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
	// example:
	//
	// {"state":1}
	Filters *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// order
	NamePattern *string `json:"NamePattern,omitempty" xml:"NamePattern,omitempty"`
	// example:
	//
	// zCXSmY0CJbybp6FZV7vo0Wjw64X-*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// name
	SortFieldName *string `json:"SortFieldName,omitempty" xml:"SortFieldName,omitempty"`
	// example:
	//
	// Desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
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
