// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveNodePoolNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v bool) *RemoveNodePoolNodesRequest
	GetConcurrency() *bool
	SetDrainNode(v bool) *RemoveNodePoolNodesRequest
	GetDrainNode() *bool
	SetInstanceIds(v []*string) *RemoveNodePoolNodesRequest
	GetInstanceIds() []*string
	SetNodes(v []*string) *RemoveNodePoolNodesRequest
	GetNodes() []*string
	SetReleaseNode(v bool) *RemoveNodePoolNodesRequest
	GetReleaseNode() *bool
}

type RemoveNodePoolNodesRequest struct {
	// Whether to remove concurrently.
	//
	// example:
	//
	// false
	Concurrency *bool `json:"concurrency,omitempty" xml:"concurrency,omitempty"`
	// Specifies whether to drain the nodes that you want to remove. Valid values:
	//
	// 	- true: drain the nodes that you want to remove.
	//
	// 	- false: do not drain the nodes that you want to remove.
	//
	// example:
	//
	// true
	DrainNode *bool `json:"drain_node,omitempty" xml:"drain_node,omitempty"`
	// A list of instances that you want to remove.
	InstanceIds []*string `json:"instance_ids,omitempty" xml:"instance_ids,omitempty" type:"Repeated"`
	// Deprecated
	//
	// This parameter is deprecated.
	//
	// A list of nodes that you want to remove.
	//
	// >  This parameter is deprecated. Use instance_ids instead.
	Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// Specifies whether to release the nodes after they are removed. Valid values:
	//
	// 	- true: release the nodes after they are removed.
	//
	// 	- false: do not release the nodes after they are removed.
	//
	// example:
	//
	// true
	ReleaseNode *bool `json:"release_node,omitempty" xml:"release_node,omitempty"`
}

func (s RemoveNodePoolNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveNodePoolNodesRequest) GoString() string {
	return s.String()
}

func (s *RemoveNodePoolNodesRequest) GetConcurrency() *bool {
	return s.Concurrency
}

func (s *RemoveNodePoolNodesRequest) GetDrainNode() *bool {
	return s.DrainNode
}

func (s *RemoveNodePoolNodesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RemoveNodePoolNodesRequest) GetNodes() []*string {
	return s.Nodes
}

func (s *RemoveNodePoolNodesRequest) GetReleaseNode() *bool {
	return s.ReleaseNode
}

func (s *RemoveNodePoolNodesRequest) SetConcurrency(v bool) *RemoveNodePoolNodesRequest {
	s.Concurrency = &v
	return s
}

func (s *RemoveNodePoolNodesRequest) SetDrainNode(v bool) *RemoveNodePoolNodesRequest {
	s.DrainNode = &v
	return s
}

func (s *RemoveNodePoolNodesRequest) SetInstanceIds(v []*string) *RemoveNodePoolNodesRequest {
	s.InstanceIds = v
	return s
}

func (s *RemoveNodePoolNodesRequest) SetNodes(v []*string) *RemoveNodePoolNodesRequest {
	s.Nodes = v
	return s
}

func (s *RemoveNodePoolNodesRequest) SetReleaseNode(v bool) *RemoveNodePoolNodesRequest {
	s.ReleaseNode = &v
	return s
}

func (s *RemoveNodePoolNodesRequest) Validate() error {
	return dara.Validate(s)
}
