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
	// This parameter is required.
	ApplicationId           *int64                                         `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	CrossBorderOptimization *string                                        `json:"CrossBorderOptimization,omitempty" xml:"CrossBorderOptimization,omitempty"`
	IpAccessRule            *string                                        `json:"IpAccessRule,omitempty" xml:"IpAccessRule,omitempty"`
	Ipv6                    *string                                        `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	KeepAliveProtection     *string                                        `json:"KeepAliveProtection,omitempty" xml:"KeepAliveProtection,omitempty"`
	Rules                   []*UpdateTransportLayerApplicationRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// This parameter is required.
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
	ClientIPPassThroughMode *string `json:"ClientIPPassThroughMode,omitempty" xml:"ClientIPPassThroughMode,omitempty"`
	Comment                 *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	EdgePort                *string `json:"EdgePort,omitempty" xml:"EdgePort,omitempty"`
	Protocol                *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Source                  *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourcePort              *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	SourceType              *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
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
