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
	SetDrainTimeoutSeconds(v int64) *DeleteClusterNodesRequest
	GetDrainTimeoutSeconds() *int64
	SetNodes(v []*string) *DeleteClusterNodesRequest
	GetNodes() []*string
	SetReleaseNode(v bool) *DeleteClusterNodesRequest
	GetReleaseNode() *bool
}

type DeleteClusterNodesRequest struct {
	// Specifies whether to automatically drain pods from the nodes. Valid values:
	//
	// - `true`: Automatically drains pods from the nodes.
	//
	// - `false`: Does not automatically drain pods from the nodes.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	DrainNode *bool `json:"drain_node,omitempty" xml:"drain_node,omitempty"`
	// The drain timeout period in seconds. Valid values: (0, 1800]. If this parameter is not specified, the default value 1800 is used. This parameter takes effect only when drain_node is set to true.
	//
	// example:
	//
	// 600
	DrainTimeoutSeconds *int64 `json:"drain_timeout_seconds,omitempty" xml:"drain_timeout_seconds,omitempty"`
	// The list of nodes to remove. The node name must be the name of the node in the cluster, such as `cn-hangzhou.192.168.xx.xx`.
	//
	// This parameter is required.
	Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// Specifies whether to release the associated ECS instances. Valid values:
	//
	// - `true`: Releases the associated ECS instances.
	//
	// - `false`: Retains the associated ECS instances.
	//
	// Default value: `false`.
	//
	// > When a node is a subscription instance, you cannot release the associated ECS instance.
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

func (s *DeleteClusterNodesRequest) GetDrainTimeoutSeconds() *int64 {
	return s.DrainTimeoutSeconds
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

func (s *DeleteClusterNodesRequest) SetDrainTimeoutSeconds(v int64) *DeleteClusterNodesRequest {
	s.DrainTimeoutSeconds = &v
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
