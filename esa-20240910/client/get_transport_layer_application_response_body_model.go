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
	SetKeepAliveProtection(v string) *GetTransportLayerApplicationResponseBody
	GetKeepAliveProtection() *string
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
	// The Layer 4 application ID.
	//
	// example:
	//
	// 17099311410****
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The CNAME domain name of the Layer 4 acceleration application. This field is non-empty only when the site is connected by using the CNAME method.
	//
	// example:
	//
	// example.com.ialicdn.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// Specifies whether to enable network access optimization for the Chinese mainland. This feature is disabled by default. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	CrossBorderOptimization *string `json:"CrossBorderOptimization,omitempty" xml:"CrossBorderOptimization,omitempty"`
	// The IP access rule switch. When enabled, IP access rules in WAF take effect for the Layer 4 application. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	IpAccessRule *string `json:"IpAccessRule,omitempty" xml:"IpAccessRule,omitempty"`
	// The IPv6 switch. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	Ipv6 *string `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	// Specifies whether to enable keep-alive protection. This feature is disabled by default. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	KeepAliveProtection *string `json:"KeepAliveProtection,omitempty" xml:"KeepAliveProtection,omitempty"`
	// The domain name of the Layer 4 application.
	//
	// example:
	//
	// test.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EB635996-1FD6-5DFD-BA57-27A849599940
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of forwarding rules.
	Rules []*GetTransportLayerApplicationResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The number of forwarding rules in the Layer 4 acceleration application.
	//
	// example:
	//
	// 1
	RulesCount *int32 `json:"RulesCount,omitempty" xml:"RulesCount,omitempty"`
	// The site ID.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Specifies whether to enable static IP. This feature is disabled by default. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// off
	StaticIp *string `json:"StaticIp,omitempty" xml:"StaticIp,omitempty"`
	// The list of static IPv4 addresses assigned to this Layer 4 application after the static IP feature is enabled.
	StaticIpV4List []*GetTransportLayerApplicationResponseBodyStaticIpV4List `json:"StaticIpV4List,omitempty" xml:"StaticIpV4List,omitempty" type:"Repeated"`
	// The Layer 4 application status. Valid values:
	//
	// - **deploying**: Being deployed. Modifications and deletions are not allowed in this state.
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

func (s *GetTransportLayerApplicationResponseBody) GetKeepAliveProtection() *string {
	return s.KeepAliveProtection
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

func (s *GetTransportLayerApplicationResponseBody) SetKeepAliveProtection(v string) *GetTransportLayerApplicationResponseBody {
	s.KeepAliveProtection = &v
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
	// The client IP pass-through protocol. Valid values:
	//
	// - **off**: disabled.
	//
	// - **PPv1**: PROXY Protocol v1, which supports client IP pass-through for TCP.
	//
	// - **PPv2**: PROXY Protocol v2, which supports client IP pass-through for TCP and UDP.
	//
	// - **SPP**: Simple Proxy Protocol, which supports client IP pass-through for UDP.
	//
	// example:
	//
	// off
	ClientIPPassThroughMode *string `json:"ClientIPPassThroughMode,omitempty" xml:"ClientIPPassThroughMode,omitempty"`
	// The comment for the rule.
	//
	// example:
	//
	// Test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The edge port. The following formats are supported:
	//
	// - A single port, such as 80.
	//
	// - A port range, such as 81-85, which represents ports 81, 82, 83, 84, and 85.
	//
	// - A combination of ports and port ranges separated by commas, such as 80,81-85,90, which represents ports 80, 81, 82, 83, 84, 85, and 90.
	//
	// example:
	//
	// 80
	EdgePort *string `json:"EdgePort,omitempty" xml:"EdgePort,omitempty"`
	// The protocol of the forwarding rule. Valid values:
	//
	// - TCP: TCP protocol.
	//
	// - UDP: UDP protocol.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 1234323***
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The specific value of the origin server, which must match the origin server type.
	//
	// example:
	//
	// 1.1.1.1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The origin server port. The following formats are supported:
	//
	// - A single port. When the origin server port is a single port, any valid edge port combination is supported.
	//
	// - A port range. The origin server port can be set to a port range only when the edge port is a port range, and the range size must be the same as the edge port range. For example, if the edge port is 90-93, you cannot set the origin server port to 81-85 because the origin server port range is 5 while the edge port range is 4, which are inconsistent.
	//
	// example:
	//
	// 80
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The origin server type. Valid values:
	//
	// - **ip**: IP address.
	//
	// - **domain**: domain name.
	//
	// - **OP**: IPAM pool.
	//
	// - **LB**: load balancing.
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
	// The IP address.
	//
	// example:
	//
	// 1.1.1.2
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The status of the IP address. Valid values:
	//
	// - healthy: Healthy.
	//
	// - unhealthy: Unhealthy.
	//
	// - unknown: The IP address is being prepared.
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
