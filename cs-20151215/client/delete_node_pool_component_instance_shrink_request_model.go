// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodePoolComponentInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchInterval(v int64) *DeleteNodePoolComponentInstanceShrinkRequest
	GetBatchInterval() *int64
	SetMaxFailedNodes(v int64) *DeleteNodePoolComponentInstanceShrinkRequest
	GetMaxFailedNodes() *int64
	SetMaxParallelism(v int64) *DeleteNodePoolComponentInstanceShrinkRequest
	GetMaxParallelism() *int64
	SetNodeNamesShrink(v string) *DeleteNodePoolComponentInstanceShrinkRequest
	GetNodeNamesShrink() *string
	SetPausePolicy(v string) *DeleteNodePoolComponentInstanceShrinkRequest
	GetPausePolicy() *string
}

type DeleteNodePoolComponentInstanceShrinkRequest struct {
	// The interval between batches, in seconds. Default value: 0.
	//
	// example:
	//
	// 10
	BatchInterval *int64 `json:"batch_interval,omitempty" xml:"batch_interval,omitempty"`
	// The maximum number of nodes that are allowed to fail during the rolling process. Default value: 0. A value of 0 indicates that the task fails if any node fails. A value greater than 0 indicates that the task fails and stops if the cumulative number of failed nodes exceeds this value.
	//
	// example:
	//
	// 0
	MaxFailedNodes *int64 `json:"max_failed_nodes,omitempty" xml:"max_failed_nodes,omitempty"`
	// The maximum number of parallel nodes per batch. Default value: 1.
	//
	// example:
	//
	// 1
	MaxParallelism *int64 `json:"max_parallelism,omitempty" xml:"max_parallelism,omitempty"`
	// The list of nodes to rotate. Default value: all nodes.
	//
	// example:
	//
	// ["cn-hangzhou.10.91.xx.xx"]
	NodeNamesShrink *string `json:"node_names,omitempty" xml:"node_names,omitempty"`
	// The auto-pause policy during the node upgrade process. Valid values:
	//
	// - NotPause
	//
	// - FirstBatch
	//
	// - EveryBatch
	//
	// Default value: NotPause.
	//
	// example:
	//
	// NotPause
	PausePolicy *string `json:"pause_policy,omitempty" xml:"pause_policy,omitempty"`
}

func (s DeleteNodePoolComponentInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodePoolComponentInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodePoolComponentInstanceShrinkRequest) GetBatchInterval() *int64 {
	return s.BatchInterval
}

func (s *DeleteNodePoolComponentInstanceShrinkRequest) GetMaxFailedNodes() *int64 {
	return s.MaxFailedNodes
}

func (s *DeleteNodePoolComponentInstanceShrinkRequest) GetMaxParallelism() *int64 {
	return s.MaxParallelism
}

func (s *DeleteNodePoolComponentInstanceShrinkRequest) GetNodeNamesShrink() *string {
	return s.NodeNamesShrink
}

func (s *DeleteNodePoolComponentInstanceShrinkRequest) GetPausePolicy() *string {
	return s.PausePolicy
}

func (s *DeleteNodePoolComponentInstanceShrinkRequest) SetBatchInterval(v int64) *DeleteNodePoolComponentInstanceShrinkRequest {
	s.BatchInterval = &v
	return s
}

func (s *DeleteNodePoolComponentInstanceShrinkRequest) SetMaxFailedNodes(v int64) *DeleteNodePoolComponentInstanceShrinkRequest {
	s.MaxFailedNodes = &v
	return s
}

func (s *DeleteNodePoolComponentInstanceShrinkRequest) SetMaxParallelism(v int64) *DeleteNodePoolComponentInstanceShrinkRequest {
	s.MaxParallelism = &v
	return s
}

func (s *DeleteNodePoolComponentInstanceShrinkRequest) SetNodeNamesShrink(v string) *DeleteNodePoolComponentInstanceShrinkRequest {
	s.NodeNamesShrink = &v
	return s
}

func (s *DeleteNodePoolComponentInstanceShrinkRequest) SetPausePolicy(v string) *DeleteNodePoolComponentInstanceShrinkRequest {
	s.PausePolicy = &v
	return s
}

func (s *DeleteNodePoolComponentInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
