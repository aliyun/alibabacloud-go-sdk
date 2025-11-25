// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrafficLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeTrafficLogResponseBodyDataList) *DescribeTrafficLogResponseBody
	GetDataList() []*DescribeTrafficLogResponseBodyDataList
	SetPageInfo(v *DescribeTrafficLogResponseBodyPageInfo) *DescribeTrafficLogResponseBody
	GetPageInfo() *DescribeTrafficLogResponseBodyPageInfo
	SetRequestId(v string) *DescribeTrafficLogResponseBody
	GetRequestId() *string
}

type DescribeTrafficLogResponseBody struct {
	DataList []*DescribeTrafficLogResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	PageInfo *DescribeTrafficLogResponseBodyPageInfo   `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// example:
	//
	// 633D92D1-768A-547F-8ADC-2870CF0A99F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTrafficLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrafficLogResponseBody) GetDataList() []*DescribeTrafficLogResponseBodyDataList {
	return s.DataList
}

func (s *DescribeTrafficLogResponseBody) GetPageInfo() *DescribeTrafficLogResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeTrafficLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTrafficLogResponseBody) SetDataList(v []*DescribeTrafficLogResponseBodyDataList) *DescribeTrafficLogResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeTrafficLogResponseBody) SetPageInfo(v *DescribeTrafficLogResponseBodyPageInfo) *DescribeTrafficLogResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeTrafficLogResponseBody) SetRequestId(v string) *DescribeTrafficLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrafficLogResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTrafficLogResponseBodyDataList struct {
	// example:
	//
	// 2
	AclPreRuleId *string `json:"AclPreRuleId,omitempty" xml:"AclPreRuleId,omitempty"`
	// example:
	//
	// test
	AclPreRuleName *string `json:"AclPreRuleName,omitempty" xml:"AclPreRuleName,omitempty"`
	// example:
	//
	// normal
	AclPreState *string `json:"AclPreState,omitempty" xml:"AclPreState,omitempty"`
	// example:
	//
	// success
	AppDpiState *string `json:"AppDpiState,omitempty" xml:"AppDpiState,omitempty"`
	// example:
	//
	// 6
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// HTTP
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// WebLogic
	AttackApp *string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty"`
	// example:
	//
	// 0
	AttackType *int32 `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// example:
	//
	// FI
	CityId *string `json:"CityId,omitempty" xml:"CityId,omitempty"`
	// example:
	//
	// tcp_fin
	CloseReason *string `json:"CloseReason,omitempty" xml:"CloseReason,omitempty"`
	// example:
	//
	// ngw-*
	CloudInstanceId *string `json:"CloudInstanceId,omitempty" xml:"CloudInstanceId,omitempty"`
	// example:
	//
	// US
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// aliyun.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// xxx.com
	DomainUrl *string `json:"DomainUrl,omitempty" xml:"DomainUrl,omitempty"`
	// example:
	//
	// 2.2.2.2
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// example:
	//
	// 80
	DstPort *int32                                        `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	DstVpc  *DescribeTrafficLogResponseBodyDataListDstVpc `json:"DstVpc,omitempty" xml:"DstVpc,omitempty" type:"Struct"`
	// example:
	//
	// 1751423363
	EndTime *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Ext     *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// 125
	InBytes *string `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// example:
	//
	// 12
	InPackets *string `json:"InPackets,omitempty" xml:"InPackets,omitempty"`
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// example:
	//
	// FOP Dmytro Nedilskyi
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// example:
	//
	// 50075069
	IspId    *string `json:"IspId,omitempty" xml:"IspId,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 14151892****7022
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// 230
	OutBytes *string `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// example:
	//
	// 11
	OutPackets *string `json:"OutPackets,omitempty" xml:"OutPackets,omitempty"`
	// example:
	//
	// 355
	PacketBytes *int64 `json:"PacketBytes,omitempty" xml:"PacketBytes,omitempty"`
	// example:
	//
	// 23
	PacketCount *int64 `json:"PacketCount,omitempty" xml:"PacketCount,omitempty"`
	// example:
	//
	// 172.21.234.XXX
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// example:
	//
	// 80
	PrivatePort *int32 `json:"PrivatePort,omitempty" xml:"PrivatePort,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-000000000000
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// pass
	RuleResult *int32 `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	// example:
	//
	// 0
	RuleSource *string                                        `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	Rules      []*DescribeTrafficLogResponseBodyDataListRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 1.1.1.1
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// example:
	//
	// 20206
	SrcPort *int32 `json:"SrcPort,omitempty" xml:"SrcPort,omitempty"`
	// example:
	//
	// 172.16.101.7
	SrcPrivateIP *string                                       `json:"SrcPrivateIP,omitempty" xml:"SrcPrivateIP,omitempty"`
	SrcVpc       *DescribeTrafficLogResponseBodyDataListSrcVpc `json:"SrcVpc,omitempty" xml:"SrcVpc,omitempty" type:"Struct"`
	// example:
	//
	// 1751423362
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// tir-xxx
	TlsRuleId *string `json:"TlsRuleId,omitempty" xml:"TlsRuleId,omitempty"`
	// example:
	//
	// test
	TlsRuleName *string `json:"TlsRuleName,omitempty" xml:"TlsRuleName,omitempty"`
	// example:
	//
	// tls-xxx
	TlsScopeId *string `json:"TlsScopeId,omitempty" xml:"TlsScopeId,omitempty"`
	// example:
	//
	// vfw-4045ca7***
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// example:
	//
	// 0
	VulLevel *int32 `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
}

func (s DescribeTrafficLogResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficLogResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeTrafficLogResponseBodyDataList) GetAclPreRuleId() *string {
	return s.AclPreRuleId
}

func (s *DescribeTrafficLogResponseBodyDataList) GetAclPreRuleName() *string {
	return s.AclPreRuleName
}

func (s *DescribeTrafficLogResponseBodyDataList) GetAclPreState() *string {
	return s.AclPreState
}

func (s *DescribeTrafficLogResponseBodyDataList) GetAppDpiState() *string {
	return s.AppDpiState
}

func (s *DescribeTrafficLogResponseBodyDataList) GetAppId() *int32 {
	return s.AppId
}

func (s *DescribeTrafficLogResponseBodyDataList) GetAppName() *string {
	return s.AppName
}

func (s *DescribeTrafficLogResponseBodyDataList) GetAttackApp() *string {
	return s.AttackApp
}

func (s *DescribeTrafficLogResponseBodyDataList) GetAttackType() *int32 {
	return s.AttackType
}

func (s *DescribeTrafficLogResponseBodyDataList) GetCityId() *string {
	return s.CityId
}

func (s *DescribeTrafficLogResponseBodyDataList) GetCloseReason() *string {
	return s.CloseReason
}

func (s *DescribeTrafficLogResponseBodyDataList) GetCloudInstanceId() *string {
	return s.CloudInstanceId
}

func (s *DescribeTrafficLogResponseBodyDataList) GetCountryId() *string {
	return s.CountryId
}

func (s *DescribeTrafficLogResponseBodyDataList) GetDirection() *string {
	return s.Direction
}

func (s *DescribeTrafficLogResponseBodyDataList) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeTrafficLogResponseBodyDataList) GetDomainUrl() *string {
	return s.DomainUrl
}

func (s *DescribeTrafficLogResponseBodyDataList) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeTrafficLogResponseBodyDataList) GetDstPort() *int32 {
	return s.DstPort
}

func (s *DescribeTrafficLogResponseBodyDataList) GetDstVpc() *DescribeTrafficLogResponseBodyDataListDstVpc {
	return s.DstVpc
}

func (s *DescribeTrafficLogResponseBodyDataList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeTrafficLogResponseBodyDataList) GetExt() *string {
	return s.Ext
}

func (s *DescribeTrafficLogResponseBodyDataList) GetInBytes() *string {
	return s.InBytes
}

func (s *DescribeTrafficLogResponseBodyDataList) GetInPackets() *string {
	return s.InPackets
}

func (s *DescribeTrafficLogResponseBodyDataList) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeTrafficLogResponseBodyDataList) GetIsp() *string {
	return s.Isp
}

func (s *DescribeTrafficLogResponseBodyDataList) GetIspId() *string {
	return s.IspId
}

func (s *DescribeTrafficLogResponseBodyDataList) GetLocation() *string {
	return s.Location
}

func (s *DescribeTrafficLogResponseBodyDataList) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeTrafficLogResponseBodyDataList) GetOutBytes() *string {
	return s.OutBytes
}

func (s *DescribeTrafficLogResponseBodyDataList) GetOutPackets() *string {
	return s.OutPackets
}

func (s *DescribeTrafficLogResponseBodyDataList) GetPacketBytes() *int64 {
	return s.PacketBytes
}

func (s *DescribeTrafficLogResponseBodyDataList) GetPacketCount() *int64 {
	return s.PacketCount
}

func (s *DescribeTrafficLogResponseBodyDataList) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *DescribeTrafficLogResponseBodyDataList) GetPrivatePort() *int32 {
	return s.PrivatePort
}

func (s *DescribeTrafficLogResponseBodyDataList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTrafficLogResponseBodyDataList) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeTrafficLogResponseBodyDataList) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeTrafficLogResponseBodyDataList) GetRuleResult() *int32 {
	return s.RuleResult
}

func (s *DescribeTrafficLogResponseBodyDataList) GetRuleSource() *string {
	return s.RuleSource
}

func (s *DescribeTrafficLogResponseBodyDataList) GetRules() []*DescribeTrafficLogResponseBodyDataListRules {
	return s.Rules
}

func (s *DescribeTrafficLogResponseBodyDataList) GetSrcIP() *string {
	return s.SrcIP
}

func (s *DescribeTrafficLogResponseBodyDataList) GetSrcPort() *int32 {
	return s.SrcPort
}

func (s *DescribeTrafficLogResponseBodyDataList) GetSrcPrivateIP() *string {
	return s.SrcPrivateIP
}

func (s *DescribeTrafficLogResponseBodyDataList) GetSrcVpc() *DescribeTrafficLogResponseBodyDataListSrcVpc {
	return s.SrcVpc
}

func (s *DescribeTrafficLogResponseBodyDataList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeTrafficLogResponseBodyDataList) GetTlsRuleId() *string {
	return s.TlsRuleId
}

func (s *DescribeTrafficLogResponseBodyDataList) GetTlsRuleName() *string {
	return s.TlsRuleName
}

func (s *DescribeTrafficLogResponseBodyDataList) GetTlsScopeId() *string {
	return s.TlsScopeId
}

func (s *DescribeTrafficLogResponseBodyDataList) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeTrafficLogResponseBodyDataList) GetVulLevel() *int32 {
	return s.VulLevel
}

func (s *DescribeTrafficLogResponseBodyDataList) SetAclPreRuleId(v string) *DescribeTrafficLogResponseBodyDataList {
	s.AclPreRuleId = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetAclPreRuleName(v string) *DescribeTrafficLogResponseBodyDataList {
	s.AclPreRuleName = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetAclPreState(v string) *DescribeTrafficLogResponseBodyDataList {
	s.AclPreState = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetAppDpiState(v string) *DescribeTrafficLogResponseBodyDataList {
	s.AppDpiState = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetAppId(v int32) *DescribeTrafficLogResponseBodyDataList {
	s.AppId = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetAppName(v string) *DescribeTrafficLogResponseBodyDataList {
	s.AppName = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetAttackApp(v string) *DescribeTrafficLogResponseBodyDataList {
	s.AttackApp = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetAttackType(v int32) *DescribeTrafficLogResponseBodyDataList {
	s.AttackType = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetCityId(v string) *DescribeTrafficLogResponseBodyDataList {
	s.CityId = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetCloseReason(v string) *DescribeTrafficLogResponseBodyDataList {
	s.CloseReason = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetCloudInstanceId(v string) *DescribeTrafficLogResponseBodyDataList {
	s.CloudInstanceId = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetCountryId(v string) *DescribeTrafficLogResponseBodyDataList {
	s.CountryId = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetDirection(v string) *DescribeTrafficLogResponseBodyDataList {
	s.Direction = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetDomainName(v string) *DescribeTrafficLogResponseBodyDataList {
	s.DomainName = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetDomainUrl(v string) *DescribeTrafficLogResponseBodyDataList {
	s.DomainUrl = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetDstIP(v string) *DescribeTrafficLogResponseBodyDataList {
	s.DstIP = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetDstPort(v int32) *DescribeTrafficLogResponseBodyDataList {
	s.DstPort = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetDstVpc(v *DescribeTrafficLogResponseBodyDataListDstVpc) *DescribeTrafficLogResponseBodyDataList {
	s.DstVpc = v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetEndTime(v int64) *DescribeTrafficLogResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetExt(v string) *DescribeTrafficLogResponseBodyDataList {
	s.Ext = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetInBytes(v string) *DescribeTrafficLogResponseBodyDataList {
	s.InBytes = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetInPackets(v string) *DescribeTrafficLogResponseBodyDataList {
	s.InPackets = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetIpProtocol(v string) *DescribeTrafficLogResponseBodyDataList {
	s.IpProtocol = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetIsp(v string) *DescribeTrafficLogResponseBodyDataList {
	s.Isp = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetIspId(v string) *DescribeTrafficLogResponseBodyDataList {
	s.IspId = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetLocation(v string) *DescribeTrafficLogResponseBodyDataList {
	s.Location = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetMemberUid(v string) *DescribeTrafficLogResponseBodyDataList {
	s.MemberUid = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetOutBytes(v string) *DescribeTrafficLogResponseBodyDataList {
	s.OutBytes = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetOutPackets(v string) *DescribeTrafficLogResponseBodyDataList {
	s.OutPackets = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetPacketBytes(v int64) *DescribeTrafficLogResponseBodyDataList {
	s.PacketBytes = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetPacketCount(v int64) *DescribeTrafficLogResponseBodyDataList {
	s.PacketCount = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetPrivateIp(v string) *DescribeTrafficLogResponseBodyDataList {
	s.PrivateIp = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetPrivatePort(v int32) *DescribeTrafficLogResponseBodyDataList {
	s.PrivatePort = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetRegionId(v string) *DescribeTrafficLogResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetRuleId(v string) *DescribeTrafficLogResponseBodyDataList {
	s.RuleId = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetRuleName(v string) *DescribeTrafficLogResponseBodyDataList {
	s.RuleName = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetRuleResult(v int32) *DescribeTrafficLogResponseBodyDataList {
	s.RuleResult = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetRuleSource(v string) *DescribeTrafficLogResponseBodyDataList {
	s.RuleSource = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetRules(v []*DescribeTrafficLogResponseBodyDataListRules) *DescribeTrafficLogResponseBodyDataList {
	s.Rules = v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetSrcIP(v string) *DescribeTrafficLogResponseBodyDataList {
	s.SrcIP = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetSrcPort(v int32) *DescribeTrafficLogResponseBodyDataList {
	s.SrcPort = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetSrcPrivateIP(v string) *DescribeTrafficLogResponseBodyDataList {
	s.SrcPrivateIP = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetSrcVpc(v *DescribeTrafficLogResponseBodyDataListSrcVpc) *DescribeTrafficLogResponseBodyDataList {
	s.SrcVpc = v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetStartTime(v int64) *DescribeTrafficLogResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetTlsRuleId(v string) *DescribeTrafficLogResponseBodyDataList {
	s.TlsRuleId = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetTlsRuleName(v string) *DescribeTrafficLogResponseBodyDataList {
	s.TlsRuleName = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetTlsScopeId(v string) *DescribeTrafficLogResponseBodyDataList {
	s.TlsScopeId = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetVpcFirewallId(v string) *DescribeTrafficLogResponseBodyDataList {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) SetVulLevel(v int32) *DescribeTrafficLogResponseBodyDataList {
	s.VulLevel = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataList) Validate() error {
	if s.DstVpc != nil {
		if err := s.DstVpc.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SrcVpc != nil {
		if err := s.SrcVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTrafficLogResponseBodyDataListDstVpc struct {
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// vpc-8vba1c1em97h0ji71b****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// yi-vpc
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeTrafficLogResponseBodyDataListDstVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficLogResponseBodyDataListDstVpc) GoString() string {
	return s.String()
}

func (s *DescribeTrafficLogResponseBodyDataListDstVpc) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeTrafficLogResponseBodyDataListDstVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeTrafficLogResponseBodyDataListDstVpc) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeTrafficLogResponseBodyDataListDstVpc) SetRegionNo(v string) *DescribeTrafficLogResponseBodyDataListDstVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataListDstVpc) SetVpcId(v string) *DescribeTrafficLogResponseBodyDataListDstVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataListDstVpc) SetVpcName(v string) *DescribeTrafficLogResponseBodyDataListDstVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataListDstVpc) Validate() error {
	return dara.Validate(s)
}

type DescribeTrafficLogResponseBodyDataListRules struct {
	// example:
	//
	// 17
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// sharepoint
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeTrafficLogResponseBodyDataListRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficLogResponseBodyDataListRules) GoString() string {
	return s.String()
}

func (s *DescribeTrafficLogResponseBodyDataListRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeTrafficLogResponseBodyDataListRules) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeTrafficLogResponseBodyDataListRules) SetRuleId(v string) *DescribeTrafficLogResponseBodyDataListRules {
	s.RuleId = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataListRules) SetRuleName(v string) *DescribeTrafficLogResponseBodyDataListRules {
	s.RuleName = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataListRules) Validate() error {
	return dara.Validate(s)
}

type DescribeTrafficLogResponseBodyDataListSrcVpc struct {
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// vpc-8vba1c1em97h0ji71****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// yi-vpc
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeTrafficLogResponseBodyDataListSrcVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficLogResponseBodyDataListSrcVpc) GoString() string {
	return s.String()
}

func (s *DescribeTrafficLogResponseBodyDataListSrcVpc) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeTrafficLogResponseBodyDataListSrcVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeTrafficLogResponseBodyDataListSrcVpc) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeTrafficLogResponseBodyDataListSrcVpc) SetRegionNo(v string) *DescribeTrafficLogResponseBodyDataListSrcVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataListSrcVpc) SetVpcId(v string) *DescribeTrafficLogResponseBodyDataListSrcVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataListSrcVpc) SetVpcName(v string) *DescribeTrafficLogResponseBodyDataListSrcVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyDataListSrcVpc) Validate() error {
	return dara.Validate(s)
}

type DescribeTrafficLogResponseBodyPageInfo struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTrafficLogResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficLogResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeTrafficLogResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeTrafficLogResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTrafficLogResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTrafficLogResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeTrafficLogResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyPageInfo) SetPageSize(v int32) *DescribeTrafficLogResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyPageInfo) SetTotalCount(v int32) *DescribeTrafficLogResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeTrafficLogResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
