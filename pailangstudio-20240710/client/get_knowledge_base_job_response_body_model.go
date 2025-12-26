// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *GetKnowledgeBaseJobResponseBody
	GetAccessibility() *string
	SetCreator(v string) *GetKnowledgeBaseJobResponseBody
	GetCreator() *string
	SetDescription(v string) *GetKnowledgeBaseJobResponseBody
	GetDescription() *string
	SetEcsSpecs(v []*GetKnowledgeBaseJobResponseBodyEcsSpecs) *GetKnowledgeBaseJobResponseBody
	GetEcsSpecs() []*GetKnowledgeBaseJobResponseBodyEcsSpecs
	SetEmbeddingConfig(v *GetKnowledgeBaseJobResponseBodyEmbeddingConfig) *GetKnowledgeBaseJobResponseBody
	GetEmbeddingConfig() *GetKnowledgeBaseJobResponseBodyEmbeddingConfig
	SetErrorMessage(v string) *GetKnowledgeBaseJobResponseBody
	GetErrorMessage() *string
	SetGmtCreateTime(v string) *GetKnowledgeBaseJobResponseBody
	GetGmtCreateTime() *string
	SetGmtFinishTime(v string) *GetKnowledgeBaseJobResponseBody
	GetGmtFinishTime() *string
	SetGmtModifiedTime(v string) *GetKnowledgeBaseJobResponseBody
	GetGmtModifiedTime() *string
	SetJobAction(v string) *GetKnowledgeBaseJobResponseBody
	GetJobAction() *string
	SetKnowledgeBaseId(v string) *GetKnowledgeBaseJobResponseBody
	GetKnowledgeBaseId() *string
	SetKnowledgeBaseJobId(v string) *GetKnowledgeBaseJobResponseBody
	GetKnowledgeBaseJobId() *string
	SetKnowledgeBaseJobResult(v *GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult) *GetKnowledgeBaseJobResponseBody
	GetKnowledgeBaseJobResult() *GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult
	SetMaxRunningTimeInSeconds(v int32) *GetKnowledgeBaseJobResponseBody
	GetMaxRunningTimeInSeconds() *int32
	SetPipelineRunInfo(v *GetKnowledgeBaseJobResponseBodyPipelineRunInfo) *GetKnowledgeBaseJobResponseBody
	GetPipelineRunInfo() *GetKnowledgeBaseJobResponseBodyPipelineRunInfo
	SetRequestId(v string) *GetKnowledgeBaseJobResponseBody
	GetRequestId() *string
	SetResourceId(v string) *GetKnowledgeBaseJobResponseBody
	GetResourceId() *string
	SetStatus(v string) *GetKnowledgeBaseJobResponseBody
	GetStatus() *string
	SetUserVpc(v *GetKnowledgeBaseJobResponseBodyUserVpc) *GetKnowledgeBaseJobResponseBody
	GetUserVpc() *GetKnowledgeBaseJobResponseBodyUserVpc
	SetWorkspaceId(v string) *GetKnowledgeBaseJobResponseBody
	GetWorkspaceId() *string
}

type GetKnowledgeBaseJobResponseBody struct {
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Creator       *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// project name pass the check
	Description     *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsSpecs        []*GetKnowledgeBaseJobResponseBodyEcsSpecs      `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	EmbeddingConfig *GetKnowledgeBaseJobResponseBodyEmbeddingConfig `json:"EmbeddingConfig,omitempty" xml:"EmbeddingConfig,omitempty" type:"Struct"`
	ErrorMessage    *string                                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 2025-09-24T07:33:53Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2025-02-08T15:45:12Z
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// example:
	//
	// 2025-06-11T08:58:35.438Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// UpdateScheduleConfig
	JobAction *string `json:"JobAction,omitempty" xml:"JobAction,omitempty"`
	// example:
	//
	// d-nacr******sxd2
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// example:
	//
	// kbjob-9m******54
	KnowledgeBaseJobId     *string                                                `json:"KnowledgeBaseJobId,omitempty" xml:"KnowledgeBaseJobId,omitempty"`
	KnowledgeBaseJobResult *GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult `json:"KnowledgeBaseJobResult,omitempty" xml:"KnowledgeBaseJobResult,omitempty" type:"Struct"`
	// example:
	//
	// 86400
	MaxRunningTimeInSeconds *int32                                          `json:"MaxRunningTimeInSeconds,omitempty" xml:"MaxRunningTimeInSeconds,omitempty"`
	PipelineRunInfo         *GetKnowledgeBaseJobResponseBodyPipelineRunInfo `json:"PipelineRunInfo,omitempty" xml:"PipelineRunInfo,omitempty" type:"Struct"`
	// example:
	//
	// C25324E3-18E6-50D8-9026-16D74AAEEB26
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// discovering
	Status  *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	UserVpc *GetKnowledgeBaseJobResponseBodyUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// example:
	//
	// 478***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetKnowledgeBaseJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseJobResponseBody) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetKnowledgeBaseJobResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetKnowledgeBaseJobResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetKnowledgeBaseJobResponseBody) GetEcsSpecs() []*GetKnowledgeBaseJobResponseBodyEcsSpecs {
	return s.EcsSpecs
}

func (s *GetKnowledgeBaseJobResponseBody) GetEmbeddingConfig() *GetKnowledgeBaseJobResponseBodyEmbeddingConfig {
	return s.EmbeddingConfig
}

func (s *GetKnowledgeBaseJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetKnowledgeBaseJobResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetKnowledgeBaseJobResponseBody) GetGmtFinishTime() *string {
	return s.GmtFinishTime
}

func (s *GetKnowledgeBaseJobResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetKnowledgeBaseJobResponseBody) GetJobAction() *string {
	return s.JobAction
}

func (s *GetKnowledgeBaseJobResponseBody) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgeBaseJobResponseBody) GetKnowledgeBaseJobId() *string {
	return s.KnowledgeBaseJobId
}

func (s *GetKnowledgeBaseJobResponseBody) GetKnowledgeBaseJobResult() *GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult {
	return s.KnowledgeBaseJobResult
}

func (s *GetKnowledgeBaseJobResponseBody) GetMaxRunningTimeInSeconds() *int32 {
	return s.MaxRunningTimeInSeconds
}

func (s *GetKnowledgeBaseJobResponseBody) GetPipelineRunInfo() *GetKnowledgeBaseJobResponseBodyPipelineRunInfo {
	return s.PipelineRunInfo
}

func (s *GetKnowledgeBaseJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKnowledgeBaseJobResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetKnowledgeBaseJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetKnowledgeBaseJobResponseBody) GetUserVpc() *GetKnowledgeBaseJobResponseBodyUserVpc {
	return s.UserVpc
}

func (s *GetKnowledgeBaseJobResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetKnowledgeBaseJobResponseBody) SetAccessibility(v string) *GetKnowledgeBaseJobResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetCreator(v string) *GetKnowledgeBaseJobResponseBody {
	s.Creator = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetDescription(v string) *GetKnowledgeBaseJobResponseBody {
	s.Description = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetEcsSpecs(v []*GetKnowledgeBaseJobResponseBodyEcsSpecs) *GetKnowledgeBaseJobResponseBody {
	s.EcsSpecs = v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetEmbeddingConfig(v *GetKnowledgeBaseJobResponseBodyEmbeddingConfig) *GetKnowledgeBaseJobResponseBody {
	s.EmbeddingConfig = v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetErrorMessage(v string) *GetKnowledgeBaseJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetGmtCreateTime(v string) *GetKnowledgeBaseJobResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetGmtFinishTime(v string) *GetKnowledgeBaseJobResponseBody {
	s.GmtFinishTime = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetGmtModifiedTime(v string) *GetKnowledgeBaseJobResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetJobAction(v string) *GetKnowledgeBaseJobResponseBody {
	s.JobAction = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetKnowledgeBaseId(v string) *GetKnowledgeBaseJobResponseBody {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetKnowledgeBaseJobId(v string) *GetKnowledgeBaseJobResponseBody {
	s.KnowledgeBaseJobId = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetKnowledgeBaseJobResult(v *GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult) *GetKnowledgeBaseJobResponseBody {
	s.KnowledgeBaseJobResult = v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetMaxRunningTimeInSeconds(v int32) *GetKnowledgeBaseJobResponseBody {
	s.MaxRunningTimeInSeconds = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetPipelineRunInfo(v *GetKnowledgeBaseJobResponseBodyPipelineRunInfo) *GetKnowledgeBaseJobResponseBody {
	s.PipelineRunInfo = v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetRequestId(v string) *GetKnowledgeBaseJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetResourceId(v string) *GetKnowledgeBaseJobResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetStatus(v string) *GetKnowledgeBaseJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetUserVpc(v *GetKnowledgeBaseJobResponseBodyUserVpc) *GetKnowledgeBaseJobResponseBody {
	s.UserVpc = v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) SetWorkspaceId(v string) *GetKnowledgeBaseJobResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBody) Validate() error {
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
	if s.KnowledgeBaseJobResult != nil {
		if err := s.KnowledgeBaseJobResult.Validate(); err != nil {
			return err
		}
	}
	if s.PipelineRunInfo != nil {
		if err := s.PipelineRunInfo.Validate(); err != nil {
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

type GetKnowledgeBaseJobResponseBodyEcsSpecs struct {
	// CPU核数
	CPU *int32 `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// 驱动版本
	//
	// example:
	//
	// 535.161.08
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

func (s GetKnowledgeBaseJobResponseBodyEcsSpecs) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseJobResponseBodyEcsSpecs) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) GetCPU() *int32 {
	return s.CPU
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) GetDriver() *string {
	return s.Driver
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) GetGPU() *int32 {
	return s.GPU
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) GetGPUType() *string {
	return s.GPUType
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) GetMemory() *int32 {
	return s.Memory
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) GetPodCount() *int32 {
	return s.PodCount
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) GetSharedMemory() *int32 {
	return s.SharedMemory
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) GetType() *string {
	return s.Type
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) SetCPU(v int32) *GetKnowledgeBaseJobResponseBodyEcsSpecs {
	s.CPU = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) SetDriver(v string) *GetKnowledgeBaseJobResponseBodyEcsSpecs {
	s.Driver = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) SetGPU(v int32) *GetKnowledgeBaseJobResponseBodyEcsSpecs {
	s.GPU = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) SetGPUType(v string) *GetKnowledgeBaseJobResponseBodyEcsSpecs {
	s.GPUType = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) SetInstanceType(v string) *GetKnowledgeBaseJobResponseBodyEcsSpecs {
	s.InstanceType = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) SetMemory(v int32) *GetKnowledgeBaseJobResponseBodyEcsSpecs {
	s.Memory = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) SetPodCount(v int32) *GetKnowledgeBaseJobResponseBodyEcsSpecs {
	s.PodCount = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) SetSharedMemory(v int32) *GetKnowledgeBaseJobResponseBodyEcsSpecs {
	s.SharedMemory = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) SetType(v string) *GetKnowledgeBaseJobResponseBodyEcsSpecs {
	s.Type = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyEcsSpecs) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseJobResponseBodyEmbeddingConfig struct {
	// Embedding分批大小
	//
	// example:
	//
	// 10
	BatchSize *int32 `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// Embedding并发数
	//
	// example:
	//
	// 2
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
}

func (s GetKnowledgeBaseJobResponseBodyEmbeddingConfig) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseJobResponseBodyEmbeddingConfig) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseJobResponseBodyEmbeddingConfig) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *GetKnowledgeBaseJobResponseBodyEmbeddingConfig) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *GetKnowledgeBaseJobResponseBodyEmbeddingConfig) SetBatchSize(v int32) *GetKnowledgeBaseJobResponseBodyEmbeddingConfig {
	s.BatchSize = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyEmbeddingConfig) SetConcurrency(v int32) *GetKnowledgeBaseJobResponseBodyEmbeddingConfig {
	s.Concurrency = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyEmbeddingConfig) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult struct {
	// 增加Chunk数量
	AddChunkCount *int32 `json:"AddChunkCount,omitempty" xml:"AddChunkCount,omitempty"`
	// 删除Chunk数量
	DeleteChunkCount *int32 `json:"DeleteChunkCount,omitempty" xml:"DeleteChunkCount,omitempty"`
	// 总处理文件数
	TotalFileCount *int32 `json:"TotalFileCount,omitempty" xml:"TotalFileCount,omitempty"`
}

func (s GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult) GetAddChunkCount() *int32 {
	return s.AddChunkCount
}

func (s *GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult) GetDeleteChunkCount() *int32 {
	return s.DeleteChunkCount
}

func (s *GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult) GetTotalFileCount() *int32 {
	return s.TotalFileCount
}

func (s *GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult) SetAddChunkCount(v int32) *GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult {
	s.AddChunkCount = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult) SetDeleteChunkCount(v int32) *GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult {
	s.DeleteChunkCount = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult) SetTotalFileCount(v int32) *GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult {
	s.TotalFileCount = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyKnowledgeBaseJobResult) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseJobResponseBodyPipelineRunInfo struct {
	// PaiFlow工作流运行ID
	//
	// example:
	//
	// flow-fi8z******g4gy
	PipelineRunId *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
}

func (s GetKnowledgeBaseJobResponseBodyPipelineRunInfo) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseJobResponseBodyPipelineRunInfo) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseJobResponseBodyPipelineRunInfo) GetPipelineRunId() *string {
	return s.PipelineRunId
}

func (s *GetKnowledgeBaseJobResponseBodyPipelineRunInfo) SetPipelineRunId(v string) *GetKnowledgeBaseJobResponseBodyPipelineRunInfo {
	s.PipelineRunId = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyPipelineRunInfo) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseJobResponseBodyUserVpc struct {
	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 交换机ID
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetKnowledgeBaseJobResponseBodyUserVpc) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseJobResponseBodyUserVpc) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseJobResponseBodyUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetKnowledgeBaseJobResponseBodyUserVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetKnowledgeBaseJobResponseBodyUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *GetKnowledgeBaseJobResponseBodyUserVpc) SetSecurityGroupId(v string) *GetKnowledgeBaseJobResponseBodyUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyUserVpc) SetVSwitchId(v string) *GetKnowledgeBaseJobResponseBodyUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyUserVpc) SetVpcId(v string) *GetKnowledgeBaseJobResponseBodyUserVpc {
	s.VpcId = &v
	return s
}

func (s *GetKnowledgeBaseJobResponseBodyUserVpc) Validate() error {
	return dara.Validate(s)
}
