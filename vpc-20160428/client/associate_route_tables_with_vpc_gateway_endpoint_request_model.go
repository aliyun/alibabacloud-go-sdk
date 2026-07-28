// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateRouteTablesWithVpcGatewayEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AssociateRouteTablesWithVpcGatewayEndpointRequest
	GetClientToken() *string
	SetDryRun(v bool) *AssociateRouteTablesWithVpcGatewayEndpointRequest
	GetDryRun() *bool
	SetEndpointId(v string) *AssociateRouteTablesWithVpcGatewayEndpointRequest
	GetEndpointId() *string
	SetOwnerAccount(v string) *AssociateRouteTablesWithVpcGatewayEndpointRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociateRouteTablesWithVpcGatewayEndpointRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AssociateRouteTablesWithVpcGatewayEndpointRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AssociateRouteTablesWithVpcGatewayEndpointRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateRouteTablesWithVpcGatewayEndpointRequest
	GetResourceOwnerId() *int64
	SetRouteTableIds(v []*string) *AssociateRouteTablesWithVpcGatewayEndpointRequest
	GetRouteTableIds() []*string
}

type AssociateRouteTablesWithVpcGatewayEndpointRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may differ for each API request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: sends a dry run request without associating route tables with the gateway endpoint. The system checks the AccessKey pair, the authorization of the Resource Access Management (RAM) user, and the required parameters. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. After the request passes the dry run, a 2xx HTTP status code is returned and the route tables are associated.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the gateway endpoint instance with which you want to associate route tables.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpce-m5e371h5clm3uadih****
	EndpointId   *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the gateway endpoint with which you want to associate route tables.
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
	// The ID of the route table to associate. Valid values of **N**: **1*	- to **20**. You can associate up to 20 route tables at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-m5elgtm3aj586iitr****
	RouteTableIds []*string `json:"RouteTableIds,omitempty" xml:"RouteTableIds,omitempty" type:"Repeated"`
}

func (s AssociateRouteTablesWithVpcGatewayEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateRouteTablesWithVpcGatewayEndpointRequest) GoString() string {
	return s.String()
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) GetRouteTableIds() []*string {
	return s.RouteTableIds
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) SetClientToken(v string) *AssociateRouteTablesWithVpcGatewayEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) SetDryRun(v bool) *AssociateRouteTablesWithVpcGatewayEndpointRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) SetEndpointId(v string) *AssociateRouteTablesWithVpcGatewayEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) SetOwnerAccount(v string) *AssociateRouteTablesWithVpcGatewayEndpointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) SetOwnerId(v int64) *AssociateRouteTablesWithVpcGatewayEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) SetRegionId(v string) *AssociateRouteTablesWithVpcGatewayEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) SetResourceOwnerAccount(v string) *AssociateRouteTablesWithVpcGatewayEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) SetResourceOwnerId(v int64) *AssociateRouteTablesWithVpcGatewayEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) SetRouteTableIds(v []*string) *AssociateRouteTablesWithVpcGatewayEndpointRequest {
	s.RouteTableIds = v
	return s
}

func (s *AssociateRouteTablesWithVpcGatewayEndpointRequest) Validate() error {
	return dara.Validate(s)
}
