// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateNodesRequest
	GetClusterId() *string
	SetComputeNode(v *NodeTemplate) *CreateNodesRequest
	GetComputeNode() *NodeTemplate
	SetCount(v int32) *CreateNodesRequest
	GetCount() *int32
	SetDeploymentSetId(v string) *CreateNodesRequest
	GetDeploymentSetId() *string
	SetHPCInterConnect(v string) *CreateNodesRequest
	GetHPCInterConnect() *string
	SetHostnamePrefix(v string) *CreateNodesRequest
	GetHostnamePrefix() *string
	SetHostnameSuffix(v string) *CreateNodesRequest
	GetHostnameSuffix() *string
	SetHostnames(v []*string) *CreateNodesRequest
	GetHostnames() []*string
	SetKeepAlive(v string) *CreateNodesRequest
	GetKeepAlive() *string
	SetMinCount(v int32) *CreateNodesRequest
	GetMinCount() *int32
	SetQueueName(v string) *CreateNodesRequest
	GetQueueName() *string
	SetRamRole(v string) *CreateNodesRequest
	GetRamRole() *string
	SetReservedNodePoolId(v string) *CreateNodesRequest
	GetReservedNodePoolId() *string
	SetVSwitchId(v string) *CreateNodesRequest
	GetVSwitchId() *string
}

type CreateNodesRequest struct {
	// The ID of the cluster.
	//
	// You can call [ListClusters](https://help.aliyun.com/document_detail/87116.html) to obtain the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies the hardware configuration of the compute node.
	ComputeNode *NodeTemplate `json:"ComputeNode,omitempty" xml:"ComputeNode,omitempty"`
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
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
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

func (s CreateNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNodesRequest) GoString() string {
	return s.String()
}

func (s *CreateNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateNodesRequest) GetComputeNode() *NodeTemplate {
	return s.ComputeNode
}

func (s *CreateNodesRequest) GetCount() *int32 {
	return s.Count
}

func (s *CreateNodesRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateNodesRequest) GetHPCInterConnect() *string {
	return s.HPCInterConnect
}

func (s *CreateNodesRequest) GetHostnamePrefix() *string {
	return s.HostnamePrefix
}

func (s *CreateNodesRequest) GetHostnameSuffix() *string {
	return s.HostnameSuffix
}

func (s *CreateNodesRequest) GetHostnames() []*string {
	return s.Hostnames
}

func (s *CreateNodesRequest) GetKeepAlive() *string {
	return s.KeepAlive
}

func (s *CreateNodesRequest) GetMinCount() *int32 {
	return s.MinCount
}

func (s *CreateNodesRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateNodesRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *CreateNodesRequest) GetReservedNodePoolId() *string {
	return s.ReservedNodePoolId
}

func (s *CreateNodesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateNodesRequest) SetClusterId(v string) *CreateNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNodesRequest) SetComputeNode(v *NodeTemplate) *CreateNodesRequest {
	s.ComputeNode = v
	return s
}

func (s *CreateNodesRequest) SetCount(v int32) *CreateNodesRequest {
	s.Count = &v
	return s
}

func (s *CreateNodesRequest) SetDeploymentSetId(v string) *CreateNodesRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateNodesRequest) SetHPCInterConnect(v string) *CreateNodesRequest {
	s.HPCInterConnect = &v
	return s
}

func (s *CreateNodesRequest) SetHostnamePrefix(v string) *CreateNodesRequest {
	s.HostnamePrefix = &v
	return s
}

func (s *CreateNodesRequest) SetHostnameSuffix(v string) *CreateNodesRequest {
	s.HostnameSuffix = &v
	return s
}

func (s *CreateNodesRequest) SetHostnames(v []*string) *CreateNodesRequest {
	s.Hostnames = v
	return s
}

func (s *CreateNodesRequest) SetKeepAlive(v string) *CreateNodesRequest {
	s.KeepAlive = &v
	return s
}

func (s *CreateNodesRequest) SetMinCount(v int32) *CreateNodesRequest {
	s.MinCount = &v
	return s
}

func (s *CreateNodesRequest) SetQueueName(v string) *CreateNodesRequest {
	s.QueueName = &v
	return s
}

func (s *CreateNodesRequest) SetRamRole(v string) *CreateNodesRequest {
	s.RamRole = &v
	return s
}

func (s *CreateNodesRequest) SetReservedNodePoolId(v string) *CreateNodesRequest {
	s.ReservedNodePoolId = &v
	return s
}

func (s *CreateNodesRequest) SetVSwitchId(v string) *CreateNodesRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateNodesRequest) Validate() error {
	if s.ComputeNode != nil {
		if err := s.ComputeNode.Validate(); err != nil {
			return err
		}
	}
	return nil
}
