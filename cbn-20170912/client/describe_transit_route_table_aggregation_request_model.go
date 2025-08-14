// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouteTableAggregationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeTransitRouteTableAggregationRequest
	GetClientToken() *string
	SetMaxResults(v int64) *DescribeTransitRouteTableAggregationRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeTransitRouteTableAggregationRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeTransitRouteTableAggregationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTransitRouteTableAggregationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeTransitRouteTableAggregationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTransitRouteTableAggregationRequest
	GetResourceOwnerId() *int64
	SetTransitRouteTableAggregationCidr(v string) *DescribeTransitRouteTableAggregationRequest
	GetTransitRouteTableAggregationCidr() *string
	SetTransitRouteTableId(v string) *DescribeTransitRouteTableAggregationRequest
	GetTransitRouteTableId() *string
}

type DescribeTransitRouteTableAggregationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// 	- If this is your first query or no subsequent query is to be sent, ignore this parameter.
	//
	// 	- If a next query is to be sent, set the value to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The destination CIDR block of the aggregate route.
	//
	// example:
	//
	// 192.168.10.0/24
	TransitRouteTableAggregationCidr *string `json:"TransitRouteTableAggregationCidr,omitempty" xml:"TransitRouteTableAggregationCidr,omitempty"`
	// The ID of the route table of the Enterprise Edition transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-6ehgc262hr170qgyc****
	TransitRouteTableId *string `json:"TransitRouteTableId,omitempty" xml:"TransitRouteTableId,omitempty"`
}

func (s DescribeTransitRouteTableAggregationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouteTableAggregationRequest) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouteTableAggregationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeTransitRouteTableAggregationRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeTransitRouteTableAggregationRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTransitRouteTableAggregationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTransitRouteTableAggregationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTransitRouteTableAggregationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTransitRouteTableAggregationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTransitRouteTableAggregationRequest) GetTransitRouteTableAggregationCidr() *string {
	return s.TransitRouteTableAggregationCidr
}

func (s *DescribeTransitRouteTableAggregationRequest) GetTransitRouteTableId() *string {
	return s.TransitRouteTableId
}

func (s *DescribeTransitRouteTableAggregationRequest) SetClientToken(v string) *DescribeTransitRouteTableAggregationRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationRequest) SetMaxResults(v int64) *DescribeTransitRouteTableAggregationRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationRequest) SetNextToken(v string) *DescribeTransitRouteTableAggregationRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationRequest) SetOwnerAccount(v string) *DescribeTransitRouteTableAggregationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationRequest) SetOwnerId(v int64) *DescribeTransitRouteTableAggregationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationRequest) SetResourceOwnerAccount(v string) *DescribeTransitRouteTableAggregationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationRequest) SetResourceOwnerId(v int64) *DescribeTransitRouteTableAggregationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationRequest) SetTransitRouteTableAggregationCidr(v string) *DescribeTransitRouteTableAggregationRequest {
	s.TransitRouteTableAggregationCidr = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationRequest) SetTransitRouteTableId(v string) *DescribeTransitRouteTableAggregationRequest {
	s.TransitRouteTableId = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationRequest) Validate() error {
	return dara.Validate(s)
}
