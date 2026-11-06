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
	SetDrainTimeoutSeconds(v int64) *RemoveNodePoolNodesShrinkRequest
	GetDrainTimeoutSeconds() *int64
	SetInstanceIdsShrink(v string) *RemoveNodePoolNodesShrinkRequest
	GetInstanceIdsShrink() *string
	SetNodesShrink(v string) *RemoveNodePoolNodesShrinkRequest
	GetNodesShrink() *string
	SetReleaseNode(v bool) *RemoveNodePoolNodesShrinkRequest
	GetReleaseNode() *bool
}

type RemoveNodePoolNodesShrinkRequest struct {
	// Specifies whether to remove nodes concurrently. Valid values:
	//
	// - true: Nodes are removed from the scaling group concurrently.
	//
	// - false: Nodes are removed from the scaling group sequentially.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Concurrency *bool `json:"concurrency,omitempty" xml:"concurrency,omitempty"`
	// Specifies whether to drain the nodes. Valid values:
	//
	// - true: Drain the nodes.
	//
	// - false: Do not drain the nodes.
	//
	// example:
	//
	// true
	DrainNode *bool `json:"drain_node,omitempty" xml:"drain_node,omitempty"`
	// The drain timeout period in seconds. Valid values: (0, 1800]. Default value: 1800. This parameter takes effect only when drain_node is set to true.
	//
	// example:
	//
	// 600
	DrainTimeoutSeconds *int64 `json:"drain_timeout_seconds,omitempty" xml:"drain_timeout_seconds,omitempty"`
	// The list of instances to be removed.
	InstanceIdsShrink *string `json:"instance_ids,omitempty" xml:"instance_ids,omitempty"`
	// Deprecated
	//
	// **[Deprecated]**
	//
	// The list of nodes to be removed.
	//
	// 	Danger: This parameter is deprecated. Use `instance_ids` instead.</danger>
	NodesShrink *string `json:"nodes,omitempty" xml:"nodes,omitempty"`
	// Specifies whether automatic release the nodes. Valid values:
	//
	// - true: Release the nodes.
	//
	// - false: Do not release the nodes.
	//
	// Removing nodes does not automatically release subscription nodes. Automatic release subscription nodes, [change the billing method of the nodes from subscription to pay-as-you-go](https://www.alibabacloud.com/help/en/ecs/change-the-billing-method-of-an-instance-from-subscription-to-pay-as-you-go-1), and then log on to the ECS console automatic release them.
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

func (s *RemoveNodePoolNodesShrinkRequest) GetDrainTimeoutSeconds() *int64 {
	return s.DrainTimeoutSeconds
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

func (s *RemoveNodePoolNodesShrinkRequest) SetDrainTimeoutSeconds(v int64) *RemoveNodePoolNodesShrinkRequest {
	s.DrainTimeoutSeconds = &v
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
