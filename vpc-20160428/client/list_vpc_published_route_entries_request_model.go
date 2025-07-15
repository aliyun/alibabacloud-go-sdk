// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcPublishedRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidrBlock(v string) *ListVpcPublishedRouteEntriesRequest
	GetDestinationCidrBlock() *string
	SetMaxResults(v int32) *ListVpcPublishedRouteEntriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcPublishedRouteEntriesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListVpcPublishedRouteEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListVpcPublishedRouteEntriesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListVpcPublishedRouteEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListVpcPublishedRouteEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListVpcPublishedRouteEntriesRequest
	GetResourceOwnerId() *int64
	SetRouteTableId(v string) *ListVpcPublishedRouteEntriesRequest
	GetRouteTableId() *string
	SetTargetInstanceId(v string) *ListVpcPublishedRouteEntriesRequest
	GetTargetInstanceId() *string
	SetTargetType(v string) *ListVpcPublishedRouteEntriesRequest
	GetTargetType() *string
}

type ListVpcPublishedRouteEntriesRequest struct {
	// The destination CIDR block of the route entry, supporting both IPv4 and IPv6 segments.
	//
	// example:
	//
	// 47.100.XX.XX/16
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The number of entries to display per batch query. Range: **1**~**500**, default value is **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Indicates whether there is a token for the next query. Values:
	//
	// - If **NextToken*	- is empty, it means there is no next query.
	//
	// - If **NextToken*	- has a return value, this value indicates the token for the start of the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the instance is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp145q7glnuzd****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The ID of the route publishing target instance.
	//
	// example:
	//
	// ecr-dhw2xsds5****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The type of the route publishing target.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECR
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListVpcPublishedRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcPublishedRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListVpcPublishedRouteEntriesRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ListVpcPublishedRouteEntriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcPublishedRouteEntriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcPublishedRouteEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListVpcPublishedRouteEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListVpcPublishedRouteEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcPublishedRouteEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListVpcPublishedRouteEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListVpcPublishedRouteEntriesRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *ListVpcPublishedRouteEntriesRequest) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *ListVpcPublishedRouteEntriesRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListVpcPublishedRouteEntriesRequest) SetDestinationCidrBlock(v string) *ListVpcPublishedRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesRequest) SetMaxResults(v int32) *ListVpcPublishedRouteEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesRequest) SetNextToken(v string) *ListVpcPublishedRouteEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesRequest) SetOwnerAccount(v string) *ListVpcPublishedRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesRequest) SetOwnerId(v int64) *ListVpcPublishedRouteEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesRequest) SetRegionId(v string) *ListVpcPublishedRouteEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesRequest) SetResourceOwnerAccount(v string) *ListVpcPublishedRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesRequest) SetResourceOwnerId(v int64) *ListVpcPublishedRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesRequest) SetRouteTableId(v string) *ListVpcPublishedRouteEntriesRequest {
	s.RouteTableId = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesRequest) SetTargetInstanceId(v string) *ListVpcPublishedRouteEntriesRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesRequest) SetTargetType(v string) *ListVpcPublishedRouteEntriesRequest {
	s.TargetType = &v
	return s
}

func (s *ListVpcPublishedRouteEntriesRequest) Validate() error {
	return dara.Validate(s)
}
