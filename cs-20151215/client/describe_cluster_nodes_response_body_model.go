// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodes(v []*DescribeClusterNodesResponseBodyNodes) *DescribeClusterNodesResponseBody
	GetNodes() []*DescribeClusterNodesResponseBodyNodes
	SetPage(v *DescribeClusterNodesResponseBodyPage) *DescribeClusterNodesResponseBody
	GetPage() *DescribeClusterNodesResponseBodyPage
}

type DescribeClusterNodesResponseBody struct {
	// The details of the nodes in the cluster.
	Nodes []*DescribeClusterNodesResponseBodyNodes `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// The pagination information.
	Page *DescribeClusterNodesResponseBodyPage `json:"page,omitempty" xml:"page,omitempty" type:"Struct"`
}

func (s DescribeClusterNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesResponseBody) GetNodes() []*DescribeClusterNodesResponseBodyNodes {
	return s.Nodes
}

func (s *DescribeClusterNodesResponseBody) GetPage() *DescribeClusterNodesResponseBodyPage {
	return s.Page
}

func (s *DescribeClusterNodesResponseBody) SetNodes(v []*DescribeClusterNodesResponseBodyNodes) *DescribeClusterNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribeClusterNodesResponseBody) SetPage(v *DescribeClusterNodesResponseBodyPage) *DescribeClusterNodesResponseBody {
	s.Page = v
	return s
}

func (s *DescribeClusterNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodesResponseBodyNodes struct {
	// The time when the node was created.
	//
	// example:
	//
	// 2020-08-25T11:25:35+08:00
	CreationTime *string `json:"creation_time,omitempty" xml:"creation_time,omitempty"`
	// The error message generated when the node was created.
	//
	// example:
	//
	// error***
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// The expiration date of the node.
	//
	// example:
	//
	// 2099-12-31T23:59:00+08:00
	ExpiredTime *string `json:"expired_time,omitempty" xml:"expired_time,omitempty"`
	// The name of the host.
	//
	// example:
	//
	// iZ2vcckdmxp7u0urj2k****
	HostName *string `json:"host_name,omitempty" xml:"host_name,omitempty"`
	// The ID of the system image that is used by the node.
	//
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200529.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The billing method of the node. Valid values:
	//
	// 	- `PrePaid`: the subscription billing method. If the value is PrePaid, make sure that you have a sufficient balance or credit in your account. Otherwise, an `InvalidPayMethod` error is returned.
	//
	// 	- `PostPaid`: the pay-as-you-go billing method.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"instance_charge_type,omitempty" xml:"instance_charge_type,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-2vcckdmxp7u0urj2****
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// The name of the instance on which the node is deployed.
	//
	// example:
	//
	// worker-k8s-for-cs-c5cdf7e3938bc4f8eb0e44b21a80f****
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name,omitempty"`
	// The role of the node. Valid values:
	//
	// 	- Master: master node
	//
	// 	- Worker: worker node
	//
	// example:
	//
	// Worker
	InstanceRole *string `json:"instance_role,omitempty" xml:"instance_role,omitempty"`
	// The status of the node.
	//
	// example:
	//
	// Running
	InstanceStatus *string `json:"instance_status,omitempty" xml:"instance_status,omitempty"`
	// The type of the node.
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// The ECS instance family of the node.
	//
	// example:
	//
	// ecs.c5
	InstanceTypeFamily *string `json:"instance_type_family,omitempty" xml:"instance_type_family,omitempty"`
	// The IP address of the node.
	IpAddress []*string `json:"ip_address,omitempty" xml:"ip_address,omitempty" type:"Repeated"`
	// Indicates whether the instance on which the node is deployed is provided by Alibaba Cloud. Valid values:
	//
	// 	- `true`: The instance is provided by Alibaba Cloud.
	//
	// 	- `false`: The instance is not provided by Alibaba Cloud.
	//
	// example:
	//
	// true
	IsAliyunNode *bool `json:"is_aliyun_node,omitempty" xml:"is_aliyun_node,omitempty"`
	// The name of the node. This name is the identifier of the node in the cluster.
	//
	// example:
	//
	// cn-chengdu.192.168.0.36
	NodeName *string `json:"node_name,omitempty" xml:"node_name,omitempty"`
	// Indicates whether the node is ready. Valid values:
	//
	// 	- `Ready`: The node is ready.
	//
	// 	- `NotReady`: The node is not ready.
	//
	// 	- `Unknown`: The status of the node is unknown.
	//
	// 	- `Offline`: The node is offline.
	//
	// example:
	//
	// Ready
	NodeStatus *string `json:"node_status,omitempty" xml:"node_status,omitempty"`
	// The node pool ID.
	//
	// example:
	//
	// np0794239424a84eb7a95327369d56****
	NodepoolId *string `json:"nodepool_id,omitempty" xml:"nodepool_id,omitempty"`
	// Indicates how the node is initialized. A node can be manually created or created by using Resource Orchestration Service (ROS).
	//
	// example:
	//
	// ess_attach
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The type of preemptible instance. Valid values:
	//
	// 	- NoSpot: a non-preemptible instance.
	//
	// 	- SpotWithPriceLimit: a preemptible instance that is configured with the highest bid price.
	//
	// 	- SpotAsPriceGo: a preemptible instance for which the system automatically bids based on the current market price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"spot_strategy,omitempty" xml:"spot_strategy,omitempty"`
	// The status of the node. Valid values:
	//
	// 	- `pending`: The node is being created.
	//
	// 	- `running`: The node is running.
	//
	// 	- `starting`: The node is being started.
	//
	// 	- `stopping`: The node is being stopped.
	//
	// 	- `stopped`: The node is stopped.
	//
	// example:
	//
	// running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s DescribeClusterNodesResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesResponseBodyNodes) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeClusterNodesResponseBodyNodes) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeClusterNodesResponseBodyNodes) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeClusterNodesResponseBodyNodes) GetHostName() *string {
	return s.HostName
}

func (s *DescribeClusterNodesResponseBodyNodes) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeClusterNodesResponseBodyNodes) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeClusterNodesResponseBodyNodes) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeClusterNodesResponseBodyNodes) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeClusterNodesResponseBodyNodes) GetInstanceRole() *string {
	return s.InstanceRole
}

func (s *DescribeClusterNodesResponseBodyNodes) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeClusterNodesResponseBodyNodes) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeClusterNodesResponseBodyNodes) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeClusterNodesResponseBodyNodes) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *DescribeClusterNodesResponseBodyNodes) GetIsAliyunNode() *bool {
	return s.IsAliyunNode
}

func (s *DescribeClusterNodesResponseBodyNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeClusterNodesResponseBodyNodes) GetNodeStatus() *string {
	return s.NodeStatus
}

func (s *DescribeClusterNodesResponseBodyNodes) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *DescribeClusterNodesResponseBodyNodes) GetSource() *string {
	return s.Source
}

func (s *DescribeClusterNodesResponseBodyNodes) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeClusterNodesResponseBodyNodes) GetState() *string {
	return s.State
}

func (s *DescribeClusterNodesResponseBodyNodes) SetCreationTime(v string) *DescribeClusterNodesResponseBodyNodes {
	s.CreationTime = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetErrorMessage(v string) *DescribeClusterNodesResponseBodyNodes {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetExpiredTime(v string) *DescribeClusterNodesResponseBodyNodes {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetHostName(v string) *DescribeClusterNodesResponseBodyNodes {
	s.HostName = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetImageId(v string) *DescribeClusterNodesResponseBodyNodes {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceChargeType(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceId(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceName(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceName = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceRole(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceRole = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceStatus(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceType(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceType = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceTypeFamily(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetIpAddress(v []*string) *DescribeClusterNodesResponseBodyNodes {
	s.IpAddress = v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetIsAliyunNode(v bool) *DescribeClusterNodesResponseBodyNodes {
	s.IsAliyunNode = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetNodeName(v string) *DescribeClusterNodesResponseBodyNodes {
	s.NodeName = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetNodeStatus(v string) *DescribeClusterNodesResponseBodyNodes {
	s.NodeStatus = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetNodepoolId(v string) *DescribeClusterNodesResponseBodyNodes {
	s.NodepoolId = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetSource(v string) *DescribeClusterNodesResponseBodyNodes {
	s.Source = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetSpotStrategy(v string) *DescribeClusterNodesResponseBodyNodes {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetState(v string) *DescribeClusterNodesResponseBodyNodes {
	s.State = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodesResponseBodyPage struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s DescribeClusterNodesResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodesResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClusterNodesResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClusterNodesResponseBodyPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeClusterNodesResponseBodyPage) SetPageNumber(v int32) *DescribeClusterNodesResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyPage) SetPageSize(v int32) *DescribeClusterNodesResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyPage) SetTotalCount(v int32) *DescribeClusterNodesResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyPage) Validate() error {
	return dara.Validate(s)
}
