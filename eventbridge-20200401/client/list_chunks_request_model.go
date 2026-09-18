// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChunksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *ListChunksRequest
	GetCatalog() *string
	SetDocumentId(v string) *ListChunksRequest
	GetDocumentId() *string
	SetEnabled(v bool) *ListChunksRequest
	GetEnabled() *bool
	SetKeyword(v string) *ListChunksRequest
	GetKeyword() *string
	SetKnowledgeBaseName(v string) *ListChunksRequest
	GetKnowledgeBaseName() *string
	SetMaxResults(v int32) *ListChunksRequest
	GetMaxResults() *int32
	SetNamespace(v string) *ListChunksRequest
	GetNamespace() *string
	SetNextToken(v string) *ListChunksRequest
	GetNextToken() *string
}

type ListChunksRequest struct {
	// The data catalog to which the knowledge base belongs. This parameter, together with Namespace and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListCatalogs to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// Optional. If specified, only chunks of the specified document are returned, sorted by ChunkSeq in ascending order within the document. If not specified, chunks of all documents in the knowledge base are returned, sorted by DocumentId in lexicographic ascending order, and within the same document by ChunkSeq in ascending order, with chunks listed contiguously. In full knowledge base mode, Keyword filtering is not supported (use SearchKnowledgeBase instead), and TotalCount is not returned. Pagination ends when NextToken is empty.
	//
	// example:
	//
	// doc-bp1xxxxxxxxxxxx
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// Optional. Set to true to return only enabled chunks, or false to return only disabled chunks. If not specified, all chunks are returned.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Optional. Active only when DocumentId is specified. Filters chunks by keyword in the chunk body. Only chunks that contain the specified keyword are returned. In full knowledge base pattern (when DocumentId is not specified), passing this parameter causes an error. To retrieve content across the full text, use SearchKnowledgeBase (set Mode to KEYWORD for full-text index).
	//
	// example:
	//
	// Installation
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The name of the knowledge base. The name is unique within a namespace and is determined at creation time. It cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-knowledge-base
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// The maximum number of results to return per page. If not specified or set to 0, the default value 20 is used. The maximum value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The namespace to which the knowledge base belongs. The namespace must belong to the specified data catalog. This parameter, together with Catalog and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListNamespaces to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Not required for the first query. For subsequent queries, pass the NextToken returned in the previous response. An empty value indicates that no more pages are available.
	//
	// example:
	//
	// ca1eb85f5d99c7d6a97e6****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListChunksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChunksRequest) GoString() string {
	return s.String()
}

func (s *ListChunksRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListChunksRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *ListChunksRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListChunksRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListChunksRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *ListChunksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListChunksRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListChunksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListChunksRequest) SetCatalog(v string) *ListChunksRequest {
	s.Catalog = &v
	return s
}

func (s *ListChunksRequest) SetDocumentId(v string) *ListChunksRequest {
	s.DocumentId = &v
	return s
}

func (s *ListChunksRequest) SetEnabled(v bool) *ListChunksRequest {
	s.Enabled = &v
	return s
}

func (s *ListChunksRequest) SetKeyword(v string) *ListChunksRequest {
	s.Keyword = &v
	return s
}

func (s *ListChunksRequest) SetKnowledgeBaseName(v string) *ListChunksRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *ListChunksRequest) SetMaxResults(v int32) *ListChunksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListChunksRequest) SetNamespace(v string) *ListChunksRequest {
	s.Namespace = &v
	return s
}

func (s *ListChunksRequest) SetNextToken(v string) *ListChunksRequest {
	s.NextToken = &v
	return s
}

func (s *ListChunksRequest) Validate() error {
	return dara.Validate(s)
}
