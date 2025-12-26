// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeBaseJob interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *KnowledgeBaseJob
	GetAccessibility() *string
	SetCreator(v string) *KnowledgeBaseJob
	GetCreator() *string
	SetDescription(v string) *KnowledgeBaseJob
	GetDescription() *string
	SetEcsSpecs(v []*KnowledgeBaseJobEcsSpecs) *KnowledgeBaseJob
	GetEcsSpecs() []*KnowledgeBaseJobEcsSpecs
	SetEmbeddingConfig(v *KnowledgeBaseJobEmbeddingConfig) *KnowledgeBaseJob
	GetEmbeddingConfig() *KnowledgeBaseJobEmbeddingConfig
	SetErrorMessage(v string) *KnowledgeBaseJob
	GetErrorMessage() *string
	SetGmtCreateTime(v string) *KnowledgeBaseJob
	GetGmtCreateTime() *string
	SetGmtFinishTime(v string) *KnowledgeBaseJob
	GetGmtFinishTime() *string
	SetGmtModifiedTime(v string) *KnowledgeBaseJob
	GetGmtModifiedTime() *string
	SetJobAction(v string) *KnowledgeBaseJob
	GetJobAction() *string
	SetKnowledgeBaseId(v string) *KnowledgeBaseJob
	GetKnowledgeBaseId() *string
	SetKnowledgeBaseJobId(v string) *KnowledgeBaseJob
	GetKnowledgeBaseJobId() *string
	SetKnowledgeBaseJobResult(v *KnowledgeBaseJobKnowledgeBaseJobResult) *KnowledgeBaseJob
	GetKnowledgeBaseJobResult() *KnowledgeBaseJobKnowledgeBaseJobResult
	SetMaxRunningTimeInSeconds(v int32) *KnowledgeBaseJob
	GetMaxRunningTimeInSeconds() *int32
	SetPipelineRunInfo(v *KnowledgeBaseJobPipelineRunInfo) *KnowledgeBaseJob
	GetPipelineRunInfo() *KnowledgeBaseJobPipelineRunInfo
	SetResourceId(v string) *KnowledgeBaseJob
	GetResourceId() *string
	SetStatus(v string) *KnowledgeBaseJob
	GetStatus() *string
	SetUserVpc(v *KnowledgeBaseJobUserVpc) *KnowledgeBaseJob
	GetUserVpc() *KnowledgeBaseJobUserVpc
	SetWorkspaceId(v string) *KnowledgeBaseJob
	GetWorkspaceId() *string
}

type KnowledgeBaseJob struct {
	Accessibility           *string                                 `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Creator                 *string                                 `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description             *string                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsSpecs                []*KnowledgeBaseJobEcsSpecs             `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	EmbeddingConfig         *KnowledgeBaseJobEmbeddingConfig        `json:"EmbeddingConfig,omitempty" xml:"EmbeddingConfig,omitempty" type:"Struct"`
	ErrorMessage            *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GmtCreateTime           *string                                 `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtFinishTime           *string                                 `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	GmtModifiedTime         *string                                 `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	JobAction               *string                                 `json:"JobAction,omitempty" xml:"JobAction,omitempty"`
	KnowledgeBaseId         *string                                 `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	KnowledgeBaseJobId      *string                                 `json:"KnowledgeBaseJobId,omitempty" xml:"KnowledgeBaseJobId,omitempty"`
	KnowledgeBaseJobResult  *KnowledgeBaseJobKnowledgeBaseJobResult `json:"KnowledgeBaseJobResult,omitempty" xml:"KnowledgeBaseJobResult,omitempty" type:"Struct"`
	MaxRunningTimeInSeconds *int32                                  `json:"MaxRunningTimeInSeconds,omitempty" xml:"MaxRunningTimeInSeconds,omitempty"`
	PipelineRunInfo         *KnowledgeBaseJobPipelineRunInfo        `json:"PipelineRunInfo,omitempty" xml:"PipelineRunInfo,omitempty" type:"Struct"`
	ResourceId              *string                                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Status                  *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	UserVpc                 *KnowledgeBaseJobUserVpc                `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	WorkspaceId             *string                                 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s KnowledgeBaseJob) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseJob) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseJob) GetAccessibility() *string {
	return s.Accessibility
}

func (s *KnowledgeBaseJob) GetCreator() *string {
	return s.Creator
}

func (s *KnowledgeBaseJob) GetDescription() *string {
	return s.Description
}

func (s *KnowledgeBaseJob) GetEcsSpecs() []*KnowledgeBaseJobEcsSpecs {
	return s.EcsSpecs
}

func (s *KnowledgeBaseJob) GetEmbeddingConfig() *KnowledgeBaseJobEmbeddingConfig {
	return s.EmbeddingConfig
}

func (s *KnowledgeBaseJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *KnowledgeBaseJob) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *KnowledgeBaseJob) GetGmtFinishTime() *string {
	return s.GmtFinishTime
}

func (s *KnowledgeBaseJob) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *KnowledgeBaseJob) GetJobAction() *string {
	return s.JobAction
}

func (s *KnowledgeBaseJob) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *KnowledgeBaseJob) GetKnowledgeBaseJobId() *string {
	return s.KnowledgeBaseJobId
}

func (s *KnowledgeBaseJob) GetKnowledgeBaseJobResult() *KnowledgeBaseJobKnowledgeBaseJobResult {
	return s.KnowledgeBaseJobResult
}

func (s *KnowledgeBaseJob) GetMaxRunningTimeInSeconds() *int32 {
	return s.MaxRunningTimeInSeconds
}

func (s *KnowledgeBaseJob) GetPipelineRunInfo() *KnowledgeBaseJobPipelineRunInfo {
	return s.PipelineRunInfo
}

func (s *KnowledgeBaseJob) GetResourceId() *string {
	return s.ResourceId
}

func (s *KnowledgeBaseJob) GetStatus() *string {
	return s.Status
}

func (s *KnowledgeBaseJob) GetUserVpc() *KnowledgeBaseJobUserVpc {
	return s.UserVpc
}

func (s *KnowledgeBaseJob) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *KnowledgeBaseJob) SetAccessibility(v string) *KnowledgeBaseJob {
	s.Accessibility = &v
	return s
}

func (s *KnowledgeBaseJob) SetCreator(v string) *KnowledgeBaseJob {
	s.Creator = &v
	return s
}

func (s *KnowledgeBaseJob) SetDescription(v string) *KnowledgeBaseJob {
	s.Description = &v
	return s
}

func (s *KnowledgeBaseJob) SetEcsSpecs(v []*KnowledgeBaseJobEcsSpecs) *KnowledgeBaseJob {
	s.EcsSpecs = v
	return s
}

func (s *KnowledgeBaseJob) SetEmbeddingConfig(v *KnowledgeBaseJobEmbeddingConfig) *KnowledgeBaseJob {
	s.EmbeddingConfig = v
	return s
}

func (s *KnowledgeBaseJob) SetErrorMessage(v string) *KnowledgeBaseJob {
	s.ErrorMessage = &v
	return s
}

func (s *KnowledgeBaseJob) SetGmtCreateTime(v string) *KnowledgeBaseJob {
	s.GmtCreateTime = &v
	return s
}

func (s *KnowledgeBaseJob) SetGmtFinishTime(v string) *KnowledgeBaseJob {
	s.GmtFinishTime = &v
	return s
}

func (s *KnowledgeBaseJob) SetGmtModifiedTime(v string) *KnowledgeBaseJob {
	s.GmtModifiedTime = &v
	return s
}

func (s *KnowledgeBaseJob) SetJobAction(v string) *KnowledgeBaseJob {
	s.JobAction = &v
	return s
}

func (s *KnowledgeBaseJob) SetKnowledgeBaseId(v string) *KnowledgeBaseJob {
	s.KnowledgeBaseId = &v
	return s
}

func (s *KnowledgeBaseJob) SetKnowledgeBaseJobId(v string) *KnowledgeBaseJob {
	s.KnowledgeBaseJobId = &v
	return s
}

func (s *KnowledgeBaseJob) SetKnowledgeBaseJobResult(v *KnowledgeBaseJobKnowledgeBaseJobResult) *KnowledgeBaseJob {
	s.KnowledgeBaseJobResult = v
	return s
}

func (s *KnowledgeBaseJob) SetMaxRunningTimeInSeconds(v int32) *KnowledgeBaseJob {
	s.MaxRunningTimeInSeconds = &v
	return s
}

func (s *KnowledgeBaseJob) SetPipelineRunInfo(v *KnowledgeBaseJobPipelineRunInfo) *KnowledgeBaseJob {
	s.PipelineRunInfo = v
	return s
}

func (s *KnowledgeBaseJob) SetResourceId(v string) *KnowledgeBaseJob {
	s.ResourceId = &v
	return s
}

func (s *KnowledgeBaseJob) SetStatus(v string) *KnowledgeBaseJob {
	s.Status = &v
	return s
}

func (s *KnowledgeBaseJob) SetUserVpc(v *KnowledgeBaseJobUserVpc) *KnowledgeBaseJob {
	s.UserVpc = v
	return s
}

func (s *KnowledgeBaseJob) SetWorkspaceId(v string) *KnowledgeBaseJob {
	s.WorkspaceId = &v
	return s
}

func (s *KnowledgeBaseJob) Validate() error {
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

type KnowledgeBaseJobEcsSpecs struct {
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

func (s KnowledgeBaseJobEcsSpecs) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseJobEcsSpecs) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseJobEcsSpecs) GetCPU() *int32 {
	return s.CPU
}

func (s *KnowledgeBaseJobEcsSpecs) GetDriver() *string {
	return s.Driver
}

func (s *KnowledgeBaseJobEcsSpecs) GetGPU() *int32 {
	return s.GPU
}

func (s *KnowledgeBaseJobEcsSpecs) GetGPUType() *string {
	return s.GPUType
}

func (s *KnowledgeBaseJobEcsSpecs) GetInstanceType() *string {
	return s.InstanceType
}

func (s *KnowledgeBaseJobEcsSpecs) GetMemory() *int32 {
	return s.Memory
}

func (s *KnowledgeBaseJobEcsSpecs) GetPodCount() *int32 {
	return s.PodCount
}

func (s *KnowledgeBaseJobEcsSpecs) GetSharedMemory() *int32 {
	return s.SharedMemory
}

func (s *KnowledgeBaseJobEcsSpecs) GetType() *string {
	return s.Type
}

func (s *KnowledgeBaseJobEcsSpecs) SetCPU(v int32) *KnowledgeBaseJobEcsSpecs {
	s.CPU = &v
	return s
}

func (s *KnowledgeBaseJobEcsSpecs) SetDriver(v string) *KnowledgeBaseJobEcsSpecs {
	s.Driver = &v
	return s
}

func (s *KnowledgeBaseJobEcsSpecs) SetGPU(v int32) *KnowledgeBaseJobEcsSpecs {
	s.GPU = &v
	return s
}

func (s *KnowledgeBaseJobEcsSpecs) SetGPUType(v string) *KnowledgeBaseJobEcsSpecs {
	s.GPUType = &v
	return s
}

func (s *KnowledgeBaseJobEcsSpecs) SetInstanceType(v string) *KnowledgeBaseJobEcsSpecs {
	s.InstanceType = &v
	return s
}

func (s *KnowledgeBaseJobEcsSpecs) SetMemory(v int32) *KnowledgeBaseJobEcsSpecs {
	s.Memory = &v
	return s
}

func (s *KnowledgeBaseJobEcsSpecs) SetPodCount(v int32) *KnowledgeBaseJobEcsSpecs {
	s.PodCount = &v
	return s
}

func (s *KnowledgeBaseJobEcsSpecs) SetSharedMemory(v int32) *KnowledgeBaseJobEcsSpecs {
	s.SharedMemory = &v
	return s
}

func (s *KnowledgeBaseJobEcsSpecs) SetType(v string) *KnowledgeBaseJobEcsSpecs {
	s.Type = &v
	return s
}

func (s *KnowledgeBaseJobEcsSpecs) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseJobEmbeddingConfig struct {
	// Embedding分批大小
	BatchSize *int32 `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	// Embedding并发数
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
}

func (s KnowledgeBaseJobEmbeddingConfig) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseJobEmbeddingConfig) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseJobEmbeddingConfig) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *KnowledgeBaseJobEmbeddingConfig) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *KnowledgeBaseJobEmbeddingConfig) SetBatchSize(v int32) *KnowledgeBaseJobEmbeddingConfig {
	s.BatchSize = &v
	return s
}

func (s *KnowledgeBaseJobEmbeddingConfig) SetConcurrency(v int32) *KnowledgeBaseJobEmbeddingConfig {
	s.Concurrency = &v
	return s
}

func (s *KnowledgeBaseJobEmbeddingConfig) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseJobKnowledgeBaseJobResult struct {
	// 增加Chunk数量
	AddChunkCount *int32 `json:"AddChunkCount,omitempty" xml:"AddChunkCount,omitempty"`
	// 删除Chunk数量
	DeleteChunkCount *int32 `json:"DeleteChunkCount,omitempty" xml:"DeleteChunkCount,omitempty"`
	// 总处理文件数
	TotalFileCount *int32 `json:"TotalFileCount,omitempty" xml:"TotalFileCount,omitempty"`
}

func (s KnowledgeBaseJobKnowledgeBaseJobResult) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseJobKnowledgeBaseJobResult) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseJobKnowledgeBaseJobResult) GetAddChunkCount() *int32 {
	return s.AddChunkCount
}

func (s *KnowledgeBaseJobKnowledgeBaseJobResult) GetDeleteChunkCount() *int32 {
	return s.DeleteChunkCount
}

func (s *KnowledgeBaseJobKnowledgeBaseJobResult) GetTotalFileCount() *int32 {
	return s.TotalFileCount
}

func (s *KnowledgeBaseJobKnowledgeBaseJobResult) SetAddChunkCount(v int32) *KnowledgeBaseJobKnowledgeBaseJobResult {
	s.AddChunkCount = &v
	return s
}

func (s *KnowledgeBaseJobKnowledgeBaseJobResult) SetDeleteChunkCount(v int32) *KnowledgeBaseJobKnowledgeBaseJobResult {
	s.DeleteChunkCount = &v
	return s
}

func (s *KnowledgeBaseJobKnowledgeBaseJobResult) SetTotalFileCount(v int32) *KnowledgeBaseJobKnowledgeBaseJobResult {
	s.TotalFileCount = &v
	return s
}

func (s *KnowledgeBaseJobKnowledgeBaseJobResult) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseJobPipelineRunInfo struct {
	// PaiFlow工作流运行ID
	PipelineRunId *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
}

func (s KnowledgeBaseJobPipelineRunInfo) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseJobPipelineRunInfo) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseJobPipelineRunInfo) GetPipelineRunId() *string {
	return s.PipelineRunId
}

func (s *KnowledgeBaseJobPipelineRunInfo) SetPipelineRunId(v string) *KnowledgeBaseJobPipelineRunInfo {
	s.PipelineRunId = &v
	return s
}

func (s *KnowledgeBaseJobPipelineRunInfo) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseJobUserVpc struct {
	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 交换机ID
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s KnowledgeBaseJobUserVpc) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseJobUserVpc) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseJobUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *KnowledgeBaseJobUserVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *KnowledgeBaseJobUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *KnowledgeBaseJobUserVpc) SetSecurityGroupId(v string) *KnowledgeBaseJobUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *KnowledgeBaseJobUserVpc) SetVSwitchId(v string) *KnowledgeBaseJobUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *KnowledgeBaseJobUserVpc) SetVpcId(v string) *KnowledgeBaseJobUserVpc {
	s.VpcId = &v
	return s
}

func (s *KnowledgeBaseJobUserVpc) Validate() error {
	return dara.Validate(s)
}
