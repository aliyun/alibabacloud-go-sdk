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
	SetDrainTimeoutSeconds(v int64) *RemoveNodePoolNodesRequest
	GetDrainTimeoutSeconds() *int64
	SetInstanceIds(v []*string) *RemoveNodePoolNodesRequest
	GetInstanceIds() []*string
	SetNodes(v []*string) *RemoveNodePoolNodesRequest
	GetNodes() []*string
	SetReleaseNode(v bool) *RemoveNodePoolNodesRequest
	GetReleaseNode() *bool
}

type RemoveNodePoolNodesRequest struct {
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
	InstanceIds []*string `json:"instance_ids,omitempty" xml:"instance_ids,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]**
	//
	// The list of nodes to be removed.
	//
	// 	Danger: This parameter is deprecated. Use `instance_ids` instead.</danger>
	Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
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

func (s *RemoveNodePoolNodesRequest) GetDrainTimeoutSeconds() *int64 {
	return s.DrainTimeoutSeconds
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

func (s *RemoveNodePoolNodesRequest) SetDrainTimeoutSeconds(v int64) *RemoveNodePoolNodesRequest {
	s.DrainTimeoutSeconds = &v
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
