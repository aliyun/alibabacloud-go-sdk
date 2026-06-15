// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHaVipAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyHaVipAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyHaVipAttributeRequest
	GetDescription() *string
	SetHaVipId(v string) *ModifyHaVipAttributeRequest
	GetHaVipId() *string
	SetOwnerAccount(v string) *ModifyHaVipAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyHaVipAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyHaVipAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyHaVipAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyHaVipAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyHaVipAttributeRequest struct {
	// A client-generated token used to ensure the idempotence of the request. Make sure that the token is unique across requests. The token can contain only ASCII characters and be no more than 64 characters long.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new description of the HaVip. The description must be 2 to 256 characters long and must not start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the HaVip that you want to modify.
	//
	// This parameter is required.
	HaVipId      *string `json:"HaVipId,omitempty" xml:"HaVipId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the HaVip is deployed. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the latest list of regions.
	//
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyHaVipAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHaVipAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyHaVipAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyHaVipAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyHaVipAttributeRequest) GetHaVipId() *string {
	return s.HaVipId
}

func (s *ModifyHaVipAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyHaVipAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyHaVipAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHaVipAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyHaVipAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyHaVipAttributeRequest) SetClientToken(v string) *ModifyHaVipAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyHaVipAttributeRequest) SetDescription(v string) *ModifyHaVipAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyHaVipAttributeRequest) SetHaVipId(v string) *ModifyHaVipAttributeRequest {
	s.HaVipId = &v
	return s
}

func (s *ModifyHaVipAttributeRequest) SetOwnerAccount(v string) *ModifyHaVipAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyHaVipAttributeRequest) SetOwnerId(v int64) *ModifyHaVipAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyHaVipAttributeRequest) SetRegionId(v string) *ModifyHaVipAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHaVipAttributeRequest) SetResourceOwnerAccount(v string) *ModifyHaVipAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyHaVipAttributeRequest) SetResourceOwnerId(v int64) *ModifyHaVipAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyHaVipAttributeRequest) Validate() error {
	return dara.Validate(s)
}
