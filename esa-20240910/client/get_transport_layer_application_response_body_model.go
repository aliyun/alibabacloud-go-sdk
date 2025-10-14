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
	SetStaticIp(v string) *GetTransportLayerApplicationResponseBody
	GetStaticIp() *string
	SetStaticIpV4List(v []*GetTransportLayerApplicationResponseBodyStaticIpV4List) *GetTransportLayerApplicationResponseBody
	GetStaticIpV4List() []*GetTransportLayerApplicationResponseBodyStaticIpV4List
	SetStatus(v string) *GetTransportLayerApplicationResponseBody
	GetStatus() *string
}

type GetTransportLayerApplicationResponseBody struct {
	// Transport layer application ID.
	//
	// example:
	//
	// 17099311410****
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The CNAME domain corresponding to the transport layer acceleration application. This field is not empty only when the site is accessed via CNAME.
	//
	// example:
	//
	// example.com.ialicdn.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// Whether to enable China mainland network access optimization, default is off. Value range:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	CrossBorderOptimization *string `json:"CrossBorderOptimization,omitempty" xml:"CrossBorderOptimization,omitempty"`
	// Switch for IP access rules. When turned on, the IP access rules in WAF take effect on the transport layer application.
	//
	// - on: Turned on.
	//
	// - off: Turned off.
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
	// The domain name of the transport layer application.
	//
	// example:
	//
	// test.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EB635996-1FD6-5DFD-BA57-27A849599940
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of forwarding rules.
	Rules []*GetTransportLayerApplicationResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// Number of forwarding rules contained in the transport layer acceleration application.
	//
	// example:
	//
	// 1
	RulesCount *int32 `json:"RulesCount,omitempty" xml:"RulesCount,omitempty"`
	// Site ID.
	//
	// example:
	//
	// 123456****
	SiteId         *int64                                                    `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	StaticIp       *string                                                   `json:"StaticIp,omitempty" xml:"StaticIp,omitempty"`
	StaticIpV4List []*GetTransportLayerApplicationResponseBodyStaticIpV4List `json:"StaticIpV4List,omitempty" xml:"StaticIpV4List,omitempty" type:"Repeated"`
	// Status of the transport layer application
	//
	// - **deploying**: Deploying. In this state, modification and deletion are not allowed.
	//
	// - **active**: Active.
	//
	// example:
	//
	// active
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

func (s *GetTransportLayerApplicationResponseBody) GetStaticIp() *string {
	return s.StaticIp
}

func (s *GetTransportLayerApplicationResponseBody) GetStaticIpV4List() []*GetTransportLayerApplicationResponseBodyStaticIpV4List {
	return s.StaticIpV4List
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

func (s *GetTransportLayerApplicationResponseBody) SetStaticIp(v string) *GetTransportLayerApplicationResponseBody {
	s.StaticIp = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBody) SetStaticIpV4List(v []*GetTransportLayerApplicationResponseBodyStaticIpV4List) *GetTransportLayerApplicationResponseBody {
	s.StaticIpV4List = v
	return s
}

func (s *GetTransportLayerApplicationResponseBody) SetStatus(v string) *GetTransportLayerApplicationResponseBody {
	s.Status = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBody) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.StaticIpV4List != nil {
		for _, item := range s.StaticIpV4List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTransportLayerApplicationResponseBodyRules struct {
	// Client IP pass-through protocol, supporting:
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
	// off
	ClientIPPassThroughMode *string `json:"ClientIPPassThroughMode,omitempty" xml:"ClientIPPassThroughMode,omitempty"`
	// Comment information of the rule.
	//
	// example:
	//
	// 测试
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Edge port. Supports:
	//
	// - A single port, such as 80.
	//
	// - Port range, such as 81-85, representing ports 81, 82, 83, 84, 85.
	//
	// - Combination of ports and port ranges, separated by commas, for example 80,81-85,90, representing ports 80, 81, 82, 83, 84, 85, 90.
	//
	// example:
	//
	// 80
	EdgePort *string `json:"EdgePort,omitempty" xml:"EdgePort,omitempty"`
	// Forwarding rule protocol, with values:
	//
	// - TCP: TCP protocol.
	//
	// - UDP: UDP protocol.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Rule ID.
	//
	// example:
	//
	// 1234323***
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// Specific value of the origin, which needs to match the type of the origin.
	//
	// example:
	//
	// 1.1.1.1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Origin port. Supports:
	//
	// - A single port, when the origin port is a single port, any valid edge port combination is supported.
	//
	// - Port range, only when the edge port is a port range, the origin port can be set as a port range and the size of the range must be consistent with the edge port. For example, if the edge port is 90-93, the origin port cannot be set to 81-85 because the origin port range is 5 and the edge port range is 3, which are inconsistent.
	//
	// example:
	//
	// 80
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// Origin type, supporting:
	//
	// - **ip**: IP.
	//
	// - **domain**: Domain name.
	//
	// - **OP**: Origin pool.
	//
	// - **LB**: Load balancer.
	//
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

type GetTransportLayerApplicationResponseBodyStaticIpV4List struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTransportLayerApplicationResponseBodyStaticIpV4List) String() string {
	return dara.Prettify(s)
}

func (s GetTransportLayerApplicationResponseBodyStaticIpV4List) GoString() string {
	return s.String()
}

func (s *GetTransportLayerApplicationResponseBodyStaticIpV4List) GetAddress() *string {
	return s.Address
}

func (s *GetTransportLayerApplicationResponseBodyStaticIpV4List) GetStatus() *string {
	return s.Status
}

func (s *GetTransportLayerApplicationResponseBodyStaticIpV4List) SetAddress(v string) *GetTransportLayerApplicationResponseBodyStaticIpV4List {
	s.Address = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBodyStaticIpV4List) SetStatus(v string) *GetTransportLayerApplicationResponseBodyStaticIpV4List {
	s.Status = &v
	return s
}

func (s *GetTransportLayerApplicationResponseBodyStaticIpV4List) Validate() error {
	return dara.Validate(s)
}
