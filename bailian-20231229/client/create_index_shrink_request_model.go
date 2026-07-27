// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndexShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryIdsShrink(v string) *CreateIndexShrinkRequest
	GetCategoryIdsShrink() *string
	SetChunkSize(v int32) *CreateIndexShrinkRequest
	GetChunkSize() *int32
	SetColumnsShrink(v string) *CreateIndexShrinkRequest
	GetColumnsShrink() *string
	SetCreateIndexType(v string) *CreateIndexShrinkRequest
	GetCreateIndexType() *string
	SetDescription(v string) *CreateIndexShrinkRequest
	GetDescription() *string
	SetDocumentIdsShrink(v string) *CreateIndexShrinkRequest
	GetDocumentIdsShrink() *string
	SetEmbeddingModelName(v string) *CreateIndexShrinkRequest
	GetEmbeddingModelName() *string
	SetEnableRewrite(v bool) *CreateIndexShrinkRequest
	GetEnableRewrite() *bool
	SetName(v string) *CreateIndexShrinkRequest
	GetName() *string
	SetOverlapSize(v int32) *CreateIndexShrinkRequest
	GetOverlapSize() *int32
	SetRerankInstruct(v string) *CreateIndexShrinkRequest
	GetRerankInstruct() *string
	SetRerankMinScore(v float64) *CreateIndexShrinkRequest
	GetRerankMinScore() *float64
	SetRerankMode(v string) *CreateIndexShrinkRequest
	GetRerankMode() *string
	SetRerankModelName(v string) *CreateIndexShrinkRequest
	GetRerankModelName() *string
	SetSeparator(v string) *CreateIndexShrinkRequest
	GetSeparator() *string
	SetSinkInstanceId(v string) *CreateIndexShrinkRequest
	GetSinkInstanceId() *string
	SetSinkRegion(v string) *CreateIndexShrinkRequest
	GetSinkRegion() *string
	SetSinkType(v string) *CreateIndexShrinkRequest
	GetSinkType() *string
	SetSourceType(v string) *CreateIndexShrinkRequest
	GetSourceType() *string
	SetStructureType(v string) *CreateIndexShrinkRequest
	GetStructureType() *string
	SetTableIdsShrink(v string) *CreateIndexShrinkRequest
	GetTableIdsShrink() *string
	SetChannelType(v string) *CreateIndexShrinkRequest
	GetChannelType() *string
	SetChunkMode(v string) *CreateIndexShrinkRequest
	GetChunkMode() *string
	SetConnectId(v string) *CreateIndexShrinkRequest
	GetConnectId() *string
	SetDatabase(v string) *CreateIndexShrinkRequest
	GetDatabase() *string
	SetDatasourceCode(v string) *CreateIndexShrinkRequest
	GetDatasourceCode() *string
	SetEnableHeaders(v bool) *CreateIndexShrinkRequest
	GetEnableHeaders() *bool
	SetKnowledgeScene(v string) *CreateIndexShrinkRequest
	GetKnowledgeScene() *string
	SetKnowledgeType(v string) *CreateIndexShrinkRequest
	GetKnowledgeType() *string
	SetMetaExtractColumnsShrink(v string) *CreateIndexShrinkRequest
	GetMetaExtractColumnsShrink() *string
	SetPipelineCommercialCu(v int32) *CreateIndexShrinkRequest
	GetPipelineCommercialCu() *int32
	SetPipelineCommercialType(v string) *CreateIndexShrinkRequest
	GetPipelineCommercialType() *string
	SetPipelineRetrieveRateLimitStrategy(v string) *CreateIndexShrinkRequest
	GetPipelineRetrieveRateLimitStrategy() *string
	SetTable(v string) *CreateIndexShrinkRequest
	GetTable() *string
}

type CreateIndexShrinkRequest struct {
	// The list of category IDs to import when creating the knowledge base. All files under the specified categories are imported. We recommend importing no more than 500 files. For remaining files, call the **SubmitIndexAddDocumentsJob*	- operation to continue importing.
	CategoryIdsShrink *string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty"`
	// <props="china">
	//
	// The chunk size, which specifies the maximum number of characters per text chunk. When this length is exceeded:
	//
	// - **Intelligent chunking*	- (when `chunkMode` is not specified): The text is likely to be truncated.
	//
	// - **Custom chunking*	- (when `chunkMode` is specified): The text is forcibly truncated.
	//
	//
	//
	// <props="intl">
	//
	// The chunk size, which specifies the maximum number of characters per text chunk. When this length is exceeded, the text is likely to be truncated.
	//
	//
	//
	// Value range: [1-6000]. If not specified, the default value is 500.
	//
	// > If `ChunkSize` is set to a value less than 100, you must also set `OverlapSize`. You can also leave both parameters unspecified, and the system uses default values.
	//
	// example:
	//
	// 128
	ChunkSize *int32 `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// <props="china">
	//
	// The structure of the data table (column names, types, etc.).
	//
	//
	// <props="intl">
	//
	// > This parameter is not available. Do not pass this parameter.
	//
	// >
	ColumnsShrink *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	// > This parameter is not available. Do not pass this parameter.
	//
	// >
	//
	// example:
	//
	// standard
	CreateIndexType *string `json:"CreateIndexType,omitempty" xml:"CreateIndexType,omitempty"`
	// The knowledge base description. The description can be up to 1000 characters in length.
	//
	// Default value: empty.
	//
	// example:
	//
	// The enterprise help document library includes important materials such as company policies and product catalogs.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of files to import when creating the knowledge base. Specify file IDs here. We recommend importing no more than 10,000 files. For remaining files, call the **SubmitIndexAddDocumentsJob*	- operation to continue importing.
	DocumentIdsShrink *string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty"`
	// <props="china">
	//
	// The embedding model used by the knowledge base. The embedding model transforms the original input prompt and knowledge text into numerical vectors for similarity comparison. The text-embedding-v4 model is a comprehensive upgrade over text-embedding-v3 in terms of language support, code snippet quantization, and vector dimensions selection, and is suitable for most scenarios. For more information, see [Vectorization](https://help.aliyun.com/document_detail/2842587.html). Valid values:
	//
	// - text-embedding-v4
	//
	// - text-embedding-v3
	//
	// Default value: empty, which uses the text-embedding-v3 model.
	//
	//
	//
	//
	// <props="intl">
	//
	// The embedding model used by the knowledge base. The embedding model transforms the original input prompt and knowledge text into numerical vectors for similarity comparison. The default text-embedding-v2 model (cannot be changed) supports Chinese, English, and multiple other languages, and performs normalization on vector results. For more information, see [Vectorization](https://help.aliyun.com/document_detail/2842587.html). Valid values:
	//
	// - text-embedding-v2
	//
	// Default value: empty, which uses the text-embedding-v2 model.
	//
	// example:
	//
	// text-embedding-v4
	EmbeddingModelName *string `json:"EmbeddingModelName,omitempty" xml:"EmbeddingModelName,omitempty"`
	// Specifies whether to enable multi-turn conversation rewriting. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// If not specified, this feature is enabled by default.
	//
	// example:
	//
	// true
	EnableRewrite *bool `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	// The knowledge base name. The name must be 1 to 20 characters in length and can contain Chinese characters, letters, digits, underscores (_), hyphens (-), periods (.), and colons (:).
	//
	// This parameter is required.
	//
	// example:
	//
	// EnterpriseHelpDocLibrary.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The chunk overlap size, which specifies the number of overlapping characters between the current text chunk and the previous text chunk. Value range: [0-1024].
	//
	// If not specified, the default value is 100.
	//
	// >`OverlapSize` must be less than `ChunkSize`. Otherwise, chunking exceptions occur.
	//
	// example:
	//
	// 16
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// <props="intl">This parameter is not available. Do not pass this parameter.
	//
	// <props="china">A natural language instruction for fine-grained control of the reranking model\\"s behavior.
	//
	// <notice>This parameter takes effect only when rerank_mode is set to "custom".
	RerankInstruct *string `json:"RerankInstruct,omitempty" xml:"RerankInstruct,omitempty"`
	// The similarity threshold. Only text chunks with similarity scores exceeding this value are recalled. This parameter filters the text chunks returned by the reranking model. Value range: [0.01-1.00].
	//
	// If not specified, the default value is 0.01.
	//
	// example:
	//
	// 0.20
	RerankMinScore *float64 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// <props="china">
	//
	// Specifies the instruction intervention mode for the reranking model to determine its scoring preference.
	//
	// **Valid values:**
	//
	// - **qa**: (Default) Q&A mode. The model tends to assign higher scores to candidates that directly answer the query. Recommended for Q&A scenarios.
	//
	// - **similar**: Similarity mode. The model tends to assign higher scores to candidates with high content consistency with the query. Recommended for matching and retrieval scenarios.
	//
	// - **custom**: Custom mode. The model\\"s ranking behavior is determined by the instruction in the rerank_instruct parameter.
	//
	//
	//
	// <props="intl">This parameter is not available. Do not pass this parameter.
	//
	// [_single.params.RerankMode.enum.similar: 相似模式。]similar: Similarity mode.
	//
	// [_single.params.RerankMode.enum.custom: 自定义模式。]custom: Custom mode.
	//
	// [_single.params.RerankMode.enum.qa:（默认值） 问答模式。]qa: (Default) Q&A mode.
	//
	// [parameters.33.schema.enumValueTitles.similar: 相似模式。]similar: Similarity mode.
	//
	// [parameters.33.schema.enumValueTitles.custom: 自定义模式。]custom: Custom mode.
	//
	// [parameters.33.schema.enumValueTitles.qa:（默认值） 问答模式。]qa: (Default) Q&A mode.
	//
	// example:
	//
	// qa
	RerankMode *string `json:"RerankMode,omitempty" xml:"RerankMode,omitempty"`
	// The reranking model used by the knowledge base. The reranking model is an external scoring system that calculates the similarity score between the user query and each text chunk in the knowledge base, sorts them in descending order, and returns the top K text chunks with the highest scores. Valid values:
	//
	//
	// <props="china">
	//
	// - qwen3-rerank-hybrid: qwen3-rerank(hybrid) reranking.
	//
	// - qwen3-rerank: qwen3-rerank reranking.
	//
	// - gte-rerank-hybrid: gte-rerank(hybrid) reranking.
	//
	// - gte-rerank: gte-rerank reranking.
	//
	//
	//
	// <props="intl">
	//
	// - gte-rerank-hybrid: official reranking.
	//
	// - gte-rerank: gte-rerank reranking.
	//
	//
	//
	//
	//
	// <props="china">
	//
	// Default value: empty, which uses qwen3-rerank.
	//
	// > If you only need semantic reranking, use `qwen3-rerank`. If you need both semantic reranking and text matching features to ensure relevance, use `qwen3-rerank-hybrid`.
	//
	// >
	//
	//
	//
	//
	// <props="intl">
	//
	// Default value: empty, which uses gte-rerank-hybrid.
	//
	// > If you only need semantic reranking, use `gte-rerank`. If you need both semantic reranking and text matching features to ensure relevance, use `gte-rerank-hybrid`.
	//
	// >
	//
	//
	//
	//
	//
	// <props="china">
	//
	// > `gte-rerank-hybrid` and `gte-rerank` will no longer be updated and are not recommended.
	//
	// >
	//
	// example:
	//
	// gte-rerank-hybrid
	RerankModelName *string `json:"RerankModelName,omitempty" xml:"RerankModelName,omitempty"`
	// <props="china">
	//
	// The sentence separator, which takes effect only when `chunkMode`=**regex*	- (it does not take effect in other modes even if specified). You can pass a single regular expression (multiple expressions are not supported) to split files into small text chunks.
	//
	// When using intelligent chunking (when `chunkMode` is not specified), keep the default empty value.
	//
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not available. Do not pass this parameter.
	//
	// example:
	//
	// (?<=。)
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// The AnalyticDB for PostgreSQL instance ID (required only when `SinkType` is set to ADB). Obtain this ID from the [AnalyticDB for PostgreSQL instance list](https://gpdbnext.console.aliyun.com/gpdb/list) page.
	//
	// example:
	//
	// gp-bp32109xxxx
	SinkInstanceId *string `json:"SinkInstanceId,omitempty" xml:"SinkInstanceId,omitempty"`
	// The region of the AnalyticDB for PostgreSQL instance (required only when `SinkType` is set to ADB). Call <props="china">[DescribeRegions](https://www.alibabacloud.com/help/en/analyticdb-for-postgresql/developer-reference/api-gpdb-2016-05-03-describeregions)<props="intl">[DescribeRegions](https://www.alibabacloud.com/help/zh/analyticdb/analyticdb-for-postgresql/developer-reference/api-gpdb-2016-05-03-describeregions?spm=a2c63.p38356.0.i3) to obtain the list of regions.
	//
	// example:
	//
	// cn-hangzhou
	SinkRegion *string `json:"SinkRegion,omitempty" xml:"SinkRegion,omitempty"`
	// The vector storage type of the knowledge base. For more information, see [Knowledge base](https://help.aliyun.com/document_detail/2807740.html). Valid values:
	//
	// - BUILT_IN: Vector data is hosted on the Alibaba Cloud Model Studio platform.
	//
	// - ADB: AnalyticDB for PostgreSQL database. If you need advanced features such as database management, auditing, and monitoring, select ADB.
	//
	// > If you have not used ADB storage on Alibaba Cloud Model Studio before, go to the <props="china">[Create Knowledge Base](https://bailian.console.aliyun.com/#/knowledge-base/create)<props="intl">[Create Knowledge Base](https://bailian.console.alibabacloud.com/#/knowledge-base/create) page, select ADB-PG as the vector storage type, and complete authorization as prompted. If you pass ADB, you must specify the `SinkInstanceId` and `SinkRegion` parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// BUILT_IN
	SinkType *string `json:"SinkType,omitempty" xml:"SinkType,omitempty"`
	// 	Notice: This parameter is required in the latest SDK. Otherwise, calling the SubmitIndexJob operation returns an error: Required parameter(data_sources) missing or invalid.
	//
	// The data source type. Valid values:
	//
	// - DATA_CENTER_CATEGORY: Category type. Imports all files under specified categories in <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center). Multiple categories can be imported simultaneously.
	//
	// - DATA_CENTER_FILE: File type. Imports specified files from <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center). Multiple files can be imported simultaneously.
	//
	// > If this parameter is set to DATA_CENTER_CATEGORY, you must specify the `CategoryIds` parameter. If this parameter is set to DATA_CENTER_FILE, you must specify the `DocumentIds` parameter.
	//
	// >
	//
	// > To create an empty knowledge base, use an empty category that contains no files: set this parameter to DATA_CENTER_CATEGORY and pass the empty category ID in `CategoryIds`.
	//
	// >
	//
	// if can be null:
	// false
	//
	// example:
	//
	// DATA_CENTER_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The knowledge base type.
	//
	// **Valid values:**
	//
	// - unstructured: A document search or audio/video knowledge base. The default scenario for document search type is basic document Q&A. <props="china">To create other scenarios, pass the knowledgeType and knowledgeScene parameters.
	//
	// <props="china">
	//
	// - structured: A data query or image-based Q&A knowledge base.
	//
	//
	//
	// > The knowledge base type cannot be changed after creation.
	//
	// >
	//
	// This parameter is required.
	//
	// example:
	//
	// unstructured
	StructureType *string `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
	// <props="china">
	//
	// Obtained by clicking the ID icon next to the table name on the Tables tab of [Data Connections](https://bailian.console.aliyun.com/cn-beijing?tab=app#/connector/list) table connector. If the list contains multiple IDs, only the first one is used.
	//
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not available. Do not pass this parameter.
	//
	// >
	TableIdsShrink *string `json:"TableIds,omitempty" xml:"TableIds,omitempty"`
	// example:
	//
	// connector
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// <props="china">
	//
	// Enables custom chunking and specifies the chunking strategy. For more information, see [Knowledge base](https://help.aliyun.com/document_detail/2807740.html).
	//
	// Valid values (only one value can be passed at a time):
	//
	// - **length**: Chunk by length. Strictly chunks according to the specified `ChunkSize` and `OverlapSize`. If these two parameters are not passed, the system uses default values (`ChunkSize` of 500 and `OverlapSize` of 100). Chunking by length does not support `Separator` (it does not take effect even if specified).
	//
	// - **page**: Chunk by page. If `ChunkSize` is specified, it is also considered during chunking (if not passed, the default value of 500 is used). Chunking by page does not support `OverlapSize` or `Separator` (they do not take effect even if specified).
	//
	// - **h1**: Chunk by first-level headings. If `ChunkSize` is specified, it is also considered during chunking (if not passed, the default value of 500 is used). Chunking by first-level headings does not support `OverlapSize` or `Separator` (they do not take effect even if specified).
	//
	// - **h2**: Chunk by second-level headings. If `ChunkSize` is specified, it is also considered during chunking (if not passed, the default value of 500 is used). Chunking by second-level headings does not support `OverlapSize` or `Separator` (they do not take effect even if specified).
	//
	// - **regex**: Chunk by regular expression. The `Separator` parameter must be specified. If `ChunkSize` is specified, it is also considered during chunking (if not passed, the default value of 500 is used). Chunking by regular expression does not support `OverlapSize` (it does not take effect even if specified).
	//
	// If not specified, intelligent chunking is used by default.
	//
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not available. Do not pass this parameter.
	//
	// example:
	//
	// regex
	ChunkMode *string `json:"chunkMode,omitempty" xml:"chunkMode,omitempty"`
	// example:
	//
	// conn_mysql_xxx_xxx
	ConnectId      *string `json:"connectId,omitempty" xml:"connectId,omitempty"`
	Database       *string `json:"database,omitempty" xml:"database,omitempty"`
	DatasourceCode *string `json:"datasourceCode,omitempty" xml:"datasourceCode,omitempty"`
	// Specifies whether to treat the first row of all xlsx and xls files as headers and concatenate them into each text chunk, preventing the large language model from treating headers as regular data rows.
	//
	//
	// > Enable this feature only when all imported files are in .xlsx or .xls format and contain headers. Otherwise, do not enable it.
	//
	// >
	//
	// Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// If not specified, this feature is disabled by default.
	//
	// example:
	//
	// false
	EnableHeaders  *bool   `json:"enableHeaders,omitempty" xml:"enableHeaders,omitempty"`
	KnowledgeScene *string `json:"knowledgeScene,omitempty" xml:"knowledgeScene,omitempty"`
	// <props="china">
	//
	// The specific knowledge type, which further specifies the type of data processed by the knowledge base.
	//
	// <notice>This parameter and knowledgeScene must be provided together or omitted together. They cannot be set independently. If both are omitted, the system uses default configurations based on structureType.
	//
	// **Settings constraint**: The value of this parameter must match the selected structureType and determines the active values for knowledgeScene.
	//
	// **Valid values**:
	//
	// - document: Document search. Must be used with structureType: unstructured.
	//
	// - table: Data query. Must be used with structureType: structured.
	//
	// - image: Image-based Q&A. Must be used with structureType: structured.
	//
	// - multimedia: Audio/video search. Must be used with structureType: unstructured.
	//
	//
	//
	//
	// <props="intl">This parameter is not available. Do not pass this parameter.
	//
	// example:
	//
	// document
	KnowledgeType *string `json:"knowledgeType,omitempty" xml:"knowledgeType,omitempty"`
	// The metadata extraction configuration. Metadata is a set of additional attributes related to unstructured data content. These attributes are integrated into text chunks as key-value pairs. For more information, see [Knowledge base](https://help.aliyun.com/document_detail/2807740.html).
	MetaExtractColumnsShrink *string `json:"metaExtractColumns,omitempty" xml:"metaExtractColumns,omitempty"`
	// <props="china">The number of RCUs for the knowledge base (required only when pipelineCommercialType is set to enterprise). Value range: [1-200].
	//
	//
	// <props="intl">
	//
	// > This parameter is not available. Do not pass this parameter.
	//
	// >
	//
	// example:
	//
	// 1
	PipelineCommercialCu *int32 `json:"pipelineCommercialCu,omitempty" xml:"pipelineCommercialCu,omitempty"`
	// <props="china">
	//
	// The [specification type](https://help.aliyun.com/document_detail/2997110.html) of the knowledge base. Valid values:
	//
	// - standard: Standard Edition.
	//
	// - enterprise: Ultimate Edition.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not available. Do not pass this parameter.
	//
	// >
	//
	// example:
	//
	// standard
	PipelineCommercialType *string `json:"pipelineCommercialType,omitempty" xml:"pipelineCommercialType,omitempty"`
	// <props="china">The rate limiting strategy for knowledge base dependent links (required only when pipelineCommercialType is set to enterprise).
	//
	// Valid values:
	//
	// downgrade: Downgrade processing (switch to lightweight link retrieval).
	//
	// If not specified, the default value is downgrade.
	//
	//
	// <props="intl">
	//
	// > This parameter is not available. Do not pass this parameter.
	//
	// >
	//
	// example:
	//
	// downgrade
	PipelineRetrieveRateLimitStrategy *string `json:"pipelineRetrieveRateLimitStrategy,omitempty" xml:"pipelineRetrieveRateLimitStrategy,omitempty"`
	Table                             *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s CreateIndexShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateIndexShrinkRequest) GetCategoryIdsShrink() *string {
	return s.CategoryIdsShrink
}

func (s *CreateIndexShrinkRequest) GetChunkSize() *int32 {
	return s.ChunkSize
}

func (s *CreateIndexShrinkRequest) GetColumnsShrink() *string {
	return s.ColumnsShrink
}

func (s *CreateIndexShrinkRequest) GetCreateIndexType() *string {
	return s.CreateIndexType
}

func (s *CreateIndexShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateIndexShrinkRequest) GetDocumentIdsShrink() *string {
	return s.DocumentIdsShrink
}

func (s *CreateIndexShrinkRequest) GetEmbeddingModelName() *string {
	return s.EmbeddingModelName
}

func (s *CreateIndexShrinkRequest) GetEnableRewrite() *bool {
	return s.EnableRewrite
}

func (s *CreateIndexShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateIndexShrinkRequest) GetOverlapSize() *int32 {
	return s.OverlapSize
}

func (s *CreateIndexShrinkRequest) GetRerankInstruct() *string {
	return s.RerankInstruct
}

func (s *CreateIndexShrinkRequest) GetRerankMinScore() *float64 {
	return s.RerankMinScore
}

func (s *CreateIndexShrinkRequest) GetRerankMode() *string {
	return s.RerankMode
}

func (s *CreateIndexShrinkRequest) GetRerankModelName() *string {
	return s.RerankModelName
}

func (s *CreateIndexShrinkRequest) GetSeparator() *string {
	return s.Separator
}

func (s *CreateIndexShrinkRequest) GetSinkInstanceId() *string {
	return s.SinkInstanceId
}

func (s *CreateIndexShrinkRequest) GetSinkRegion() *string {
	return s.SinkRegion
}

func (s *CreateIndexShrinkRequest) GetSinkType() *string {
	return s.SinkType
}

func (s *CreateIndexShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateIndexShrinkRequest) GetStructureType() *string {
	return s.StructureType
}

func (s *CreateIndexShrinkRequest) GetTableIdsShrink() *string {
	return s.TableIdsShrink
}

func (s *CreateIndexShrinkRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *CreateIndexShrinkRequest) GetChunkMode() *string {
	return s.ChunkMode
}

func (s *CreateIndexShrinkRequest) GetConnectId() *string {
	return s.ConnectId
}

func (s *CreateIndexShrinkRequest) GetDatabase() *string {
	return s.Database
}

func (s *CreateIndexShrinkRequest) GetDatasourceCode() *string {
	return s.DatasourceCode
}

func (s *CreateIndexShrinkRequest) GetEnableHeaders() *bool {
	return s.EnableHeaders
}

func (s *CreateIndexShrinkRequest) GetKnowledgeScene() *string {
	return s.KnowledgeScene
}

func (s *CreateIndexShrinkRequest) GetKnowledgeType() *string {
	return s.KnowledgeType
}

func (s *CreateIndexShrinkRequest) GetMetaExtractColumnsShrink() *string {
	return s.MetaExtractColumnsShrink
}

func (s *CreateIndexShrinkRequest) GetPipelineCommercialCu() *int32 {
	return s.PipelineCommercialCu
}

func (s *CreateIndexShrinkRequest) GetPipelineCommercialType() *string {
	return s.PipelineCommercialType
}

func (s *CreateIndexShrinkRequest) GetPipelineRetrieveRateLimitStrategy() *string {
	return s.PipelineRetrieveRateLimitStrategy
}

func (s *CreateIndexShrinkRequest) GetTable() *string {
	return s.Table
}

func (s *CreateIndexShrinkRequest) SetCategoryIdsShrink(v string) *CreateIndexShrinkRequest {
	s.CategoryIdsShrink = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetChunkSize(v int32) *CreateIndexShrinkRequest {
	s.ChunkSize = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetColumnsShrink(v string) *CreateIndexShrinkRequest {
	s.ColumnsShrink = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetCreateIndexType(v string) *CreateIndexShrinkRequest {
	s.CreateIndexType = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetDescription(v string) *CreateIndexShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetDocumentIdsShrink(v string) *CreateIndexShrinkRequest {
	s.DocumentIdsShrink = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetEmbeddingModelName(v string) *CreateIndexShrinkRequest {
	s.EmbeddingModelName = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetEnableRewrite(v bool) *CreateIndexShrinkRequest {
	s.EnableRewrite = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetName(v string) *CreateIndexShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetOverlapSize(v int32) *CreateIndexShrinkRequest {
	s.OverlapSize = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetRerankInstruct(v string) *CreateIndexShrinkRequest {
	s.RerankInstruct = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetRerankMinScore(v float64) *CreateIndexShrinkRequest {
	s.RerankMinScore = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetRerankMode(v string) *CreateIndexShrinkRequest {
	s.RerankMode = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetRerankModelName(v string) *CreateIndexShrinkRequest {
	s.RerankModelName = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetSeparator(v string) *CreateIndexShrinkRequest {
	s.Separator = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetSinkInstanceId(v string) *CreateIndexShrinkRequest {
	s.SinkInstanceId = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetSinkRegion(v string) *CreateIndexShrinkRequest {
	s.SinkRegion = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetSinkType(v string) *CreateIndexShrinkRequest {
	s.SinkType = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetSourceType(v string) *CreateIndexShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetStructureType(v string) *CreateIndexShrinkRequest {
	s.StructureType = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetTableIdsShrink(v string) *CreateIndexShrinkRequest {
	s.TableIdsShrink = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetChannelType(v string) *CreateIndexShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetChunkMode(v string) *CreateIndexShrinkRequest {
	s.ChunkMode = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetConnectId(v string) *CreateIndexShrinkRequest {
	s.ConnectId = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetDatabase(v string) *CreateIndexShrinkRequest {
	s.Database = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetDatasourceCode(v string) *CreateIndexShrinkRequest {
	s.DatasourceCode = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetEnableHeaders(v bool) *CreateIndexShrinkRequest {
	s.EnableHeaders = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetKnowledgeScene(v string) *CreateIndexShrinkRequest {
	s.KnowledgeScene = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetKnowledgeType(v string) *CreateIndexShrinkRequest {
	s.KnowledgeType = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetMetaExtractColumnsShrink(v string) *CreateIndexShrinkRequest {
	s.MetaExtractColumnsShrink = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetPipelineCommercialCu(v int32) *CreateIndexShrinkRequest {
	s.PipelineCommercialCu = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetPipelineCommercialType(v string) *CreateIndexShrinkRequest {
	s.PipelineCommercialType = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetPipelineRetrieveRateLimitStrategy(v string) *CreateIndexShrinkRequest {
	s.PipelineRetrieveRateLimitStrategy = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetTable(v string) *CreateIndexShrinkRequest {
	s.Table = &v
	return s
}

func (s *CreateIndexShrinkRequest) Validate() error {
	return dara.Validate(s)
}
