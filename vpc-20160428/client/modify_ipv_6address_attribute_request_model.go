// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpv6AddressAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyIpv6AddressAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyIpv6AddressAttributeRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyIpv6AddressAttributeRequest
	GetDryRun() *bool
	SetIpv6AddressId(v string) *ModifyIpv6AddressAttributeRequest
	GetIpv6AddressId() *string
	SetName(v string) *ModifyIpv6AddressAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyIpv6AddressAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyIpv6AddressAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyIpv6AddressAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyIpv6AddressAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyIpv6AddressAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyIpv6AddressAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// d7d24a21-f4ba-4454-9173-b3828dae492b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the IPv6 address.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPv6 address.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6-hp32vv2klzw4yerdf****
	Ipv6AddressId *string `json:"Ipv6AddressId,omitempty" xml:"Ipv6AddressId,omitempty"`
	// The name of the IPv6 address.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IPv6 address. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyIpv6AddressAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpv6AddressAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpv6AddressAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyIpv6AddressAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyIpv6AddressAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyIpv6AddressAttributeRequest) GetIpv6AddressId() *string {
	return s.Ipv6AddressId
}

func (s *ModifyIpv6AddressAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyIpv6AddressAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyIpv6AddressAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyIpv6AddressAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyIpv6AddressAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyIpv6AddressAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyIpv6AddressAttributeRequest) SetClientToken(v string) *ModifyIpv6AddressAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyIpv6AddressAttributeRequest) SetDescription(v string) *ModifyIpv6AddressAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyIpv6AddressAttributeRequest) SetDryRun(v bool) *ModifyIpv6AddressAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyIpv6AddressAttributeRequest) SetIpv6AddressId(v string) *ModifyIpv6AddressAttributeRequest {
	s.Ipv6AddressId = &v
	return s
}

func (s *ModifyIpv6AddressAttributeRequest) SetName(v string) *ModifyIpv6AddressAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyIpv6AddressAttributeRequest) SetOwnerAccount(v string) *ModifyIpv6AddressAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyIpv6AddressAttributeRequest) SetOwnerId(v int64) *ModifyIpv6AddressAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyIpv6AddressAttributeRequest) SetRegionId(v string) *ModifyIpv6AddressAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyIpv6AddressAttributeRequest) SetResourceOwnerAccount(v string) *ModifyIpv6AddressAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyIpv6AddressAttributeRequest) SetResourceOwnerId(v int64) *ModifyIpv6AddressAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyIpv6AddressAttributeRequest) Validate() error {
	return dara.Validate(s)
}
