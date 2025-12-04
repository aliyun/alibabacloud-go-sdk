// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeConfigHistoryRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeConfigHistoryRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeConfigHistoryRequest
	GetStartTime() *string
}

type DescribeConfigHistoryRequest struct {
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all the clusters that are deployed in a specific region. The information includes the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1p816075e21****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-08-22T10:06:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in Coordinated Universal Time (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-08-11T06:27:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeConfigHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigHistoryRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeConfigHistoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeConfigHistoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeConfigHistoryRequest) SetDBClusterId(v string) *DescribeConfigHistoryRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeConfigHistoryRequest) SetEndTime(v string) *DescribeConfigHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeConfigHistoryRequest) SetStartTime(v string) *DescribeConfigHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeConfigHistoryRequest) Validate() error {
	return dara.Validate(s)
}
