// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogRunMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMetrics(v []*RunMetric) *LogRunMetricsRequest
	GetMetrics() []*RunMetric
}

type LogRunMetricsRequest struct {
	// The metrics.
	Metrics []*RunMetric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
}

func (s LogRunMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s LogRunMetricsRequest) GoString() string {
	return s.String()
}

func (s *LogRunMetricsRequest) GetMetrics() []*RunMetric {
	return s.Metrics
}

func (s *LogRunMetricsRequest) SetMetrics(v []*RunMetric) *LogRunMetricsRequest {
	s.Metrics = v
	return s
}

func (s *LogRunMetricsRequest) Validate() error {
	return dara.Validate(s)
}
