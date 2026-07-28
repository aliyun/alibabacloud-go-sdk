// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBgpPeerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBfdMultiHop(v int32) *ModifyBgpPeerAttributeRequest
	GetBfdMultiHop() *int32
	SetBgpGroupId(v string) *ModifyBgpPeerAttributeRequest
	GetBgpGroupId() *string
	SetBgpPeerId(v string) *ModifyBgpPeerAttributeRequest
	GetBgpPeerId() *string
	SetClientToken(v string) *ModifyBgpPeerAttributeRequest
	GetClientToken() *string
	SetEnableBfd(v bool) *ModifyBgpPeerAttributeRequest
	GetEnableBfd() *bool
	SetOwnerAccount(v string) *ModifyBgpPeerAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyBgpPeerAttributeRequest
	GetOwnerId() *int64
	SetPeerIpAddress(v string) *ModifyBgpPeerAttributeRequest
	GetPeerIpAddress() *string
	SetRegionId(v string) *ModifyBgpPeerAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyBgpPeerAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyBgpPeerAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyBgpPeerAttributeRequest struct {
	// The BFD hop count. Valid values: **1*	- to **255**.
	//
	// This parameter is required when BFD is enabled.
	//
	// Enter the BFD hop count, which specifies the maximum number of devices that data passes through from the source to the destination. You can configure different hop counts based on the actual physical link conditions.
	//
	// > If you use BFD in a multi-cloud environment or a direct fiber connection topology where no bridging devices exist, change the default BFD hop count from **255*	- to **1**.
	//
	// example:
	//
	// 3
	BfdMultiHop *int32 `json:"BfdMultiHop,omitempty" xml:"BfdMultiHop,omitempty"`
	// The ID of the BGP group to which the BGP peer belongs.
	//
	// example:
	//
	// bgpg-m5eo12jxuw2hc0uqq****
	BgpGroupId *string `json:"BgpGroupId,omitempty" xml:"BgpGroupId,omitempty"`
	// The ID of the BGP peer whose attributes you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// bgp-m5eoyp2mwegk8ce9v****
	BgpPeerId *string `json:"BgpPeerId,omitempty" xml:"BgpPeerId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// The client generates the value. The value must be unique among different requests and cannot exceed 64 ASCII characters in length.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable Bidirectional Forwarding Detection (BFD). Valid values:
	//
	// - **true**: enables BFD.
	//
	// - **false*	- (default): does not enable BFD.
	//
	// example:
	//
	// false
	EnableBfd    *bool   `json:"EnableBfd,omitempty" xml:"EnableBfd,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The IP address of the BGP peer.
	//
	// example:
	//
	// 116.62.XX.XX
	PeerIpAddress *string `json:"PeerIpAddress,omitempty" xml:"PeerIpAddress,omitempty"`
	// The region ID of the BGP group to which the BGP peer belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
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

func (s ModifyBgpPeerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBgpPeerAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyBgpPeerAttributeRequest) GetBfdMultiHop() *int32 {
	return s.BfdMultiHop
}

func (s *ModifyBgpPeerAttributeRequest) GetBgpGroupId() *string {
	return s.BgpGroupId
}

func (s *ModifyBgpPeerAttributeRequest) GetBgpPeerId() *string {
	return s.BgpPeerId
}

func (s *ModifyBgpPeerAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyBgpPeerAttributeRequest) GetEnableBfd() *bool {
	return s.EnableBfd
}

func (s *ModifyBgpPeerAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyBgpPeerAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyBgpPeerAttributeRequest) GetPeerIpAddress() *string {
	return s.PeerIpAddress
}

func (s *ModifyBgpPeerAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyBgpPeerAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyBgpPeerAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyBgpPeerAttributeRequest) SetBfdMultiHop(v int32) *ModifyBgpPeerAttributeRequest {
	s.BfdMultiHop = &v
	return s
}

func (s *ModifyBgpPeerAttributeRequest) SetBgpGroupId(v string) *ModifyBgpPeerAttributeRequest {
	s.BgpGroupId = &v
	return s
}

func (s *ModifyBgpPeerAttributeRequest) SetBgpPeerId(v string) *ModifyBgpPeerAttributeRequest {
	s.BgpPeerId = &v
	return s
}

func (s *ModifyBgpPeerAttributeRequest) SetClientToken(v string) *ModifyBgpPeerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyBgpPeerAttributeRequest) SetEnableBfd(v bool) *ModifyBgpPeerAttributeRequest {
	s.EnableBfd = &v
	return s
}

func (s *ModifyBgpPeerAttributeRequest) SetOwnerAccount(v string) *ModifyBgpPeerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBgpPeerAttributeRequest) SetOwnerId(v int64) *ModifyBgpPeerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBgpPeerAttributeRequest) SetPeerIpAddress(v string) *ModifyBgpPeerAttributeRequest {
	s.PeerIpAddress = &v
	return s
}

func (s *ModifyBgpPeerAttributeRequest) SetRegionId(v string) *ModifyBgpPeerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBgpPeerAttributeRequest) SetResourceOwnerAccount(v string) *ModifyBgpPeerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBgpPeerAttributeRequest) SetResourceOwnerId(v int64) *ModifyBgpPeerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBgpPeerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
