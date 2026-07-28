// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerGatewayId(v string) *DescribeVpnConnectionsRequest
	GetCustomerGatewayId() *string
	SetOwnerAccount(v string) *DescribeVpnConnectionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpnConnectionsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVpnConnectionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnConnectionsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVpnConnectionsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeVpnConnectionsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeVpnConnectionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpnConnectionsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeVpnConnectionsRequestTag) *DescribeVpnConnectionsRequest
	GetTag() []*DescribeVpnConnectionsRequestTag
	SetVpnConnectionId(v string) *DescribeVpnConnectionsRequest
	GetVpnConnectionId() *string
	SetVpnGatewayId(v string) *DescribeVpnConnectionsRequest
	GetVpnGatewayId() *string
}

type DescribeVpnConnectionsRequest struct {
	// The ID of the customer gateway.
	//
	// example:
	//
	// cgw-bp1mvj4g9kogw****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paging query. Default value: **10**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the IPsec-VPN connection.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the IPsec-VPN connection belongs.
	//
	// You can call [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) to query resource group IDs.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags bound to the IPsec-VPN connection.
	//
	// You can specify up to 20 tags at a time.
	Tag []*DescribeVpnConnectionsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the IPsec-VPN connection.
	//
	// example:
	//
	// vco-bp10lz7aejumd****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
	// The instance ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp1q8bgx4xnkx****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s DescribeVpnConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsRequest) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *DescribeVpnConnectionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpnConnectionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpnConnectionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnConnectionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnConnectionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpnConnectionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVpnConnectionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpnConnectionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpnConnectionsRequest) GetTag() []*DescribeVpnConnectionsRequestTag {
	return s.Tag
}

func (s *DescribeVpnConnectionsRequest) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *DescribeVpnConnectionsRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DescribeVpnConnectionsRequest) SetCustomerGatewayId(v string) *DescribeVpnConnectionsRequest {
	s.CustomerGatewayId = &v
	return s
}

func (s *DescribeVpnConnectionsRequest) SetOwnerAccount(v string) *DescribeVpnConnectionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpnConnectionsRequest) SetOwnerId(v int64) *DescribeVpnConnectionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpnConnectionsRequest) SetPageNumber(v int32) *DescribeVpnConnectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnConnectionsRequest) SetPageSize(v int32) *DescribeVpnConnectionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnConnectionsRequest) SetRegionId(v string) *DescribeVpnConnectionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpnConnectionsRequest) SetResourceGroupId(v string) *DescribeVpnConnectionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVpnConnectionsRequest) SetResourceOwnerAccount(v string) *DescribeVpnConnectionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpnConnectionsRequest) SetResourceOwnerId(v int64) *DescribeVpnConnectionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpnConnectionsRequest) SetTag(v []*DescribeVpnConnectionsRequestTag) *DescribeVpnConnectionsRequest {
	s.Tag = v
	return s
}

func (s *DescribeVpnConnectionsRequest) SetVpnConnectionId(v string) *DescribeVpnConnectionsRequest {
	s.VpnConnectionId = &v
	return s
}

func (s *DescribeVpnConnectionsRequest) SetVpnGatewayId(v string) *DescribeVpnConnectionsRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *DescribeVpnConnectionsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpnConnectionsRequestTag struct {
	// The tag key. If you specify this parameter, the value cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be up to 128 characters in length and can be an empty string. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each tag key corresponds to one tag value. You can specify up to 20 tag values at a time.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVpnConnectionsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnConnectionsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVpnConnectionsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVpnConnectionsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVpnConnectionsRequestTag) SetKey(v string) *DescribeVpnConnectionsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVpnConnectionsRequestTag) SetValue(v string) *DescribeVpnConnectionsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeVpnConnectionsRequestTag) Validate() error {
	return dara.Validate(s)
}
