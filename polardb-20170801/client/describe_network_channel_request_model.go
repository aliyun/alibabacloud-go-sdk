// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *DescribeNetworkChannelRequest
	GetChannelName() *string
	SetClientToken(v string) *DescribeNetworkChannelRequest
	GetClientToken() *string
	SetDBClusterId(v string) *DescribeNetworkChannelRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeNetworkChannelRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeNetworkChannelRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeNetworkChannelRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeNetworkChannelRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeNetworkChannelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeNetworkChannelRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *DescribeNetworkChannelRequest
	GetVpcId() *string
}

type DescribeNetworkChannelRequest struct {
	// example:
	//
	// ch4
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// example:
	//
	// vpc-bp1qpo0kug3a20qqe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeNetworkChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkChannelRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkChannelRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *DescribeNetworkChannelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeNetworkChannelRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeNetworkChannelRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeNetworkChannelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNetworkChannelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNetworkChannelRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeNetworkChannelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeNetworkChannelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeNetworkChannelRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNetworkChannelRequest) SetChannelName(v string) *DescribeNetworkChannelRequest {
	s.ChannelName = &v
	return s
}

func (s *DescribeNetworkChannelRequest) SetClientToken(v string) *DescribeNetworkChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeNetworkChannelRequest) SetDBClusterId(v string) *DescribeNetworkChannelRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeNetworkChannelRequest) SetOwnerAccount(v string) *DescribeNetworkChannelRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeNetworkChannelRequest) SetOwnerId(v int64) *DescribeNetworkChannelRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNetworkChannelRequest) SetRegionId(v string) *DescribeNetworkChannelRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkChannelRequest) SetResourceGroupId(v string) *DescribeNetworkChannelRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeNetworkChannelRequest) SetResourceOwnerAccount(v string) *DescribeNetworkChannelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNetworkChannelRequest) SetResourceOwnerId(v int64) *DescribeNetworkChannelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeNetworkChannelRequest) SetVpcId(v string) *DescribeNetworkChannelRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeNetworkChannelRequest) Validate() error {
	return dara.Validate(s)
}
