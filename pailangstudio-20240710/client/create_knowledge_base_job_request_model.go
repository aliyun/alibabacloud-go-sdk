// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreateKnowledgeBaseJobRequest
	GetAccessibility() *string
	SetDescription(v string) *CreateKnowledgeBaseJobRequest
	GetDescription() *string
	SetEcsSpecs(v []*CreateKnowledgeBaseJobRequestEcsSpecs) *CreateKnowledgeBaseJobRequest
	GetEcsSpecs() []*CreateKnowledgeBaseJobRequestEcsSpecs
	SetEmbeddingConfig(v *CreateKnowledgeBaseJobRequestEmbeddingConfig) *CreateKnowledgeBaseJobRequest
	GetEmbeddingConfig() *CreateKnowledgeBaseJobRequestEmbeddingConfig
	SetJobAction(v string) *CreateKnowledgeBaseJobRequest
	GetJobAction() *string
	SetMaxRunningTimeInSeconds(v int32) *CreateKnowledgeBaseJobRequest
	GetMaxRunningTimeInSeconds() *int32
	SetResourceId(v string) *CreateKnowledgeBaseJobRequest
	GetResourceId() *string
	SetUserVpc(v *CreateKnowledgeBaseJobRequestUserVpc) *CreateKnowledgeBaseJobRequest
	GetUserVpc() *CreateKnowledgeBaseJobRequestUserVpc
	SetWorkspaceId(v string) *CreateKnowledgeBaseJobRequest
	GetWorkspaceId() *string
}

type CreateKnowledgeBaseJobRequest struct {
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// This is a description of the knowledge base job.
	Description     *string                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsSpecs        []*CreateKnowledgeBaseJobRequestEcsSpecs      `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	EmbeddingConfig *CreateKnowledgeBaseJobRequestEmbeddingConfig `json:"EmbeddingConfig,omitempty" xml:"EmbeddingConfig,omitempty" type:"Struct"`
	JobAction       *string                                       `json:"JobAction,omitempty" xml:"JobAction,omitempty"`
	// example:
	//
	// 86400
	MaxRunningTimeInSeconds *int32 `json:"MaxRunningTimeInSeconds,omitempty" xml:"MaxRunningTimeInSeconds,omitempty"`
	// example:
	//
	// public-cluster
	ResourceId *string                               `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	UserVpc    *CreateKnowledgeBaseJobRequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateKnowledgeBaseJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseJobRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseJobRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateKnowledgeBaseJobRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseJobRequest) GetEcsSpecs() []*CreateKnowledgeBaseJobRequestEcsSpecs {
	return s.EcsSpecs
}

func (s *CreateKnowledgeBaseJobRequest) GetEmbeddingConfig() *CreateKnowledgeBaseJobRequestEmbeddingConfig {
	return s.EmbeddingConfig
}

func (s *CreateKnowledgeBaseJobRequest) GetJobAction() *string {
	return s.JobAction
}

func (s *CreateKnowledgeBaseJobRequest) GetMaxRunningTimeInSeconds() *int32 {
	return s.MaxRunningTimeInSeconds
}

func (s *CreateKnowledgeBaseJobRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateKnowledgeBaseJobRequest) GetUserVpc() *CreateKnowledgeBaseJobRequestUserVpc {
	return s.UserVpc
}

func (s *CreateKnowledgeBaseJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateKnowledgeBaseJobRequest) SetAccessibility(v string) *CreateKnowledgeBaseJobRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequest) SetDescription(v string) *CreateKnowledgeBaseJobRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequest) SetEcsSpecs(v []*CreateKnowledgeBaseJobRequestEcsSpecs) *CreateKnowledgeBaseJobRequest {
	s.EcsSpecs = v
	return s
}

func (s *CreateKnowledgeBaseJobRequest) SetEmbeddingConfig(v *CreateKnowledgeBaseJobRequestEmbeddingConfig) *CreateKnowledgeBaseJobRequest {
	s.EmbeddingConfig = v
	return s
}

func (s *CreateKnowledgeBaseJobRequest) SetJobAction(v string) *CreateKnowledgeBaseJobRequest {
	s.JobAction = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequest) SetMaxRunningTimeInSeconds(v int32) *CreateKnowledgeBaseJobRequest {
	s.MaxRunningTimeInSeconds = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequest) SetResourceId(v string) *CreateKnowledgeBaseJobRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequest) SetUserVpc(v *CreateKnowledgeBaseJobRequestUserVpc) *CreateKnowledgeBaseJobRequest {
	s.UserVpc = v
	return s
}

func (s *CreateKnowledgeBaseJobRequest) SetWorkspaceId(v string) *CreateKnowledgeBaseJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequest) Validate() error {
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

type CreateKnowledgeBaseJobRequestEcsSpecs struct {
	// CPU核数
	//
	// example:
	//
	// 2
	CPU *int32 `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// 驱动版本
	//
	// example:
	//
	// 535.161.08
	Driver *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	// GPU卡数
	//
	// example:
	//
	// 1
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
	//
	// example:
	//
	// 8
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// 副本数量
	//
	// example:
	//
	// 1
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

func (s CreateKnowledgeBaseJobRequestEcsSpecs) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseJobRequestEcsSpecs) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) GetCPU() *int32 {
	return s.CPU
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) GetDriver() *string {
	return s.Driver
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) GetGPU() *int32 {
	return s.GPU
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) GetGPUType() *string {
	return s.GPUType
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) GetPodCount() *int32 {
	return s.PodCount
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) GetSharedMemory() *int32 {
	return s.SharedMemory
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) GetType() *string {
	return s.Type
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) SetCPU(v int32) *CreateKnowledgeBaseJobRequestEcsSpecs {
	s.CPU = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) SetDriver(v string) *CreateKnowledgeBaseJobRequestEcsSpecs {
	s.Driver = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) SetGPU(v int32) *CreateKnowledgeBaseJobRequestEcsSpecs {
	s.GPU = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) SetGPUType(v string) *CreateKnowledgeBaseJobRequestEcsSpecs {
	s.GPUType = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) SetInstanceType(v string) *CreateKnowledgeBaseJobRequestEcsSpecs {
	s.InstanceType = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) SetMemory(v int32) *CreateKnowledgeBaseJobRequestEcsSpecs {
	s.Memory = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) SetPodCount(v int32) *CreateKnowledgeBaseJobRequestEcsSpecs {
	s.PodCount = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) SetSharedMemory(v int32) *CreateKnowledgeBaseJobRequestEcsSpecs {
	s.SharedMemory = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) SetType(v string) *CreateKnowledgeBaseJobRequestEcsSpecs {
	s.Type = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequestEcsSpecs) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseJobRequestEmbeddingConfig struct {
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

func (s CreateKnowledgeBaseJobRequestEmbeddingConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseJobRequestEmbeddingConfig) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseJobRequestEmbeddingConfig) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *CreateKnowledgeBaseJobRequestEmbeddingConfig) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *CreateKnowledgeBaseJobRequestEmbeddingConfig) SetBatchSize(v int32) *CreateKnowledgeBaseJobRequestEmbeddingConfig {
	s.BatchSize = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequestEmbeddingConfig) SetConcurrency(v int32) *CreateKnowledgeBaseJobRequestEmbeddingConfig {
	s.Concurrency = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequestEmbeddingConfig) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseJobRequestUserVpc struct {
	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 交换机ID
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateKnowledgeBaseJobRequestUserVpc) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseJobRequestUserVpc) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseJobRequestUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateKnowledgeBaseJobRequestUserVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateKnowledgeBaseJobRequestUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateKnowledgeBaseJobRequestUserVpc) SetSecurityGroupId(v string) *CreateKnowledgeBaseJobRequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequestUserVpc) SetVSwitchId(v string) *CreateKnowledgeBaseJobRequestUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequestUserVpc) SetVpcId(v string) *CreateKnowledgeBaseJobRequestUserVpc {
	s.VpcId = &v
	return s
}

func (s *CreateKnowledgeBaseJobRequestUserVpc) Validate() error {
	return dara.Validate(s)
}
