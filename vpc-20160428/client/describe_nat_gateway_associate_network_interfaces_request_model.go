// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatGatewayAssociateNetworkInterfacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequest
	GetClientToken() *string
	SetFilter(v []*DescribeNatGatewayAssociateNetworkInterfacesRequestFilter) *DescribeNatGatewayAssociateNetworkInterfacesRequest
	GetFilter() []*DescribeNatGatewayAssociateNetworkInterfacesRequestFilter
	SetMaxResults(v int32) *DescribeNatGatewayAssociateNetworkInterfacesRequest
	GetMaxResults() *int32
	SetNatGatewayId(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequest
	GetNatGatewayId() *string
	SetNextToken(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeNatGatewayAssociateNetworkInterfacesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeNatGatewayAssociateNetworkInterfacesRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeNatGatewayAssociateNetworkInterfacesRequestTag) *DescribeNatGatewayAssociateNetworkInterfacesRequest
	GetTag() []*DescribeNatGatewayAssociateNetworkInterfacesRequestTag
}

type DescribeNatGatewayAssociateNetworkInterfacesRequest struct {
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string                                                      `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Filter      []*DescribeNatGatewayAssociateNetworkInterfacesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmxazdjdhd****
	ResourceGroupId      *string                                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string                                                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tag                  []*DescribeNatGatewayAssociateNetworkInterfacesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeNatGatewayAssociateNetworkInterfacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewayAssociateNetworkInterfacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) GetFilter() []*DescribeNatGatewayAssociateNetworkInterfacesRequestFilter {
	return s.Filter
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) GetTag() []*DescribeNatGatewayAssociateNetworkInterfacesRequestTag {
	return s.Tag
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) SetClientToken(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) SetFilter(v []*DescribeNatGatewayAssociateNetworkInterfacesRequestFilter) *DescribeNatGatewayAssociateNetworkInterfacesRequest {
	s.Filter = v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) SetMaxResults(v int32) *DescribeNatGatewayAssociateNetworkInterfacesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) SetNatGatewayId(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) SetNextToken(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) SetOwnerAccount(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) SetOwnerId(v int64) *DescribeNatGatewayAssociateNetworkInterfacesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) SetRegionId(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) SetResourceGroupId(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) SetResourceOwnerAccount(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) SetResourceOwnerId(v int64) *DescribeNatGatewayAssociateNetworkInterfacesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) SetTag(v []*DescribeNatGatewayAssociateNetworkInterfacesRequestTag) *DescribeNatGatewayAssociateNetworkInterfacesRequest {
	s.Tag = v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeNatGatewayAssociateNetworkInterfacesRequestFilter struct {
	// example:
	//
	// ResourceId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// ep-8psre8c8936596cd****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNatGatewayAssociateNetworkInterfacesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewayAssociateNetworkInterfacesRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequestFilter) GetValue() *string {
	return s.Value
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequestFilter) SetKey(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequestFilter) SetValue(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequestFilter {
	s.Value = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequestFilter) Validate() error {
	return dara.Validate(s)
}

type DescribeNatGatewayAssociateNetworkInterfacesRequestTag struct {
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNatGatewayAssociateNetworkInterfacesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewayAssociateNetworkInterfacesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequestTag) SetKey(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequestTag) SetValue(v string) *DescribeNatGatewayAssociateNetworkInterfacesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesRequestTag) Validate() error {
	return dara.Validate(s)
}
