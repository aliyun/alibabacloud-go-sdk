// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterVpnAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTransitRouterVpnAttachmentRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTransitRouterVpnAttachmentRequest
	GetDryRun() *bool
	SetForce(v bool) *DeleteTransitRouterVpnAttachmentRequest
	GetForce() *bool
	SetOwnerAccount(v string) *DeleteTransitRouterVpnAttachmentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTransitRouterVpnAttachmentRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTransitRouterVpnAttachmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTransitRouterVpnAttachmentRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentId(v string) *DeleteTransitRouterVpnAttachmentRequest
	GetTransitRouterAttachmentId() *string
}

type DeleteTransitRouterVpnAttachmentRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without deleting the VPN connection. The system checks the required parameters, request syntax, and business restrictions. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): performs a dry run and then deletes the VPN connection after the request passes the dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to force delete the VPN connection. Valid values:
	//
	// - **false*	- (default): checks whether related resource dependencies exist before deleting the VPN connection, such as associated forwarding and routing learning. If dependencies exist, the VPN connection is not deleted and the corresponding error is returned.
	//
	// - **true**: deletes the VPN connection along with all related dependencies.
	//
	// example:
	//
	// false
	Force                *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPN connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-b9xj1dv69600kj****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s DeleteTransitRouterVpnAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterVpnAttachmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterVpnAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTransitRouterVpnAttachmentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTransitRouterVpnAttachmentRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteTransitRouterVpnAttachmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTransitRouterVpnAttachmentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTransitRouterVpnAttachmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTransitRouterVpnAttachmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTransitRouterVpnAttachmentRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *DeleteTransitRouterVpnAttachmentRequest) SetClientToken(v string) *DeleteTransitRouterVpnAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterVpnAttachmentRequest) SetDryRun(v bool) *DeleteTransitRouterVpnAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterVpnAttachmentRequest) SetForce(v bool) *DeleteTransitRouterVpnAttachmentRequest {
	s.Force = &v
	return s
}

func (s *DeleteTransitRouterVpnAttachmentRequest) SetOwnerAccount(v string) *DeleteTransitRouterVpnAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterVpnAttachmentRequest) SetOwnerId(v int64) *DeleteTransitRouterVpnAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterVpnAttachmentRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterVpnAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterVpnAttachmentRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterVpnAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterVpnAttachmentRequest) SetTransitRouterAttachmentId(v string) *DeleteTransitRouterVpnAttachmentRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *DeleteTransitRouterVpnAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
