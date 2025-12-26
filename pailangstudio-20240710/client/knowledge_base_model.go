// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeBase interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *KnowledgeBase
	GetAccessibility() *string
	SetAutoUpdateConfig(v *KnowledgeBaseAutoUpdateConfig) *KnowledgeBase
	GetAutoUpdateConfig() *KnowledgeBaseAutoUpdateConfig
	SetChunkConfig(v *KnowledgeBaseChunkConfig) *KnowledgeBase
	GetChunkConfig() *KnowledgeBaseChunkConfig
	SetCreator(v string) *KnowledgeBase
	GetCreator() *string
	SetDataSources(v []*KnowledgeBaseDataSources) *KnowledgeBase
	GetDataSources() []*KnowledgeBaseDataSources
	SetDatasetId(v string) *KnowledgeBase
	GetDatasetId() *string
	SetDescription(v string) *KnowledgeBase
	GetDescription() *string
	SetEmbeddingConfig(v *KnowledgeBaseEmbeddingConfig) *KnowledgeBase
	GetEmbeddingConfig() *KnowledgeBaseEmbeddingConfig
	SetGmtCreateTime(v string) *KnowledgeBase
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *KnowledgeBase
	GetGmtModifiedTime() *string
	SetIndexColumnConfig(v *KnowledgeBaseIndexColumnConfig) *KnowledgeBase
	GetIndexColumnConfig() *KnowledgeBaseIndexColumnConfig
	SetKnowledgeBaseId(v string) *KnowledgeBase
	GetKnowledgeBaseId() *string
	SetKnowledgeBaseType(v string) *KnowledgeBase
	GetKnowledgeBaseType() *string
	SetMetaDataConfig(v *KnowledgeBaseMetaDataConfig) *KnowledgeBase
	GetMetaDataConfig() *KnowledgeBaseMetaDataConfig
	SetName(v string) *KnowledgeBase
	GetName() *string
	SetOutputDir(v string) *KnowledgeBase
	GetOutputDir() *string
	SetRuntimeId(v string) *KnowledgeBase
	GetRuntimeId() *string
	SetVectorDBConfig(v *KnowledgeBaseVectorDBConfig) *KnowledgeBase
	GetVectorDBConfig() *KnowledgeBaseVectorDBConfig
	SetVersionName(v string) *KnowledgeBase
	GetVersionName() *string
	SetWorkspaceId(v string) *KnowledgeBase
	GetWorkspaceId() *string
}

type KnowledgeBase struct {
	Accessibility     *string                         `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	AutoUpdateConfig  *KnowledgeBaseAutoUpdateConfig  `json:"AutoUpdateConfig,omitempty" xml:"AutoUpdateConfig,omitempty" type:"Struct"`
	ChunkConfig       *KnowledgeBaseChunkConfig       `json:"ChunkConfig,omitempty" xml:"ChunkConfig,omitempty" type:"Struct"`
	Creator           *string                         `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DataSources       []*KnowledgeBaseDataSources     `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	DatasetId         *string                         `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Description       *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	EmbeddingConfig   *KnowledgeBaseEmbeddingConfig   `json:"EmbeddingConfig,omitempty" xml:"EmbeddingConfig,omitempty" type:"Struct"`
	GmtCreateTime     *string                         `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime   *string                         `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	IndexColumnConfig *KnowledgeBaseIndexColumnConfig `json:"IndexColumnConfig,omitempty" xml:"IndexColumnConfig,omitempty" type:"Struct"`
	KnowledgeBaseId   *string                         `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	KnowledgeBaseType *string                         `json:"KnowledgeBaseType,omitempty" xml:"KnowledgeBaseType,omitempty"`
	MetaDataConfig    *KnowledgeBaseMetaDataConfig    `json:"MetaDataConfig,omitempty" xml:"MetaDataConfig,omitempty" type:"Struct"`
	Name              *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	OutputDir         *string                         `json:"OutputDir,omitempty" xml:"OutputDir,omitempty"`
	RuntimeId         *string                         `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	VectorDBConfig    *KnowledgeBaseVectorDBConfig    `json:"VectorDBConfig,omitempty" xml:"VectorDBConfig,omitempty" type:"Struct"`
	VersionName       *string                         `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	WorkspaceId       *string                         `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s KnowledgeBase) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBase) GoString() string {
	return s.String()
}

func (s *KnowledgeBase) GetAccessibility() *string {
	return s.Accessibility
}

func (s *KnowledgeBase) GetAutoUpdateConfig() *KnowledgeBaseAutoUpdateConfig {
	return s.AutoUpdateConfig
}

func (s *KnowledgeBase) GetChunkConfig() *KnowledgeBaseChunkConfig {
	return s.ChunkConfig
}

func (s *KnowledgeBase) GetCreator() *string {
	return s.Creator
}

func (s *KnowledgeBase) GetDataSources() []*KnowledgeBaseDataSources {
	return s.DataSources
}

func (s *KnowledgeBase) GetDatasetId() *string {
	return s.DatasetId
}

func (s *KnowledgeBase) GetDescription() *string {
	return s.Description
}

func (s *KnowledgeBase) GetEmbeddingConfig() *KnowledgeBaseEmbeddingConfig {
	return s.EmbeddingConfig
}

func (s *KnowledgeBase) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *KnowledgeBase) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *KnowledgeBase) GetIndexColumnConfig() *KnowledgeBaseIndexColumnConfig {
	return s.IndexColumnConfig
}

func (s *KnowledgeBase) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *KnowledgeBase) GetKnowledgeBaseType() *string {
	return s.KnowledgeBaseType
}

func (s *KnowledgeBase) GetMetaDataConfig() *KnowledgeBaseMetaDataConfig {
	return s.MetaDataConfig
}

func (s *KnowledgeBase) GetName() *string {
	return s.Name
}

func (s *KnowledgeBase) GetOutputDir() *string {
	return s.OutputDir
}

func (s *KnowledgeBase) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *KnowledgeBase) GetVectorDBConfig() *KnowledgeBaseVectorDBConfig {
	return s.VectorDBConfig
}

func (s *KnowledgeBase) GetVersionName() *string {
	return s.VersionName
}

func (s *KnowledgeBase) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *KnowledgeBase) SetAccessibility(v string) *KnowledgeBase {
	s.Accessibility = &v
	return s
}

func (s *KnowledgeBase) SetAutoUpdateConfig(v *KnowledgeBaseAutoUpdateConfig) *KnowledgeBase {
	s.AutoUpdateConfig = v
	return s
}

func (s *KnowledgeBase) SetChunkConfig(v *KnowledgeBaseChunkConfig) *KnowledgeBase {
	s.ChunkConfig = v
	return s
}

func (s *KnowledgeBase) SetCreator(v string) *KnowledgeBase {
	s.Creator = &v
	return s
}

func (s *KnowledgeBase) SetDataSources(v []*KnowledgeBaseDataSources) *KnowledgeBase {
	s.DataSources = v
	return s
}

func (s *KnowledgeBase) SetDatasetId(v string) *KnowledgeBase {
	s.DatasetId = &v
	return s
}

func (s *KnowledgeBase) SetDescription(v string) *KnowledgeBase {
	s.Description = &v
	return s
}

func (s *KnowledgeBase) SetEmbeddingConfig(v *KnowledgeBaseEmbeddingConfig) *KnowledgeBase {
	s.EmbeddingConfig = v
	return s
}

func (s *KnowledgeBase) SetGmtCreateTime(v string) *KnowledgeBase {
	s.GmtCreateTime = &v
	return s
}

func (s *KnowledgeBase) SetGmtModifiedTime(v string) *KnowledgeBase {
	s.GmtModifiedTime = &v
	return s
}

func (s *KnowledgeBase) SetIndexColumnConfig(v *KnowledgeBaseIndexColumnConfig) *KnowledgeBase {
	s.IndexColumnConfig = v
	return s
}

func (s *KnowledgeBase) SetKnowledgeBaseId(v string) *KnowledgeBase {
	s.KnowledgeBaseId = &v
	return s
}

func (s *KnowledgeBase) SetKnowledgeBaseType(v string) *KnowledgeBase {
	s.KnowledgeBaseType = &v
	return s
}

func (s *KnowledgeBase) SetMetaDataConfig(v *KnowledgeBaseMetaDataConfig) *KnowledgeBase {
	s.MetaDataConfig = v
	return s
}

func (s *KnowledgeBase) SetName(v string) *KnowledgeBase {
	s.Name = &v
	return s
}

func (s *KnowledgeBase) SetOutputDir(v string) *KnowledgeBase {
	s.OutputDir = &v
	return s
}

func (s *KnowledgeBase) SetRuntimeId(v string) *KnowledgeBase {
	s.RuntimeId = &v
	return s
}

func (s *KnowledgeBase) SetVectorDBConfig(v *KnowledgeBaseVectorDBConfig) *KnowledgeBase {
	s.VectorDBConfig = v
	return s
}

func (s *KnowledgeBase) SetVersionName(v string) *KnowledgeBase {
	s.VersionName = &v
	return s
}

func (s *KnowledgeBase) SetWorkspaceId(v string) *KnowledgeBase {
	s.WorkspaceId = &v
	return s
}

func (s *KnowledgeBase) Validate() error {
	if s.AutoUpdateConfig != nil {
		if err := s.AutoUpdateConfig.Validate(); err != nil {
			return err
		}
	}
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

type KnowledgeBaseAutoUpdateConfig struct {
	// 运行资源配置
	EcsSpecs []*KnowledgeBaseAutoUpdateConfigEcsSpecs `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	// Embedding配置
	EmbeddingConfig *KnowledgeBaseAutoUpdateConfigEmbeddingConfig `json:"EmbeddingConfig,omitempty" xml:"EmbeddingConfig,omitempty" type:"Struct"`
	// 任务最大运行时间
	MaxRunningTimeInSeconds *int32 `json:"MaxRunningTimeInSeconds,omitempty" xml:"MaxRunningTimeInSeconds,omitempty"`
	// 资源组ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 知识库自动更新状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 用户VPC配置
	UserVpc *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
}

func (s KnowledgeBaseAutoUpdateConfig) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseAutoUpdateConfig) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseAutoUpdateConfig) GetEcsSpecs() []*KnowledgeBaseAutoUpdateConfigEcsSpecs {
	return s.EcsSpecs
}

func (s *KnowledgeBaseAutoUpdateConfig) GetEmbeddingConfig() *KnowledgeBaseAutoUpdateConfigEmbeddingConfig {
	return s.EmbeddingConfig
}

func (s *KnowledgeBaseAutoUpdateConfig) GetMaxRunningTimeInSeconds() *int32 {
	return s.MaxRunningTimeInSeconds
}

func (s *KnowledgeBaseAutoUpdateConfig) GetResourceId() *string {
	return s.ResourceId
}

func (s *KnowledgeBaseAutoUpdateConfig) GetStatus() *string {
	return s.Status
}

func (s *KnowledgeBaseAutoUpdateConfig) GetUserVpc() *UserVpc {
	return s.UserVpc
}

func (s *KnowledgeBaseAutoUpdateConfig) SetEcsSpecs(v []*KnowledgeBaseAutoUpdateConfigEcsSpecs) *KnowledgeBaseAutoUpdateConfig {
	s.EcsSpecs = v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfig) SetEmbeddingConfig(v *KnowledgeBaseAutoUpdateConfigEmbeddingConfig) *KnowledgeBaseAutoUpdateConfig {
	s.EmbeddingConfig = v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfig) SetMaxRunningTimeInSeconds(v int32) *KnowledgeBaseAutoUpdateConfig {
	s.MaxRunningTimeInSeconds = &v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfig) SetResourceId(v string) *KnowledgeBaseAutoUpdateConfig {
	s.ResourceId = &v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfig) SetStatus(v string) *KnowledgeBaseAutoUpdateConfig {
	s.Status = &v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfig) SetUserVpc(v *UserVpc) *KnowledgeBaseAutoUpdateConfig {
	s.UserVpc = v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfig) Validate() error {
	if s.EcsSpecs != nil {
		for _, item := range s.EcsSpecs {
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
	if s.UserVpc != nil {
		if err := s.UserVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KnowledgeBaseAutoUpdateConfigEcsSpecs struct {
	// CPU核数
	CPU *int32 `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// 驱动版本
	Driver *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	// GPU卡数
	GPU *int32 `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// GPU类型
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// 机型名称
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 内存大小
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// 副本数量
	PodCount *int32 `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
	// 共享内存容量
	SharedMemory *int32 `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
	// 节点类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s KnowledgeBaseAutoUpdateConfigEcsSpecs) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseAutoUpdateConfigEcsSpecs) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) GetCPU() *int32 {
	return s.CPU
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) GetDriver() *string {
	return s.Driver
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) GetGPU() *int32 {
	return s.GPU
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) GetGPUType() *string {
	return s.GPUType
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) GetInstanceType() *string {
	return s.InstanceType
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) GetMemory() *int32 {
	return s.Memory
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) GetPodCount() *int32 {
	return s.PodCount
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) GetSharedMemory() *int32 {
	return s.SharedMemory
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) GetType() *string {
	return s.Type
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) SetCPU(v int32) *KnowledgeBaseAutoUpdateConfigEcsSpecs {
	s.CPU = &v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) SetDriver(v string) *KnowledgeBaseAutoUpdateConfigEcsSpecs {
	s.Driver = &v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) SetGPU(v int32) *KnowledgeBaseAutoUpdateConfigEcsSpecs {
	s.GPU = &v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) SetGPUType(v string) *KnowledgeBaseAutoUpdateConfigEcsSpecs {
	s.GPUType = &v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) SetInstanceType(v string) *KnowledgeBaseAutoUpdateConfigEcsSpecs {
	s.InstanceType = &v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) SetMemory(v int32) *KnowledgeBaseAutoUpdateConfigEcsSpecs {
	s.Memory = &v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) SetPodCount(v int32) *KnowledgeBaseAutoUpdateConfigEcsSpecs {
	s.PodCount = &v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) SetSharedMemory(v int32) *KnowledgeBaseAutoUpdateConfigEcsSpecs {
	s.SharedMemory = &v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) SetType(v string) *KnowledgeBaseAutoUpdateConfigEcsSpecs {
	s.Type = &v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfigEcsSpecs) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseAutoUpdateConfigEmbeddingConfig struct {
	// Embedding分批大小
	BatchSize *int32 `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// Embedding并发数
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
}

func (s KnowledgeBaseAutoUpdateConfigEmbeddingConfig) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseAutoUpdateConfigEmbeddingConfig) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseAutoUpdateConfigEmbeddingConfig) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *KnowledgeBaseAutoUpdateConfigEmbeddingConfig) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *KnowledgeBaseAutoUpdateConfigEmbeddingConfig) SetBatchSize(v int32) *KnowledgeBaseAutoUpdateConfigEmbeddingConfig {
	s.BatchSize = &v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfigEmbeddingConfig) SetConcurrency(v int32) *KnowledgeBaseAutoUpdateConfigEmbeddingConfig {
	s.Concurrency = &v
	return s
}

func (s *KnowledgeBaseAutoUpdateConfigEmbeddingConfig) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseChunkConfig struct {
	// 分块时长
	ChunkDuration *int32 `json:"ChunkDuration,omitempty" xml:"ChunkDuration,omitempty"`
	// 分块重叠大小
	ChunkOverlap *int32 `json:"ChunkOverlap,omitempty" xml:"ChunkOverlap,omitempty"`
	// 分块大小
	ChunkSize *int32 `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// 分块策略
	ChunkStrategy *string `json:"ChunkStrategy,omitempty" xml:"ChunkStrategy,omitempty"`
}

func (s KnowledgeBaseChunkConfig) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseChunkConfig) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseChunkConfig) GetChunkDuration() *int32 {
	return s.ChunkDuration
}

func (s *KnowledgeBaseChunkConfig) GetChunkOverlap() *int32 {
	return s.ChunkOverlap
}

func (s *KnowledgeBaseChunkConfig) GetChunkSize() *int32 {
	return s.ChunkSize
}

func (s *KnowledgeBaseChunkConfig) GetChunkStrategy() *string {
	return s.ChunkStrategy
}

func (s *KnowledgeBaseChunkConfig) SetChunkDuration(v int32) *KnowledgeBaseChunkConfig {
	s.ChunkDuration = &v
	return s
}

func (s *KnowledgeBaseChunkConfig) SetChunkOverlap(v int32) *KnowledgeBaseChunkConfig {
	s.ChunkOverlap = &v
	return s
}

func (s *KnowledgeBaseChunkConfig) SetChunkSize(v int32) *KnowledgeBaseChunkConfig {
	s.ChunkSize = &v
	return s
}

func (s *KnowledgeBaseChunkConfig) SetChunkStrategy(v string) *KnowledgeBaseChunkConfig {
	s.ChunkStrategy = &v
	return s
}

func (s *KnowledgeBaseChunkConfig) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseDataSources struct {
	// 统一资源识别码
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s KnowledgeBaseDataSources) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseDataSources) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseDataSources) GetUri() *string {
	return s.Uri
}

func (s *KnowledgeBaseDataSources) SetUri(v string) *KnowledgeBaseDataSources {
	s.Uri = &v
	return s
}

func (s *KnowledgeBaseDataSources) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseEmbeddingConfig struct {
	// Embedding连接ID
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// Embedding连接名称
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// 模型
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s KnowledgeBaseEmbeddingConfig) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseEmbeddingConfig) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseEmbeddingConfig) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *KnowledgeBaseEmbeddingConfig) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *KnowledgeBaseEmbeddingConfig) GetModel() *string {
	return s.Model
}

func (s *KnowledgeBaseEmbeddingConfig) SetConnectionId(v string) *KnowledgeBaseEmbeddingConfig {
	s.ConnectionId = &v
	return s
}

func (s *KnowledgeBaseEmbeddingConfig) SetConnectionName(v string) *KnowledgeBaseEmbeddingConfig {
	s.ConnectionName = &v
	return s
}

func (s *KnowledgeBaseEmbeddingConfig) SetModel(v string) *KnowledgeBaseEmbeddingConfig {
	s.Model = &v
	return s
}

func (s *KnowledgeBaseEmbeddingConfig) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseIndexColumnConfig struct {
	// 所有列名
	ColumnDefinitions []*KnowledgeBaseIndexColumnConfigColumnDefinitions `json:"ColumnDefinitions,omitempty" xml:"ColumnDefinitions,omitempty" type:"Repeated"`
	// 内容检索列
	ContentColumns []*KnowledgeBaseIndexColumnConfigContentColumns `json:"ContentColumns,omitempty" xml:"ContentColumns,omitempty" type:"Repeated"`
	// Embedding列
	EmbeddingColumns []*KnowledgeBaseIndexColumnConfigEmbeddingColumns `json:"EmbeddingColumns,omitempty" xml:"EmbeddingColumns,omitempty" type:"Repeated"`
}

func (s KnowledgeBaseIndexColumnConfig) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseIndexColumnConfig) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseIndexColumnConfig) GetColumnDefinitions() []*KnowledgeBaseIndexColumnConfigColumnDefinitions {
	return s.ColumnDefinitions
}

func (s *KnowledgeBaseIndexColumnConfig) GetContentColumns() []*KnowledgeBaseIndexColumnConfigContentColumns {
	return s.ContentColumns
}

func (s *KnowledgeBaseIndexColumnConfig) GetEmbeddingColumns() []*KnowledgeBaseIndexColumnConfigEmbeddingColumns {
	return s.EmbeddingColumns
}

func (s *KnowledgeBaseIndexColumnConfig) SetColumnDefinitions(v []*KnowledgeBaseIndexColumnConfigColumnDefinitions) *KnowledgeBaseIndexColumnConfig {
	s.ColumnDefinitions = v
	return s
}

func (s *KnowledgeBaseIndexColumnConfig) SetContentColumns(v []*KnowledgeBaseIndexColumnConfigContentColumns) *KnowledgeBaseIndexColumnConfig {
	s.ContentColumns = v
	return s
}

func (s *KnowledgeBaseIndexColumnConfig) SetEmbeddingColumns(v []*KnowledgeBaseIndexColumnConfigEmbeddingColumns) *KnowledgeBaseIndexColumnConfig {
	s.EmbeddingColumns = v
	return s
}

func (s *KnowledgeBaseIndexColumnConfig) Validate() error {
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

type KnowledgeBaseIndexColumnConfigColumnDefinitions struct {
	// 列Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s KnowledgeBaseIndexColumnConfigColumnDefinitions) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseIndexColumnConfigColumnDefinitions) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseIndexColumnConfigColumnDefinitions) GetKey() *string {
	return s.Key
}

func (s *KnowledgeBaseIndexColumnConfigColumnDefinitions) SetKey(v string) *KnowledgeBaseIndexColumnConfigColumnDefinitions {
	s.Key = &v
	return s
}

func (s *KnowledgeBaseIndexColumnConfigColumnDefinitions) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseIndexColumnConfigContentColumns struct {
	// 列Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s KnowledgeBaseIndexColumnConfigContentColumns) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseIndexColumnConfigContentColumns) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseIndexColumnConfigContentColumns) GetKey() *string {
	return s.Key
}

func (s *KnowledgeBaseIndexColumnConfigContentColumns) SetKey(v string) *KnowledgeBaseIndexColumnConfigContentColumns {
	s.Key = &v
	return s
}

func (s *KnowledgeBaseIndexColumnConfigContentColumns) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseIndexColumnConfigEmbeddingColumns struct {
	// 列Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s KnowledgeBaseIndexColumnConfigEmbeddingColumns) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseIndexColumnConfigEmbeddingColumns) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseIndexColumnConfigEmbeddingColumns) GetKey() *string {
	return s.Key
}

func (s *KnowledgeBaseIndexColumnConfigEmbeddingColumns) SetKey(v string) *KnowledgeBaseIndexColumnConfigEmbeddingColumns {
	s.Key = &v
	return s
}

func (s *KnowledgeBaseIndexColumnConfigEmbeddingColumns) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseMetaDataConfig struct {
	// 自定义元数据
	CustomMetaData []*KnowledgeBaseMetaDataConfigCustomMetaData `json:"CustomMetaData,omitempty" xml:"CustomMetaData,omitempty" type:"Repeated"`
}

func (s KnowledgeBaseMetaDataConfig) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseMetaDataConfig) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseMetaDataConfig) GetCustomMetaData() []*KnowledgeBaseMetaDataConfigCustomMetaData {
	return s.CustomMetaData
}

func (s *KnowledgeBaseMetaDataConfig) SetCustomMetaData(v []*KnowledgeBaseMetaDataConfigCustomMetaData) *KnowledgeBaseMetaDataConfig {
	s.CustomMetaData = v
	return s
}

func (s *KnowledgeBaseMetaDataConfig) Validate() error {
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

type KnowledgeBaseMetaDataConfigCustomMetaData struct {
	// 元数据Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 引用次数
	ReferenceCount *int32 `json:"ReferenceCount,omitempty" xml:"ReferenceCount,omitempty"`
	// 值的个数
	ValueCount *int32 `json:"ValueCount,omitempty" xml:"ValueCount,omitempty"`
	// 元数据Value类型
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s KnowledgeBaseMetaDataConfigCustomMetaData) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseMetaDataConfigCustomMetaData) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseMetaDataConfigCustomMetaData) GetKey() *string {
	return s.Key
}

func (s *KnowledgeBaseMetaDataConfigCustomMetaData) GetReferenceCount() *int32 {
	return s.ReferenceCount
}

func (s *KnowledgeBaseMetaDataConfigCustomMetaData) GetValueCount() *int32 {
	return s.ValueCount
}

func (s *KnowledgeBaseMetaDataConfigCustomMetaData) GetValueType() *string {
	return s.ValueType
}

func (s *KnowledgeBaseMetaDataConfigCustomMetaData) SetKey(v string) *KnowledgeBaseMetaDataConfigCustomMetaData {
	s.Key = &v
	return s
}

func (s *KnowledgeBaseMetaDataConfigCustomMetaData) SetReferenceCount(v int32) *KnowledgeBaseMetaDataConfigCustomMetaData {
	s.ReferenceCount = &v
	return s
}

func (s *KnowledgeBaseMetaDataConfigCustomMetaData) SetValueCount(v int32) *KnowledgeBaseMetaDataConfigCustomMetaData {
	s.ValueCount = &v
	return s
}

func (s *KnowledgeBaseMetaDataConfigCustomMetaData) SetValueType(v string) *KnowledgeBaseMetaDataConfigCustomMetaData {
	s.ValueType = &v
	return s
}

func (s *KnowledgeBaseMetaDataConfigCustomMetaData) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseVectorDBConfig struct {
	// Collectioin名称
	CollectionName *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
	// Embedding连接ID
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// VectorDB连接名称
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// VectorDB类型
	VectorDBType *string `json:"VectorDBType,omitempty" xml:"VectorDBType,omitempty"`
}

func (s KnowledgeBaseVectorDBConfig) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseVectorDBConfig) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseVectorDBConfig) GetCollectionName() *string {
	return s.CollectionName
}

func (s *KnowledgeBaseVectorDBConfig) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *KnowledgeBaseVectorDBConfig) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *KnowledgeBaseVectorDBConfig) GetVectorDBType() *string {
	return s.VectorDBType
}

func (s *KnowledgeBaseVectorDBConfig) SetCollectionName(v string) *KnowledgeBaseVectorDBConfig {
	s.CollectionName = &v
	return s
}

func (s *KnowledgeBaseVectorDBConfig) SetConnectionId(v string) *KnowledgeBaseVectorDBConfig {
	s.ConnectionId = &v
	return s
}

func (s *KnowledgeBaseVectorDBConfig) SetConnectionName(v string) *KnowledgeBaseVectorDBConfig {
	s.ConnectionName = &v
	return s
}

func (s *KnowledgeBaseVectorDBConfig) SetVectorDBType(v string) *KnowledgeBaseVectorDBConfig {
	s.VectorDBType = &v
	return s
}

func (s *KnowledgeBaseVectorDBConfig) Validate() error {
	return dara.Validate(s)
}
