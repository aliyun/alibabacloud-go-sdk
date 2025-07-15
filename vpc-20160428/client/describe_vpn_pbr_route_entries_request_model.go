// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnPbrRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeVpnPbrRouteEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpnPbrRouteEntriesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVpnPbrRouteEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnPbrRouteEntriesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVpnPbrRouteEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVpnPbrRouteEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpnPbrRouteEntriesRequest
	GetResourceOwnerId() *int64
	SetVpnGatewayId(v string) *DescribeVpnPbrRouteEntriesRequest
	GetVpnGatewayId() *string
}

type DescribeVpnPbrRouteEntriesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the VPN gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp1a3kqjiiq9legfx****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s DescribeVpnPbrRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnPbrRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpnPbrRouteEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpnPbrRouteEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpnPbrRouteEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnPbrRouteEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnPbrRouteEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpnPbrRouteEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpnPbrRouteEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpnPbrRouteEntriesRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DescribeVpnPbrRouteEntriesRequest) SetOwnerAccount(v string) *DescribeVpnPbrRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesRequest) SetOwnerId(v int64) *DescribeVpnPbrRouteEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesRequest) SetPageNumber(v int32) *DescribeVpnPbrRouteEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesRequest) SetPageSize(v int32) *DescribeVpnPbrRouteEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesRequest) SetRegionId(v string) *DescribeVpnPbrRouteEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesRequest) SetResourceOwnerAccount(v string) *DescribeVpnPbrRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesRequest) SetResourceOwnerId(v int64) *DescribeVpnPbrRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesRequest) SetVpnGatewayId(v string) *DescribeVpnPbrRouteEntriesRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
