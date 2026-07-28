// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowQueryStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSlowQueryStatsRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeSlowQueryStatsRequest
	GetEndTime() *string
	SetRegionId(v string) *DescribeSlowQueryStatsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeSlowQueryStatsRequest
	GetStartTime() *string
	SetThresholdMs(v int64) *DescribeSlowQueryStatsRequest
	GetThresholdMs() *int64
	SetTopN(v int32) *DescribeSlowQueryStatsRequest
	GetTopN() *int32
}

type DescribeSlowQueryStatsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end time. Must be later than the start time. Defaults to the current time.
	//
	// example:
	//
	// 2026-04-08 16:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. Defaults to 24 hours before the current time.
	//
	// example:
	//
	// 2026-04-07 16:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The slow query threshold, in milliseconds. The default value is 5000.
	//
	// example:
	//
	// 5000
	ThresholdMs *int64 `json:"ThresholdMs,omitempty" xml:"ThresholdMs,omitempty"`
	// The number of top slow queries to return. The default value is 10.
	//
	// example:
	//
	// 10
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
}

func (s DescribeSlowQueryStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowQueryStatsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowQueryStatsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSlowQueryStatsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSlowQueryStatsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSlowQueryStatsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSlowQueryStatsRequest) GetThresholdMs() *int64 {
	return s.ThresholdMs
}

func (s *DescribeSlowQueryStatsRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *DescribeSlowQueryStatsRequest) SetDBInstanceId(v string) *DescribeSlowQueryStatsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlowQueryStatsRequest) SetEndTime(v string) *DescribeSlowQueryStatsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowQueryStatsRequest) SetRegionId(v string) *DescribeSlowQueryStatsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlowQueryStatsRequest) SetStartTime(v string) *DescribeSlowQueryStatsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowQueryStatsRequest) SetThresholdMs(v int64) *DescribeSlowQueryStatsRequest {
	s.ThresholdMs = &v
	return s
}

func (s *DescribeSlowQueryStatsRequest) SetTopN(v int32) *DescribeSlowQueryStatsRequest {
	s.TopN = &v
	return s
}

func (s *DescribeSlowQueryStatsRequest) Validate() error {
	return dara.Validate(s)
}
