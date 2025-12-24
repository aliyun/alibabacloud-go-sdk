// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPredictiveMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPredictiveMetrics(v *QueryPredictiveMetricResponseBodyPredictiveMetrics) *QueryPredictiveMetricResponseBody
	GetPredictiveMetrics() *QueryPredictiveMetricResponseBodyPredictiveMetrics
	SetRequestId(v string) *QueryPredictiveMetricResponseBody
	GetRequestId() *string
}

type QueryPredictiveMetricResponseBody struct {
	PredictiveMetrics *QueryPredictiveMetricResponseBodyPredictiveMetrics `json:"PredictiveMetrics,omitempty" xml:"PredictiveMetrics,omitempty" type:"Struct"`
	// example:
	//
	// CC107349-57B7-4405-B1BF-9BF5AF7F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPredictiveMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPredictiveMetricResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPredictiveMetricResponseBody) GetPredictiveMetrics() *QueryPredictiveMetricResponseBodyPredictiveMetrics {
	return s.PredictiveMetrics
}

func (s *QueryPredictiveMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPredictiveMetricResponseBody) SetPredictiveMetrics(v *QueryPredictiveMetricResponseBodyPredictiveMetrics) *QueryPredictiveMetricResponseBody {
	s.PredictiveMetrics = v
	return s
}

func (s *QueryPredictiveMetricResponseBody) SetRequestId(v string) *QueryPredictiveMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPredictiveMetricResponseBody) Validate() error {
	if s.PredictiveMetrics != nil {
		if err := s.PredictiveMetrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPredictiveMetricResponseBodyPredictiveMetrics struct {
	PredictiveMetric []*QueryPredictiveMetricResponseBodyPredictiveMetricsPredictiveMetric `json:"PredictiveMetric,omitempty" xml:"PredictiveMetric,omitempty" type:"Repeated"`
}

func (s QueryPredictiveMetricResponseBodyPredictiveMetrics) String() string {
	return dara.Prettify(s)
}

func (s QueryPredictiveMetricResponseBodyPredictiveMetrics) GoString() string {
	return s.String()
}

func (s *QueryPredictiveMetricResponseBodyPredictiveMetrics) GetPredictiveMetric() []*QueryPredictiveMetricResponseBodyPredictiveMetricsPredictiveMetric {
	return s.PredictiveMetric
}

func (s *QueryPredictiveMetricResponseBodyPredictiveMetrics) SetPredictiveMetric(v []*QueryPredictiveMetricResponseBodyPredictiveMetricsPredictiveMetric) *QueryPredictiveMetricResponseBodyPredictiveMetrics {
	s.PredictiveMetric = v
	return s
}

func (s *QueryPredictiveMetricResponseBodyPredictiveMetrics) Validate() error {
	if s.PredictiveMetric != nil {
		for _, item := range s.PredictiveMetric {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPredictiveMetricResponseBodyPredictiveMetricsPredictiveMetric struct {
	// example:
	//
	// 10.0
	MetricValue *string `json:"MetricValue,omitempty" xml:"MetricValue,omitempty"`
	// example:
	//
	// 2025-12-17T16:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryPredictiveMetricResponseBodyPredictiveMetricsPredictiveMetric) String() string {
	return dara.Prettify(s)
}

func (s QueryPredictiveMetricResponseBodyPredictiveMetricsPredictiveMetric) GoString() string {
	return s.String()
}

func (s *QueryPredictiveMetricResponseBodyPredictiveMetricsPredictiveMetric) GetMetricValue() *string {
	return s.MetricValue
}

func (s *QueryPredictiveMetricResponseBodyPredictiveMetricsPredictiveMetric) GetTime() *string {
	return s.Time
}

func (s *QueryPredictiveMetricResponseBodyPredictiveMetricsPredictiveMetric) SetMetricValue(v string) *QueryPredictiveMetricResponseBodyPredictiveMetricsPredictiveMetric {
	s.MetricValue = &v
	return s
}

func (s *QueryPredictiveMetricResponseBodyPredictiveMetricsPredictiveMetric) SetTime(v string) *QueryPredictiveMetricResponseBodyPredictiveMetricsPredictiveMetric {
	s.Time = &v
	return s
}

func (s *QueryPredictiveMetricResponseBodyPredictiveMetricsPredictiveMetric) Validate() error {
	return dara.Validate(s)
}
