// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSQLLogCountRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeSQLLogCountRequest
	GetDatabase() *string
	SetEndTime(v string) *DescribeSQLLogCountRequest
	GetEndTime() *string
	SetExecuteCost(v string) *DescribeSQLLogCountRequest
	GetExecuteCost() *string
	SetExecuteState(v string) *DescribeSQLLogCountRequest
	GetExecuteState() *string
	SetMaxExecuteCost(v string) *DescribeSQLLogCountRequest
	GetMaxExecuteCost() *string
	SetMinExecuteCost(v string) *DescribeSQLLogCountRequest
	GetMinExecuteCost() *string
	SetOperationClass(v string) *DescribeSQLLogCountRequest
	GetOperationClass() *string
	SetOperationType(v string) *DescribeSQLLogCountRequest
	GetOperationType() *string
	SetQueryKeywords(v string) *DescribeSQLLogCountRequest
	GetQueryKeywords() *string
	SetSourceIP(v string) *DescribeSQLLogCountRequest
	GetSourceIP() *string
	SetStartTime(v string) *DescribeSQLLogCountRequest
	GetStartTime() *string
	SetUser(v string) *DescribeSQLLogCountRequest
	GetUser() *string
}

type DescribeSQLLogCountRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
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
	// testdb
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// > The end time must be later than the start time. The maximum time range that can be specified is seven days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-12-14T11:22Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The execution duration of the SQL statement. Unit: seconds.
	//
	// example:
	//
	// 100
	ExecuteCost *string `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	// The execution status of the query. Valid values:
	//
	// 	- 1: successful.
	//
	// 	- 0: failed.
	//
	// 	- 0,1 or 1,0: all.
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
	// 10
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
	// 	- If you specify **OperationClass**, the value of **OperationType*	- must be of the corresponding query language. For example, if you set **OperationClass*	- to **DQL**, the value of **OperationType*	- must be a **DQL*	- statement such as **SELECT**.
	//
	// 	- If you leave **OperationClass*	- empty, the value of **OperationType*	- can be an SQL statement of any query language.
	//
	// 	- If you leave **OperationClass*	- and **OperationType*	- empty, all types of SQL statements are returned.
	//
	// example:
	//
	// SELECT
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The keywords that are used to query audit logs.
	//
	// example:
	//
	// test
	QueryKeywords *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 10.**.**.13
	SourceIP *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-12-12T11:22Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the database account that is used to connect to the database.
	//
	// example:
	//
	// adbpgadmin
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLLogCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSQLLogCountRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeSQLLogCountRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSQLLogCountRequest) GetExecuteCost() *string {
	return s.ExecuteCost
}

func (s *DescribeSQLLogCountRequest) GetExecuteState() *string {
	return s.ExecuteState
}

func (s *DescribeSQLLogCountRequest) GetMaxExecuteCost() *string {
	return s.MaxExecuteCost
}

func (s *DescribeSQLLogCountRequest) GetMinExecuteCost() *string {
	return s.MinExecuteCost
}

func (s *DescribeSQLLogCountRequest) GetOperationClass() *string {
	return s.OperationClass
}

func (s *DescribeSQLLogCountRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeSQLLogCountRequest) GetQueryKeywords() *string {
	return s.QueryKeywords
}

func (s *DescribeSQLLogCountRequest) GetSourceIP() *string {
	return s.SourceIP
}

func (s *DescribeSQLLogCountRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSQLLogCountRequest) GetUser() *string {
	return s.User
}

func (s *DescribeSQLLogCountRequest) SetDBInstanceId(v string) *DescribeSQLLogCountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetDatabase(v string) *DescribeSQLLogCountRequest {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetEndTime(v string) *DescribeSQLLogCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetExecuteCost(v string) *DescribeSQLLogCountRequest {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetExecuteState(v string) *DescribeSQLLogCountRequest {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetMaxExecuteCost(v string) *DescribeSQLLogCountRequest {
	s.MaxExecuteCost = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetMinExecuteCost(v string) *DescribeSQLLogCountRequest {
	s.MinExecuteCost = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetOperationClass(v string) *DescribeSQLLogCountRequest {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetOperationType(v string) *DescribeSQLLogCountRequest {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetQueryKeywords(v string) *DescribeSQLLogCountRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetSourceIP(v string) *DescribeSQLLogCountRequest {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetStartTime(v string) *DescribeSQLLogCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetUser(v string) *DescribeSQLLogCountRequest {
	s.User = &v
	return s
}

func (s *DescribeSQLLogCountRequest) Validate() error {
	return dara.Validate(s)
}
