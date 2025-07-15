// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBgpPeerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBgpPeerId(v string) *DeleteBgpPeerRequest
	GetBgpPeerId() *string
	SetClientToken(v string) *DeleteBgpPeerRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteBgpPeerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteBgpPeerRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteBgpPeerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteBgpPeerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteBgpPeerRequest
	GetResourceOwnerId() *int64
}

type DeleteBgpPeerRequest struct {
	// The ID of the BGP peer.
	//
	// This parameter is required.
	//
	// example:
	//
	// bgp-wz977wcrmb69a********
	BgpPeerId *string `json:"BgpPeerId,omitempty" xml:"BgpPeerId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the BGP group.
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

func (s DeleteBgpPeerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBgpPeerRequest) GoString() string {
	return s.String()
}

func (s *DeleteBgpPeerRequest) GetBgpPeerId() *string {
	return s.BgpPeerId
}

func (s *DeleteBgpPeerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteBgpPeerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteBgpPeerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteBgpPeerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBgpPeerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteBgpPeerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteBgpPeerRequest) SetBgpPeerId(v string) *DeleteBgpPeerRequest {
	s.BgpPeerId = &v
	return s
}

func (s *DeleteBgpPeerRequest) SetClientToken(v string) *DeleteBgpPeerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBgpPeerRequest) SetOwnerAccount(v string) *DeleteBgpPeerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteBgpPeerRequest) SetOwnerId(v int64) *DeleteBgpPeerRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteBgpPeerRequest) SetRegionId(v string) *DeleteBgpPeerRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBgpPeerRequest) SetResourceOwnerAccount(v string) *DeleteBgpPeerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteBgpPeerRequest) SetResourceOwnerId(v int64) *DeleteBgpPeerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteBgpPeerRequest) Validate() error {
	return dara.Validate(s)
}
