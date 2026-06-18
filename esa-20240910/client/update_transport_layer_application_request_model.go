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
	// The transport layer application ID. You can obtain this ID by calling the [ListTransportLayerApplications](~~ListTransportLayerApplications~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 165503967****
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Specifies whether to enable cross-border optimization for network access from the Chinese mainland. This feature is disabled by default. Valid values:
	//
	// - on: Enables the feature.
	//
	// - off: Disables the feature.
	//
	// example:
	//
	// on
	CrossBorderOptimization *string `json:"CrossBorderOptimization,omitempty" xml:"CrossBorderOptimization,omitempty"`
	// Specifies whether to enable IP access rules. If enabled, the IP access rules in WAF apply to the transport layer application. Valid values:
	//
	// - on: Enables the feature.
	//
	// - off: Disables the feature.
	//
	// example:
	//
	// on
	IpAccessRule *string `json:"IpAccessRule,omitempty" xml:"IpAccessRule,omitempty"`
	// Specifies whether to enable IPv6. Valid values: `on` and `off`.
	//
	// example:
	//
	// on
	Ipv6                *string `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	KeepAliveProtection *string `json:"KeepAliveProtection,omitempty" xml:"KeepAliveProtection,omitempty"`
	// A list of forwarding rules. For each rule, all parameters are required except for `Comment`.
	Rules []*UpdateTransportLayerApplicationRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The site ID. You can obtain this ID by calling the [ListSites](~~ListSites~~) operation.
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
	// Specifies the protocol for client IP pass-through. Valid values:
	//
	// - **off**: Disables client IP pass-through.
	//
	// - **PPv1**: PROXY Protocol v1. Supports client IP pass-through for the TCP protocol.
	//
	// - **PPv2**: PROXY Protocol v2. Supports client IP pass-through for both TCP and UDP protocols.
	//
	// - **SPP**: Simple Proxy Protocol. Supports client IP pass-through for the UDP protocol.
	//
	// example:
	//
	// SPP
	ClientIPPassThroughMode *string `json:"ClientIPPassThroughMode,omitempty" xml:"ClientIPPassThroughMode,omitempty"`
	// An optional comment for the forwarding rule.
	//
	// example:
	//
	// 123
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The edge port. The following formats are supported:
	//
	// - A single port, for example, `80`.
	//
	// - A port range, for example, `81-85`. This range includes ports 81, 82, 83, 84, and 85.
	//
	// - A combination of ports and port ranges separated by commas, for example, `80,81-85,90`. This includes ports 80, 81, 82, 83, 84, 85, and 90.
	//
	// - Edge ports cannot overlap within a single rule or across multiple rules.
	//
	// example:
	//
	// 80
	EdgePort *string `json:"EdgePort,omitempty" xml:"EdgePort,omitempty"`
	// The forwarding protocol. Valid values:
	//
	// - TCP: Transmission Control Protocol.
	//
	// - UDP: User Datagram Protocol.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source, which must correspond to the specified `SourceType`. For example, if `SourceType` is `ip`, this parameter must be an IP address.
	//
	// example:
	//
	// 1.1.1.1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The source port. The following formats are supported:
	//
	// - A single port. When a single source port is used, any valid format can be used for the edge port.
	//
	// - A port range. You can specify a port range for the source port only if the edge port is also a port range, and their sizes must match. For example, if `EdgePort` is `90-93`, you cannot set `SourcePort` to `81-85` because their sizes (4 and 5 ports, respectively) do not match.
	//
	// example:
	//
	// 80
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The type of the source. Valid values:
	//
	// - **ip**: An IP address.
	//
	// - **domain**: A domain name.
	//
	// - **OP**: An origin pool.
	//
	// - **LB**: A load balancer.
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
