// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcGatewayEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateVpcGatewayEndpointRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateVpcGatewayEndpointRequest
	GetDryRun() *bool
	SetEndpointDescription(v string) *CreateVpcGatewayEndpointRequest
	GetEndpointDescription() *string
	SetEndpointName(v string) *CreateVpcGatewayEndpointRequest
	GetEndpointName() *string
	SetOwnerAccount(v string) *CreateVpcGatewayEndpointRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVpcGatewayEndpointRequest
	GetOwnerId() *int64
	SetPolicyDocument(v string) *CreateVpcGatewayEndpointRequest
	GetPolicyDocument() *string
	SetRegionId(v string) *CreateVpcGatewayEndpointRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateVpcGatewayEndpointRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateVpcGatewayEndpointRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVpcGatewayEndpointRequest
	GetResourceOwnerId() *int64
	SetServiceName(v string) *CreateVpcGatewayEndpointRequest
	GetServiceName() *string
	SetTag(v []*CreateVpcGatewayEndpointRequestTag) *CreateVpcGatewayEndpointRequest
	GetTag() []*CreateVpcGatewayEndpointRequestTag
	SetVpcId(v string) *CreateVpcGatewayEndpointRequest
	GetVpcId() *string
}

type CreateVpcGatewayEndpointRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: performs only a dry run and does not create the gateway endpoint. The system checks whether your AccessKey pair is valid, the authorization of the Resource Access Management (RAM) user, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the dry run, a 2xx HTTP status code is returned and the gateway endpoint is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the gateway endpoint.
	//
	// The description must be 1 to 255 characters in length.
	//
	// example:
	//
	// test
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The name of the gateway endpoint.
	//
	// The name must be 1 to 128 characters in length.
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The access policy for the cloud service.
	//
	// For more information about the syntax and structure of access policies, see [Policy syntax and structure](https://help.aliyun.com/document_detail/93739.html).
	//
	// >
	//
	// > - This parameter is required if the selected endpoint service supports access policies.
	//
	// > - This parameter must be left empty if the selected endpoint service does not support access policies.
	//
	// > - You can call the [ListVpcEndpointServicesByEndUser](https://help.aliyun.com/document_detail/448920.html) operation to check whether the endpoint service supports access policies. If the value of SupportPolicy is true, the endpoint service supports access policies. Otherwise, it does not.
	//
	// example:
	//
	// {   "Version" : "1",   "Statement" : [ {     "Effect" : "Allow",     "Resource" : [ "*" ],     "Action" : [ "*" ],     "Principal" : [ "*" ]   } ] }
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The region ID of the gateway endpoint that you want to create.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the gateway endpoint belongs.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the endpoint service.
	//
	// This parameter is required.
	//
	// example:
	//
	// com.aliyun.cn-hangzhou.oss
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The tags of the resource.
	Tag []*CreateVpcGatewayEndpointRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) where you want to create the gateway endpoint.
	//
	// The VPC must be in the same region as the gateway endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1gsk7h12ew7oegk****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateVpcGatewayEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcGatewayEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcGatewayEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVpcGatewayEndpointRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVpcGatewayEndpointRequest) GetEndpointDescription() *string {
	return s.EndpointDescription
}

func (s *CreateVpcGatewayEndpointRequest) GetEndpointName() *string {
	return s.EndpointName
}

func (s *CreateVpcGatewayEndpointRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVpcGatewayEndpointRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVpcGatewayEndpointRequest) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *CreateVpcGatewayEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpcGatewayEndpointRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVpcGatewayEndpointRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVpcGatewayEndpointRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVpcGatewayEndpointRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateVpcGatewayEndpointRequest) GetTag() []*CreateVpcGatewayEndpointRequestTag {
	return s.Tag
}

func (s *CreateVpcGatewayEndpointRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateVpcGatewayEndpointRequest) SetClientToken(v string) *CreateVpcGatewayEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpcGatewayEndpointRequest) SetDryRun(v bool) *CreateVpcGatewayEndpointRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVpcGatewayEndpointRequest) SetEndpointDescription(v string) *CreateVpcGatewayEndpointRequest {
	s.EndpointDescription = &v
	return s
}

func (s *CreateVpcGatewayEndpointRequest) SetEndpointName(v string) *CreateVpcGatewayEndpointRequest {
	s.EndpointName = &v
	return s
}

func (s *CreateVpcGatewayEndpointRequest) SetOwnerAccount(v string) *CreateVpcGatewayEndpointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVpcGatewayEndpointRequest) SetOwnerId(v int64) *CreateVpcGatewayEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVpcGatewayEndpointRequest) SetPolicyDocument(v string) *CreateVpcGatewayEndpointRequest {
	s.PolicyDocument = &v
	return s
}

func (s *CreateVpcGatewayEndpointRequest) SetRegionId(v string) *CreateVpcGatewayEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpcGatewayEndpointRequest) SetResourceGroupId(v string) *CreateVpcGatewayEndpointRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVpcGatewayEndpointRequest) SetResourceOwnerAccount(v string) *CreateVpcGatewayEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVpcGatewayEndpointRequest) SetResourceOwnerId(v int64) *CreateVpcGatewayEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVpcGatewayEndpointRequest) SetServiceName(v string) *CreateVpcGatewayEndpointRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateVpcGatewayEndpointRequest) SetTag(v []*CreateVpcGatewayEndpointRequestTag) *CreateVpcGatewayEndpointRequest {
	s.Tag = v
	return s
}

func (s *CreateVpcGatewayEndpointRequest) SetVpcId(v string) *CreateVpcGatewayEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateVpcGatewayEndpointRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateVpcGatewayEndpointRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys. If you specify this parameter, you cannot specify an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. If you specify this parameter, you can specify an empty string.
	//
	// A tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVpcGatewayEndpointRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcGatewayEndpointRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVpcGatewayEndpointRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVpcGatewayEndpointRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVpcGatewayEndpointRequestTag) SetKey(v string) *CreateVpcGatewayEndpointRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVpcGatewayEndpointRequestTag) SetValue(v string) *CreateVpcGatewayEndpointRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVpcGatewayEndpointRequestTag) Validate() error {
	return dara.Validate(s)
}
