// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSwitchLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceSwitchLogRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeDBInstanceSwitchLogRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeDBInstanceSwitchLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstanceSwitchLogRequest
	GetPageSize() *int32
	SetResourceOwnerId(v int64) *DescribeDBInstanceSwitchLogRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeDBInstanceSwitchLogRequest
	GetStartTime() *string
}

type DescribeDBInstanceSwitchLogRequest struct {
	// The instance ID.
	//
	// example:
	//
	// dds-uf68f1b5a57exxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-mm-dd*T*hh:mm*Z format. The time must be in UTC. The end time must be later than the start time.
	//
	// example:
	//
	// 2023-05-28T02:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0 and less than or equal to the maximum value supported by the integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **30, 50, and 100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize        *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-mm-dd*T*hh:mm*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2023-05-27T02:46Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstanceSwitchLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSwitchLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSwitchLogRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceSwitchLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBInstanceSwitchLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstanceSwitchLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstanceSwitchLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceSwitchLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBInstanceSwitchLogRequest) SetDBInstanceId(v string) *DescribeDBInstanceSwitchLogRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) SetEndTime(v string) *DescribeDBInstanceSwitchLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) SetPageNumber(v int32) *DescribeDBInstanceSwitchLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) SetPageSize(v int32) *DescribeDBInstanceSwitchLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceSwitchLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) SetStartTime(v string) *DescribeDBInstanceSwitchLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstanceSwitchLogRequest) Validate() error {
	return dara.Validate(s)
}
