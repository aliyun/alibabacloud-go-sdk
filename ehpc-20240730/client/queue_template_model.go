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
	// example:
	//
	// PriorityInstanceType
	AllocationStrategy *string         `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	ComputeNodes       []*NodeTemplate `json:"ComputeNodes,omitempty" xml:"ComputeNodes,omitempty" type:"Repeated"`
	// example:
	//
	// true
	EnableScaleIn *bool `json:"EnableScaleIn,omitempty" xml:"EnableScaleIn,omitempty"`
	// example:
	//
	// true
	EnableScaleOut *bool `json:"EnableScaleOut,omitempty" xml:"EnableScaleOut,omitempty"`
	// example:
	//
	// compute
	HostnamePrefix *string `json:"HostnamePrefix,omitempty" xml:"HostnamePrefix,omitempty"`
	HostnameSuffix *string `json:"HostnameSuffix,omitempty" xml:"HostnameSuffix,omitempty"`
	// example:
	//
	// 0
	InitialCount *int32 `json:"InitialCount,omitempty" xml:"InitialCount,omitempty"`
	// example:
	//
	// erdma
	InterConnect   *string   `json:"InterConnect,omitempty" xml:"InterConnect,omitempty"`
	KeepAliveNodes []*string `json:"KeepAliveNodes,omitempty" xml:"KeepAliveNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 1000
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// example:
	//
	// 99
	MaxCountPerCycle *int64 `json:"MaxCountPerCycle,omitempty" xml:"MaxCountPerCycle,omitempty"`
	// example:
	//
	// 0
	MinCount *int32 `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// comp
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// example:
	//
	// AliyunECSInstanceForEHPCRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// example:
	//
	// rnp-756vlp7a
	ReservedNodePoolId *string   `json:"ReservedNodePoolId,omitempty" xml:"ReservedNodePoolId,omitempty"`
	VSwitchIds         []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
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
