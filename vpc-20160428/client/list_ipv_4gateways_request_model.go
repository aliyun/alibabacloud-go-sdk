// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpv4GatewaysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpv4GatewayId(v string) *ListIpv4GatewaysRequest
	GetIpv4GatewayId() *string
	SetIpv4GatewayName(v string) *ListIpv4GatewaysRequest
	GetIpv4GatewayName() *string
	SetMaxResults(v int32) *ListIpv4GatewaysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIpv4GatewaysRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListIpv4GatewaysRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListIpv4GatewaysRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListIpv4GatewaysRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListIpv4GatewaysRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ListIpv4GatewaysRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListIpv4GatewaysRequest
	GetResourceOwnerId() *int64
	SetTags(v []*ListIpv4GatewaysRequestTags) *ListIpv4GatewaysRequest
	GetTags() []*ListIpv4GatewaysRequestTags
	SetVpcId(v string) *ListIpv4GatewaysRequest
	GetVpcId() *string
}

type ListIpv4GatewaysRequest struct {
	// The ID of the IPv4 gateway.
	//
	// example:
	//
	// ipv4gw-5tsnc6s4ogsedtp3k****
	Ipv4GatewayId *string `json:"Ipv4GatewayId,omitempty" xml:"Ipv4GatewayId,omitempty"`
	// The name of the IPv4 gateway.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// name
	Ipv4GatewayName *string `json:"Ipv4GatewayName,omitempty" xml:"Ipv4GatewayName,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If this is your first query and no next queries are to be sent, ignore this parameter.
	//
	// 	- If a next query is to be sent, set the value to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPv4 gateways to be queried are deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-6
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the IPv4 gateway belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the resource.
	Tags []*ListIpv4GatewaysRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) with which the IPv4 gateway is associated.
	//
	// example:
	//
	// vpc-5tsrxlw7dv074gci4****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListIpv4GatewaysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIpv4GatewaysRequest) GoString() string {
	return s.String()
}

func (s *ListIpv4GatewaysRequest) GetIpv4GatewayId() *string {
	return s.Ipv4GatewayId
}

func (s *ListIpv4GatewaysRequest) GetIpv4GatewayName() *string {
	return s.Ipv4GatewayName
}

func (s *ListIpv4GatewaysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpv4GatewaysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpv4GatewaysRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListIpv4GatewaysRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListIpv4GatewaysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpv4GatewaysRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIpv4GatewaysRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListIpv4GatewaysRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListIpv4GatewaysRequest) GetTags() []*ListIpv4GatewaysRequestTags {
	return s.Tags
}

func (s *ListIpv4GatewaysRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListIpv4GatewaysRequest) SetIpv4GatewayId(v string) *ListIpv4GatewaysRequest {
	s.Ipv4GatewayId = &v
	return s
}

func (s *ListIpv4GatewaysRequest) SetIpv4GatewayName(v string) *ListIpv4GatewaysRequest {
	s.Ipv4GatewayName = &v
	return s
}

func (s *ListIpv4GatewaysRequest) SetMaxResults(v int32) *ListIpv4GatewaysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpv4GatewaysRequest) SetNextToken(v string) *ListIpv4GatewaysRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpv4GatewaysRequest) SetOwnerAccount(v string) *ListIpv4GatewaysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListIpv4GatewaysRequest) SetOwnerId(v int64) *ListIpv4GatewaysRequest {
	s.OwnerId = &v
	return s
}

func (s *ListIpv4GatewaysRequest) SetRegionId(v string) *ListIpv4GatewaysRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpv4GatewaysRequest) SetResourceGroupId(v string) *ListIpv4GatewaysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpv4GatewaysRequest) SetResourceOwnerAccount(v string) *ListIpv4GatewaysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListIpv4GatewaysRequest) SetResourceOwnerId(v int64) *ListIpv4GatewaysRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpv4GatewaysRequest) SetTags(v []*ListIpv4GatewaysRequestTags) *ListIpv4GatewaysRequest {
	s.Tags = v
	return s
}

func (s *ListIpv4GatewaysRequest) SetVpcId(v string) *ListIpv4GatewaysRequest {
	s.VpcId = &v
	return s
}

func (s *ListIpv4GatewaysRequest) Validate() error {
	return dara.Validate(s)
}

type ListIpv4GatewaysRequestTags struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The key cannot exceed 64 characters in length, and can contain digits, periods (.), underscores (_), and hyphens (-). The key must start with a letter but cannot start with `aliyun` or `acs:`. The key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. It can be an empty string.
	//
	// The tag value cannot exceed 128 characters in length, and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter but cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpv4GatewaysRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListIpv4GatewaysRequestTags) GoString() string {
	return s.String()
}

func (s *ListIpv4GatewaysRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListIpv4GatewaysRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListIpv4GatewaysRequestTags) SetKey(v string) *ListIpv4GatewaysRequestTags {
	s.Key = &v
	return s
}

func (s *ListIpv4GatewaysRequestTags) SetValue(v string) *ListIpv4GatewaysRequestTags {
	s.Value = &v
	return s
}

func (s *ListIpv4GatewaysRequestTags) Validate() error {
	return dara.Validate(s)
}
