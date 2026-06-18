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
	// You can import files when you create a knowledge base. Specify category IDs to import all files under the corresponding categories. We recommend importing no more than 10,000 files. If you have more files, you can call the **SubmitIndexAddDocumentsJob*	- operation to import them later.
	CategoryIdsShrink *string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty"`
	// <props="china">
	//
	// The chunk size, which is the maximum number of characters for each text chunk. If this length is exceeded:
	//
	// - **Smart chunking*	- (the \\`chunkMode\\` parameter is not specified): The text is likely to be truncated.
	//
	// - **Custom chunking*	- (the \\`chunkMode\\` parameter is specified): The text is forcibly truncated.
	//
	//
	//
	// <props="intl">
	//
	// The chunk size, which is the maximum number of characters for each text chunk. If this length is exceeded, the text is likely to be truncated.
	//
	//
	//
	// The value must be between 1 and 6000. If you do not specify this parameter, the default value 500 is used.
	//
	// > If you set \\`ChunkSize\\` to a value less than 100, you must also set \\`OverlapSize\\`. You can also leave both parameters unspecified, and the system will use the default values.
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
	//
	// <props="intl">
	//
	// > This parameter is not yet available. Do not specify it.
	ColumnsShrink *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	// > This parameter is not yet available. Do not specify it.
	//
	// example:
	//
	// standard
	CreateIndexType *string `json:"CreateIndexType,omitempty" xml:"CreateIndexType,omitempty"`
	// The description of the knowledge base. The description can be 0 to 1,000 English or Chinese characters in length.
	//
	// The default value is empty.
	//
	// example:
	//
	// 企业帮助文档库包括了公司制度、产品清单等重要资料。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// You can import files when you create a knowledge base. Specify a list of files to import by providing their IDs. We recommend importing no more than 10,000 files. If you have more files, you can call the **SubmitIndexAddDocumentsJob*	- operation to import them later.
	DocumentIdsShrink *string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty"`
	// <props="china">
	//
	// The vector model used by the knowledge base. A vector model converts the original input prompt and knowledge text into numerical vectors to compare their similarity. The text-embedding-v4 model is a comprehensive upgrade over the text-embedding-v3 model in terms of language support, vectorization of code snippets, and vector dimension selection. It is suitable for most scenarios. For more information, see [Vectorization](https://help.aliyun.com/document_detail/2842587.html). Valid values:
	//
	// - text-embedding-v4
	//
	// - text-embedding-v3
	//
	// If you do not specify this parameter, \\`text-embedding-v3\\` is used.
	//
	//
	//
	// <props="intl">
	//
	// - The vector model used by the knowledge base. A vector model converts the original input prompt and knowledge text into numerical vectors to compare their similarity. The default text-embedding-v2 model (which cannot be changed for now) supports both Chinese and English, along with multiple other languages, and normalizes the vector results. For more information, see [Vectorization](https://help.aliyun.com/document_detail/2842587.html). Valid values:
	//
	//
	//
	//
	// - text-embedding-v2
	//
	// If you do not specify this parameter, \\`text-embedding-v2\\` is used.
	EmbeddingModelName *string `json:"EmbeddingModelName,omitempty" xml:"EmbeddingModelName,omitempty"`
	// Specifies whether to enable multi-turn conversation rewriting. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// If you do not specify this parameter, this feature is enabled by default.
	//
	// example:
	//
	// true
	EnableRewrite *bool `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	// The name of the knowledge base. The name can be 1 to 20 characters in length and can contain Chinese characters, letters, digits, underscores (_), hyphens (-), periods (.), and colons (:).
	//
	// This parameter is required.
	//
	// example:
	//
	// 企业帮助文档库
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The overlap size, which is the number of overlapping characters between the current text chunk and the previous one. The value must be between 0 and 1024.
	//
	// If you do not specify this parameter, the default value 100 is used.
	//
	// > \\`OverlapSize\\` must be smaller than \\`ChunkSize\\`. Otherwise, chunking errors will occur.
	//
	// example:
	//
	// 16
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// The name of the database. This parameter is required when creating a data query knowledge base.
	//
	// The database must exist in the data source specified by \\`datasourceCode\\`.
	RerankInstruct *string `json:"RerankInstruct,omitempty" xml:"RerankInstruct,omitempty"`
	// The similarity threshold. Only text chunks with a similarity score greater than this value are recalled. This is used to filter the text chunks returned by the reranking model. The value must be between 0.01 and 1.00.
	//
	// If you do not specify this parameter, the default value 0.01 is used.
	//
	// example:
	//
	// 0.20
	RerankMinScore *float64 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// The name of the data table. This parameter is required when creating a data query knowledge base.
	//
	// The data table must exist in the data source specified by \\`connectId\\` or \\`datasourceCode\\`.
	//
	// example:
	//
	// qa
	RerankMode *string `json:"RerankMode,omitempty" xml:"RerankMode,omitempty"`
	// The reranking model used by the knowledge base. The reranking model is an external scoring system that calculates a similarity score between the user\\"s question and each text chunk in the knowledge base, sorts them in descending order, and returns the top K text chunks. Valid values:
	//
	// <props="china">
	//
	// - qwen3-rerank-hybrid: qwen3-rerank (hybrid) reranking.
	//
	// - qwen3-rerank: qwen3-rerank reranking.
	//
	// - gte-rerank-hybrid: gte-rerank (hybrid) reranking.
	//
	// - gte-rerank: gte-rerank reranking.
	//
	//
	//
	// <props="intl">
	//
	// - gte-rerank-hybrid: Official reranking.
	//
	// - gte-rerank: gte-rerank reranking.
	//
	//
	//
	// <props="china">
	//
	// If you do not specify this parameter, \\`qwen3-rerank\\` is used.
	//
	// > Use \\`qwen3-rerank\\` if you only need semantic sorting. Use \\`qwen3-rerank-hybrid\\` if you need both semantic sorting and text-matching features to ensure relevance.
	//
	//
	//
	// <props="intl">
	//
	// If you do not specify this parameter, \\`gte-rerank-hybrid\\` is used.
	//
	// > Use \\`gte-rerank\\` if you only need semantic sorting. Use \\`gte-rerank-hybrid\\` if you need both semantic sorting and text-matching features to ensure relevance.
	//
	//
	//
	// <props="china">
	//
	// > The \\`gte-rerank-hybrid\\` and \\`gte-rerank\\` models are no longer updated and are not recommended.
	//
	// example:
	//
	// gte-rerank-hybrid
	RerankModelName *string `json:"RerankModelName,omitempty" xml:"RerankModelName,omitempty"`
	// <props="china">
	//
	// The sentence separator. This parameter takes effect only when \\`chunkMode\\` is set to **regex**. It is ignored in other modes, even if specified. You can enter a regular expression (multiple expressions are not supported) to split the file into smaller text chunks.
	//
	// For smart chunking (the \\`chunkMode\\` parameter is not specified), you can leave this parameter empty.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not yet available. Do not specify it.
	//
	// example:
	//
	// (?<=。)
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// The ID of the AnalyticDB for PostgreSQL instance. This parameter is required only when \\`SinkType\\` is set to ADB. Go to the [AnalyticDB for PostgreSQL instance list](https://gpdbnext.console.aliyun.com/gpdb/list) page to obtain this ID.
	//
	// example:
	//
	// gp-bp32109xxxx
	SinkInstanceId *string `json:"SinkInstanceId,omitempty" xml:"SinkInstanceId,omitempty"`
	// The region where the AnalyticDB for PostgreSQL instance is located. This parameter is required only when \\`SinkType\\` is set to ADB. You can call the <props="intl">[DescribeRegions ](https://www.alibabacloud.com/help/zh/analyticdb/analyticdb-for-postgresql/developer-reference/api-gpdb-2016-05-03-describeregions?spm=a2c63.p38356.0.i3)operation to obtain a list of regions.
	//
	// example:
	//
	// cn-hangzhou
	SinkRegion *string `json:"SinkRegion,omitempty" xml:"SinkRegion,omitempty"`
	// The storage class for the knowledge base vectors. For more information, see [Knowledge bases](https://help.aliyun.com/document_detail/2807740.html). Valid values:
	//
	// - BUILT_IN: Hosts the vector data on the Alibaba Cloud Model Studio platform.
	//
	// - ADB: AnalyticDB for PostgreSQL. We recommend choosing ADB if you need advanced features such as database management, auditing, and monitoring.
	//
	// > If you have not used ADB storage on Alibaba Cloud Model Studio, go to the <props="intl">[Create Knowledge Base](https://bailian.console.alibabacloud.com/#/knowledge-base/create) page, set the vector storage class to ADB-PG, and follow the on-screen instructions to grant the required permissions. If you set this parameter to ADB, you must specify the \\`SinkInstanceId\\` and \\`SinkRegion\\` parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// BUILT_IN
	SinkType *string `json:"SinkType,omitempty" xml:"SinkType,omitempty"`
	// 	Notice:
	//
	// In the latest SDK version, this parameter is required. Otherwise, calling the SubmitIndexJob operation will result in the error: Required parameter(data_sources) missing or invalid.
	//
	//
	//
	// The source of the imported data. Valid values:
	//
	// - DATA_CENTER_CATEGORY: Category type. Imports all files under the specified categories in <props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center). You can import multiple categories at the same time.
	//
	// - DATA_CENTER_FILE: File type. Imports the specified files from <props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center). You can import multiple files at the same time.
	//
	// > If you set this parameter to DATA_CENTER_CATEGORY, you must specify the \\`CategoryIds\\` parameter. If you set this parameter to DATA_CENTER_FILE, you must specify the \\`DocumentIds\\` parameter.
	//
	// > To create an empty knowledge base, use an empty category that contains no files. Set this parameter to DATA_CENTER_CATEGORY and specify the ID of the empty category for \\`CategoryIds\\`.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// DATA_CENTER_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The type of the knowledge base.
	//
	// **Valid values**:
	//
	// - unstructured: A knowledge base for document search, audio, or video. The default scenario for document search is basic document Q\\&A.
	//
	// <props="china">
	//
	// - structured: A knowledge base for data query or image Q\\&A.
	//
	//
	//
	//
	// > The type of a knowledge base cannot be changed after it is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// unstructured
	StructureType *string `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
	// <props="china">
	//
	// Obtain the table ID on the Tables tab of the table connector in Data Connections by clicking the ID icon next to the table name. If the list contains multiple IDs, only the first one is used.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not yet available. Do not specify it.
	TableIdsShrink *string `json:"TableIds,omitempty" xml:"TableIds,omitempty"`
	// example:
	//
	// connector
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// <props="china">
	//
	// Enables custom chunking and specifies the chunking policy. For more information, see [Knowledge bases](https://help.aliyun.com/document_detail/2807740.html).
	//
	// Possible values (only one value can be specified at a time):
	//
	// - **length**: Chunks by length. The text is strictly chunked according to the \\`ChunkSize\\` and \\`OverlapSize\\` you specify. If you do not specify these two parameters, the system uses the default values (\\`ChunkSize\\` is 500, \\`OverlapSize\\` is 100). Chunking by length does not support \\`Separator\\` (it is ignored even if specified).
	//
	// - **page**: Chunks by page. If \\`ChunkSize\\` is specified, it is also considered during chunking (if not specified, the default value 500 is used). Chunking by page does not support \\`OverlapSize\\` or \\`Separator\\` (they are ignored even if specified).
	//
	// - **h1**: Chunks by level-1 heading. If \\`ChunkSize\\` is specified, it is also considered during chunking (if not specified, the default value 500 is used). Chunking by level-1 heading does not support \\`OverlapSize\\` or \\`Separator\\` (they are ignored even if specified).
	//
	// - **h2**: Chunks by level-2 heading. If \\`ChunkSize\\` is specified, it is also considered during chunking (if not specified, the default value 500 is used). Chunking by level-2 heading does not support \\`OverlapSize\\` or \\`Separator\\` (they are ignored even if specified).
	//
	// - **regex**: Chunks by regular expression. You must specify the \\`Separator\\` parameter. If \\`ChunkSize\\` is specified, it is also considered during chunking (if not specified, the default value 500 is used). Chunking by regular expression does not support \\`OverlapSize\\` (it is ignored even if specified).
	//
	// If you do not specify this parameter, smart chunking is used by default.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not yet available. Do not specify it.
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
	// Specifies whether to treat the first row of all .xlsx and .xls files as the table header and append it to each text chunk. This prevents the LLM from treating the header as a regular data row.
	//
	// > We recommend enabling this feature only when all imported files are in .xlsx or .xls format and contain a header. Otherwise, do not enable it.
	//
	// Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// If you do not specify this parameter, this feature is disabled by default.
	//
	// example:
	//
	// false
	EnableHeaders  *bool   `json:"enableHeaders,omitempty" xml:"enableHeaders,omitempty"`
	KnowledgeScene *string `json:"knowledgeScene,omitempty" xml:"knowledgeScene,omitempty"`
	// The data source code. This parameter is required when creating a data query knowledge base and is used with \\`table\\` and \\`database\\`.
	//
	// <props="china">
	//
	// We recommend using the new \\`connectId\\` parameter, which you can obtain from the data connector card on the [Data Connections](https://modelstudio.console.alibabacloud.com/?tab=app#/connector/list) page. This parameter is still compatible but will no longer be maintained in the future.
	//
	//
	//
	// > - This operation does not support associating custom databases. Use the Alibaba Cloud Model Studio console to create them.
	//
	// example:
	//
	// document
	KnowledgeType *string `json:"knowledgeType,omitempty" xml:"knowledgeType,omitempty"`
	// The metadata extraction configuration. Metadata is a series of additional attributes related to unstructured data content. These attributes are integrated into text chunks as key-value pairs. For more information, see [Knowledge bases](https://help.aliyun.com/document_detail/2807740.html).
	MetaExtractColumnsShrink *string `json:"metaExtractColumns,omitempty" xml:"metaExtractColumns,omitempty"`
	// <props="china">
	//
	// The number of RCUs for the knowledge base. This parameter is required only when \\`pipelineCommercialType\\` is set to \\`enterprise\\`. The value must be between 1 and 200.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not yet available. Do not specify it.
	//
	// example:
	//
	// 1
	PipelineCommercialCu *int32 `json:"pipelineCommercialCu,omitempty" xml:"pipelineCommercialCu,omitempty"`
	// <props="china">
	//
	// The [edition type](https://help.aliyun.com/document_detail/2997110.html) of the knowledge base. Valid values:
	//
	// - standard: Standard Edition
	//
	// - enterprise: Ultimate Edition
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not yet available. Do not specify it.
	//
	// example:
	//
	// standard
	PipelineCommercialType *string `json:"pipelineCommercialType,omitempty" xml:"pipelineCommercialType,omitempty"`
	// <props="china">
	//
	// The rate limiting policy for the knowledge base dependency chain. This parameter is required only when \\`pipelineCommercialType\\` is set to \\`enterprise\\`.
	//
	// Value:
	//
	// downgrade: Degrades the service (switches to using a lightweight retrieval chain).
	//
	// If you do not specify this parameter, the default value \\`downgrade\\` is used.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not yet available. Do not specify it.
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
