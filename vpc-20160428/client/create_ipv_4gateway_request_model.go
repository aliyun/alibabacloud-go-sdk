// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpv4GatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateIpv4GatewayRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateIpv4GatewayRequest
	GetDryRun() *bool
	SetIpv4GatewayDescription(v string) *CreateIpv4GatewayRequest
	GetIpv4GatewayDescription() *string
	SetIpv4GatewayName(v string) *CreateIpv4GatewayRequest
	GetIpv4GatewayName() *string
	SetOwnerAccount(v string) *CreateIpv4GatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateIpv4GatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateIpv4GatewayRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateIpv4GatewayRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateIpv4GatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateIpv4GatewayRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateIpv4GatewayRequestTag) *CreateIpv4GatewayRequest
	GetTag() []*CreateIpv4GatewayRequestTag
	SetVpcId(v string) *CreateIpv4GatewayRequest
	GetVpcId() *string
}

type CreateIpv4GatewayRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system sets **ClientToken*	- to the value of **RequestId**. The value of **RequestId*	- for each API request is different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the IPv4 gateway.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Ipv4GatewayDescription *string `json:"Ipv4GatewayDescription,omitempty" xml:"Ipv4GatewayDescription,omitempty"`
	// The name of the IPv4 gateway.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// ipv4
	Ipv4GatewayName *string `json:"Ipv4GatewayName,omitempty" xml:"Ipv4GatewayName,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where you want to create the IPv4 gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-6
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the resource.
	Tag []*CreateIpv4GatewayRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC where you want to create the IPv4 gateway.
	//
	// You can create only one IPv4 gateway in a VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-5tss06uvoyps5xoya****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateIpv4GatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIpv4GatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateIpv4GatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateIpv4GatewayRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateIpv4GatewayRequest) GetIpv4GatewayDescription() *string {
	return s.Ipv4GatewayDescription
}

func (s *CreateIpv4GatewayRequest) GetIpv4GatewayName() *string {
	return s.Ipv4GatewayName
}

func (s *CreateIpv4GatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateIpv4GatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateIpv4GatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIpv4GatewayRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateIpv4GatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateIpv4GatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateIpv4GatewayRequest) GetTag() []*CreateIpv4GatewayRequestTag {
	return s.Tag
}

func (s *CreateIpv4GatewayRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateIpv4GatewayRequest) SetClientToken(v string) *CreateIpv4GatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpv4GatewayRequest) SetDryRun(v bool) *CreateIpv4GatewayRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIpv4GatewayRequest) SetIpv4GatewayDescription(v string) *CreateIpv4GatewayRequest {
	s.Ipv4GatewayDescription = &v
	return s
}

func (s *CreateIpv4GatewayRequest) SetIpv4GatewayName(v string) *CreateIpv4GatewayRequest {
	s.Ipv4GatewayName = &v
	return s
}

func (s *CreateIpv4GatewayRequest) SetOwnerAccount(v string) *CreateIpv4GatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateIpv4GatewayRequest) SetOwnerId(v int64) *CreateIpv4GatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIpv4GatewayRequest) SetRegionId(v string) *CreateIpv4GatewayRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIpv4GatewayRequest) SetResourceGroupId(v string) *CreateIpv4GatewayRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateIpv4GatewayRequest) SetResourceOwnerAccount(v string) *CreateIpv4GatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateIpv4GatewayRequest) SetResourceOwnerId(v int64) *CreateIpv4GatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateIpv4GatewayRequest) SetTag(v []*CreateIpv4GatewayRequestTag) *CreateIpv4GatewayRequest {
	s.Tag = v
	return s
}

func (s *CreateIpv4GatewayRequest) SetVpcId(v string) *CreateIpv4GatewayRequest {
	s.VpcId = &v
	return s
}

func (s *CreateIpv4GatewayRequest) Validate() error {
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

type CreateIpv4GatewayRequestTag struct {
	// The key of tag N to add to the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIpv4GatewayRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateIpv4GatewayRequestTag) GoString() string {
	return s.String()
}

func (s *CreateIpv4GatewayRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateIpv4GatewayRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateIpv4GatewayRequestTag) SetKey(v string) *CreateIpv4GatewayRequestTag {
	s.Key = &v
	return s
}

func (s *CreateIpv4GatewayRequestTag) SetValue(v string) *CreateIpv4GatewayRequestTag {
	s.Value = &v
	return s
}

func (s *CreateIpv4GatewayRequestTag) Validate() error {
	return dara.Validate(s)
}
