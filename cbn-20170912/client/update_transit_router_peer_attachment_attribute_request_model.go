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
	// Specifies whether to enable the local Enterprise Edition transit router to automatically advertise the routes of the inter-region connection to the peer transit router. Valid values:
	//
	// 	- **false*	- (default): no
	//
	// 	- **true**: yes
	//
	// example:
	//
	// false
	AutoPublishRouteEnabled *bool `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	// The bandwidth value of the inter-region connection. Unit: Mbit/s.
	//
	// 	- This parameter specifies the maximum bandwidth value for the inter-region connection if you set **BandwidthType*	- to **BandwidthPackage**.
	//
	// 	- This parameter specifies the bandwidth throttling threshold for the inter-region connection if you set **BandwidthType*	- to **DataTransfer**.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The bandwidth allocation method. Valid values:
	//
	// 	- **BandwidthPackage**: allocates bandwidth from a bandwidth plan.
	//
	// 	- **DataTransfer**: bandwidth is billed based on the pay-by-data-transfer metering method.
	//
	// example:
	//
	// BandwidthPackage
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The ID of the bandwidth plan that is used to allocate bandwidth to the inter-region connection.
	//
	// >  If you set **BandwidthType*	- to **DataTransfer**, you do not need to set this parameter.
	//
	// example:
	//
	// cenbwp-3xrxupouolw5ou****
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The default line type.
	//
	// Valid values: Platinum and Gold.
	//
	// Platinum is supported only when BandwidthType is set to DataTransfer.
	//
	// example:
	//
	// Gold
	DefaultLinkType *string `json:"DefaultLinkType,omitempty" xml:"DefaultLinkType,omitempty"`
	// Specifies whether to perform a dry run to check information such as the permissions and the instance status. Default values:
	//
	// 	- **false*	- (default): performs a dry run and sends the request.
	//
	// 	- **true**: performs a dry run. The system checks the required parameters and request syntax. If the request fails the dry run, an error message is returned. If the request passes the dry run, the system returns the ID of the request.
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
	// This parameter is optional. If you enter a description, it must be 1 to 256 characters in length, and cannot start with http:// or https://.
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
