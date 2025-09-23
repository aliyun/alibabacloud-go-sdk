// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsVpcsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeRdsVpcsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRdsVpcsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeRdsVpcsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeRdsVpcsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeRdsVpcsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRdsVpcsRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeRdsVpcsRequest
	GetSecurityToken() *string
	SetZoneId(v string) *DescribeRdsVpcsRequest
	GetZoneId() *string
}

type DescribeRdsVpcsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRdsVpcsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVpcsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRdsVpcsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRdsVpcsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRdsVpcsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRdsVpcsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRdsVpcsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRdsVpcsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeRdsVpcsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRdsVpcsRequest) SetOwnerAccount(v string) *DescribeRdsVpcsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetOwnerId(v int64) *DescribeRdsVpcsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetRegionId(v string) *DescribeRdsVpcsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetResourceGroupId(v string) *DescribeRdsVpcsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetResourceOwnerAccount(v string) *DescribeRdsVpcsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetResourceOwnerId(v int64) *DescribeRdsVpcsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetSecurityToken(v string) *DescribeRdsVpcsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetZoneId(v string) *DescribeRdsVpcsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) Validate() error {
	return dara.Validate(s)
}
