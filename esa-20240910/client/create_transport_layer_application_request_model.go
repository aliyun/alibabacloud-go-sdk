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
	// Whether to enable China mainland network access optimization, default is disabled. Value range:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	CrossBorderOptimization *string `json:"CrossBorderOptimization,omitempty" xml:"CrossBorderOptimization,omitempty"`
	// IP access rule switch. When enabled, the WAF\\"s IP access rules apply to the transport layer application.
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	IpAccessRule *string `json:"IpAccessRule,omitempty" xml:"IpAccessRule,omitempty"`
	// IPv6 switch.
	//
	// example:
	//
	// off
	Ipv6 *string `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	// Domain name of the transport layer application.
	//
	// This parameter is required.
	//
	// example:
	//
	// aaa.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// List of forwarding rules.
	//
	// This parameter is required.
	Rules []*CreateTransportLayerApplicationRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
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
	// Client IP pass-through protocol, supporting:
	//
	// - **off**: No pass-through.
	//
	// - **PPv1**: PROXY Protocol v1, supports client IP pass-through for TCP protocol.
	//
	// - **PPv2**: PROXY Protocol v2, supports client IP pass-through for TCP and UDP protocols.
	//
	// - **SPP**: Simple Proxy Protocol, supports client IP pass-through for UDP protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// SPP
	ClientIPPassThroughMode *string `json:"ClientIPPassThroughMode,omitempty" xml:"ClientIPPassThroughMode,omitempty"`
	// Comment information for the rule (optional).
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Edge port. Supports:
	//
	// - A single port, such as 80.
	//
	// - Port range, such as 81-85, representing ports 81, 82, 83, 84, and 85.
	//
	// - Combination of ports and port ranges, separated by commas, such as 80,81-85,90, representing ports 80, 81, 82, 83, 84, 85, and 90.
	//
	// Edge ports within a single rule and between multiple rules must not overlap.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	EdgePort *string `json:"EdgePort,omitempty" xml:"EdgePort,omitempty"`
	// Forwarding rule protocol, with values:
	//
	// - TCP: TCP protocol.
	//
	// - UDP: UDP protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Specific value of the origin, which needs to match the origin type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.1.1.1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Origin port. Supports:
	//
	// - A single port, when the origin port is a single port, any valid combination of edge ports is supported.
	//
	// - Port range, only when the edge port is a port range, the origin port can be set to a port range, and the size of the range must match that of the edge port. For example, if the edge port is 90-93, the origin port cannot be set to 81-85 because the origin port range is 5 and the edge port range is 3, which do not match.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// Origin type, supporting:
	//
	// - **ip**: IP address.
	//
	// - **domain**: Domain name.
	//
	// - **OP**: Origin pool.
	//
	// - **LB**: Load balancer.
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
