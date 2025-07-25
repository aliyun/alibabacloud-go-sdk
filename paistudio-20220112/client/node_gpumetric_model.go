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
	AcceleratorType *string      `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	GPUCount        *int32       `json:"GPUCount,omitempty" xml:"GPUCount,omitempty"`
	GPUMetrics      []*GPUMetric `json:"GPUMetrics,omitempty" xml:"GPUMetrics,omitempty" type:"Repeated"`
	GPUType         *string      `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	MemoryUtil      *float32     `json:"MemoryUtil,omitempty" xml:"MemoryUtil,omitempty"`
	NodeId          *string      `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeType        *string      `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	TotalMemory     *float32     `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	UsedMemory      *float32     `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
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
	return dara.Validate(s)
}
