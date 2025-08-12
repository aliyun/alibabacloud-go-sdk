// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMetricName(v string) *DescribeSiteMonitorStatisticsRequest
	GetMetricName() *string
	SetRegionId(v string) *DescribeSiteMonitorStatisticsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeSiteMonitorStatisticsRequest
	GetStartTime() *string
	SetTaskId(v string) *DescribeSiteMonitorStatisticsRequest
	GetTaskId() *string
	SetTimeRange(v string) *DescribeSiteMonitorStatisticsRequest
	GetTimeRange() *string
}

type DescribeSiteMonitorStatisticsRequest struct {
	// The metric name. Valid values:
	//
	// 	- Availability
	//
	// 	- ErrorRate
	//
	// 	- ResponseTime
	//
	// This parameter is required.
	//
	// example:
	//
	// Availability
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query.
	//
	// Unit: milliseconds. The default value is 1 hour ahead of the current time.
	//
	// example:
	//
	// 1576142850527
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the site monitoring task.
	//
	// For more information about how to obtain the ID of a site monitoring task, see [DescribeSiteMonitorList](https://help.aliyun.com/document_detail/115052.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ef4cdc8b-9dc7-43e7-810e-f950e56c****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The statistical period.
	//
	// Unit: minutes. Default value: 1440 (one day). Maximum value: 43200 (30 days).
	//
	// example:
	//
	// 1440
	TimeRange *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s DescribeSiteMonitorStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorStatisticsRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeSiteMonitorStatisticsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSiteMonitorStatisticsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteMonitorStatisticsRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeSiteMonitorStatisticsRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *DescribeSiteMonitorStatisticsRequest) SetMetricName(v string) *DescribeSiteMonitorStatisticsRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsRequest) SetRegionId(v string) *DescribeSiteMonitorStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsRequest) SetStartTime(v string) *DescribeSiteMonitorStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsRequest) SetTaskId(v string) *DescribeSiteMonitorStatisticsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsRequest) SetTimeRange(v string) *DescribeSiteMonitorStatisticsRequest {
	s.TimeRange = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
