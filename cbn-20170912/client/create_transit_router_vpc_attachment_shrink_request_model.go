// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterVpcAttachmentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPublishRouteEnabled(v bool) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetAutoPublishRouteEnabled() *bool
	SetCenId(v string) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetCenId() *string
	SetChargeType(v string) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateTransitRouterVpcAttachmentShrinkRequestTag) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetTag() []*CreateTransitRouterVpcAttachmentShrinkRequestTag
	SetTransitRouterAttachmentDescription(v string) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetTransitRouterAttachmentDescription() *string
	SetTransitRouterAttachmentName(v string) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetTransitRouterAttachmentName() *string
	SetTransitRouterId(v string) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetTransitRouterId() *string
	SetTransitRouterVPCAttachmentOptionsShrink(v string) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetTransitRouterVPCAttachmentOptionsShrink() *string
	SetVpcId(v string) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetVpcId() *string
	SetVpcOwnerId(v int64) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetVpcOwnerId() *int64
	SetZoneMappings(v []*CreateTransitRouterVpcAttachmentShrinkRequestZoneMappings) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetZoneMappings() []*CreateTransitRouterVpcAttachmentShrinkRequestZoneMappings
}

type CreateTransitRouterVpcAttachmentShrinkRequest struct {
	// Specifies whether to enable the Enterprise Edition transit router to automatically advertise routes to VPCs. Valid values:
	//
	// 	- **false:*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// true
	AutoPublishRouteEnabled *bool `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The billing method. The default value is **POSTPAY**, which specifies the pay-as-you-go billing method.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **false*	- (default): performs a dry run and sends the request.
	//
	// 	- **true**: performs a dry run. The system checks the required parameters and request syntax. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the VPC is deployed.
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
	Tag []*CreateTransitRouterVpcAttachmentShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The description of the VPC connection.
	//
	// The description must be 1 to 256 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// testname
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The name of the VPC connection.
	//
	// The name must be 1 to 128 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// testname
	TransitRouterAttachmentName *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	// The ID of the Enterprise Edition transit router.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// Feature configurations of the VPC connection.
	TransitRouterVPCAttachmentOptionsShrink *string `json:"TransitRouterVPCAttachmentOptions,omitempty" xml:"TransitRouterVPCAttachmentOptions,omitempty"`
	// The VPC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1kbjcre9vtsebo1****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the Alibaba Cloud account to which the VPC belongs. The default value is the ID of the current Alibaba Cloud account.
	//
	// > If the network instance and CEN instance belong to different Alibaba Cloud accounts, this parameter is required.
	//
	// example:
	//
	// 1250123456123456
	VpcOwnerId *int64 `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
	// A zone that supports Enterprise Edition transit routers.
	//
	// You can specify at most 10 zones.
	//
	// This parameter is required.
	ZoneMappings []*CreateTransitRouterVpcAttachmentShrinkRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s CreateTransitRouterVpcAttachmentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVpcAttachmentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetAutoPublishRouteEnabled() *bool {
	return s.AutoPublishRouteEnabled
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetTag() []*CreateTransitRouterVpcAttachmentShrinkRequestTag {
	return s.Tag
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetTransitRouterVPCAttachmentOptionsShrink() *string {
	return s.TransitRouterVPCAttachmentOptionsShrink
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetVpcOwnerId() *int64 {
	return s.VpcOwnerId
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetZoneMappings() []*CreateTransitRouterVpcAttachmentShrinkRequestZoneMappings {
	return s.ZoneMappings
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetAutoPublishRouteEnabled(v bool) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetCenId(v string) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.CenId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetChargeType(v string) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetClientToken(v string) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetDryRun(v bool) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetOwnerAccount(v string) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetOwnerId(v int64) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetRegionId(v string) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetResourceOwnerId(v int64) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetTag(v []*CreateTransitRouterVpcAttachmentShrinkRequestTag) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetTransitRouterAttachmentDescription(v string) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetTransitRouterAttachmentName(v string) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetTransitRouterId(v string) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetTransitRouterVPCAttachmentOptionsShrink(v string) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.TransitRouterVPCAttachmentOptionsShrink = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetVpcId(v string) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetVpcOwnerId(v int64) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.VpcOwnerId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetZoneMappings(v []*CreateTransitRouterVpcAttachmentShrinkRequestZoneMappings) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.ZoneMappings = v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTransitRouterVpcAttachmentShrinkRequestTag struct {
	// The tag key.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys.
	//
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be 0 to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// Each tag key must have a unique tag value. You can specify at most 20 tag values in each call.
	//
	// example:
	//
	// tagtest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTransitRouterVpcAttachmentShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVpcAttachmentShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequestTag) SetKey(v string) *CreateTransitRouterVpcAttachmentShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequestTag) SetValue(v string) *CreateTransitRouterVpcAttachmentShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateTransitRouterVpcAttachmentShrinkRequestZoneMappings struct {
	// A vSwitch that is deployed in the zone that supports Enterprise Edition transit routers.
	//
	// You can specify vSwitches for at most 10 zones in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1a214sbus8z3b54****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone that supports Enterprise Edition transit routers.
	//
	// You can call the [DescribeZones](https://help.aliyun.com/document_detail/36064.html) operation to query the most recent zone list.
	//
	// You can specify at most 10 zones in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateTransitRouterVpcAttachmentShrinkRequestZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVpcAttachmentShrinkRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequestZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequestZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequestZoneMappings) SetVSwitchId(v string) *CreateTransitRouterVpcAttachmentShrinkRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequestZoneMappings) SetZoneId(v string) *CreateTransitRouterVpcAttachmentShrinkRequestZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentShrinkRequestZoneMappings) Validate() error {
	return dara.Validate(s)
}
