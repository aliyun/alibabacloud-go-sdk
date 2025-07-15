// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGlobalAccelerationInstanceIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalAccelerationInstanceId(v string) *AddGlobalAccelerationInstanceIpRequest
	GetGlobalAccelerationInstanceId() *string
	SetIpInstanceId(v string) *AddGlobalAccelerationInstanceIpRequest
	GetIpInstanceId() *string
	SetOwnerAccount(v string) *AddGlobalAccelerationInstanceIpRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddGlobalAccelerationInstanceIpRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddGlobalAccelerationInstanceIpRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddGlobalAccelerationInstanceIpRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddGlobalAccelerationInstanceIpRequest
	GetResourceOwnerId() *int64
}

type AddGlobalAccelerationInstanceIpRequest struct {
	// The ID of the shared-bandwidth GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-Ldefrgbttnyyf****
	GlobalAccelerationInstanceId *string `json:"GlobalAccelerationInstanceId,omitempty" xml:"GlobalAccelerationInstanceId,omitempty"`
	// The EIP ID. You can call the [DescribeEipAddresses](https://help.aliyun.com/document_detail/36018.html) operation to query EIP IDs.
	//
	// >  Make sure that the billing method of the EIP is pay-as-you-go, and the EIP and the shared-bandwidth GA instance belong to the same region.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-rw434rwfdeaf****
	IpInstanceId *string `json:"IpInstanceId,omitempty" xml:"IpInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region of the shared-bandwidth GA instance.
	//
	// You can call the **DescribeRegions*	- operation to query the most recent region list.
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

func (s AddGlobalAccelerationInstanceIpRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGlobalAccelerationInstanceIpRequest) GoString() string {
	return s.String()
}

func (s *AddGlobalAccelerationInstanceIpRequest) GetGlobalAccelerationInstanceId() *string {
	return s.GlobalAccelerationInstanceId
}

func (s *AddGlobalAccelerationInstanceIpRequest) GetIpInstanceId() *string {
	return s.IpInstanceId
}

func (s *AddGlobalAccelerationInstanceIpRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddGlobalAccelerationInstanceIpRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddGlobalAccelerationInstanceIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddGlobalAccelerationInstanceIpRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddGlobalAccelerationInstanceIpRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddGlobalAccelerationInstanceIpRequest) SetGlobalAccelerationInstanceId(v string) *AddGlobalAccelerationInstanceIpRequest {
	s.GlobalAccelerationInstanceId = &v
	return s
}

func (s *AddGlobalAccelerationInstanceIpRequest) SetIpInstanceId(v string) *AddGlobalAccelerationInstanceIpRequest {
	s.IpInstanceId = &v
	return s
}

func (s *AddGlobalAccelerationInstanceIpRequest) SetOwnerAccount(v string) *AddGlobalAccelerationInstanceIpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddGlobalAccelerationInstanceIpRequest) SetOwnerId(v int64) *AddGlobalAccelerationInstanceIpRequest {
	s.OwnerId = &v
	return s
}

func (s *AddGlobalAccelerationInstanceIpRequest) SetRegionId(v string) *AddGlobalAccelerationInstanceIpRequest {
	s.RegionId = &v
	return s
}

func (s *AddGlobalAccelerationInstanceIpRequest) SetResourceOwnerAccount(v string) *AddGlobalAccelerationInstanceIpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddGlobalAccelerationInstanceIpRequest) SetResourceOwnerId(v int64) *AddGlobalAccelerationInstanceIpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddGlobalAccelerationInstanceIpRequest) Validate() error {
	return dara.Validate(s)
}
