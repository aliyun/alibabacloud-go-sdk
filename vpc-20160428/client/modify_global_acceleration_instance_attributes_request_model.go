// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalAccelerationInstanceAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyGlobalAccelerationInstanceAttributesRequest
	GetDescription() *string
	SetGlobalAccelerationInstanceId(v string) *ModifyGlobalAccelerationInstanceAttributesRequest
	GetGlobalAccelerationInstanceId() *string
	SetName(v string) *ModifyGlobalAccelerationInstanceAttributesRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyGlobalAccelerationInstanceAttributesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyGlobalAccelerationInstanceAttributesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyGlobalAccelerationInstanceAttributesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyGlobalAccelerationInstanceAttributesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyGlobalAccelerationInstanceAttributesRequest
	GetResourceOwnerId() *int64
}

type ModifyGlobalAccelerationInstanceAttributesRequest struct {
	// The description of the GA instance.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// My GA
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-14fdsf3****
	GlobalAccelerationInstanceId *string `json:"GlobalAccelerationInstanceId,omitempty" xml:"GlobalAccelerationInstanceId,omitempty"`
	// The name of the GA instance.
	//
	// The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// GA-1
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the GA instance.
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

func (s ModifyGlobalAccelerationInstanceAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalAccelerationInstanceAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) GetGlobalAccelerationInstanceId() *string {
	return s.GlobalAccelerationInstanceId
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) GetName() *string {
	return s.Name
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) SetDescription(v string) *ModifyGlobalAccelerationInstanceAttributesRequest {
	s.Description = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) SetGlobalAccelerationInstanceId(v string) *ModifyGlobalAccelerationInstanceAttributesRequest {
	s.GlobalAccelerationInstanceId = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) SetName(v string) *ModifyGlobalAccelerationInstanceAttributesRequest {
	s.Name = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) SetOwnerAccount(v string) *ModifyGlobalAccelerationInstanceAttributesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) SetOwnerId(v int64) *ModifyGlobalAccelerationInstanceAttributesRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) SetRegionId(v string) *ModifyGlobalAccelerationInstanceAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) SetResourceOwnerAccount(v string) *ModifyGlobalAccelerationInstanceAttributesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) SetResourceOwnerId(v int64) *ModifyGlobalAccelerationInstanceAttributesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyGlobalAccelerationInstanceAttributesRequest) Validate() error {
	return dara.Validate(s)
}
