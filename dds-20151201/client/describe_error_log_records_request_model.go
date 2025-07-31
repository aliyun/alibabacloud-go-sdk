// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeErrorLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeErrorLogRecordsRequest
	GetDBInstanceId() *string
	SetDBName(v string) *DescribeErrorLogRecordsRequest
	GetDBName() *string
	SetEndTime(v string) *DescribeErrorLogRecordsRequest
	GetEndTime() *string
	SetLogicalOperator(v string) *DescribeErrorLogRecordsRequest
	GetLogicalOperator() *string
	SetNodeId(v string) *DescribeErrorLogRecordsRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DescribeErrorLogRecordsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeErrorLogRecordsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeErrorLogRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeErrorLogRecordsRequest
	GetPageSize() *int32
	SetQueryKeywords(v string) *DescribeErrorLogRecordsRequest
	GetQueryKeywords() *string
	SetResourceGroupId(v string) *DescribeErrorLogRecordsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeErrorLogRecordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeErrorLogRecordsRequest
	GetResourceOwnerId() *int64
	SetRoleType(v string) *DescribeErrorLogRecordsRequest
	GetRoleType() *string
	SetStartTime(v string) *DescribeErrorLogRecordsRequest
	GetStartTime() *string
}

type DescribeErrorLogRecordsRequest struct {
	// The ID of the instance.
	//
	// >  If you set this parameter to the ID of a sharded cluster instance, you must also specify the **NodeId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp12c5b040dc****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// mongodbtest
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time and within 24 hours from the start time. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-01-02T12:10Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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
	// The ID of the mongos node or shard node whose error logs you want to query in the instance. If the instance is a sharded cluster instance, you must specify this parameter.
	//
	// >  This parameter is valid only when **DBInstanceId*	- is set to the ID of a sharded cluster instance.
	//
	// example:
	//
	// d-bp128a003436****
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
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
	// test test1
	QueryKeywords *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	// The ID of the resource group. For more information, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The role of the node whose error logs you want to query in the instance. Valid values:
	//
	// 	- **primary**
	//
	// 	- **secondary**
	//
	// >  If you set the **NodeId*	- parameter to the ID of a mongos node, the RoleType parameter must be set to **primary**.
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

func (s DescribeErrorLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeErrorLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeErrorLogRecordsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeErrorLogRecordsRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeErrorLogRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeErrorLogRecordsRequest) GetLogicalOperator() *string {
	return s.LogicalOperator
}

func (s *DescribeErrorLogRecordsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeErrorLogRecordsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeErrorLogRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeErrorLogRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeErrorLogRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeErrorLogRecordsRequest) GetQueryKeywords() *string {
	return s.QueryKeywords
}

func (s *DescribeErrorLogRecordsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeErrorLogRecordsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeErrorLogRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeErrorLogRecordsRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeErrorLogRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeErrorLogRecordsRequest) SetDBInstanceId(v string) *DescribeErrorLogRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetDBName(v string) *DescribeErrorLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetEndTime(v string) *DescribeErrorLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetLogicalOperator(v string) *DescribeErrorLogRecordsRequest {
	s.LogicalOperator = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetNodeId(v string) *DescribeErrorLogRecordsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetOwnerAccount(v string) *DescribeErrorLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetOwnerId(v int64) *DescribeErrorLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetPageNumber(v int32) *DescribeErrorLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetPageSize(v int32) *DescribeErrorLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetQueryKeywords(v string) *DescribeErrorLogRecordsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetResourceGroupId(v string) *DescribeErrorLogRecordsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeErrorLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeErrorLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetRoleType(v string) *DescribeErrorLogRecordsRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) SetStartTime(v string) *DescribeErrorLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeErrorLogRecordsRequest) Validate() error {
	return dara.Validate(s)
}
