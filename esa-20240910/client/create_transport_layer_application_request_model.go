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
}

type CreateTransportLayerApplicationRequest struct {
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
	Rules []*CreateTransportLayerApplicationRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
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

func (s *CreateTransportLayerApplicationRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTransportLayerApplicationRequestRules struct {
	// This parameter is required.
	//
	// example:
	//
	// TOA
	ClientIPPassThroughMode *string `json:"ClientIPPassThroughMode,omitempty" xml:"ClientIPPassThroughMode,omitempty"`
	Comment                 *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 80
	EdgePort *string `json:"EdgePort,omitempty" xml:"EdgePort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.1.1.1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 80
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
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
