// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoleZoneInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNode(v *DescribeRoleZoneInfoResponseBodyNode) *DescribeRoleZoneInfoResponseBody
	GetNode() *DescribeRoleZoneInfoResponseBodyNode
	SetPageNumber(v int32) *DescribeRoleZoneInfoResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRoleZoneInfoResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRoleZoneInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeRoleZoneInfoResponseBody
	GetTotalCount() *int32
}

type DescribeRoleZoneInfoResponseBody struct {
	// Details about each node in the instance.
	Node *DescribeRoleZoneInfoResponseBodyNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 224B97FB-A275-4EAC-86E9-8922FEA2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRoleZoneInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBody) GetNode() *DescribeRoleZoneInfoResponseBodyNode {
	return s.Node
}

func (s *DescribeRoleZoneInfoResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRoleZoneInfoResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRoleZoneInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRoleZoneInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRoleZoneInfoResponseBody) SetNode(v *DescribeRoleZoneInfoResponseBodyNode) *DescribeRoleZoneInfoResponseBody {
	s.Node = v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetPageNumber(v int32) *DescribeRoleZoneInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetPageSize(v int32) *DescribeRoleZoneInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetRequestId(v string) *DescribeRoleZoneInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetTotalCount(v int32) *DescribeRoleZoneInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRoleZoneInfoResponseBodyNode struct {
	NodeInfo []*DescribeRoleZoneInfoResponseBodyNodeNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s DescribeRoleZoneInfoResponseBodyNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBodyNode) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBodyNode) GetNodeInfo() []*DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	return s.NodeInfo
}

func (s *DescribeRoleZoneInfoResponseBodyNode) SetNodeInfo(v []*DescribeRoleZoneInfoResponseBodyNodeNodeInfo) *DescribeRoleZoneInfoResponseBodyNode {
	s.NodeInfo = v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNode) Validate() error {
	return dara.Validate(s)
}

type DescribeRoleZoneInfoResponseBodyNodeNodeInfo struct {
	// The current bandwidth of the node, which consists of the default bandwidth and the increased bandwidth. Unit: MB/s.
	//
	// > 	- You can call the [EnableAdditionalBandwidth](https://help.aliyun.com/document_detail/473771.html) operation to specify the increased bandwidth.
	//
	// > 	- You can also use this parameter to calculate the increased bandwidth. For example, if the default bandwidth of the node is 96 MB/s and the returned value of this parameter is 100, the increased bandwidth is 4 MB/s.
	//
	// example:
	//
	// 100
	CurrentBandWidth *int64 `json:"CurrentBandWidth,omitempty" xml:"CurrentBandWidth,omitempty"`
	// The minor version of the node.
	//
	// example:
	//
	// redis-5.0_0.3.10
	CurrentMinorVersion *string `json:"CurrentMinorVersion,omitempty" xml:"CurrentMinorVersion,omitempty"`
	// The ID of the data shard.
	//
	// example:
	//
	// 30381****
	CustinsId *string `json:"CustinsId,omitempty" xml:"CustinsId,omitempty"`
	// The default bandwidth of the node. Unit: MB/s.
	//
	// example:
	//
	// 96
	DefaultBandWidth *int64 `json:"DefaultBandWidth,omitempty" xml:"DefaultBandWidth,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// r-t4nlenc2p04uvb****
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// Indicates whether the node is a read replica. If the node is a read replica, **3*	- is returned.
	//
	// >  If the node is not a read replica, no value is returned.
	//
	// example:
	//
	// 3
	InsType *int32 `json:"InsType,omitempty" xml:"InsType,omitempty"`
	// Indicates whether the minor version is the latest version. Valid values:
	//
	// 	- **0**: The minor version is not the latest version.
	//
	// 	- **1**: The minor version is the latest version.
	//
	// >  To update the minor version, call the [ModifyInstanceMinorVersion](https://help.aliyun.com/document_detail/473777.html) operation.
	//
	// example:
	//
	// 1
	IsLatestVersion *int32 `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// Indicates whether the bandwidth of the node is increased. Valid values:
	//
	// 	- **true**: The bandwidth of the node is not increased.
	//
	// 	- **false**: The bandwidth of the node is increased.
	//
	// example:
	//
	// true
	IsOpenBandWidthService *bool `json:"IsOpenBandWidthService,omitempty" xml:"IsOpenBandWidthService,omitempty"`
	// This parameter is used only for internal maintenance of instances.
	//
	// example:
	//
	// 10065****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node type. Valid values:
	//
	// 	- **db**: data node.
	//
	// 	- **proxy**: proxy node.
	//
	// 	- **normal**: regular node. This value is returned when the instance runs in the standard architecture.
	//
	// example:
	//
	// normal
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The role of the node. Valid values:
	//
	// 	- **master**: master node
	//
	// 	- **slave**: replica node
	//
	// example:
	//
	// master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRoleZoneInfoResponseBodyNodeNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetCurrentBandWidth() *int64 {
	return s.CurrentBandWidth
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetCurrentMinorVersion() *string {
	return s.CurrentMinorVersion
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetCustinsId() *string {
	return s.CustinsId
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetDefaultBandWidth() *int64 {
	return s.DefaultBandWidth
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetInsName() *string {
	return s.InsName
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetInsType() *int32 {
	return s.InsType
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetIsLatestVersion() *int32 {
	return s.IsLatestVersion
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetIsOpenBandWidthService() *bool {
	return s.IsOpenBandWidthService
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetRole() *string {
	return s.Role
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetCurrentBandWidth(v int64) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.CurrentBandWidth = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetCurrentMinorVersion(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.CurrentMinorVersion = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetCustinsId(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.CustinsId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetDefaultBandWidth(v int64) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.DefaultBandWidth = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetInsName(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.InsName = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetInsType(v int32) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.InsType = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetIsLatestVersion(v int32) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetIsOpenBandWidthService(v bool) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.IsOpenBandWidthService = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetNodeId(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.NodeId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetNodeType(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.NodeType = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetRole(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.Role = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetZoneId(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.ZoneId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) Validate() error {
	return dara.Validate(s)
}
