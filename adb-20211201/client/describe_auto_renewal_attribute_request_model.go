// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoRenewalAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v []*string) *DescribeAutoRenewalAttributeRequest
	GetDBClusterId() []*string
	SetOwnerAccount(v string) *DescribeAutoRenewalAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAutoRenewalAttributeRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAutoRenewalAttributeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoRenewalAttributeRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAutoRenewalAttributeRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeAutoRenewalAttributeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeAutoRenewalAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAutoRenewalAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeAutoRenewalAttributeRequest struct {
	DBClusterId  []*string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAutoRenewalAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRenewalAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewalAttributeRequest) GetDBClusterId() []*string {
	return s.DBClusterId
}

func (s *DescribeAutoRenewalAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAutoRenewalAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAutoRenewalAttributeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoRenewalAttributeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoRenewalAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoRenewalAttributeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAutoRenewalAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAutoRenewalAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAutoRenewalAttributeRequest) SetDBClusterId(v []*string) *DescribeAutoRenewalAttributeRequest {
	s.DBClusterId = v
	return s
}

func (s *DescribeAutoRenewalAttributeRequest) SetOwnerAccount(v string) *DescribeAutoRenewalAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAutoRenewalAttributeRequest) SetOwnerId(v int64) *DescribeAutoRenewalAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoRenewalAttributeRequest) SetPageNumber(v int32) *DescribeAutoRenewalAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoRenewalAttributeRequest) SetPageSize(v int32) *DescribeAutoRenewalAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoRenewalAttributeRequest) SetRegionId(v string) *DescribeAutoRenewalAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoRenewalAttributeRequest) SetResourceGroupId(v string) *DescribeAutoRenewalAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAutoRenewalAttributeRequest) SetResourceOwnerAccount(v string) *DescribeAutoRenewalAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAutoRenewalAttributeRequest) SetResourceOwnerId(v int64) *DescribeAutoRenewalAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAutoRenewalAttributeRequest) Validate() error {
	return dara.Validate(s)
}
