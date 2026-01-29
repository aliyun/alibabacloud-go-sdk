// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRayClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateRayClusterRequest
	GetDescription() *string
	SetDisplayReleaseVersion(v string) *CreateRayClusterRequest
	GetDisplayReleaseVersion() *string
	SetExtraParam(v string) *CreateRayClusterRequest
	GetExtraParam() *string
	SetHeadSpec(v *CreateRayClusterRequestHeadSpec) *CreateRayClusterRequest
	GetHeadSpec() *CreateRayClusterRequestHeadSpec
	SetName(v string) *CreateRayClusterRequest
	GetName() *string
	SetNetworkServiceName(v string) *CreateRayClusterRequest
	GetNetworkServiceName() *string
	SetWorkerSpec(v []*CreateRayClusterRequestWorkerSpec) *CreateRayClusterRequest
	GetWorkerSpec() []*CreateRayClusterRequestWorkerSpec
}

type CreateRayClusterRequest struct {
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
	HeadSpec   *CreateRayClusterRequestHeadSpec `json:"headSpec,omitempty" xml:"headSpec,omitempty" type:"Struct"`
	// example:
	//
	// testRayCluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// vpc
	NetworkServiceName *string                              `json:"networkServiceName,omitempty" xml:"networkServiceName,omitempty"`
	WorkerSpec         []*CreateRayClusterRequestWorkerSpec `json:"workerSpec,omitempty" xml:"workerSpec,omitempty" type:"Repeated"`
}

func (s CreateRayClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRayClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateRayClusterRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRayClusterRequest) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *CreateRayClusterRequest) GetExtraParam() *string {
	return s.ExtraParam
}

func (s *CreateRayClusterRequest) GetHeadSpec() *CreateRayClusterRequestHeadSpec {
	return s.HeadSpec
}

func (s *CreateRayClusterRequest) GetName() *string {
	return s.Name
}

func (s *CreateRayClusterRequest) GetNetworkServiceName() *string {
	return s.NetworkServiceName
}

func (s *CreateRayClusterRequest) GetWorkerSpec() []*CreateRayClusterRequestWorkerSpec {
	return s.WorkerSpec
}

func (s *CreateRayClusterRequest) SetDescription(v string) *CreateRayClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateRayClusterRequest) SetDisplayReleaseVersion(v string) *CreateRayClusterRequest {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *CreateRayClusterRequest) SetExtraParam(v string) *CreateRayClusterRequest {
	s.ExtraParam = &v
	return s
}

func (s *CreateRayClusterRequest) SetHeadSpec(v *CreateRayClusterRequestHeadSpec) *CreateRayClusterRequest {
	s.HeadSpec = v
	return s
}

func (s *CreateRayClusterRequest) SetName(v string) *CreateRayClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateRayClusterRequest) SetNetworkServiceName(v string) *CreateRayClusterRequest {
	s.NetworkServiceName = &v
	return s
}

func (s *CreateRayClusterRequest) SetWorkerSpec(v []*CreateRayClusterRequestWorkerSpec) *CreateRayClusterRequest {
	s.WorkerSpec = v
	return s
}

func (s *CreateRayClusterRequest) Validate() error {
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

type CreateRayClusterRequestHeadSpec struct {
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

func (s CreateRayClusterRequestHeadSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateRayClusterRequestHeadSpec) GoString() string {
	return s.String()
}

func (s *CreateRayClusterRequestHeadSpec) GetCpu() *string {
	return s.Cpu
}

func (s *CreateRayClusterRequestHeadSpec) GetEnableAutoScaling() *bool {
	return s.EnableAutoScaling
}

func (s *CreateRayClusterRequestHeadSpec) GetIdleTimeoutSeconds() *int32 {
	return s.IdleTimeoutSeconds
}

func (s *CreateRayClusterRequestHeadSpec) GetMemory() *string {
	return s.Memory
}

func (s *CreateRayClusterRequestHeadSpec) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateRayClusterRequestHeadSpec) SetCpu(v string) *CreateRayClusterRequestHeadSpec {
	s.Cpu = &v
	return s
}

func (s *CreateRayClusterRequestHeadSpec) SetEnableAutoScaling(v bool) *CreateRayClusterRequestHeadSpec {
	s.EnableAutoScaling = &v
	return s
}

func (s *CreateRayClusterRequestHeadSpec) SetIdleTimeoutSeconds(v int32) *CreateRayClusterRequestHeadSpec {
	s.IdleTimeoutSeconds = &v
	return s
}

func (s *CreateRayClusterRequestHeadSpec) SetMemory(v string) *CreateRayClusterRequestHeadSpec {
	s.Memory = &v
	return s
}

func (s *CreateRayClusterRequestHeadSpec) SetQueueName(v string) *CreateRayClusterRequestHeadSpec {
	s.QueueName = &v
	return s
}

func (s *CreateRayClusterRequestHeadSpec) Validate() error {
	return dara.Validate(s)
}

type CreateRayClusterRequestWorkerSpec struct {
	// example:
	//
	// 4
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

func (s CreateRayClusterRequestWorkerSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateRayClusterRequestWorkerSpec) GoString() string {
	return s.String()
}

func (s *CreateRayClusterRequestWorkerSpec) GetCpu() *string {
	return s.Cpu
}

func (s *CreateRayClusterRequestWorkerSpec) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateRayClusterRequestWorkerSpec) GetMaxReplica() *int32 {
	return s.MaxReplica
}

func (s *CreateRayClusterRequestWorkerSpec) GetMemory() *string {
	return s.Memory
}

func (s *CreateRayClusterRequestWorkerSpec) GetMinReplica() *int32 {
	return s.MinReplica
}

func (s *CreateRayClusterRequestWorkerSpec) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateRayClusterRequestWorkerSpec) GetReplica() *int32 {
	return s.Replica
}

func (s *CreateRayClusterRequestWorkerSpec) GetWorkerType() *string {
	return s.WorkerType
}

func (s *CreateRayClusterRequestWorkerSpec) SetCpu(v string) *CreateRayClusterRequestWorkerSpec {
	s.Cpu = &v
	return s
}

func (s *CreateRayClusterRequestWorkerSpec) SetGroupName(v string) *CreateRayClusterRequestWorkerSpec {
	s.GroupName = &v
	return s
}

func (s *CreateRayClusterRequestWorkerSpec) SetMaxReplica(v int32) *CreateRayClusterRequestWorkerSpec {
	s.MaxReplica = &v
	return s
}

func (s *CreateRayClusterRequestWorkerSpec) SetMemory(v string) *CreateRayClusterRequestWorkerSpec {
	s.Memory = &v
	return s
}

func (s *CreateRayClusterRequestWorkerSpec) SetMinReplica(v int32) *CreateRayClusterRequestWorkerSpec {
	s.MinReplica = &v
	return s
}

func (s *CreateRayClusterRequestWorkerSpec) SetQueueName(v string) *CreateRayClusterRequestWorkerSpec {
	s.QueueName = &v
	return s
}

func (s *CreateRayClusterRequestWorkerSpec) SetReplica(v int32) *CreateRayClusterRequestWorkerSpec {
	s.Replica = &v
	return s
}

func (s *CreateRayClusterRequestWorkerSpec) SetWorkerType(v string) *CreateRayClusterRequestWorkerSpec {
	s.WorkerType = &v
	return s
}

func (s *CreateRayClusterRequestWorkerSpec) Validate() error {
	return dara.Validate(s)
}
