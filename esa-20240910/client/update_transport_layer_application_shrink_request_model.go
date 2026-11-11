// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransportLayerApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v int64) *UpdateTransportLayerApplicationShrinkRequest
	GetApplicationId() *int64
	SetCrossBorderOptimization(v string) *UpdateTransportLayerApplicationShrinkRequest
	GetCrossBorderOptimization() *string
	SetIpAccessRule(v string) *UpdateTransportLayerApplicationShrinkRequest
	GetIpAccessRule() *string
	SetIpv6(v string) *UpdateTransportLayerApplicationShrinkRequest
	GetIpv6() *string
	SetKeepAliveProtection(v string) *UpdateTransportLayerApplicationShrinkRequest
	GetKeepAliveProtection() *string
	SetRulesShrink(v string) *UpdateTransportLayerApplicationShrinkRequest
	GetRulesShrink() *string
	SetSiteId(v int64) *UpdateTransportLayerApplicationShrinkRequest
	GetSiteId() *int64
	SetStaticIp(v string) *UpdateTransportLayerApplicationShrinkRequest
	GetStaticIp() *string
}

type UpdateTransportLayerApplicationShrinkRequest struct {
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
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
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

func (s UpdateTransportLayerApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransportLayerApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransportLayerApplicationShrinkRequest) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *UpdateTransportLayerApplicationShrinkRequest) GetCrossBorderOptimization() *string {
	return s.CrossBorderOptimization
}

func (s *UpdateTransportLayerApplicationShrinkRequest) GetIpAccessRule() *string {
	return s.IpAccessRule
}

func (s *UpdateTransportLayerApplicationShrinkRequest) GetIpv6() *string {
	return s.Ipv6
}

func (s *UpdateTransportLayerApplicationShrinkRequest) GetKeepAliveProtection() *string {
	return s.KeepAliveProtection
}

func (s *UpdateTransportLayerApplicationShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *UpdateTransportLayerApplicationShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateTransportLayerApplicationShrinkRequest) GetStaticIp() *string {
	return s.StaticIp
}

func (s *UpdateTransportLayerApplicationShrinkRequest) SetApplicationId(v int64) *UpdateTransportLayerApplicationShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateTransportLayerApplicationShrinkRequest) SetCrossBorderOptimization(v string) *UpdateTransportLayerApplicationShrinkRequest {
	s.CrossBorderOptimization = &v
	return s
}

func (s *UpdateTransportLayerApplicationShrinkRequest) SetIpAccessRule(v string) *UpdateTransportLayerApplicationShrinkRequest {
	s.IpAccessRule = &v
	return s
}

func (s *UpdateTransportLayerApplicationShrinkRequest) SetIpv6(v string) *UpdateTransportLayerApplicationShrinkRequest {
	s.Ipv6 = &v
	return s
}

func (s *UpdateTransportLayerApplicationShrinkRequest) SetKeepAliveProtection(v string) *UpdateTransportLayerApplicationShrinkRequest {
	s.KeepAliveProtection = &v
	return s
}

func (s *UpdateTransportLayerApplicationShrinkRequest) SetRulesShrink(v string) *UpdateTransportLayerApplicationShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *UpdateTransportLayerApplicationShrinkRequest) SetSiteId(v int64) *UpdateTransportLayerApplicationShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateTransportLayerApplicationShrinkRequest) SetStaticIp(v string) *UpdateTransportLayerApplicationShrinkRequest {
	s.StaticIp = &v
	return s
}

func (s *UpdateTransportLayerApplicationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
