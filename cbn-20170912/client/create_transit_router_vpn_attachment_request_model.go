// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterVpnAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPublishRouteEnabled(v bool) *CreateTransitRouterVpnAttachmentRequest
	GetAutoPublishRouteEnabled() *bool
	SetCenId(v string) *CreateTransitRouterVpnAttachmentRequest
	GetCenId() *string
	SetChargeType(v string) *CreateTransitRouterVpnAttachmentRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateTransitRouterVpnAttachmentRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTransitRouterVpnAttachmentRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *CreateTransitRouterVpnAttachmentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTransitRouterVpnAttachmentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateTransitRouterVpnAttachmentRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateTransitRouterVpnAttachmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTransitRouterVpnAttachmentRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateTransitRouterVpnAttachmentRequestTag) *CreateTransitRouterVpnAttachmentRequest
	GetTag() []*CreateTransitRouterVpnAttachmentRequestTag
	SetTransitRouterAttachmentDescription(v string) *CreateTransitRouterVpnAttachmentRequest
	GetTransitRouterAttachmentDescription() *string
	SetTransitRouterAttachmentName(v string) *CreateTransitRouterVpnAttachmentRequest
	GetTransitRouterAttachmentName() *string
	SetTransitRouterId(v string) *CreateTransitRouterVpnAttachmentRequest
	GetTransitRouterId() *string
	SetVpnId(v string) *CreateTransitRouterVpnAttachmentRequest
	GetVpnId() *string
	SetVpnOwnerId(v int64) *CreateTransitRouterVpnAttachmentRequest
	GetVpnOwnerId() *int64
	SetZone(v []*CreateTransitRouterVpnAttachmentRequestZone) *CreateTransitRouterVpnAttachmentRequest
	GetZone() []*CreateTransitRouterVpnAttachmentRequestZone
}

type CreateTransitRouterVpnAttachmentRequest struct {
	// Specifies whether to allow the transit router instance to automatically publish route entries to the IPsec connection. Valid values:
	//
	// - **true*	- (default): allowed.
	//
	// - **false**: not allowed.
	//
	// example:
	//
	// true
	AutoPublishRouteEnabled *bool `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	// The Cloud Enterprise Network (CEN) instance ID.
	//
	// example:
	//
	// cen-rsgxs8ng2awen2****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The billing method.
	//
	// Set the value to **POSTPAY*	- (default), which specifies the pay-as-you-go billable method based on usage.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, the VPN connection is created.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the transit router instance.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag information list.
	//
	// You can specify up to 20 tags at a time.
	Tag []*CreateTransitRouterVpnAttachmentRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The description of the VPN connection.
	//
	// The description can be empty or 1 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The name of the VPN connection.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	TransitRouterAttachmentName *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	// The transit router instance ID.
	//
	// example:
	//
	// tr-p0wm740vjnbaprv0m****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The ID of the IPsec connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-p0w042cqwvlhl4zyw****
	VpnId *string `json:"VpnId,omitempty" xml:"VpnId,omitempty"`
	// The Alibaba Cloud account ID of the Alibaba Cloud account to which the IPsec connection belongs.
	//
	// - If you do not specify this parameter, the Alibaba Cloud account ID of the current logon account is used by default.
	//
	// - This parameter is required if you want to connect to an IPsec connection that belongs to a different account.
	//
	// example:
	//
	// 1210123456123456
	VpnOwnerId *int64 `json:"VpnOwnerId,omitempty" xml:"VpnOwnerId,omitempty"`
	// The zone ID in the current region.
	//
	// The system creates resources in the zone that you specify.
	//
	// > If the bindeded IPsec connection uses the dual-tunnel mode, leave this parameter empty.
	Zone []*CreateTransitRouterVpnAttachmentRequestZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s CreateTransitRouterVpnAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVpnAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetAutoPublishRouteEnabled() *bool {
	return s.AutoPublishRouteEnabled
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetTag() []*CreateTransitRouterVpnAttachmentRequestTag {
	return s.Tag
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetVpnId() *string {
	return s.VpnId
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetVpnOwnerId() *int64 {
	return s.VpnOwnerId
}

func (s *CreateTransitRouterVpnAttachmentRequest) GetZone() []*CreateTransitRouterVpnAttachmentRequestZone {
	return s.Zone
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetAutoPublishRouteEnabled(v bool) *CreateTransitRouterVpnAttachmentRequest {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetCenId(v string) *CreateTransitRouterVpnAttachmentRequest {
	s.CenId = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetChargeType(v string) *CreateTransitRouterVpnAttachmentRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetClientToken(v string) *CreateTransitRouterVpnAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetDryRun(v bool) *CreateTransitRouterVpnAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetOwnerAccount(v string) *CreateTransitRouterVpnAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetOwnerId(v int64) *CreateTransitRouterVpnAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetRegionId(v string) *CreateTransitRouterVpnAttachmentRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterVpnAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetResourceOwnerId(v int64) *CreateTransitRouterVpnAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetTag(v []*CreateTransitRouterVpnAttachmentRequestTag) *CreateTransitRouterVpnAttachmentRequest {
	s.Tag = v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetTransitRouterAttachmentDescription(v string) *CreateTransitRouterVpnAttachmentRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetTransitRouterAttachmentName(v string) *CreateTransitRouterVpnAttachmentRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetTransitRouterId(v string) *CreateTransitRouterVpnAttachmentRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetVpnId(v string) *CreateTransitRouterVpnAttachmentRequest {
	s.VpnId = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetVpnOwnerId(v int64) *CreateTransitRouterVpnAttachmentRequest {
	s.VpnOwnerId = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) SetZone(v []*CreateTransitRouterVpnAttachmentRequestZone) *CreateTransitRouterVpnAttachmentRequest {
	s.Zone = v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Zone != nil {
		for _, item := range s.Zone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTransitRouterVpnAttachmentRequestTag struct {
	// The tag key of the resource.
	//
	// Once specified, the tag key cannot be an empty string. The tag key can be up to 64 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// Once specified, the tag value cannot be empty. The tag value can be up to 128 characters in length, and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// Each tag key corresponds to one tag value. You can specify up to 20 tag values at a time.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTransitRouterVpnAttachmentRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVpnAttachmentRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpnAttachmentRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTransitRouterVpnAttachmentRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTransitRouterVpnAttachmentRequestTag) SetKey(v string) *CreateTransitRouterVpnAttachmentRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequestTag) SetValue(v string) *CreateTransitRouterVpnAttachmentRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateTransitRouterVpnAttachmentRequestZone struct {
	// The zone ID.
	//
	// You can call the [ListTransitRouterAvailableResource](https://help.aliyun.com/document_detail/261356.html) operation to query zone IDs.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateTransitRouterVpnAttachmentRequestZone) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVpnAttachmentRequestZone) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpnAttachmentRequestZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateTransitRouterVpnAttachmentRequestZone) SetZoneId(v string) *CreateTransitRouterVpnAttachmentRequestZone {
	s.ZoneId = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentRequestZone) Validate() error {
	return dara.Validate(s)
}
