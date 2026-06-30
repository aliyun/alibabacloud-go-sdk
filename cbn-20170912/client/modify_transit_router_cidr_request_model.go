// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTransitRouterCidrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidr(v string) *ModifyTransitRouterCidrRequest
	GetCidr() *string
	SetClientToken(v string) *ModifyTransitRouterCidrRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyTransitRouterCidrRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyTransitRouterCidrRequest
	GetDryRun() *bool
	SetName(v string) *ModifyTransitRouterCidrRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyTransitRouterCidrRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyTransitRouterCidrRequest
	GetOwnerId() *int64
	SetPublishCidrRoute(v bool) *ModifyTransitRouterCidrRequest
	GetPublishCidrRoute() *bool
	SetRegionId(v string) *ModifyTransitRouterCidrRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyTransitRouterCidrRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyTransitRouterCidrRequest
	GetResourceOwnerId() *int64
	SetTransitRouterCidrId(v string) *ModifyTransitRouterCidrRequest
	GetTransitRouterCidrId() *string
	SetTransitRouterId(v string) *ModifyTransitRouterCidrRequest
	GetTransitRouterId() *string
}

type ModifyTransitRouterCidrRequest struct {
	// The new CIDR block.
	//
	// example:
	//
	// 192.168.10.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a client token to make sure that the token is unique for each request. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID is different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new description of the CIDR block.
	//
	// The description can be empty or 1 to 256 characters in length. It cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run. The valid values are:
	//
	// - **true**: Sends a check request but does not modify the CIDR block. The system checks the required parameters, request format, and service limits. If the request fails the check, the corresponding error is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): Sends a normal request. The CIDR block is modified after the request passes the check.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The new name of the CIDR block.
	//
	// The name can be empty or 1 to 128 characters in length. It cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// nametest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to automatically add a route that points to the CIDR block to the route table of the transit router.
	//
	// - **true**: Yes.
	//
	//   This blackhole route is advertised only to the route tables of virtual border routers (VBRs) that are attached to the transit router.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	PublishCidrRoute *bool `json:"PublishCidrRoute,omitempty" xml:"PublishCidrRoute,omitempty"`
	// The ID of the region where the Transit Router instance is deployed.
	//
	// Call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the CIDR block.
	//
	// Call the [ListTransitRouterCidr](https://help.aliyun.com/document_detail/462772.html) operation to query the ID of the CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// cidr-0zv0q9crqpntzz****
	TransitRouterCidrId *string `json:"TransitRouterCidrId,omitempty" xml:"TransitRouterCidrId,omitempty"`
	// The ID of the Transit Router instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-gw8ergozrv77rtbjd****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ModifyTransitRouterCidrRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouterCidrRequest) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouterCidrRequest) GetCidr() *string {
	return s.Cidr
}

func (s *ModifyTransitRouterCidrRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyTransitRouterCidrRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyTransitRouterCidrRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyTransitRouterCidrRequest) GetName() *string {
	return s.Name
}

func (s *ModifyTransitRouterCidrRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyTransitRouterCidrRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyTransitRouterCidrRequest) GetPublishCidrRoute() *bool {
	return s.PublishCidrRoute
}

func (s *ModifyTransitRouterCidrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyTransitRouterCidrRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyTransitRouterCidrRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyTransitRouterCidrRequest) GetTransitRouterCidrId() *string {
	return s.TransitRouterCidrId
}

func (s *ModifyTransitRouterCidrRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ModifyTransitRouterCidrRequest) SetCidr(v string) *ModifyTransitRouterCidrRequest {
	s.Cidr = &v
	return s
}

func (s *ModifyTransitRouterCidrRequest) SetClientToken(v string) *ModifyTransitRouterCidrRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyTransitRouterCidrRequest) SetDescription(v string) *ModifyTransitRouterCidrRequest {
	s.Description = &v
	return s
}

func (s *ModifyTransitRouterCidrRequest) SetDryRun(v bool) *ModifyTransitRouterCidrRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyTransitRouterCidrRequest) SetName(v string) *ModifyTransitRouterCidrRequest {
	s.Name = &v
	return s
}

func (s *ModifyTransitRouterCidrRequest) SetOwnerAccount(v string) *ModifyTransitRouterCidrRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyTransitRouterCidrRequest) SetOwnerId(v int64) *ModifyTransitRouterCidrRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyTransitRouterCidrRequest) SetPublishCidrRoute(v bool) *ModifyTransitRouterCidrRequest {
	s.PublishCidrRoute = &v
	return s
}

func (s *ModifyTransitRouterCidrRequest) SetRegionId(v string) *ModifyTransitRouterCidrRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyTransitRouterCidrRequest) SetResourceOwnerAccount(v string) *ModifyTransitRouterCidrRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyTransitRouterCidrRequest) SetResourceOwnerId(v int64) *ModifyTransitRouterCidrRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyTransitRouterCidrRequest) SetTransitRouterCidrId(v string) *ModifyTransitRouterCidrRequest {
	s.TransitRouterCidrId = &v
	return s
}

func (s *ModifyTransitRouterCidrRequest) SetTransitRouterId(v string) *ModifyTransitRouterCidrRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ModifyTransitRouterCidrRequest) Validate() error {
	return dara.Validate(s)
}
