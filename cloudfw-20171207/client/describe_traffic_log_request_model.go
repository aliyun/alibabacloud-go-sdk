// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrafficLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclPreRuleId(v string) *DescribeTrafficLogRequest
	GetAclPreRuleId() *string
	SetAclPreState(v string) *DescribeTrafficLogRequest
	GetAclPreState() *string
	SetAppDpiState(v string) *DescribeTrafficLogRequest
	GetAppDpiState() *string
	SetAppId(v string) *DescribeTrafficLogRequest
	GetAppId() *string
	SetAssetRegion(v string) *DescribeTrafficLogRequest
	GetAssetRegion() *string
	SetAttackType(v string) *DescribeTrafficLogRequest
	GetAttackType() *string
	SetCurrentPage(v string) *DescribeTrafficLogRequest
	GetCurrentPage() *string
	SetDirection(v string) *DescribeTrafficLogRequest
	GetDirection() *string
	SetDomainName(v string) *DescribeTrafficLogRequest
	GetDomainName() *string
	SetDomainUrl(v string) *DescribeTrafficLogRequest
	GetDomainUrl() *string
	SetDstIP(v string) *DescribeTrafficLogRequest
	GetDstIP() *string
	SetDstPort(v string) *DescribeTrafficLogRequest
	GetDstPort() *string
	SetDstVpcId(v string) *DescribeTrafficLogRequest
	GetDstVpcId() *string
	SetDstVpcRegionNo(v string) *DescribeTrafficLogRequest
	GetDstVpcRegionNo() *string
	SetEndTime(v string) *DescribeTrafficLogRequest
	GetEndTime() *string
	SetFirewallType(v string) *DescribeTrafficLogRequest
	GetFirewallType() *string
	SetFlowType(v string) *DescribeTrafficLogRequest
	GetFlowType() *string
	SetIpProtocol(v string) *DescribeTrafficLogRequest
	GetIpProtocol() *string
	SetIpVersion(v string) *DescribeTrafficLogRequest
	GetIpVersion() *string
	SetIsp(v string) *DescribeTrafficLogRequest
	GetIsp() *string
	SetLang(v string) *DescribeTrafficLogRequest
	GetLang() *string
	SetLocation(v string) *DescribeTrafficLogRequest
	GetLocation() *string
	SetMemberUid(v int64) *DescribeTrafficLogRequest
	GetMemberUid() *int64
	SetNatFirewallId(v string) *DescribeTrafficLogRequest
	GetNatFirewallId() *string
	SetNatGatewayId(v string) *DescribeTrafficLogRequest
	GetNatGatewayId() *string
	SetPageSize(v string) *DescribeTrafficLogRequest
	GetPageSize() *string
	SetRuleId(v string) *DescribeTrafficLogRequest
	GetRuleId() *string
	SetRuleResult(v string) *DescribeTrafficLogRequest
	GetRuleResult() *string
	SetRuleSource(v string) *DescribeTrafficLogRequest
	GetRuleSource() *string
	SetRuleSourceFinal(v string) *DescribeTrafficLogRequest
	GetRuleSourceFinal() *string
	SetSourceCode(v string) *DescribeTrafficLogRequest
	GetSourceCode() *string
	SetSourceIp(v string) *DescribeTrafficLogRequest
	GetSourceIp() *string
	SetSrcIP(v string) *DescribeTrafficLogRequest
	GetSrcIP() *string
	SetSrcPort(v string) *DescribeTrafficLogRequest
	GetSrcPort() *string
	SetSrcPrivateIP(v string) *DescribeTrafficLogRequest
	GetSrcPrivateIP() *string
	SetSrcVpcId(v string) *DescribeTrafficLogRequest
	GetSrcVpcId() *string
	SetSrcVpcRegionNo(v string) *DescribeTrafficLogRequest
	GetSrcVpcRegionNo() *string
	SetStartTime(v string) *DescribeTrafficLogRequest
	GetStartTime() *string
	SetTlsScopeId(v string) *DescribeTrafficLogRequest
	GetTlsScopeId() *string
	SetVpcFirewallId(v string) *DescribeTrafficLogRequest
	GetVpcFirewallId() *string
	SetVulLevel(v string) *DescribeTrafficLogRequest
	GetVulLevel() *string
}

type DescribeTrafficLogRequest struct {
	// The ID of the pre-matched ACL rule.
	//
	// example:
	//
	// 00000000-0000-0000-0000-000000000000
	AclPreRuleId *string `json:"AclPreRuleId,omitempty" xml:"AclPreRuleId,omitempty"`
	// The pre-matching status of the ACL.
	//
	// example:
	//
	// normal
	AclPreState *string `json:"AclPreState,omitempty" xml:"AclPreState,omitempty"`
	// The status of deep packet inspection.
	//
	// example:
	//
	// success
	AppDpiState *string `json:"AppDpiState,omitempty" xml:"AppDpiState,omitempty"`
	// The application ID.
	//
	// example:
	//
	// 7
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	AssetRegion *string `json:"AssetRegion,omitempty" xml:"AssetRegion,omitempty"`
	// The attack type.
	//
	// example:
	//
	// 1
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The page number of the query.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The direction.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The URL in the flow log.
	//
	// example:
	//
	// example.com
	DomainUrl *string `json:"DomainUrl,omitempty" xml:"DomainUrl,omitempty"`
	// The destination IP address.
	//
	// example:
	//
	// 182.92.206.XXX
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 9876
	DstPort *string `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// The destination VPC ID.
	//
	// example:
	//
	// vpc-wz95m1aq9b0h****vk1yb
	DstVpcId *string `json:"DstVpcId,omitempty" xml:"DstVpcId,omitempty"`
	// The region of the destination VPC asset.
	//
	// example:
	//
	// cn-shenzhen
	DstVpcRegionNo *string `json:"DstVpcRegionNo,omitempty" xml:"DstVpcRegionNo,omitempty"`
	// The end time. This value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1742926322
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The firewall type.
	//
	// example:
	//
	// VpcFirewall
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// The flow log type.
	//
	// example:
	//
	// All
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// icmp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The IP version.
	//
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The Internet Service Provider (ISP).
	//
	// example:
	//
	// telecom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region of the source or destination IP address.
	//
	// example:
	//
	// Hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The UID of the member account.
	//
	// example:
	//
	// 128599825273****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the NAT firewall.
	//
	// example:
	//
	// vfw-tr-7a9c8901ed394****
	NatFirewallId *string `json:"NatFirewallId,omitempty" xml:"NatFirewallId,omitempty"`
	// The ID of the NAT Gateway.
	//
	// example:
	//
	// ngw-2zew6yn017hhzbm****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The page number to query. The maximum value is 20.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 8b115ae3-da64-4b80-81c1-1cd2dd42****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The action of the rule.
	//
	// example:
	//
	// 0
	RuleResult *string `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	// The source of the rule.
	//
	// example:
	//
	// 1
	RuleSource *string `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	// The module that takes effect in the end.
	//
	// example:
	//
	// 1
	RuleSourceFinal *string `json:"RuleSourceFinal,omitempty" xml:"RuleSourceFinal,omitempty"`
	// The traceability code.
	//
	// This parameter is required.
	//
	// example:
	//
	// yundun
	SourceCode *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 139.217.234.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 10.68.60.XXX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The source port.
	//
	// example:
	//
	// 8082
	SrcPort *string `json:"SrcPort,omitempty" xml:"SrcPort,omitempty"`
	// The private source IP address.
	//
	// example:
	//
	// 10.100.134.XX
	SrcPrivateIP *string `json:"SrcPrivateIP,omitempty" xml:"SrcPrivateIP,omitempty"`
	// The source VPC ID.
	//
	// example:
	//
	// vpc-wz9309pkwe06lv****tk4
	SrcVpcId *string `json:"SrcVpcId,omitempty" xml:"SrcVpcId,omitempty"`
	// The region of the source VPC asset.
	//
	// example:
	//
	// cn-beijing
	SrcVpcRegionNo *string `json:"SrcVpcRegionNo,omitempty" xml:"SrcVpcRegionNo,omitempty"`
	// The start time. This value is a UNIX timestamp. Unit: seconds. You can query data within the last 7 days. The time range for a single query cannot exceed one day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1730946241
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the TLS inspection scope.
	//
	// example:
	//
	// tis-98fd64c5****
	TlsScopeId *string `json:"TlsScopeId,omitempty" xml:"TlsScopeId,omitempty"`
	// The instance ID of the VPC border firewall.
	//
	// example:
	//
	// vfw-a42bbb7b887148c9****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The vulnerability level.
	//
	// example:
	//
	// 1
	VulLevel *string `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
}

func (s DescribeTrafficLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrafficLogRequest) GetAclPreRuleId() *string {
	return s.AclPreRuleId
}

func (s *DescribeTrafficLogRequest) GetAclPreState() *string {
	return s.AclPreState
}

func (s *DescribeTrafficLogRequest) GetAppDpiState() *string {
	return s.AppDpiState
}

func (s *DescribeTrafficLogRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeTrafficLogRequest) GetAssetRegion() *string {
	return s.AssetRegion
}

func (s *DescribeTrafficLogRequest) GetAttackType() *string {
	return s.AttackType
}

func (s *DescribeTrafficLogRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeTrafficLogRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeTrafficLogRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeTrafficLogRequest) GetDomainUrl() *string {
	return s.DomainUrl
}

func (s *DescribeTrafficLogRequest) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeTrafficLogRequest) GetDstPort() *string {
	return s.DstPort
}

func (s *DescribeTrafficLogRequest) GetDstVpcId() *string {
	return s.DstVpcId
}

func (s *DescribeTrafficLogRequest) GetDstVpcRegionNo() *string {
	return s.DstVpcRegionNo
}

func (s *DescribeTrafficLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTrafficLogRequest) GetFirewallType() *string {
	return s.FirewallType
}

func (s *DescribeTrafficLogRequest) GetFlowType() *string {
	return s.FlowType
}

func (s *DescribeTrafficLogRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeTrafficLogRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeTrafficLogRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeTrafficLogRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTrafficLogRequest) GetLocation() *string {
	return s.Location
}

func (s *DescribeTrafficLogRequest) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribeTrafficLogRequest) GetNatFirewallId() *string {
	return s.NatFirewallId
}

func (s *DescribeTrafficLogRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeTrafficLogRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeTrafficLogRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeTrafficLogRequest) GetRuleResult() *string {
	return s.RuleResult
}

func (s *DescribeTrafficLogRequest) GetRuleSource() *string {
	return s.RuleSource
}

func (s *DescribeTrafficLogRequest) GetRuleSourceFinal() *string {
	return s.RuleSourceFinal
}

func (s *DescribeTrafficLogRequest) GetSourceCode() *string {
	return s.SourceCode
}

func (s *DescribeTrafficLogRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeTrafficLogRequest) GetSrcIP() *string {
	return s.SrcIP
}

func (s *DescribeTrafficLogRequest) GetSrcPort() *string {
	return s.SrcPort
}

func (s *DescribeTrafficLogRequest) GetSrcPrivateIP() *string {
	return s.SrcPrivateIP
}

func (s *DescribeTrafficLogRequest) GetSrcVpcId() *string {
	return s.SrcVpcId
}

func (s *DescribeTrafficLogRequest) GetSrcVpcRegionNo() *string {
	return s.SrcVpcRegionNo
}

func (s *DescribeTrafficLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTrafficLogRequest) GetTlsScopeId() *string {
	return s.TlsScopeId
}

func (s *DescribeTrafficLogRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeTrafficLogRequest) GetVulLevel() *string {
	return s.VulLevel
}

func (s *DescribeTrafficLogRequest) SetAclPreRuleId(v string) *DescribeTrafficLogRequest {
	s.AclPreRuleId = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetAclPreState(v string) *DescribeTrafficLogRequest {
	s.AclPreState = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetAppDpiState(v string) *DescribeTrafficLogRequest {
	s.AppDpiState = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetAppId(v string) *DescribeTrafficLogRequest {
	s.AppId = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetAssetRegion(v string) *DescribeTrafficLogRequest {
	s.AssetRegion = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetAttackType(v string) *DescribeTrafficLogRequest {
	s.AttackType = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetCurrentPage(v string) *DescribeTrafficLogRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetDirection(v string) *DescribeTrafficLogRequest {
	s.Direction = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetDomainName(v string) *DescribeTrafficLogRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetDomainUrl(v string) *DescribeTrafficLogRequest {
	s.DomainUrl = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetDstIP(v string) *DescribeTrafficLogRequest {
	s.DstIP = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetDstPort(v string) *DescribeTrafficLogRequest {
	s.DstPort = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetDstVpcId(v string) *DescribeTrafficLogRequest {
	s.DstVpcId = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetDstVpcRegionNo(v string) *DescribeTrafficLogRequest {
	s.DstVpcRegionNo = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetEndTime(v string) *DescribeTrafficLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetFirewallType(v string) *DescribeTrafficLogRequest {
	s.FirewallType = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetFlowType(v string) *DescribeTrafficLogRequest {
	s.FlowType = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetIpProtocol(v string) *DescribeTrafficLogRequest {
	s.IpProtocol = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetIpVersion(v string) *DescribeTrafficLogRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetIsp(v string) *DescribeTrafficLogRequest {
	s.Isp = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetLang(v string) *DescribeTrafficLogRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetLocation(v string) *DescribeTrafficLogRequest {
	s.Location = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetMemberUid(v int64) *DescribeTrafficLogRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetNatFirewallId(v string) *DescribeTrafficLogRequest {
	s.NatFirewallId = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetNatGatewayId(v string) *DescribeTrafficLogRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetPageSize(v string) *DescribeTrafficLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetRuleId(v string) *DescribeTrafficLogRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetRuleResult(v string) *DescribeTrafficLogRequest {
	s.RuleResult = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetRuleSource(v string) *DescribeTrafficLogRequest {
	s.RuleSource = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetRuleSourceFinal(v string) *DescribeTrafficLogRequest {
	s.RuleSourceFinal = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetSourceCode(v string) *DescribeTrafficLogRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetSourceIp(v string) *DescribeTrafficLogRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetSrcIP(v string) *DescribeTrafficLogRequest {
	s.SrcIP = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetSrcPort(v string) *DescribeTrafficLogRequest {
	s.SrcPort = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetSrcPrivateIP(v string) *DescribeTrafficLogRequest {
	s.SrcPrivateIP = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetSrcVpcId(v string) *DescribeTrafficLogRequest {
	s.SrcVpcId = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetSrcVpcRegionNo(v string) *DescribeTrafficLogRequest {
	s.SrcVpcRegionNo = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetStartTime(v string) *DescribeTrafficLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetTlsScopeId(v string) *DescribeTrafficLogRequest {
	s.TlsScopeId = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetVpcFirewallId(v string) *DescribeTrafficLogRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeTrafficLogRequest) SetVulLevel(v string) *DescribeTrafficLogRequest {
	s.VulLevel = &v
	return s
}

func (s *DescribeTrafficLogRequest) Validate() error {
	return dara.Validate(s)
}
