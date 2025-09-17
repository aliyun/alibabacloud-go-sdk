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
}

type CreateTransportLayerApplicationShrinkRequest struct {
	CrossBorderOptimization *string `json:"CrossBorderOptimization,omitempty" xml:"CrossBorderOptimization,omitempty"`
	IpAccessRule            *string `json:"IpAccessRule,omitempty" xml:"IpAccessRule,omitempty"`
	Ipv6                    *string `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aaa.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// This parameter is required.
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
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

func (s *CreateTransportLayerApplicationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
