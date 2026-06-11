// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentChunksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkTitlePattern(v string) *ListDocumentChunksRequest
	GetChunkTitlePattern() *string
	SetDocumentName(v string) *ListDocumentChunksRequest
	GetDocumentName() *string
	SetKbUuid(v string) *ListDocumentChunksRequest
	GetKbUuid() *string
	SetMaxResults(v int32) *ListDocumentChunksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDocumentChunksRequest
	GetNextToken() *string
	SetSortFieldName(v string) *ListDocumentChunksRequest
	GetSortFieldName() *string
	SetSortOrder(v string) *ListDocumentChunksRequest
	GetSortOrder() *string
}

type ListDocumentChunksRequest struct {
	// A filter pattern. The operation returns only the chunks whose titles contain this pattern.
	//
	// example:
	//
	// test
	ChunkTitlePattern *string `json:"ChunkTitlePattern,omitempty" xml:"ChunkTitlePattern,omitempty"`
	// The name of the document.
	//
	// example:
	//
	// test.md
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// The ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results. Omit this parameter to get the first page. If the response returns a `NextToken`, it indicates that more results are available. To get the next page, pass this `NextToken` value in the `NextToken` parameter of your next request. A null value for NextToken indicates that all results have been retrieved.
	//
	// example:
	//
	// zCXSmY0CJbybp6FZV7vo0Wjw64X-*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort field. Valid values:
	//
	// - **id*	- (default): The chunk ID.
	//
	// - **hits**: The number of hits.
	//
	// - **modifyTime**: The modification time.
	//
	// example:
	//
	// hits
	SortFieldName *string `json:"SortFieldName,omitempty" xml:"SortFieldName,omitempty"`
	// The sort order. Valid values:
	//
	// - **ASC*	- (default): ascending order.
	//
	// - **DESC**: descending order.
	//
	// example:
	//
	// DESC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListDocumentChunksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentChunksRequest) GoString() string {
	return s.String()
}

func (s *ListDocumentChunksRequest) GetChunkTitlePattern() *string {
	return s.ChunkTitlePattern
}

func (s *ListDocumentChunksRequest) GetDocumentName() *string {
	return s.DocumentName
}

func (s *ListDocumentChunksRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *ListDocumentChunksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDocumentChunksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDocumentChunksRequest) GetSortFieldName() *string {
	return s.SortFieldName
}

func (s *ListDocumentChunksRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListDocumentChunksRequest) SetChunkTitlePattern(v string) *ListDocumentChunksRequest {
	s.ChunkTitlePattern = &v
	return s
}

func (s *ListDocumentChunksRequest) SetDocumentName(v string) *ListDocumentChunksRequest {
	s.DocumentName = &v
	return s
}

func (s *ListDocumentChunksRequest) SetKbUuid(v string) *ListDocumentChunksRequest {
	s.KbUuid = &v
	return s
}

func (s *ListDocumentChunksRequest) SetMaxResults(v int32) *ListDocumentChunksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDocumentChunksRequest) SetNextToken(v string) *ListDocumentChunksRequest {
	s.NextToken = &v
	return s
}

func (s *ListDocumentChunksRequest) SetSortFieldName(v string) *ListDocumentChunksRequest {
	s.SortFieldName = &v
	return s
}

func (s *ListDocumentChunksRequest) SetSortOrder(v string) *ListDocumentChunksRequest {
	s.SortOrder = &v
	return s
}

func (s *ListDocumentChunksRequest) Validate() error {
	return dara.Validate(s)
}
