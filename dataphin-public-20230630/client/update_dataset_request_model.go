// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDatasetRequest
	GetOpTenantId() *int64
	SetProjectId(v string) *UpdateDatasetRequest
	GetProjectId() *string
	SetUpdateCommand(v *UpdateDatasetRequestUpdateCommand) *UpdateDatasetRequest
	GetUpdateCommand() *UpdateDatasetRequestUpdateCommand
}

type UpdateDatasetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7273382541481536
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateDatasetRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDatasetRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateDatasetRequest) GetUpdateCommand() *UpdateDatasetRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateDatasetRequest) SetOpTenantId(v int64) *UpdateDatasetRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDatasetRequest) SetProjectId(v string) *UpdateDatasetRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDatasetRequest) SetUpdateCommand(v *UpdateDatasetRequestUpdateCommand) *UpdateDatasetRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateDatasetRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommand struct {
	// example:
	//
	// GENERAL
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// 78201
	DataCellId *string `json:"DataCellId,omitempty" xml:"DataCellId,omitempty"`
	// example:
	//
	// 测试数据集
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7261110566632832
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7280832407583104
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// POSTGRESQL
	MetadataStorageType *string `json:"MetadataStorageType,omitempty" xml:"MetadataStorageType,omitempty"`
	// example:
	//
	// audio_dataset
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 300001391
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// OFFLINE
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// example:
	//
	// HYBRID
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// V1
	Version       *string                                         `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionConfig *UpdateDatasetRequestUpdateCommandVersionConfig `json:"VersionConfig,omitempty" xml:"VersionConfig,omitempty" type:"Struct"`
}

func (s UpdateDatasetRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommand) GetContentType() *string {
	return s.ContentType
}

func (s *UpdateDatasetRequestUpdateCommand) GetDataCellId() *string {
	return s.DataCellId
}

func (s *UpdateDatasetRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateDatasetRequestUpdateCommand) GetFileId() *string {
	return s.FileId
}

func (s *UpdateDatasetRequestUpdateCommand) GetId() *int64 {
	return s.Id
}

func (s *UpdateDatasetRequestUpdateCommand) GetMetadataStorageType() *string {
	return s.MetadataStorageType
}

func (s *UpdateDatasetRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateDatasetRequestUpdateCommand) GetOwner() *string {
	return s.Owner
}

func (s *UpdateDatasetRequestUpdateCommand) GetScenario() *string {
	return s.Scenario
}

func (s *UpdateDatasetRequestUpdateCommand) GetStorageType() *string {
	return s.StorageType
}

func (s *UpdateDatasetRequestUpdateCommand) GetType() *string {
	return s.Type
}

func (s *UpdateDatasetRequestUpdateCommand) GetVersion() *string {
	return s.Version
}

func (s *UpdateDatasetRequestUpdateCommand) GetVersionConfig() *UpdateDatasetRequestUpdateCommandVersionConfig {
	return s.VersionConfig
}

func (s *UpdateDatasetRequestUpdateCommand) SetContentType(v string) *UpdateDatasetRequestUpdateCommand {
	s.ContentType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetDataCellId(v string) *UpdateDatasetRequestUpdateCommand {
	s.DataCellId = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetDescription(v string) *UpdateDatasetRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetFileId(v string) *UpdateDatasetRequestUpdateCommand {
	s.FileId = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetId(v int64) *UpdateDatasetRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetMetadataStorageType(v string) *UpdateDatasetRequestUpdateCommand {
	s.MetadataStorageType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetName(v string) *UpdateDatasetRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetOwner(v string) *UpdateDatasetRequestUpdateCommand {
	s.Owner = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetScenario(v string) *UpdateDatasetRequestUpdateCommand {
	s.Scenario = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetStorageType(v string) *UpdateDatasetRequestUpdateCommand {
	s.StorageType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetType(v string) *UpdateDatasetRequestUpdateCommand {
	s.Type = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetVersion(v string) *UpdateDatasetRequestUpdateCommand {
	s.Version = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) SetVersionConfig(v *UpdateDatasetRequestUpdateCommandVersionConfig) *UpdateDatasetRequestUpdateCommand {
	s.VersionConfig = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommand) Validate() error {
	if s.VersionConfig != nil {
		if err := s.VersionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandVersionConfig struct {
	FileStorageConfig       *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig       `json:"FileStorageConfig,omitempty" xml:"FileStorageConfig,omitempty" type:"Struct"`
	MetadataStorageConfig   *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig   `json:"MetadataStorageConfig,omitempty" xml:"MetadataStorageConfig,omitempty" type:"Struct"`
	RealtimeMetaTableConfig *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig `json:"RealtimeMetaTableConfig,omitempty" xml:"RealtimeMetaTableConfig,omitempty" type:"Struct"`
	// example:
	//
	// 测试数据集版本
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) GetFileStorageConfig() *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig {
	return s.FileStorageConfig
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) GetMetadataStorageConfig() *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	return s.MetadataStorageConfig
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) GetRealtimeMetaTableConfig() *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig {
	return s.RealtimeMetaTableConfig
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) SetFileStorageConfig(v *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) *UpdateDatasetRequestUpdateCommandVersionConfig {
	s.FileStorageConfig = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) SetMetadataStorageConfig(v *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) *UpdateDatasetRequestUpdateCommandVersionConfig {
	s.MetadataStorageConfig = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) SetRealtimeMetaTableConfig(v *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) *UpdateDatasetRequestUpdateCommandVersionConfig {
	s.RealtimeMetaTableConfig = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) SetVersionDescription(v string) *UpdateDatasetRequestUpdateCommandVersionConfig {
	s.VersionDescription = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfig) Validate() error {
	if s.FileStorageConfig != nil {
		if err := s.FileStorageConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MetadataStorageConfig != nil {
		if err := s.MetadataStorageConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RealtimeMetaTableConfig != nil {
		if err := s.RealtimeMetaTableConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// 7445343860022804608
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// 测试数据源
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// example:
	//
	// HTML正文提取/test423/
	DevPath *string `json:"DevPath,omitempty" xml:"DevPath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /var/run/openresty/cache/corp
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HTML正文提取/test423/
	ProdPath *string `json:"ProdPath,omitempty" xml:"ProdPath,omitempty"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) GetDevPath() *string {
	return s.DevPath
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) GetMountPath() *string {
	return s.MountPath
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) GetProdPath() *string {
	return s.ProdPath
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) SetDataSourceId(v string) *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig {
	s.DataSourceId = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) SetDataSourceName(v string) *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig {
	s.DataSourceName = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) SetDevPath(v string) *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig {
	s.DevPath = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) SetMountPath(v string) *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig {
	s.MountPath = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) SetProdPath(v string) *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig {
	s.ProdPath = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigFileStorageConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// 7429133693081710272
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// 测试数据源
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// example:
	//
	// HTML正文提取/test423/
	DevSchema *string `json:"DevSchema,omitempty" xml:"DevSchema,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CREATE
	MetadataStorageMode *string `json:"MetadataStorageMode,omitempty" xml:"MetadataStorageMode,omitempty"`
	// example:
	//
	// MILVUS
	MetadataStorageType *string `json:"MetadataStorageType,omitempty" xml:"MetadataStorageType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HTML正文提取/test423/
	ProdSchema *string `json:"ProdSchema,omitempty" xml:"ProdSchema,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// s_crm_all_plt_jala_shop
	TableName   *string                                                                         `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableSchema *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetDevSchema() *string {
	return s.DevSchema
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetMetadataStorageMode() *string {
	return s.MetadataStorageMode
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetMetadataStorageType() *string {
	return s.MetadataStorageType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetProdSchema() *string {
	return s.ProdSchema
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetTableName() *string {
	return s.TableName
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) GetTableSchema() *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema {
	return s.TableSchema
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetDataSourceId(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.DataSourceId = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetDataSourceName(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.DataSourceName = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetDevSchema(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.DevSchema = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetMetadataStorageMode(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.MetadataStorageMode = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetMetadataStorageType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.MetadataStorageType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetProdSchema(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.ProdSchema = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetTableName(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.TableName = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) SetTableSchema(v *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig {
	s.TableSchema = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfig) Validate() error {
	if s.TableSchema != nil {
		if err := s.TableSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema struct {
	Columns []*UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema) GetColumns() []*UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	return s.Columns
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema) SetColumns(v []*UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema {
	s.Columns = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchema) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns struct {
	// example:
	//
	// primary key
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// INT64
	ElementType *string `json:"ElementType,omitempty" xml:"ElementType,omitempty"`
	// example:
	//
	// 250
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	Pk *bool `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// int8
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// false
	Url               *bool                                                                                                   `json:"Url,omitempty" xml:"Url,omitempty"`
	VectorIndexConfig *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig `json:"VectorIndexConfig,omitempty" xml:"VectorIndexConfig,omitempty" type:"Struct"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetComment() *string {
	return s.Comment
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetElementType() *string {
	return s.ElementType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetName() *string {
	return s.Name
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetPk() *bool {
	return s.Pk
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetType() *string {
	return s.Type
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetUrl() *bool {
	return s.Url
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) GetVectorIndexConfig() *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	return s.VectorIndexConfig
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetComment(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Comment = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetElementType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.ElementType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetMaxCapacity(v int32) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.MaxCapacity = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetName(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Name = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetPk(v bool) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Pk = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Type = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetUrl(v bool) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.Url = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) SetVectorIndexConfig(v *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns {
	s.VectorIndexConfig = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumns) Validate() error {
	if s.VectorIndexConfig != nil {
		if err := s.VectorIndexConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// instance:mongodb
	Dimension *int64 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MultiModal-Embedding
	EmbeddingModel *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	// example:
	//
	// {M:30, efConstruction:360}
	IndexParams map[string]interface{} `json:"IndexParams,omitempty" xml:"IndexParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AUTOINDEX
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// COSINE
	SimilarityType *string `json:"SimilarityType,omitempty" xml:"SimilarityType,omitempty"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetDimension() *int64 {
	return s.Dimension
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetIndexParams() map[string]interface{} {
	return s.IndexParams
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetIndexType() *string {
	return s.IndexType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) GetSimilarityType() *string {
	return s.SimilarityType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetDimension(v int64) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.Dimension = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetEmbeddingModel(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.EmbeddingModel = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetIndexParams(v map[string]interface{}) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexParams = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetIndexType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) SetSimilarityType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig {
	s.SimilarityType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigMetadataStorageConfigTableSchemaColumnsVectorIndexConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// KAFKA
	DatasourceType *string `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试元表
	MetaTableName *string `json:"MetaTableName,omitempty" xml:"MetaTableName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7255013756724992
	ProjectId   *int64                                                                            `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	TableSchema *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) GetDatasourceType() *string {
	return s.DatasourceType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) GetMetaTableName() *string {
	return s.MetaTableName
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) GetTableSchema() *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema {
	return s.TableSchema
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) SetDatasourceType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig {
	s.DatasourceType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) SetMetaTableName(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig {
	s.MetaTableName = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) SetProjectId(v int64) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig {
	s.ProjectId = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) SetTableSchema(v *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig {
	s.TableSchema = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfig) Validate() error {
	if s.TableSchema != nil {
		if err := s.TableSchema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema struct {
	Columns []*UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema) GetColumns() []*UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	return s.Columns
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema) SetColumns(v []*UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema {
	s.Columns = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchema) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns struct {
	// example:
	//
	// happen time
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// INT64
	ElementType *string `json:"ElementType,omitempty" xml:"ElementType,omitempty"`
	// example:
	//
	// 10
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// happen_time
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	Pk *bool `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// date
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// false
	Url               *bool                                                                                                     `json:"Url,omitempty" xml:"Url,omitempty"`
	VectorIndexConfig *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig `json:"VectorIndexConfig,omitempty" xml:"VectorIndexConfig,omitempty" type:"Struct"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetComment() *string {
	return s.Comment
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetElementType() *string {
	return s.ElementType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetName() *string {
	return s.Name
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetPk() *bool {
	return s.Pk
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetType() *string {
	return s.Type
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetUrl() *bool {
	return s.Url
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) GetVectorIndexConfig() *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	return s.VectorIndexConfig
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetComment(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Comment = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetElementType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.ElementType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetMaxCapacity(v int32) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.MaxCapacity = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetName(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Name = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetPk(v bool) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Pk = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Type = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetUrl(v bool) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.Url = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) SetVectorIndexConfig(v *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns {
	s.VectorIndexConfig = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumns) Validate() error {
	if s.VectorIndexConfig != nil {
		if err := s.VectorIndexConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// instance
	Dimension *int64 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// multimodal-embedding-v1
	EmbeddingModel *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	// example:
	//
	// {M:30, efConstruction:360}
	IndexParams map[string]interface{} `json:"IndexParams,omitempty" xml:"IndexParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AUTOINDEX
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// COSINE
	SimilarityType *string `json:"SimilarityType,omitempty" xml:"SimilarityType,omitempty"`
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetDimension() *int64 {
	return s.Dimension
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetIndexParams() map[string]interface{} {
	return s.IndexParams
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetIndexType() *string {
	return s.IndexType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) GetSimilarityType() *string {
	return s.SimilarityType
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetDimension(v int64) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.Dimension = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetEmbeddingModel(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.EmbeddingModel = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetIndexParams(v map[string]interface{}) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexParams = v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetIndexType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.IndexType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) SetSimilarityType(v string) *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig {
	s.SimilarityType = &v
	return s
}

func (s *UpdateDatasetRequestUpdateCommandVersionConfigRealtimeMetaTableConfigTableSchemaColumnsVectorIndexConfig) Validate() error {
	return dara.Validate(s)
}
