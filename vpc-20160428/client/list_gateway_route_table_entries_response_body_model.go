// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayRouteTableEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayRouteEntryModels(v []*ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) *ListGatewayRouteTableEntriesResponseBody
	GetGatewayRouteEntryModels() []*ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels
	SetNextToken(v string) *ListGatewayRouteTableEntriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListGatewayRouteTableEntriesResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListGatewayRouteTableEntriesResponseBody
	GetTotalCount() *string
}

type ListGatewayRouteTableEntriesResponseBody struct {
	// The details of the routes in the gateway route table.
	GatewayRouteEntryModels []*ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels `json:"GatewayRouteEntryModels,omitempty" xml:"GatewayRouteEntryModels,omitempty" type:"Repeated"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is used to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGatewayRouteTableEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteTableEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteTableEntriesResponseBody) GetGatewayRouteEntryModels() []*ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels {
	return s.GatewayRouteEntryModels
}

func (s *ListGatewayRouteTableEntriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGatewayRouteTableEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayRouteTableEntriesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListGatewayRouteTableEntriesResponseBody) SetGatewayRouteEntryModels(v []*ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) *ListGatewayRouteTableEntriesResponseBody {
	s.GatewayRouteEntryModels = v
	return s
}

func (s *ListGatewayRouteTableEntriesResponseBody) SetNextToken(v string) *ListGatewayRouteTableEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGatewayRouteTableEntriesResponseBody) SetRequestId(v string) *ListGatewayRouteTableEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayRouteTableEntriesResponseBody) SetTotalCount(v string) *ListGatewayRouteTableEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGatewayRouteTableEntriesResponseBody) Validate() error {
	if s.GatewayRouteEntryModels != nil {
		for _, item := range s.GatewayRouteEntryModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels struct {
	// The name of the route entry.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block of the route.
	//
	// example:
	//
	// 192.168.0.5
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The name of the route entry.
	//
	// The name must be 2 to 128 characters in length and can contain letter, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the next hop.
	//
	// example:
	//
	// i-bp11gcl0sm85t9bi****
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The type of the next hop. Valid values:
	//
	// 	- **EcsInstance**: Elastic Compute Service (ECS) instance
	//
	// 	- **NetworkInterface**: elastic network interfaces (ENIs).
	//
	// 	- **Local**: local next hop
	//
	// example:
	//
	// EcsInstance
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The information about the next hop.
	NextHops []*ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops `json:"NextHops,omitempty" xml:"NextHops,omitempty" type:"Repeated"`
	// The status of the route entry. Valid values:
	//
	// 	- **Pending**
	//
	// 	- **Available**
	//
	// 	- **Modifying**
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) GetDescription() *string {
	return s.Description
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) GetName() *string {
	return s.Name
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) GetNextHopId() *string {
	return s.NextHopId
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) GetNextHopType() *string {
	return s.NextHopType
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) GetNextHops() []*ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops {
	return s.NextHops
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) GetStatus() *string {
	return s.Status
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) SetDescription(v string) *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels {
	s.Description = &v
	return s
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) SetDestinationCidrBlock(v string) *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels {
	s.DestinationCidrBlock = &v
	return s
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) SetName(v string) *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels {
	s.Name = &v
	return s
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) SetNextHopId(v string) *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels {
	s.NextHopId = &v
	return s
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) SetNextHopType(v string) *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels {
	s.NextHopType = &v
	return s
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) SetNextHops(v []*ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops) *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels {
	s.NextHops = v
	return s
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) SetStatus(v string) *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels {
	s.Status = &v
	return s
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModels) Validate() error {
	if s.NextHops != nil {
		for _, item := range s.NextHops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops struct {
	// Indicates whether the route is available. Valid values:
	//
	// 	- **0**: unavailable
	//
	// 	- **1**: available
	//
	// example:
	//
	// 1
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the next hop.
	//
	// example:
	//
	// vpn-bp10zyaph5cc8b7c7****
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The type of the next hop. Valid values:
	//
	// 	- **Instance*	- (default): an ECS instance
	//
	// 	- **HaVip**: a high-availability virtual IP address (HaVip).
	//
	// 	- **VpnGateway**: a VPN gateway
	//
	// 	- **NatGateway**: a NAT gateway
	//
	// 	- **NetworkInterface**: a secondary ENI
	//
	// 	- **RouterInterface**: a router interface
	//
	// 	- **IPv6Gateway**: an IPv6 gateway
	//
	// 	- **Attachment**: a transit router
	//
	// example:
	//
	// Instance
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The weight of the route.
	//
	// example:
	//
	// 100
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops) GetEnabled() *string {
	return s.Enabled
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops) GetNextHopId() *string {
	return s.NextHopId
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops) GetNextHopType() *string {
	return s.NextHopType
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops) GetWeight() *string {
	return s.Weight
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops) SetEnabled(v string) *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops {
	s.Enabled = &v
	return s
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops) SetNextHopId(v string) *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops {
	s.NextHopId = &v
	return s
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops) SetNextHopType(v string) *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops {
	s.NextHopType = &v
	return s
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops) SetWeight(v string) *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops {
	s.Weight = &v
	return s
}

func (s *ListGatewayRouteTableEntriesResponseBodyGatewayRouteEntryModelsNextHops) Validate() error {
	return dara.Validate(s)
}
