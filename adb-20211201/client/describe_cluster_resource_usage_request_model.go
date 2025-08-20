// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterResourceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeClusterResourceUsageRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeClusterResourceUsageRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeClusterResourceUsageRequest
	GetStartTime() *string
}

type DescribeClusterResourceUsageRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-08-22T01:06:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-11-29T10:20Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeClusterResourceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceUsageRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeClusterResourceUsageRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeClusterResourceUsageRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeClusterResourceUsageRequest) SetDBClusterId(v string) *DescribeClusterResourceUsageRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeClusterResourceUsageRequest) SetEndTime(v string) *DescribeClusterResourceUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeClusterResourceUsageRequest) SetStartTime(v string) *DescribeClusterResourceUsageRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeClusterResourceUsageRequest) Validate() error {
	return dara.Validate(s)
}
