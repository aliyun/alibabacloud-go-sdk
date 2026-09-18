// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *UpdateKnowledgeBaseRequest
	GetCatalog() *string
	SetChunkConfiguration(v *UpdateKnowledgeBaseRequestChunkConfiguration) *UpdateKnowledgeBaseRequest
	GetChunkConfiguration() *UpdateKnowledgeBaseRequestChunkConfiguration
	SetDescription(v string) *UpdateKnowledgeBaseRequest
	GetDescription() *string
	SetKnowledgeBaseName(v string) *UpdateKnowledgeBaseRequest
	GetKnowledgeBaseName() *string
	SetNamespace(v string) *UpdateKnowledgeBaseRequest
	GetNamespace() *string
	SetSearchConfiguration(v *UpdateKnowledgeBaseRequestSearchConfiguration) *UpdateKnowledgeBaseRequest
	GetSearchConfiguration() *UpdateKnowledgeBaseRequestSearchConfiguration
}

type UpdateKnowledgeBaseRequest struct {
	// The data catalog to which the knowledge base belongs. This parameter, together with Namespace and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListCatalogs to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// Optional. Updates the default chunking strategy of the knowledge base. The update takes effect only for documents uploaded after the update. Existing documents are not re-chunked. If this parameter is not specified, the configuration remains unchanged.
	ChunkConfiguration *UpdateKnowledgeBaseRequestChunkConfiguration `json:"ChunkConfiguration,omitempty" xml:"ChunkConfiguration,omitempty" type:"Struct"`
	// The description of the knowledge base to update. If this parameter is not specified, the description remains unchanged.
	//
	// example:
	//
	// Product documentation knowledge base
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the knowledge base. The name must be unique within the namespace. The name is specified during creation and cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-knowledge-base
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// The namespace to which the knowledge base belongs. The namespace must belong to the specified data catalog. This parameter, together with Catalog and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListNamespaces to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Optional. Updates the default search configuration at the knowledge base level. The update takes effect immediately for subsequent search requests. If this parameter is not specified, the configuration remains unchanged.
	SearchConfiguration *UpdateKnowledgeBaseRequestSearchConfiguration `json:"SearchConfiguration,omitempty" xml:"SearchConfiguration,omitempty" type:"Struct"`
}

func (s UpdateKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *UpdateKnowledgeBaseRequest) GetChunkConfiguration() *UpdateKnowledgeBaseRequestChunkConfiguration {
	return s.ChunkConfiguration
}

func (s *UpdateKnowledgeBaseRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeBaseRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *UpdateKnowledgeBaseRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateKnowledgeBaseRequest) GetSearchConfiguration() *UpdateKnowledgeBaseRequestSearchConfiguration {
	return s.SearchConfiguration
}

func (s *UpdateKnowledgeBaseRequest) SetCatalog(v string) *UpdateKnowledgeBaseRequest {
	s.Catalog = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetChunkConfiguration(v *UpdateKnowledgeBaseRequestChunkConfiguration) *UpdateKnowledgeBaseRequest {
	s.ChunkConfiguration = v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetDescription(v string) *UpdateKnowledgeBaseRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetKnowledgeBaseName(v string) *UpdateKnowledgeBaseRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetNamespace(v string) *UpdateKnowledgeBaseRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetSearchConfiguration(v *UpdateKnowledgeBaseRequestSearchConfiguration) *UpdateKnowledgeBaseRequest {
	s.SearchConfiguration = v
	return s
}

func (s *UpdateKnowledgeBaseRequest) Validate() error {
	if s.ChunkConfiguration != nil {
		if err := s.ChunkConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.SearchConfiguration != nil {
		if err := s.SearchConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeBaseRequestChunkConfiguration struct {
	// Required for the BY_HEADING strategy. Valid values: 1 to 6. This parameter is ignored for other strategies. Headings at or above the specified level are used as split boundaries. For example, if you set this parameter to 2, both H1 and H2 headings are used as split boundaries. Deeper-level headings are not used for splitting and are retained in the chunk body. If the content within a section exceeds MaxChunkSize, the content is split by paragraph or sentence as a fallback. Documents without headings fall back to intelligent chunking.
	//
	// example:
	//
	// 2
	HeadingLevel *int32 `json:"HeadingLevel,omitempty" xml:"HeadingLevel,omitempty"`
	// The maximum token length of a single chunk.
	//
	// example:
	//
	// 512
	MaxChunkSize *int32 `json:"MaxChunkSize,omitempty" xml:"MaxChunkSize,omitempty"`
	// The overlap token length between adjacent chunks.
	//
	// example:
	//
	// 6
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// The preprocessing rules.
	PreprocessRules *UpdateKnowledgeBaseRequestChunkConfigurationPreprocessRules `json:"PreprocessRules,omitempty" xml:"PreprocessRules,omitempty" type:"Struct"`
	// The segment identifier for the LINE_BREAK strategy, such as a line feed.
	//
	// example:
	//
	// \\\\n\\\\n
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// The chunking strategy. Valid values:
	//
	// - AUTO: automatic chunking.
	//
	// - LINE_BREAK: chunking by segment identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// LINE_BREAK
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s UpdateKnowledgeBaseRequestChunkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseRequestChunkConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseRequestChunkConfiguration) GetHeadingLevel() *int32 {
	return s.HeadingLevel
}

func (s *UpdateKnowledgeBaseRequestChunkConfiguration) GetMaxChunkSize() *int32 {
	return s.MaxChunkSize
}

func (s *UpdateKnowledgeBaseRequestChunkConfiguration) GetOverlapSize() *int32 {
	return s.OverlapSize
}

func (s *UpdateKnowledgeBaseRequestChunkConfiguration) GetPreprocessRules() *UpdateKnowledgeBaseRequestChunkConfigurationPreprocessRules {
	return s.PreprocessRules
}

func (s *UpdateKnowledgeBaseRequestChunkConfiguration) GetSeparator() *string {
	return s.Separator
}

func (s *UpdateKnowledgeBaseRequestChunkConfiguration) GetStrategy() *string {
	return s.Strategy
}

func (s *UpdateKnowledgeBaseRequestChunkConfiguration) SetHeadingLevel(v int32) *UpdateKnowledgeBaseRequestChunkConfiguration {
	s.HeadingLevel = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestChunkConfiguration) SetMaxChunkSize(v int32) *UpdateKnowledgeBaseRequestChunkConfiguration {
	s.MaxChunkSize = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestChunkConfiguration) SetOverlapSize(v int32) *UpdateKnowledgeBaseRequestChunkConfiguration {
	s.OverlapSize = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestChunkConfiguration) SetPreprocessRules(v *UpdateKnowledgeBaseRequestChunkConfigurationPreprocessRules) *UpdateKnowledgeBaseRequestChunkConfiguration {
	s.PreprocessRules = v
	return s
}

func (s *UpdateKnowledgeBaseRequestChunkConfiguration) SetSeparator(v string) *UpdateKnowledgeBaseRequestChunkConfiguration {
	s.Separator = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestChunkConfiguration) SetStrategy(v string) *UpdateKnowledgeBaseRequestChunkConfiguration {
	s.Strategy = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestChunkConfiguration) Validate() error {
	if s.PreprocessRules != nil {
		if err := s.PreprocessRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeBaseRequestChunkConfigurationPreprocessRules struct {
	// Specifies whether to remove URLs and email addresses during parsing.
	//
	// example:
	//
	// false
	RemoveUrlsAndEmails *bool `json:"RemoveUrlsAndEmails,omitempty" xml:"RemoveUrlsAndEmails,omitempty"`
	// Specifies whether to replace consecutive whitespace characters (spaces, line breaks, and tab characters) with a single space.
	//
	// example:
	//
	// true
	ReplaceConsecutiveWhitespace *bool `json:"ReplaceConsecutiveWhitespace,omitempty" xml:"ReplaceConsecutiveWhitespace,omitempty"`
}

func (s UpdateKnowledgeBaseRequestChunkConfigurationPreprocessRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseRequestChunkConfigurationPreprocessRules) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseRequestChunkConfigurationPreprocessRules) GetRemoveUrlsAndEmails() *bool {
	return s.RemoveUrlsAndEmails
}

func (s *UpdateKnowledgeBaseRequestChunkConfigurationPreprocessRules) GetReplaceConsecutiveWhitespace() *bool {
	return s.ReplaceConsecutiveWhitespace
}

func (s *UpdateKnowledgeBaseRequestChunkConfigurationPreprocessRules) SetRemoveUrlsAndEmails(v bool) *UpdateKnowledgeBaseRequestChunkConfigurationPreprocessRules {
	s.RemoveUrlsAndEmails = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestChunkConfigurationPreprocessRules) SetReplaceConsecutiveWhitespace(v bool) *UpdateKnowledgeBaseRequestChunkConfigurationPreprocessRules {
	s.ReplaceConsecutiveWhitespace = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestChunkConfigurationPreprocessRules) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeBaseRequestSearchConfiguration struct {
	// The retrieve mode. Valid values:
	//
	// - KEYWORD: keyword retrieve.
	//
	// - VECTOR: vector retrieve.
	//
	// - HYBRID: hybrid retrieve.
	//
	// example:
	//
	// HYBRID
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// Takes effect only in hybrid search mode. Valid values:
	//
	// - RRF: reciprocal rank fusion.
	//
	// example:
	//
	// RRF
	RankAlgorithm *string `json:"RankAlgorithm,omitempty" xml:"RankAlgorithm,omitempty"`
	// Takes effect only in hybrid search mode. This default value is used when the search request does not specify Rerank.
	//
	// example:
	//
	// false
	RerankEnabled *bool `json:"RerankEnabled,omitempty" xml:"RerankEnabled,omitempty"`
	// The default reranking model used when the search request does not specify RerankModel. Valid values: qwen3-rerank. Default value: qwen3-rerank.
	//
	// example:
	//
	// qwen3-rerank
	RerankModel *string `json:"RerankModel,omitempty" xml:"RerankModel,omitempty"`
	// The parameter k for the RRF fusion algorithm. Default value: 60. The value must be greater than 0.
	//
	// example:
	//
	// 60
	RrfK *int32 `json:"RrfK,omitempty" xml:"RrfK,omitempty"`
	// The default number of results to return.
	//
	// example:
	//
	// 20
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// The vector path weight for the WEIGHTED fusion algorithm. Valid values: 0 to 1. The keyword path weight equals 1 minus this value. Default value: 0.7.
	//
	// example:
	//
	// 0.7
	VectorWeight *float64 `json:"VectorWeight,omitempty" xml:"VectorWeight,omitempty"`
}

func (s UpdateKnowledgeBaseRequestSearchConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseRequestSearchConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseRequestSearchConfiguration) GetMode() *string {
	return s.Mode
}

func (s *UpdateKnowledgeBaseRequestSearchConfiguration) GetRankAlgorithm() *string {
	return s.RankAlgorithm
}

func (s *UpdateKnowledgeBaseRequestSearchConfiguration) GetRerankEnabled() *bool {
	return s.RerankEnabled
}

func (s *UpdateKnowledgeBaseRequestSearchConfiguration) GetRerankModel() *string {
	return s.RerankModel
}

func (s *UpdateKnowledgeBaseRequestSearchConfiguration) GetRrfK() *int32 {
	return s.RrfK
}

func (s *UpdateKnowledgeBaseRequestSearchConfiguration) GetTopK() *int32 {
	return s.TopK
}

func (s *UpdateKnowledgeBaseRequestSearchConfiguration) GetVectorWeight() *float64 {
	return s.VectorWeight
}

func (s *UpdateKnowledgeBaseRequestSearchConfiguration) SetMode(v string) *UpdateKnowledgeBaseRequestSearchConfiguration {
	s.Mode = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestSearchConfiguration) SetRankAlgorithm(v string) *UpdateKnowledgeBaseRequestSearchConfiguration {
	s.RankAlgorithm = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestSearchConfiguration) SetRerankEnabled(v bool) *UpdateKnowledgeBaseRequestSearchConfiguration {
	s.RerankEnabled = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestSearchConfiguration) SetRerankModel(v string) *UpdateKnowledgeBaseRequestSearchConfiguration {
	s.RerankModel = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestSearchConfiguration) SetRrfK(v int32) *UpdateKnowledgeBaseRequestSearchConfiguration {
	s.RrfK = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestSearchConfiguration) SetTopK(v int32) *UpdateKnowledgeBaseRequestSearchConfiguration {
	s.TopK = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestSearchConfiguration) SetVectorWeight(v float64) *UpdateKnowledgeBaseRequestSearchConfiguration {
	s.VectorWeight = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestSearchConfiguration) Validate() error {
	return dara.Validate(s)
}
