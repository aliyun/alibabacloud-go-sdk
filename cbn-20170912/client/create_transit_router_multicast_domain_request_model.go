// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterMulticastDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *CreateTransitRouterMulticastDomainRequest
	GetCenId() *string
	SetClientToken(v string) *CreateTransitRouterMulticastDomainRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTransitRouterMulticastDomainRequest
	GetDryRun() *bool
	SetOptions(v *CreateTransitRouterMulticastDomainRequestOptions) *CreateTransitRouterMulticastDomainRequest
	GetOptions() *CreateTransitRouterMulticastDomainRequestOptions
	SetOwnerAccount(v string) *CreateTransitRouterMulticastDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTransitRouterMulticastDomainRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateTransitRouterMulticastDomainRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateTransitRouterMulticastDomainRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTransitRouterMulticastDomainRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateTransitRouterMulticastDomainRequestTag) *CreateTransitRouterMulticastDomainRequest
	GetTag() []*CreateTransitRouterMulticastDomainRequestTag
	SetTransitRouterId(v string) *CreateTransitRouterMulticastDomainRequest
	GetTransitRouterId() *string
	SetTransitRouterMulticastDomainDescription(v string) *CreateTransitRouterMulticastDomainRequest
	GetTransitRouterMulticastDomainDescription() *string
	SetTransitRouterMulticastDomainName(v string) *CreateTransitRouterMulticastDomainRequest
	GetTransitRouterMulticastDomainName() *string
}

type CreateTransitRouterMulticastDomainRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-a7syd349kne38g****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request format, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Multicast domain feature.
	Options      *CreateTransitRouterMulticastDomainRequestOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	OwnerAccount *string                                           `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the transit router.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The information about the tags.
	//
	// You can specify at most 20 tags in each call.
	Tag []*CreateTransitRouterMulticastDomainRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-p0wr9p28r92d598y6****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The description of the multicast domain.
	//
	// The description must be 1 to 256 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// desctest
	TransitRouterMulticastDomainDescription *string `json:"TransitRouterMulticastDomainDescription,omitempty" xml:"TransitRouterMulticastDomainDescription,omitempty"`
	// The name of the multicast domain.
	//
	// The name must be 1 to 128 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// nametest
	TransitRouterMulticastDomainName *string `json:"TransitRouterMulticastDomainName,omitempty" xml:"TransitRouterMulticastDomainName,omitempty"`
}

func (s CreateTransitRouterMulticastDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterMulticastDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterMulticastDomainRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateTransitRouterMulticastDomainRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouterMulticastDomainRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouterMulticastDomainRequest) GetOptions() *CreateTransitRouterMulticastDomainRequestOptions {
	return s.Options
}

func (s *CreateTransitRouterMulticastDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTransitRouterMulticastDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTransitRouterMulticastDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTransitRouterMulticastDomainRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTransitRouterMulticastDomainRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTransitRouterMulticastDomainRequest) GetTag() []*CreateTransitRouterMulticastDomainRequestTag {
	return s.Tag
}

func (s *CreateTransitRouterMulticastDomainRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateTransitRouterMulticastDomainRequest) GetTransitRouterMulticastDomainDescription() *string {
	return s.TransitRouterMulticastDomainDescription
}

func (s *CreateTransitRouterMulticastDomainRequest) GetTransitRouterMulticastDomainName() *string {
	return s.TransitRouterMulticastDomainName
}

func (s *CreateTransitRouterMulticastDomainRequest) SetCenId(v string) *CreateTransitRouterMulticastDomainRequest {
	s.CenId = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequest) SetClientToken(v string) *CreateTransitRouterMulticastDomainRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequest) SetDryRun(v bool) *CreateTransitRouterMulticastDomainRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequest) SetOptions(v *CreateTransitRouterMulticastDomainRequestOptions) *CreateTransitRouterMulticastDomainRequest {
	s.Options = v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequest) SetOwnerAccount(v string) *CreateTransitRouterMulticastDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequest) SetOwnerId(v int64) *CreateTransitRouterMulticastDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequest) SetRegionId(v string) *CreateTransitRouterMulticastDomainRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterMulticastDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequest) SetResourceOwnerId(v int64) *CreateTransitRouterMulticastDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequest) SetTag(v []*CreateTransitRouterMulticastDomainRequestTag) *CreateTransitRouterMulticastDomainRequest {
	s.Tag = v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequest) SetTransitRouterId(v string) *CreateTransitRouterMulticastDomainRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequest) SetTransitRouterMulticastDomainDescription(v string) *CreateTransitRouterMulticastDomainRequest {
	s.TransitRouterMulticastDomainDescription = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequest) SetTransitRouterMulticastDomainName(v string) *CreateTransitRouterMulticastDomainRequest {
	s.TransitRouterMulticastDomainName = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTransitRouterMulticastDomainRequestOptions struct {
	// Indicates whether the IGMP feature is enabled for the multicast domain. Once enabled, hosts can dynamically join or leave multicast groups by using IGMP protocol. Valid values:
	//
	// 	- **enable**: enables IGMP.
	//
	// 	- **disable**(default): disables IGMP.
	//
	// > 	- The IGMP feature is in beta testing. To use it, contact your account manager.
	//
	// > 	- If you select this option, you cannot disable IPv6 after the VBR is created.
	//
	// example:
	//
	// enable
	Igmpv2Support *string `json:"Igmpv2Support,omitempty" xml:"Igmpv2Support,omitempty"`
}

func (s CreateTransitRouterMulticastDomainRequestOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterMulticastDomainRequestOptions) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterMulticastDomainRequestOptions) GetIgmpv2Support() *string {
	return s.Igmpv2Support
}

func (s *CreateTransitRouterMulticastDomainRequestOptions) SetIgmpv2Support(v string) *CreateTransitRouterMulticastDomainRequestOptions {
	s.Igmpv2Support = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequestOptions) Validate() error {
	return dara.Validate(s)
}

type CreateTransitRouterMulticastDomainRequestTag struct {
	// The tag key.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be 0 to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// Each tag key must have a unique tag value. You can specify at most 20 tag values in each call.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTransitRouterMulticastDomainRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterMulticastDomainRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterMulticastDomainRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTransitRouterMulticastDomainRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTransitRouterMulticastDomainRequestTag) SetKey(v string) *CreateTransitRouterMulticastDomainRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequestTag) SetValue(v string) *CreateTransitRouterMulticastDomainRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainRequestTag) Validate() error {
	return dara.Validate(s)
}
