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
	// >You can call the [describedbclusters](https://help.aliyun.com/document_detail/98094.html) operation to view the details of all clusters under your account, such as the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The end time must be later than the start time. The format is YYYY-MM-DDThh:mmZ (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-09-23T01:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies the time granularity of the monitoring data to retrieve. Valid values: **60*	- (minutes) and **3600*	- (hours).
	//
	// - **Interval**=**60**: You can query data for the last month at most. The maximum time span for a single query is 7 days.
	//
	// - **Interval**=**3600**: You can query data for the last month at most. The maximum time span for a single query is 7 days.
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the monitoring metric.
	//
	// This parameter is required.
	//
	// example:
	//
	// PolarDBAIModelCall
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The start time of the query. The format is YYYY-MM-DDThh:mmZ (UTC).
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
