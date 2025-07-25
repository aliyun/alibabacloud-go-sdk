// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsProductInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDnsProducts(v *DescribeDnsProductInstancesResponseBodyDnsProducts) *DescribeDnsProductInstancesResponseBody
	GetDnsProducts() *DescribeDnsProductInstancesResponseBodyDnsProducts
	SetDomainType(v string) *DescribeDnsProductInstancesResponseBody
	GetDomainType() *string
	SetPageNumber(v int64) *DescribeDnsProductInstancesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDnsProductInstancesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDnsProductInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDnsProductInstancesResponseBody
	GetTotalCount() *int64
}

type DescribeDnsProductInstancesResponseBody struct {
	// The paid Alibaba Cloud DNS instances.
	DnsProducts *DescribeDnsProductInstancesResponseBodyDnsProducts `json:"DnsProducts,omitempty" xml:"DnsProducts,omitempty" type:"Struct"`
	// The type of the domain name. Valid values:
	//
	// 	- PUBLIC (default): hosted public domain name
	//
	// 	- CACHE: cached public domain name
	//
	// example:
	//
	// PUBLIC
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 2
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of domain names.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDnsProductInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsProductInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsProductInstancesResponseBody) GetDnsProducts() *DescribeDnsProductInstancesResponseBodyDnsProducts {
	return s.DnsProducts
}

func (s *DescribeDnsProductInstancesResponseBody) GetDomainType() *string {
	return s.DomainType
}

func (s *DescribeDnsProductInstancesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDnsProductInstancesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDnsProductInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsProductInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDnsProductInstancesResponseBody) SetDnsProducts(v *DescribeDnsProductInstancesResponseBodyDnsProducts) *DescribeDnsProductInstancesResponseBody {
	s.DnsProducts = v
	return s
}

func (s *DescribeDnsProductInstancesResponseBody) SetDomainType(v string) *DescribeDnsProductInstancesResponseBody {
	s.DomainType = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBody) SetPageNumber(v int64) *DescribeDnsProductInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBody) SetPageSize(v int64) *DescribeDnsProductInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBody) SetRequestId(v string) *DescribeDnsProductInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBody) SetTotalCount(v int64) *DescribeDnsProductInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsProductInstancesResponseBodyDnsProducts struct {
	DnsProduct []*DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct `json:"DnsProduct,omitempty" xml:"DnsProduct,omitempty" type:"Repeated"`
}

func (s DescribeDnsProductInstancesResponseBodyDnsProducts) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsProductInstancesResponseBodyDnsProducts) GoString() string {
	return s.String()
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) GetDnsProduct() []*DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	return s.DnsProduct
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) SetDnsProduct(v []*DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) *DescribeDnsProductInstancesResponseBodyDnsProducts {
	s.DnsProduct = v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProducts) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct struct {
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
	// The number of times you can change domain names that are bound to the DNS instance. It can be specified by the user who uses Alibaba Cloud DNS of the custom version.
	//
	// example:
	//
	// 3
	BindCount *int64 `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	// The number of domain names that can be bound to the DNS instance.
	//
	// example:
	//
	// 5
	BindDomainCount *int64 `json:"BindDomainCount,omitempty" xml:"BindDomainCount,omitempty"`
	// The number of domain names that have been bound to the DNS instance.
	//
	// example:
	//
	// 3
	BindDomainUsedCount *int64 `json:"BindDomainUsedCount,omitempty" xml:"BindDomainUsedCount,omitempty"`
	// The number of times you have changed domain names that are bound to the DNS instance. It can be specified by the user who uses Alibaba Cloud DNS of the custom version.
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
	// The DDoS protection frequency. Unit: 10,000 QPS.
	//
	// example:
	//
	// 50
	DDosDefendQuery *int64 `json:"DDosDefendQuery,omitempty" xml:"DDosDefendQuery,omitempty"`
	// The number of IP addresses supported by a domain name or line.
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
	// no
	DnsSecurity *string `json:"DnsSecurity,omitempty" xml:"DnsSecurity,omitempty"`
	// The bound domain name.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The time at which the instance expired.
	//
	// example:
	//
	// 2015-12-12T00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The UNIX timestamp representing the expiration time of the instance.
	//
	// example:
	//
	// 1474335170000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// Indicates whether global server load balancing (GSLB) is supported.
	//
	// 	- true: GSLB is supported.
	//
	// 	- false: GSLB is not supported.
	//
	// example:
	//
	// true
	Gslb *bool `json:"Gslb,omitempty" xml:"Gslb,omitempty"`
	// The ISP resolution lines.
	//
	// 	- China Telecom
	//
	// 	- China Mobile
	//
	// 	- China Unicom
	//
	// 	- CERNET
	//
	// 	- China Broadcasting Network (CBN)
	//
	// 	- Dr Peng Telecom & Media Group
	//
	// example:
	//
	// China Mobile,China Unicom,China Telecom,China Edu
	ISPLines *string `json:"ISPLines,omitempty" xml:"ISPLines,omitempty"`
	// The regional ISP resolution lines. Valid values:
	//
	// 	- China Telecom (province)
	//
	// 	- China Mobile (province)
	//
	// 	- China Unicom (province)
	//
	// 	- China Education and Research Network (CERNET) (province)
	//
	// example:
	//
	// Telecom (Province), Mobile (Province), Unicom (Province), Education Network (Province)
	ISPRegionLines *string `json:"ISPRegionLines,omitempty" xml:"ISPRegionLines,omitempty"`
	// Indicates whether the Domain Name System (DNS) servers stopped responding to all requests. Valid values:
	//
	// 	- true: The DNS servers stopped responding to all requests.
	//
	// 	- false: The DNS servers did not stop responding to all requests.
	//
	// example:
	//
	// false
	InBlackHole *bool `json:"InBlackHole,omitempty" xml:"InBlackHole,omitempty"`
	// Indicates whether the request for domain name resolution was being cleared.
	//
	// example:
	//
	// false
	InClean *bool `json:"InClean,omitempty" xml:"InClean,omitempty"`
	// The ID of the Alibaba Cloud DNS instance.
	//
	// example:
	//
	// i-8fj
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The monitoring frequency. Unit: minutes.
	//
	// example:
	//
	// 50
	MonitorFrequency *int64 `json:"MonitorFrequency,omitempty" xml:"MonitorFrequency,omitempty"`
	// The number of monitored nodes.
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
	// DDoS protection traffic outside China. Unit: GB.
	//
	// example:
	//
	// 1
	OverseaDDosDefendFlow *int64 `json:"OverseaDDosDefendFlow,omitempty" xml:"OverseaDDosDefendFlow,omitempty"`
	// The type of the overseas line.
	//
	// example:
	//
	// Countries
	OverseaLine *string `json:"OverseaLine,omitempty" xml:"OverseaLine,omitempty"`
	// The billing method.
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// Indicates whether the DNS request lines are regional lines.
	//
	// 	- true: The DNS request lines are regional lines.
	//
	// 	- false: The DNS request lines are not regional lines.
	//
	// example:
	//
	// true
	RegionLines *bool `json:"RegionLines,omitempty" xml:"RegionLines,omitempty"`
	// The search engine resolution lines. Valid values:
	//
	// 	- Google
	//
	// 	- Baidu
	//
	// 	- Bing
	//
	// 	- Youdao
	//
	// example:
	//
	// Search Engine Robots,Google Robots,Baidu Robots,Bing Robots
	SearchEngineLines *string `json:"SearchEngineLines,omitempty" xml:"SearchEngineLines,omitempty"`
	// The time when the DNS instance was purchased.
	//
	// example:
	//
	// 2015-11-12T09:23Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The UNIX timestamp representing when the DNS instance was purchased.
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
	// The minimum TTL. Unit: seconds.
	//
	// example:
	//
	// 10
	TTLMinValue *int64 `json:"TTLMinValue,omitempty" xml:"TTLMinValue,omitempty"`
	// The URL forwarding quantity.
	//
	// example:
	//
	// 20
	URLForwardCount *int64 `json:"URLForwardCount,omitempty" xml:"URLForwardCount,omitempty"`
	// The version code of the Alibaba Cloud DNS instance.
	//
	// example:
	//
	// version1
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	// The version name of the Alibaba Cloud DNS instance.
	//
	// example:
	//
	// Alibaba Cloud DNS
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GoString() string {
	return s.String()
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetBindCount() *int64 {
	return s.BindCount
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetBindDomainCount() *int64 {
	return s.BindDomainCount
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetBindDomainUsedCount() *int64 {
	return s.BindDomainUsedCount
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetBindUsedCount() *int64 {
	return s.BindUsedCount
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetDDosDefendFlow() *int64 {
	return s.DDosDefendFlow
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetDDosDefendQuery() *int64 {
	return s.DDosDefendQuery
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetDnsSLBCount() *int64 {
	return s.DnsSLBCount
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetDnsSecurity() *string {
	return s.DnsSecurity
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetGslb() *bool {
	return s.Gslb
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetISPLines() *string {
	return s.ISPLines
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetISPRegionLines() *string {
	return s.ISPRegionLines
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetInBlackHole() *bool {
	return s.InBlackHole
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetInClean() *bool {
	return s.InClean
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetMonitorFrequency() *int64 {
	return s.MonitorFrequency
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetMonitorNodeCount() *int64 {
	return s.MonitorNodeCount
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetMonitorTaskCount() *int64 {
	return s.MonitorTaskCount
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetOverseaDDosDefendFlow() *int64 {
	return s.OverseaDDosDefendFlow
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetOverseaLine() *string {
	return s.OverseaLine
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetPaymentType() *string {
	return s.PaymentType
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetRegionLines() *bool {
	return s.RegionLines
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetSearchEngineLines() *string {
	return s.SearchEngineLines
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetSubDomainLevel() *int64 {
	return s.SubDomainLevel
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetTTLMinValue() *int64 {
	return s.TTLMinValue
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetURLForwardCount() *int64 {
	return s.URLForwardCount
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) GetVersionName() *string {
	return s.VersionName
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetAutoRenewal(v bool) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetBindCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.BindCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetBindDomainCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.BindDomainCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetBindDomainUsedCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.BindDomainUsedCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetBindUsedCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.BindUsedCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetDDosDefendFlow(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.DDosDefendFlow = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetDDosDefendQuery(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.DDosDefendQuery = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetDnsSLBCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.DnsSLBCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetDnsSecurity(v string) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.DnsSecurity = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetDomain(v string) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.Domain = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetEndTime(v string) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.EndTime = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetEndTimestamp(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetGslb(v bool) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.Gslb = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetISPLines(v string) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.ISPLines = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetISPRegionLines(v string) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.ISPRegionLines = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetInBlackHole(v bool) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.InBlackHole = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetInClean(v bool) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.InClean = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetInstanceId(v string) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetMonitorFrequency(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.MonitorFrequency = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetMonitorNodeCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.MonitorNodeCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetMonitorTaskCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.MonitorTaskCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetOverseaDDosDefendFlow(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.OverseaDDosDefendFlow = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetOverseaLine(v string) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.OverseaLine = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetPaymentType(v string) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.PaymentType = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetRegionLines(v bool) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.RegionLines = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetSearchEngineLines(v string) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.SearchEngineLines = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetStartTime(v string) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.StartTime = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetStartTimestamp(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetSubDomainLevel(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.SubDomainLevel = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetTTLMinValue(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.TTLMinValue = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetURLForwardCount(v int64) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.URLForwardCount = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetVersionCode(v string) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.VersionCode = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) SetVersionName(v string) *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct {
	s.VersionName = &v
	return s
}

func (s *DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct) Validate() error {
	return dara.Validate(s)
}
