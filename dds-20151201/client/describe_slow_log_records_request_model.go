// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSlowLogRecordsRequest
	GetDBInstanceId() *string
	SetDBName(v string) *DescribeSlowLogRecordsRequest
	GetDBName() *string
	SetEndTime(v string) *DescribeSlowLogRecordsRequest
	GetEndTime() *string
	SetLogicalOperator(v string) *DescribeSlowLogRecordsRequest
	GetLogicalOperator() *string
	SetNodeId(v string) *DescribeSlowLogRecordsRequest
	GetNodeId() *string
	SetOrderType(v string) *DescribeSlowLogRecordsRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *DescribeSlowLogRecordsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSlowLogRecordsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSlowLogRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSlowLogRecordsRequest
	GetPageSize() *int32
	SetQueryKeywords(v string) *DescribeSlowLogRecordsRequest
	GetQueryKeywords() *string
	SetResourceGroupId(v string) *DescribeSlowLogRecordsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeSlowLogRecordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSlowLogRecordsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeSlowLogRecordsRequest
	GetStartTime() *string
}

type DescribeSlowLogRecordsRequest struct {
	// The instance ID.
	//
	// > If you set this parameter to the ID of a sharded cluster instance, you must also specify the `NodeId` parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1fc7e65108****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// mongodbtest
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// >
	//
	// 	- The end time must be later than the start time.
	//
	// 	- The end time must be within 24 hours from the start time. Otherwise, the query fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-08-16T14:13Z
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
	// The ID of the shard node.
	//
	// > This parameter is required only when you specify the `DBInstanceId` parameter to the ID of a sharded cluster instance.
	//
	// example:
	//
	// d-bp18b06ebc21****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The order of time in which the log entries to return are sorted. Valid values:
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
	// The page number of the page to return. The value must be a positive integer that does not exceed the maximum value of the INTEGER data type. Default value: **1**.
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
	// test test1
	QueryKeywords *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	// The ID of the resource group to which the instances you want to query belong.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-08-15T14:13Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSlowLogRecordsRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSlowLogRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSlowLogRecordsRequest) GetLogicalOperator() *string {
	return s.LogicalOperator
}

func (s *DescribeSlowLogRecordsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSlowLogRecordsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeSlowLogRecordsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSlowLogRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSlowLogRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSlowLogRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSlowLogRecordsRequest) GetQueryKeywords() *string {
	return s.QueryKeywords
}

func (s *DescribeSlowLogRecordsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSlowLogRecordsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSlowLogRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSlowLogRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSlowLogRecordsRequest) SetDBInstanceId(v string) *DescribeSlowLogRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBName(v string) *DescribeSlowLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetEndTime(v string) *DescribeSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetLogicalOperator(v string) *DescribeSlowLogRecordsRequest {
	s.LogicalOperator = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetNodeId(v string) *DescribeSlowLogRecordsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOrderType(v string) *DescribeSlowLogRecordsRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOwnerAccount(v string) *DescribeSlowLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOwnerId(v int64) *DescribeSlowLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageNumber(v int32) *DescribeSlowLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageSize(v int32) *DescribeSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetQueryKeywords(v string) *DescribeSlowLogRecordsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetResourceGroupId(v string) *DescribeSlowLogRecordsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeSlowLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeSlowLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetStartTime(v string) *DescribeSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) Validate() error {
	return dara.Validate(s)
}
