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
	SetEnableHeaders(v bool) *CreateIndexShrinkRequest
	GetEnableHeaders() *bool
	SetMetaExtractColumnsShrink(v string) *CreateIndexShrinkRequest
	GetMetaExtractColumnsShrink() *string
}

type CreateIndexShrinkRequest struct {
	// The list of primary key IDs of the categories to be imported into the knowledge base.
	CategoryIdsShrink *string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty"`
	// The estimated length of chunks. The maximum number of characters for a chunk. Texts exceeding this limit are splited. For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base). Valid values: [1-2048].
	//
	// The default value is empty, which means using the intelligent splitting method.
	//
	// >  If you specify the `ChunkSize` parameter, you must also specify the `OverlapSize` and `Separator` parameters. If you do not specify these three parameters, the system uses the intelligent splitting method by default.
	//
	// example:
	//
	// 128
	ChunkSize       *int32  `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	ColumnsShrink   *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	CreateIndexType *string `json:"CreateIndexType,omitempty" xml:"CreateIndexType,omitempty"`
	// >  This parameter is not available. Do not specify this parameter.
	DataSourceShrink *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// The description of the knowledge base. The description must be 0 to 1,000 characters in length. This parameter is empty by default.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of primary key IDs of the documents to be imported into the knowledge base.
	DocumentIdsShrink *string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty"`
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
	StructureType            *string `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
	TableIdsShrink           *string `json:"TableIds,omitempty" xml:"TableIds,omitempty"`
	ChunkMode                *string `json:"chunkMode,omitempty" xml:"chunkMode,omitempty"`
	EnableHeaders            *bool   `json:"enableHeaders,omitempty" xml:"enableHeaders,omitempty"`
	MetaExtractColumnsShrink *string `json:"metaExtractColumns,omitempty" xml:"metaExtractColumns,omitempty"`
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

func (s *CreateIndexShrinkRequest) GetEnableHeaders() *bool {
	return s.EnableHeaders
}

func (s *CreateIndexShrinkRequest) GetMetaExtractColumnsShrink() *string {
	return s.MetaExtractColumnsShrink
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

func (s *CreateIndexShrinkRequest) SetEnableHeaders(v bool) *CreateIndexShrinkRequest {
	s.EnableHeaders = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetMetaExtractColumnsShrink(v string) *CreateIndexShrinkRequest {
	s.MetaExtractColumnsShrink = &v
	return s
}

func (s *CreateIndexShrinkRequest) Validate() error {
	return dara.Validate(s)
}
