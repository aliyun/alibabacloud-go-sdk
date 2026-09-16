// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeBase interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *KnowledgeBase
	GetCatalog() *string
	SetChunkConfiguration(v *KnowledgeBaseChunkConfiguration) *KnowledgeBase
	GetChunkConfiguration() *KnowledgeBaseChunkConfiguration
	SetCreatedAt(v string) *KnowledgeBase
	GetCreatedAt() *string
	SetDescription(v string) *KnowledgeBase
	GetDescription() *string
	SetEmbeddingDimension(v int32) *KnowledgeBase
	GetEmbeddingDimension() *int32
	SetEmbeddingModel(v string) *KnowledgeBase
	GetEmbeddingModel() *string
	SetFailureReason(v string) *KnowledgeBase
	GetFailureReason() *string
	SetKnowledgeBaseName(v string) *KnowledgeBase
	GetKnowledgeBaseName() *string
	SetMetadataSchema(v []*MetadataSchemaField) *KnowledgeBase
	GetMetadataSchema() []*MetadataSchemaField
	SetNamespace(v string) *KnowledgeBase
	GetNamespace() *string
	SetSearchConfiguration(v *KnowledgeBaseSearchConfiguration) *KnowledgeBase
	GetSearchConfiguration() *KnowledgeBaseSearchConfiguration
	SetStatus(v string) *KnowledgeBase
	GetStatus() *string
	SetUpdatedAt(v string) *KnowledgeBase
	GetUpdatedAt() *string
}

type KnowledgeBase struct {
	// The EventHouse data catalog to which the knowledge base belongs. This value cannot be modified after the knowledge base is created.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The default chunking strategy of the knowledge base. This configuration takes effect only for documents uploaded after the configuration is updated. Existing documents are not re-chunked.
	ChunkConfiguration *KnowledgeBaseChunkConfiguration `json:"ChunkConfiguration,omitempty" xml:"ChunkConfiguration,omitempty" type:"Struct"`
	// The time when the knowledge base was created.
	//
	// example:
	//
	// 2026-08-24T10:00:00Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The description of the knowledge base.
	//
	// example:
	//
	// Product documentation knowledge base
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The embedding vector dimension specified during creation or the default dimension of the model. This value cannot be modified after the knowledge base is created.
	//
	// example:
	//
	// 1024
	EmbeddingDimension *int32 `json:"EmbeddingDimension,omitempty" xml:"EmbeddingDimension,omitempty"`
	// The embedding model specified during creation. This value cannot be modified after the knowledge base is created.
	//
	// example:
	//
	// text-embedding-v4
	EmbeddingModel *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	// The brief reason for the most recent creation or deletion failure. This parameter is returned only when the status is CREATE_FAILED or DELETE_FAILED.
	//
	// example:
	//
	// OssException: BucketAlreadyExists ...
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	// The name of the knowledge base, which is unique within the namespace.
	//
	// example:
	//
	// product-docs
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// The metadata fields declared when the knowledge base was created. These fields cannot be modified after the knowledge base is created.
	//
	// example:
	//
	// [{"Name":"department","Type":"STRING"}]
	MetadataSchema []*MetadataSchemaField `json:"MetadataSchema,omitempty" xml:"MetadataSchema,omitempty" type:"Repeated"`
	// The EventHouse namespace to which the knowledge base belongs. This value cannot be modified after the knowledge base is created.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The default search configuration at the knowledge base level. This configuration takes effect when the corresponding parameters are not specified in a search request. You can modify this configuration by calling the UpdateKnowledgeBase operation.
	SearchConfiguration *KnowledgeBaseSearchConfiguration `json:"SearchConfiguration,omitempty" xml:"SearchConfiguration,omitempty" type:"Struct"`
	// The current status of the knowledge base. Valid values:
	//
	// - CREATING: The knowledge base is being created.
	//
	// - ACTIVE: The knowledge base is available.
	//
	// - CREATE_FAILED: The knowledge base failed to be created.
	//
	// - DELETING: The knowledge base is being deleted.
	//
	// - DELETE_FAILED: The knowledge base failed to be deleted.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the knowledge base was last updated.
	//
	// example:
	//
	// 2026-08-24T10:00:00Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s KnowledgeBase) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBase) GoString() string {
	return s.String()
}

func (s *KnowledgeBase) GetCatalog() *string {
	return s.Catalog
}

func (s *KnowledgeBase) GetChunkConfiguration() *KnowledgeBaseChunkConfiguration {
	return s.ChunkConfiguration
}

func (s *KnowledgeBase) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *KnowledgeBase) GetDescription() *string {
	return s.Description
}

func (s *KnowledgeBase) GetEmbeddingDimension() *int32 {
	return s.EmbeddingDimension
}

func (s *KnowledgeBase) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *KnowledgeBase) GetFailureReason() *string {
	return s.FailureReason
}

func (s *KnowledgeBase) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *KnowledgeBase) GetMetadataSchema() []*MetadataSchemaField {
	return s.MetadataSchema
}

func (s *KnowledgeBase) GetNamespace() *string {
	return s.Namespace
}

func (s *KnowledgeBase) GetSearchConfiguration() *KnowledgeBaseSearchConfiguration {
	return s.SearchConfiguration
}

func (s *KnowledgeBase) GetStatus() *string {
	return s.Status
}

func (s *KnowledgeBase) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *KnowledgeBase) SetCatalog(v string) *KnowledgeBase {
	s.Catalog = &v
	return s
}

func (s *KnowledgeBase) SetChunkConfiguration(v *KnowledgeBaseChunkConfiguration) *KnowledgeBase {
	s.ChunkConfiguration = v
	return s
}

func (s *KnowledgeBase) SetCreatedAt(v string) *KnowledgeBase {
	s.CreatedAt = &v
	return s
}

func (s *KnowledgeBase) SetDescription(v string) *KnowledgeBase {
	s.Description = &v
	return s
}

func (s *KnowledgeBase) SetEmbeddingDimension(v int32) *KnowledgeBase {
	s.EmbeddingDimension = &v
	return s
}

func (s *KnowledgeBase) SetEmbeddingModel(v string) *KnowledgeBase {
	s.EmbeddingModel = &v
	return s
}

func (s *KnowledgeBase) SetFailureReason(v string) *KnowledgeBase {
	s.FailureReason = &v
	return s
}

func (s *KnowledgeBase) SetKnowledgeBaseName(v string) *KnowledgeBase {
	s.KnowledgeBaseName = &v
	return s
}

func (s *KnowledgeBase) SetMetadataSchema(v []*MetadataSchemaField) *KnowledgeBase {
	s.MetadataSchema = v
	return s
}

func (s *KnowledgeBase) SetNamespace(v string) *KnowledgeBase {
	s.Namespace = &v
	return s
}

func (s *KnowledgeBase) SetSearchConfiguration(v *KnowledgeBaseSearchConfiguration) *KnowledgeBase {
	s.SearchConfiguration = v
	return s
}

func (s *KnowledgeBase) SetStatus(v string) *KnowledgeBase {
	s.Status = &v
	return s
}

func (s *KnowledgeBase) SetUpdatedAt(v string) *KnowledgeBase {
	s.UpdatedAt = &v
	return s
}

func (s *KnowledgeBase) Validate() error {
	if s.ChunkConfiguration != nil {
		if err := s.ChunkConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.MetadataSchema != nil {
		for _, item := range s.MetadataSchema {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SearchConfiguration != nil {
		if err := s.SearchConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KnowledgeBaseChunkConfiguration struct {
	// The heading level (1 to 6) used for splitting in the BY_HEADING strategy. Headings at or above this level serve as split boundaries. Deeper-level headings are retained in the chunk body.
	//
	// example:
	//
	// 2
	HeadingLevel *int32 `json:"HeadingLevel,omitempty" xml:"HeadingLevel,omitempty"`
	// The maximum character length of a single chunk. Starting from revision 22, this value is character-based. Valid values: 1 to 6000.
	//
	// example:
	//
	// 600
	MaxChunkSize *int32 `json:"MaxChunkSize,omitempty" xml:"MaxChunkSize,omitempty"`
	// The overlap character length between adjacent chunks. This parameter takes effect only for the BY_LENGTH strategy. When the value is greater than 0, the beginning of the next chunk repeats the content from the end of the previous chunk within this window. The overlap does not cause a chunk to exceed MaxChunkSize. A value of 0 indicates no overlap.
	//
	// example:
	//
	// 40
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// The preprocessing rules that take effect during document parsing.
	PreprocessRules *KnowledgeBaseChunkConfigurationPreprocessRules `json:"PreprocessRules,omitempty" xml:"PreprocessRules,omitempty" type:"Struct"`
	// The separator used in the BY_SEPARATOR strategy. The separator is matched as a literal string (not a regular expression). The maximum length is 32 characters.
	//
	// example:
	//
	// \\\\n\\\\n
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// The chunking strategy. Valid values:
	//
	// - AUTO: intelligent splitting (heading-aware + paragraph packing).
	//
	// - BY_LENGTH: sliding window splitting by length. You can specify OverlapSize.
	//
	// - BY_SEPARATOR: splitting by separator. You must specify Separator.
	//
	// - BY_HEADING: splitting by heading level. You must specify HeadingLevel.
	//
	// example:
	//
	// BY_SEPARATOR
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s KnowledgeBaseChunkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseChunkConfiguration) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseChunkConfiguration) GetHeadingLevel() *int32 {
	return s.HeadingLevel
}

func (s *KnowledgeBaseChunkConfiguration) GetMaxChunkSize() *int32 {
	return s.MaxChunkSize
}

func (s *KnowledgeBaseChunkConfiguration) GetOverlapSize() *int32 {
	return s.OverlapSize
}

func (s *KnowledgeBaseChunkConfiguration) GetPreprocessRules() *KnowledgeBaseChunkConfigurationPreprocessRules {
	return s.PreprocessRules
}

func (s *KnowledgeBaseChunkConfiguration) GetSeparator() *string {
	return s.Separator
}

func (s *KnowledgeBaseChunkConfiguration) GetStrategy() *string {
	return s.Strategy
}

func (s *KnowledgeBaseChunkConfiguration) SetHeadingLevel(v int32) *KnowledgeBaseChunkConfiguration {
	s.HeadingLevel = &v
	return s
}

func (s *KnowledgeBaseChunkConfiguration) SetMaxChunkSize(v int32) *KnowledgeBaseChunkConfiguration {
	s.MaxChunkSize = &v
	return s
}

func (s *KnowledgeBaseChunkConfiguration) SetOverlapSize(v int32) *KnowledgeBaseChunkConfiguration {
	s.OverlapSize = &v
	return s
}

func (s *KnowledgeBaseChunkConfiguration) SetPreprocessRules(v *KnowledgeBaseChunkConfigurationPreprocessRules) *KnowledgeBaseChunkConfiguration {
	s.PreprocessRules = v
	return s
}

func (s *KnowledgeBaseChunkConfiguration) SetSeparator(v string) *KnowledgeBaseChunkConfiguration {
	s.Separator = &v
	return s
}

func (s *KnowledgeBaseChunkConfiguration) SetStrategy(v string) *KnowledgeBaseChunkConfiguration {
	s.Strategy = &v
	return s
}

func (s *KnowledgeBaseChunkConfiguration) Validate() error {
	if s.PreprocessRules != nil {
		if err := s.PreprocessRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KnowledgeBaseChunkConfigurationPreprocessRules struct {
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

func (s KnowledgeBaseChunkConfigurationPreprocessRules) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseChunkConfigurationPreprocessRules) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseChunkConfigurationPreprocessRules) GetRemoveUrlsAndEmails() *bool {
	return s.RemoveUrlsAndEmails
}

func (s *KnowledgeBaseChunkConfigurationPreprocessRules) GetReplaceConsecutiveWhitespace() *bool {
	return s.ReplaceConsecutiveWhitespace
}

func (s *KnowledgeBaseChunkConfigurationPreprocessRules) SetRemoveUrlsAndEmails(v bool) *KnowledgeBaseChunkConfigurationPreprocessRules {
	s.RemoveUrlsAndEmails = &v
	return s
}

func (s *KnowledgeBaseChunkConfigurationPreprocessRules) SetReplaceConsecutiveWhitespace(v bool) *KnowledgeBaseChunkConfigurationPreprocessRules {
	s.ReplaceConsecutiveWhitespace = &v
	return s
}

func (s *KnowledgeBaseChunkConfigurationPreprocessRules) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseSearchConfiguration struct {
	// The retrieval mode. Valid values:
	//
	// - KEYWORD: keyword retrieval.
	//
	// - VECTOR: vector retrieval.
	//
	// - HYBRID: hybrid retrieval.
	//
	// example:
	//
	// HYBRID
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The fusion algorithm for hybrid search. This parameter takes effect only in hybrid search mode. Valid values:
	//
	// - RRF: reciprocal rank fusion.
	//
	// - WEIGHTED: weighted normalization fusion. Use this value together with VectorWeight.
	//
	// Default value: RRF.
	//
	// example:
	//
	// RRF
	RankAlgorithm *string `json:"RankAlgorithm,omitempty" xml:"RankAlgorithm,omitempty"`
	// Specifies whether reranking is enabled by default. This parameter takes effect for all search modes (KEYWORD, VECTOR, and HYBRID). This default value is used when the Rerank parameter is not specified in a search request.
	//
	// example:
	//
	// false
	RerankEnabled *bool `json:"RerankEnabled,omitempty" xml:"RerankEnabled,omitempty"`
	// The default reranking model used when the RerankModel parameter is not specified in a search request. Valid values: qwen3-rerank, gte-rerank-v2, and qwen3-vl-rerank. Default value: qwen3-rerank. Score distributions vary across models and cannot be compared. Use the same model consistently within a knowledge base.
	//
	// example:
	//
	// qwen3-rerank
	RerankModel *string `json:"RerankModel,omitempty" xml:"RerankModel,omitempty"`
	// The k parameter of the RRF fusion algorithm. The value must be greater than 0. Default value: 60.
	//
	// example:
	//
	// 60
	RrfK *int32 `json:"RrfK,omitempty" xml:"RrfK,omitempty"`
	// The maximum number of results returned by default for a search request.
	//
	// example:
	//
	// 20
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// The weight of the vector path in the WEIGHTED fusion algorithm. Valid values: 0 to 1. The keyword path weight equals 1 minus this value. Default value: 0.7.
	//
	// example:
	//
	// 0.7
	VectorWeight *float64 `json:"VectorWeight,omitempty" xml:"VectorWeight,omitempty"`
}

func (s KnowledgeBaseSearchConfiguration) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseSearchConfiguration) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseSearchConfiguration) GetMode() *string {
	return s.Mode
}

func (s *KnowledgeBaseSearchConfiguration) GetRankAlgorithm() *string {
	return s.RankAlgorithm
}

func (s *KnowledgeBaseSearchConfiguration) GetRerankEnabled() *bool {
	return s.RerankEnabled
}

func (s *KnowledgeBaseSearchConfiguration) GetRerankModel() *string {
	return s.RerankModel
}

func (s *KnowledgeBaseSearchConfiguration) GetRrfK() *int32 {
	return s.RrfK
}

func (s *KnowledgeBaseSearchConfiguration) GetTopK() *int32 {
	return s.TopK
}

func (s *KnowledgeBaseSearchConfiguration) GetVectorWeight() *float64 {
	return s.VectorWeight
}

func (s *KnowledgeBaseSearchConfiguration) SetMode(v string) *KnowledgeBaseSearchConfiguration {
	s.Mode = &v
	return s
}

func (s *KnowledgeBaseSearchConfiguration) SetRankAlgorithm(v string) *KnowledgeBaseSearchConfiguration {
	s.RankAlgorithm = &v
	return s
}

func (s *KnowledgeBaseSearchConfiguration) SetRerankEnabled(v bool) *KnowledgeBaseSearchConfiguration {
	s.RerankEnabled = &v
	return s
}

func (s *KnowledgeBaseSearchConfiguration) SetRerankModel(v string) *KnowledgeBaseSearchConfiguration {
	s.RerankModel = &v
	return s
}

func (s *KnowledgeBaseSearchConfiguration) SetRrfK(v int32) *KnowledgeBaseSearchConfiguration {
	s.RrfK = &v
	return s
}

func (s *KnowledgeBaseSearchConfiguration) SetTopK(v int32) *KnowledgeBaseSearchConfiguration {
	s.TopK = &v
	return s
}

func (s *KnowledgeBaseSearchConfiguration) SetVectorWeight(v float64) *KnowledgeBaseSearchConfiguration {
	s.VectorWeight = &v
	return s
}

func (s *KnowledgeBaseSearchConfiguration) Validate() error {
	return dara.Validate(s)
}
