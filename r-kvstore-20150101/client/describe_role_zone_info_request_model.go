// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoleZoneInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeRoleZoneInfoRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeRoleZoneInfoRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRoleZoneInfoRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeRoleZoneInfoRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRoleZoneInfoRequest
	GetPageSize() *int32
	SetQueryType(v int32) *DescribeRoleZoneInfoRequest
	GetQueryType() *int32
	SetResourceOwnerAccount(v string) *DescribeRoleZoneInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRoleZoneInfoRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeRoleZoneInfoRequest
	GetSecurityToken() *string
}

type DescribeRoleZoneInfoRequest struct {
	// The ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-t4nlenc2p04uvb****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be an integer that is greater than **0*	- and less than or equal to the maximum value supported by the integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **10**, **20**, and **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the node to query. Default value: 1. Valid values:
	//
	// 	- **0**: proxy node
	//
	//     **
	//
	//     **Note*	- This parameter is supported only for cluster instances and read/write splitting instances.
	//
	// 	- **1**: data node
	//
	// example:
	//
	// 0
	QueryType            *int32  `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeRoleZoneInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleZoneInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRoleZoneInfoRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRoleZoneInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRoleZoneInfoRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRoleZoneInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRoleZoneInfoRequest) GetQueryType() *int32 {
	return s.QueryType
}

func (s *DescribeRoleZoneInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRoleZoneInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRoleZoneInfoRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeRoleZoneInfoRequest) SetInstanceId(v string) *DescribeRoleZoneInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetOwnerAccount(v string) *DescribeRoleZoneInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetOwnerId(v int64) *DescribeRoleZoneInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetPageNumber(v int32) *DescribeRoleZoneInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetPageSize(v int32) *DescribeRoleZoneInfoRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetQueryType(v int32) *DescribeRoleZoneInfoRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetResourceOwnerAccount(v string) *DescribeRoleZoneInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetResourceOwnerId(v int64) *DescribeRoleZoneInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) SetSecurityToken(v string) *DescribeRoleZoneInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRoleZoneInfoRequest) Validate() error {
	return dara.Validate(s)
}
