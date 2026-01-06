// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoRenewalAttributeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterIdShrink(v string) *DescribeAutoRenewalAttributeShrinkRequest
	GetDBClusterIdShrink() *string
	SetOwnerAccount(v string) *DescribeAutoRenewalAttributeShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAutoRenewalAttributeShrinkRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAutoRenewalAttributeShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoRenewalAttributeShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAutoRenewalAttributeShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeAutoRenewalAttributeShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeAutoRenewalAttributeShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAutoRenewalAttributeShrinkRequest
	GetResourceOwnerId() *int64
}

type DescribeAutoRenewalAttributeShrinkRequest struct {
	// The cluster IDs.
	DBClusterIdShrink *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. A positive integer greater than 0 and not exceeding the maximum value of the Integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page. Valid values:
	//
	// 	- **30**(Default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAutoRenewalAttributeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRenewalAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) GetDBClusterIdShrink() *string {
	return s.DBClusterIdShrink
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) SetDBClusterIdShrink(v string) *DescribeAutoRenewalAttributeShrinkRequest {
	s.DBClusterIdShrink = &v
	return s
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) SetOwnerAccount(v string) *DescribeAutoRenewalAttributeShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) SetOwnerId(v int64) *DescribeAutoRenewalAttributeShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) SetPageNumber(v int32) *DescribeAutoRenewalAttributeShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) SetPageSize(v int32) *DescribeAutoRenewalAttributeShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) SetRegionId(v string) *DescribeAutoRenewalAttributeShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) SetResourceGroupId(v string) *DescribeAutoRenewalAttributeShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) SetResourceOwnerAccount(v string) *DescribeAutoRenewalAttributeShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) SetResourceOwnerId(v int64) *DescribeAutoRenewalAttributeShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAutoRenewalAttributeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
