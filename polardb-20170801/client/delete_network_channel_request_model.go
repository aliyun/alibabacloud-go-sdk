// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *DeleteNetworkChannelRequest
	GetChannelName() *string
	SetClientToken(v string) *DeleteNetworkChannelRequest
	GetClientToken() *string
	SetDBClusterId(v string) *DeleteNetworkChannelRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DeleteNetworkChannelRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteNetworkChannelRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteNetworkChannelRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteNetworkChannelRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DeleteNetworkChannelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteNetworkChannelRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *DeleteNetworkChannelRequest
	GetVpcId() *string
}

type DeleteNetworkChannelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ch4
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
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

func (s DeleteNetworkChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkChannelRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *DeleteNetworkChannelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteNetworkChannelRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteNetworkChannelRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteNetworkChannelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteNetworkChannelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNetworkChannelRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteNetworkChannelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteNetworkChannelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteNetworkChannelRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DeleteNetworkChannelRequest) SetChannelName(v string) *DeleteNetworkChannelRequest {
	s.ChannelName = &v
	return s
}

func (s *DeleteNetworkChannelRequest) SetClientToken(v string) *DeleteNetworkChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteNetworkChannelRequest) SetDBClusterId(v string) *DeleteNetworkChannelRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteNetworkChannelRequest) SetOwnerAccount(v string) *DeleteNetworkChannelRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteNetworkChannelRequest) SetOwnerId(v int64) *DeleteNetworkChannelRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteNetworkChannelRequest) SetRegionId(v string) *DeleteNetworkChannelRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNetworkChannelRequest) SetResourceGroupId(v string) *DeleteNetworkChannelRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteNetworkChannelRequest) SetResourceOwnerAccount(v string) *DeleteNetworkChannelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteNetworkChannelRequest) SetResourceOwnerId(v int64) *DeleteNetworkChannelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteNetworkChannelRequest) SetVpcId(v string) *DeleteNetworkChannelRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteNetworkChannelRequest) Validate() error {
	return dara.Validate(s)
}
