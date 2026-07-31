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
	// <props="china">The ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all clusters in a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-t4nj8619bz2w3****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database on which the SQL statement was executed.
	//
	// example:
	//
	// adb_demo
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. The time must be in UTC and in the `yyyy-MM-ddTHH:mmZ` format.
	//
	// > - The end time must be later than the start time.
	//
	// >
	//
	// > - The time range cannot exceed 24 hours.
	//
	// example:
	//
	// 2022-08-12T17:08Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The client IP address and port number.
	//
	// example:
	//
	// 100.104.XX.XX:43908
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// Specifies the fields for sorting the results. The value is a JSON string that is an array of objects. The order of objects in the array defines the sort priority. Each object contains the`Field` and`Type` parameters. Example: `[{"Field":"ExecutionStartTime","Type":"Desc"},{"Field":"ScanRows","Type":"Asc"}]`.
	//
	// - `Field`: the field by which to sort the results. Valid values:
	//
	//   - **HostAddress**: the client IP address.
	//
	//   - **UserName**: the username.
	//
	//   - **ExecutionStartTime**: the execution start time of the SQL statement.
	//
	//   - **QueryTime**: the execution duration.
	//
	//   - **PeakMemoryUsage**: the peak memory usage of the SQL statement.
	//
	//   - **ScanRows**: the number of rows scanned by a task that involves a data source.
	//
	//   - **ScanSize**: the amount of data scanned.
	//
	//   - **ScanTime**: the time taken for the data scan.
	//
	//   - **PlanningTime**: the time taken to generate the execution plan.
	//
	//   - **WallTime**: the total CPU time of all operators on all nodes.
	//
	//   - **ProcessID**: the process ID.
	//
	// - `Type`: the sort order. Valid values:
	//
	//   - **Desc**: descending order.
	//
	//   - **Asc**: ascending order.
	//
	// example:
	//
	// [{"Field":"ExecuteTime","Type":"Desc"},{"Field":"HostAddress","Type":"Asc"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The sort order for the results based on execution time. Valid values:
	//
	// - **asc**: ascending order.
	//
	// - **desc**: descending order.
	//
	// example:
	//
	// asc
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size. Valid values:
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
	// A reserved parameter.
	//
	// example:
	//
	// 无
	ProxyUser *string `json:"ProxyUser,omitempty" xml:"ProxyUser,omitempty"`
	// A keyword used to perform a fuzzy search on the returned results.
	//
	// example:
	//
	// adb
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the SQL statement. Valid values:
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
	// > You can specify only one type per request. If this parameter is not specified, all types are queried by default.
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The start of the time range to query. The time must be in UTC and in the `yyyy-MM-ddTHH:mmZ` format.
	//
	// > You can query SQL audit logs only when this feature is enabled. Logs are available for the last 30 days. If you disable and then re-enable SQL audit, only logs generated after the feature was re-enabled are returned.
	//
	// example:
	//
	// 2022-08-12T04:17Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates whether the SQL statement was successfully executed. Valid values:
	//
	// - **true**: The SQL statement succeeded.
	//
	// - **false**: The SQL statement failed.
	//
	// example:
	//
	// true
	Succeed *string `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// The username that executed the SQL statement.
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
