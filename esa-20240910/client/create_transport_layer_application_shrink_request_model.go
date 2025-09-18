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
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
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
