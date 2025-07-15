// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEipAddressAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *ModifyEipAddressAttributeRequest
	GetAllocationId() *string
	SetBandwidth(v string) *ModifyEipAddressAttributeRequest
	GetBandwidth() *string
	SetDescription(v string) *ModifyEipAddressAttributeRequest
	GetDescription() *string
	SetName(v string) *ModifyEipAddressAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyEipAddressAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyEipAddressAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyEipAddressAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyEipAddressAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyEipAddressAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyEipAddressAttributeRequest struct {
	// The ID of the pay-as-you-go EIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-2zeerraiwb7uj6i0d****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The new maximum bandwidth of the EIP. Valid values:
	//
	// 	- **1*	- to **200*	- if the metering method is pay-by-data-transfer. Unit: Mbit/s.
	//
	// 	- **1*	- to **500*	- if the metering method is pay-by-bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The new description of the EIP.
	//
	// The description must be 2 to 256 characters in length and start with a letter. The description cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The new name of the EIP.
	//
	// The name must be 1 to 128 characters in length, and can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// Test123
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the EIP.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyEipAddressAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEipAddressAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyEipAddressAttributeRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *ModifyEipAddressAttributeRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *ModifyEipAddressAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyEipAddressAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyEipAddressAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyEipAddressAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyEipAddressAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyEipAddressAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyEipAddressAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyEipAddressAttributeRequest) SetAllocationId(v string) *ModifyEipAddressAttributeRequest {
	s.AllocationId = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) SetBandwidth(v string) *ModifyEipAddressAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) SetDescription(v string) *ModifyEipAddressAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) SetName(v string) *ModifyEipAddressAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) SetOwnerAccount(v string) *ModifyEipAddressAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) SetOwnerId(v int64) *ModifyEipAddressAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) SetRegionId(v string) *ModifyEipAddressAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) SetResourceOwnerAccount(v string) *ModifyEipAddressAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) SetResourceOwnerId(v int64) *ModifyEipAddressAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) Validate() error {
	return dara.Validate(s)
}
