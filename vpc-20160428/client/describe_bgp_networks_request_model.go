// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBgpNetworksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeBgpNetworksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBgpNetworksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeBgpNetworksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBgpNetworksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeBgpNetworksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeBgpNetworksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBgpNetworksRequest
	GetResourceOwnerId() *int64
	SetRouterId(v string) *DescribeBgpNetworksRequest
	GetRouterId() *string
}

type DescribeBgpNetworksRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the returned page. Default value: **1**.
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
	// The region ID of the BGP group.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VBR.
	//
	// example:
	//
	// vrt-bp1lhl0taikrteen8****
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
}

func (s DescribeBgpNetworksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpNetworksRequest) GoString() string {
	return s.String()
}

func (s *DescribeBgpNetworksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBgpNetworksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBgpNetworksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBgpNetworksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBgpNetworksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBgpNetworksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBgpNetworksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBgpNetworksRequest) GetRouterId() *string {
	return s.RouterId
}

func (s *DescribeBgpNetworksRequest) SetOwnerAccount(v string) *DescribeBgpNetworksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBgpNetworksRequest) SetOwnerId(v int64) *DescribeBgpNetworksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBgpNetworksRequest) SetPageNumber(v int32) *DescribeBgpNetworksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBgpNetworksRequest) SetPageSize(v int32) *DescribeBgpNetworksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBgpNetworksRequest) SetRegionId(v string) *DescribeBgpNetworksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBgpNetworksRequest) SetResourceOwnerAccount(v string) *DescribeBgpNetworksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBgpNetworksRequest) SetResourceOwnerId(v int64) *DescribeBgpNetworksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBgpNetworksRequest) SetRouterId(v string) *DescribeBgpNetworksRequest {
	s.RouterId = &v
	return s
}

func (s *DescribeBgpNetworksRequest) Validate() error {
	return dara.Validate(s)
}
