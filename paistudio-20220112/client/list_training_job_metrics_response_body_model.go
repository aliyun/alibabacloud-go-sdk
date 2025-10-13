// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetrics(v []*ListTrainingJobMetricsResponseBodyMetrics) *ListTrainingJobMetricsResponseBody
	GetMetrics() []*ListTrainingJobMetricsResponseBodyMetrics
	SetRequestId(v string) *ListTrainingJobMetricsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListTrainingJobMetricsResponseBody
	GetTotalCount() *int64
}

type ListTrainingJobMetricsResponseBody struct {
	Metrics []*ListTrainingJobMetricsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTrainingJobMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainingJobMetricsResponseBody) GetMetrics() []*ListTrainingJobMetricsResponseBodyMetrics {
	return s.Metrics
}

func (s *ListTrainingJobMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrainingJobMetricsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTrainingJobMetricsResponseBody) SetMetrics(v []*ListTrainingJobMetricsResponseBodyMetrics) *ListTrainingJobMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *ListTrainingJobMetricsResponseBody) SetRequestId(v string) *ListTrainingJobMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrainingJobMetricsResponseBody) SetTotalCount(v int64) *ListTrainingJobMetricsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrainingJobMetricsResponseBody) Validate() error {
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

type ListTrainingJobMetricsResponseBodyMetrics struct {
	// example:
	//
	// accuracy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-04-18T22:20:55Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 0.97
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrainingJobMetricsResponseBodyMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobMetricsResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *ListTrainingJobMetricsResponseBodyMetrics) GetName() *string {
	return s.Name
}

func (s *ListTrainingJobMetricsResponseBodyMetrics) GetTimestamp() *string {
	return s.Timestamp
}

func (s *ListTrainingJobMetricsResponseBodyMetrics) GetValue() *float64 {
	return s.Value
}

func (s *ListTrainingJobMetricsResponseBodyMetrics) SetName(v string) *ListTrainingJobMetricsResponseBodyMetrics {
	s.Name = &v
	return s
}

func (s *ListTrainingJobMetricsResponseBodyMetrics) SetTimestamp(v string) *ListTrainingJobMetricsResponseBodyMetrics {
	s.Timestamp = &v
	return s
}

func (s *ListTrainingJobMetricsResponseBodyMetrics) SetValue(v float64) *ListTrainingJobMetricsResponseBodyMetrics {
	s.Value = &v
	return s
}

func (s *ListTrainingJobMetricsResponseBodyMetrics) Validate() error {
	return dara.Validate(s)
}
