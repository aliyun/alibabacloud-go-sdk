// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRayJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActiveDeadlineSeconds(v int32) *GetRayJobResponseBody
	GetActiveDeadlineSeconds() *int32
	SetBackoffLimit(v int32) *GetRayJobResponseBody
	GetBackoffLimit() *int32
	SetClusterState(v string) *GetRayJobResponseBody
	GetClusterState() *string
	SetCreatorName(v string) *GetRayJobResponseBody
	GetCreatorName() *string
	SetCuHours(v float64) *GetRayJobResponseBody
	GetCuHours() *float64
	SetDashboardUrl(v string) *GetRayJobResponseBody
	GetDashboardUrl() *string
	SetDashboardUrlExtra(v []*string) *GetRayJobResponseBody
	GetDashboardUrlExtra() []*string
	SetDisplayReleaseVersion(v string) *GetRayJobResponseBody
	GetDisplayReleaseVersion() *string
	SetDuration(v int64) *GetRayJobResponseBody
	GetDuration() *int64
	SetEndTime(v int64) *GetRayJobResponseBody
	GetEndTime() *int64
	SetEntrypoint(v string) *GetRayJobResponseBody
	GetEntrypoint() *string
	SetEntrypointMemory(v string) *GetRayJobResponseBody
	GetEntrypointMemory() *string
	SetEntrypointNumCpus(v string) *GetRayJobResponseBody
	GetEntrypointNumCpus() *string
	SetEntrypointNumGpus(v string) *GetRayJobResponseBody
	GetEntrypointNumGpus() *string
	SetEntrypointResources(v string) *GetRayJobResponseBody
	GetEntrypointResources() *string
	SetExtraParam(v string) *GetRayJobResponseBody
	GetExtraParam() *string
	SetGuHours(v *GetRayJobResponseBodyGuHours) *GetRayJobResponseBody
	GetGuHours() *GetRayJobResponseBodyGuHours
	SetHeadSpec(v *GetRayJobResponseBodyHeadSpec) *GetRayJobResponseBody
	GetHeadSpec() *GetRayJobResponseBodyHeadSpec
	SetLogBucketName(v string) *GetRayJobResponseBody
	GetLogBucketName() *string
	SetLogPath(v string) *GetRayJobResponseBody
	GetLogPath() *string
	SetMessage(v string) *GetRayJobResponseBody
	GetMessage() *string
	SetMetadataJson(v string) *GetRayJobResponseBody
	GetMetadataJson() *string
	SetName(v string) *GetRayJobResponseBody
	GetName() *string
	SetNetworkServiceName(v string) *GetRayJobResponseBody
	GetNetworkServiceName() *string
	SetRequestId(v string) *GetRayJobResponseBody
	GetRequestId() *string
	SetRuntimeEnvJson(v string) *GetRayJobResponseBody
	GetRuntimeEnvJson() *string
	SetShutdownAfterJobFinishes(v bool) *GetRayJobResponseBody
	GetShutdownAfterJobFinishes() *bool
	SetStartTime(v int64) *GetRayJobResponseBody
	GetStartTime() *int64
	SetStatus(v string) *GetRayJobResponseBody
	GetStatus() *string
	SetSubmissionId(v string) *GetRayJobResponseBody
	GetSubmissionId() *string
	SetSubmissionMode(v string) *GetRayJobResponseBody
	GetSubmissionMode() *string
	SetSubmitTime(v int64) *GetRayJobResponseBody
	GetSubmitTime() *int64
	SetTags(v []*Tag) *GetRayJobResponseBody
	GetTags() []*Tag
	SetTaskBizId(v string) *GetRayJobResponseBody
	GetTaskBizId() *string
	SetTtlSecondsAfterFinished(v int32) *GetRayJobResponseBody
	GetTtlSecondsAfterFinished() *int32
	SetVolumeIds(v []*string) *GetRayJobResponseBody
	GetVolumeIds() []*string
	SetWorkerSpecs(v []*GetRayJobResponseBodyWorkerSpecs) *GetRayJobResponseBody
	GetWorkerSpecs() []*GetRayJobResponseBodyWorkerSpecs
	SetWorkingDir(v string) *GetRayJobResponseBody
	GetWorkingDir() *string
}

type GetRayJobResponseBody struct {
	// example:
	//
	// 3600
	ActiveDeadlineSeconds *int32 `json:"activeDeadlineSeconds,omitempty" xml:"activeDeadlineSeconds,omitempty"`
	// example:
	//
	// 2
	BackoffLimit *int32 `json:"backoffLimit,omitempty" xml:"backoffLimit,omitempty"`
	// example:
	//
	// Running
	ClusterState *string `json:"clusterState,omitempty" xml:"clusterState,omitempty"`
	// example:
	//
	// Alice
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// 1899
	CuHours *float64 `json:"cuHours,omitempty" xml:"cuHours,omitempty"`
	// example:
	//
	// https://emr-ray-gateway-cn-hangzhou.aliyuncs.com/workspace/w-xxxxxxxx/raycluster/ray-xxxxxx/dashboard?token=xxxxxx
	DashboardUrl      *string   `json:"dashboardUrl,omitempty" xml:"dashboardUrl,omitempty"`
	DashboardUrlExtra []*string `json:"dashboardUrlExtra,omitempty" xml:"dashboardUrlExtra,omitempty" type:"Repeated"`
	// example:
	//
	// err-1.2.0 (Ray 2.55.1, Python 3.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// example:
	//
	// 2459764
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 1762949372000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// python main.py
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
	ExtraParam *string                        `json:"extraParam,omitempty" xml:"extraParam,omitempty"`
	GuHours    *GetRayJobResponseBodyGuHours  `json:"guHours,omitempty" xml:"guHours,omitempty" type:"Struct"`
	HeadSpec   *GetRayJobResponseBodyHeadSpec `json:"headSpec,omitempty" xml:"headSpec,omitempty" type:"Struct"`
	// example:
	//
	// ss-ray-cn-hangzhou
	LogBucketName *string `json:"logBucketName,omitempty" xml:"logBucketName,omitempty"`
	// example:
	//
	// w-xxxxxxx/ray/logs/xxxxxx/
	LogPath *string `json:"logPath,omitempty" xml:"logPath,omitempty"`
	// example:
	//
	// Job finished successfully.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// {"owner": "alice"}
	MetadataJson *string `json:"metadataJson,omitempty" xml:"metadataJson,omitempty"`
	// example:
	//
	// myRayCluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// vpc
	NetworkServiceName *string `json:"networkServiceName,omitempty" xml:"networkServiceName,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	// 1750327083303
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// rj-xxxxxxxxxx
	SubmissionId *string `json:"submissionId,omitempty" xml:"submissionId,omitempty"`
	// example:
	//
	// HTTPMode
	SubmissionMode *string `json:"submissionMode,omitempty" xml:"submissionMode,omitempty"`
	// example:
	//
	// 1750327082303
	SubmitTime *int64 `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	Tags       []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// TSK-682e0112f6f24d9f9305b92174846985
	TaskBizId *string `json:"taskBizId,omitempty" xml:"taskBizId,omitempty"`
	// example:
	//
	// 60
	TtlSecondsAfterFinished *int32                              `json:"ttlSecondsAfterFinished,omitempty" xml:"ttlSecondsAfterFinished,omitempty"`
	VolumeIds               []*string                           `json:"volumeIds,omitempty" xml:"volumeIds,omitempty" type:"Repeated"`
	WorkerSpecs             []*GetRayJobResponseBodyWorkerSpecs `json:"workerSpecs,omitempty" xml:"workerSpecs,omitempty" type:"Repeated"`
	// example:
	//
	// oss://mybucket/hello.zip
	WorkingDir *string `json:"workingDir,omitempty" xml:"workingDir,omitempty"`
}

func (s GetRayJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRayJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetRayJobResponseBody) GetActiveDeadlineSeconds() *int32 {
	return s.ActiveDeadlineSeconds
}

func (s *GetRayJobResponseBody) GetBackoffLimit() *int32 {
	return s.BackoffLimit
}

func (s *GetRayJobResponseBody) GetClusterState() *string {
	return s.ClusterState
}

func (s *GetRayJobResponseBody) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetRayJobResponseBody) GetCuHours() *float64 {
	return s.CuHours
}

func (s *GetRayJobResponseBody) GetDashboardUrl() *string {
	return s.DashboardUrl
}

func (s *GetRayJobResponseBody) GetDashboardUrlExtra() []*string {
	return s.DashboardUrlExtra
}

func (s *GetRayJobResponseBody) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *GetRayJobResponseBody) GetDuration() *int64 {
	return s.Duration
}

func (s *GetRayJobResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetRayJobResponseBody) GetEntrypoint() *string {
	return s.Entrypoint
}

func (s *GetRayJobResponseBody) GetEntrypointMemory() *string {
	return s.EntrypointMemory
}

func (s *GetRayJobResponseBody) GetEntrypointNumCpus() *string {
	return s.EntrypointNumCpus
}

func (s *GetRayJobResponseBody) GetEntrypointNumGpus() *string {
	return s.EntrypointNumGpus
}

func (s *GetRayJobResponseBody) GetEntrypointResources() *string {
	return s.EntrypointResources
}

func (s *GetRayJobResponseBody) GetExtraParam() *string {
	return s.ExtraParam
}

func (s *GetRayJobResponseBody) GetGuHours() *GetRayJobResponseBodyGuHours {
	return s.GuHours
}

func (s *GetRayJobResponseBody) GetHeadSpec() *GetRayJobResponseBodyHeadSpec {
	return s.HeadSpec
}

func (s *GetRayJobResponseBody) GetLogBucketName() *string {
	return s.LogBucketName
}

func (s *GetRayJobResponseBody) GetLogPath() *string {
	return s.LogPath
}

func (s *GetRayJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRayJobResponseBody) GetMetadataJson() *string {
	return s.MetadataJson
}

func (s *GetRayJobResponseBody) GetName() *string {
	return s.Name
}

func (s *GetRayJobResponseBody) GetNetworkServiceName() *string {
	return s.NetworkServiceName
}

func (s *GetRayJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRayJobResponseBody) GetRuntimeEnvJson() *string {
	return s.RuntimeEnvJson
}

func (s *GetRayJobResponseBody) GetShutdownAfterJobFinishes() *bool {
	return s.ShutdownAfterJobFinishes
}

func (s *GetRayJobResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetRayJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetRayJobResponseBody) GetSubmissionId() *string {
	return s.SubmissionId
}

func (s *GetRayJobResponseBody) GetSubmissionMode() *string {
	return s.SubmissionMode
}

func (s *GetRayJobResponseBody) GetSubmitTime() *int64 {
	return s.SubmitTime
}

func (s *GetRayJobResponseBody) GetTags() []*Tag {
	return s.Tags
}

func (s *GetRayJobResponseBody) GetTaskBizId() *string {
	return s.TaskBizId
}

func (s *GetRayJobResponseBody) GetTtlSecondsAfterFinished() *int32 {
	return s.TtlSecondsAfterFinished
}

func (s *GetRayJobResponseBody) GetVolumeIds() []*string {
	return s.VolumeIds
}

func (s *GetRayJobResponseBody) GetWorkerSpecs() []*GetRayJobResponseBodyWorkerSpecs {
	return s.WorkerSpecs
}

func (s *GetRayJobResponseBody) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *GetRayJobResponseBody) SetActiveDeadlineSeconds(v int32) *GetRayJobResponseBody {
	s.ActiveDeadlineSeconds = &v
	return s
}

func (s *GetRayJobResponseBody) SetBackoffLimit(v int32) *GetRayJobResponseBody {
	s.BackoffLimit = &v
	return s
}

func (s *GetRayJobResponseBody) SetClusterState(v string) *GetRayJobResponseBody {
	s.ClusterState = &v
	return s
}

func (s *GetRayJobResponseBody) SetCreatorName(v string) *GetRayJobResponseBody {
	s.CreatorName = &v
	return s
}

func (s *GetRayJobResponseBody) SetCuHours(v float64) *GetRayJobResponseBody {
	s.CuHours = &v
	return s
}

func (s *GetRayJobResponseBody) SetDashboardUrl(v string) *GetRayJobResponseBody {
	s.DashboardUrl = &v
	return s
}

func (s *GetRayJobResponseBody) SetDashboardUrlExtra(v []*string) *GetRayJobResponseBody {
	s.DashboardUrlExtra = v
	return s
}

func (s *GetRayJobResponseBody) SetDisplayReleaseVersion(v string) *GetRayJobResponseBody {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *GetRayJobResponseBody) SetDuration(v int64) *GetRayJobResponseBody {
	s.Duration = &v
	return s
}

func (s *GetRayJobResponseBody) SetEndTime(v int64) *GetRayJobResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetRayJobResponseBody) SetEntrypoint(v string) *GetRayJobResponseBody {
	s.Entrypoint = &v
	return s
}

func (s *GetRayJobResponseBody) SetEntrypointMemory(v string) *GetRayJobResponseBody {
	s.EntrypointMemory = &v
	return s
}

func (s *GetRayJobResponseBody) SetEntrypointNumCpus(v string) *GetRayJobResponseBody {
	s.EntrypointNumCpus = &v
	return s
}

func (s *GetRayJobResponseBody) SetEntrypointNumGpus(v string) *GetRayJobResponseBody {
	s.EntrypointNumGpus = &v
	return s
}

func (s *GetRayJobResponseBody) SetEntrypointResources(v string) *GetRayJobResponseBody {
	s.EntrypointResources = &v
	return s
}

func (s *GetRayJobResponseBody) SetExtraParam(v string) *GetRayJobResponseBody {
	s.ExtraParam = &v
	return s
}

func (s *GetRayJobResponseBody) SetGuHours(v *GetRayJobResponseBodyGuHours) *GetRayJobResponseBody {
	s.GuHours = v
	return s
}

func (s *GetRayJobResponseBody) SetHeadSpec(v *GetRayJobResponseBodyHeadSpec) *GetRayJobResponseBody {
	s.HeadSpec = v
	return s
}

func (s *GetRayJobResponseBody) SetLogBucketName(v string) *GetRayJobResponseBody {
	s.LogBucketName = &v
	return s
}

func (s *GetRayJobResponseBody) SetLogPath(v string) *GetRayJobResponseBody {
	s.LogPath = &v
	return s
}

func (s *GetRayJobResponseBody) SetMessage(v string) *GetRayJobResponseBody {
	s.Message = &v
	return s
}

func (s *GetRayJobResponseBody) SetMetadataJson(v string) *GetRayJobResponseBody {
	s.MetadataJson = &v
	return s
}

func (s *GetRayJobResponseBody) SetName(v string) *GetRayJobResponseBody {
	s.Name = &v
	return s
}

func (s *GetRayJobResponseBody) SetNetworkServiceName(v string) *GetRayJobResponseBody {
	s.NetworkServiceName = &v
	return s
}

func (s *GetRayJobResponseBody) SetRequestId(v string) *GetRayJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRayJobResponseBody) SetRuntimeEnvJson(v string) *GetRayJobResponseBody {
	s.RuntimeEnvJson = &v
	return s
}

func (s *GetRayJobResponseBody) SetShutdownAfterJobFinishes(v bool) *GetRayJobResponseBody {
	s.ShutdownAfterJobFinishes = &v
	return s
}

func (s *GetRayJobResponseBody) SetStartTime(v int64) *GetRayJobResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetRayJobResponseBody) SetStatus(v string) *GetRayJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetRayJobResponseBody) SetSubmissionId(v string) *GetRayJobResponseBody {
	s.SubmissionId = &v
	return s
}

func (s *GetRayJobResponseBody) SetSubmissionMode(v string) *GetRayJobResponseBody {
	s.SubmissionMode = &v
	return s
}

func (s *GetRayJobResponseBody) SetSubmitTime(v int64) *GetRayJobResponseBody {
	s.SubmitTime = &v
	return s
}

func (s *GetRayJobResponseBody) SetTags(v []*Tag) *GetRayJobResponseBody {
	s.Tags = v
	return s
}

func (s *GetRayJobResponseBody) SetTaskBizId(v string) *GetRayJobResponseBody {
	s.TaskBizId = &v
	return s
}

func (s *GetRayJobResponseBody) SetTtlSecondsAfterFinished(v int32) *GetRayJobResponseBody {
	s.TtlSecondsAfterFinished = &v
	return s
}

func (s *GetRayJobResponseBody) SetVolumeIds(v []*string) *GetRayJobResponseBody {
	s.VolumeIds = v
	return s
}

func (s *GetRayJobResponseBody) SetWorkerSpecs(v []*GetRayJobResponseBodyWorkerSpecs) *GetRayJobResponseBody {
	s.WorkerSpecs = v
	return s
}

func (s *GetRayJobResponseBody) SetWorkingDir(v string) *GetRayJobResponseBody {
	s.WorkingDir = &v
	return s
}

func (s *GetRayJobResponseBody) Validate() error {
	if s.GuHours != nil {
		if err := s.GuHours.Validate(); err != nil {
			return err
		}
	}
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
	if s.WorkerSpecs != nil {
		for _, item := range s.WorkerSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRayJobResponseBodyGuHours struct {
	// example:
	//
	// 2.6
	GpuHours *float64 `json:"gpuHours,omitempty" xml:"gpuHours,omitempty"`
	// example:
	//
	// ecs.gn6i-c4g1.xlarge
	GpuSpec *string `json:"gpuSpec,omitempty" xml:"gpuSpec,omitempty"`
}

func (s GetRayJobResponseBodyGuHours) String() string {
	return dara.Prettify(s)
}

func (s GetRayJobResponseBodyGuHours) GoString() string {
	return s.String()
}

func (s *GetRayJobResponseBodyGuHours) GetGpuHours() *float64 {
	return s.GpuHours
}

func (s *GetRayJobResponseBodyGuHours) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *GetRayJobResponseBodyGuHours) SetGpuHours(v float64) *GetRayJobResponseBodyGuHours {
	s.GpuHours = &v
	return s
}

func (s *GetRayJobResponseBodyGuHours) SetGpuSpec(v string) *GetRayJobResponseBodyGuHours {
	s.GpuSpec = &v
	return s
}

func (s *GetRayJobResponseBodyGuHours) Validate() error {
	return dara.Validate(s)
}

type GetRayJobResponseBodyHeadSpec struct {
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
	// example:
	//
	// 1
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
}

func (s GetRayJobResponseBodyHeadSpec) String() string {
	return dara.Prettify(s)
}

func (s GetRayJobResponseBodyHeadSpec) GoString() string {
	return s.String()
}

func (s *GetRayJobResponseBodyHeadSpec) GetCpu() *string {
	return s.Cpu
}

func (s *GetRayJobResponseBodyHeadSpec) GetEnableAutoScaling() *bool {
	return s.EnableAutoScaling
}

func (s *GetRayJobResponseBodyHeadSpec) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *GetRayJobResponseBodyHeadSpec) GetIdleTimeoutSeconds() *int32 {
	return s.IdleTimeoutSeconds
}

func (s *GetRayJobResponseBodyHeadSpec) GetMemory() *string {
	return s.Memory
}

func (s *GetRayJobResponseBodyHeadSpec) GetQueueName() *string {
	return s.QueueName
}

func (s *GetRayJobResponseBodyHeadSpec) GetReplica() *int32 {
	return s.Replica
}

func (s *GetRayJobResponseBodyHeadSpec) SetCpu(v string) *GetRayJobResponseBodyHeadSpec {
	s.Cpu = &v
	return s
}

func (s *GetRayJobResponseBodyHeadSpec) SetEnableAutoScaling(v bool) *GetRayJobResponseBodyHeadSpec {
	s.EnableAutoScaling = &v
	return s
}

func (s *GetRayJobResponseBodyHeadSpec) SetGpuSpec(v string) *GetRayJobResponseBodyHeadSpec {
	s.GpuSpec = &v
	return s
}

func (s *GetRayJobResponseBodyHeadSpec) SetIdleTimeoutSeconds(v int32) *GetRayJobResponseBodyHeadSpec {
	s.IdleTimeoutSeconds = &v
	return s
}

func (s *GetRayJobResponseBodyHeadSpec) SetMemory(v string) *GetRayJobResponseBodyHeadSpec {
	s.Memory = &v
	return s
}

func (s *GetRayJobResponseBodyHeadSpec) SetQueueName(v string) *GetRayJobResponseBodyHeadSpec {
	s.QueueName = &v
	return s
}

func (s *GetRayJobResponseBodyHeadSpec) SetReplica(v int32) *GetRayJobResponseBodyHeadSpec {
	s.Replica = &v
	return s
}

func (s *GetRayJobResponseBodyHeadSpec) Validate() error {
	return dara.Validate(s)
}

type GetRayJobResponseBodyWorkerSpecs struct {
	// example:
	//
	// 2
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
	// 8Gi
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
	// 1
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
}

func (s GetRayJobResponseBodyWorkerSpecs) String() string {
	return dara.Prettify(s)
}

func (s GetRayJobResponseBodyWorkerSpecs) GoString() string {
	return s.String()
}

func (s *GetRayJobResponseBodyWorkerSpecs) GetCpu() *string {
	return s.Cpu
}

func (s *GetRayJobResponseBodyWorkerSpecs) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *GetRayJobResponseBodyWorkerSpecs) GetGroupName() *string {
	return s.GroupName
}

func (s *GetRayJobResponseBodyWorkerSpecs) GetMaxReplica() *int32 {
	return s.MaxReplica
}

func (s *GetRayJobResponseBodyWorkerSpecs) GetMemory() *string {
	return s.Memory
}

func (s *GetRayJobResponseBodyWorkerSpecs) GetMinReplica() *int32 {
	return s.MinReplica
}

func (s *GetRayJobResponseBodyWorkerSpecs) GetQueueName() *string {
	return s.QueueName
}

func (s *GetRayJobResponseBodyWorkerSpecs) GetReplica() *int32 {
	return s.Replica
}

func (s *GetRayJobResponseBodyWorkerSpecs) SetCpu(v string) *GetRayJobResponseBodyWorkerSpecs {
	s.Cpu = &v
	return s
}

func (s *GetRayJobResponseBodyWorkerSpecs) SetGpuSpec(v string) *GetRayJobResponseBodyWorkerSpecs {
	s.GpuSpec = &v
	return s
}

func (s *GetRayJobResponseBodyWorkerSpecs) SetGroupName(v string) *GetRayJobResponseBodyWorkerSpecs {
	s.GroupName = &v
	return s
}

func (s *GetRayJobResponseBodyWorkerSpecs) SetMaxReplica(v int32) *GetRayJobResponseBodyWorkerSpecs {
	s.MaxReplica = &v
	return s
}

func (s *GetRayJobResponseBodyWorkerSpecs) SetMemory(v string) *GetRayJobResponseBodyWorkerSpecs {
	s.Memory = &v
	return s
}

func (s *GetRayJobResponseBodyWorkerSpecs) SetMinReplica(v int32) *GetRayJobResponseBodyWorkerSpecs {
	s.MinReplica = &v
	return s
}

func (s *GetRayJobResponseBodyWorkerSpecs) SetQueueName(v string) *GetRayJobResponseBodyWorkerSpecs {
	s.QueueName = &v
	return s
}

func (s *GetRayJobResponseBodyWorkerSpecs) SetReplica(v int32) *GetRayJobResponseBodyWorkerSpecs {
	s.Replica = &v
	return s
}

func (s *GetRayJobResponseBodyWorkerSpecs) Validate() error {
	return dara.Validate(s)
}
