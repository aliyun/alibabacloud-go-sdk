// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagRouteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSagRouteListResponseBody
	GetRequestId() *string
	SetRoutes(v []*DescribeSagRouteListResponseBodyRoutes) *DescribeSagRouteListResponseBody
	GetRoutes() []*DescribeSagRouteListResponseBodyRoutes
}

type DescribeSagRouteListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CE6642D4-21EB-4168-9BF9-F217953F9892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of routes.
	Routes []*DescribeSagRouteListResponseBodyRoutes `json:"Routes,omitempty" xml:"Routes,omitempty" type:"Repeated"`
}

func (s DescribeSagRouteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRouteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagRouteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagRouteListResponseBody) GetRoutes() []*DescribeSagRouteListResponseBodyRoutes {
	return s.Routes
}

func (s *DescribeSagRouteListResponseBody) SetRequestId(v string) *DescribeSagRouteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagRouteListResponseBody) SetRoutes(v []*DescribeSagRouteListResponseBodyRoutes) *DescribeSagRouteListResponseBody {
	s.Routes = v
	return s
}

func (s *DescribeSagRouteListResponseBody) Validate() error {
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

type DescribeSagRouteListResponseBodyRoutes struct {
	// The list of conflicting CIDR blocks.
	ConflictCidrs []*string `json:"ConflictCidrs,omitempty" xml:"ConflictCidrs,omitempty" type:"Repeated"`
	// The routing cost value. The number to the left of the forward slash (/) indicates the administrative distance (AD) of the routing protocol. The number to the right of the forward slash (/) indicates the metric.
	//
	// example:
	//
	// [110/11]
	Cost *string `json:"Cost,omitempty" xml:"Cost,omitempty"`
	// The destination CIDR block.
	//
	// example:
	//
	// 6.XX.XX.6/32
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The next hop.
	//
	// example:
	//
	// 192.XX.XX.1
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The port name. A value of -1 indicates that data traffic enters Alibaba Cloud through a VPN connection.
	//
	// example:
	//
	// 2
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The routing protocol type. Valid values:
	//
	// - **STATIC**: Static routing protocol.
	//
	// - **OSPF**: OSPF dynamic routing protocol.
	//
	// - **BGP**: BGP dynamic routing protocol.
	//
	// example:
	//
	// STATIC
	RouteProtocol *string `json:"RouteProtocol,omitempty" xml:"RouteProtocol,omitempty"`
}

func (s DescribeSagRouteListResponseBodyRoutes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRouteListResponseBodyRoutes) GoString() string {
	return s.String()
}

func (s *DescribeSagRouteListResponseBodyRoutes) GetConflictCidrs() []*string {
	return s.ConflictCidrs
}

func (s *DescribeSagRouteListResponseBodyRoutes) GetCost() *string {
	return s.Cost
}

func (s *DescribeSagRouteListResponseBodyRoutes) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *DescribeSagRouteListResponseBodyRoutes) GetNextHop() *string {
	return s.NextHop
}

func (s *DescribeSagRouteListResponseBodyRoutes) GetPortName() *string {
	return s.PortName
}

func (s *DescribeSagRouteListResponseBodyRoutes) GetRouteProtocol() *string {
	return s.RouteProtocol
}

func (s *DescribeSagRouteListResponseBodyRoutes) SetConflictCidrs(v []*string) *DescribeSagRouteListResponseBodyRoutes {
	s.ConflictCidrs = v
	return s
}

func (s *DescribeSagRouteListResponseBodyRoutes) SetCost(v string) *DescribeSagRouteListResponseBodyRoutes {
	s.Cost = &v
	return s
}

func (s *DescribeSagRouteListResponseBodyRoutes) SetDestinationCidr(v string) *DescribeSagRouteListResponseBodyRoutes {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeSagRouteListResponseBodyRoutes) SetNextHop(v string) *DescribeSagRouteListResponseBodyRoutes {
	s.NextHop = &v
	return s
}

func (s *DescribeSagRouteListResponseBodyRoutes) SetPortName(v string) *DescribeSagRouteListResponseBodyRoutes {
	s.PortName = &v
	return s
}

func (s *DescribeSagRouteListResponseBodyRoutes) SetRouteProtocol(v string) *DescribeSagRouteListResponseBodyRoutes {
	s.RouteProtocol = &v
	return s
}

func (s *DescribeSagRouteListResponseBodyRoutes) Validate() error {
	return dara.Validate(s)
}
