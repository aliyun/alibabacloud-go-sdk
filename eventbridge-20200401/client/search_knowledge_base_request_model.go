// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *SearchKnowledgeBaseRequest
	GetCatalog() *string
	SetKnowledgeBaseName(v string) *SearchKnowledgeBaseRequest
	GetKnowledgeBaseName() *string
	SetMetadataFilter(v string) *SearchKnowledgeBaseRequest
	GetMetadataFilter() *string
	SetMode(v string) *SearchKnowledgeBaseRequest
	GetMode() *string
	SetNamespace(v string) *SearchKnowledgeBaseRequest
	GetNamespace() *string
	SetQuery(v string) *SearchKnowledgeBaseRequest
	GetQuery() *string
	SetRankAlgorithm(v string) *SearchKnowledgeBaseRequest
	GetRankAlgorithm() *string
	SetRerank(v bool) *SearchKnowledgeBaseRequest
	GetRerank() *bool
	SetRerankModel(v string) *SearchKnowledgeBaseRequest
	GetRerankModel() *string
	SetRrfK(v int32) *SearchKnowledgeBaseRequest
	GetRrfK() *int32
	SetTopK(v int32) *SearchKnowledgeBaseRequest
	GetTopK() *int32
	SetVectorWeight(v float64) *SearchKnowledgeBaseRequest
	GetVectorWeight() *float64
}

type SearchKnowledgeBaseRequest struct {
	// The data catalog to which the knowledge base belongs. This parameter, together with Namespace and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListCatalogs to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The name of the knowledge base. The name is unique within a namespace and is specified at creation time. It cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-knowledge-base
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// Filters search results by document metadata. The value is a JSON expression tree. Nodes are classified into two types: logical nodes {"AndAll": [child nodes...]} (all conditions must be met), {"OrAll": [child nodes...]} (any condition must be met), or {"NotAll": [child nodes...]} (none of the conditions must be met), which support arbitrary nesting; and leaf conditions {"Key": field name, "Operator": operator, "Value": value}, where the In and NotIn operators use "Values": [values...]. Valid values of Operator: Equals, NotEquals, In, NotIn, GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual (the last four range operators apply only to LONG, DOUBLE, or DATETIME fields), StartsWith, and StringContains (applies only to STRING fields). Fields must be declared in the knowledge base schema. The nesting depth cannot exceed 5, and the total number of leaf conditions cannot exceed 20. DATETIME field value format: ISO-8601 with time zone, yyyy-MM-dd HH:mm:ss, or yyyy-MM-dd.
	//
	// example:
	//
	// {"AndAll": [{"Key": "env", "Operator": "In", "Values": ["prod", "staging"]}, {"OrAll": [{"Key": "score", "Operator": "GreaterThan", "Value": "0.8"}, {"Key": "owner", "Operator": "Equals", "Value": "alice"}]}]}
	MetadataFilter *string `json:"MetadataFilter,omitempty" xml:"MetadataFilter,omitempty"`
	// The retrieval mode. KEYWORD indicates keyword retrieval. VECTOR indicates vector retrieval. HYBRID indicates hybrid retrieval.
	//
	// example:
	//
	// HYBRID
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The namespace to which the knowledge base belongs. The namespace must belong to the specified data catalog. This parameter, together with Catalog and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListNamespaces to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The query text. Both keyword retrieval and vector retrieval are based on this text.
	//
	// This parameter is required.
	//
	// example:
	//
	// How to configure event rules
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Optional. Overrides the fusion algorithm configured for the knowledge base for this request. This parameter takes effect only in hybrid search (HYBRID) mode. RRF indicates reciprocal rank fusion. WEIGHTED indicates weighted normalization fusion (used together with VectorWeight). If you set this parameter to WEIGHTED in keyword-only or vector-only mode, an InvalidParameter error is returned.
	//
	// example:
	//
	// RRF
	RankAlgorithm *string `json:"RankAlgorithm,omitempty" xml:"RankAlgorithm,omitempty"`
	// Specifies whether to enable reranking for search results. Default value: false.
	//
	// example:
	//
	// false
	Rerank *bool `json:"Rerank,omitempty" xml:"Rerank,omitempty"`
	// Optional. Overrides the reranking model configured for the knowledge base for this request. This parameter takes effect when reranking is enabled (in all search modes). Valid values: qwen3-rerank, gte-rerank-v2, and qwen3-vl-rerank. If this parameter is not specified, the value configured for the knowledge base is used. If no value is configured, the default value qwen3-rerank is used. Score distributions differ across models and cannot be compared. Use the same model consistently for a given knowledge base.
	//
	// example:
	//
	// qwen3-rerank
	RerankModel *string `json:"RerankModel,omitempty" xml:"RerankModel,omitempty"`
	// Optional. Overrides the reciprocal rank fusion (RRF) parameter k configured for the knowledge base for this request. This parameter takes effect only in hybrid search (HYBRID) mode. The value must be greater than 0.
	//
	// example:
	//
	// 60
	RrfK *int32 `json:"RrfK,omitempty" xml:"RrfK,omitempty"`
	// The number of most relevant results to return. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// Optional. The vector weight for the WEIGHTED fusion algorithm. Valid values: 0 to 1. The keyword weight equals 1 minus this value. If this parameter is not specified, the value configured for the knowledge base is used. If no value is configured, the default value 0.7 is used.
	//
	// example:
	//
	// 0.7
	VectorWeight *float64 `json:"VectorWeight,omitempty" xml:"VectorWeight,omitempty"`
}

func (s SearchKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *SearchKnowledgeBaseRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *SearchKnowledgeBaseRequest) GetMetadataFilter() *string {
	return s.MetadataFilter
}

func (s *SearchKnowledgeBaseRequest) GetMode() *string {
	return s.Mode
}

func (s *SearchKnowledgeBaseRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *SearchKnowledgeBaseRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchKnowledgeBaseRequest) GetRankAlgorithm() *string {
	return s.RankAlgorithm
}

func (s *SearchKnowledgeBaseRequest) GetRerank() *bool {
	return s.Rerank
}

func (s *SearchKnowledgeBaseRequest) GetRerankModel() *string {
	return s.RerankModel
}

func (s *SearchKnowledgeBaseRequest) GetRrfK() *int32 {
	return s.RrfK
}

func (s *SearchKnowledgeBaseRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *SearchKnowledgeBaseRequest) GetVectorWeight() *float64 {
	return s.VectorWeight
}

func (s *SearchKnowledgeBaseRequest) SetCatalog(v string) *SearchKnowledgeBaseRequest {
	s.Catalog = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetKnowledgeBaseName(v string) *SearchKnowledgeBaseRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetMetadataFilter(v string) *SearchKnowledgeBaseRequest {
	s.MetadataFilter = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetMode(v string) *SearchKnowledgeBaseRequest {
	s.Mode = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetNamespace(v string) *SearchKnowledgeBaseRequest {
	s.Namespace = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetQuery(v string) *SearchKnowledgeBaseRequest {
	s.Query = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetRankAlgorithm(v string) *SearchKnowledgeBaseRequest {
	s.RankAlgorithm = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetRerank(v bool) *SearchKnowledgeBaseRequest {
	s.Rerank = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetRerankModel(v string) *SearchKnowledgeBaseRequest {
	s.RerankModel = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetRrfK(v int32) *SearchKnowledgeBaseRequest {
	s.RrfK = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetTopK(v int32) *SearchKnowledgeBaseRequest {
	s.TopK = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetVectorWeight(v float64) *SearchKnowledgeBaseRequest {
	s.VectorWeight = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
