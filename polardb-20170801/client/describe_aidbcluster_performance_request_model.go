// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAIDBClusterPerformanceRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeAIDBClusterPerformanceRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeAIDBClusterPerformanceRequest
	GetInterval() *string
	SetKey(v string) *DescribeAIDBClusterPerformanceRequest
	GetKey() *string
	SetStartTime(v string) *DescribeAIDBClusterPerformanceRequest
	GetStartTime() *string
}

type DescribeAIDBClusterPerformanceRequest struct {
	// The cluster ID.
	//
	// > Call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of all clusters under your account, including the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the `YYYY-MM-DDThh:mmZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-09-23T01:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the data to retrieve. Valid values: **60*	- (minutes) and **3600*	- (hours).
	//
	// - If you set **Interval*	- to **60**, you can query data from the last month. The maximum time range for a single query is 7 days.
	//
	// - If you set **Interval*	- to **3600**, you can query data from the last month. The maximum time range for a single query is 7 days.
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the metric.
	//
	// This parameter is required.
	//
	// example:
	//
	// PolarDBAIModelCall
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The beginning of the time range to query. Specify the time in the `YYYY-MM-DDThh:mmZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-09-17T02:18:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAIDBClusterPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterPerformanceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAIDBClusterPerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAIDBClusterPerformanceRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeAIDBClusterPerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeAIDBClusterPerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAIDBClusterPerformanceRequest) SetDBClusterId(v string) *DescribeAIDBClusterPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceRequest) SetEndTime(v string) *DescribeAIDBClusterPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceRequest) SetInterval(v string) *DescribeAIDBClusterPerformanceRequest {
	s.Interval = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceRequest) SetKey(v string) *DescribeAIDBClusterPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceRequest) SetStartTime(v string) *DescribeAIDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAIDBClusterPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
