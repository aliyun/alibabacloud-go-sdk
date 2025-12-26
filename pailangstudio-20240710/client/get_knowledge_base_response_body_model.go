// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *GetKnowledgeBaseResponseBody
	GetAccessibility() *string
	SetAutoUpdateConfig(v *GetKnowledgeBaseResponseBodyAutoUpdateConfig) *GetKnowledgeBaseResponseBody
	GetAutoUpdateConfig() *GetKnowledgeBaseResponseBodyAutoUpdateConfig
	SetChunkConfig(v *GetKnowledgeBaseResponseBodyChunkConfig) *GetKnowledgeBaseResponseBody
	GetChunkConfig() *GetKnowledgeBaseResponseBodyChunkConfig
	SetCreator(v string) *GetKnowledgeBaseResponseBody
	GetCreator() *string
	SetDataSources(v []*GetKnowledgeBaseResponseBodyDataSources) *GetKnowledgeBaseResponseBody
	GetDataSources() []*GetKnowledgeBaseResponseBodyDataSources
	SetDatasetId(v string) *GetKnowledgeBaseResponseBody
	GetDatasetId() *string
	SetDescription(v string) *GetKnowledgeBaseResponseBody
	GetDescription() *string
	SetEmbeddingConfig(v *GetKnowledgeBaseResponseBodyEmbeddingConfig) *GetKnowledgeBaseResponseBody
	GetEmbeddingConfig() *GetKnowledgeBaseResponseBodyEmbeddingConfig
	SetGmtCreateTime(v string) *GetKnowledgeBaseResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetKnowledgeBaseResponseBody
	GetGmtModifiedTime() *string
	SetIndexColumnConfig(v *GetKnowledgeBaseResponseBodyIndexColumnConfig) *GetKnowledgeBaseResponseBody
	GetIndexColumnConfig() *GetKnowledgeBaseResponseBodyIndexColumnConfig
	SetKnowledgeBaseId(v string) *GetKnowledgeBaseResponseBody
	GetKnowledgeBaseId() *string
	SetKnowledgeBaseType(v string) *GetKnowledgeBaseResponseBody
	GetKnowledgeBaseType() *string
	SetMetaDataConfig(v *GetKnowledgeBaseResponseBodyMetaDataConfig) *GetKnowledgeBaseResponseBody
	GetMetaDataConfig() *GetKnowledgeBaseResponseBodyMetaDataConfig
	SetName(v string) *GetKnowledgeBaseResponseBody
	GetName() *string
	SetOutputDir(v string) *GetKnowledgeBaseResponseBody
	GetOutputDir() *string
	SetRequestId(v string) *GetKnowledgeBaseResponseBody
	GetRequestId() *string
	SetRuntimeId(v string) *GetKnowledgeBaseResponseBody
	GetRuntimeId() *string
	SetVectorDBConfig(v *GetKnowledgeBaseResponseBodyVectorDBConfig) *GetKnowledgeBaseResponseBody
	GetVectorDBConfig() *GetKnowledgeBaseResponseBodyVectorDBConfig
	SetVersionName(v string) *GetKnowledgeBaseResponseBody
	GetVersionName() *string
	SetWorkspaceId(v string) *GetKnowledgeBaseResponseBody
	GetWorkspaceId() *string
}

type GetKnowledgeBaseResponseBody struct {
	// example:
	//
	// PRIVATE
	Accessibility    *string                                       `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	AutoUpdateConfig *GetKnowledgeBaseResponseBodyAutoUpdateConfig `json:"AutoUpdateConfig,omitempty" xml:"AutoUpdateConfig,omitempty" type:"Struct"`
	ChunkConfig      *GetKnowledgeBaseResponseBodyChunkConfig      `json:"ChunkConfig,omitempty" xml:"ChunkConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2485765****023475
	Creator     *string                                    `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DataSources []*GetKnowledgeBaseResponseBodyDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// example:
	//
	// d-cupbwkk5us9xpjz870
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// This is a description of the knowledge base.
	Description     *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	EmbeddingConfig *GetKnowledgeBaseResponseBodyEmbeddingConfig `json:"EmbeddingConfig,omitempty" xml:"EmbeddingConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2024-12-15T14:46:23Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2025-12-18T19:32:58Z
	GmtModifiedTime   *string                                        `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	IndexColumnConfig *GetKnowledgeBaseResponseBodyIndexColumnConfig `json:"IndexColumnConfig,omitempty" xml:"IndexColumnConfig,omitempty" type:"Struct"`
	// example:
	//
	// d-ksicx823d
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// example:
	//
	// TEXT
	KnowledgeBaseType *string                                     `json:"KnowledgeBaseType,omitempty" xml:"KnowledgeBaseType,omitempty"`
	MetaDataConfig    *GetKnowledgeBaseResponseBodyMetaDataConfig `json:"MetaDataConfig,omitempty" xml:"MetaDataConfig,omitempty" type:"Struct"`
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou-internal.aliyuncs.com/langstudio/output/
	OutputDir *string `json:"OutputDir,omitempty" xml:"OutputDir,omitempty"`
	// example:
	//
	// 963BD7F9-0C02-5594-9550-BCC6DD43E3C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rtime-apje******beaz
	RuntimeId      *string                                     `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	VectorDBConfig *GetKnowledgeBaseResponseBodyVectorDBConfig `json:"VectorDBConfig,omitempty" xml:"VectorDBConfig,omitempty" type:"Struct"`
	VersionName    *string                                     `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetKnowledgeBaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBody) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetKnowledgeBaseResponseBody) GetAutoUpdateConfig() *GetKnowledgeBaseResponseBodyAutoUpdateConfig {
	return s.AutoUpdateConfig
}

func (s *GetKnowledgeBaseResponseBody) GetChunkConfig() *GetKnowledgeBaseResponseBodyChunkConfig {
	return s.ChunkConfig
}

func (s *GetKnowledgeBaseResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetKnowledgeBaseResponseBody) GetDataSources() []*GetKnowledgeBaseResponseBodyDataSources {
	return s.DataSources
}

func (s *GetKnowledgeBaseResponseBody) GetDatasetId() *string {
	return s.DatasetId
}

func (s *GetKnowledgeBaseResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetKnowledgeBaseResponseBody) GetEmbeddingConfig() *GetKnowledgeBaseResponseBodyEmbeddingConfig {
	return s.EmbeddingConfig
}

func (s *GetKnowledgeBaseResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetKnowledgeBaseResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetKnowledgeBaseResponseBody) GetIndexColumnConfig() *GetKnowledgeBaseResponseBodyIndexColumnConfig {
	return s.IndexColumnConfig
}

func (s *GetKnowledgeBaseResponseBody) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgeBaseResponseBody) GetKnowledgeBaseType() *string {
	return s.KnowledgeBaseType
}

func (s *GetKnowledgeBaseResponseBody) GetMetaDataConfig() *GetKnowledgeBaseResponseBodyMetaDataConfig {
	return s.MetaDataConfig
}

func (s *GetKnowledgeBaseResponseBody) GetName() *string {
	return s.Name
}

func (s *GetKnowledgeBaseResponseBody) GetOutputDir() *string {
	return s.OutputDir
}

func (s *GetKnowledgeBaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKnowledgeBaseResponseBody) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *GetKnowledgeBaseResponseBody) GetVectorDBConfig() *GetKnowledgeBaseResponseBodyVectorDBConfig {
	return s.VectorDBConfig
}

func (s *GetKnowledgeBaseResponseBody) GetVersionName() *string {
	return s.VersionName
}

func (s *GetKnowledgeBaseResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetKnowledgeBaseResponseBody) SetAccessibility(v string) *GetKnowledgeBaseResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetAutoUpdateConfig(v *GetKnowledgeBaseResponseBodyAutoUpdateConfig) *GetKnowledgeBaseResponseBody {
	s.AutoUpdateConfig = v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetChunkConfig(v *GetKnowledgeBaseResponseBodyChunkConfig) *GetKnowledgeBaseResponseBody {
	s.ChunkConfig = v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetCreator(v string) *GetKnowledgeBaseResponseBody {
	s.Creator = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetDataSources(v []*GetKnowledgeBaseResponseBodyDataSources) *GetKnowledgeBaseResponseBody {
	s.DataSources = v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetDatasetId(v string) *GetKnowledgeBaseResponseBody {
	s.DatasetId = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetDescription(v string) *GetKnowledgeBaseResponseBody {
	s.Description = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetEmbeddingConfig(v *GetKnowledgeBaseResponseBodyEmbeddingConfig) *GetKnowledgeBaseResponseBody {
	s.EmbeddingConfig = v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetGmtCreateTime(v string) *GetKnowledgeBaseResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetGmtModifiedTime(v string) *GetKnowledgeBaseResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetIndexColumnConfig(v *GetKnowledgeBaseResponseBodyIndexColumnConfig) *GetKnowledgeBaseResponseBody {
	s.IndexColumnConfig = v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetKnowledgeBaseId(v string) *GetKnowledgeBaseResponseBody {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetKnowledgeBaseType(v string) *GetKnowledgeBaseResponseBody {
	s.KnowledgeBaseType = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetMetaDataConfig(v *GetKnowledgeBaseResponseBodyMetaDataConfig) *GetKnowledgeBaseResponseBody {
	s.MetaDataConfig = v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetName(v string) *GetKnowledgeBaseResponseBody {
	s.Name = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetOutputDir(v string) *GetKnowledgeBaseResponseBody {
	s.OutputDir = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetRequestId(v string) *GetKnowledgeBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetRuntimeId(v string) *GetKnowledgeBaseResponseBody {
	s.RuntimeId = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetVectorDBConfig(v *GetKnowledgeBaseResponseBodyVectorDBConfig) *GetKnowledgeBaseResponseBody {
	s.VectorDBConfig = v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetVersionName(v string) *GetKnowledgeBaseResponseBody {
	s.VersionName = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) SetWorkspaceId(v string) *GetKnowledgeBaseResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetKnowledgeBaseResponseBody) Validate() error {
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

type GetKnowledgeBaseResponseBodyAutoUpdateConfig struct {
	// 运行资源配置
	EcsSpecs []*GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	// Embedding配置
	EmbeddingConfig *GetKnowledgeBaseResponseBodyAutoUpdateConfigEmbeddingConfig `json:"EmbeddingConfig,omitempty" xml:"EmbeddingConfig,omitempty" type:"Struct"`
	// 任务最大运行时间
	//
	// example:
	//
	// 86400
	MaxRunningTimeInSeconds *int32 `json:"MaxRunningTimeInSeconds,omitempty" xml:"MaxRunningTimeInSeconds,omitempty"`
	// 资源组ID
	//
	// example:
	//
	// public-cluster
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 知识库自动更新状态
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 用户VPC配置
	UserVpc *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
}

func (s GetKnowledgeBaseResponseBodyAutoUpdateConfig) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBodyAutoUpdateConfig) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfig) GetEcsSpecs() []*GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs {
	return s.EcsSpecs
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfig) GetEmbeddingConfig() *GetKnowledgeBaseResponseBodyAutoUpdateConfigEmbeddingConfig {
	return s.EmbeddingConfig
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfig) GetMaxRunningTimeInSeconds() *int32 {
	return s.MaxRunningTimeInSeconds
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfig) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfig) GetStatus() *string {
	return s.Status
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfig) GetUserVpc() *UserVpc {
	return s.UserVpc
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfig) SetEcsSpecs(v []*GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) *GetKnowledgeBaseResponseBodyAutoUpdateConfig {
	s.EcsSpecs = v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfig) SetEmbeddingConfig(v *GetKnowledgeBaseResponseBodyAutoUpdateConfigEmbeddingConfig) *GetKnowledgeBaseResponseBodyAutoUpdateConfig {
	s.EmbeddingConfig = v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfig) SetMaxRunningTimeInSeconds(v int32) *GetKnowledgeBaseResponseBodyAutoUpdateConfig {
	s.MaxRunningTimeInSeconds = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfig) SetResourceId(v string) *GetKnowledgeBaseResponseBodyAutoUpdateConfig {
	s.ResourceId = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfig) SetStatus(v string) *GetKnowledgeBaseResponseBodyAutoUpdateConfig {
	s.Status = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfig) SetUserVpc(v *UserVpc) *GetKnowledgeBaseResponseBodyAutoUpdateConfig {
	s.UserVpc = v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfig) Validate() error {
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

type GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs struct {
	// CPU核数
	CPU *int32 `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// 驱动版本
	Driver *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	// GPU卡数
	GPU *int32 `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// GPU类型
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// 机型名称
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 内存大小
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// 副本数量
	PodCount *int32 `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
	// 共享内存容量
	SharedMemory *int32 `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
	// 节点类型
	//
	// example:
	//
	// Worker
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) GetCPU() *int32 {
	return s.CPU
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) GetDriver() *string {
	return s.Driver
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) GetGPU() *int32 {
	return s.GPU
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) GetGPUType() *string {
	return s.GPUType
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) GetMemory() *int32 {
	return s.Memory
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) GetPodCount() *int32 {
	return s.PodCount
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) GetSharedMemory() *int32 {
	return s.SharedMemory
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) GetType() *string {
	return s.Type
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) SetCPU(v int32) *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs {
	s.CPU = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) SetDriver(v string) *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs {
	s.Driver = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) SetGPU(v int32) *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs {
	s.GPU = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) SetGPUType(v string) *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs {
	s.GPUType = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) SetInstanceType(v string) *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs {
	s.InstanceType = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) SetMemory(v int32) *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs {
	s.Memory = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) SetPodCount(v int32) *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs {
	s.PodCount = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) SetSharedMemory(v int32) *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs {
	s.SharedMemory = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) SetType(v string) *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs {
	s.Type = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEcsSpecs) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseResponseBodyAutoUpdateConfigEmbeddingConfig struct {
	// Embedding分批大小
	//
	// example:
	//
	// 8
	BatchSize *int32 `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// Embedding并发数
	//
	// example:
	//
	// 1
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
}

func (s GetKnowledgeBaseResponseBodyAutoUpdateConfigEmbeddingConfig) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBodyAutoUpdateConfigEmbeddingConfig) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEmbeddingConfig) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEmbeddingConfig) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEmbeddingConfig) SetBatchSize(v int32) *GetKnowledgeBaseResponseBodyAutoUpdateConfigEmbeddingConfig {
	s.BatchSize = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEmbeddingConfig) SetConcurrency(v int32) *GetKnowledgeBaseResponseBodyAutoUpdateConfigEmbeddingConfig {
	s.Concurrency = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyAutoUpdateConfigEmbeddingConfig) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseResponseBodyChunkConfig struct {
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

func (s GetKnowledgeBaseResponseBodyChunkConfig) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBodyChunkConfig) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBodyChunkConfig) GetChunkDuration() *int32 {
	return s.ChunkDuration
}

func (s *GetKnowledgeBaseResponseBodyChunkConfig) GetChunkOverlap() *int32 {
	return s.ChunkOverlap
}

func (s *GetKnowledgeBaseResponseBodyChunkConfig) GetChunkSize() *int32 {
	return s.ChunkSize
}

func (s *GetKnowledgeBaseResponseBodyChunkConfig) GetChunkStrategy() *string {
	return s.ChunkStrategy
}

func (s *GetKnowledgeBaseResponseBodyChunkConfig) SetChunkDuration(v int32) *GetKnowledgeBaseResponseBodyChunkConfig {
	s.ChunkDuration = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyChunkConfig) SetChunkOverlap(v int32) *GetKnowledgeBaseResponseBodyChunkConfig {
	s.ChunkOverlap = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyChunkConfig) SetChunkSize(v int32) *GetKnowledgeBaseResponseBodyChunkConfig {
	s.ChunkSize = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyChunkConfig) SetChunkStrategy(v string) *GetKnowledgeBaseResponseBodyChunkConfig {
	s.ChunkStrategy = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyChunkConfig) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseResponseBodyDataSources struct {
	// 统一资源识别码
	//
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou-internal.aliyuncs.com/langstudio/source/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s GetKnowledgeBaseResponseBodyDataSources) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBodyDataSources) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBodyDataSources) GetUri() *string {
	return s.Uri
}

func (s *GetKnowledgeBaseResponseBodyDataSources) SetUri(v string) *GetKnowledgeBaseResponseBodyDataSources {
	s.Uri = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyDataSources) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseResponseBodyEmbeddingConfig struct {
	// Embedding连接ID
	//
	// example:
	//
	// conn-r3o7******38bh
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// Embedding连接名称
	//
	// example:
	//
	// myEmbeddingConn
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// 模型
	//
	// example:
	//
	// text-embedding-v4
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s GetKnowledgeBaseResponseBodyEmbeddingConfig) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBodyEmbeddingConfig) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBodyEmbeddingConfig) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *GetKnowledgeBaseResponseBodyEmbeddingConfig) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *GetKnowledgeBaseResponseBodyEmbeddingConfig) GetModel() *string {
	return s.Model
}

func (s *GetKnowledgeBaseResponseBodyEmbeddingConfig) SetConnectionId(v string) *GetKnowledgeBaseResponseBodyEmbeddingConfig {
	s.ConnectionId = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyEmbeddingConfig) SetConnectionName(v string) *GetKnowledgeBaseResponseBodyEmbeddingConfig {
	s.ConnectionName = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyEmbeddingConfig) SetModel(v string) *GetKnowledgeBaseResponseBodyEmbeddingConfig {
	s.Model = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyEmbeddingConfig) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseResponseBodyIndexColumnConfig struct {
	// 所有列名
	ColumnDefinitions []*GetKnowledgeBaseResponseBodyIndexColumnConfigColumnDefinitions `json:"ColumnDefinitions,omitempty" xml:"ColumnDefinitions,omitempty" type:"Repeated"`
	// 内容检索列
	ContentColumns []*GetKnowledgeBaseResponseBodyIndexColumnConfigContentColumns `json:"ContentColumns,omitempty" xml:"ContentColumns,omitempty" type:"Repeated"`
	// Embedding列
	EmbeddingColumns []*GetKnowledgeBaseResponseBodyIndexColumnConfigEmbeddingColumns `json:"EmbeddingColumns,omitempty" xml:"EmbeddingColumns,omitempty" type:"Repeated"`
}

func (s GetKnowledgeBaseResponseBodyIndexColumnConfig) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBodyIndexColumnConfig) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfig) GetColumnDefinitions() []*GetKnowledgeBaseResponseBodyIndexColumnConfigColumnDefinitions {
	return s.ColumnDefinitions
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfig) GetContentColumns() []*GetKnowledgeBaseResponseBodyIndexColumnConfigContentColumns {
	return s.ContentColumns
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfig) GetEmbeddingColumns() []*GetKnowledgeBaseResponseBodyIndexColumnConfigEmbeddingColumns {
	return s.EmbeddingColumns
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfig) SetColumnDefinitions(v []*GetKnowledgeBaseResponseBodyIndexColumnConfigColumnDefinitions) *GetKnowledgeBaseResponseBodyIndexColumnConfig {
	s.ColumnDefinitions = v
	return s
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfig) SetContentColumns(v []*GetKnowledgeBaseResponseBodyIndexColumnConfigContentColumns) *GetKnowledgeBaseResponseBodyIndexColumnConfig {
	s.ContentColumns = v
	return s
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfig) SetEmbeddingColumns(v []*GetKnowledgeBaseResponseBodyIndexColumnConfigEmbeddingColumns) *GetKnowledgeBaseResponseBodyIndexColumnConfig {
	s.EmbeddingColumns = v
	return s
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfig) Validate() error {
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

type GetKnowledgeBaseResponseBodyIndexColumnConfigColumnDefinitions struct {
	// 列Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s GetKnowledgeBaseResponseBodyIndexColumnConfigColumnDefinitions) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBodyIndexColumnConfigColumnDefinitions) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfigColumnDefinitions) GetKey() *string {
	return s.Key
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfigColumnDefinitions) SetKey(v string) *GetKnowledgeBaseResponseBodyIndexColumnConfigColumnDefinitions {
	s.Key = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfigColumnDefinitions) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseResponseBodyIndexColumnConfigContentColumns struct {
	// 列Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s GetKnowledgeBaseResponseBodyIndexColumnConfigContentColumns) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBodyIndexColumnConfigContentColumns) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfigContentColumns) GetKey() *string {
	return s.Key
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfigContentColumns) SetKey(v string) *GetKnowledgeBaseResponseBodyIndexColumnConfigContentColumns {
	s.Key = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfigContentColumns) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseResponseBodyIndexColumnConfigEmbeddingColumns struct {
	// 列Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s GetKnowledgeBaseResponseBodyIndexColumnConfigEmbeddingColumns) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBodyIndexColumnConfigEmbeddingColumns) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfigEmbeddingColumns) GetKey() *string {
	return s.Key
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfigEmbeddingColumns) SetKey(v string) *GetKnowledgeBaseResponseBodyIndexColumnConfigEmbeddingColumns {
	s.Key = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyIndexColumnConfigEmbeddingColumns) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseResponseBodyMetaDataConfig struct {
	// 自定义元数据
	CustomMetaData []*GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData `json:"CustomMetaData,omitempty" xml:"CustomMetaData,omitempty" type:"Repeated"`
}

func (s GetKnowledgeBaseResponseBodyMetaDataConfig) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBodyMetaDataConfig) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBodyMetaDataConfig) GetCustomMetaData() []*GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData {
	return s.CustomMetaData
}

func (s *GetKnowledgeBaseResponseBodyMetaDataConfig) SetCustomMetaData(v []*GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData) *GetKnowledgeBaseResponseBodyMetaDataConfig {
	s.CustomMetaData = v
	return s
}

func (s *GetKnowledgeBaseResponseBodyMetaDataConfig) Validate() error {
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

type GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData struct {
	// 元数据Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 引用次数
	ReferenceCount *int32 `json:"ReferenceCount,omitempty" xml:"ReferenceCount,omitempty"`
	// 值的个数
	ValueCount *int32 `json:"ValueCount,omitempty" xml:"ValueCount,omitempty"`
	// 元数据Value类型
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData) GetKey() *string {
	return s.Key
}

func (s *GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData) GetReferenceCount() *int32 {
	return s.ReferenceCount
}

func (s *GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData) GetValueCount() *int32 {
	return s.ValueCount
}

func (s *GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData) GetValueType() *string {
	return s.ValueType
}

func (s *GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData) SetKey(v string) *GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData {
	s.Key = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData) SetReferenceCount(v int32) *GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData {
	s.ReferenceCount = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData) SetValueCount(v int32) *GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData {
	s.ValueCount = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData) SetValueType(v string) *GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData {
	s.ValueType = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyMetaDataConfigCustomMetaData) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseResponseBodyVectorDBConfig struct {
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
	// VectorDB连接名称
	//
	// example:
	//
	// myVectorConn
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// VectorDB类型
	//
	// example:
	//
	// Milvus
	VectorDBType *string `json:"VectorDBType,omitempty" xml:"VectorDBType,omitempty"`
}

func (s GetKnowledgeBaseResponseBodyVectorDBConfig) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponseBodyVectorDBConfig) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponseBodyVectorDBConfig) GetCollectionName() *string {
	return s.CollectionName
}

func (s *GetKnowledgeBaseResponseBodyVectorDBConfig) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *GetKnowledgeBaseResponseBodyVectorDBConfig) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *GetKnowledgeBaseResponseBodyVectorDBConfig) GetVectorDBType() *string {
	return s.VectorDBType
}

func (s *GetKnowledgeBaseResponseBodyVectorDBConfig) SetCollectionName(v string) *GetKnowledgeBaseResponseBodyVectorDBConfig {
	s.CollectionName = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyVectorDBConfig) SetConnectionId(v string) *GetKnowledgeBaseResponseBodyVectorDBConfig {
	s.ConnectionId = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyVectorDBConfig) SetConnectionName(v string) *GetKnowledgeBaseResponseBodyVectorDBConfig {
	s.ConnectionName = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyVectorDBConfig) SetVectorDBType(v string) *GetKnowledgeBaseResponseBodyVectorDBConfig {
	s.VectorDBType = &v
	return s
}

func (s *GetKnowledgeBaseResponseBodyVectorDBConfig) Validate() error {
	return dara.Validate(s)
}
