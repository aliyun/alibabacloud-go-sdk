// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceTransitRouterRouteTableAssociationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ReplaceTransitRouterRouteTableAssociationRequest
	GetClientToken() *string
	SetDryRun(v bool) *ReplaceTransitRouterRouteTableAssociationRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ReplaceTransitRouterRouteTableAssociationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReplaceTransitRouterRouteTableAssociationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReplaceTransitRouterRouteTableAssociationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReplaceTransitRouterRouteTableAssociationRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentId(v string) *ReplaceTransitRouterRouteTableAssociationRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterRouteTableId(v string) *ReplaceTransitRouterRouteTableAssociationRequest
	GetTransitRouterRouteTableId() *string
}

type ReplaceTransitRouterRouteTableAssociationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the API request. Valid values:
	//
	// 	- **true**: prechecks the request but does not associate the network instance connection with another route table. The system checks whether the required parameters are set, whether the formats of the values are valid, and the service limits. If the request fails to pass the precheck, the corresponding error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: sends the request. After the request passes the authentication, the network instance connection is associated with the specified route table. This is the default value.
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
	// tr-attach-071g5j5tefg4x6****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the route table with which you want to associate the network instance connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp1cprmc6xmzjd66i****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s ReplaceTransitRouterRouteTableAssociationRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceTransitRouterRouteTableAssociationRequest) GoString() string {
	return s.String()
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetClientToken(v string) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetDryRun(v bool) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetOwnerAccount(v string) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetOwnerId(v int64) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.OwnerId = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetResourceOwnerAccount(v string) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetResourceOwnerId(v int64) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetTransitRouterAttachmentId(v string) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) SetTransitRouterRouteTableId(v string) *ReplaceTransitRouterRouteTableAssociationRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationRequest) Validate() error {
	return dara.Validate(s)
}
