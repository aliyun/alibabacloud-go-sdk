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
	// The end of the time range to query.
	//
	// Unit: milliseconds.
	//
	// > - Unix timestamp: the number of milliseconds that have elapsed since 00:00:00 on January 1, 1970. The value is in the YYYY-MM-DDThh:mm:ssZ format. For example, 2023-01-01T00:00:00Z indicates 00:00:00 on January 1, 2023 (GMT).
	//
	// - If you do not specify an end time, the end time is unlimited. You do not need to specify this parameter when you export data in real time.
	//
	// - The time to live (TTL) of monitoring data varies based on the statistical granularity in CloudMonitor. Configure a proper time range based on the TTL of the data that corresponds to the `Period` parameter.
	//
	// example:
	//
	// 1641645000000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The dimension information of the metric.
	MatchersShrink *string `json:"Matchers,omitempty" xml:"Matchers,omitempty"`
	// The metric name of the cloud service.
	//
	// For information about how to obtain the metric name of a cloud service, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_idle
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The data namespace of the cloud service.
	//
	// For information about how to obtain the data namespace of a cloud service, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The statistical period of the metric.
	//
	// Unit: seconds.
	//
	// > The statistical period of a metric is typically 60 seconds. For special values, see the `Period` parameter in [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The beginning of the time range to query.
	//
	// The value is in the YYYY-MM-DDThh:mm:ssZ format. For example, 2023-01-01T00:00:00Z indicates 00:00:00 on January 1, 2023 (GMT).
	//
	// > The time to live (TTL) of monitoring data varies based on the statistical granularity in CloudMonitor. Configure a proper time range based on the TTL of the data that corresponds to the `Period` parameter.
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
