// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpv6EgressOnlyRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeIpv6EgressOnlyRulesRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeIpv6EgressOnlyRulesRequest
	GetInstanceType() *string
	SetIpv6EgressOnlyRuleId(v string) *DescribeIpv6EgressOnlyRulesRequest
	GetIpv6EgressOnlyRuleId() *string
	SetIpv6GatewayId(v string) *DescribeIpv6EgressOnlyRulesRequest
	GetIpv6GatewayId() *string
	SetName(v string) *DescribeIpv6EgressOnlyRulesRequest
	GetName() *string
	SetOwnerAccount(v string) *DescribeIpv6EgressOnlyRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeIpv6EgressOnlyRulesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeIpv6EgressOnlyRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIpv6EgressOnlyRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeIpv6EgressOnlyRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeIpv6EgressOnlyRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeIpv6EgressOnlyRulesRequest
	GetResourceOwnerId() *int64
}

type DescribeIpv6EgressOnlyRulesRequest struct {
	// The ID of the instance that is associated with the IPv6 address to which the egress-only rule is applied.
	//
	// example:
	//
	// ipv6gw-bp1rhhs9zjlxukc5e****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the instance to which you want to apply the egress-only rule. Valid values:
	//
	// 	- IPv6Address (default)
	//
	// 	- IPv6Prefix
	//
	// example:
	//
	// Ipv6Address
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the egress-only rule that you want to query.
	//
	// example:
	//
	// ipv6py-bp1rr7fq1md8pbb3k****
	Ipv6EgressOnlyRuleId *string `json:"Ipv6EgressOnlyRuleId,omitempty" xml:"Ipv6EgressOnlyRuleId,omitempty"`
	// The ID of the IPv6 gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6gw-bp1rhhs9zjlxukc5e****
	Ipv6GatewayId *string `json:"Ipv6GatewayId,omitempty" xml:"Ipv6GatewayId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// rulename
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// The ID of the region where the IPv6 gateway is deployed. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeIpv6EgressOnlyRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6EgressOnlyRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpv6EgressOnlyRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIpv6EgressOnlyRulesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeIpv6EgressOnlyRulesRequest) GetIpv6EgressOnlyRuleId() *string {
	return s.Ipv6EgressOnlyRuleId
}

func (s *DescribeIpv6EgressOnlyRulesRequest) GetIpv6GatewayId() *string {
	return s.Ipv6GatewayId
}

func (s *DescribeIpv6EgressOnlyRulesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeIpv6EgressOnlyRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeIpv6EgressOnlyRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeIpv6EgressOnlyRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIpv6EgressOnlyRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIpv6EgressOnlyRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIpv6EgressOnlyRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeIpv6EgressOnlyRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeIpv6EgressOnlyRulesRequest) SetInstanceId(v string) *DescribeIpv6EgressOnlyRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesRequest) SetInstanceType(v string) *DescribeIpv6EgressOnlyRulesRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesRequest) SetIpv6EgressOnlyRuleId(v string) *DescribeIpv6EgressOnlyRulesRequest {
	s.Ipv6EgressOnlyRuleId = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesRequest) SetIpv6GatewayId(v string) *DescribeIpv6EgressOnlyRulesRequest {
	s.Ipv6GatewayId = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesRequest) SetName(v string) *DescribeIpv6EgressOnlyRulesRequest {
	s.Name = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesRequest) SetOwnerAccount(v string) *DescribeIpv6EgressOnlyRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesRequest) SetOwnerId(v int64) *DescribeIpv6EgressOnlyRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesRequest) SetPageNumber(v int32) *DescribeIpv6EgressOnlyRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesRequest) SetPageSize(v int32) *DescribeIpv6EgressOnlyRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesRequest) SetRegionId(v string) *DescribeIpv6EgressOnlyRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesRequest) SetResourceOwnerAccount(v string) *DescribeIpv6EgressOnlyRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesRequest) SetResourceOwnerId(v int64) *DescribeIpv6EgressOnlyRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesRequest) Validate() error {
	return dara.Validate(s)
}
