// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterPolicyBasedRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressFamily(v string) *CreateTransitRouterPolicyBasedRouteRequest
	GetAddressFamily() *string
	SetClientToken(v string) *CreateTransitRouterPolicyBasedRouteRequest
	GetClientToken() *string
	SetDescription(v string) *CreateTransitRouterPolicyBasedRouteRequest
	GetDescription() *string
	SetDestPortRange(v string) *CreateTransitRouterPolicyBasedRouteRequest
	GetDestPortRange() *string
	SetDestinationCidr(v string) *CreateTransitRouterPolicyBasedRouteRequest
	GetDestinationCidr() *string
	SetDryRun(v bool) *CreateTransitRouterPolicyBasedRouteRequest
	GetDryRun() *bool
	SetDscp(v string) *CreateTransitRouterPolicyBasedRouteRequest
	GetDscp() *string
	SetName(v string) *CreateTransitRouterPolicyBasedRouteRequest
	GetName() *string
	SetPolicyTableId(v string) *CreateTransitRouterPolicyBasedRouteRequest
	GetPolicyTableId() *string
	SetPriority(v int32) *CreateTransitRouterPolicyBasedRouteRequest
	GetPriority() *int32
	SetProtocol(v string) *CreateTransitRouterPolicyBasedRouteRequest
	GetProtocol() *string
	SetSourceCidr(v string) *CreateTransitRouterPolicyBasedRouteRequest
	GetSourceCidr() *string
	SetSourcePortRange(v string) *CreateTransitRouterPolicyBasedRouteRequest
	GetSourcePortRange() *string
	SetTargetRouteTableId(v string) *CreateTransitRouterPolicyBasedRouteRequest
	GetTargetRouteTableId() *string
}

type CreateTransitRouterPolicyBasedRouteRequest struct {
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
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the policy-based route.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port. Valid values: **-1*	- or **1*	- to **65535**.
	//
	// This parameter matches traffic whose port numbers fall within the destination port range. If you do not specify this parameter, traffic to any destination port is matched.
	//
	// The port range is specified in the following formats:
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
	// The destination CIDR block of the route entry. Both IPv4 and IPv6 addresses are supported.
	//
	// example:
	//
	// 172.19.2.0/24
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: prechecks the request without creating the routing rule entry in the policy route table. The system checks whether required parameters are specified, the request format, and service limits. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends the request. After the request passes the precheck, the routing rule entry is created in the policy route table.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The Differentiated Services Code Point (DSCP) value of the traffic packet to be matched. Valid values: **0*	- to **63**.
	//
	// The DSCP range is specified in the following formats:
	//
	// - 1: traffic with a DSCP value of 1.
	//
	// - 1/10: traffic with a DSCP value from 1 to 10.
	//
	// - -1/-1: traffic with any DSCP value.
	//
	// example:
	//
	// 10
	Dscp *string `json:"Dscp,omitempty" xml:"Dscp,omitempty"`
	// The name of the policy-based route.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The instance ID of the policy routing table.
	//
	// This parameter is required.
	//
	// example:
	//
	// ptbl-abc123xyz
	PolicyTableId *string `json:"PolicyTableId,omitempty" xml:"PolicyTableId,omitempty"`
	// The priority of the policy-based route. Valid values: **1*	- to **100**.
	//
	// The priority of each policy-based route in a policy route table of a transit router must be unique. A smaller value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The protocol type of the traffic packet.
	//
	// Traffic classification rules support matching traffic of various protocol types, such as **HTTP**, **HTTPS**, **TCP**, **UDP**, **SSH**, and **Telnet**. For more protocol types, log on to the [Cloud Enterprise Network (CEN) console](https://cen.console.aliyun.com/cen/list).
	//
	// example:
	//
	// ICMP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block. The source CIDR block must be in CIDR format. Example: 192.168.1.0/24.
	//
	// example:
	//
	// 10.0.0.0/8
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The source port. Valid values: **-1*	- or **1*	- to **65535**.
	//
	// This parameter matches traffic whose port numbers fall within the source port range. If you do not specify this parameter, traffic from any source port is matched.
	//
	// The port range is specified in the following formats:
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
	// The destination route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-123abcxyz
	TargetRouteTableId *string `json:"TargetRouteTableId,omitempty" xml:"TargetRouteTableId,omitempty"`
}

func (s CreateTransitRouterPolicyBasedRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterPolicyBasedRouteRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) GetAddressFamily() *string {
	return s.AddressFamily
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) GetDestPortRange() *string {
	return s.DestPortRange
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) GetDscp() *string {
	return s.Dscp
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) GetName() *string {
	return s.Name
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) GetPolicyTableId() *string {
	return s.PolicyTableId
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) GetTargetRouteTableId() *string {
	return s.TargetRouteTableId
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) SetAddressFamily(v string) *CreateTransitRouterPolicyBasedRouteRequest {
	s.AddressFamily = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) SetClientToken(v string) *CreateTransitRouterPolicyBasedRouteRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) SetDescription(v string) *CreateTransitRouterPolicyBasedRouteRequest {
	s.Description = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) SetDestPortRange(v string) *CreateTransitRouterPolicyBasedRouteRequest {
	s.DestPortRange = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) SetDestinationCidr(v string) *CreateTransitRouterPolicyBasedRouteRequest {
	s.DestinationCidr = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) SetDryRun(v bool) *CreateTransitRouterPolicyBasedRouteRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) SetDscp(v string) *CreateTransitRouterPolicyBasedRouteRequest {
	s.Dscp = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) SetName(v string) *CreateTransitRouterPolicyBasedRouteRequest {
	s.Name = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) SetPolicyTableId(v string) *CreateTransitRouterPolicyBasedRouteRequest {
	s.PolicyTableId = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) SetPriority(v int32) *CreateTransitRouterPolicyBasedRouteRequest {
	s.Priority = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) SetProtocol(v string) *CreateTransitRouterPolicyBasedRouteRequest {
	s.Protocol = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) SetSourceCidr(v string) *CreateTransitRouterPolicyBasedRouteRequest {
	s.SourceCidr = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) SetSourcePortRange(v string) *CreateTransitRouterPolicyBasedRouteRequest {
	s.SourcePortRange = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) SetTargetRouteTableId(v string) *CreateTransitRouterPolicyBasedRouteRequest {
	s.TargetRouteTableId = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteRequest) Validate() error {
	return dara.Validate(s)
}
