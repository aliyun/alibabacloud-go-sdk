// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStoragePlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeStoragePlanRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeStoragePlanRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeStoragePlanRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeStoragePlanRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeStoragePlanRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeStoragePlanRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeStoragePlanRequest
	GetResourceOwnerId() *int64
}

type DescribeStoragePlanRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be a positive integer that does not exceed the maximum value of the Integer data type.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeStoragePlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStoragePlanRequest) GoString() string {
	return s.String()
}

func (s *DescribeStoragePlanRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeStoragePlanRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeStoragePlanRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeStoragePlanRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStoragePlanRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeStoragePlanRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeStoragePlanRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeStoragePlanRequest) SetOwnerAccount(v string) *DescribeStoragePlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeStoragePlanRequest) SetOwnerId(v int64) *DescribeStoragePlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStoragePlanRequest) SetPageNumber(v int32) *DescribeStoragePlanRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStoragePlanRequest) SetPageSize(v int32) *DescribeStoragePlanRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStoragePlanRequest) SetResourceGroupId(v string) *DescribeStoragePlanRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeStoragePlanRequest) SetResourceOwnerAccount(v string) *DescribeStoragePlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeStoragePlanRequest) SetResourceOwnerId(v int64) *DescribeStoragePlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeStoragePlanRequest) Validate() error {
	return dara.Validate(s)
}
