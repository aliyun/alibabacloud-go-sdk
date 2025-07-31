// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRunningLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeRunningLogRecordsRequest
	GetDBInstanceId() *string
	SetDBName(v string) *DescribeRunningLogRecordsRequest
	GetDBName() *string
	SetEndTime(v string) *DescribeRunningLogRecordsRequest
	GetEndTime() *string
	SetLogicalOperator(v string) *DescribeRunningLogRecordsRequest
	GetLogicalOperator() *string
	SetNodeId(v string) *DescribeRunningLogRecordsRequest
	GetNodeId() *string
	SetOrderType(v string) *DescribeRunningLogRecordsRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *DescribeRunningLogRecordsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRunningLogRecordsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeRunningLogRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRunningLogRecordsRequest
	GetPageSize() *int32
	SetQueryKeywords(v string) *DescribeRunningLogRecordsRequest
	GetQueryKeywords() *string
	SetResourceGroupId(v string) *DescribeRunningLogRecordsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeRunningLogRecordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRunningLogRecordsRequest
	GetResourceOwnerId() *int64
	SetRoleId(v string) *DescribeRunningLogRecordsRequest
	GetRoleId() *string
	SetRoleType(v string) *DescribeRunningLogRecordsRequest
	GetRoleType() *string
	SetStartTime(v string) *DescribeRunningLogRecordsRequest
	GetStartTime() *string
}

type DescribeRunningLogRecordsRequest struct {
	// The ID of the instance.
	//
	// >  If you set this parameter to the ID of a sharded cluster instance, you must also specify the **NodeId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// mongodbtest
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// >  The end time must be later than the start time and within 24 hours from the start time. Otherwise, the query fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-01-01T13:10Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The logical relationship among multiple keywords.
	//
	// 	- **or**
	//
	// 	- **and*	- (default value)
	//
	// example:
	//
	// and
	LogicalOperator *string `json:"LogicalOperator,omitempty" xml:"LogicalOperator,omitempty"`
	// The ID of the mongos node or shard node whose operational logs you want to query in the instance. If the instance is a sharded cluster instance, you must specify this parameter.
	//
	// >  This parameter is valid only when **DBInstanceId*	- is set to the ID of a sharded cluster instance.
	//
	// example:
	//
	// d-bpxxxxxxxx
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The order of time in which the operational log entries to return are sorted. Valid values:
	//
	// 	- asc: The log entries are sorted by time in ascending order.
	//
	// 	- desc: The log entries are sorted by time in descending order.
	//
	// example:
	//
	// asc
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **30*	- to **100**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keywords used for query. You can enter up to 10 keywords at a time. If you enter multiple keywords, separate the keywords with spaces.
	//
	// example:
	//
	// test test2
	QueryKeywords *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// sg-bpxxxxxxxxxxxxxxxxxx
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The role ID of the node. You can call the [DescribeReplicaSetRole](https://help.aliyun.com/document_detail/62134.html) operation to query the role ID.
	//
	// example:
	//
	// 651xxxxx
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The role of the node whose error logs you want to query in the instance. Valid values:
	//
	// 	- **primary**
	//
	// 	- **secondary**
	//
	// >  If you set the **NodeId*	- parameter to the ID of a mongos node, the **RoleType*	- parameter must be set to **primary**.
	//
	// example:
	//
	// primary
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The beginning of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-01-01T12:10Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRunningLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRunningLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRunningLogRecordsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeRunningLogRecordsRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeRunningLogRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRunningLogRecordsRequest) GetLogicalOperator() *string {
	return s.LogicalOperator
}

func (s *DescribeRunningLogRecordsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeRunningLogRecordsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeRunningLogRecordsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRunningLogRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRunningLogRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRunningLogRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRunningLogRecordsRequest) GetQueryKeywords() *string {
	return s.QueryKeywords
}

func (s *DescribeRunningLogRecordsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRunningLogRecordsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRunningLogRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRunningLogRecordsRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *DescribeRunningLogRecordsRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeRunningLogRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRunningLogRecordsRequest) SetDBInstanceId(v string) *DescribeRunningLogRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetDBName(v string) *DescribeRunningLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetEndTime(v string) *DescribeRunningLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetLogicalOperator(v string) *DescribeRunningLogRecordsRequest {
	s.LogicalOperator = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetNodeId(v string) *DescribeRunningLogRecordsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetOrderType(v string) *DescribeRunningLogRecordsRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetOwnerAccount(v string) *DescribeRunningLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetOwnerId(v int64) *DescribeRunningLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetPageNumber(v int32) *DescribeRunningLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetPageSize(v int32) *DescribeRunningLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetQueryKeywords(v string) *DescribeRunningLogRecordsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetResourceGroupId(v string) *DescribeRunningLogRecordsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeRunningLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeRunningLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetRoleId(v string) *DescribeRunningLogRecordsRequest {
	s.RoleId = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetRoleType(v string) *DescribeRunningLogRecordsRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetStartTime(v string) *DescribeRunningLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) Validate() error {
	return dara.Validate(s)
}
