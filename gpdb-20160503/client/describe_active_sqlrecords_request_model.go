// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveSQLRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeActiveSQLRecordsRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeActiveSQLRecordsRequest
	GetDatabase() *string
	SetEndTime(v string) *DescribeActiveSQLRecordsRequest
	GetEndTime() *string
	SetKeyword(v string) *DescribeActiveSQLRecordsRequest
	GetKeyword() *string
	SetMaxDuration(v string) *DescribeActiveSQLRecordsRequest
	GetMaxDuration() *string
	SetMinDuration(v string) *DescribeActiveSQLRecordsRequest
	GetMinDuration() *string
	SetOrder(v string) *DescribeActiveSQLRecordsRequest
	GetOrder() *string
	SetStartTime(v string) *DescribeActiveSQLRecordsRequest
	GetStartTime() *string
	SetUser(v string) *DescribeActiveSQLRecordsRequest
	GetUser() *string
}

type DescribeActiveSQLRecordsRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// testdb
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The end time must be later than the start time.
	//
	// example:
	//
	// 2022-05-07T07:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword used to filter queries.
	//
	// example:
	//
	// SELECT
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The maxmum amount of time consumed by traces. Unit: milliseconds.
	//
	// example:
	//
	// 600
	MaxDuration *string `json:"MaxDuration,omitempty" xml:"MaxDuration,omitempty"`
	// The minimum amount of time consumed by traces. Unit: milliseconds.
	//
	// example:
	//
	// 300
	MinDuration *string `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	// The field used to sort lock diagnostics records and the sorting order.
	//
	// Default value: `{"Field":"StartTime","Type":"Desc"}`, which indicates that lock diagnostics records are sorted by the start time in descending order. No other values are supported.
	//
	// example:
	//
	// {"Field":"StartTime","Type":"Desc"}
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-08-03T09:30Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// testuser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeActiveSQLRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveSQLRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveSQLRecordsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeActiveSQLRecordsRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeActiveSQLRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeActiveSQLRecordsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeActiveSQLRecordsRequest) GetMaxDuration() *string {
	return s.MaxDuration
}

func (s *DescribeActiveSQLRecordsRequest) GetMinDuration() *string {
	return s.MinDuration
}

func (s *DescribeActiveSQLRecordsRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeActiveSQLRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeActiveSQLRecordsRequest) GetUser() *string {
	return s.User
}

func (s *DescribeActiveSQLRecordsRequest) SetDBInstanceId(v string) *DescribeActiveSQLRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeActiveSQLRecordsRequest) SetDatabase(v string) *DescribeActiveSQLRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeActiveSQLRecordsRequest) SetEndTime(v string) *DescribeActiveSQLRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeActiveSQLRecordsRequest) SetKeyword(v string) *DescribeActiveSQLRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeActiveSQLRecordsRequest) SetMaxDuration(v string) *DescribeActiveSQLRecordsRequest {
	s.MaxDuration = &v
	return s
}

func (s *DescribeActiveSQLRecordsRequest) SetMinDuration(v string) *DescribeActiveSQLRecordsRequest {
	s.MinDuration = &v
	return s
}

func (s *DescribeActiveSQLRecordsRequest) SetOrder(v string) *DescribeActiveSQLRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeActiveSQLRecordsRequest) SetStartTime(v string) *DescribeActiveSQLRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeActiveSQLRecordsRequest) SetUser(v string) *DescribeActiveSQLRecordsRequest {
	s.User = &v
	return s
}

func (s *DescribeActiveSQLRecordsRequest) Validate() error {
	return dara.Validate(s)
}
