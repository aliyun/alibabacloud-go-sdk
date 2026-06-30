// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterVpcAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPublishRouteEnabled(v bool) *CreateTransitRouterVpcAttachmentRequest
	GetAutoPublishRouteEnabled() *bool
	SetCenId(v string) *CreateTransitRouterVpcAttachmentRequest
	GetCenId() *string
	SetChargeType(v string) *CreateTransitRouterVpcAttachmentRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateTransitRouterVpcAttachmentRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTransitRouterVpcAttachmentRequest
	GetDryRun() *bool
	SetOptions(v *CreateTransitRouterVpcAttachmentRequestOptions) *CreateTransitRouterVpcAttachmentRequest
	GetOptions() *CreateTransitRouterVpcAttachmentRequestOptions
	SetOwnerAccount(v string) *CreateTransitRouterVpcAttachmentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTransitRouterVpcAttachmentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateTransitRouterVpcAttachmentRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateTransitRouterVpcAttachmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTransitRouterVpcAttachmentRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateTransitRouterVpcAttachmentRequestTag) *CreateTransitRouterVpcAttachmentRequest
	GetTag() []*CreateTransitRouterVpcAttachmentRequestTag
	SetTransitRouterAttachmentDescription(v string) *CreateTransitRouterVpcAttachmentRequest
	GetTransitRouterAttachmentDescription() *string
	SetTransitRouterAttachmentName(v string) *CreateTransitRouterVpcAttachmentRequest
	GetTransitRouterAttachmentName() *string
	SetTransitRouterId(v string) *CreateTransitRouterVpcAttachmentRequest
	GetTransitRouterId() *string
	SetTransitRouterVPCAttachmentOptions(v map[string]*string) *CreateTransitRouterVpcAttachmentRequest
	GetTransitRouterVPCAttachmentOptions() map[string]*string
	SetVpcId(v string) *CreateTransitRouterVpcAttachmentRequest
	GetVpcId() *string
	SetVpcOwnerId(v int64) *CreateTransitRouterVpcAttachmentRequest
	GetVpcOwnerId() *int64
	SetZoneMappings(v []*CreateTransitRouterVpcAttachmentRequestZoneMappings) *CreateTransitRouterVpcAttachmentRequest
	GetZoneMappings() []*CreateTransitRouterVpcAttachmentRequestZoneMappings
}

type CreateTransitRouterVpcAttachmentRequest struct {
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
	DryRun       *bool                                           `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Options      *CreateTransitRouterVpcAttachmentRequestOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	OwnerAccount *string                                         `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	Tag []*CreateTransitRouterVpcAttachmentRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	TransitRouterVPCAttachmentOptions map[string]*string `json:"TransitRouterVPCAttachmentOptions,omitempty" xml:"TransitRouterVPCAttachmentOptions,omitempty"`
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
	ZoneMappings []*CreateTransitRouterVpcAttachmentRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s CreateTransitRouterVpcAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVpcAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetAutoPublishRouteEnabled() *bool {
	return s.AutoPublishRouteEnabled
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetOptions() *CreateTransitRouterVpcAttachmentRequestOptions {
	return s.Options
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetTag() []*CreateTransitRouterVpcAttachmentRequestTag {
	return s.Tag
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetTransitRouterVPCAttachmentOptions() map[string]*string {
	return s.TransitRouterVPCAttachmentOptions
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetVpcOwnerId() *int64 {
	return s.VpcOwnerId
}

func (s *CreateTransitRouterVpcAttachmentRequest) GetZoneMappings() []*CreateTransitRouterVpcAttachmentRequestZoneMappings {
	return s.ZoneMappings
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetAutoPublishRouteEnabled(v bool) *CreateTransitRouterVpcAttachmentRequest {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetCenId(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.CenId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetChargeType(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetClientToken(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetDryRun(v bool) *CreateTransitRouterVpcAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetOptions(v *CreateTransitRouterVpcAttachmentRequestOptions) *CreateTransitRouterVpcAttachmentRequest {
	s.Options = v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetOwnerAccount(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetOwnerId(v int64) *CreateTransitRouterVpcAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetRegionId(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetResourceOwnerId(v int64) *CreateTransitRouterVpcAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetTag(v []*CreateTransitRouterVpcAttachmentRequestTag) *CreateTransitRouterVpcAttachmentRequest {
	s.Tag = v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetTransitRouterAttachmentDescription(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetTransitRouterAttachmentName(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetTransitRouterId(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetTransitRouterVPCAttachmentOptions(v map[string]*string) *CreateTransitRouterVpcAttachmentRequest {
	s.TransitRouterVPCAttachmentOptions = v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetVpcId(v string) *CreateTransitRouterVpcAttachmentRequest {
	s.VpcId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetVpcOwnerId(v int64) *CreateTransitRouterVpcAttachmentRequest {
	s.VpcOwnerId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) SetZoneMappings(v []*CreateTransitRouterVpcAttachmentRequestZoneMappings) *CreateTransitRouterVpcAttachmentRequest {
	s.ZoneMappings = v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequest) Validate() error {
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			return err
		}
	}
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

type CreateTransitRouterVpcAttachmentRequestOptions struct {
	// example:
	//
	// enable
	ApplianceModeSupport *string `json:"ApplianceModeSupport,omitempty" xml:"ApplianceModeSupport,omitempty"`
	// example:
	//
	// enable
	Ipv6Support *string `json:"Ipv6Support,omitempty" xml:"Ipv6Support,omitempty"`
}

func (s CreateTransitRouterVpcAttachmentRequestOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVpcAttachmentRequestOptions) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpcAttachmentRequestOptions) GetApplianceModeSupport() *string {
	return s.ApplianceModeSupport
}

func (s *CreateTransitRouterVpcAttachmentRequestOptions) GetIpv6Support() *string {
	return s.Ipv6Support
}

func (s *CreateTransitRouterVpcAttachmentRequestOptions) SetApplianceModeSupport(v string) *CreateTransitRouterVpcAttachmentRequestOptions {
	s.ApplianceModeSupport = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequestOptions) SetIpv6Support(v string) *CreateTransitRouterVpcAttachmentRequestOptions {
	s.Ipv6Support = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequestOptions) Validate() error {
	return dara.Validate(s)
}

type CreateTransitRouterVpcAttachmentRequestTag struct {
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

func (s CreateTransitRouterVpcAttachmentRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVpcAttachmentRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpcAttachmentRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTransitRouterVpcAttachmentRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTransitRouterVpcAttachmentRequestTag) SetKey(v string) *CreateTransitRouterVpcAttachmentRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequestTag) SetValue(v string) *CreateTransitRouterVpcAttachmentRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateTransitRouterVpcAttachmentRequestZoneMappings struct {
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

func (s CreateTransitRouterVpcAttachmentRequestZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVpcAttachmentRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpcAttachmentRequestZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateTransitRouterVpcAttachmentRequestZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateTransitRouterVpcAttachmentRequestZoneMappings) SetVSwitchId(v string) *CreateTransitRouterVpcAttachmentRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequestZoneMappings) SetZoneId(v string) *CreateTransitRouterVpcAttachmentRequestZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentRequestZoneMappings) Validate() error {
	return dara.Validate(s)
}
