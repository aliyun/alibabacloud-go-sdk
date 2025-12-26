// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutoUpdateConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEcsSpecs(v []*AutoUpdateConfigEcsSpecs) *AutoUpdateConfig
	GetEcsSpecs() []*AutoUpdateConfigEcsSpecs
	SetEmbeddingConfig(v *AutoUpdateConfigEmbeddingConfig) *AutoUpdateConfig
	GetEmbeddingConfig() *AutoUpdateConfigEmbeddingConfig
	SetMaxRunningTimeInSeconds(v int32) *AutoUpdateConfig
	GetMaxRunningTimeInSeconds() *int32
	SetResourceId(v string) *AutoUpdateConfig
	GetResourceId() *string
	SetStatus(v string) *AutoUpdateConfig
	GetStatus() *string
	SetUserVpc(v *UserVpc) *AutoUpdateConfig
	GetUserVpc() *UserVpc
}

type AutoUpdateConfig struct {
	// 运行资源配置
	EcsSpecs []*AutoUpdateConfigEcsSpecs `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	// Embedding配置
	EmbeddingConfig *AutoUpdateConfigEmbeddingConfig `json:"EmbeddingConfig,omitempty" xml:"EmbeddingConfig,omitempty" type:"Struct"`
	// 任务最大运行时间
	MaxRunningTimeInSeconds *int32 `json:"MaxRunningTimeInSeconds,omitempty" xml:"MaxRunningTimeInSeconds,omitempty"`
	// 资源组ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 知识库自动更新状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 用户VPC配置
	UserVpc *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
}

func (s AutoUpdateConfig) String() string {
	return dara.Prettify(s)
}

func (s AutoUpdateConfig) GoString() string {
	return s.String()
}

func (s *AutoUpdateConfig) GetEcsSpecs() []*AutoUpdateConfigEcsSpecs {
	return s.EcsSpecs
}

func (s *AutoUpdateConfig) GetEmbeddingConfig() *AutoUpdateConfigEmbeddingConfig {
	return s.EmbeddingConfig
}

func (s *AutoUpdateConfig) GetMaxRunningTimeInSeconds() *int32 {
	return s.MaxRunningTimeInSeconds
}

func (s *AutoUpdateConfig) GetResourceId() *string {
	return s.ResourceId
}

func (s *AutoUpdateConfig) GetStatus() *string {
	return s.Status
}

func (s *AutoUpdateConfig) GetUserVpc() *UserVpc {
	return s.UserVpc
}

func (s *AutoUpdateConfig) SetEcsSpecs(v []*AutoUpdateConfigEcsSpecs) *AutoUpdateConfig {
	s.EcsSpecs = v
	return s
}

func (s *AutoUpdateConfig) SetEmbeddingConfig(v *AutoUpdateConfigEmbeddingConfig) *AutoUpdateConfig {
	s.EmbeddingConfig = v
	return s
}

func (s *AutoUpdateConfig) SetMaxRunningTimeInSeconds(v int32) *AutoUpdateConfig {
	s.MaxRunningTimeInSeconds = &v
	return s
}

func (s *AutoUpdateConfig) SetResourceId(v string) *AutoUpdateConfig {
	s.ResourceId = &v
	return s
}

func (s *AutoUpdateConfig) SetStatus(v string) *AutoUpdateConfig {
	s.Status = &v
	return s
}

func (s *AutoUpdateConfig) SetUserVpc(v *UserVpc) *AutoUpdateConfig {
	s.UserVpc = v
	return s
}

func (s *AutoUpdateConfig) Validate() error {
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

type AutoUpdateConfigEcsSpecs struct {
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

func (s AutoUpdateConfigEcsSpecs) String() string {
	return dara.Prettify(s)
}

func (s AutoUpdateConfigEcsSpecs) GoString() string {
	return s.String()
}

func (s *AutoUpdateConfigEcsSpecs) GetCPU() *int32 {
	return s.CPU
}

func (s *AutoUpdateConfigEcsSpecs) GetDriver() *string {
	return s.Driver
}

func (s *AutoUpdateConfigEcsSpecs) GetGPU() *int32 {
	return s.GPU
}

func (s *AutoUpdateConfigEcsSpecs) GetGPUType() *string {
	return s.GPUType
}

func (s *AutoUpdateConfigEcsSpecs) GetInstanceType() *string {
	return s.InstanceType
}

func (s *AutoUpdateConfigEcsSpecs) GetMemory() *int32 {
	return s.Memory
}

func (s *AutoUpdateConfigEcsSpecs) GetPodCount() *int32 {
	return s.PodCount
}

func (s *AutoUpdateConfigEcsSpecs) GetSharedMemory() *int32 {
	return s.SharedMemory
}

func (s *AutoUpdateConfigEcsSpecs) GetType() *string {
	return s.Type
}

func (s *AutoUpdateConfigEcsSpecs) SetCPU(v int32) *AutoUpdateConfigEcsSpecs {
	s.CPU = &v
	return s
}

func (s *AutoUpdateConfigEcsSpecs) SetDriver(v string) *AutoUpdateConfigEcsSpecs {
	s.Driver = &v
	return s
}

func (s *AutoUpdateConfigEcsSpecs) SetGPU(v int32) *AutoUpdateConfigEcsSpecs {
	s.GPU = &v
	return s
}

func (s *AutoUpdateConfigEcsSpecs) SetGPUType(v string) *AutoUpdateConfigEcsSpecs {
	s.GPUType = &v
	return s
}

func (s *AutoUpdateConfigEcsSpecs) SetInstanceType(v string) *AutoUpdateConfigEcsSpecs {
	s.InstanceType = &v
	return s
}

func (s *AutoUpdateConfigEcsSpecs) SetMemory(v int32) *AutoUpdateConfigEcsSpecs {
	s.Memory = &v
	return s
}

func (s *AutoUpdateConfigEcsSpecs) SetPodCount(v int32) *AutoUpdateConfigEcsSpecs {
	s.PodCount = &v
	return s
}

func (s *AutoUpdateConfigEcsSpecs) SetSharedMemory(v int32) *AutoUpdateConfigEcsSpecs {
	s.SharedMemory = &v
	return s
}

func (s *AutoUpdateConfigEcsSpecs) SetType(v string) *AutoUpdateConfigEcsSpecs {
	s.Type = &v
	return s
}

func (s *AutoUpdateConfigEcsSpecs) Validate() error {
	return dara.Validate(s)
}

type AutoUpdateConfigEmbeddingConfig struct {
	// Embedding分批大小
	BatchSize *int32 `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// Embedding并发数
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
}

func (s AutoUpdateConfigEmbeddingConfig) String() string {
	return dara.Prettify(s)
}

func (s AutoUpdateConfigEmbeddingConfig) GoString() string {
	return s.String()
}

func (s *AutoUpdateConfigEmbeddingConfig) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *AutoUpdateConfigEmbeddingConfig) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *AutoUpdateConfigEmbeddingConfig) SetBatchSize(v int32) *AutoUpdateConfigEmbeddingConfig {
	s.BatchSize = &v
	return s
}

func (s *AutoUpdateConfigEmbeddingConfig) SetConcurrency(v int32) *AutoUpdateConfigEmbeddingConfig {
	s.Concurrency = &v
	return s
}

func (s *AutoUpdateConfigEmbeddingConfig) Validate() error {
	return dara.Validate(s)
}
