// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpv6GatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateIpv6GatewayRequest
	GetClientToken() *string
	SetDescription(v string) *CreateIpv6GatewayRequest
	GetDescription() *string
	SetName(v string) *CreateIpv6GatewayRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateIpv6GatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateIpv6GatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateIpv6GatewayRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateIpv6GatewayRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateIpv6GatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateIpv6GatewayRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateIpv6GatewayRequestTag) *CreateIpv6GatewayRequest
	GetTag() []*CreateIpv6GatewayRequestTag
	SetVpcId(v string) *CreateIpv6GatewayRequest
	GetVpcId() *string
}

type CreateIpv6GatewayRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the IPv6 gateway.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// ipv6gatewayforVPC1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the IPv6 gateway.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// ipv6GW
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IPv6 gateway. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the IPv6 gateway belongs.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*CreateIpv6GatewayRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC in which you want to create the IPv6 gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-123sedrfswd23****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateIpv6GatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIpv6GatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateIpv6GatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateIpv6GatewayRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateIpv6GatewayRequest) GetName() *string {
	return s.Name
}

func (s *CreateIpv6GatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateIpv6GatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateIpv6GatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIpv6GatewayRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateIpv6GatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateIpv6GatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateIpv6GatewayRequest) GetTag() []*CreateIpv6GatewayRequestTag {
	return s.Tag
}

func (s *CreateIpv6GatewayRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateIpv6GatewayRequest) SetClientToken(v string) *CreateIpv6GatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpv6GatewayRequest) SetDescription(v string) *CreateIpv6GatewayRequest {
	s.Description = &v
	return s
}

func (s *CreateIpv6GatewayRequest) SetName(v string) *CreateIpv6GatewayRequest {
	s.Name = &v
	return s
}

func (s *CreateIpv6GatewayRequest) SetOwnerAccount(v string) *CreateIpv6GatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateIpv6GatewayRequest) SetOwnerId(v int64) *CreateIpv6GatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIpv6GatewayRequest) SetRegionId(v string) *CreateIpv6GatewayRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIpv6GatewayRequest) SetResourceGroupId(v string) *CreateIpv6GatewayRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateIpv6GatewayRequest) SetResourceOwnerAccount(v string) *CreateIpv6GatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateIpv6GatewayRequest) SetResourceOwnerId(v int64) *CreateIpv6GatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateIpv6GatewayRequest) SetTag(v []*CreateIpv6GatewayRequestTag) *CreateIpv6GatewayRequest {
	s.Tag = v
	return s
}

func (s *CreateIpv6GatewayRequest) SetVpcId(v string) *CreateIpv6GatewayRequest {
	s.VpcId = &v
	return s
}

func (s *CreateIpv6GatewayRequest) Validate() error {
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

type CreateIpv6GatewayRequestTag struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIpv6GatewayRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateIpv6GatewayRequestTag) GoString() string {
	return s.String()
}

func (s *CreateIpv6GatewayRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateIpv6GatewayRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateIpv6GatewayRequestTag) SetKey(v string) *CreateIpv6GatewayRequestTag {
	s.Key = &v
	return s
}

func (s *CreateIpv6GatewayRequestTag) SetValue(v string) *CreateIpv6GatewayRequestTag {
	s.Value = &v
	return s
}

func (s *CreateIpv6GatewayRequestTag) Validate() error {
	return dara.Validate(s)
}
