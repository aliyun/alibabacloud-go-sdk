// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayRouteTableEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidrBlock(v string) *ListGatewayRouteTableEntriesRequest
	GetDestinationCidrBlock() *string
	SetGatewayRouteTableId(v string) *ListGatewayRouteTableEntriesRequest
	GetGatewayRouteTableId() *string
	SetMaxResults(v int32) *ListGatewayRouteTableEntriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListGatewayRouteTableEntriesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListGatewayRouteTableEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListGatewayRouteTableEntriesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListGatewayRouteTableEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListGatewayRouteTableEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListGatewayRouteTableEntriesRequest
	GetResourceOwnerId() *int64
}

type ListGatewayRouteTableEntriesRequest struct {
	// The destination CIDR block of the route entry in the gateway route table.
	//
	// example:
	//
	// 192.168.0.5
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the gateway route table that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-5ts0ohchwkp3dydt2****
	GatewayRouteTableId *string `json:"GatewayRouteTableId,omitempty" xml:"GatewayRouteTableId,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- If a value is returned for NextToken, specify the value in the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the gateway route table.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent list of regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-6
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListGatewayRouteTableEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteTableEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteTableEntriesRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ListGatewayRouteTableEntriesRequest) GetGatewayRouteTableId() *string {
	return s.GatewayRouteTableId
}

func (s *ListGatewayRouteTableEntriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGatewayRouteTableEntriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGatewayRouteTableEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListGatewayRouteTableEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListGatewayRouteTableEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListGatewayRouteTableEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListGatewayRouteTableEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListGatewayRouteTableEntriesRequest) SetDestinationCidrBlock(v string) *ListGatewayRouteTableEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListGatewayRouteTableEntriesRequest) SetGatewayRouteTableId(v string) *ListGatewayRouteTableEntriesRequest {
	s.GatewayRouteTableId = &v
	return s
}

func (s *ListGatewayRouteTableEntriesRequest) SetMaxResults(v int32) *ListGatewayRouteTableEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGatewayRouteTableEntriesRequest) SetNextToken(v string) *ListGatewayRouteTableEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListGatewayRouteTableEntriesRequest) SetOwnerAccount(v string) *ListGatewayRouteTableEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListGatewayRouteTableEntriesRequest) SetOwnerId(v int64) *ListGatewayRouteTableEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListGatewayRouteTableEntriesRequest) SetRegionId(v string) *ListGatewayRouteTableEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListGatewayRouteTableEntriesRequest) SetResourceOwnerAccount(v string) *ListGatewayRouteTableEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListGatewayRouteTableEntriesRequest) SetResourceOwnerId(v int64) *ListGatewayRouteTableEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListGatewayRouteTableEntriesRequest) Validate() error {
	return dara.Validate(s)
}
