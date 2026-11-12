// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseErrorLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseInstanceId(v string) *DescribeDatabaseErrorLogsRequest
	GetDatabaseInstanceId() *string
	SetEndTime(v string) *DescribeDatabaseErrorLogsRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeDatabaseErrorLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDatabaseErrorLogsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDatabaseErrorLogsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeDatabaseErrorLogsRequest
	GetStartTime() *string
}

type DescribeDatabaseErrorLogsRequest struct {
	// The ID of the lightweight database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The end of the time range to query. The time must be in UTC and formatted as `yyyy-MM-ddTHH:mmZ` according to the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard. This value must be later than the start time.
	//
	// > The time displayed in the console is 8 hours ahead of the specified UTC time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-08T04:04Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of error logs to return per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the lightweight database instance resides.
	//
	// You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start of the time range to query. The time must be in UTC and formatted as `yyyy-MM-ddTHH:mmZ` according to the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard.
	//
	// > The time displayed in the console is 8 hours ahead of the specified UTC time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-07T04:04Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDatabaseErrorLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseErrorLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseErrorLogsRequest) GetDatabaseInstanceId() *string {
	return s.DatabaseInstanceId
}

func (s *DescribeDatabaseErrorLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDatabaseErrorLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDatabaseErrorLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDatabaseErrorLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDatabaseErrorLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDatabaseErrorLogsRequest) SetDatabaseInstanceId(v string) *DescribeDatabaseErrorLogsRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) SetEndTime(v string) *DescribeDatabaseErrorLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) SetPageNumber(v int32) *DescribeDatabaseErrorLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) SetPageSize(v int32) *DescribeDatabaseErrorLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) SetRegionId(v string) *DescribeDatabaseErrorLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) SetStartTime(v string) *DescribeDatabaseErrorLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) Validate() error {
	return dara.Validate(s)
}
