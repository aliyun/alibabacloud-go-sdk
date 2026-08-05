// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodePodInfo interface {
	dara.Model
	String() string
	GoString() string
	SetGPUIndexes(v []*int32) *NodePodInfo
	GetGPUIndexes() []*int32
	SetPhase(v string) *NodePodInfo
	GetPhase() *string
	SetPodIP(v string) *NodePodInfo
	GetPodIP() *string
	SetPodName(v string) *NodePodInfo
	GetPodName() *string
	SetPodNamespace(v string) *NodePodInfo
	GetPodNamespace() *string
	SetResourceSpec(v *ResourceAmount) *NodePodInfo
	GetResourceSpec() *ResourceAmount
	SetWorkloadId(v string) *NodePodInfo
	GetWorkloadId() *string
	SetWorkloadType(v string) *NodePodInfo
	GetWorkloadType() *string
}

type NodePodInfo struct {
	// The occupied GPU indexes.
	GPUIndexes []*int32 `json:"GPUIndexes,omitempty" xml:"GPUIndexes,omitempty" type:"Repeated"`
	// The Pod status.
	//
	// example:
	//
	// Running
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The IP address of the Pod.
	//
	// example:
	//
	// 192.168.2.2
	PodIP *string `json:"PodIP,omitempty" xml:"PodIP,omitempty"`
	// The Pod name.
	//
	// example:
	//
	// test
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// The namespace of the Pod.
	//
	// example:
	//
	// test
	PodNamespace *string `json:"PodNamespace,omitempty" xml:"PodNamespace,omitempty"`
	// The resource usage information.
	ResourceSpec *ResourceAmount `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
	// The task ID or service ID.
	//
	// example:
	//
	// dlc19de9s6vn3acr
	WorkloadId *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// The sub-product that the Pod belongs to.
	//
	// example:
	//
	// dlc
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
}

func (s NodePodInfo) String() string {
	return dara.Prettify(s)
}

func (s NodePodInfo) GoString() string {
	return s.String()
}

func (s *NodePodInfo) GetGPUIndexes() []*int32 {
	return s.GPUIndexes
}

func (s *NodePodInfo) GetPhase() *string {
	return s.Phase
}

func (s *NodePodInfo) GetPodIP() *string {
	return s.PodIP
}

func (s *NodePodInfo) GetPodName() *string {
	return s.PodName
}

func (s *NodePodInfo) GetPodNamespace() *string {
	return s.PodNamespace
}

func (s *NodePodInfo) GetResourceSpec() *ResourceAmount {
	return s.ResourceSpec
}

func (s *NodePodInfo) GetWorkloadId() *string {
	return s.WorkloadId
}

func (s *NodePodInfo) GetWorkloadType() *string {
	return s.WorkloadType
}

func (s *NodePodInfo) SetGPUIndexes(v []*int32) *NodePodInfo {
	s.GPUIndexes = v
	return s
}

func (s *NodePodInfo) SetPhase(v string) *NodePodInfo {
	s.Phase = &v
	return s
}

func (s *NodePodInfo) SetPodIP(v string) *NodePodInfo {
	s.PodIP = &v
	return s
}

func (s *NodePodInfo) SetPodName(v string) *NodePodInfo {
	s.PodName = &v
	return s
}

func (s *NodePodInfo) SetPodNamespace(v string) *NodePodInfo {
	s.PodNamespace = &v
	return s
}

func (s *NodePodInfo) SetResourceSpec(v *ResourceAmount) *NodePodInfo {
	s.ResourceSpec = v
	return s
}

func (s *NodePodInfo) SetWorkloadId(v string) *NodePodInfo {
	s.WorkloadId = &v
	return s
}

func (s *NodePodInfo) SetWorkloadType(v string) *NodePodInfo {
	s.WorkloadType = &v
	return s
}

func (s *NodePodInfo) Validate() error {
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}
