// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterPeerAttachmentAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPublishRouteEnabled(v bool) *UpdateTransitRouterPeerAttachmentAttributeRequest
	GetAutoPublishRouteEnabled() *bool
	SetBandwidth(v int32) *UpdateTransitRouterPeerAttachmentAttributeRequest
	GetBandwidth() *int32
	SetBandwidthType(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest
	GetBandwidthType() *string
	SetCenBandwidthPackageId(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest
	GetCenBandwidthPackageId() *string
	SetClientToken(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest
	GetClientToken() *string
	SetDefaultLinkType(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest
	GetDefaultLinkType() *string
	SetDryRun(v bool) *UpdateTransitRouterPeerAttachmentAttributeRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateTransitRouterPeerAttachmentAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTransitRouterPeerAttachmentAttributeRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentDescription(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest
	GetTransitRouterAttachmentDescription() *string
	SetTransitRouterAttachmentId(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterAttachmentName(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest
	GetTransitRouterAttachmentName() *string
}

type UpdateTransitRouterPeerAttachmentAttributeRequest struct {
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
	// - If **BandwidthType*	- is set to **BandwidthPackage**, this parameter specifies the bandwidth that the inter-region connection can use.
	//
	// - If **BandwidthType*	- is set to **DataTransfer**, this parameter specifies the bandwidth limit of the inter-region connection.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The bandwidth allocation method. Valid values:
	//
	// - **BandwidthPackage**: allocates bandwidth from a bandwidth package.
	//
	// - **DataTransfer**: does not allocate bandwidth to the inter-region connection. Billing is based on the traffic volume.
	//
	// example:
	//
	// BandwidthPackage
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The ID of the bandwidth package to be associated with the inter-region connection.
	//
	// <props="china">If you do not specify a bandwidth package ID, the test bandwidth is used. The default test bandwidth is 1 Kbit/s and is intended only for testing (IPv4) network connectivity.
	//
	// >If **BandwidthType*	- is set to **DataTransfer**, you do not need to configure this parameter.
	//
	// example:
	//
	// cenbwp-3xrxupouolw5ou****
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The default link type.
	//
	// Valid values: Platinum and Gold. Default value: Gold.
	//
	// The value can be set to Platinum only when the bandwidth allocation method is pay-by-data-transfer.
	//
	// example:
	//
	// Gold
	DefaultLinkType *string `json:"DefaultLinkType,omitempty" xml:"DefaultLinkType,omitempty"`
	// Specifies whether to perform a dry run, including permission and instance status verification. Valid values:
	//
	// - **false*	- (default): sends a normal request and directly modifies the configuration of the inter-region connection after the request passes the check.
	//
	// - **true**: sends a check request. Only the check is performed and the configuration of the inter-region connection is not modified. The check items include whether required parameters are specified and the request format. If the check fails, the corresponding error is returned. If the check succeeds, the corresponding request ID is returned.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new description of the inter-region connection.
	//
	// The description can be empty or 1 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// testdesc
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The ID of the inter-region connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-ft94dcrbc3e5taun3x
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The new name of the inter-region connection.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// testname
	TransitRouterAttachmentName *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
}

func (s UpdateTransitRouterPeerAttachmentAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterPeerAttachmentAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) GetAutoPublishRouteEnabled() *bool {
	return s.AutoPublishRouteEnabled
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) GetCenBandwidthPackageId() *string {
	return s.CenBandwidthPackageId
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) GetDefaultLinkType() *string {
	return s.DefaultLinkType
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetAutoPublishRouteEnabled(v bool) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetBandwidth(v int32) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetBandwidthType(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.BandwidthType = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetCenBandwidthPackageId(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetClientToken(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetDefaultLinkType(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.DefaultLinkType = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetDryRun(v bool) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetOwnerAccount(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetOwnerId(v int64) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetResourceOwnerAccount(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetResourceOwnerId(v int64) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetTransitRouterAttachmentDescription(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetTransitRouterAttachmentId(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) SetTransitRouterAttachmentName(v string) *UpdateTransitRouterPeerAttachmentAttributeRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *UpdateTransitRouterPeerAttachmentAttributeRequest) Validate() error {
	return dara.Validate(s)
}
