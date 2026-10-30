// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCursorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *CursorRequest
	GetEndTime() *string
	SetMatchers(v []*Matcher) *CursorRequest
	GetMatchers() []*Matcher
	SetMetric(v string) *CursorRequest
	GetMetric() *string
	SetNamespace(v string) *CursorRequest
	GetNamespace() *string
	SetPeriod(v int32) *CursorRequest
	GetPeriod() *int32
	SetStartTime(v string) *CursorRequest
	GetStartTime() *string
}

type CursorRequest struct {
	// The end timestamp.
	//
	// Unit: milliseconds.
	//
	// > - UNIX timestamp: the number of milliseconds that have elapsed since January 1, 1970. The format YYYY-MM-DDThh:mm:ssZ is supported. Example: 2023-01-01T00:00:00Z, which specifies 00:00:00 on January 1, 2023 (GMT).
	//
	// - If you do not specify the end time, the end time is considered infinite. In real-time export scenarios, you can leave this parameter empty.
	//
	// - The TTL of monitoring data in CloudMonitor varies with the time granularity. Refer to the data TTL corresponding to the Period request parameter to set a reasonable time interval.
	//
	// example:
	//
	// 1641645000000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The dimension information of the metric.
	Matchers []*Matcher `json:"Matchers,omitempty" xml:"Matchers,omitempty" type:"Repeated"`
	// The name of the metric for the cloud service.
	//
	// For more information about how to obtain the metric name of a cloud service, see [Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_idle
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The data namespace of the cloud service.
	//
	// For more information about how to obtain the data namespace of a cloud service, see [Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The time interval of the metric.
	//
	// Unit: seconds.
	//
	// > The time interval of a metric is typically 60 seconds. For special values, see the Period parameter in [Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The start timestamp.
	//
	// The format YYYY-MM-DDThh:mm:ssZ is supported. Example: 2023-01-01T00:00:00Z, which specifies 00:00:00 on January 1, 2023 (GMT).
	//
	// > The time-to-live (TTL) of monitoring data in CloudMonitor varies with the time granularity. Refer to the data TTL corresponding to the Period request parameter to set a reasonable time interval.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1641627000000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CursorRequest) String() string {
	return dara.Prettify(s)
}

func (s CursorRequest) GoString() string {
	return s.String()
}

func (s *CursorRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CursorRequest) GetMatchers() []*Matcher {
	return s.Matchers
}

func (s *CursorRequest) GetMetric() *string {
	return s.Metric
}

func (s *CursorRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CursorRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CursorRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CursorRequest) SetEndTime(v string) *CursorRequest {
	s.EndTime = &v
	return s
}

func (s *CursorRequest) SetMatchers(v []*Matcher) *CursorRequest {
	s.Matchers = v
	return s
}

func (s *CursorRequest) SetMetric(v string) *CursorRequest {
	s.Metric = &v
	return s
}

func (s *CursorRequest) SetNamespace(v string) *CursorRequest {
	s.Namespace = &v
	return s
}

func (s *CursorRequest) SetPeriod(v int32) *CursorRequest {
	s.Period = &v
	return s
}

func (s *CursorRequest) SetStartTime(v string) *CursorRequest {
	s.StartTime = &v
	return s
}

func (s *CursorRequest) Validate() error {
	if s.Matchers != nil {
		for _, item := range s.Matchers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
