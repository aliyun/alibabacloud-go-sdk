// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayRoutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ViewSmartAccessGatewayRoutesResponseBody
	GetRequestId() *string
	SetRoutes(v []*ViewSmartAccessGatewayRoutesResponseBodyRoutes) *ViewSmartAccessGatewayRoutesResponseBody
	GetRoutes() []*ViewSmartAccessGatewayRoutesResponseBodyRoutes
}

type ViewSmartAccessGatewayRoutesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F3D21D0B-5258-5412-AD1C-3929D297286B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the queried routes.
	Routes []*ViewSmartAccessGatewayRoutesResponseBodyRoutes `json:"Routes,omitempty" xml:"Routes,omitempty" type:"Repeated"`
}

func (s ViewSmartAccessGatewayRoutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayRoutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ViewSmartAccessGatewayRoutesResponseBody) GetRoutes() []*ViewSmartAccessGatewayRoutesResponseBodyRoutes {
	return s.Routes
}

func (s *ViewSmartAccessGatewayRoutesResponseBody) SetRequestId(v string) *ViewSmartAccessGatewayRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ViewSmartAccessGatewayRoutesResponseBody) SetRoutes(v []*ViewSmartAccessGatewayRoutesResponseBodyRoutes) *ViewSmartAccessGatewayRoutesResponseBody {
	s.Routes = v
	return s
}

func (s *ViewSmartAccessGatewayRoutesResponseBody) Validate() error {
	if s.Routes != nil {
		for _, item := range s.Routes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ViewSmartAccessGatewayRoutesResponseBodyRoutes struct {
	// The CIDR blocks that overlap with each other.
	ConflictCidrs []*string `json:"ConflictCidrs,omitempty" xml:"ConflictCidrs,omitempty" type:"Repeated"`
	// The route cost.
	//
	// The first **0*	- represents the administrative distance (AD).
	//
	// The second **0*	- represents the router metric.
	//
	// example:
	//
	// [0/0]
	Cost *string `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// The destination CIDR block.
	//
	// example:
	//
	// 172.1.1.0/24
	Dst *string `json:"Dst,omitempty" xml:"Dst,omitempty"`
	// The port number. A value of -1 indicates that the next hop points to a VPN tunnel.
	//
	// Valid values: **-1*	- to **4294967295**.
	//
	// example:
	//
	// 2
	Idx *int32 `json:"Idx,omitempty" xml:"Idx,omitempty"`
	// The next hop.
	//
	// example:
	//
	// 1.XX.XX.1
	Nexthop *string `json:"Nexthop,omitempty" xml:"Nexthop,omitempty"`
	// The port role.
	//
	// example:
	//
	// WAN
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The route type. Valid values:
	//
	// 	- **static**
	//
	// 	- **bgp**
	//
	// 	- **ospf**
	//
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ViewSmartAccessGatewayRoutesResponseBodyRoutes) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayRoutesResponseBodyRoutes) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayRoutesResponseBodyRoutes) GetConflictCidrs() []*string {
	return s.ConflictCidrs
}

func (s *ViewSmartAccessGatewayRoutesResponseBodyRoutes) GetCost() *string {
	return s.Cost
}

func (s *ViewSmartAccessGatewayRoutesResponseBodyRoutes) GetDst() *string {
	return s.Dst
}

func (s *ViewSmartAccessGatewayRoutesResponseBodyRoutes) GetIdx() *int32 {
	return s.Idx
}

func (s *ViewSmartAccessGatewayRoutesResponseBodyRoutes) GetNexthop() *string {
	return s.Nexthop
}

func (s *ViewSmartAccessGatewayRoutesResponseBodyRoutes) GetRole() *string {
	return s.Role
}

func (s *ViewSmartAccessGatewayRoutesResponseBodyRoutes) GetType() *string {
	return s.Type
}

func (s *ViewSmartAccessGatewayRoutesResponseBodyRoutes) SetConflictCidrs(v []*string) *ViewSmartAccessGatewayRoutesResponseBodyRoutes {
	s.ConflictCidrs = v
	return s
}

func (s *ViewSmartAccessGatewayRoutesResponseBodyRoutes) SetCost(v string) *ViewSmartAccessGatewayRoutesResponseBodyRoutes {
	s.Cost = &v
	return s
}

func (s *ViewSmartAccessGatewayRoutesResponseBodyRoutes) SetDst(v string) *ViewSmartAccessGatewayRoutesResponseBodyRoutes {
	s.Dst = &v
	return s
}

func (s *ViewSmartAccessGatewayRoutesResponseBodyRoutes) SetIdx(v int32) *ViewSmartAccessGatewayRoutesResponseBodyRoutes {
	s.Idx = &v
	return s
}

func (s *ViewSmartAccessGatewayRoutesResponseBodyRoutes) SetNexthop(v string) *ViewSmartAccessGatewayRoutesResponseBodyRoutes {
	s.Nexthop = &v
	return s
}

func (s *ViewSmartAccessGatewayRoutesResponseBodyRoutes) SetRole(v string) *ViewSmartAccessGatewayRoutesResponseBodyRoutes {
	s.Role = &v
	return s
}

func (s *ViewSmartAccessGatewayRoutesResponseBodyRoutes) SetType(v string) *ViewSmartAccessGatewayRoutesResponseBodyRoutes {
	s.Type = &v
	return s
}

func (s *ViewSmartAccessGatewayRoutesResponseBodyRoutes) Validate() error {
	return dara.Validate(s)
}
