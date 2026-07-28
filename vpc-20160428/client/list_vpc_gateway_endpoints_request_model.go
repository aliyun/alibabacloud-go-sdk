// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcGatewayEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *ListVpcGatewayEndpointsRequest
	GetEndpointId() *string
	SetEndpointName(v string) *ListVpcGatewayEndpointsRequest
	GetEndpointName() *string
	SetMaxResults(v int64) *ListVpcGatewayEndpointsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListVpcGatewayEndpointsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListVpcGatewayEndpointsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListVpcGatewayEndpointsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListVpcGatewayEndpointsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListVpcGatewayEndpointsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ListVpcGatewayEndpointsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListVpcGatewayEndpointsRequest
	GetResourceOwnerId() *int64
	SetServiceName(v string) *ListVpcGatewayEndpointsRequest
	GetServiceName() *string
	SetTags(v []*ListVpcGatewayEndpointsRequestTags) *ListVpcGatewayEndpointsRequest
	GetTags() []*ListVpcGatewayEndpointsRequestTags
	SetVpcId(v string) *ListVpcGatewayEndpointsRequest
	GetVpcId() *string
}

type ListVpcGatewayEndpointsRequest struct {
	// The endpoint instance ID of the gateway endpoint.
	//
	// example:
	//
	// vpce-bp1i1212ss2whuwyw****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The name of the gateway endpoint.
	//
	// The name must be 1 to 128 characters in length.
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The number of entries per page for a paged query. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Specifies whether a next query token is available. Valid values:
	//
	// - Leave this parameter empty for the first query or when no more results are available.
	//
	// - If more results are available, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the gateway endpoint that you want to query.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the gateway endpoint belongs.
	//
	// example:
	//
	// rg-acfmxvfvazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service name of the endpoint service.
	//
	// example:
	//
	// com.aliyun.cn-hangzhou.oss
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The list of tags.
	Tags []*ListVpcGatewayEndpointsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the VPC to which the gateway endpoint belongs.
	//
	// example:
	//
	// vpc-bp1gsk7h12ew7oegk****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVpcGatewayEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcGatewayEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListVpcGatewayEndpointsRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListVpcGatewayEndpointsRequest) GetEndpointName() *string {
	return s.EndpointName
}

func (s *ListVpcGatewayEndpointsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListVpcGatewayEndpointsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcGatewayEndpointsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListVpcGatewayEndpointsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListVpcGatewayEndpointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcGatewayEndpointsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpcGatewayEndpointsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListVpcGatewayEndpointsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListVpcGatewayEndpointsRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListVpcGatewayEndpointsRequest) GetTags() []*ListVpcGatewayEndpointsRequestTags {
	return s.Tags
}

func (s *ListVpcGatewayEndpointsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVpcGatewayEndpointsRequest) SetEndpointId(v string) *ListVpcGatewayEndpointsRequest {
	s.EndpointId = &v
	return s
}

func (s *ListVpcGatewayEndpointsRequest) SetEndpointName(v string) *ListVpcGatewayEndpointsRequest {
	s.EndpointName = &v
	return s
}

func (s *ListVpcGatewayEndpointsRequest) SetMaxResults(v int64) *ListVpcGatewayEndpointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcGatewayEndpointsRequest) SetNextToken(v string) *ListVpcGatewayEndpointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcGatewayEndpointsRequest) SetOwnerAccount(v string) *ListVpcGatewayEndpointsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListVpcGatewayEndpointsRequest) SetOwnerId(v int64) *ListVpcGatewayEndpointsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListVpcGatewayEndpointsRequest) SetRegionId(v string) *ListVpcGatewayEndpointsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcGatewayEndpointsRequest) SetResourceGroupId(v string) *ListVpcGatewayEndpointsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcGatewayEndpointsRequest) SetResourceOwnerAccount(v string) *ListVpcGatewayEndpointsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListVpcGatewayEndpointsRequest) SetResourceOwnerId(v int64) *ListVpcGatewayEndpointsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListVpcGatewayEndpointsRequest) SetServiceName(v string) *ListVpcGatewayEndpointsRequest {
	s.ServiceName = &v
	return s
}

func (s *ListVpcGatewayEndpointsRequest) SetTags(v []*ListVpcGatewayEndpointsRequestTags) *ListVpcGatewayEndpointsRequest {
	s.Tags = v
	return s
}

func (s *ListVpcGatewayEndpointsRequest) SetVpcId(v string) *ListVpcGatewayEndpointsRequest {
	s.VpcId = &v
	return s
}

func (s *ListVpcGatewayEndpointsRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVpcGatewayEndpointsRequestTags struct {
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcGatewayEndpointsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListVpcGatewayEndpointsRequestTags) GoString() string {
	return s.String()
}

func (s *ListVpcGatewayEndpointsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListVpcGatewayEndpointsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListVpcGatewayEndpointsRequestTags) SetKey(v string) *ListVpcGatewayEndpointsRequestTags {
	s.Key = &v
	return s
}

func (s *ListVpcGatewayEndpointsRequestTags) SetValue(v string) *ListVpcGatewayEndpointsRequestTags {
	s.Value = &v
	return s
}

func (s *ListVpcGatewayEndpointsRequestTags) Validate() error {
	return dara.Validate(s)
}
