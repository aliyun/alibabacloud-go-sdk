// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentGovernanceMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGenerateTime(v string) *ListRecentGovernanceMetricsResponseBody
	GetGenerateTime() *string
	SetGovernanceMetrics(v *ListRecentGovernanceMetricsResponseBodyGovernanceMetrics) *ListRecentGovernanceMetricsResponseBody
	GetGovernanceMetrics() *ListRecentGovernanceMetricsResponseBodyGovernanceMetrics
	SetRequestId(v string) *ListRecentGovernanceMetricsResponseBody
	GetRequestId() *string
}

type ListRecentGovernanceMetricsResponseBody struct {
	// The time when the report was generated.
	//
	// example:
	//
	// 2025-02-10T02:11:23Z
	GenerateTime *string `json:"GenerateTime,omitempty" xml:"GenerateTime,omitempty"`
	// The metric values of all governance items. The value of the parameter is an array, and each row in the array contains the metric value of a governance item.
	GovernanceMetrics *ListRecentGovernanceMetricsResponseBodyGovernanceMetrics `json:"GovernanceMetrics,omitempty" xml:"GovernanceMetrics,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 49846A91-C1C5-5C2B-BC64-8B0B7BADB4C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRecentGovernanceMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecentGovernanceMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecentGovernanceMetricsResponseBody) GetGenerateTime() *string {
	return s.GenerateTime
}

func (s *ListRecentGovernanceMetricsResponseBody) GetGovernanceMetrics() *ListRecentGovernanceMetricsResponseBodyGovernanceMetrics {
	return s.GovernanceMetrics
}

func (s *ListRecentGovernanceMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecentGovernanceMetricsResponseBody) SetGenerateTime(v string) *ListRecentGovernanceMetricsResponseBody {
	s.GenerateTime = &v
	return s
}

func (s *ListRecentGovernanceMetricsResponseBody) SetGovernanceMetrics(v *ListRecentGovernanceMetricsResponseBodyGovernanceMetrics) *ListRecentGovernanceMetricsResponseBody {
	s.GovernanceMetrics = v
	return s
}

func (s *ListRecentGovernanceMetricsResponseBody) SetRequestId(v string) *ListRecentGovernanceMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecentGovernanceMetricsResponseBody) Validate() error {
	if s.GovernanceMetrics != nil {
		if err := s.GovernanceMetrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRecentGovernanceMetricsResponseBodyGovernanceMetrics struct {
	GovernanceMetric []*ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric `json:"GovernanceMetric,omitempty" xml:"GovernanceMetric,omitempty" type:"Repeated"`
}

func (s ListRecentGovernanceMetricsResponseBodyGovernanceMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListRecentGovernanceMetricsResponseBodyGovernanceMetrics) GoString() string {
	return s.String()
}

func (s *ListRecentGovernanceMetricsResponseBodyGovernanceMetrics) GetGovernanceMetric() []*ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric {
	return s.GovernanceMetric
}

func (s *ListRecentGovernanceMetricsResponseBodyGovernanceMetrics) SetGovernanceMetric(v []*ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) *ListRecentGovernanceMetricsResponseBodyGovernanceMetrics {
	s.GovernanceMetric = v
	return s
}

func (s *ListRecentGovernanceMetricsResponseBodyGovernanceMetrics) Validate() error {
	if s.GovernanceMetric != nil {
		for _, item := range s.GovernanceMetric {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric struct {
	// The name of the governance item.
	//
	// example:
	//
	// RecentAccountLoginTimes
	GovernanceItem *string `json:"GovernanceItem,omitempty" xml:"GovernanceItem,omitempty"`
	// The type of the metric value. Valid values:
	//
	// 	- Number
	//
	// 	- String
	//
	// 	- Boolean
	//
	// example:
	//
	// Number
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The metric value. The type of the metric value is determined by `MetricType`.
	//
	// example:
	//
	// 5
	MetricValue interface{} `json:"MetricValue,omitempty" xml:"MetricValue,omitempty"`
}

func (s ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) String() string {
	return dara.Prettify(s)
}

func (s ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) GoString() string {
	return s.String()
}

func (s *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) GetGovernanceItem() *string {
	return s.GovernanceItem
}

func (s *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) GetMetricType() *string {
	return s.MetricType
}

func (s *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) GetMetricValue() interface{} {
	return s.MetricValue
}

func (s *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) SetGovernanceItem(v string) *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric {
	s.GovernanceItem = &v
	return s
}

func (s *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) SetMetricType(v string) *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric {
	s.MetricType = &v
	return s
}

func (s *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) SetMetricValue(v interface{}) *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric {
	s.MetricValue = v
	return s
}

func (s *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) Validate() error {
	return dara.Validate(s)
}
