// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSlowSQLRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListSlowSQLRecordsRequest
	GetDBInstanceId() *string
	SetDBName(v string) *ListSlowSQLRecordsRequest
	GetDBName() *string
	SetEndTime(v string) *ListSlowSQLRecordsRequest
	GetEndTime() *string
	SetKeyword(v string) *ListSlowSQLRecordsRequest
	GetKeyword() *string
	SetMaxDuration(v string) *ListSlowSQLRecordsRequest
	GetMaxDuration() *string
	SetMinDuration(v string) *ListSlowSQLRecordsRequest
	GetMinDuration() *string
	SetOrderBy(v string) *ListSlowSQLRecordsRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *ListSlowSQLRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSlowSQLRecordsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListSlowSQLRecordsRequest
	GetRegionId() *string
	SetStartTime(v string) *ListSlowSQLRecordsRequest
	GetStartTime() *string
	SetUserName(v string) *ListSlowSQLRecordsRequest
	GetUserName() *string
}

type ListSlowSQLRecordsRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// testdb01
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. The time must be in UTC and formatted as *yyyy-MM-dd*T*HH:mm*Z. The time must be in UTC. The end time must be later than the start time.
	//
	// Defaults to the current time.
	//
	// example:
	//
	// 2024-12-04T17:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// SELECT
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The longest execution duration. Unit: seconds.
	//
	// example:
	//
	// 600
	MaxDuration *string `json:"MaxDuration,omitempty" xml:"MaxDuration,omitempty"`
	// The minimum execution duration. Unit: seconds.
	//
	// example:
	//
	// 10
	MinDuration *string `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	// Sort criteria (JSON string). {"Field":"QueryStartTime","Type":"Desc"}
	//
	// Field: Specifies the field to sort by.
	//
	// 	- QueryID: Query ID.
	//
	// 	- UserName: The username.
	//
	// 	- DBName: The name of the database.
	//
	// 	- QueryStartTime: The start time.
	//
	// 	- QueryEndTime: The end time.
	//
	// 	- DurationTime: The execution duration of the query.
	//
	// 	- Optimizer.
	//
	// 	- LockWaitTime: The lock waiting time.
	//
	// 	- QueueWaitTime: Time in Queue.
	//
	// 	- CpuTimeMs: CPU time.
	//
	// 	- MemBytes: The memory usage.
	//
	// 	- SpillBytes: Spill File Size.
	//
	// Type: Sort order.
	//
	// 	- Desc: descending.
	//
	// 	- Asc: ascending.
	//
	// example:
	//
	// {"Field":"QueryStartTime","Type":"Desc"}
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of the page to return. The value must be a positive integer in the range [1, 100]. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. The time must be in UTC and formatted as yyyy-MM-ddTHH:mmZ.
	//
	// Defaults to the current time minus 5 minutes.
	//
	// example:
	//
	// 2024-12-04T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The account name.
	//
	// example:
	//
	// test_user
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListSlowSQLRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSlowSQLRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListSlowSQLRecordsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListSlowSQLRecordsRequest) GetDBName() *string {
	return s.DBName
}

func (s *ListSlowSQLRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListSlowSQLRecordsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListSlowSQLRecordsRequest) GetMaxDuration() *string {
	return s.MaxDuration
}

func (s *ListSlowSQLRecordsRequest) GetMinDuration() *string {
	return s.MinDuration
}

func (s *ListSlowSQLRecordsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListSlowSQLRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSlowSQLRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSlowSQLRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSlowSQLRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListSlowSQLRecordsRequest) GetUserName() *string {
	return s.UserName
}

func (s *ListSlowSQLRecordsRequest) SetDBInstanceId(v string) *ListSlowSQLRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListSlowSQLRecordsRequest) SetDBName(v string) *ListSlowSQLRecordsRequest {
	s.DBName = &v
	return s
}

func (s *ListSlowSQLRecordsRequest) SetEndTime(v string) *ListSlowSQLRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *ListSlowSQLRecordsRequest) SetKeyword(v string) *ListSlowSQLRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *ListSlowSQLRecordsRequest) SetMaxDuration(v string) *ListSlowSQLRecordsRequest {
	s.MaxDuration = &v
	return s
}

func (s *ListSlowSQLRecordsRequest) SetMinDuration(v string) *ListSlowSQLRecordsRequest {
	s.MinDuration = &v
	return s
}

func (s *ListSlowSQLRecordsRequest) SetOrderBy(v string) *ListSlowSQLRecordsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListSlowSQLRecordsRequest) SetPageNumber(v int32) *ListSlowSQLRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSlowSQLRecordsRequest) SetPageSize(v int32) *ListSlowSQLRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSlowSQLRecordsRequest) SetRegionId(v string) *ListSlowSQLRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSlowSQLRecordsRequest) SetStartTime(v string) *ListSlowSQLRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *ListSlowSQLRecordsRequest) SetUserName(v string) *ListSlowSQLRecordsRequest {
	s.UserName = &v
	return s
}

func (s *ListSlowSQLRecordsRequest) Validate() error {
	return dara.Validate(s)
}
