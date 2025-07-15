// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcGatewayEndpointAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateVpcGatewayEndpointAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateVpcGatewayEndpointAttributeRequest
	GetDryRun() *bool
	SetEndpointDescription(v string) *UpdateVpcGatewayEndpointAttributeRequest
	GetEndpointDescription() *string
	SetEndpointId(v string) *UpdateVpcGatewayEndpointAttributeRequest
	GetEndpointId() *string
	SetEndpointName(v string) *UpdateVpcGatewayEndpointAttributeRequest
	GetEndpointName() *string
	SetOwnerAccount(v string) *UpdateVpcGatewayEndpointAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateVpcGatewayEndpointAttributeRequest
	GetOwnerId() *int64
	SetPolicyDocument(v string) *UpdateVpcGatewayEndpointAttributeRequest
	GetPolicyDocument() *string
	SetRegionId(v string) *UpdateVpcGatewayEndpointAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateVpcGatewayEndpointAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateVpcGatewayEndpointAttributeRequest
	GetResourceOwnerId() *int64
}

type UpdateVpcGatewayEndpointAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among all requests. The token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId*	- as **ClientToken**. The value of **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks your AccessKey pair, the RAM user permissions, and the required parameters If the request fails the dry run, the corresponding error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The new description of the gateway endpoint.
	//
	// The description must be 1 to 255 characters in length.
	//
	// example:
	//
	// updateendpoint
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The ID of the gateway endpoint that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpce-bp1w1dmdqjpwul0v3****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The new name of the gateway endpoint.
	//
	// The name must be 1 to 128 characters in length.
	//
	// example:
	//
	// update
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The access policy for the cloud service.
	//
	// For more information about the syntax and structure of the access policy, see [Policy syntax and structure](https://help.aliyun.com/document_detail/93739.html).
	//
	// example:
	//
	// {   "Version" : "1",   "Statement" : [ {     "Effect" : "Allow",     "Resource" : [ "*" ],     "Action" : [ "*" ],     "Principal" : [ "*" ]   } ] }
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
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

func (s UpdateVpcGatewayEndpointAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcGatewayEndpointAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) GetEndpointDescription() *string {
	return s.EndpointDescription
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) GetEndpointName() *string {
	return s.EndpointName
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) SetClientToken(v string) *UpdateVpcGatewayEndpointAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) SetDryRun(v bool) *UpdateVpcGatewayEndpointAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) SetEndpointDescription(v string) *UpdateVpcGatewayEndpointAttributeRequest {
	s.EndpointDescription = &v
	return s
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) SetEndpointId(v string) *UpdateVpcGatewayEndpointAttributeRequest {
	s.EndpointId = &v
	return s
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) SetEndpointName(v string) *UpdateVpcGatewayEndpointAttributeRequest {
	s.EndpointName = &v
	return s
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) SetOwnerAccount(v string) *UpdateVpcGatewayEndpointAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) SetOwnerId(v int64) *UpdateVpcGatewayEndpointAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) SetPolicyDocument(v string) *UpdateVpcGatewayEndpointAttributeRequest {
	s.PolicyDocument = &v
	return s
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) SetRegionId(v string) *UpdateVpcGatewayEndpointAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) SetResourceOwnerAccount(v string) *UpdateVpcGatewayEndpointAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) SetResourceOwnerId(v int64) *UpdateVpcGatewayEndpointAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateVpcGatewayEndpointAttributeRequest) Validate() error {
	return dara.Validate(s)
}
