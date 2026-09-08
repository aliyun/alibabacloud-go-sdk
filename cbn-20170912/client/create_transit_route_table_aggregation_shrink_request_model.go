// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouteTableAggregationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTransitRouteTableAggregationShrinkRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTransitRouteTableAggregationShrinkRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *CreateTransitRouteTableAggregationShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTransitRouteTableAggregationShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateTransitRouteTableAggregationShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTransitRouteTableAggregationShrinkRequest
	GetResourceOwnerId() *int64
	SetTransitRouteTableAggregationCidr(v string) *CreateTransitRouteTableAggregationShrinkRequest
	GetTransitRouteTableAggregationCidr() *string
	SetTransitRouteTableAggregationDescription(v string) *CreateTransitRouteTableAggregationShrinkRequest
	GetTransitRouteTableAggregationDescription() *string
	SetTransitRouteTableAggregationName(v string) *CreateTransitRouteTableAggregationShrinkRequest
	GetTransitRouteTableAggregationName() *string
	SetTransitRouteTableAggregationScope(v string) *CreateTransitRouteTableAggregationShrinkRequest
	GetTransitRouteTableAggregationScope() *string
	SetTransitRouteTableAggregationScopeListShrink(v string) *CreateTransitRouteTableAggregationShrinkRequest
	GetTransitRouteTableAggregationScopeListShrink() *string
	SetTransitRouteTableId(v string) *CreateTransitRouteTableAggregationShrinkRequest
	GetTransitRouteTableId() *string
}

type CreateTransitRouteTableAggregationShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. The dry run checks items such as permissions and instance status. Valid values:
	//
	// - **false*	- (default): sends a normal request and directly creates the aggregate route after the request passes the check.
	//
	// - **true**: sends a check request without creating the aggregate route. The check items include required parameters and request format. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The destination CIDR block of the aggregate route.
	//
	// > The following CIDR blocks are not supported:
	//
	// - CIDR blocks that start with "0" or "100.64"
	//
	// - Multicast addresses (224.0.0.1 to 239.255.255.254)
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.10.0/24
	TransitRouteTableAggregationCidr *string `json:"TransitRouteTableAggregationCidr,omitempty" xml:"TransitRouteTableAggregationCidr,omitempty"`
	// The description of the aggregate route.
	//
	// The description can be empty or 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	TransitRouteTableAggregationDescription *string `json:"TransitRouteTableAggregationDescription,omitempty" xml:"TransitRouteTableAggregationDescription,omitempty"`
	// The name of the aggregate route.
	//
	// The name can be empty or 1 to 128 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	TransitRouteTableAggregationName *string `json:"TransitRouteTableAggregationName,omitempty" xml:"TransitRouteTableAggregationName,omitempty"`
	// The propagation scope of the aggregate route.
	//
	// Set the value to **VPC**, which indicates that the aggregate route is propagated to all VPC-connected instances that have established an associated forwarding relationship with the current Enterprise Edition transit router route table and have the route synchronization feature enabled.
	//
	// example:
	//
	// VPC
	TransitRouteTableAggregationScope *string `json:"TransitRouteTableAggregationScope,omitempty" xml:"TransitRouteTableAggregationScope,omitempty"`
	// The propagation scope list of the aggregate route.
	//
	// >You must specify at least one of the propagation scope and the propagation scope list. We recommend that you use the propagation scope list. The elements in the propagation scope list cannot duplicate the value of the propagation scope.
	TransitRouteTableAggregationScopeListShrink *string `json:"TransitRouteTableAggregationScopeList,omitempty" xml:"TransitRouteTableAggregationScopeList,omitempty"`
	// The ID of the Enterprise Edition transit router route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-iq8qgruq1ry8jc7vt****
	TransitRouteTableId *string `json:"TransitRouteTableId,omitempty" xml:"TransitRouteTableId,omitempty"`
}

func (s CreateTransitRouteTableAggregationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouteTableAggregationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) GetTransitRouteTableAggregationCidr() *string {
	return s.TransitRouteTableAggregationCidr
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) GetTransitRouteTableAggregationDescription() *string {
	return s.TransitRouteTableAggregationDescription
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) GetTransitRouteTableAggregationName() *string {
	return s.TransitRouteTableAggregationName
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) GetTransitRouteTableAggregationScope() *string {
	return s.TransitRouteTableAggregationScope
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) GetTransitRouteTableAggregationScopeListShrink() *string {
	return s.TransitRouteTableAggregationScopeListShrink
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) GetTransitRouteTableId() *string {
	return s.TransitRouteTableId
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) SetClientToken(v string) *CreateTransitRouteTableAggregationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) SetDryRun(v bool) *CreateTransitRouteTableAggregationShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) SetOwnerAccount(v string) *CreateTransitRouteTableAggregationShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) SetOwnerId(v int64) *CreateTransitRouteTableAggregationShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) SetResourceOwnerAccount(v string) *CreateTransitRouteTableAggregationShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) SetResourceOwnerId(v int64) *CreateTransitRouteTableAggregationShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) SetTransitRouteTableAggregationCidr(v string) *CreateTransitRouteTableAggregationShrinkRequest {
	s.TransitRouteTableAggregationCidr = &v
	return s
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) SetTransitRouteTableAggregationDescription(v string) *CreateTransitRouteTableAggregationShrinkRequest {
	s.TransitRouteTableAggregationDescription = &v
	return s
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) SetTransitRouteTableAggregationName(v string) *CreateTransitRouteTableAggregationShrinkRequest {
	s.TransitRouteTableAggregationName = &v
	return s
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) SetTransitRouteTableAggregationScope(v string) *CreateTransitRouteTableAggregationShrinkRequest {
	s.TransitRouteTableAggregationScope = &v
	return s
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) SetTransitRouteTableAggregationScopeListShrink(v string) *CreateTransitRouteTableAggregationShrinkRequest {
	s.TransitRouteTableAggregationScopeListShrink = &v
	return s
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) SetTransitRouteTableId(v string) *CreateTransitRouteTableAggregationShrinkRequest {
	s.TransitRouteTableId = &v
	return s
}

func (s *CreateTransitRouteTableAggregationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
