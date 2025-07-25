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
	// Indicates whether auto-renewal was enabled. Valid values:
	//
	// 	- true: Auto-renewal was enabled.
	//
	// 	- false: Auto-renewal was not enabled.
	//
	// example:
	//
	// true
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The number of times that you can change the domain names that are bound to the paid Alibaba Cloud DNS instance. This parameter applies to Alibaba Cloud DNS instances of the custom edition.
	//
	// example:
	//
	// 3
	BindCount *int64 `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	// The number of domain names that can be bound to the paid Alibaba Cloud DNS instance. This parameter applies to Alibaba Cloud DNS instances of Personal Edition, Enterprise Standard Edition, and Enterprise Ultimate Edition.
	//
	// example:
	//
	// 5
	BindDomainCount *int64 `json:"BindDomainCount,omitempty" xml:"BindDomainCount,omitempty"`
	// The number of domain names that are bound to the paid Alibaba Cloud DNS instance. This parameter applies to Alibaba Cloud DNS instances of Personal Edition, Enterprise Standard Edition, and Enterprise Ultimate Edition.
	//
	// example:
	//
	// 3
	BindDomainUsedCount *int64 `json:"BindDomainUsedCount,omitempty" xml:"BindDomainUsedCount,omitempty"`
	// The number of times that you have changed the domain names that are bound to the paid Alibaba Cloud DNS instance. This parameter applies to Alibaba Cloud DNS instances of the custom edition.
	//
	// example:
	//
	// 1
	BindUsedCount *int64 `json:"BindUsedCount,omitempty" xml:"BindUsedCount,omitempty"`
	// The DDoS protection traffic. Unit: GB.
	//
	// example:
	//
	// 50
	DDosDefendFlow *int64 `json:"DDosDefendFlow,omitempty" xml:"DDosDefendFlow,omitempty"`
	// The DDoS protection frequency. Unit: 10,000 QPS. This parameter applies to Alibaba Cloud DNS instances of the custom edition.
	//
	// example:
	//
	// 50
	DDosDefendQuery *int64 `json:"DDosDefendQuery,omitempty" xml:"DDosDefendQuery,omitempty"`
	// The maximum number of IP addresses that are used for load balancing in a single line of a domain name.
	//
	// example:
	//
	// 15
	DnsSLBCount *int64 `json:"DnsSLBCount,omitempty" xml:"DnsSLBCount,omitempty"`
	// The level of DNS protection. Valid values:
	//
	// 	- no: No DNS protection is provided.
	//
	// 	- basic: Basic DNS protection is provided.
	//
	// 	- advanced: Advanced DNS protection is provided.
	//
	// example:
	//
	// advanced
	DnsSecurity *string `json:"DnsSecurity,omitempty" xml:"DnsSecurity,omitempty"`
	// The DNS servers configured for the domain names.
	DnsServers *DescribeDnsProductInstanceResponseBodyDnsServers `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Struct"`
	// The domain name that is bound to the paid instance.
	//
	// If no value is returned for this parameter, no domain name is bound to the paid instance.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The type of the instance. Valid values:
	//
	// 	- PUBLIC: authoritative domain name
	//
	// 	- CACHE: cache-accelerated domain name
	//
	// example:
	//
	// PUBLIC
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The time when the instance expired. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-12T09:23Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the instance expired. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1474335170000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// Indicates whether global server load balancing (GSLB) is supported. Valid values:
	//
	// 	- true: GSLB is supported.
	//
	// 	- false: GSLB is not supported.
	//
	// example:
	//
	// true
	Gslb *bool `json:"Gslb,omitempty" xml:"Gslb,omitempty"`
	// The ISP resolution lines. Valid values:
	//
	// 	- China Telecom
	//
	// 	- China Mobile
	//
	// 	- China Unicom
	//
	// 	- China Education and Research Network (CERNET)
	//
	// 	- China Broadcasting Network (CBN)
	//
	// 	- Dr Peng Telecom & Media Group
	ISPLines *string `json:"ISPLines,omitempty" xml:"ISPLines,omitempty"`
	// The regional ISP resolution lines. Valid values:
	//
	// 	- China Telecom (province)
	//
	// 	- China Mobile (province)
	//
	// 	- China Unicom (province)
	//
	// 	- CERNET (province)
	ISPRegionLines *string `json:"ISPRegionLines,omitempty" xml:"ISPRegionLines,omitempty"`
	// Indicates whether the Domain Name System (DNS) servers stopped responding to all DNS requests. Valid values:
	//
	// 	- true: The DNS servers stopped responding to all DNS requests.
	//
	// 	- false: The DNS servers did not stop responding to all DNS requests.
	//
	// example:
	//
	// false
	InBlackHole *bool `json:"InBlackHole,omitempty" xml:"InBlackHole,omitempty"`
	// Indicates whether the DNS servers stopped responding to abnormal requests sent to the domain names.
	//
	// 	- true: The DNS servers stopped responding to abnormal requests sent to the domain names.
	//
	// 	- false: The DNS servers did not stop responding to abnormal requests sent to the domain names.
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
	// The interval at which the instance is monitored. Unit: minutes.
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
	// The number of monitoring tasks.
	//
	// example:
	//
	// 2
	MonitorTaskCount *int64 `json:"MonitorTaskCount,omitempty" xml:"MonitorTaskCount,omitempty"`
	// The DDoS protection traffic outside the Chinese mainland. Unit: GB.
	//
	// example:
	//
	// 1
	OverseaDDosDefendFlow *int64 `json:"OverseaDDosDefendFlow,omitempty" xml:"OverseaDDosDefendFlow,omitempty"`
	// The line outside the Chinese mainland.
	OverseaLine *string `json:"OverseaLine,omitempty" xml:"OverseaLine,omitempty"`
	// The billing method.
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// Indicates whether the DNS request lines are regional lines. Valid values:
	//
	// 	- true: The DNS request lines are regional lines.
	//
	// 	- false: The DNS request lines are not regional lines.
	//
	// example:
	//
	// true
	RegionLines *bool `json:"RegionLines,omitempty" xml:"RegionLines,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The search engine resolution lines. Valid values:
	//
	// 	- Google
	//
	// 	- Baidu
	//
	// 	- Bing
	//
	// 	- Youdao
	SearchEngineLines *string `json:"SearchEngineLines,omitempty" xml:"SearchEngineLines,omitempty"`
	// The time when the instance was purchased. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-12T09:23Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time when the instance was purchased. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1474335170000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The number of subdomain name levels.
	//
	// example:
	//
	// 6
	SubDomainLevel *int64 `json:"SubDomainLevel,omitempty" xml:"SubDomainLevel,omitempty"`
	// The minimum time-to-live (TTL) period. Unit: seconds.
	//
	// example:
	//
	// 10
	TTLMinValue *int64 `json:"TTLMinValue,omitempty" xml:"TTLMinValue,omitempty"`
	// The number of the forwarded URLs.
	//
	// example:
	//
	// 20
	URLForwardCount *int64 `json:"URLForwardCount,omitempty" xml:"URLForwardCount,omitempty"`
	// The version code of Alibaba Cloud DNS.
	//
	// example:
	//
	// version1
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	// The edition of Alibaba Cloud DNS.
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
	return dara.Validate(s)
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
