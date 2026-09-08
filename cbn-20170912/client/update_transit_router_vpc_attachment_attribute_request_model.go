// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterVpcAttachmentAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPublishRouteEnabled(v bool) *UpdateTransitRouterVpcAttachmentAttributeRequest
	GetAutoPublishRouteEnabled() *bool
	SetClientToken(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateTransitRouterVpcAttachmentAttributeRequest
	GetDryRun() *bool
	SetOptions(v *UpdateTransitRouterVpcAttachmentAttributeRequestOptions) *UpdateTransitRouterVpcAttachmentAttributeRequest
	GetOptions() *UpdateTransitRouterVpcAttachmentAttributeRequestOptions
	SetOrderType(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateTransitRouterVpcAttachmentAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTransitRouterVpcAttachmentAttributeRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentDescription(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest
	GetTransitRouterAttachmentDescription() *string
	SetTransitRouterAttachmentId(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterAttachmentName(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest
	GetTransitRouterAttachmentName() *string
	SetTransitRouterVPCAttachmentOptions(v map[string]*string) *UpdateTransitRouterVpcAttachmentAttributeRequest
	GetTransitRouterVPCAttachmentOptions() map[string]*string
}

type UpdateTransitRouterVpcAttachmentAttributeRequest struct {
	// Specifies whether to allow the Enterprise Edition forward router to automatically publish route entry to the VPC-connected instance.
	//
	// - **false**: no.
	//
	// - **true**: yes.
	//
	// example:
	//
	// true
	AutoPublishRouteEnabled *bool `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
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
	// Specifies whether to perform a dry run, including permission and instance status verification. Valid values:
	//
	// - **false*	- (default): Sends a normal request. If the request passes the check, the name and description of the VPC connection are modified.
	//
	// - **true**: Sends a check request. Only the verification is performed, and the name and description of the VPC connection are not modified. The system checks whether the required parameters are specified and whether the request format is valid. If the check fails, the corresponding error is returned. If the check passes, the corresponding request ID is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The collection of feature attributes.
	Options *UpdateTransitRouterVpcAttachmentAttributeRequestOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	// The payer of the network instance.
	//
	// example:
	//
	// PayByCenOwner
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The description of the VPC connection.
	//
	// The description can be empty or 1 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// testdesc
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The ID of the VPC connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-nls9fzkfat8934****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The name of the VPC connection.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// testname
	TransitRouterAttachmentName *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	// The list of feature attributes of the VPC connection (to be deprecated. Use the new parameter Options instead).
	TransitRouterVPCAttachmentOptions map[string]*string `json:"TransitRouterVPCAttachmentOptions,omitempty" xml:"TransitRouterVPCAttachmentOptions,omitempty"`
}

func (s UpdateTransitRouterVpcAttachmentAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterVpcAttachmentAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) GetAutoPublishRouteEnabled() *bool {
	return s.AutoPublishRouteEnabled
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) GetOptions() *UpdateTransitRouterVpcAttachmentAttributeRequestOptions {
	return s.Options
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) GetTransitRouterVPCAttachmentOptions() map[string]*string {
	return s.TransitRouterVPCAttachmentOptions
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetAutoPublishRouteEnabled(v bool) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetClientToken(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetDryRun(v bool) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetOptions(v *UpdateTransitRouterVpcAttachmentAttributeRequestOptions) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.Options = v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetOrderType(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.OrderType = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetOwnerAccount(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetOwnerId(v int64) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetResourceOwnerAccount(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetResourceOwnerId(v int64) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetTransitRouterAttachmentDescription(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetTransitRouterAttachmentId(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetTransitRouterAttachmentName(v string) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) SetTransitRouterVPCAttachmentOptions(v map[string]*string) *UpdateTransitRouterVpcAttachmentAttributeRequest {
	s.TransitRouterVPCAttachmentOptions = v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequest) Validate() error {
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTransitRouterVpcAttachmentAttributeRequestOptions struct {
	// Specifies whether to enable the appliance mode for traffic redirection.
	//
	// - **disable*	- (default): no.
	//
	// - **enable**: yes.
	//
	// example:
	//
	// enable
	ApplianceModeSupport *string `json:"ApplianceModeSupport,omitempty" xml:"ApplianceModeSupport,omitempty"`
	// Specifies whether IPv6 is supported.
	//
	// - **disable*	- (default): no.
	//
	// - **enable**: yes.
	//
	// example:
	//
	// enable
	Ipv6Support *string `json:"Ipv6Support,omitempty" xml:"Ipv6Support,omitempty"`
}

func (s UpdateTransitRouterVpcAttachmentAttributeRequestOptions) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterVpcAttachmentAttributeRequestOptions) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequestOptions) GetApplianceModeSupport() *string {
	return s.ApplianceModeSupport
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequestOptions) GetIpv6Support() *string {
	return s.Ipv6Support
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequestOptions) SetApplianceModeSupport(v string) *UpdateTransitRouterVpcAttachmentAttributeRequestOptions {
	s.ApplianceModeSupport = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequestOptions) SetIpv6Support(v string) *UpdateTransitRouterVpcAttachmentAttributeRequestOptions {
	s.Ipv6Support = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeRequestOptions) Validate() error {
	return dara.Validate(s)
}
