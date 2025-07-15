// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcGatewayEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVpcGatewayEndpointRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteVpcGatewayEndpointRequest
	GetDryRun() *bool
	SetEndpointId(v string) *DeleteVpcGatewayEndpointRequest
	GetEndpointId() *string
	SetOwnerAccount(v string) *DeleteVpcGatewayEndpointRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVpcGatewayEndpointRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVpcGatewayEndpointRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVpcGatewayEndpointRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVpcGatewayEndpointRequest
	GetResourceOwnerId() *int64
}

type DeleteVpcGatewayEndpointRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the request for potential issues, including the AccessKey pair, the permissions of the RAM user, and the required parameters. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the gateway endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpce-bp1w1dmdqjpwul0v3****
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
}

func (s DeleteVpcGatewayEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcGatewayEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcGatewayEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVpcGatewayEndpointRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteVpcGatewayEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DeleteVpcGatewayEndpointRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVpcGatewayEndpointRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVpcGatewayEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVpcGatewayEndpointRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVpcGatewayEndpointRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVpcGatewayEndpointRequest) SetClientToken(v string) *DeleteVpcGatewayEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVpcGatewayEndpointRequest) SetDryRun(v bool) *DeleteVpcGatewayEndpointRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteVpcGatewayEndpointRequest) SetEndpointId(v string) *DeleteVpcGatewayEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *DeleteVpcGatewayEndpointRequest) SetOwnerAccount(v string) *DeleteVpcGatewayEndpointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVpcGatewayEndpointRequest) SetOwnerId(v int64) *DeleteVpcGatewayEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVpcGatewayEndpointRequest) SetRegionId(v string) *DeleteVpcGatewayEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVpcGatewayEndpointRequest) SetResourceOwnerAccount(v string) *DeleteVpcGatewayEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVpcGatewayEndpointRequest) SetResourceOwnerId(v int64) *DeleteVpcGatewayEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVpcGatewayEndpointRequest) Validate() error {
	return dara.Validate(s)
}
