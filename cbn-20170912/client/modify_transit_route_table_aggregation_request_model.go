// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTransitRouteTableAggregationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyTransitRouteTableAggregationRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyTransitRouteTableAggregationRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ModifyTransitRouteTableAggregationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyTransitRouteTableAggregationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyTransitRouteTableAggregationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyTransitRouteTableAggregationRequest
	GetResourceOwnerId() *int64
	SetTransitRouteTableAggregationCidr(v string) *ModifyTransitRouteTableAggregationRequest
	GetTransitRouteTableAggregationCidr() *string
	SetTransitRouteTableAggregationDescription(v string) *ModifyTransitRouteTableAggregationRequest
	GetTransitRouteTableAggregationDescription() *string
	SetTransitRouteTableAggregationName(v string) *ModifyTransitRouteTableAggregationRequest
	GetTransitRouteTableAggregationName() *string
	SetTransitRouteTableAggregationScope(v string) *ModifyTransitRouteTableAggregationRequest
	GetTransitRouteTableAggregationScope() *string
	SetTransitRouteTableAggregationScopeList(v []*string) *ModifyTransitRouteTableAggregationRequest
	GetTransitRouteTableAggregationScopeList() []*string
	SetTransitRouteTableId(v string) *ModifyTransitRouteTableAggregationRequest
	GetTransitRouteTableId() *string
}

type ModifyTransitRouteTableAggregationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// Use the client to generate the token, but make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- may be different.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run to check information such as the permissions and instance status. Valid values:
	//
	// 	- **false*	- (default): sends the request. If the request passes the check, an Enterprise Edition transit router is created.
	//
	// 	- **true**: checks the request but does not create the Enterprise Edition transit router. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
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
	// The description can be empty or 0 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	TransitRouteTableAggregationDescription *string `json:"TransitRouteTableAggregationDescription,omitempty" xml:"TransitRouteTableAggregationDescription,omitempty"`
	// The name of the aggregate route.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	TransitRouteTableAggregationName *string `json:"TransitRouteTableAggregationName,omitempty" xml:"TransitRouteTableAggregationName,omitempty"`
	// The scope of networks that you want to advertise the aggregate route.
	//
	// The valid value is **VPC**, which indicates that the aggregate route is advertised to all VPCs that have associated forwarding correlation with the Enterprise Edition transit router and have route synchronization enabled.
	//
	// example:
	//
	// VPC
	TransitRouteTableAggregationScope *string `json:"TransitRouteTableAggregationScope,omitempty" xml:"TransitRouteTableAggregationScope,omitempty"`
	// The scope of networks to which the aggregate route is advertised.
	//
	// >  You must select at least one attribute from either the Aggregate Route Propagation Range or the Aggregate Route Propagation Range List. We recommend using the latter. The elements of the two attributes cannot duplicate.
	TransitRouteTableAggregationScopeList []*string `json:"TransitRouteTableAggregationScopeList,omitempty" xml:"TransitRouteTableAggregationScopeList,omitempty" type:"Repeated"`
	// The list of route table IDs of the Enterprise Edition transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-6ehgc262hr170qgyc****
	TransitRouteTableId *string `json:"TransitRouteTableId,omitempty" xml:"TransitRouteTableId,omitempty"`
}

func (s ModifyTransitRouteTableAggregationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouteTableAggregationRequest) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouteTableAggregationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyTransitRouteTableAggregationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyTransitRouteTableAggregationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyTransitRouteTableAggregationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyTransitRouteTableAggregationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyTransitRouteTableAggregationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyTransitRouteTableAggregationRequest) GetTransitRouteTableAggregationCidr() *string {
	return s.TransitRouteTableAggregationCidr
}

func (s *ModifyTransitRouteTableAggregationRequest) GetTransitRouteTableAggregationDescription() *string {
	return s.TransitRouteTableAggregationDescription
}

func (s *ModifyTransitRouteTableAggregationRequest) GetTransitRouteTableAggregationName() *string {
	return s.TransitRouteTableAggregationName
}

func (s *ModifyTransitRouteTableAggregationRequest) GetTransitRouteTableAggregationScope() *string {
	return s.TransitRouteTableAggregationScope
}

func (s *ModifyTransitRouteTableAggregationRequest) GetTransitRouteTableAggregationScopeList() []*string {
	return s.TransitRouteTableAggregationScopeList
}

func (s *ModifyTransitRouteTableAggregationRequest) GetTransitRouteTableId() *string {
	return s.TransitRouteTableId
}

func (s *ModifyTransitRouteTableAggregationRequest) SetClientToken(v string) *ModifyTransitRouteTableAggregationRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationRequest) SetDryRun(v bool) *ModifyTransitRouteTableAggregationRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationRequest) SetOwnerAccount(v string) *ModifyTransitRouteTableAggregationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationRequest) SetOwnerId(v int64) *ModifyTransitRouteTableAggregationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationRequest) SetResourceOwnerAccount(v string) *ModifyTransitRouteTableAggregationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationRequest) SetResourceOwnerId(v int64) *ModifyTransitRouteTableAggregationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationRequest) SetTransitRouteTableAggregationCidr(v string) *ModifyTransitRouteTableAggregationRequest {
	s.TransitRouteTableAggregationCidr = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationRequest) SetTransitRouteTableAggregationDescription(v string) *ModifyTransitRouteTableAggregationRequest {
	s.TransitRouteTableAggregationDescription = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationRequest) SetTransitRouteTableAggregationName(v string) *ModifyTransitRouteTableAggregationRequest {
	s.TransitRouteTableAggregationName = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationRequest) SetTransitRouteTableAggregationScope(v string) *ModifyTransitRouteTableAggregationRequest {
	s.TransitRouteTableAggregationScope = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationRequest) SetTransitRouteTableAggregationScopeList(v []*string) *ModifyTransitRouteTableAggregationRequest {
	s.TransitRouteTableAggregationScopeList = v
	return s
}

func (s *ModifyTransitRouteTableAggregationRequest) SetTransitRouteTableId(v string) *ModifyTransitRouteTableAggregationRequest {
	s.TransitRouteTableId = &v
	return s
}

func (s *ModifyTransitRouteTableAggregationRequest) Validate() error {
	return dara.Validate(s)
}
