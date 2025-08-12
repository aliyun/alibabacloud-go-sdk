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
	// >
	//
	// 	- This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. For example, 2023-01-01T00:00:00Z indicates January 1, 2023, 00:00:00 UTC.
	//
	// 	- If you do not set the end time, the end time is infinite. You can leave this parameter empty in real-time export scenarios.
	//
	// 	- In CloudMonitor, the TTL of monitoring data varies with the time granularity. Specify a proper time interval based on the TTL corresponding to the value of the `Period` parameter.
	//
	// example:
	//
	// 1641645000000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The dimension information of the metric.
	MatchersShrink *string `json:"Matchers,omitempty" xml:"Matchers,omitempty"`
	// The metric that is used to monitor the cloud service.
	//
	// For more information about the metrics of cloud services, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_idle
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The namespace of the cloud service.
	//
	// For more information about the namespaces of cloud services, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The time interval based on which the metric value is measured.
	//
	// Unit: seconds.
	//
	// >  Generally, the time interval is 60 seconds. For more information about specific values, see the `Period` parameter in [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. For example, 2023-01-01T00:00:00Z indicates January 1, 2023, 00:00:00 UTC.
	//
	// >  In CloudMonitor, the TTL of monitoring data varies with the time granularity. Specify a proper time interval based on the TTL corresponding to the value of the `Period` parameter.
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
