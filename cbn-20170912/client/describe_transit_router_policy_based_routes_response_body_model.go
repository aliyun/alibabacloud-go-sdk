// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouterPolicyBasedRoutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeTransitRouterPolicyBasedRoutesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBody
	GetNextToken() *string
	SetPolicyBasedRoutes(v []*DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) *DescribeTransitRouterPolicyBasedRoutesResponseBody
	GetPolicyBasedRoutes() []*DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes
	SetRequestId(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeTransitRouterPolicyBasedRoutesResponseBody
	GetTotalCount() *int32
}

type DescribeTransitRouterPolicyBasedRoutesResponseBody struct {
	// The number of entries returned per page. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If **NextToken*	- is empty, no next query is available.
	//
	// - If **NextToken*	- returns a value, the value is the token used to start the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of policy-based route rules.
	PolicyBasedRoutes []*DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes `json:"PolicyBasedRoutes,omitempty" xml:"PolicyBasedRoutes,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// F7DDDC17-FA06-4AC2-8F35-59D2470FCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTransitRouterPolicyBasedRoutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyBasedRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBody) GetPolicyBasedRoutes() []*DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes {
	return s.PolicyBasedRoutes
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBody) SetMaxResults(v int32) *DescribeTransitRouterPolicyBasedRoutesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBody) SetNextToken(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBody) SetPolicyBasedRoutes(v []*DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) *DescribeTransitRouterPolicyBasedRoutesResponseBody {
	s.PolicyBasedRoutes = v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBody) SetRequestId(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBody) SetTotalCount(v int32) *DescribeTransitRouterPolicyBasedRoutesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBody) Validate() error {
	if s.PolicyBasedRoutes != nil {
		for _, item := range s.PolicyBasedRoutes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes struct {
	// The address family. Valid values:
	//
	// - IPv4
	//
	// - IPv6
	//
	// example:
	//
	// IPv4
	AddressFamily *string `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	// The description of the policy-based route table.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port. Valid values: **-1*	- or **1*	- to **65535**.
	//
	// Matches traffic whose port numbers fall within the destination port range. If this parameter is not specified, traffic to any destination port is matched.
	//
	// The port range format is described as follows:
	//
	// - 80: port 80.
	//
	// - 80/80: port 80.
	//
	// - 1/200: ports 1 to 200.
	//
	// - -1/-1: all ports.
	//
	// example:
	//
	// 80/80
	DestPortRange *string `json:"DestPortRange,omitempty" xml:"DestPortRange,omitempty"`
	// The destination CIDR block.
	//
	// example:
	//
	// 10.185.33.0/24
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The Differentiated Services Code Point (DSCP) value of the traffic packets to match. Valid values: **0*	- to **63**.
	//
	// The DSCP value format is described as follows:
	//
	// - 1: matches traffic with a DSCP value of 1.
	//
	// - 1/10: matches traffic with DSCP values from 1 to 10.
	//
	// - -1/-1: matches traffic with any DSCP value.
	//
	// example:
	//
	// 1/10
	Dscp *string `json:"Dscp,omitempty" xml:"Dscp,omitempty"`
	// The name of the policy-based route rule.
	//
	// The name can be empty or 1 to 128 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// None
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The instance ID of the policy-based routing rule.
	//
	// example:
	//
	// pbr-123qwead
	PolicyBasedRouteId *string `json:"PolicyBasedRouteId,omitempty" xml:"PolicyBasedRouteId,omitempty"`
	// The instance ID of the policy-based routing table.
	//
	// example:
	//
	// ptbl-123adsfzc
	PolicyTableId *string `json:"PolicyTableId,omitempty" xml:"PolicyTableId,omitempty"`
	// The priority of the policy-based route. Valid values: **1*	- to **100**.
	//
	// The priority of each policy-based route in a policy-based route table of a transit router must be unique. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The protocol type of the traffic packets.
	//
	// Traffic classification rules support matching traffic of various protocol types, such as **HTTP**, **HTTPS**, **TCP**, **UDP**, **SSH**, and **Telnet**. For more protocol types, log on to the [Cloud Enterprise Network (CEN) console](https://cen.console.aliyun.com/cen/list).
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block.
	//
	// example:
	//
	// 10.0.0.0/8
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The source port. Valid values: **-1*	- or **1*	- to **65535**.
	//
	// Matches traffic whose port numbers fall within the source port range. If this parameter is not specified, traffic from any source port is matched.
	//
	// The port range format is described as follows:
	//
	// - 80: port 80.
	//
	// - 80/80: port 80.
	//
	// - 1/200: ports 1 to 200.
	//
	// - -1/-1: all ports.
	//
	// example:
	//
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The status of the policy-based route. Valid values:
	//
	// - **Creating**: The policy-based route is being created.
	//
	// - **Active**: The policy-based route is available.
	//
	// - **Deleting**: The policy-based route is being deleted.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The destination route table.
	//
	// example:
	//
	// vtb-123adf
	TargetRouteTableId *string `json:"TargetRouteTableId,omitempty" xml:"TargetRouteTableId,omitempty"`
}

func (s DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) GetAddressFamily() *string {
	return s.AddressFamily
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) GetDescription() *string {
	return s.Description
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) GetDestPortRange() *string {
	return s.DestPortRange
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) GetDscp() *string {
	return s.Dscp
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) GetName() *string {
	return s.Name
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) GetPolicyBasedRouteId() *string {
	return s.PolicyBasedRouteId
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) GetPolicyTableId() *string {
	return s.PolicyTableId
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) GetStatus() *string {
	return s.Status
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) GetTargetRouteTableId() *string {
	return s.TargetRouteTableId
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) SetAddressFamily(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes {
	s.AddressFamily = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) SetDescription(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes {
	s.Description = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) SetDestPortRange(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes {
	s.DestPortRange = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) SetDestinationCidr(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) SetDscp(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes {
	s.Dscp = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) SetName(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes {
	s.Name = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) SetPolicyBasedRouteId(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes {
	s.PolicyBasedRouteId = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) SetPolicyTableId(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes {
	s.PolicyTableId = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) SetPriority(v int32) *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes {
	s.Priority = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) SetProtocol(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes {
	s.Protocol = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) SetSourceCidr(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes {
	s.SourceCidr = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) SetSourcePortRange(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes {
	s.SourcePortRange = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) SetStatus(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes {
	s.Status = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) SetTargetRouteTableId(v string) *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes {
	s.TargetRouteTableId = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponseBodyPolicyBasedRoutes) Validate() error {
	return dara.Validate(s)
}
