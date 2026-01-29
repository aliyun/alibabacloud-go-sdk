// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRayClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetRayClusterResponseBody
	GetClusterId() *string
	SetCreateTime(v int64) *GetRayClusterResponseBody
	GetCreateTime() *int64
	SetCreator(v string) *GetRayClusterResponseBody
	GetCreator() *string
	SetCreatorName(v string) *GetRayClusterResponseBody
	GetCreatorName() *string
	SetDashboardUrl(v string) *GetRayClusterResponseBody
	GetDashboardUrl() *string
	SetDescription(v string) *GetRayClusterResponseBody
	GetDescription() *string
	SetDisplayReleaseVersion(v string) *GetRayClusterResponseBody
	GetDisplayReleaseVersion() *string
	SetExtraParam(v string) *GetRayClusterResponseBody
	GetExtraParam() *string
	SetGrpcEndpoint(v string) *GetRayClusterResponseBody
	GetGrpcEndpoint() *string
	SetHeadSpec(v *GetRayClusterResponseBodyHeadSpec) *GetRayClusterResponseBody
	GetHeadSpec() *GetRayClusterResponseBodyHeadSpec
	SetInstanceId(v string) *GetRayClusterResponseBody
	GetInstanceId() *string
	SetInstances(v []*GetRayClusterResponseBodyInstances) *GetRayClusterResponseBody
	GetInstances() []*GetRayClusterResponseBodyInstances
	SetJobUrl(v string) *GetRayClusterResponseBody
	GetJobUrl() *string
	SetMessage(v string) *GetRayClusterResponseBody
	GetMessage() *string
	SetModified(v bool) *GetRayClusterResponseBody
	GetModified() *bool
	SetModifiedTime(v int64) *GetRayClusterResponseBody
	GetModifiedTime() *int64
	SetModifier(v string) *GetRayClusterResponseBody
	GetModifier() *string
	SetModifierName(v string) *GetRayClusterResponseBody
	GetModifierName() *string
	SetName(v string) *GetRayClusterResponseBody
	GetName() *string
	SetNetworkServiceName(v string) *GetRayClusterResponseBody
	GetNetworkServiceName() *string
	SetRequestId(v string) *GetRayClusterResponseBody
	GetRequestId() *string
	SetStartTime(v int64) *GetRayClusterResponseBody
	GetStartTime() *int64
	SetState(v string) *GetRayClusterResponseBody
	GetState() *string
	SetSubmitToken(v string) *GetRayClusterResponseBody
	GetSubmitToken() *string
	SetUserId(v string) *GetRayClusterResponseBody
	GetUserId() *string
	SetWorkerSpec(v []*GetRayClusterResponseBodyWorkerSpec) *GetRayClusterResponseBody
	GetWorkerSpec() []*GetRayClusterResponseBodyWorkerSpec
}

type GetRayClusterResponseBody struct {
	// example:
	//
	// ray-k7nm8ahl5te4tg91
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// example:
	//
	// 1750327083303
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 202077646755523991
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// Alice
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// https://emr-ray-gateway-cn-hangzhou.aliyuncs.com/workspace/w-xxxxxxxx/raycluster/ray-xxxxxx/dashboard?token=xxxxxx
	DashboardUrl *string `json:"dashboardUrl,omitempty" xml:"dashboardUrl,omitempty"`
	// example:
	//
	// Ray Cluster for dev.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// ray-1.0.0 (Ray 2.47.1, Python 3.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// example:
	//
	// {}
	ExtraParam *string `json:"extraParam,omitempty" xml:"extraParam,omitempty"`
	// example:
	//
	// ray://emr-spark-ray-gateway-cn-hangzhou-internal.emr.aliyuncs.com:80
	GrpcEndpoint *string                            `json:"grpcEndpoint,omitempty" xml:"grpcEndpoint,omitempty"`
	HeadSpec     *GetRayClusterResponseBodyHeadSpec `json:"headSpec,omitempty" xml:"headSpec,omitempty" type:"Struct"`
	// example:
	//
	// ray-k7nm8ahl5te4tg93-xxxxxxx
	InstanceId *string                               `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Instances  []*GetRayClusterResponseBodyInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// example:
	//
	// https://emr-spark-ray-gateway-cn-hangzhou.aliyuncs.com
	JobUrl *string `json:"jobUrl,omitempty" xml:"jobUrl,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// false
	Modified *bool `json:"modified,omitempty" xml:"modified,omitempty"`
	// example:
	//
	// 1754274541693
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// 202077646755523991
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// Alice
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
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
	// 1750327083303
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// Running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// 1d06484d3b424f7fa4ab7082a4076da2
	SubmitToken *string `json:"submitToken,omitempty" xml:"submitToken,omitempty"`
	// example:
	//
	// 123456789012
	UserId     *string                                `json:"userId,omitempty" xml:"userId,omitempty"`
	WorkerSpec []*GetRayClusterResponseBodyWorkerSpec `json:"workerSpec,omitempty" xml:"workerSpec,omitempty" type:"Repeated"`
}

func (s GetRayClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRayClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetRayClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetRayClusterResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetRayClusterResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetRayClusterResponseBody) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetRayClusterResponseBody) GetDashboardUrl() *string {
	return s.DashboardUrl
}

func (s *GetRayClusterResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetRayClusterResponseBody) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *GetRayClusterResponseBody) GetExtraParam() *string {
	return s.ExtraParam
}

func (s *GetRayClusterResponseBody) GetGrpcEndpoint() *string {
	return s.GrpcEndpoint
}

func (s *GetRayClusterResponseBody) GetHeadSpec() *GetRayClusterResponseBodyHeadSpec {
	return s.HeadSpec
}

func (s *GetRayClusterResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRayClusterResponseBody) GetInstances() []*GetRayClusterResponseBodyInstances {
	return s.Instances
}

func (s *GetRayClusterResponseBody) GetJobUrl() *string {
	return s.JobUrl
}

func (s *GetRayClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRayClusterResponseBody) GetModified() *bool {
	return s.Modified
}

func (s *GetRayClusterResponseBody) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *GetRayClusterResponseBody) GetModifier() *string {
	return s.Modifier
}

func (s *GetRayClusterResponseBody) GetModifierName() *string {
	return s.ModifierName
}

func (s *GetRayClusterResponseBody) GetName() *string {
	return s.Name
}

func (s *GetRayClusterResponseBody) GetNetworkServiceName() *string {
	return s.NetworkServiceName
}

func (s *GetRayClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRayClusterResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetRayClusterResponseBody) GetState() *string {
	return s.State
}

func (s *GetRayClusterResponseBody) GetSubmitToken() *string {
	return s.SubmitToken
}

func (s *GetRayClusterResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetRayClusterResponseBody) GetWorkerSpec() []*GetRayClusterResponseBodyWorkerSpec {
	return s.WorkerSpec
}

func (s *GetRayClusterResponseBody) SetClusterId(v string) *GetRayClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetRayClusterResponseBody) SetCreateTime(v int64) *GetRayClusterResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetRayClusterResponseBody) SetCreator(v string) *GetRayClusterResponseBody {
	s.Creator = &v
	return s
}

func (s *GetRayClusterResponseBody) SetCreatorName(v string) *GetRayClusterResponseBody {
	s.CreatorName = &v
	return s
}

func (s *GetRayClusterResponseBody) SetDashboardUrl(v string) *GetRayClusterResponseBody {
	s.DashboardUrl = &v
	return s
}

func (s *GetRayClusterResponseBody) SetDescription(v string) *GetRayClusterResponseBody {
	s.Description = &v
	return s
}

func (s *GetRayClusterResponseBody) SetDisplayReleaseVersion(v string) *GetRayClusterResponseBody {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *GetRayClusterResponseBody) SetExtraParam(v string) *GetRayClusterResponseBody {
	s.ExtraParam = &v
	return s
}

func (s *GetRayClusterResponseBody) SetGrpcEndpoint(v string) *GetRayClusterResponseBody {
	s.GrpcEndpoint = &v
	return s
}

func (s *GetRayClusterResponseBody) SetHeadSpec(v *GetRayClusterResponseBodyHeadSpec) *GetRayClusterResponseBody {
	s.HeadSpec = v
	return s
}

func (s *GetRayClusterResponseBody) SetInstanceId(v string) *GetRayClusterResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetRayClusterResponseBody) SetInstances(v []*GetRayClusterResponseBodyInstances) *GetRayClusterResponseBody {
	s.Instances = v
	return s
}

func (s *GetRayClusterResponseBody) SetJobUrl(v string) *GetRayClusterResponseBody {
	s.JobUrl = &v
	return s
}

func (s *GetRayClusterResponseBody) SetMessage(v string) *GetRayClusterResponseBody {
	s.Message = &v
	return s
}

func (s *GetRayClusterResponseBody) SetModified(v bool) *GetRayClusterResponseBody {
	s.Modified = &v
	return s
}

func (s *GetRayClusterResponseBody) SetModifiedTime(v int64) *GetRayClusterResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetRayClusterResponseBody) SetModifier(v string) *GetRayClusterResponseBody {
	s.Modifier = &v
	return s
}

func (s *GetRayClusterResponseBody) SetModifierName(v string) *GetRayClusterResponseBody {
	s.ModifierName = &v
	return s
}

func (s *GetRayClusterResponseBody) SetName(v string) *GetRayClusterResponseBody {
	s.Name = &v
	return s
}

func (s *GetRayClusterResponseBody) SetNetworkServiceName(v string) *GetRayClusterResponseBody {
	s.NetworkServiceName = &v
	return s
}

func (s *GetRayClusterResponseBody) SetRequestId(v string) *GetRayClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRayClusterResponseBody) SetStartTime(v int64) *GetRayClusterResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetRayClusterResponseBody) SetState(v string) *GetRayClusterResponseBody {
	s.State = &v
	return s
}

func (s *GetRayClusterResponseBody) SetSubmitToken(v string) *GetRayClusterResponseBody {
	s.SubmitToken = &v
	return s
}

func (s *GetRayClusterResponseBody) SetUserId(v string) *GetRayClusterResponseBody {
	s.UserId = &v
	return s
}

func (s *GetRayClusterResponseBody) SetWorkerSpec(v []*GetRayClusterResponseBodyWorkerSpec) *GetRayClusterResponseBody {
	s.WorkerSpec = v
	return s
}

func (s *GetRayClusterResponseBody) Validate() error {
	if s.HeadSpec != nil {
		if err := s.HeadSpec.Validate(); err != nil {
			return err
		}
	}
	if s.Instances != nil {
		for _, item := range s.Instances {
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

type GetRayClusterResponseBodyHeadSpec struct {
	// example:
	//
	// 2
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// false
	EnableAutoScaling *bool `json:"enableAutoScaling,omitempty" xml:"enableAutoScaling,omitempty"`
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

func (s GetRayClusterResponseBodyHeadSpec) String() string {
	return dara.Prettify(s)
}

func (s GetRayClusterResponseBodyHeadSpec) GoString() string {
	return s.String()
}

func (s *GetRayClusterResponseBodyHeadSpec) GetCpu() *string {
	return s.Cpu
}

func (s *GetRayClusterResponseBodyHeadSpec) GetEnableAutoScaling() *bool {
	return s.EnableAutoScaling
}

func (s *GetRayClusterResponseBodyHeadSpec) GetIdleTimeoutSeconds() *int32 {
	return s.IdleTimeoutSeconds
}

func (s *GetRayClusterResponseBodyHeadSpec) GetMemory() *string {
	return s.Memory
}

func (s *GetRayClusterResponseBodyHeadSpec) GetQueueName() *string {
	return s.QueueName
}

func (s *GetRayClusterResponseBodyHeadSpec) GetReplica() *int32 {
	return s.Replica
}

func (s *GetRayClusterResponseBodyHeadSpec) SetCpu(v string) *GetRayClusterResponseBodyHeadSpec {
	s.Cpu = &v
	return s
}

func (s *GetRayClusterResponseBodyHeadSpec) SetEnableAutoScaling(v bool) *GetRayClusterResponseBodyHeadSpec {
	s.EnableAutoScaling = &v
	return s
}

func (s *GetRayClusterResponseBodyHeadSpec) SetIdleTimeoutSeconds(v int32) *GetRayClusterResponseBodyHeadSpec {
	s.IdleTimeoutSeconds = &v
	return s
}

func (s *GetRayClusterResponseBodyHeadSpec) SetMemory(v string) *GetRayClusterResponseBodyHeadSpec {
	s.Memory = &v
	return s
}

func (s *GetRayClusterResponseBodyHeadSpec) SetQueueName(v string) *GetRayClusterResponseBodyHeadSpec {
	s.QueueName = &v
	return s
}

func (s *GetRayClusterResponseBodyHeadSpec) SetReplica(v int32) *GetRayClusterResponseBodyHeadSpec {
	s.Replica = &v
	return s
}

func (s *GetRayClusterResponseBodyHeadSpec) Validate() error {
	return dara.Validate(s)
}

type GetRayClusterResponseBodyInstances struct {
	// example:
	//
	// 137
	ContainerExitCode *int32 `json:"containerExitCode,omitempty" xml:"containerExitCode,omitempty"`
	// example:
	//
	// Running
	ContainerState *string `json:"containerState,omitempty" xml:"containerState,omitempty"`
	// example:
	//
	// ContainerExit
	ContainerStateMessage *string `json:"containerStateMessage,omitempty" xml:"containerStateMessage,omitempty"`
	// example:
	//
	// ok
	ContainerStateReason *string `json:"containerStateReason,omitempty" xml:"containerStateReason,omitempty"`
	// example:
	//
	// 1735870116167
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// ray-uiulpgow9xljim10-head-7cgta
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// Running
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// example:
	//
	// OOMKilled
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// example:
	//
	// 1762946698000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// Head
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetRayClusterResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s GetRayClusterResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *GetRayClusterResponseBodyInstances) GetContainerExitCode() *int32 {
	return s.ContainerExitCode
}

func (s *GetRayClusterResponseBodyInstances) GetContainerState() *string {
	return s.ContainerState
}

func (s *GetRayClusterResponseBodyInstances) GetContainerStateMessage() *string {
	return s.ContainerStateMessage
}

func (s *GetRayClusterResponseBodyInstances) GetContainerStateReason() *string {
	return s.ContainerStateReason
}

func (s *GetRayClusterResponseBodyInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetRayClusterResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRayClusterResponseBodyInstances) GetMessage() *string {
	return s.Message
}

func (s *GetRayClusterResponseBodyInstances) GetPhase() *string {
	return s.Phase
}

func (s *GetRayClusterResponseBodyInstances) GetReason() *string {
	return s.Reason
}

func (s *GetRayClusterResponseBodyInstances) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetRayClusterResponseBodyInstances) GetType() *string {
	return s.Type
}

func (s *GetRayClusterResponseBodyInstances) SetContainerExitCode(v int32) *GetRayClusterResponseBodyInstances {
	s.ContainerExitCode = &v
	return s
}

func (s *GetRayClusterResponseBodyInstances) SetContainerState(v string) *GetRayClusterResponseBodyInstances {
	s.ContainerState = &v
	return s
}

func (s *GetRayClusterResponseBodyInstances) SetContainerStateMessage(v string) *GetRayClusterResponseBodyInstances {
	s.ContainerStateMessage = &v
	return s
}

func (s *GetRayClusterResponseBodyInstances) SetContainerStateReason(v string) *GetRayClusterResponseBodyInstances {
	s.ContainerStateReason = &v
	return s
}

func (s *GetRayClusterResponseBodyInstances) SetCreateTime(v int64) *GetRayClusterResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *GetRayClusterResponseBodyInstances) SetInstanceId(v string) *GetRayClusterResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *GetRayClusterResponseBodyInstances) SetMessage(v string) *GetRayClusterResponseBodyInstances {
	s.Message = &v
	return s
}

func (s *GetRayClusterResponseBodyInstances) SetPhase(v string) *GetRayClusterResponseBodyInstances {
	s.Phase = &v
	return s
}

func (s *GetRayClusterResponseBodyInstances) SetReason(v string) *GetRayClusterResponseBodyInstances {
	s.Reason = &v
	return s
}

func (s *GetRayClusterResponseBodyInstances) SetStartTime(v int64) *GetRayClusterResponseBodyInstances {
	s.StartTime = &v
	return s
}

func (s *GetRayClusterResponseBodyInstances) SetType(v string) *GetRayClusterResponseBodyInstances {
	s.Type = &v
	return s
}

func (s *GetRayClusterResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type GetRayClusterResponseBodyWorkerSpec struct {
	// example:
	//
	// 2
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
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
	// 2
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
	// example:
	//
	// CPU
	WorkerType *string `json:"workerType,omitempty" xml:"workerType,omitempty"`
}

func (s GetRayClusterResponseBodyWorkerSpec) String() string {
	return dara.Prettify(s)
}

func (s GetRayClusterResponseBodyWorkerSpec) GoString() string {
	return s.String()
}

func (s *GetRayClusterResponseBodyWorkerSpec) GetCpu() *string {
	return s.Cpu
}

func (s *GetRayClusterResponseBodyWorkerSpec) GetGroupName() *string {
	return s.GroupName
}

func (s *GetRayClusterResponseBodyWorkerSpec) GetMaxReplica() *int32 {
	return s.MaxReplica
}

func (s *GetRayClusterResponseBodyWorkerSpec) GetMemory() *string {
	return s.Memory
}

func (s *GetRayClusterResponseBodyWorkerSpec) GetMinReplica() *int32 {
	return s.MinReplica
}

func (s *GetRayClusterResponseBodyWorkerSpec) GetQueueName() *string {
	return s.QueueName
}

func (s *GetRayClusterResponseBodyWorkerSpec) GetReplica() *int32 {
	return s.Replica
}

func (s *GetRayClusterResponseBodyWorkerSpec) GetWorkerType() *string {
	return s.WorkerType
}

func (s *GetRayClusterResponseBodyWorkerSpec) SetCpu(v string) *GetRayClusterResponseBodyWorkerSpec {
	s.Cpu = &v
	return s
}

func (s *GetRayClusterResponseBodyWorkerSpec) SetGroupName(v string) *GetRayClusterResponseBodyWorkerSpec {
	s.GroupName = &v
	return s
}

func (s *GetRayClusterResponseBodyWorkerSpec) SetMaxReplica(v int32) *GetRayClusterResponseBodyWorkerSpec {
	s.MaxReplica = &v
	return s
}

func (s *GetRayClusterResponseBodyWorkerSpec) SetMemory(v string) *GetRayClusterResponseBodyWorkerSpec {
	s.Memory = &v
	return s
}

func (s *GetRayClusterResponseBodyWorkerSpec) SetMinReplica(v int32) *GetRayClusterResponseBodyWorkerSpec {
	s.MinReplica = &v
	return s
}

func (s *GetRayClusterResponseBodyWorkerSpec) SetQueueName(v string) *GetRayClusterResponseBodyWorkerSpec {
	s.QueueName = &v
	return s
}

func (s *GetRayClusterResponseBodyWorkerSpec) SetReplica(v int32) *GetRayClusterResponseBodyWorkerSpec {
	s.Replica = &v
	return s
}

func (s *GetRayClusterResponseBodyWorkerSpec) SetWorkerType(v string) *GetRayClusterResponseBodyWorkerSpec {
	s.WorkerType = &v
	return s
}

func (s *GetRayClusterResponseBodyWorkerSpec) Validate() error {
	return dara.Validate(s)
}
