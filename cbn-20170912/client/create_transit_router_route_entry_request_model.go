// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTransitRouterRouteEntryRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTransitRouterRouteEntryRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *CreateTransitRouterRouteEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTransitRouterRouteEntryRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateTransitRouterRouteEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTransitRouterRouteEntryRequest
	GetResourceOwnerId() *int64
	SetTransitRouterRouteEntryDescription(v string) *CreateTransitRouterRouteEntryRequest
	GetTransitRouterRouteEntryDescription() *string
	SetTransitRouterRouteEntryDestinationCidrBlock(v string) *CreateTransitRouterRouteEntryRequest
	GetTransitRouterRouteEntryDestinationCidrBlock() *string
	SetTransitRouterRouteEntryName(v string) *CreateTransitRouterRouteEntryRequest
	GetTransitRouterRouteEntryName() *string
	SetTransitRouterRouteEntryNextHopId(v string) *CreateTransitRouterRouteEntryRequest
	GetTransitRouterRouteEntryNextHopId() *string
	SetTransitRouterRouteEntryNextHopType(v string) *CreateTransitRouterRouteEntryRequest
	GetTransitRouterRouteEntryNextHopType() *string
	SetTransitRouterRouteTableId(v string) *CreateTransitRouterRouteEntryRequest
	GetTransitRouterRouteTableId() *string
}

type CreateTransitRouterRouteEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- is different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The description of the route.
	//
	// The description must be 1 to 256 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// testdesc
	TransitRouterRouteEntryDescription *string `json:"TransitRouterRouteEntryDescription,omitempty" xml:"TransitRouterRouteEntryDescription,omitempty"`
	// The destination CIDR block of the route entry. IPv4 and IPv6 addresses are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/24
	TransitRouterRouteEntryDestinationCidrBlock *string `json:"TransitRouterRouteEntryDestinationCidrBlock,omitempty" xml:"TransitRouterRouteEntryDestinationCidrBlock,omitempty"`
	// The name of the route.
	//
	// The name must be 1 to 128 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// testname
	TransitRouterRouteEntryName *string `json:"TransitRouterRouteEntryName,omitempty" xml:"TransitRouterRouteEntryName,omitempty"`
	// The ID of the network instance connection that you want to specify as the next hop.
	//
	// example:
	//
	// tr-attach-nls9fzkfat8934****
	TransitRouterRouteEntryNextHopId *string `json:"TransitRouterRouteEntryNextHopId,omitempty" xml:"TransitRouterRouteEntryNextHopId,omitempty"`
	// The type of the next hop. Valid values:
	//
	// 	- **BlackHole**: routes network traffic to a black hole. All packets that match this route are dropped. If you select this option, you do not need to specify the next hop information.
	//
	// 	- **Attachment**: routes network traffic to a network instance connection. If you select this option, you must specify the ID of the network instance connection. All packets that match this route are routed to the specified network instance connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// BlackHole
	TransitRouterRouteEntryNextHopType *string `json:"TransitRouterRouteEntryNextHopType,omitempty" xml:"TransitRouterRouteEntryNextHopType,omitempty"`
	// The ID of the route table of the Enterprise Edition transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp1dudbh2d5na6b50****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s CreateTransitRouterRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRouteEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouterRouteEntryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouterRouteEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTransitRouterRouteEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTransitRouterRouteEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTransitRouterRouteEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTransitRouterRouteEntryRequest) GetTransitRouterRouteEntryDescription() *string {
	return s.TransitRouterRouteEntryDescription
}

func (s *CreateTransitRouterRouteEntryRequest) GetTransitRouterRouteEntryDestinationCidrBlock() *string {
	return s.TransitRouterRouteEntryDestinationCidrBlock
}

func (s *CreateTransitRouterRouteEntryRequest) GetTransitRouterRouteEntryName() *string {
	return s.TransitRouterRouteEntryName
}

func (s *CreateTransitRouterRouteEntryRequest) GetTransitRouterRouteEntryNextHopId() *string {
	return s.TransitRouterRouteEntryNextHopId
}

func (s *CreateTransitRouterRouteEntryRequest) GetTransitRouterRouteEntryNextHopType() *string {
	return s.TransitRouterRouteEntryNextHopType
}

func (s *CreateTransitRouterRouteEntryRequest) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *CreateTransitRouterRouteEntryRequest) SetClientToken(v string) *CreateTransitRouterRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetDryRun(v bool) *CreateTransitRouterRouteEntryRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetOwnerAccount(v string) *CreateTransitRouterRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetOwnerId(v int64) *CreateTransitRouterRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetResourceOwnerId(v int64) *CreateTransitRouterRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryDescription(v string) *CreateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryDescription = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryDestinationCidrBlock(v string) *CreateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryDestinationCidrBlock = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryName(v string) *CreateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryName = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryNextHopId(v string) *CreateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryNextHopId = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryNextHopType(v string) *CreateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryNextHopType = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) SetTransitRouterRouteTableId(v string) *CreateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *CreateTransitRouterRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
