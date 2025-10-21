// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceConfigChangeLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceConfigChangeLogRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeDBInstanceConfigChangeLogRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeDBInstanceConfigChangeLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstanceConfigChangeLogRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDBInstanceConfigChangeLogRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeDBInstanceConfigChangeLogRequest
	GetStartTime() *string
}

type DescribeDBInstanceConfigChangeLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-uf6lkzf*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-01-01 10:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-01-01 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstanceConfigChangeLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConfigChangeLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigChangeLogRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceConfigChangeLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBInstanceConfigChangeLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstanceConfigChangeLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstanceConfigChangeLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceConfigChangeLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBInstanceConfigChangeLogRequest) SetDBInstanceId(v string) *DescribeDBInstanceConfigChangeLogRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogRequest) SetEndTime(v string) *DescribeDBInstanceConfigChangeLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogRequest) SetPageNumber(v int32) *DescribeDBInstanceConfigChangeLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogRequest) SetPageSize(v int32) *DescribeDBInstanceConfigChangeLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogRequest) SetRegionId(v string) *DescribeDBInstanceConfigChangeLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogRequest) SetStartTime(v string) *DescribeDBInstanceConfigChangeLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogRequest) Validate() error {
	return dara.Validate(s)
}
