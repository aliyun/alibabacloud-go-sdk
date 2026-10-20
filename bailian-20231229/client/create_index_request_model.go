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
	// You can import files synchronously when you create a knowledge base. Here, you can import all files under a specified category by specifying the category ID. We recommend that you import no more than 500 files. If you have remaining files, you can call the **SubmitIndexAddDocumentsJob*	- operation to continue importing them later.
	CategoryIds []*string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	// <props="china">
	//
	// The chunk size, which is the maximum number of characters in each text chunk. If the text exceeds this length:
	//
	// - **Intelligent chunking*	- (chunkMode is not specified): The text is highly likely to be truncated.
	//
	// - **Custom chunking*	- (chunkMode is specified): The text is forcibly truncated.
	//
	//
	//
	// <props="intl">
	//
	// The chunk size, which is the maximum number of characters in each text chunk. If the text exceeds this length, it is highly likely to be truncated.
	//
	//
	//
	// Valid values: [1, 6000]. Default value: 500.
	//
	// > If you set ChunkSize to a value less than 100, you must also set OverlapSize. You can also leave these two parameters unspecified, and the system uses the default values.
	//
	// example:
	//
	// 128
	ChunkSize *int32 `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// <props="china">
	//
	// The structure of the data table, such as column names and types.
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
	// The description of the knowledge base. The description must be 0 to 200 characters in length and can contain English and Chinese characters.
	//
	// Default value: empty.
	//
	// example:
	//
	// The enterprise help document library includes important materials such as company policies and product lists
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// You can import files synchronously when you create a knowledge base. Here, you can specify the list of files to be imported by passing file IDs. We recommend that you import no more than 10,000 files. If you have remaining files, you can call the **SubmitIndexAddDocumentsJob*	- operation to continue importing them later.
	DocumentIds []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
	// <props="china">
	//
	// The embedding model used by the knowledge base. The embedding model is used to transform the original input prompt and knowledge text into numerical vectors for similarity comparison. Compared with the text-embedding-v3 model, the text-embedding-v4 model is comprehensively upgraded in terms of language support, code snippet vectorization effect, and selection of vector dimensions, and is suitable for most scenarios. For more information, refer to [Quantization](https://help.aliyun.com/document_detail/2842587.html). Valid values:
	//
	// - text-embedding-v4
	//
	// - text-embedding-v3
	//
	// Default value: empty. In this case, the text-embedding-v3 model is used.
	//
	//
	//
	//
	// <props="intl">
	//
	// The embedding model used by the knowledge base. The embedding model is used to transform the original input prompt and knowledge text into numerical vectors for similarity comparison. The default text-embedding-v2 model (which cannot be changed at the moment) supports not only Chinese and English but also multiple other languages, and performs normalization on the vector results. For more information, refer to [Quantization](https://help.aliyun.com/document_detail/2842587.html). Valid values:
	//
	// - text-embedding-v2
	//
	// Default value: empty. In this case, the text-embedding-v2 model is used.
	//
	// example:
	//
	// text-embedding-v4
	EmbeddingModelName *string `json:"EmbeddingModelName,omitempty" xml:"EmbeddingModelName,omitempty"`
	// Specifies whether to enable multi-turn conversation rewriting. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	EnableRewrite *bool `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	// The name of the knowledge base. The name must be 1 to 20 characters in length and can contain Chinese characters, English letters, digits, underscores (_), hyphens (-), periods (.), and colons (:).
	//
	// This parameter is required.
	//
	// example:
	//
	// Enterprise help document library
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The chunk overlap size, which indicates the number of overlapping characters between the current text chunk and the previous text chunk. Valid values: [0, 1024].
	//
	// Default value: 100.
	//
	// > OverlapSize must be less than ChunkSize. Otherwise, chunking exceptions occur.
	//
	// example:
	//
	// 16
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// <props="intl">This parameter is not available. Do not specify this parameter.
	//
	// <props="china">Provides a natural language instruction to fine-tune the behavior of the reranking model.
	//
	// <notice>This parameter takes effect only when rerank_mode is set to "custom".</notice>
	//
	// example:
	//
	// Given a web search query, retrieve relevant passages that answer the query.
	RerankInstruct *string `json:"RerankInstruct,omitempty" xml:"RerankInstruct,omitempty"`
	// The similarity threshold. Only text chunks with a similarity score greater than this value are recalled. This parameter is used to filter the text chunks returned by the reranking model. Valid values: [0.01, 1.00].
	//
	// Default value: 0.01.
	//
	// example:
	//
	// 0.20
	RerankMinScore *float64 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// <props="china">
	//
	// Specifies the instruction intervention mode of the reranking model to determine its scoring preference. Valid values:
	//
	// - **qa**: question answering mode. This is the default value. The model tends to assign higher scores to candidates that can directly answer the query. We recommend that you use this mode in question answering scenarios.
	//
	// - **similar**: similarity mode. The model tends to assign higher scores to candidates that are highly consistent with the query content. We recommend that you use this mode in matching and retrieval scenarios.
	//
	// - **custom**: custom mode. The ranking behavior of the model is determined by the instruction in the rerank_instruct parameter.
	//
	//
	//
	// <props="intl">This parameter is not available. Do not specify this parameter.
	//
	// example:
	//
	// qa
	RerankMode *string `json:"RerankMode,omitempty" xml:"RerankMode,omitempty"`
	// The reranking model used by the knowledge base. The reranking model is a scoring system external to the knowledge base. It calculates the similarity score between the user\\"s question and each text chunk in the knowledge base, sorts the chunks in descending order of the scores, and returns the top K text chunks with the highest scores. Valid values:
	//
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
	// Default value: empty. In this case, qwen3-rerank is used.
	//
	// > If you only need semantic reranking, you can use qwen3-rerank. If you need both semantic reranking and text matching features to ensure relevance, use qwen3-rerank-hybrid.
	//
	// >
	//
	//
	//
	//
	// <props="intl">
	//
	// Default value: empty. In this case, gte-rerank-hybrid is used.
	//
	// > If you only need semantic reranking, you can use gte-rerank. If you need both semantic reranking and text matching features to ensure relevance, use gte-rerank-hybrid.
	//
	// >
	//
	//
	//
	//
	//
	// <props="china">
	//
	// > gte-rerank-hybrid and gte-rerank will stop receiving updates and are not recommended.
	//
	// >
	//
	// example:
	//
	// gte-rerank-hybrid
	RerankModelName *string `json:"RerankModelName,omitempty" xml:"RerankModelName,omitempty"`
	// <props="china">
	//
	// The sentence separator. This parameter takes effect only when chunkMode is set to **regex**. In other modes, this parameter does not take effect even if it is passed. You can pass a regular expression (multiple regular expressions are not supported) to split the file into small text chunks.
	//
	// When you use intelligent chunking (chunkMode is not specified), leave this parameter empty.
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
	// The ID of the AnalyticDB for PostgreSQL instance. This parameter is required only when SinkType is set to ADB. Go to the [AnalyticDB for PostgreSQL Instance List](https://gpdbnext.console.aliyun.com/gpdb/list) page to obtain the ID.
	//
	// example:
	//
	// gp-bp32109xxxx
	SinkInstanceId *string `json:"SinkInstanceId,omitempty" xml:"SinkInstanceId,omitempty"`
	// The region where the AnalyticDB for PostgreSQL instance resides. This parameter is required only when SinkType is set to ADB. You can call the <props="china">[DescribeRegions](https://www.alibabacloud.com/help/en/analyticdb-for-postgresql/developer-reference/api-gpdb-2016-05-03-describeregions)<props="intl">[DescribeRegions](https://www.alibabacloud.com/help/zh/analyticdb/analyticdb-for-postgresql/developer-reference/api-gpdb-2016-05-03-describeregions?spm=a2c63.p38356.0.i3) operation to obtain the region list.
	//
	// example:
	//
	// cn-hangzhou
	SinkRegion *string `json:"SinkRegion,omitempty" xml:"SinkRegion,omitempty"`
	// The vector storage type of the knowledge base. For more information, refer to [Knowledge base](https://help.aliyun.com/document_detail/2807740.html). Valid values:
	//
	// - BUILT_IN: stores vector data in Alibaba Cloud Model Studio.
	//
	// - ADB: AnalyticDB for PostgreSQL database. If you need advanced features such as managing, auditing, and monitoring the database, select ADB.
	//
	// > If you have not used ADB storage in Alibaba Cloud Model Studio, go to the <props="china">[Create Knowledge Base](https://bailian.console.aliyun.com/#/knowledge-base/create)<props="intl">[Create Knowledge Base](https://bailian.console.alibabacloud.com/#/knowledge-base/create) page, select ADB-PG as the vector storage type, and complete the authorization as prompted. If you pass ADB, you must specify the SinkInstanceId and SinkRegion parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// BUILT_IN
	SinkType *string `json:"SinkType,omitempty" xml:"SinkType,omitempty"`
	// 	Notice: This parameter is required in the latest SDK version. Otherwise, an error is reported when you call the SubmitIndexJob operation: Required parameter(data_sources) missing or invalid.</notice>
	//
	// The source of the imported data. Valid values:
	//
	// - DATA_CENTER_CATEGORY: the category type. This value indicates that all files under a specified category in <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center) are imported. You can import multiple categories at a time.
	//
	// - DATA_CENTER_FILE: the file type. This value indicates that specified files in <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center) are imported. You can import multiple files at a time.
	//
	// > If you set this parameter to DATA_CENTER_CATEGORY, you must specify the CategoryIds parameter. If you set this parameter to DATA_CENTER_FILE, you must specify the DocumentIds parameter.
	//
	// >
	//
	// > To create an empty knowledge base, you can use an empty category that contains no files: set this parameter to DATA_CENTER_CATEGORY and set CategoryIds to the ID of the empty category.
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
	// The type of the knowledge base.
	//
	// **Valid values**:
	//
	// - unstructured: a knowledge base for document search or audio and video. The default use scenario for the document search type is basic document question answering. <props="china">To create a knowledge base for other scenarios, you must also pass the knowledgeType and knowledgeScene parameters.
	//
	// <props="china">
	//
	// - structured: a knowledge base for data querying or image question answering.
	//
	//
	//
	// > Note that the type of a knowledge base cannot be changed after it is created.
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
	// You can obtain the ID by clicking the ID icon next to the table name on the Tables tab of the Table Connector in [Data Connection](https://bailian.console.aliyun.com/cn-beijing?tab=app#/connector/list). If the list contains multiple IDs, only the first one is used.
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
	// Enables custom chunking and specifies the chunking strategy. For more information, refer to [Knowledge base](https://help.aliyun.com/document_detail/2807740.html).
	//
	// Valid values (you cannot pass multiple values at a time):
	//
	// - **length**: chunks by length. The text is strictly chunked according to the ChunkSize and OverlapSize you specify. If you do not pass these two parameters, the system uses the default values (ChunkSize is 500 and OverlapSize is 100). Chunking by length does not support Separator (it does not take effect even if passed).
	//
	// - **page**: chunks by page. If ChunkSize is specified, it is considered during chunking (if not passed, the default value 500 is used). Chunking by page does not support OverlapSize and Separator (they do not take effect even if passed).
	//
	// - **h1**: chunks by level-1 heading. If ChunkSize is specified, it is considered during chunking (if not passed, the default value 500 is used). Chunking by level-1 heading does not support OverlapSize and Separator (they do not take effect even if passed).
	//
	// - **h2**: chunks by level-2 heading. If ChunkSize is specified, it is considered during chunking (if not passed, the default value 500 is used). Chunking by level-2 heading does not support OverlapSize and Separator (they do not take effect even if passed).
	//
	// - **regex**: chunks by regular expression. In this case, you must specify the Separator parameter. If ChunkSize is specified, it is considered during chunking (if not passed, the default value 500 is used). Chunking by regular expression does not support OverlapSize (it does not take effect even if passed).
	//
	// Default value: empty. In this case, intelligent chunking is used.
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
	// Specifies whether to use the first row of all .xlsx and .xls files as the header and concatenate it to each text chunk to prevent the large language model from mistakenly treating the header as a normal data row.
	//
	//
	// > Enable this feature only when all imported files are in .xlsx or .xls format and contain headers. Otherwise, do not enable it.
	//
	// >
	//
	// Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableHeaders  *bool   `json:"enableHeaders,omitempty" xml:"enableHeaders,omitempty"`
	KnowledgeScene *string `json:"knowledgeScene,omitempty" xml:"knowledgeScene,omitempty"`
	// <props="china">
	//
	// The specific knowledge type, which is used to further specify the type of data processed by the knowledge base.
	//
	// <notice>This parameter must be provided or omitted together with knowledgeScene. It cannot be set separately. If both are omitted, the system uses the default configurations based on structureType.</notice>
	//
	// **Settings*	- constraints: The value of this parameter must match the structureType you select and determines the active values for knowledgeScene.
	//
	// **Valid values**:
	//
	// - document: document search. Must be used with structureType: unstructured.
	//
	// - table: data querying. Must be used with structureType: structured.
	//
	// - image: image question answering. Must be used with structureType: structured.
	//
	// - multimedia: audio and video search. Must be used with structureType: unstructured.
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
	// The metadata extraction configuration. Metadata is a series of additional attributes related to the content of unstructured data. These attributes are integrated into text chunks in the form of key-value pairs. For more information, refer to [Knowledge base](https://help.aliyun.com/document_detail/2807740.html).
	MetaExtractColumns []*CreateIndexRequestMetaExtractColumns `json:"metaExtractColumns,omitempty" xml:"metaExtractColumns,omitempty" type:"Repeated"`
	// <props="china">The number of RCUs for the knowledge base. This parameter is required only when pipelineCommercialType is set to enterprise. Valid values: [1, 200].
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
	// The [edition](https://help.aliyun.com/document_detail/2997110.html) of the knowledge base. Valid values:
	//
	// - standard: Standard Edition
	//
	// - enterprise: Ultimate Edition
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
	// <props="china">The throttling handling strategy for the knowledge base dependency link. This parameter is required only when pipelineCommercialType is set to enterprise.
	//
	// Valid values:
	//
	// downgrade: downgrade processing (switches to lightweight link retrieval).
	//
	// Default value: downgrade.
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
	// Specifies whether the column participates in model responses. If this parameter is set to true, the retrieval results of this column are used as input information when the large language model generates responses. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
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
	// Specifies whether the column participates in knowledge base retrieval. If this parameter is set to true, the knowledge base is allowed to search in the data of this column. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
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
	// The field name. The name must be consistent with the header of the data table created in Application Data.
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
	// School
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// <props="china">
	//
	// The field type. The type must be consistent with the header of the data table created in Application Data. Valid values:
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
	// The Chinese description of the metadata field. The description must be 0 to 1,000 characters in length and can contain Chinese characters, English letters, digits, underscores (_), hyphens (-), periods (.), and colons (:). Default value: empty.
	//
	// example:
	//
	// Author name
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// Specifies whether the metadata field and its value participate in the response generation process of the large language model together with the content of the text chunk. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableLlm *bool `json:"EnableLlm,omitempty" xml:"EnableLlm,omitempty"`
	// Specifies whether the metadata field and its value participate in knowledge base retrieval together with the content of the text chunk. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableSearch *bool `json:"EnableSearch,omitempty" xml:"EnableSearch,omitempty"`
	// The metadata field. The field must be 1 to 50 characters in length and can contain only English letters and underscores (_). If you specify this parameter, you must also specify the Value and Type parameters.
	//
	// example:
	//
	// author
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value extraction method for the metadata field. Valid values:
	//
	// - constant: constant.
	//
	// - variable: variable.
	//
	// - custom_prompt: large language model.
	//
	// - regular: regular expression.
	//
	// - keywords: keyword search.
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
