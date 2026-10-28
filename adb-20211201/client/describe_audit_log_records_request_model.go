// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAuditLogRecordsRequest
	GetDBClusterId() *string
	SetDBName(v string) *DescribeAuditLogRecordsRequest
	GetDBName() *string
	SetEndTime(v string) *DescribeAuditLogRecordsRequest
	GetEndTime() *string
	SetEngineType(v string) *DescribeAuditLogRecordsRequest
	GetEngineType() *string
	SetHostAddress(v string) *DescribeAuditLogRecordsRequest
	GetHostAddress() *string
	SetOrder(v string) *DescribeAuditLogRecordsRequest
	GetOrder() *string
	SetOrderType(v string) *DescribeAuditLogRecordsRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *DescribeAuditLogRecordsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAuditLogRecordsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAuditLogRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAuditLogRecordsRequest
	GetPageSize() *int32
	SetProcessId(v string) *DescribeAuditLogRecordsRequest
	GetProcessId() *string
	SetProxyUser(v string) *DescribeAuditLogRecordsRequest
	GetProxyUser() *string
	SetQueryKeyword(v string) *DescribeAuditLogRecordsRequest
	GetQueryKeyword() *string
	SetRegionId(v string) *DescribeAuditLogRecordsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAuditLogRecordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAuditLogRecordsRequest
	GetResourceOwnerId() *int64
	SetSqlType(v string) *DescribeAuditLogRecordsRequest
	GetSqlType() *string
	SetStartTime(v string) *DescribeAuditLogRecordsRequest
	GetStartTime() *string
	SetSucceed(v string) *DescribeAuditLogRecordsRequest
	GetSucceed() *string
	SetUser(v string) *DescribeAuditLogRecordsRequest
	GetUser() *string
}

type DescribeAuditLogRecordsRequest struct {
	// <props="china">The cluster ID of an Enterprise Edition, Basic Edition, or Lakehouse Edition cluster.
	//
	// <props="intl">The cluster ID of a Lakehouse Edition cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all clusters in a specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-t4nj8619bz2w3****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database on which the target SQL statement is executed.
	//
	// example:
	//
	// adb_demo
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end time of the query. The time is displayed in UTC and in the yyyy-MM-ddTHH:mmZ format.
	//
	// > - The end time must be later than the start time.
	//
	// > - The interval between the start time and the end time cannot exceed 24 hours.
	//
	// example:
	//
	// 2022-08-12T17:08Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The engine type. Valid values:
	//
	// - XIHE: audit logs of the default compute engine.
	//
	// - AGENT_SERVERLESS: audit logs of the Serverless analytics feature.
	//
	// If this parameter is not specified, the default value is XIHE.
	//
	// example:
	//
	// XIHE
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The IP address and port number of the client that executes the target SQL statement.
	//
	// example:
	//
	// 100.104.XX.XX:43908
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The sort configuration in JSON format. The value is an ordered JSON array that specifies composite sorting based on the input array order. Each element contains the `Field` and `Type` parameters. Example: `[{"Field":"ExecutionStartTime","Type":"Desc"},{"Field":"ScanRows","Type":"Asc"}]`. Details:
	//
	// 	- `Field` specifies the field name for sorting. Valid values:
	//
	//     	- **HostAddress**: the client address used to connect to the database.
	//
	//     	- **UserName**: the username.
	//
	//     	- **ExecutionStartTime**: the start time of the target SQL statement execution.
	//
	//     	- **QueryTime**: the execution duration of the target SQL statement.
	//
	//     	- **PeakMemoryUsage**: the peak memory usage during the execution of the target SQL statement.
	//
	//     	- **ScanRows**: the number of scanned rows for tasks with a data source.
	//
	//     	- **ScanSize**: the volume of scanned data.
	//
	//     	- **ScanTime**: the total time consumed by data scanning.
	//
	//     	- **PlanningTime**: the time consumed by execution plan generation.
	//
	//     	- **WallTime**: the cumulative CPU time of all operators in the query across all nodes.
	//
	//     	- **ProcessID**: the process ID.
	//
	// 	- `Type` specifies the sort order. Valid values:
	//
	//     	- **Desc**: descending order.
	//
	//     	- **Asc**: ascending order.
	//
	// example:
	//
	// [{"Field":"ExecuteTime","Type":"Desc"},{"Field":"HostAddress","Type":"Asc"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The sort order of SQL statements by execution time. Valid values:
	//
	// 	- **asc**: ascending order.
	//
	// 	- **desc**: descending order.
	//
	// example:
	//
	// asc
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be a positive integer that does not exceed the maximum value of the INTEGER data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **10*	- (default)
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The process ID of the SQL query, used to filter audit logs of a specified process.
	//
	// example:
	//
	// 20210608175202172016166249034536****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// None
	ProxyUser *string `json:"ProxyUser,omitempty" xml:"ProxyUser,omitempty"`
	// A string value that is used as a keyword to filter the returned results.
	//
	// example:
	//
	// adb
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The SQL statement type. Valid values:
	//
	// - **DELETE**
	//
	// - **SELECT**
	//
	// - **UPDATE**
	//
	// - **INSERT INTO SELECT**
	//
	// - **ALTER**
	//
	// - **DROP**
	//
	// - **CREATE**
	//
	// > Only one type can be specified at a time. If this parameter is left empty, all types are queried by default.
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The start time of the query. The time is displayed in UTC and in the yyyy-MM-ddTHH:mmZ format.
	//
	// > SQL audit logs can be queried only when SQL Audit is enabled. Only SQL audit logs within the last 30 days can be queried. If SQL Audit is disabled and then re-enabled, only SQL audit logs generated after the re-enablement can be queried.
	//
	// example:
	//
	// 2022-08-12T04:17Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates whether the target SQL statement is executed successfully. Valid values:
	//
	// 	- **true**: The SQL statement is executed successfully.
	//
	// 	- **false**: The SQL statement fails to be executed.
	//
	// example:
	//
	// true
	Succeed *string `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// The username that executes the target SQL statement.
	//
	// example:
	//
	// test
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeAuditLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogRecordsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAuditLogRecordsRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeAuditLogRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAuditLogRecordsRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeAuditLogRecordsRequest) GetHostAddress() *string {
	return s.HostAddress
}

func (s *DescribeAuditLogRecordsRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeAuditLogRecordsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeAuditLogRecordsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAuditLogRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAuditLogRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAuditLogRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAuditLogRecordsRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *DescribeAuditLogRecordsRequest) GetProxyUser() *string {
	return s.ProxyUser
}

func (s *DescribeAuditLogRecordsRequest) GetQueryKeyword() *string {
	return s.QueryKeyword
}

func (s *DescribeAuditLogRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAuditLogRecordsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAuditLogRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAuditLogRecordsRequest) GetSqlType() *string {
	return s.SqlType
}

func (s *DescribeAuditLogRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAuditLogRecordsRequest) GetSucceed() *string {
	return s.Succeed
}

func (s *DescribeAuditLogRecordsRequest) GetUser() *string {
	return s.User
}

func (s *DescribeAuditLogRecordsRequest) SetDBClusterId(v string) *DescribeAuditLogRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetDBName(v string) *DescribeAuditLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetEndTime(v string) *DescribeAuditLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetEngineType(v string) *DescribeAuditLogRecordsRequest {
	s.EngineType = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetHostAddress(v string) *DescribeAuditLogRecordsRequest {
	s.HostAddress = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOrder(v string) *DescribeAuditLogRecordsRequest {
	s.Order = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOrderType(v string) *DescribeAuditLogRecordsRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOwnerAccount(v string) *DescribeAuditLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetOwnerId(v int64) *DescribeAuditLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetPageNumber(v int32) *DescribeAuditLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetPageSize(v int32) *DescribeAuditLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetProcessId(v string) *DescribeAuditLogRecordsRequest {
	s.ProcessId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetProxyUser(v string) *DescribeAuditLogRecordsRequest {
	s.ProxyUser = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetQueryKeyword(v string) *DescribeAuditLogRecordsRequest {
	s.QueryKeyword = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetRegionId(v string) *DescribeAuditLogRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeAuditLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeAuditLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetSqlType(v string) *DescribeAuditLogRecordsRequest {
	s.SqlType = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetStartTime(v string) *DescribeAuditLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetSucceed(v string) *DescribeAuditLogRecordsRequest {
	s.Succeed = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) SetUser(v string) *DescribeAuditLogRecordsRequest {
	s.User = &v
	return s
}

func (s *DescribeAuditLogRecordsRequest) Validate() error {
	return dara.Validate(s)
}
