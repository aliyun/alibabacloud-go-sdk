// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsProductInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenewal(v bool) *DescribeDnsProductInstanceResponseBody
	GetAutoRenewal() *bool
	SetBindCount(v int64) *DescribeDnsProductInstanceResponseBody
	GetBindCount() *int64
	SetBindDomainCount(v int64) *DescribeDnsProductInstanceResponseBody
	GetBindDomainCount() *int64
	SetBindDomainUsedCount(v int64) *DescribeDnsProductInstanceResponseBody
	GetBindDomainUsedCount() *int64
	SetBindUsedCount(v int64) *DescribeDnsProductInstanceResponseBody
	GetBindUsedCount() *int64
	SetDDosDefendFlow(v int64) *DescribeDnsProductInstanceResponseBody
	GetDDosDefendFlow() *int64
	SetDDosDefendQuery(v int64) *DescribeDnsProductInstanceResponseBody
	GetDDosDefendQuery() *int64
	SetDnsSLBCount(v int64) *DescribeDnsProductInstanceResponseBody
	GetDnsSLBCount() *int64
	SetDnsSecurity(v string) *DescribeDnsProductInstanceResponseBody
	GetDnsSecurity() *string
	SetDnsServers(v *DescribeDnsProductInstanceResponseBodyDnsServers) *DescribeDnsProductInstanceResponseBody
	GetDnsServers() *DescribeDnsProductInstanceResponseBodyDnsServers
	SetDomain(v string) *DescribeDnsProductInstanceResponseBody
	GetDomain() *string
	SetDomainType(v string) *DescribeDnsProductInstanceResponseBody
	GetDomainType() *string
	SetEndTime(v string) *DescribeDnsProductInstanceResponseBody
	GetEndTime() *string
	SetEndTimestamp(v int64) *DescribeDnsProductInstanceResponseBody
	GetEndTimestamp() *int64
	SetGslb(v bool) *DescribeDnsProductInstanceResponseBody
	GetGslb() *bool
	SetISPLines(v string) *DescribeDnsProductInstanceResponseBody
	GetISPLines() *string
	SetISPRegionLines(v string) *DescribeDnsProductInstanceResponseBody
	GetISPRegionLines() *string
	SetInBlackHole(v bool) *DescribeDnsProductInstanceResponseBody
	GetInBlackHole() *bool
	SetInClean(v bool) *DescribeDnsProductInstanceResponseBody
	GetInClean() *bool
	SetInstanceId(v string) *DescribeDnsProductInstanceResponseBody
	GetInstanceId() *string
	SetMonitorFrequency(v int64) *DescribeDnsProductInstanceResponseBody
	GetMonitorFrequency() *int64
	SetMonitorNodeCount(v int64) *DescribeDnsProductInstanceResponseBody
	GetMonitorNodeCount() *int64
	SetMonitorTaskCount(v int64) *DescribeDnsProductInstanceResponseBody
	GetMonitorTaskCount() *int64
	SetOverseaDDosDefendFlow(v int64) *DescribeDnsProductInstanceResponseBody
	GetOverseaDDosDefendFlow() *int64
	SetOverseaLine(v string) *DescribeDnsProductInstanceResponseBody
	GetOverseaLine() *string
	SetPaymentType(v string) *DescribeDnsProductInstanceResponseBody
	GetPaymentType() *string
	SetRegionLines(v bool) *DescribeDnsProductInstanceResponseBody
	GetRegionLines() *bool
	SetRequestId(v string) *DescribeDnsProductInstanceResponseBody
	GetRequestId() *string
	SetSearchEngineLines(v string) *DescribeDnsProductInstanceResponseBody
	GetSearchEngineLines() *string
	SetStartTime(v string) *DescribeDnsProductInstanceResponseBody
	GetStartTime() *string
	SetStartTimestamp(v int64) *DescribeDnsProductInstanceResponseBody
	GetStartTimestamp() *int64
	SetSubDomainLevel(v int64) *DescribeDnsProductInstanceResponseBody
	GetSubDomainLevel() *int64
	SetTTLMinValue(v int64) *DescribeDnsProductInstanceResponseBody
	GetTTLMinValue() *int64
	SetURLForwardCount(v int64) *DescribeDnsProductInstanceResponseBody
	GetURLForwardCount() *int64
	SetVersionCode(v string) *DescribeDnsProductInstanceResponseBody
	GetVersionCode() *string
	SetVersionName(v string) *DescribeDnsProductInstanceResponseBody
	GetVersionName() *string
}

type DescribeDnsProductInstanceResponseBody struct {
	// Indicates whether auto-renewal is enabled for the instance.
	//
	// - true: Auto-renewal is enabled.
	//
	// - false: Auto-renewal is disabled.
	//
	// example:
	//
	// true
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The number of times the domain name can be changed for the paid DNS instance. This parameter is available for the Custom Edition.
	//
	// example:
	//
	// 3
	BindCount *int64 `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	// The number of domain names that can be attached to the paid DNS instance. This parameter is available for the Personal and Ultimate editions.
	//
	// example:
	//
	// 5
	BindDomainCount *int64 `json:"BindDomainCount,omitempty" xml:"BindDomainCount,omitempty"`
	// The number of domain names that are attached to the paid DNS instance. This parameter is available for the Personal and Ultimate editions.
	//
	// example:
	//
	// 3
	BindDomainUsedCount *int64 `json:"BindDomainUsedCount,omitempty" xml:"BindDomainUsedCount,omitempty"`
	// The number of times the domain name has been changed for the paid DNS instance. This parameter is available for the Custom Edition.
	//
	// example:
	//
	// 1
	BindUsedCount *int64 `json:"BindUsedCount,omitempty" xml:"BindUsedCount,omitempty"`
	// The DDoS protection bandwidth. Unit: Gbit/s.
	//
	// example:
	//
	// 50
	DDosDefendFlow *int64 `json:"DDosDefendFlow,omitempty" xml:"DDosDefendFlow,omitempty"`
	// The DDoS protection capacity in queries per second (QPS). The unit is 10,000 QPS. This parameter is available for the Custom Edition.
	//
	// example:
	//
	// 50
	DDosDefendQuery *int64 `json:"DDosDefendQuery,omitempty" xml:"DDosDefendQuery,omitempty"`
	// The Server Load Balancer (SLB) capacity. This is the number of IP addresses that can be configured for a domain name on a single line.
	//
	// example:
	//
	// 15
	DnsSLBCount *int64 `json:"DnsSLBCount,omitempty" xml:"DnsSLBCount,omitempty"`
	// The DNS security level. Valid values:
	//
	// - no: Not required
	//
	// - basic: Basic DNS attack protection
	//
	// - advanced: Advanced DNS attack protection
	//
	// example:
	//
	// advanced
	DnsSecurity *string                                           `json:"DnsSecurity,omitempty" xml:"DnsSecurity,omitempty"`
	DnsServers  *DescribeDnsProductInstanceResponseBodyDnsServers `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Struct"`
	// The attached domain name.
	//
	// If this parameter is empty, no domain name is attached.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The type of the instance:
	//
	// - PUBLIC: An instance for an authoritative domain name.
	//
	// - CACHE: An instance for a recursive DNS proxy.
	//
	// example:
	//
	// PUBLIC
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 2015-12-12T09:23Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the instance expires. This is a UNIX timestamp.
	//
	// example:
	//
	// 1474335170000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// Indicates whether Global Server Load Balancer (GSLB) is allowed.
	//
	// - true: Allowed
	//
	// - false: Not allowed
	//
	// example:
	//
	// true
	Gslb *bool `json:"Gslb,omitempty" xml:"Gslb,omitempty"`
	// The carrier line from which the DNS request was initiated. Valid values:
	//
	// - China Telecom
	//
	// - China Mobile
	//
	// - China Unicom
	//
	// - China Education and Research Network
	//
	// - China Broadcasting Network
	//
	// - Dr. Peng Group
	//
	// example:
	//
	// 中国电信
	ISPLines *string `json:"ISPLines,omitempty" xml:"ISPLines,omitempty"`
	// The carrier line and province from which the DNS request was initiated. Valid values:
	//
	// - China Telecom (by province)
	//
	// - China Mobile (by province)
	//
	// - China Unicom (by province)
	//
	// - China Education and Research Network (by province)
	//
	// example:
	//
	// 电信_浙江
	ISPRegionLines *string `json:"ISPRegionLines,omitempty" xml:"ISPRegionLines,omitempty"`
	// Indicates whether the domain name is in a blackhole filtering status.
	//
	// - true: The domain name is in a blackhole filtering status.
	//
	// - false: The domain name is not in a blackhole filtering status.
	//
	// example:
	//
	// false
	InBlackHole *bool `json:"InBlackHole,omitempty" xml:"InBlackHole,omitempty"`
	// Indicates whether the domain name is undergoing traffic scrubbing.
	//
	// - true: Traffic scrubbing is in progress.
	//
	// - false: Traffic scrubbing is not in progress.
	//
	// example:
	//
	// false
	InClean *bool `json:"InClean,omitempty" xml:"InClean,omitempty"`
	// The ID of the Alibaba Cloud DNS instance.
	//
	// example:
	//
	// i-8fxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The monitoring frequency. Unit: minutes.
	//
	// example:
	//
	// 50
	MonitorFrequency *int64 `json:"MonitorFrequency,omitempty" xml:"MonitorFrequency,omitempty"`
	// The number of monitoring nodes.
	//
	// example:
	//
	// 5
	MonitorNodeCount *int64 `json:"MonitorNodeCount,omitempty" xml:"MonitorNodeCount,omitempty"`
	// The number of monitoring jobs.
	//
	// example:
	//
	// 2
	MonitorTaskCount *int64 `json:"MonitorTaskCount,omitempty" xml:"MonitorTaskCount,omitempty"`
	// The DDoS protection bandwidth for regions outside China. Unit: Gbit/s.
	//
	// example:
	//
	// 1
	OverseaDDosDefendFlow *int64 `json:"OverseaDDosDefendFlow,omitempty" xml:"OverseaDDosDefendFlow,omitempty"`
	// The line for regions outside China.
	//
	// example:
	//
	// 海外大洲
	OverseaLine *string `json:"OverseaLine,omitempty" xml:"OverseaLine,omitempty"`
	// The billing method.
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// Indicates whether regional lines are used.
	//
	// - true: Regional lines are used.
	//
	// - false: Regional lines are not used.
	//
	// example:
	//
	// true
	RegionLines *bool `json:"RegionLines,omitempty" xml:"RegionLines,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-xxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The search engine line. Valid values:
	//
	// - Google
	//
	// - Baidu
	//
	// - Bing
	//
	// - Youdao
	//
	// example:
	//
	// 谷歌
	SearchEngineLines *string `json:"SearchEngineLines,omitempty" xml:"SearchEngineLines,omitempty"`
	// The time when the instance was purchased.
	//
	// example:
	//
	// 2015-12-12T09:23Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time when the instance was purchased. This is a UNIX timestamp.
	//
	// example:
	//
	// 1474335170000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The number of subdomain levels.
	//
	// example:
	//
	// 6
	SubDomainLevel *int64 `json:"SubDomainLevel,omitempty" xml:"SubDomainLevel,omitempty"`
	// The minimum Time to Live (TTL) value. Unit: seconds.
	//
	// example:
	//
	// 10
	TTLMinValue *int64 `json:"TTLMinValue,omitempty" xml:"TTLMinValue,omitempty"`
	// The number of URL forwards.
	//
	// example:
	//
	// 20
	URLForwardCount *int64 `json:"URLForwardCount,omitempty" xml:"URLForwardCount,omitempty"`
	// The code of the Alibaba Cloud DNS edition.
	//
	// example:
	//
	// version1
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	// The name of the Alibaba Cloud DNS edition.
	//
	// example:
	//
	// 企业旗舰版
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s DescribeDnsProductInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsProductInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsProductInstanceResponseBody) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *DescribeDnsProductInstanceResponseBody) GetBindCount() *int64 {
	return s.BindCount
}

func (s *DescribeDnsProductInstanceResponseBody) GetBindDomainCount() *int64 {
	return s.BindDomainCount
}

func (s *DescribeDnsProductInstanceResponseBody) GetBindDomainUsedCount() *int64 {
	return s.BindDomainUsedCount
}

func (s *DescribeDnsProductInstanceResponseBody) GetBindUsedCount() *int64 {
	return s.BindUsedCount
}

func (s *DescribeDnsProductInstanceResponseBody) GetDDosDefendFlow() *int64 {
	return s.DDosDefendFlow
}

func (s *DescribeDnsProductInstanceResponseBody) GetDDosDefendQuery() *int64 {
	return s.DDosDefendQuery
}

func (s *DescribeDnsProductInstanceResponseBody) GetDnsSLBCount() *int64 {
	return s.DnsSLBCount
}

func (s *DescribeDnsProductInstanceResponseBody) GetDnsSecurity() *string {
	return s.DnsSecurity
}

func (s *DescribeDnsProductInstanceResponseBody) GetDnsServers() *DescribeDnsProductInstanceResponseBodyDnsServers {
	return s.DnsServers
}

func (s *DescribeDnsProductInstanceResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDnsProductInstanceResponseBody) GetDomainType() *string {
	return s.DomainType
}

func (s *DescribeDnsProductInstanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDnsProductInstanceResponseBody) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeDnsProductInstanceResponseBody) GetGslb() *bool {
	return s.Gslb
}

func (s *DescribeDnsProductInstanceResponseBody) GetISPLines() *string {
	return s.ISPLines
}

func (s *DescribeDnsProductInstanceResponseBody) GetISPRegionLines() *string {
	return s.ISPRegionLines
}

func (s *DescribeDnsProductInstanceResponseBody) GetInBlackHole() *bool {
	return s.InBlackHole
}

func (s *DescribeDnsProductInstanceResponseBody) GetInClean() *bool {
	return s.InClean
}

func (s *DescribeDnsProductInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsProductInstanceResponseBody) GetMonitorFrequency() *int64 {
	return s.MonitorFrequency
}

func (s *DescribeDnsProductInstanceResponseBody) GetMonitorNodeCount() *int64 {
	return s.MonitorNodeCount
}

func (s *DescribeDnsProductInstanceResponseBody) GetMonitorTaskCount() *int64 {
	return s.MonitorTaskCount
}

func (s *DescribeDnsProductInstanceResponseBody) GetOverseaDDosDefendFlow() *int64 {
	return s.OverseaDDosDefendFlow
}

func (s *DescribeDnsProductInstanceResponseBody) GetOverseaLine() *string {
	return s.OverseaLine
}

func (s *DescribeDnsProductInstanceResponseBody) GetPaymentType() *string {
	return s.PaymentType
}

func (s *DescribeDnsProductInstanceResponseBody) GetRegionLines() *bool {
	return s.RegionLines
}

func (s *DescribeDnsProductInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsProductInstanceResponseBody) GetSearchEngineLines() *string {
	return s.SearchEngineLines
}

func (s *DescribeDnsProductInstanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDnsProductInstanceResponseBody) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeDnsProductInstanceResponseBody) GetSubDomainLevel() *int64 {
	return s.SubDomainLevel
}

func (s *DescribeDnsProductInstanceResponseBody) GetTTLMinValue() *int64 {
	return s.TTLMinValue
}

func (s *DescribeDnsProductInstanceResponseBody) GetURLForwardCount() *int64 {
	return s.URLForwardCount
}

func (s *DescribeDnsProductInstanceResponseBody) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeDnsProductInstanceResponseBody) GetVersionName() *string {
	return s.VersionName
}

func (s *DescribeDnsProductInstanceResponseBody) SetAutoRenewal(v bool) *DescribeDnsProductInstanceResponseBody {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetBindCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.BindCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetBindDomainCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.BindDomainCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetBindDomainUsedCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.BindDomainUsedCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetBindUsedCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.BindUsedCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetDDosDefendFlow(v int64) *DescribeDnsProductInstanceResponseBody {
	s.DDosDefendFlow = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetDDosDefendQuery(v int64) *DescribeDnsProductInstanceResponseBody {
	s.DDosDefendQuery = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetDnsSLBCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.DnsSLBCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetDnsSecurity(v string) *DescribeDnsProductInstanceResponseBody {
	s.DnsSecurity = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetDnsServers(v *DescribeDnsProductInstanceResponseBodyDnsServers) *DescribeDnsProductInstanceResponseBody {
	s.DnsServers = v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetDomain(v string) *DescribeDnsProductInstanceResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetDomainType(v string) *DescribeDnsProductInstanceResponseBody {
	s.DomainType = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetEndTime(v string) *DescribeDnsProductInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetEndTimestamp(v int64) *DescribeDnsProductInstanceResponseBody {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetGslb(v bool) *DescribeDnsProductInstanceResponseBody {
	s.Gslb = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetISPLines(v string) *DescribeDnsProductInstanceResponseBody {
	s.ISPLines = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetISPRegionLines(v string) *DescribeDnsProductInstanceResponseBody {
	s.ISPRegionLines = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetInBlackHole(v bool) *DescribeDnsProductInstanceResponseBody {
	s.InBlackHole = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetInClean(v bool) *DescribeDnsProductInstanceResponseBody {
	s.InClean = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetInstanceId(v string) *DescribeDnsProductInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetMonitorFrequency(v int64) *DescribeDnsProductInstanceResponseBody {
	s.MonitorFrequency = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetMonitorNodeCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.MonitorNodeCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetMonitorTaskCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.MonitorTaskCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetOverseaDDosDefendFlow(v int64) *DescribeDnsProductInstanceResponseBody {
	s.OverseaDDosDefendFlow = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetOverseaLine(v string) *DescribeDnsProductInstanceResponseBody {
	s.OverseaLine = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetPaymentType(v string) *DescribeDnsProductInstanceResponseBody {
	s.PaymentType = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetRegionLines(v bool) *DescribeDnsProductInstanceResponseBody {
	s.RegionLines = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetRequestId(v string) *DescribeDnsProductInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetSearchEngineLines(v string) *DescribeDnsProductInstanceResponseBody {
	s.SearchEngineLines = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetStartTime(v string) *DescribeDnsProductInstanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetStartTimestamp(v int64) *DescribeDnsProductInstanceResponseBody {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetSubDomainLevel(v int64) *DescribeDnsProductInstanceResponseBody {
	s.SubDomainLevel = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetTTLMinValue(v int64) *DescribeDnsProductInstanceResponseBody {
	s.TTLMinValue = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetURLForwardCount(v int64) *DescribeDnsProductInstanceResponseBody {
	s.URLForwardCount = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetVersionCode(v string) *DescribeDnsProductInstanceResponseBody {
	s.VersionCode = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) SetVersionName(v string) *DescribeDnsProductInstanceResponseBody {
	s.VersionName = &v
	return s
}

func (s *DescribeDnsProductInstanceResponseBody) Validate() error {
	if s.DnsServers != nil {
		if err := s.DnsServers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsProductInstanceResponseBodyDnsServers struct {
	DnsServer []*string `json:"DnsServer,omitempty" xml:"DnsServer,omitempty" type:"Repeated"`
}

func (s DescribeDnsProductInstanceResponseBodyDnsServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsProductInstanceResponseBodyDnsServers) GoString() string {
	return s.String()
}

func (s *DescribeDnsProductInstanceResponseBodyDnsServers) GetDnsServer() []*string {
	return s.DnsServer
}

func (s *DescribeDnsProductInstanceResponseBodyDnsServers) SetDnsServer(v []*string) *DescribeDnsProductInstanceResponseBodyDnsServers {
	s.DnsServer = v
	return s
}

func (s *DescribeDnsProductInstanceResponseBodyDnsServers) Validate() error {
	return dara.Validate(s)
}
