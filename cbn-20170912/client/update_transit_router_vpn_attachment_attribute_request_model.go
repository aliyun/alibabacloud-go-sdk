// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterVpnAttachmentAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPublishRouteEnabled(v bool) *UpdateTransitRouterVpnAttachmentAttributeRequest
	GetAutoPublishRouteEnabled() *bool
	SetClientToken(v string) *UpdateTransitRouterVpnAttachmentAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateTransitRouterVpnAttachmentAttributeRequest
	GetDryRun() *bool
	SetOrderType(v string) *UpdateTransitRouterVpnAttachmentAttributeRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *UpdateTransitRouterVpnAttachmentAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateTransitRouterVpnAttachmentAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateTransitRouterVpnAttachmentAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTransitRouterVpnAttachmentAttributeRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentDescription(v string) *UpdateTransitRouterVpnAttachmentAttributeRequest
	GetTransitRouterAttachmentDescription() *string
	SetTransitRouterAttachmentId(v string) *UpdateTransitRouterVpnAttachmentAttributeRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterAttachmentName(v string) *UpdateTransitRouterVpnAttachmentAttributeRequest
	GetTransitRouterAttachmentName() *string
}

type UpdateTransitRouterVpnAttachmentAttributeRequest struct {
	// Specifies whether to allow the Enterprise Edition transit router to automatically publish routes to the IPsec-VPN connection. Valid values:
	//
	// - **true**: The Enterprise Edition transit router automatically publishes routes to the IPsec-VPN connection.
	//
	// - **false**: The Enterprise Edition transit router does not automatically publish routes to the IPsec-VPN connection.
	//
	// example:
	//
	// true
	AutoPublishRouteEnabled *bool `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	// A client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **ClientToken**. The **RequestId*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Performs a dry run to check the request without modifying the attributes of the VPN connection. The system checks the required parameters, request format, and service limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (Default): Sends a normal request. If the request passes the check, the system modifies the attributes of the VPN connection.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// PayByCenOwner
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new description of the VPN connection.
	//
	// The description can be empty or 1 to 256 characters in length, and cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// desctest
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The ID of the VPN connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-a6p8voaodog5c0****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The new name of the VPN connection.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// nametest
	TransitRouterAttachmentName *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
}

func (s UpdateTransitRouterVpnAttachmentAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterVpnAttachmentAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) GetAutoPublishRouteEnabled() *bool {
	return s.AutoPublishRouteEnabled
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) SetAutoPublishRouteEnabled(v bool) *UpdateTransitRouterVpnAttachmentAttributeRequest {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) SetClientToken(v string) *UpdateTransitRouterVpnAttachmentAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) SetDryRun(v bool) *UpdateTransitRouterVpnAttachmentAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) SetOrderType(v string) *UpdateTransitRouterVpnAttachmentAttributeRequest {
	s.OrderType = &v
	return s
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) SetOwnerAccount(v string) *UpdateTransitRouterVpnAttachmentAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) SetOwnerId(v int64) *UpdateTransitRouterVpnAttachmentAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) SetResourceOwnerAccount(v string) *UpdateTransitRouterVpnAttachmentAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) SetResourceOwnerId(v int64) *UpdateTransitRouterVpnAttachmentAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) SetTransitRouterAttachmentDescription(v string) *UpdateTransitRouterVpnAttachmentAttributeRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) SetTransitRouterAttachmentId(v string) *UpdateTransitRouterVpnAttachmentAttributeRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) SetTransitRouterAttachmentName(v string) *UpdateTransitRouterVpnAttachmentAttributeRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *UpdateTransitRouterVpnAttachmentAttributeRequest) Validate() error {
	return dara.Validate(s)
}
