// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceErrorLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceErrorLogRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeDBInstanceErrorLogRequest
	GetDatabase() *string
	SetEndTime(v string) *DescribeDBInstanceErrorLogRequest
	GetEndTime() *string
	SetHost(v string) *DescribeDBInstanceErrorLogRequest
	GetHost() *string
	SetKeywords(v string) *DescribeDBInstanceErrorLogRequest
	GetKeywords() *string
	SetLogLevel(v string) *DescribeDBInstanceErrorLogRequest
	GetLogLevel() *string
	SetPageNumber(v int32) *DescribeDBInstanceErrorLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstanceErrorLogRequest
	GetPageSize() *int32
	SetStartTime(v string) *DescribeDBInstanceErrorLogRequest
	GetStartTime() *string
	SetUser(v string) *DescribeDBInstanceErrorLogRequest
	GetUser() *string
}

type DescribeDBInstanceErrorLogRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including the instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database name.
	//
	// example:
	//
	// adbtest
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end time of the query. The end time must be later than the start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm</i>Z format. The time must be in UTC.
	//
	// example:
	//
	// 2022-04-25T06:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is not supported in the public cloud. You do not need to specify this parameter.
	//
	// example:
	//
	// null
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The keyword used to search for error logs.
	//
	// example:
	//
	// error
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// The log level. Valid values:
	//
	// - **ALL**: displays logs of all levels.
	//
	// - **PANIC**: displays only panic-level logs.
	//
	// - **FATAL**: displays only fatal-level logs.
	//
	// - **ERROR**: displays only error-level logs.
	//
	// example:
	//
	// ALL
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// The page number. The value must be greater than 0 and cannot exceed the maximum value of the Integer type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **20**
	//
	// - **50**
	//
	// - **100**
	//
	// Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time of the query. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm</i>Z format. The time must be in UTC.
	//
	// example:
	//
	// 2022-04-24T06:59Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The username.
	//
	// example:
	//
	// adbpguser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDBInstanceErrorLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceErrorLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceErrorLogRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceErrorLogRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeDBInstanceErrorLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBInstanceErrorLogRequest) GetHost() *string {
	return s.Host
}

func (s *DescribeDBInstanceErrorLogRequest) GetKeywords() *string {
	return s.Keywords
}

func (s *DescribeDBInstanceErrorLogRequest) GetLogLevel() *string {
	return s.LogLevel
}

func (s *DescribeDBInstanceErrorLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstanceErrorLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstanceErrorLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBInstanceErrorLogRequest) GetUser() *string {
	return s.User
}

func (s *DescribeDBInstanceErrorLogRequest) SetDBInstanceId(v string) *DescribeDBInstanceErrorLogRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetDatabase(v string) *DescribeDBInstanceErrorLogRequest {
	s.Database = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetEndTime(v string) *DescribeDBInstanceErrorLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetHost(v string) *DescribeDBInstanceErrorLogRequest {
	s.Host = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetKeywords(v string) *DescribeDBInstanceErrorLogRequest {
	s.Keywords = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetLogLevel(v string) *DescribeDBInstanceErrorLogRequest {
	s.LogLevel = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetPageNumber(v int32) *DescribeDBInstanceErrorLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetPageSize(v int32) *DescribeDBInstanceErrorLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetStartTime(v string) *DescribeDBInstanceErrorLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) SetUser(v string) *DescribeDBInstanceErrorLogRequest {
	s.User = &v
	return s
}

func (s *DescribeDBInstanceErrorLogRequest) Validate() error {
	return dara.Validate(s)
}
