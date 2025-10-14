// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransportLayerApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v int64) *UpdateTransportLayerApplicationRequest
	GetApplicationId() *int64
	SetCrossBorderOptimization(v string) *UpdateTransportLayerApplicationRequest
	GetCrossBorderOptimization() *string
	SetIpAccessRule(v string) *UpdateTransportLayerApplicationRequest
	GetIpAccessRule() *string
	SetIpv6(v string) *UpdateTransportLayerApplicationRequest
	GetIpv6() *string
	SetRules(v []*UpdateTransportLayerApplicationRequestRules) *UpdateTransportLayerApplicationRequest
	GetRules() []*UpdateTransportLayerApplicationRequestRules
	SetSiteId(v int64) *UpdateTransportLayerApplicationRequest
	GetSiteId() *int64
	SetStaticIp(v string) *UpdateTransportLayerApplicationRequest
	GetStaticIp() *string
}

type UpdateTransportLayerApplicationRequest struct {
	// Transport layer application ID, which can be obtained by calling the [ListTransportLayerApplications](~~ListTransportLayerApplications~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 165503967****
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
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
	// IP access rule switch. When enabled, the IP access rules in WAF will take effect on the transport layer application.
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
	// on
	Ipv6 *string `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	// Forwarding rule list. Details of each rule. Except for the comment, all other parameters are required.
	Rules []*UpdateTransportLayerApplicationRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
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

func (s UpdateTransportLayerApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransportLayerApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransportLayerApplicationRequest) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *UpdateTransportLayerApplicationRequest) GetCrossBorderOptimization() *string {
	return s.CrossBorderOptimization
}

func (s *UpdateTransportLayerApplicationRequest) GetIpAccessRule() *string {
	return s.IpAccessRule
}

func (s *UpdateTransportLayerApplicationRequest) GetIpv6() *string {
	return s.Ipv6
}

func (s *UpdateTransportLayerApplicationRequest) GetRules() []*UpdateTransportLayerApplicationRequestRules {
	return s.Rules
}

func (s *UpdateTransportLayerApplicationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateTransportLayerApplicationRequest) GetStaticIp() *string {
	return s.StaticIp
}

func (s *UpdateTransportLayerApplicationRequest) SetApplicationId(v int64) *UpdateTransportLayerApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateTransportLayerApplicationRequest) SetCrossBorderOptimization(v string) *UpdateTransportLayerApplicationRequest {
	s.CrossBorderOptimization = &v
	return s
}

func (s *UpdateTransportLayerApplicationRequest) SetIpAccessRule(v string) *UpdateTransportLayerApplicationRequest {
	s.IpAccessRule = &v
	return s
}

func (s *UpdateTransportLayerApplicationRequest) SetIpv6(v string) *UpdateTransportLayerApplicationRequest {
	s.Ipv6 = &v
	return s
}

func (s *UpdateTransportLayerApplicationRequest) SetRules(v []*UpdateTransportLayerApplicationRequestRules) *UpdateTransportLayerApplicationRequest {
	s.Rules = v
	return s
}

func (s *UpdateTransportLayerApplicationRequest) SetSiteId(v int64) *UpdateTransportLayerApplicationRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateTransportLayerApplicationRequest) SetStaticIp(v string) *UpdateTransportLayerApplicationRequest {
	s.StaticIp = &v
	return s
}

func (s *UpdateTransportLayerApplicationRequest) Validate() error {
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

type UpdateTransportLayerApplicationRequestRules struct {
	// Client IP pass-through protocol, supports:
	//
	// - **off**: No pass-through.
	//
	// - **PPv1**: PROXY Protocol v1, supports client IP pass-through for TCP protocol.
	//
	// - **PPv2**: PROXY Protocol v2, supports client IP pass-through for TCP and UDP protocols.
	//
	// - **SPP**: Simple Proxy Protocol, supports client IP pass-through for UDP protocol.
	//
	// example:
	//
	// SPP
	ClientIPPassThroughMode *string `json:"ClientIPPassThroughMode,omitempty" xml:"ClientIPPassThroughMode,omitempty"`
	// Comment information for the rule.
	//
	// example:
	//
	// 123
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Edge port. Supports:
	//
	// - A single port, e.g., 80.
	//
	// - Port range, e.g., 81-85, representing ports 81, 82, 83, 84, 85.
	//
	// - Combination of ports and port ranges, separated by commas, e.g., 80,81-85,90, representing ports 80, 81, 82, 83, 84, 85, 90.
	//
	// - Edge ports within a single rule and between multiple rules must not overlap.
	//
	// example:
	//
	// 80
	EdgePort *string `json:"EdgePort,omitempty" xml:"EdgePort,omitempty"`
	// Forwarding rule protocol, supports:
	//
	// - TCP: TCP protocol.
	//
	// - UDP: UDP protocol.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Specific value of the source.
	//
	// example:
	//
	// 1.1.1.1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Source port. Supports:
	//
	// - A single port, when the source port is a single port, any valid edge port combination is supported.
	//
	// - Port range, only when the edge port is a port range, the source port can be set as a port range, and the size of the range must match that of the edge port. For example, if the edge port is 90-93, the source port cannot be set to 81-85 because the source port range is 5 and the edge port range is 3, which do not match.
	//
	// example:
	//
	// 80
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// Source type, supports:
	//
	// - **ip**: IP address.
	//
	// - **domain**: Domain name.
	//
	// - **OP**: Origin pool.
	//
	// - **LB**: Load balancer.
	//
	// example:
	//
	// ip
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s UpdateTransportLayerApplicationRequestRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransportLayerApplicationRequestRules) GoString() string {
	return s.String()
}

func (s *UpdateTransportLayerApplicationRequestRules) GetClientIPPassThroughMode() *string {
	return s.ClientIPPassThroughMode
}

func (s *UpdateTransportLayerApplicationRequestRules) GetComment() *string {
	return s.Comment
}

func (s *UpdateTransportLayerApplicationRequestRules) GetEdgePort() *string {
	return s.EdgePort
}

func (s *UpdateTransportLayerApplicationRequestRules) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateTransportLayerApplicationRequestRules) GetSource() *string {
	return s.Source
}

func (s *UpdateTransportLayerApplicationRequestRules) GetSourcePort() *string {
	return s.SourcePort
}

func (s *UpdateTransportLayerApplicationRequestRules) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateTransportLayerApplicationRequestRules) SetClientIPPassThroughMode(v string) *UpdateTransportLayerApplicationRequestRules {
	s.ClientIPPassThroughMode = &v
	return s
}

func (s *UpdateTransportLayerApplicationRequestRules) SetComment(v string) *UpdateTransportLayerApplicationRequestRules {
	s.Comment = &v
	return s
}

func (s *UpdateTransportLayerApplicationRequestRules) SetEdgePort(v string) *UpdateTransportLayerApplicationRequestRules {
	s.EdgePort = &v
	return s
}

func (s *UpdateTransportLayerApplicationRequestRules) SetProtocol(v string) *UpdateTransportLayerApplicationRequestRules {
	s.Protocol = &v
	return s
}

func (s *UpdateTransportLayerApplicationRequestRules) SetSource(v string) *UpdateTransportLayerApplicationRequestRules {
	s.Source = &v
	return s
}

func (s *UpdateTransportLayerApplicationRequestRules) SetSourcePort(v string) *UpdateTransportLayerApplicationRequestRules {
	s.SourcePort = &v
	return s
}

func (s *UpdateTransportLayerApplicationRequestRules) SetSourceType(v string) *UpdateTransportLayerApplicationRequestRules {
	s.SourceType = &v
	return s
}

func (s *UpdateTransportLayerApplicationRequestRules) Validate() error {
	return dara.Validate(s)
}
