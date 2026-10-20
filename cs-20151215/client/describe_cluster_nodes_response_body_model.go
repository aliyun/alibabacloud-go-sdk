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
	// The list of node details.
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
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterNodesResponseBodyNodes struct {
	// The time when the node was created. The time follows the ISO 8601 standard and includes the time zone.
	//
	// example:
	//
	// 2025-04-16T07:28:00Z
	CreationTime *string `json:"creation_time,omitempty" xml:"creation_time,omitempty"`
	// The error message returned when the node was created.
	//
	// example:
	//
	// error***
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// The time when the node expires. The time follows the ISO 8601 standard and includes the time zone.
	//
	// example:
	//
	// 2099-12-31T15:59:00Z
	ExpiredTime *string `json:"expired_time,omitempty" xml:"expired_time,omitempty"`
	// The hostname of the node.
	//
	// example:
	//
	// iZ2vcckdmxp7u0urj2k****
	HostName *string `json:"host_name,omitempty" xml:"host_name,omitempty"`
	// The ID of the system image used by the node.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20241218.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The billing method of the node. Valid values:
	//
	// - `PrePaid`: subscription.
	//
	// - `PostPaid`: pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"instance_charge_type,omitempty" xml:"instance_charge_type,omitempty"`
	// The instance ID of the node.
	//
	// example:
	//
	// i-2vcckdmxp7u0urj2****
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// The name of the node in the cluster.
	//
	// example:
	//
	// worker-k8s-for-cs-c5cdf7e3938bc4f8eb0e44b21a80f****
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name,omitempty"`
	// The role of the node. Valid values:
	//
	// - Master: the master node of the cluster.
	//
	// - Worker: the worker node of the cluster.
	//
	// example:
	//
	// Worker
	InstanceRole *string `json:"instance_role,omitempty" xml:"instance_role,omitempty"`
	// The ECS instance status of the node. Valid values:
	//
	// - `Pending`: the instance is being created.
	//
	// - `Running`: the instance is running.
	//
	// - `Starting`: the instance is starting.
	//
	// - `Stopping`: the instance is stopping.
	//
	// - `Stopped`: the instance is stopped.
	//
	// example:
	//
	// Running
	InstanceStatus *string `json:"instance_status,omitempty" xml:"instance_status,omitempty"`
	// The node specifications.
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"instance_type,omitempty" xml:"instance_type,omitempty"`
	// The name of the ECS instance family to which the node belongs.
	//
	// example:
	//
	// ecs.c5
	InstanceTypeFamily *string `json:"instance_type_family,omitempty" xml:"instance_type_family,omitempty"`
	// The IP addresses of the node.
	IpAddress []*string `json:"ip_address,omitempty" xml:"ip_address,omitempty" type:"Repeated"`
	// Specifies whether the node is an Alibaba Cloud instance. Valid values:
	//
	// - `true`: the node is an Alibaba Cloud instance.
	//
	// - `false`: the node is not an Alibaba Cloud instance.
	//
	// example:
	//
	// true
	IsAliyunNode *bool `json:"is_aliyun_node,omitempty" xml:"is_aliyun_node,omitempty"`
	// The name of the node, which is the node identifier in the cluster.
	//
	// example:
	//
	// cn-chengdu.192.168xx.xx
	NodeName *string `json:"node_name,omitempty" xml:"node_name,omitempty"`
	// Specifies whether the node is ready. Valid values:
	//
	// - `Ready`: the node is ready.
	//
	// - `NotReady`: the node is not ready.
	//
	// - `Unknown`: the status is unknown.
	//
	// - `Offline`: the node is offline.
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
	// The method used to initialize the node, such as manual creation or creation by Resource Orchestration Service (ROS).
	//
	// example:
	//
	// ess
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The type of the spot instance. Valid values:
	//
	// - NoSpot: the instance is not a spot instance.
	//
	// - SpotWithPriceLimit: the instance is a spot instance with a maximum price limit.
	//
	// - SpotAsPriceGo: the instance is a spot instance whose price is automatically set based on the current market price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"spot_strategy,omitempty" xml:"spot_strategy,omitempty"`
	// The running state of the node. Valid values:
	//
	// - `pending`: the node is being created.
	//
	// - `running`: the node is running.
	//
	// - `starting`: the node is starting.
	//
	// - `stopping`: the node is stopping.
	//
	// - `stopped`: the node is stopped.
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
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The maximum number of records to return per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of results.
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
