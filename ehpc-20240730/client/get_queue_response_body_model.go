// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueue(v *GetQueueResponseBodyQueue) *GetQueueResponseBody
	GetQueue() *GetQueueResponseBodyQueue
	SetRequestId(v string) *GetQueueResponseBody
	GetRequestId() *string
}

type GetQueueResponseBody struct {
	// The queue configurations.
	Queue *GetQueueResponseBodyQueue `json:"Queue,omitempty" xml:"Queue,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueueResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueueResponseBody) GetQueue() *GetQueueResponseBodyQueue {
	return s.Queue
}

func (s *GetQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueueResponseBody) SetQueue(v *GetQueueResponseBodyQueue) *GetQueueResponseBody {
	s.Queue = v
	return s
}

func (s *GetQueueResponseBody) SetRequestId(v string) *GetQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueueResponseBody) Validate() error {
	if s.Queue != nil {
		if err := s.Queue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQueueResponseBodyQueue struct {
	// The auto scale-out policy of the queue.
	//
	// example:
	//
	// PriorityInstanceType
	AllocationStrategy *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	// The hardware configurations of the compute nodes in the queue.
	ComputeNodes []*NodeTemplate `json:"ComputeNodes,omitempty" xml:"ComputeNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-01-01T12:05:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether auto scale-in is enabled for the queue. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableScaleIn *bool `json:"EnableScaleIn,omitempty" xml:"EnableScaleIn,omitempty"`
	// Indicates whether auto scale-out is enabled for the queue. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableScaleOut *bool `json:"EnableScaleOut,omitempty" xml:"EnableScaleOut,omitempty"`
	// The hostname prefix of the compute nodes in the queue.
	//
	// example:
	//
	// compute
	HostnamePrefix *string `json:"HostnamePrefix,omitempty" xml:"HostnamePrefix,omitempty"`
	// The hostname suffix of the compute nodes in the queue.
	//
	// example:
	//
	// demo
	HostnameSuffix *string `json:"HostnameSuffix,omitempty" xml:"HostnameSuffix,omitempty"`
	// The initial number of nodes in the queue.
	//
	// example:
	//
	// 0
	InitialCount *int32 `json:"InitialCount,omitempty" xml:"InitialCount,omitempty"`
	// The type of the network between compute nodes in the queue. Valid values:
	//
	// 	- vpc
	//
	// 	- eRDMA
	//
	// example:
	//
	// erdma
	InterConnect *string `json:"InterConnect,omitempty" xml:"InterConnect,omitempty"`
	// The nodes for which deletion protection is enabled in the queue.
	KeepAliveNodes []*string `json:"KeepAliveNodes,omitempty" xml:"KeepAliveNodes,omitempty" type:"Repeated"`
	// The maximum number of compute nodes that the queue can contain.
	//
	// example:
	//
	// 1000
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The minimum number of nodes that are delivered to the queue in each scale-out cycle.
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
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// Preset node pool ID.
	//
	// example:
	//
	// rnp-756vlp7a
	ReservedNodePoolId *string `json:"ReservedNodePoolId,omitempty" xml:"ReservedNodePoolId,omitempty"`
	// example:
	//
	// 2025-01-01T12:05:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The available vSwitches for compute nodes in the queue. Valid values of N: 1 to 5.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s GetQueueResponseBodyQueue) String() string {
	return dara.Prettify(s)
}

func (s GetQueueResponseBodyQueue) GoString() string {
	return s.String()
}

func (s *GetQueueResponseBodyQueue) GetAllocationStrategy() *string {
	return s.AllocationStrategy
}

func (s *GetQueueResponseBodyQueue) GetComputeNodes() []*NodeTemplate {
	return s.ComputeNodes
}

func (s *GetQueueResponseBodyQueue) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQueueResponseBodyQueue) GetEnableScaleIn() *bool {
	return s.EnableScaleIn
}

func (s *GetQueueResponseBodyQueue) GetEnableScaleOut() *bool {
	return s.EnableScaleOut
}

func (s *GetQueueResponseBodyQueue) GetHostnamePrefix() *string {
	return s.HostnamePrefix
}

func (s *GetQueueResponseBodyQueue) GetHostnameSuffix() *string {
	return s.HostnameSuffix
}

func (s *GetQueueResponseBodyQueue) GetInitialCount() *int32 {
	return s.InitialCount
}

func (s *GetQueueResponseBodyQueue) GetInterConnect() *string {
	return s.InterConnect
}

func (s *GetQueueResponseBodyQueue) GetKeepAliveNodes() []*string {
	return s.KeepAliveNodes
}

func (s *GetQueueResponseBodyQueue) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *GetQueueResponseBodyQueue) GetMaxCountPerCycle() *int64 {
	return s.MaxCountPerCycle
}

func (s *GetQueueResponseBodyQueue) GetMinCount() *int32 {
	return s.MinCount
}

func (s *GetQueueResponseBodyQueue) GetQueueName() *string {
	return s.QueueName
}

func (s *GetQueueResponseBodyQueue) GetRamRole() *string {
	return s.RamRole
}

func (s *GetQueueResponseBodyQueue) GetReservedNodePoolId() *string {
	return s.ReservedNodePoolId
}

func (s *GetQueueResponseBodyQueue) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetQueueResponseBodyQueue) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *GetQueueResponseBodyQueue) SetAllocationStrategy(v string) *GetQueueResponseBodyQueue {
	s.AllocationStrategy = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetComputeNodes(v []*NodeTemplate) *GetQueueResponseBodyQueue {
	s.ComputeNodes = v
	return s
}

func (s *GetQueueResponseBodyQueue) SetCreateTime(v string) *GetQueueResponseBodyQueue {
	s.CreateTime = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetEnableScaleIn(v bool) *GetQueueResponseBodyQueue {
	s.EnableScaleIn = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetEnableScaleOut(v bool) *GetQueueResponseBodyQueue {
	s.EnableScaleOut = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetHostnamePrefix(v string) *GetQueueResponseBodyQueue {
	s.HostnamePrefix = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetHostnameSuffix(v string) *GetQueueResponseBodyQueue {
	s.HostnameSuffix = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetInitialCount(v int32) *GetQueueResponseBodyQueue {
	s.InitialCount = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetInterConnect(v string) *GetQueueResponseBodyQueue {
	s.InterConnect = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetKeepAliveNodes(v []*string) *GetQueueResponseBodyQueue {
	s.KeepAliveNodes = v
	return s
}

func (s *GetQueueResponseBodyQueue) SetMaxCount(v int32) *GetQueueResponseBodyQueue {
	s.MaxCount = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetMaxCountPerCycle(v int64) *GetQueueResponseBodyQueue {
	s.MaxCountPerCycle = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetMinCount(v int32) *GetQueueResponseBodyQueue {
	s.MinCount = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetQueueName(v string) *GetQueueResponseBodyQueue {
	s.QueueName = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetRamRole(v string) *GetQueueResponseBodyQueue {
	s.RamRole = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetReservedNodePoolId(v string) *GetQueueResponseBodyQueue {
	s.ReservedNodePoolId = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetUpdateTime(v string) *GetQueueResponseBodyQueue {
	s.UpdateTime = &v
	return s
}

func (s *GetQueueResponseBodyQueue) SetVSwitchIds(v []*string) *GetQueueResponseBodyQueue {
	s.VSwitchIds = v
	return s
}

func (s *GetQueueResponseBodyQueue) Validate() error {
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
