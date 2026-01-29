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
	SetDataSource(v *CreateIndexRequestDataSource) *CreateIndexRequest
	GetDataSource() *CreateIndexRequestDataSource
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
	SetRerankMinScore(v float64) *CreateIndexRequest
	GetRerankMinScore() *float64
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
	// The files to imported to the knowledge base. Specify the category IDs. All files under the categories will be imported (up to 10,000 files). To add more files later, call **SubmitIndexAddDocumentsJob**.
	CategoryIds []*string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	// The chunk size, which is the maximum number of characters in each chunk. Text exceeding this length may be truncated.
	//
	// Valid values: 1 to 6000. Default value: 500.
	//
	// > If `ChunkSize` is set to a value less than 100, `OverlapSize` is required. Or, if you do not pass these two parameters, the system uses the default values of the two.
	//
	// example:
	//
	// 128
	ChunkSize *int32                       `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	Columns   []*CreateIndexRequestColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// > This parameter is not available. Do not specify this parameter.
	CreateIndexType *string `json:"CreateIndexType,omitempty" xml:"CreateIndexType,omitempty"`
	// >  This parameter is not available. Do not specify this parameter.
	DataSource *CreateIndexRequestDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description of the knowledge base. The description must be 0 to 1,000 characters in length. This parameter is empty by default.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The files to imported to the knowledge base. Specify the file IDs to import (up to 10,000 files). To add more files later, call **SubmitIndexAddDocumentsJob**.
	DocumentIds []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
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
	TableIds []*string `json:"TableIds,omitempty" xml:"TableIds,omitempty" type:"Repeated"`
	// > This parameter is not available. Do not specify this parameter.
	//
	// example:
	//
	// regex
	ChunkMode      *string `json:"chunkMode,omitempty" xml:"chunkMode,omitempty"`
	ConnectId      *string `json:"connectId,omitempty" xml:"connectId,omitempty"`
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
	EnableHeaders  *bool   `json:"enableHeaders,omitempty" xml:"enableHeaders,omitempty"`
	KnowledgeScene *string `json:"knowledgeScene,omitempty" xml:"knowledgeScene,omitempty"`
	// The metadata extraction configurations. Metadata refers to a set of additional attributes associated with unstructured data, which are integrated into text chunks in key-value pairs. For more information, see [Knowledge base](https://help.aliyun.com/document_detail/2807740.html).
	MetaExtractColumns []*CreateIndexRequestMetaExtractColumns `json:"metaExtractColumns,omitempty" xml:"metaExtractColumns,omitempty" type:"Repeated"`
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

func (s *CreateIndexRequest) GetDataSource() *CreateIndexRequestDataSource {
	return s.DataSource
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

func (s *CreateIndexRequest) GetRerankMinScore() *float64 {
	return s.RerankMinScore
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

func (s *CreateIndexRequest) SetDataSource(v *CreateIndexRequestDataSource) *CreateIndexRequest {
	s.DataSource = v
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

func (s *CreateIndexRequest) SetRerankMinScore(v float64) *CreateIndexRequest {
	s.RerankMinScore = &v
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
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
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
	Column   *string `json:"Column,omitempty" xml:"Column,omitempty"`
	IsRecall *bool   `json:"IsRecall,omitempty" xml:"IsRecall,omitempty"`
	IsSearch *bool   `json:"IsSearch,omitempty" xml:"IsSearch,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

type CreateIndexRequestDataSource struct {
	// >  This parameter is not available. Do not specify this parameter.
	CredentialId *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	// >  This parameter is not available. Do not specify this parameter.
	CredentialKey *string `json:"CredentialKey,omitempty" xml:"CredentialKey,omitempty"`
	// >  This parameter is not available. Do not specify this parameter.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// >  This parameter is not available. Do not specify this parameter.
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// >  This parameter is not available. Do not specify this parameter.
	IsPrivateLink *bool `json:"IsPrivateLink,omitempty" xml:"IsPrivateLink,omitempty"`
	// >  This parameter is not available. Do not specify this parameter.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// >  This parameter is not available. Do not specify this parameter.
	SubPath *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
	// >  This parameter is not available. Do not specify this parameter.
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// >  This parameter is not available. Do not specify this parameter.
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
	// >  This parameter is not available. Do not specify this parameter.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateIndexRequestDataSource) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexRequestDataSource) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestDataSource) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateIndexRequestDataSource) GetCredentialKey() *string {
	return s.CredentialKey
}

func (s *CreateIndexRequestDataSource) GetDatabase() *string {
	return s.Database
}

func (s *CreateIndexRequestDataSource) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateIndexRequestDataSource) GetIsPrivateLink() *bool {
	return s.IsPrivateLink
}

func (s *CreateIndexRequestDataSource) GetRegion() *string {
	return s.Region
}

func (s *CreateIndexRequestDataSource) GetSubPath() *string {
	return s.SubPath
}

func (s *CreateIndexRequestDataSource) GetSubType() *string {
	return s.SubType
}

func (s *CreateIndexRequestDataSource) GetTable() *string {
	return s.Table
}

func (s *CreateIndexRequestDataSource) GetType() *string {
	return s.Type
}

func (s *CreateIndexRequestDataSource) SetCredentialId(v string) *CreateIndexRequestDataSource {
	s.CredentialId = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetCredentialKey(v string) *CreateIndexRequestDataSource {
	s.CredentialKey = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetDatabase(v string) *CreateIndexRequestDataSource {
	s.Database = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetEndpoint(v string) *CreateIndexRequestDataSource {
	s.Endpoint = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetIsPrivateLink(v bool) *CreateIndexRequestDataSource {
	s.IsPrivateLink = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetRegion(v string) *CreateIndexRequestDataSource {
	s.Region = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetSubPath(v string) *CreateIndexRequestDataSource {
	s.SubPath = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetSubType(v string) *CreateIndexRequestDataSource {
	s.SubType = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetTable(v string) *CreateIndexRequestDataSource {
	s.Table = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetType(v string) *CreateIndexRequestDataSource {
	s.Type = &v
	return s
}

func (s *CreateIndexRequestDataSource) Validate() error {
	return dara.Validate(s)
}

type CreateIndexRequestMetaExtractColumns struct {
	// The description of the metadata field. The description must be 0 to 1,000 characters in length, and can contain Chinese characters, letters, digits, underscores (_), hyphens (-), periods (.), and colons (:). This parameter is left empty by default.
	//
	// example:
	//
	// AuthorName
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// When set to true, the key and value of this metadata filed will participate in the generation process of the model, together with the chunk. Valid values:
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
	EnableLlm *bool `json:"EnableLlm,omitempty" xml:"EnableLlm,omitempty"`
	// When set to true, the key and value of this metadata filed will participate in the knowledge base retrieval, together with the chunk. Valid values:
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
	EnableSearch *bool `json:"EnableSearch,omitempty" xml:"EnableSearch,omitempty"`
	// The metadata key. It must be 1 to 50 characters in length and must be English letters or underscores. If you specify this parameter, the `Value` and `Type` parameters are required.
	//
	// example:
	//
	// author
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The type of the metadata field. Valid values:
	//
	// 	- constant
	//
	// 	- variable
	//
	// 	- custom_prompt
	//
	// 	- regular
	//
	// 	- keywords
	//
	// Enumerated value:
	//
	// 	- constant: constant extraction.
	//
	// 	- keywords: keyword extraction.
	//
	// 	- custom_prompt: LLM.
	//
	// 	- variable: variable extraction.
	//
	// 	- regular: regular expression.
	//
	// example:
	//
	// constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The metadata value.
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
