// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnsRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateEnsRouteEntryRequest
	GetDescription() *string
	SetDestinationCidrBlock(v string) *CreateEnsRouteEntryRequest
	GetDestinationCidrBlock() *string
	SetNextHopId(v string) *CreateEnsRouteEntryRequest
	GetNextHopId() *string
	SetNextHopType(v string) *CreateEnsRouteEntryRequest
	GetNextHopType() *string
	SetRouteEntryName(v string) *CreateEnsRouteEntryRequest
	GetRouteEntryName() *string
	SetRouteTableId(v string) *CreateEnsRouteEntryRequest
	GetRouteTableId() *string
	SetSourceCidrBlock(v string) *CreateEnsRouteEntryRequest
	GetSourceCidrBlock() *string
}

type CreateEnsRouteEntryRequest struct {
	// The description of the custom route entry.
	//
	// example:
	//
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block of the custom route entry. Make sure that the following requirements are met:
	//
	// 	- The destination CIDR block cannot point or belong to 100.64.0.0/10.
	//
	// 	- The destination CIDR blocks of the custom route entries in the same route table cannot overlap.
	//
	// 	- 0.0.0.0/0 indicates the default CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the next hop of the custom route entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-5vb7leks9z4mxy1ay258
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The next hop type of the custom route. Valid values:
	//
	// 	- Instance: an ENS instance.
	//
	// 	- HaVip: a high-availability virtual IP address (HAVIP).
	//
	// 	- NetworkPeer: VPC peering connection.
	//
	// example:
	//
	// Instance
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The name of the custom route entry that you want to add. The name must be 1 to 128 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// test
	RouteEntryName *string `json:"RouteEntryName,omitempty" xml:"RouteEntryName,omitempty"`
	// The ID of the route table to which you want to add a custom route entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp1cifr72dioje82lse2j
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The source CIDR block is available when you configure a route entry in the gateway route table, but is not unavailable when you configure a route entry in the vSwitch route table.
	//
	// example:
	//
	// 172.XXX.XXX.0/24
	SourceCidrBlock *string `json:"SourceCidrBlock,omitempty" xml:"SourceCidrBlock,omitempty"`
}

func (s CreateEnsRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnsRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateEnsRouteEntryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEnsRouteEntryRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *CreateEnsRouteEntryRequest) GetNextHopId() *string {
	return s.NextHopId
}

func (s *CreateEnsRouteEntryRequest) GetNextHopType() *string {
	return s.NextHopType
}

func (s *CreateEnsRouteEntryRequest) GetRouteEntryName() *string {
	return s.RouteEntryName
}

func (s *CreateEnsRouteEntryRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *CreateEnsRouteEntryRequest) GetSourceCidrBlock() *string {
	return s.SourceCidrBlock
}

func (s *CreateEnsRouteEntryRequest) SetDescription(v string) *CreateEnsRouteEntryRequest {
	s.Description = &v
	return s
}

func (s *CreateEnsRouteEntryRequest) SetDestinationCidrBlock(v string) *CreateEnsRouteEntryRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateEnsRouteEntryRequest) SetNextHopId(v string) *CreateEnsRouteEntryRequest {
	s.NextHopId = &v
	return s
}

func (s *CreateEnsRouteEntryRequest) SetNextHopType(v string) *CreateEnsRouteEntryRequest {
	s.NextHopType = &v
	return s
}

func (s *CreateEnsRouteEntryRequest) SetRouteEntryName(v string) *CreateEnsRouteEntryRequest {
	s.RouteEntryName = &v
	return s
}

func (s *CreateEnsRouteEntryRequest) SetRouteTableId(v string) *CreateEnsRouteEntryRequest {
	s.RouteTableId = &v
	return s
}

func (s *CreateEnsRouteEntryRequest) SetSourceCidrBlock(v string) *CreateEnsRouteEntryRequest {
	s.SourceCidrBlock = &v
	return s
}

func (s *CreateEnsRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
