// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseSlowLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseInstanceId(v string) *DescribeDatabaseSlowLogRecordsRequest
	GetDatabaseInstanceId() *string
	SetEndTime(v string) *DescribeDatabaseSlowLogRecordsRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeDatabaseSlowLogRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDatabaseSlowLogRecordsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDatabaseSlowLogRecordsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeDatabaseSlowLogRecordsRequest
	GetStartTime() *string
}

type DescribeDatabaseSlowLogRecordsRequest struct {
	// The ID of the lightweight database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. The interval between the start time and the end time cannot be longer than 7 days.
	//
	// The time is in UTC and must be formatted as `yyyy-MM-ddTHH:mm:ssZ` as specified by the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard.
	//
	// > The console displays this time in the UTC+8 time zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-08T04:04:44Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// Starting value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 30 to 100.
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the lightweight database instance.
	//
	// You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to get the latest list of supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query.
	//
	// The time is in UTC and must be formatted as `yyyy-MM-ddTHH:mm:ssZ` as specified by the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard.
	//
	// > The console displays this time in the UTC+8 time zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-07T04:04:44Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDatabaseSlowLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseSlowLogRecordsRequest) GetDatabaseInstanceId() *string {
	return s.DatabaseInstanceId
}

func (s *DescribeDatabaseSlowLogRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDatabaseSlowLogRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDatabaseSlowLogRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDatabaseSlowLogRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDatabaseSlowLogRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetDatabaseInstanceId(v string) *DescribeDatabaseSlowLogRecordsRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetEndTime(v string) *DescribeDatabaseSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetPageNumber(v int32) *DescribeDatabaseSlowLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetPageSize(v int32) *DescribeDatabaseSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetRegionId(v string) *DescribeDatabaseSlowLogRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetStartTime(v string) *DescribeDatabaseSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) Validate() error {
	return dara.Validate(s)
}
