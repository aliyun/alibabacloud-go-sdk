// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeVpcsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpcsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeVpcsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVpcsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpcsRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeVpcsRequest
	GetSecurityToken() *string
	SetVpcId(v string) *DescribeVpcsRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeVpcsRequest
	GetZoneId() *string
}

type DescribeVpcsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp13h7uzhulpuxvnp****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVpcsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpcsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpcsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpcsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpcsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpcsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVpcsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVpcsRequest) SetOwnerAccount(v string) *DescribeVpcsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpcsRequest) SetOwnerId(v int64) *DescribeVpcsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcsRequest) SetRegionId(v string) *DescribeVpcsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpcsRequest) SetResourceOwnerAccount(v string) *DescribeVpcsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpcsRequest) SetResourceOwnerId(v int64) *DescribeVpcsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpcsRequest) SetSecurityToken(v string) *DescribeVpcsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVpcsRequest) SetVpcId(v string) *DescribeVpcsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcsRequest) SetZoneId(v string) *DescribeVpcsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeVpcsRequest) Validate() error {
	return dara.Validate(s)
}
