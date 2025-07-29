// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveNodePoolNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v bool) *RemoveNodePoolNodesShrinkRequest
	GetConcurrency() *bool
	SetDrainNode(v bool) *RemoveNodePoolNodesShrinkRequest
	GetDrainNode() *bool
	SetInstanceIdsShrink(v string) *RemoveNodePoolNodesShrinkRequest
	GetInstanceIdsShrink() *string
	SetNodesShrink(v string) *RemoveNodePoolNodesShrinkRequest
	GetNodesShrink() *string
	SetReleaseNode(v bool) *RemoveNodePoolNodesShrinkRequest
	GetReleaseNode() *bool
}

type RemoveNodePoolNodesShrinkRequest struct {
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
	InstanceIdsShrink *string `json:"instance_ids,omitempty" xml:"instance_ids,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
	//
	// A list of nodes that you want to remove.
	//
	// >  This parameter is deprecated. Use instance_ids instead.
	NodesShrink *string `json:"nodes,omitempty" xml:"nodes,omitempty"`
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

func (s RemoveNodePoolNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveNodePoolNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveNodePoolNodesShrinkRequest) GetConcurrency() *bool {
	return s.Concurrency
}

func (s *RemoveNodePoolNodesShrinkRequest) GetDrainNode() *bool {
	return s.DrainNode
}

func (s *RemoveNodePoolNodesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *RemoveNodePoolNodesShrinkRequest) GetNodesShrink() *string {
	return s.NodesShrink
}

func (s *RemoveNodePoolNodesShrinkRequest) GetReleaseNode() *bool {
	return s.ReleaseNode
}

func (s *RemoveNodePoolNodesShrinkRequest) SetConcurrency(v bool) *RemoveNodePoolNodesShrinkRequest {
	s.Concurrency = &v
	return s
}

func (s *RemoveNodePoolNodesShrinkRequest) SetDrainNode(v bool) *RemoveNodePoolNodesShrinkRequest {
	s.DrainNode = &v
	return s
}

func (s *RemoveNodePoolNodesShrinkRequest) SetInstanceIdsShrink(v string) *RemoveNodePoolNodesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *RemoveNodePoolNodesShrinkRequest) SetNodesShrink(v string) *RemoveNodePoolNodesShrinkRequest {
	s.NodesShrink = &v
	return s
}

func (s *RemoveNodePoolNodesShrinkRequest) SetReleaseNode(v bool) *RemoveNodePoolNodesShrinkRequest {
	s.ReleaseNode = &v
	return s
}

func (s *RemoveNodePoolNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
