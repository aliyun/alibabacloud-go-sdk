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
	SetKeepAlive(v string) *CreateNodesShrinkRequest
	GetKeepAlive() *string
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
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The hardware configurations of the compute nodes.
	ComputeNodeShrink *string `json:"ComputeNode,omitempty" xml:"ComputeNode,omitempty"`
	// The number of compute nodes that you want to add. Valid values: 1 to 99. The MinCount value must be smaller than the Count value.
	//
	// 	- If the number of available Elastic Compute Service (ECS) instances is smaller than the MinCount value, the nodes fail to be added.
	//
	// 	- If the number of available ECS instances is larger than the MinCount value but smaller than the Count value, nodes are added based on the MinCount value.
	//
	// 	- If the number of available ECS instances is larger than the Count value, nodes are added based on the Count value.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Deployment set ID. You can obtain the deployment set ID through [DescribeDeploymentSets](https://help.aliyun.com/document_detail/91313.html). Currently, only deployment sets with a low network latency strategy are supported.
	//
	// example:
	//
	// ds-bp1frxuzdg87zh4pzq****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The type of the network between compute nodes. Valid values:
	//
	// 	- vpc
	//
	// 	- eRDMA
	//
	// example:
	//
	// vpc
	HPCInterConnect *string `json:"HPCInterConnect,omitempty" xml:"HPCInterConnect,omitempty"`
	// The hostname prefix of the added compute nodes.
	//
	// example:
	//
	// compute
	HostnamePrefix *string `json:"HostnamePrefix,omitempty" xml:"HostnamePrefix,omitempty"`
	// The hostname suffix of the added compute nodes.
	//
	// example:
	//
	// demo
	HostnameSuffix *string `json:"HostnameSuffix,omitempty" xml:"HostnameSuffix,omitempty"`
	// Specifies whether to enable deletion protection for the added compute nodes.
	//
	// example:
	//
	// false
	KeepAlive *string `json:"KeepAlive,omitempty" xml:"KeepAlive,omitempty"`
	// The name of the queue for which you want to create compute nodes.
	//
	// example:
	//
	// test1
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The Resource Access Management (RAM) role to be assumed by the added nodes.
	//
	// example:
	//
	// AliyunServiceRoleForOOSBandwidthScheduler
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// Preset node pool ID.
	//
	// example:
	//
	// rnp-756vlp7a
	ReservedNodePoolId *string `json:"ReservedNodePoolId,omitempty" xml:"ReservedNodePoolId,omitempty"`
	// The ID of the vSwitch to be used by the added nodes.
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

func (s *CreateNodesShrinkRequest) GetKeepAlive() *string {
	return s.KeepAlive
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

func (s *CreateNodesShrinkRequest) SetKeepAlive(v string) *CreateNodesShrinkRequest {
	s.KeepAlive = &v
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
