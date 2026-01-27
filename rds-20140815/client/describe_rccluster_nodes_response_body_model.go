// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCClusterNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodes(v []*DescribeRCClusterNodesResponseBodyNodes) *DescribeRCClusterNodesResponseBody
	GetNodes() []*DescribeRCClusterNodesResponseBodyNodes
	SetPage(v *DescribeRCClusterNodesResponseBodyPage) *DescribeRCClusterNodesResponseBody
	GetPage() *DescribeRCClusterNodesResponseBodyPage
	SetRequestId(v string) *DescribeRCClusterNodesResponseBody
	GetRequestId() *string
}

type DescribeRCClusterNodesResponseBody struct {
	// The details of the nodes.
	Nodes []*DescribeRCClusterNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The pagination information.
	Page *DescribeRCClusterNodesResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 16C62438-491B-5C02-9B49-BA924A1372A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRCClusterNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCClusterNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCClusterNodesResponseBody) GetNodes() []*DescribeRCClusterNodesResponseBodyNodes {
	return s.Nodes
}

func (s *DescribeRCClusterNodesResponseBody) GetPage() *DescribeRCClusterNodesResponseBodyPage {
	return s.Page
}

func (s *DescribeRCClusterNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCClusterNodesResponseBody) SetNodes(v []*DescribeRCClusterNodesResponseBodyNodes) *DescribeRCClusterNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribeRCClusterNodesResponseBody) SetPage(v *DescribeRCClusterNodesResponseBodyPage) *DescribeRCClusterNodesResponseBody {
	s.Page = v
	return s
}

func (s *DescribeRCClusterNodesResponseBody) SetRequestId(v string) *DescribeRCClusterNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBody) Validate() error {
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

type DescribeRCClusterNodesResponseBodyNodes struct {
	// The time when the node was created.
	//
	// example:
	//
	// 2024-10-21T07:20:09Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The container version.
	//
	// example:
	//
	// 1.0
	DockerVersion *string `json:"DockerVersion,omitempty" xml:"DockerVersion,omitempty"`
	// The image ID of the node.
	//
	// example:
	//
	// m-2oqiu973jwcxe****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The node ID.
	//
	// example:
	//
	// rc-u79597n5f54s5bnz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node role. Valid values:
	//
	// 	- **Master**: master node
	//
	// 	- **Worker**: worker node
	//
	// example:
	//
	// Master
	InstanceRole *string `json:"InstanceRole,omitempty" xml:"InstanceRole,omitempty"`
	// The IP address.
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// Indicates whether the node is provided by Alibaba Cloud. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsAliyunNode *bool `json:"IsAliyunNode,omitempty" xml:"IsAliyunNode,omitempty"`
	// The node name, which is the identifier of the RDS Custom node in the cluster.
	//
	// example:
	//
	// cn-hangzhou.192.168.XXX.XXX
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The node pool ID.
	//
	// example:
	//
	// None
	NodePoolId *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	// Indicates whether the node is ready. Valid values:
	//
	// 	- **Ready**: The node is ready.
	//
	// 	- **NotReady**: The node is not ready.
	//
	// 	- **Unknown**: The status of the node is unknown.
	//
	// 	- **Offline**: The node is offline.
	//
	// example:
	//
	// Ready
	NodeStatus *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	// example:
	//
	// 1
	PodCount *int64 `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
	// The runtime of the ACK cluster.
	//
	// example:
	//
	// 2024-10-21T07:20:09Z
	RuntimeVersion *string `json:"RuntimeVersion,omitempty" xml:"RuntimeVersion,omitempty"`
	// The node status. Valid values:
	//
	// 	- **pending**
	//
	// 	- **running**
	//
	// 	- **starting**
	//
	// 	- **stopping**
	//
	// 	- **stopped**
	//
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeRCClusterNodesResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCClusterNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribeRCClusterNodesResponseBodyNodes) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRCClusterNodesResponseBodyNodes) GetDockerVersion() *string {
	return s.DockerVersion
}

func (s *DescribeRCClusterNodesResponseBodyNodes) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeRCClusterNodesResponseBodyNodes) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCClusterNodesResponseBodyNodes) GetInstanceRole() *string {
	return s.InstanceRole
}

func (s *DescribeRCClusterNodesResponseBodyNodes) GetIpAddresses() []*string {
	return s.IpAddresses
}

func (s *DescribeRCClusterNodesResponseBodyNodes) GetIsAliyunNode() *bool {
	return s.IsAliyunNode
}

func (s *DescribeRCClusterNodesResponseBodyNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeRCClusterNodesResponseBodyNodes) GetNodePoolId() *string {
	return s.NodePoolId
}

func (s *DescribeRCClusterNodesResponseBodyNodes) GetNodeStatus() *string {
	return s.NodeStatus
}

func (s *DescribeRCClusterNodesResponseBodyNodes) GetPodCount() *int64 {
	return s.PodCount
}

func (s *DescribeRCClusterNodesResponseBodyNodes) GetRuntimeVersion() *string {
	return s.RuntimeVersion
}

func (s *DescribeRCClusterNodesResponseBodyNodes) GetState() *string {
	return s.State
}

func (s *DescribeRCClusterNodesResponseBodyNodes) SetCreationTime(v string) *DescribeRCClusterNodesResponseBodyNodes {
	s.CreationTime = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyNodes) SetDockerVersion(v string) *DescribeRCClusterNodesResponseBodyNodes {
	s.DockerVersion = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyNodes) SetImageId(v string) *DescribeRCClusterNodesResponseBodyNodes {
	s.ImageId = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyNodes) SetInstanceId(v string) *DescribeRCClusterNodesResponseBodyNodes {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyNodes) SetInstanceRole(v string) *DescribeRCClusterNodesResponseBodyNodes {
	s.InstanceRole = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyNodes) SetIpAddresses(v []*string) *DescribeRCClusterNodesResponseBodyNodes {
	s.IpAddresses = v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyNodes) SetIsAliyunNode(v bool) *DescribeRCClusterNodesResponseBodyNodes {
	s.IsAliyunNode = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyNodes) SetNodeName(v string) *DescribeRCClusterNodesResponseBodyNodes {
	s.NodeName = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyNodes) SetNodePoolId(v string) *DescribeRCClusterNodesResponseBodyNodes {
	s.NodePoolId = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyNodes) SetNodeStatus(v string) *DescribeRCClusterNodesResponseBodyNodes {
	s.NodeStatus = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyNodes) SetPodCount(v int64) *DescribeRCClusterNodesResponseBodyNodes {
	s.PodCount = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyNodes) SetRuntimeVersion(v string) *DescribeRCClusterNodesResponseBodyNodes {
	s.RuntimeVersion = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyNodes) SetState(v string) *DescribeRCClusterNodesResponseBodyNodes {
	s.State = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeRCClusterNodesResponseBodyPage struct {
	// The page number.
	//
	// example:
	//
	// 2
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRCClusterNodesResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCClusterNodesResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribeRCClusterNodesResponseBodyPage) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeRCClusterNodesResponseBodyPage) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeRCClusterNodesResponseBodyPage) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeRCClusterNodesResponseBodyPage) SetPageNumber(v int64) *DescribeRCClusterNodesResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyPage) SetPageSize(v int64) *DescribeRCClusterNodesResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyPage) SetTotalCount(v int64) *DescribeRCClusterNodesResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *DescribeRCClusterNodesResponseBodyPage) Validate() error {
	return dara.Validate(s)
}
