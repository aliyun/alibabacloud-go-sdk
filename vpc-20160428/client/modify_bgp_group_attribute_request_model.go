// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBgpGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthKey(v string) *ModifyBgpGroupAttributeRequest
	GetAuthKey() *string
	SetBgpGroupId(v string) *ModifyBgpGroupAttributeRequest
	GetBgpGroupId() *string
	SetClearAuthKey(v bool) *ModifyBgpGroupAttributeRequest
	GetClearAuthKey() *bool
	SetClientToken(v string) *ModifyBgpGroupAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyBgpGroupAttributeRequest
	GetDescription() *string
	SetIsFakeAsn(v bool) *ModifyBgpGroupAttributeRequest
	GetIsFakeAsn() *bool
	SetLocalAsn(v int64) *ModifyBgpGroupAttributeRequest
	GetLocalAsn() *int64
	SetName(v string) *ModifyBgpGroupAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyBgpGroupAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyBgpGroupAttributeRequest
	GetOwnerId() *int64
	SetPeerAsn(v int64) *ModifyBgpGroupAttributeRequest
	GetPeerAsn() *int64
	SetRegionId(v string) *ModifyBgpGroupAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyBgpGroupAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyBgpGroupAttributeRequest
	GetResourceOwnerId() *int64
	SetRouteQuota(v int32) *ModifyBgpGroupAttributeRequest
	GetRouteQuota() *int32
}

type ModifyBgpGroupAttributeRequest struct {
	// The authentication key of the BGP group.
	//
	// example:
	//
	// !PWZ2****
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The BGP group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bgpg-wz9f62v4fbg2g****
	BgpGroupId *string `json:"BgpGroupId,omitempty" xml:"BgpGroupId,omitempty"`
	// Specifies whether to clear the secret key. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ClearAuthKey *bool `json:"ClearAuthKey,omitempty" xml:"ClearAuthKey,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The BGP group description.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// BGP
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to use a fake AS number. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// > A router that runs BGP typically belongs to only one AS. If you need to replace an AS with a new one, but you cannot immediately modify BGP configurations due to business requirements, you can specify a fake AS number to establish a connection with the local end. This ensures service continuity in scenarios such as AS migration or AS merging.
	//
	// example:
	//
	// false
	IsFakeAsn *bool `json:"IsFakeAsn,omitempty" xml:"IsFakeAsn,omitempty"`
	// The custom autonomous system number (ASN) of the BGP on the Alibaba Cloud side. Valid values:
	//
	// 	- **45104**
	//
	// 	- **64512~65534**
	//
	// 	- **4200000000~4294967294**
	//
	// >  **65025*	- is reserved by Alibaba Cloud. Alibaba Cloud uses **45104*	- as the **local ASN*	- by default. Custom **local ASNs*	- may cause loops in multi-line scenarios. Proceed with caution.
	//
	// example:
	//
	// 45104
	LocalAsn *int64 `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// The BGP group name.
	//
	// The name must be 2 to 128 characters in length, and can contain digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ASN of the gateway device in the data center.
	//
	// example:
	//
	// 1****
	PeerAsn *int64 `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	// The region ID of the BGP group.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The maximum number of routes supported by a BGP peer. Default value: **110**.
	//
	// example:
	//
	// 110
	RouteQuota *int32 `json:"RouteQuota,omitempty" xml:"RouteQuota,omitempty"`
}

func (s ModifyBgpGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBgpGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyBgpGroupAttributeRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *ModifyBgpGroupAttributeRequest) GetBgpGroupId() *string {
	return s.BgpGroupId
}

func (s *ModifyBgpGroupAttributeRequest) GetClearAuthKey() *bool {
	return s.ClearAuthKey
}

func (s *ModifyBgpGroupAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyBgpGroupAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyBgpGroupAttributeRequest) GetIsFakeAsn() *bool {
	return s.IsFakeAsn
}

func (s *ModifyBgpGroupAttributeRequest) GetLocalAsn() *int64 {
	return s.LocalAsn
}

func (s *ModifyBgpGroupAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyBgpGroupAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyBgpGroupAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyBgpGroupAttributeRequest) GetPeerAsn() *int64 {
	return s.PeerAsn
}

func (s *ModifyBgpGroupAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyBgpGroupAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyBgpGroupAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyBgpGroupAttributeRequest) GetRouteQuota() *int32 {
	return s.RouteQuota
}

func (s *ModifyBgpGroupAttributeRequest) SetAuthKey(v string) *ModifyBgpGroupAttributeRequest {
	s.AuthKey = &v
	return s
}

func (s *ModifyBgpGroupAttributeRequest) SetBgpGroupId(v string) *ModifyBgpGroupAttributeRequest {
	s.BgpGroupId = &v
	return s
}

func (s *ModifyBgpGroupAttributeRequest) SetClearAuthKey(v bool) *ModifyBgpGroupAttributeRequest {
	s.ClearAuthKey = &v
	return s
}

func (s *ModifyBgpGroupAttributeRequest) SetClientToken(v string) *ModifyBgpGroupAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyBgpGroupAttributeRequest) SetDescription(v string) *ModifyBgpGroupAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyBgpGroupAttributeRequest) SetIsFakeAsn(v bool) *ModifyBgpGroupAttributeRequest {
	s.IsFakeAsn = &v
	return s
}

func (s *ModifyBgpGroupAttributeRequest) SetLocalAsn(v int64) *ModifyBgpGroupAttributeRequest {
	s.LocalAsn = &v
	return s
}

func (s *ModifyBgpGroupAttributeRequest) SetName(v string) *ModifyBgpGroupAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyBgpGroupAttributeRequest) SetOwnerAccount(v string) *ModifyBgpGroupAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBgpGroupAttributeRequest) SetOwnerId(v int64) *ModifyBgpGroupAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBgpGroupAttributeRequest) SetPeerAsn(v int64) *ModifyBgpGroupAttributeRequest {
	s.PeerAsn = &v
	return s
}

func (s *ModifyBgpGroupAttributeRequest) SetRegionId(v string) *ModifyBgpGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBgpGroupAttributeRequest) SetResourceOwnerAccount(v string) *ModifyBgpGroupAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBgpGroupAttributeRequest) SetResourceOwnerId(v int64) *ModifyBgpGroupAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBgpGroupAttributeRequest) SetRouteQuota(v int32) *ModifyBgpGroupAttributeRequest {
	s.RouteQuota = &v
	return s
}

func (s *ModifyBgpGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}
