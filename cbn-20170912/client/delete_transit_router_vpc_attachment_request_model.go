// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterVpcAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTransitRouterVpcAttachmentRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTransitRouterVpcAttachmentRequest
	GetDryRun() *bool
	SetForce(v bool) *DeleteTransitRouterVpcAttachmentRequest
	GetForce() *bool
	SetOwnerAccount(v string) *DeleteTransitRouterVpcAttachmentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTransitRouterVpcAttachmentRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTransitRouterVpcAttachmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTransitRouterVpcAttachmentRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentId(v string) *DeleteTransitRouterVpcAttachmentRequest
	GetTransitRouterAttachmentId() *string
}

type DeleteTransitRouterVpcAttachmentRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among all requests. The token can contain only ASCII characters.
	//
	// > If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Default values:
	//
	// 	- **false*	- (default): performs a dry run and sends the request.
	//
	// 	- **true**: performs a dry run. The system checks the required parameters and request syntax. If the request fails the dry run, an error message is returned. If the request passes the dry run, the system returns the ID of the request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully delete the VPC connection. Valid values:
	//
	// 	- **false*	- (default): checks resources such as associated forwarding correlations and route learning policies that are related to the VPC connection before it is deleted. If such a resource exists, the VPC connection is not deleted and an error message is returned.
	//
	// 	- **true**: deletes the VPC connection and all resources that are related to the VPC connection.
	//
	// example:
	//
	// false
	Force                *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPC connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-ia340z7xis7t5s****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s DeleteTransitRouterVpcAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterVpcAttachmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterVpcAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTransitRouterVpcAttachmentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTransitRouterVpcAttachmentRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteTransitRouterVpcAttachmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTransitRouterVpcAttachmentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTransitRouterVpcAttachmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTransitRouterVpcAttachmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTransitRouterVpcAttachmentRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *DeleteTransitRouterVpcAttachmentRequest) SetClientToken(v string) *DeleteTransitRouterVpcAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentRequest) SetDryRun(v bool) *DeleteTransitRouterVpcAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentRequest) SetForce(v bool) *DeleteTransitRouterVpcAttachmentRequest {
	s.Force = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentRequest) SetOwnerAccount(v string) *DeleteTransitRouterVpcAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentRequest) SetOwnerId(v int64) *DeleteTransitRouterVpcAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterVpcAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterVpcAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentRequest) SetTransitRouterAttachmentId(v string) *DeleteTransitRouterVpcAttachmentRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
