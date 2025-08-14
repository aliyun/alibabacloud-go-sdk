// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTransitRouterRouteEntryRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTransitRouterRouteEntryRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteTransitRouterRouteEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTransitRouterRouteEntryRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTransitRouterRouteEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTransitRouterRouteEntryRequest
	GetResourceOwnerId() *int64
	SetTransitRouterRouteEntryDestinationCidrBlock(v string) *DeleteTransitRouterRouteEntryRequest
	GetTransitRouterRouteEntryDestinationCidrBlock() *string
	SetTransitRouterRouteEntryId(v string) *DeleteTransitRouterRouteEntryRequest
	GetTransitRouterRouteEntryId() *string
	SetTransitRouterRouteEntryNextHopId(v string) *DeleteTransitRouterRouteEntryRequest
	GetTransitRouterRouteEntryNextHopId() *string
	SetTransitRouterRouteEntryNextHopType(v string) *DeleteTransitRouterRouteEntryRequest
	GetTransitRouterRouteEntryNextHopType() *string
	SetTransitRouterRouteTableId(v string) *DeleteTransitRouterRouteEntryRequest
	GetTransitRouterRouteTableId() *string
}

type DeleteTransitRouterRouteEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the value of **RequestId*	- as the value of **ClientToken**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request but not perform the operation. The system checks the permissions and the status of the specified instances. Valid values:
	//
	// 	- **false*	- (default): sends the request. If the request passes the precheck, the route is deleted.
	//
	// 	- **true**: sends a precheck request. The route is not deleted after the request passes the precheck. If you use this value, the system checks the required parameters and the request syntax. If the check fails, the corresponding error message is returned. If the request passes the check, the system returns the ID of the request.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The destination CIDR block.
	//
	// example:
	//
	// 192.168.0.0/24
	TransitRouterRouteEntryDestinationCidrBlock *string `json:"TransitRouterRouteEntryDestinationCidrBlock,omitempty" xml:"TransitRouterRouteEntryDestinationCidrBlock,omitempty"`
	// The ID of the route entry.
	//
	// example:
	//
	// rte-75eg4jprkvk0pw****
	TransitRouterRouteEntryId *string `json:"TransitRouterRouteEntryId,omitempty" xml:"TransitRouterRouteEntryId,omitempty"`
	// The ID of the network instance connection that you want to specify as the next hop.
	//
	// example:
	//
	// tr-attach-nls9fzkfat8934****
	TransitRouterRouteEntryNextHopId *string `json:"TransitRouterRouteEntryNextHopId,omitempty" xml:"TransitRouterRouteEntryNextHopId,omitempty"`
	// The type of the next hop. Valid values:
	//
	// 	- **BlackHole**: a blackhole route. You do not need to specify a next hop.
	//
	// 	- **Attachment**: a network instance connection. You must specify a network instance connection as the next hop.
	//
	// example:
	//
	// BlackHole
	TransitRouterRouteEntryNextHopType *string `json:"TransitRouterRouteEntryNextHopType,omitempty" xml:"TransitRouterRouteEntryNextHopType,omitempty"`
	// The ID of the route table of the Enterprise Edition transit router.
	//
	// example:
	//
	// vtb-bp1dudbh2d5na6b50****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s DeleteTransitRouterRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterRouteEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTransitRouterRouteEntryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTransitRouterRouteEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTransitRouterRouteEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTransitRouterRouteEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTransitRouterRouteEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTransitRouterRouteEntryRequest) GetTransitRouterRouteEntryDestinationCidrBlock() *string {
	return s.TransitRouterRouteEntryDestinationCidrBlock
}

func (s *DeleteTransitRouterRouteEntryRequest) GetTransitRouterRouteEntryId() *string {
	return s.TransitRouterRouteEntryId
}

func (s *DeleteTransitRouterRouteEntryRequest) GetTransitRouterRouteEntryNextHopId() *string {
	return s.TransitRouterRouteEntryNextHopId
}

func (s *DeleteTransitRouterRouteEntryRequest) GetTransitRouterRouteEntryNextHopType() *string {
	return s.TransitRouterRouteEntryNextHopType
}

func (s *DeleteTransitRouterRouteEntryRequest) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *DeleteTransitRouterRouteEntryRequest) SetClientToken(v string) *DeleteTransitRouterRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetDryRun(v bool) *DeleteTransitRouterRouteEntryRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetOwnerAccount(v string) *DeleteTransitRouterRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetOwnerId(v int64) *DeleteTransitRouterRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryDestinationCidrBlock(v string) *DeleteTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryDestinationCidrBlock = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryId(v string) *DeleteTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryId = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryNextHopId(v string) *DeleteTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryNextHopId = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryNextHopType(v string) *DeleteTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryNextHopType = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) SetTransitRouterRouteTableId(v string) *DeleteTransitRouterRouteEntryRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *DeleteTransitRouterRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
