// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRayClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateRayClusterRequest
	GetDescription() *string
	SetDisplayReleaseVersion(v string) *UpdateRayClusterRequest
	GetDisplayReleaseVersion() *string
	SetExtraParam(v string) *UpdateRayClusterRequest
	GetExtraParam() *string
	SetHeadSpec(v *UpdateRayClusterRequestHeadSpec) *UpdateRayClusterRequest
	GetHeadSpec() *UpdateRayClusterRequestHeadSpec
	SetName(v string) *UpdateRayClusterRequest
	GetName() *string
	SetNetworkServiceName(v string) *UpdateRayClusterRequest
	GetNetworkServiceName() *string
	SetWorkerSpec(v []*UpdateRayClusterRequestWorkerSpec) *UpdateRayClusterRequest
	GetWorkerSpec() []*UpdateRayClusterRequestWorkerSpec
}

type UpdateRayClusterRequest struct {
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
	ExtraParam *string                          `json:"extraParam,omitempty" xml:"extraParam,omitempty"`
	HeadSpec   *UpdateRayClusterRequestHeadSpec `json:"headSpec,omitempty" xml:"headSpec,omitempty" type:"Struct"`
	// example:
	//
	// myRayCluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// vpc
	NetworkServiceName *string                              `json:"networkServiceName,omitempty" xml:"networkServiceName,omitempty"`
	WorkerSpec         []*UpdateRayClusterRequestWorkerSpec `json:"workerSpec,omitempty" xml:"workerSpec,omitempty" type:"Repeated"`
}

func (s UpdateRayClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRayClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateRayClusterRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRayClusterRequest) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *UpdateRayClusterRequest) GetExtraParam() *string {
	return s.ExtraParam
}

func (s *UpdateRayClusterRequest) GetHeadSpec() *UpdateRayClusterRequestHeadSpec {
	return s.HeadSpec
}

func (s *UpdateRayClusterRequest) GetName() *string {
	return s.Name
}

func (s *UpdateRayClusterRequest) GetNetworkServiceName() *string {
	return s.NetworkServiceName
}

func (s *UpdateRayClusterRequest) GetWorkerSpec() []*UpdateRayClusterRequestWorkerSpec {
	return s.WorkerSpec
}

func (s *UpdateRayClusterRequest) SetDescription(v string) *UpdateRayClusterRequest {
	s.Description = &v
	return s
}

func (s *UpdateRayClusterRequest) SetDisplayReleaseVersion(v string) *UpdateRayClusterRequest {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *UpdateRayClusterRequest) SetExtraParam(v string) *UpdateRayClusterRequest {
	s.ExtraParam = &v
	return s
}

func (s *UpdateRayClusterRequest) SetHeadSpec(v *UpdateRayClusterRequestHeadSpec) *UpdateRayClusterRequest {
	s.HeadSpec = v
	return s
}

func (s *UpdateRayClusterRequest) SetName(v string) *UpdateRayClusterRequest {
	s.Name = &v
	return s
}

func (s *UpdateRayClusterRequest) SetNetworkServiceName(v string) *UpdateRayClusterRequest {
	s.NetworkServiceName = &v
	return s
}

func (s *UpdateRayClusterRequest) SetWorkerSpec(v []*UpdateRayClusterRequestWorkerSpec) *UpdateRayClusterRequest {
	s.WorkerSpec = v
	return s
}

func (s *UpdateRayClusterRequest) Validate() error {
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

type UpdateRayClusterRequestHeadSpec struct {
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
}

func (s UpdateRayClusterRequestHeadSpec) String() string {
	return dara.Prettify(s)
}

func (s UpdateRayClusterRequestHeadSpec) GoString() string {
	return s.String()
}

func (s *UpdateRayClusterRequestHeadSpec) GetCpu() *string {
	return s.Cpu
}

func (s *UpdateRayClusterRequestHeadSpec) GetEnableAutoScaling() *bool {
	return s.EnableAutoScaling
}

func (s *UpdateRayClusterRequestHeadSpec) GetIdleTimeoutSeconds() *int32 {
	return s.IdleTimeoutSeconds
}

func (s *UpdateRayClusterRequestHeadSpec) GetMemory() *string {
	return s.Memory
}

func (s *UpdateRayClusterRequestHeadSpec) GetQueueName() *string {
	return s.QueueName
}

func (s *UpdateRayClusterRequestHeadSpec) SetCpu(v string) *UpdateRayClusterRequestHeadSpec {
	s.Cpu = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) SetEnableAutoScaling(v bool) *UpdateRayClusterRequestHeadSpec {
	s.EnableAutoScaling = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) SetIdleTimeoutSeconds(v int32) *UpdateRayClusterRequestHeadSpec {
	s.IdleTimeoutSeconds = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) SetMemory(v string) *UpdateRayClusterRequestHeadSpec {
	s.Memory = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) SetQueueName(v string) *UpdateRayClusterRequestHeadSpec {
	s.QueueName = &v
	return s
}

func (s *UpdateRayClusterRequestHeadSpec) Validate() error {
	return dara.Validate(s)
}

type UpdateRayClusterRequestWorkerSpec struct {
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

func (s UpdateRayClusterRequestWorkerSpec) String() string {
	return dara.Prettify(s)
}

func (s UpdateRayClusterRequestWorkerSpec) GoString() string {
	return s.String()
}

func (s *UpdateRayClusterRequestWorkerSpec) GetCpu() *string {
	return s.Cpu
}

func (s *UpdateRayClusterRequestWorkerSpec) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateRayClusterRequestWorkerSpec) GetMaxReplica() *int32 {
	return s.MaxReplica
}

func (s *UpdateRayClusterRequestWorkerSpec) GetMemory() *string {
	return s.Memory
}

func (s *UpdateRayClusterRequestWorkerSpec) GetMinReplica() *int32 {
	return s.MinReplica
}

func (s *UpdateRayClusterRequestWorkerSpec) GetQueueName() *string {
	return s.QueueName
}

func (s *UpdateRayClusterRequestWorkerSpec) GetReplica() *int32 {
	return s.Replica
}

func (s *UpdateRayClusterRequestWorkerSpec) GetWorkerType() *string {
	return s.WorkerType
}

func (s *UpdateRayClusterRequestWorkerSpec) SetCpu(v string) *UpdateRayClusterRequestWorkerSpec {
	s.Cpu = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetGroupName(v string) *UpdateRayClusterRequestWorkerSpec {
	s.GroupName = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetMaxReplica(v int32) *UpdateRayClusterRequestWorkerSpec {
	s.MaxReplica = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetMemory(v string) *UpdateRayClusterRequestWorkerSpec {
	s.Memory = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetMinReplica(v int32) *UpdateRayClusterRequestWorkerSpec {
	s.MinReplica = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetQueueName(v string) *UpdateRayClusterRequestWorkerSpec {
	s.QueueName = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetReplica(v int32) *UpdateRayClusterRequestWorkerSpec {
	s.Replica = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) SetWorkerType(v string) *UpdateRayClusterRequestWorkerSpec {
	s.WorkerType = &v
	return s
}

func (s *UpdateRayClusterRequestWorkerSpec) Validate() error {
	return dara.Validate(s)
}
