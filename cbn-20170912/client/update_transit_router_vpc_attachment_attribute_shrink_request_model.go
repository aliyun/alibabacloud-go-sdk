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
	// Specifies whether to allow the Enterprise Edition transit router to automatically advertise routes to the VPC.
	//
	// - **false**: The transit router does not automatically advertise routes.
	//
	// - **true**: The transit router automatically advertises routes.
	//
	// example:
	//
	// true
	AutoPublishRouteEnabled *bool `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the token from your client, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **ClientToken**. The **RequestId*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, which checks for issues such as permissions and instance status. Valid values:
	//
	// - **false*	- (default): sends a normal request. After the request passes the check, the system modifies the name and description of the VPC connection.
	//
	// - **true**: sends a check request. The system validates the request without modifying the VPC connection. If the check passes, the system returns the ID of the request. Otherwise, the system returns an error.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The billing method.
	OptionsShrink *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The billing method.
	//
	// example:
	//
	// PayByCenOwner
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new description of the VPC connection.
	//
	// The description can be empty or 1 to 256 characters in length, and cannot start with http\\:// or https\\://.
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
	// The new name of the VPC connection.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// testname
	TransitRouterAttachmentName *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	// The feature properties of the VPC connection. This parameter is deprecated. We recommend that you use the `Options` parameter.
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
