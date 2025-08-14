// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableTransitRouterRouteTablePropagationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DisableTransitRouterRouteTablePropagationRequest
	GetClientToken() *string
	SetDryRun(v bool) *DisableTransitRouterRouteTablePropagationRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DisableTransitRouterRouteTablePropagationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DisableTransitRouterRouteTablePropagationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DisableTransitRouterRouteTablePropagationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DisableTransitRouterRouteTablePropagationRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentId(v string) *DisableTransitRouterRouteTablePropagationRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterRouteTableId(v string) *DisableTransitRouterRouteTablePropagationRequest
	GetTransitRouterRouteTableId() *string
}

type DisableTransitRouterRouteTablePropagationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Default values:
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
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
	// tr-attach-vx6iwhjr1x1j78****
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

func (s DisableTransitRouterRouteTablePropagationRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableTransitRouterRouteTablePropagationRequest) GoString() string {
	return s.String()
}

func (s *DisableTransitRouterRouteTablePropagationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableTransitRouterRouteTablePropagationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DisableTransitRouterRouteTablePropagationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DisableTransitRouterRouteTablePropagationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableTransitRouterRouteTablePropagationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisableTransitRouterRouteTablePropagationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisableTransitRouterRouteTablePropagationRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *DisableTransitRouterRouteTablePropagationRequest) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetClientToken(v string) *DisableTransitRouterRouteTablePropagationRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetDryRun(v bool) *DisableTransitRouterRouteTablePropagationRequest {
	s.DryRun = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetOwnerAccount(v string) *DisableTransitRouterRouteTablePropagationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetOwnerId(v int64) *DisableTransitRouterRouteTablePropagationRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetResourceOwnerAccount(v string) *DisableTransitRouterRouteTablePropagationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetResourceOwnerId(v int64) *DisableTransitRouterRouteTablePropagationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetTransitRouterAttachmentId(v string) *DisableTransitRouterRouteTablePropagationRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationRequest) SetTransitRouterRouteTableId(v string) *DisableTransitRouterRouteTablePropagationRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *DisableTransitRouterRouteTablePropagationRequest) Validate() error {
	return dara.Validate(s)
}
