// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchLumaKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *SearchLumaKnowledgeBaseRequest
	GetAgentName() *string
	SetCatalog(v string) *SearchLumaKnowledgeBaseRequest
	GetCatalog() *string
	SetKnowledgeBaseName(v string) *SearchLumaKnowledgeBaseRequest
	GetKnowledgeBaseName() *string
	SetMetadataFilter(v string) *SearchLumaKnowledgeBaseRequest
	GetMetadataFilter() *string
	SetMode(v string) *SearchLumaKnowledgeBaseRequest
	GetMode() *string
	SetNamespace(v string) *SearchLumaKnowledgeBaseRequest
	GetNamespace() *string
	SetQuery(v string) *SearchLumaKnowledgeBaseRequest
	GetQuery() *string
	SetRerank(v bool) *SearchLumaKnowledgeBaseRequest
	GetRerank() *bool
	SetTopK(v int32) *SearchLumaKnowledgeBaseRequest
	GetTopK() *int32
}

type SearchLumaKnowledgeBaseRequest struct {
	// The name of the Agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The name of the data catalog bound to the Agent. You can call ListLumaCatalogs to obtain the catalog name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The name of the knowledge base bound to the Agent. You can call ListLumaKnowledgeBases to obtain the knowledge base name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-knowledge-base
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// A JSON string that filters the retrieval scope based on document metadata. For available fields, refer to the MetadataSchema of the knowledge base.
	//
	// example:
	//
	// {"category":"faq"}
	MetadataFilter *string `json:"MetadataFilter,omitempty" xml:"MetadataFilter,omitempty"`
	// Valid values: vector (AISearch), keyword (keyword match), hybrid (hybrid search). If not specified, the retrieve configuration of the knowledge base is used.
	//
	// example:
	//
	// hybrid
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name of the namespace bound to the Agent. You can call ListLumaNamespaces to obtain the namespace name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The natural language query for retrieval.
	//
	// This parameter is required.
	//
	// example:
	//
	// How to configure event rules
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Specifies whether to enable reranking for the retrieved results. Reranking improves accuracy but increases latency. If not specified, the retrieval configuration of the knowledge base is used.
	//
	// example:
	//
	// true
	Rerank *bool `json:"Rerank,omitempty" xml:"Rerank,omitempty"`
	// Valid values: 1 to 100. If not specified, the retrieval configuration of the knowledge base is used.
	//
	// example:
	//
	// 5
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s SearchLumaKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchLumaKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *SearchLumaKnowledgeBaseRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *SearchLumaKnowledgeBaseRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *SearchLumaKnowledgeBaseRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *SearchLumaKnowledgeBaseRequest) GetMetadataFilter() *string {
	return s.MetadataFilter
}

func (s *SearchLumaKnowledgeBaseRequest) GetMode() *string {
	return s.Mode
}

func (s *SearchLumaKnowledgeBaseRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *SearchLumaKnowledgeBaseRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchLumaKnowledgeBaseRequest) GetRerank() *bool {
	return s.Rerank
}

func (s *SearchLumaKnowledgeBaseRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *SearchLumaKnowledgeBaseRequest) SetAgentName(v string) *SearchLumaKnowledgeBaseRequest {
	s.AgentName = &v
	return s
}

func (s *SearchLumaKnowledgeBaseRequest) SetCatalog(v string) *SearchLumaKnowledgeBaseRequest {
	s.Catalog = &v
	return s
}

func (s *SearchLumaKnowledgeBaseRequest) SetKnowledgeBaseName(v string) *SearchLumaKnowledgeBaseRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *SearchLumaKnowledgeBaseRequest) SetMetadataFilter(v string) *SearchLumaKnowledgeBaseRequest {
	s.MetadataFilter = &v
	return s
}

func (s *SearchLumaKnowledgeBaseRequest) SetMode(v string) *SearchLumaKnowledgeBaseRequest {
	s.Mode = &v
	return s
}

func (s *SearchLumaKnowledgeBaseRequest) SetNamespace(v string) *SearchLumaKnowledgeBaseRequest {
	s.Namespace = &v
	return s
}

func (s *SearchLumaKnowledgeBaseRequest) SetQuery(v string) *SearchLumaKnowledgeBaseRequest {
	s.Query = &v
	return s
}

func (s *SearchLumaKnowledgeBaseRequest) SetRerank(v bool) *SearchLumaKnowledgeBaseRequest {
	s.Rerank = &v
	return s
}

func (s *SearchLumaKnowledgeBaseRequest) SetTopK(v int32) *SearchLumaKnowledgeBaseRequest {
	s.TopK = &v
	return s
}

func (s *SearchLumaKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
