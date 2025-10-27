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
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-t4nj8619bz2w3****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database on which you want to execute the SQL statement.
	//
	// example:
	//
	// adb_demo
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// > - The end time must be later than the start time.
	//
	// > - The maximum time range that can be specified is 24 hours.
	//
	// example:
	//
	// 2022-01-23T22:18Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address and port number of the client that is used to execute the SQL statement.
	//
	// example:
	//
	// 100.104.XX.XX:43908
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The order in which specified fields are sorted. Specify this parameter as an ordered JSON array that consists of the Field and Type fields.
	//
	// 	- Field specifies the field that is used to sort the retrieved entries. Valid values:
	//
	//     	- HostAddress: the IP address of the client that is used to connect to the database.
	//
	//     	- Succeed: specifies whether the SQL statement is successfully executed.
	//
	//     	- TotalTime: the total amount of time that is consumed to execute the SQL statement.
	//
	//     	- DBName: the name of the database on which the SQL statement is executed.
	//
	//     	- SQLType: the type of the SQL statement.
	//
	//     	- User: the username that is used to execute the SQL statement.
	//
	//     	- ExecuteTime: the time to start executing the SQL statement.
	//
	// 	- Type specifies the sorting order. Valid values:
	//
	//     	- Desc: descending order.
	//
	//     	- Asc: ascending order.
	//
	// example:
	//
	// [{"Field":"ExecuteTime","Type":"Desc"},{"Field":"HostAddress","Type":"Asc"}]
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The sorting order of the retrieved entries. Valid values:
	//
	// 	- **asc**: sorts the retrieved entries by time in ascending order.
	//
	// 	- **desc**: sorts the retrieved entries by time in descending order.
	//
	// example:
	//
	// asc
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value is an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **10**
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// >  If you leave this parameter empty, the value 10 is used.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keywords that are included in the SQL statement to query.
	//
	// example:
	//
	// adb
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
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
	// 	- **DELETE**
	//
	// 	- **SELECT**
	//
	// 	- **UPDATE**
	//
	// 	- **INSERT_INTO_SELECT**
	//
	// 	- **ALTER**
	//
	// 	- **DROP**
	//
	// 	- **CREATE**
	//
	// > You can query only a single type of SQL statements at a time. If you leave this parameter empty, the **SELECT*	- statements are queried.
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// > SQL audit logs can be queried only when SQL audit is enabled. Only SQL audit logs within the last 30 days can be queried. If SQL audit was disabled and re-enabled, only SQL audit logs from the time when SQL audit was re-enabled can be queried.
	//
	// example:
	//
	// 2022-01-23T02:18Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Specifies whether the execution of the SQL statement succeeds. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Succeed *string `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// The name of the user who executed the SQL statement.
	//
	// example:
	//
	// test_user
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
