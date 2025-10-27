// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputingGroupId(v string) *DescribeSlowLogTrendRequest
	GetComputingGroupId() *string
	SetDBInstanceId(v string) *DescribeSlowLogTrendRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeSlowLogTrendRequest
	GetEndTime() *string
	SetProduct(v string) *DescribeSlowLogTrendRequest
	GetProduct() *string
	SetQueryDurationMs(v string) *DescribeSlowLogTrendRequest
	GetQueryDurationMs() *string
	SetRegionId(v string) *DescribeSlowLogTrendRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeSlowLogTrendRequest
	GetStartTime() *string
}

type DescribeSlowLogTrendRequest struct {
	ComputingGroupId *string `json:"ComputingGroupId,omitempty" xml:"ComputingGroupId,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-dd hh:mm:ss format. The time must be in UTC.
	//
	// example:
	//
	// 2023-06-07 10:03:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// clickhouse
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The execution duration of slow SQL queries. Minimum value: **1000**. Unit: milliseconds.
	//
	// example:
	//
	// 3000
	QueryDurationMs *string `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start of the time range to query. Specify the time in the yyyy-MM-dd hh:mm:ss format. The time must be in UTC.
	//
	// example:
	//
	// 2023-04-13 17:48:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogTrendRequest) GetComputingGroupId() *string {
	return s.ComputingGroupId
}

func (s *DescribeSlowLogTrendRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSlowLogTrendRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSlowLogTrendRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeSlowLogTrendRequest) GetQueryDurationMs() *string {
	return s.QueryDurationMs
}

func (s *DescribeSlowLogTrendRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSlowLogTrendRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSlowLogTrendRequest) SetComputingGroupId(v string) *DescribeSlowLogTrendRequest {
	s.ComputingGroupId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetDBInstanceId(v string) *DescribeSlowLogTrendRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetEndTime(v string) *DescribeSlowLogTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetProduct(v string) *DescribeSlowLogTrendRequest {
	s.Product = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetQueryDurationMs(v string) *DescribeSlowLogTrendRequest {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetRegionId(v string) *DescribeSlowLogTrendRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) SetStartTime(v string) *DescribeSlowLogTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogTrendRequest) Validate() error {
	return dara.Validate(s)
}
