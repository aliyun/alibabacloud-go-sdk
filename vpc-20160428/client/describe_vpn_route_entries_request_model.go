// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeVpnRouteEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpnRouteEntriesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVpnRouteEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnRouteEntriesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVpnRouteEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVpnRouteEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpnRouteEntriesRequest
	GetResourceOwnerId() *int64
	SetRouteEntryType(v string) *DescribeVpnRouteEntriesRequest
	GetRouteEntryType() *string
	SetVpnGatewayId(v string) *DescribeVpnRouteEntriesRequest
	GetVpnGatewayId() *string
}

type DescribeVpnRouteEntriesRequest struct {
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
	// The ID of the region where the VPN gateway is created.
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
	// The type of the route entry. Valid values:
	//
	// 	- **Custom**: custom
	//
	// 	- **System**: system
	//
	// example:
	//
	// System
	RouteEntryType *string `json:"RouteEntryType,omitempty" xml:"RouteEntryType,omitempty"`
	// The ID of the VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp1cmw7jh1nfe43m9****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s DescribeVpnRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpnRouteEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpnRouteEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpnRouteEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnRouteEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnRouteEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpnRouteEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpnRouteEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpnRouteEntriesRequest) GetRouteEntryType() *string {
	return s.RouteEntryType
}

func (s *DescribeVpnRouteEntriesRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DescribeVpnRouteEntriesRequest) SetOwnerAccount(v string) *DescribeVpnRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpnRouteEntriesRequest) SetOwnerId(v int64) *DescribeVpnRouteEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpnRouteEntriesRequest) SetPageNumber(v int32) *DescribeVpnRouteEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnRouteEntriesRequest) SetPageSize(v int32) *DescribeVpnRouteEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnRouteEntriesRequest) SetRegionId(v string) *DescribeVpnRouteEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpnRouteEntriesRequest) SetResourceOwnerAccount(v string) *DescribeVpnRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpnRouteEntriesRequest) SetResourceOwnerId(v int64) *DescribeVpnRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpnRouteEntriesRequest) SetRouteEntryType(v string) *DescribeVpnRouteEntriesRequest {
	s.RouteEntryType = &v
	return s
}

func (s *DescribeVpnRouteEntriesRequest) SetVpnGatewayId(v string) *DescribeVpnRouteEntriesRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *DescribeVpnRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
