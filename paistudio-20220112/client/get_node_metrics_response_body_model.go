// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetricType(v string) *GetNodeMetricsResponseBody
	GetMetricType() *string
	SetNodesMetrics(v []*NodeMetric) *GetNodeMetricsResponseBody
	GetNodesMetrics() []*NodeMetric
	SetResourceGroupID(v string) *GetNodeMetricsResponseBody
	GetResourceGroupID() *string
}

type GetNodeMetricsResponseBody struct {
	// example:
	//
	// DiskWriteRate
	MetricType   *string       `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	NodesMetrics []*NodeMetric `json:"NodesMetrics,omitempty" xml:"NodesMetrics,omitempty" type:"Repeated"`
	// example:
	//
	// rgf0zhfqn1d4ity2
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
}

func (s GetNodeMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeMetricsResponseBody) GetMetricType() *string {
	return s.MetricType
}

func (s *GetNodeMetricsResponseBody) GetNodesMetrics() []*NodeMetric {
	return s.NodesMetrics
}

func (s *GetNodeMetricsResponseBody) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *GetNodeMetricsResponseBody) SetMetricType(v string) *GetNodeMetricsResponseBody {
	s.MetricType = &v
	return s
}

func (s *GetNodeMetricsResponseBody) SetNodesMetrics(v []*NodeMetric) *GetNodeMetricsResponseBody {
	s.NodesMetrics = v
	return s
}

func (s *GetNodeMetricsResponseBody) SetResourceGroupID(v string) *GetNodeMetricsResponseBody {
	s.ResourceGroupID = &v
	return s
}

func (s *GetNodeMetricsResponseBody) Validate() error {
	return dara.Validate(s)
}
