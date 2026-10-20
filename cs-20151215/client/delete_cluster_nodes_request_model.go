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
	// Specifies whether to automatically drain Pods on the nodes. Valid values:
	//
	// - `true`: Automatically drains Pods on the nodes.
	//
	// - `false`: Does not automatically drain Pods on the nodes.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	DrainNode *bool `json:"drain_node,omitempty" xml:"drain_node,omitempty"`
	// The list of nodes to remove. The node name must be the name of the node in the cluster, such as `cn-hangzhou.192.168.xx.xx`.
	//
	// This parameter is required.
	Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// Specifies whether to release the ECS instances at the same time. Valid values:
	//
	// - `true`: Releases the ECS instances at the same time.
	//
	// - `false`: Retains the ECS instances.
	//
	// Default value: `false`.
	//
	// > Releasing ECS instances at the same time is not supported when the nodes are subscription instances.
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
