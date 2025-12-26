// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUpdateConfig(v *UpdateKnowledgeBaseRequestAutoUpdateConfig) *UpdateKnowledgeBaseRequest
	GetAutoUpdateConfig() *UpdateKnowledgeBaseRequestAutoUpdateConfig
	SetBindRuntime(v bool) *UpdateKnowledgeBaseRequest
	GetBindRuntime() *bool
	SetDescription(v string) *UpdateKnowledgeBaseRequest
	GetDescription() *string
	SetMetaDataConfig(v *UpdateKnowledgeBaseRequestMetaDataConfig) *UpdateKnowledgeBaseRequest
	GetMetaDataConfig() *UpdateKnowledgeBaseRequestMetaDataConfig
	SetRuntimeId(v string) *UpdateKnowledgeBaseRequest
	GetRuntimeId() *string
	SetWorkspaceId(v string) *UpdateKnowledgeBaseRequest
	GetWorkspaceId() *string
}

type UpdateKnowledgeBaseRequest struct {
	AutoUpdateConfig *UpdateKnowledgeBaseRequestAutoUpdateConfig `json:"AutoUpdateConfig,omitempty" xml:"AutoUpdateConfig,omitempty" type:"Struct"`
	BindRuntime      *bool                                       `json:"BindRuntime,omitempty" xml:"BindRuntime,omitempty"`
	Description      *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	MetaDataConfig   *UpdateKnowledgeBaseRequestMetaDataConfig   `json:"MetaDataConfig,omitempty" xml:"MetaDataConfig,omitempty" type:"Struct"`
	RuntimeId        *string                                     `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseRequest) GetAutoUpdateConfig() *UpdateKnowledgeBaseRequestAutoUpdateConfig {
	return s.AutoUpdateConfig
}

func (s *UpdateKnowledgeBaseRequest) GetBindRuntime() *bool {
	return s.BindRuntime
}

func (s *UpdateKnowledgeBaseRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeBaseRequest) GetMetaDataConfig() *UpdateKnowledgeBaseRequestMetaDataConfig {
	return s.MetaDataConfig
}

func (s *UpdateKnowledgeBaseRequest) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *UpdateKnowledgeBaseRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateKnowledgeBaseRequest) SetAutoUpdateConfig(v *UpdateKnowledgeBaseRequestAutoUpdateConfig) *UpdateKnowledgeBaseRequest {
	s.AutoUpdateConfig = v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetBindRuntime(v bool) *UpdateKnowledgeBaseRequest {
	s.BindRuntime = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetDescription(v string) *UpdateKnowledgeBaseRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetMetaDataConfig(v *UpdateKnowledgeBaseRequestMetaDataConfig) *UpdateKnowledgeBaseRequest {
	s.MetaDataConfig = v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetRuntimeId(v string) *UpdateKnowledgeBaseRequest {
	s.RuntimeId = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) SetWorkspaceId(v string) *UpdateKnowledgeBaseRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateKnowledgeBaseRequest) Validate() error {
	if s.AutoUpdateConfig != nil {
		if err := s.AutoUpdateConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MetaDataConfig != nil {
		if err := s.MetaDataConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeBaseRequestAutoUpdateConfig struct {
	// 运行资源配置
	EcsSpecs []*UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	// Embedding配置
	EmbeddingConfig *UpdateKnowledgeBaseRequestAutoUpdateConfigEmbeddingConfig `json:"EmbeddingConfig,omitempty" xml:"EmbeddingConfig,omitempty" type:"Struct"`
	// 任务最大运行时间
	//
	// example:
	//
	// 86400
	MaxRunningTimeInSeconds *int32 `json:"MaxRunningTimeInSeconds,omitempty" xml:"MaxRunningTimeInSeconds,omitempty"`
	// 资源组ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 知识库自动更新状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 用户VPC配置
	UserVpc *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
}

func (s UpdateKnowledgeBaseRequestAutoUpdateConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseRequestAutoUpdateConfig) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfig) GetEcsSpecs() []*UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs {
	return s.EcsSpecs
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfig) GetEmbeddingConfig() *UpdateKnowledgeBaseRequestAutoUpdateConfigEmbeddingConfig {
	return s.EmbeddingConfig
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfig) GetMaxRunningTimeInSeconds() *int32 {
	return s.MaxRunningTimeInSeconds
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfig) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfig) GetStatus() *string {
	return s.Status
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfig) GetUserVpc() *UserVpc {
	return s.UserVpc
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfig) SetEcsSpecs(v []*UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) *UpdateKnowledgeBaseRequestAutoUpdateConfig {
	s.EcsSpecs = v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfig) SetEmbeddingConfig(v *UpdateKnowledgeBaseRequestAutoUpdateConfigEmbeddingConfig) *UpdateKnowledgeBaseRequestAutoUpdateConfig {
	s.EmbeddingConfig = v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfig) SetMaxRunningTimeInSeconds(v int32) *UpdateKnowledgeBaseRequestAutoUpdateConfig {
	s.MaxRunningTimeInSeconds = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfig) SetResourceId(v string) *UpdateKnowledgeBaseRequestAutoUpdateConfig {
	s.ResourceId = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfig) SetStatus(v string) *UpdateKnowledgeBaseRequestAutoUpdateConfig {
	s.Status = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfig) SetUserVpc(v *UserVpc) *UpdateKnowledgeBaseRequestAutoUpdateConfig {
	s.UserVpc = v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfig) Validate() error {
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

type UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs struct {
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

func (s UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) GetCPU() *int32 {
	return s.CPU
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) GetDriver() *string {
	return s.Driver
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) GetGPU() *int32 {
	return s.GPU
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) GetGPUType() *string {
	return s.GPUType
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) GetMemory() *int32 {
	return s.Memory
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) GetPodCount() *int32 {
	return s.PodCount
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) GetSharedMemory() *int32 {
	return s.SharedMemory
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) GetType() *string {
	return s.Type
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) SetCPU(v int32) *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs {
	s.CPU = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) SetDriver(v string) *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs {
	s.Driver = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) SetGPU(v int32) *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs {
	s.GPU = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) SetGPUType(v string) *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs {
	s.GPUType = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) SetInstanceType(v string) *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs {
	s.InstanceType = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) SetMemory(v int32) *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs {
	s.Memory = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) SetPodCount(v int32) *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs {
	s.PodCount = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) SetSharedMemory(v int32) *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs {
	s.SharedMemory = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) SetType(v string) *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs {
	s.Type = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEcsSpecs) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeBaseRequestAutoUpdateConfigEmbeddingConfig struct {
	// Embedding分批大小
	BatchSize *int32 `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// Embedding并发数
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
}

func (s UpdateKnowledgeBaseRequestAutoUpdateConfigEmbeddingConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseRequestAutoUpdateConfigEmbeddingConfig) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEmbeddingConfig) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEmbeddingConfig) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEmbeddingConfig) SetBatchSize(v int32) *UpdateKnowledgeBaseRequestAutoUpdateConfigEmbeddingConfig {
	s.BatchSize = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEmbeddingConfig) SetConcurrency(v int32) *UpdateKnowledgeBaseRequestAutoUpdateConfigEmbeddingConfig {
	s.Concurrency = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestAutoUpdateConfigEmbeddingConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeBaseRequestMetaDataConfig struct {
	// 自定义元数据
	CustomMetaData []*UpdateKnowledgeBaseRequestMetaDataConfigCustomMetaData `json:"CustomMetaData,omitempty" xml:"CustomMetaData,omitempty" type:"Repeated"`
}

func (s UpdateKnowledgeBaseRequestMetaDataConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseRequestMetaDataConfig) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseRequestMetaDataConfig) GetCustomMetaData() []*UpdateKnowledgeBaseRequestMetaDataConfigCustomMetaData {
	return s.CustomMetaData
}

func (s *UpdateKnowledgeBaseRequestMetaDataConfig) SetCustomMetaData(v []*UpdateKnowledgeBaseRequestMetaDataConfigCustomMetaData) *UpdateKnowledgeBaseRequestMetaDataConfig {
	s.CustomMetaData = v
	return s
}

func (s *UpdateKnowledgeBaseRequestMetaDataConfig) Validate() error {
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

type UpdateKnowledgeBaseRequestMetaDataConfigCustomMetaData struct {
	// 元数据Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 元数据Value类型
	//
	// example:
	//
	// String
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s UpdateKnowledgeBaseRequestMetaDataConfigCustomMetaData) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseRequestMetaDataConfigCustomMetaData) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseRequestMetaDataConfigCustomMetaData) GetKey() *string {
	return s.Key
}

func (s *UpdateKnowledgeBaseRequestMetaDataConfigCustomMetaData) GetValueType() *string {
	return s.ValueType
}

func (s *UpdateKnowledgeBaseRequestMetaDataConfigCustomMetaData) SetKey(v string) *UpdateKnowledgeBaseRequestMetaDataConfigCustomMetaData {
	s.Key = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestMetaDataConfigCustomMetaData) SetValueType(v string) *UpdateKnowledgeBaseRequestMetaDataConfigCustomMetaData {
	s.ValueType = &v
	return s
}

func (s *UpdateKnowledgeBaseRequestMetaDataConfigCustomMetaData) Validate() error {
	return dara.Validate(s)
}
