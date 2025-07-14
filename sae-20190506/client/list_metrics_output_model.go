// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetricsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListMetricsOutput
	GetRequestId() *string
	SetMetrics(v map[string][]*MetricInfo) *ListMetricsOutput
	GetMetrics() map[string][]*MetricInfo
}

type ListMetricsOutput struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Metrics   map[string][]*MetricInfo `json:"metrics,omitempty" xml:"metrics,omitempty"`
}

func (s ListMetricsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListMetricsOutput) GoString() string {
	return s.String()
}

func (s *ListMetricsOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMetricsOutput) GetMetrics() map[string][]*MetricInfo {
	return s.Metrics
}

func (s *ListMetricsOutput) SetRequestId(v string) *ListMetricsOutput {
	s.RequestId = &v
	return s
}

func (s *ListMetricsOutput) SetMetrics(v map[string][]*MetricInfo) *ListMetricsOutput {
	s.Metrics = v
	return s
}

func (s *ListMetricsOutput) Validate() error {
	return dara.Validate(s)
}
