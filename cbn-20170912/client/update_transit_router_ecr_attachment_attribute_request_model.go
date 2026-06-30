// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterEcrAttachmentAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateTransitRouterEcrAttachmentAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateTransitRouterEcrAttachmentAttributeRequest
	GetDryRun() *bool
	SetOrderType(v string) *UpdateTransitRouterEcrAttachmentAttributeRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *UpdateTransitRouterEcrAttachmentAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateTransitRouterEcrAttachmentAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateTransitRouterEcrAttachmentAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTransitRouterEcrAttachmentAttributeRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentDescription(v string) *UpdateTransitRouterEcrAttachmentAttributeRequest
	GetTransitRouterAttachmentDescription() *string
	SetTransitRouterAttachmentId(v string) *UpdateTransitRouterEcrAttachmentAttributeRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterAttachmentName(v string) *UpdateTransitRouterEcrAttachmentAttributeRequest
	GetTransitRouterAttachmentName() *string
}

type UpdateTransitRouterEcrAttachmentAttributeRequest struct {
	// The client token that ensures the idempotence of the request.
	//
	// You can generate a token from your client, but you must ensure that it is unique across requests. The `ClientToken` can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run to check for potential issues, including permissions and instance status. Valid values:
	//
	// - **false*	- (default): Sends a normal request. The system modifies the ECR attachment attributes if the request passes the check.
	//
	// - **true**: Sends a check request only. The system validates the request but does not modify the ECR attachment attributes. If the check fails, an error is returned. If the check passes, the system returns the request ID.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The payer for the network instance. This operation does not support changing the payer for an ECR attachment.
	//
	// example:
	//
	// PayByCenOwner
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new description of the ECR attachment.
	//
	// The description can be empty or 1 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// desctest
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The ID of the ECR attachment.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-r6g0m3epjehw57****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The new name of the ECR attachment.
	//
	// The name can be empty or 1 to 128 characters in length. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// nametest
	TransitRouterAttachmentName *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
}

func (s UpdateTransitRouterEcrAttachmentAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterEcrAttachmentAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) SetClientToken(v string) *UpdateTransitRouterEcrAttachmentAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) SetDryRun(v bool) *UpdateTransitRouterEcrAttachmentAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) SetOrderType(v string) *UpdateTransitRouterEcrAttachmentAttributeRequest {
	s.OrderType = &v
	return s
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) SetOwnerAccount(v string) *UpdateTransitRouterEcrAttachmentAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) SetOwnerId(v int64) *UpdateTransitRouterEcrAttachmentAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) SetResourceOwnerAccount(v string) *UpdateTransitRouterEcrAttachmentAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) SetResourceOwnerId(v int64) *UpdateTransitRouterEcrAttachmentAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) SetTransitRouterAttachmentDescription(v string) *UpdateTransitRouterEcrAttachmentAttributeRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) SetTransitRouterAttachmentId(v string) *UpdateTransitRouterEcrAttachmentAttributeRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) SetTransitRouterAttachmentName(v string) *UpdateTransitRouterEcrAttachmentAttributeRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *UpdateTransitRouterEcrAttachmentAttributeRequest) Validate() error {
	return dara.Validate(s)
}
