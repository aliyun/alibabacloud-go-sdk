// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVSwitchCidrReservationAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyVSwitchCidrReservationAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyVSwitchCidrReservationAttributeRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ModifyVSwitchCidrReservationAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVSwitchCidrReservationAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyVSwitchCidrReservationAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyVSwitchCidrReservationAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVSwitchCidrReservationAttributeRequest
	GetResourceOwnerId() *int64
	SetVSwitchCidrReservationDescription(v string) *ModifyVSwitchCidrReservationAttributeRequest
	GetVSwitchCidrReservationDescription() *string
	SetVSwitchCidrReservationId(v string) *ModifyVSwitchCidrReservationAttributeRequest
	GetVSwitchCidrReservationId() *string
	SetVSwitchCidrReservationName(v string) *ModifyVSwitchCidrReservationAttributeRequest
	GetVSwitchCidrReservationName() *string
}

type ModifyVSwitchCidrReservationAttributeRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the vSwitch is deployed.
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
	// The new description of the reserved CIDR block. The default value is empty.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// ReservationDescription
	VSwitchCidrReservationDescription *string `json:"VSwitchCidrReservationDescription,omitempty" xml:"VSwitchCidrReservationDescription,omitempty"`
	// The ID of the reserved CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// vcr-bp1m12saqteraw3rp****
	VSwitchCidrReservationId *string `json:"VSwitchCidrReservationId,omitempty" xml:"VSwitchCidrReservationId,omitempty"`
	// The new name of the reserved CIDR block.
	//
	// The name must be 2 to 128 characters in length and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// ReservationName
	VSwitchCidrReservationName *string `json:"VSwitchCidrReservationName,omitempty" xml:"VSwitchCidrReservationName,omitempty"`
}

func (s ModifyVSwitchCidrReservationAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVSwitchCidrReservationAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) GetVSwitchCidrReservationDescription() *string {
	return s.VSwitchCidrReservationDescription
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) GetVSwitchCidrReservationId() *string {
	return s.VSwitchCidrReservationId
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) GetVSwitchCidrReservationName() *string {
	return s.VSwitchCidrReservationName
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) SetClientToken(v string) *ModifyVSwitchCidrReservationAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) SetDryRun(v bool) *ModifyVSwitchCidrReservationAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) SetOwnerAccount(v string) *ModifyVSwitchCidrReservationAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) SetOwnerId(v int64) *ModifyVSwitchCidrReservationAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) SetRegionId(v string) *ModifyVSwitchCidrReservationAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) SetResourceOwnerAccount(v string) *ModifyVSwitchCidrReservationAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) SetResourceOwnerId(v int64) *ModifyVSwitchCidrReservationAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) SetVSwitchCidrReservationDescription(v string) *ModifyVSwitchCidrReservationAttributeRequest {
	s.VSwitchCidrReservationDescription = &v
	return s
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) SetVSwitchCidrReservationId(v string) *ModifyVSwitchCidrReservationAttributeRequest {
	s.VSwitchCidrReservationId = &v
	return s
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) SetVSwitchCidrReservationName(v string) *ModifyVSwitchCidrReservationAttributeRequest {
	s.VSwitchCidrReservationName = &v
	return s
}

func (s *ModifyVSwitchCidrReservationAttributeRequest) Validate() error {
	return dara.Validate(s)
}
