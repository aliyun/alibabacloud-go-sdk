// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostGroupId(v string) *DescribeDedicatedHostGroupsRequest
	GetDedicatedHostGroupId() *string
	SetImageCategory(v string) *DescribeDedicatedHostGroupsRequest
	GetImageCategory() *string
	SetOwnerId(v int64) *DescribeDedicatedHostGroupsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDedicatedHostGroupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDedicatedHostGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDedicatedHostGroupsRequest
	GetResourceOwnerId() *int64
}

type DescribeDedicatedHostGroupsRequest struct {
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	ImageCategory        *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDedicatedHostGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostGroupsRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *DescribeDedicatedHostGroupsRequest) GetImageCategory() *string {
	return s.ImageCategory
}

func (s *DescribeDedicatedHostGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDedicatedHostGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedHostGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDedicatedHostGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDedicatedHostGroupsRequest) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostGroupsRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) SetImageCategory(v string) *DescribeDedicatedHostGroupsRequest {
	s.ImageCategory = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) SetOwnerId(v int64) *DescribeDedicatedHostGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) SetRegionId(v string) *DescribeDedicatedHostGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedHostGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) SetResourceOwnerId(v int64) *DescribeDedicatedHostGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostGroupsRequest) Validate() error {
	return dara.Validate(s)
}
