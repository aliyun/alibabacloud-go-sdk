// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHaVipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteHaVipRequest
	GetClientToken() *string
	SetHaVipId(v string) *DeleteHaVipRequest
	GetHaVipId() *string
	SetOwnerAccount(v string) *DeleteHaVipRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteHaVipRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteHaVipRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteHaVipRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteHaVipRequest
	GetResourceOwnerId() *int64
}

type DeleteHaVipRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId*	- as **ClientToken**. The value of **RequestId*	- in each API request may be different.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the HaVip that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// havip-2zeo05qre24nhrqpy****
	HaVipId      *string `json:"HaVipId,omitempty" xml:"HaVipId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the HaVip resides. Call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region list.
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

func (s DeleteHaVipRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHaVipRequest) GoString() string {
	return s.String()
}

func (s *DeleteHaVipRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteHaVipRequest) GetHaVipId() *string {
	return s.HaVipId
}

func (s *DeleteHaVipRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteHaVipRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteHaVipRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHaVipRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteHaVipRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteHaVipRequest) SetClientToken(v string) *DeleteHaVipRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteHaVipRequest) SetHaVipId(v string) *DeleteHaVipRequest {
	s.HaVipId = &v
	return s
}

func (s *DeleteHaVipRequest) SetOwnerAccount(v string) *DeleteHaVipRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteHaVipRequest) SetOwnerId(v int64) *DeleteHaVipRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteHaVipRequest) SetRegionId(v string) *DeleteHaVipRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHaVipRequest) SetResourceOwnerAccount(v string) *DeleteHaVipRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteHaVipRequest) SetResourceOwnerId(v int64) *DeleteHaVipRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteHaVipRequest) Validate() error {
	return dara.Validate(s)
}
