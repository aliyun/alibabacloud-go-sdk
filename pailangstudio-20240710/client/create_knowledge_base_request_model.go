// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreateKnowledgeBaseRequest
	GetAccessibility() *string
	SetChunkConfig(v *CreateKnowledgeBaseRequestChunkConfig) *CreateKnowledgeBaseRequest
	GetChunkConfig() *CreateKnowledgeBaseRequestChunkConfig
	SetDataSources(v []*CreateKnowledgeBaseRequestDataSources) *CreateKnowledgeBaseRequest
	GetDataSources() []*CreateKnowledgeBaseRequestDataSources
	SetDescription(v string) *CreateKnowledgeBaseRequest
	GetDescription() *string
	SetEmbeddingConfig(v *CreateKnowledgeBaseRequestEmbeddingConfig) *CreateKnowledgeBaseRequest
	GetEmbeddingConfig() *CreateKnowledgeBaseRequestEmbeddingConfig
	SetIndexColumnConfig(v *CreateKnowledgeBaseRequestIndexColumnConfig) *CreateKnowledgeBaseRequest
	GetIndexColumnConfig() *CreateKnowledgeBaseRequestIndexColumnConfig
	SetKnowledgeBaseType(v string) *CreateKnowledgeBaseRequest
	GetKnowledgeBaseType() *string
	SetMetaDataConfig(v *CreateKnowledgeBaseRequestMetaDataConfig) *CreateKnowledgeBaseRequest
	GetMetaDataConfig() *CreateKnowledgeBaseRequestMetaDataConfig
	SetName(v string) *CreateKnowledgeBaseRequest
	GetName() *string
	SetOutputDir(v string) *CreateKnowledgeBaseRequest
	GetOutputDir() *string
	SetRuntimeId(v string) *CreateKnowledgeBaseRequest
	GetRuntimeId() *string
	SetVectorDBConfig(v *CreateKnowledgeBaseRequestVectorDBConfig) *CreateKnowledgeBaseRequest
	GetVectorDBConfig() *CreateKnowledgeBaseRequestVectorDBConfig
	SetWorkspaceId(v string) *CreateKnowledgeBaseRequest
	GetWorkspaceId() *string
}

type CreateKnowledgeBaseRequest struct {
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// This parameter is required.
	ChunkConfig *CreateKnowledgeBaseRequestChunkConfig `json:"ChunkConfig,omitempty" xml:"ChunkConfig,omitempty" type:"Struct"`
	// This parameter is required.
	DataSources []*CreateKnowledgeBaseRequestDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// example:
	//
	// This is a description of the knowledge base.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	EmbeddingConfig   *CreateKnowledgeBaseRequestEmbeddingConfig   `json:"EmbeddingConfig,omitempty" xml:"EmbeddingConfig,omitempty" type:"Struct"`
	IndexColumnConfig *CreateKnowledgeBaseRequestIndexColumnConfig `json:"IndexColumnConfig,omitempty" xml:"IndexColumnConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// TEXT
	KnowledgeBaseType *string                                   `json:"KnowledgeBaseType,omitempty" xml:"KnowledgeBaseType,omitempty"`
	MetaDataConfig    *CreateKnowledgeBaseRequestMetaDataConfig `json:"MetaDataConfig,omitempty" xml:"MetaDataConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou-internal.aliyuncs.com/langstudio/output/
	OutputDir *string `json:"OutputDir,omitempty" xml:"OutputDir,omitempty"`
	// example:
	//
	// rtime-apje******beaz
	RuntimeId *string `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// This parameter is required.
	VectorDBConfig *CreateKnowledgeBaseRequestVectorDBConfig `json:"VectorDBConfig,omitempty" xml:"VectorDBConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateKnowledgeBaseRequest) GetChunkConfig() *CreateKnowledgeBaseRequestChunkConfig {
	return s.ChunkConfig
}

func (s *CreateKnowledgeBaseRequest) GetDataSources() []*CreateKnowledgeBaseRequestDataSources {
	return s.DataSources
}

func (s *CreateKnowledgeBaseRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseRequest) GetEmbeddingConfig() *CreateKnowledgeBaseRequestEmbeddingConfig {
	return s.EmbeddingConfig
}

func (s *CreateKnowledgeBaseRequest) GetIndexColumnConfig() *CreateKnowledgeBaseRequestIndexColumnConfig {
	return s.IndexColumnConfig
}

func (s *CreateKnowledgeBaseRequest) GetKnowledgeBaseType() *string {
	return s.KnowledgeBaseType
}

func (s *CreateKnowledgeBaseRequest) GetMetaDataConfig() *CreateKnowledgeBaseRequestMetaDataConfig {
	return s.MetaDataConfig
}

func (s *CreateKnowledgeBaseRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseRequest) GetOutputDir() *string {
	return s.OutputDir
}

func (s *CreateKnowledgeBaseRequest) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *CreateKnowledgeBaseRequest) GetVectorDBConfig() *CreateKnowledgeBaseRequestVectorDBConfig {
	return s.VectorDBConfig
}

func (s *CreateKnowledgeBaseRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateKnowledgeBaseRequest) SetAccessibility(v string) *CreateKnowledgeBaseRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetChunkConfig(v *CreateKnowledgeBaseRequestChunkConfig) *CreateKnowledgeBaseRequest {
	s.ChunkConfig = v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetDataSources(v []*CreateKnowledgeBaseRequestDataSources) *CreateKnowledgeBaseRequest {
	s.DataSources = v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetDescription(v string) *CreateKnowledgeBaseRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetEmbeddingConfig(v *CreateKnowledgeBaseRequestEmbeddingConfig) *CreateKnowledgeBaseRequest {
	s.EmbeddingConfig = v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetIndexColumnConfig(v *CreateKnowledgeBaseRequestIndexColumnConfig) *CreateKnowledgeBaseRequest {
	s.IndexColumnConfig = v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetKnowledgeBaseType(v string) *CreateKnowledgeBaseRequest {
	s.KnowledgeBaseType = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetMetaDataConfig(v *CreateKnowledgeBaseRequestMetaDataConfig) *CreateKnowledgeBaseRequest {
	s.MetaDataConfig = v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetName(v string) *CreateKnowledgeBaseRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetOutputDir(v string) *CreateKnowledgeBaseRequest {
	s.OutputDir = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetRuntimeId(v string) *CreateKnowledgeBaseRequest {
	s.RuntimeId = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetVectorDBConfig(v *CreateKnowledgeBaseRequestVectorDBConfig) *CreateKnowledgeBaseRequest {
	s.VectorDBConfig = v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetWorkspaceId(v string) *CreateKnowledgeBaseRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) Validate() error {
	if s.ChunkConfig != nil {
		if err := s.ChunkConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DataSources != nil {
		for _, item := range s.DataSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EmbeddingConfig != nil {
		if err := s.EmbeddingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.IndexColumnConfig != nil {
		if err := s.IndexColumnConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MetaDataConfig != nil {
		if err := s.MetaDataConfig.Validate(); err != nil {
			return err
		}
	}
	if s.VectorDBConfig != nil {
		if err := s.VectorDBConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKnowledgeBaseRequestChunkConfig struct {
	// 分块时长
	//
	// example:
	//
	// 30
	ChunkDuration *int32 `json:"ChunkDuration,omitempty" xml:"ChunkDuration,omitempty"`
	// 分块重叠大小
	//
	// example:
	//
	// 200
	ChunkOverlap *int32 `json:"ChunkOverlap,omitempty" xml:"ChunkOverlap,omitempty"`
	// 分块大小
	//
	// example:
	//
	// 1024
	ChunkSize *int32 `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// 分块策略
	//
	// example:
	//
	// Default
	ChunkStrategy *string `json:"ChunkStrategy,omitempty" xml:"ChunkStrategy,omitempty"`
}

func (s CreateKnowledgeBaseRequestChunkConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequestChunkConfig) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequestChunkConfig) GetChunkDuration() *int32 {
	return s.ChunkDuration
}

func (s *CreateKnowledgeBaseRequestChunkConfig) GetChunkOverlap() *int32 {
	return s.ChunkOverlap
}

func (s *CreateKnowledgeBaseRequestChunkConfig) GetChunkSize() *int32 {
	return s.ChunkSize
}

func (s *CreateKnowledgeBaseRequestChunkConfig) GetChunkStrategy() *string {
	return s.ChunkStrategy
}

func (s *CreateKnowledgeBaseRequestChunkConfig) SetChunkDuration(v int32) *CreateKnowledgeBaseRequestChunkConfig {
	s.ChunkDuration = &v
	return s
}

func (s *CreateKnowledgeBaseRequestChunkConfig) SetChunkOverlap(v int32) *CreateKnowledgeBaseRequestChunkConfig {
	s.ChunkOverlap = &v
	return s
}

func (s *CreateKnowledgeBaseRequestChunkConfig) SetChunkSize(v int32) *CreateKnowledgeBaseRequestChunkConfig {
	s.ChunkSize = &v
	return s
}

func (s *CreateKnowledgeBaseRequestChunkConfig) SetChunkStrategy(v string) *CreateKnowledgeBaseRequestChunkConfig {
	s.ChunkStrategy = &v
	return s
}

func (s *CreateKnowledgeBaseRequestChunkConfig) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseRequestDataSources struct {
	// 统一资源识别码
	//
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou-internal.aliyuncs.com/langstudio/source/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateKnowledgeBaseRequestDataSources) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequestDataSources) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequestDataSources) GetUri() *string {
	return s.Uri
}

func (s *CreateKnowledgeBaseRequestDataSources) SetUri(v string) *CreateKnowledgeBaseRequestDataSources {
	s.Uri = &v
	return s
}

func (s *CreateKnowledgeBaseRequestDataSources) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseRequestEmbeddingConfig struct {
	// Embedding连接ID
	//
	// This parameter is required.
	//
	// example:
	//
	// conn-r3o7******38bh
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// 模型
	//
	// This parameter is required.
	//
	// example:
	//
	// text-embedding-v4
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s CreateKnowledgeBaseRequestEmbeddingConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequestEmbeddingConfig) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequestEmbeddingConfig) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *CreateKnowledgeBaseRequestEmbeddingConfig) GetModel() *string {
	return s.Model
}

func (s *CreateKnowledgeBaseRequestEmbeddingConfig) SetConnectionId(v string) *CreateKnowledgeBaseRequestEmbeddingConfig {
	s.ConnectionId = &v
	return s
}

func (s *CreateKnowledgeBaseRequestEmbeddingConfig) SetModel(v string) *CreateKnowledgeBaseRequestEmbeddingConfig {
	s.Model = &v
	return s
}

func (s *CreateKnowledgeBaseRequestEmbeddingConfig) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseRequestIndexColumnConfig struct {
	// 所有列名
	ColumnDefinitions []*CreateKnowledgeBaseRequestIndexColumnConfigColumnDefinitions `json:"ColumnDefinitions,omitempty" xml:"ColumnDefinitions,omitempty" type:"Repeated"`
	// 内容检索列
	ContentColumns []*CreateKnowledgeBaseRequestIndexColumnConfigContentColumns `json:"ContentColumns,omitempty" xml:"ContentColumns,omitempty" type:"Repeated"`
	// Embedding列
	EmbeddingColumns []*CreateKnowledgeBaseRequestIndexColumnConfigEmbeddingColumns `json:"EmbeddingColumns,omitempty" xml:"EmbeddingColumns,omitempty" type:"Repeated"`
}

func (s CreateKnowledgeBaseRequestIndexColumnConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequestIndexColumnConfig) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfig) GetColumnDefinitions() []*CreateKnowledgeBaseRequestIndexColumnConfigColumnDefinitions {
	return s.ColumnDefinitions
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfig) GetContentColumns() []*CreateKnowledgeBaseRequestIndexColumnConfigContentColumns {
	return s.ContentColumns
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfig) GetEmbeddingColumns() []*CreateKnowledgeBaseRequestIndexColumnConfigEmbeddingColumns {
	return s.EmbeddingColumns
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfig) SetColumnDefinitions(v []*CreateKnowledgeBaseRequestIndexColumnConfigColumnDefinitions) *CreateKnowledgeBaseRequestIndexColumnConfig {
	s.ColumnDefinitions = v
	return s
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfig) SetContentColumns(v []*CreateKnowledgeBaseRequestIndexColumnConfigContentColumns) *CreateKnowledgeBaseRequestIndexColumnConfig {
	s.ContentColumns = v
	return s
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfig) SetEmbeddingColumns(v []*CreateKnowledgeBaseRequestIndexColumnConfigEmbeddingColumns) *CreateKnowledgeBaseRequestIndexColumnConfig {
	s.EmbeddingColumns = v
	return s
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfig) Validate() error {
	if s.ColumnDefinitions != nil {
		for _, item := range s.ColumnDefinitions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ContentColumns != nil {
		for _, item := range s.ContentColumns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EmbeddingColumns != nil {
		for _, item := range s.EmbeddingColumns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateKnowledgeBaseRequestIndexColumnConfigColumnDefinitions struct {
	// 列Key
	//
	// example:
	//
	// column1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s CreateKnowledgeBaseRequestIndexColumnConfigColumnDefinitions) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequestIndexColumnConfigColumnDefinitions) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfigColumnDefinitions) GetKey() *string {
	return s.Key
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfigColumnDefinitions) SetKey(v string) *CreateKnowledgeBaseRequestIndexColumnConfigColumnDefinitions {
	s.Key = &v
	return s
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfigColumnDefinitions) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseRequestIndexColumnConfigContentColumns struct {
	// 列Key
	//
	// example:
	//
	// column1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s CreateKnowledgeBaseRequestIndexColumnConfigContentColumns) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequestIndexColumnConfigContentColumns) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfigContentColumns) GetKey() *string {
	return s.Key
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfigContentColumns) SetKey(v string) *CreateKnowledgeBaseRequestIndexColumnConfigContentColumns {
	s.Key = &v
	return s
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfigContentColumns) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseRequestIndexColumnConfigEmbeddingColumns struct {
	// 列Key
	//
	// example:
	//
	// column1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s CreateKnowledgeBaseRequestIndexColumnConfigEmbeddingColumns) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequestIndexColumnConfigEmbeddingColumns) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfigEmbeddingColumns) GetKey() *string {
	return s.Key
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfigEmbeddingColumns) SetKey(v string) *CreateKnowledgeBaseRequestIndexColumnConfigEmbeddingColumns {
	s.Key = &v
	return s
}

func (s *CreateKnowledgeBaseRequestIndexColumnConfigEmbeddingColumns) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseRequestMetaDataConfig struct {
	// 自定义元数据
	CustomMetaData []*CreateKnowledgeBaseRequestMetaDataConfigCustomMetaData `json:"CustomMetaData,omitempty" xml:"CustomMetaData,omitempty" type:"Repeated"`
}

func (s CreateKnowledgeBaseRequestMetaDataConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequestMetaDataConfig) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequestMetaDataConfig) GetCustomMetaData() []*CreateKnowledgeBaseRequestMetaDataConfigCustomMetaData {
	return s.CustomMetaData
}

func (s *CreateKnowledgeBaseRequestMetaDataConfig) SetCustomMetaData(v []*CreateKnowledgeBaseRequestMetaDataConfigCustomMetaData) *CreateKnowledgeBaseRequestMetaDataConfig {
	s.CustomMetaData = v
	return s
}

func (s *CreateKnowledgeBaseRequestMetaDataConfig) Validate() error {
	if s.CustomMetaData != nil {
		for _, item := range s.CustomMetaData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateKnowledgeBaseRequestMetaDataConfigCustomMetaData struct {
	// 元数据Key
	//
	// example:
	//
	// author
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 元数据Value类型
	//
	// example:
	//
	// String
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s CreateKnowledgeBaseRequestMetaDataConfigCustomMetaData) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequestMetaDataConfigCustomMetaData) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequestMetaDataConfigCustomMetaData) GetKey() *string {
	return s.Key
}

func (s *CreateKnowledgeBaseRequestMetaDataConfigCustomMetaData) GetValueType() *string {
	return s.ValueType
}

func (s *CreateKnowledgeBaseRequestMetaDataConfigCustomMetaData) SetKey(v string) *CreateKnowledgeBaseRequestMetaDataConfigCustomMetaData {
	s.Key = &v
	return s
}

func (s *CreateKnowledgeBaseRequestMetaDataConfigCustomMetaData) SetValueType(v string) *CreateKnowledgeBaseRequestMetaDataConfigCustomMetaData {
	s.ValueType = &v
	return s
}

func (s *CreateKnowledgeBaseRequestMetaDataConfigCustomMetaData) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseRequestVectorDBConfig struct {
	// Collectioin名称
	//
	// example:
	//
	// my_collection
	CollectionName *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
	// Embedding连接ID
	//
	// example:
	//
	// conn-7y5y******jja7
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// VectorDB类型
	//
	// This parameter is required.
	//
	// example:
	//
	// Milvus
	VectorDBType *string `json:"VectorDBType,omitempty" xml:"VectorDBType,omitempty"`
}

func (s CreateKnowledgeBaseRequestVectorDBConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequestVectorDBConfig) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequestVectorDBConfig) GetCollectionName() *string {
	return s.CollectionName
}

func (s *CreateKnowledgeBaseRequestVectorDBConfig) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *CreateKnowledgeBaseRequestVectorDBConfig) GetVectorDBType() *string {
	return s.VectorDBType
}

func (s *CreateKnowledgeBaseRequestVectorDBConfig) SetCollectionName(v string) *CreateKnowledgeBaseRequestVectorDBConfig {
	s.CollectionName = &v
	return s
}

func (s *CreateKnowledgeBaseRequestVectorDBConfig) SetConnectionId(v string) *CreateKnowledgeBaseRequestVectorDBConfig {
	s.ConnectionId = &v
	return s
}

func (s *CreateKnowledgeBaseRequestVectorDBConfig) SetVectorDBType(v string) *CreateKnowledgeBaseRequestVectorDBConfig {
	s.VectorDBType = &v
	return s
}

func (s *CreateKnowledgeBaseRequestVectorDBConfig) Validate() error {
	return dara.Validate(s)
}
