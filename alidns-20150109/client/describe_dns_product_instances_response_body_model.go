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
	if s.DnsProducts != nil {
		if err := s.DnsProducts.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.DnsProduct != nil {
		for _, item := range s.DnsProduct {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsProductInstancesResponseBodyDnsProductsDnsProduct struct {
	AutoRenewal           *bool   `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	BindCount             *int64  `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	BindDomainCount       *int64  `json:"BindDomainCount,omitempty" xml:"BindDomainCount,omitempty"`
	BindDomainUsedCount   *int64  `json:"BindDomainUsedCount,omitempty" xml:"BindDomainUsedCount,omitempty"`
	BindUsedCount         *int64  `json:"BindUsedCount,omitempty" xml:"BindUsedCount,omitempty"`
	DDosDefendFlow        *int64  `json:"DDosDefendFlow,omitempty" xml:"DDosDefendFlow,omitempty"`
	DDosDefendQuery       *int64  `json:"DDosDefendQuery,omitempty" xml:"DDosDefendQuery,omitempty"`
	DnsSLBCount           *int64  `json:"DnsSLBCount,omitempty" xml:"DnsSLBCount,omitempty"`
	DnsSecurity           *string `json:"DnsSecurity,omitempty" xml:"DnsSecurity,omitempty"`
	Domain                *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EndTime               *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimestamp          *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	Gslb                  *bool   `json:"Gslb,omitempty" xml:"Gslb,omitempty"`
	ISPLines              *string `json:"ISPLines,omitempty" xml:"ISPLines,omitempty"`
	ISPRegionLines        *string `json:"ISPRegionLines,omitempty" xml:"ISPRegionLines,omitempty"`
	InBlackHole           *bool   `json:"InBlackHole,omitempty" xml:"InBlackHole,omitempty"`
	InClean               *bool   `json:"InClean,omitempty" xml:"InClean,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MonitorFrequency      *int64  `json:"MonitorFrequency,omitempty" xml:"MonitorFrequency,omitempty"`
	MonitorNodeCount      *int64  `json:"MonitorNodeCount,omitempty" xml:"MonitorNodeCount,omitempty"`
	MonitorTaskCount      *int64  `json:"MonitorTaskCount,omitempty" xml:"MonitorTaskCount,omitempty"`
	OverseaDDosDefendFlow *int64  `json:"OverseaDDosDefendFlow,omitempty" xml:"OverseaDDosDefendFlow,omitempty"`
	OverseaLine           *string `json:"OverseaLine,omitempty" xml:"OverseaLine,omitempty"`
	PaymentType           *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	RegionLines           *bool   `json:"RegionLines,omitempty" xml:"RegionLines,omitempty"`
	SearchEngineLines     *string `json:"SearchEngineLines,omitempty" xml:"SearchEngineLines,omitempty"`
	StartTime             *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimestamp        *int64  `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	SubDomainLevel        *int64  `json:"SubDomainLevel,omitempty" xml:"SubDomainLevel,omitempty"`
	TTLMinValue           *int64  `json:"TTLMinValue,omitempty" xml:"TTLMinValue,omitempty"`
	URLForwardCount       *int64  `json:"URLForwardCount,omitempty" xml:"URLForwardCount,omitempty"`
	VersionCode           *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VersionName           *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
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
