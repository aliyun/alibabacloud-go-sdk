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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the gateway endpoint to be disassociated from the route table.
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
	// The ID of the route table. Valid values of **N*	- are **1*	- to **20**, which specifies that you can disassociate a gateway endpoint from at most 20 route tables at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-m5elgtm3aj586iitr****
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
