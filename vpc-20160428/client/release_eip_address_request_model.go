// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseEipAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *ReleaseEipAddressRequest
	GetAllocationId() *string
	SetOwnerAccount(v string) *ReleaseEipAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReleaseEipAddressRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ReleaseEipAddressRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ReleaseEipAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleaseEipAddressRequest
	GetResourceOwnerId() *int64
}

type ReleaseEipAddressRequest struct {
	// The ID of the EIP that you want to release.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-2zeerraiwb7uj6i0d****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the EIP belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleaseEipAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseEipAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleaseEipAddressRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *ReleaseEipAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReleaseEipAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseEipAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleaseEipAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleaseEipAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseEipAddressRequest) SetAllocationId(v string) *ReleaseEipAddressRequest {
	s.AllocationId = &v
	return s
}

func (s *ReleaseEipAddressRequest) SetOwnerAccount(v string) *ReleaseEipAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseEipAddressRequest) SetOwnerId(v int64) *ReleaseEipAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseEipAddressRequest) SetRegionId(v string) *ReleaseEipAddressRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseEipAddressRequest) SetResourceOwnerAccount(v string) *ReleaseEipAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseEipAddressRequest) SetResourceOwnerId(v int64) *ReleaseEipAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseEipAddressRequest) Validate() error {
	return dara.Validate(s)
}
