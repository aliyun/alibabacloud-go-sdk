// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateRouteTableWithGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AssociateRouteTableWithGatewayRequest
	GetClientToken() *string
	SetDryRun(v bool) *AssociateRouteTableWithGatewayRequest
	GetDryRun() *bool
	SetGatewayId(v string) *AssociateRouteTableWithGatewayRequest
	GetGatewayId() *string
	SetGatewayType(v string) *AssociateRouteTableWithGatewayRequest
	GetGatewayType() *string
	SetOwnerAccount(v string) *AssociateRouteTableWithGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociateRouteTableWithGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AssociateRouteTableWithGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AssociateRouteTableWithGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateRouteTableWithGatewayRequest
	GetResourceOwnerId() *int64
	SetRouteTableId(v string) *AssociateRouteTableWithGatewayRequest
	GetRouteTableId() *string
}

type AssociateRouteTableWithGatewayRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not set this parameter, the system automatically uses **RequestId*	- as **ClientToken**. **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// 	- **true**: prechecks the request without performing the operation. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the request. After the request passes the precheck, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPv4 gateway.
	//
	// The IPv4 gateway must be in the **Activated*	- state.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv4gw-5tsnc6s4ogsedtp3k****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The type of a gateway to be associated with a route table.
	//
	// example:
	//
	// Ipv4Gateway
	GatewayType  *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IPv4 gateway with which you want to associate the gateway route table.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-6
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the gateway route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-5ts0ohchwkp3dydt2****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s AssociateRouteTableWithGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateRouteTableWithGatewayRequest) GoString() string {
	return s.String()
}

func (s *AssociateRouteTableWithGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateRouteTableWithGatewayRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AssociateRouteTableWithGatewayRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *AssociateRouteTableWithGatewayRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *AssociateRouteTableWithGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociateRouteTableWithGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateRouteTableWithGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateRouteTableWithGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateRouteTableWithGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateRouteTableWithGatewayRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *AssociateRouteTableWithGatewayRequest) SetClientToken(v string) *AssociateRouteTableWithGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateRouteTableWithGatewayRequest) SetDryRun(v bool) *AssociateRouteTableWithGatewayRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateRouteTableWithGatewayRequest) SetGatewayId(v string) *AssociateRouteTableWithGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *AssociateRouteTableWithGatewayRequest) SetGatewayType(v string) *AssociateRouteTableWithGatewayRequest {
	s.GatewayType = &v
	return s
}

func (s *AssociateRouteTableWithGatewayRequest) SetOwnerAccount(v string) *AssociateRouteTableWithGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateRouteTableWithGatewayRequest) SetOwnerId(v int64) *AssociateRouteTableWithGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateRouteTableWithGatewayRequest) SetRegionId(v string) *AssociateRouteTableWithGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateRouteTableWithGatewayRequest) SetResourceOwnerAccount(v string) *AssociateRouteTableWithGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateRouteTableWithGatewayRequest) SetResourceOwnerId(v int64) *AssociateRouteTableWithGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateRouteTableWithGatewayRequest) SetRouteTableId(v string) *AssociateRouteTableWithGatewayRequest {
	s.RouteTableId = &v
	return s
}

func (s *AssociateRouteTableWithGatewayRequest) Validate() error {
	return dara.Validate(s)
}
