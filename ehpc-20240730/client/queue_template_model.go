// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueueTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationStrategy(v string) *QueueTemplate
	GetAllocationStrategy() *string
	SetComputeNodes(v []*NodeTemplate) *QueueTemplate
	GetComputeNodes() []*NodeTemplate
	SetEnableScaleIn(v bool) *QueueTemplate
	GetEnableScaleIn() *bool
	SetEnableScaleOut(v bool) *QueueTemplate
	GetEnableScaleOut() *bool
	SetHostnamePrefix(v string) *QueueTemplate
	GetHostnamePrefix() *string
	SetHostnameSuffix(v string) *QueueTemplate
	GetHostnameSuffix() *string
	SetInitialCount(v int32) *QueueTemplate
	GetInitialCount() *int32
	SetInterConnect(v string) *QueueTemplate
	GetInterConnect() *string
	SetKeepAliveNodes(v []*string) *QueueTemplate
	GetKeepAliveNodes() []*string
	SetMaxCount(v int32) *QueueTemplate
	GetMaxCount() *int32
	SetMaxCountPerCycle(v int64) *QueueTemplate
	GetMaxCountPerCycle() *int64
	SetMinCount(v int32) *QueueTemplate
	GetMinCount() *int32
	SetQueueName(v string) *QueueTemplate
	GetQueueName() *string
	SetRamRole(v string) *QueueTemplate
	GetRamRole() *string
	SetReservedNodePoolId(v string) *QueueTemplate
	GetReservedNodePoolId() *string
	SetVSwitchIds(v []*string) *QueueTemplate
	GetVSwitchIds() []*string
}

type QueueTemplate struct {
	// The auto scale-out policy for the queue.
	//
	// example:
	//
	// PriorityInstanceType
	AllocationStrategy *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	// The list of hardware configurations for the compute nodes in the queue. You can specify 0 to 10 configurations.
	ComputeNodes []*NodeTemplate `json:"ComputeNodes,omitempty" xml:"ComputeNodes,omitempty" type:"Repeated"`
	// Indicates whether to enable auto scale-in for the queue. Valid values:
	//
	// - true: enabled
	//
	// - false: disabled
	//
	// example:
	//
	// true
	EnableScaleIn *bool `json:"EnableScaleIn,omitempty" xml:"EnableScaleIn,omitempty"`
	// Indicates whether to enable auto scale-out for the queue. Valid values:
	//
	// - true: enabled
	//
	// - false: disabled
	//
	// example:
	//
	// true
	EnableScaleOut *bool `json:"EnableScaleOut,omitempty" xml:"EnableScaleOut,omitempty"`
	// The hostname prefix of the compute nodes in the queue.
	//
	// 	Notice:
	//
	// The prefix can be up to 8 characters in length for Windows operating systems and up to 32 characters in length for Linux operating systems. The prefix can contain only lowercase letters, digits, and hyphens (-).
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
	// The initial number of compute nodes in the queue.
	//
	// example:
	//
	// 0
	InitialCount *int32 `json:"InitialCount,omitempty" xml:"InitialCount,omitempty"`
	// The network type of the compute nodes in the queue. Valid values:
	//
	// - vpc
	//
	// - eRDMA
	//
	// example:
	//
	// erdma
	InterConnect *string `json:"InterConnect,omitempty" xml:"InterConnect,omitempty"`
	// The list of nodes in the queue that have deletion protection enabled.
	KeepAliveNodes []*string `json:"KeepAliveNodes,omitempty" xml:"KeepAliveNodes,omitempty" type:"Repeated"`
	// The maximum number of compute nodes in the queue.
	//
	// example:
	//
	// 1000
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The maximum number of compute nodes that can be scaled out in each cycle.
	//
	// example:
	//
	// 99
	MaxCountPerCycle *int64 `json:"MaxCountPerCycle,omitempty" xml:"MaxCountPerCycle,omitempty"`
	// The minimum number of compute nodes in the queue.
	//
	// example:
	//
	// 0
	MinCount *int32 `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// The queue name. The name must be 1 to 15 characters long. It can contain letters from the Unicode letter category, such as English letters and digits, and periods (.).
	//
	// This parameter is required.
	//
	// example:
	//
	// comp
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The name of the instance role attached to the compute nodes in the queue.
	//
	// example:
	//
	// AliyunECSInstanceForEHPCRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// The ID of the reserved node pool that the queue uses.
	//
	// > If this parameter is specified, allocatable nodes from the reserved node pool are used to create compute nodes. The `VSwitchIds`, `HostnamePrefix`, and `HostnameSuffix` parameters are ignored.
	//
	// example:
	//
	// rnp-756vlp7a
	ReservedNodePoolId *string `json:"ReservedNodePoolId,omitempty" xml:"ReservedNodePoolId,omitempty"`
	// A list of virtual switches available to the compute nodes in the queue. You can specify 1 to 5 virtual switches.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s QueueTemplate) String() string {
	return dara.Prettify(s)
}

func (s QueueTemplate) GoString() string {
	return s.String()
}

func (s *QueueTemplate) GetAllocationStrategy() *string {
	return s.AllocationStrategy
}

func (s *QueueTemplate) GetComputeNodes() []*NodeTemplate {
	return s.ComputeNodes
}

func (s *QueueTemplate) GetEnableScaleIn() *bool {
	return s.EnableScaleIn
}

func (s *QueueTemplate) GetEnableScaleOut() *bool {
	return s.EnableScaleOut
}

func (s *QueueTemplate) GetHostnamePrefix() *string {
	return s.HostnamePrefix
}

func (s *QueueTemplate) GetHostnameSuffix() *string {
	return s.HostnameSuffix
}

func (s *QueueTemplate) GetInitialCount() *int32 {
	return s.InitialCount
}

func (s *QueueTemplate) GetInterConnect() *string {
	return s.InterConnect
}

func (s *QueueTemplate) GetKeepAliveNodes() []*string {
	return s.KeepAliveNodes
}

func (s *QueueTemplate) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *QueueTemplate) GetMaxCountPerCycle() *int64 {
	return s.MaxCountPerCycle
}

func (s *QueueTemplate) GetMinCount() *int32 {
	return s.MinCount
}

func (s *QueueTemplate) GetQueueName() *string {
	return s.QueueName
}

func (s *QueueTemplate) GetRamRole() *string {
	return s.RamRole
}

func (s *QueueTemplate) GetReservedNodePoolId() *string {
	return s.ReservedNodePoolId
}

func (s *QueueTemplate) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *QueueTemplate) SetAllocationStrategy(v string) *QueueTemplate {
	s.AllocationStrategy = &v
	return s
}

func (s *QueueTemplate) SetComputeNodes(v []*NodeTemplate) *QueueTemplate {
	s.ComputeNodes = v
	return s
}

func (s *QueueTemplate) SetEnableScaleIn(v bool) *QueueTemplate {
	s.EnableScaleIn = &v
	return s
}

func (s *QueueTemplate) SetEnableScaleOut(v bool) *QueueTemplate {
	s.EnableScaleOut = &v
	return s
}

func (s *QueueTemplate) SetHostnamePrefix(v string) *QueueTemplate {
	s.HostnamePrefix = &v
	return s
}

func (s *QueueTemplate) SetHostnameSuffix(v string) *QueueTemplate {
	s.HostnameSuffix = &v
	return s
}

func (s *QueueTemplate) SetInitialCount(v int32) *QueueTemplate {
	s.InitialCount = &v
	return s
}

func (s *QueueTemplate) SetInterConnect(v string) *QueueTemplate {
	s.InterConnect = &v
	return s
}

func (s *QueueTemplate) SetKeepAliveNodes(v []*string) *QueueTemplate {
	s.KeepAliveNodes = v
	return s
}

func (s *QueueTemplate) SetMaxCount(v int32) *QueueTemplate {
	s.MaxCount = &v
	return s
}

func (s *QueueTemplate) SetMaxCountPerCycle(v int64) *QueueTemplate {
	s.MaxCountPerCycle = &v
	return s
}

func (s *QueueTemplate) SetMinCount(v int32) *QueueTemplate {
	s.MinCount = &v
	return s
}

func (s *QueueTemplate) SetQueueName(v string) *QueueTemplate {
	s.QueueName = &v
	return s
}

func (s *QueueTemplate) SetRamRole(v string) *QueueTemplate {
	s.RamRole = &v
	return s
}

func (s *QueueTemplate) SetReservedNodePoolId(v string) *QueueTemplate {
	s.ReservedNodePoolId = &v
	return s
}

func (s *QueueTemplate) SetVSwitchIds(v []*string) *QueueTemplate {
	s.VSwitchIds = v
	return s
}

func (s *QueueTemplate) Validate() error {
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
