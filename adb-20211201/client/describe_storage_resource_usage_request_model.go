// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageResourceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeStorageResourceUsageRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeStorageResourceUsageRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeStorageResourceUsageRequest
	GetStartTime() *string
}

type DescribeStorageResourceUsageRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp10yt0gva71ei7d
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-08-23T01:06:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-08-22T01:06:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeStorageResourceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeStorageResourceUsageRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeStorageResourceUsageRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeStorageResourceUsageRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeStorageResourceUsageRequest) SetDBClusterId(v string) *DescribeStorageResourceUsageRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeStorageResourceUsageRequest) SetEndTime(v string) *DescribeStorageResourceUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeStorageResourceUsageRequest) SetStartTime(v string) *DescribeStorageResourceUsageRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeStorageResourceUsageRequest) Validate() error {
	return dara.Validate(s)
}
