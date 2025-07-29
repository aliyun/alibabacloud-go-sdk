// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveClusterNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrainNode(v bool) *RemoveClusterNodesRequest
	GetDrainNode() *bool
	SetNodes(v []*string) *RemoveClusterNodesRequest
	GetNodes() []*string
	SetReleaseNode(v bool) *RemoveClusterNodesRequest
	GetReleaseNode() *bool
}

type RemoveClusterNodesRequest struct {
	// Specifies whether to evict all pods from the nodes that you want to remove.
	DrainNode *bool `json:"drain_node,omitempty" xml:"drain_node,omitempty"`
	// The list of nodes to be removed.
	//
	// This parameter is required.
	Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// Specifies whether to release the Elastic Compute Service (ECS) instances when they are removed from the cluster.
	ReleaseNode *bool `json:"release_node,omitempty" xml:"release_node,omitempty"`
}

func (s RemoveClusterNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *RemoveClusterNodesRequest) GetDrainNode() *bool {
	return s.DrainNode
}

func (s *RemoveClusterNodesRequest) GetNodes() []*string {
	return s.Nodes
}

func (s *RemoveClusterNodesRequest) GetReleaseNode() *bool {
	return s.ReleaseNode
}

func (s *RemoveClusterNodesRequest) SetDrainNode(v bool) *RemoveClusterNodesRequest {
	s.DrainNode = &v
	return s
}

func (s *RemoveClusterNodesRequest) SetNodes(v []*string) *RemoveClusterNodesRequest {
	s.Nodes = v
	return s
}

func (s *RemoveClusterNodesRequest) SetReleaseNode(v bool) *RemoveClusterNodesRequest {
	s.ReleaseNode = &v
	return s
}

func (s *RemoveClusterNodesRequest) Validate() error {
	return dara.Validate(s)
}
