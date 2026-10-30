// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCursorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *CursorShrinkRequest
	GetEndTime() *string
	SetMatchersShrink(v string) *CursorShrinkRequest
	GetMatchersShrink() *string
	SetMetric(v string) *CursorShrinkRequest
	GetMetric() *string
	SetNamespace(v string) *CursorShrinkRequest
	GetNamespace() *string
	SetPeriod(v int32) *CursorShrinkRequest
	GetPeriod() *int32
	SetStartTime(v string) *CursorShrinkRequest
	GetStartTime() *string
}

type CursorShrinkRequest struct {
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
	MatchersShrink *string `json:"Matchers,omitempty" xml:"Matchers,omitempty"`
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

func (s CursorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CursorShrinkRequest) GoString() string {
	return s.String()
}

func (s *CursorShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CursorShrinkRequest) GetMatchersShrink() *string {
	return s.MatchersShrink
}

func (s *CursorShrinkRequest) GetMetric() *string {
	return s.Metric
}

func (s *CursorShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CursorShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CursorShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CursorShrinkRequest) SetEndTime(v string) *CursorShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *CursorShrinkRequest) SetMatchersShrink(v string) *CursorShrinkRequest {
	s.MatchersShrink = &v
	return s
}

func (s *CursorShrinkRequest) SetMetric(v string) *CursorShrinkRequest {
	s.Metric = &v
	return s
}

func (s *CursorShrinkRequest) SetNamespace(v string) *CursorShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *CursorShrinkRequest) SetPeriod(v int32) *CursorShrinkRequest {
	s.Period = &v
	return s
}

func (s *CursorShrinkRequest) SetStartTime(v string) *CursorShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *CursorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
