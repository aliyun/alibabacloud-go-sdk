// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobInstanceMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceMetrics(v []*ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) *ListTrainingJobInstanceMetricsResponseBody
	GetInstanceMetrics() []*ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics
	SetRequestId(v string) *ListTrainingJobInstanceMetricsResponseBody
	GetRequestId() *string
}

type ListTrainingJobInstanceMetricsResponseBody struct {
	// List of all monitoring metrics that match the filter condition.
	InstanceMetrics []*ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics `json:"InstanceMetrics,omitempty" xml:"InstanceMetrics,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// F082BD0D-21E1-5F9B-81A0-AB07485B03CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTrainingJobInstanceMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobInstanceMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceMetricsResponseBody) GetInstanceMetrics() []*ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics {
	return s.InstanceMetrics
}

func (s *ListTrainingJobInstanceMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrainingJobInstanceMetricsResponseBody) SetInstanceMetrics(v []*ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) *ListTrainingJobInstanceMetricsResponseBody {
	s.InstanceMetrics = v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponseBody) SetRequestId(v string) *ListTrainingJobInstanceMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponseBody) Validate() error {
	if s.InstanceMetrics != nil {
		for _, item := range s.InstanceMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics struct {
	// Instance ID.
	//
	// example:
	//
	// trainkxen7qjyg6y-master-0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// List of instance monitoring metrics.
	Metrics []*ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// Node name.
	//
	// example:
	//
	// trains930928remn-master-0
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) GetMetrics() []*ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics {
	return s.Metrics
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) GetNodeName() *string {
	return s.NodeName
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) SetInstanceId(v string) *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics {
	s.InstanceId = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) SetMetrics(v []*ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics) *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics {
	s.Metrics = v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) SetNodeName(v string) *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics {
	s.NodeName = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetrics) Validate() error {
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

type ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics struct {
	// UTC time in ISO 8601 format.
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// Metric value.
	//
	// example:
	//
	// 1
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics) GetTime() *string {
	return s.Time
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics) GetValue() *float64 {
	return s.Value
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics) SetTime(v string) *ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics {
	s.Time = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics) SetValue(v float64) *ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics {
	s.Value = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponseBodyInstanceMetricsMetrics) Validate() error {
	return dara.Validate(s)
}
