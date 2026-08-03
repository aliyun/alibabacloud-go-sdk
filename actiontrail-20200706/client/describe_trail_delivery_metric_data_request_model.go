// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrailDeliveryMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeTrailDeliveryMetricDataRequest
	GetEndTime() *string
	SetMetricName(v string) *DescribeTrailDeliveryMetricDataRequest
	GetMetricName() *string
	SetPeriod(v int64) *DescribeTrailDeliveryMetricDataRequest
	GetPeriod() *int64
	SetStartTime(v string) *DescribeTrailDeliveryMetricDataRequest
	GetStartTime() *string
	SetTrailName(v string) *DescribeTrailDeliveryMetricDataRequest
	GetTrailName() *string
}

type DescribeTrailDeliveryMetricDataRequest struct {
	// The end of the time window for the query. Specify the time in ISO 8601 format: \\"YYYY-MM-DDThh:mm:ssZ\\". The \\"Z\\" indicates UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-04-10T01:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the delivery monitoring metric. Valid values:
	//
	// - `delivery_sls_success_count`: The number of logs successfully delivered to SLS.
	//
	// - `delivery_sls_fail_count`: The number of logs that failed to be delivered to SLS.
	//
	// - `delivery_oss_success_count`: The number of logs successfully delivered to OSS.
	//
	// - `delivery_oss_fail_count`: The number of logs that failed to be delivered to OSS.
	//
	// This parameter is required.
	//
	// example:
	//
	// delivery_sls_success_count
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The statistical period for the metric data, in seconds. The value must be 60 or a multiple of 60.
	//
	// Recommended values: 60, 900, and 3600.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3600
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The start of the time window for the query. Specify the time in ISO 8601 format: \\"YYYY-MM-DDThh:mm:ssZ\\". The \\"Z\\" indicates UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-04-09T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the trail.
	//
	// This parameter is required.
	//
	// example:
	//
	// trail-name
	TrailName *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
}

func (s DescribeTrailDeliveryMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrailDeliveryMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrailDeliveryMetricDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTrailDeliveryMetricDataRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeTrailDeliveryMetricDataRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *DescribeTrailDeliveryMetricDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTrailDeliveryMetricDataRequest) GetTrailName() *string {
	return s.TrailName
}

func (s *DescribeTrailDeliveryMetricDataRequest) SetEndTime(v string) *DescribeTrailDeliveryMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTrailDeliveryMetricDataRequest) SetMetricName(v string) *DescribeTrailDeliveryMetricDataRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeTrailDeliveryMetricDataRequest) SetPeriod(v int64) *DescribeTrailDeliveryMetricDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeTrailDeliveryMetricDataRequest) SetStartTime(v string) *DescribeTrailDeliveryMetricDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTrailDeliveryMetricDataRequest) SetTrailName(v string) *DescribeTrailDeliveryMetricDataRequest {
	s.TrailName = &v
	return s
}

func (s *DescribeTrailDeliveryMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
