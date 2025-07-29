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
	// Specifies whether to remove all pods from the nodes you want to remove. Valid values:
	//
	// 	- `true`: removes all pods automatically.
	//
	// 	- `false`: skips removing pods.
	//
	// Default value: `false`
	//
	// example:
	//
	// true
	DrainNode *bool `json:"drain_node,omitempty" xml:"drain_node,omitempty"`
	// The list of nodes to remove. You must specify the node names used in the cluster, for example, `cn-hangzhou.192.168.xx.xx`.
	//
	// This parameter is required.
	Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// Specifies whether to release the ECS instances. Valid values:
	//
	// 	- `true`: releases the ECS instances.
	//
	// 	- `false`: retains the ECS instances.
	//
	// Default value: `false`
	//
	// **
	//
	// **Notes*	- Unsupported for subscription ECS instances.
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
