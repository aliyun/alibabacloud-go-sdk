// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateTransitRouterAttachmentWithRouteTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AssociateTransitRouterAttachmentWithRouteTableRequest
	GetClientToken() *string
	SetDryRun(v bool) *AssociateTransitRouterAttachmentWithRouteTableRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *AssociateTransitRouterAttachmentWithRouteTableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociateTransitRouterAttachmentWithRouteTableRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AssociateTransitRouterAttachmentWithRouteTableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateTransitRouterAttachmentWithRouteTableRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentId(v string) *AssociateTransitRouterAttachmentWithRouteTableRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterRouteTableId(v string) *AssociateTransitRouterAttachmentWithRouteTableRequest
	GetTransitRouterRouteTableId() *string
}

type AssociateTransitRouterAttachmentWithRouteTableRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a parameter value from your client to make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **ClientToken**. The **RequestId*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run to check for potential issues, such as permissions and instance status. Valid values:
	//
	// - **false*	- (default): sends a normal request. An association is created after the request passes the check.
	//
	// - **true**: sends a check request to perform a dry run. The system checks the required parameters, request format, and other items. No association is created. If the check fails, an error message is returned. If the check passes, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the network instance connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-nls9fzkfat8934****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the route table of the Enterprise Edition transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp1dudbh2d5na6b50****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s AssociateTransitRouterAttachmentWithRouteTableRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateTransitRouterAttachmentWithRouteTableRequest) GoString() string {
	return s.String()
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetClientToken(v string) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetDryRun(v bool) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetOwnerAccount(v string) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetOwnerId(v int64) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetResourceOwnerAccount(v string) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetResourceOwnerId(v int64) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetTransitRouterAttachmentId(v string) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) SetTransitRouterRouteTableId(v string) *AssociateTransitRouterAttachmentWithRouteTableRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *AssociateTransitRouterAttachmentWithRouteTableRequest) Validate() error {
	return dara.Validate(s)
}
