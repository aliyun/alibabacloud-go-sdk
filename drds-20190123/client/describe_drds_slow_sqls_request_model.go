// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsSlowSqlsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeDrdsSlowSqlsRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeDrdsSlowSqlsRequest
	GetDrdsInstanceId() *string
	SetEndTime(v int64) *DescribeDrdsSlowSqlsRequest
	GetEndTime() *int64
	SetExeTime(v int64) *DescribeDrdsSlowSqlsRequest
	GetExeTime() *int64
	SetPageNumber(v int32) *DescribeDrdsSlowSqlsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDrdsSlowSqlsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDrdsSlowSqlsRequest
	GetRegionId() *string
	SetStartTime(v int64) *DescribeDrdsSlowSqlsRequest
	GetStartTime() *int64
}

type DescribeDrdsSlowSqlsRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_db
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds***************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The start time of the SQL query. Specify the time in the UNIX timestamp format. The time must be in UTC. Unit: ms.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1568267711000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The SQL execution time. Unit: ms.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	ExeTime *int64 `json:"ExeTime,omitempty" xml:"ExeTime,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The end time of the SQL query. Specify the time in the UNIX timestamp format. The time must be in UTC. Unit: ms.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1568269711000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDrdsSlowSqlsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsSlowSqlsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsSlowSqlsRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDrdsSlowSqlsRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsSlowSqlsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDrdsSlowSqlsRequest) GetExeTime() *int64 {
	return s.ExeTime
}

func (s *DescribeDrdsSlowSqlsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDrdsSlowSqlsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDrdsSlowSqlsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDrdsSlowSqlsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDrdsSlowSqlsRequest) SetDbName(v string) *DescribeDrdsSlowSqlsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetDrdsInstanceId(v string) *DescribeDrdsSlowSqlsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetEndTime(v int64) *DescribeDrdsSlowSqlsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetExeTime(v int64) *DescribeDrdsSlowSqlsRequest {
	s.ExeTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetPageNumber(v int32) *DescribeDrdsSlowSqlsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetPageSize(v int32) *DescribeDrdsSlowSqlsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetRegionId(v string) *DescribeDrdsSlowSqlsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) SetStartTime(v int64) *DescribeDrdsSlowSqlsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDrdsSlowSqlsRequest) Validate() error {
	return dara.Validate(s)
}
