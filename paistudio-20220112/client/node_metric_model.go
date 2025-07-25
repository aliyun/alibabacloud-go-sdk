// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeMetric interface {
	dara.Model
	String() string
	GoString() string
	SetGPUType(v string) *NodeMetric
	GetGPUType() *string
	SetMetrics(v []*Metric) *NodeMetric
	GetMetrics() []*Metric
	SetNodeID(v string) *NodeMetric
	GetNodeID() *string
}

type NodeMetric struct {
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// 23000
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// -i121212node
	NodeID *string `json:"NodeID,omitempty" xml:"NodeID,omitempty"`
}

func (s NodeMetric) String() string {
	return dara.Prettify(s)
}

func (s NodeMetric) GoString() string {
	return s.String()
}

func (s *NodeMetric) GetGPUType() *string {
	return s.GPUType
}

func (s *NodeMetric) GetMetrics() []*Metric {
	return s.Metrics
}

func (s *NodeMetric) GetNodeID() *string {
	return s.NodeID
}

func (s *NodeMetric) SetGPUType(v string) *NodeMetric {
	s.GPUType = &v
	return s
}

func (s *NodeMetric) SetMetrics(v []*Metric) *NodeMetric {
	s.Metrics = v
	return s
}

func (s *NodeMetric) SetNodeID(v string) *NodeMetric {
	s.NodeID = &v
	return s
}

func (s *NodeMetric) Validate() error {
	return dara.Validate(s)
}
