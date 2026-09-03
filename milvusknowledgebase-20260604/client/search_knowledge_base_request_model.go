// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentIds(v []*string) *SearchKnowledgeBaseRequest
	GetDocumentIds() []*string
	SetEnableKnowledgeGraph(v bool) *SearchKnowledgeBaseRequest
	GetEnableKnowledgeGraph() *bool
	SetImage(v *SearchKnowledgeBaseRequestImage) *SearchKnowledgeBaseRequest
	GetImage() *SearchKnowledgeBaseRequestImage
	SetPageNumber(v int32) *SearchKnowledgeBaseRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchKnowledgeBaseRequest
	GetPageSize() *int32
	SetQuery(v string) *SearchKnowledgeBaseRequest
	GetQuery() *string
	SetRerankModelId(v int64) *SearchKnowledgeBaseRequest
	GetRerankModelId() *int64
	SetRerankModelName(v string) *SearchKnowledgeBaseRequest
	GetRerankModelName() *string
	SetRetrievalConfig(v *SearchKnowledgeBaseRequestRetrievalConfig) *SearchKnowledgeBaseRequest
	GetRetrievalConfig() *SearchKnowledgeBaseRequestRetrievalConfig
	SetTagFilter(v *SearchKnowledgeBaseRequestTagFilter) *SearchKnowledgeBaseRequest
	GetTagFilter() *SearchKnowledgeBaseRequestTagFilter
	SetVersion(v string) *SearchKnowledgeBaseRequest
	GetVersion() *string
}

type SearchKnowledgeBaseRequest struct {
	// The list of document IDs.
	DocumentIds []*string `json:"documentIds,omitempty" xml:"documentIds,omitempty" type:"Repeated"`
	// Specifies whether to enable the knowledge graph.
	//
	// example:
	//
	// false
	EnableKnowledgeGraph *bool `json:"enableKnowledgeGraph,omitempty" xml:"enableKnowledgeGraph,omitempty"`
	// The image retrieval input.
	Image *SearchKnowledgeBaseRequestImage `json:"image,omitempty" xml:"image,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The retrieval query.
	//
	// example:
	//
	// What is the tax amount on the invoice?
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The reranking model ID.
	//
	// example:
	//
	// 123
	RerankModelId *int64 `json:"rerankModelId,omitempty" xml:"rerankModelId,omitempty"`
	// The name of a reranking model that the tenant has activated. If both rerankModelName and rerankModelId are specified, this parameter takes precedence.
	//
	// example:
	//
	// qwen3-rerank
	RerankModelName *string `json:"rerankModelName,omitempty" xml:"rerankModelName,omitempty"`
	// The retrieval configuration.
	RetrievalConfig *SearchKnowledgeBaseRequestRetrievalConfig `json:"retrievalConfig,omitempty" xml:"retrievalConfig,omitempty" type:"Struct"`
	// The tag filter.
	TagFilter *SearchKnowledgeBaseRequestTagFilter `json:"tagFilter,omitempty" xml:"tagFilter,omitempty" type:"Struct"`
	// The knowledge base version.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SearchKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseRequest) GetDocumentIds() []*string {
	return s.DocumentIds
}

func (s *SearchKnowledgeBaseRequest) GetEnableKnowledgeGraph() *bool {
	return s.EnableKnowledgeGraph
}

func (s *SearchKnowledgeBaseRequest) GetImage() *SearchKnowledgeBaseRequestImage {
	return s.Image
}

func (s *SearchKnowledgeBaseRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchKnowledgeBaseRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchKnowledgeBaseRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchKnowledgeBaseRequest) GetRerankModelId() *int64 {
	return s.RerankModelId
}

func (s *SearchKnowledgeBaseRequest) GetRerankModelName() *string {
	return s.RerankModelName
}

func (s *SearchKnowledgeBaseRequest) GetRetrievalConfig() *SearchKnowledgeBaseRequestRetrievalConfig {
	return s.RetrievalConfig
}

func (s *SearchKnowledgeBaseRequest) GetTagFilter() *SearchKnowledgeBaseRequestTagFilter {
	return s.TagFilter
}

func (s *SearchKnowledgeBaseRequest) GetVersion() *string {
	return s.Version
}

func (s *SearchKnowledgeBaseRequest) SetDocumentIds(v []*string) *SearchKnowledgeBaseRequest {
	s.DocumentIds = v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetEnableKnowledgeGraph(v bool) *SearchKnowledgeBaseRequest {
	s.EnableKnowledgeGraph = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetImage(v *SearchKnowledgeBaseRequestImage) *SearchKnowledgeBaseRequest {
	s.Image = v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetPageNumber(v int32) *SearchKnowledgeBaseRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetPageSize(v int32) *SearchKnowledgeBaseRequest {
	s.PageSize = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetQuery(v string) *SearchKnowledgeBaseRequest {
	s.Query = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetRerankModelId(v int64) *SearchKnowledgeBaseRequest {
	s.RerankModelId = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetRerankModelName(v string) *SearchKnowledgeBaseRequest {
	s.RerankModelName = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetRetrievalConfig(v *SearchKnowledgeBaseRequestRetrievalConfig) *SearchKnowledgeBaseRequest {
	s.RetrievalConfig = v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetTagFilter(v *SearchKnowledgeBaseRequestTagFilter) *SearchKnowledgeBaseRequest {
	s.TagFilter = v
	return s
}

func (s *SearchKnowledgeBaseRequest) SetVersion(v string) *SearchKnowledgeBaseRequest {
	s.Version = &v
	return s
}

func (s *SearchKnowledgeBaseRequest) Validate() error {
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			return err
		}
	}
	if s.RetrievalConfig != nil {
		if err := s.RetrievalConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TagFilter != nil {
		if err := s.TagFilter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchKnowledgeBaseRequestImage struct {
	// The Base64-encoded image.
	//
	// example:
	//
	// data:image/png;base64,iVBORw0KGgoAAA...
	Base64 *string `json:"base64,omitempty" xml:"base64,omitempty"`
	// The object key of the image.
	//
	// example:
	//
	// uploaded/invoice.png
	ObjectKey *string `json:"objectKey,omitempty" xml:"objectKey,omitempty"`
	// The image URL.
	//
	// example:
	//
	// https://example.com/invoice.png
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SearchKnowledgeBaseRequestImage) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseRequestImage) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseRequestImage) GetBase64() *string {
	return s.Base64
}

func (s *SearchKnowledgeBaseRequestImage) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *SearchKnowledgeBaseRequestImage) GetUrl() *string {
	return s.Url
}

func (s *SearchKnowledgeBaseRequestImage) SetBase64(v string) *SearchKnowledgeBaseRequestImage {
	s.Base64 = &v
	return s
}

func (s *SearchKnowledgeBaseRequestImage) SetObjectKey(v string) *SearchKnowledgeBaseRequestImage {
	s.ObjectKey = &v
	return s
}

func (s *SearchKnowledgeBaseRequestImage) SetUrl(v string) *SearchKnowledgeBaseRequestImage {
	s.Url = &v
	return s
}

func (s *SearchKnowledgeBaseRequestImage) Validate() error {
	return dara.Validate(s)
}

type SearchKnowledgeBaseRequestRetrievalConfig struct {
	// The number of candidate recall results.
	//
	// example:
	//
	// 5
	CandidateCount *int32 `json:"candidateCount,omitempty" xml:"candidateCount,omitempty"`
	// Specifies whether to enable query expansion.
	//
	// example:
	//
	// true
	EnableQueryExpansion *bool `json:"enableQueryExpansion,omitempty" xml:"enableQueryExpansion,omitempty"`
	// The minimum relevance score.
	//
	// example:
	//
	// 0.2
	MinScore *float32 `json:"minScore,omitempty" xml:"minScore,omitempty"`
	// The weight of semantic relevance.
	//
	// example:
	//
	// 0.5
	SemanticWeight *float32 `json:"semanticWeight,omitempty" xml:"semanticWeight,omitempty"`
	// The list of translation languages.
	TranslationLanguages []*string `json:"translationLanguages,omitempty" xml:"translationLanguages,omitempty" type:"Repeated"`
}

func (s SearchKnowledgeBaseRequestRetrievalConfig) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseRequestRetrievalConfig) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseRequestRetrievalConfig) GetCandidateCount() *int32 {
	return s.CandidateCount
}

func (s *SearchKnowledgeBaseRequestRetrievalConfig) GetEnableQueryExpansion() *bool {
	return s.EnableQueryExpansion
}

func (s *SearchKnowledgeBaseRequestRetrievalConfig) GetMinScore() *float32 {
	return s.MinScore
}

func (s *SearchKnowledgeBaseRequestRetrievalConfig) GetSemanticWeight() *float32 {
	return s.SemanticWeight
}

func (s *SearchKnowledgeBaseRequestRetrievalConfig) GetTranslationLanguages() []*string {
	return s.TranslationLanguages
}

func (s *SearchKnowledgeBaseRequestRetrievalConfig) SetCandidateCount(v int32) *SearchKnowledgeBaseRequestRetrievalConfig {
	s.CandidateCount = &v
	return s
}

func (s *SearchKnowledgeBaseRequestRetrievalConfig) SetEnableQueryExpansion(v bool) *SearchKnowledgeBaseRequestRetrievalConfig {
	s.EnableQueryExpansion = &v
	return s
}

func (s *SearchKnowledgeBaseRequestRetrievalConfig) SetMinScore(v float32) *SearchKnowledgeBaseRequestRetrievalConfig {
	s.MinScore = &v
	return s
}

func (s *SearchKnowledgeBaseRequestRetrievalConfig) SetSemanticWeight(v float32) *SearchKnowledgeBaseRequestRetrievalConfig {
	s.SemanticWeight = &v
	return s
}

func (s *SearchKnowledgeBaseRequestRetrievalConfig) SetTranslationLanguages(v []*string) *SearchKnowledgeBaseRequestRetrievalConfig {
	s.TranslationLanguages = v
	return s
}

func (s *SearchKnowledgeBaseRequestRetrievalConfig) Validate() error {
	return dara.Validate(s)
}

type SearchKnowledgeBaseRequestTagFilter struct {
	// The list of tag conditions.
	Conditions []*SearchKnowledgeBaseRequestTagFilterConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The logical relation between conditions.
	//
	// example:
	//
	// or
	Relation *string `json:"relation,omitempty" xml:"relation,omitempty"`
}

func (s SearchKnowledgeBaseRequestTagFilter) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseRequestTagFilter) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseRequestTagFilter) GetConditions() []*SearchKnowledgeBaseRequestTagFilterConditions {
	return s.Conditions
}

func (s *SearchKnowledgeBaseRequestTagFilter) GetRelation() *string {
	return s.Relation
}

func (s *SearchKnowledgeBaseRequestTagFilter) SetConditions(v []*SearchKnowledgeBaseRequestTagFilterConditions) *SearchKnowledgeBaseRequestTagFilter {
	s.Conditions = v
	return s
}

func (s *SearchKnowledgeBaseRequestTagFilter) SetRelation(v string) *SearchKnowledgeBaseRequestTagFilter {
	s.Relation = &v
	return s
}

func (s *SearchKnowledgeBaseRequestTagFilter) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchKnowledgeBaseRequestTagFilterConditions struct {
	// The tag field.
	//
	// example:
	//
	// category
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// The operator.
	//
	// example:
	//
	// in
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// The tag value.
	//
	// example:
	//
	// invoice
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SearchKnowledgeBaseRequestTagFilterConditions) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseRequestTagFilterConditions) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseRequestTagFilterConditions) GetField() *string {
	return s.Field
}

func (s *SearchKnowledgeBaseRequestTagFilterConditions) GetOp() *string {
	return s.Op
}

func (s *SearchKnowledgeBaseRequestTagFilterConditions) GetValue() interface{} {
	return s.Value
}

func (s *SearchKnowledgeBaseRequestTagFilterConditions) SetField(v string) *SearchKnowledgeBaseRequestTagFilterConditions {
	s.Field = &v
	return s
}

func (s *SearchKnowledgeBaseRequestTagFilterConditions) SetOp(v string) *SearchKnowledgeBaseRequestTagFilterConditions {
	s.Op = &v
	return s
}

func (s *SearchKnowledgeBaseRequestTagFilterConditions) SetValue(v interface{}) *SearchKnowledgeBaseRequestTagFilterConditions {
	s.Value = v
	return s
}

func (s *SearchKnowledgeBaseRequestTagFilterConditions) Validate() error {
	return dara.Validate(s)
}
