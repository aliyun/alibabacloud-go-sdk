// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrainNode(v bool) *DeleteClusterNodesRequest
	GetDrainNode() *bool
	SetNodes(v []*string) *DeleteClusterNodesRequest
	GetNodes() []*string
	SetReleaseNode(v bool) *DeleteClusterNodesRequest
	GetReleaseNode() *bool
}

type DeleteClusterNodesRequest struct {
	// Whether to automatically drain Pods on the node. Valid values:
	//
	// - `true`: Automatically drain Pods on the node.
	//
	// - `false`: Do not automatically drain Pods on the node.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	DrainNode *bool `json:"drain_node,omitempty" xml:"drain_node,omitempty"`
	// List of nodes to remove. The node names must be the names of the nodes in the cluster, for example: `cn-hangzhou.192.168.xx.xx`.
	//
	// This parameter is required.
	Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// Whether to simultaneously remove ECS instances. Valid values:
	//
	// - `true`: Simultaneously remove ECS instances.
	//
	// - `false`: Retain ECS instances.
	//
	// Default value: `false`.
	//
	// > Simultaneous removal of ECS instances is not supported when the nodes are subscription instances.
	//
	// example:
	//
	// true
	ReleaseNode *bool `json:"release_node,omitempty" xml:"release_node,omitempty"`
}

func (s DeleteClusterNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodesRequest) GetDrainNode() *bool {
	return s.DrainNode
}

func (s *DeleteClusterNodesRequest) GetNodes() []*string {
	return s.Nodes
}

func (s *DeleteClusterNodesRequest) GetReleaseNode() *bool {
	return s.ReleaseNode
}

func (s *DeleteClusterNodesRequest) SetDrainNode(v bool) *DeleteClusterNodesRequest {
	s.DrainNode = &v
	return s
}

func (s *DeleteClusterNodesRequest) SetNodes(v []*string) *DeleteClusterNodesRequest {
	s.Nodes = v
	return s
}

func (s *DeleteClusterNodesRequest) SetReleaseNode(v bool) *DeleteClusterNodesRequest {
	s.ReleaseNode = &v
	return s
}

func (s *DeleteClusterNodesRequest) Validate() error {
	return dara.Validate(s)
}
