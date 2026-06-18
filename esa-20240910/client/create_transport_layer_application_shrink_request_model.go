// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransportLayerApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossBorderOptimization(v string) *CreateTransportLayerApplicationShrinkRequest
	GetCrossBorderOptimization() *string
	SetIpAccessRule(v string) *CreateTransportLayerApplicationShrinkRequest
	GetIpAccessRule() *string
	SetIpv6(v string) *CreateTransportLayerApplicationShrinkRequest
	GetIpv6() *string
	SetKeepAliveProtection(v string) *CreateTransportLayerApplicationShrinkRequest
	GetKeepAliveProtection() *string
	SetRecordName(v string) *CreateTransportLayerApplicationShrinkRequest
	GetRecordName() *string
	SetRulesShrink(v string) *CreateTransportLayerApplicationShrinkRequest
	GetRulesShrink() *string
	SetSiteId(v int64) *CreateTransportLayerApplicationShrinkRequest
	GetSiteId() *int64
	SetStaticIp(v string) *CreateTransportLayerApplicationShrinkRequest
	GetStaticIp() *string
}

type CreateTransportLayerApplicationShrinkRequest struct {
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
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
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

func (s CreateTransportLayerApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransportLayerApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTransportLayerApplicationShrinkRequest) GetCrossBorderOptimization() *string {
	return s.CrossBorderOptimization
}

func (s *CreateTransportLayerApplicationShrinkRequest) GetIpAccessRule() *string {
	return s.IpAccessRule
}

func (s *CreateTransportLayerApplicationShrinkRequest) GetIpv6() *string {
	return s.Ipv6
}

func (s *CreateTransportLayerApplicationShrinkRequest) GetKeepAliveProtection() *string {
	return s.KeepAliveProtection
}

func (s *CreateTransportLayerApplicationShrinkRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *CreateTransportLayerApplicationShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *CreateTransportLayerApplicationShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateTransportLayerApplicationShrinkRequest) GetStaticIp() *string {
	return s.StaticIp
}

func (s *CreateTransportLayerApplicationShrinkRequest) SetCrossBorderOptimization(v string) *CreateTransportLayerApplicationShrinkRequest {
	s.CrossBorderOptimization = &v
	return s
}

func (s *CreateTransportLayerApplicationShrinkRequest) SetIpAccessRule(v string) *CreateTransportLayerApplicationShrinkRequest {
	s.IpAccessRule = &v
	return s
}

func (s *CreateTransportLayerApplicationShrinkRequest) SetIpv6(v string) *CreateTransportLayerApplicationShrinkRequest {
	s.Ipv6 = &v
	return s
}

func (s *CreateTransportLayerApplicationShrinkRequest) SetKeepAliveProtection(v string) *CreateTransportLayerApplicationShrinkRequest {
	s.KeepAliveProtection = &v
	return s
}

func (s *CreateTransportLayerApplicationShrinkRequest) SetRecordName(v string) *CreateTransportLayerApplicationShrinkRequest {
	s.RecordName = &v
	return s
}

func (s *CreateTransportLayerApplicationShrinkRequest) SetRulesShrink(v string) *CreateTransportLayerApplicationShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *CreateTransportLayerApplicationShrinkRequest) SetSiteId(v int64) *CreateTransportLayerApplicationShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateTransportLayerApplicationShrinkRequest) SetStaticIp(v string) *CreateTransportLayerApplicationShrinkRequest {
	s.StaticIp = &v
	return s
}

func (s *CreateTransportLayerApplicationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
