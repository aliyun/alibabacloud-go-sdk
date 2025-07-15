// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBgpPeerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBfdMultiHop(v int32) *CreateBgpPeerRequest
	GetBfdMultiHop() *int32
	SetBgpGroupId(v string) *CreateBgpPeerRequest
	GetBgpGroupId() *string
	SetClientToken(v string) *CreateBgpPeerRequest
	GetClientToken() *string
	SetEnableBfd(v bool) *CreateBgpPeerRequest
	GetEnableBfd() *bool
	SetIpVersion(v string) *CreateBgpPeerRequest
	GetIpVersion() *string
	SetOwnerAccount(v string) *CreateBgpPeerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateBgpPeerRequest
	GetOwnerId() *int64
	SetPeerIpAddress(v string) *CreateBgpPeerRequest
	GetPeerIpAddress() *string
	SetRegionId(v string) *CreateBgpPeerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateBgpPeerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateBgpPeerRequest
	GetResourceOwnerId() *int64
}

type CreateBgpPeerRequest struct {
	// The BFD hop count. Valid values: **1*	- to **255**.
	//
	// This parameter is required only if you enable BFD. The parameter specifies the maximum number of network devices that a packet can traverse from the source to the destination. Set a value based on your network topology.
	//
	// > If you use BFD in a multi-cloud environment or a fiber-optic direct connection network without any bridge device, you need to change the default BFD hop count from **255*	- to **1**.
	//
	// example:
	//
	// 3
	BfdMultiHop *int32 `json:"BfdMultiHop,omitempty" xml:"BfdMultiHop,omitempty"`
	// The ID of the BGP group.
	//
	// This parameter is required.
	//
	// example:
	//
	// bgpg-wz9f62v4fbg****
	BgpGroupId *string `json:"BgpGroupId,omitempty" xml:"BgpGroupId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not set this parameter, the system uses the value of **RequestId*	- as **ClientToken**. The value of **RequestId*	- for each API request is different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable the Bidirectional Forwarding Detection (BFD) feature. Valid values:
	//
	// 	- **true**: enables BFD.
	//
	// 	- **false**: disables BFD.
	//
	// example:
	//
	// true
	EnableBfd *bool `json:"EnableBfd,omitempty" xml:"EnableBfd,omitempty"`
	// The IP version. Valid values:
	//
	// 	- **IPv4**: This is the default value.
	//
	// 	- **IPv6**: IPv6 is supported only if the VBR for which you want to create the BGP group has IPv6 enabled.
	//
	// example:
	//
	// IPv4
	IpVersion    *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The IP address of the BGP peer.
	//
	// example:
	//
	// 116.62.XX.XX
	PeerIpAddress *string `json:"PeerIpAddress,omitempty" xml:"PeerIpAddress,omitempty"`
	// The ID of the region to which the BGP group belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateBgpPeerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBgpPeerRequest) GoString() string {
	return s.String()
}

func (s *CreateBgpPeerRequest) GetBfdMultiHop() *int32 {
	return s.BfdMultiHop
}

func (s *CreateBgpPeerRequest) GetBgpGroupId() *string {
	return s.BgpGroupId
}

func (s *CreateBgpPeerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateBgpPeerRequest) GetEnableBfd() *bool {
	return s.EnableBfd
}

func (s *CreateBgpPeerRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CreateBgpPeerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateBgpPeerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateBgpPeerRequest) GetPeerIpAddress() *string {
	return s.PeerIpAddress
}

func (s *CreateBgpPeerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBgpPeerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateBgpPeerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateBgpPeerRequest) SetBfdMultiHop(v int32) *CreateBgpPeerRequest {
	s.BfdMultiHop = &v
	return s
}

func (s *CreateBgpPeerRequest) SetBgpGroupId(v string) *CreateBgpPeerRequest {
	s.BgpGroupId = &v
	return s
}

func (s *CreateBgpPeerRequest) SetClientToken(v string) *CreateBgpPeerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBgpPeerRequest) SetEnableBfd(v bool) *CreateBgpPeerRequest {
	s.EnableBfd = &v
	return s
}

func (s *CreateBgpPeerRequest) SetIpVersion(v string) *CreateBgpPeerRequest {
	s.IpVersion = &v
	return s
}

func (s *CreateBgpPeerRequest) SetOwnerAccount(v string) *CreateBgpPeerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateBgpPeerRequest) SetOwnerId(v int64) *CreateBgpPeerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateBgpPeerRequest) SetPeerIpAddress(v string) *CreateBgpPeerRequest {
	s.PeerIpAddress = &v
	return s
}

func (s *CreateBgpPeerRequest) SetRegionId(v string) *CreateBgpPeerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBgpPeerRequest) SetResourceOwnerAccount(v string) *CreateBgpPeerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateBgpPeerRequest) SetResourceOwnerId(v int64) *CreateBgpPeerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateBgpPeerRequest) Validate() error {
	return dara.Validate(s)
}
