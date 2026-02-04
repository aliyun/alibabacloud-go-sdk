// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterCidrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCidrLists(v []*ListTransitRouterCidrResponseBodyCidrLists) *ListTransitRouterCidrResponseBody
	GetCidrLists() []*ListTransitRouterCidrResponseBodyCidrLists
	SetRequestId(v string) *ListTransitRouterCidrResponseBody
	GetRequestId() *string
}

type ListTransitRouterCidrResponseBody struct {
	// The information about the CIDR block.
	CidrLists []*ListTransitRouterCidrResponseBodyCidrLists `json:"CidrLists,omitempty" xml:"CidrLists,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0876E54E-3E36-5C31-89F0-9EE8A9266F9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTransitRouterCidrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterCidrResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterCidrResponseBody) GetCidrLists() []*ListTransitRouterCidrResponseBodyCidrLists {
	return s.CidrLists
}

func (s *ListTransitRouterCidrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterCidrResponseBody) SetCidrLists(v []*ListTransitRouterCidrResponseBodyCidrLists) *ListTransitRouterCidrResponseBody {
	s.CidrLists = v
	return s
}

func (s *ListTransitRouterCidrResponseBody) SetRequestId(v string) *ListTransitRouterCidrResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterCidrResponseBody) Validate() error {
	if s.CidrLists != nil {
		for _, item := range s.CidrLists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransitRouterCidrResponseBodyCidrLists struct {
	// The CIDR block of the transit router.
	//
	// example:
	//
	// 192.168.10.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The description of the CIDR block.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the CIDR block.
	//
	// The value is **IPv4**, which indicates that the CIDR block is of the IPv4 type.
	//
	// example:
	//
	// IPv4
	Family *string `json:"Family,omitempty" xml:"Family,omitempty"`
	// The name of the CIDR block.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the system is allowed to automatically add a route to the route table of the transit router. Valid values:
	//
	// 	- **true**
	//
	//     A value of **true*	- indicates that after you create a private VPN connection and enable route learning for the connection, the system automatically adds a blackhole route to the route table of the transit router to which the VPN connection is attached.
	//
	//     The destination CIDR block of the blackhole route is the CIDR block of the transit router. The CIDR block of the transit router refers to the CIDR block from which gateway IP addresses are allocated to IPsec-VPN connections.
	//
	//     The blackhole route is advertised only to the route table of the virtual border router (VBR) that is connected to the transit router.
	//
	// 	- **false**
	//
	// example:
	//
	// true
	PublishCidrRoute *bool `json:"PublishCidrRoute,omitempty" xml:"PublishCidrRoute,omitempty"`
	// The ID of the transit router CIDR block.
	//
	// example:
	//
	// cidr-0zv0q9crqpntzz****
	TransitRouterCidrId *string `json:"TransitRouterCidrId,omitempty" xml:"TransitRouterCidrId,omitempty"`
	// The transit router ID.
	//
	// example:
	//
	// tr-p0w3x8c9em72a40nw****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTransitRouterCidrResponseBodyCidrLists) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterCidrResponseBodyCidrLists) GoString() string {
	return s.String()
}

func (s *ListTransitRouterCidrResponseBodyCidrLists) GetCidr() *string {
	return s.Cidr
}

func (s *ListTransitRouterCidrResponseBodyCidrLists) GetDescription() *string {
	return s.Description
}

func (s *ListTransitRouterCidrResponseBodyCidrLists) GetFamily() *string {
	return s.Family
}

func (s *ListTransitRouterCidrResponseBodyCidrLists) GetName() *string {
	return s.Name
}

func (s *ListTransitRouterCidrResponseBodyCidrLists) GetPublishCidrRoute() *bool {
	return s.PublishCidrRoute
}

func (s *ListTransitRouterCidrResponseBodyCidrLists) GetTransitRouterCidrId() *string {
	return s.TransitRouterCidrId
}

func (s *ListTransitRouterCidrResponseBodyCidrLists) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterCidrResponseBodyCidrLists) SetCidr(v string) *ListTransitRouterCidrResponseBodyCidrLists {
	s.Cidr = &v
	return s
}

func (s *ListTransitRouterCidrResponseBodyCidrLists) SetDescription(v string) *ListTransitRouterCidrResponseBodyCidrLists {
	s.Description = &v
	return s
}

func (s *ListTransitRouterCidrResponseBodyCidrLists) SetFamily(v string) *ListTransitRouterCidrResponseBodyCidrLists {
	s.Family = &v
	return s
}

func (s *ListTransitRouterCidrResponseBodyCidrLists) SetName(v string) *ListTransitRouterCidrResponseBodyCidrLists {
	s.Name = &v
	return s
}

func (s *ListTransitRouterCidrResponseBodyCidrLists) SetPublishCidrRoute(v bool) *ListTransitRouterCidrResponseBodyCidrLists {
	s.PublishCidrRoute = &v
	return s
}

func (s *ListTransitRouterCidrResponseBodyCidrLists) SetTransitRouterCidrId(v string) *ListTransitRouterCidrResponseBodyCidrLists {
	s.TransitRouterCidrId = &v
	return s
}

func (s *ListTransitRouterCidrResponseBodyCidrLists) SetTransitRouterId(v string) *ListTransitRouterCidrResponseBodyCidrLists {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterCidrResponseBodyCidrLists) Validate() error {
	return dara.Validate(s)
}
