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
	SetResourceGroupId(v string) *DescribeRdsVpcsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeRdsVpcsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRdsVpcsRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeRdsVpcsRequest
	GetZoneId() *string
}

type DescribeRdsVpcsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// cn-beijing-l
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

func (s *DescribeRdsVpcsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRdsVpcsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRdsVpcsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
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

func (s *DescribeRdsVpcsRequest) SetZoneId(v string) *DescribeRdsVpcsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) Validate() error {
	return dara.Validate(s)
}
