// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudProvider(v string) *DescribeAutoRenewAttributeRequest
	GetCloudProvider() *string
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
	SetResourceGroupId(v string) *DescribeAutoRenewAttributeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeAutoRenewAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAutoRenewAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeAutoRenewAttributeRequest struct {
	// example:
	//
	// ENS
	CloudProvider *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// The ID of the cluster. If you need to specify multiple cluster IDs, separate the cluster IDs with commas (,).
	//
	// example:
	//
	// pc-****************
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 30, 50, and 100. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the region ID details.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-re*********
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeRequest) GetCloudProvider() *string {
	return s.CloudProvider
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

func (s *DescribeAutoRenewAttributeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAutoRenewAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAutoRenewAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAutoRenewAttributeRequest) SetCloudProvider(v string) *DescribeAutoRenewAttributeRequest {
	s.CloudProvider = &v
	return s
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

func (s *DescribeAutoRenewAttributeRequest) SetResourceGroupId(v string) *DescribeAutoRenewAttributeRequest {
	s.ResourceGroupId = &v
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
