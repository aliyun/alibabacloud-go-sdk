// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePatternPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribePatternPerformanceRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribePatternPerformanceRequest
	GetEndTime() *string
	SetPatternId(v string) *DescribePatternPerformanceRequest
	GetPatternId() *string
	SetRegionId(v string) *DescribePatternPerformanceRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribePatternPerformanceRequest
	GetStartTime() *string
}

type DescribePatternPerformanceRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-18T18:05:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The SQL pattern ID.
	//
	// > You can call the [DescribeSQLPatterns](https://help.aliyun.com/document_detail/321868.html) operation to query the information about all SQL patterns in an AnalyticDB for MySQL cluster within a period of time, including SQL pattern IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3847585356974******
	PatternId *string `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// >
	//
	// 	- Only data within the last 14 days can be queried. For example, if the current date is November 22 (UTC+8), you can query data on a day as early as November 9 by setting StartTime to 2021-11-08T16:00:00Z. If you set StartTime to a value earlier than 2021-11-08T16:00:00Z, the Performances parameter is empty.
	//
	// 	- The maximum time range that can be specified is 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-18T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePatternPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePatternPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribePatternPerformanceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribePatternPerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePatternPerformanceRequest) GetPatternId() *string {
	return s.PatternId
}

func (s *DescribePatternPerformanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePatternPerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePatternPerformanceRequest) SetDBClusterId(v string) *DescribePatternPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetEndTime(v string) *DescribePatternPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetPatternId(v string) *DescribePatternPerformanceRequest {
	s.PatternId = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetRegionId(v string) *DescribePatternPerformanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePatternPerformanceRequest) SetStartTime(v string) *DescribePatternPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePatternPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
