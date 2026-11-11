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
	SetKeepAliveProtection(v string) *UpdateTransportLayerApplicationRequest
	GetKeepAliveProtection() *string
	SetRules(v []*UpdateTransportLayerApplicationRequestRules) *UpdateTransportLayerApplicationRequest
	GetRules() []*UpdateTransportLayerApplicationRequestRules
	SetSiteId(v int64) *UpdateTransportLayerApplicationRequest
	GetSiteId() *int64
	SetStaticIp(v string) *UpdateTransportLayerApplicationRequest
	GetStaticIp() *string
}

type UpdateTransportLayerApplicationRequest struct {
	// The Layer 4 application ID. You can obtain the ID by calling the [ListTransportLayerApplications](~~ListTransportLayerApplications~~) API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 165503967****
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Specifies whether to enable network access optimization for the Chinese mainland. Default value: off. Valid values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	CrossBorderOptimization *string `json:"CrossBorderOptimization,omitempty" xml:"CrossBorderOptimization,omitempty"`
	// Specifies whether to enable the IP access rule. After this feature is enabled, the IP access rules in WAF take effect for the Layer 4 application. Valid values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	IpAccessRule *string `json:"IpAccessRule,omitempty" xml:"IpAccessRule,omitempty"`
	// Specifies whether to enable IPv6. Valid values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	Ipv6 *string `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	// Specifies whether to enable connection keep-alive protection. Default value: off. Valid values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// off
	KeepAliveProtection *string `json:"KeepAliveProtection,omitempty" xml:"KeepAliveProtection,omitempty"`
	// The list of forwarding rules. For rule details, all parameters except the comment are required for each rule.
	Rules []*UpdateTransportLayerApplicationRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The site ID. You can obtain the ID by calling the [ListSites](~~ListSites~~) API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Specifies whether to enable the static IP address feature. Default value: off. Valid values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// off
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

func (s *UpdateTransportLayerApplicationRequest) GetKeepAliveProtection() *string {
	return s.KeepAliveProtection
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

func (s *UpdateTransportLayerApplicationRequest) SetKeepAliveProtection(v string) *UpdateTransportLayerApplicationRequest {
	s.KeepAliveProtection = &v
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
	// The client IP address pass-through protocol. Valid values:
	//
	// - **off**: Do not pass through.
	//
	// - **PPv1**: PROXY protocol v1, which supports client IP address pass-through for the TCP protocol.
	//
	// - **PPv2**: PROXY protocol v2, which supports client IP address pass-through for the TCP and UDP protocols.
	//
	// - **SPP**: Simple Proxy Protocol, which supports client IP address pass-through for the UDP protocol.
	//
	// example:
	//
	// SPP
	ClientIPPassThroughMode *string `json:"ClientIPPassThroughMode,omitempty" xml:"ClientIPPassThroughMode,omitempty"`
	// The comment for the rule.
	//
	// example:
	//
	// 123
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The edge port. Valid values:
	//
	// - A single port, such as 80.
	//
	// - A port range, such as 81-85, which represents ports 81, 82, 83, 84, and 85.
	//
	// - A combination of ports and port ranges, separated by commas, such as 80,81-85,90, which represents ports 80, 81, 82, 83, 84, 85, and 90.
	//
	// - The edge ports within a single rule and the edge ports across multiple rules must not overlap.
	//
	// example:
	//
	// 80
	EdgePort *string `json:"EdgePort,omitempty" xml:"EdgePort,omitempty"`
	// The protocol of the forwarding rule. Valid values:
	//
	// - TCP: TCP.
	//
	// - UDP: UDP.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The specific value of the origin server.
	//
	// example:
	//
	// 1.1.1.1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The origin server port. Valid values:
	//
	// - A single port. If the origin server port is a single port, any valid combination of edge ports is supported.
	//
	// - A port range. The origin server port can be set to a port range only when the edge port is a port range, and the range size must be the same as that of the edge port. For example, if the edge port is 90-93, you cannot set the origin server port to 81-85 because the origin server port range is 5 and the edge port range is 3, which are inconsistent.
	//
	// example:
	//
	// 80
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The origin server type. Valid values:
	//
	// - **ip**: IP address.
	//
	// - **domain**: Domain name.
	//
	// - **OP**: Origin IPAM pool.
	//
	// - **LB**: Load balancing.
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
