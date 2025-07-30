// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbType(v string) *DescribeParameterGroupsRequest
	GetDbType() *string
	SetOwnerAccount(v string) *DescribeParameterGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeParameterGroupsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeParameterGroupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeParameterGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeParameterGroupsRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeParameterGroupsRequest
	GetSecurityToken() *string
}

type DescribeParameterGroupsRequest struct {
	// The engine type. Valid values:
	//
	// 	- **redis**: Redis Open-Source Edition or Tair (In-Memory)
	//
	// 	- **tair_pena**: Tair (On NVM)
	//
	// 	- **tair_pdb**: Tair (On Disk)
	//
	// example:
	//
	// redis
	DbType       *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeParameterGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsRequest) GetDbType() *string {
	return s.DbType
}

func (s *DescribeParameterGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeParameterGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParameterGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeParameterGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeParameterGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeParameterGroupsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeParameterGroupsRequest) SetDbType(v string) *DescribeParameterGroupsRequest {
	s.DbType = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetOwnerAccount(v string) *DescribeParameterGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetOwnerId(v int64) *DescribeParameterGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetRegionId(v string) *DescribeParameterGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetResourceOwnerAccount(v string) *DescribeParameterGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetResourceOwnerId(v int64) *DescribeParameterGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetSecurityToken(v string) *DescribeParameterGroupsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeParameterGroupsRequest) Validate() error {
	return dara.Validate(s)
}
