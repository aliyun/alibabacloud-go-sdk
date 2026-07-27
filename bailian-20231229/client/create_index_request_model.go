// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryIds(v []*string) *CreateIndexRequest
	GetCategoryIds() []*string
	SetChunkSize(v int32) *CreateIndexRequest
	GetChunkSize() *int32
	SetColumns(v []*CreateIndexRequestColumns) *CreateIndexRequest
	GetColumns() []*CreateIndexRequestColumns
	SetCreateIndexType(v string) *CreateIndexRequest
	GetCreateIndexType() *string
	SetDescription(v string) *CreateIndexRequest
	GetDescription() *string
	SetDocumentIds(v []*string) *CreateIndexRequest
	GetDocumentIds() []*string
	SetEmbeddingModelName(v string) *CreateIndexRequest
	GetEmbeddingModelName() *string
	SetEnableRewrite(v bool) *CreateIndexRequest
	GetEnableRewrite() *bool
	SetName(v string) *CreateIndexRequest
	GetName() *string
	SetOverlapSize(v int32) *CreateIndexRequest
	GetOverlapSize() *int32
	SetRerankInstruct(v string) *CreateIndexRequest
	GetRerankInstruct() *string
	SetRerankMinScore(v float64) *CreateIndexRequest
	GetRerankMinScore() *float64
	SetRerankMode(v string) *CreateIndexRequest
	GetRerankMode() *string
	SetRerankModelName(v string) *CreateIndexRequest
	GetRerankModelName() *string
	SetSeparator(v string) *CreateIndexRequest
	GetSeparator() *string
	SetSinkInstanceId(v string) *CreateIndexRequest
	GetSinkInstanceId() *string
	SetSinkRegion(v string) *CreateIndexRequest
	GetSinkRegion() *string
	SetSinkType(v string) *CreateIndexRequest
	GetSinkType() *string
	SetSourceType(v string) *CreateIndexRequest
	GetSourceType() *string
	SetStructureType(v string) *CreateIndexRequest
	GetStructureType() *string
	SetTableIds(v []*string) *CreateIndexRequest
	GetTableIds() []*string
	SetChannelType(v string) *CreateIndexRequest
	GetChannelType() *string
	SetChunkMode(v string) *CreateIndexRequest
	GetChunkMode() *string
	SetConnectId(v string) *CreateIndexRequest
	GetConnectId() *string
	SetDatabase(v string) *CreateIndexRequest
	GetDatabase() *string
	SetDatasourceCode(v string) *CreateIndexRequest
	GetDatasourceCode() *string
	SetEnableHeaders(v bool) *CreateIndexRequest
	GetEnableHeaders() *bool
	SetKnowledgeScene(v string) *CreateIndexRequest
	GetKnowledgeScene() *string
	SetKnowledgeType(v string) *CreateIndexRequest
	GetKnowledgeType() *string
	SetMetaExtractColumns(v []*CreateIndexRequestMetaExtractColumns) *CreateIndexRequest
	GetMetaExtractColumns() []*CreateIndexRequestMetaExtractColumns
	SetPipelineCommercialCu(v int32) *CreateIndexRequest
	GetPipelineCommercialCu() *int32
	SetPipelineCommercialType(v string) *CreateIndexRequest
	GetPipelineCommercialType() *string
	SetPipelineRetrieveRateLimitStrategy(v string) *CreateIndexRequest
	GetPipelineRetrieveRateLimitStrategy() *string
	SetTable(v string) *CreateIndexRequest
	GetTable() *string
}

type CreateIndexRequest struct {
	// The list of category IDs to import when creating the knowledge base. All files under the specified categories are imported. We recommend importing no more than 500 files. For remaining files, call the **SubmitIndexAddDocumentsJob*	- operation to continue importing.
	CategoryIds []*string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
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
	Columns []*CreateIndexRequestColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
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
	DocumentIds []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
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
	TableIds []*string `json:"TableIds,omitempty" xml:"TableIds,omitempty" type:"Repeated"`
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
	MetaExtractColumns []*CreateIndexRequestMetaExtractColumns `json:"metaExtractColumns,omitempty" xml:"metaExtractColumns,omitempty" type:"Repeated"`
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

func (s CreateIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateIndexRequest) GetCategoryIds() []*string {
	return s.CategoryIds
}

func (s *CreateIndexRequest) GetChunkSize() *int32 {
	return s.ChunkSize
}

func (s *CreateIndexRequest) GetColumns() []*CreateIndexRequestColumns {
	return s.Columns
}

func (s *CreateIndexRequest) GetCreateIndexType() *string {
	return s.CreateIndexType
}

func (s *CreateIndexRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateIndexRequest) GetDocumentIds() []*string {
	return s.DocumentIds
}

func (s *CreateIndexRequest) GetEmbeddingModelName() *string {
	return s.EmbeddingModelName
}

func (s *CreateIndexRequest) GetEnableRewrite() *bool {
	return s.EnableRewrite
}

func (s *CreateIndexRequest) GetName() *string {
	return s.Name
}

func (s *CreateIndexRequest) GetOverlapSize() *int32 {
	return s.OverlapSize
}

func (s *CreateIndexRequest) GetRerankInstruct() *string {
	return s.RerankInstruct
}

func (s *CreateIndexRequest) GetRerankMinScore() *float64 {
	return s.RerankMinScore
}

func (s *CreateIndexRequest) GetRerankMode() *string {
	return s.RerankMode
}

func (s *CreateIndexRequest) GetRerankModelName() *string {
	return s.RerankModelName
}

func (s *CreateIndexRequest) GetSeparator() *string {
	return s.Separator
}

func (s *CreateIndexRequest) GetSinkInstanceId() *string {
	return s.SinkInstanceId
}

func (s *CreateIndexRequest) GetSinkRegion() *string {
	return s.SinkRegion
}

func (s *CreateIndexRequest) GetSinkType() *string {
	return s.SinkType
}

func (s *CreateIndexRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateIndexRequest) GetStructureType() *string {
	return s.StructureType
}

func (s *CreateIndexRequest) GetTableIds() []*string {
	return s.TableIds
}

func (s *CreateIndexRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *CreateIndexRequest) GetChunkMode() *string {
	return s.ChunkMode
}

func (s *CreateIndexRequest) GetConnectId() *string {
	return s.ConnectId
}

func (s *CreateIndexRequest) GetDatabase() *string {
	return s.Database
}

func (s *CreateIndexRequest) GetDatasourceCode() *string {
	return s.DatasourceCode
}

func (s *CreateIndexRequest) GetEnableHeaders() *bool {
	return s.EnableHeaders
}

func (s *CreateIndexRequest) GetKnowledgeScene() *string {
	return s.KnowledgeScene
}

func (s *CreateIndexRequest) GetKnowledgeType() *string {
	return s.KnowledgeType
}

func (s *CreateIndexRequest) GetMetaExtractColumns() []*CreateIndexRequestMetaExtractColumns {
	return s.MetaExtractColumns
}

func (s *CreateIndexRequest) GetPipelineCommercialCu() *int32 {
	return s.PipelineCommercialCu
}

func (s *CreateIndexRequest) GetPipelineCommercialType() *string {
	return s.PipelineCommercialType
}

func (s *CreateIndexRequest) GetPipelineRetrieveRateLimitStrategy() *string {
	return s.PipelineRetrieveRateLimitStrategy
}

func (s *CreateIndexRequest) GetTable() *string {
	return s.Table
}

func (s *CreateIndexRequest) SetCategoryIds(v []*string) *CreateIndexRequest {
	s.CategoryIds = v
	return s
}

func (s *CreateIndexRequest) SetChunkSize(v int32) *CreateIndexRequest {
	s.ChunkSize = &v
	return s
}

func (s *CreateIndexRequest) SetColumns(v []*CreateIndexRequestColumns) *CreateIndexRequest {
	s.Columns = v
	return s
}

func (s *CreateIndexRequest) SetCreateIndexType(v string) *CreateIndexRequest {
	s.CreateIndexType = &v
	return s
}

func (s *CreateIndexRequest) SetDescription(v string) *CreateIndexRequest {
	s.Description = &v
	return s
}

func (s *CreateIndexRequest) SetDocumentIds(v []*string) *CreateIndexRequest {
	s.DocumentIds = v
	return s
}

func (s *CreateIndexRequest) SetEmbeddingModelName(v string) *CreateIndexRequest {
	s.EmbeddingModelName = &v
	return s
}

func (s *CreateIndexRequest) SetEnableRewrite(v bool) *CreateIndexRequest {
	s.EnableRewrite = &v
	return s
}

func (s *CreateIndexRequest) SetName(v string) *CreateIndexRequest {
	s.Name = &v
	return s
}

func (s *CreateIndexRequest) SetOverlapSize(v int32) *CreateIndexRequest {
	s.OverlapSize = &v
	return s
}

func (s *CreateIndexRequest) SetRerankInstruct(v string) *CreateIndexRequest {
	s.RerankInstruct = &v
	return s
}

func (s *CreateIndexRequest) SetRerankMinScore(v float64) *CreateIndexRequest {
	s.RerankMinScore = &v
	return s
}

func (s *CreateIndexRequest) SetRerankMode(v string) *CreateIndexRequest {
	s.RerankMode = &v
	return s
}

func (s *CreateIndexRequest) SetRerankModelName(v string) *CreateIndexRequest {
	s.RerankModelName = &v
	return s
}

func (s *CreateIndexRequest) SetSeparator(v string) *CreateIndexRequest {
	s.Separator = &v
	return s
}

func (s *CreateIndexRequest) SetSinkInstanceId(v string) *CreateIndexRequest {
	s.SinkInstanceId = &v
	return s
}

func (s *CreateIndexRequest) SetSinkRegion(v string) *CreateIndexRequest {
	s.SinkRegion = &v
	return s
}

func (s *CreateIndexRequest) SetSinkType(v string) *CreateIndexRequest {
	s.SinkType = &v
	return s
}

func (s *CreateIndexRequest) SetSourceType(v string) *CreateIndexRequest {
	s.SourceType = &v
	return s
}

func (s *CreateIndexRequest) SetStructureType(v string) *CreateIndexRequest {
	s.StructureType = &v
	return s
}

func (s *CreateIndexRequest) SetTableIds(v []*string) *CreateIndexRequest {
	s.TableIds = v
	return s
}

func (s *CreateIndexRequest) SetChannelType(v string) *CreateIndexRequest {
	s.ChannelType = &v
	return s
}

func (s *CreateIndexRequest) SetChunkMode(v string) *CreateIndexRequest {
	s.ChunkMode = &v
	return s
}

func (s *CreateIndexRequest) SetConnectId(v string) *CreateIndexRequest {
	s.ConnectId = &v
	return s
}

func (s *CreateIndexRequest) SetDatabase(v string) *CreateIndexRequest {
	s.Database = &v
	return s
}

func (s *CreateIndexRequest) SetDatasourceCode(v string) *CreateIndexRequest {
	s.DatasourceCode = &v
	return s
}

func (s *CreateIndexRequest) SetEnableHeaders(v bool) *CreateIndexRequest {
	s.EnableHeaders = &v
	return s
}

func (s *CreateIndexRequest) SetKnowledgeScene(v string) *CreateIndexRequest {
	s.KnowledgeScene = &v
	return s
}

func (s *CreateIndexRequest) SetKnowledgeType(v string) *CreateIndexRequest {
	s.KnowledgeType = &v
	return s
}

func (s *CreateIndexRequest) SetMetaExtractColumns(v []*CreateIndexRequestMetaExtractColumns) *CreateIndexRequest {
	s.MetaExtractColumns = v
	return s
}

func (s *CreateIndexRequest) SetPipelineCommercialCu(v int32) *CreateIndexRequest {
	s.PipelineCommercialCu = &v
	return s
}

func (s *CreateIndexRequest) SetPipelineCommercialType(v string) *CreateIndexRequest {
	s.PipelineCommercialType = &v
	return s
}

func (s *CreateIndexRequest) SetPipelineRetrieveRateLimitStrategy(v string) *CreateIndexRequest {
	s.PipelineRetrieveRateLimitStrategy = &v
	return s
}

func (s *CreateIndexRequest) SetTable(v string) *CreateIndexRequest {
	s.Table = &v
	return s
}

func (s *CreateIndexRequest) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MetaExtractColumns != nil {
		for _, item := range s.MetaExtractColumns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateIndexRequestColumns struct {
	// > This parameter is not available. Do not pass this parameter.
	//
	// >
	//
	// example:
	//
	// school
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// <props="china">
	//
	// Specifies whether this column participates in model responses. When enabled, the search results of this column are used as input for the large language model to generate answers. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
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
	// true
	IsRecall *bool `json:"IsRecall,omitempty" xml:"IsRecall,omitempty"`
	// <props="china">
	//
	// Specifies whether this column participates in knowledge base retrieval. When enabled, the knowledge base can search within the data of this column. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
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
	// true
	IsSearch *bool `json:"IsSearch,omitempty" xml:"IsSearch,omitempty"`
	// <props="china">
	//
	// The field name. Must be consistent with the header of the data table created in Application Data.
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
	// School.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// <props="china">
	//
	// The field type. Must be consistent with the header of the data table created in Application Data. Valid values:
	//
	// - string
	//
	// - double
	//
	// - long
	//
	// - datetime
	//
	// - image_url
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
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateIndexRequestColumns) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexRequestColumns) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestColumns) GetColumn() *string {
	return s.Column
}

func (s *CreateIndexRequestColumns) GetIsRecall() *bool {
	return s.IsRecall
}

func (s *CreateIndexRequestColumns) GetIsSearch() *bool {
	return s.IsSearch
}

func (s *CreateIndexRequestColumns) GetName() *string {
	return s.Name
}

func (s *CreateIndexRequestColumns) GetType() *string {
	return s.Type
}

func (s *CreateIndexRequestColumns) SetColumn(v string) *CreateIndexRequestColumns {
	s.Column = &v
	return s
}

func (s *CreateIndexRequestColumns) SetIsRecall(v bool) *CreateIndexRequestColumns {
	s.IsRecall = &v
	return s
}

func (s *CreateIndexRequestColumns) SetIsSearch(v bool) *CreateIndexRequestColumns {
	s.IsSearch = &v
	return s
}

func (s *CreateIndexRequestColumns) SetName(v string) *CreateIndexRequestColumns {
	s.Name = &v
	return s
}

func (s *CreateIndexRequestColumns) SetType(v string) *CreateIndexRequestColumns {
	s.Type = &v
	return s
}

func (s *CreateIndexRequestColumns) Validate() error {
	return dara.Validate(s)
}

type CreateIndexRequestMetaExtractColumns struct {
	// The Chinese description of the metadata field. The description can be up to 1000 characters in length and can contain Chinese characters, letters, digits, underscores (_), hyphens (-), periods (.), and colons (:). Default value: empty.
	//
	// example:
	//
	// AuthorName.
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// Specifies whether this metadata field and its value participate in the large language model\\"s answer generation process along with the text chunk content. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableLlm *bool `json:"EnableLlm,omitempty" xml:"EnableLlm,omitempty"`
	// Specifies whether this metadata field and its value participate in knowledge base retrieval along with the text chunk content. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableSearch *bool `json:"EnableSearch,omitempty" xml:"EnableSearch,omitempty"`
	// The metadata field. The field must be 1 to 50 characters in length and can contain only letters and underscores. If this parameter is specified, you must also specify the `Value` and `Type` parameters.
	//
	// example:
	//
	// author
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The extraction method for the metadata field. Valid values:
	//
	// - constant: Constant.
	//
	// - variable: Variable.
	//
	// - custom_prompt: Large language model.
	//
	// - regular: Regular expression.
	//
	// - keywords: Keyword search.
	//
	// example:
	//
	// constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the metadata field.
	//
	// example:
	//
	// Tim
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIndexRequestMetaExtractColumns) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexRequestMetaExtractColumns) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestMetaExtractColumns) GetDesc() *string {
	return s.Desc
}

func (s *CreateIndexRequestMetaExtractColumns) GetEnableLlm() *bool {
	return s.EnableLlm
}

func (s *CreateIndexRequestMetaExtractColumns) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *CreateIndexRequestMetaExtractColumns) GetKey() *string {
	return s.Key
}

func (s *CreateIndexRequestMetaExtractColumns) GetType() *string {
	return s.Type
}

func (s *CreateIndexRequestMetaExtractColumns) GetValue() *string {
	return s.Value
}

func (s *CreateIndexRequestMetaExtractColumns) SetDesc(v string) *CreateIndexRequestMetaExtractColumns {
	s.Desc = &v
	return s
}

func (s *CreateIndexRequestMetaExtractColumns) SetEnableLlm(v bool) *CreateIndexRequestMetaExtractColumns {
	s.EnableLlm = &v
	return s
}

func (s *CreateIndexRequestMetaExtractColumns) SetEnableSearch(v bool) *CreateIndexRequestMetaExtractColumns {
	s.EnableSearch = &v
	return s
}

func (s *CreateIndexRequestMetaExtractColumns) SetKey(v string) *CreateIndexRequestMetaExtractColumns {
	s.Key = &v
	return s
}

func (s *CreateIndexRequestMetaExtractColumns) SetType(v string) *CreateIndexRequestMetaExtractColumns {
	s.Type = &v
	return s
}

func (s *CreateIndexRequestMetaExtractColumns) SetValue(v string) *CreateIndexRequestMetaExtractColumns {
	s.Value = &v
	return s
}

func (s *CreateIndexRequestMetaExtractColumns) Validate() error {
	return dara.Validate(s)
}
