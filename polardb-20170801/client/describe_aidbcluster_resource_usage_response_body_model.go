// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterResourceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeAIDBClusterResourceUsageResponseBody
	GetCreateTime() *string
	SetDBClusterId(v string) *DescribeAIDBClusterResourceUsageResponseBody
	GetDBClusterId() *string
	SetEngine(v string) *DescribeAIDBClusterResourceUsageResponseBody
	GetEngine() *string
	SetEngineVersion(v string) *DescribeAIDBClusterResourceUsageResponseBody
	GetEngineVersion() *string
	SetNodeItems(v []*DescribeAIDBClusterResourceUsageResponseBodyNodeItems) *DescribeAIDBClusterResourceUsageResponseBody
	GetNodeItems() []*DescribeAIDBClusterResourceUsageResponseBodyNodeItems
	SetRequestId(v string) *DescribeAIDBClusterResourceUsageResponseBody
	GetRequestId() *string
}

type DescribeAIDBClusterResourceUsageResponseBody struct {
	// The time when the cluster was created, in ISO 8601 UTC format.
	//
	// example:
	//
	// 2021-09-13T15:45:27Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the AI database cluster.
	//
	// example:
	//
	// pc-2zeu0t5v0p98*****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine of the cluster.
	//
	// example:
	//
	// AI
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the cluster.
	//
	// example:
	//
	// 3.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The list of resource usage details for the nodes in the cluster.
	NodeItems []*DescribeAIDBClusterResourceUsageResponseBodyNodeItems `json:"NodeItems,omitempty" xml:"NodeItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAIDBClusterResourceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterResourceUsageResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAIDBClusterResourceUsageResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAIDBClusterResourceUsageResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeAIDBClusterResourceUsageResponseBody) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeAIDBClusterResourceUsageResponseBody) GetNodeItems() []*DescribeAIDBClusterResourceUsageResponseBodyNodeItems {
	return s.NodeItems
}

func (s *DescribeAIDBClusterResourceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAIDBClusterResourceUsageResponseBody) SetCreateTime(v string) *DescribeAIDBClusterResourceUsageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBody) SetDBClusterId(v string) *DescribeAIDBClusterResourceUsageResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBody) SetEngine(v string) *DescribeAIDBClusterResourceUsageResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBody) SetEngineVersion(v string) *DescribeAIDBClusterResourceUsageResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBody) SetNodeItems(v []*DescribeAIDBClusterResourceUsageResponseBodyNodeItems) *DescribeAIDBClusterResourceUsageResponseBody {
	s.NodeItems = v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBody) SetRequestId(v string) *DescribeAIDBClusterResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBody) Validate() error {
	if s.NodeItems != nil {
		for _, item := range s.NodeItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAIDBClusterResourceUsageResponseBodyNodeItems struct {
	// The total CPU resources of the node.
	//
	// example:
	//
	// 4
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// The amount of CPU resources used on the node.
	//
	// example:
	//
	// 2
	CPUUsed *string `json:"CPUUsed,omitempty" xml:"CPUUsed,omitempty"`
	// The ID of the database node.
	//
	// example:
	//
	// pi-2ze7sztqv1*****
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The total GPU resources of the node.
	//
	// example:
	//
	// 1
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// The GPU model of the node.
	//
	// example:
	//
	// NVIDIA A10
	GPUModel *string `json:"GPUModel,omitempty" xml:"GPUModel,omitempty"`
	// The amount of GPU resources used on the node.
	//
	// example:
	//
	// 0
	GPUUsed *string `json:"GPUUsed,omitempty" xml:"GPUUsed,omitempty"`
	// The total memory of the node.
	//
	// example:
	//
	// 16Gi
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The amount of memory used on the node.
	//
	// example:
	//
	// 8Gi
	MemoryUsed *string `json:"MemoryUsed,omitempty" xml:"MemoryUsed,omitempty"`
	// The list of resource information for the pods on the node.
	PodItems []*DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems `json:"PodItems,omitempty" xml:"PodItems,omitempty" type:"Repeated"`
	// The ID of the virtual node corresponding to the database node.
	//
	// example:
	//
	// vnode-001
	VNodeId *string `json:"VNodeId,omitempty" xml:"VNodeId,omitempty"`
}

func (s DescribeAIDBClusterResourceUsageResponseBodyNodeItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterResourceUsageResponseBodyNodeItems) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) GetCPU() *string {
	return s.CPU
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) GetCPUUsed() *string {
	return s.CPUUsed
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) GetGPU() *string {
	return s.GPU
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) GetGPUModel() *string {
	return s.GPUModel
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) GetGPUUsed() *string {
	return s.GPUUsed
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) GetMemory() *string {
	return s.Memory
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) GetMemoryUsed() *string {
	return s.MemoryUsed
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) GetPodItems() []*DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems {
	return s.PodItems
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) GetVNodeId() *string {
	return s.VNodeId
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) SetCPU(v string) *DescribeAIDBClusterResourceUsageResponseBodyNodeItems {
	s.CPU = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) SetCPUUsed(v string) *DescribeAIDBClusterResourceUsageResponseBodyNodeItems {
	s.CPUUsed = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) SetDBNodeId(v string) *DescribeAIDBClusterResourceUsageResponseBodyNodeItems {
	s.DBNodeId = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) SetGPU(v string) *DescribeAIDBClusterResourceUsageResponseBodyNodeItems {
	s.GPU = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) SetGPUModel(v string) *DescribeAIDBClusterResourceUsageResponseBodyNodeItems {
	s.GPUModel = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) SetGPUUsed(v string) *DescribeAIDBClusterResourceUsageResponseBodyNodeItems {
	s.GPUUsed = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) SetMemory(v string) *DescribeAIDBClusterResourceUsageResponseBodyNodeItems {
	s.Memory = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) SetMemoryUsed(v string) *DescribeAIDBClusterResourceUsageResponseBodyNodeItems {
	s.MemoryUsed = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) SetPodItems(v []*DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) *DescribeAIDBClusterResourceUsageResponseBodyNodeItems {
	s.PodItems = v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) SetVNodeId(v string) *DescribeAIDBClusterResourceUsageResponseBodyNodeItems {
	s.VNodeId = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItems) Validate() error {
	if s.PodItems != nil {
		for _, item := range s.PodItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems struct {
	// The time when the pod was created, in ISO 8601 UTC format.
	//
	// example:
	//
	// 2021-09-13T15:45:27Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the pod.
	//
	// example:
	//
	// ai-db-pod-001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace to which the pod belongs.
	//
	// example:
	//
	// default
	NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	// The IP address of the pod.
	//
	// example:
	//
	// 10.0.0.12
	PodIP *string `json:"PodIP,omitempty" xml:"PodIP,omitempty"`
	// The resource configuration requested by the pod.
	//
	// example:
	//
	// cpu=1,memory=2Gi
	ResourceRequest *string `json:"ResourceRequest,omitempty" xml:"ResourceRequest,omitempty"`
	// The running status of the pod.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) GetName() *string {
	return s.Name
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) GetNameSpace() *string {
	return s.NameSpace
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) GetPodIP() *string {
	return s.PodIP
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) GetResourceRequest() *string {
	return s.ResourceRequest
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) SetCreateTime(v string) *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) SetName(v string) *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems {
	s.Name = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) SetNameSpace(v string) *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems {
	s.NameSpace = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) SetPodIP(v string) *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems {
	s.PodIP = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) SetResourceRequest(v string) *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems {
	s.ResourceRequest = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) SetStatus(v string) *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems {
	s.Status = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageResponseBodyNodeItemsPodItems) Validate() error {
	return dara.Validate(s)
}
