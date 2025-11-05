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
	// Aggregation method over time. Possible values include:
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
	// - SUM_OVER_TIME_LCRO: Sum over a left-closed, right-open interval
	//
	// - AVG_OVER_TIME_LCRO: Average over a left-closed, right-open interval
	//
	// - SUM_OVER_TIME_LORC: Sum over a left-open, right-closed interval
	//
	// - AVG_OVER_TIME_LORC: Average over a left-open, right-closed interval
	//
	// example:
	//
	// AVG_OVER_TIME
	AggreOps *string `json:"AggreOps,omitempty" xml:"AggreOps,omitempty"`
	// Aggregation method between lines. Possible values include:
	//
	// - NON: No aggregation
	//
	// - SUM: Sum
	//
	// - AVG: Average
	//
	// - COUNT: Count
	//
	// - MAX: Maximum
	//
	// - MIN: Minimum
	//
	// example:
	//
	// NON
	AggreOverLineOps *string `json:"AggreOverLineOps,omitempty" xml:"AggreOverLineOps,omitempty"`
	// The dimension map, in the JSON format. Valid values:
	//
	// 	- DiskId: the disk name. Example: d-xxx.
	//
	// 	- DeviceType: the disk type. system indicates the system disk, and data indicates the data disk.
	//
	// 	- DeviceCategory: the disk category. Example: cloud_essd.
	//
	// 	- EcsInstanceId: the ECS instance name. Example: i-xxx.
	//
	// 	- Azone: the zone, such as cn-hangzhou-a.
	//
	// The returned result is the intersection of all dimension filtering conditions.
	//
	// example:
	//
	// {"DiskId":["d-bp14xxxx","d-bp11xxxx"], "DeviceCategory": ["cloud_essd"]}
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The end time point for obtaining metric data. It should not be later than the current moment. Represented according to the ISO 8601 standard, using UTC +0 time, in the format yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2023-11-21T02:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of fields used for grouping and aggregation.
	GroupByLabels []*string `json:"GroupByLabels,omitempty" xml:"GroupByLabels,omitempty" type:"Repeated"`
	// Metric name. Possible values include:
	//
	//
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
	// The granularity at which data is collected for the metric. Unit: seconds. Default value: 5. Valid values:
	//
	// 	- 5: 5 seconds. The query time range can be up to 12 hours.
	//
	// 	- 10: 10 seconds. The query time range can be up to 24 hours.
	//
	// 	- 60: 60 seconds. The query time range can be up to 7 days.
	//
	// 	- 300: 300 seconds. The query time range can be up to 30 days.
	//
	// 	- 600: 600 seconds. The query time range can be up to 30 days.
	//
	// 	- 3600: 3,600 seconds. The query time range can be up to 30 days.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. You can specify a point in time that is up to 30 days before the current time. If both StartTime and EndTime are left empty, the monitoring metric data of the most recent statistical period is queried. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
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
