// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *CreateNetworkChannelRequest
	GetChannelName() *string
	SetClientToken(v string) *CreateNetworkChannelRequest
	GetClientToken() *string
	SetDBClusterId(v string) *CreateNetworkChannelRequest
	GetDBClusterId() *string
	SetNotes(v string) *CreateNetworkChannelRequest
	GetNotes() *string
	SetOwnerAccount(v string) *CreateNetworkChannelRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateNetworkChannelRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateNetworkChannelRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateNetworkChannelRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateNetworkChannelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateNetworkChannelRequest
	GetResourceOwnerId() *int64
	SetTargetDBClusterId(v string) *CreateNetworkChannelRequest
	GetTargetDBClusterId() *string
	SetTargetIp(v string) *CreateNetworkChannelRequest
	GetTargetIp() *string
	SetTargetPort(v string) *CreateNetworkChannelRequest
	GetTargetPort() *string
	SetVpcId(v string) *CreateNetworkChannelRequest
	GetVpcId() *string
}

type CreateNetworkChannelRequest struct {
	// The name of the network channel. The name must consist of lowercase letters, digits, and underscores (_). It must start and end with a letter or a digit. The name can be up to 64 characters long.
	//
	// This parameter is required.
	//
	// example:
	//
	// ch4
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// A client token that ensures the request is idempotent. The client generates this token. The token must be unique for each request. It is case-sensitive and can be up to 64 ASCII characters long.
	//
	// example:
	//
	// 6000170000591aed949d0f5********************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the source instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The notes.
	//
	// example:
	//
	// test
	Notes        *string `json:"Notes,omitempty" xml:"Notes,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-re*********
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the destination instance.
	//
	// example:
	//
	// pc-*****************
	TargetDBClusterId *string `json:"TargetDBClusterId,omitempty" xml:"TargetDBClusterId,omitempty"`
	// The IP address of the destination instance.
	//
	// example:
	//
	// 192.**.**.46
	TargetIp *string `json:"TargetIp,omitempty" xml:"TargetIp,omitempty"`
	// The port of the destination instance.
	//
	// example:
	//
	// 9032
	TargetPort *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	// The ID of the virtual private cloud (VPC) where the endpoint is located.
	//
	// example:
	//
	// vpc-25cdvfeq58pl****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateNetworkChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkChannelRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *CreateNetworkChannelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNetworkChannelRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateNetworkChannelRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateNetworkChannelRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateNetworkChannelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateNetworkChannelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNetworkChannelRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateNetworkChannelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateNetworkChannelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateNetworkChannelRequest) GetTargetDBClusterId() *string {
	return s.TargetDBClusterId
}

func (s *CreateNetworkChannelRequest) GetTargetIp() *string {
	return s.TargetIp
}

func (s *CreateNetworkChannelRequest) GetTargetPort() *string {
	return s.TargetPort
}

func (s *CreateNetworkChannelRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateNetworkChannelRequest) SetChannelName(v string) *CreateNetworkChannelRequest {
	s.ChannelName = &v
	return s
}

func (s *CreateNetworkChannelRequest) SetClientToken(v string) *CreateNetworkChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNetworkChannelRequest) SetDBClusterId(v string) *CreateNetworkChannelRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateNetworkChannelRequest) SetNotes(v string) *CreateNetworkChannelRequest {
	s.Notes = &v
	return s
}

func (s *CreateNetworkChannelRequest) SetOwnerAccount(v string) *CreateNetworkChannelRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateNetworkChannelRequest) SetOwnerId(v int64) *CreateNetworkChannelRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNetworkChannelRequest) SetRegionId(v string) *CreateNetworkChannelRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkChannelRequest) SetResourceGroupId(v string) *CreateNetworkChannelRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateNetworkChannelRequest) SetResourceOwnerAccount(v string) *CreateNetworkChannelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNetworkChannelRequest) SetResourceOwnerId(v int64) *CreateNetworkChannelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateNetworkChannelRequest) SetTargetDBClusterId(v string) *CreateNetworkChannelRequest {
	s.TargetDBClusterId = &v
	return s
}

func (s *CreateNetworkChannelRequest) SetTargetIp(v string) *CreateNetworkChannelRequest {
	s.TargetIp = &v
	return s
}

func (s *CreateNetworkChannelRequest) SetTargetPort(v string) *CreateNetworkChannelRequest {
	s.TargetPort = &v
	return s
}

func (s *CreateNetworkChannelRequest) SetVpcId(v string) *CreateNetworkChannelRequest {
	s.VpcId = &v
	return s
}

func (s *CreateNetworkChannelRequest) Validate() error {
	return dara.Validate(s)
}
