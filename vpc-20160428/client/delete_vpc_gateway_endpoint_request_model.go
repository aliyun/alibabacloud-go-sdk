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
	// Ensures the idempotence of the request. You can generate a parameter value on your client to ensure that the value is unique across different requests. **ClientToken*	- supports only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId of each API request may be different.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: performs a dry run and does not delete the gateway endpoint. The system checks whether your AccessKey pair is valid, the authorization of the Resource Access Management (RAM) user, and whether the required parameters are specified. If the request fails the dry run, the corresponding error is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, a 2xx HTTP status code is returned and the gateway endpoint is deleted.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the gateway endpoint instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpce-bp1w1dmdqjpwul0v3****
	EndpointId   *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the gateway endpoint resides.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
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
