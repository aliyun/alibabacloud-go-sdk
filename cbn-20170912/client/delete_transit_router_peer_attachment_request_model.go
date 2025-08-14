// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterPeerAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTransitRouterPeerAttachmentRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTransitRouterPeerAttachmentRequest
	GetDryRun() *bool
	SetForce(v bool) *DeleteTransitRouterPeerAttachmentRequest
	GetForce() *bool
	SetOwnerAccount(v string) *DeleteTransitRouterPeerAttachmentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTransitRouterPeerAttachmentRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTransitRouterPeerAttachmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTransitRouterPeerAttachmentRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentId(v string) *DeleteTransitRouterPeerAttachmentRequest
	GetTransitRouterAttachmentId() *string
}

type DeleteTransitRouterPeerAttachmentRequest struct {
	// The client token that you want to use to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **false*	- (default): performs a dry run and sends the request.
	//
	// 	- **true**: performs a dry run. The system checks the required parameters and request syntax. If the request fails to pass the check, an error message is returned. If the request passes the dry run, the system returns the ID of the request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully delete the inter-region connection. Valid values:
	//
	// 	- **false*	- (default): Check for relevant resources, including associated forwarding and route learning, before deleting the inter-region connection. If such a resource exists, the VPC connection is not deleted and an error message is returned.
	//
	// 	- **true**: Delete the inter-region connection and all relevant resources.
	//
	// example:
	//
	// false
	Force                *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the inter-region connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-gyjhtx9sjsxhm6****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s DeleteTransitRouterPeerAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterPeerAttachmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPeerAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTransitRouterPeerAttachmentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTransitRouterPeerAttachmentRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteTransitRouterPeerAttachmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTransitRouterPeerAttachmentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTransitRouterPeerAttachmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTransitRouterPeerAttachmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTransitRouterPeerAttachmentRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *DeleteTransitRouterPeerAttachmentRequest) SetClientToken(v string) *DeleteTransitRouterPeerAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentRequest) SetDryRun(v bool) *DeleteTransitRouterPeerAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentRequest) SetForce(v bool) *DeleteTransitRouterPeerAttachmentRequest {
	s.Force = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentRequest) SetOwnerAccount(v string) *DeleteTransitRouterPeerAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentRequest) SetOwnerId(v int64) *DeleteTransitRouterPeerAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterPeerAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterPeerAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentRequest) SetTransitRouterAttachmentId(v string) *DeleteTransitRouterPeerAttachmentRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *DeleteTransitRouterPeerAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
