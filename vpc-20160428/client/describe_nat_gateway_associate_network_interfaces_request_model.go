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
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the value of **RequestId*	- is used.***	- The **RequestId*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The filter information. You can specify a filter key and a filter value.
	Filter []*DescribeNatGatewayAssociateNetworkInterfacesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return per page. Valid values: **1 to 100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The ID of the NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid value:
	//
	// 	- If no value is returned for NetToken, you do not need to specify this parameter.
	//
	// 	- If a value is returned for NextToken, you must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Internet NAT gateway.
	//
	// Call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxazdjdhd****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The information about resource tags.
	Tag []*DescribeNatGatewayAssociateNetworkInterfacesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type DescribeNatGatewayAssociateNetworkInterfacesRequestFilter struct {
	// The filter key.
	//
	// 	- ResourceId
	//
	// >  Specify the service resource ID in the Value field.
	//
	// 	- NetworkInterfaceId
	//
	// >  Specify the ENI ID in the Value field.
	//
	// 	- ResourceOwnerId
	//
	// >  Specify the UID of the account to which the service resource belongs.
	//
	// example:
	//
	// ResourceId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Separate multiple values with commas (,).
	//
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
	// The tag key You can specify at most 20 tag keys. It cannot be an empty string,
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag key. You can specify at most 20 tag keys. It cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
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
