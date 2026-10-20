// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransportLayerApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListTransportLayerApplicationsResponseBodyApplications) *ListTransportLayerApplicationsResponseBody
	GetApplications() []*ListTransportLayerApplicationsResponseBodyApplications
	SetPageNumber(v int32) *ListTransportLayerApplicationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTransportLayerApplicationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTransportLayerApplicationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransportLayerApplicationsResponseBody
	GetTotalCount() *int32
}

type ListTransportLayerApplicationsResponseBody struct {
	// The list of Layer 4 applications.
	Applications []*ListTransportLayerApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The current page number, which is the same as the PageNumber request parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of Layer 4 applications.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTransportLayerApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransportLayerApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransportLayerApplicationsResponseBody) GetApplications() []*ListTransportLayerApplicationsResponseBodyApplications {
	return s.Applications
}

func (s *ListTransportLayerApplicationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTransportLayerApplicationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTransportLayerApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransportLayerApplicationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransportLayerApplicationsResponseBody) SetApplications(v []*ListTransportLayerApplicationsResponseBodyApplications) *ListTransportLayerApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListTransportLayerApplicationsResponseBody) SetPageNumber(v int32) *ListTransportLayerApplicationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBody) SetPageSize(v int32) *ListTransportLayerApplicationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBody) SetRequestId(v string) *ListTransportLayerApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBody) SetTotalCount(v int32) *ListTransportLayerApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBody) Validate() error {
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransportLayerApplicationsResponseBodyApplications struct {
	// The ID of the Layer 4 application.
	//
	// example:
	//
	// 170997271816****
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The CNAME domain name corresponding to the Layer 4 acceleration application. This field is not empty only when the site is onboarded by using a CNAME record.
	//
	// example:
	//
	// example.com.ialicdn.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// Specifies whether to enable network access optimization for the Chinese mainland. Default value: off. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	CrossBorderOptimization *string `json:"CrossBorderOptimization,omitempty" xml:"CrossBorderOptimization,omitempty"`
	// Specifies whether to enable the IP access rule. After this feature is enabled, the IP access rules in WAF take effect for the Layer 4 application. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	IpAccessRule *string `json:"IpAccessRule,omitempty" xml:"IpAccessRule,omitempty"`
	// Specifies whether to enable IPv6. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	Ipv6 *string `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	// Specifies whether to enable connection keep-alive protection. Default value: off. Valid values:
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
	// The list of forwarding rules.
	Rules []*ListTransportLayerApplicationsResponseBodyApplicationsRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The number of forwarding rules included in the Layer 4 acceleration application.
	//
	// example:
	//
	// 1
	RulesCount *int32 `json:"RulesCount,omitempty" xml:"RulesCount,omitempty"`
	// The site ID.
	//
	// example:
	//
	// 36556540048****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Specifies whether to enable the static IP address feature. Default value: off. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	StaticIp *string `json:"StaticIp,omitempty" xml:"StaticIp,omitempty"`
	// The list of static IPv4 addresses assigned to the Layer 4 application after the static IP address feature is enabled.
	//
	// This parameter is required.
	StaticIpV4List []*ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List `json:"StaticIpV4List,omitempty" xml:"StaticIpV4List,omitempty" type:"Repeated"`
	// The status of the Layer 4 application. Valid values:
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

func (s ListTransportLayerApplicationsResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListTransportLayerApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetCname() *string {
	return s.Cname
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetCrossBorderOptimization() *string {
	return s.CrossBorderOptimization
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetIpAccessRule() *string {
	return s.IpAccessRule
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetIpv6() *string {
	return s.Ipv6
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetKeepAliveProtection() *string {
	return s.KeepAliveProtection
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetRecordName() *string {
	return s.RecordName
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetRules() []*ListTransportLayerApplicationsResponseBodyApplicationsRules {
	return s.Rules
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetRulesCount() *int32 {
	return s.RulesCount
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetStaticIp() *string {
	return s.StaticIp
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetStaticIpV4List() []*ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List {
	return s.StaticIpV4List
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) GetStatus() *string {
	return s.Status
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetApplicationId(v int64) *ListTransportLayerApplicationsResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetCname(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.Cname = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetCrossBorderOptimization(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.CrossBorderOptimization = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetIpAccessRule(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.IpAccessRule = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetIpv6(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.Ipv6 = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetKeepAliveProtection(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.KeepAliveProtection = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetRecordName(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.RecordName = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetRules(v []*ListTransportLayerApplicationsResponseBodyApplicationsRules) *ListTransportLayerApplicationsResponseBodyApplications {
	s.Rules = v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetRulesCount(v int32) *ListTransportLayerApplicationsResponseBodyApplications {
	s.RulesCount = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetSiteId(v int64) *ListTransportLayerApplicationsResponseBodyApplications {
	s.SiteId = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetStaticIp(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.StaticIp = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetStaticIpV4List(v []*ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) *ListTransportLayerApplicationsResponseBodyApplications {
	s.StaticIpV4List = v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) SetStatus(v string) *ListTransportLayerApplicationsResponseBodyApplications {
	s.Status = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplications) Validate() error {
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

type ListTransportLayerApplicationsResponseBodyApplicationsRules struct {
	// The protocol used to pass through the client IP address. Valid values:
	//
	// - **off**: Does not pass through the client IP address.
	//
	// - **PPv1**: PROXY protocol version 1, which supports passing through the client IP address over TCP.
	//
	// - **PPv2**: PROXY protocol version 2, which supports passing through the client IP address over TCP and UDP.
	//
	// - **SPP**: Simple Proxy Protocol, which supports passing through the client IP address over UDP.
	//
	// example:
	//
	// off
	ClientIPPassThroughMode *string `json:"ClientIPPassThroughMode,omitempty" xml:"ClientIPPassThroughMode,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// Test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The edge port. Valid values:
	//
	// - A single port, such as 80.
	//
	// - A port range, such as 81-85, which represents ports 81, 82, 83, 84, and 85.
	//
	// - A combination of ports and port ranges, separated by commas, such as 80,81-85,90, which represents ports 80, 81, 82, 83, 84, 85, and 90.
	//
	// example:
	//
	// 80
	EdgePort *string `json:"EdgePort,omitempty" xml:"EdgePort,omitempty"`
	// The protocol of the forwarding rule. Valid values:
	//
	// - TCP: TCP
	//
	// - UDP: UDP
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the Layer 4 acceleration rule.
	//
	// example:
	//
	// 20258028****
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The specific value of the origin server, which must match the origin server type.
	//
	// example:
	//
	// 1.1.1.1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The origin server port. Valid values:
	//
	// - A single port. If the origin server port is a single port, any valid combination of edge ports is supported.
	//
	// - A port range. The origin server port can be set to a port range only when the edge port is a port range, and the range size must be the same as that of the edge port. For example, if the edge port is 90-93, you cannot set the origin server port to 81-85 because the origin server port range includes 5 ports and the edge port range includes 4 ports, which do not match.
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
	// - **OP**: origin IPAM pool.
	//
	// - **LB**: load balancing.
	//
	// example:
	//
	// ip
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListTransportLayerApplicationsResponseBodyApplicationsRules) String() string {
	return dara.Prettify(s)
}

func (s ListTransportLayerApplicationsResponseBodyApplicationsRules) GoString() string {
	return s.String()
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetClientIPPassThroughMode() *string {
	return s.ClientIPPassThroughMode
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetComment() *string {
	return s.Comment
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetEdgePort() *string {
	return s.EdgePort
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetProtocol() *string {
	return s.Protocol
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetSource() *string {
	return s.Source
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetSourcePort() *string {
	return s.SourcePort
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) GetSourceType() *string {
	return s.SourceType
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetClientIPPassThroughMode(v string) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.ClientIPPassThroughMode = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetComment(v string) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.Comment = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetEdgePort(v string) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.EdgePort = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetProtocol(v string) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.Protocol = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetRuleId(v int64) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.RuleId = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetSource(v string) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.Source = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetSourcePort(v string) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.SourcePort = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) SetSourceType(v string) *ListTransportLayerApplicationsResponseBodyApplicationsRules {
	s.SourceType = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsRules) Validate() error {
	return dara.Validate(s)
}

type ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List struct {
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

func (s ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) String() string {
	return dara.Prettify(s)
}

func (s ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) GoString() string {
	return s.String()
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) GetAddress() *string {
	return s.Address
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) GetStatus() *string {
	return s.Status
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) SetAddress(v string) *ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List {
	s.Address = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) SetStatus(v string) *ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List {
	s.Status = &v
	return s
}

func (s *ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List) Validate() error {
	return dara.Validate(s)
}
