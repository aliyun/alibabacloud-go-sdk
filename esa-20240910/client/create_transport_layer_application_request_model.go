// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransportLayerApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossBorderOptimization(v string) *CreateTransportLayerApplicationRequest
	GetCrossBorderOptimization() *string
	SetIpAccessRule(v string) *CreateTransportLayerApplicationRequest
	GetIpAccessRule() *string
	SetIpv6(v string) *CreateTransportLayerApplicationRequest
	GetIpv6() *string
	SetKeepAliveProtection(v string) *CreateTransportLayerApplicationRequest
	GetKeepAliveProtection() *string
	SetRecordName(v string) *CreateTransportLayerApplicationRequest
	GetRecordName() *string
	SetRules(v []*CreateTransportLayerApplicationRequestRules) *CreateTransportLayerApplicationRequest
	GetRules() []*CreateTransportLayerApplicationRequestRules
	SetSiteId(v int64) *CreateTransportLayerApplicationRequest
	GetSiteId() *int64
	SetStaticIp(v string) *CreateTransportLayerApplicationRequest
	GetStaticIp() *string
}

type CreateTransportLayerApplicationRequest struct {
	// Enables or disables network optimization for access from the Chinese mainland. This feature is disabled by default. Valid values:
	//
	// - `on`: Enables the optimization.
	//
	// - `off`: Disables the optimization.
	//
	// example:
	//
	// on
	CrossBorderOptimization *string `json:"CrossBorderOptimization,omitempty" xml:"CrossBorderOptimization,omitempty"`
	// Applies IP access rules from Web Application Firewall (WAF) to this Transport Layer Application. Valid values:
	//
	// - `on`: Enables the feature.
	//
	// - `off`: Disables the feature.
	//
	// example:
	//
	// on
	IpAccessRule *string `json:"IpAccessRule,omitempty" xml:"IpAccessRule,omitempty"`
	// Enables or disables IPv6 support.
	//
	// example:
	//
	// off
	Ipv6 *string `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	// Enables or disables keep-alive protection.
	KeepAliveProtection *string `json:"KeepAliveProtection,omitempty" xml:"KeepAliveProtection,omitempty"`
	// The domain name of the Transport Layer Application.
	//
	// This parameter is required.
	//
	// example:
	//
	// aaa.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The list of forwarding rules.
	//
	// This parameter is required.
	Rules []*CreateTransportLayerApplicationRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId   *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	StaticIp *string `json:"StaticIp,omitempty" xml:"StaticIp,omitempty"`
}

func (s CreateTransportLayerApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransportLayerApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateTransportLayerApplicationRequest) GetCrossBorderOptimization() *string {
	return s.CrossBorderOptimization
}

func (s *CreateTransportLayerApplicationRequest) GetIpAccessRule() *string {
	return s.IpAccessRule
}

func (s *CreateTransportLayerApplicationRequest) GetIpv6() *string {
	return s.Ipv6
}

func (s *CreateTransportLayerApplicationRequest) GetKeepAliveProtection() *string {
	return s.KeepAliveProtection
}

func (s *CreateTransportLayerApplicationRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *CreateTransportLayerApplicationRequest) GetRules() []*CreateTransportLayerApplicationRequestRules {
	return s.Rules
}

func (s *CreateTransportLayerApplicationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateTransportLayerApplicationRequest) GetStaticIp() *string {
	return s.StaticIp
}

func (s *CreateTransportLayerApplicationRequest) SetCrossBorderOptimization(v string) *CreateTransportLayerApplicationRequest {
	s.CrossBorderOptimization = &v
	return s
}

func (s *CreateTransportLayerApplicationRequest) SetIpAccessRule(v string) *CreateTransportLayerApplicationRequest {
	s.IpAccessRule = &v
	return s
}

func (s *CreateTransportLayerApplicationRequest) SetIpv6(v string) *CreateTransportLayerApplicationRequest {
	s.Ipv6 = &v
	return s
}

func (s *CreateTransportLayerApplicationRequest) SetKeepAliveProtection(v string) *CreateTransportLayerApplicationRequest {
	s.KeepAliveProtection = &v
	return s
}

func (s *CreateTransportLayerApplicationRequest) SetRecordName(v string) *CreateTransportLayerApplicationRequest {
	s.RecordName = &v
	return s
}

func (s *CreateTransportLayerApplicationRequest) SetRules(v []*CreateTransportLayerApplicationRequestRules) *CreateTransportLayerApplicationRequest {
	s.Rules = v
	return s
}

func (s *CreateTransportLayerApplicationRequest) SetSiteId(v int64) *CreateTransportLayerApplicationRequest {
	s.SiteId = &v
	return s
}

func (s *CreateTransportLayerApplicationRequest) SetStaticIp(v string) *CreateTransportLayerApplicationRequest {
	s.StaticIp = &v
	return s
}

func (s *CreateTransportLayerApplicationRequest) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTransportLayerApplicationRequestRules struct {
	// The client IP passthrough protocol. Valid values:
	//
	// - **off**: Disables client IP passthrough.
	//
	// - **PPv1**: PROXY Protocol v1. Preserves the client IP address for TCP connections.
	//
	// - **PPv2**: PROXY Protocol v2. Preserves the client IP address for TCP and UDP connections.
	//
	// - **SPP**: Simple Proxy Protocol. Preserves the client IP address for UDP connections.
	//
	// This parameter is required.
	//
	// example:
	//
	// SPP
	ClientIPPassThroughMode *string `json:"ClientIPPassThroughMode,omitempty" xml:"ClientIPPassThroughMode,omitempty"`
	// An optional comment for the rule.
	//
	// example:
	//
	// Test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The edge port. Supported formats:
	//
	// - A single port, for example, `80`.
	//
	// - A port range, for example, `81-85`, which includes ports 81, 82, 83, 84, and 85.
	//
	// - A combination of ports and port ranges separated by commas, for example, `80,81-85,90`, which includes ports 80, 81, 82, 83, 84, 85, and 90.
	//
	// Edge ports cannot overlap within a single rule or across multiple rules.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	EdgePort *string `json:"EdgePort,omitempty" xml:"EdgePort,omitempty"`
	// The forwarding protocol. Valid values:
	//
	// - `TCP`: The TCP protocol.
	//
	// - `UDP`: The UDP protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The origin address. The value must match the specified `SourceType`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.1.1.1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The origin port. Supported formats:
	//
	// - A single port, which supports any valid combination of edge ports.
	//
	// - A port range. You can specify a port range only if the edge port is also a port range. The origin and edge port ranges must be the same size. For example, if the edge port range is `90-93` (4 ports), the origin port range cannot be `81-85` (5 ports).
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The origin type. Valid values:
	//
	// - **ip**: An IP address.
	//
	// - **domain**: A domain name.
	//
	// - **OP**: An origin pool.
	//
	// - **LB**: A load balancer.
	//
	// This parameter is required.
	//
	// example:
	//
	// ip
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s CreateTransportLayerApplicationRequestRules) String() string {
	return dara.Prettify(s)
}

func (s CreateTransportLayerApplicationRequestRules) GoString() string {
	return s.String()
}

func (s *CreateTransportLayerApplicationRequestRules) GetClientIPPassThroughMode() *string {
	return s.ClientIPPassThroughMode
}

func (s *CreateTransportLayerApplicationRequestRules) GetComment() *string {
	return s.Comment
}

func (s *CreateTransportLayerApplicationRequestRules) GetEdgePort() *string {
	return s.EdgePort
}

func (s *CreateTransportLayerApplicationRequestRules) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateTransportLayerApplicationRequestRules) GetSource() *string {
	return s.Source
}

func (s *CreateTransportLayerApplicationRequestRules) GetSourcePort() *string {
	return s.SourcePort
}

func (s *CreateTransportLayerApplicationRequestRules) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateTransportLayerApplicationRequestRules) SetClientIPPassThroughMode(v string) *CreateTransportLayerApplicationRequestRules {
	s.ClientIPPassThroughMode = &v
	return s
}

func (s *CreateTransportLayerApplicationRequestRules) SetComment(v string) *CreateTransportLayerApplicationRequestRules {
	s.Comment = &v
	return s
}

func (s *CreateTransportLayerApplicationRequestRules) SetEdgePort(v string) *CreateTransportLayerApplicationRequestRules {
	s.EdgePort = &v
	return s
}

func (s *CreateTransportLayerApplicationRequestRules) SetProtocol(v string) *CreateTransportLayerApplicationRequestRules {
	s.Protocol = &v
	return s
}

func (s *CreateTransportLayerApplicationRequestRules) SetSource(v string) *CreateTransportLayerApplicationRequestRules {
	s.Source = &v
	return s
}

func (s *CreateTransportLayerApplicationRequestRules) SetSourcePort(v string) *CreateTransportLayerApplicationRequestRules {
	s.SourcePort = &v
	return s
}

func (s *CreateTransportLayerApplicationRequestRules) SetSourceType(v string) *CreateTransportLayerApplicationRequestRules {
	s.SourceType = &v
	return s
}

func (s *CreateTransportLayerApplicationRequestRules) Validate() error {
	return dara.Validate(s)
}
