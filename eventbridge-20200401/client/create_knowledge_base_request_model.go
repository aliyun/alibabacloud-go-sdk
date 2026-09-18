// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *CreateKnowledgeBaseRequest
	GetCatalog() *string
	SetChunkConfiguration(v *CreateKnowledgeBaseRequestChunkConfiguration) *CreateKnowledgeBaseRequest
	GetChunkConfiguration() *CreateKnowledgeBaseRequestChunkConfiguration
	SetDescription(v string) *CreateKnowledgeBaseRequest
	GetDescription() *string
	SetEmbeddingDimension(v int32) *CreateKnowledgeBaseRequest
	GetEmbeddingDimension() *int32
	SetEmbeddingModel(v string) *CreateKnowledgeBaseRequest
	GetEmbeddingModel() *string
	SetKnowledgeBaseName(v string) *CreateKnowledgeBaseRequest
	GetKnowledgeBaseName() *string
	SetMetadataSchema(v []*CreateKnowledgeBaseRequestMetadataSchema) *CreateKnowledgeBaseRequest
	GetMetadataSchema() []*CreateKnowledgeBaseRequestMetadataSchema
	SetNamespace(v string) *CreateKnowledgeBaseRequest
	GetNamespace() *string
	SetSearchConfiguration(v *CreateKnowledgeBaseRequestSearchConfiguration) *CreateKnowledgeBaseRequest
	GetSearchConfiguration() *CreateKnowledgeBaseRequestSearchConfiguration
}

type CreateKnowledgeBaseRequest struct {
	// The EventHouse catalog to which the knowledge base belongs. This parameter, together with Namespace and KnowledgeBaseName, uniquely identifies the knowledge base. This parameter cannot be modified after the knowledge base is created. System catalogs cannot be bound.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// Optional. The default chunking strategy for the knowledge base. This strategy applies only to documents uploaded after the configuration is set. If this parameter is not specified, the system default chunking strategy is used.
	ChunkConfiguration *CreateKnowledgeBaseRequestChunkConfiguration `json:"ChunkConfiguration,omitempty" xml:"ChunkConfiguration,omitempty" type:"Struct"`
	// The description of the knowledge base.
	//
	// example:
	//
	// Product documentation knowledge base
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Optional. The vector dimensions. The value is validated against the embedding model: text-embedding-v3 supports 64, 128, 256, 512, 768, and 1024. text-embedding-v4 supports 64, 128, 256, 512, 768, 1024, 1536, and 2048. qwen3.7-text-embedding supports 256, 512, 768, 1024, 1536, 2048, and 2560. qwen3.7-text-embedding-flash supports 256, 512, 768, and 1024. Default value: 1024 (the default dimension of the model). This parameter cannot be modified after the knowledge base is created. Even if the dimensions are the same, you must rebuild the knowledge base when switching models.
	//
	// example:
	//
	// 1024
	EmbeddingDimension *int32 `json:"EmbeddingDimension,omitempty" xml:"EmbeddingDimension,omitempty"`
	// Optional. The embedding model used for vectorization. This parameter cannot be modified after the knowledge base is created. Valid values: text-embedding-v3, text-embedding-v4, qwen3.7-text-embedding, and qwen3.7-text-embedding-flash. Only Bailian Tongyi models are supported. Third-party models are not supported. Default value: text-embedding-v4.
	//
	// example:
	//
	// text-embedding-v4
	EmbeddingModel *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	// The name of the knowledge base. The name must be unique within the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// product-docs
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// Declares the metadata fields of the knowledge base. When you upload documents, the metadata can contain only the fields declared here. This parameter cannot be modified after the knowledge base is created.
	//
	// example:
	//
	// [{"Name":"department","Type":"STRING"}]
	MetadataSchema []*CreateKnowledgeBaseRequestMetadataSchema `json:"MetadataSchema,omitempty" xml:"MetadataSchema,omitempty" type:"Repeated"`
	// The EventHouse namespace to which the knowledge base belongs. The namespace must belong to the specified catalog. This parameter, together with Catalog and KnowledgeBaseName, uniquely identifies the knowledge base. This parameter cannot be modified after the knowledge base is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Optional. The default search configuration at the knowledge base level. This configuration takes effect when the corresponding parameters are not specified in search requests. You can modify this configuration after creation by calling UpdateKnowledgeBase.
	SearchConfiguration *CreateKnowledgeBaseRequestSearchConfiguration `json:"SearchConfiguration,omitempty" xml:"SearchConfiguration,omitempty" type:"Struct"`
}

func (s CreateKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *CreateKnowledgeBaseRequest) GetChunkConfiguration() *CreateKnowledgeBaseRequestChunkConfiguration {
	return s.ChunkConfiguration
}

func (s *CreateKnowledgeBaseRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseRequest) GetEmbeddingDimension() *int32 {
	return s.EmbeddingDimension
}

func (s *CreateKnowledgeBaseRequest) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *CreateKnowledgeBaseRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *CreateKnowledgeBaseRequest) GetMetadataSchema() []*CreateKnowledgeBaseRequestMetadataSchema {
	return s.MetadataSchema
}

func (s *CreateKnowledgeBaseRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateKnowledgeBaseRequest) GetSearchConfiguration() *CreateKnowledgeBaseRequestSearchConfiguration {
	return s.SearchConfiguration
}

func (s *CreateKnowledgeBaseRequest) SetCatalog(v string) *CreateKnowledgeBaseRequest {
	s.Catalog = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetChunkConfiguration(v *CreateKnowledgeBaseRequestChunkConfiguration) *CreateKnowledgeBaseRequest {
	s.ChunkConfiguration = v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetDescription(v string) *CreateKnowledgeBaseRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetEmbeddingDimension(v int32) *CreateKnowledgeBaseRequest {
	s.EmbeddingDimension = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetEmbeddingModel(v string) *CreateKnowledgeBaseRequest {
	s.EmbeddingModel = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetKnowledgeBaseName(v string) *CreateKnowledgeBaseRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetMetadataSchema(v []*CreateKnowledgeBaseRequestMetadataSchema) *CreateKnowledgeBaseRequest {
	s.MetadataSchema = v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetNamespace(v string) *CreateKnowledgeBaseRequest {
	s.Namespace = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetSearchConfiguration(v *CreateKnowledgeBaseRequestSearchConfiguration) *CreateKnowledgeBaseRequest {
	s.SearchConfiguration = v
	return s
}

func (s *CreateKnowledgeBaseRequest) Validate() error {
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

type CreateKnowledgeBaseRequestChunkConfiguration struct {
	// Required for the BY_HEADING strategy. Valid values: 1 to 6. This parameter is ignored for other strategies. Headings at or above the specified level are used as split boundaries. For example, if you set this parameter to 2, both H1 and H2 headings are used as split boundaries. Deeper-level headings are not used for splitting and are retained in the chunk body. If the content within a section exceeds MaxChunkSize, the system falls back to splitting by paragraphs or sentences. Documents without headings fall back to intelligent chunking.
	//
	// example:
	//
	// 2
	HeadingLevel *int32 `json:"HeadingLevel,omitempty" xml:"HeadingLevel,omitempty"`
	// The maximum character length of a single chunk. Valid values: 1 to 6000 (characters). An error is returned if the value exceeds the limit.
	//
	// example:
	//
	// 512
	MaxChunkSize *int32 `json:"MaxChunkSize,omitempty" xml:"MaxChunkSize,omitempty"`
	// Takes effect only for the BY_LENGTH strategy. This parameter is ignored for other strategies. Specifies the overlap length (in characters) between adjacent chunks. If the value is greater than 0, the beginning of the next chunk repeats the content from the end of the previous chunk within this window. The overlap does not cause a chunk to exceed MaxChunkSize. Default value: 0, which indicates no overlap.
	//
	// example:
	//
	// 40
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// The pre-processing rules.
	PreprocessRules *CreateKnowledgeBaseRequestChunkConfigurationPreprocessRules `json:"PreprocessRules,omitempty" xml:"PreprocessRules,omitempty" type:"Struct"`
	// Required for the BY_SEPARATOR strategy. This parameter is ignored for other strategies. The system splits content by matching the literal string as a whole (not as a regular expression). The maximum length is 32 characters. Example: \\
	//
	// \\
	//
	//  for paragraph separators.
	//
	// example:
	//
	// \\\\n\\\\n
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// Valid values:
	//
	// - AUTO: Intelligent chunking (heading-aware + paragraph packing).
	//
	// - BY_LENGTH: Sliding window chunking by length. You can specify OverlapSize.
	//
	// - BY_SEPARATOR: Chunking by separator. You must specify Separator.
	//
	// - BY_HEADING: Chunking by heading level. You must specify HeadingLevel.
	//
	// This parameter is required.
	//
	// example:
	//
	// BY_SEPARATOR
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s CreateKnowledgeBaseRequestChunkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequestChunkConfiguration) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequestChunkConfiguration) GetHeadingLevel() *int32 {
	return s.HeadingLevel
}

func (s *CreateKnowledgeBaseRequestChunkConfiguration) GetMaxChunkSize() *int32 {
	return s.MaxChunkSize
}

func (s *CreateKnowledgeBaseRequestChunkConfiguration) GetOverlapSize() *int32 {
	return s.OverlapSize
}

func (s *CreateKnowledgeBaseRequestChunkConfiguration) GetPreprocessRules() *CreateKnowledgeBaseRequestChunkConfigurationPreprocessRules {
	return s.PreprocessRules
}

func (s *CreateKnowledgeBaseRequestChunkConfiguration) GetSeparator() *string {
	return s.Separator
}

func (s *CreateKnowledgeBaseRequestChunkConfiguration) GetStrategy() *string {
	return s.Strategy
}

func (s *CreateKnowledgeBaseRequestChunkConfiguration) SetHeadingLevel(v int32) *CreateKnowledgeBaseRequestChunkConfiguration {
	s.HeadingLevel = &v
	return s
}

func (s *CreateKnowledgeBaseRequestChunkConfiguration) SetMaxChunkSize(v int32) *CreateKnowledgeBaseRequestChunkConfiguration {
	s.MaxChunkSize = &v
	return s
}

func (s *CreateKnowledgeBaseRequestChunkConfiguration) SetOverlapSize(v int32) *CreateKnowledgeBaseRequestChunkConfiguration {
	s.OverlapSize = &v
	return s
}

func (s *CreateKnowledgeBaseRequestChunkConfiguration) SetPreprocessRules(v *CreateKnowledgeBaseRequestChunkConfigurationPreprocessRules) *CreateKnowledgeBaseRequestChunkConfiguration {
	s.PreprocessRules = v
	return s
}

func (s *CreateKnowledgeBaseRequestChunkConfiguration) SetSeparator(v string) *CreateKnowledgeBaseRequestChunkConfiguration {
	s.Separator = &v
	return s
}

func (s *CreateKnowledgeBaseRequestChunkConfiguration) SetStrategy(v string) *CreateKnowledgeBaseRequestChunkConfiguration {
	s.Strategy = &v
	return s
}

func (s *CreateKnowledgeBaseRequestChunkConfiguration) Validate() error {
	if s.PreprocessRules != nil {
		if err := s.PreprocessRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKnowledgeBaseRequestChunkConfigurationPreprocessRules struct {
	// Specifies whether to remove URLs and email addresses during parsing.
	//
	// example:
	//
	// false
	RemoveUrlsAndEmails *bool `json:"RemoveUrlsAndEmails,omitempty" xml:"RemoveUrlsAndEmails,omitempty"`
	// Specifies whether to replace consecutive whitespace characters (spaces, line breaks, and tabs) with a single space.
	//
	// example:
	//
	// true
	ReplaceConsecutiveWhitespace *bool `json:"ReplaceConsecutiveWhitespace,omitempty" xml:"ReplaceConsecutiveWhitespace,omitempty"`
}

func (s CreateKnowledgeBaseRequestChunkConfigurationPreprocessRules) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequestChunkConfigurationPreprocessRules) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequestChunkConfigurationPreprocessRules) GetRemoveUrlsAndEmails() *bool {
	return s.RemoveUrlsAndEmails
}

func (s *CreateKnowledgeBaseRequestChunkConfigurationPreprocessRules) GetReplaceConsecutiveWhitespace() *bool {
	return s.ReplaceConsecutiveWhitespace
}

func (s *CreateKnowledgeBaseRequestChunkConfigurationPreprocessRules) SetRemoveUrlsAndEmails(v bool) *CreateKnowledgeBaseRequestChunkConfigurationPreprocessRules {
	s.RemoveUrlsAndEmails = &v
	return s
}

func (s *CreateKnowledgeBaseRequestChunkConfigurationPreprocessRules) SetReplaceConsecutiveWhitespace(v bool) *CreateKnowledgeBaseRequestChunkConfigurationPreprocessRules {
	s.ReplaceConsecutiveWhitespace = &v
	return s
}

func (s *CreateKnowledgeBaseRequestChunkConfigurationPreprocessRules) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseRequestMetadataSchema struct {
	// The name of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// department
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Valid values: STRING, LONG, DOUBLE, BOOLEAN, and DATETIME.
	//
	// This parameter is required.
	//
	// example:
	//
	// STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// When ValueMode is set to CONSTANT, this parameter specifies a fixed value. An empty value indicates that the value can be assigned during upload. When ValueMode is set to SYSTEM_VARIABLE, this parameter specifies a system variable name, such as DOCUMENT_NAME, FILE_TYPE, FILE_SIZE, DOCUMENT_UPLOAD_TIME, SOURCE_TYPE, SOURCE_URI, or SOURCE_MODIFIED_TIME.
	//
	// example:
	//
	// EventHouse
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// Valid values:
	//
	// - CONSTANT: Constant. If Value is not empty, all documents use the fixed value. If Value is empty, the value can be assigned during upload.
	//
	// - SYSTEM_VARIABLE: System variable. Value specifies the variable name. The system automatically generates the value, and the value cannot be overridden during upload.
	//
	// This parameter is required.
	//
	// example:
	//
	// CONSTANT
	ValueMode *string `json:"ValueMode,omitempty" xml:"ValueMode,omitempty"`
}

func (s CreateKnowledgeBaseRequestMetadataSchema) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequestMetadataSchema) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequestMetadataSchema) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseRequestMetadataSchema) GetType() *string {
	return s.Type
}

func (s *CreateKnowledgeBaseRequestMetadataSchema) GetValue() *string {
	return s.Value
}

func (s *CreateKnowledgeBaseRequestMetadataSchema) GetValueMode() *string {
	return s.ValueMode
}

func (s *CreateKnowledgeBaseRequestMetadataSchema) SetName(v string) *CreateKnowledgeBaseRequestMetadataSchema {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseRequestMetadataSchema) SetType(v string) *CreateKnowledgeBaseRequestMetadataSchema {
	s.Type = &v
	return s
}

func (s *CreateKnowledgeBaseRequestMetadataSchema) SetValue(v string) *CreateKnowledgeBaseRequestMetadataSchema {
	s.Value = &v
	return s
}

func (s *CreateKnowledgeBaseRequestMetadataSchema) SetValueMode(v string) *CreateKnowledgeBaseRequestMetadataSchema {
	s.ValueMode = &v
	return s
}

func (s *CreateKnowledgeBaseRequestMetadataSchema) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseRequestSearchConfiguration struct {
	// Valid values:
	//
	// - KEYWORD: Keyword retrieval.
	//
	// - VECTOR: Vector retrieval.
	//
	// - HYBRID: Hybrid retrieval.
	//
	// example:
	//
	// HYBRID
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// Takes effect only in hybrid search mode. Valid values:
	//
	// - RRF: Reciprocal rank fusion.
	//
	// - WEIGHTED: Weighted normalization fusion. Use this value together with VectorWeight.
	//
	// Default value: RRF.
	//
	// example:
	//
	// RRF
	RankAlgorithm *string `json:"RankAlgorithm,omitempty" xml:"RankAlgorithm,omitempty"`
	// Takes effect only in hybrid search mode. Specifies whether to enable reranking by default when the search request does not specify a Rerank parameter.
	//
	// example:
	//
	// false
	RerankEnabled *bool `json:"RerankEnabled,omitempty" xml:"RerankEnabled,omitempty"`
	// The default reranking model used when the search request does not specify a RerankModel parameter. Valid values: qwen3-rerank. Default value: qwen3-rerank.
	//
	// example:
	//
	// qwen3-rerank
	RerankModel *string `json:"RerankModel,omitempty" xml:"RerankModel,omitempty"`
	// The k parameter for the RRF fusion algorithm. The value must be greater than 0. Default value: 60.
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
	// The vector weight for the WEIGHTED fusion algorithm. Valid values: 0 to 1. The keyword weight equals 1 minus this value. Default value: 0.7.
	//
	// example:
	//
	// 0.7
	VectorWeight *float64 `json:"VectorWeight,omitempty" xml:"VectorWeight,omitempty"`
}

func (s CreateKnowledgeBaseRequestSearchConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequestSearchConfiguration) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequestSearchConfiguration) GetMode() *string {
	return s.Mode
}

func (s *CreateKnowledgeBaseRequestSearchConfiguration) GetRankAlgorithm() *string {
	return s.RankAlgorithm
}

func (s *CreateKnowledgeBaseRequestSearchConfiguration) GetRerankEnabled() *bool {
	return s.RerankEnabled
}

func (s *CreateKnowledgeBaseRequestSearchConfiguration) GetRerankModel() *string {
	return s.RerankModel
}

func (s *CreateKnowledgeBaseRequestSearchConfiguration) GetRrfK() *int32 {
	return s.RrfK
}

func (s *CreateKnowledgeBaseRequestSearchConfiguration) GetTopK() *int32 {
	return s.TopK
}

func (s *CreateKnowledgeBaseRequestSearchConfiguration) GetVectorWeight() *float64 {
	return s.VectorWeight
}

func (s *CreateKnowledgeBaseRequestSearchConfiguration) SetMode(v string) *CreateKnowledgeBaseRequestSearchConfiguration {
	s.Mode = &v
	return s
}

func (s *CreateKnowledgeBaseRequestSearchConfiguration) SetRankAlgorithm(v string) *CreateKnowledgeBaseRequestSearchConfiguration {
	s.RankAlgorithm = &v
	return s
}

func (s *CreateKnowledgeBaseRequestSearchConfiguration) SetRerankEnabled(v bool) *CreateKnowledgeBaseRequestSearchConfiguration {
	s.RerankEnabled = &v
	return s
}

func (s *CreateKnowledgeBaseRequestSearchConfiguration) SetRerankModel(v string) *CreateKnowledgeBaseRequestSearchConfiguration {
	s.RerankModel = &v
	return s
}

func (s *CreateKnowledgeBaseRequestSearchConfiguration) SetRrfK(v int32) *CreateKnowledgeBaseRequestSearchConfiguration {
	s.RrfK = &v
	return s
}

func (s *CreateKnowledgeBaseRequestSearchConfiguration) SetTopK(v int32) *CreateKnowledgeBaseRequestSearchConfiguration {
	s.TopK = &v
	return s
}

func (s *CreateKnowledgeBaseRequestSearchConfiguration) SetVectorWeight(v float64) *CreateKnowledgeBaseRequestSearchConfiguration {
	s.VectorWeight = &v
	return s
}

func (s *CreateKnowledgeBaseRequestSearchConfiguration) Validate() error {
	return dara.Validate(s)
}
