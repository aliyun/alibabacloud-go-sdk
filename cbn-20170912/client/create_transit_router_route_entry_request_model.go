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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, including permission and instance status verification. Valid values:
	//
	// - **false*	- (default): Sends a normal request and creates the route entry after the request passes the verification.
	//
	// - **true**: Sends a check request. Only the verification is performed. No route entry is created. The system checks whether the required parameters are specified and whether the request format is valid. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The description of the route entry.
	//
	// The description can be empty or 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// testdesc
	TransitRouterRouteEntryDescription *string `json:"TransitRouterRouteEntryDescription,omitempty" xml:"TransitRouterRouteEntryDescription,omitempty"`
	// The destination CIDR block of the route entry. Both IPv4 and IPv6 addresses are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/24
	TransitRouterRouteEntryDestinationCidrBlock *string `json:"TransitRouterRouteEntryDestinationCidrBlock,omitempty" xml:"TransitRouterRouteEntryDestinationCidrBlock,omitempty"`
	// The name of the route entry.
	//
	// The name can be empty or 1 to 128 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// testname
	TransitRouterRouteEntryName *string `json:"TransitRouterRouteEntryName,omitempty" xml:"TransitRouterRouteEntryName,omitempty"`
	// The ID of the network instance connection associated with the next hop of the route entry.
	//
	// example:
	//
	// tr-attach-nls9fzkfat8934****
	TransitRouterRouteEntryNextHopId *string `json:"TransitRouterRouteEntryNextHopId,omitempty" xml:"TransitRouterRouteEntryNextHopId,omitempty"`
	// The next hop type. Valid values:
	//
	// - **BlackHole**: specifies the route entry as a blackhole route. All packets destined for the destination CIDR block are dropped. You do not need to specify next hop information.
	//
	// - **Attachment**: specifies the next hop of the route entry as a network instance connection. You must also specify the network instance connection ID. All packets destined for the destination CIDR block are forwarded to the specified network instance connection.
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
