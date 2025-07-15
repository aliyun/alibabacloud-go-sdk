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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate a client token. Make sure that a unique client token is used for each request. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks your AccessKey pair, the RAM user permissions, and the required parameters. If the request fails the dry run, the DryRunOperation error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the gateway endpoint to be associated with the route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpce-m5e371h5clm3uadih****
	EndpointId   *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
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
	// The ID of the route table. Valid values of **N*	- are **1*	- to **20**, which specifies that you can associate a gateway endpoint with at most 20 route tables at a time.
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
