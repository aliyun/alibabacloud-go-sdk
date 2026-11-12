// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseInstanceMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseInstanceId(v string) *DescribeDatabaseInstanceMetricDataRequest
	GetDatabaseInstanceId() *string
	SetEndTime(v string) *DescribeDatabaseInstanceMetricDataRequest
	GetEndTime() *string
	SetMetricName(v string) *DescribeDatabaseInstanceMetricDataRequest
	GetMetricName() *string
	SetRegionId(v string) *DescribeDatabaseInstanceMetricDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeDatabaseInstanceMetricDataRequest
	GetStartTime() *string
}

type DescribeDatabaseInstanceMetricDataRequest struct {
	// The ID of the lightweight database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
	//
	// > The time in the console is displayed in UTC+8.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-07T04:04Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the metric. Valid values:
	//
	// - `MySQL_MemCpuUsage`: The CPU utilization and memory usage of the instance, as a percentage of the total resources of the operating system.
	//
	// - `MySQL_DetailedSpaceUsage`: The total, data, log, temporary, and system disk space usage of the instance.
	//
	// - `MySQL_Sessions`: The number of active and total connections.
	//
	// - `MySQL_IOPS`: The IOPS of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL_MemCpuUsage
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The region ID of the lightweight database instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to get the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
	//
	// > The time in the console is displayed in UTC+8.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-06T04:04Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDatabaseInstanceMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseInstanceMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceMetricDataRequest) GetDatabaseInstanceId() *string {
	return s.DatabaseInstanceId
}

func (s *DescribeDatabaseInstanceMetricDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDatabaseInstanceMetricDataRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeDatabaseInstanceMetricDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDatabaseInstanceMetricDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDatabaseInstanceMetricDataRequest) SetDatabaseInstanceId(v string) *DescribeDatabaseInstanceMetricDataRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataRequest) SetEndTime(v string) *DescribeDatabaseInstanceMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataRequest) SetMetricName(v string) *DescribeDatabaseInstanceMetricDataRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataRequest) SetRegionId(v string) *DescribeDatabaseInstanceMetricDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataRequest) SetStartTime(v string) *DescribeDatabaseInstanceMetricDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
