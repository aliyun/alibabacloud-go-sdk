// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggreOps(v string) *DescribeMetricDataShrinkRequest
	GetAggreOps() *string
	SetAggreOverLineOps(v string) *DescribeMetricDataShrinkRequest
	GetAggreOverLineOps() *string
	SetDimensions(v string) *DescribeMetricDataShrinkRequest
	GetDimensions() *string
	SetEndTime(v string) *DescribeMetricDataShrinkRequest
	GetEndTime() *string
	SetGroupByLabelsShrink(v string) *DescribeMetricDataShrinkRequest
	GetGroupByLabelsShrink() *string
	SetMetricName(v string) *DescribeMetricDataShrinkRequest
	GetMetricName() *string
	SetPeriod(v int32) *DescribeMetricDataShrinkRequest
	GetPeriod() *int32
	SetRegionId(v string) *DescribeMetricDataShrinkRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeMetricDataShrinkRequest
	GetStartTime() *string
}

type DescribeMetricDataShrinkRequest struct {
	// The method for aggregating data over time. Valid values:
	//
	// - SUM_OVER_TIME
	//
	// - COUNT_OVER_TIME
	//
	// - AVG_OVER_TIME
	//
	// - MAX_OVER_TIME
	//
	// - MIN_OVER_TIME
	//
	// - SUM_OVER_TIME_LCRO: The sum of values in a left-closed, right-open interval.
	//
	// - AVG_OVER_TIME_LCRO: The average of values in a left-closed, right-open interval.
	//
	// - SUM_OVER_TIME_LORC: The sum of values in a left-open, right-closed interval.
	//
	// - AVG_OVER_TIME_LORC: The average of values in a left-open, right-closed interval.
	//
	// example:
	//
	// AVG_OVER_TIME
	AggreOps *string `json:"AggreOps,omitempty" xml:"AggreOps,omitempty"`
	// The method for aggregating data across different lines. Valid values:
	//
	// - NON: No aggregation is performed.
	//
	// - SUM: The sum of values.
	//
	// - AVG: The average of values.
	//
	// - COUNT: The number of values.
	//
	// - MAX: The maximum value.
	//
	// - MIN: The minimum value.
	//
	// example:
	//
	// NON
	AggreOverLineOps *string `json:"AggreOverLineOps,omitempty" xml:"AggreOverLineOps,omitempty"`
	// A map of dimensions in the JSON format. The map specifies the dimensions to query. The following keys are supported:
	//
	// - DiskId: The disk name, such as d-xxx.
	//
	// - DeviceType: The disk category. \\`system\\` indicates a system disk and \\`data\\` indicates a data disk.
	//
	// - DeviceCategory: The disk type, such as cloud_essd.
	//
	// - EcsInstanceId: The name of the ECS instance to which the disk is attached, such as i-xxx.
	//
	// - Azone: The zone, such as cn-hangzhou-a.
	//
	// The returned results are the intersection of all specified dimension-based filter conditions.
	//
	// example:
	//
	// {"DiskId":["d-bp14xxxx","d-bp11xxxx"], "DeviceCategory": ["cloud_essd"]}
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The end of the time range to query metric data. The time cannot be later than the current time. The time must be in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-11-21T02:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// A list of fields for grouping and aggregation.
	GroupByLabelsShrink *string `json:"GroupByLabels,omitempty" xml:"GroupByLabels,omitempty"`
	// The name of the metric. Valid values:
	//
	// - disk_bps_percent
	//
	// - disk_iops_percent
	//
	// - disk_read_block_size
	//
	// - disk_read_bps
	//
	// - disk_read_iops
	//
	// - disk_write_block_size
	//
	// - disk_write_bps
	//
	// - disk_write_iops
	//
	// This parameter is required.
	//
	// example:
	//
	// disk_bps_percent
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The interval at which to query metric data. Unit: seconds. The default value is 5. Valid values:
	//
	// - 5: 5-second precision. You can query data within a 12-hour time range.
	//
	// - 10: 10-second precision. You can query data within a 24-hour time range.
	//
	// - 60: 60-second precision. You can query data within a 7-day time range.
	//
	// - 300: 300-second precision. You can query data within a 30-day time range.
	//
	// - 600: 600-second precision. You can query data within a 30-day time range.
	//
	// - 3600: 3600-second precision. You can query data within a 30-day time range.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query metric data. The start time can be up to 30 days before the current time. If you leave both the StartTime and EndTime parameters empty, the system queries the metrics for the most recent period. The time must be in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-11-21T01:50:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMetricDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricDataShrinkRequest) GetAggreOps() *string {
	return s.AggreOps
}

func (s *DescribeMetricDataShrinkRequest) GetAggreOverLineOps() *string {
	return s.AggreOverLineOps
}

func (s *DescribeMetricDataShrinkRequest) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeMetricDataShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMetricDataShrinkRequest) GetGroupByLabelsShrink() *string {
	return s.GroupByLabelsShrink
}

func (s *DescribeMetricDataShrinkRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricDataShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeMetricDataShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMetricDataShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMetricDataShrinkRequest) SetAggreOps(v string) *DescribeMetricDataShrinkRequest {
	s.AggreOps = &v
	return s
}

func (s *DescribeMetricDataShrinkRequest) SetAggreOverLineOps(v string) *DescribeMetricDataShrinkRequest {
	s.AggreOverLineOps = &v
	return s
}

func (s *DescribeMetricDataShrinkRequest) SetDimensions(v string) *DescribeMetricDataShrinkRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricDataShrinkRequest) SetEndTime(v string) *DescribeMetricDataShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricDataShrinkRequest) SetGroupByLabelsShrink(v string) *DescribeMetricDataShrinkRequest {
	s.GroupByLabelsShrink = &v
	return s
}

func (s *DescribeMetricDataShrinkRequest) SetMetricName(v string) *DescribeMetricDataShrinkRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricDataShrinkRequest) SetPeriod(v int32) *DescribeMetricDataShrinkRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricDataShrinkRequest) SetRegionId(v string) *DescribeMetricDataShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMetricDataShrinkRequest) SetStartTime(v string) *DescribeMetricDataShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
