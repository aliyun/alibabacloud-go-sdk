// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterNetwork(v *DescribeClusterNetworkResponseBodyClusterNetwork) *DescribeClusterNetworkResponseBody
	GetClusterNetwork() *DescribeClusterNetworkResponseBodyClusterNetwork
	SetRequestId(v string) *DescribeClusterNetworkResponseBody
	GetRequestId() *string
}

type DescribeClusterNetworkResponseBody struct {
	// Information about the network topology edge in the cluster.
	ClusterNetwork *DescribeClusterNetworkResponseBodyClusterNetwork `json:"ClusterNetwork,omitempty" xml:"ClusterNetwork,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C286491D-4A2F-589A-B63B-D2AD3DA9BD71
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetworkResponseBody) GetClusterNetwork() *DescribeClusterNetworkResponseBodyClusterNetwork {
	return s.ClusterNetwork
}

func (s *DescribeClusterNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterNetworkResponseBody) SetClusterNetwork(v *DescribeClusterNetworkResponseBodyClusterNetwork) *DescribeClusterNetworkResponseBody {
	s.ClusterNetwork = v
	return s
}

func (s *DescribeClusterNetworkResponseBody) SetRequestId(v string) *DescribeClusterNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNetworkResponseBodyClusterNetwork struct {
	// An array that consists of information about the topology edge.
	Edge []*DescribeClusterNetworkResponseBodyClusterNetworkEdge `json:"Edge,omitempty" xml:"Edge,omitempty" type:"Repeated"`
	// An array that consists of information about the node.
	Node []*DescribeClusterNetworkResponseBodyClusterNetworkNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Repeated"`
}

func (s DescribeClusterNetworkResponseBodyClusterNetwork) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNetworkResponseBodyClusterNetwork) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetworkResponseBodyClusterNetwork) GetEdge() []*DescribeClusterNetworkResponseBodyClusterNetworkEdge {
	return s.Edge
}

func (s *DescribeClusterNetworkResponseBodyClusterNetwork) GetNode() []*DescribeClusterNetworkResponseBodyClusterNetworkNode {
	return s.Node
}

func (s *DescribeClusterNetworkResponseBodyClusterNetwork) SetEdge(v []*DescribeClusterNetworkResponseBodyClusterNetworkEdge) *DescribeClusterNetworkResponseBodyClusterNetwork {
	s.Edge = v
	return s
}

func (s *DescribeClusterNetworkResponseBodyClusterNetwork) SetNode(v []*DescribeClusterNetworkResponseBodyClusterNetworkNode) *DescribeClusterNetworkResponseBodyClusterNetwork {
	s.Node = v
	return s
}

func (s *DescribeClusterNetworkResponseBodyClusterNetwork) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNetworkResponseBodyClusterNetworkEdge struct {
	// The ID of the destination node.
	//
	// example:
	//
	// cfb41a869c71e4678a97021582dd8****
	DstNodeId *string `json:"DstNodeId,omitempty" xml:"DstNodeId,omitempty"`
	// The type of the destination node. Valid values:
	//
	// 	- Set the value to **cluster**.
	//
	// example:
	//
	// cluster
	DstNodeType *string `json:"DstNodeType,omitempty" xml:"DstNodeType,omitempty"`
	// The ID of the topology edge.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The port number of the topology edge.
	//
	// example:
	//
	// 6164
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the source node.
	//
	// example:
	//
	// cfb41a869c71e4678a97021582dd8****
	SrcNodeId *string `json:"SrcNodeId,omitempty" xml:"SrcNodeId,omitempty"`
	// The type of the source node. Valid values:
	//
	// 	- **cluster**: a cluster.
	//
	// 	- **internet**: a network node outside the cluster
	//
	// example:
	//
	// cluster
	SrcNodeType *string `json:"SrcNodeType,omitempty" xml:"SrcNodeType,omitempty"`
}

func (s DescribeClusterNetworkResponseBodyClusterNetworkEdge) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNetworkResponseBodyClusterNetworkEdge) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkEdge) GetDstNodeId() *string {
	return s.DstNodeId
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkEdge) GetDstNodeType() *string {
	return s.DstNodeType
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkEdge) GetId() *string {
	return s.Id
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkEdge) GetPort() *string {
	return s.Port
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkEdge) GetSrcNodeId() *string {
	return s.SrcNodeId
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkEdge) GetSrcNodeType() *string {
	return s.SrcNodeType
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkEdge) SetDstNodeId(v string) *DescribeClusterNetworkResponseBodyClusterNetworkEdge {
	s.DstNodeId = &v
	return s
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkEdge) SetDstNodeType(v string) *DescribeClusterNetworkResponseBodyClusterNetworkEdge {
	s.DstNodeType = &v
	return s
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkEdge) SetId(v string) *DescribeClusterNetworkResponseBodyClusterNetworkEdge {
	s.Id = &v
	return s
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkEdge) SetPort(v string) *DescribeClusterNetworkResponseBodyClusterNetworkEdge {
	s.Port = &v
	return s
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkEdge) SetSrcNodeId(v string) *DescribeClusterNetworkResponseBodyClusterNetworkEdge {
	s.SrcNodeId = &v
	return s
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkEdge) SetSrcNodeType(v string) *DescribeClusterNetworkResponseBodyClusterNetworkEdge {
	s.SrcNodeType = &v
	return s
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkEdge) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterNetworkResponseBodyClusterNetworkNode struct {
	// The status of the microsegmentation switch. Valid values:
	//
	// 	- **0**: off.
	//
	// 	- **1**: on.
	//
	// example:
	//
	// 1
	CnnfSwitch *int32 `json:"CnnfSwitch,omitempty" xml:"CnnfSwitch,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// cfeb7a9f99ce740e98c5595d0fe37****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The network type. Valid values:
	//
	// 	- **0**: classic network.
	//
	// 	- **1**: virtual private cloud (VPC).
	//
	// example:
	//
	// 1
	InterceptionType *int32 `json:"InterceptionType,omitempty" xml:"InterceptionType,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// testwww
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the network topology switch. Valid values:
	//
	// 	- **ON**
	//
	// 	- **OFF**
	//
	// example:
	//
	// ON
	NetTopoSwitch *string `json:"NetTopoSwitch,omitempty" xml:"NetTopoSwitch,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **3**: high.
	//
	// 	- **2**: medium.
	//
	// 	- **1**: low.
	//
	// 	- **0**: secure.
	//
	// 	- **-1**: unknown.
	//
	// example:
	//
	// 3
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The type of the node. Valid values:
	//
	// 	- **cluster**: a cluster.
	//
	// 	- **internet**: a network node outside the cluster.
	//
	// example:
	//
	// cluster
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeClusterNetworkResponseBodyClusterNetworkNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNetworkResponseBodyClusterNetworkNode) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkNode) GetCnnfSwitch() *int32 {
	return s.CnnfSwitch
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkNode) GetId() *string {
	return s.Id
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkNode) GetInterceptionType() *int32 {
	return s.InterceptionType
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkNode) GetName() *string {
	return s.Name
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkNode) GetNetTopoSwitch() *string {
	return s.NetTopoSwitch
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkNode) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkNode) GetType() *string {
	return s.Type
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkNode) SetCnnfSwitch(v int32) *DescribeClusterNetworkResponseBodyClusterNetworkNode {
	s.CnnfSwitch = &v
	return s
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkNode) SetId(v string) *DescribeClusterNetworkResponseBodyClusterNetworkNode {
	s.Id = &v
	return s
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkNode) SetInterceptionType(v int32) *DescribeClusterNetworkResponseBodyClusterNetworkNode {
	s.InterceptionType = &v
	return s
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkNode) SetName(v string) *DescribeClusterNetworkResponseBodyClusterNetworkNode {
	s.Name = &v
	return s
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkNode) SetNetTopoSwitch(v string) *DescribeClusterNetworkResponseBodyClusterNetworkNode {
	s.NetTopoSwitch = &v
	return s
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkNode) SetRiskLevel(v string) *DescribeClusterNetworkResponseBodyClusterNetworkNode {
	s.RiskLevel = &v
	return s
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkNode) SetType(v string) *DescribeClusterNetworkResponseBodyClusterNetworkNode {
	s.Type = &v
	return s
}

func (s *DescribeClusterNetworkResponseBodyClusterNetworkNode) Validate() error {
	return dara.Validate(s)
}
