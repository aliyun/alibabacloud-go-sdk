// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointServicesByEndUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListVpcEndpointServicesByEndUserRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListVpcEndpointServicesByEndUserRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListVpcEndpointServicesByEndUserRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListVpcEndpointServicesByEndUserRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListVpcEndpointServicesByEndUserRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListVpcEndpointServicesByEndUserRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListVpcEndpointServicesByEndUserRequest
	GetResourceOwnerId() *int64
	SetServiceName(v string) *ListVpcEndpointServicesByEndUserRequest
	GetServiceName() *string
}

type ListVpcEndpointServicesByEndUserRequest struct {
	// The number of entries to return per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If this is your first query and no next queries are to be sent, ignore this parameter.
	//
	// 	- If a next query is to be performed, set the value to the NextToken value returned in the last call to the ListListenerCertificates operation.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the gateway endpoint.
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
	// The name of the endpoint service that you want to query.
	//
	// example:
	//
	// com.aliyun.cn-hangzhou.oss
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListVpcEndpointServicesByEndUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListVpcEndpointServicesByEndUserRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetMaxResults(v int64) *ListVpcEndpointServicesByEndUserRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetNextToken(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetOwnerAccount(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetOwnerId(v int64) *ListVpcEndpointServicesByEndUserRequest {
	s.OwnerId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetRegionId(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetResourceOwnerAccount(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetResourceOwnerId(v int64) *ListVpcEndpointServicesByEndUserRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) SetServiceName(v string) *ListVpcEndpointServicesByEndUserRequest {
	s.ServiceName = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserRequest) Validate() error {
	return dara.Validate(s)
}
