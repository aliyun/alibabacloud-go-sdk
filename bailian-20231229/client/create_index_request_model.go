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
	// You can import files when you create a knowledge base. Specify category IDs to import all files under the corresponding categories. We recommend importing no more than 10,000 files. If you have more files, you can call the **SubmitIndexAddDocumentsJob*	- operation to import them later.
	CategoryIds []*string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
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
	Columns []*CreateIndexRequestColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
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
	DocumentIds []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
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
	TableIds []*string `json:"TableIds,omitempty" xml:"TableIds,omitempty" type:"Repeated"`
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
	MetaExtractColumns []*CreateIndexRequestMetaExtractColumns `json:"metaExtractColumns,omitempty" xml:"metaExtractColumns,omitempty" type:"Repeated"`
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
	// > This parameter is not yet available. Do not specify it.
	//
	// example:
	//
	// school
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// <props="china">
	//
	// Specifies whether to participate in model response generation. If enabled, the retrieval results from this column are used as input for the LLM to generate an answer. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not yet available. Do not specify it.
	//
	// example:
	//
	// true
	IsRecall *bool `json:"IsRecall,omitempty" xml:"IsRecall,omitempty"`
	// <props="china">
	//
	// Specifies whether to participate in knowledge base retrieval. If enabled, the knowledge base is allowed to search for data in this column. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not yet available. Do not specify it.
	//
	// example:
	//
	// true
	IsSearch *bool `json:"IsSearch,omitempty" xml:"IsSearch,omitempty"`
	// <props="china">
	//
	// The field name. It must be consistent with the table header of the data table created in Application Data.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not yet available. Do not specify it.
	//
	// example:
	//
	// 学校
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// <props="china">
	//
	// The field type. It must be consistent with the table header of the data table created in Application Data. Valid values:
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
	// > This parameter is not yet available. Do not specify it.
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
	// The Chinese description of the metadata field. The description can be 0 to 1,000 characters in length and can contain Chinese characters, letters, digits, underscores (_), hyphens (-), periods (.), and colons (:). The default value is empty.
	//
	// example:
	//
	// 作者名
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// If enabled, the metadata field and its value are used along with the text chunk content in the answer generation process of the LLM. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// The default value is false.
	//
	// example:
	//
	// false
	EnableLlm *bool `json:"EnableLlm,omitempty" xml:"EnableLlm,omitempty"`
	// If enabled, the metadata field and its value are used along with the text chunk content in the knowledge base retrieval process. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// The default value is false.
	//
	// example:
	//
	// false
	EnableSearch *bool `json:"EnableSearch,omitempty" xml:"EnableSearch,omitempty"`
	// The metadata field. The field name can be 1 to 50 characters in length and must consist of letters or underscores. If you specify this parameter, you must also specify the \\`Value\\` and \\`Type\\` parameters.
	//
	// example:
	//
	// author
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The method for obtaining the value of the metadata field. Valid values:
	//
	// - constant: Constant.
	//
	// - variable: Variable.
	//
	// - custom_prompt: Large Language Model (LLM).
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
