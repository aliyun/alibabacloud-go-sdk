// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRayClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRayClusters(v []*ListRayClusterResponseBodyRayClusters) *ListRayClusterResponseBody
	GetRayClusters() []*ListRayClusterResponseBodyRayClusters
	SetRequestId(v string) *ListRayClusterResponseBody
	GetRequestId() *string
}

type ListRayClusterResponseBody struct {
	RayClusters []*ListRayClusterResponseBodyRayClusters `json:"rayClusters,omitempty" xml:"rayClusters,omitempty" type:"Repeated"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRayClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRayClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ListRayClusterResponseBody) GetRayClusters() []*ListRayClusterResponseBodyRayClusters {
	return s.RayClusters
}

func (s *ListRayClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRayClusterResponseBody) SetRayClusters(v []*ListRayClusterResponseBodyRayClusters) *ListRayClusterResponseBody {
	s.RayClusters = v
	return s
}

func (s *ListRayClusterResponseBody) SetRequestId(v string) *ListRayClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRayClusterResponseBody) Validate() error {
	if s.RayClusters != nil {
		for _, item := range s.RayClusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRayClusterResponseBodyRayClusters struct {
	// example:
	//
	// ray-uiulpgow9xljimm1
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// example:
	//
	// 1723722279800
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 202077646755123991
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// Alice
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// https://emr-spark-ray-gateway-cn-hangzhou.aliyuncs.com?token=xxxxxxxxx
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
	// ray://emr-spark-ray-gateway-cn-hanghzou-internal.emr.aliyuncs.com:80
	GrpcEndpoint *string                                        `json:"grpcEndpoint,omitempty" xml:"grpcEndpoint,omitempty"`
	HeadSpec     *ListRayClusterResponseBodyRayClustersHeadSpec `json:"headSpec,omitempty" xml:"headSpec,omitempty" type:"Struct"`
	// example:
	//
	// ray-uiulpgow9xljimm1-xxxxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
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
	// 1723722279800
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// 202077646755123991
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// Alice
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// example:
	//
	// testRayCluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// vpc
	NetworkServiceName *string `json:"networkServiceName,omitempty" xml:"networkServiceName,omitempty"`
	// example:
	//
	// 1723722279800
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// Running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// 1234567890
	UserId     *string                                            `json:"userId,omitempty" xml:"userId,omitempty"`
	WorkerSpec []*ListRayClusterResponseBodyRayClustersWorkerSpec `json:"workerSpec,omitempty" xml:"workerSpec,omitempty" type:"Repeated"`
}

func (s ListRayClusterResponseBodyRayClusters) String() string {
	return dara.Prettify(s)
}

func (s ListRayClusterResponseBodyRayClusters) GoString() string {
	return s.String()
}

func (s *ListRayClusterResponseBodyRayClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListRayClusterResponseBodyRayClusters) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListRayClusterResponseBodyRayClusters) GetCreator() *string {
	return s.Creator
}

func (s *ListRayClusterResponseBodyRayClusters) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListRayClusterResponseBodyRayClusters) GetDashboardUrl() *string {
	return s.DashboardUrl
}

func (s *ListRayClusterResponseBodyRayClusters) GetDescription() *string {
	return s.Description
}

func (s *ListRayClusterResponseBodyRayClusters) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *ListRayClusterResponseBodyRayClusters) GetGrpcEndpoint() *string {
	return s.GrpcEndpoint
}

func (s *ListRayClusterResponseBodyRayClusters) GetHeadSpec() *ListRayClusterResponseBodyRayClustersHeadSpec {
	return s.HeadSpec
}

func (s *ListRayClusterResponseBodyRayClusters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRayClusterResponseBodyRayClusters) GetMessage() *string {
	return s.Message
}

func (s *ListRayClusterResponseBodyRayClusters) GetModified() *bool {
	return s.Modified
}

func (s *ListRayClusterResponseBodyRayClusters) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListRayClusterResponseBodyRayClusters) GetModifier() *string {
	return s.Modifier
}

func (s *ListRayClusterResponseBodyRayClusters) GetModifierName() *string {
	return s.ModifierName
}

func (s *ListRayClusterResponseBodyRayClusters) GetName() *string {
	return s.Name
}

func (s *ListRayClusterResponseBodyRayClusters) GetNetworkServiceName() *string {
	return s.NetworkServiceName
}

func (s *ListRayClusterResponseBodyRayClusters) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListRayClusterResponseBodyRayClusters) GetState() *string {
	return s.State
}

func (s *ListRayClusterResponseBodyRayClusters) GetUserId() *string {
	return s.UserId
}

func (s *ListRayClusterResponseBodyRayClusters) GetWorkerSpec() []*ListRayClusterResponseBodyRayClustersWorkerSpec {
	return s.WorkerSpec
}

func (s *ListRayClusterResponseBodyRayClusters) SetClusterId(v string) *ListRayClusterResponseBodyRayClusters {
	s.ClusterId = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetCreateTime(v int64) *ListRayClusterResponseBodyRayClusters {
	s.CreateTime = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetCreator(v string) *ListRayClusterResponseBodyRayClusters {
	s.Creator = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetCreatorName(v string) *ListRayClusterResponseBodyRayClusters {
	s.CreatorName = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetDashboardUrl(v string) *ListRayClusterResponseBodyRayClusters {
	s.DashboardUrl = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetDescription(v string) *ListRayClusterResponseBodyRayClusters {
	s.Description = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetDisplayReleaseVersion(v string) *ListRayClusterResponseBodyRayClusters {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetGrpcEndpoint(v string) *ListRayClusterResponseBodyRayClusters {
	s.GrpcEndpoint = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetHeadSpec(v *ListRayClusterResponseBodyRayClustersHeadSpec) *ListRayClusterResponseBodyRayClusters {
	s.HeadSpec = v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetInstanceId(v string) *ListRayClusterResponseBodyRayClusters {
	s.InstanceId = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetMessage(v string) *ListRayClusterResponseBodyRayClusters {
	s.Message = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetModified(v bool) *ListRayClusterResponseBodyRayClusters {
	s.Modified = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetModifiedTime(v int64) *ListRayClusterResponseBodyRayClusters {
	s.ModifiedTime = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetModifier(v string) *ListRayClusterResponseBodyRayClusters {
	s.Modifier = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetModifierName(v string) *ListRayClusterResponseBodyRayClusters {
	s.ModifierName = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetName(v string) *ListRayClusterResponseBodyRayClusters {
	s.Name = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetNetworkServiceName(v string) *ListRayClusterResponseBodyRayClusters {
	s.NetworkServiceName = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetStartTime(v int64) *ListRayClusterResponseBodyRayClusters {
	s.StartTime = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetState(v string) *ListRayClusterResponseBodyRayClusters {
	s.State = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetUserId(v string) *ListRayClusterResponseBodyRayClusters {
	s.UserId = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) SetWorkerSpec(v []*ListRayClusterResponseBodyRayClustersWorkerSpec) *ListRayClusterResponseBodyRayClusters {
	s.WorkerSpec = v
	return s
}

func (s *ListRayClusterResponseBodyRayClusters) Validate() error {
	if s.HeadSpec != nil {
		if err := s.HeadSpec.Validate(); err != nil {
			return err
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

type ListRayClusterResponseBodyRayClustersHeadSpec struct {
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

func (s ListRayClusterResponseBodyRayClustersHeadSpec) String() string {
	return dara.Prettify(s)
}

func (s ListRayClusterResponseBodyRayClustersHeadSpec) GoString() string {
	return s.String()
}

func (s *ListRayClusterResponseBodyRayClustersHeadSpec) GetCpu() *string {
	return s.Cpu
}

func (s *ListRayClusterResponseBodyRayClustersHeadSpec) GetEnableAutoScaling() *bool {
	return s.EnableAutoScaling
}

func (s *ListRayClusterResponseBodyRayClustersHeadSpec) GetIdleTimeoutSeconds() *int32 {
	return s.IdleTimeoutSeconds
}

func (s *ListRayClusterResponseBodyRayClustersHeadSpec) GetMemory() *string {
	return s.Memory
}

func (s *ListRayClusterResponseBodyRayClustersHeadSpec) GetQueueName() *string {
	return s.QueueName
}

func (s *ListRayClusterResponseBodyRayClustersHeadSpec) GetReplica() *int32 {
	return s.Replica
}

func (s *ListRayClusterResponseBodyRayClustersHeadSpec) SetCpu(v string) *ListRayClusterResponseBodyRayClustersHeadSpec {
	s.Cpu = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClustersHeadSpec) SetEnableAutoScaling(v bool) *ListRayClusterResponseBodyRayClustersHeadSpec {
	s.EnableAutoScaling = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClustersHeadSpec) SetIdleTimeoutSeconds(v int32) *ListRayClusterResponseBodyRayClustersHeadSpec {
	s.IdleTimeoutSeconds = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClustersHeadSpec) SetMemory(v string) *ListRayClusterResponseBodyRayClustersHeadSpec {
	s.Memory = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClustersHeadSpec) SetQueueName(v string) *ListRayClusterResponseBodyRayClustersHeadSpec {
	s.QueueName = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClustersHeadSpec) SetReplica(v int32) *ListRayClusterResponseBodyRayClustersHeadSpec {
	s.Replica = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClustersHeadSpec) Validate() error {
	return dara.Validate(s)
}

type ListRayClusterResponseBodyRayClustersWorkerSpec struct {
	// example:
	//
	// 2
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// Group1
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

func (s ListRayClusterResponseBodyRayClustersWorkerSpec) String() string {
	return dara.Prettify(s)
}

func (s ListRayClusterResponseBodyRayClustersWorkerSpec) GoString() string {
	return s.String()
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) GetCpu() *string {
	return s.Cpu
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) GetGroupName() *string {
	return s.GroupName
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) GetMaxReplica() *int32 {
	return s.MaxReplica
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) GetMemory() *string {
	return s.Memory
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) GetMinReplica() *int32 {
	return s.MinReplica
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) GetQueueName() *string {
	return s.QueueName
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) GetReplica() *int32 {
	return s.Replica
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) GetWorkerType() *string {
	return s.WorkerType
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) SetCpu(v string) *ListRayClusterResponseBodyRayClustersWorkerSpec {
	s.Cpu = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) SetGroupName(v string) *ListRayClusterResponseBodyRayClustersWorkerSpec {
	s.GroupName = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) SetMaxReplica(v int32) *ListRayClusterResponseBodyRayClustersWorkerSpec {
	s.MaxReplica = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) SetMemory(v string) *ListRayClusterResponseBodyRayClustersWorkerSpec {
	s.Memory = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) SetMinReplica(v int32) *ListRayClusterResponseBodyRayClustersWorkerSpec {
	s.MinReplica = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) SetQueueName(v string) *ListRayClusterResponseBodyRayClustersWorkerSpec {
	s.QueueName = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) SetReplica(v int32) *ListRayClusterResponseBodyRayClustersWorkerSpec {
	s.Replica = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) SetWorkerType(v string) *ListRayClusterResponseBodyRayClustersWorkerSpec {
	s.WorkerType = &v
	return s
}

func (s *ListRayClusterResponseBodyRayClustersWorkerSpec) Validate() error {
	return dara.Validate(s)
}
