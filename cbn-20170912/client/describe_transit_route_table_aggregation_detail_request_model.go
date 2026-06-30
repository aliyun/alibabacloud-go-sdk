// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouteTableAggregationDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeTransitRouteTableAggregationDetailRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DescribeTransitRouteTableAggregationDetailRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTransitRouteTableAggregationDetailRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeTransitRouteTableAggregationDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTransitRouteTableAggregationDetailRequest
	GetResourceOwnerId() *int64
	SetTransitRouteTableAggregationCidr(v string) *DescribeTransitRouteTableAggregationDetailRequest
	GetTransitRouteTableAggregationCidr() *string
	SetTransitRouteTableId(v string) *DescribeTransitRouteTableAggregationDetailRequest
	GetTransitRouteTableId() *string
}

type DescribeTransitRouteTableAggregationDetailRequest struct {
	// A client token to ensure the idempotence of the request.
	//
	// Generate a unique value from your client for each request. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The destination CIDR block of the aggregate route.
	//
	// This parameter is required.
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

func (s DescribeTransitRouteTableAggregationDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouteTableAggregationDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouteTableAggregationDetailRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeTransitRouteTableAggregationDetailRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTransitRouteTableAggregationDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTransitRouteTableAggregationDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTransitRouteTableAggregationDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTransitRouteTableAggregationDetailRequest) GetTransitRouteTableAggregationCidr() *string {
	return s.TransitRouteTableAggregationCidr
}

func (s *DescribeTransitRouteTableAggregationDetailRequest) GetTransitRouteTableId() *string {
	return s.TransitRouteTableId
}

func (s *DescribeTransitRouteTableAggregationDetailRequest) SetClientToken(v string) *DescribeTransitRouteTableAggregationDetailRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailRequest) SetOwnerAccount(v string) *DescribeTransitRouteTableAggregationDetailRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailRequest) SetOwnerId(v int64) *DescribeTransitRouteTableAggregationDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailRequest) SetResourceOwnerAccount(v string) *DescribeTransitRouteTableAggregationDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailRequest) SetResourceOwnerId(v int64) *DescribeTransitRouteTableAggregationDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailRequest) SetTransitRouteTableAggregationCidr(v string) *DescribeTransitRouteTableAggregationDetailRequest {
	s.TransitRouteTableAggregationCidr = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailRequest) SetTransitRouteTableId(v string) *DescribeTransitRouteTableAggregationDetailRequest {
	s.TransitRouteTableId = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationDetailRequest) Validate() error {
	return dara.Validate(s)
}
