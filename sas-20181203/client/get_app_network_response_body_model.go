// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppNetwork(v *GetAppNetworkResponseBodyAppNetwork) *GetAppNetworkResponseBody
	GetAppNetwork() *GetAppNetworkResponseBodyAppNetwork
	SetRequestId(v string) *GetAppNetworkResponseBody
	GetRequestId() *string
}

type GetAppNetworkResponseBody struct {
	// The information about the application network topology.
	AppNetwork *GetAppNetworkResponseBodyAppNetwork `json:"AppNetwork,omitempty" xml:"AppNetwork,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7D46EDB0-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppNetworkResponseBody) GetAppNetwork() *GetAppNetworkResponseBodyAppNetwork {
	return s.AppNetwork
}

func (s *GetAppNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppNetworkResponseBody) SetAppNetwork(v *GetAppNetworkResponseBodyAppNetwork) *GetAppNetworkResponseBody {
	s.AppNetwork = v
	return s
}

func (s *GetAppNetworkResponseBody) SetRequestId(v string) *GetAppNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppNetworkResponseBody) Validate() error {
	if s.AppNetwork != nil {
		if err := s.AppNetwork.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppNetworkResponseBodyAppNetwork struct {
	// The information about the topology edge.
	Edge []*GetAppNetworkResponseBodyAppNetworkEdge `json:"Edge,omitempty" xml:"Edge,omitempty" type:"Repeated"`
	// The namespace.
	Namespace []*GetAppNetworkResponseBodyAppNetworkNamespace `json:"Namespace,omitempty" xml:"Namespace,omitempty" type:"Repeated"`
	// The information about the application node.
	Node []*GetAppNetworkResponseBodyAppNetworkNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Repeated"`
}

func (s GetAppNetworkResponseBodyAppNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetAppNetworkResponseBodyAppNetwork) GoString() string {
	return s.String()
}

func (s *GetAppNetworkResponseBodyAppNetwork) GetEdge() []*GetAppNetworkResponseBodyAppNetworkEdge {
	return s.Edge
}

func (s *GetAppNetworkResponseBodyAppNetwork) GetNamespace() []*GetAppNetworkResponseBodyAppNetworkNamespace {
	return s.Namespace
}

func (s *GetAppNetworkResponseBodyAppNetwork) GetNode() []*GetAppNetworkResponseBodyAppNetworkNode {
	return s.Node
}

func (s *GetAppNetworkResponseBodyAppNetwork) SetEdge(v []*GetAppNetworkResponseBodyAppNetworkEdge) *GetAppNetworkResponseBodyAppNetwork {
	s.Edge = v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetwork) SetNamespace(v []*GetAppNetworkResponseBodyAppNetworkNamespace) *GetAppNetworkResponseBodyAppNetwork {
	s.Namespace = v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetwork) SetNode(v []*GetAppNetworkResponseBodyAppNetworkNode) *GetAppNetworkResponseBodyAppNetwork {
	s.Node = v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetwork) Validate() error {
	if s.Edge != nil {
		for _, item := range s.Edge {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Namespace != nil {
		for _, item := range s.Namespace {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Node != nil {
		for _, item := range s.Node {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAppNetworkResponseBodyAppNetworkEdge struct {
	// The ID of the destination node.
	//
	// example:
	//
	// 102
	DstNodeId *string `json:"DstNodeId,omitempty" xml:"DstNodeId,omitempty"`
	// The type of the destination node. Valid values:
	//
	// 	- **app**: an application
	//
	// 	- **internet**: a network node in another cluster
	//
	// example:
	//
	// app
	DstNodeType *string `json:"DstNodeType,omitempty" xml:"DstNodeType,omitempty"`
	// The ID of the edge.
	//
	// example:
	//
	// 3534
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of the destination port.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the source node.
	//
	// example:
	//
	// 101
	SrcNodeId *string `json:"SrcNodeId,omitempty" xml:"SrcNodeId,omitempty"`
	// The type of the source node. Valid values:
	//
	// 	- **app**: an application
	//
	// 	- **internet**: a network node in another cluster
	//
	// example:
	//
	// app
	SrcNodeType *string `json:"SrcNodeType,omitempty" xml:"SrcNodeType,omitempty"`
}

func (s GetAppNetworkResponseBodyAppNetworkEdge) String() string {
	return dara.Prettify(s)
}

func (s GetAppNetworkResponseBodyAppNetworkEdge) GoString() string {
	return s.String()
}

func (s *GetAppNetworkResponseBodyAppNetworkEdge) GetDstNodeId() *string {
	return s.DstNodeId
}

func (s *GetAppNetworkResponseBodyAppNetworkEdge) GetDstNodeType() *string {
	return s.DstNodeType
}

func (s *GetAppNetworkResponseBodyAppNetworkEdge) GetId() *string {
	return s.Id
}

func (s *GetAppNetworkResponseBodyAppNetworkEdge) GetPort() *string {
	return s.Port
}

func (s *GetAppNetworkResponseBodyAppNetworkEdge) GetSrcNodeId() *string {
	return s.SrcNodeId
}

func (s *GetAppNetworkResponseBodyAppNetworkEdge) GetSrcNodeType() *string {
	return s.SrcNodeType
}

func (s *GetAppNetworkResponseBodyAppNetworkEdge) SetDstNodeId(v string) *GetAppNetworkResponseBodyAppNetworkEdge {
	s.DstNodeId = &v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetworkEdge) SetDstNodeType(v string) *GetAppNetworkResponseBodyAppNetworkEdge {
	s.DstNodeType = &v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetworkEdge) SetId(v string) *GetAppNetworkResponseBodyAppNetworkEdge {
	s.Id = &v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetworkEdge) SetPort(v string) *GetAppNetworkResponseBodyAppNetworkEdge {
	s.Port = &v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetworkEdge) SetSrcNodeId(v string) *GetAppNetworkResponseBodyAppNetworkEdge {
	s.SrcNodeId = &v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetworkEdge) SetSrcNodeType(v string) *GetAppNetworkResponseBodyAppNetworkEdge {
	s.SrcNodeType = &v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetworkEdge) Validate() error {
	return dara.Validate(s)
}

type GetAppNetworkResponseBodyAppNetworkNamespace struct {
	// The ID of the namespace.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the custom namespace.
	//
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAppNetworkResponseBodyAppNetworkNamespace) String() string {
	return dara.Prettify(s)
}

func (s GetAppNetworkResponseBodyAppNetworkNamespace) GoString() string {
	return s.String()
}

func (s *GetAppNetworkResponseBodyAppNetworkNamespace) GetId() *string {
	return s.Id
}

func (s *GetAppNetworkResponseBodyAppNetworkNamespace) GetName() *string {
	return s.Name
}

func (s *GetAppNetworkResponseBodyAppNetworkNamespace) SetId(v string) *GetAppNetworkResponseBodyAppNetworkNamespace {
	s.Id = &v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetworkNamespace) SetName(v string) *GetAppNetworkResponseBodyAppNetworkNamespace {
	s.Name = &v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetworkNamespace) Validate() error {
	return dara.Validate(s)
}

type GetAppNetworkResponseBodyAppNetworkNode struct {
	// The list of the container IDs.
	ContainerIds []*string `json:"ContainerIds,omitempty" xml:"ContainerIds,omitempty" type:"Repeated"`
	// The ID of the node.
	//
	// example:
	//
	// 1274
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// console
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 1
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **3**: high
	//
	// 	- **2**: medium
	//
	// 	- **1**: low
	//
	// 	- **0**: warning
	//
	// 	- **-1**: unknown
	//
	// example:
	//
	// 0
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The type of the node. Valid values:
	//
	// 	- **app**: an application
	//
	// 	- **internet**: a network node in another cluster
	//
	// example:
	//
	// app
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAppNetworkResponseBodyAppNetworkNode) String() string {
	return dara.Prettify(s)
}

func (s GetAppNetworkResponseBodyAppNetworkNode) GoString() string {
	return s.String()
}

func (s *GetAppNetworkResponseBodyAppNetworkNode) GetContainerIds() []*string {
	return s.ContainerIds
}

func (s *GetAppNetworkResponseBodyAppNetworkNode) GetId() *string {
	return s.Id
}

func (s *GetAppNetworkResponseBodyAppNetworkNode) GetName() *string {
	return s.Name
}

func (s *GetAppNetworkResponseBodyAppNetworkNode) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetAppNetworkResponseBodyAppNetworkNode) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetAppNetworkResponseBodyAppNetworkNode) GetType() *string {
	return s.Type
}

func (s *GetAppNetworkResponseBodyAppNetworkNode) SetContainerIds(v []*string) *GetAppNetworkResponseBodyAppNetworkNode {
	s.ContainerIds = v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetworkNode) SetId(v string) *GetAppNetworkResponseBodyAppNetworkNode {
	s.Id = &v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetworkNode) SetName(v string) *GetAppNetworkResponseBodyAppNetworkNode {
	s.Name = &v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetworkNode) SetNamespaceId(v string) *GetAppNetworkResponseBodyAppNetworkNode {
	s.NamespaceId = &v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetworkNode) SetRiskLevel(v string) *GetAppNetworkResponseBodyAppNetworkNode {
	s.RiskLevel = &v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetworkNode) SetType(v string) *GetAppNetworkResponseBodyAppNetworkNode {
	s.Type = &v
	return s
}

func (s *GetAppNetworkResponseBodyAppNetworkNode) Validate() error {
	return dara.Validate(s)
}
