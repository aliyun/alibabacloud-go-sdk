// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIPv6TranslatorAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyIPv6TranslatorAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyIPv6TranslatorAttributeRequest
	GetDescription() *string
	SetIpv6TranslatorId(v string) *ModifyIPv6TranslatorAttributeRequest
	GetIpv6TranslatorId() *string
	SetName(v string) *ModifyIPv6TranslatorAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyIPv6TranslatorAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyIPv6TranslatorAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyIPv6TranslatorAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyIPv6TranslatorAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyIPv6TranslatorAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyIPv6TranslatorAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// sha1111
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of IPv6 Translation Service. This parameter is empty by default. It must be 2 to 100 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It cannot start with http:// or [https://](https://。).
	//
	// example:
	//
	// instancedescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the IPv6 Translation Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6trans-bp1858ys****
	Ipv6TranslatorId *string `json:"Ipv6TranslatorId,omitempty" xml:"Ipv6TranslatorId,omitempty"`
	// The name of the IPv6 Translation Service instance. The default name is the instance ID. It must be 2 to 100 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It cannot start with http:// or [https://](https://。).
	//
	// example:
	//
	// instancename
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region of the IPv6 Translation Service instance. You can call the **DescribeRegions*	- operation to query the most recent region list.
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

func (s ModifyIPv6TranslatorAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIPv6TranslatorAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyIPv6TranslatorAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyIPv6TranslatorAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyIPv6TranslatorAttributeRequest) GetIpv6TranslatorId() *string {
	return s.Ipv6TranslatorId
}

func (s *ModifyIPv6TranslatorAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyIPv6TranslatorAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyIPv6TranslatorAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyIPv6TranslatorAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyIPv6TranslatorAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyIPv6TranslatorAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyIPv6TranslatorAttributeRequest) SetClientToken(v string) *ModifyIPv6TranslatorAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyIPv6TranslatorAttributeRequest) SetDescription(v string) *ModifyIPv6TranslatorAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyIPv6TranslatorAttributeRequest) SetIpv6TranslatorId(v string) *ModifyIPv6TranslatorAttributeRequest {
	s.Ipv6TranslatorId = &v
	return s
}

func (s *ModifyIPv6TranslatorAttributeRequest) SetName(v string) *ModifyIPv6TranslatorAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyIPv6TranslatorAttributeRequest) SetOwnerAccount(v string) *ModifyIPv6TranslatorAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyIPv6TranslatorAttributeRequest) SetOwnerId(v int64) *ModifyIPv6TranslatorAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyIPv6TranslatorAttributeRequest) SetRegionId(v string) *ModifyIPv6TranslatorAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyIPv6TranslatorAttributeRequest) SetResourceOwnerAccount(v string) *ModifyIPv6TranslatorAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyIPv6TranslatorAttributeRequest) SetResourceOwnerId(v int64) *ModifyIPv6TranslatorAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyIPv6TranslatorAttributeRequest) Validate() error {
	return dara.Validate(s)
}
