// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBgpGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBgpGroupId(v string) *DescribeBgpGroupsRequest
	GetBgpGroupId() *string
	SetIsDefault(v bool) *DescribeBgpGroupsRequest
	GetIsDefault() *bool
	SetOwnerAccount(v string) *DescribeBgpGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBgpGroupsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeBgpGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBgpGroupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeBgpGroupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeBgpGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBgpGroupsRequest
	GetResourceOwnerId() *int64
	SetRouterId(v string) *DescribeBgpGroupsRequest
	GetRouterId() *string
}

type DescribeBgpGroupsRequest struct {
	// The ID of the BGP group.
	//
	// example:
	//
	// bgpg-bp1k25cyp26cllath****
	BgpGroupId *string `json:"BgpGroupId,omitempty" xml:"BgpGroupId,omitempty"`
	// Specifies whether the BGP group is the default one. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// false
	IsDefault    *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The maximum value is **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the VBR is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the virtual border router (VBR) that is associated with the BGP group.
	//
	// example:
	//
	// vbr-bp1ctxy813985gkuk****
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
}

func (s DescribeBgpGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBgpGroupsRequest) GetBgpGroupId() *string {
	return s.BgpGroupId
}

func (s *DescribeBgpGroupsRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeBgpGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBgpGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBgpGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBgpGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBgpGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBgpGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBgpGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBgpGroupsRequest) GetRouterId() *string {
	return s.RouterId
}

func (s *DescribeBgpGroupsRequest) SetBgpGroupId(v string) *DescribeBgpGroupsRequest {
	s.BgpGroupId = &v
	return s
}

func (s *DescribeBgpGroupsRequest) SetIsDefault(v bool) *DescribeBgpGroupsRequest {
	s.IsDefault = &v
	return s
}

func (s *DescribeBgpGroupsRequest) SetOwnerAccount(v string) *DescribeBgpGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBgpGroupsRequest) SetOwnerId(v int64) *DescribeBgpGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBgpGroupsRequest) SetPageNumber(v int32) *DescribeBgpGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBgpGroupsRequest) SetPageSize(v int32) *DescribeBgpGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBgpGroupsRequest) SetRegionId(v string) *DescribeBgpGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBgpGroupsRequest) SetResourceOwnerAccount(v string) *DescribeBgpGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBgpGroupsRequest) SetResourceOwnerId(v int64) *DescribeBgpGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBgpGroupsRequest) SetRouterId(v string) *DescribeBgpGroupsRequest {
	s.RouterId = &v
	return s
}

func (s *DescribeBgpGroupsRequest) Validate() error {
	return dara.Validate(s)
}
