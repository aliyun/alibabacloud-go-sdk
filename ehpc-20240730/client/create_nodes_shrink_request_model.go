// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateNodesShrinkRequest
	GetClusterId() *string
	SetComputeNodeShrink(v string) *CreateNodesShrinkRequest
	GetComputeNodeShrink() *string
	SetCount(v int32) *CreateNodesShrinkRequest
	GetCount() *int32
	SetDeploymentSetId(v string) *CreateNodesShrinkRequest
	GetDeploymentSetId() *string
	SetHPCInterConnect(v string) *CreateNodesShrinkRequest
	GetHPCInterConnect() *string
	SetHostnamePrefix(v string) *CreateNodesShrinkRequest
	GetHostnamePrefix() *string
	SetHostnameSuffix(v string) *CreateNodesShrinkRequest
	GetHostnameSuffix() *string
	SetHostnamesShrink(v string) *CreateNodesShrinkRequest
	GetHostnamesShrink() *string
	SetKeepAlive(v string) *CreateNodesShrinkRequest
	GetKeepAlive() *string
	SetMinCount(v int32) *CreateNodesShrinkRequest
	GetMinCount() *int32
	SetQueueName(v string) *CreateNodesShrinkRequest
	GetQueueName() *string
	SetRamRole(v string) *CreateNodesShrinkRequest
	GetRamRole() *string
	SetReservedNodePoolId(v string) *CreateNodesShrinkRequest
	GetReservedNodePoolId() *string
	SetVSwitchId(v string) *CreateNodesShrinkRequest
	GetVSwitchId() *string
}

type CreateNodesShrinkRequest struct {
	// The ID of the cluster.
	//
	// You can call [ListClusters](https://help.aliyun.com/document_detail/87116.html) to obtain the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies the hardware configuration of the compute node.
	ComputeNodeShrink *string `json:"ComputeNode,omitempty" xml:"ComputeNode,omitempty"`
	// The number of compute nodes to add. Valid values: 1 to 99. The value of MinCount must be less than the value of Count.
	//
	// - If the ECS inventory is less than MinCount, the operation fails.
	//
	// - If the ECS inventory is between MinCount and Count, the number of nodes specified by MinCount is added.
	//
	// - If the ECS inventory is greater than Count, the number of nodes specified by Count is added.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the deployment set. You can call the [DescribeDeploymentSets](https://help.aliyun.com/document_detail/91313.html) operation to obtain the ID. Only deployment sets that use the low-latency network policy are supported.
	//
	// example:
	//
	// ds-bp1frxuzdg87zh4pzq****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// Specifies the network type for communication between compute nodes. Valid values:
	//
	// - vpc
	//
	// - eRDMA
	//
	// example:
	//
	// vpc
	HPCInterConnect *string `json:"HPCInterConnect,omitempty" xml:"HPCInterConnect,omitempty"`
	// The hostname prefix for the compute nodes in the queue.
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
	// The ID of the reserved node pool.
	HostnamesShrink *string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty"`
	// Specifies whether deletion protection is enabled for the compute node.
	//
	// example:
	//
	// false
	KeepAlive *string `json:"KeepAlive,omitempty" xml:"KeepAlive,omitempty"`
	// example:
	//
	// 10
	MinCount *int32 `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// The name of the queue to which the compute nodes belong.
	//
	// example:
	//
	// test1
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The name of the authorized instance role to be attached to the compute nodes in the queue.
	//
	// example:
	//
	// AliyunServiceRoleForOOSBandwidthScheduler
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// The ID of the reserved node pool.
	//
	// example:
	//
	// rnp-756vlp7a
	ReservedNodePoolId *string `json:"ReservedNodePoolId,omitempty" xml:"ReservedNodePoolId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1lfcjbfb099rrjn****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNodesShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateNodesShrinkRequest) GetComputeNodeShrink() *string {
	return s.ComputeNodeShrink
}

func (s *CreateNodesShrinkRequest) GetCount() *int32 {
	return s.Count
}

func (s *CreateNodesShrinkRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateNodesShrinkRequest) GetHPCInterConnect() *string {
	return s.HPCInterConnect
}

func (s *CreateNodesShrinkRequest) GetHostnamePrefix() *string {
	return s.HostnamePrefix
}

func (s *CreateNodesShrinkRequest) GetHostnameSuffix() *string {
	return s.HostnameSuffix
}

func (s *CreateNodesShrinkRequest) GetHostnamesShrink() *string {
	return s.HostnamesShrink
}

func (s *CreateNodesShrinkRequest) GetKeepAlive() *string {
	return s.KeepAlive
}

func (s *CreateNodesShrinkRequest) GetMinCount() *int32 {
	return s.MinCount
}

func (s *CreateNodesShrinkRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateNodesShrinkRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *CreateNodesShrinkRequest) GetReservedNodePoolId() *string {
	return s.ReservedNodePoolId
}

func (s *CreateNodesShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateNodesShrinkRequest) SetClusterId(v string) *CreateNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetComputeNodeShrink(v string) *CreateNodesShrinkRequest {
	s.ComputeNodeShrink = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetCount(v int32) *CreateNodesShrinkRequest {
	s.Count = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetDeploymentSetId(v string) *CreateNodesShrinkRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetHPCInterConnect(v string) *CreateNodesShrinkRequest {
	s.HPCInterConnect = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetHostnamePrefix(v string) *CreateNodesShrinkRequest {
	s.HostnamePrefix = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetHostnameSuffix(v string) *CreateNodesShrinkRequest {
	s.HostnameSuffix = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetHostnamesShrink(v string) *CreateNodesShrinkRequest {
	s.HostnamesShrink = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetKeepAlive(v string) *CreateNodesShrinkRequest {
	s.KeepAlive = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetMinCount(v int32) *CreateNodesShrinkRequest {
	s.MinCount = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetQueueName(v string) *CreateNodesShrinkRequest {
	s.QueueName = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetRamRole(v string) *CreateNodesShrinkRequest {
	s.RamRole = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetReservedNodePoolId(v string) *CreateNodesShrinkRequest {
	s.ReservedNodePoolId = &v
	return s
}

func (s *CreateNodesShrinkRequest) SetVSwitchId(v string) *CreateNodesShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
