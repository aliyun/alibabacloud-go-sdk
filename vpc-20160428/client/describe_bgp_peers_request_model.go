// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBgpPeersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBgpGroupId(v string) *DescribeBgpPeersRequest
	GetBgpGroupId() *string
	SetBgpPeerId(v string) *DescribeBgpPeersRequest
	GetBgpPeerId() *string
	SetIsDefault(v bool) *DescribeBgpPeersRequest
	GetIsDefault() *bool
	SetOwnerAccount(v string) *DescribeBgpPeersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBgpPeersRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeBgpPeersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBgpPeersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeBgpPeersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeBgpPeersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBgpPeersRequest
	GetResourceOwnerId() *int64
	SetRouterId(v string) *DescribeBgpPeersRequest
	GetRouterId() *string
}

type DescribeBgpPeersRequest struct {
	// The ID of the BGP group to which the BGP peer that you want to query belongs.
	//
	// example:
	//
	// bgpg-2zev8h2wo414sfh****
	BgpGroupId *string `json:"BgpGroupId,omitempty" xml:"BgpGroupId,omitempty"`
	// The ID of the BGP peer that you want to query.
	//
	// example:
	//
	// bgp-2ze3un0ft1jd1xd****
	BgpPeerId *string `json:"BgpPeerId,omitempty" xml:"BgpPeerId,omitempty"`
	// Specifies whether the BGP group is the default group. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
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
	// The number of entries per page. Valid values: **1 to 50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the BGP group to which the BGP peer that you want to query belongs.
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
	// The ID of the virtual border router (VBR) that is associated with the BGP peer that you want to query.
	//
	// example:
	//
	// vbr-2zecmmvg5gvu8i4te****
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
}

func (s DescribeBgpPeersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpPeersRequest) GoString() string {
	return s.String()
}

func (s *DescribeBgpPeersRequest) GetBgpGroupId() *string {
	return s.BgpGroupId
}

func (s *DescribeBgpPeersRequest) GetBgpPeerId() *string {
	return s.BgpPeerId
}

func (s *DescribeBgpPeersRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeBgpPeersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBgpPeersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBgpPeersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBgpPeersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBgpPeersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBgpPeersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBgpPeersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBgpPeersRequest) GetRouterId() *string {
	return s.RouterId
}

func (s *DescribeBgpPeersRequest) SetBgpGroupId(v string) *DescribeBgpPeersRequest {
	s.BgpGroupId = &v
	return s
}

func (s *DescribeBgpPeersRequest) SetBgpPeerId(v string) *DescribeBgpPeersRequest {
	s.BgpPeerId = &v
	return s
}

func (s *DescribeBgpPeersRequest) SetIsDefault(v bool) *DescribeBgpPeersRequest {
	s.IsDefault = &v
	return s
}

func (s *DescribeBgpPeersRequest) SetOwnerAccount(v string) *DescribeBgpPeersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBgpPeersRequest) SetOwnerId(v int64) *DescribeBgpPeersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBgpPeersRequest) SetPageNumber(v int32) *DescribeBgpPeersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBgpPeersRequest) SetPageSize(v int32) *DescribeBgpPeersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBgpPeersRequest) SetRegionId(v string) *DescribeBgpPeersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBgpPeersRequest) SetResourceOwnerAccount(v string) *DescribeBgpPeersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBgpPeersRequest) SetResourceOwnerId(v int64) *DescribeBgpPeersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBgpPeersRequest) SetRouterId(v string) *DescribeBgpPeersRequest {
	s.RouterId = &v
	return s
}

func (s *DescribeBgpPeersRequest) Validate() error {
	return dara.Validate(s)
}
