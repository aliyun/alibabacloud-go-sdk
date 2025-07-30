// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeMetric interface {
	dara.Model
	String() string
	GoString() string
	SetMetrics(v []*Metric) *NodeMetric
	GetMetrics() []*Metric
	SetNodeName(v string) *NodeMetric
	GetNodeName() *string
}

type NodeMetric struct {
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// asi_xxx
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s NodeMetric) String() string {
	return dara.Prettify(s)
}

func (s NodeMetric) GoString() string {
	return s.String()
}

func (s *NodeMetric) GetMetrics() []*Metric {
	return s.Metrics
}

func (s *NodeMetric) GetNodeName() *string {
	return s.NodeName
}

func (s *NodeMetric) SetMetrics(v []*Metric) *NodeMetric {
	s.Metrics = v
	return s
}

func (s *NodeMetric) SetNodeName(v string) *NodeMetric {
	s.NodeName = &v
	return s
}

func (s *NodeMetric) Validate() error {
	return dara.Validate(s)
}
