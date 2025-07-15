// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBgpGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthKey(v string) *CreateBgpGroupRequest
	GetAuthKey() *string
	SetClientToken(v string) *CreateBgpGroupRequest
	GetClientToken() *string
	SetDescription(v string) *CreateBgpGroupRequest
	GetDescription() *string
	SetIpVersion(v string) *CreateBgpGroupRequest
	GetIpVersion() *string
	SetIsFakeAsn(v bool) *CreateBgpGroupRequest
	GetIsFakeAsn() *bool
	SetLocalAsn(v int64) *CreateBgpGroupRequest
	GetLocalAsn() *int64
	SetName(v string) *CreateBgpGroupRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateBgpGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateBgpGroupRequest
	GetOwnerId() *int64
	SetPeerAsn(v int64) *CreateBgpGroupRequest
	GetPeerAsn() *int64
	SetRegionId(v string) *CreateBgpGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateBgpGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateBgpGroupRequest
	GetResourceOwnerId() *int64
	SetRouteQuota(v int32) *CreateBgpGroupRequest
	GetRouteQuota() *int32
	SetRouterId(v string) *CreateBgpGroupRequest
	GetRouterId() *string
}

type CreateBgpGroupRequest struct {
	// The authentication key of the BGP group.
	//
	// example:
	//
	// !PWZ2****
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the BGP group.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// BGP
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IP version. Valid values:
	//
	// 	- **IPv4**: This is the default value.
	//
	// 	- **IPv6**: IPv6 is supported only if the VBR for which you want to create the BGP group has IPv6 enabled.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// Specifies whether to use a fake ASN. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// >  A router that runs BGP typically belongs to only one AS. If you need to replace an existing AS with a new AS and you cannot immediately modify BGP configurations, you can use fake ASNs to ensure service continuity.
	//
	// example:
	//
	// true
	IsFakeAsn *bool `json:"IsFakeAsn,omitempty" xml:"IsFakeAsn,omitempty"`
	// The custom ASN on the Alibaba Cloud side. Valid values:
	//
	// 	- **45104**
	//
	// 	- **64512~65534**
	//
	// 	- **4200000000~4294967294**
	//
	// >  **65025*	- is reserved by Alibaba Cloud. By default, Alibaba Cloud uses **45104*	- as **LocalAsn**. If you use custom **LocalAsn*	- in multi-line access scenarios, loops in BGP may occur.
	//
	// example:
	//
	// 45104
	LocalAsn *int64 `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// The name of the BGP group.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ASN of the gateway device in the data center.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1****
	PeerAsn *int64 `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	// The region ID of the VBR.
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
	// The ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp1ctxy813985gkuk****
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
}

func (s CreateBgpGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBgpGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateBgpGroupRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *CreateBgpGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateBgpGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateBgpGroupRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CreateBgpGroupRequest) GetIsFakeAsn() *bool {
	return s.IsFakeAsn
}

func (s *CreateBgpGroupRequest) GetLocalAsn() *int64 {
	return s.LocalAsn
}

func (s *CreateBgpGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateBgpGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateBgpGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateBgpGroupRequest) GetPeerAsn() *int64 {
	return s.PeerAsn
}

func (s *CreateBgpGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBgpGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateBgpGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateBgpGroupRequest) GetRouteQuota() *int32 {
	return s.RouteQuota
}

func (s *CreateBgpGroupRequest) GetRouterId() *string {
	return s.RouterId
}

func (s *CreateBgpGroupRequest) SetAuthKey(v string) *CreateBgpGroupRequest {
	s.AuthKey = &v
	return s
}

func (s *CreateBgpGroupRequest) SetClientToken(v string) *CreateBgpGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBgpGroupRequest) SetDescription(v string) *CreateBgpGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateBgpGroupRequest) SetIpVersion(v string) *CreateBgpGroupRequest {
	s.IpVersion = &v
	return s
}

func (s *CreateBgpGroupRequest) SetIsFakeAsn(v bool) *CreateBgpGroupRequest {
	s.IsFakeAsn = &v
	return s
}

func (s *CreateBgpGroupRequest) SetLocalAsn(v int64) *CreateBgpGroupRequest {
	s.LocalAsn = &v
	return s
}

func (s *CreateBgpGroupRequest) SetName(v string) *CreateBgpGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateBgpGroupRequest) SetOwnerAccount(v string) *CreateBgpGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateBgpGroupRequest) SetOwnerId(v int64) *CreateBgpGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateBgpGroupRequest) SetPeerAsn(v int64) *CreateBgpGroupRequest {
	s.PeerAsn = &v
	return s
}

func (s *CreateBgpGroupRequest) SetRegionId(v string) *CreateBgpGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBgpGroupRequest) SetResourceOwnerAccount(v string) *CreateBgpGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateBgpGroupRequest) SetResourceOwnerId(v int64) *CreateBgpGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateBgpGroupRequest) SetRouteQuota(v int32) *CreateBgpGroupRequest {
	s.RouteQuota = &v
	return s
}

func (s *CreateBgpGroupRequest) SetRouterId(v string) *CreateBgpGroupRequest {
	s.RouterId = &v
	return s
}

func (s *CreateBgpGroupRequest) Validate() error {
	return dara.Validate(s)
}
