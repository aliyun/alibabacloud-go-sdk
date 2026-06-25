// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSslVpnServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DescribeSslVpnServersRequest
	GetName() *string
	SetOwnerAccount(v string) *DescribeSslVpnServersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSslVpnServersRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSslVpnServersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSslVpnServersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSslVpnServersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSslVpnServersRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeSslVpnServersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSslVpnServersRequest
	GetResourceOwnerId() *int64
	SetSslVpnServerId(v string) *DescribeSslVpnServersRequest
	GetSslVpnServerId() *string
	SetVpnGatewayId(v string) *DescribeSslVpnServersRequest
	GetVpnGatewayId() *string
}

type DescribeSslVpnServersRequest struct {
	// The name of the SSL server.
	//
	// The name must be 1 to 100 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in paging query. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the SSL server.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the SSL server belongs.
	//
	// The SSL server belongs to the same resource group as the associated VPN gateway instance. You can call the [DescribeVpnGateway](https://help.aliyun.com/document_detail/2794055.html) operation to query the resource group ID of the VPN gateway instance.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SSL server.
	//
	// example:
	//
	// vss-bp15j3du13gq1dgey****
	SslVpnServerId *string `json:"SslVpnServerId,omitempty" xml:"SslVpnServerId,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp1on0xae9d771ggi****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s DescribeSslVpnServersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnServersRequest) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnServersRequest) GetName() *string {
	return s.Name
}

func (s *DescribeSslVpnServersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSslVpnServersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSslVpnServersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSslVpnServersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSslVpnServersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSslVpnServersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSslVpnServersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSslVpnServersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSslVpnServersRequest) GetSslVpnServerId() *string {
	return s.SslVpnServerId
}

func (s *DescribeSslVpnServersRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DescribeSslVpnServersRequest) SetName(v string) *DescribeSslVpnServersRequest {
	s.Name = &v
	return s
}

func (s *DescribeSslVpnServersRequest) SetOwnerAccount(v string) *DescribeSslVpnServersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSslVpnServersRequest) SetOwnerId(v int64) *DescribeSslVpnServersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSslVpnServersRequest) SetPageNumber(v int32) *DescribeSslVpnServersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSslVpnServersRequest) SetPageSize(v int32) *DescribeSslVpnServersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSslVpnServersRequest) SetRegionId(v string) *DescribeSslVpnServersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSslVpnServersRequest) SetResourceGroupId(v string) *DescribeSslVpnServersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSslVpnServersRequest) SetResourceOwnerAccount(v string) *DescribeSslVpnServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSslVpnServersRequest) SetResourceOwnerId(v int64) *DescribeSslVpnServersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSslVpnServersRequest) SetSslVpnServerId(v string) *DescribeSslVpnServersRequest {
	s.SslVpnServerId = &v
	return s
}

func (s *DescribeSslVpnServersRequest) SetVpnGatewayId(v string) *DescribeSslVpnServersRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *DescribeSslVpnServersRequest) Validate() error {
	return dara.Validate(s)
}
