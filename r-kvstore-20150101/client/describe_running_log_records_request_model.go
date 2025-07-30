// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRunningLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCharacterType(v string) *DescribeRunningLogRecordsRequest
	GetCharacterType() *string
	SetDBName(v string) *DescribeRunningLogRecordsRequest
	GetDBName() *string
	SetEndTime(v string) *DescribeRunningLogRecordsRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeRunningLogRecordsRequest
	GetInstanceId() *string
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
	SetQueryKeyword(v string) *DescribeRunningLogRecordsRequest
	GetQueryKeyword() *string
	SetResourceGroupId(v string) *DescribeRunningLogRecordsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeRunningLogRecordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRunningLogRecordsRequest
	GetResourceOwnerId() *int64
	SetRoleType(v string) *DescribeRunningLogRecordsRequest
	GetRoleType() *string
	SetSecurityToken(v string) *DescribeRunningLogRecordsRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeRunningLogRecordsRequest
	GetStartTime() *string
}

type DescribeRunningLogRecordsRequest struct {
	// The shard type of the cluster instance. Valid values:
	//
	// 	- **proxy**: proxy node
	//
	// 	- **db**: data node
	//
	// 	- **cs**: config server node
	//
	// >  If you set this parameter, you must also set the **NodeId*	- parameter.
	//
	// example:
	//
	// proxy
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// 0
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. The time range cannot exceed one day. We recommend that you specify 1 hour. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-12-03T08:01Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the node in the instance. You can set this parameter to query the operational logs of a specified node.
	//
	// >
	//
	// 	- This parameter is available only for read/write splitting and cluster instances.
	//
	// 	- If you set this parameter, you must also set the **CharacterType*	- parameter.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The method that is used to sort the returned log entries. Valid values:
	//
	// 	- **asc**: ascending order
	//
	// 	- **desc**: descending order
	//
	// example:
	//
	// asc
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than **0*	- and less than or equal to the maximum value supported by the integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **30**, **50**, and **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword that is used to query operational logs.
	//
	// example:
	//
	// aof
	QueryKeyword *string `json:"QueryKeyword,omitempty" xml:"QueryKeyword,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The role of the data shard. Default value: master. Valid values:
	//
	// 	- **master**: master node
	//
	// 	- **slave**: replica node
	//
	// example:
	//
	// master
	RoleType      *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The beginning of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-12-03T07:01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRunningLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRunningLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRunningLogRecordsRequest) GetCharacterType() *string {
	return s.CharacterType
}

func (s *DescribeRunningLogRecordsRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeRunningLogRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRunningLogRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
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

func (s *DescribeRunningLogRecordsRequest) GetQueryKeyword() *string {
	return s.QueryKeyword
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

func (s *DescribeRunningLogRecordsRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeRunningLogRecordsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeRunningLogRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRunningLogRecordsRequest) SetCharacterType(v string) *DescribeRunningLogRecordsRequest {
	s.CharacterType = &v
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

func (s *DescribeRunningLogRecordsRequest) SetInstanceId(v string) *DescribeRunningLogRecordsRequest {
	s.InstanceId = &v
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

func (s *DescribeRunningLogRecordsRequest) SetQueryKeyword(v string) *DescribeRunningLogRecordsRequest {
	s.QueryKeyword = &v
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

func (s *DescribeRunningLogRecordsRequest) SetRoleType(v string) *DescribeRunningLogRecordsRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetSecurityToken(v string) *DescribeRunningLogRecordsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) SetStartTime(v string) *DescribeRunningLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRunningLogRecordsRequest) Validate() error {
	return dara.Validate(s)
}
