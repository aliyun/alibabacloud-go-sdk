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
	SetOptionsShrink(v string) *CreateTransitRouterVpcAttachmentShrinkRequest
	GetOptionsShrink() *string
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
	// Specifies whether to enable the Enterprise Edition transit router to automatically advertise routes to the VPC.
	//
	// - **false*	- (default): Do not automatically advertise routes.
	//
	// - **true**: Automatically advertise routes.
	//
	// example:
	//
	// true
	AutoPublishRouteEnabled *bool `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The billing method. The default value is **POSTPAY*	- (pay-as-you-go).
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token used to ensure request idempotency.
	//
	// You must generate a value on your client that is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token, which is unique for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run to check the validity of the request without creating the resource. The check includes permissions and instance status. Valid values:
	//
	// - **false*	- (default): Sends a normal request. The system creates the VPC connection if the request is valid.
	//
	// - **true**: Sends only a check request. The system checks required parameters, request format, and permissions. The VPC connection is not created. If the check fails, an error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun        *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OptionsShrink *string `json:"Options,omitempty" xml:"Options,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the VPC is located.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags to add to the VPC connection.
	//
	// You can add up to 20 tags.
	Tag []*CreateTransitRouterVpcAttachmentShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The description of the VPC connection.
	//
	// The description can be empty or 1 to 256 characters long, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testname
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The name of the VPC connection.
	//
	// The name can be empty or 1 to 128 characters long, and cannot start with `http://` or `https://`.
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
	// The properties of the VPC connection. This parameter is deprecated. We recommend that you use the `Options` parameter instead.
	TransitRouterVPCAttachmentOptionsShrink *string `json:"TransitRouterVPCAttachmentOptions,omitempty" xml:"TransitRouterVPCAttachmentOptions,omitempty"`
	// The ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1kbjcre9vtsebo1****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the VPC. By default, this is the ID of the current Alibaba Cloud account.
	//
	// > This parameter is required if you want to attach a cross-account network instance.
	//
	// example:
	//
	// 1250123456123456
	VpcOwnerId *int64 `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
	// The zone mappings for the VPC connection. For each mapping, you must specify a vSwitch in a zone that is supported by the Enterprise Edition transit router.
	//
	// You can specify up to 10 zone mappings.
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

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) GetOptionsShrink() *string {
	return s.OptionsShrink
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

func (s *CreateTransitRouterVpcAttachmentShrinkRequest) SetOptionsShrink(v string) *CreateTransitRouterVpcAttachmentShrinkRequest {
	s.OptionsShrink = &v
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ZoneMappings != nil {
		for _, item := range s.ZoneMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTransitRouterVpcAttachmentShrinkRequestTag struct {
	// The key of the tag.
	//
	// The tag key cannot be an empty string. The key can be up to 64 characters long and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// The tag value can be an empty string or a string up to 128 characters long. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
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
	// The ID of the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1a214sbus8z3b54****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone. The zone must be supported by the Enterprise Edition transit router.
	//
	// You can call the [DescribeZones](https://help.aliyun.com/document_detail/36064.html) operation to query available zones.
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
