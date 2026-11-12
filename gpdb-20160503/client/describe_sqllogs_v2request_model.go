// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogsV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSQLLogsV2Request
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeSQLLogsV2Request
	GetDatabase() *string
	SetEndTime(v string) *DescribeSQLLogsV2Request
	GetEndTime() *string
	SetExecuteCost(v string) *DescribeSQLLogsV2Request
	GetExecuteCost() *string
	SetExecuteState(v string) *DescribeSQLLogsV2Request
	GetExecuteState() *string
	SetMaxExecuteCost(v string) *DescribeSQLLogsV2Request
	GetMaxExecuteCost() *string
	SetMinExecuteCost(v string) *DescribeSQLLogsV2Request
	GetMinExecuteCost() *string
	SetOperationClass(v string) *DescribeSQLLogsV2Request
	GetOperationClass() *string
	SetOperationType(v string) *DescribeSQLLogsV2Request
	GetOperationType() *string
	SetPageNumber(v string) *DescribeSQLLogsV2Request
	GetPageNumber() *string
	SetPageSize(v string) *DescribeSQLLogsV2Request
	GetPageSize() *string
	SetQueryKeywords(v string) *DescribeSQLLogsV2Request
	GetQueryKeywords() *string
	SetRegionId(v string) *DescribeSQLLogsV2Request
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSQLLogsV2Request
	GetResourceGroupId() *string
	SetSourceIP(v string) *DescribeSQLLogsV2Request
	GetSourceIP() *string
	SetStartTime(v string) *DescribeSQLLogsV2Request
	GetStartTime() *string
	SetUser(v string) *DescribeSQLLogsV2Request
	GetUser() *string
}

type DescribeSQLLogsV2Request struct {
	// The ID of instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
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
	// >  The end time must be later than the start time. The interval cannot be more than 24 hours.
	//
	// example:
	//
	// 2022-03-17T06:30Z
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
	// The maximum amount of time consumed by a slow query. Minimum value: 0. Unit: seconds.
	//
	// example:
	//
	// 1000
	MaxExecuteCost *string `json:"MaxExecuteCost,omitempty" xml:"MaxExecuteCost,omitempty"`
	// The minimum amount of time consumed by a slow query. Minimum value: 0. Unit: seconds.
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
	// > 	- If the **OperationClass*	- parameter is specified, the **OperationType*	- value must belong to the corresponding query language. For example, if the **OperationClass*	- value is **DQL**, the **OperationType*	- value must be a **DQL*	- SQL statement such as **SELECT**.
	//
	// >	- If the **OperationClass*	- parameter is not specified, the **OperationType*	- value can be an SQL statement of all query languages.
	//
	// >	- If neither of the **OperationClass*	- and **OperationType*	- parameters is specified, all types of SQL statements are returned.
	//
	// example:
	//
	// SELECT
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of the page to return. The maximum value is 200.
	//
	// example:
	//
	// 1
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keywords of the SQL statement.
	//
	// example:
	//
	// select 1
	QueryKeywords *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	// The region ID of the instance.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 100.XX.XX.90
	SourceIP *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	// The beginning of the time range. Specify the time in the ISO 8601 standard in the *yyyy-MM-ddTHH:mmZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2022-03-10T06:30Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// testadmin
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeSQLLogsV2Request) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogsV2Request) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsV2Request) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSQLLogsV2Request) GetDatabase() *string {
	return s.Database
}

func (s *DescribeSQLLogsV2Request) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSQLLogsV2Request) GetExecuteCost() *string {
	return s.ExecuteCost
}

func (s *DescribeSQLLogsV2Request) GetExecuteState() *string {
	return s.ExecuteState
}

func (s *DescribeSQLLogsV2Request) GetMaxExecuteCost() *string {
	return s.MaxExecuteCost
}

func (s *DescribeSQLLogsV2Request) GetMinExecuteCost() *string {
	return s.MinExecuteCost
}

func (s *DescribeSQLLogsV2Request) GetOperationClass() *string {
	return s.OperationClass
}

func (s *DescribeSQLLogsV2Request) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeSQLLogsV2Request) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeSQLLogsV2Request) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSQLLogsV2Request) GetQueryKeywords() *string {
	return s.QueryKeywords
}

func (s *DescribeSQLLogsV2Request) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSQLLogsV2Request) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSQLLogsV2Request) GetSourceIP() *string {
	return s.SourceIP
}

func (s *DescribeSQLLogsV2Request) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSQLLogsV2Request) GetUser() *string {
	return s.User
}

func (s *DescribeSQLLogsV2Request) SetDBInstanceId(v string) *DescribeSQLLogsV2Request {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetDatabase(v string) *DescribeSQLLogsV2Request {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetEndTime(v string) *DescribeSQLLogsV2Request {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetExecuteCost(v string) *DescribeSQLLogsV2Request {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetExecuteState(v string) *DescribeSQLLogsV2Request {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetMaxExecuteCost(v string) *DescribeSQLLogsV2Request {
	s.MaxExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetMinExecuteCost(v string) *DescribeSQLLogsV2Request {
	s.MinExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetOperationClass(v string) *DescribeSQLLogsV2Request {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetOperationType(v string) *DescribeSQLLogsV2Request {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetPageNumber(v string) *DescribeSQLLogsV2Request {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetPageSize(v string) *DescribeSQLLogsV2Request {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetQueryKeywords(v string) *DescribeSQLLogsV2Request {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetRegionId(v string) *DescribeSQLLogsV2Request {
	s.RegionId = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetResourceGroupId(v string) *DescribeSQLLogsV2Request {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetSourceIP(v string) *DescribeSQLLogsV2Request {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetStartTime(v string) *DescribeSQLLogsV2Request {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogsV2Request) SetUser(v string) *DescribeSQLLogsV2Request {
	s.User = &v
	return s
}

func (s *DescribeSQLLogsV2Request) Validate() error {
	return dara.Validate(s)
}
