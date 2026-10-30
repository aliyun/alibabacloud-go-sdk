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
	// The data catalog to which the knowledge base belongs. This parameter, together with Namespace and KnowledgeBaseName, uniquely identifies the knowledge base. You can call ListCatalogs to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The name of the knowledge base. The name must be unique within the namespace. The name is specified upon creation and cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-knowledge-base
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// Filters search results by document metadata. The value is a JSON expression tree. Nodes are divided into two types: logical nodes and leaf conditions. Logical nodes include {"AndAll": [child nodes...]} (all conditions met), {"OrAll": [child nodes...]} (any condition met), and {"NotAll": [child nodes...]} (no conditions met). Arbitrary nesting is supported. Leaf conditions are in the format of {"Key": field name, "Operator": operator, "Value": value}. For the In and NotIn operators, use "Values": [values...]. Valid values for Operator: Equals, NotEquals, In, NotIn, GreaterThan, GreaterThanOrEqual, LessThan, and LessThanOrEqual (the last four range operators apply only to LONG, DOUBLE, and DATETIME fields), StartsWith, and StringContains (applies only to STRING fields). The field must be declared in the knowledge base schema. The nesting depth cannot exceed 5, and the total number of leaf conditions cannot exceed 20. Valid formats for DATETIME field values: ISO-8601 with time zone, yyyy-MM-dd HH:mm:ss, or yyyy-MM-dd.
	//
	// example:
	//
	// {"AndAll": [{"Key": "env", "Operator": "In", "Values": ["prod", "staging"]}, {"OrAll": [{"Key": "score", "Operator": "GreaterThan", "Value": "0.8"}, {"Key": "owner", "Operator": "Equals", "Value": "alice"}]}]}
	MetadataFilter *string `json:"MetadataFilter,omitempty" xml:"MetadataFilter,omitempty"`
	// The retrieve mode. Valid values:
	//
	// - KEYWORD: keyword retrieve.
	//
	// - VECTOR: AISearch.
	//
	// - HYBRID: hybrid retrieve.
	//
	// example:
	//
	// HYBRID
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The namespace to which the knowledge base belongs. The namespace must belong to the specified data catalog. This parameter, together with Catalog and KnowledgeBaseName, uniquely identifies the knowledge base. You can call ListNamespaces to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The retrieve text. Both keyword retrieve and AISearch are based on this text.
	//
	// This parameter is required.
	//
	// example:
	//
	// How to configure event rules
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Optional. Overrides the fusion algorithm configured for the knowledge base in this request. This parameter takes effect only in HYBRID search mode. Valid values: RRF for reciprocal rank fusion, and WEIGHTED for weighted normalization fusion (used with VectorWeight). If you set this parameter to WEIGHTED in KEYWORD or VECTOR single mode, the InvalidParameter error is returned.
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
	// Optional. Overrides the reranking model configured for the knowledge base in this request. This parameter takes effect when reranking is enabled across all search modes. Valid values: qwen3-rerank, gte-rerank-v2, and qwen3-vl-rerank. If this parameter is not specified, the value configured for the knowledge base is used. If no value is configured, the default value is qwen3-rerank. Score distributions of different models cannot be compared horizontally. We recommend that you use a fixed model for each knowledge base.
	//
	// example:
	//
	// qwen3-rerank
	RerankModel *string `json:"RerankModel,omitempty" xml:"RerankModel,omitempty"`
	// Optional. Overrides the RRF fusion parameter k configured for the knowledge base in this request. This parameter takes effect only in HYBRID search mode and must be greater than 0.
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
	// Optional. The vector path weight for the WEIGHTED fusion algorithm. Valid values: 0 to 1. The keyword path weight is 1 minus this value. If this parameter is not specified, the value configured for the knowledge base is used. If no value is configured, the default value is 0.7.
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
