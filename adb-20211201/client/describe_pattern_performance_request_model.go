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
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the details of all clusters in a region, including their cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-uf6li1r3do8m****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in UTC in the *yyyy-MM-ddTHH:mm:ssZ	- format.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2022-08-22T01:06:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the SQL pattern.
	//
	// > You can call the [DescribeSQLPatterns](https://help.aliyun.com/document_detail/321868.html) operation to query information about all SQL patterns in a cluster within a specified time range, including the ID of each SQL pattern.
	//
	// example:
	//
	// 3847585356974******
	PatternId *string `json:"PatternId,omitempty" xml:"PatternId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in UTC in the *yyyy-MM-ddTHH:mm:ssZ	- format.
	//
	// > - You can query data from the last 14 days. If you specify a start time earlier than this period, an empty value is returned. For example, if the current date is August 22, 2022 (China Standard Time), the earliest valid start time is 2022-08-08T16:00:00Z.
	//
	// - The interval between the start time and the end time cannot exceed 24 hours.
	//
	// example:
	//
	// 2022-08-21T02:15:00Z
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
