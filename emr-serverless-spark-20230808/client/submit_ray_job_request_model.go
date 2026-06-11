// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitRayJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActiveDeadlineSeconds(v int32) *SubmitRayJobRequest
	GetActiveDeadlineSeconds() *int32
	SetDisplayReleaseVersion(v string) *SubmitRayJobRequest
	GetDisplayReleaseVersion() *string
	SetEntrypoint(v string) *SubmitRayJobRequest
	GetEntrypoint() *string
	SetEntrypointMemory(v string) *SubmitRayJobRequest
	GetEntrypointMemory() *string
	SetEntrypointNumCpus(v string) *SubmitRayJobRequest
	GetEntrypointNumCpus() *string
	SetEntrypointNumGpus(v string) *SubmitRayJobRequest
	GetEntrypointNumGpus() *string
	SetEntrypointResources(v string) *SubmitRayJobRequest
	GetEntrypointResources() *string
	SetExtraParam(v string) *SubmitRayJobRequest
	GetExtraParam() *string
	SetHeadSpec(v *SubmitRayJobRequestHeadSpec) *SubmitRayJobRequest
	GetHeadSpec() *SubmitRayJobRequestHeadSpec
	SetMetadataJson(v string) *SubmitRayJobRequest
	GetMetadataJson() *string
	SetName(v string) *SubmitRayJobRequest
	GetName() *string
	SetNetworkServiceName(v string) *SubmitRayJobRequest
	GetNetworkServiceName() *string
	SetRuntimeEnvJson(v string) *SubmitRayJobRequest
	GetRuntimeEnvJson() *string
	SetShutdownAfterJobFinishes(v bool) *SubmitRayJobRequest
	GetShutdownAfterJobFinishes() *bool
	SetSubmissionMode(v string) *SubmitRayJobRequest
	GetSubmissionMode() *string
	SetTags(v []*SubmitRayJobRequestTags) *SubmitRayJobRequest
	GetTags() []*SubmitRayJobRequestTags
	SetTtlSecondsAfterFinished(v int32) *SubmitRayJobRequest
	GetTtlSecondsAfterFinished() *int32
	SetVolumeIds(v []*string) *SubmitRayJobRequest
	GetVolumeIds() []*string
	SetWorkerSpec(v []*SubmitRayJobRequestWorkerSpec) *SubmitRayJobRequest
	GetWorkerSpec() []*SubmitRayJobRequestWorkerSpec
	SetWorkingDir(v string) *SubmitRayJobRequest
	GetWorkingDir() *string
}

type SubmitRayJobRequest struct {
	// example:
	//
	// 3600
	ActiveDeadlineSeconds *int32 `json:"activeDeadlineSeconds,omitempty" xml:"activeDeadlineSeconds,omitempty"`
	// example:
	//
	// err-1.2.0 (Ray 2.55.1, Python 3.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// example:
	//
	// python -c "print(\\"hello ray job\\")"
	Entrypoint *string `json:"entrypoint,omitempty" xml:"entrypoint,omitempty"`
	// example:
	//
	// 4Gi
	EntrypointMemory *string `json:"entrypointMemory,omitempty" xml:"entrypointMemory,omitempty"`
	// example:
	//
	// 1
	EntrypointNumCpus *string `json:"entrypointNumCpus,omitempty" xml:"entrypointNumCpus,omitempty"`
	// example:
	//
	// 0
	EntrypointNumGpus *string `json:"entrypointNumGpus,omitempty" xml:"entrypointNumGpus,omitempty"`
	// example:
	//
	// {"fpu": 1}
	EntrypointResources *string `json:"entrypointResources,omitempty" xml:"entrypointResources,omitempty"`
	// example:
	//
	// {"userDefinedFiles": "oss://mybucket/artifact/config.json,oss://mybucket/artifact/config2.json", "userRequirementsFile": "oss://mybucket/requirements.txt"}
	ExtraParam *string                      `json:"extraParam,omitempty" xml:"extraParam,omitempty"`
	HeadSpec   *SubmitRayJobRequestHeadSpec `json:"headSpec,omitempty" xml:"headSpec,omitempty" type:"Struct"`
	// example:
	//
	// {"resourceName": "test"}
	MetadataJson *string `json:"metadataJson,omitempty" xml:"metadataJson,omitempty"`
	// example:
	//
	// my-job
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// vpc
	NetworkServiceName *string `json:"networkServiceName,omitempty" xml:"networkServiceName,omitempty"`
	// example:
	//
	// {"pip":["requests==2.26.0","pendulum==2.1.2"],"env_vars":{"KEY":"VALUE"}}
	RuntimeEnvJson *string `json:"runtimeEnvJson,omitempty" xml:"runtimeEnvJson,omitempty"`
	// example:
	//
	// true
	ShutdownAfterJobFinishes *bool `json:"shutdownAfterJobFinishes,omitempty" xml:"shutdownAfterJobFinishes,omitempty"`
	// example:
	//
	// HTTPMode
	SubmissionMode *string                    `json:"submissionMode,omitempty" xml:"submissionMode,omitempty"`
	Tags           []*SubmitRayJobRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 60
	TtlSecondsAfterFinished *int32                           `json:"ttlSecondsAfterFinished,omitempty" xml:"ttlSecondsAfterFinished,omitempty"`
	VolumeIds               []*string                        `json:"volumeIds,omitempty" xml:"volumeIds,omitempty" type:"Repeated"`
	WorkerSpec              []*SubmitRayJobRequestWorkerSpec `json:"workerSpec,omitempty" xml:"workerSpec,omitempty" type:"Repeated"`
	// example:
	//
	// oss://mybucket/rayjob.zip
	WorkingDir *string `json:"workingDir,omitempty" xml:"workingDir,omitempty"`
}

func (s SubmitRayJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitRayJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitRayJobRequest) GetActiveDeadlineSeconds() *int32 {
	return s.ActiveDeadlineSeconds
}

func (s *SubmitRayJobRequest) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *SubmitRayJobRequest) GetEntrypoint() *string {
	return s.Entrypoint
}

func (s *SubmitRayJobRequest) GetEntrypointMemory() *string {
	return s.EntrypointMemory
}

func (s *SubmitRayJobRequest) GetEntrypointNumCpus() *string {
	return s.EntrypointNumCpus
}

func (s *SubmitRayJobRequest) GetEntrypointNumGpus() *string {
	return s.EntrypointNumGpus
}

func (s *SubmitRayJobRequest) GetEntrypointResources() *string {
	return s.EntrypointResources
}

func (s *SubmitRayJobRequest) GetExtraParam() *string {
	return s.ExtraParam
}

func (s *SubmitRayJobRequest) GetHeadSpec() *SubmitRayJobRequestHeadSpec {
	return s.HeadSpec
}

func (s *SubmitRayJobRequest) GetMetadataJson() *string {
	return s.MetadataJson
}

func (s *SubmitRayJobRequest) GetName() *string {
	return s.Name
}

func (s *SubmitRayJobRequest) GetNetworkServiceName() *string {
	return s.NetworkServiceName
}

func (s *SubmitRayJobRequest) GetRuntimeEnvJson() *string {
	return s.RuntimeEnvJson
}

func (s *SubmitRayJobRequest) GetShutdownAfterJobFinishes() *bool {
	return s.ShutdownAfterJobFinishes
}

func (s *SubmitRayJobRequest) GetSubmissionMode() *string {
	return s.SubmissionMode
}

func (s *SubmitRayJobRequest) GetTags() []*SubmitRayJobRequestTags {
	return s.Tags
}

func (s *SubmitRayJobRequest) GetTtlSecondsAfterFinished() *int32 {
	return s.TtlSecondsAfterFinished
}

func (s *SubmitRayJobRequest) GetVolumeIds() []*string {
	return s.VolumeIds
}

func (s *SubmitRayJobRequest) GetWorkerSpec() []*SubmitRayJobRequestWorkerSpec {
	return s.WorkerSpec
}

func (s *SubmitRayJobRequest) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *SubmitRayJobRequest) SetActiveDeadlineSeconds(v int32) *SubmitRayJobRequest {
	s.ActiveDeadlineSeconds = &v
	return s
}

func (s *SubmitRayJobRequest) SetDisplayReleaseVersion(v string) *SubmitRayJobRequest {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *SubmitRayJobRequest) SetEntrypoint(v string) *SubmitRayJobRequest {
	s.Entrypoint = &v
	return s
}

func (s *SubmitRayJobRequest) SetEntrypointMemory(v string) *SubmitRayJobRequest {
	s.EntrypointMemory = &v
	return s
}

func (s *SubmitRayJobRequest) SetEntrypointNumCpus(v string) *SubmitRayJobRequest {
	s.EntrypointNumCpus = &v
	return s
}

func (s *SubmitRayJobRequest) SetEntrypointNumGpus(v string) *SubmitRayJobRequest {
	s.EntrypointNumGpus = &v
	return s
}

func (s *SubmitRayJobRequest) SetEntrypointResources(v string) *SubmitRayJobRequest {
	s.EntrypointResources = &v
	return s
}

func (s *SubmitRayJobRequest) SetExtraParam(v string) *SubmitRayJobRequest {
	s.ExtraParam = &v
	return s
}

func (s *SubmitRayJobRequest) SetHeadSpec(v *SubmitRayJobRequestHeadSpec) *SubmitRayJobRequest {
	s.HeadSpec = v
	return s
}

func (s *SubmitRayJobRequest) SetMetadataJson(v string) *SubmitRayJobRequest {
	s.MetadataJson = &v
	return s
}

func (s *SubmitRayJobRequest) SetName(v string) *SubmitRayJobRequest {
	s.Name = &v
	return s
}

func (s *SubmitRayJobRequest) SetNetworkServiceName(v string) *SubmitRayJobRequest {
	s.NetworkServiceName = &v
	return s
}

func (s *SubmitRayJobRequest) SetRuntimeEnvJson(v string) *SubmitRayJobRequest {
	s.RuntimeEnvJson = &v
	return s
}

func (s *SubmitRayJobRequest) SetShutdownAfterJobFinishes(v bool) *SubmitRayJobRequest {
	s.ShutdownAfterJobFinishes = &v
	return s
}

func (s *SubmitRayJobRequest) SetSubmissionMode(v string) *SubmitRayJobRequest {
	s.SubmissionMode = &v
	return s
}

func (s *SubmitRayJobRequest) SetTags(v []*SubmitRayJobRequestTags) *SubmitRayJobRequest {
	s.Tags = v
	return s
}

func (s *SubmitRayJobRequest) SetTtlSecondsAfterFinished(v int32) *SubmitRayJobRequest {
	s.TtlSecondsAfterFinished = &v
	return s
}

func (s *SubmitRayJobRequest) SetVolumeIds(v []*string) *SubmitRayJobRequest {
	s.VolumeIds = v
	return s
}

func (s *SubmitRayJobRequest) SetWorkerSpec(v []*SubmitRayJobRequestWorkerSpec) *SubmitRayJobRequest {
	s.WorkerSpec = v
	return s
}

func (s *SubmitRayJobRequest) SetWorkingDir(v string) *SubmitRayJobRequest {
	s.WorkingDir = &v
	return s
}

func (s *SubmitRayJobRequest) Validate() error {
	if s.HeadSpec != nil {
		if err := s.HeadSpec.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WorkerSpec != nil {
		for _, item := range s.WorkerSpec {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitRayJobRequestHeadSpec struct {
	// example:
	//
	// 2
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// true
	EnableAutoScaling *bool `json:"enableAutoScaling,omitempty" xml:"enableAutoScaling,omitempty"`
	// example:
	//
	// ecs.gn6i-c4g1.xlarge
	GpuSpec *string `json:"gpuSpec,omitempty" xml:"gpuSpec,omitempty"`
	// example:
	//
	// 60
	IdleTimeoutSeconds *int32 `json:"idleTimeoutSeconds,omitempty" xml:"idleTimeoutSeconds,omitempty"`
	// example:
	//
	// 8Gi
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
}

func (s SubmitRayJobRequestHeadSpec) String() string {
	return dara.Prettify(s)
}

func (s SubmitRayJobRequestHeadSpec) GoString() string {
	return s.String()
}

func (s *SubmitRayJobRequestHeadSpec) GetCpu() *string {
	return s.Cpu
}

func (s *SubmitRayJobRequestHeadSpec) GetEnableAutoScaling() *bool {
	return s.EnableAutoScaling
}

func (s *SubmitRayJobRequestHeadSpec) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *SubmitRayJobRequestHeadSpec) GetIdleTimeoutSeconds() *int32 {
	return s.IdleTimeoutSeconds
}

func (s *SubmitRayJobRequestHeadSpec) GetMemory() *string {
	return s.Memory
}

func (s *SubmitRayJobRequestHeadSpec) GetQueueName() *string {
	return s.QueueName
}

func (s *SubmitRayJobRequestHeadSpec) SetCpu(v string) *SubmitRayJobRequestHeadSpec {
	s.Cpu = &v
	return s
}

func (s *SubmitRayJobRequestHeadSpec) SetEnableAutoScaling(v bool) *SubmitRayJobRequestHeadSpec {
	s.EnableAutoScaling = &v
	return s
}

func (s *SubmitRayJobRequestHeadSpec) SetGpuSpec(v string) *SubmitRayJobRequestHeadSpec {
	s.GpuSpec = &v
	return s
}

func (s *SubmitRayJobRequestHeadSpec) SetIdleTimeoutSeconds(v int32) *SubmitRayJobRequestHeadSpec {
	s.IdleTimeoutSeconds = &v
	return s
}

func (s *SubmitRayJobRequestHeadSpec) SetMemory(v string) *SubmitRayJobRequestHeadSpec {
	s.Memory = &v
	return s
}

func (s *SubmitRayJobRequestHeadSpec) SetQueueName(v string) *SubmitRayJobRequestHeadSpec {
	s.QueueName = &v
	return s
}

func (s *SubmitRayJobRequestHeadSpec) Validate() error {
	return dara.Validate(s)
}

type SubmitRayJobRequestTags struct {
	// example:
	//
	// jobname
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SubmitRayJobRequestTags) String() string {
	return dara.Prettify(s)
}

func (s SubmitRayJobRequestTags) GoString() string {
	return s.String()
}

func (s *SubmitRayJobRequestTags) GetKey() *string {
	return s.Key
}

func (s *SubmitRayJobRequestTags) GetValue() *string {
	return s.Value
}

func (s *SubmitRayJobRequestTags) SetKey(v string) *SubmitRayJobRequestTags {
	s.Key = &v
	return s
}

func (s *SubmitRayJobRequestTags) SetValue(v string) *SubmitRayJobRequestTags {
	s.Value = &v
	return s
}

func (s *SubmitRayJobRequestTags) Validate() error {
	return dara.Validate(s)
}

type SubmitRayJobRequestWorkerSpec struct {
	// example:
	//
	// 4
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// ecs.gn6i-c4g1.xlarge
	GpuSpec *string `json:"gpuSpec,omitempty" xml:"gpuSpec,omitempty"`
	// example:
	//
	// WorkerGroup1
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// 10
	MaxReplica *int32 `json:"maxReplica,omitempty" xml:"maxReplica,omitempty"`
	// example:
	//
	// 16Gi
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
	// example:
	//
	// 1
	MinReplica *int32 `json:"minReplica,omitempty" xml:"minReplica,omitempty"`
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// example:
	//
	// 2
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
	// example:
	//
	// CPU
	WorkerType *string `json:"workerType,omitempty" xml:"workerType,omitempty"`
}

func (s SubmitRayJobRequestWorkerSpec) String() string {
	return dara.Prettify(s)
}

func (s SubmitRayJobRequestWorkerSpec) GoString() string {
	return s.String()
}

func (s *SubmitRayJobRequestWorkerSpec) GetCpu() *string {
	return s.Cpu
}

func (s *SubmitRayJobRequestWorkerSpec) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *SubmitRayJobRequestWorkerSpec) GetGroupName() *string {
	return s.GroupName
}

func (s *SubmitRayJobRequestWorkerSpec) GetMaxReplica() *int32 {
	return s.MaxReplica
}

func (s *SubmitRayJobRequestWorkerSpec) GetMemory() *string {
	return s.Memory
}

func (s *SubmitRayJobRequestWorkerSpec) GetMinReplica() *int32 {
	return s.MinReplica
}

func (s *SubmitRayJobRequestWorkerSpec) GetQueueName() *string {
	return s.QueueName
}

func (s *SubmitRayJobRequestWorkerSpec) GetReplica() *int32 {
	return s.Replica
}

func (s *SubmitRayJobRequestWorkerSpec) GetWorkerType() *string {
	return s.WorkerType
}

func (s *SubmitRayJobRequestWorkerSpec) SetCpu(v string) *SubmitRayJobRequestWorkerSpec {
	s.Cpu = &v
	return s
}

func (s *SubmitRayJobRequestWorkerSpec) SetGpuSpec(v string) *SubmitRayJobRequestWorkerSpec {
	s.GpuSpec = &v
	return s
}

func (s *SubmitRayJobRequestWorkerSpec) SetGroupName(v string) *SubmitRayJobRequestWorkerSpec {
	s.GroupName = &v
	return s
}

func (s *SubmitRayJobRequestWorkerSpec) SetMaxReplica(v int32) *SubmitRayJobRequestWorkerSpec {
	s.MaxReplica = &v
	return s
}

func (s *SubmitRayJobRequestWorkerSpec) SetMemory(v string) *SubmitRayJobRequestWorkerSpec {
	s.Memory = &v
	return s
}

func (s *SubmitRayJobRequestWorkerSpec) SetMinReplica(v int32) *SubmitRayJobRequestWorkerSpec {
	s.MinReplica = &v
	return s
}

func (s *SubmitRayJobRequestWorkerSpec) SetQueueName(v string) *SubmitRayJobRequestWorkerSpec {
	s.QueueName = &v
	return s
}

func (s *SubmitRayJobRequestWorkerSpec) SetReplica(v int32) *SubmitRayJobRequestWorkerSpec {
	s.Replica = &v
	return s
}

func (s *SubmitRayJobRequestWorkerSpec) SetWorkerType(v string) *SubmitRayJobRequestWorkerSpec {
	s.WorkerType = &v
	return s
}

func (s *SubmitRayJobRequestWorkerSpec) Validate() error {
	return dara.Validate(s)
}
