// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveGlobalAccelerationInstanceIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalAccelerationInstanceId(v string) *RemoveGlobalAccelerationInstanceIpRequest
	GetGlobalAccelerationInstanceId() *string
	SetIpInstanceId(v string) *RemoveGlobalAccelerationInstanceIpRequest
	GetIpInstanceId() *string
	SetOwnerAccount(v string) *RemoveGlobalAccelerationInstanceIpRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveGlobalAccelerationInstanceIpRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveGlobalAccelerationInstanceIpRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RemoveGlobalAccelerationInstanceIpRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveGlobalAccelerationInstanceIpRequest
	GetResourceOwnerId() *int64
}

type RemoveGlobalAccelerationInstanceIpRequest struct {
	// The ID of the shared-bandwidth instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-m5ex47zwya1sejyni****
	GlobalAccelerationInstanceId *string `json:"GlobalAccelerationInstanceId,omitempty" xml:"GlobalAccelerationInstanceId,omitempty"`
	// The ID of the EIP.
	//
	// To query the EIP ID, call DescribeEipAddresses.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-bp13e9i2qst4g6jzi****
	IpInstanceId *string `json:"IpInstanceId,omitempty" xml:"IpInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the shared-bandwidth instance is located.
	//
	// To query the region ID, call DescribeRegions.
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

func (s RemoveGlobalAccelerationInstanceIpRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveGlobalAccelerationInstanceIpRequest) GoString() string {
	return s.String()
}

func (s *RemoveGlobalAccelerationInstanceIpRequest) GetGlobalAccelerationInstanceId() *string {
	return s.GlobalAccelerationInstanceId
}

func (s *RemoveGlobalAccelerationInstanceIpRequest) GetIpInstanceId() *string {
	return s.IpInstanceId
}

func (s *RemoveGlobalAccelerationInstanceIpRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveGlobalAccelerationInstanceIpRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveGlobalAccelerationInstanceIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveGlobalAccelerationInstanceIpRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveGlobalAccelerationInstanceIpRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveGlobalAccelerationInstanceIpRequest) SetGlobalAccelerationInstanceId(v string) *RemoveGlobalAccelerationInstanceIpRequest {
	s.GlobalAccelerationInstanceId = &v
	return s
}

func (s *RemoveGlobalAccelerationInstanceIpRequest) SetIpInstanceId(v string) *RemoveGlobalAccelerationInstanceIpRequest {
	s.IpInstanceId = &v
	return s
}

func (s *RemoveGlobalAccelerationInstanceIpRequest) SetOwnerAccount(v string) *RemoveGlobalAccelerationInstanceIpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveGlobalAccelerationInstanceIpRequest) SetOwnerId(v int64) *RemoveGlobalAccelerationInstanceIpRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveGlobalAccelerationInstanceIpRequest) SetRegionId(v string) *RemoveGlobalAccelerationInstanceIpRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveGlobalAccelerationInstanceIpRequest) SetResourceOwnerAccount(v string) *RemoveGlobalAccelerationInstanceIpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveGlobalAccelerationInstanceIpRequest) SetResourceOwnerId(v int64) *RemoveGlobalAccelerationInstanceIpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveGlobalAccelerationInstanceIpRequest) Validate() error {
	return dara.Validate(s)
}
