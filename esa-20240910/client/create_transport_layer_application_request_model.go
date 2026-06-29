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
	// Specifies whether to enable network access optimization for the Chinese mainland. This feature is disabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	CrossBorderOptimization *string `json:"CrossBorderOptimization,omitempty" xml:"CrossBorderOptimization,omitempty"`
	// The IP access rule switch. When enabled, IP access rules in WAF take effect for the Layer 4 application. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	IpAccessRule *string `json:"IpAccessRule,omitempty" xml:"IpAccessRule,omitempty"`
	// Specifies whether to enable IPv6. This feature is disabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// off
	Ipv6 *string `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	// Specifies whether to enable keep-alive protection. This feature is disabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// off
	KeepAliveProtection *string `json:"KeepAliveProtection,omitempty" xml:"KeepAliveProtection,omitempty"`
	// The domain name of the Layer 4 application.
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
	// The site ID. You can call the [ListSites](~~ListSites~~) operation to obtain the site ID. The site must be activated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Specifies whether to enable static IP. This feature is disabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// off
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
	// The client IP pass-through protocol. Valid values:
	//
	// - **off**: disabled.
	//
	// - **PPv1**: PROXY Protocol v1. Supports client IP pass-through for TCP.
	//
	// - **PPv2**: PROXY Protocol v2. Supports client IP pass-through for TCP and UDP.
	//
	// - **SPP**: Simple Proxy Protocol. Supports client IP pass-through for UDP.
	//
	// This parameter is required.
	//
	// example:
	//
	// SPP
	ClientIPPassThroughMode *string `json:"ClientIPPassThroughMode,omitempty" xml:"ClientIPPassThroughMode,omitempty"`
	// The comment for the rule (optional).
	//
	// example:
	//
	// Test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The edge port. The following formats are supported:
	//
	// - A single port, such as 80.
	//
	// - A port range, such as 81-85, which represents ports 81, 82, 83, 84, and 85.
	//
	// - A combination of ports and port ranges separated by commas, such as 80,81-85,90, which represents ports 80, 81, 82, 83, 84, 85, and 90.
	//
	// Edge ports within a single rule and across multiple rules must not overlap.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	EdgePort *string `json:"EdgePort,omitempty" xml:"EdgePort,omitempty"`
	// The forwarding rule protocol. Valid values:
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
	// The specific value of the origin, which must match the origin type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.1.1.1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Origin Server Port. The following formats are supported:
	//
	// - A single port. When Origin Server Port is a single port, any valid edge port combination is supported.
	//
	// - A port range. Origin Server Port can be set to a port range only when the edge port is also a port range, and the range size must match the edge port range. For example, if the edge port is 90-93, you cannot set Origin Server Port to 81-85 because Origin Server Port range size is 5 while the edge port range size is 4, which is inconsistent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The origin type. Valid values:
	//
	// - **ip**: IP address.
	//
	// - **domain**: domain name.
	//
	// - **OP**: origin IPAM pool.
	//
	// - **LB**: load balancing.
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
