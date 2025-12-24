// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoricalMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHistoricalMetrics(v *QueryHistoricalMetricResponseBodyHistoricalMetrics) *QueryHistoricalMetricResponseBody
	GetHistoricalMetrics() *QueryHistoricalMetricResponseBodyHistoricalMetrics
	SetRequestId(v string) *QueryHistoricalMetricResponseBody
	GetRequestId() *string
}

type QueryHistoricalMetricResponseBody struct {
	HistoricalMetrics *QueryHistoricalMetricResponseBodyHistoricalMetrics `json:"HistoricalMetrics,omitempty" xml:"HistoricalMetrics,omitempty" type:"Struct"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryHistoricalMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoricalMetricResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHistoricalMetricResponseBody) GetHistoricalMetrics() *QueryHistoricalMetricResponseBodyHistoricalMetrics {
	return s.HistoricalMetrics
}

func (s *QueryHistoricalMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryHistoricalMetricResponseBody) SetHistoricalMetrics(v *QueryHistoricalMetricResponseBodyHistoricalMetrics) *QueryHistoricalMetricResponseBody {
	s.HistoricalMetrics = v
	return s
}

func (s *QueryHistoricalMetricResponseBody) SetRequestId(v string) *QueryHistoricalMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHistoricalMetricResponseBody) Validate() error {
	if s.HistoricalMetrics != nil {
		if err := s.HistoricalMetrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryHistoricalMetricResponseBodyHistoricalMetrics struct {
	HistoricalMetric []*QueryHistoricalMetricResponseBodyHistoricalMetricsHistoricalMetric `json:"HistoricalMetric,omitempty" xml:"HistoricalMetric,omitempty" type:"Repeated"`
}

func (s QueryHistoricalMetricResponseBodyHistoricalMetrics) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoricalMetricResponseBodyHistoricalMetrics) GoString() string {
	return s.String()
}

func (s *QueryHistoricalMetricResponseBodyHistoricalMetrics) GetHistoricalMetric() []*QueryHistoricalMetricResponseBodyHistoricalMetricsHistoricalMetric {
	return s.HistoricalMetric
}

func (s *QueryHistoricalMetricResponseBodyHistoricalMetrics) SetHistoricalMetric(v []*QueryHistoricalMetricResponseBodyHistoricalMetricsHistoricalMetric) *QueryHistoricalMetricResponseBodyHistoricalMetrics {
	s.HistoricalMetric = v
	return s
}

func (s *QueryHistoricalMetricResponseBodyHistoricalMetrics) Validate() error {
	if s.HistoricalMetric != nil {
		for _, item := range s.HistoricalMetric {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryHistoricalMetricResponseBodyHistoricalMetricsHistoricalMetric struct {
	// example:
	//
	// 10.0
	MetricValue *string `json:"MetricValue,omitempty" xml:"MetricValue,omitempty"`
	// example:
	//
	// 2025-12-17T16:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryHistoricalMetricResponseBodyHistoricalMetricsHistoricalMetric) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoricalMetricResponseBodyHistoricalMetricsHistoricalMetric) GoString() string {
	return s.String()
}

func (s *QueryHistoricalMetricResponseBodyHistoricalMetricsHistoricalMetric) GetMetricValue() *string {
	return s.MetricValue
}

func (s *QueryHistoricalMetricResponseBodyHistoricalMetricsHistoricalMetric) GetTime() *string {
	return s.Time
}

func (s *QueryHistoricalMetricResponseBodyHistoricalMetricsHistoricalMetric) SetMetricValue(v string) *QueryHistoricalMetricResponseBodyHistoricalMetricsHistoricalMetric {
	s.MetricValue = &v
	return s
}

func (s *QueryHistoricalMetricResponseBodyHistoricalMetricsHistoricalMetric) SetTime(v string) *QueryHistoricalMetricResponseBodyHistoricalMetricsHistoricalMetric {
	s.Time = &v
	return s
}

func (s *QueryHistoricalMetricResponseBodyHistoricalMetricsHistoricalMetric) Validate() error {
	return dara.Validate(s)
}
