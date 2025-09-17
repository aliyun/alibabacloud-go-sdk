// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransportLayerApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v int64) *GetTransportLayerApplicationResponseBody
	GetApplicationId() *int64
	SetCname(v string) *GetTransportLayerApplicationResponseBody
	GetCname() *string
	SetCrossBorderOptimization(v string) *GetTransportLayerApplicationResponseBody
	GetCrossBorderOptimization() *string
	SetIpAccessRule(v string) *GetTransportLayerApplicationResponseBody
	GetIpAccessRule() *string
	SetIpv6(v string) *GetTransportLayerApplicationResponseBody
	GetIpv6() *string
	SetRecordName(v string) *GetTransportLayerApplicationResponseBody
	GetRecordName() *string
	SetRequestId(v string) *GetTransportLayerApplicationResponseBody
	GetRequestId() *string
	SetRules(v []*GetTransportLayerApplicationResponseBodyRules) *GetTransportLayerApplicationResponseBody
	GetRules() []*GetTransportLayerApplicationResponseBodyRules
	SetRulesCount(v int32) *GetTransportLayerApplicationResponseBody
	GetRulesCount() *int32
	SetSiteId(v int64) *GetTransportLayerApplicationResponseBody
	GetSiteId() *int64
	SetStatus(v string) *GetTransportLayerApplicationResponseBody
	GetStatus() *string
}

type GetTransportLayerApplicationResponseBody struct {
	// example:
	//
	// 17099311410****
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// example.com.ialicdn.com
	Cname                   *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	CrossBorderOptimization *string `json:"CrossBorderOptimization,omitempty" xml:"CrossBorderOptimization,omitempty"`
	IpAccessRule            *string `json:"IpAccessRule,omitempty" xml:"IpAccessRule,omitempty"`
	Ipv6                    *string `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	// example:
	//
	// test.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EB635996-1FD6-5DFD-BA57-27A849599940
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*GetTransportLayerApplicationResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	RulesCount *int32 `json:"RulesCount,omitempty" xml:"RulesCount,omitempty"`
	// example:
	//
	// 123456****
	SiteId *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTransportLayerApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTransportLayerApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetTransportLayerApplicationResponseBody) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *GetTransportLayerApplicationResponseBody) GetCname() *string {
	return s.Cname
}

func (s *GetTransportLayerApplicationResponseBody) GetCrossBorderOptimization() *string {
	return s.CrossBorderOptimization
}

func (s *GetTransportLayerApplicationResponseBody) GetIpAccessRule() *string {
	return s.IpAccessRule
}

func (s *GetTransportLayerApplicationResponseBody) GetIpv6() *string {
	return s.Ipv6
}

func (s *GetTransportLayerApplicationResponseBody) GetRecordName() *string {
	return s.RecordName
}

func (s *GetTransportLayerApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTransportLayerApplicationResponseBody) GetRules() []*GetTransportLayerApplicationResponseBodyRules {
	return s.Rules
}

func (s *GetTransportLayerApplicationResponseBody) GetRulesCount() *int32 {
	return s.RulesCount
}

func (s *GetTransportLayerApplicationResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetTransportLayerApplicationResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetTransportLayerApplicationResponseBody) SetApplicationId(v int64) *GetTransportLayerApplicationResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBody) SetCname(v string) *GetTransportLayerApplicationResponseBody {
	s.Cname = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBody) SetCrossBorderOptimization(v string) *GetTransportLayerApplicationResponseBody {
	s.CrossBorderOptimization = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBody) SetIpAccessRule(v string) *GetTransportLayerApplicationResponseBody {
	s.IpAccessRule = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBody) SetIpv6(v string) *GetTransportLayerApplicationResponseBody {
	s.Ipv6 = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBody) SetRecordName(v string) *GetTransportLayerApplicationResponseBody {
	s.RecordName = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBody) SetRequestId(v string) *GetTransportLayerApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBody) SetRules(v []*GetTransportLayerApplicationResponseBodyRules) *GetTransportLayerApplicationResponseBody {
	s.Rules = v
	return s
}

func (s *GetTransportLayerApplicationResponseBody) SetRulesCount(v int32) *GetTransportLayerApplicationResponseBody {
	s.RulesCount = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBody) SetSiteId(v int64) *GetTransportLayerApplicationResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBody) SetStatus(v string) *GetTransportLayerApplicationResponseBody {
	s.Status = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTransportLayerApplicationResponseBodyRules struct {
	// example:
	//
	// off
	ClientIPPassThroughMode *string `json:"ClientIPPassThroughMode,omitempty" xml:"ClientIPPassThroughMode,omitempty"`
	Comment                 *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 80
	EdgePort *string `json:"EdgePort,omitempty" xml:"EdgePort,omitempty"`
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 1234323***
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 1.1.1.1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 80
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// example:
	//
	// domain
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s GetTransportLayerApplicationResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s GetTransportLayerApplicationResponseBodyRules) GoString() string {
	return s.String()
}

func (s *GetTransportLayerApplicationResponseBodyRules) GetClientIPPassThroughMode() *string {
	return s.ClientIPPassThroughMode
}

func (s *GetTransportLayerApplicationResponseBodyRules) GetComment() *string {
	return s.Comment
}

func (s *GetTransportLayerApplicationResponseBodyRules) GetEdgePort() *string {
	return s.EdgePort
}

func (s *GetTransportLayerApplicationResponseBodyRules) GetProtocol() *string {
	return s.Protocol
}

func (s *GetTransportLayerApplicationResponseBodyRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetTransportLayerApplicationResponseBodyRules) GetSource() *string {
	return s.Source
}

func (s *GetTransportLayerApplicationResponseBodyRules) GetSourcePort() *string {
	return s.SourcePort
}

func (s *GetTransportLayerApplicationResponseBodyRules) GetSourceType() *string {
	return s.SourceType
}

func (s *GetTransportLayerApplicationResponseBodyRules) SetClientIPPassThroughMode(v string) *GetTransportLayerApplicationResponseBodyRules {
	s.ClientIPPassThroughMode = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBodyRules) SetComment(v string) *GetTransportLayerApplicationResponseBodyRules {
	s.Comment = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBodyRules) SetEdgePort(v string) *GetTransportLayerApplicationResponseBodyRules {
	s.EdgePort = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBodyRules) SetProtocol(v string) *GetTransportLayerApplicationResponseBodyRules {
	s.Protocol = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBodyRules) SetRuleId(v int64) *GetTransportLayerApplicationResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBodyRules) SetSource(v string) *GetTransportLayerApplicationResponseBodyRules {
	s.Source = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBodyRules) SetSourcePort(v string) *GetTransportLayerApplicationResponseBodyRules {
	s.SourcePort = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBodyRules) SetSourceType(v string) *GetTransportLayerApplicationResponseBodyRules {
	s.SourceType = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
