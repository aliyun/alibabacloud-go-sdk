// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggreOps(v string) *DescribeMetricDataRequest
	GetAggreOps() *string
	SetAggreOverLineOps(v string) *DescribeMetricDataRequest
	GetAggreOverLineOps() *string
	SetDimensions(v string) *DescribeMetricDataRequest
	GetDimensions() *string
	SetEndTime(v string) *DescribeMetricDataRequest
	GetEndTime() *string
	SetGroupByLabels(v []*string) *DescribeMetricDataRequest
	GetGroupByLabels() []*string
	SetMetricName(v string) *DescribeMetricDataRequest
	GetMetricName() *string
	SetPeriod(v int32) *DescribeMetricDataRequest
	GetPeriod() *int32
	SetRegionId(v string) *DescribeMetricDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeMetricDataRequest
	GetStartTime() *string
}

type DescribeMetricDataRequest struct {
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
	GroupByLabels []*string `json:"GroupByLabels,omitempty" xml:"GroupByLabels,omitempty" type:"Repeated"`
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

func (s DescribeMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricDataRequest) GetAggreOps() *string {
	return s.AggreOps
}

func (s *DescribeMetricDataRequest) GetAggreOverLineOps() *string {
	return s.AggreOverLineOps
}

func (s *DescribeMetricDataRequest) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeMetricDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMetricDataRequest) GetGroupByLabels() []*string {
	return s.GroupByLabels
}

func (s *DescribeMetricDataRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricDataRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeMetricDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMetricDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMetricDataRequest) SetAggreOps(v string) *DescribeMetricDataRequest {
	s.AggreOps = &v
	return s
}

func (s *DescribeMetricDataRequest) SetAggreOverLineOps(v string) *DescribeMetricDataRequest {
	s.AggreOverLineOps = &v
	return s
}

func (s *DescribeMetricDataRequest) SetDimensions(v string) *DescribeMetricDataRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricDataRequest) SetEndTime(v string) *DescribeMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricDataRequest) SetGroupByLabels(v []*string) *DescribeMetricDataRequest {
	s.GroupByLabels = v
	return s
}

func (s *DescribeMetricDataRequest) SetMetricName(v string) *DescribeMetricDataRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricDataRequest) SetPeriod(v int32) *DescribeMetricDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricDataRequest) SetRegionId(v string) *DescribeMetricDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMetricDataRequest) SetStartTime(v string) *DescribeMetricDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
