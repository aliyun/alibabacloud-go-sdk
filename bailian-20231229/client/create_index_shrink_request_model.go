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
	SetDataSourceShrink(v string) *CreateIndexShrinkRequest
	GetDataSourceShrink() *string
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
	SetRerankMinScore(v float64) *CreateIndexShrinkRequest
	GetRerankMinScore() *float64
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
	SetChunkMode(v string) *CreateIndexShrinkRequest
	GetChunkMode() *string
	SetDatabase(v string) *CreateIndexShrinkRequest
	GetDatabase() *string
	SetDatasourceCode(v string) *CreateIndexShrinkRequest
	GetDatasourceCode() *string
	SetEnableHeaders(v bool) *CreateIndexShrinkRequest
	GetEnableHeaders() *bool
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
	// The files to imported to the knowledge base. Specify the category IDs. All files under the categories will be imported (up to 10,000 files). To add more files later, call **SubmitIndexAddDocumentsJob**.
	CategoryIdsShrink *string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty"`
	// The chunk size, which is the maximum number of characters in each chunk. Text exceeding this length may be truncated.
	//
	// Valid values: 1 to 6000. Default value: 500.
	//
	// > If `ChunkSize` is set to a value less than 100, `OverlapSize` is required. Or, if you do not pass these two parameters, the system uses the default values of the two.
	//
	// example:
	//
	// 128
	ChunkSize     *int32  `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	ColumnsShrink *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	// > This parameter is not available. Do not specify this parameter.
	CreateIndexType *string `json:"CreateIndexType,omitempty" xml:"CreateIndexType,omitempty"`
	// >  This parameter is not available. Do not specify this parameter.
	DataSourceShrink *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// The description of the knowledge base. The description must be 0 to 1,000 characters in length. This parameter is empty by default.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The files to imported to the knowledge base. Specify the file IDs to import (up to 10,000 files). To add more files later, call **SubmitIndexAddDocumentsJob**.
	DocumentIdsShrink *string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty"`
	// The embedding model used in the knowledge base. The embedding model converts the original input prompt and knowledge text into numerical embeddings for similarity comparison. The default and only model available is text-embedding-v2. It supports multiple languages including Chinese and English and normalizes the embedding results. For more information, see [Embedding](https://help.aliyun.com/document_detail/2842587.html). Valid values:
	//
	// 	- text-embedding-v2
	//
	// The default value is null, which means using text-embedding-v2.
	//
	// example:
	//
	// text-embedding-v2
	EmbeddingModelName *string `json:"EmbeddingModelName,omitempty" xml:"EmbeddingModelName,omitempty"`
	// Whether to enable rewriting for multi-turn conversations. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: true.
	//
	// example:
	//
	// true
	EnableRewrite *bool `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	// The name of the knowledge base. The name must be 1 to 20 characters in length, and can contain Chinese characters, letters, digits, underscores (_), hyphens (-), periods (.), and colons (:).
	//
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The overlap size, which is the number of overlapping characters between two consecutive chunks. Valid values: 0 to 1024.
	//
	// Default value: 100.
	//
	// > `OverlapSize` must be less than `ChunkSize`. Otherwise, chunking errors may occur.
	//
	// example:
	//
	// 16
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// The similarity threshold. Only chunks with a similarity score higher than this value can be recalled. This parameter is used to filter chunks returned by the re-rank model. Valid values: 0.01 to 1.00.
	//
	// Default value: 0.01.
	//
	// example:
	//
	// 0.20
	RerankMinScore *float64 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// The re-ranking model used in the knowledge base. The re-rank model is a scoring system outside the knowledge base. It calculates the similarity score of the query and text chunks in the knowledge base and ranks them in descending order. Then, the model returns the top K chunks with the highest scores. Valid values:
	//
	// 	- gte-rerank-hybrid
	//
	// 	- gte-rerank
	//
	// The default value is empty, which means using gte-rerank-hybrid.
	//
	// > If you need only semantic ranking, we recommend gte-rerank. If you need both semantic ranking and text matching features to ensure relevance, we recommend gte-rerank-hybrid.
	//
	// example:
	//
	// gte-rerank-hybrid
	RerankModelName *string `json:"RerankModelName,omitempty" xml:"RerankModelName,omitempty"`
	// > This parameter is not available. Do not specify this parameter.
	//
	// example:
	//
	// ,
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// The ID of the AnalyticDB for PostgreSQL instance. Required only when `SinkType` is set to ADB. Get the ID on the [Instances](https://gpdbnext.console.aliyun.com/gpdb/list) page of AnalyticDB for PostgreSQL.
	//
	// example:
	//
	// gp-bp321093j84
	SinkInstanceId *string `json:"SinkInstanceId,omitempty" xml:"SinkInstanceId,omitempty"`
	// The region of the AnalyticDB for PostgreSQL instance. Required only when `SinkType` is set to ADB. Call [DescribeRegions](https://www.alibabacloud.com/help/zh/analyticdb/analyticdb-for-postgresql/developer-reference/api-gpdb-2016-05-03-describeregions?spm=a2c63.p38356.0.i3) to obtain the region list.
	//
	// example:
	//
	// cn-hangzhou
	SinkRegion *string `json:"SinkRegion,omitempty" xml:"SinkRegion,omitempty"`
	// The vector storage type of the knowledge base. For more information, see [Knowledge base](https://help.aliyun.com/document_detail/2807740.html). Valid values:
	//
	// 	- BUILT_IN: The vector data is hosted by Alibaba Cloud Model Studio.
	//
	// 	- ADB: AnalyticDB for PostgreSQL database. If you need advanced features, such as managing, auditing, and monitoring, we recommend ADB.
	//
	// > If you have not used AnalyticDB for AnalyticDB in Model Studio before, go to the [Create Knowledge Base](https://bailian.console.alibabacloud.com/#/knowledge-base/create) page, select ADB-PG as Vector Storage Type, and follow the instructions to grant permissions. If you specify ADB, the `SinkInstanceId` and `SinkRegion` parameters are required.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEFAULT
	SinkType *string `json:"SinkType,omitempty" xml:"SinkType,omitempty"`
	// > This parameter is required in the latest version of the SDK. Otherwise, when you call SubmitIndexJob, an error will occur: Required parameter(data_sources) missing or invalid.
	//
	// The source of the imported data. Valid values:
	//
	// 	- DATA_CENTER_CATEGORY: The category type, that is to import all files in one or more specified categories in [Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center).
	//
	// 	- DATA_CENTER_FILE: The file type, that is to import one or more specified files in [Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center).
	//
	// > If set to DATA_CENTER_CATEGORY, `CategoryIds` is required. If set to DATA_CENTER_FILE, `DocumentIds` is required.
	//
	// > To create an empty knowledge base, you can use an empty category with no files: Set this parameter to DATA_CENTER_CATEGORY, and `CategoryIds` to the ID of an empty category.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// DATA_CENTER_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The type of the knowledge base. Valid values:
	//
	// 	- unstructured: The document search type.
	//
	// > After you create a knowledge base, its type cannot be changed. This operation does not support data query and image Q\\&A types. Use the console instead.
	//
	// This parameter is required.
	//
	// example:
	//
	// structured
	StructureType *string `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
	// > This parameter is not available. Do not specify this parameter.
	TableIdsShrink *string `json:"TableIds,omitempty" xml:"TableIds,omitempty"`
	// > This parameter is not available. Do not specify this parameter.
	//
	// example:
	//
	// regex
	ChunkMode      *string `json:"chunkMode,omitempty" xml:"chunkMode,omitempty"`
	Database       *string `json:"database,omitempty" xml:"database,omitempty"`
	DatasourceCode *string `json:"datasourceCode,omitempty" xml:"datasourceCode,omitempty"`
	// Whether to treat the first row of all .xlsx and .xls files as headers and concatenate them into each text chunk. This prevents the models from mistakenly interpreting headers as regular data rows.
	//
	// > Enable this feature only when all imported files are in .xlsx or .xls format and contain headers. Otherwise, leave it disabled.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableHeaders *bool `json:"enableHeaders,omitempty" xml:"enableHeaders,omitempty"`
	// The metadata extraction configurations. Metadata refers to a set of additional attributes associated with unstructured data, which are integrated into text chunks in key-value pairs. For more information, see [Knowledge base](https://help.aliyun.com/document_detail/2807740.html).
	MetaExtractColumnsShrink *string `json:"metaExtractColumns,omitempty" xml:"metaExtractColumns,omitempty"`
	// example:
	//
	// 1
	PipelineCommercialCu *int32 `json:"pipelineCommercialCu,omitempty" xml:"pipelineCommercialCu,omitempty"`
	// example:
	//
	// standard
	PipelineCommercialType *string `json:"pipelineCommercialType,omitempty" xml:"pipelineCommercialType,omitempty"`
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

func (s *CreateIndexShrinkRequest) GetDataSourceShrink() *string {
	return s.DataSourceShrink
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

func (s *CreateIndexShrinkRequest) GetRerankMinScore() *float64 {
	return s.RerankMinScore
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

func (s *CreateIndexShrinkRequest) GetChunkMode() *string {
	return s.ChunkMode
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

func (s *CreateIndexShrinkRequest) SetDataSourceShrink(v string) *CreateIndexShrinkRequest {
	s.DataSourceShrink = &v
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

func (s *CreateIndexShrinkRequest) SetRerankMinScore(v float64) *CreateIndexShrinkRequest {
	s.RerankMinScore = &v
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

func (s *CreateIndexShrinkRequest) SetChunkMode(v string) *CreateIndexShrinkRequest {
	s.ChunkMode = &v
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
