// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpv6GatewaysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6GatewayId(v string) *DescribeIpv6GatewaysRequest
	GetIpv6GatewayId() *string
	SetName(v string) *DescribeIpv6GatewaysRequest
	GetName() *string
	SetOwnerAccount(v string) *DescribeIpv6GatewaysRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeIpv6GatewaysRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeIpv6GatewaysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIpv6GatewaysRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeIpv6GatewaysRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeIpv6GatewaysRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeIpv6GatewaysRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeIpv6GatewaysRequest
	GetResourceOwnerId() *int64
	SetTags(v []*DescribeIpv6GatewaysRequestTags) *DescribeIpv6GatewaysRequest
	GetTags() []*DescribeIpv6GatewaysRequestTags
	SetVpcId(v string) *DescribeIpv6GatewaysRequest
	GetVpcId() *string
}

type DescribeIpv6GatewaysRequest struct {
	// The ID of the IPv6 gateway.
	//
	// example:
	//
	// ipv6gw-hp3rwmtmfhgis****
	Ipv6GatewayId *string `json:"Ipv6GatewayId,omitempty" xml:"Ipv6GatewayId,omitempty"`
	// The name of the IPv6 gateway.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. The name must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// ipv6GW
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the IPv6 gateway is deployed. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the IPv6 gateway belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The information about the tags.
	Tags []*DescribeIpv6GatewaysRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the IPv6 gateway belongs.
	//
	// example:
	//
	// vpc-123sedrfswd23****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeIpv6GatewaysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6GatewaysRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpv6GatewaysRequest) GetIpv6GatewayId() *string {
	return s.Ipv6GatewayId
}

func (s *DescribeIpv6GatewaysRequest) GetName() *string {
	return s.Name
}

func (s *DescribeIpv6GatewaysRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeIpv6GatewaysRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeIpv6GatewaysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIpv6GatewaysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIpv6GatewaysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIpv6GatewaysRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeIpv6GatewaysRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeIpv6GatewaysRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeIpv6GatewaysRequest) GetTags() []*DescribeIpv6GatewaysRequestTags {
	return s.Tags
}

func (s *DescribeIpv6GatewaysRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeIpv6GatewaysRequest) SetIpv6GatewayId(v string) *DescribeIpv6GatewaysRequest {
	s.Ipv6GatewayId = &v
	return s
}

func (s *DescribeIpv6GatewaysRequest) SetName(v string) *DescribeIpv6GatewaysRequest {
	s.Name = &v
	return s
}

func (s *DescribeIpv6GatewaysRequest) SetOwnerAccount(v string) *DescribeIpv6GatewaysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeIpv6GatewaysRequest) SetOwnerId(v int64) *DescribeIpv6GatewaysRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeIpv6GatewaysRequest) SetPageNumber(v int32) *DescribeIpv6GatewaysRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeIpv6GatewaysRequest) SetPageSize(v int32) *DescribeIpv6GatewaysRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIpv6GatewaysRequest) SetRegionId(v string) *DescribeIpv6GatewaysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeIpv6GatewaysRequest) SetResourceGroupId(v string) *DescribeIpv6GatewaysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeIpv6GatewaysRequest) SetResourceOwnerAccount(v string) *DescribeIpv6GatewaysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeIpv6GatewaysRequest) SetResourceOwnerId(v int64) *DescribeIpv6GatewaysRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeIpv6GatewaysRequest) SetTags(v []*DescribeIpv6GatewaysRequestTags) *DescribeIpv6GatewaysRequest {
	s.Tags = v
	return s
}

func (s *DescribeIpv6GatewaysRequest) SetVpcId(v string) *DescribeIpv6GatewaysRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeIpv6GatewaysRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeIpv6GatewaysRequestTags struct {
	// The tag keys of the resources. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values of the resources. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIpv6GatewaysRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6GatewaysRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeIpv6GatewaysRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeIpv6GatewaysRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeIpv6GatewaysRequestTags) SetKey(v string) *DescribeIpv6GatewaysRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeIpv6GatewaysRequestTags) SetValue(v string) *DescribeIpv6GatewaysRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeIpv6GatewaysRequestTags) Validate() error {
	return dara.Validate(s)
}
