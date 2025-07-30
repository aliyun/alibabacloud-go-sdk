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
	SetEnableHeaders(v bool) *CreateIndexRequest
	GetEnableHeaders() *bool
	SetMetaExtractColumns(v []*CreateIndexRequestMetaExtractColumns) *CreateIndexRequest
	GetMetaExtractColumns() []*CreateIndexRequestMetaExtractColumns
}

type CreateIndexRequest struct {
	// The list of primary key IDs of the categories to be imported into the knowledge base.
	CategoryIds []*string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	// The estimated length of chunks. The maximum number of characters for a chunk. Texts exceeding this limit are splited. For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base). Valid values: [1-2048].
	//
	// The default value is empty, which means using the intelligent splitting method.
	//
	// >  If you specify the `ChunkSize` parameter, you must also specify the `OverlapSize` and `Separator` parameters. If you do not specify these three parameters, the system uses the intelligent splitting method by default.
	//
	// example:
	//
	// 128
	ChunkSize       *int32                       `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	Columns         []*CreateIndexRequestColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	CreateIndexType *string                      `json:"CreateIndexType,omitempty" xml:"CreateIndexType,omitempty"`
	// >  This parameter is not available. Do not specify this parameter.
	DataSource *CreateIndexRequestDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description of the knowledge base. The description must be 0 to 1,000 characters in length. This parameter is empty by default.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of primary key IDs of the documents to be imported into the knowledge base.
	DocumentIds []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
	// The name of the embedding model. The embedding model converts the original input prompt and knowledge text into numerical vectors for similarity comparison. The default and only model available is DashScope text-embedding-v2. It supports multiple languages including Chinese and English and normalizes the vector results. For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base). Valid value:
	//
	// 	- text-embedding-v2
	//
	// The default value is null, which means using the text-embedding-v2 model.
	//
	// example:
	//
	// text-embedding-v2
	EmbeddingModelName *string `json:"EmbeddingModelName,omitempty" xml:"EmbeddingModelName,omitempty"`
	EnableRewrite      *bool   `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	// The name of the knowledge base. The name must be 1 to 20 characters in length and can contain characters classified as letter in Unicode, including English letters, Chinese characters, digits, among others. The name can also contain colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The overlap length. The number of overlapping characters between two consecutive chunks. For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base). Valid values: 0 to 1024.
	//
	// The default value is empty, which means using the intelligent splitting method.
	//
	// example:
	//
	// 16
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// Similarity Threshold. The lowest similarity score of chunks that can be returned. This parameter is used to filter text chunks returned by the rank model. For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base). Valid values: [0.01-1.00].
	//
	// Default value: 0.20.
	//
	// example:
	//
	// 0.20
	RerankMinScore *float64 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// The name of the rank model. The rank model is a scoring system outside the knowledge base. It calculates the similarity score of each text chunk in the input question and knowledge base and ranks them in descending order. Then, the model returns the top K chunks with the highest scores. For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base). Valid values:
	//
	// 	- gte-rerank-hybrid
	//
	// 	- gte-rerank
	//
	// The default value is empty, which means using the official gte-rerank-hybrid model.
	//
	// >  If you need only semantic ranking, we recommend that you use gte-rerank. If you need both semantic ranking and text matching features to ensure relevance, we recommend that you use gte-rerank-hybrid.
	//
	// example:
	//
	// gte-rerank-hybrid
	RerankModelName *string `json:"RerankModelName,omitempty" xml:"RerankModelName,omitempty"`
	// The clause identifier. The document is split into chunks based on this identifier. For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base). You can specify multiple identifiers and do not need to add any other characters to separate them. For example: !,\\\\\\n. Valid values:
	//
	// 	- \\n: line break
	//
	// 	- ，: Chinese comma
	//
	// 	- ,: English comma
	//
	// 	- 。 : Chinese full stop
	//
	// 	- .: English full stop
	//
	// 	- ！ : Chinese exclamation point
	//
	// 	- ! : English exclamation point
	//
	// 	- ；: Chinese semicolon
	//
	// 	- ;: English semicolon
	//
	// 	- ？: Chinese question mark
	//
	// 	- ?: English question mark
	//
	// The default value is empty, which means using the intelligent splitting method.
	//
	// example:
	//
	// ,
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// The ID of the vector storage instance. This parameter is available only when SinkType is set to ADB. You can view the ID on the [Instances](https://gpdbnext.console.aliyun.com/gpdb/list) page of AnalyticDB for PostgreSQL.
	//
	// example:
	//
	// gp-bp321093j84
	SinkInstanceId *string `json:"SinkInstanceId,omitempty" xml:"SinkInstanceId,omitempty"`
	// The region of the vector storage instance. This parameter is available only when SinkType is set to ADB. You can call the [DescribeRegions](https://www.alibabacloud.com/help/en/analyticdb/analyticdb-for-postgresql/developer-reference/api-gpdb-2016-05-03-describeregions) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	SinkRegion *string `json:"SinkRegion,omitempty" xml:"SinkRegion,omitempty"`
	// The vector storage type of the knowledge base. For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base). Valid values:
	//
	// 	- DEFAULT: The built-in vector database.
	//
	// 	- ADB: AnalyticDB for PostgreSQL database. If you need advanced features, such as managing, auditing, and monitoring, we recommend that you specify ADB.
	//
	// >  If you have not used AnalyticDB for AnalyticDB in Model Studio before, go to the [Create Knowledge Base](https://bailian.console.aliyun.com/#/knowledge-base/create) page, select ADB-PG as Vector Storage Type, and follow the instructions to grant permissions. If you specify ADB, you must also specify the `SinkInstanceId` and `SinkRegion` parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEFAULT
	SinkType *string `json:"SinkType,omitempty" xml:"SinkType,omitempty"`
	// The data type of [Data Management](https://bailian.console.aliyun.com/#/data-center). For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base). Valid values:
	//
	// 	- DATA_CENTER_CATEGORY: The category type. Import all documents from one or more categories in Data Center.
	//
	// 	- DATA_CENTER_FILE: The document type. Import one or more documents from Data Center.
	//
	// >  If this parameter is set to DATA_CENTER_CATEGORY, you must specify the `CategoryIds` parameter. If this parameter is set to DATA_CENTER_FILE, you must specify the `DocumentIds` parameter.
	//
	// >  If you want to create an empty knowledge base, you can use an empty category. Set this parameter to DATA_CENTER_CATEGORY. And specify the ID of an empty category for the `CategoryIds` parameter.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// DATA_CENTER_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The data type of the knowledge base. For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base). Valid value:
	//
	// 	- unstructured
	//
	// >  After a knowledge base is created, its data type cannot be changed. You cannot create a structured knowledge base by calling an API operation. Use the console instead.
	//
	// This parameter is required.
	//
	// example:
	//
	// structured
	StructureType      *string                                 `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
	TableIds           []*string                               `json:"TableIds,omitempty" xml:"TableIds,omitempty" type:"Repeated"`
	ChunkMode          *string                                 `json:"chunkMode,omitempty" xml:"chunkMode,omitempty"`
	EnableHeaders      *bool                                   `json:"enableHeaders,omitempty" xml:"enableHeaders,omitempty"`
	MetaExtractColumns []*CreateIndexRequestMetaExtractColumns `json:"metaExtractColumns,omitempty" xml:"metaExtractColumns,omitempty" type:"Repeated"`
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

func (s *CreateIndexRequest) GetEnableHeaders() *bool {
	return s.EnableHeaders
}

func (s *CreateIndexRequest) GetMetaExtractColumns() []*CreateIndexRequestMetaExtractColumns {
	return s.MetaExtractColumns
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

func (s *CreateIndexRequest) SetEnableHeaders(v bool) *CreateIndexRequest {
	s.EnableHeaders = &v
	return s
}

func (s *CreateIndexRequest) SetMetaExtractColumns(v []*CreateIndexRequestMetaExtractColumns) *CreateIndexRequest {
	s.MetaExtractColumns = v
	return s
}

func (s *CreateIndexRequest) Validate() error {
	return dara.Validate(s)
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
	Desc         *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	EnableLlm    *bool   `json:"EnableLlm,omitempty" xml:"EnableLlm,omitempty"`
	EnableSearch *bool   `json:"EnableSearch,omitempty" xml:"EnableSearch,omitempty"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
