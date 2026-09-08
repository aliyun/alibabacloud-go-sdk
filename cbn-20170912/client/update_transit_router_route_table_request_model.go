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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, including permission and instance status verification. Valid values:
	//
	// - **false*	- (default): Sends a normal request. If the request passes the check, the name and description of the route table are modified.
	//
	// - **true**: Sends a check request. Only the validation is performed, and the name and description of the route table are not modified. The check items include whether required parameters are specified and the request format. If the check fails, the corresponding error is returned. If the check passes, the error code `DryRunOperation` is returned.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The route table feature options.
	RouteTableOptions *UpdateTransitRouterRouteTableRequestRouteTableOptions `json:"RouteTableOptions,omitempty" xml:"RouteTableOptions,omitempty" type:"Struct"`
	// The description of the route table.
	//
	// The description can be empty or 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// testdesc
	TransitRouterRouteTableDescription *string `json:"TransitRouterRouteTableDescription,omitempty" xml:"TransitRouterRouteTableDescription,omitempty"`
	// The ID of the Enterprise Edition transit router route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp1dudbh2d5na6b50****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
	// The name of the route table.
	//
	// The name can be empty or 1 to 128 characters in length and cannot start with http:// or https://.
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
	if s.RouteTableOptions != nil {
		if err := s.RouteTableOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTransitRouterRouteTableRequestRouteTableOptions struct {
	// Multi-region equal-cost multi-path (ECMP) routing. Valid values:
	//
	// - **disable**: Disables multi-region ECMP routing. After multi-region ECMP routing is disabled, routes with the same prefix learned from different regions will prefer the transit router (TR) with the smallest Region ID (sorted alphabetically) as the next hop when other route attributes are the same. In this case, the traffic latency and bandwidth consumed between different regions may change. Make sure that you have fully evaluated the impact before disabling this feature.
	//
	// - **enable**: Enables multi-region ECMP routing. After multi-region ECMP routing is enabled, routes with the same prefix learned from different regions will form equal-cost routes when other route attributes are the same. In this case, the traffic latency and bandwidth consumed between different regions may change. Make sure that you have fully evaluated the impact before enabling this feature.
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
