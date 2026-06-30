// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterVbrAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTransitRouterVbrAttachmentRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTransitRouterVbrAttachmentRequest
	GetDryRun() *bool
	SetForce(v bool) *DeleteTransitRouterVbrAttachmentRequest
	GetForce() *bool
	SetOwnerAccount(v string) *DeleteTransitRouterVbrAttachmentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTransitRouterVbrAttachmentRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTransitRouterVbrAttachmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTransitRouterVbrAttachmentRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentId(v string) *DeleteTransitRouterVbrAttachmentRequest
	GetTransitRouterAttachmentId() *string
}

type DeleteTransitRouterVbrAttachmentRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// Use the client to generate the token, but you must make sure that the token is unique among requests. The token can contain only ASCII characters.
	//
	// > When left empty, the system automatically uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- is different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run to check permissions and instance status. Valid values:
	//
	// - **false*	- (default): sends a request and deletes the VBR connection.
	//
	// - **true**: sends a check request without deleting the VBR connection. The system checks the required parameters and request syntax. If the request fails the dry run, an error code is returned. If the request passes the dry run, a request ID is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcibly delete the VBR connection. Valid values:
	//
	// - **false*	- (default): The system checks resources, such as forwarding associations or route learning. If there are such resources, the VBR connection is not deleted and an error code is returned.
	//
	// - **true**: When the VBR connection is deleted, all associated resources are also deleted.
	//
	// example:
	//
	// false
	Force                *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VBR connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-9nlnjv7by7n7a****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s DeleteTransitRouterVbrAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterVbrAttachmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterVbrAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTransitRouterVbrAttachmentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTransitRouterVbrAttachmentRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteTransitRouterVbrAttachmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTransitRouterVbrAttachmentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTransitRouterVbrAttachmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTransitRouterVbrAttachmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTransitRouterVbrAttachmentRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *DeleteTransitRouterVbrAttachmentRequest) SetClientToken(v string) *DeleteTransitRouterVbrAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentRequest) SetDryRun(v bool) *DeleteTransitRouterVbrAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentRequest) SetForce(v bool) *DeleteTransitRouterVbrAttachmentRequest {
	s.Force = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentRequest) SetOwnerAccount(v string) *DeleteTransitRouterVbrAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentRequest) SetOwnerId(v int64) *DeleteTransitRouterVbrAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterVbrAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterVbrAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentRequest) SetTransitRouterAttachmentId(v string) *DeleteTransitRouterVbrAttachmentRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
