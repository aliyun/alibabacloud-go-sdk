// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVRouterAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyVRouterAttributeRequest
	GetDescription() *string
	SetOwnerAccount(v string) *ModifyVRouterAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVRouterAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyVRouterAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyVRouterAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVRouterAttributeRequest
	GetResourceOwnerId() *int64
	SetVRouterId(v string) *ModifyVRouterAttributeRequest
	GetVRouterId() *string
	SetVRouterName(v string) *ModifyVRouterAttributeRequest
	GetVRouterName() *string
}

type ModifyVRouterAttributeRequest struct {
	// The description of the vRouter.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// My VRouter
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the vRouter.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The vRouter ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vrt-gw0p7zigake75znix****
	VRouterId *string `json:"VRouterId,omitempty" xml:"VRouterId,omitempty"`
	// The name of the vRouter.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// VRouter-1
	VRouterName *string `json:"VRouterName,omitempty" xml:"VRouterName,omitempty"`
}

func (s ModifyVRouterAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVRouterAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyVRouterAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyVRouterAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVRouterAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVRouterAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVRouterAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVRouterAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVRouterAttributeRequest) GetVRouterId() *string {
	return s.VRouterId
}

func (s *ModifyVRouterAttributeRequest) GetVRouterName() *string {
	return s.VRouterName
}

func (s *ModifyVRouterAttributeRequest) SetDescription(v string) *ModifyVRouterAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyVRouterAttributeRequest) SetOwnerAccount(v string) *ModifyVRouterAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVRouterAttributeRequest) SetOwnerId(v int64) *ModifyVRouterAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVRouterAttributeRequest) SetRegionId(v string) *ModifyVRouterAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVRouterAttributeRequest) SetResourceOwnerAccount(v string) *ModifyVRouterAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVRouterAttributeRequest) SetResourceOwnerId(v int64) *ModifyVRouterAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVRouterAttributeRequest) SetVRouterId(v string) *ModifyVRouterAttributeRequest {
	s.VRouterId = &v
	return s
}

func (s *ModifyVRouterAttributeRequest) SetVRouterName(v string) *ModifyVRouterAttributeRequest {
	s.VRouterName = &v
	return s
}

func (s *ModifyVRouterAttributeRequest) Validate() error {
	return dara.Validate(s)
}
