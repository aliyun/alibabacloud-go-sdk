// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterPeerAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPublishRouteEnabled(v bool) *CreateTransitRouterPeerAttachmentRequest
	GetAutoPublishRouteEnabled() *bool
	SetBandwidth(v int32) *CreateTransitRouterPeerAttachmentRequest
	GetBandwidth() *int32
	SetBandwidthType(v string) *CreateTransitRouterPeerAttachmentRequest
	GetBandwidthType() *string
	SetCenBandwidthPackageId(v string) *CreateTransitRouterPeerAttachmentRequest
	GetCenBandwidthPackageId() *string
	SetCenId(v string) *CreateTransitRouterPeerAttachmentRequest
	GetCenId() *string
	SetClientToken(v string) *CreateTransitRouterPeerAttachmentRequest
	GetClientToken() *string
	SetDefaultLinkType(v string) *CreateTransitRouterPeerAttachmentRequest
	GetDefaultLinkType() *string
	SetDryRun(v bool) *CreateTransitRouterPeerAttachmentRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *CreateTransitRouterPeerAttachmentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTransitRouterPeerAttachmentRequest
	GetOwnerId() *int64
	SetPeerTransitRouterId(v string) *CreateTransitRouterPeerAttachmentRequest
	GetPeerTransitRouterId() *string
	SetPeerTransitRouterRegionId(v string) *CreateTransitRouterPeerAttachmentRequest
	GetPeerTransitRouterRegionId() *string
	SetRegionId(v string) *CreateTransitRouterPeerAttachmentRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateTransitRouterPeerAttachmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTransitRouterPeerAttachmentRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateTransitRouterPeerAttachmentRequestTag) *CreateTransitRouterPeerAttachmentRequest
	GetTag() []*CreateTransitRouterPeerAttachmentRequestTag
	SetTransitRouterAttachmentDescription(v string) *CreateTransitRouterPeerAttachmentRequest
	GetTransitRouterAttachmentDescription() *string
	SetTransitRouterAttachmentName(v string) *CreateTransitRouterPeerAttachmentRequest
	GetTransitRouterAttachmentName() *string
	SetTransitRouterId(v string) *CreateTransitRouterPeerAttachmentRequest
	GetTransitRouterId() *string
}

type CreateTransitRouterPeerAttachmentRequest struct {
	// Specifies whether to allow the Enterprise Edition transit router to automatically advertise routes of the inter-region connection to the peer region.
	//
	// - **false*	- (default): no.
	//
	// - **true**: yes.
	//
	// example:
	//
	// false
	AutoPublishRouteEnabled *bool `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	// The bandwidth value of the inter-region connection. Unit: Mbit/s.
	//
	// - If **BandwidthType*	- is set to **BandwidthPackage**, this parameter specifies the bandwidth that can be used by the inter-region connection.
	//
	// - If **BandwidthType*	- is set to **DataTransfer**, this parameter specifies the bandwidth limit of the inter-region connection.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The bandwidth allocation method of the inter-region connection. Valid values:
	//
	// - **BandwidthPackage**: allocates bandwidth from a bandwidth package.
	//
	// - **DataTransfer**: does not allocate bandwidth to the inter-region connection. The system charges you based on the actual traffic.
	//
	// example:
	//
	// BandwidthPackage
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The ID of the bandwidth package to be associated with the inter-region connection.
	//
	// > If **BandwidthType*	- is set to **DataTransfer**, you do not need to configure this parameter.
	//
	// example:
	//
	// cenbwp-3xrxupouolw5ou****
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	// The Cloud Enterprise Network (CEN) instance ID.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The default link type.
	//
	// Valid values: Platinum and Gold. Default value: Gold.
	//
	// The link type can be set to Platinum only when the bandwidth allocation method is pay-by-data-transfer.
	//
	// example:
	//
	// Gold
	DefaultLinkType *string `json:"DefaultLinkType,omitempty" xml:"DefaultLinkType,omitempty"`
	// Specifies whether to perform a dry run, including permission and instance status verification. Valid values:
	//
	// - **false*	- (default): sends a normal request. If the request passes the verification, the inter-region connection is created.
	//
	// - **true**: sends a check request. Only the verification is performed. No inter-region connection is created. The system checks whether the required parameters are specified, and validates the request format. If the check fails, the corresponding error is returned. If the check succeeds, the corresponding request ID is returned.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the peer transit router instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-m5eq27g6bndum7e88****
	PeerTransitRouterId *string `json:"PeerTransitRouterId,omitempty" xml:"PeerTransitRouterId,omitempty"`
	// The region ID of the peer transit router instance.
	//
	// example:
	//
	// cn-qingdao
	PeerTransitRouterRegionId *string `json:"PeerTransitRouterRegionId,omitempty" xml:"PeerTransitRouterRegionId,omitempty"`
	// The region ID of the local Enterprise Edition transit router instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags at a time.
	Tag []*CreateTransitRouterPeerAttachmentRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The description of the inter-region connection.
	//
	// The description can be empty or 1 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// testdesc
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The name of the inter-region connection.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// testname
	TransitRouterAttachmentName *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	// The ID of the local Enterprise Edition transit router instance.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s CreateTransitRouterPeerAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterPeerAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetAutoPublishRouteEnabled() *bool {
	return s.AutoPublishRouteEnabled
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetCenBandwidthPackageId() *string {
	return s.CenBandwidthPackageId
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetDefaultLinkType() *string {
	return s.DefaultLinkType
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetPeerTransitRouterId() *string {
	return s.PeerTransitRouterId
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetPeerTransitRouterRegionId() *string {
	return s.PeerTransitRouterRegionId
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetTag() []*CreateTransitRouterPeerAttachmentRequestTag {
	return s.Tag
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *CreateTransitRouterPeerAttachmentRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetAutoPublishRouteEnabled(v bool) *CreateTransitRouterPeerAttachmentRequest {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetBandwidth(v int32) *CreateTransitRouterPeerAttachmentRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetBandwidthType(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.BandwidthType = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetCenBandwidthPackageId(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetCenId(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.CenId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetClientToken(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetDefaultLinkType(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.DefaultLinkType = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetDryRun(v bool) *CreateTransitRouterPeerAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetOwnerAccount(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetOwnerId(v int64) *CreateTransitRouterPeerAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetPeerTransitRouterId(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.PeerTransitRouterId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetPeerTransitRouterRegionId(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.PeerTransitRouterRegionId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetRegionId(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetResourceOwnerId(v int64) *CreateTransitRouterPeerAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetTag(v []*CreateTransitRouterPeerAttachmentRequestTag) *CreateTransitRouterPeerAttachmentRequest {
	s.Tag = v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetTransitRouterAttachmentDescription(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetTransitRouterAttachmentName(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) SetTransitRouterId(v string) *CreateTransitRouterPeerAttachmentRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequest) Validate() error {
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

type CreateTransitRouterPeerAttachmentRequestTag struct {
	// The tag key of the resource.
	//
	// Once specified, the tag key cannot be an empty string. The tag key can be up to 64 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// tag_A1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// Once specified, the tag value cannot be empty. The tag value can be up to 128 characters in length, and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// Each tag key corresponds to one tag value. You can specify up to 20 tag values at a time.
	//
	// example:
	//
	// value_A1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTransitRouterPeerAttachmentRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterPeerAttachmentRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPeerAttachmentRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTransitRouterPeerAttachmentRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTransitRouterPeerAttachmentRequestTag) SetKey(v string) *CreateTransitRouterPeerAttachmentRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequestTag) SetValue(v string) *CreateTransitRouterPeerAttachmentRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentRequestTag) Validate() error {
	return dara.Validate(s)
}
