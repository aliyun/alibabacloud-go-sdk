// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobResourceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeJobResourceUsageRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeJobResourceUsageRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeJobResourceUsageRequest
	GetStartTime() *string
}

type DescribeJobResourceUsageRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6g8w25jacm7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-03-17T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-04T03:45:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeJobResourceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobResourceUsageRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeJobResourceUsageRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeJobResourceUsageRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeJobResourceUsageRequest) SetDBClusterId(v string) *DescribeJobResourceUsageRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeJobResourceUsageRequest) SetEndTime(v string) *DescribeJobResourceUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeJobResourceUsageRequest) SetStartTime(v string) *DescribeJobResourceUsageRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeJobResourceUsageRequest) Validate() error {
	return dara.Validate(s)
}
