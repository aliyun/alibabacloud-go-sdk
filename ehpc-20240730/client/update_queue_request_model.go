// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateQueueRequest
	GetClusterId() *string
	SetQueue(v *UpdateQueueRequestQueue) *UpdateQueueRequest
	GetQueue() *UpdateQueueRequestQueue
}

type UpdateQueueRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the queue to be updated.
	Queue *UpdateQueueRequestQueue `json:"Queue,omitempty" xml:"Queue,omitempty" type:"Struct"`
}

func (s UpdateQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQueueRequest) GoString() string {
	return s.String()
}

func (s *UpdateQueueRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateQueueRequest) GetQueue() *UpdateQueueRequestQueue {
	return s.Queue
}

func (s *UpdateQueueRequest) SetClusterId(v string) *UpdateQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateQueueRequest) SetQueue(v *UpdateQueueRequestQueue) *UpdateQueueRequest {
	s.Queue = v
	return s
}

func (s *UpdateQueueRequest) Validate() error {
	if s.Queue != nil {
		if err := s.Queue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateQueueRequestQueue struct {
	// The policy based on which instance types are selected for compute nodes during auto scale-outs. Valid values:
	//
	// 	- PriorityInstanceType
	//
	// example:
	//
	// PriorityInstanceType
	AllocationStrategy *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	// The hardware configurations of the compute nodes in the queue. Valid values of N: 1 to 10.
	ComputeNodes []*NodeTemplate `json:"ComputeNodes,omitempty" xml:"ComputeNodes,omitempty" type:"Repeated"`
	// Specifies whether to enable auto scale-in for the queue. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableScaleIn *bool `json:"EnableScaleIn,omitempty" xml:"EnableScaleIn,omitempty"`
	// Specifies whether to enable auto scale-out for the queue. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableScaleOut *bool `json:"EnableScaleOut,omitempty" xml:"EnableScaleOut,omitempty"`
	// The hostname prefix of the added compute nodes.
	//
	// example:
	//
	// compute
	HostnamePrefix *string `json:"HostnamePrefix,omitempty" xml:"HostnamePrefix,omitempty"`
	// The hostname suffix of the compute nodes in the queue.
	//
	// example:
	//
	// hpc
	HostnameSuffix *string `json:"HostnameSuffix,omitempty" xml:"HostnameSuffix,omitempty"`
	// The initial number of compute nodes in the queue.
	//
	// example:
	//
	// 0
	InitialCount *int32 `json:"InitialCount,omitempty" xml:"InitialCount,omitempty"`
	// The type of the network for interconnecting compute nodes in the queue.
	//
	// example:
	//
	// erdma
	InterConnect *string `json:"InterConnect,omitempty" xml:"InterConnect,omitempty"`
	// List of excluded compute nodes in the queue.
	KeepAliveNodes []*string `json:"KeepAliveNodes,omitempty" xml:"KeepAliveNodes,omitempty" type:"Repeated"`
	// The maximum number of compute nodes that the queue can contain.
	//
	// example:
	//
	// 1000
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The minimum number of compute nodes that are added to the queue during an automatic scale-out.
	//
	// example:
	//
	// 99
	MaxCountPerCycle *int64 `json:"MaxCountPerCycle,omitempty" xml:"MaxCountPerCycle,omitempty"`
	// The minimum number of compute nodes that the queue must contain.
	//
	// example:
	//
	// 0
	MinCount *int32 `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// The queue name.
	//
	// This parameter is required.
	//
	// example:
	//
	// comp
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The Resource Access Management (RAM) role that is assumed by compute nodes in the queue.
	//
	// example:
	//
	// AliyunECSInstanceForEHPCRole
	RamRole            *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	ReservedNodePoolId *string `json:"ReservedNodePoolId,omitempty" xml:"ReservedNodePoolId,omitempty"`
	// The vSwitches available for use by compute nodes in the queue.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s UpdateQueueRequestQueue) String() string {
	return dara.Prettify(s)
}

func (s UpdateQueueRequestQueue) GoString() string {
	return s.String()
}

func (s *UpdateQueueRequestQueue) GetAllocationStrategy() *string {
	return s.AllocationStrategy
}

func (s *UpdateQueueRequestQueue) GetComputeNodes() []*NodeTemplate {
	return s.ComputeNodes
}

func (s *UpdateQueueRequestQueue) GetEnableScaleIn() *bool {
	return s.EnableScaleIn
}

func (s *UpdateQueueRequestQueue) GetEnableScaleOut() *bool {
	return s.EnableScaleOut
}

func (s *UpdateQueueRequestQueue) GetHostnamePrefix() *string {
	return s.HostnamePrefix
}

func (s *UpdateQueueRequestQueue) GetHostnameSuffix() *string {
	return s.HostnameSuffix
}

func (s *UpdateQueueRequestQueue) GetInitialCount() *int32 {
	return s.InitialCount
}

func (s *UpdateQueueRequestQueue) GetInterConnect() *string {
	return s.InterConnect
}

func (s *UpdateQueueRequestQueue) GetKeepAliveNodes() []*string {
	return s.KeepAliveNodes
}

func (s *UpdateQueueRequestQueue) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *UpdateQueueRequestQueue) GetMaxCountPerCycle() *int64 {
	return s.MaxCountPerCycle
}

func (s *UpdateQueueRequestQueue) GetMinCount() *int32 {
	return s.MinCount
}

func (s *UpdateQueueRequestQueue) GetQueueName() *string {
	return s.QueueName
}

func (s *UpdateQueueRequestQueue) GetRamRole() *string {
	return s.RamRole
}

func (s *UpdateQueueRequestQueue) GetReservedNodePoolId() *string {
	return s.ReservedNodePoolId
}

func (s *UpdateQueueRequestQueue) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *UpdateQueueRequestQueue) SetAllocationStrategy(v string) *UpdateQueueRequestQueue {
	s.AllocationStrategy = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetComputeNodes(v []*NodeTemplate) *UpdateQueueRequestQueue {
	s.ComputeNodes = v
	return s
}

func (s *UpdateQueueRequestQueue) SetEnableScaleIn(v bool) *UpdateQueueRequestQueue {
	s.EnableScaleIn = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetEnableScaleOut(v bool) *UpdateQueueRequestQueue {
	s.EnableScaleOut = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetHostnamePrefix(v string) *UpdateQueueRequestQueue {
	s.HostnamePrefix = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetHostnameSuffix(v string) *UpdateQueueRequestQueue {
	s.HostnameSuffix = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetInitialCount(v int32) *UpdateQueueRequestQueue {
	s.InitialCount = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetInterConnect(v string) *UpdateQueueRequestQueue {
	s.InterConnect = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetKeepAliveNodes(v []*string) *UpdateQueueRequestQueue {
	s.KeepAliveNodes = v
	return s
}

func (s *UpdateQueueRequestQueue) SetMaxCount(v int32) *UpdateQueueRequestQueue {
	s.MaxCount = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetMaxCountPerCycle(v int64) *UpdateQueueRequestQueue {
	s.MaxCountPerCycle = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetMinCount(v int32) *UpdateQueueRequestQueue {
	s.MinCount = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetQueueName(v string) *UpdateQueueRequestQueue {
	s.QueueName = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetRamRole(v string) *UpdateQueueRequestQueue {
	s.RamRole = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetReservedNodePoolId(v string) *UpdateQueueRequestQueue {
	s.ReservedNodePoolId = &v
	return s
}

func (s *UpdateQueueRequestQueue) SetVSwitchIds(v []*string) *UpdateQueueRequestQueue {
	s.VSwitchIds = v
	return s
}

func (s *UpdateQueueRequestQueue) Validate() error {
	if s.ComputeNodes != nil {
		for _, item := range s.ComputeNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
