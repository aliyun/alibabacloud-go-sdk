// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSQLLogsRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeSQLLogsRequest
	GetDatabase() *string
	SetEndTime(v string) *DescribeSQLLogsRequest
	GetEndTime() *string
	SetExecuteCost(v string) *DescribeSQLLogsRequest
	GetExecuteCost() *string
	SetExecuteState(v string) *DescribeSQLLogsRequest
	GetExecuteState() *string
	SetMaxExecuteCost(v string) *DescribeSQLLogsRequest
	GetMaxExecuteCost() *string
	SetMinExecuteCost(v string) *DescribeSQLLogsRequest
	GetMinExecuteCost() *string
	SetOperationClass(v string) *DescribeSQLLogsRequest
	GetOperationClass() *string
	SetOperationType(v string) *DescribeSQLLogsRequest
	GetOperationType() *string
	SetPageNumber(v int32) *DescribeSQLLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSQLLogsRequest
	GetPageSize() *int32
	SetQueryKeywords(v string) *DescribeSQLLogsRequest
	GetQueryKeywords() *string
	SetSourceIP(v string) *DescribeSQLLogsRequest
	GetSourceIP() *string
	SetStartTime(v string) *DescribeSQLLogsRequest
	GetStartTime() *string
	SetUser(v string) *DescribeSQLLogsRequest
	GetUser() *string
}

type DescribeSQLLogsRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adbpgadmin
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC.
	//
	// > The end time must be later than the start time. The maximum time range that can be specified is seven days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-03-17T06:30Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The execution duration of the SQL statement. Unit: seconds.
	//
	// example:
	//
	// 1
	ExecuteCost *string `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	// The execution status of the SQL statement. Valid values:
	//
	// 	- **1**: successful.
	//
	// 	- **0**: failed.
	//
	// example:
	//
	// success
	ExecuteState *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	// The maximum amount of time consumed by a slow query. Unit: seconds. Minimum value: 0.
	//
	// example:
	//
	// 1000
	MaxExecuteCost *string `json:"MaxExecuteCost,omitempty" xml:"MaxExecuteCost,omitempty"`
	// The minimum amount of time consumed by a slow query. Unit: seconds. Minimum value: 0.
	//
	// example:
	//
	// 1
	MinExecuteCost *string `json:"MinExecuteCost,omitempty" xml:"MinExecuteCost,omitempty"`
	// The type of the query language. Valid values:
	//
	// 	- **DQL**
	//
	// 	- **DML**
	//
	// 	- **DDL**
	//
	// 	- **DCL**
	//
	// 	- **TCL**
	//
	// example:
	//
	// DQL
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	// The type of the SQL statement.
	//
	// >
	//
	// 	- If **OperationClass*	- is specified, the value of **OperationType*	- must belong to the corresponding query language. For example, if **OperationClass*	- is set to **DQL**, the value of **OperationType*	- must be a **DQL*	- statement such as **SELECT**.
	//
	// 	- If **OperationClass*	- is not specified, the value of **OperationType*	- can be an SQL statement of any query language.
	//
	// 	- If **OperationClass*	- and **OperationType*	- are not specified, all types of SQL statements are returned.
	//
	// example:
	//
	// SELECT
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
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
	// The keywords of the SQL statement.
	//
	// example:
	//
	// select 1
	QueryKeywords *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 100.**.**.90
	SourceIP *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-03-10T06:30Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// testadmin
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSQLLogsRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeSQLLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSQLLogsRequest) GetExecuteCost() *string {
	return s.ExecuteCost
}

func (s *DescribeSQLLogsRequest) GetExecuteState() *string {
	return s.ExecuteState
}

func (s *DescribeSQLLogsRequest) GetMaxExecuteCost() *string {
	return s.MaxExecuteCost
}

func (s *DescribeSQLLogsRequest) GetMinExecuteCost() *string {
	return s.MinExecuteCost
}

func (s *DescribeSQLLogsRequest) GetOperationClass() *string {
	return s.OperationClass
}

func (s *DescribeSQLLogsRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeSQLLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSQLLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSQLLogsRequest) GetQueryKeywords() *string {
	return s.QueryKeywords
}

func (s *DescribeSQLLogsRequest) GetSourceIP() *string {
	return s.SourceIP
}

func (s *DescribeSQLLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSQLLogsRequest) GetUser() *string {
	return s.User
}

func (s *DescribeSQLLogsRequest) SetDBInstanceId(v string) *DescribeSQLLogsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetDatabase(v string) *DescribeSQLLogsRequest {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetEndTime(v string) *DescribeSQLLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetExecuteCost(v string) *DescribeSQLLogsRequest {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetExecuteState(v string) *DescribeSQLLogsRequest {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetMaxExecuteCost(v string) *DescribeSQLLogsRequest {
	s.MaxExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetMinExecuteCost(v string) *DescribeSQLLogsRequest {
	s.MinExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetOperationClass(v string) *DescribeSQLLogsRequest {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetOperationType(v string) *DescribeSQLLogsRequest {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetPageNumber(v int32) *DescribeSQLLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetPageSize(v int32) *DescribeSQLLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetQueryKeywords(v string) *DescribeSQLLogsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetSourceIP(v string) *DescribeSQLLogsRequest {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetStartTime(v string) *DescribeSQLLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetUser(v string) *DescribeSQLLogsRequest {
	s.User = &v
	return s
}

func (s *DescribeSQLLogsRequest) Validate() error {
	return dara.Validate(s)
}
