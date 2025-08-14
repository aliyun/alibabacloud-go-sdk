// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterCidrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidr(v string) *CreateTransitRouterCidrRequest
	GetCidr() *string
	SetClientToken(v string) *CreateTransitRouterCidrRequest
	GetClientToken() *string
	SetDescription(v string) *CreateTransitRouterCidrRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateTransitRouterCidrRequest
	GetDryRun() *bool
	SetName(v string) *CreateTransitRouterCidrRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateTransitRouterCidrRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTransitRouterCidrRequest
	GetOwnerId() *int64
	SetPublishCidrRoute(v bool) *CreateTransitRouterCidrRequest
	GetPublishCidrRoute() *bool
	SetRegionId(v string) *CreateTransitRouterCidrRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateTransitRouterCidrRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTransitRouterCidrRequest
	GetResourceOwnerId() *int64
	SetTransitRouterId(v string) *CreateTransitRouterCidrRequest
	GetTransitRouterId() *string
}

type CreateTransitRouterCidrRequest struct {
	// The CIDR block of the transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.10.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the transit router CIDR block.
	//
	// The description must be 1 to 256 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the transit router CIDR block.
	//
	// The name must be 1 to 128 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// nametest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to allow the system to automatically add a route that points to the CIDR block to the route table of the transit router.
	//
	// 	- **true*	- (default)
	//
	//     If you set the value to true, after you create a VPN attachment on a private VPN gateway and enable route learning for the VPN attachment, the system automatically adds the following route to the route table of the transit router that is in route learning relationship with the VPN attachment:
	//
	//     A blackhole route whose destination CIDR block is the transit router CIDR block, which refers to the CIDR block from which gateway IP addresses are allocated to the IPsec-VPN connection. The blackhole route is advertised only to the route tables of virtual border routers (VBRs) connected to the transit router.
	//
	// 	- **false**
	//
	// example:
	//
	// true
	PublishCidrRoute *bool `json:"PublishCidrRoute,omitempty" xml:"PublishCidrRoute,omitempty"`
	// The ID of the region to which the transfer router belongs.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-p0w3x8c9em72a40nw****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s CreateTransitRouterCidrRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterCidrRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterCidrRequest) GetCidr() *string {
	return s.Cidr
}

func (s *CreateTransitRouterCidrRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouterCidrRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTransitRouterCidrRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouterCidrRequest) GetName() *string {
	return s.Name
}

func (s *CreateTransitRouterCidrRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTransitRouterCidrRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTransitRouterCidrRequest) GetPublishCidrRoute() *bool {
	return s.PublishCidrRoute
}

func (s *CreateTransitRouterCidrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTransitRouterCidrRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTransitRouterCidrRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTransitRouterCidrRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateTransitRouterCidrRequest) SetCidr(v string) *CreateTransitRouterCidrRequest {
	s.Cidr = &v
	return s
}

func (s *CreateTransitRouterCidrRequest) SetClientToken(v string) *CreateTransitRouterCidrRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterCidrRequest) SetDescription(v string) *CreateTransitRouterCidrRequest {
	s.Description = &v
	return s
}

func (s *CreateTransitRouterCidrRequest) SetDryRun(v bool) *CreateTransitRouterCidrRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterCidrRequest) SetName(v string) *CreateTransitRouterCidrRequest {
	s.Name = &v
	return s
}

func (s *CreateTransitRouterCidrRequest) SetOwnerAccount(v string) *CreateTransitRouterCidrRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterCidrRequest) SetOwnerId(v int64) *CreateTransitRouterCidrRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterCidrRequest) SetPublishCidrRoute(v bool) *CreateTransitRouterCidrRequest {
	s.PublishCidrRoute = &v
	return s
}

func (s *CreateTransitRouterCidrRequest) SetRegionId(v string) *CreateTransitRouterCidrRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransitRouterCidrRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterCidrRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterCidrRequest) SetResourceOwnerId(v int64) *CreateTransitRouterCidrRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterCidrRequest) SetTransitRouterId(v string) *CreateTransitRouterCidrRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTransitRouterCidrRequest) Validate() error {
	return dara.Validate(s)
}
