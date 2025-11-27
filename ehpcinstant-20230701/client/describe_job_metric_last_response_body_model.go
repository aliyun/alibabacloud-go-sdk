// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobMetricLastResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetrics(v []*DescribeJobMetricLastResponseBodyMetrics) *DescribeJobMetricLastResponseBody
	GetMetrics() []*DescribeJobMetricLastResponseBodyMetrics
	SetRequestId(v string) *DescribeJobMetricLastResponseBody
	GetRequestId() *string
}

type DescribeJobMetricLastResponseBody struct {
	// The list of the JobMetric details.
	Metrics []*DescribeJobMetricLastResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeJobMetricLastResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobMetricLastResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricLastResponseBody) GetMetrics() []*DescribeJobMetricLastResponseBodyMetrics {
	return s.Metrics
}

func (s *DescribeJobMetricLastResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJobMetricLastResponseBody) SetMetrics(v []*DescribeJobMetricLastResponseBodyMetrics) *DescribeJobMetricLastResponseBody {
	s.Metrics = v
	return s
}

func (s *DescribeJobMetricLastResponseBody) SetRequestId(v string) *DescribeJobMetricLastResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobMetricLastResponseBody) Validate() error {
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeJobMetricLastResponseBodyMetrics struct {
	// The index of the array job.
	//
	// example:
	//
	// 1
	ArrayIndex *int32 `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty"`
	// The monitoring item information corresponding to the array job index.
	//
	// example:
	//
	// {"memory_utilization": 37.42,"cpu_utilization": 1.008, "disk_utilization": 3.562}
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
}

func (s DescribeJobMetricLastResponseBodyMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobMetricLastResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricLastResponseBodyMetrics) GetArrayIndex() *int32 {
	return s.ArrayIndex
}

func (s *DescribeJobMetricLastResponseBodyMetrics) GetMetric() *string {
	return s.Metric
}

func (s *DescribeJobMetricLastResponseBodyMetrics) SetArrayIndex(v int32) *DescribeJobMetricLastResponseBodyMetrics {
	s.ArrayIndex = &v
	return s
}

func (s *DescribeJobMetricLastResponseBodyMetrics) SetMetric(v string) *DescribeJobMetricLastResponseBodyMetrics {
	s.Metric = &v
	return s
}

func (s *DescribeJobMetricLastResponseBodyMetrics) Validate() error {
	return dara.Validate(s)
}
