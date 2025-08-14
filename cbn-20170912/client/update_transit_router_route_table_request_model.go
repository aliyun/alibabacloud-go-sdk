// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterRouteTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateTransitRouterRouteTableRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateTransitRouterRouteTableRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *UpdateTransitRouterRouteTableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateTransitRouterRouteTableRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateTransitRouterRouteTableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTransitRouterRouteTableRequest
	GetResourceOwnerId() *int64
	SetRouteTableOptions(v *UpdateTransitRouterRouteTableRequestRouteTableOptions) *UpdateTransitRouterRouteTableRequest
	GetRouteTableOptions() *UpdateTransitRouterRouteTableRequestRouteTableOptions
	SetTransitRouterRouteTableDescription(v string) *UpdateTransitRouterRouteTableRequest
	GetTransitRouterRouteTableDescription() *string
	SetTransitRouterRouteTableId(v string) *UpdateTransitRouterRouteTableRequest
	GetTransitRouterRouteTableId() *string
	SetTransitRouterRouteTableName(v string) *UpdateTransitRouterRouteTableRequest
	GetTransitRouterRouteTableName() *string
}

type UpdateTransitRouterRouteTableRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Default values:
	//
	// 	- **false*	- (default): performs a dry run and sends the request.
	//
	// 	- **true**: performs a dry run. The system checks the required parameters and the request syntax. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The features of the route table.
	RouteTableOptions *UpdateTransitRouterRouteTableRequestRouteTableOptions `json:"RouteTableOptions,omitempty" xml:"RouteTableOptions,omitempty" type:"Struct"`
	// The description of the route table.
	//
	// The description must be 1 to 256 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// testdesc
	TransitRouterRouteTableDescription *string `json:"TransitRouterRouteTableDescription,omitempty" xml:"TransitRouterRouteTableDescription,omitempty"`
	// The ID of the route table of the Enterprise Edition transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp1dudbh2d5na6b50****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
	// The name of the route table.
	//
	// The name must be 1 to 128 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// testname
	TransitRouterRouteTableName *string `json:"TransitRouterRouteTableName,omitempty" xml:"TransitRouterRouteTableName,omitempty"`
}

func (s UpdateTransitRouterRouteTableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterRouteTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterRouteTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTransitRouterRouteTableRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateTransitRouterRouteTableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateTransitRouterRouteTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTransitRouterRouteTableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTransitRouterRouteTableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTransitRouterRouteTableRequest) GetRouteTableOptions() *UpdateTransitRouterRouteTableRequestRouteTableOptions {
	return s.RouteTableOptions
}

func (s *UpdateTransitRouterRouteTableRequest) GetTransitRouterRouteTableDescription() *string {
	return s.TransitRouterRouteTableDescription
}

func (s *UpdateTransitRouterRouteTableRequest) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *UpdateTransitRouterRouteTableRequest) GetTransitRouterRouteTableName() *string {
	return s.TransitRouterRouteTableName
}

func (s *UpdateTransitRouterRouteTableRequest) SetClientToken(v string) *UpdateTransitRouterRouteTableRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetDryRun(v bool) *UpdateTransitRouterRouteTableRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetOwnerAccount(v string) *UpdateTransitRouterRouteTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetOwnerId(v int64) *UpdateTransitRouterRouteTableRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetResourceOwnerAccount(v string) *UpdateTransitRouterRouteTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetResourceOwnerId(v int64) *UpdateTransitRouterRouteTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetRouteTableOptions(v *UpdateTransitRouterRouteTableRequestRouteTableOptions) *UpdateTransitRouterRouteTableRequest {
	s.RouteTableOptions = v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetTransitRouterRouteTableDescription(v string) *UpdateTransitRouterRouteTableRequest {
	s.TransitRouterRouteTableDescription = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetTransitRouterRouteTableId(v string) *UpdateTransitRouterRouteTableRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) SetTransitRouterRouteTableName(v string) *UpdateTransitRouterRouteTableRequest {
	s.TransitRouterRouteTableName = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateTransitRouterRouteTableRequestRouteTableOptions struct {
	// Indicates whether multi-region ECMP routing is enabled. Valid values:
	//
	// - **disable**: If multi-region ECMP routing is disabled, routes that are learned from different regions but have the same prefix and attributes select the transit router with the smallest region ID as the next hop. Region IDs are sorted in alphabetic order. The network latency and bandwidth consumption also vary based on the region. Proceed with caution.
	//
	// - **enable**: If multi-region ECMP routing is enabled, routes that are learned from different regions but have the same prefix and attributes form an ECMP route. The network latency and bandwidth consumption also vary based on the region. Proceed with caution.
	//
	// example:
	//
	// disable
	MultiRegionECMP *string `json:"MultiRegionECMP,omitempty" xml:"MultiRegionECMP,omitempty"`
}

func (s UpdateTransitRouterRouteTableRequestRouteTableOptions) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterRouteTableRequestRouteTableOptions) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterRouteTableRequestRouteTableOptions) GetMultiRegionECMP() *string {
	return s.MultiRegionECMP
}

func (s *UpdateTransitRouterRouteTableRequestRouteTableOptions) SetMultiRegionECMP(v string) *UpdateTransitRouterRouteTableRequestRouteTableOptions {
	s.MultiRegionECMP = &v
	return s
}

func (s *UpdateTransitRouterRouteTableRequestRouteTableOptions) Validate() error {
	return dara.Validate(s)
}
