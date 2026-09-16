// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaChunksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *ListLumaChunksRequest
	GetAgentName() *string
	SetCatalog(v string) *ListLumaChunksRequest
	GetCatalog() *string
	SetDocumentId(v string) *ListLumaChunksRequest
	GetDocumentId() *string
	SetEnabled(v bool) *ListLumaChunksRequest
	GetEnabled() *bool
	SetKeyword(v string) *ListLumaChunksRequest
	GetKeyword() *string
	SetKnowledgeBaseName(v string) *ListLumaChunksRequest
	GetKnowledgeBaseName() *string
	SetMaxResults(v int32) *ListLumaChunksRequest
	GetMaxResults() *int32
	SetNamespace(v string) *ListLumaChunksRequest
	GetNamespace() *string
	SetNextToken(v string) *ListLumaChunksRequest
	GetNextToken() *string
}

type ListLumaChunksRequest struct {
	// The name of the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The name of the data catalog bound to the agent. You can call ListLumaCatalogs to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The ID of the document used to filter text chunks. If this parameter is not specified, text chunks of all documents in the knowledge base are returned.
	//
	// example:
	//
	// doc-a1b2c3d4
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// Specifies whether to return only enabled text chunks. If this parameter is not specified, text chunks are not filtered by enabled status.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The keyword used to filter text chunks by content.
	//
	// example:
	//
	// Event rule
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The name of the knowledge base bound to the agent. You can call ListLumaKnowledgeBases to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-knowledge-base
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// The maximum number of records to return. Valid values: 1 to 100. If this parameter is not specified, the server uses a default value.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the namespace bound to the agent. You can call ListLumaNamespaces to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The pagination token. Do not specify this parameter for the first request. For subsequent requests, use the NextToken value returned in the previous response. This value is an opaque string. Do not parse it.
	//
	// example:
	//
	// ca1eb85f5d99c7d6a97e6****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListLumaChunksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLumaChunksRequest) GoString() string {
	return s.String()
}

func (s *ListLumaChunksRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ListLumaChunksRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListLumaChunksRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *ListLumaChunksRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListLumaChunksRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListLumaChunksRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *ListLumaChunksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLumaChunksRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListLumaChunksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLumaChunksRequest) SetAgentName(v string) *ListLumaChunksRequest {
	s.AgentName = &v
	return s
}

func (s *ListLumaChunksRequest) SetCatalog(v string) *ListLumaChunksRequest {
	s.Catalog = &v
	return s
}

func (s *ListLumaChunksRequest) SetDocumentId(v string) *ListLumaChunksRequest {
	s.DocumentId = &v
	return s
}

func (s *ListLumaChunksRequest) SetEnabled(v bool) *ListLumaChunksRequest {
	s.Enabled = &v
	return s
}

func (s *ListLumaChunksRequest) SetKeyword(v string) *ListLumaChunksRequest {
	s.Keyword = &v
	return s
}

func (s *ListLumaChunksRequest) SetKnowledgeBaseName(v string) *ListLumaChunksRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *ListLumaChunksRequest) SetMaxResults(v int32) *ListLumaChunksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLumaChunksRequest) SetNamespace(v string) *ListLumaChunksRequest {
	s.Namespace = &v
	return s
}

func (s *ListLumaChunksRequest) SetNextToken(v string) *ListLumaChunksRequest {
	s.NextToken = &v
	return s
}

func (s *ListLumaChunksRequest) Validate() error {
	return dara.Validate(s)
}
