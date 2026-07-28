// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateRouteTablesFromVpcGatewayEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DissociateRouteTablesFromVpcGatewayEndpointRequest
	GetClientToken() *string
	SetDryRun(v bool) *DissociateRouteTablesFromVpcGatewayEndpointRequest
	GetDryRun() *bool
	SetEndpointId(v string) *DissociateRouteTablesFromVpcGatewayEndpointRequest
	GetEndpointId() *string
	SetOwnerAccount(v string) *DissociateRouteTablesFromVpcGatewayEndpointRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DissociateRouteTablesFromVpcGatewayEndpointRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DissociateRouteTablesFromVpcGatewayEndpointRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DissociateRouteTablesFromVpcGatewayEndpointRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DissociateRouteTablesFromVpcGatewayEndpointRequest
	GetResourceOwnerId() *int64
	SetRouteTableIds(v []*string) *DissociateRouteTablesFromVpcGatewayEndpointRequest
	GetRouteTableIds() []*string
}

type DissociateRouteTablesFromVpcGatewayEndpointRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// TF-DissociateRouteTablesFromVpcGatewayEndpoint-1634369235-8f****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run.
	//
	// - **false*	- (default): sends the request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The endpoint instance ID of the gateway endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpce-bp11cwae3p6z7ftbm****
	EndpointId   *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the endpoint. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of route tables.
	//
	// This parameter is required.
	RouteTableIds []*string `json:"RouteTableIds,omitempty" xml:"RouteTableIds,omitempty" type:"Repeated"`
}

func (s DissociateRouteTablesFromVpcGatewayEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateRouteTablesFromVpcGatewayEndpointRequest) GoString() string {
	return s.String()
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) GetRouteTableIds() []*string {
	return s.RouteTableIds
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) SetClientToken(v string) *DissociateRouteTablesFromVpcGatewayEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) SetDryRun(v bool) *DissociateRouteTablesFromVpcGatewayEndpointRequest {
	s.DryRun = &v
	return s
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) SetEndpointId(v string) *DissociateRouteTablesFromVpcGatewayEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) SetOwnerAccount(v string) *DissociateRouteTablesFromVpcGatewayEndpointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) SetOwnerId(v int64) *DissociateRouteTablesFromVpcGatewayEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) SetRegionId(v string) *DissociateRouteTablesFromVpcGatewayEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) SetResourceOwnerAccount(v string) *DissociateRouteTablesFromVpcGatewayEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) SetResourceOwnerId(v int64) *DissociateRouteTablesFromVpcGatewayEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) SetRouteTableIds(v []*string) *DissociateRouteTablesFromVpcGatewayEndpointRequest {
	s.RouteTableIds = v
	return s
}

func (s *DissociateRouteTablesFromVpcGatewayEndpointRequest) Validate() error {
	return dara.Validate(s)
}
