// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterVpcAttachmentAttributeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPublishRouteEnabled(v bool) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest
	GetAutoPublishRouteEnabled() *bool
	SetClientToken(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest
	GetDryRun() *bool
	SetOptionsShrink(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest
	GetOptionsShrink() *string
	SetOrderType(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentDescription(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest
	GetTransitRouterAttachmentDescription() *string
	SetTransitRouterAttachmentId(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterAttachmentName(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest
	GetTransitRouterAttachmentName() *string
	SetTransitRouterVPCAttachmentOptionsShrink(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest
	GetTransitRouterVPCAttachmentOptionsShrink() *string
}

type UpdateTransitRouterVpcAttachmentAttributeShrinkRequest struct {
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
	OptionsShrink *string `json:"Options,omitempty" xml:"Options,omitempty"`
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
	TransitRouterVPCAttachmentOptionsShrink *string `json:"TransitRouterVPCAttachmentOptions,omitempty" xml:"TransitRouterVPCAttachmentOptions,omitempty"`
}

func (s UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) GetAutoPublishRouteEnabled() *bool {
	return s.AutoPublishRouteEnabled
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) GetOptionsShrink() *string {
	return s.OptionsShrink
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) GetTransitRouterVPCAttachmentOptionsShrink() *string {
	return s.TransitRouterVPCAttachmentOptionsShrink
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) SetAutoPublishRouteEnabled(v bool) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) SetClientToken(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) SetDryRun(v bool) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) SetOptionsShrink(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest {
	s.OptionsShrink = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) SetOrderType(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) SetOwnerAccount(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) SetOwnerId(v int64) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) SetResourceOwnerAccount(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) SetResourceOwnerId(v int64) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) SetTransitRouterAttachmentDescription(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) SetTransitRouterAttachmentId(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) SetTransitRouterAttachmentName(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) SetTransitRouterVPCAttachmentOptionsShrink(v string) *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest {
	s.TransitRouterVPCAttachmentOptionsShrink = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentAttributeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
