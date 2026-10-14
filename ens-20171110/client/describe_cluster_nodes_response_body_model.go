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
	SetRequestId(v string) *DescribeClusterNodesResponseBody
	GetRequestId() *string
}

type DescribeClusterNodesResponseBody struct {
	// The node information.
	Nodes []*DescribeClusterNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The pagination information.
	Page *DescribeClusterNodesResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DescribeClusterNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterNodesResponseBody) SetNodes(v []*DescribeClusterNodesResponseBodyNodes) *DescribeClusterNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribeClusterNodesResponseBody) SetPage(v *DescribeClusterNodesResponseBodyPage) *DescribeClusterNodesResponseBody {
	s.Page = v
	return s
}

func (s *DescribeClusterNodesResponseBody) SetRequestId(v string) *DescribeClusterNodesResponseBody {
	s.RequestId = &v
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
	// The cluster ID.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ENS edge zone ID.
	//
	// example:
	//
	// cn-fuzhou-23
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The instance image ID.
	//
	// example:
	//
	// m-680cki2ruj1q2bm0mz1k9tnbz
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-51mctytg7tv4yw4amu3oczxsx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the ENS instance.
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
	// The node information.
	NodeInfo *DescribeClusterNodesResponseBodyNodesNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// The node pool ID.
	//
	// example:
	//
	// np861febb748f84f5f9f1c76819eff7f56
	NodepoolId *string `json:"NodepoolId,omitempty" xml:"NodepoolId,omitempty"`
	// The billing method. Valid values:
	//
	// - PostPaid: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The status of the node pool. Valid values:
	//
	// - initial: The node pool is being created.
	//
	// - running: The node pool is created.
	//
	// - removing: The node pool is being removed.
	//
	// example:
	//
	// initial
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeClusterNodesResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesResponseBodyNodes) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterNodesResponseBodyNodes) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeClusterNodesResponseBodyNodes) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeClusterNodesResponseBodyNodes) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeClusterNodesResponseBodyNodes) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *DescribeClusterNodesResponseBodyNodes) GetNodeInfo() *DescribeClusterNodesResponseBodyNodesNodeInfo {
	return s.NodeInfo
}

func (s *DescribeClusterNodesResponseBodyNodes) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *DescribeClusterNodesResponseBodyNodes) GetPayType() *string {
	return s.PayType
}

func (s *DescribeClusterNodesResponseBodyNodes) GetState() *string {
	return s.State
}

func (s *DescribeClusterNodesResponseBodyNodes) SetClusterId(v string) *DescribeClusterNodesResponseBodyNodes {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetEnsRegionId(v string) *DescribeClusterNodesResponseBodyNodes {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetImageId(v string) *DescribeClusterNodesResponseBodyNodes {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetInstanceId(v string) *DescribeClusterNodesResponseBodyNodes {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetIpAddress(v []*string) *DescribeClusterNodesResponseBodyNodes {
	s.IpAddress = v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetNodeInfo(v *DescribeClusterNodesResponseBodyNodesNodeInfo) *DescribeClusterNodesResponseBodyNodes {
	s.NodeInfo = v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetNodepoolId(v string) *DescribeClusterNodesResponseBodyNodes {
	s.NodepoolId = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetPayType(v string) *DescribeClusterNodesResponseBodyNodes {
	s.PayType = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) SetState(v string) *DescribeClusterNodesResponseBodyNodes {
	s.State = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodes) Validate() error {
	if s.NodeInfo != nil {
		if err := s.NodeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterNodesResponseBodyNodesNodeInfo struct {
	// The container runtime version.
	//
	// example:
	//
	// containerd://1.6.28
	ContainerRuntimeVersion *string `json:"ContainerRuntimeVersion,omitempty" xml:"ContainerRuntimeVersion,omitempty"`
	// The kubelet version of the node.
	//
	// example:
	//
	// v1.30.7-aliyun.1
	KubeletVersion *string `json:"KubeletVersion,omitempty" xml:"KubeletVersion,omitempty"`
	// The operating system image of the node.
	//
	// example:
	//
	// CentOS Linux 7 (Core)
	OsImage *string `json:"OsImage,omitempty" xml:"OsImage,omitempty"`
}

func (s DescribeClusterNodesResponseBodyNodesNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodesResponseBodyNodesNodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesResponseBodyNodesNodeInfo) GetContainerRuntimeVersion() *string {
	return s.ContainerRuntimeVersion
}

func (s *DescribeClusterNodesResponseBodyNodesNodeInfo) GetKubeletVersion() *string {
	return s.KubeletVersion
}

func (s *DescribeClusterNodesResponseBodyNodesNodeInfo) GetOsImage() *string {
	return s.OsImage
}

func (s *DescribeClusterNodesResponseBodyNodesNodeInfo) SetContainerRuntimeVersion(v string) *DescribeClusterNodesResponseBodyNodesNodeInfo {
	s.ContainerRuntimeVersion = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodesNodeInfo) SetKubeletVersion(v string) *DescribeClusterNodesResponseBodyNodesNodeInfo {
	s.KubeletVersion = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodesNodeInfo) SetOsImage(v string) *DescribeClusterNodesResponseBodyNodesNodeInfo {
	s.OsImage = &v
	return s
}

func (s *DescribeClusterNodesResponseBodyNodesNodeInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNodesResponseBodyPage struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
