// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterIds(v string) *DescribeAutoRenewAttributeRequest
	GetDBClusterIds() *string
	SetOwnerAccount(v string) *DescribeAutoRenewAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAutoRenewAttributeRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAutoRenewAttributeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoRenewAttributeRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAutoRenewAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAutoRenewAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAutoRenewAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeAutoRenewAttributeRequest struct {
	// example:
	//
	// cc-uf6g4417bo*****
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeRequest) GetDBClusterIds() *string {
	return s.DBClusterIds
}

func (s *DescribeAutoRenewAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAutoRenewAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAutoRenewAttributeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoRenewAttributeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoRenewAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoRenewAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAutoRenewAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAutoRenewAttributeRequest) SetDBClusterIds(v string) *DescribeAutoRenewAttributeRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetOwnerAccount(v string) *DescribeAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetOwnerId(v int64) *DescribeAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetPageNumber(v int32) *DescribeAutoRenewAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetPageSize(v int32) *DescribeAutoRenewAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetRegionId(v string) *DescribeAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *DescribeAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *DescribeAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) Validate() error {
	return dara.Validate(s)
}
