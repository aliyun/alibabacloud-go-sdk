// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v string) *ListDocumentsRequest
	GetFilters() *string
	SetKbUuid(v string) *ListDocumentsRequest
	GetKbUuid() *string
	SetMaxResults(v int32) *ListDocumentsRequest
	GetMaxResults() *int32
	SetNamePattern(v string) *ListDocumentsRequest
	GetNamePattern() *string
	SetNextToken(v string) *ListDocumentsRequest
	GetNextToken() *string
	SetSortFieldName(v string) *ListDocumentsRequest
	GetSortFieldName() *string
	SetSortOrder(v string) *ListDocumentsRequest
	GetSortOrder() *string
}

type ListDocumentsRequest struct {
	// The document property filter. The following properties are supported:
	//
	// - `fileExt`: The document extension.
	//
	// - `state`: The document status.
	//
	// example:
	//
	// {"fileExt":"pdf","state":0}
	Filters *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	// The maximum number of documents to return per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The document name filter. Returns only documents whose names contain this value.
	//
	// example:
	//
	// test
	NamePattern *string `json:"NamePattern,omitempty" xml:"NamePattern,omitempty"`
	// The pagination token to retrieve the next page of results. Omit this parameter to retrieve the first page. A `NextToken` value in the response indicates that more results are available. To fetch the next page, pass this value in a subsequent request. A `null` value indicates that all results have been retrieved.
	//
	// example:
	//
	// zCXSmY0CJbybp6FZV7vo0Wjw64X-*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort field. Valid values:
	//
	// - `id` (default): The document ID.
	//
	// - `hits`: The number of hits.
	//
	// - `modifyTime`: The modification time.
	//
	// example:
	//
	// hits
	SortFieldName *string `json:"SortFieldName,omitempty" xml:"SortFieldName,omitempty"`
	// The sort order. Valid values:
	//
	// - **ASC*	- (default): Sorts in ascending order.
	//
	// - **DESC**: Sorts in descending order.
	//
	// example:
	//
	// DESC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsRequest) GoString() string {
	return s.String()
}

func (s *ListDocumentsRequest) GetFilters() *string {
	return s.Filters
}

func (s *ListDocumentsRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *ListDocumentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDocumentsRequest) GetNamePattern() *string {
	return s.NamePattern
}

func (s *ListDocumentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDocumentsRequest) GetSortFieldName() *string {
	return s.SortFieldName
}

func (s *ListDocumentsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListDocumentsRequest) SetFilters(v string) *ListDocumentsRequest {
	s.Filters = &v
	return s
}

func (s *ListDocumentsRequest) SetKbUuid(v string) *ListDocumentsRequest {
	s.KbUuid = &v
	return s
}

func (s *ListDocumentsRequest) SetMaxResults(v int32) *ListDocumentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDocumentsRequest) SetNamePattern(v string) *ListDocumentsRequest {
	s.NamePattern = &v
	return s
}

func (s *ListDocumentsRequest) SetNextToken(v string) *ListDocumentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDocumentsRequest) SetSortFieldName(v string) *ListDocumentsRequest {
	s.SortFieldName = &v
	return s
}

func (s *ListDocumentsRequest) SetSortOrder(v string) *ListDocumentsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListDocumentsRequest) Validate() error {
	return dara.Validate(s)
}
