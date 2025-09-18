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
	// List of transport layer applications.
	Applications []*ListTransportLayerApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// Current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of transport layer applications.
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
	return dara.Validate(s)
}

type ListTransportLayerApplicationsResponseBodyApplications struct {
	// Layer 4 application ID.
	//
	// example:
	//
	// 170997271816****
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// CNAME domain name corresponding to the Layer 4 acceleration application. This field is not empty only when the site is accessed via CNAME.
	//
	// example:
	//
	// example.com.ialicdn.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// Whether to enable China mainland network access optimization. It is disabled by default. The value range is:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	CrossBorderOptimization *string `json:"CrossBorderOptimization,omitempty" xml:"CrossBorderOptimization,omitempty"`
	// IP access rule switch. When enabled, the IP access rules in WAF take effect on the Layer 4 application.
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
	// on
	Ipv6 *string `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	// Domain name of the Layer 4 application.
	//
	// example:
	//
	// test.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// List of forwarding rules.
	Rules []*ListTransportLayerApplicationsResponseBodyApplicationsRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// Number of forwarding rules contained in the Layer 4 acceleration application.
	//
	// example:
	//
	// 1
	RulesCount *int32 `json:"RulesCount,omitempty" xml:"RulesCount,omitempty"`
	// Site ID.
	//
	// example:
	//
	// 36556540048****
	SiteId   *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	StaticIp *string `json:"StaticIp,omitempty" xml:"StaticIp,omitempty"`
	// This parameter is required.
	StaticIpV4List []*ListTransportLayerApplicationsResponseBodyApplicationsStaticIpV4List `json:"StaticIpV4List,omitempty" xml:"StaticIpV4List,omitempty" type:"Repeated"`
	// Status of the Layer 4 application
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
	return dara.Validate(s)
}

type ListTransportLayerApplicationsResponseBodyApplicationsRules struct {
	// Client IP pass-through protocol, supports:
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
	// Comment information for the rule.
	//
	// example:
	//
	// 测试
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Edge port. Supports:
	//
	// - A single port, e.g., 80.
	//
	// - Port range, e.g., 81-85, representing ports 81, 82, 83, 84, 85.
	//
	// - Combination of ports and port ranges, separated by commas, e.g., 80,81-85,90, representing ports 80, 81, 82, 83, 84, 85, 90.
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
	// Layer 4 acceleration rule ID.
	//
	// example:
	//
	// 20258028****
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// Specific value of the source, which needs to match the source type.
	//
	// example:
	//
	// 1.1.1.1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Source port. Supports:
	//
	// - A single port, when the source port is a single port, any valid combination of edge ports is supported.
	//
	// - Port range, only when the edge port is a port range, the source port can be set to a port range, and the range size must be consistent with the edge port. For example, if the edge port is 90-93, the source port cannot be set to 81-85 because the source port range is 5 and the edge port range is 3, which are inconsistent.
	//
	// example:
	//
	// 80
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// Source type, supports:
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
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
