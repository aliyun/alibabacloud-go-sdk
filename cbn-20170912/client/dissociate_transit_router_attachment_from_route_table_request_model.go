// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateTransitRouterAttachmentFromRouteTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DissociateTransitRouterAttachmentFromRouteTableRequest
	GetClientToken() *string
	SetDryRun(v bool) *DissociateTransitRouterAttachmentFromRouteTableRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DissociateTransitRouterAttachmentFromRouteTableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DissociateTransitRouterAttachmentFromRouteTableRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DissociateTransitRouterAttachmentFromRouteTableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DissociateTransitRouterAttachmentFromRouteTableRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentId(v string) *DissociateTransitRouterAttachmentFromRouteTableRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterRouteTableId(v string) *DissociateTransitRouterAttachmentFromRouteTableRequest
	GetTransitRouterRouteTableId() *string
}

type DissociateTransitRouterAttachmentFromRouteTableRequest struct {
	// A client token that ensures the idempotence of the request.
	//
	// Generate a unique token on your client. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **false*	- (default): Sends a normal request. If the request passes the check, the network instance connection is dissociated from the route table.
	//
	// - **true**: Sends a dry run request to check the required parameters, request format, and permissions. An error message is returned if the request fails the dry run. The corresponding request ID is returned if the request passes the dry run.
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

func (s DissociateTransitRouterAttachmentFromRouteTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateTransitRouterAttachmentFromRouteTableRequest) GoString() string {
	return s.String()
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetClientToken(v string) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetDryRun(v bool) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.DryRun = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetOwnerAccount(v string) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetOwnerId(v int64) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.OwnerId = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetResourceOwnerAccount(v string) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetResourceOwnerId(v int64) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetTransitRouterAttachmentId(v string) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) SetTransitRouterRouteTableId(v string) *DissociateTransitRouterAttachmentFromRouteTableRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *DissociateTransitRouterAttachmentFromRouteTableRequest) Validate() error {
	return dara.Validate(s)
}
