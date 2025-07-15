// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateRouteTableFromGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DissociateRouteTableFromGatewayRequest
	GetClientToken() *string
	SetDryRun(v bool) *DissociateRouteTableFromGatewayRequest
	GetDryRun() *bool
	SetGatewayId(v string) *DissociateRouteTableFromGatewayRequest
	GetGatewayId() *string
	SetGatewayType(v string) *DissociateRouteTableFromGatewayRequest
	GetGatewayType() *string
	SetOwnerAccount(v string) *DissociateRouteTableFromGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DissociateRouteTableFromGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DissociateRouteTableFromGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DissociateRouteTableFromGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DissociateRouteTableFromGatewayRequest
	GetResourceOwnerId() *int64
	SetRouteTableId(v string) *DissociateRouteTableFromGatewayRequest
	GetRouteTableId() *string
}

type DissociateRouteTableFromGatewayRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not set this parameter, the system automatically uses **RequestId*	- as **ClientToken**. **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// 	- **true**: prechecks the request without performing the operation. The system prechecks the required parameters, request syntax, and limits. If the request fails to pass the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the request. After the request passes the precheck, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPv4 gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv4gw-5tsnc6s4ogsedtp3k****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The type of a gateway to be disassociated from a route table.
	//
	// example:
	//
	// Ipv4Gateway
	GatewayType  *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IPv4 gateway from which you want to disassociate the gateway route table.
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

func (s DissociateRouteTableFromGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateRouteTableFromGatewayRequest) GoString() string {
	return s.String()
}

func (s *DissociateRouteTableFromGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DissociateRouteTableFromGatewayRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DissociateRouteTableFromGatewayRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *DissociateRouteTableFromGatewayRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *DissociateRouteTableFromGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DissociateRouteTableFromGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DissociateRouteTableFromGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DissociateRouteTableFromGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DissociateRouteTableFromGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DissociateRouteTableFromGatewayRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DissociateRouteTableFromGatewayRequest) SetClientToken(v string) *DissociateRouteTableFromGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateRouteTableFromGatewayRequest) SetDryRun(v bool) *DissociateRouteTableFromGatewayRequest {
	s.DryRun = &v
	return s
}

func (s *DissociateRouteTableFromGatewayRequest) SetGatewayId(v string) *DissociateRouteTableFromGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *DissociateRouteTableFromGatewayRequest) SetGatewayType(v string) *DissociateRouteTableFromGatewayRequest {
	s.GatewayType = &v
	return s
}

func (s *DissociateRouteTableFromGatewayRequest) SetOwnerAccount(v string) *DissociateRouteTableFromGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DissociateRouteTableFromGatewayRequest) SetOwnerId(v int64) *DissociateRouteTableFromGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *DissociateRouteTableFromGatewayRequest) SetRegionId(v string) *DissociateRouteTableFromGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *DissociateRouteTableFromGatewayRequest) SetResourceOwnerAccount(v string) *DissociateRouteTableFromGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DissociateRouteTableFromGatewayRequest) SetResourceOwnerId(v int64) *DissociateRouteTableFromGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DissociateRouteTableFromGatewayRequest) SetRouteTableId(v string) *DissociateRouteTableFromGatewayRequest {
	s.RouteTableId = &v
	return s
}

func (s *DissociateRouteTableFromGatewayRequest) Validate() error {
	return dara.Validate(s)
}
