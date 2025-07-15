// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBgpGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBgpGroupId(v string) *DeleteBgpGroupRequest
	GetBgpGroupId() *string
	SetClientToken(v string) *DeleteBgpGroupRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteBgpGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteBgpGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteBgpGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteBgpGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteBgpGroupRequest
	GetResourceOwnerId() *int64
}

type DeleteBgpGroupRequest struct {
	// The ID of the BGP group.
	//
	// This parameter is required.
	//
	// example:
	//
	// bgpg-bp1k25cyp26cllath****
	BgpGroupId *string `json:"BgpGroupId,omitempty" xml:"BgpGroupId,omitempty"`
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
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteBgpGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBgpGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteBgpGroupRequest) GetBgpGroupId() *string {
	return s.BgpGroupId
}

func (s *DeleteBgpGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteBgpGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteBgpGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteBgpGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBgpGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteBgpGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteBgpGroupRequest) SetBgpGroupId(v string) *DeleteBgpGroupRequest {
	s.BgpGroupId = &v
	return s
}

func (s *DeleteBgpGroupRequest) SetClientToken(v string) *DeleteBgpGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBgpGroupRequest) SetOwnerAccount(v string) *DeleteBgpGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteBgpGroupRequest) SetOwnerId(v int64) *DeleteBgpGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteBgpGroupRequest) SetRegionId(v string) *DeleteBgpGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBgpGroupRequest) SetResourceOwnerAccount(v string) *DeleteBgpGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteBgpGroupRequest) SetResourceOwnerId(v int64) *DeleteBgpGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteBgpGroupRequest) Validate() error {
	return dara.Validate(s)
}
