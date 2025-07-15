// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSslVpnClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSslVpnClientsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSslVpnClientsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSslVpnClientsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSslVpnClientsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSslVpnClientsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSslVpnClientsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSslVpnClientsRequest
	GetResourceOwnerId() *int64
	SetVpnGatewayId(v string) *DescribeSslVpnClientsRequest
	GetVpnGatewayId() *string
}

type DescribeSslVpnClientsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **50**. Default value: **10**.
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
	// eu-central-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-gw8gfb947ctddabja****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s DescribeSslVpnClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnClientsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSslVpnClientsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSslVpnClientsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSslVpnClientsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSslVpnClientsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSslVpnClientsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSslVpnClientsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSslVpnClientsRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DescribeSslVpnClientsRequest) SetOwnerAccount(v string) *DescribeSslVpnClientsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSslVpnClientsRequest) SetOwnerId(v int64) *DescribeSslVpnClientsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSslVpnClientsRequest) SetPageNumber(v int32) *DescribeSslVpnClientsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSslVpnClientsRequest) SetPageSize(v int32) *DescribeSslVpnClientsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSslVpnClientsRequest) SetRegionId(v string) *DescribeSslVpnClientsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSslVpnClientsRequest) SetResourceOwnerAccount(v string) *DescribeSslVpnClientsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSslVpnClientsRequest) SetResourceOwnerId(v int64) *DescribeSslVpnClientsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSslVpnClientsRequest) SetVpnGatewayId(v string) *DescribeSslVpnClientsRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *DescribeSslVpnClientsRequest) Validate() error {
	return dara.Validate(s)
}
