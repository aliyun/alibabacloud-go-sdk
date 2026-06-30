// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTransitRouteTableAggregationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyTransitRouteTableAggregationShrinkRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyTransitRouteTableAggregationShrinkRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ModifyTransitRouteTableAggregationShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyTransitRouteTableAggregationShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyTransitRouteTableAggregationShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyTransitRouteTableAggregationShrinkRequest
	GetResourceOwnerId() *int64
	SetTransitRouteTableAggregationCidr(v string) *ModifyTransitRouteTableAggregationShrinkRequest
	GetTransitRouteTableAggregationCidr() *string
	SetTransitRouteTableAggregationDescription(v string) *ModifyTransitRouteTableAggregationShrinkRequest
	GetTransitRouteTableAggregationDescription() *string
	SetTransitRouteTableAggregationName(v string) *ModifyTransitRouteTableAggregationShrinkRequest
	GetTransitRouteTableAggregationName() *string
	SetTransitRouteTableAggregationScope(v string) *ModifyTransitRouteTableAggregationShrinkRequest
	GetTransitRouteTableAggregationScope() *string
	SetTransitRouteTableAggregationScopeListShrink(v string) *ModifyTransitRouteTableAggregationShrinkRequest
	GetTransitRouteTableAggregationScopeListShrink() *string
	SetTransitRouteTableId(v string) *ModifyTransitRouteTableAggregationShrinkRequest
	GetTransitRouteTableId() *string
}

type ModifyTransitRouteTableAggregationShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a token from your client to make sure that the token is unique among different requests. The \\`ClientToken\\` parameter can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **false*	- (default): sends a normal request and modifies the aggregate route after the request passes the check.
	//
	// - **true**: sends a check request to perform a dry run. The system checks the required parameters, request format, and permissions. If the check fails, the corresponding error is returned. If the check passes, the \\`DryRunOperation\\` error code is returned. In this case, the aggregate route is not modified.
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
	// This parameter is required.
	//
	// example:
	//
	// 192.168.10.0/24
	TransitRouteTableAggregationCidr *string `json:"TransitRouteTableAggregationCidr,omitempty" xml:"TransitRouteTableAggregationCidr,omitempty"`
	// The description of the aggregate route.
	//
	// The description can be empty or 0 to 256 characters in length. It cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// desctest
	TransitRouteTableAggregationDescription *string `json:"TransitRouteTableAggregationDescription,omitempty" xml:"TransitRouteTableAggregationDescription,omitempty"`
	// The name of the aggregate route.
	//
	// The name can be empty or 1 to 128 characters in length. It cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// nametest
	TransitRouteTableAggregationName *string `json:"TransitRouteTableAggregationName,omitempty" xml:"TransitRouteTableAggregationName,omitempty"`
	// The propagation scope of the aggregate route.
	//
	// The only valid value is **VPC**. This value specifies that the aggregate route is propagated to all VPC instances that are associated with the route table of the Enterprise Edition transit router and have route synchronization enabled.
	//
	// example:
	//
	// VPC
	TransitRouteTableAggregationScope *string `json:"TransitRouteTableAggregationScope,omitempty" xml:"TransitRouteTableAggregationScope,omitempty"`
	// The list of propagation scopes for the aggregate route.
	//
	// > You must specify this parameter or \\`TransitRouteTableAggregationScope\\`. We recommend that you specify this parameter. The elements in this list cannot be the same as the value of \\`TransitRouteTableAggregationScope\\`.
	TransitRouteTableAggregationScopeListShrink *string `json:"TransitRouteTableAggregationScopeList,omitempty" xml:"TransitRouteTableAggregationScopeList,omitempty"`
	// The ID of the route table of the Enterprise Edition transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-6ehgc262hr170qgyc****
	TransitRouteTableId *string `json:"TransitRouteTableId,omitempty" xml:"TransitRouteTableId,omitempty"`
}

func (s ModifyTransitRouteTableAggregationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouteTableAggregationShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) GetTransitRouteTableAggregationCidr() *string {
	return s.TransitRouteTableAggregationCidr
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) GetTransitRouteTableAggregationDescription() *string {
	return s.TransitRouteTableAggregationDescription
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) GetTransitRouteTableAggregationName() *string {
	return s.TransitRouteTableAggregationName
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) GetTransitRouteTableAggregationScope() *string {
	return s.TransitRouteTableAggregationScope
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) GetTransitRouteTableAggregationScopeListShrink() *string {
	return s.TransitRouteTableAggregationScopeListShrink
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) GetTransitRouteTableId() *string {
	return s.TransitRouteTableId
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) SetClientToken(v string) *ModifyTransitRouteTableAggregationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) SetDryRun(v bool) *ModifyTransitRouteTableAggregationShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) SetOwnerAccount(v string) *ModifyTransitRouteTableAggregationShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) SetOwnerId(v int64) *ModifyTransitRouteTableAggregationShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) SetResourceOwnerAccount(v string) *ModifyTransitRouteTableAggregationShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) SetResourceOwnerId(v int64) *ModifyTransitRouteTableAggregationShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) SetTransitRouteTableAggregationCidr(v string) *ModifyTransitRouteTableAggregationShrinkRequest {
	s.TransitRouteTableAggregationCidr = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) SetTransitRouteTableAggregationDescription(v string) *ModifyTransitRouteTableAggregationShrinkRequest {
	s.TransitRouteTableAggregationDescription = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) SetTransitRouteTableAggregationName(v string) *ModifyTransitRouteTableAggregationShrinkRequest {
	s.TransitRouteTableAggregationName = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) SetTransitRouteTableAggregationScope(v string) *ModifyTransitRouteTableAggregationShrinkRequest {
	s.TransitRouteTableAggregationScope = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) SetTransitRouteTableAggregationScopeListShrink(v string) *ModifyTransitRouteTableAggregationShrinkRequest {
	s.TransitRouteTableAggregationScopeListShrink = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) SetTransitRouteTableId(v string) *ModifyTransitRouteTableAggregationShrinkRequest {
	s.TransitRouteTableId = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
