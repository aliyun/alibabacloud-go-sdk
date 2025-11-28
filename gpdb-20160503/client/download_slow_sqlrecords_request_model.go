// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSlowSQLRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DownloadSlowSQLRecordsRequest
	GetDBInstanceId() *string
	SetDBName(v string) *DownloadSlowSQLRecordsRequest
	GetDBName() *string
	SetEndTime(v string) *DownloadSlowSQLRecordsRequest
	GetEndTime() *string
	SetKeyword(v string) *DownloadSlowSQLRecordsRequest
	GetKeyword() *string
	SetMaxDuration(v string) *DownloadSlowSQLRecordsRequest
	GetMaxDuration() *string
	SetMinDuration(v string) *DownloadSlowSQLRecordsRequest
	GetMinDuration() *string
	SetOrderBy(v string) *DownloadSlowSQLRecordsRequest
	GetOrderBy() *string
	SetRegionId(v string) *DownloadSlowSQLRecordsRequest
	GetRegionId() *string
	SetStartTime(v string) *DownloadSlowSQLRecordsRequest
	GetStartTime() *string
	SetUserName(v string) *DownloadSlowSQLRecordsRequest
	GetUserName() *string
}

type DownloadSlowSQLRecordsRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database name.
	//
	// example:
	//
	// testdb01
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. The time must be in UTC and formatted as *yyyy-MM-dd*T*HH:mm*Z. The time must be in UTC. The end time must be later than the start time.
	//
	// Defaults to the current time
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
	// Sort criteria.
	//
	// example:
	//
	// {Field: SchemaName, Type: Desc}
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. The time must be in UTC and formatted as *yyyy-MM-dd*T*HH:mm*Z.
	//
	// Defaults to the current time minus 5 minutes.
	//
	// example:
	//
	// 2024-12-04T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The database account.
	//
	// example:
	//
	// test_user
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DownloadSlowSQLRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadSlowSQLRecordsRequest) GoString() string {
	return s.String()
}

func (s *DownloadSlowSQLRecordsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DownloadSlowSQLRecordsRequest) GetDBName() *string {
	return s.DBName
}

func (s *DownloadSlowSQLRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DownloadSlowSQLRecordsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DownloadSlowSQLRecordsRequest) GetMaxDuration() *string {
	return s.MaxDuration
}

func (s *DownloadSlowSQLRecordsRequest) GetMinDuration() *string {
	return s.MinDuration
}

func (s *DownloadSlowSQLRecordsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DownloadSlowSQLRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DownloadSlowSQLRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DownloadSlowSQLRecordsRequest) GetUserName() *string {
	return s.UserName
}

func (s *DownloadSlowSQLRecordsRequest) SetDBInstanceId(v string) *DownloadSlowSQLRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DownloadSlowSQLRecordsRequest) SetDBName(v string) *DownloadSlowSQLRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DownloadSlowSQLRecordsRequest) SetEndTime(v string) *DownloadSlowSQLRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DownloadSlowSQLRecordsRequest) SetKeyword(v string) *DownloadSlowSQLRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DownloadSlowSQLRecordsRequest) SetMaxDuration(v string) *DownloadSlowSQLRecordsRequest {
	s.MaxDuration = &v
	return s
}

func (s *DownloadSlowSQLRecordsRequest) SetMinDuration(v string) *DownloadSlowSQLRecordsRequest {
	s.MinDuration = &v
	return s
}

func (s *DownloadSlowSQLRecordsRequest) SetOrderBy(v string) *DownloadSlowSQLRecordsRequest {
	s.OrderBy = &v
	return s
}

func (s *DownloadSlowSQLRecordsRequest) SetRegionId(v string) *DownloadSlowSQLRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DownloadSlowSQLRecordsRequest) SetStartTime(v string) *DownloadSlowSQLRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DownloadSlowSQLRecordsRequest) SetUserName(v string) *DownloadSlowSQLRecordsRequest {
	s.UserName = &v
	return s
}

func (s *DownloadSlowSQLRecordsRequest) Validate() error {
	return dara.Validate(s)
}
