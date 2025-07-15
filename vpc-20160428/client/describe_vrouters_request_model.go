// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVRoutersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeVRoutersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVRoutersRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVRoutersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVRoutersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVRoutersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVRoutersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVRoutersRequest
	GetResourceOwnerId() *int64
	SetVRouterId(v string) *DescribeVRoutersRequest
	GetVRouterId() *string
}

type DescribeVRoutersRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the vRouter that you want to query.
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
	// The ID of the vRouter.
	//
	// example:
	//
	// vrt-bp1lhl0taikrteen****
	VRouterId *string `json:"VRouterId,omitempty" xml:"VRouterId,omitempty"`
}

func (s DescribeVRoutersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVRoutersRequest) GoString() string {
	return s.String()
}

func (s *DescribeVRoutersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVRoutersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVRoutersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVRoutersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVRoutersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVRoutersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVRoutersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVRoutersRequest) GetVRouterId() *string {
	return s.VRouterId
}

func (s *DescribeVRoutersRequest) SetOwnerAccount(v string) *DescribeVRoutersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVRoutersRequest) SetOwnerId(v int64) *DescribeVRoutersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVRoutersRequest) SetPageNumber(v int32) *DescribeVRoutersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVRoutersRequest) SetPageSize(v int32) *DescribeVRoutersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVRoutersRequest) SetRegionId(v string) *DescribeVRoutersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVRoutersRequest) SetResourceOwnerAccount(v string) *DescribeVRoutersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVRoutersRequest) SetResourceOwnerId(v int64) *DescribeVRoutersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVRoutersRequest) SetVRouterId(v string) *DescribeVRoutersRequest {
	s.VRouterId = &v
	return s
}

func (s *DescribeVRoutersRequest) Validate() error {
	return dara.Validate(s)
}
