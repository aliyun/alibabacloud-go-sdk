// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeAuditRecordsRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeAuditRecordsRequest
	GetDatabase() *string
	SetEndTime(v string) *DescribeAuditRecordsRequest
	GetEndTime() *string
	SetForm(v string) *DescribeAuditRecordsRequest
	GetForm() *string
	SetLogicalOperator(v string) *DescribeAuditRecordsRequest
	GetLogicalOperator() *string
	SetNodeId(v string) *DescribeAuditRecordsRequest
	GetNodeId() *string
	SetOrderType(v string) *DescribeAuditRecordsRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *DescribeAuditRecordsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAuditRecordsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAuditRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAuditRecordsRequest
	GetPageSize() *int32
	SetQueryKeywords(v string) *DescribeAuditRecordsRequest
	GetQueryKeywords() *string
	SetResourceOwnerAccount(v string) *DescribeAuditRecordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAuditRecordsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeAuditRecordsRequest
	GetStartTime() *string
	SetUser(v string) *DescribeAuditRecordsRequest
	GetUser() *string
}

type DescribeAuditRecordsRequest struct {
	// The instance ID.
	//
	// > If you set this parameter to the ID of a sharded cluster instance, you must also specify the **NodeId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp12c5b040dc****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database to be queried. By default, all databases are queried.
	//
	// example:
	//
	// database****
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// > The end time must be within 24 hours from the start time. Otherwise, the query fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-03-13T13:11:14Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The form of the audit log that the operation returns. Valid values:
	//
	// 	- **File**: triggers the generation of audit logs. If this parameter is set to File, only common parameters are returned.
	//
	// 	- **Stream*	- (default): returns data streams.
	//
	// example:
	//
	// Stream
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The logical relationship between multiple keywords. Valid values:
	//
	// 	- **or**
	//
	// 	- **and*	- (default value)
	//
	// example:
	//
	// and
	LogicalOperator *string `json:"LogicalOperator,omitempty" xml:"LogicalOperator,omitempty"`
	// The ID of the mongos node or shard node in the instance.
	//
	// > This parameter takes effect only when you set the **DBInstanceId*	- parameter to the ID of a sharded cluster instance.
	//
	// example:
	//
	// d-bp128a003436****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The order of time in which the log entries to return are sorted. Valid values:
	//
	// 	- **asc**: The log entries are sorted by time in ascending order.
	//
	// 	- **desc**: The log entries are sorted by time in descending order.
	//
	// example:
	//
	// asc
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the page to return. The valid value must be a positive integer that does not exceed the maximum value of the INTEGER data type. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Default value: 30. Valid values: **30**, **50**, and **100**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keywords used for query. You can enter up to 10 keywords at a time. If you enter multiple keywords, separate the keywords with spaces.
	//
	// example:
	//
	// slow
	QueryKeywords        *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-03-13T12:11:14Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The user of the database. If you do not specify this parameter, this operation returns records of all users.
	//
	// example:
	//
	// test
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeAuditRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditRecordsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAuditRecordsRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeAuditRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAuditRecordsRequest) GetForm() *string {
	return s.Form
}

func (s *DescribeAuditRecordsRequest) GetLogicalOperator() *string {
	return s.LogicalOperator
}

func (s *DescribeAuditRecordsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeAuditRecordsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeAuditRecordsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAuditRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAuditRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAuditRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAuditRecordsRequest) GetQueryKeywords() *string {
	return s.QueryKeywords
}

func (s *DescribeAuditRecordsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAuditRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAuditRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAuditRecordsRequest) GetUser() *string {
	return s.User
}

func (s *DescribeAuditRecordsRequest) SetDBInstanceId(v string) *DescribeAuditRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetDatabase(v string) *DescribeAuditRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetEndTime(v string) *DescribeAuditRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetForm(v string) *DescribeAuditRecordsRequest {
	s.Form = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetLogicalOperator(v string) *DescribeAuditRecordsRequest {
	s.LogicalOperator = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetNodeId(v string) *DescribeAuditRecordsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetOrderType(v string) *DescribeAuditRecordsRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetOwnerAccount(v string) *DescribeAuditRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetOwnerId(v int64) *DescribeAuditRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetPageNumber(v int32) *DescribeAuditRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetPageSize(v int32) *DescribeAuditRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetQueryKeywords(v string) *DescribeAuditRecordsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetResourceOwnerAccount(v string) *DescribeAuditRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetResourceOwnerId(v int64) *DescribeAuditRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetStartTime(v string) *DescribeAuditRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAuditRecordsRequest) SetUser(v string) *DescribeAuditRecordsRequest {
	s.User = &v
	return s
}

func (s *DescribeAuditRecordsRequest) Validate() error {
	return dara.Validate(s)
}
