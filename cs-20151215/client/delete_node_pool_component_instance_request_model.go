// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodePoolComponentInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchInterval(v int64) *DeleteNodePoolComponentInstanceRequest
	GetBatchInterval() *int64
	SetMaxFailedNodes(v int64) *DeleteNodePoolComponentInstanceRequest
	GetMaxFailedNodes() *int64
	SetMaxParallelism(v int64) *DeleteNodePoolComponentInstanceRequest
	GetMaxParallelism() *int64
	SetNodeNames(v []*string) *DeleteNodePoolComponentInstanceRequest
	GetNodeNames() []*string
	SetPausePolicy(v string) *DeleteNodePoolComponentInstanceRequest
	GetPausePolicy() *string
}

type DeleteNodePoolComponentInstanceRequest struct {
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
	NodeNames []*string `json:"node_names,omitempty" xml:"node_names,omitempty" type:"Repeated"`
	// example:
	//
	// NotPause
	PausePolicy *string `json:"pause_policy,omitempty" xml:"pause_policy,omitempty"`
}

func (s DeleteNodePoolComponentInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodePoolComponentInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodePoolComponentInstanceRequest) GetBatchInterval() *int64 {
	return s.BatchInterval
}

func (s *DeleteNodePoolComponentInstanceRequest) GetMaxFailedNodes() *int64 {
	return s.MaxFailedNodes
}

func (s *DeleteNodePoolComponentInstanceRequest) GetMaxParallelism() *int64 {
	return s.MaxParallelism
}

func (s *DeleteNodePoolComponentInstanceRequest) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *DeleteNodePoolComponentInstanceRequest) GetPausePolicy() *string {
	return s.PausePolicy
}

func (s *DeleteNodePoolComponentInstanceRequest) SetBatchInterval(v int64) *DeleteNodePoolComponentInstanceRequest {
	s.BatchInterval = &v
	return s
}

func (s *DeleteNodePoolComponentInstanceRequest) SetMaxFailedNodes(v int64) *DeleteNodePoolComponentInstanceRequest {
	s.MaxFailedNodes = &v
	return s
}

func (s *DeleteNodePoolComponentInstanceRequest) SetMaxParallelism(v int64) *DeleteNodePoolComponentInstanceRequest {
	s.MaxParallelism = &v
	return s
}

func (s *DeleteNodePoolComponentInstanceRequest) SetNodeNames(v []*string) *DeleteNodePoolComponentInstanceRequest {
	s.NodeNames = v
	return s
}

func (s *DeleteNodePoolComponentInstanceRequest) SetPausePolicy(v string) *DeleteNodePoolComponentInstanceRequest {
	s.PausePolicy = &v
	return s
}

func (s *DeleteNodePoolComponentInstanceRequest) Validate() error {
	return dara.Validate(s)
}
