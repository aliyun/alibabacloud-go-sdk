// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeGPUMetric interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorType(v string) *NodeGPUMetric
	GetAcceleratorType() *string
	SetGPUCount(v int32) *NodeGPUMetric
	GetGPUCount() *int32
	SetGPUMetrics(v []*GPUMetric) *NodeGPUMetric
	GetGPUMetrics() []*GPUMetric
	SetGPUType(v string) *NodeGPUMetric
	GetGPUType() *string
	SetMemoryUtil(v float32) *NodeGPUMetric
	GetMemoryUtil() *float32
	SetNodeId(v string) *NodeGPUMetric
	GetNodeId() *string
	SetNodeType(v string) *NodeGPUMetric
	GetNodeType() *string
	SetTotalMemory(v float32) *NodeGPUMetric
	GetTotalMemory() *float32
	SetUsedMemory(v float32) *NodeGPUMetric
	GetUsedMemory() *float32
}

type NodeGPUMetric struct {
	// The accelerator type.
	//
	// example:
	//
	// GPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// The GPU count.
	//
	// example:
	//
	// 1
	GPUCount *int32 `json:"GPUCount,omitempty" xml:"GPUCount,omitempty"`
	// The GPU metrics.
	GPUMetrics []*GPUMetric `json:"GPUMetrics,omitempty" xml:"GPUMetrics,omitempty" type:"Repeated"`
	// The GPU type.
	//
	// example:
	//
	// V100
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// The memory utilization.
	//
	// example:
	//
	// 0
	MemoryUtil *float32 `json:"MemoryUtil,omitempty" xml:"MemoryUtil,omitempty"`
	// The node ID.
	//
	// example:
	//
	// lingj19q90jp66nq-mg1a2kc92f9nn8yd-20
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node type.
	//
	// example:
	//
	// ecs.c8i.xlarge
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The total memory.
	TotalMemory *float32 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	// The used memory.
	//
	// example:
	//
	// 0
	UsedMemory *float32 `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
}

func (s NodeGPUMetric) String() string {
	return dara.Prettify(s)
}

func (s NodeGPUMetric) GoString() string {
	return s.String()
}

func (s *NodeGPUMetric) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *NodeGPUMetric) GetGPUCount() *int32 {
	return s.GPUCount
}

func (s *NodeGPUMetric) GetGPUMetrics() []*GPUMetric {
	return s.GPUMetrics
}

func (s *NodeGPUMetric) GetGPUType() *string {
	return s.GPUType
}

func (s *NodeGPUMetric) GetMemoryUtil() *float32 {
	return s.MemoryUtil
}

func (s *NodeGPUMetric) GetNodeId() *string {
	return s.NodeId
}

func (s *NodeGPUMetric) GetNodeType() *string {
	return s.NodeType
}

func (s *NodeGPUMetric) GetTotalMemory() *float32 {
	return s.TotalMemory
}

func (s *NodeGPUMetric) GetUsedMemory() *float32 {
	return s.UsedMemory
}

func (s *NodeGPUMetric) SetAcceleratorType(v string) *NodeGPUMetric {
	s.AcceleratorType = &v
	return s
}

func (s *NodeGPUMetric) SetGPUCount(v int32) *NodeGPUMetric {
	s.GPUCount = &v
	return s
}

func (s *NodeGPUMetric) SetGPUMetrics(v []*GPUMetric) *NodeGPUMetric {
	s.GPUMetrics = v
	return s
}

func (s *NodeGPUMetric) SetGPUType(v string) *NodeGPUMetric {
	s.GPUType = &v
	return s
}

func (s *NodeGPUMetric) SetMemoryUtil(v float32) *NodeGPUMetric {
	s.MemoryUtil = &v
	return s
}

func (s *NodeGPUMetric) SetNodeId(v string) *NodeGPUMetric {
	s.NodeId = &v
	return s
}

func (s *NodeGPUMetric) SetNodeType(v string) *NodeGPUMetric {
	s.NodeType = &v
	return s
}

func (s *NodeGPUMetric) SetTotalMemory(v float32) *NodeGPUMetric {
	s.TotalMemory = &v
	return s
}

func (s *NodeGPUMetric) SetUsedMemory(v float32) *NodeGPUMetric {
	s.UsedMemory = &v
	return s
}

func (s *NodeGPUMetric) Validate() error {
	if s.GPUMetrics != nil {
		for _, item := range s.GPUMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
