// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrainingJobLatestMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetrics(v []*GetTrainingJobLatestMetricsResponseBodyMetrics) *GetTrainingJobLatestMetricsResponseBody
	GetMetrics() []*GetTrainingJobLatestMetricsResponseBodyMetrics
	SetRequestId(v string) *GetTrainingJobLatestMetricsResponseBody
	GetRequestId() *string
}

type GetTrainingJobLatestMetricsResponseBody struct {
	Metrics []*GetTrainingJobLatestMetricsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// 18D5A1C6-14B8-545E-8408-0A7DDB4C6B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTrainingJobLatestMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobLatestMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrainingJobLatestMetricsResponseBody) GetMetrics() []*GetTrainingJobLatestMetricsResponseBodyMetrics {
	return s.Metrics
}

func (s *GetTrainingJobLatestMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrainingJobLatestMetricsResponseBody) SetMetrics(v []*GetTrainingJobLatestMetricsResponseBodyMetrics) *GetTrainingJobLatestMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *GetTrainingJobLatestMetricsResponseBody) SetRequestId(v string) *GetTrainingJobLatestMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrainingJobLatestMetricsResponseBody) Validate() error {
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

type GetTrainingJobLatestMetricsResponseBodyMetrics struct {
	// example:
	//
	// loss
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

func (s GetTrainingJobLatestMetricsResponseBodyMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobLatestMetricsResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *GetTrainingJobLatestMetricsResponseBodyMetrics) GetName() *string {
	return s.Name
}

func (s *GetTrainingJobLatestMetricsResponseBodyMetrics) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetTrainingJobLatestMetricsResponseBodyMetrics) GetValue() *float64 {
	return s.Value
}

func (s *GetTrainingJobLatestMetricsResponseBodyMetrics) SetName(v string) *GetTrainingJobLatestMetricsResponseBodyMetrics {
	s.Name = &v
	return s
}

func (s *GetTrainingJobLatestMetricsResponseBodyMetrics) SetTimestamp(v string) *GetTrainingJobLatestMetricsResponseBodyMetrics {
	s.Timestamp = &v
	return s
}

func (s *GetTrainingJobLatestMetricsResponseBodyMetrics) SetValue(v float64) *GetTrainingJobLatestMetricsResponseBodyMetrics {
	s.Value = &v
	return s
}

func (s *GetTrainingJobLatestMetricsResponseBodyMetrics) Validate() error {
	return dara.Validate(s)
}
