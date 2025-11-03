// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateIpv6AddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressType(v string) *AllocateIpv6AddressRequest
	GetAddressType() *string
	SetClientToken(v string) *AllocateIpv6AddressRequest
	GetClientToken() *string
	SetDryRun(v bool) *AllocateIpv6AddressRequest
	GetDryRun() *bool
	SetIpv6Address(v string) *AllocateIpv6AddressRequest
	GetIpv6Address() *string
	SetIpv6AddressDescription(v string) *AllocateIpv6AddressRequest
	GetIpv6AddressDescription() *string
	SetIpv6AddressName(v string) *AllocateIpv6AddressRequest
	GetIpv6AddressName() *string
	SetOwnerAccount(v string) *AllocateIpv6AddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateIpv6AddressRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AllocateIpv6AddressRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *AllocateIpv6AddressRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *AllocateIpv6AddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateIpv6AddressRequest
	GetResourceOwnerId() *int64
	SetTag(v []*AllocateIpv6AddressRequestTag) *AllocateIpv6AddressRequest
	GetTag() []*AllocateIpv6AddressRequestTag
	SetVSwitchId(v string) *AllocateIpv6AddressRequest
	GetVSwitchId() *string
}

type AllocateIpv6AddressRequest struct {
	// The type of the IPv6 address. Valid values:
	//
	// 	- IPv6Address (default): an IPv6 address.
	//
	// 	- IPv6Prefix: an IPv6 CIDR block.
	//
	// example:
	//
	// IPv6Address
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- false (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IPv6 address. The IPv6 address must be an idle one that falls within the vSwitch CIDR block.
	//
	// example:
	//
	// 2408:XXXX:153:3921:851c:c435:7b12:1c5f
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	// The description of the IPv6 address.
	//
	// example:
	//
	// ipv6-description
	Ipv6AddressDescription *string `json:"Ipv6AddressDescription,omitempty" xml:"Ipv6AddressDescription,omitempty"`
	// The name of the IPv6 address.
	//
	// example:
	//
	// ipv6-name
	Ipv6AddressName *string `json:"Ipv6AddressName,omitempty" xml:"Ipv6AddressName,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. For more information about resource groups, see related documentation.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag list.
	Tag []*AllocateIpv6AddressRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the IPv6 address belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-asdfjlnaue4g****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s AllocateIpv6AddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateIpv6AddressRequest) GoString() string {
	return s.String()
}

func (s *AllocateIpv6AddressRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *AllocateIpv6AddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AllocateIpv6AddressRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AllocateIpv6AddressRequest) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *AllocateIpv6AddressRequest) GetIpv6AddressDescription() *string {
	return s.Ipv6AddressDescription
}

func (s *AllocateIpv6AddressRequest) GetIpv6AddressName() *string {
	return s.Ipv6AddressName
}

func (s *AllocateIpv6AddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateIpv6AddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateIpv6AddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocateIpv6AddressRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AllocateIpv6AddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateIpv6AddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateIpv6AddressRequest) GetTag() []*AllocateIpv6AddressRequestTag {
	return s.Tag
}

func (s *AllocateIpv6AddressRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *AllocateIpv6AddressRequest) SetAddressType(v string) *AllocateIpv6AddressRequest {
	s.AddressType = &v
	return s
}

func (s *AllocateIpv6AddressRequest) SetClientToken(v string) *AllocateIpv6AddressRequest {
	s.ClientToken = &v
	return s
}

func (s *AllocateIpv6AddressRequest) SetDryRun(v bool) *AllocateIpv6AddressRequest {
	s.DryRun = &v
	return s
}

func (s *AllocateIpv6AddressRequest) SetIpv6Address(v string) *AllocateIpv6AddressRequest {
	s.Ipv6Address = &v
	return s
}

func (s *AllocateIpv6AddressRequest) SetIpv6AddressDescription(v string) *AllocateIpv6AddressRequest {
	s.Ipv6AddressDescription = &v
	return s
}

func (s *AllocateIpv6AddressRequest) SetIpv6AddressName(v string) *AllocateIpv6AddressRequest {
	s.Ipv6AddressName = &v
	return s
}

func (s *AllocateIpv6AddressRequest) SetOwnerAccount(v string) *AllocateIpv6AddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateIpv6AddressRequest) SetOwnerId(v int64) *AllocateIpv6AddressRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateIpv6AddressRequest) SetRegionId(v string) *AllocateIpv6AddressRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateIpv6AddressRequest) SetResourceGroupId(v string) *AllocateIpv6AddressRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AllocateIpv6AddressRequest) SetResourceOwnerAccount(v string) *AllocateIpv6AddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateIpv6AddressRequest) SetResourceOwnerId(v int64) *AllocateIpv6AddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateIpv6AddressRequest) SetTag(v []*AllocateIpv6AddressRequestTag) *AllocateIpv6AddressRequest {
	s.Tag = v
	return s
}

func (s *AllocateIpv6AddressRequest) SetVSwitchId(v string) *AllocateIpv6AddressRequest {
	s.VSwitchId = &v
	return s
}

func (s *AllocateIpv6AddressRequest) Validate() error {
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

type AllocateIpv6AddressRequestTag struct {
	// The key of tag N to add to the resource. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AllocateIpv6AddressRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AllocateIpv6AddressRequestTag) GoString() string {
	return s.String()
}

func (s *AllocateIpv6AddressRequestTag) GetKey() *string {
	return s.Key
}

func (s *AllocateIpv6AddressRequestTag) GetValue() *string {
	return s.Value
}

func (s *AllocateIpv6AddressRequestTag) SetKey(v string) *AllocateIpv6AddressRequestTag {
	s.Key = &v
	return s
}

func (s *AllocateIpv6AddressRequestTag) SetValue(v string) *AllocateIpv6AddressRequestTag {
	s.Value = &v
	return s
}

func (s *AllocateIpv6AddressRequestTag) Validate() error {
	return dara.Validate(s)
}
