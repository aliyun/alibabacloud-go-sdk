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
	// example:
	//
	// 10
	BatchInterval *int64 `json:"batch_interval,omitempty" xml:"batch_interval,omitempty"`
	// example:
	//
	// 0
	MaxFailedNodes *int64 `json:"max_failed_nodes,omitempty" xml:"max_failed_nodes,omitempty"`
	// example:
	//
	// 1
	MaxParallelism *int64 `json:"max_parallelism,omitempty" xml:"max_parallelism,omitempty"`
	// example:
	//
	// ["cn-hangzhou.10.91.xx.xx"]
	NodeNamesShrink *string `json:"node_names,omitempty" xml:"node_names,omitempty"`
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
